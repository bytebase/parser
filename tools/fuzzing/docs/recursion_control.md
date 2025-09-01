# Recursion Control in Grammar-Aware Fuzzing

## Overview

This document describes our dependency graph-based approach to handle recursion in ANTLR 4 grammars for the fuzzing system. The strategy ensures valid output generation while preventing infinite loops and stack overflows.

## Our Strategy: Dependency Graph with Terminal Reachability

### Core Approach

1. **Build dependency graph** during grammar parsing
2. **Analyze terminal reachability** for each rule
3. **Force terminal alternatives** when hitting recursion/depth limits

### Key Principles

- **Rule = Graph Node**: Each grammar rule becomes a node
- **Reference = Graph Edge**: `a -> b` when rule `a` references rule `b`
- **Terminal Reachability**: Every rule must have at least one path to terminal nodes
- **Alternative Classification**: Mark which alternatives can terminate without recursion

## Graph Structure

### Node Definition

```go
type GraphNode struct {
    RuleName                 string        // Rule name (e.g., "selectStmt", "expr")
    HasTerminalAlternatives  bool          // Can reach terminal without recursion
    Alternatives            []Alternative  // All alternatives for this rule
    TerminalAlternativeIndex []int         // Indices of alternatives that terminate
}

type DependencyGraph struct {
    Nodes map[string]*GraphNode
}
```

### Edge Types

- **Self-Reference**: `expr -> expr` (direct recursion)
- **Cross-Reference**: `selectStmt -> whereClause` (potential indirect recursion)
- **Terminal Reference**: `expr -> NUMBER` (terminates)

## Implementation Algorithm

### Step 1: Build Graph During Parsing

```go
func BuildDependencyGraph(grammar *ParsedGrammar) *DependencyGraph {
    graph := &DependencyGraph{Nodes: make(map[string]*GraphNode)}
    
    // Create nodes for all rules
    for ruleName, rule := range grammar.GetAllRules() {
        node := &GraphNode{
            RuleName:     ruleName,
            Alternatives: rule.Alternatives,
        }
        graph.Nodes[ruleName] = node
    }
    
    // Analyze each rule for terminal reachability
    analyzeTerminalReachability(graph)
    
    return graph
}
```

### Step 2: Terminal Reachability Analysis

```go
func analyzeTerminalReachability(graph *DependencyGraph) {
    // Phase 1: Mark lexer rules as terminal
    for _, node := range graph.Nodes {
        if isLexerRule(node.RuleName) {
            node.HasTerminalAlternatives = true
            // All lexer alternatives are terminal
            for i := range node.Alternatives {
                node.TerminalAlternativeIndex = append(node.TerminalAlternativeIndex, i)
            }
        }
    }
    
    // Phase 2: Propagate terminal reachability
    changed := true
    for changed {
        changed = false
        for _, node := range graph.Nodes {
            if node.HasTerminalAlternatives {
                continue
            }
            
            // Check each alternative
            for altIndex, alt := range node.Alternatives {
                if canAlternativeTerminate(alt, graph) {
                    if !node.HasTerminalAlternatives {
                        node.HasTerminalAlternatives = true
                        changed = true
                    }
                    node.TerminalAlternativeIndex = append(node.TerminalAlternativeIndex, altIndex)
                }
            }
        }
    }
}

func canAlternativeTerminate(alt Alternative, graph *DependencyGraph) bool {
    for _, element := range alt.Elements {
        if element.IsRule() {
            referencedNode := graph.Nodes[element.RuleName]
            if referencedNode == nil || !referencedNode.HasTerminalAlternatives {
                return false
            }
        }
        // Literals and lexer rules are always terminal
    }
    return true
}
```

### Step 3: Generation with Terminal Forcing

```go
func (g *Generator) generateFromRule(ruleName string, activeRules map[string]bool, depth int) string {
    node := g.dependencyGraph.Nodes[ruleName]
    
    // Grammar validation: ensure rule can terminate
    if !node.HasTerminalAlternatives {
        return "", fmt.Errorf("unsupported grammar: rule '%s' has no terminal alternatives", ruleName)
    }
    
    // Force terminal alternatives when hitting limits
    if activeRules[ruleName] || depth >= g.config.MaxDepth {
        return g.forceTerminalGeneration(node)
    }
    
    // Normal generation
    activeRules[ruleName] = true
    defer delete(activeRules, ruleName)
    
    altIndex := g.random.Intn(len(node.Alternatives))
    return g.generateFromAlternative(node.Alternatives[altIndex], activeRules, depth+1)
}

func (g *Generator) forceTerminalGeneration(node *GraphNode) string {
    // Choose randomly from terminal alternatives only
    terminalIndex := g.random.Intn(len(node.TerminalAlternativeIndex))
    altIndex := node.TerminalAlternativeIndex[terminalIndex]
    
    // Generate with fresh context to avoid recursion
    return g.generateFromAlternative(node.Alternatives[altIndex], make(map[string]bool), 0)
}
```

## Special Cases

### Empty Alternatives (ε-transitions)

```antlr
optionalClause: whereClause | /* empty */ ;
```

**Handling**: Create implicit ε-node for empty alternatives:
```go
// Empty alternatives are always terminal
if len(alt.Elements) == 0 {
    node.TerminalAlternativeIndex = append(node.TerminalAlternativeIndex, altIndex)
}
```

### Quantified Elements

```antlr
stmt: 'BEGIN' stmt* 'END';  // stmt* can be 0 occurrences
```

**Handling**: Quantifiers `*` and `?` create implicit terminal paths:
```go
func canElementTerminate(element Element, graph *DependencyGraph) bool {
    if element.Quantifier == ZERO_MORE || element.Quantifier == OPTIONAL_Q {
        return true  // Can generate 0 occurrences
    }
    // Check if referenced rule can terminate
    return graph.Nodes[element.RuleName].HasTerminalAlternatives
}
```

### Grammar Validation

**Unsupported Grammars**: Rules with no terminal alternatives:
```antlr
// This will cause validation error
expr: '(' expr ')';  // No base case!
```

**Error Handling**:
```go
func ValidateGrammar(graph *DependencyGraph) error {
    for ruleName, node := range graph.Nodes {
        if !node.HasTerminalAlternatives {
            return fmt.Errorf("grammar error: rule '%s' has no terminal alternatives", ruleName)
        }
    }
    return nil
}
```

## Example: PostgreSQL Expression Rule

```antlr
a_expr: a_expr '+' a_expr    // Alternative 0: NON-TERMINAL (recursive)
      | a_expr '*' a_expr    // Alternative 1: NON-TERMINAL (recursive)
      | '(' a_expr ')'       // Alternative 2: NON-TERMINAL (depends on a_expr)
      | c_expr               // Alternative 3: TERMINAL (if c_expr terminates)
      ;

c_expr: columnref            // Alternative 0: TERMINAL (lexer rule)
      | '(' a_expr ')'       // Alternative 1: NON-TERMINAL (recursive)
      ;

columnref: IDENTIFIER;       // TERMINAL (lexer rule)
```

**Analysis Result**:
```go
a_expr.HasTerminalAlternatives = true
a_expr.TerminalAlternativeIndex = [3]  // Only c_expr alternative

c_expr.HasTerminalAlternatives = true
c_expr.TerminalAlternativeIndex = [0]  // Only columnref alternative
```

**Generation Behavior**:
- **Normal case**: Choose any alternative randomly
- **Recursion/MaxDepth**: Force choice from `TerminalAlternativeIndex` only
- **Result**: Always generates valid expressions without stack overflow

## Benefits

1. **No Stack Overflow**: Guaranteed termination via terminal forcing
2. **Valid Output**: No placeholders, always generates parseable content
3. **Grammar Coverage**: Supports all ANTLR 4 constructs including quantifiers
4. **Early Validation**: Detects unsupported grammars during initialization
5. **Efficient**: O(1) lookup for terminal alternatives during generation