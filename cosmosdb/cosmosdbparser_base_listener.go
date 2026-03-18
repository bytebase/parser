// Code generated from CosmosDBParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package cosmosdb // CosmosDBParser
import "github.com/antlr4-go/antlr/v4"

// BaseCosmosDBParserListener is a complete listener for a parse tree produced by CosmosDBParser.
type BaseCosmosDBParserListener struct{}

var _ CosmosDBParserListener = &BaseCosmosDBParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseCosmosDBParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseCosmosDBParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseCosmosDBParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseCosmosDBParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterRoot is called when production root is entered.
func (s *BaseCosmosDBParserListener) EnterRoot(ctx *RootContext) {}

// ExitRoot is called when production root is exited.
func (s *BaseCosmosDBParserListener) ExitRoot(ctx *RootContext) {}

// EnterSelect is called when production select is entered.
func (s *BaseCosmosDBParserListener) EnterSelect(ctx *SelectContext) {}

// ExitSelect is called when production select is exited.
func (s *BaseCosmosDBParserListener) ExitSelect(ctx *SelectContext) {}

// EnterSelect_clause is called when production select_clause is entered.
func (s *BaseCosmosDBParserListener) EnterSelect_clause(ctx *Select_clauseContext) {}

// ExitSelect_clause is called when production select_clause is exited.
func (s *BaseCosmosDBParserListener) ExitSelect_clause(ctx *Select_clauseContext) {}

// EnterTop_clause is called when production top_clause is entered.
func (s *BaseCosmosDBParserListener) EnterTop_clause(ctx *Top_clauseContext) {}

// ExitTop_clause is called when production top_clause is exited.
func (s *BaseCosmosDBParserListener) ExitTop_clause(ctx *Top_clauseContext) {}

// EnterSelect_specification is called when production select_specification is entered.
func (s *BaseCosmosDBParserListener) EnterSelect_specification(ctx *Select_specificationContext) {}

// ExitSelect_specification is called when production select_specification is exited.
func (s *BaseCosmosDBParserListener) ExitSelect_specification(ctx *Select_specificationContext) {}

// EnterFrom_clause is called when production from_clause is entered.
func (s *BaseCosmosDBParserListener) EnterFrom_clause(ctx *From_clauseContext) {}

// ExitFrom_clause is called when production from_clause is exited.
func (s *BaseCosmosDBParserListener) ExitFrom_clause(ctx *From_clauseContext) {}

// EnterWhere_clause is called when production where_clause is entered.
func (s *BaseCosmosDBParserListener) EnterWhere_clause(ctx *Where_clauseContext) {}

// ExitWhere_clause is called when production where_clause is exited.
func (s *BaseCosmosDBParserListener) ExitWhere_clause(ctx *Where_clauseContext) {}

// EnterGroup_by_clause is called when production group_by_clause is entered.
func (s *BaseCosmosDBParserListener) EnterGroup_by_clause(ctx *Group_by_clauseContext) {}

// ExitGroup_by_clause is called when production group_by_clause is exited.
func (s *BaseCosmosDBParserListener) ExitGroup_by_clause(ctx *Group_by_clauseContext) {}

// EnterHaving_clause is called when production having_clause is entered.
func (s *BaseCosmosDBParserListener) EnterHaving_clause(ctx *Having_clauseContext) {}

// ExitHaving_clause is called when production having_clause is exited.
func (s *BaseCosmosDBParserListener) ExitHaving_clause(ctx *Having_clauseContext) {}

// EnterOrder_by_clause is called when production order_by_clause is entered.
func (s *BaseCosmosDBParserListener) EnterOrder_by_clause(ctx *Order_by_clauseContext) {}

// ExitOrder_by_clause is called when production order_by_clause is exited.
func (s *BaseCosmosDBParserListener) ExitOrder_by_clause(ctx *Order_by_clauseContext) {}

// EnterSort_expression is called when production sort_expression is entered.
func (s *BaseCosmosDBParserListener) EnterSort_expression(ctx *Sort_expressionContext) {}

// ExitSort_expression is called when production sort_expression is exited.
func (s *BaseCosmosDBParserListener) ExitSort_expression(ctx *Sort_expressionContext) {}

// EnterOffset_limit_clause is called when production offset_limit_clause is entered.
func (s *BaseCosmosDBParserListener) EnterOffset_limit_clause(ctx *Offset_limit_clauseContext) {}

// ExitOffset_limit_clause is called when production offset_limit_clause is exited.
func (s *BaseCosmosDBParserListener) ExitOffset_limit_clause(ctx *Offset_limit_clauseContext) {}

// EnterFrom_specification is called when production from_specification is entered.
func (s *BaseCosmosDBParserListener) EnterFrom_specification(ctx *From_specificationContext) {}

// ExitFrom_specification is called when production from_specification is exited.
func (s *BaseCosmosDBParserListener) ExitFrom_specification(ctx *From_specificationContext) {}

// EnterFrom_source is called when production from_source is entered.
func (s *BaseCosmosDBParserListener) EnterFrom_source(ctx *From_sourceContext) {}

// ExitFrom_source is called when production from_source is exited.
func (s *BaseCosmosDBParserListener) ExitFrom_source(ctx *From_sourceContext) {}

// EnterContainer_expression is called when production container_expression is entered.
func (s *BaseCosmosDBParserListener) EnterContainer_expression(ctx *Container_expressionContext) {}

// ExitContainer_expression is called when production container_expression is exited.
func (s *BaseCosmosDBParserListener) ExitContainer_expression(ctx *Container_expressionContext) {}

// EnterJoin_clause is called when production join_clause is entered.
func (s *BaseCosmosDBParserListener) EnterJoin_clause(ctx *Join_clauseContext) {}

// ExitJoin_clause is called when production join_clause is exited.
func (s *BaseCosmosDBParserListener) ExitJoin_clause(ctx *Join_clauseContext) {}

// EnterContainer_name is called when production container_name is entered.
func (s *BaseCosmosDBParserListener) EnterContainer_name(ctx *Container_nameContext) {}

// ExitContainer_name is called when production container_name is exited.
func (s *BaseCosmosDBParserListener) ExitContainer_name(ctx *Container_nameContext) {}

// EnterObject_property_list is called when production object_property_list is entered.
func (s *BaseCosmosDBParserListener) EnterObject_property_list(ctx *Object_property_listContext) {}

// ExitObject_property_list is called when production object_property_list is exited.
func (s *BaseCosmosDBParserListener) ExitObject_property_list(ctx *Object_property_listContext) {}

// EnterObject_property is called when production object_property is entered.
func (s *BaseCosmosDBParserListener) EnterObject_property(ctx *Object_propertyContext) {}

// ExitObject_property is called when production object_property is exited.
func (s *BaseCosmosDBParserListener) ExitObject_property(ctx *Object_propertyContext) {}

// EnterProperty_alias is called when production property_alias is entered.
func (s *BaseCosmosDBParserListener) EnterProperty_alias(ctx *Property_aliasContext) {}

// ExitProperty_alias is called when production property_alias is exited.
func (s *BaseCosmosDBParserListener) ExitProperty_alias(ctx *Property_aliasContext) {}

// EnterScalar_expression is called when production scalar_expression is entered.
func (s *BaseCosmosDBParserListener) EnterScalar_expression(ctx *Scalar_expressionContext) {}

// ExitScalar_expression is called when production scalar_expression is exited.
func (s *BaseCosmosDBParserListener) ExitScalar_expression(ctx *Scalar_expressionContext) {}

// EnterCreate_array_expression is called when production create_array_expression is entered.
func (s *BaseCosmosDBParserListener) EnterCreate_array_expression(ctx *Create_array_expressionContext) {
}

// ExitCreate_array_expression is called when production create_array_expression is exited.
func (s *BaseCosmosDBParserListener) ExitCreate_array_expression(ctx *Create_array_expressionContext) {
}

// EnterCreate_object_expression is called when production create_object_expression is entered.
func (s *BaseCosmosDBParserListener) EnterCreate_object_expression(ctx *Create_object_expressionContext) {
}

// ExitCreate_object_expression is called when production create_object_expression is exited.
func (s *BaseCosmosDBParserListener) ExitCreate_object_expression(ctx *Create_object_expressionContext) {
}

// EnterObject_field_pair is called when production object_field_pair is entered.
func (s *BaseCosmosDBParserListener) EnterObject_field_pair(ctx *Object_field_pairContext) {}

// ExitObject_field_pair is called when production object_field_pair is exited.
func (s *BaseCosmosDBParserListener) ExitObject_field_pair(ctx *Object_field_pairContext) {}

// EnterScalar_function_expression is called when production scalar_function_expression is entered.
func (s *BaseCosmosDBParserListener) EnterScalar_function_expression(ctx *Scalar_function_expressionContext) {
}

// ExitScalar_function_expression is called when production scalar_function_expression is exited.
func (s *BaseCosmosDBParserListener) ExitScalar_function_expression(ctx *Scalar_function_expressionContext) {
}

// EnterUdf_scalar_function_expression is called when production udf_scalar_function_expression is entered.
func (s *BaseCosmosDBParserListener) EnterUdf_scalar_function_expression(ctx *Udf_scalar_function_expressionContext) {
}

// ExitUdf_scalar_function_expression is called when production udf_scalar_function_expression is exited.
func (s *BaseCosmosDBParserListener) ExitUdf_scalar_function_expression(ctx *Udf_scalar_function_expressionContext) {
}

// EnterBuiltin_function_expression is called when production builtin_function_expression is entered.
func (s *BaseCosmosDBParserListener) EnterBuiltin_function_expression(ctx *Builtin_function_expressionContext) {
}

// ExitBuiltin_function_expression is called when production builtin_function_expression is exited.
func (s *BaseCosmosDBParserListener) ExitBuiltin_function_expression(ctx *Builtin_function_expressionContext) {
}

// EnterMultiplicative_operator is called when production multiplicative_operator is entered.
func (s *BaseCosmosDBParserListener) EnterMultiplicative_operator(ctx *Multiplicative_operatorContext) {
}

// ExitMultiplicative_operator is called when production multiplicative_operator is exited.
func (s *BaseCosmosDBParserListener) ExitMultiplicative_operator(ctx *Multiplicative_operatorContext) {
}

// EnterAdditive_operator is called when production additive_operator is entered.
func (s *BaseCosmosDBParserListener) EnterAdditive_operator(ctx *Additive_operatorContext) {}

// ExitAdditive_operator is called when production additive_operator is exited.
func (s *BaseCosmosDBParserListener) ExitAdditive_operator(ctx *Additive_operatorContext) {}

// EnterShift_operator is called when production shift_operator is entered.
func (s *BaseCosmosDBParserListener) EnterShift_operator(ctx *Shift_operatorContext) {}

// ExitShift_operator is called when production shift_operator is exited.
func (s *BaseCosmosDBParserListener) ExitShift_operator(ctx *Shift_operatorContext) {}

// EnterComparison_operator is called when production comparison_operator is entered.
func (s *BaseCosmosDBParserListener) EnterComparison_operator(ctx *Comparison_operatorContext) {}

// ExitComparison_operator is called when production comparison_operator is exited.
func (s *BaseCosmosDBParserListener) ExitComparison_operator(ctx *Comparison_operatorContext) {}

// EnterUnary_operator is called when production unary_operator is entered.
func (s *BaseCosmosDBParserListener) EnterUnary_operator(ctx *Unary_operatorContext) {}

// ExitUnary_operator is called when production unary_operator is exited.
func (s *BaseCosmosDBParserListener) ExitUnary_operator(ctx *Unary_operatorContext) {}

// EnterParameter_name is called when production parameter_name is entered.
func (s *BaseCosmosDBParserListener) EnterParameter_name(ctx *Parameter_nameContext) {}

// ExitParameter_name is called when production parameter_name is exited.
func (s *BaseCosmosDBParserListener) ExitParameter_name(ctx *Parameter_nameContext) {}

// EnterConstant is called when production constant is entered.
func (s *BaseCosmosDBParserListener) EnterConstant(ctx *ConstantContext) {}

// ExitConstant is called when production constant is exited.
func (s *BaseCosmosDBParserListener) ExitConstant(ctx *ConstantContext) {}

// EnterUndefined_constant is called when production undefined_constant is entered.
func (s *BaseCosmosDBParserListener) EnterUndefined_constant(ctx *Undefined_constantContext) {}

// ExitUndefined_constant is called when production undefined_constant is exited.
func (s *BaseCosmosDBParserListener) ExitUndefined_constant(ctx *Undefined_constantContext) {}

// EnterNull_constant is called when production null_constant is entered.
func (s *BaseCosmosDBParserListener) EnterNull_constant(ctx *Null_constantContext) {}

// ExitNull_constant is called when production null_constant is exited.
func (s *BaseCosmosDBParserListener) ExitNull_constant(ctx *Null_constantContext) {}

// EnterBoolean_constant is called when production boolean_constant is entered.
func (s *BaseCosmosDBParserListener) EnterBoolean_constant(ctx *Boolean_constantContext) {}

// ExitBoolean_constant is called when production boolean_constant is exited.
func (s *BaseCosmosDBParserListener) ExitBoolean_constant(ctx *Boolean_constantContext) {}

// EnterNumber_constant is called when production number_constant is entered.
func (s *BaseCosmosDBParserListener) EnterNumber_constant(ctx *Number_constantContext) {}

// ExitNumber_constant is called when production number_constant is exited.
func (s *BaseCosmosDBParserListener) ExitNumber_constant(ctx *Number_constantContext) {}

// EnterString_constant is called when production string_constant is entered.
func (s *BaseCosmosDBParserListener) EnterString_constant(ctx *String_constantContext) {}

// ExitString_constant is called when production string_constant is exited.
func (s *BaseCosmosDBParserListener) ExitString_constant(ctx *String_constantContext) {}

// EnterString_literal is called when production string_literal is entered.
func (s *BaseCosmosDBParserListener) EnterString_literal(ctx *String_literalContext) {}

// ExitString_literal is called when production string_literal is exited.
func (s *BaseCosmosDBParserListener) ExitString_literal(ctx *String_literalContext) {}

// EnterDecimal_literal is called when production decimal_literal is entered.
func (s *BaseCosmosDBParserListener) EnterDecimal_literal(ctx *Decimal_literalContext) {}

// ExitDecimal_literal is called when production decimal_literal is exited.
func (s *BaseCosmosDBParserListener) ExitDecimal_literal(ctx *Decimal_literalContext) {}

// EnterHexadecimal_literal is called when production hexadecimal_literal is entered.
func (s *BaseCosmosDBParserListener) EnterHexadecimal_literal(ctx *Hexadecimal_literalContext) {}

// ExitHexadecimal_literal is called when production hexadecimal_literal is exited.
func (s *BaseCosmosDBParserListener) ExitHexadecimal_literal(ctx *Hexadecimal_literalContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseCosmosDBParserListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseCosmosDBParserListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterProperty_name is called when production property_name is entered.
func (s *BaseCosmosDBParserListener) EnterProperty_name(ctx *Property_nameContext) {}

// ExitProperty_name is called when production property_name is exited.
func (s *BaseCosmosDBParserListener) ExitProperty_name(ctx *Property_nameContext) {}

// EnterArray_index is called when production array_index is entered.
func (s *BaseCosmosDBParserListener) EnterArray_index(ctx *Array_indexContext) {}

// ExitArray_index is called when production array_index is exited.
func (s *BaseCosmosDBParserListener) ExitArray_index(ctx *Array_indexContext) {}

// EnterInput_alias is called when production input_alias is entered.
func (s *BaseCosmosDBParserListener) EnterInput_alias(ctx *Input_aliasContext) {}

// ExitInput_alias is called when production input_alias is exited.
func (s *BaseCosmosDBParserListener) ExitInput_alias(ctx *Input_aliasContext) {}
