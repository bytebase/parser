-- Valid. _5abc and dataField are valid identifiers.
SELECT _5abc.dataField FROM _5abc;

-- Valid. `5abc` and dataField are valid identifiers.
SELECT `5abc`.dataField FROM `5abc`;

-- Invalid. 5abc is an invalid identifier because it is unquoted and starts
-- with a number rather than a letter or underscore.
-- SELECT 5abc.dataField FROM 5abc;

-- Valid. abc5 and dataField are valid identifiers.
SELECT abc5.dataField FROM abc5;

-- Invalid. abc5! is an invalid identifier because it is unquoted and contains
-- a character that is not a letter, number, or underscore.
-- SELECT abc5!.dataField FROM abc5!;

-- Valid. `GROUP` and dataField are valid identifiers.
SELECT `GROUP`.dataField FROM `GROUP`;

-- Invalid. GROUP is an invalid identifier because it is unquoted and is a
-- stand-alone reserved keyword.
-- SELECT GROUP.dataField FROM GROUP;

-- Valid. abc5 and GROUP are valid identifiers.
SELECT abc5.GROUP FROM abc5;