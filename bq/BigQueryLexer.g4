lexer grammar BigQueryLexer;

// ARRAY and STRUCT included in the list of BQ keywords instead of here
QUOTE : '\'' ;
DQOUTE : '"';
SEMI : ';';

/*
 * BigQuery Keywords:
 * Based off the list of BigQuery Reserved Keywords at:
 * https://cloud.google.com/bigquery/docs/reference/standard-sql/lexical 
 */
ALL : A L L ;
AND : A N D ;
ANY : A N Y ;
ARRAY : A R R A Y ;
AS : A S ;
ASC : A S C ;
ASSERT_ROWS_MODIFIED : A S S E R T [_] R O W S [_] M O D I F I E D ;
AT : A T ;
BETWEEN : B E T W E E N ;
BY : B Y ;
CASE : C A S E ;
CAST : C A S T ;
COLLATE : C O L L A T E ;
CONTAINS : C O N T A I N S ;
CREATE : C R E A T E ;
CROSS : C R O S S ;
CUBE : C U B E ;
CURRENT : C U R R E N T ;
DEFAULT : D E F A U L T ;
DEFINE : D E F I N E ;
DESC : D E S C ;
DISTINCT : D I S T I N C T ;
ELSE : E L S E ;
END : E N D ;
ENUM : E N U M ;
ESCAPE : E S C A P E ;
EXCEPT : E X C E P T ;
EXCLUDE : E X C L U D E ;
EXISTS : E X I S T S ;
EXTRACT : E X T R A C T ;
FALSE : F A L S E ;
FETCH : F E T C H ;
FOLLOWING : F O L L O W I N G ;
FOR : F O R ;
FROM : F R O M ;
FULL : F U L L ;
GROUP : G R O U P ;
GROUPING : G R O U P I N G ;
GROUPS : G R O U P S ;
HASH : H A S H ;
HAVING : H A V I N G ;
IF : I F ;
IGNORE : I G N O R E ;
IN : I N ;
INNER : I N N E R ;
INTERSECT : I N T E R S E C T ;
INTERVAL : I N T E R V A L ;
INTO : I N T O ;
IS : I S ;
JOIN : J O I N ;
LATERAL : L A T E R A L ;
LEFT : L E F T ;
LIKE : L I K E ;
LIMIT : L I M I T ;
LOOKUP : L O O K U P ;
MERGE : M E R G E ;
NATURAL : N A T U R A L ;
NEW : N E W ;
NO : N O ;
NOT : N O T ;
S_NULL : N U L L ;
NULLS : N U L L S ;
OF : O F ;
OFFSET : O F F S E T;
ON : O N ;
OR : O R ;
ORDER : O R D E R ;
ORDINAL : O R D I N A L;
OUTER : O U T E R ;
OVER : O V E R ;
PARTITION : P A R T I T I O N ;
PRECEDING : P R E C E D I N G ;
PROTO : P R O T O ;
RANGE : R A N G E ;
RECURSIVE : R E C U R S I V E ;
REPLACE : R E P L A C E;
RESPECT : R E S P E C T ;
RIGHT : R I G H T ;
ROLLUP : R O L L U P ;
ROWS : R O W S ;
SAFE_OFFSET : S A F E '_' O F F S E T ;
SAFE_ORDINAL : S A F E '_' O R D I N A L ;
SELECT : S E L E C T ;
SET : S E T ;
SOME : S O M E ;
SSTRUCT : S T R U C T ;
SYSTEM : S Y S T E M ;
TABLESAMPLE : T A B L E S A M P L E ;
THEN : T H E N ;
TIME : T I M E ;
TO : T O ;
TREAT : T R E A T ;
TRUE : T R U E ;
UNBOUNDED : U N B O U N D E D ;
UNION : U N I O N ;
UNNEST : U N N E S T ;
USING : U S I N G ;
WHEN : W H E N ;
WHERE : W H E R E ;
WINDOW : W I N D O W ;
WITH : W I T H ;
WITHIN : W I T H I N ;

INT : ('+' | '-')? ('0x')? DIGITS;
FLOAT : ('+' | '-')? DIGITS '.' DIGITS? ('e' ('+' | '-') DIGITS)? 
	  | DIGITS? '.' DIGITS ('e' ('+' | '-') DIGITS)?
	  | DIGITS 'e' ('+' | '-')? DIGITS;
DIGITS : DIGIT+ ;

// Whitespace
WS : [ \t\r\n]+ -> channel(HIDDEN) ;
// Comments
CMT : '--' ~[\r\n]* -> channel(HIDDEN) ;
M_CMT : '/*' .*? '*/' -> channel(HIDDEN) ;
// Quoted String
QUOTED_STRING : '"' (~'"' | '\\' '"')* '"' 
			  | '\'' (~'\'' | '\\' '\'' )* '\'' ;
TRIPLE_QUOTED_STRING : QUOTE QUOTE QUOTE .*? ~'\\' QUOTE QUOTE QUOTE 
					 | DQOUTE DQOUTE DQOUTE .*? ~'\\' DQOUTE DQOUTE DQOUTE ;
RAW_STRING : R (QUOTED_STRING | TRIPLE_QUOTED_STRING) ;
BYTE_STRING : B (QUOTED_STRING | TRIPLE_QUOTED_STRING) ;
RAW_BYTE_STRING : RB (QUOTED_STRING | TRIPLE_QUOTED_STRING) ;
// ID regex
QUOTED_ID : BACKTICK (('\\'? .))+ BACKTICK ;
ID : [a-zA-Z_][A-Za-z_0-9]* ;
RB : [rR][bB] | [bB][rR] ;

DOT:                 '.';
BACKTICK:            '`';
LR_BRACKET:          '(';
RR_BRACKET:          ')';
LSB:                 '[';
RSB:                 ']';
EQ:                  '=';
LT:                  '<';
GT:                  '>';
DOUBLE_LT:           '<<';
DOUBLE_GT:           '>>';
LE:                  '<=';
GE:                  '>=';
NE:                  '!=';
LTGT:                '<>';
BIT_AND:             '&';
BIT_OR:              '|';
STAR:                '*';
DIVIDE:              '/';
PLUS:                '+';
MINUS:               '-';
COMMA:               ',';
TILDE:               '~';
CARET:               '^';

fragment DIGIT : [0-9] ;
// Fragments for each letter of the alphabet. This is necessary because SQL keywords are case-insensitive.
fragment A : [aA];
fragment B : [bB];
fragment C : [cC];
fragment D : [dD];
fragment E : [eE];
fragment F : [fF];
fragment G : [gG];
fragment H : [hH];
fragment I : [iI];
fragment J : [jJ];
fragment K : [kK];
fragment L : [lL];
fragment M : [mM];
fragment N : [nN];
fragment O : [oO];
fragment P : [pP];
fragment Q : [qQ];
fragment R : [rR];
fragment S : [sS];
fragment T : [tT];
fragment U : [uU];
fragment V : [vV];
fragment W : [wW];
fragment X : [xX];
fragment Y : [yY];
fragment Z : [zZ];