# Grammar-Aware Fuzzing Tool Design

## Overview

A simple fuzzing tool that generates SQL inputs from ANTLR grammar rules to test parser performance on specific constructs.

## Core Problems & Solutions

### 1. Target Specific Rules
**Problem**: Performance issues often occur in specific rules (e.g., `createProcedureStatement`)
**Solution**: Allow users to specify starting rule chains

```bash
./fuzzer --grammar postgresql --start-rule createProcedureStatement --count 100
./fuzzer --grammar cql --start-rule selectStatement.whereClause --count 50
```

### 2. Recursion Control  
**Problem**: Grammar rules can be recursive, causing infinite loops during generation
**Solution**: Limit recursion depth per rule (proven to handle all ANTLR recursion types)

#### ANTLR 4 Recursion Types

**Direct Left Recursion:**
```antlr
expr: expr '+' expr | INT    // expr directly refers to itself on left
```

**Direct Right Recursion:**
```antlr
expr: INT '+' expr | INT     // expr directly refers to itself on right  
```

**Indirect Recursion (Non-Left):**
```antlr
selectStmt: SELECT columns fromClause whereClause?
whereClause: WHERE expr
expr: '(' selectStmt ')' | INT   // Indirect: expr -> selectStmt -> whereClause -> expr
```
*Note: ANTLR 4 does NOT support mutually left recursive grammars. This example is valid because the recursion is not left-recursive (selectStmt doesn't start with selectStmt).*

**Self-Recursion with Alternatives:**
```antlr
stmt: ifStmt | whileStmt | blockStmt
blockStmt: '{' stmt* '}'         // blockStmt contains multiple stmt references
```

#### Why Depth Control Works

**Theorem**: Any grammar rule expansion terminates in finite steps with depth limiting.

**Proof by Contradiction:**
1. Assume infinite expansion despite depth limit `D`
2. Each recursive call increases depth: `depth(rule_n) = depth(rule_{n-1}) + 1`
3. When `depth ≥ D`, generator forces terminal selection
4. Therefore, maximum expansion depth is bounded by `D`
5. Since each rule has finite alternatives and finite elements, total expansion is finite ∎

#### Depth Control Implementation

```go
func (g *Generator) GenerateFromRule(ruleName string, currentDepth int) string {
    // Base case: exceed depth limit -> force terminal
    if currentDepth >= g.maxDepth {
        return g.forceTerminal(ruleName)
    }
    
    rule := g.grammar.GetRule(ruleName)
    
    // Prefer non-recursive alternatives as depth increases
    alternative := g.selectAlternativeWithDepthBias(rule, currentDepth)
    
    result := ""
    for _, element := range alternative {
        if element.IsRule() {
            // Recursive call with incremented depth
            result += g.GenerateFromRule(element.Name, currentDepth+1)
        } else {
            result += element.Literal
        }
    }
    return result
}

func (g *Generator) forceTerminal(ruleName string) string {
    rule := g.grammar.GetRule(ruleName)
    
    // Find non-recursive alternatives (containing only terminals)
    for _, alt := range rule.Alternatives {
        if !alt.ContainsRecursion() {
            return g.expandAlternative(alt, g.maxDepth)
        }
    }
    
    // Fallback: use default terminal for this rule type
    return g.getDefaultTerminal(ruleName)
}
```

#### Examples with Depth Control

```bash
./fuzzer --start-rule expr --max-depth 3 --count 5
```

**Generated sequences:**
- Depth 0: `INT` (terminal)
- Depth 1: `INT + INT` 
- Depth 2: `(INT + INT) + INT`
- Depth 3: `((INT + INT) + INT) + INT` (max depth reached)

**Complex mutual recursion:**
```bash  
./fuzzer --start-rule selectStmt --max-depth 4 --count 3
```

**Expansion trace:**
```
selectStmt (depth=0)
├── SELECT columns FROM table whereClause (depth=0)
    └── whereClause (depth=1)  
        └── WHERE expr (depth=1)
            └── '(' selectStmt ')' (depth=2)
                └── selectStmt (depth=2)
                    └── SELECT columns FROM table (depth=2, no whereClause to avoid depth=4)
```

#### Depth Strategy Options

**Conservative (Early Termination):**
- Lower max depth (3-5)
- Bias toward terminals as depth increases
- Prevents deep nesting, faster generation

**Aggressive (Deep Testing):**  
- Higher max depth (10-15)
- Equal probability until max depth
- Tests parser limits, slower generation

```bash
# Conservative - quick, shallow testing
./fuzzer --start-rule expr --max-depth 3 --depth-strategy conservative

# Aggressive - deep parser stress testing  
./fuzzer --start-rule createProcedureStmt --max-depth 12 --depth-strategy aggressive
```

### 3. Optional Rule Probability
**Problem**: Optional rules (`selectStmt: SELECT columns FROM table whereClause?`) need probability control
**Solution**: Configure probability for optional elements (standard in grammar-based fuzzing)

### 4. Quantified Rule Generation
**Problem**: Quantified rules (`stmt*`, `expr+`, `column{1,5}`) need count control
**Solution**: Configure generation counts for quantified elements

#### ANTLR 4 Quantifier Types

**Zero or More (`rule*`):**
```antlr
blockStmt: '{' stmt* '}'        // Generate 0 to N statements
selectList: column (',' column)*  // Generate 1 to N columns
```

**One or More (`rule+`):**  
```antlr
identifier: LETTER (LETTER | DIGIT)+  // Generate 1 to N characters
```

**Note**: ANTLR v4 does not support `{n}` or `{n,m}` quantifier syntax. These are regex-style quantifiers not supported in ANTLR grammar files.

#### Quantifier Control Strategy

**Count Distribution Options:**
- **Uniform**: Equal probability for each count in range
- **Exponential**: Higher probability for lower counts (realistic)  
- **Fixed**: Always generate specific count

```bash
# Basic usage - user specifies max count
./fuzzer --start-rule blockStmt --max-quantifier 10 --count 100

# User controls both min and max for quantifiers  
./fuzzer --start-rule selectList --min-quantifier 1 --max-quantifier 5 --count 50

# Fixed count for performance testing
./fuzzer --start-rule selectStmt --quantifier-count 100 --count 10
```

#### Implementation Logic

```go
type QuantifierConfig struct {
    Strategy   string // "uniform", "exponential", "fixed"
    MinRepeat  int    // Minimum repetitions (overrides grammar min)
    MaxRepeat  int    // Maximum repetitions (overrides grammar max)  
    FixedCount int    // Fixed count for "fixed" strategy
}

func (g *Generator) generateQuantified(element *GrammarElement, config QuantifierConfig) string {
    var count int
    
    switch element.Quantifier {
    case "*": // Zero or more
        min := max(0, config.MinRepeat)
        max := min(config.MaxRepeat, 50) // Reasonable default limit
        count = g.selectCount(min, max, config.Strategy)
        
    case "+": // One or more  
        min := max(1, config.MinRepeat)
        max := min(config.MaxRepeat, 50)
        count = g.selectCount(min, max, config.Strategy)
        
    // Note: ANTLR v4 does not support {n} or {min,max} syntax
    }
    
    result := ""
    for i := 0; i < count; i++ {
        if element.IsRule() {
            result += g.GenerateFromRule(element.RuleName, g.currentDepth+1)
        } else {
            result += element.Literal
        }
        
        // Add separators for lists (e.g., comma-separated)
        if i < count-1 && element.HasSeparator() {
            result += element.Separator
        }
    }
    return result
}

func (g *Generator) selectCount(min, max int, strategy string) int {
    if min > max {
        return min
    }
    
    switch strategy {
    case "fixed":
        return min // Use minimum as fixed value
        
    case "uniform":
        return min + g.random.Intn(max-min+1)
        
    case "exponential":
        // Exponential decay: higher probability for lower counts
        range_size := max - min + 1
        // Generate exponentially distributed number, then map to range
        lambda := 2.0 / float64(range_size)
        exp_val := g.random.ExpFloat64() / lambda
        count := min + int(exp_val)
        if count > max {
            count = max
        }
        return count
        
    default:
        return min + g.random.Intn(max-min+1)
    }
}
```

#### Examples with Quantifier Control

**Block statement with multiple statements:**
```bash
./fuzzer --start-rule blockStmt --quantifier-strategy exponential --max-repeat 8
```
**Generated:**
- 70% chance: `{ stmt; }` (1 statement)
- 20% chance: `{ stmt; stmt; }` (2 statements)  
- 7% chance: `{ stmt; stmt; stmt; }` (3 statements)
- 3% chance: 4+ statements

**Column list generation:**
```bash  
./fuzzer --start-rule selectList --quantifier-strategy uniform --min-repeat 3 --max-repeat 7
```
**Generated:**
- Equal probability: `col1, col2, col3` to `col1, col2, col3, col4, col5, col6, col7`

**Performance testing with large lists:**
```bash
./fuzzer --start-rule selectStmt --quantifier-count 100 --count 5
```
**Generated:**
- Always generates exactly 100 columns to test parser performance on large SELECT lists

**Simple user control:**
```bash
./fuzzer --start-rule blockStmt --max-quantifier 3 --count 10
```
**Generated:**
- `stmt*` generates 0-3 statements
- `expr+` generates 1-3 expressions  
- User controls maximum without complex strategy options

```bash
./fuzzer --start-rule selectStmt --optional-prob 0.7 --count 100
# 70% chance to include optional whereClause
```

## Simple Architecture

```
tools/fuzzing/
├── main.go              # CLI entry point
├── generator.go         # Core generation logic
└── grammar_parser.go    # Reuse tools/grammar/ 
```

## Core Logic

```go
type Generator struct {
    grammar     *ParsedGrammar
    maxDepth    int
    optionalProb float64
    random      *rand.Rand
}

func (g *Generator) GenerateFromRule(ruleName string, currentDepth int) string {
    if currentDepth > g.maxDepth {
        return g.generateTerminal() // Stop recursion
    }
    
    rule := g.grammar.GetRule(ruleName)
    alternative := g.selectAlternative(rule)
    
    result := ""
    for _, element := range alternative {
        if element.IsOptional() && g.random.Float64() > g.optionalProb {
            continue // Skip optional element
        }
        if element.IsRule() {
            result += g.GenerateFromRule(element.Name, currentDepth+1)
        } else {
            result += element.Literal
        }
    }
    return result
}
```

## CLI Interface

```bash
# Basic usage - generate from specific rule
./fuzzer --grammar postgresql --start-rule selectStmt --count 10

# Control recursion depth  
./fuzzer --grammar cql --start-rule expr --max-depth 3 --count 5

# Control optional probability
./fuzzer --grammar postgresql --start-rule createStmt --optional-prob 0.8 --count 10

# Control quantifier max count (for rule*, rule+)
./fuzzer --grammar postgresql --start-rule blockStmt --max-quantifier 8 --count 20

# Control all parameters together
./fuzzer --grammar cql --start-rule selectStmt \
  --max-depth 5 \
  --optional-prob 0.7 \
  --max-quantifier 10 \
  --count 50

# Output to file
./fuzzer --grammar postgresql --start-rule selectStmt --count 100 --output queries.sql
```

## Implementation Steps

### Step 1: Basic Generator
- Parse grammar using existing `tools/grammar/`
- Simple rule expansion with depth limit
- CLI with `--start-rule`, `--max-depth`, `--count`

### Step 2: Optional Control  
- Add `--optional-prob` flag
- Detect optional elements in grammar rules
- Apply probability during generation

### Step 3: Integration
- Test generated queries against parsers
- Add basic performance timing
- CI integration for regression testing

## Common Fuzzing Techniques Used

1. **Grammar-based generation** - Generate from formal grammar rules
2. **Depth limiting** - Prevent infinite recursion in recursive grammars  
3. **Probability-based selection** - Control optional rule inclusion
4. **Targeted fuzzing** - Focus on specific rule paths instead of full grammar

This approach is much simpler but addresses your specific needs for testing parser performance on particular constructs.