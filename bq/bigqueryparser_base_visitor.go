// Code generated from BigQueryParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package bq // BigQueryParser
import "github.com/antlr4-go/antlr/v4"

type BaseBigQueryParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseBigQueryParserVisitor) VisitRoot(ctx *RootContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigQueryParserVisitor) VisitStmtmulti(ctx *StmtmultiContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigQueryParserVisitor) VisitStmt(ctx *StmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigQueryParserVisitor) VisitQuery_statement(ctx *Query_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigQueryParserVisitor) VisitQuery_expr(ctx *Query_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigQueryParserVisitor) VisitSelect_statement(ctx *Select_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigQueryParserVisitor) VisitFrom_statement(ctx *From_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigQueryParserVisitor) VisitFrom_item(ctx *From_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigQueryParserVisitor) VisitWhere_statement(ctx *Where_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigQueryParserVisitor) VisitGroup_statement(ctx *Group_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigQueryParserVisitor) VisitHaving_statement(ctx *Having_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigQueryParserVisitor) VisitWindow_statement(ctx *Window_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigQueryParserVisitor) VisitOrder_clause(ctx *Order_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigQueryParserVisitor) VisitLimit_clause(ctx *Limit_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigQueryParserVisitor) VisitUnary_operator(ctx *Unary_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigQueryParserVisitor) VisitExpr(ctx *ExprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigQueryParserVisitor) VisitCast_expr(ctx *Cast_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigQueryParserVisitor) VisitColumn_expr(ctx *Column_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigQueryParserVisitor) VisitExcept_statement(ctx *Except_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigQueryParserVisitor) VisitReplace_statement(ctx *Replace_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigQueryParserVisitor) VisitJoin_type(ctx *Join_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigQueryParserVisitor) VisitOn_clause(ctx *On_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigQueryParserVisitor) VisitSet_op(ctx *Set_opContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigQueryParserVisitor) VisitUsing_clause(ctx *Using_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigQueryParserVisitor) VisitField_path(ctx *Field_pathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigQueryParserVisitor) VisitSstruct(ctx *SstructContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigQueryParserVisitor) VisitArray_expr(ctx *Array_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigQueryParserVisitor) VisitArray_path(ctx *Array_pathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigQueryParserVisitor) VisitBool_expression(ctx *Bool_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigQueryParserVisitor) VisitWindow_name(ctx *Window_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigQueryParserVisitor) VisitWindow_definition(ctx *Window_definitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigQueryParserVisitor) VisitCount(ctx *CountContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigQueryParserVisitor) VisitSkip_rows(ctx *Skip_rowsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigQueryParserVisitor) VisitWith_statement(ctx *With_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigQueryParserVisitor) VisitName(ctx *NameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigQueryParserVisitor) VisitDot_name(ctx *Dot_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigQueryParserVisitor) VisitAlias_name(ctx *Alias_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigQueryParserVisitor) VisitArray_name(ctx *Array_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigQueryParserVisitor) VisitColumn_name(ctx *Column_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigQueryParserVisitor) VisitCte_name(ctx *Cte_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigQueryParserVisitor) VisitDataset_name(ctx *Dataset_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigQueryParserVisitor) VisitDatatype_name(ctx *Datatype_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigQueryParserVisitor) VisitFunction_name(ctx *Function_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigQueryParserVisitor) VisitJoin_name(ctx *Join_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigQueryParserVisitor) VisitMember_name(ctx *Member_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigQueryParserVisitor) VisitProject_name(ctx *Project_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigQueryParserVisitor) VisitStruct_name(ctx *Struct_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigQueryParserVisitor) VisitTable_name(ctx *Table_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigQueryParserVisitor) VisitTable_expr(ctx *Table_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigQueryParserVisitor) VisitNumber(ctx *NumberContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigQueryParserVisitor) VisitInteger_type(ctx *Integer_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigQueryParserVisitor) VisitFloat_type(ctx *Float_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigQueryParserVisitor) VisitString(ctx *StringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigQueryParserVisitor) VisitQuoted_string(ctx *Quoted_stringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigQueryParserVisitor) VisitTriple_quoted_string(ctx *Triple_quoted_stringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigQueryParserVisitor) VisitRaw_string(ctx *Raw_stringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigQueryParserVisitor) VisitByte_string(ctx *Byte_stringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigQueryParserVisitor) VisitRaw_byte_string(ctx *Raw_byte_stringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigQueryParserVisitor) VisitSpecial_string(ctx *Special_stringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseBigQueryParserVisitor) VisitKeyword(ctx *KeywordContext) interface{} {
	return v.VisitChildren(ctx)
}
