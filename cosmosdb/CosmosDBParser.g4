parser grammar CosmosDBParser;

options {
	tokenVocab = CosmosDBLexer;
}

root: select EOF;

select:
	select_clause from_clause? where_clause? group_by_clause? having_clause? order_by_clause?
		offset_limit_clause?;

select_clause: SELECT_SYMBOL top_clause? select_specification;

top_clause: TOP_SYMBOL DECIMAL;

select_specification:
	MULTIPLY_OPERATOR
	| DISTINCT_SYMBOL? VALUE_SYMBOL? object_property_list;

from_clause: FROM_SYMBOL from_specification;

where_clause: WHERE_SYMBOL scalar_expression;

group_by_clause:
	GROUP_SYMBOL BY_SYMBOL scalar_expression (
		COMMA_SYMBOL scalar_expression
	)*;

having_clause: HAVING_SYMBOL scalar_expression;

order_by_clause:
	ORDER_SYMBOL BY_SYMBOL sort_expression (
		COMMA_SYMBOL sort_expression
	)*;

sort_expression: scalar_expression (ASC_SYMBOL | DESC_SYMBOL)?;

offset_limit_clause: OFFSET_SYMBOL DECIMAL LIMIT_SYMBOL DECIMAL;

from_specification: from_source;

from_source: container_expression (join_clause)*;

container_expression: container_name (AS_SYMBOL? identifier)?;

join_clause:
	JOIN_SYMBOL identifier IN_SYMBOL scalar_expression;

container_name: identifier;

object_property_list:
	object_property (COMMA_SYMBOL object_property)*;

object_property:
	scalar_expression (AS_SYMBOL? property_alias)?;

property_alias: identifier;

// Unified scalar_expression - used in both SELECT projections and WHERE clause.
// https://learn.microsoft.com/en-us/azure/cosmos-db/nosql/query/scalar-expressions
scalar_expression:
	constant
	| input_alias
	| parameter_name
	| scalar_expression AND_SYMBOL scalar_expression
	| scalar_expression OR_SYMBOL scalar_expression
	| scalar_expression DOT_SYMBOL property_name
	| scalar_expression LS_BRACKET_SYMBOL (
		DOUBLE_QUOTE_STRING_LITERAL
		| SINGLE_QUOTE_STRING_LITERAL
		| array_index
	) RS_BRACKET_SYMBOL
	| unary_operator scalar_expression
	| NOT_SYMBOL scalar_expression
	| scalar_expression binary_operator scalar_expression
	| scalar_expression NOT_SYMBOL? IN_SYMBOL LR_BRACKET_SYMBOL (
		scalar_expression (COMMA_SYMBOL scalar_expression)*
	)? RR_BRACKET_SYMBOL
	| scalar_expression NOT_SYMBOL? BETWEEN_SYMBOL scalar_expression AND_SYMBOL scalar_expression
	| scalar_expression NOT_SYMBOL? LIKE_SYMBOL scalar_expression
	| EXISTS_SYMBOL LR_BRACKET_SYMBOL select RR_BRACKET_SYMBOL
	| scalar_expression QUESTION_MARK_SYMBOL scalar_expression COLON_SYMBOL scalar_expression
	| scalar_function_expression
	| create_object_expression
	| create_array_expression
	| LR_BRACKET_SYMBOL scalar_expression RR_BRACKET_SYMBOL
	| LR_BRACKET_SYMBOL select RR_BRACKET_SYMBOL;

create_array_expression:
	LS_BRACKET_SYMBOL (
		scalar_expression (COMMA_SYMBOL scalar_expression)*
	)? RS_BRACKET_SYMBOL;

create_object_expression:
	LC_BRACKET_SYMBOL (
		object_field_pair (COMMA_SYMBOL object_field_pair)*
	)? RC_BRACKET_SYMBOL;

object_field_pair:
	(string_literal | property_name) COLON_SYMBOL scalar_expression;

scalar_function_expression:
	udf_scalar_function_expression
	| builtin_function_expression;

udf_scalar_function_expression:
	UDF_SYMBOL DOT_SYMBOL identifier LR_BRACKET_SYMBOL (
		scalar_expression (COMMA_SYMBOL scalar_expression)*
	)? RR_BRACKET_SYMBOL;

builtin_function_expression:
	identifier LR_BRACKET_SYMBOL (
		(MULTIPLY_OPERATOR | scalar_expression) (
			COMMA_SYMBOL scalar_expression
		)*
	)? RR_BRACKET_SYMBOL;

binary_operator:
	MULTIPLY_OPERATOR
	| DIVIDE_SYMBOL
	| MODULO_SYMBOL
	| PLUS_SYMBOL
	| MINUS_SYMBOL
	| BIT_AND_SYMBOL
	| BIT_XOR_SYMBOL
	| BIT_OR_SYMBOL
	| DOUBLE_BAR_SYMBOL
	| EQUAL_SYMBOL
	| NOT_EQUAL_OPERATOR
	| LESS_THAN_OPERATOR
	| LESS_THAN_EQUAL_OPERATOR
	| GREATER_THAN_OPERATOR
	| GREATER_THAN_EQUAL_OPERATOR
	| LEFT_SHIFT_OPERATOR
	| RIGHT_SHIFT_OPERATOR
	| ZERO_FILL_RIGHT_SHIFT_OPERATOR;

unary_operator: BIT_NOT_SYMBOL | PLUS_SYMBOL | MINUS_SYMBOL;

parameter_name: AT_SYMBOL identifier;

// https://learn.microsoft.com/en-us/azure/cosmos-db/nosql/query/constants
constant:
	undefined_constant
	| null_constant
	| boolean_constant
	| number_constant
	| string_constant;

undefined_constant: UNDEFINED_SYMBOL;

null_constant: NULL_SYMBOL;

boolean_constant: TRUE_SYMBOL | FALSE_SYMBOL;

number_constant: decimal_literal | hexadecimal_literal;

string_constant: string_literal;

string_literal:
	SINGLE_QUOTE_STRING_LITERAL
	| DOUBLE_QUOTE_STRING_LITERAL;

decimal_literal: DECIMAL | REAL | FLOAT;

hexadecimal_literal: HEXADECIMAL;

// Allow keywords to be used as identifiers (property names, aliases, etc.)
// This is necessary because CosmosDB allows keywords as property names.
identifier:
	IDENTIFIER
	| IN_SYMBOL
	| BETWEEN_SYMBOL
	| TOP_SYMBOL
	| VALUE_SYMBOL
	| ORDER_SYMBOL
	| BY_SYMBOL
	| GROUP_SYMBOL
	| OFFSET_SYMBOL
	| LIMIT_SYMBOL
	| ASC_SYMBOL
	| DESC_SYMBOL
	| EXISTS_SYMBOL
	| LIKE_SYMBOL
	| HAVING_SYMBOL
	| JOIN_SYMBOL;

property_name: identifier;

array_index: DECIMAL;

input_alias: identifier;
