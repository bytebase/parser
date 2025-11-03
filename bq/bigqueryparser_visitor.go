// Code generated from BigQueryParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package bq // BigQueryParser
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by BigQueryParser.
type BigQueryParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by BigQueryParser#root.
	VisitRoot(ctx *RootContext) interface{}

	// Visit a parse tree produced by BigQueryParser#stmtmulti.
	VisitStmtmulti(ctx *StmtmultiContext) interface{}

	// Visit a parse tree produced by BigQueryParser#stmt.
	VisitStmt(ctx *StmtContext) interface{}

	// Visit a parse tree produced by BigQueryParser#query_statement.
	VisitQuery_statement(ctx *Query_statementContext) interface{}

	// Visit a parse tree produced by BigQueryParser#query_expr.
	VisitQuery_expr(ctx *Query_exprContext) interface{}

	// Visit a parse tree produced by BigQueryParser#select_statement.
	VisitSelect_statement(ctx *Select_statementContext) interface{}

	// Visit a parse tree produced by BigQueryParser#from_statement.
	VisitFrom_statement(ctx *From_statementContext) interface{}

	// Visit a parse tree produced by BigQueryParser#from_item.
	VisitFrom_item(ctx *From_itemContext) interface{}

	// Visit a parse tree produced by BigQueryParser#where_statement.
	VisitWhere_statement(ctx *Where_statementContext) interface{}

	// Visit a parse tree produced by BigQueryParser#group_statement.
	VisitGroup_statement(ctx *Group_statementContext) interface{}

	// Visit a parse tree produced by BigQueryParser#having_statement.
	VisitHaving_statement(ctx *Having_statementContext) interface{}

	// Visit a parse tree produced by BigQueryParser#window_statement.
	VisitWindow_statement(ctx *Window_statementContext) interface{}

	// Visit a parse tree produced by BigQueryParser#order_clause.
	VisitOrder_clause(ctx *Order_clauseContext) interface{}

	// Visit a parse tree produced by BigQueryParser#limit_clause.
	VisitLimit_clause(ctx *Limit_clauseContext) interface{}

	// Visit a parse tree produced by BigQueryParser#unary_operator.
	VisitUnary_operator(ctx *Unary_operatorContext) interface{}

	// Visit a parse tree produced by BigQueryParser#expr.
	VisitExpr(ctx *ExprContext) interface{}

	// Visit a parse tree produced by BigQueryParser#cast_expr.
	VisitCast_expr(ctx *Cast_exprContext) interface{}

	// Visit a parse tree produced by BigQueryParser#column_expr.
	VisitColumn_expr(ctx *Column_exprContext) interface{}

	// Visit a parse tree produced by BigQueryParser#except_statement.
	VisitExcept_statement(ctx *Except_statementContext) interface{}

	// Visit a parse tree produced by BigQueryParser#replace_statement.
	VisitReplace_statement(ctx *Replace_statementContext) interface{}

	// Visit a parse tree produced by BigQueryParser#join_type.
	VisitJoin_type(ctx *Join_typeContext) interface{}

	// Visit a parse tree produced by BigQueryParser#on_clause.
	VisitOn_clause(ctx *On_clauseContext) interface{}

	// Visit a parse tree produced by BigQueryParser#set_op.
	VisitSet_op(ctx *Set_opContext) interface{}

	// Visit a parse tree produced by BigQueryParser#using_clause.
	VisitUsing_clause(ctx *Using_clauseContext) interface{}

	// Visit a parse tree produced by BigQueryParser#field_path.
	VisitField_path(ctx *Field_pathContext) interface{}

	// Visit a parse tree produced by BigQueryParser#sstruct.
	VisitSstruct(ctx *SstructContext) interface{}

	// Visit a parse tree produced by BigQueryParser#array_expr.
	VisitArray_expr(ctx *Array_exprContext) interface{}

	// Visit a parse tree produced by BigQueryParser#array_path.
	VisitArray_path(ctx *Array_pathContext) interface{}

	// Visit a parse tree produced by BigQueryParser#bool_expression.
	VisitBool_expression(ctx *Bool_expressionContext) interface{}

	// Visit a parse tree produced by BigQueryParser#window_name.
	VisitWindow_name(ctx *Window_nameContext) interface{}

	// Visit a parse tree produced by BigQueryParser#window_definition.
	VisitWindow_definition(ctx *Window_definitionContext) interface{}

	// Visit a parse tree produced by BigQueryParser#count.
	VisitCount(ctx *CountContext) interface{}

	// Visit a parse tree produced by BigQueryParser#skip_rows.
	VisitSkip_rows(ctx *Skip_rowsContext) interface{}

	// Visit a parse tree produced by BigQueryParser#with_statement.
	VisitWith_statement(ctx *With_statementContext) interface{}

	// Visit a parse tree produced by BigQueryParser#name.
	VisitName(ctx *NameContext) interface{}

	// Visit a parse tree produced by BigQueryParser#dot_name.
	VisitDot_name(ctx *Dot_nameContext) interface{}

	// Visit a parse tree produced by BigQueryParser#alias_name.
	VisitAlias_name(ctx *Alias_nameContext) interface{}

	// Visit a parse tree produced by BigQueryParser#array_name.
	VisitArray_name(ctx *Array_nameContext) interface{}

	// Visit a parse tree produced by BigQueryParser#column_name.
	VisitColumn_name(ctx *Column_nameContext) interface{}

	// Visit a parse tree produced by BigQueryParser#cte_name.
	VisitCte_name(ctx *Cte_nameContext) interface{}

	// Visit a parse tree produced by BigQueryParser#dataset_name.
	VisitDataset_name(ctx *Dataset_nameContext) interface{}

	// Visit a parse tree produced by BigQueryParser#datatype_name.
	VisitDatatype_name(ctx *Datatype_nameContext) interface{}

	// Visit a parse tree produced by BigQueryParser#function_name.
	VisitFunction_name(ctx *Function_nameContext) interface{}

	// Visit a parse tree produced by BigQueryParser#join_name.
	VisitJoin_name(ctx *Join_nameContext) interface{}

	// Visit a parse tree produced by BigQueryParser#member_name.
	VisitMember_name(ctx *Member_nameContext) interface{}

	// Visit a parse tree produced by BigQueryParser#project_name.
	VisitProject_name(ctx *Project_nameContext) interface{}

	// Visit a parse tree produced by BigQueryParser#struct_name.
	VisitStruct_name(ctx *Struct_nameContext) interface{}

	// Visit a parse tree produced by BigQueryParser#table_name.
	VisitTable_name(ctx *Table_nameContext) interface{}

	// Visit a parse tree produced by BigQueryParser#table_expr.
	VisitTable_expr(ctx *Table_exprContext) interface{}

	// Visit a parse tree produced by BigQueryParser#number.
	VisitNumber(ctx *NumberContext) interface{}

	// Visit a parse tree produced by BigQueryParser#integer_type.
	VisitInteger_type(ctx *Integer_typeContext) interface{}

	// Visit a parse tree produced by BigQueryParser#float_type.
	VisitFloat_type(ctx *Float_typeContext) interface{}

	// Visit a parse tree produced by BigQueryParser#string.
	VisitString(ctx *StringContext) interface{}

	// Visit a parse tree produced by BigQueryParser#quoted_string.
	VisitQuoted_string(ctx *Quoted_stringContext) interface{}

	// Visit a parse tree produced by BigQueryParser#triple_quoted_string.
	VisitTriple_quoted_string(ctx *Triple_quoted_stringContext) interface{}

	// Visit a parse tree produced by BigQueryParser#raw_string.
	VisitRaw_string(ctx *Raw_stringContext) interface{}

	// Visit a parse tree produced by BigQueryParser#byte_string.
	VisitByte_string(ctx *Byte_stringContext) interface{}

	// Visit a parse tree produced by BigQueryParser#raw_byte_string.
	VisitRaw_byte_string(ctx *Raw_byte_stringContext) interface{}

	// Visit a parse tree produced by BigQueryParser#special_string.
	VisitSpecial_string(ctx *Special_stringContext) interface{}

	// Visit a parse tree produced by BigQueryParser#keyword.
	VisitKeyword(ctx *KeywordContext) interface{}
}
