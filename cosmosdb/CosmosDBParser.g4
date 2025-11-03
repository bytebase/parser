parser grammar CosmosDBParser;

options {
	tokenVocab = CosmosDBLexer;
}

root: select EOF;

select: select_clause from_clause where_clause?;

select_clause: SELECT_SYMBOL select_specification;

select_specification:
	MULTIPLY_OPERATOR
	| DISTINCT_SYMBOL? object_property_list;

from_clause: FROM_SYMBOL from_specification;

where_clause: WHERE_SYMBOL scalar_expression_in_where;

from_specification: from_source;

from_source: container_expression;

container_expression: container_name (AS_SYMBOL? IDENTIFIER)?;

container_name: IDENTIFIER;

object_property_list:
	object_property (COMMA_SYMBOL object_property)*;

object_property: scalar_expression (AS_SYMBOL? property_alias)?;

property_alias: IDENTIFIER;

// scalar_expression: https://learn.microsoft.com/en-us/azure/cosmos-db/nosql/query/scalar-expressions
scalar_expression:
	input_alias
	| scalar_expression DOT_SYMBOL property_name
	| scalar_expression LS_BRACKET_SYMBOL (
		(DOUBLE_QUOTE_STRING_LITERAL)
		| (array_index)
	) RS_BRACKET_SYMBOL
	| unary_operator scalar_expression;

// TODO(zp): Merge scalar_expression and scalar_expression_in_where while supporting the project
// fully. https://learn.microsoft.com/en-us/azure/cosmos-db/nosql/query/scalar-expressions
scalar_expression_in_where:
	constant
	| input_alias
	| parameter_name
	| scalar_expression_in_where AND_SYMBOL scalar_expression_in_where
	| scalar_expression_in_where OR_SYMBOL scalar_expression_in_where
	| scalar_expression_in_where DOT_SYMBOL property_name
	| scalar_expression_in_where LS_BRACKET_SYMBOL (
		(DOUBLE_QUOTE_STRING_LITERAL)
		| (array_index)
	) RS_BRACKET_SYMBOL
	| unary_operator scalar_expression_in_where
	| scalar_expression_in_where binary_operator scalar_expression_in_where
	| scalar_expression_in_where QUESTION_MARK_SYMBOL scalar_expression_in_where COLON_SYMBOL
		scalar_expression_in_where
	| scalar_function_expression
	| create_object_expression
	| create_array_expression
	| LR_BRACKET_SYMBOL scalar_expression_in_where RR_BRACKET_SYMBOL;

create_array_expression: array_constant;

create_object_expression: object_constant;

scalar_function_expression:
	udf_scalar_function_expression
	| builtin_function_expression;

udf_scalar_function_expression:
	UDF_SYMBOL DOT_SYMBOL IDENTIFIER LR_BRACKET_SYMBOL (
		scalar_expression_in_where (
			COMMA_SYMBOL scalar_expression_in_where
		)*
	) RR_BRACKET_SYMBOL;

builtin_function_expression:
	IDENTIFIER LR_BRACKET_SYMBOL (
		scalar_expression_in_where (
			COMMA_SYMBOL scalar_expression_in_where
		)*
	) RR_BRACKET_SYMBOL;

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
	| EQUAL_SYMBOL;

unary_operator: BIT_NOT_SYMBOL | PLUS_SYMBOL | MINUS_SYMBOL;

parameter_name: AT_SYMBOL IDENTIFIER;

// https://learn.microsoft.com/en-us/azure/cosmos-db/nosql/query/constants
constant:
	undefined_constant
	| null_constant
	| boolean_constant
	| number_constant
	| string_constant
	| array_constant
	| object_constant;

object_constant:
	LC_BRACKET_SYMBOL (
		object_constant_field_pair (
			COMMA_SYMBOL object_constant_field_pair
		)*
	) RC_BRACKET_SYMBOL;

object_constant_field_pair: (
		property_name
		| (DOUBLE_QUOTE_SYMBOL property_name DOUBLE_QUOTE_SYMBOL)
	) COMMA_SYMBOL constant;

array_constant:
	LS_BRACKET_SYMBOL (constant (COMMA_SYMBOL constant)*)? RS_BRACKET_SYMBOL;

string_constant: string_literal;

undefined_constant: UNDEFINED_SYMBOL;

null_constant: NULL_SYMBOL;

boolean_constant: TRUE_SYMBOL | FALSE_SYMBOL;

number_constant: decimal_literal | hexadecimal_literal;

string_literal:
	SINGLE_QUOTE_STRING_LITERAL
	| DOUBLE_QUOTE_STRING_LITERAL;

decimal_literal: DECIMAL | REAL | FLOAT;

hexadecimal_literal: HEXADECIMAL;

property_name: IDENTIFIER;

array_index: DECIMAL;

input_alias: IDENTIFIER;