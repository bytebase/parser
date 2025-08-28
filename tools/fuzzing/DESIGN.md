# Grammar-Aware Fuzzing Tool Design

## Overview

A fuzzing tool that generates valid SQL inputs by analyzing ANTLR v4 grammar files, ensuring comprehensive parser testing with syntactically correct queries that can stress-test parsing performance and correctness.

## Goals

- **Valid Input Generation**: Generate syntactically correct SQL queries based on grammar rules
- **Performance Testing**: Create complex queries to test parser performance limits  
- **Coverage Maximization**: Exercise all grammar rules and edge cases
- **Automated Testing**: Integrate with CI for continuous parser validation

## Architecture

```
tools/fuzzing/
├── generator/           # Core generation logic
│   ├── grammar_analyzer.go    # Parse ANTLR grammar files
│   ├── rule_expander.go       # Expand grammar rules to concrete syntax
│   └── query_builder.go       # Build SQL queries from rule expansions
├── strategies/          # Different generation strategies
│   ├── depth_first.go         # Generate deeply nested structures
│   ├── breadth_first.go       # Generate wide, complex queries
│   └── weighted.go            # Probability-based rule selection
├── corpus/              # Generated test cases and seeds
│   ├── seeds/                 # Hand-crafted seed inputs
│   └── generated/             # Auto-generated test cases
└── cmd/                 # CLI tools
    └── fuzzer/               # Main fuzzer executable
```

## Core Components

### 1. Grammar Analyzer

Leverages the existing `tools/grammar/` ANTLR v4 parser to:
- Parse target grammar files (e.g., `postgresql.g4`, `cql.g4`) 
- Extract production rules and their alternatives
- Build dependency graph between rules
- Identify terminal vs non-terminal symbols

```go
type GrammarAnalyzer struct {
    parser *grammar.ANTLRv4Parser
    rules  map[string]*Rule
}

type Rule struct {
    Name         string
    Alternatives []Alternative
    Type         RuleType // LEXER, PARSER, FRAGMENT
}
```

### 2. Rule Expander

Recursively expands grammar rules into concrete syntax trees:
- Handles rule recursion with configurable depth limits
- Supports probability-weighted alternative selection  
- Manages lexer rules and literal generation
- Tracks generation context for smart decisions

```go
type RuleExpander struct {
    grammar    *ParsedGrammar
    maxDepth   int
    weights    map[string]float64
    random     *rand.Rand
}
```

### 3. Query Builder

Converts syntax trees to executable SQL strings:
- Handles whitespace and formatting
- Manages identifier generation (table names, columns)
- Ensures semantic consistency where possible
- Outputs parseable query strings

## Generation Strategies

### Depth-First Strategy
- Generates deeply nested subqueries, expressions
- Tests parser stack limits and recursion handling
- Focuses on structural complexity

### Breadth-First Strategy  
- Creates wide queries with many clauses, joins, columns
- Tests parser memory usage and performance
- Focuses on query size and breadth

### Weighted Strategy
- Uses probability weights for rule selection
- Biases toward commonly used constructs
- Configurable via weight files per dialect

## Integration Points

### With Existing Grammar Parser
```go
// Reuse tools/grammar/ for parsing target grammars
analyzer := NewGrammarAnalyzer()
targetGrammar, err := analyzer.ParseGrammarFile("postgresql/PostgreSQLLexer.g4")
```

### With Parser Testing
```go
// Generate test cases for specific parser
fuzzer := NewFuzzer(postgresqlGrammar)
queries := fuzzer.GenerateQueries(1000)

for _, query := range queries {
    // Test against postgresql parser
    result := postgresqlParser.Parse(query)
    // Collect metrics, detect crashes
}
```

## Configuration

### Fuzzer Config
```yaml
target_grammar: "postgresql"
strategies:
  - name: "depth_first"
    weight: 0.3
    max_depth: 15
  - name: "breadth_first" 
    weight: 0.4
    max_width: 50
  - name: "weighted"
    weight: 0.3
    weights_file: "postgresql_weights.yaml"

generation:
  count: 10000
  max_query_length: 100000
  seed: 42

output:
  format: "sql"
  directory: "corpus/generated"
```

### Grammar Weights
```yaml
# postgresql_weights.yaml
rules:
  selectStmt: 0.4
  insertStmt: 0.2  
  updateStmt: 0.2
  deleteStmt: 0.1
  createStmt: 0.1
  
  # Bias toward complex expressions
  expr:
    binaryOp: 0.4
    functionCall: 0.3
    subquery: 0.2
    literal: 0.1
```

## CLI Interface

```bash
# Generate queries for PostgreSQL
./fuzzer generate --grammar postgresql --count 1000 --strategy weighted

# Run continuous fuzzing with performance metrics
./fuzzer fuzz --grammar cql --duration 1h --metrics

# Validate existing corpus against parser
./fuzzer validate --grammar postgresql --corpus corpus/postgresql/
```

## Performance Metrics

### Generation Metrics
- Queries generated per second
- Grammar rule coverage percentage
- Distribution of query complexity (depth, width)

### Parser Testing Metrics  
- Parse success rate
- Average parse time per query
- Memory usage during parsing
- Parser crash/error detection

## Implementation Phases

### Phase 1: Foundation (Week 1-2)
- Basic grammar analyzer using existing ANTLR parser
- Simple rule expander with depth-first strategy
- Command-line interface for manual testing

### Phase 2: Core Features (Week 3-4)
- Multiple generation strategies
- Configuration system
- Basic corpus management
- Integration with existing parser tests

### Phase 3: Advanced Features (Week 5-6)
- Weighted generation with probability tuning
- Performance metrics collection
- CI integration for continuous fuzzing
- Corpus minimization and deduplication

### Phase 4: Optimization (Week 7-8)
- Generation performance optimization
- Advanced semantic awareness
- Custom mutation strategies
- Comprehensive documentation

## Future Enhancements

- **Semantic Awareness**: Generate queries with valid schema references
- **Mutation-Based Fuzzing**: Mutate existing queries to explore edge cases
- **Differential Testing**: Compare parser outputs across database dialects
- **Performance Regression Detection**: Track parser performance over time
- **Grammar Evolution**: Adapt fuzzing as grammars evolve

## Dependencies

- Existing `tools/grammar/` ANTLR v4 parser
- Go standard library (`rand`, `fmt`, `strings`)
- YAML configuration parsing
- CLI framework (e.g., `cobra`)

This design provides a solid foundation for grammar-aware fuzzing while leveraging our existing ANTLR infrastructure.