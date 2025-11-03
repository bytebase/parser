// Code generated from CosmosDBParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package cosmosdb // CosmosDBParser
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by CosmosDBParser.
type CosmosDBParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by CosmosDBParser#root.
	VisitRoot(ctx *RootContext) interface{}

	// Visit a parse tree produced by CosmosDBParser#select.
	VisitSelect(ctx *SelectContext) interface{}

	// Visit a parse tree produced by CosmosDBParser#select_clause.
	VisitSelect_clause(ctx *Select_clauseContext) interface{}

	// Visit a parse tree produced by CosmosDBParser#select_specification.
	VisitSelect_specification(ctx *Select_specificationContext) interface{}

	// Visit a parse tree produced by CosmosDBParser#from_clause.
	VisitFrom_clause(ctx *From_clauseContext) interface{}

	// Visit a parse tree produced by CosmosDBParser#where_clause.
	VisitWhere_clause(ctx *Where_clauseContext) interface{}

	// Visit a parse tree produced by CosmosDBParser#from_specification.
	VisitFrom_specification(ctx *From_specificationContext) interface{}

	// Visit a parse tree produced by CosmosDBParser#from_source.
	VisitFrom_source(ctx *From_sourceContext) interface{}

	// Visit a parse tree produced by CosmosDBParser#container_expression.
	VisitContainer_expression(ctx *Container_expressionContext) interface{}

	// Visit a parse tree produced by CosmosDBParser#container_name.
	VisitContainer_name(ctx *Container_nameContext) interface{}

	// Visit a parse tree produced by CosmosDBParser#object_property_list.
	VisitObject_property_list(ctx *Object_property_listContext) interface{}

	// Visit a parse tree produced by CosmosDBParser#object_property.
	VisitObject_property(ctx *Object_propertyContext) interface{}

	// Visit a parse tree produced by CosmosDBParser#property_alias.
	VisitProperty_alias(ctx *Property_aliasContext) interface{}

	// Visit a parse tree produced by CosmosDBParser#scalar_expression.
	VisitScalar_expression(ctx *Scalar_expressionContext) interface{}

	// Visit a parse tree produced by CosmosDBParser#scalar_expression_in_where.
	VisitScalar_expression_in_where(ctx *Scalar_expression_in_whereContext) interface{}

	// Visit a parse tree produced by CosmosDBParser#create_array_expression.
	VisitCreate_array_expression(ctx *Create_array_expressionContext) interface{}

	// Visit a parse tree produced by CosmosDBParser#create_object_expression.
	VisitCreate_object_expression(ctx *Create_object_expressionContext) interface{}

	// Visit a parse tree produced by CosmosDBParser#scalar_function_expression.
	VisitScalar_function_expression(ctx *Scalar_function_expressionContext) interface{}

	// Visit a parse tree produced by CosmosDBParser#udf_scalar_function_expression.
	VisitUdf_scalar_function_expression(ctx *Udf_scalar_function_expressionContext) interface{}

	// Visit a parse tree produced by CosmosDBParser#builtin_function_expression.
	VisitBuiltin_function_expression(ctx *Builtin_function_expressionContext) interface{}

	// Visit a parse tree produced by CosmosDBParser#binary_operator.
	VisitBinary_operator(ctx *Binary_operatorContext) interface{}

	// Visit a parse tree produced by CosmosDBParser#unary_operator.
	VisitUnary_operator(ctx *Unary_operatorContext) interface{}

	// Visit a parse tree produced by CosmosDBParser#parameter_name.
	VisitParameter_name(ctx *Parameter_nameContext) interface{}

	// Visit a parse tree produced by CosmosDBParser#constant.
	VisitConstant(ctx *ConstantContext) interface{}

	// Visit a parse tree produced by CosmosDBParser#object_constant.
	VisitObject_constant(ctx *Object_constantContext) interface{}

	// Visit a parse tree produced by CosmosDBParser#object_constant_field_pair.
	VisitObject_constant_field_pair(ctx *Object_constant_field_pairContext) interface{}

	// Visit a parse tree produced by CosmosDBParser#array_constant.
	VisitArray_constant(ctx *Array_constantContext) interface{}

	// Visit a parse tree produced by CosmosDBParser#string_constant.
	VisitString_constant(ctx *String_constantContext) interface{}

	// Visit a parse tree produced by CosmosDBParser#undefined_constant.
	VisitUndefined_constant(ctx *Undefined_constantContext) interface{}

	// Visit a parse tree produced by CosmosDBParser#null_constant.
	VisitNull_constant(ctx *Null_constantContext) interface{}

	// Visit a parse tree produced by CosmosDBParser#boolean_constant.
	VisitBoolean_constant(ctx *Boolean_constantContext) interface{}

	// Visit a parse tree produced by CosmosDBParser#number_constant.
	VisitNumber_constant(ctx *Number_constantContext) interface{}

	// Visit a parse tree produced by CosmosDBParser#string_literal.
	VisitString_literal(ctx *String_literalContext) interface{}

	// Visit a parse tree produced by CosmosDBParser#decimal_literal.
	VisitDecimal_literal(ctx *Decimal_literalContext) interface{}

	// Visit a parse tree produced by CosmosDBParser#hexadecimal_literal.
	VisitHexadecimal_literal(ctx *Hexadecimal_literalContext) interface{}

	// Visit a parse tree produced by CosmosDBParser#property_name.
	VisitProperty_name(ctx *Property_nameContext) interface{}

	// Visit a parse tree produced by CosmosDBParser#array_index.
	VisitArray_index(ctx *Array_indexContext) interface{}

	// Visit a parse tree produced by CosmosDBParser#input_alias.
	VisitInput_alias(ctx *Input_aliasContext) interface{}
}
