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
lexer grammar RedshiftLexer;
/* Reference:
 * http://www.postgresql.org/docs/9.3/static/sql-syntax-lexical.html
 */

options {
   superClass = RedshiftLexerBase;
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

ALL
   : 'ALL'
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

ASYMMETRIC
   : 'ASYMMETRIC'
   ;

BOTH
   : 'BOTH'
   ;

CASE
   : 'CASE'
   ;

CAST
   : 'CAST'
   ;

CHECK
   : 'CHECK'
   ;

COLLATE
   : 'COLLATE'
   ;

COLUMN
   : 'COLUMN'
   ;

CONSTRAINT
   : 'CONSTRAINT'
   ;

CREATE
   : 'CREATE'
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

CURRENT_TIME
   : 'CURRENT_TIME'
   ;

CURRENT_TIMESTAMP
   : 'CURRENT_TIMESTAMP'
   ;

CURRENT_USER
   : 'CURRENT_USER'
   ;

DEFAULT
   : 'DEFAULT'
   ;

DEFERRABLE
   : 'DEFERRABLE'
   ;

DESC
   : 'DESC'
   ;

DISTINCT
   : 'DISTINCT'
   ;

DO
   : 'DO'
   ;

ELSE
   : 'ELSE'
   ;

EXCEPT
   : 'EXCEPT'
   ;

FALSE_P
   : 'FALSE'
   ;

FETCH
   : 'FETCH'
   ;

FOR
   : 'FOR'
   ;

FOREIGN
   : 'FOREIGN'
   ;

FROM
   : 'FROM'
   ;

GRANT
   : 'GRANT'
   ;

GROUP_P
   : 'GROUP'
   ;

HAVING
   : 'HAVING'
   ;

IN_P
   : 'IN'
   ;

INITIALLY
   : 'INITIALLY'
   ;

INTERSECT
   : 'INTERSECT'
   ;

INTO
   : 'INTO'
   ;

LATERAL_P
   : 'LATERAL'
   ;

LEADING
   : 'LEADING'
   ;

LIMIT
   : 'LIMIT'
   ;

LOCALTIME
   : 'LOCALTIME'
   ;

LOCALTIMESTAMP
   : 'LOCALTIMESTAMP'
   ;

NOT
   : 'NOT'
   ;

NULL_P
   : 'NULL'
   ;

OFFSET
   : 'OFFSET'
   ;

ON
   : 'ON'
   ;

ONLY
   : 'ONLY'
   ;

OR
   : 'OR'
   ;

ORDER
   : 'ORDER'
   ;

PLACING
   : 'PLACING'
   ;

PRIMARY
   : 'PRIMARY'
   ;

PUBLIC
   : 'PUBLIC'
   ;

REFERENCES
   : 'REFERENCES'
   ;

RETURNING
   : 'RETURNING'
   ;

SELECT
   : 'SELECT'
   ;

SESSION_USER
   : 'SESSION_USER'
   ;

SOME
   : 'SOME'
   ;

SYMMETRIC
   : 'SYMMETRIC'
   ;

TABLE
   : 'TABLE'
   ;

THEN
   : 'THEN'
   ;

TO
   : 'TO'
   ;

TRAILING
   : 'TRAILING'
   ;

TRUE_P
   : 'TRUE'
   ;

UNION
   : 'UNION'
   ;

UNIQUE
   : 'UNIQUE'
   ;

USER
   : 'USER'
   ;

USING
   : 'USING'
   ;

VARIADIC
   : 'VARIADIC'
   ;

WHEN
   : 'WHEN'
   ;

WHERE
   : 'WHERE'
   ;

WINDOW
   : 'WINDOW'
   ;

WITH
   : 'WITH'
   ;

   //

   // reserved keywords (can be function or type)

   //

AUTHORIZATION
   : 'AUTHORIZATION'
   ;

BINARY
   : 'BINARY'
   ;

BINDING
   : 'BINDING'
   ;

COLLATION
   : 'COLLATION'
   ;

CONCURRENTLY
   : 'CONCURRENTLY'
   ;

CROSS
   : 'CROSS'
   ;

CURRENT_SCHEMA
   : 'CURRENT_SCHEMA'
   ;

FREEZE
   : 'FREEZE'
   ;

FULL
   : 'FULL'
   ;

ILIKE
   : 'ILIKE'
   ;

INNER_P
   : 'INNER'
   ;

IS
   : 'IS'
   ;

ISNULL
   : 'ISNULL'
   ;

JOIN
   : 'JOIN'
   ;

LEFT
   : 'LEFT'
   ;

LIKE
   : 'LIKE'
   ;

NATURAL
   : 'NATURAL'
   ;

NOTNULL
   : 'NOTNULL'
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

RIGHT
   : 'RIGHT'
   ;

SIMILAR
   : 'SIMILAR'
   ;

VERBOSE
   : 'VERBOSE'
   ;
   //

   // non-reserved keywords

   //

ABORT_P
   : 'ABORT'
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

ALSO
   : 'ALSO'
   ;

ALTER
   : 'ALTER'
   ;

ALWAYS
   : 'ALWAYS'
   ;

ASSERTION
   : 'ASSERTION'
   ;

ASSIGNMENT
   : 'ASSIGNMENT'
   ;

AT
   : 'AT'
   ;

ATTRIBUTE
   : 'ATTRIBUTE'
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

BY
   : 'BY'
   ;

CACHE
   : 'CACHE'
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

CATALOG
   : 'CATALOG'
   ;

CHAIN
   : 'CHAIN'
   ;

CHARACTERISTICS
   : 'CHARACTERISTICS'
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

CONFIGURATION
   : 'CONFIGURATION'
   ;

CONNECTION
   : 'CONNECTION'
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

CSV
   : 'CSV'
   ;

JSON
   : 'JSON'
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

DATA_CATALOG
   : 'DATA_CATALOG'
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

DECLARE
   : 'DECLARE'
   ;

DEFAULTS
   : 'DEFAULTS'
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

DICTIONARY
   : 'DICTIONARY'
   ;

DISABLE_P
   : 'DISABLE'
   ;

DISCARD
   : 'DISCARD'
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

ENABLE_P
   : 'ENABLE'
   ;

ENCODING
   : 'ENCODING'
   ;

ENCRYPTED
   : 'ENCRYPTED'
   ;

ENUM_P
   : 'ENUM'
   ;

ESCAPE
   : 'ESCAPE'
   ;

EVENT
   : 'EVENT'
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

EXPLAIN
   : 'EXPLAIN'
   ;

EXTENSION
   : 'EXTENSION'
   ;

EXTERNAL
   : 'EXTERNAL'
   ;

FAMILY
   : 'FAMILY'
   ;

FIRST_P
   : 'FIRST'
   ;

FOLLOWING
   : 'FOLLOWING'
   ;

FORCE
   : 'FORCE'
   ;

FORWARD
   : 'FORWARD'
   ;

FUNCTION
   : 'FUNCTION'
   ;

FUNCTIONS
   : 'FUNCTIONS'
   ;

GLOBAL
   : 'GLOBAL'
   ;

GRANTED
   : 'GRANTED'
   ;

HANDLER
   : 'HANDLER'
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

IMMEDIATE
   : 'IMMEDIATE'
   ;

IMMUTABLE
   : 'IMMUTABLE'
   ;

IMPLICIT_P
   : 'IMPLICIT'
   ;

INCLUDING
   : 'INCLUDING'
   ;

INCREMENT
   : 'INCREMENT'
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

INLINE_P
   : 'INLINE'
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

INVOKER
   : 'INVOKER'
   ;

ISOLATION
   : 'ISOLATION'
   ;

KEY
   : 'KEY'
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
   //LC_COLLATE			: 'LC'_'COLLATE;

   //LC_CTYPE			: 'LC'_'CTYPE;

LEAKPROOF
   : 'LEAKPROOF'
   ;

LEVEL
   : 'LEVEL'
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

LOCATION
   : 'LOCATION'
   ;

LOCK_P
   : 'LOCK'
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

NEXT
   : 'NEXT'
   ;

NO
   : 'NO'
   ;

NOTHING
   : 'NOTHING'
   ;

NOTIFY
   : 'NOTIFY'
   ;

NOWAIT
   : 'NOWAIT'
   ;

NULLS_P
   : 'NULLS'
   ;

OBJECT_P
   : 'OBJECT'
   ;

OF
   : 'OF'
   ;

OFF
   : 'OFF'
   ;

OIDS
   : 'OIDS'
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

OWNED
   : 'OWNED'
   ;

OWNER
   : 'OWNER'
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

PLANS
   : 'PLANS'
   ;

PRECEDING
   : 'PRECEDING'
   ;

PREDICATE
   : 'PREDICATE'
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

PROGRAM
   : 'PROGRAM'
   ;

QUOTE
   : 'QUOTE'
   ;

RANGE
   : 'RANGE'
   ;

READ
   : 'READ'
   ;

REASSIGN
   : 'REASSIGN'
   ;

RECHECK
   : 'RECHECK'
   ;

RECURSIVE
   : 'RECURSIVE'
   ;

REF
   : 'REF'
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

RETURNS
   : 'RETURNS'
   ;

REVOKE
   : 'REVOKE'
   ;

ROLE
   : 'ROLE'
   ;

ROLLBACK
   : 'ROLLBACK'
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

SCHEMA
   : 'SCHEMA'
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

SET
   : 'SET'
   ;

SHARE
   : 'SHARE'
   ;

SHOW
   : 'SHOW'
   ;

SIMPLE
   : 'SIMPLE'
   ;

SNAPSHOT
   : 'SNAPSHOT'
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

STRICT_P
   : 'STRICT'
   ;

STRIP_P
   : 'STRIP'
   ;

SYSID
   : 'SYSID'
   ;

SYSTEM_P
   : 'SYSTEM'
   ;

TABLES
   : 'TABLES'
   ;

TABLESPACE
   : 'TABLESPACE'
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

TRANSACTION
   : 'TRANSACTION'
   ;

TRIGGER
   : 'TRIGGER'
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

UNBOUNDED
   : 'UNBOUNDED'
   ;

UNCOMMITTED
   : 'UNCOMMITTED'
   ;

UNENCRYPTED
   : 'UNENCRYPTED'
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
   //VALUE				: 'VALUE;

VARYING
   : 'VARYING'
   ;

VERSION_P
   : 'VERSION'
   ;

VIEW
   : 'VIEW'
   ;

VOLATILE
   : 'VOLATILE'
   ;

WHITESPACE_P
   : 'WHITESPACE'
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

YEAR_P
   : 'YEAR'
   ;

YES_P
   : 'YES'
   ;

ZONE
   : 'ZONE'
   ;

// REDSHIFT-SPECIFIC KEYWORDS
// These tokens are specific to Amazon Redshift and not part of standard PostgreSQL

// Datashare Commands
QUALIFY
   : 'QUALIFY'
   ;

CONNECT
   : 'CONNECT'
   ;

TOP
   : 'TOP'
   ;

VARBYTE
   : 'VARBYTE'
   ;

VARBINARY
   : 'VARBINARY'
   ;

CONJUNCTION
   : 'CONJUNCTION'
   ;

DEFINITION
   : 'DEFINITION'
   ;

DATASHARE
   : 'DATASHARE'
   ;

FILE
   : 'FILE'
   ;

PUBLICACCESSIBLE
   : 'PUBLICACCESSIBLE'
   ;

INCLUDENEW
   : 'INCLUDENEW'
   ;

// External Data Sources
IAM_ROLE
   : 'IAM_ROLE'
   ;

CATALOG_ROLE
   : 'CATALOG_ROLE'
   ;

CATALOG_ID
   : 'CATALOG_ID'
   ;

HIVE
   : 'HIVE'
   ;

METASTORE
   : 'METASTORE'
   ;

URI
   : 'URI'
   ;

POSTGRES
   : 'POSTGRES'
   ;

MYSQL
   : 'MYSQL'
   ;

SECRET_ARN
   : 'SECRET_ARN'
   ;

KINESIS
   : 'KINESIS'
   ;

KAFKA
   : 'KAFKA'
   ;

MSK
   : 'MSK'
   ;

AUTHENTICATION
   : 'AUTHENTICATION'
   ;

AUTHENTICATION_ARN
   : 'AUTHENTICATION_ARN'
   ;

SESSION_TOKEN
   : 'SESSION'
   ;

MTLS
   : 'MTLS'
   ;

// Security/Policy
MASKING
   : 'MASKING'
   ;

RLS
   : 'RLS'
   ;

IDENTITY
   : 'IDENTITY'
   ;

PROVIDER
   : 'PROVIDER'
   ;

PROTECTED
   : 'PROTECTED'
   ;

// ML/Model Commands
MODEL
   : 'MODEL'
   ;

TARGET
   : 'TARGET'
   ;

SAGEMAKER
   : 'SAGEMAKER'
   ;

AUTO
   : 'AUTO'
   ;

MODEL_TYPE
   : 'MODEL_TYPE'
   ;

PROBLEM_TYPE
   : 'PROBLEM_TYPE'
   ;

OBJECTIVE
   : 'OBJECTIVE'
   ;

PREPROCESSORS
   : 'PREPROCESSORS'
   ;

HYPERPARAMETERS
   : 'HYPERPARAMETERS'
   ;

XGBOOST
   : 'XGBOOST'
   ;

MLP
   : 'MLP'
   ;

LINEAR_LEARNER
   : 'LINEAR_LEARNER'
   ;

KMEANS
   : 'KMEANS'
   ;

FORECAST
   : 'FORECAST'
   ;

REGRESSION
   : 'REGRESSION'
   ;

BINARY_CLASSIFICATION
   : 'BINARY_CLASSIFICATION'
   ;

MULTICLASS_CLASSIFICATION
   : 'MULTICLASS_CLASSIFICATION'
   ;

S3_BUCKET
   : 'S3_BUCKET'
   ;

TAGS
   : 'TAGS'
   ;

KMS_KEY_ID
   : 'KMS_KEY_ID'
   ;

S3_GARBAGE_COLLECT
   : 'S3_GARBAGE_COLLECT'
   ;

MAX_CELLS
   : 'MAX_CELLS'
   ;

MAX_RUNTIME
   : 'MAX_RUNTIME'
   ;

HORIZON
   : 'HORIZON'
   ;

FREQUENCY
   : 'FREQUENCY'
   ;

PERCENTILES
   : 'PERCENTILES'
   ;

MAX_BATCH_ROWS
   : 'MAX_BATCH_ROWS'
   ;

// UNLOAD Command
UNLOAD
   : 'UNLOAD'
   ;

MANIFEST
   : 'MANIFEST'
   ;

ADDQUOTES
   : 'ADDQUOTES'
   ;

ALLOWOVERWRITE
   : 'ALLOWOVERWRITE'
   ;

CLEANPATH
   : 'CLEANPATH'
   ;

MAXFILESIZE
   : 'MAXFILESIZE'
   ;

ROWGROUPSIZE
   : 'ROWGROUPSIZE'
   ;

BZIP2
   : 'BZIP2'
   ;

GZIP
   : 'GZIP'
   ;

ZSTD
   : 'ZSTD'
   ;

// Show Commands
DATABASES
   : 'DATABASES'
   ;

DATASHARES
   : 'DATASHARES'
   ;

GRANTS
   : 'GRANTS'
   ;

// Session/Connection
USE
   : 'USE'
   ;

CANCEL
   : 'CANCEL'
   ;

SESSION_AUTHORIZATION
   : 'SESSION_AUTHORIZATION'
   ;

SESSION_CHARACTERISTICS
   : 'SESSION_CHARACTERISTICS'
   ;

// General
COMPRESSION
   : 'COMPRESSION'
   ;

LIBRARY
   : 'LIBRARY'
   ;

APPEND
   : 'APPEND'
   ;

// Size Units
MB
   : 'MB'
   ;

GB
   : 'GB'
   ;

// Common Redshift Keywords
ACCOUNT
   : 'ACCOUNT'
   ;

NAMESPACE
   : 'NAMESPACE'
   ;

DESCRIBE
   : 'DESCRIBE'
   ;

// Additional Redshift-specific tokens
NONATOMIC
    : 'NONATOMIC'
    ;

MANAGEDBY
   : 'MANAGEDBY'
   ;

ADX
   : 'ADX'
   ;

REMOVE
   : 'REMOVE'
   ;

DUPLICATES
   : 'DUPLICATES'
   ;

BEDROCK
   : 'BEDROCK'
   ;

MODEL_ID
   : 'MODEL_ID'
   ;

PROMPT
   : 'PROMPT'
   ;

SUFFIX
   : 'SUFFIX'
   ;

REQUEST_TYPE
   : 'REQUEST_TYPE'
   ;

RESPONSE_TYPE
   : 'RESPONSE_TYPE'
   ;

RAW
   : 'RAW'
   ;

UNIFIED
   : 'UNIFIED'
   ;

SUPER
   : 'SUPER'
   ;

CI
   : 'CI'
   ;

CS
   : 'CS'
   ;

PLPYTHONU
   : 'PLPYTHONU'
   ;

FILLTARGET
   : 'FILLTARGET'
   ;

IGNOREEXTRA
   : 'IGNOREEXTRA'
   ;

CREATEUSER
   : 'CREATEUSER'
   ;

NOCREATEUSER
   : 'NOCREATEUSER'
   ;

// Additional commonly used tokens (avoiding conflicts with existing _P tokens and identifiers)
REGION
   : 'REGION'
   ;

PORT
   : 'PORT'
   ;

REDSHIFT
   : 'REDSHIFT'
   ;

IAM
   : 'IAM'
   ;

CREATEDB
   : 'CREATEDB'
   ;

NOCREATEDB
   : 'NOCREATEDB'
   ;

RESTRICTED
   : 'RESTRICTED'
   ;

UNLIMITED
   : 'UNLIMITED'
   ;

EXTERNALID
   : 'EXTERNALID'
   ;

TIMEOUT
   : 'TIMEOUT'
   ;

SYSLOG
   : 'SYSLOG'
   ;

CREDENTIALS
   : 'CREDENTIALS'
   ;

UNRESTRICTED
   : 'UNRESTRICTED'
   ;

PARAMETERS
   : 'PARAMETERS'
   ;

APPLICATION_ARN
   : 'APPLICATION_ARN'
   ;

AUTO_CREATE_ROLES
   : 'AUTO_CREATE_ROLES'
   ;

COMPROWS
   : 'COMPROWS'
   ;

PROVIDER_URL
   : 'PROVIDER_URL'
   ;

PROVIDER_URL_PORT
   : 'PROVIDER_URL_PORT'
   ;

ATTRIBUTE_MAP
   : 'ATTRIBUTE_MAP'
   ;

PROVIDER_ARN
   : 'PROVIDER_ARN'
   ;

ASSUME_ROLE_ARN
   : 'ASSUME_ROLE_ARN'
   ;


PROPERTIES
   : 'PROPERTIES'
   ;

AVRO
   : 'AVRO'
   ;

RCFILE
   : 'RCFILE'
   ;

SEQUENCEFILE
   : 'SEQUENCEFILE'
   ;

TEXTFILE
   : 'TEXTFILE'
   ;

ORC
   : 'ORC'
   ;

ION
   : 'ION'
   ;

LAMBDA
   : 'LAMBDA'
   ;

FIXEDWIDTH
   : 'FIXEDWIDTH'
   ;

// Missing tokens causing implicit definition warnings (avoiding conflicts with _P versions)
PARQUET
   : 'PARQUET'
   ;

LZOP
   : 'LZOP'
   ;

REMOVEQUOTES
   : 'REMOVEQUOTES'
   ;

TRUNCATECOLUMNS
   : 'TRUNCATECOLUMNS'
   ;

FILLRECORD
   : 'FILLRECORD'
   ;

BLANKSASNULL
   : 'BLANKSASNULL'
   ;

EMPTYASNULL
   : 'EMPTYASNULL'
   ;

MAXERROR
   : 'MAXERROR'
   ;

DATEFORMAT
   : 'DATEFORMAT'
   ;

TIMEFORMAT
   : 'TIMEFORMAT'
   ;

ACCEPTINVCHARS
   : 'ACCEPTINVCHARS'
   ;

ACCEPTANYDATE
   : 'ACCEPTANYDATE'
   ;

IGNOREHEADER
   : 'IGNOREHEADER'
   ;

IGNOREBLANKLINES
   : 'IGNOREBLANKLINES'
   ;

COMPUPDATE
   : 'COMPUPDATE'
   ;

STATUPDATE
   : 'STATUPDATE'
   ;

EXPLICIT_IDS
   : 'EXPLICIT_IDS'
   ;

READRATIO
   : 'READRATIO'
   ;

ROUNDEC
   : 'ROUNDEC'
   ;

TRIMBLANKS
   : 'TRIMBLANKS'
   ;

PRESET
   : 'PRESET'
   ;

ACCESS_KEY_ID
   : 'ACCESS_KEY_ID'
   ;

SECRET_ACCESS_KEY
   : 'SECRET_ACCESS_KEY'
   ;

SESSION_TOKEN_KW
   : 'SESSION_TOKEN'
   ;

HEADER
   : 'HEADER'
   ;

SETTINGS
   : 'SETTINGS'
   ;

FUNCTION_NAME
   : 'FUNCTION_NAME'
   ;

   //

   // non-reserved keywords (can not be function or type)

   //
   
ATOMIC_P
   : 'ATOMIC'
   ;

BETWEEN
   : 'BETWEEN'
   ;

BIGINT
   : 'BIGINT'
   ;

BIT
   : 'BIT'
   ;

BOOLEAN_P
   : 'BOOLEAN'
   ;

CHAR_P
   : 'CHAR'
   ;

CHARACTER
   : 'CHARACTER'
   ;

COALESCE
   : 'COALESCE'
   ;

DEC
   : 'DEC'
   ;

DECIMAL_P
   : 'DECIMAL'
   ;

EXISTS
   : 'EXISTS'
   ;

EXTRACT
   : 'EXTRACT'
   ;

FLOAT_P
   : 'FLOAT'
   ;

GREATEST
   : 'GREATEST'
   ;

INOUT
   : 'INOUT'
   ;

INT_P
   : 'INT'
   ;

INTEGER
   : 'INTEGER'
   ;

INTERVAL
   : 'INTERVAL'
   ;

LEAST
   : 'LEAST'
   ;

NATIONAL
   : 'NATIONAL'
   ;

NCHAR
   : 'NCHAR'
   ;

NONE
   : 'NONE'
   ;

NULLIF
   : 'NULLIF'
   ;

NUMERIC
   : 'NUMERIC'
   ;

OVERLAY
   : 'OVERLAY'
   ;
   
PARAMETER
   : 'PARAMETER'
   ;

POSITION
   : 'POSITION'
   ;

PRECISION
   : 'PRECISION'
   ;

REAL
   : 'REAL'
   ;

ROW
   : 'ROW'
   ;

SETOF
   : 'SETOF'
   ;

SMALLINT
   : 'SMALLINT'
   ;

SUBSTRING
   : 'SUBSTRING'
   ;

TIME
   : 'TIME'
   ;

TIMESTAMP
   : 'TIMESTAMP'
   ;

TREAT
   : 'TREAT'
   ;

TRIM
   : 'TRIM'
   ;

VALUES
   : 'VALUES'
   ;

VARCHAR
   : 'VARCHAR'
   ;

XMLATTRIBUTES
   : 'XMLATTRIBUTES'
   ;

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
   //MISSED

CALL
   : 'CALL'
   ;

CURRENT_P
   : 'CURRENT'
   ;

ATTACH
   : 'ATTACH'
   ;

DETACH
   : 'DETACH'
   ;

EXPRESSION
   : 'EXPRESSION'
   ;

GENERATED
   : 'GENERATED'
   ;

LOGGED
   : 'LOGGED'
   ;

STORED
   : 'STORED'
   ;

// External table format tokens
SERDE
   : 'SERDE'
   ;

SERDEPROPERTIES
   : 'SERDEPROPERTIES'
   ;

INPUTFORMAT
   : 'INPUTFORMAT'
   ;

OUTPUTFORMAT
   : 'OUTPUTFORMAT'
   ;

FIELDS
   : 'FIELDS'
   ;

COLLECTION
   : 'COLLECTION'
   ;

ITEMS
   : 'ITEMS'
   ;

TERMINATED
   : 'TERMINATED'
   ;

ESCAPED
   : 'ESCAPED'
   ;

DEFINED
   : 'DEFINED'
   ;

LINES
   : 'LINES'
   ;

KEYS
   : 'KEYS'
   ;

PARTITIONED
   : 'PARTITIONED'
   ;

// Complex type tokens
STRUCT
   : 'STRUCT'
   ;

MAP
   : 'MAP'
   ;

STRING
   : 'STRING'
   ;

DELIMITED
   : 'DELIMITED'
   ;

// Privilges and roles tokens
USAGE
   : 'USAGE'
   ;

IGNORE
    : 'IGNORE'
    ;

LANGUAGES
   : 'LANGUAGES'
   ;

JOB
   : 'JOB'
   ;

JOBS
    : 'JOBS'
    ;

VIA
   : 'VIA'
   ;

ASSUMEROLE
    : 'ASSUMEROLE'
    ;

// External function tokens
RETRY_TIMEOUT
   : 'RETRY_TIMEOUT'
   ;

MAX_BATCH_SIZE
   : 'MAX_BATCH_SIZE'
   ;

MAX_PAYLOAD_IN_MB
   : 'MAX_PAYLOAD_IN_MB'
   ;

KB
   : 'KB'
   ;


INCLUDE
   : 'INCLUDE'
   ;

ROUTINE
   : 'ROUTINE'
   ;

TRANSFORM
   : 'TRANSFORM'
   ;

IMPORT_P
   : 'IMPORT'
   ;

POLICY
   : 'POLICY'
   ;

PRIORITY
   : 'PRIORITY'
   ;

METHOD
   : 'METHOD'
   ;

REFERENCING
   : 'REFERENCING'
   ;

NEW
   : 'NEW'
   ;

OLD
   : 'OLD'
   ;

VALUE_P
   : 'VALUE'
   ;

SUBSCRIPTION
   : 'SUBSCRIPTION'
   ;

PUBLICATION
   : 'PUBLICATION'
   ;

OUT_P
   : 'OUT'
   ;

END_P
   : 'END'
   ;

ROUTINES
   : 'ROUTINES'
   ;

SCHEMAS
   : 'SCHEMAS'
   ;

PROCEDURES
   : 'PROCEDURES'
   ;

INPUT_P
   : 'INPUT'
   ;

SUPPORT
   : 'SUPPORT'
   ;

PARALLEL
   : 'PARALLEL'
   ;

SQL_P
   : 'SQL'
   ;

DEPENDS
   : 'DEPENDS'
   ;

OVERRIDING
   : 'OVERRIDING'
   ;

CONFLICT
   : 'CONFLICT'
   ;

SKIP_P
   : 'SKIP'
   ;

LOCKED
   : 'LOCKED'
   ;

TIES
   : 'TIES'
   ;

ROLLUP
   : 'ROLLUP'
   ;

CUBE
   : 'CUBE'
   ;

GROUPING
   : 'GROUPING'
   ;

SETS
   : 'SETS'
   ;

TABLESAMPLE
   : 'TABLESAMPLE'
   ;

ORDINALITY
   : 'ORDINALITY'
   ;

XMLTABLE
   : 'XMLTABLE'
   ;

COLUMNS
   : 'COLUMNS'
   ;

XMLNAMESPACES
   : 'XMLNAMESPACES'
   ;

ROWTYPE
   : 'ROWTYPE'
   ;

NORMALIZED
   : 'NORMALIZED'
   ;

WITHIN
   : 'WITHIN'
   ;

FILTER
   : 'FILTER'
   ;

GROUPS
   : 'GROUPS'
   ;

OTHERS
   : 'OTHERS'
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

UESCAPE
   : 'UESCAPE'
   ;

VIEWS
   : 'VIEWS'
   ;

NORMALIZE
   : 'NORMALIZE'
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

ERROR
   : 'ERROR'
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

RETURN
   : 'RETURN'
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
   //

   // IDENTIFIERS (4.1.1)

   //

ABS
   : 'ABS'
   ;

CBRT
   : 'CBRT'
   ;

CEIL
   : 'CEIL'
   ;

CEILING
   : 'CEILING'
   ;

DEGREES
   : 'DEGREES'
   ;

DIV
   : 'DIV'
   ;

EXP
   : 'EXP'
   ;

FACTORIAL
   : 'FACTORIAL'
   ;

FLOOR
   : 'FLOOR'
   ;

GCD
   : 'GCD'
   ;

LCM
   : 'LCM'
   ;

LN
   : 'LN'
   ;

LOG10
   : 'LOG10'
   ;

MIN_SCALE
   : 'MIN_SCALE'
   ;

MOD
   : 'MOD'
   ;

PI
   : 'PI'
   ;

POWER
   : 'POWER'
   ;

RADIANS
   : 'RADIANS'
   ;

ROUND
   : 'ROUND'
   ;

SCALE
   : 'SCALE'
   ;

SIGN
   : 'SIGN'
   ;

SQRT
   : 'SQRT'
   ;

TRIM_SCALE
   : 'TRIM_SCALE'
   ;

TRUNC
   : 'TRUNC'
   ;

WIDTH_BUCKET
   : 'WIDTH_BUCKET'
   ;

RANDOM
   : 'RANDOM'
   ;

SETSEED
   : 'SETSEED'
   ;

ACOS
   : 'ACOS'
   ;

ACOSD
   : 'ACOSD'
   ;

ASIN
   : 'ASIN'
   ;

ASIND
   : 'ASIND'
   ;

ATAN
   : 'ATAN'
   ;

ATAND
   : 'ATAND'
   ;

ATAN2
   : 'ATAN2'
   ;

ATAN2D
   : 'ATAN2D'
   ;

COS
   : 'COS'
   ;

COSD
   : 'COSD'
   ;

COT
   : 'COT'
   ;

COTD
   : 'COTD'
   ;

SIN
   : 'SIN'
   ;

SIND
   : 'SIND'
   ;

TAN
   : 'TAN'
   ;

TAND
   : 'TAND'
   ;

SINH
   : 'SINH'
   ;

COSH
   : 'COSH'
   ;

TANH
   : 'TANH'
   ;

ASINH
   : 'ASINH'
   ;

ACOSH
   : 'ACOSH'
   ;

ATANH
   : 'ATANH'
   ;

BIT_LENGTH
   : 'BIT_LENGTH'
   ;

CHAR_LENGTH
   : 'CHAR_LENGTH'
   ;

CHARACTER_LENGTH
   : 'CHARACTER_LENGTH'
   ;

LOWER
   : 'LOWER'
   ;

OCTET_LENGTH
   : 'OCTET_LENGTH'
   ;

UPPER
   : 'UPPER'
   ;

ASCII
   : 'ASCII'
   ;

BTRIM
   : 'BTRIM'
   ;

CHR
   : 'CHR'
   ;

CONCAT
   : 'CONCAT'
   ;

CONCAT_WS
   : 'CONCAT_WS'
   ;

FORMAT
   : 'FORMAT'
   ;

INITCAP
   : 'INITCAP'
   ;

LENGTH
   : 'LENGTH'
   ;

LPAD
   : 'LPAD'
   ;

LTRIM
   : 'LTRIM'
   ;

MD5
   : 'MD5'
   ;

PARSE_IDENT
   : 'PARSE_IDENT'
   ;

PG_CLIENT_ENCODING
   : 'PG_CLIENT_ENCODING'
   ;

QUOTE_IDENT
   : 'QUOTE_IDENT'
   ;

QUOTE_LITERAL
   : 'QUOTE_LITERAL'
   ;

QUOTE_NULLABLE
   : 'QUOTE_NULLABLE'
   ;

REGEXP_COUNT
   : 'REGEXP_COUNT'
   ;

REGEXP_INSTR
   : 'REGEXP_INSTR'
   ;

REGEXP_LIKE
   : 'REGEXP_LIKE'
   ;

REGEXP_MATCH
   : 'REGEXP_MATCH'
   ;

REGEXP_MATCHES
   : 'REGEXP_MATCHES'
   ;

REGEXP_REPLACE
   : 'REGEXP_REPLACE'
   ;

REGEXP_SPLIT_TO_ARRAY
   : 'REGEXP_SPLIT_TO_ARRAY'
   ;

REGEXP_SPLIT_TO_TABLE
   : 'REGEXP_SPLIT_TO_TABLE'
   ;

REGEXP_SUBSTR
   : 'REGEXP_SUBSTR'
   ;

REPEAT
   : 'REPEAT'
   ;

RPAD
   : 'RPAD'
   ;

RTRIM
   : 'RTRIM'
   ;

SPLIT_PART
   : 'SPLIT_PART'
   ;

STARTS_WITH
   : 'STARTS_WITH'
   ;

STRING_TO_ARRAY
   : 'STRING_TO_ARRAY'
   ;

STRING_TO_TABLE
   : 'STRING_TO_TABLE'
   ;

STRPOS
   : 'STRPOS'
   ;

SUBSTR
   : 'SUBSTR'
   ;

TO_ASCII
   : 'TO_ASCII'
   ;

TO_HEX
   : 'TO_HEX'
   ;

TRANSLATE
   : 'TRANSLATE'
   ;

UNISTR
   : 'UNISTR'
   ;

AGE
   : 'AGE'
   ;

CLOCK_TIMESTAMP
   : 'CLOCK_TIMESTAMP'
   ;

DATE_BIN
   : 'DATE_BIN'
   ;

DATE_PART
   : 'DATE_PART'
   ;

DATE_TRUNC
   : 'DATE_TRUNC'
   ;

ISFINITE
   : 'ISFINITE'
   ;

JUSTIFY_DAYS
   : 'JUSTIFY_DAYS'
   ;

JUSTIFY_HOURS
   : 'JUSTIFY_HOURS'
   ;

JUSTIFY_INTERVAL
   : 'JUSTIFY_INTERVAL'
   ;

MAKE_DATE
   : 'MAKE_DATE'
   ;

MAKE_INTERVAL
   : 'MAKE_INTERVAL'
   ;

MAKE_TIME
   : 'MAKE_TIME'
   ;

MAKE_TIMESTAMP
   : 'MAKE_TIMESTAMP'
   ;

MAKE_TIMESTAMPTZ
   : 'MAKE_TIMESTAMPTZ'
   ;

NOW
   : 'NOW'
   ;

STATEMENT_TIMESTAMP
   : 'STATEMENT_TIMESTAMP'
   ;

TIMEOFDAY
   : 'TIMEOFDAY'
   ;

TRANSACTION_TIMESTAMP
   : 'TRANSACTION_TIMESTAMP'
   ;

TO_TIMESTAMP
   : 'TO_TIMESTAMP'
   ;

TO_CHAR
   : 'TO_CHAR'
   ;

TO_DATE
   : 'TO_DATE'
   ;

TO_NUMBER
   : 'TO_NUMBER'
   ;

ENCODE
   : 'ENCODE'
   ;

DISTKEY
   : 'DISTKEY'
   ;

SORTKEY
   : 'SORTKEY'
   ;

DISTSTYLE
   : 'DISTSTYLE'
   ;

BACKUP
   : 'BACKUP'
   ;

COMPOUND
   : 'COMPOUND'
   ;

INTERLEAVED
   : 'INTERLEAVED'
   ;

EVEN
   : 'EVEN'
   ;

CASE_SENSITIVE
   : 'CASE_SENSITIVE'
   ;

QUOTA
   : 'QUOTA'
   ;

TB
   : 'TB'
   ;

BOOST
   : 'BOOST'
   ;

RECLUSTER
   : 'RECLUSTER'
   ;

SORT
   : 'SORT'
   ;

PERCENT_WORD
   : 'PERCENT'
   ;

CASE_INSENSITIVE
   : 'CASE_INSENSITIVE'
   ;

Identifier
   : IdentifierStartChar IdentifierChar*
   ;

TemporaryIdentifier
   : '#' Identifier
   ;

NamespaceUser
   : Identifier ':' Identifier
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
