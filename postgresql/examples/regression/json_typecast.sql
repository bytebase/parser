-- Regression test: JSON and JSONB type casting
-- Issue: Missing jsontype rule caused ::json and ::jsonb typecasts to fail
-- Fix: Added jsontype rule to simpletypename and consttypename

-- Basic JSON typecast
SELECT '{"a": 1}'::json;
SELECT '{"b": 2}'::jsonb;

-- JSON with unicode escapes
SELECT '"\u0000"'::json;
SELECT '"\uaBcD"'::json;

-- JSON with surrogate pairs
SELECT '{ "a":  "\ud83d\ude04\ud83d\udc36" }'::json;

-- JSON in table definition
CREATE TABLE json_test (
    data json,
    data_b jsonb
);

-- JSON typecast in expressions
SELECT '{"a": 1}'::json -> 'a';
SELECT '{"b": 2}'::jsonb -> 'b';
