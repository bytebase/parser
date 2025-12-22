-- Test bare column labels (aliases without AS keyword)
-- These should parse successfully because 'name', 'value', 'action' etc are bare_label_keywords

-- Basic bare labels with unreserved keywords
SELECT 1 name;
SELECT 1 value;
SELECT 1 action;
SELECT 1 data;

-- Bare labels with reserved keywords that are also bare_label_keywords
SELECT 1 all;
SELECT 1 table;
SELECT 1 select;

-- Bare labels with col_name_keywords
SELECT 1 int;
SELECT 1 timestamp;
SELECT 1 json;

-- Bare labels with type_func_name_keywords
SELECT 1 left;
SELECT 1 right;
SELECT 1 join;

-- Multiple columns with bare labels
SELECT 1 name, 2 value, 3 action;

-- Mixed with and without AS
SELECT 1 AS alias1, 2 alias2, 3 AS alias3;

-- Expression with bare label
SELECT 1 + 2 result;
SELECT a.id name FROM t a;

-- Subquery with bare label
SELECT * FROM (SELECT 1 name) sub;

-- NOTE: The following require AS keyword (AS_LABEL keywords):
-- SELECT 1 AS year;   -- 'year' is AS_LABEL, needs AS
-- SELECT 1 AS month;  -- 'month' is AS_LABEL, needs AS
-- SELECT 1 AS day;    -- 'day' is AS_LABEL, needs AS
-- SELECT 1 AS hour;   -- 'hour' is AS_LABEL, needs AS
-- SELECT 1 AS char;   -- 'char' is AS_LABEL, needs AS
