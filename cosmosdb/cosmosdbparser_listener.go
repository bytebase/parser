// Code generated from CosmosDBParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package cosmosdb // CosmosDBParser
import "github.com/antlr4-go/antlr/v4"

// CosmosDBParserListener is a complete listener for a parse tree produced by CosmosDBParser.
type CosmosDBParserListener interface {
	antlr.ParseTreeListener

	// EnterRoot is called when entering the root production.
	EnterRoot(c *RootContext)

	// EnterSelect is called when entering the select production.
	EnterSelect(c *SelectContext)

	// EnterSelect_clause is called when entering the select_clause production.
	EnterSelect_clause(c *Select_clauseContext)

	// EnterSelect_specification is called when entering the select_specification production.
	EnterSelect_specification(c *Select_specificationContext)

	// EnterFrom_clause is called when entering the from_clause production.
	EnterFrom_clause(c *From_clauseContext)

	// EnterWhere_clause is called when entering the where_clause production.
	EnterWhere_clause(c *Where_clauseContext)

	// EnterFrom_specification is called when entering the from_specification production.
	EnterFrom_specification(c *From_specificationContext)

	// EnterFrom_source is called when entering the from_source production.
	EnterFrom_source(c *From_sourceContext)

	// EnterContainer_expression is called when entering the container_expression production.
	EnterContainer_expression(c *Container_expressionContext)

	// EnterContainer_name is called when entering the container_name production.
	EnterContainer_name(c *Container_nameContext)

	// EnterObject_property_list is called when entering the object_property_list production.
	EnterObject_property_list(c *Object_property_listContext)

	// EnterObject_property is called when entering the object_property production.
	EnterObject_property(c *Object_propertyContext)

	// EnterProperty_alias is called when entering the property_alias production.
	EnterProperty_alias(c *Property_aliasContext)

	// EnterScalar_expression is called when entering the scalar_expression production.
	EnterScalar_expression(c *Scalar_expressionContext)

	// EnterScalar_expression_in_where is called when entering the scalar_expression_in_where production.
	EnterScalar_expression_in_where(c *Scalar_expression_in_whereContext)

	// EnterCreate_array_expression is called when entering the create_array_expression production.
	EnterCreate_array_expression(c *Create_array_expressionContext)

	// EnterCreate_object_expression is called when entering the create_object_expression production.
	EnterCreate_object_expression(c *Create_object_expressionContext)

	// EnterScalar_function_expression is called when entering the scalar_function_expression production.
	EnterScalar_function_expression(c *Scalar_function_expressionContext)

	// EnterUdf_scalar_function_expression is called when entering the udf_scalar_function_expression production.
	EnterUdf_scalar_function_expression(c *Udf_scalar_function_expressionContext)

	// EnterBuiltin_function_expression is called when entering the builtin_function_expression production.
	EnterBuiltin_function_expression(c *Builtin_function_expressionContext)

	// EnterBinary_operator is called when entering the binary_operator production.
	EnterBinary_operator(c *Binary_operatorContext)

	// EnterUnary_operator is called when entering the unary_operator production.
	EnterUnary_operator(c *Unary_operatorContext)

	// EnterParameter_name is called when entering the parameter_name production.
	EnterParameter_name(c *Parameter_nameContext)

	// EnterConstant is called when entering the constant production.
	EnterConstant(c *ConstantContext)

	// EnterObject_constant is called when entering the object_constant production.
	EnterObject_constant(c *Object_constantContext)

	// EnterObject_constant_field_pair is called when entering the object_constant_field_pair production.
	EnterObject_constant_field_pair(c *Object_constant_field_pairContext)

	// EnterArray_constant is called when entering the array_constant production.
	EnterArray_constant(c *Array_constantContext)

	// EnterString_constant is called when entering the string_constant production.
	EnterString_constant(c *String_constantContext)

	// EnterUndefined_constant is called when entering the undefined_constant production.
	EnterUndefined_constant(c *Undefined_constantContext)

	// EnterNull_constant is called when entering the null_constant production.
	EnterNull_constant(c *Null_constantContext)

	// EnterBoolean_constant is called when entering the boolean_constant production.
	EnterBoolean_constant(c *Boolean_constantContext)

	// EnterNumber_constant is called when entering the number_constant production.
	EnterNumber_constant(c *Number_constantContext)

	// EnterString_literal is called when entering the string_literal production.
	EnterString_literal(c *String_literalContext)

	// EnterDecimal_literal is called when entering the decimal_literal production.
	EnterDecimal_literal(c *Decimal_literalContext)

	// EnterHexadecimal_literal is called when entering the hexadecimal_literal production.
	EnterHexadecimal_literal(c *Hexadecimal_literalContext)

	// EnterProperty_name is called when entering the property_name production.
	EnterProperty_name(c *Property_nameContext)

	// EnterArray_index is called when entering the array_index production.
	EnterArray_index(c *Array_indexContext)

	// EnterInput_alias is called when entering the input_alias production.
	EnterInput_alias(c *Input_aliasContext)

	// ExitRoot is called when exiting the root production.
	ExitRoot(c *RootContext)

	// ExitSelect is called when exiting the select production.
	ExitSelect(c *SelectContext)

	// ExitSelect_clause is called when exiting the select_clause production.
	ExitSelect_clause(c *Select_clauseContext)

	// ExitSelect_specification is called when exiting the select_specification production.
	ExitSelect_specification(c *Select_specificationContext)

	// ExitFrom_clause is called when exiting the from_clause production.
	ExitFrom_clause(c *From_clauseContext)

	// ExitWhere_clause is called when exiting the where_clause production.
	ExitWhere_clause(c *Where_clauseContext)

	// ExitFrom_specification is called when exiting the from_specification production.
	ExitFrom_specification(c *From_specificationContext)

	// ExitFrom_source is called when exiting the from_source production.
	ExitFrom_source(c *From_sourceContext)

	// ExitContainer_expression is called when exiting the container_expression production.
	ExitContainer_expression(c *Container_expressionContext)

	// ExitContainer_name is called when exiting the container_name production.
	ExitContainer_name(c *Container_nameContext)

	// ExitObject_property_list is called when exiting the object_property_list production.
	ExitObject_property_list(c *Object_property_listContext)

	// ExitObject_property is called when exiting the object_property production.
	ExitObject_property(c *Object_propertyContext)

	// ExitProperty_alias is called when exiting the property_alias production.
	ExitProperty_alias(c *Property_aliasContext)

	// ExitScalar_expression is called when exiting the scalar_expression production.
	ExitScalar_expression(c *Scalar_expressionContext)

	// ExitScalar_expression_in_where is called when exiting the scalar_expression_in_where production.
	ExitScalar_expression_in_where(c *Scalar_expression_in_whereContext)

	// ExitCreate_array_expression is called when exiting the create_array_expression production.
	ExitCreate_array_expression(c *Create_array_expressionContext)

	// ExitCreate_object_expression is called when exiting the create_object_expression production.
	ExitCreate_object_expression(c *Create_object_expressionContext)

	// ExitScalar_function_expression is called when exiting the scalar_function_expression production.
	ExitScalar_function_expression(c *Scalar_function_expressionContext)

	// ExitUdf_scalar_function_expression is called when exiting the udf_scalar_function_expression production.
	ExitUdf_scalar_function_expression(c *Udf_scalar_function_expressionContext)

	// ExitBuiltin_function_expression is called when exiting the builtin_function_expression production.
	ExitBuiltin_function_expression(c *Builtin_function_expressionContext)

	// ExitBinary_operator is called when exiting the binary_operator production.
	ExitBinary_operator(c *Binary_operatorContext)

	// ExitUnary_operator is called when exiting the unary_operator production.
	ExitUnary_operator(c *Unary_operatorContext)

	// ExitParameter_name is called when exiting the parameter_name production.
	ExitParameter_name(c *Parameter_nameContext)

	// ExitConstant is called when exiting the constant production.
	ExitConstant(c *ConstantContext)

	// ExitObject_constant is called when exiting the object_constant production.
	ExitObject_constant(c *Object_constantContext)

	// ExitObject_constant_field_pair is called when exiting the object_constant_field_pair production.
	ExitObject_constant_field_pair(c *Object_constant_field_pairContext)

	// ExitArray_constant is called when exiting the array_constant production.
	ExitArray_constant(c *Array_constantContext)

	// ExitString_constant is called when exiting the string_constant production.
	ExitString_constant(c *String_constantContext)

	// ExitUndefined_constant is called when exiting the undefined_constant production.
	ExitUndefined_constant(c *Undefined_constantContext)

	// ExitNull_constant is called when exiting the null_constant production.
	ExitNull_constant(c *Null_constantContext)

	// ExitBoolean_constant is called when exiting the boolean_constant production.
	ExitBoolean_constant(c *Boolean_constantContext)

	// ExitNumber_constant is called when exiting the number_constant production.
	ExitNumber_constant(c *Number_constantContext)

	// ExitString_literal is called when exiting the string_literal production.
	ExitString_literal(c *String_literalContext)

	// ExitDecimal_literal is called when exiting the decimal_literal production.
	ExitDecimal_literal(c *Decimal_literalContext)

	// ExitHexadecimal_literal is called when exiting the hexadecimal_literal production.
	ExitHexadecimal_literal(c *Hexadecimal_literalContext)

	// ExitProperty_name is called when exiting the property_name production.
	ExitProperty_name(c *Property_nameContext)

	// ExitArray_index is called when exiting the array_index production.
	ExitArray_index(c *Array_indexContext)

	// ExitInput_alias is called when exiting the input_alias production.
	ExitInput_alias(c *Input_aliasContext)
}
