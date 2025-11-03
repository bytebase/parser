lexer grammar CosmosDBLexer;

options {
	caseInsensitive = true;
}

fragment A: [a];
fragment B: [b];
fragment C: [c];
fragment D: [d];
fragment E: [e];
fragment F: [f];
fragment G: [g];
fragment H: [h];
fragment I: [i];
fragment J: [j];
fragment K: [k];
fragment L: [l];
fragment M: [m];
fragment N: [n];
fragment O: [o];
fragment P: [p];
fragment Q: [q];
fragment R: [r];
fragment S: [s];
fragment T: [t];
fragment U: [u];
fragment V: [v];
fragment W: [w];
fragment X: [x];
fragment Y: [y];
fragment Z: [z];

MULTIPLY_OPERATOR: '*';

AS_SYMBOL: 'AS';
SELECT_SYMBOL: 'SELECT';
FROM_SYMBOL: 'FROM';
DISTINCT_SYMBOL: 'DISTINCT';
UNDEFINED_SYMBOL: 'UNDEFINED';
NULL_SYMBOL: 'NULL';
FALSE_SYMBOL: 'FALSE';
TRUE_SYMBOL: 'TRUE';
NOT_SYMBOL: 'NOT';
UDF_SYMBOL: 'UDF';
WHERE_SYMBOL: 'WHERE';
AND_SYMBOL: 'AND';
OR_SYMBOL: 'OR';

AT_SYMBOL: '@';
LC_BRACKET_SYMBOL: '{';
RC_BRACKET_SYMBOL: '}';
LS_BRACKET_SYMBOL: '[';
RS_BRACKET_SYMBOL: ']';
LR_BRACKET_SYMBOL: '(';
RR_BRACKET_SYMBOL: ')';
SINGLE_QUOTE_SYMBOL: '\'';
DOUBLE_QUOTE_SYMBOL: '"';
COMMA_SYMBOL: ',';
DOT_SYMBOL: '.';
QUESTION_MARK_SYMBOL: '?';
COLON_SYMBOL: ':';
PLUS_SYMBOL: '+';
MINUS_SYMBOL: '-';
BIT_NOT_SYMBOL: '~';
DIVIDE_SYMBOL: '/';
MODULO_SYMBOL: '%';
BIT_AND_SYMBOL: '&';
BIT_OR_SYMBOL: '|';
DOUBLE_BAR_SYMBOL: '||';
BIT_XOR_SYMBOL: '^';
EQUAL_SYMBOL: '=';

/* Identifiers */
IDENTIFIER: [a-z] [a-z_0-9]*;

// White space handling
WHITESPACE:
	[ \t\f\r\n] -> channel(HIDDEN); // Ignore whitespaces.

// Decimal literal.
fragment DEC_DIGIT: [0-9];
fragment DEC_DOT_DEC: (
		DEC_DIGIT+ '.' DEC_DIGIT+
		| DEC_DIGIT+ '.'
		| '.' DEC_DIGIT+
	);

DECIMAL: DEC_DIGIT+;
REAL: (DECIMAL | DEC_DOT_DEC) ('E' [+-]? DEC_DIGIT+);
FLOAT: DEC_DOT_DEC;

// Hexadecimal literal.
fragment HEX_DIGIT: [0-9A-F];
HEXADECIMAL: '0' 'X' HEX_DIGIT+;

fragment FullWidthLetter options {
	caseInsensitive = false;
}:
	'\u00c0' ..'\u00d6'
	| '\u00d8' ..'\u00f6'
	| '\u00f8' ..'\u00ff'
	| '\u0100' ..'\u1fff'
	| '\u2c00' ..'\u2fff'
	| '\u3040' ..'\u318f'
	| '\u3300' ..'\u337f'
	| '\u3400' ..'\u3fff'
	| '\u4e00' ..'\u9fff'
	| '\ua000' ..'\ud7ff'
	| '\uf900' ..'\ufaff'
	| '\uff00' ..'\ufff0';
// | '\u10000'..'\u1F9FF' //not support four bytes chars | '\u20000'..'\u2FA1F'

// String literal.
fragment ESCAPE_SEQUENCE:
	'\\' [btnrf"'\\/] // Basic escape sequences: \b, \t, \n, \r, \f, ", ', \, /
	| '\\u' HEX_DIGIT HEX_DIGIT HEX_DIGIT HEX_DIGIT; // Unicode escape: \uXXXX

fragment STRING_CHAR:
	ESCAPE_SEQUENCE
	| ~[\\"'\r\n]; // Any Unicode character EXCEPT: \, ", ', \r, \n

// String literals
SINGLE_QUOTE_STRING_LITERAL:
    SINGLE_QUOTE_SYMBOL STRING_CHAR* SINGLE_QUOTE_SYMBOL;


DOUBLE_QUOTE_STRING_LITERAL:
    DOUBLE_QUOTE_SYMBOL STRING_CHAR* DOUBLE_QUOTE_SYMBOL;
