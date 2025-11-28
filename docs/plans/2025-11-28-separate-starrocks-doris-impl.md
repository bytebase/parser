# Separate StarRocks and Doris Parsers - Implementation Plan

> **For Claude:** REQUIRED SUB-SKILL: Use superpowers:executing-plans to implement this plan task-by-task.

**Goal:** Rename current doris/ to starrocks/ and create new doris/ with official Apache Doris grammar.

**Architecture:** The current doris/ uses StarRocks grammar. We rename it to starrocks/, then create a new doris/ directory with the official Apache Doris grammar from their GitHub repository.

**Tech Stack:** ANTLR 4, Go, GitHub Actions

**Worktree:** `~/.config/superpowers/worktrees/parser/separate-starrocks-doris`

---

## Task 1: Rename Grammar Files in doris/ to StarRocks naming

**Files:**
- Rename: `doris/DorisSQL.g4` → `doris/StarRocksSQL.g4`
- Rename: `doris/DorisSQLLex.g4` → `doris/StarRocksSQLLex.g4`

**Step 1: Rename the grammar files**

```bash
cd ~/.config/superpowers/worktrees/parser/separate-starrocks-doris
mv doris/DorisSQL.g4 doris/StarRocksSQL.g4
mv doris/DorisSQLLex.g4 doris/StarRocksSQLLex.g4
```

**Step 2: Update grammar declarations inside the files**

In `doris/StarRocksSQL.g4`, change:
```
grammar DorisSQL;
```
to:
```
grammar StarRocksSQL;
```

In `doris/StarRocksSQLLex.g4`, change:
```
lexer grammar DorisSQLLex;
```
to:
```
lexer grammar StarRocksSQLLex;
```

Also update any `import DorisSQLLex` to `import StarRocksSQLLex` in StarRocksSQL.g4.

---

## Task 2: Update Makefile and README in doris/

**Files:**
- Modify: `doris/Makefile`
- Modify: `doris/README.md`

**Step 1: Update Makefile**

Change `doris/Makefile` from:
```makefile
all: build test

build:
	antlr -Dlanguage=Go -package doris -visitor -o . DorisSQLLex.g4 DorisSQL.g4

test:
	go test -v -run TestDorisSQLParser
```

to:
```makefile
all: build test

build:
	antlr -Dlanguage=Go -package starrocks -visitor -o . StarRocksSQLLex.g4 StarRocksSQL.g4

test:
	go test -v -run TestStarRocksSQLParser
```

**Step 2: Update README.md**

Replace all references to "Doris" with "StarRocks" in `doris/README.md`:
- Title: "StarRocks SQL Parser"
- Test name: "TestStarRocksSQLParser"
- Description: parser for StarRocks SQL

---

## Task 3: Update parser_test.go in doris/

**Files:**
- Modify: `doris/parser_test.go`

**Step 1: Update package import and function names**

Change the test file to use StarRocks naming:

```go
package starrocks_test

import (
	"io/ioutil"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/antlr4-go/antlr/v4"
	starrocks "github.com/bytebase/parser/starrocks"
	"github.com/stretchr/testify/require"
)
```

**Step 2: Update lexer/parser instantiation**

Change:
```go
lexer := doris.NewDorisSQLLexer(input)
// ...
p := doris.NewDorisSQLParser(stream)
// ...
tree := p.SqlStatements()
```

to:
```go
lexer := starrocks.NewStarRocksSQLLexer(input)
// ...
p := starrocks.NewStarRocksSQLParser(stream)
// ...
tree := p.SqlStatements()
```

**Step 3: Update test function name**

Change `TestDorisSQLParser` to `TestStarRocksSQLParser`.

---

## Task 4: Delete generated files and rename directory

**Files:**
- Delete: All `doris/dorissql*.go` and `doris/*.interp` and `doris/*.tokens` files
- Rename: `doris/` → `starrocks/`

**Step 1: Delete old generated files**

```bash
cd ~/.config/superpowers/worktrees/parser/separate-starrocks-doris
rm -f doris/dorissql*.go doris/dorissqllex*.go doris/*.interp doris/*.tokens
```

**Step 2: Rename directory**

```bash
git mv doris starrocks
```

---

## Task 5: Regenerate StarRocks parser and verify

**Files:**
- Generate: `starrocks/*.go` files

**Step 1: Build the parser**

```bash
cd ~/.config/superpowers/worktrees/parser/separate-starrocks-doris/starrocks
make build
```

Expected: ANTLR generates new Go files with `starrocks` package name.

**Step 2: Run tests**

```bash
make test
```

Expected: `TestStarRocksSQLParser` passes.

**Step 3: Commit the rename**

```bash
cd ~/.config/superpowers/worktrees/parser/separate-starrocks-doris
git add -A
git commit -m "refactor: rename doris parser to starrocks

The existing doris/ directory actually contains StarRocks SQL grammar.
Rename to accurately reflect the parser source."
```

---

## Task 6: Create new doris/ directory structure

**Files:**
- Create: `doris/` directory
- Create: `doris/examples/` directory

**Step 1: Create directories**

```bash
cd ~/.config/superpowers/worktrees/parser/separate-starrocks-doris
mkdir -p doris/examples
```

---

## Task 7: Download Apache Doris grammar files

**Files:**
- Create: `doris/DorisLexer.g4`
- Create: `doris/DorisParser.g4`

**Step 1: Download grammar files**

```bash
cd ~/.config/superpowers/worktrees/parser/separate-starrocks-doris/doris
curl -o DorisLexer.g4 https://raw.githubusercontent.com/apache/doris/master/fe/fe-core/src/main/antlr4/org/apache/doris/nereids/DorisLexer.g4
curl -o DorisParser.g4 https://raw.githubusercontent.com/apache/doris/master/fe/fe-core/src/main/antlr4/org/apache/doris/nereids/DorisParser.g4
```

**Step 2: Verify files downloaded**

```bash
head -5 DorisLexer.g4 DorisParser.g4
```

Expected: Apache License headers visible.

---

## Task 8: Create Makefile for new doris parser

**Files:**
- Create: `doris/Makefile`

**Step 1: Create Makefile**

```makefile
all: build test

build:
	antlr -Dlanguage=Go -package doris -visitor -o . DorisLexer.g4 DorisParser.g4

test:
	go test -v -run TestDorisParser
```

---

## Task 9: Create README for new doris parser

**Files:**
- Create: `doris/README.md`

**Step 1: Create README.md**

```markdown
# Apache Doris SQL Parser

The doris-parser is a parser for Apache Doris SQL. It is based on [ANTLR4](https://github.com/antlr/antlr4).

## Grammar Source

The grammar files are from the official Apache Doris repository:
https://github.com/apache/doris/tree/master/fe/fe-core/src/main/antlr4/org/apache/doris/nereids

## Build

Before building, you need to install ANTLR4.

Requirements:
- https://github.com/antlr/antlr4/blob/master/doc/getting-started.md
- https://github.com/antlr/antlr4/blob/master/doc/go-target.md

```bash
make build
```

## Test the parser

```bash
make test
```

## References

- Apache Doris: https://doris.apache.org/
- ANTLR4 Getting Started: https://github.com/antlr/antlr4/blob/master/doc/getting-started.md
- ANTLR4 Go Target: https://github.com/antlr/antlr4/blob/master/doc/go-target.md
```

---

## Task 10: Create parser_test.go for new doris parser

**Files:**
- Create: `doris/parser_test.go`

**Step 1: Create the test file**

```go
package doris_test

import (
	"os"
	"path"
	"strings"
	"testing"

	"github.com/antlr4-go/antlr/v4"
	doris "github.com/bytebase/parser/doris"
	"github.com/stretchr/testify/require"
)

type CustomErrorListener struct {
	errors int
}

func NewCustomErrorListener() *CustomErrorListener {
	return new(CustomErrorListener)
}

func (l *CustomErrorListener) SyntaxError(recognizer antlr.Recognizer, offendingSymbol interface{}, line, column int, msg string, e antlr.RecognitionException) {
	l.errors += 1
	antlr.ConsoleErrorListenerINSTANCE.SyntaxError(recognizer, offendingSymbol, line, column, msg, e)
}

func (l *CustomErrorListener) ReportAmbiguity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, exact bool, ambigAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
	antlr.ConsoleErrorListenerINSTANCE.ReportAmbiguity(recognizer, dfa, startIndex, stopIndex, exact, ambigAlts, configs)
}

func (l *CustomErrorListener) ReportAttemptingFullContext(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex int, conflictingAlts *antlr.BitSet, configs *antlr.ATNConfigSet) {
	antlr.ConsoleErrorListenerINSTANCE.ReportAttemptingFullContext(recognizer, dfa, startIndex, stopIndex, conflictingAlts, configs)
}

func (l *CustomErrorListener) ReportContextSensitivity(recognizer antlr.Parser, dfa *antlr.DFA, startIndex, stopIndex, prediction int, configs *antlr.ATNConfigSet) {
	antlr.ConsoleErrorListenerINSTANCE.ReportContextSensitivity(recognizer, dfa, startIndex, stopIndex, prediction, configs)
}

func TestDorisParser(t *testing.T) {
	examples, err := os.ReadDir("examples")
	require.NoError(t, err)

	for _, file := range examples {
		if !strings.HasSuffix(file.Name(), ".sql") {
			continue
		}
		filePath := path.Join("examples", file.Name())
		t.Run(filePath, func(t *testing.T) {
			t.Parallel()
			data, err := os.ReadFile(filePath)
			require.NoError(t, err)

			dataString := strings.TrimRight(string(data), " \t\r\n;") + "\n;"

			input := antlr.NewInputStream(dataString)

			lexer := doris.NewDorisLexer(input)

			stream := antlr.NewCommonTokenStream(lexer, 0)
			p := doris.NewDorisParser(stream)

			lexerErrors := &CustomErrorListener{}
			lexer.RemoveErrorListeners()
			lexer.AddErrorListener(lexerErrors)

			parserErrors := &CustomErrorListener{}
			p.RemoveErrorListeners()
			p.AddErrorListener(parserErrors)

			p.BuildParseTrees = true

			tree := p.MultiStatements()

			require.Equal(t, 0, lexerErrors.errors)
			require.Equal(t, 0, parserErrors.errors)

			require.Equal(t, dataString, stream.GetTextFromRuleContext(tree))
		})
	}
}
```

Note: The entry point is `p.MultiStatements()` (from Apache Doris grammar) instead of `p.SqlStatements()`.

---

## Task 11: Create test SQL examples - SELECT statements

**Files:**
- Create: `doris/examples/select.sql`

**Step 1: Create select.sql with examples from Doris documentation**

```sql
SELECT Name FROM student WHERE age IN (18,20,25);
SELECT * EXCEPT(age) FROM student;
SELECT type, AVG(price) FROM tb_book GROUP BY type;
SELECT DISTINCT type FROM tb_book;
SELECT * FROM tb_book ORDER BY id DESC LIMIT 3;
SELECT * FROM tb_book WHERE name LIKE('_h%');
SELECT * FROM tb_book ORDER BY price DESC LIMIT 3;
SELECT id, CONCAT(name, ":", price) AS info, type FROM tb_book;
SELECT SUM(price) AS total, type FROM tb_book GROUP BY type;
SELECT *, (price * 0.8) AS "20%" FROM tb_book;
WITH cte AS (SELECT 1 AS col1, 2 AS col2 UNION ALL SELECT 3, 4) SELECT col1, col2 FROM cte;
SELECT * FROM t1 LEFT JOIN (t2, t3, t4) ON (t2.a = t1.a AND t3.b = t1.b AND t4.c = t1.c);
SELECT * FROM t1 LEFT JOIN (t2 CROSS JOIN t3 CROSS JOIN t4) ON (t2.a = t1.a AND t3.b = t1.b AND t4.c = t1.c);
SELECT t1.name, t2.salary FROM employee AS t1 INNER JOIN info AS t2 ON t1.name = t2.name;
SELECT left_tbl.* FROM left_tbl LEFT JOIN right_tbl ON left_tbl.id = right_tbl.id WHERE right_tbl.id IS NULL;
SELECT * FROM t1 RIGHT JOIN t2 ON (t1.a = t2.a);
SELECT college, region, seed FROM tournament ORDER BY region, seed;
SELECT college, region AS r, seed AS s FROM tournament ORDER BY r, s;
SELECT college, region, seed FROM tournament ORDER BY 2, 3;
SELECT a, COUNT(b) FROM test_table GROUP BY a ORDER BY NULL;
SELECT COUNT(col1) AS col2 FROM t GROUP BY col2 HAVING col2 = 2;
SELECT user, MAX(salary) FROM users GROUP BY user HAVING MAX(salary) > 10;
SELECT * FROM tbl LIMIT 5, 10;
SELECT * FROM (SELECT age FROM student_01 UNION ALL SELECT age FROM student_02) AS t1 ORDER BY age LIMIT 4
```

---

## Task 12: Create test SQL examples - INSERT statements

**Files:**
- Create: `doris/examples/insert.sql`

**Step 1: Create insert.sql with examples from Doris documentation**

```sql
INSERT INTO test VALUES (1, 2);
INSERT INTO test (c1, c2) VALUES (1, 2);
INSERT INTO test (c1, c2) VALUES (1, DEFAULT);
INSERT INTO test (c1) VALUES (1);
INSERT INTO test VALUES (1, 2), (3, 2 + 2);
INSERT INTO test (c1, c2) VALUES (1, 2), (3, 2 * 2);
INSERT INTO test (c1) VALUES (1), (3);
INSERT INTO test (c1, c2) VALUES (1, DEFAULT), (3, DEFAULT);
INSERT INTO test SELECT * FROM test2;
INSERT INTO test (c1, c2) SELECT * FROM test2
```

---

## Task 13: Create test SQL examples - CREATE TABLE statements

**Files:**
- Create: `doris/examples/create_table.sql`

**Step 1: Create create_table.sql with examples from Doris documentation**

```sql
CREATE TABLE t1(c1 INT, c2 STRING)
DUPLICATE KEY(c1)
DISTRIBUTED BY HASH(c1)
PROPERTIES('replication_num' = '1');

CREATE TABLE t2(c1 INT, c2 INT MAX)
AGGREGATE KEY(c1)
DISTRIBUTED BY HASH(c1)
PROPERTIES('replication_num' = '1');

CREATE TABLE t3(c1 INT, c2 INT)
UNIQUE KEY(c1)
DISTRIBUTED BY HASH(c1)
PROPERTIES('replication_num' = '1');

CREATE TABLE t5(c1 INT, c2 INT DEFAULT 10)
DUPLICATE KEY(c1)
DISTRIBUTED BY HASH(c1)
PROPERTIES('replication_num' = '1');

CREATE TABLE t6(c1 INT, c2 INT)
DUPLICATE KEY(c1)
DISTRIBUTED BY RANDOM
PROPERTIES('replication_num' = '1');

CREATE TABLE example_db.table_hash(k1 BIGINT, k2 LARGEINT, v1 VARCHAR(2048), v2 SMALLINT DEFAULT "10")
UNIQUE KEY(k1, k2)
DISTRIBUTED BY HASH(k1, k2) BUCKETS 32
PROPERTIES("storage_medium" = "SSD", "storage_cooldown_time" = "2015-06-04 00:00:00");

CREATE TABLE t10
PROPERTIES('replication_num' = '1')
AS SELECT * FROM t1;

CREATE TABLE t11 LIKE t10
```

---

## Task 14: Build the new Doris parser

**Files:**
- Generate: `doris/*.go` files

**Step 1: Run ANTLR to generate parser**

```bash
cd ~/.config/superpowers/worktrees/parser/separate-starrocks-doris/doris
make build
```

Expected: ANTLR generates Go files for the Doris parser.

**Step 2: Check if additional grammar files are needed**

If build fails with missing imports, download additional grammar files from Apache Doris repo.

---

## Task 15: Test the new Doris parser

**Files:**
- Test: `doris/parser_test.go`

**Step 1: Run tests**

```bash
cd ~/.config/superpowers/worktrees/parser/separate-starrocks-doris/doris
make test
```

**Step 2: Fix any parsing errors**

If tests fail, examine the error messages and:
1. Check if the SQL examples match the grammar's expected syntax
2. Adjust examples or investigate grammar issues

**Step 3: Commit the new parser**

```bash
cd ~/.config/superpowers/worktrees/parser/separate-starrocks-doris
git add doris/
git commit -m "feat: add Apache Doris parser with official grammar

Add new doris/ directory using the official Apache Doris grammar from:
https://github.com/apache/doris/tree/master/fe/fe-core/src/main/antlr4/org/apache/doris/nereids

Includes test examples from Doris documentation for:
- SELECT statements
- INSERT statements
- CREATE TABLE statements"
```

---

## Task 16: Update GitHub Actions workflow

**Files:**
- Modify: `.github/workflows/tests.yml`

**Step 1: Add starrocks to the parser list**

In `.github/workflows/tests.yml`, line 37, change:
```
ALL_PARSERS="redshift postgresql cql snowflake tsql doris trino plsql googlesql mysql partiql tidb mariadb cosmosdb"
```

to:
```
ALL_PARSERS="redshift postgresql cql snowflake tsql doris starrocks trino plsql googlesql mysql partiql tidb mariadb cosmosdb"
```

**Step 2: Commit the CI update**

```bash
cd ~/.config/superpowers/worktrees/parser/separate-starrocks-doris
git add .github/workflows/tests.yml
git commit -m "ci: add starrocks parser to test matrix"
```

---

## Task 17: Final verification

**Step 1: Run all tests**

```bash
cd ~/.config/superpowers/worktrees/parser/separate-starrocks-doris
cd starrocks && make test && cd ..
cd doris && make test && cd ..
```

Expected: Both parsers pass all tests.

**Step 2: Verify git status is clean**

```bash
git status
```

Expected: Working tree clean, all changes committed.

**Step 3: Review commit history**

```bash
git log --oneline -5
```

Expected: 4 commits:
1. docs: add design for separating StarRocks and Doris parsers
2. refactor: rename doris parser to starrocks
3. feat: add Apache Doris parser with official grammar
4. ci: add starrocks parser to test matrix

---

## Summary

| Task | Description |
|------|-------------|
| 1-5 | Rename doris/ to starrocks/ with all internal updates |
| 6-13 | Create new doris/ with Apache Doris grammar and test examples |
| 14-15 | Build and test the new Doris parser |
| 16 | Update CI to include starrocks |
| 17 | Final verification |
