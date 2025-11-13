-- Regression test: Function names should be usable as identifiers
-- Issue: After removing builtin_function_name tokens, mathematical function names
-- should work as regular identifiers (table names, column names, etc.)

-- Mathematical functions as column names
CREATE TABLE test_math (
    exp INT,
    div INT,
    floor INT,
    mod INT,
    power INT,
    sqrt INT,
    log INT
);

-- String function REVERSE as column name
CREATE TABLE test_string (
    reverse VARCHAR(100)
);

-- Function names as table names
CREATE TABLE exp (id INT);
CREATE TABLE div (id INT);

-- Select from tables with function names
SELECT exp, div, floor, mod FROM test_math;
SELECT reverse FROM test_string;
SELECT * FROM exp;
