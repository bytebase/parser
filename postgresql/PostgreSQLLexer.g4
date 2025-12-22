/*
based on
https://github.com/tunnelvisionlabs/antlr4-grammar-postgresql/blob/master/src/com/tunnelvisionlabs/postgresql/PostgreSqlLexer.g4
*/

/*
 * [The "MIT license"]
 * Copyright (C) 2014 Sam Harwell, Tunnel Vision Laboratories, LLC
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * 1. The above copyright notice and this permission notice shall be included in
 *    all copies or substantial portions of the Software.
 * 2. THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 *    IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 *    FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL
 *    THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 *    LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING
 *    FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER
 *    DEALINGS IN THE SOFTWARE.
 * 3. Except as contained in this notice, the name of Tunnel Vision
 *    Laboratories, LLC. shall not be used in advertising or otherwise to
 *    promote the sale, use or other dealings in this Software without prior
 *    written authorization from Tunnel Vision Laboratories, LLC.
 */
lexer grammar PostgreSQLLexer;
/* Reference:
 * http://www.postgresql.org/docs/9.3/static/sql-syntax-lexical.html
 */

options {
   superClass = PostgreSQLLexerBase;
   caseInsensitive = true;
}

@ header
{
}
@ members
{
/* This field stores the tags which are used to detect the end of a dollar-quoted string literal.
 */
}

// ============================================================================
// BEGIN AUTO-GENERATED KEYWORDS
//
// Source: PostgreSQL REL_18_STABLE kwlist.h
// URL: https://raw.githubusercontent.com/postgres/postgres/REL_18_STABLE/src/include/parser/kwlist.h
// Generated: 2025-12-22T17:51:41+08:00
// Total Keywords: 494
//
// NOTE: These keyword rules must appear BEFORE the Identifier rule
// to ensure keywords are matched with higher priority than identifiers.
//

ABORT_P
   : 'ABORT'
   ;

ABSENT
   : 'ABSENT'
   ;

ABSOLUTE_P
   : 'ABSOLUTE'
   ;

ACCESS
   : 'ACCESS'
   ;

ACTION
   : 'ACTION'
   ;

ADD_P
   : 'ADD'
   ;

ADMIN
   : 'ADMIN'
   ;

AFTER
   : 'AFTER'
   ;

AGGREGATE
   : 'AGGREGATE'
   ;

ALL
   : 'ALL'
   ;

ALSO
   : 'ALSO'
   ;

ALTER
   : 'ALTER'
   ;

ALWAYS
   : 'ALWAYS'
   ;

ANALYSE
   : 'ANALYSE'
   ;

ANALYZE
   : 'ANALYZE'
   ;

AND
   : 'AND'
   ;

ANY
   : 'ANY'
   ;

ARRAY
   : 'ARRAY'
   ;

AS
   : 'AS'
   ;

ASC
   : 'ASC'
   ;

ASENSITIVE
   : 'ASENSITIVE'
   ;

ASSERTION
   : 'ASSERTION'
   ;

ASSIGNMENT
   : 'ASSIGNMENT'
   ;

ASYMMETRIC
   : 'ASYMMETRIC'
   ;

AT
   : 'AT'
   ;

ATOMIC
   : 'ATOMIC'
   ;

ATTACH
   : 'ATTACH'
   ;

ATTRIBUTE
   : 'ATTRIBUTE'
   ;

AUTHORIZATION
   : 'AUTHORIZATION'
   ;

BACKWARD
   : 'BACKWARD'
   ;

BEFORE
   : 'BEFORE'
   ;

BEGIN_P
   : 'BEGIN'
   ;

BETWEEN
   : 'BETWEEN'
   ;

BIGINT
   : 'BIGINT'
   ;

BINARY
   : 'BINARY'
   ;

BIT
   : 'BIT'
   ;

BOOLEAN_P
   : 'BOOLEAN'
   ;

BOTH
   : 'BOTH'
   ;

BREADTH
   : 'BREADTH'
   ;

BY
   : 'BY'
   ;

CACHE
   : 'CACHE'
   ;

CALL
   : 'CALL'
   ;

CALLED
   : 'CALLED'
   ;

CASCADE
   : 'CASCADE'
   ;

CASCADED
   : 'CASCADED'
   ;

CASE
   : 'CASE'
   ;

CAST
   : 'CAST'
   ;

CATALOG_P
   : 'CATALOG'
   ;

CHAIN
   : 'CHAIN'
   ;

CHAR_P
   : 'CHAR'
   ;

CHARACTER
   : 'CHARACTER'
   ;

CHARACTERISTICS
   : 'CHARACTERISTICS'
   ;

CHECK
   : 'CHECK'
   ;

CHECKPOINT
   : 'CHECKPOINT'
   ;

CLASS
   : 'CLASS'
   ;

CLOSE
   : 'CLOSE'
   ;

CLUSTER
   : 'CLUSTER'
   ;

COALESCE
   : 'COALESCE'
   ;

COLLATE
   : 'COLLATE'
   ;

COLLATION
   : 'COLLATION'
   ;

COLUMN
   : 'COLUMN'
   ;

COLUMNS
   : 'COLUMNS'
   ;

COMMENT
   : 'COMMENT'
   ;

COMMENTS
   : 'COMMENTS'
   ;

COMMIT
   : 'COMMIT'
   ;

COMMITTED
   : 'COMMITTED'
   ;

COMPRESSION
   : 'COMPRESSION'
   ;

CONCURRENTLY
   : 'CONCURRENTLY'
   ;

CONDITIONAL
   : 'CONDITIONAL'
   ;

CONFIGURATION
   : 'CONFIGURATION'
   ;

CONFLICT
   : 'CONFLICT'
   ;

CONNECTION
   : 'CONNECTION'
   ;

CONSTRAINT
   : 'CONSTRAINT'
   ;

CONSTRAINTS
   : 'CONSTRAINTS'
   ;

CONTENT_P
   : 'CONTENT'
   ;

CONTINUE_P
   : 'CONTINUE'
   ;

CONVERSION_P
   : 'CONVERSION'
   ;

COPY
   : 'COPY'
   ;

COST
   : 'COST'
   ;

CREATE
   : 'CREATE'
   ;

CROSS
   : 'CROSS'
   ;

CSV
   : 'CSV'
   ;

CUBE
   : 'CUBE'
   ;

CURRENT_P
   : 'CURRENT'
   ;

CURRENT_CATALOG
   : 'CURRENT_CATALOG'
   ;

CURRENT_DATE
   : 'CURRENT_DATE'
   ;

CURRENT_ROLE
   : 'CURRENT_ROLE'
   ;

CURRENT_SCHEMA
   : 'CURRENT_SCHEMA'
   ;

CURRENT_TIME
   : 'CURRENT_TIME'
   ;

CURRENT_TIMESTAMP
   : 'CURRENT_TIMESTAMP'
   ;

CURRENT_USER
   : 'CURRENT_USER'
   ;

CURSOR
   : 'CURSOR'
   ;

CYCLE
   : 'CYCLE'
   ;

DATA_P
   : 'DATA'
   ;

DATABASE
   : 'DATABASE'
   ;

DAY_P
   : 'DAY'
   ;

DEALLOCATE
   : 'DEALLOCATE'
   ;

DEC
   : 'DEC'
   ;

DECIMAL_P
   : 'DECIMAL'
   ;

DECLARE
   : 'DECLARE'
   ;

DEFAULT
   : 'DEFAULT'
   ;

DEFAULTS
   : 'DEFAULTS'
   ;

DEFERRABLE
   : 'DEFERRABLE'
   ;

DEFERRED
   : 'DEFERRED'
   ;

DEFINER
   : 'DEFINER'
   ;

DELETE_P
   : 'DELETE'
   ;

DELIMITER
   : 'DELIMITER'
   ;

DELIMITERS
   : 'DELIMITERS'
   ;

DEPENDS
   : 'DEPENDS'
   ;

DEPTH
   : 'DEPTH'
   ;

DESC
   : 'DESC'
   ;

DETACH
   : 'DETACH'
   ;

DICTIONARY
   : 'DICTIONARY'
   ;

DISABLE_P
   : 'DISABLE'
   ;

DISCARD
   : 'DISCARD'
   ;

DISTINCT
   : 'DISTINCT'
   ;

DO
   : 'DO'
   ;

DOCUMENT_P
   : 'DOCUMENT'
   ;

DOMAIN_P
   : 'DOMAIN'
   ;

DOUBLE_P
   : 'DOUBLE'
   ;

DROP
   : 'DROP'
   ;

EACH
   : 'EACH'
   ;

ELSE
   : 'ELSE'
   ;

EMPTY_P
   : 'EMPTY'
   ;

ENABLE_P
   : 'ENABLE'
   ;

ENCODING
   : 'ENCODING'
   ;

ENCRYPTED
   : 'ENCRYPTED'
   ;

END_P
   : 'END'
   ;

ENFORCED
   : 'ENFORCED'
   ;

ENUM_P
   : 'ENUM'
   ;

ERROR_P
   : 'ERROR'
   ;

ESCAPE
   : 'ESCAPE'
   ;

EVENT
   : 'EVENT'
   ;

EXCEPT
   : 'EXCEPT'
   ;

EXCLUDE
   : 'EXCLUDE'
   ;

EXCLUDING
   : 'EXCLUDING'
   ;

EXCLUSIVE
   : 'EXCLUSIVE'
   ;

EXECUTE
   : 'EXECUTE'
   ;

EXISTS
   : 'EXISTS'
   ;

EXPLAIN
   : 'EXPLAIN'
   ;

EXPRESSION
   : 'EXPRESSION'
   ;

EXTENSION
   : 'EXTENSION'
   ;

EXTERNAL
   : 'EXTERNAL'
   ;

EXTRACT
   : 'EXTRACT'
   ;

FALSE_P
   : 'FALSE'
   ;

FAMILY
   : 'FAMILY'
   ;

FETCH
   : 'FETCH'
   ;

FILTER
   : 'FILTER'
   ;

FINALIZE
   : 'FINALIZE'
   ;

FIRST_P
   : 'FIRST'
   ;

FLOAT_P
   : 'FLOAT'
   ;

FOLLOWING
   : 'FOLLOWING'
   ;

FOR
   : 'FOR'
   ;

FORCE
   : 'FORCE'
   ;

FOREIGN
   : 'FOREIGN'
   ;

FORMAT
   : 'FORMAT'
   ;

FORWARD
   : 'FORWARD'
   ;

FREEZE
   : 'FREEZE'
   ;

FROM
   : 'FROM'
   ;

FULL
   : 'FULL'
   ;

FUNCTION
   : 'FUNCTION'
   ;

FUNCTIONS
   : 'FUNCTIONS'
   ;

GENERATED
   : 'GENERATED'
   ;

GLOBAL
   : 'GLOBAL'
   ;

GRANT
   : 'GRANT'
   ;

GRANTED
   : 'GRANTED'
   ;

GREATEST
   : 'GREATEST'
   ;

GROUP_P
   : 'GROUP'
   ;

GROUPING
   : 'GROUPING'
   ;

GROUPS
   : 'GROUPS'
   ;

HANDLER
   : 'HANDLER'
   ;

HAVING
   : 'HAVING'
   ;

HEADER_P
   : 'HEADER'
   ;

HOLD
   : 'HOLD'
   ;

HOUR_P
   : 'HOUR'
   ;

IDENTITY_P
   : 'IDENTITY'
   ;

IF_P
   : 'IF'
   ;

ILIKE
   : 'ILIKE'
   ;

IMMEDIATE
   : 'IMMEDIATE'
   ;

IMMUTABLE
   : 'IMMUTABLE'
   ;

IMPLICIT_P
   : 'IMPLICIT'
   ;

IMPORT_P
   : 'IMPORT'
   ;

IN_P
   : 'IN'
   ;

INCLUDE
   : 'INCLUDE'
   ;

INCLUDING
   : 'INCLUDING'
   ;

INCREMENT
   : 'INCREMENT'
   ;

INDENT
   : 'INDENT'
   ;

INDEX
   : 'INDEX'
   ;

INDEXES
   : 'INDEXES'
   ;

INHERIT
   : 'INHERIT'
   ;

INHERITS
   : 'INHERITS'
   ;

INITIALLY
   : 'INITIALLY'
   ;

INLINE_P
   : 'INLINE'
   ;

INNER_P
   : 'INNER'
   ;

INOUT
   : 'INOUT'
   ;

INPUT_P
   : 'INPUT'
   ;

INSENSITIVE
   : 'INSENSITIVE'
   ;

INSERT
   : 'INSERT'
   ;

INSTEAD
   : 'INSTEAD'
   ;

INT_P
   : 'INT'
   ;

INTEGER
   : 'INTEGER'
   ;

INTERSECT
   : 'INTERSECT'
   ;

INTERVAL
   : 'INTERVAL'
   ;

INTO
   : 'INTO'
   ;

INVOKER
   : 'INVOKER'
   ;

IS
   : 'IS'
   ;

ISNULL
   : 'ISNULL'
   ;

ISOLATION
   : 'ISOLATION'
   ;

JOIN
   : 'JOIN'
   ;

JSON
   : 'JSON'
   ;

JSON_ARRAY
   : 'JSON_ARRAY'
   ;

JSON_ARRAYAGG
   : 'JSON_ARRAYAGG'
   ;

JSON_EXISTS
   : 'JSON_EXISTS'
   ;

JSON_OBJECT
   : 'JSON_OBJECT'
   ;

JSON_OBJECTAGG
   : 'JSON_OBJECTAGG'
   ;

JSON_QUERY
   : 'JSON_QUERY'
   ;

JSON_SCALAR
   : 'JSON_SCALAR'
   ;

JSON_SERIALIZE
   : 'JSON_SERIALIZE'
   ;

JSON_TABLE
   : 'JSON_TABLE'
   ;

JSON_VALUE
   : 'JSON_VALUE'
   ;

KEEP
   : 'KEEP'
   ;

KEY
   : 'KEY'
   ;

KEYS
   : 'KEYS'
   ;

LABEL
   : 'LABEL'
   ;

LANGUAGE
   : 'LANGUAGE'
   ;

LARGE_P
   : 'LARGE'
   ;

LAST_P
   : 'LAST'
   ;

LATERAL_P
   : 'LATERAL'
   ;

LEADING
   : 'LEADING'
   ;

LEAKPROOF
   : 'LEAKPROOF'
   ;

LEAST
   : 'LEAST'
   ;

LEFT
   : 'LEFT'
   ;

LEVEL
   : 'LEVEL'
   ;

LIKE
   : 'LIKE'
   ;

LIMIT
   : 'LIMIT'
   ;

LISTEN
   : 'LISTEN'
   ;

LOAD
   : 'LOAD'
   ;

LOCAL
   : 'LOCAL'
   ;

LOCALTIME
   : 'LOCALTIME'
   ;

LOCALTIMESTAMP
   : 'LOCALTIMESTAMP'
   ;

LOCATION
   : 'LOCATION'
   ;

LOCK_P
   : 'LOCK'
   ;

LOCKED
   : 'LOCKED'
   ;

LOGGED
   : 'LOGGED'
   ;

MAPPING
   : 'MAPPING'
   ;

MATCH
   : 'MATCH'
   ;

MATCHED
   : 'MATCHED'
   ;

MATERIALIZED
   : 'MATERIALIZED'
   ;

MAXVALUE
   : 'MAXVALUE'
   ;

MERGE
   : 'MERGE'
   ;

MERGE_ACTION
   : 'MERGE_ACTION'
   ;

METHOD
   : 'METHOD'
   ;

MINUTE_P
   : 'MINUTE'
   ;

MINVALUE
   : 'MINVALUE'
   ;

MODE
   : 'MODE'
   ;

MONTH_P
   : 'MONTH'
   ;

MOVE
   : 'MOVE'
   ;

NAME_P
   : 'NAME'
   ;

NAMES
   : 'NAMES'
   ;

NATIONAL
   : 'NATIONAL'
   ;

NATURAL
   : 'NATURAL'
   ;

NCHAR
   : 'NCHAR'
   ;

NESTED
   : 'NESTED'
   ;

NEW
   : 'NEW'
   ;

NEXT
   : 'NEXT'
   ;

NFC
   : 'NFC'
   ;

NFD
   : 'NFD'
   ;

NFKC
   : 'NFKC'
   ;

NFKD
   : 'NFKD'
   ;

NO
   : 'NO'
   ;

NONE
   : 'NONE'
   ;

NORMALIZE
   : 'NORMALIZE'
   ;

NORMALIZED
   : 'NORMALIZED'
   ;

NOT
   : 'NOT'
   ;

NOTHING
   : 'NOTHING'
   ;

NOTIFY
   : 'NOTIFY'
   ;

NOTNULL
   : 'NOTNULL'
   ;

NOWAIT
   : 'NOWAIT'
   ;

NULL_P
   : 'NULL'
   ;

NULLIF
   : 'NULLIF'
   ;

NULLS_P
   : 'NULLS'
   ;

NUMERIC
   : 'NUMERIC'
   ;

OBJECT_P
   : 'OBJECT'
   ;

OBJECTS_P
   : 'OBJECTS'
   ;

OF
   : 'OF'
   ;

OFF
   : 'OFF'
   ;

OFFSET
   : 'OFFSET'
   ;

OIDS
   : 'OIDS'
   ;

OLD
   : 'OLD'
   ;

OMIT
   : 'OMIT'
   ;

ON
   : 'ON'
   ;

ONLY
   : 'ONLY'
   ;

OPERATOR
   : 'OPERATOR'
   ;

OPTION
   : 'OPTION'
   ;

OPTIONS
   : 'OPTIONS'
   ;

OR
   : 'OR'
   ;

ORDER
   : 'ORDER'
   ;

ORDINALITY
   : 'ORDINALITY'
   ;

OTHERS
   : 'OTHERS'
   ;

OUT_P
   : 'OUT'
   ;

OUTER_P
   : 'OUTER'
   ;

OVER
   : 'OVER'
   ;

OVERLAPS
   : 'OVERLAPS'
   ;

OVERLAY
   : 'OVERLAY'
   ;

OVERRIDING
   : 'OVERRIDING'
   ;

OWNED
   : 'OWNED'
   ;

OWNER
   : 'OWNER'
   ;

PARALLEL
   : 'PARALLEL'
   ;

PARAMETER
   : 'PARAMETER'
   ;

PARSER
   : 'PARSER'
   ;

PARTIAL
   : 'PARTIAL'
   ;

PARTITION
   : 'PARTITION'
   ;

PASSING
   : 'PASSING'
   ;

PASSWORD
   : 'PASSWORD'
   ;

PATH
   : 'PATH'
   ;

PERIOD
   : 'PERIOD'
   ;

PLACING
   : 'PLACING'
   ;

PLAN
   : 'PLAN'
   ;

PLANS
   : 'PLANS'
   ;

POLICY
   : 'POLICY'
   ;

POSITION
   : 'POSITION'
   ;

PRECEDING
   : 'PRECEDING'
   ;

PRECISION
   : 'PRECISION'
   ;

PREPARE
   : 'PREPARE'
   ;

PREPARED
   : 'PREPARED'
   ;

PRESERVE
   : 'PRESERVE'
   ;

PRIMARY
   : 'PRIMARY'
   ;

PRIOR
   : 'PRIOR'
   ;

PRIVILEGES
   : 'PRIVILEGES'
   ;

PROCEDURAL
   : 'PROCEDURAL'
   ;

PROCEDURE
   : 'PROCEDURE'
   ;

PROCEDURES
   : 'PROCEDURES'
   ;

PROGRAM
   : 'PROGRAM'
   ;

PUBLICATION
   : 'PUBLICATION'
   ;

QUOTE
   : 'QUOTE'
   ;

QUOTES
   : 'QUOTES'
   ;

RANGE
   : 'RANGE'
   ;

READ
   : 'READ'
   ;

REAL
   : 'REAL'
   ;

REASSIGN
   : 'REASSIGN'
   ;

RECURSIVE
   : 'RECURSIVE'
   ;

REF_P
   : 'REF'
   ;

REFERENCES
   : 'REFERENCES'
   ;

REFERENCING
   : 'REFERENCING'
   ;

REFRESH
   : 'REFRESH'
   ;

REINDEX
   : 'REINDEX'
   ;

RELATIVE_P
   : 'RELATIVE'
   ;

RELEASE
   : 'RELEASE'
   ;

RENAME
   : 'RENAME'
   ;

REPEATABLE
   : 'REPEATABLE'
   ;

REPLACE
   : 'REPLACE'
   ;

REPLICA
   : 'REPLICA'
   ;

RESET
   : 'RESET'
   ;

RESTART
   : 'RESTART'
   ;

RESTRICT
   : 'RESTRICT'
   ;

RETURN
   : 'RETURN'
   ;

RETURNING
   : 'RETURNING'
   ;

RETURNS
   : 'RETURNS'
   ;

REVOKE
   : 'REVOKE'
   ;

RIGHT
   : 'RIGHT'
   ;

ROLE
   : 'ROLE'
   ;

ROLLBACK
   : 'ROLLBACK'
   ;

ROLLUP
   : 'ROLLUP'
   ;

ROUTINE
   : 'ROUTINE'
   ;

ROUTINES
   : 'ROUTINES'
   ;

ROW
   : 'ROW'
   ;

ROWS
   : 'ROWS'
   ;

RULE
   : 'RULE'
   ;

SAVEPOINT
   : 'SAVEPOINT'
   ;

SCALAR
   : 'SCALAR'
   ;

SCHEMA
   : 'SCHEMA'
   ;

SCHEMAS
   : 'SCHEMAS'
   ;

SCROLL
   : 'SCROLL'
   ;

SEARCH
   : 'SEARCH'
   ;

SECOND_P
   : 'SECOND'
   ;

SECURITY
   : 'SECURITY'
   ;

SELECT
   : 'SELECT'
   ;

SEQUENCE
   : 'SEQUENCE'
   ;

SEQUENCES
   : 'SEQUENCES'
   ;

SERIALIZABLE
   : 'SERIALIZABLE'
   ;

SERVER
   : 'SERVER'
   ;

SESSION
   : 'SESSION'
   ;

SESSION_USER
   : 'SESSION_USER'
   ;

SET
   : 'SET'
   ;

SETOF
   : 'SETOF'
   ;

SETS
   : 'SETS'
   ;

SHARE
   : 'SHARE'
   ;

SHOW
   : 'SHOW'
   ;

SIMILAR
   : 'SIMILAR'
   ;

SIMPLE
   : 'SIMPLE'
   ;

SKIP_P
   : 'SKIP'
   ;

SMALLINT
   : 'SMALLINT'
   ;

SNAPSHOT
   : 'SNAPSHOT'
   ;

SOME
   : 'SOME'
   ;

SOURCE
   : 'SOURCE'
   ;

SQL_P
   : 'SQL'
   ;

STABLE
   : 'STABLE'
   ;

STANDALONE_P
   : 'STANDALONE'
   ;

START
   : 'START'
   ;

STATEMENT
   : 'STATEMENT'
   ;

STATISTICS
   : 'STATISTICS'
   ;

STDIN
   : 'STDIN'
   ;

STDOUT
   : 'STDOUT'
   ;

STORAGE
   : 'STORAGE'
   ;

STORED
   : 'STORED'
   ;

STRICT_P
   : 'STRICT'
   ;

STRING_P
   : 'STRING'
   ;

STRIP_P
   : 'STRIP'
   ;

SUBSCRIPTION
   : 'SUBSCRIPTION'
   ;

SUBSTRING
   : 'SUBSTRING'
   ;

SUPPORT
   : 'SUPPORT'
   ;

SYMMETRIC
   : 'SYMMETRIC'
   ;

SYSID
   : 'SYSID'
   ;

SYSTEM_P
   : 'SYSTEM'
   ;

SYSTEM_USER
   : 'SYSTEM_USER'
   ;

TABLE
   : 'TABLE'
   ;

TABLES
   : 'TABLES'
   ;

TABLESAMPLE
   : 'TABLESAMPLE'
   ;

TABLESPACE
   : 'TABLESPACE'
   ;

TARGET
   : 'TARGET'
   ;

TEMP
   : 'TEMP'
   ;

TEMPLATE
   : 'TEMPLATE'
   ;

TEMPORARY
   : 'TEMPORARY'
   ;

TEXT_P
   : 'TEXT'
   ;

THEN
   : 'THEN'
   ;

TIES
   : 'TIES'
   ;

TIME
   : 'TIME'
   ;

TIMESTAMP
   : 'TIMESTAMP'
   ;

TO
   : 'TO'
   ;

TRAILING
   : 'TRAILING'
   ;

TRANSACTION
   : 'TRANSACTION'
   ;

TRANSFORM
   : 'TRANSFORM'
   ;

TREAT
   : 'TREAT'
   ;

TRIGGER
   : 'TRIGGER'
   ;

TRIM
   : 'TRIM'
   ;

TRUE_P
   : 'TRUE'
   ;

TRUNCATE
   : 'TRUNCATE'
   ;

TRUSTED
   : 'TRUSTED'
   ;

TYPE_P
   : 'TYPE'
   ;

TYPES_P
   : 'TYPES'
   ;

UESCAPE
   : 'UESCAPE'
   ;

UNBOUNDED
   : 'UNBOUNDED'
   ;

UNCOMMITTED
   : 'UNCOMMITTED'
   ;

UNCONDITIONAL
   : 'UNCONDITIONAL'
   ;

UNENCRYPTED
   : 'UNENCRYPTED'
   ;

UNION
   : 'UNION'
   ;

UNIQUE
   : 'UNIQUE'
   ;

UNKNOWN
   : 'UNKNOWN'
   ;

UNLISTEN
   : 'UNLISTEN'
   ;

UNLOGGED
   : 'UNLOGGED'
   ;

UNTIL
   : 'UNTIL'
   ;

UPDATE
   : 'UPDATE'
   ;

USER
   : 'USER'
   ;

USING
   : 'USING'
   ;

VACUUM
   : 'VACUUM'
   ;

VALID
   : 'VALID'
   ;

VALIDATE
   : 'VALIDATE'
   ;

VALIDATOR
   : 'VALIDATOR'
   ;

VALUE_P
   : 'VALUE'
   ;

VALUES
   : 'VALUES'
   ;

VARCHAR
   : 'VARCHAR'
   ;

VARIADIC
   : 'VARIADIC'
   ;

VARYING
   : 'VARYING'
   ;

VERBOSE
   : 'VERBOSE'
   ;

VERSION_P
   : 'VERSION'
   ;

VIEW
   : 'VIEW'
   ;

VIEWS
   : 'VIEWS'
   ;

VIRTUAL
   : 'VIRTUAL'
   ;

VOLATILE
   : 'VOLATILE'
   ;

WHEN
   : 'WHEN'
   ;

WHERE
   : 'WHERE'
   ;

WHITESPACE_P
   : 'WHITESPACE'
   ;

WINDOW
   : 'WINDOW'
   ;

WITH
   : 'WITH'
   ;

WITHIN
   : 'WITHIN'
   ;

WITHOUT
   : 'WITHOUT'
   ;

WORK
   : 'WORK'
   ;

WRAPPER
   : 'WRAPPER'
   ;

WRITE
   : 'WRITE'
   ;

XML_P
   : 'XML'
   ;

XMLATTRIBUTES
   : 'XMLATTRIBUTES'
   ;

XMLCONCAT
   : 'XMLCONCAT'
   ;

XMLELEMENT
   : 'XMLELEMENT'
   ;

XMLEXISTS
   : 'XMLEXISTS'
   ;

XMLFOREST
   : 'XMLFOREST'
   ;

XMLNAMESPACES
   : 'XMLNAMESPACES'
   ;

XMLPARSE
   : 'XMLPARSE'
   ;

XMLPI
   : 'XMLPI'
   ;

XMLROOT
   : 'XMLROOT'
   ;

XMLSERIALIZE
   : 'XMLSERIALIZE'
   ;

XMLTABLE
   : 'XMLTABLE'
   ;

YEAR_P
   : 'YEAR'
   ;

YES_P
   : 'YES'
   ;

ZONE
   : 'ZONE'
   ;

// ============================================================================
// Automatically Renamed Tokens (ANTLR Compatibility)
// ============================================================================
// The following tokens were renamed to avoid ANTLR reserved name conflicts:
//
//   SKIP → SKIP_P (keyword: 'skip')
// ============================================================================

// END AUTO-GENERATED KEYWORDS
// ============================================================================

//

// SPECIAL CHARACTERS (4.1.4)

//

// Note that Asterisk is a valid operator, but does not have the type Operator due to its syntactic use in locations

// that are not expressions.

Dollar
   : '$'
   ;

OPEN_PAREN
   : '('
   ;

CLOSE_PAREN
   : ')'
   ;

OPEN_BRACKET
   : '['
   ;

CLOSE_BRACKET
   : ']'
   ;

COMMA
   : ','
   ;

SEMI
   : ';'
   ;

COLON
   : ':'
   ;

STAR
   : '*'
   ;

EQUAL
   : '='
   ;

DOT
   : '.'
   ;
   //NamedArgument	: ':=';

PLUS
   : '+'
   ;

MINUS
   : '-'
   ;

SLASH
   : '/'
   ;

CARET
   : '^'
   ;

LT
   : '<'
   ;

GT
   : '>'
   ;

LESS_LESS
   : '<<'
   ;

GREATER_GREATER
   : '>>'
   ;

COLON_EQUALS
   : ':='
   ;

LESS_EQUALS
   : '<='
   ;

EQUALS_GREATER
   : '=>'
   ;

GREATER_EQUALS
   : '>='
   ;

DOT_DOT
   : '..'
   ;

NOT_EQUALS
   : '<>'
   ;

TYPECAST
   : '::'
   ;

PERCENT
   : '%'
   ;

PARAM
   : '$' ([0-9])+
   ;
   //

   // OPERATORS (4.1.3)

   //

   // this rule does not allow + or - at the end of a multi-character operator

Operator
   : ((OperatorCharacter | ('+' | '-'
   {p.checkLA('-')}?)+ (OperatorCharacter | '/'
   {p.checkLA('*')}?) | '/'
   {p.checkLA('*')}?)+ | // special handling for the single-character operators + and -
   [+-])
   //TODO somehow rewrite this part without using Actions

   {
    l.HandleLessLessGreaterGreater();
   }
   ;
/* This rule handles operators which end with + or -, and sets the token type to Operator. It is comprised of four
 * parts, in order:
 *
 *   1. A prefix, which does not contain a character from the required set which allows + or - to appear at the end of
 *      the operator.
 *   2. A character from the required set which allows + or - to appear at the end of the operator.
 *   3. An optional sub-token which takes the form of an operator which does not include a + or - at the end of the
 *      sub-token.
 *   4. A suffix sequence of + and - characters.
 */


OperatorEndingWithPlusMinus
   : (OperatorCharacterNotAllowPlusMinusAtEnd | '-'
   {p.checkLA('-')}? | '/'
   {p.checkLA('*')}?)* OperatorCharacterAllowPlusMinusAtEnd Operator? ('+' | '-'
   {p.checkLA('-')}?)+ -> type (Operator)
   ;
   // Each of the following fragment rules omits the +, -, and / characters, which must always be handled in a special way

   // by the operator rules above.

fragment OperatorCharacter
   : [*<>=~!@%^&|`?#]
   ;
   // these are the operator characters that don't count towards one ending with + or -

fragment OperatorCharacterNotAllowPlusMinusAtEnd
   : [*<>=+]
   ;
   // an operator may end with + or - if it contains one of these characters

fragment OperatorCharacterAllowPlusMinusAtEnd
   : [~!@%^&|`?#]
   ;
   //

   // KEYWORDS (Appendix C)

   //

   //

   // reserved keywords

   //

   //

   // reserved keywords (can be function or type)

   //

   //

   // non-reserved keywords

   //

   //LC_COLLATE			: 'LC'_'COLLATE;

   //LC_CTYPE			: 'LC'_'CTYPE;

RECHECK
   : 'RECHECK'
   ;

   //VALUE				: 'VALUE;

   //

   // non-reserved keywords (can not be function or type)

   //
   
   
XMLCOMMENT
   : 'XMLCOMMENT'
   ;

XMLAGG
   : 'XMLAGG'
   ;

XML_IS_WELL_FORMED
   : 'XML_IS_WELL_FORMED'
   ;

XML_IS_WELL_FORMED_DOCUMENT
   : 'XML_IS_WELL_FORMED_DOCUMENT'
   ;

XML_IS_WELL_FORMED_CONTENT
   : 'XML_IS_WELL_FORMED_CONTENT'
   ;

XPATH
   : 'XPATH'
   ;

XPATH_EXISTS
   : 'XPATH_EXISTS'
   ;

   //MISSED

ROWTYPE
   : 'ROWTYPE'
   ;

DUMP
   : 'DUMP'
   ;

PRINT_STRICT_PARAMS
   : 'PRINT_STRICT_PARAMS'
   ;

VARIABLE_CONFLICT
   : 'VARIABLE_CONFLICT'
   ;

USE_VARIABLE
   : 'USE_VARIABLE'
   ;

USE_COLUMN
   : 'USE_COLUMN'
   ;

ALIAS
   : 'ALIAS'
   ;

CONSTANT
   : 'CONSTANT'
   ;

PERFORM
   : 'PERFORM'
   ;

GET
   : 'GET'
   ;

DIAGNOSTICS
   : 'DIAGNOSTICS'
   ;

STACKED
   : 'STACKED'
   ;

ELSIF
   : 'ELSIF'
   ;

WHILE
   : 'WHILE'
   ;

REVERSE
   : 'REVERSE'
   ;

FOREACH
   : 'FOREACH'
   ;

SLICE
   : 'SLICE'
   ;

EXIT
   : 'EXIT'
   ;

QUERY
   : 'QUERY'
   ;

RAISE
   : 'RAISE'
   ;

SQLSTATE
   : 'SQLSTATE'
   ;

DEBUG
   : 'DEBUG'
   ;

LOG
   : 'LOG'
   ;

INFO
   : 'INFO'
   ;

NOTICE
   : 'NOTICE'
   ;

WARNING
   : 'WARNING'
   ;

EXCEPTION
   : 'EXCEPTION'
   ;

ASSERT
   : 'ASSERT'
   ;

LOOP
   : 'LOOP'
   ;

OPEN
   : 'OPEN'
   ;

Identifier
   : IdentifierStartChar IdentifierChar*
   ;

fragment IdentifierStartChar options { caseInsensitive=false; }
   : // these are the valid identifier start characters below 0x7F
   [a-zA-Z_]
   | // these are the valid characters from 0x80 to 0xFF
   [\u00AA\u00B5\u00BA\u00C0-\u00D6\u00D8-\u00F6\u00F8-\u00FF]
   | // these are the letters above 0xFF which only need a single UTF-16 code unit
   [\u0100-\uD7FF\uE000-\uFFFF]
   {p.charIsLetter()}?
   | // letters which require multiple UTF-16 code units
   [\uD800-\uDBFF] [\uDC00-\uDFFF]
   {    p.CheckIfUtf32Letter()   }?

   ;

fragment IdentifierChar
   : StrictIdentifierChar
   | '$'
   ;

fragment StrictIdentifierChar
   : IdentifierStartChar
   | [0-9]
   ;
/* Quoted Identifiers
 *
 *   These are divided into four separate tokens, allowing distinction of valid quoted identifiers from invalid quoted
 *   identifiers without sacrificing the ability of the lexer to reliably recover from lexical errors in the input.
 */


QuotedIdentifier
   : UnterminatedQuotedIdentifier '"'
   ;
   // This is a quoted identifier which only contains valid characters but is not terminated

UnterminatedQuotedIdentifier
   : '"' ('""' | ~ [\u0000"])*
   ;
   // This is a quoted identifier which is terminated but contains a \u0000 character

InvalidQuotedIdentifier
   : InvalidUnterminatedQuotedIdentifier '"'
   ;
   // This is a quoted identifier which is unterminated and contains a \u0000 character

InvalidUnterminatedQuotedIdentifier
   : '"' ('""' | ~ '"')*
   ;
/* Unicode Quoted Identifiers
 *
 *   These are divided into four separate tokens, allowing distinction of valid Unicode quoted identifiers from invalid
 *   Unicode quoted identifiers without sacrificing the ability of the lexer to reliably recover from lexical errors in
 *   the input. Note that escape sequences are never checked as part of this determination due to the ability of users
 *   to change the escape character with a UESCAPE clause following the Unicode quoted identifier.
 *
 * TODO: these rules assume "" is still a valid escape sequence within a Unicode quoted identifier.
 */


UnicodeQuotedIdentifier
   : 'U' '&' QuotedIdentifier
   ;
   // This is a Unicode quoted identifier which only contains valid characters but is not terminated

UnterminatedUnicodeQuotedIdentifier
   : 'U' '&' UnterminatedQuotedIdentifier
   ;
   // This is a Unicode quoted identifier which is terminated but contains a \u0000 character

InvalidUnicodeQuotedIdentifier
   : 'U' '&' InvalidQuotedIdentifier
   ;
   // This is a Unicode quoted identifier which is unterminated and contains a \u0000 character

InvalidUnterminatedUnicodeQuotedIdentifier
   : 'U' '&' InvalidUnterminatedQuotedIdentifier
   ;
   //

   // CONSTANTS (4.1.2)

   //

   // String Constants (4.1.2.1)

StringConstant
   : UnterminatedStringConstant '\''
   ;

UnterminatedStringConstant
   : '\'' ('\'\'' | ~ '\'')*
   ;
   // String Constants with C-style Escapes (4.1.2.2)

BeginEscapeStringConstant
   : 'E' '\'' -> more , pushMode (EscapeStringConstantMode)
   ;
   // String Constants with Unicode Escapes (4.1.2.3)

   //

   //   Note that escape sequences are never checked as part of this token due to the ability of users to change the escape

   //   character with a UESCAPE clause following the Unicode string constant.

   //

   // TODO: these rules assume '' is still a valid escape sequence within a Unicode string constant.

UnicodeEscapeStringConstant
   : UnterminatedUnicodeEscapeStringConstant '\''
   ;

UnterminatedUnicodeEscapeStringConstant
   : 'U' '&' UnterminatedStringConstant
   ;
   // Dollar-quoted String Constants (4.1.2.4)

BeginDollarStringConstant
   : '$' Tag? '$'
   {l.pushTag();} -> pushMode (DollarQuotedStringMode)
   ;
/* "The tag, if any, of a dollar-quoted string follows the same rules as an
 * unquoted identifier, except that it cannot contain a dollar sign."
 */


fragment Tag
   : IdentifierStartChar StrictIdentifierChar*
   ;
   // Bit-strings Constants (4.1.2.5)

BinaryStringConstant
   : UnterminatedBinaryStringConstant '\''
   ;

UnterminatedBinaryStringConstant
   : 'B' '\'' [01]*
   ;

InvalidBinaryStringConstant
   : InvalidUnterminatedBinaryStringConstant '\''
   ;

InvalidUnterminatedBinaryStringConstant
   : 'B' UnterminatedStringConstant
   ;

HexadecimalStringConstant
   : UnterminatedHexadecimalStringConstant '\''
   ;

UnterminatedHexadecimalStringConstant
   : 'X' '\'' [0-9A-F]*
   ;

InvalidHexadecimalStringConstant
   : InvalidUnterminatedHexadecimalStringConstant '\''
   ;

InvalidUnterminatedHexadecimalStringConstant
   : 'X' UnterminatedStringConstant
   ;
   // Numeric Constants (4.1.2.6)

Integral
   : Digits
   ;

NumericFail
   : Digits '..'
   {l.HandleNumericFail();}
   ;

Numeric
   : Digits '.' Digits? /*? replaced with + to solve problem with DOT_DOT .. but this surely must be rewriten */

   ('E' [+-]? Digits)?
   | '.' Digits ('E' [+-]? Digits)?
   | Digits 'E' [+-]? Digits
   ;

fragment Digits
   : [0-9]+
   ;

PLSQLVARIABLENAME
   : ':' [A-Z_] [A-Z_0-9$]*
   ;

PLSQLIDENTIFIER
   : ':"' ('\\' . | '""' | ~ ('"' | '\\'))* '"'
   ;
   //

   // WHITESPACE (4.1)

   //

Whitespace
   : [ \t] -> channel (HIDDEN)
   ;

Newline
   : ('\r' '\n'? | '\n') -> channel (HIDDEN)
   ;
   //

   // COMMENTS (4.1.5)

   //

LineComment
   : '--' ~ [\r\n]* -> channel (HIDDEN)
   ;

BlockComment
   : ('/*' ('/'* BlockComment | ~ [/*] | '/'+ ~ [/*] | '*'+ ~ [/*])* '*'* '*/') -> channel (HIDDEN)
   ;

UnterminatedBlockComment
   : '/*' ('/'* BlockComment | // these characters are not part of special sequences in a block comment
   ~ [/*] | // handle / or * characters which are not part of /* or */ and do not appear at the end of the file
   ('/'+ ~ [/*] | '*'+ ~ [/*]))*
   // Handle the case of / or * characters at the end of the file, or a nested unterminated block comment
   ('/'+ | '*'+ | '/'* UnterminatedBlockComment)?
   // Optional assertion to make sure this rule is working as intended

   {
            l.UnterminatedBlockCommentDebugAssert();
   }
   ;
   //

   // META-COMMANDS

   //

   // http://www.postgresql.org/docs/9.3/static/app-psql.html

MetaCommand
   : '\\' (~ [\r\n\\"] | '"' ~ [\r\n"]* '"')* ('"' ~ [\r\n"]*)?
   ;

EndMetaCommand
   : '\\\\'
   ;
   //

   // ERROR

   //

   // Any character which does not match one of the above rules will appear in the token stream as an ErrorCharacter token.

   // This ensures the lexer itself will never encounter a syntax error, so all error handling may be performed by the

   // parser.

ErrorCharacter
   : .
   ;

mode EscapeStringConstantMode;
EscapeStringConstant
   : EscapeStringText '\'' -> mode (AfterEscapeStringConstantMode)
   ;

UnterminatedEscapeStringConstant
   : EscapeStringText
   // Handle a final unmatched \ character appearing at the end of the file
   '\\'? EOF
   ;

fragment EscapeStringText options { caseInsensitive=false; }
   : ('\'\'' | '\\' ( // two-digit hex escapes are still valid when treated as single-digit escapes
   'x' [0-9a-fA-F] |
   'u' [0-9a-fA-F] [0-9a-fA-F] [0-9a-fA-F] [0-9a-fA-F] |
   'U' [0-9a-fA-F] [0-9a-fA-F] [0-9a-fA-F] [0-9a-fA-F] [0-9a-fA-F] [0-9a-fA-F] [0-9a-fA-F] [0-9a-fA-F] | // Any character other than the Unicode escapes can follow a backslash. Some have special meaning,
   // but that doesn't affect the syntax.
   ~ [xuU]) | ~ ['\\])*
   ;

InvalidEscapeStringConstant
   : InvalidEscapeStringText '\'' -> mode (AfterEscapeStringConstantMode)
   ;

InvalidUnterminatedEscapeStringConstant
   : InvalidEscapeStringText
   // Handle a final unmatched \ character appearing at the end of the file
   '\\'? EOF
   ;

fragment InvalidEscapeStringText
   : ('\'\'' | '\\' . | ~ ['\\])*
   ;

mode AfterEscapeStringConstantMode;
AfterEscapeStringConstantMode_Whitespace
   : Whitespace -> type (Whitespace) , channel (HIDDEN)
   ;

AfterEscapeStringConstantMode_Newline
   : Newline -> type (Newline) , channel (HIDDEN) , mode (AfterEscapeStringConstantWithNewlineMode)
   ;

AfterEscapeStringConstantMode_NotContinued
   :
   {} // intentionally empty
   -> skip , popMode
   ;

mode AfterEscapeStringConstantWithNewlineMode;
AfterEscapeStringConstantWithNewlineMode_Whitespace
   : Whitespace -> type (Whitespace) , channel (HIDDEN)
   ;

AfterEscapeStringConstantWithNewlineMode_Newline
   : Newline -> type (Newline) , channel (HIDDEN)
   ;

AfterEscapeStringConstantWithNewlineMode_Continued
   : '\'' -> more , mode (EscapeStringConstantMode)
   ;

AfterEscapeStringConstantWithNewlineMode_NotContinued
   :
   {} // intentionally empty
   -> skip , popMode
   ;

mode DollarQuotedStringMode;
DollarText
   : ~ '$'+
   //| '$'([0-9])+
   | // this alternative improves the efficiency of handling $ characters within a dollar-quoted string which are

   // not part of the ending tag.
   '$' ~ '$'*
   ;

EndDollarStringConstant
   : ('$' Tag? '$')
   {p.isTag()}?
   {l.popTag();} -> popMode
   ;
