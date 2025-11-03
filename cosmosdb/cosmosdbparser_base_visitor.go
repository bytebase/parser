// Code generated from CosmosDBParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package cosmosdb // CosmosDBParser
import "github.com/antlr4-go/antlr/v4"

type BaseCosmosDBParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseCosmosDBParserVisitor) VisitRoot(ctx *RootContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCosmosDBParserVisitor) VisitSelect(ctx *SelectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCosmosDBParserVisitor) VisitSelect_clause(ctx *Select_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCosmosDBParserVisitor) VisitSelect_specification(ctx *Select_specificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCosmosDBParserVisitor) VisitFrom_clause(ctx *From_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCosmosDBParserVisitor) VisitWhere_clause(ctx *Where_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCosmosDBParserVisitor) VisitFrom_specification(ctx *From_specificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCosmosDBParserVisitor) VisitFrom_source(ctx *From_sourceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCosmosDBParserVisitor) VisitContainer_expression(ctx *Container_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCosmosDBParserVisitor) VisitContainer_name(ctx *Container_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCosmosDBParserVisitor) VisitObject_property_list(ctx *Object_property_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCosmosDBParserVisitor) VisitObject_property(ctx *Object_propertyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCosmosDBParserVisitor) VisitProperty_alias(ctx *Property_aliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCosmosDBParserVisitor) VisitScalar_expression(ctx *Scalar_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCosmosDBParserVisitor) VisitScalar_expression_in_where(ctx *Scalar_expression_in_whereContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCosmosDBParserVisitor) VisitCreate_array_expression(ctx *Create_array_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCosmosDBParserVisitor) VisitCreate_object_expression(ctx *Create_object_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCosmosDBParserVisitor) VisitScalar_function_expression(ctx *Scalar_function_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCosmosDBParserVisitor) VisitUdf_scalar_function_expression(ctx *Udf_scalar_function_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCosmosDBParserVisitor) VisitBuiltin_function_expression(ctx *Builtin_function_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCosmosDBParserVisitor) VisitBinary_operator(ctx *Binary_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCosmosDBParserVisitor) VisitUnary_operator(ctx *Unary_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCosmosDBParserVisitor) VisitParameter_name(ctx *Parameter_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCosmosDBParserVisitor) VisitConstant(ctx *ConstantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCosmosDBParserVisitor) VisitObject_constant(ctx *Object_constantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCosmosDBParserVisitor) VisitObject_constant_field_pair(ctx *Object_constant_field_pairContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCosmosDBParserVisitor) VisitArray_constant(ctx *Array_constantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCosmosDBParserVisitor) VisitString_constant(ctx *String_constantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCosmosDBParserVisitor) VisitUndefined_constant(ctx *Undefined_constantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCosmosDBParserVisitor) VisitNull_constant(ctx *Null_constantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCosmosDBParserVisitor) VisitBoolean_constant(ctx *Boolean_constantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCosmosDBParserVisitor) VisitNumber_constant(ctx *Number_constantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCosmosDBParserVisitor) VisitString_literal(ctx *String_literalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCosmosDBParserVisitor) VisitDecimal_literal(ctx *Decimal_literalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCosmosDBParserVisitor) VisitHexadecimal_literal(ctx *Hexadecimal_literalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCosmosDBParserVisitor) VisitProperty_name(ctx *Property_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCosmosDBParserVisitor) VisitArray_index(ctx *Array_indexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseCosmosDBParserVisitor) VisitInput_alias(ctx *Input_aliasContext) interface{} {
	return v.VisitChildren(ctx)
}
