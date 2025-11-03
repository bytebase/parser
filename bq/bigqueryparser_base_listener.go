// Code generated from BigQueryParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package bq // BigQueryParser
import "github.com/antlr4-go/antlr/v4"

// BaseBigQueryParserListener is a complete listener for a parse tree produced by BigQueryParser.
type BaseBigQueryParserListener struct{}

var _ BigQueryParserListener = &BaseBigQueryParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseBigQueryParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseBigQueryParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseBigQueryParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseBigQueryParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterRoot is called when production root is entered.
func (s *BaseBigQueryParserListener) EnterRoot(ctx *RootContext) {}

// ExitRoot is called when production root is exited.
func (s *BaseBigQueryParserListener) ExitRoot(ctx *RootContext) {}

// EnterStmtmulti is called when production stmtmulti is entered.
func (s *BaseBigQueryParserListener) EnterStmtmulti(ctx *StmtmultiContext) {}

// ExitStmtmulti is called when production stmtmulti is exited.
func (s *BaseBigQueryParserListener) ExitStmtmulti(ctx *StmtmultiContext) {}

// EnterStmt is called when production stmt is entered.
func (s *BaseBigQueryParserListener) EnterStmt(ctx *StmtContext) {}

// ExitStmt is called when production stmt is exited.
func (s *BaseBigQueryParserListener) ExitStmt(ctx *StmtContext) {}

// EnterQuery_statement is called when production query_statement is entered.
func (s *BaseBigQueryParserListener) EnterQuery_statement(ctx *Query_statementContext) {}

// ExitQuery_statement is called when production query_statement is exited.
func (s *BaseBigQueryParserListener) ExitQuery_statement(ctx *Query_statementContext) {}

// EnterQuery_expr is called when production query_expr is entered.
func (s *BaseBigQueryParserListener) EnterQuery_expr(ctx *Query_exprContext) {}

// ExitQuery_expr is called when production query_expr is exited.
func (s *BaseBigQueryParserListener) ExitQuery_expr(ctx *Query_exprContext) {}

// EnterSelect_statement is called when production select_statement is entered.
func (s *BaseBigQueryParserListener) EnterSelect_statement(ctx *Select_statementContext) {}

// ExitSelect_statement is called when production select_statement is exited.
func (s *BaseBigQueryParserListener) ExitSelect_statement(ctx *Select_statementContext) {}

// EnterFrom_statement is called when production from_statement is entered.
func (s *BaseBigQueryParserListener) EnterFrom_statement(ctx *From_statementContext) {}

// ExitFrom_statement is called when production from_statement is exited.
func (s *BaseBigQueryParserListener) ExitFrom_statement(ctx *From_statementContext) {}

// EnterFrom_item is called when production from_item is entered.
func (s *BaseBigQueryParserListener) EnterFrom_item(ctx *From_itemContext) {}

// ExitFrom_item is called when production from_item is exited.
func (s *BaseBigQueryParserListener) ExitFrom_item(ctx *From_itemContext) {}

// EnterWhere_statement is called when production where_statement is entered.
func (s *BaseBigQueryParserListener) EnterWhere_statement(ctx *Where_statementContext) {}

// ExitWhere_statement is called when production where_statement is exited.
func (s *BaseBigQueryParserListener) ExitWhere_statement(ctx *Where_statementContext) {}

// EnterGroup_statement is called when production group_statement is entered.
func (s *BaseBigQueryParserListener) EnterGroup_statement(ctx *Group_statementContext) {}

// ExitGroup_statement is called when production group_statement is exited.
func (s *BaseBigQueryParserListener) ExitGroup_statement(ctx *Group_statementContext) {}

// EnterHaving_statement is called when production having_statement is entered.
func (s *BaseBigQueryParserListener) EnterHaving_statement(ctx *Having_statementContext) {}

// ExitHaving_statement is called when production having_statement is exited.
func (s *BaseBigQueryParserListener) ExitHaving_statement(ctx *Having_statementContext) {}

// EnterWindow_statement is called when production window_statement is entered.
func (s *BaseBigQueryParserListener) EnterWindow_statement(ctx *Window_statementContext) {}

// ExitWindow_statement is called when production window_statement is exited.
func (s *BaseBigQueryParserListener) ExitWindow_statement(ctx *Window_statementContext) {}

// EnterOrder_clause is called when production order_clause is entered.
func (s *BaseBigQueryParserListener) EnterOrder_clause(ctx *Order_clauseContext) {}

// ExitOrder_clause is called when production order_clause is exited.
func (s *BaseBigQueryParserListener) ExitOrder_clause(ctx *Order_clauseContext) {}

// EnterLimit_clause is called when production limit_clause is entered.
func (s *BaseBigQueryParserListener) EnterLimit_clause(ctx *Limit_clauseContext) {}

// ExitLimit_clause is called when production limit_clause is exited.
func (s *BaseBigQueryParserListener) ExitLimit_clause(ctx *Limit_clauseContext) {}

// EnterUnary_operator is called when production unary_operator is entered.
func (s *BaseBigQueryParserListener) EnterUnary_operator(ctx *Unary_operatorContext) {}

// ExitUnary_operator is called when production unary_operator is exited.
func (s *BaseBigQueryParserListener) ExitUnary_operator(ctx *Unary_operatorContext) {}

// EnterExpr is called when production expr is entered.
func (s *BaseBigQueryParserListener) EnterExpr(ctx *ExprContext) {}

// ExitExpr is called when production expr is exited.
func (s *BaseBigQueryParserListener) ExitExpr(ctx *ExprContext) {}

// EnterCast_expr is called when production cast_expr is entered.
func (s *BaseBigQueryParserListener) EnterCast_expr(ctx *Cast_exprContext) {}

// ExitCast_expr is called when production cast_expr is exited.
func (s *BaseBigQueryParserListener) ExitCast_expr(ctx *Cast_exprContext) {}

// EnterColumn_expr is called when production column_expr is entered.
func (s *BaseBigQueryParserListener) EnterColumn_expr(ctx *Column_exprContext) {}

// ExitColumn_expr is called when production column_expr is exited.
func (s *BaseBigQueryParserListener) ExitColumn_expr(ctx *Column_exprContext) {}

// EnterExcept_statement is called when production except_statement is entered.
func (s *BaseBigQueryParserListener) EnterExcept_statement(ctx *Except_statementContext) {}

// ExitExcept_statement is called when production except_statement is exited.
func (s *BaseBigQueryParserListener) ExitExcept_statement(ctx *Except_statementContext) {}

// EnterReplace_statement is called when production replace_statement is entered.
func (s *BaseBigQueryParserListener) EnterReplace_statement(ctx *Replace_statementContext) {}

// ExitReplace_statement is called when production replace_statement is exited.
func (s *BaseBigQueryParserListener) ExitReplace_statement(ctx *Replace_statementContext) {}

// EnterJoin_type is called when production join_type is entered.
func (s *BaseBigQueryParserListener) EnterJoin_type(ctx *Join_typeContext) {}

// ExitJoin_type is called when production join_type is exited.
func (s *BaseBigQueryParserListener) ExitJoin_type(ctx *Join_typeContext) {}

// EnterOn_clause is called when production on_clause is entered.
func (s *BaseBigQueryParserListener) EnterOn_clause(ctx *On_clauseContext) {}

// ExitOn_clause is called when production on_clause is exited.
func (s *BaseBigQueryParserListener) ExitOn_clause(ctx *On_clauseContext) {}

// EnterSet_op is called when production set_op is entered.
func (s *BaseBigQueryParserListener) EnterSet_op(ctx *Set_opContext) {}

// ExitSet_op is called when production set_op is exited.
func (s *BaseBigQueryParserListener) ExitSet_op(ctx *Set_opContext) {}

// EnterUsing_clause is called when production using_clause is entered.
func (s *BaseBigQueryParserListener) EnterUsing_clause(ctx *Using_clauseContext) {}

// ExitUsing_clause is called when production using_clause is exited.
func (s *BaseBigQueryParserListener) ExitUsing_clause(ctx *Using_clauseContext) {}

// EnterField_path is called when production field_path is entered.
func (s *BaseBigQueryParserListener) EnterField_path(ctx *Field_pathContext) {}

// ExitField_path is called when production field_path is exited.
func (s *BaseBigQueryParserListener) ExitField_path(ctx *Field_pathContext) {}

// EnterSstruct is called when production sstruct is entered.
func (s *BaseBigQueryParserListener) EnterSstruct(ctx *SstructContext) {}

// ExitSstruct is called when production sstruct is exited.
func (s *BaseBigQueryParserListener) ExitSstruct(ctx *SstructContext) {}

// EnterArray_expr is called when production array_expr is entered.
func (s *BaseBigQueryParserListener) EnterArray_expr(ctx *Array_exprContext) {}

// ExitArray_expr is called when production array_expr is exited.
func (s *BaseBigQueryParserListener) ExitArray_expr(ctx *Array_exprContext) {}

// EnterArray_path is called when production array_path is entered.
func (s *BaseBigQueryParserListener) EnterArray_path(ctx *Array_pathContext) {}

// ExitArray_path is called when production array_path is exited.
func (s *BaseBigQueryParserListener) ExitArray_path(ctx *Array_pathContext) {}

// EnterBool_expression is called when production bool_expression is entered.
func (s *BaseBigQueryParserListener) EnterBool_expression(ctx *Bool_expressionContext) {}

// ExitBool_expression is called when production bool_expression is exited.
func (s *BaseBigQueryParserListener) ExitBool_expression(ctx *Bool_expressionContext) {}

// EnterWindow_name is called when production window_name is entered.
func (s *BaseBigQueryParserListener) EnterWindow_name(ctx *Window_nameContext) {}

// ExitWindow_name is called when production window_name is exited.
func (s *BaseBigQueryParserListener) ExitWindow_name(ctx *Window_nameContext) {}

// EnterWindow_definition is called when production window_definition is entered.
func (s *BaseBigQueryParserListener) EnterWindow_definition(ctx *Window_definitionContext) {}

// ExitWindow_definition is called when production window_definition is exited.
func (s *BaseBigQueryParserListener) ExitWindow_definition(ctx *Window_definitionContext) {}

// EnterCount is called when production count is entered.
func (s *BaseBigQueryParserListener) EnterCount(ctx *CountContext) {}

// ExitCount is called when production count is exited.
func (s *BaseBigQueryParserListener) ExitCount(ctx *CountContext) {}

// EnterSkip_rows is called when production skip_rows is entered.
func (s *BaseBigQueryParserListener) EnterSkip_rows(ctx *Skip_rowsContext) {}

// ExitSkip_rows is called when production skip_rows is exited.
func (s *BaseBigQueryParserListener) ExitSkip_rows(ctx *Skip_rowsContext) {}

// EnterWith_statement is called when production with_statement is entered.
func (s *BaseBigQueryParserListener) EnterWith_statement(ctx *With_statementContext) {}

// ExitWith_statement is called when production with_statement is exited.
func (s *BaseBigQueryParserListener) ExitWith_statement(ctx *With_statementContext) {}

// EnterName is called when production name is entered.
func (s *BaseBigQueryParserListener) EnterName(ctx *NameContext) {}

// ExitName is called when production name is exited.
func (s *BaseBigQueryParserListener) ExitName(ctx *NameContext) {}

// EnterDot_name is called when production dot_name is entered.
func (s *BaseBigQueryParserListener) EnterDot_name(ctx *Dot_nameContext) {}

// ExitDot_name is called when production dot_name is exited.
func (s *BaseBigQueryParserListener) ExitDot_name(ctx *Dot_nameContext) {}

// EnterAlias_name is called when production alias_name is entered.
func (s *BaseBigQueryParserListener) EnterAlias_name(ctx *Alias_nameContext) {}

// ExitAlias_name is called when production alias_name is exited.
func (s *BaseBigQueryParserListener) ExitAlias_name(ctx *Alias_nameContext) {}

// EnterArray_name is called when production array_name is entered.
func (s *BaseBigQueryParserListener) EnterArray_name(ctx *Array_nameContext) {}

// ExitArray_name is called when production array_name is exited.
func (s *BaseBigQueryParserListener) ExitArray_name(ctx *Array_nameContext) {}

// EnterColumn_name is called when production column_name is entered.
func (s *BaseBigQueryParserListener) EnterColumn_name(ctx *Column_nameContext) {}

// ExitColumn_name is called when production column_name is exited.
func (s *BaseBigQueryParserListener) ExitColumn_name(ctx *Column_nameContext) {}

// EnterCte_name is called when production cte_name is entered.
func (s *BaseBigQueryParserListener) EnterCte_name(ctx *Cte_nameContext) {}

// ExitCte_name is called when production cte_name is exited.
func (s *BaseBigQueryParserListener) ExitCte_name(ctx *Cte_nameContext) {}

// EnterDataset_name is called when production dataset_name is entered.
func (s *BaseBigQueryParserListener) EnterDataset_name(ctx *Dataset_nameContext) {}

// ExitDataset_name is called when production dataset_name is exited.
func (s *BaseBigQueryParserListener) ExitDataset_name(ctx *Dataset_nameContext) {}

// EnterDatatype_name is called when production datatype_name is entered.
func (s *BaseBigQueryParserListener) EnterDatatype_name(ctx *Datatype_nameContext) {}

// ExitDatatype_name is called when production datatype_name is exited.
func (s *BaseBigQueryParserListener) ExitDatatype_name(ctx *Datatype_nameContext) {}

// EnterFunction_name is called when production function_name is entered.
func (s *BaseBigQueryParserListener) EnterFunction_name(ctx *Function_nameContext) {}

// ExitFunction_name is called when production function_name is exited.
func (s *BaseBigQueryParserListener) ExitFunction_name(ctx *Function_nameContext) {}

// EnterJoin_name is called when production join_name is entered.
func (s *BaseBigQueryParserListener) EnterJoin_name(ctx *Join_nameContext) {}

// ExitJoin_name is called when production join_name is exited.
func (s *BaseBigQueryParserListener) ExitJoin_name(ctx *Join_nameContext) {}

// EnterMember_name is called when production member_name is entered.
func (s *BaseBigQueryParserListener) EnterMember_name(ctx *Member_nameContext) {}

// ExitMember_name is called when production member_name is exited.
func (s *BaseBigQueryParserListener) ExitMember_name(ctx *Member_nameContext) {}

// EnterProject_name is called when production project_name is entered.
func (s *BaseBigQueryParserListener) EnterProject_name(ctx *Project_nameContext) {}

// ExitProject_name is called when production project_name is exited.
func (s *BaseBigQueryParserListener) ExitProject_name(ctx *Project_nameContext) {}

// EnterStruct_name is called when production struct_name is entered.
func (s *BaseBigQueryParserListener) EnterStruct_name(ctx *Struct_nameContext) {}

// ExitStruct_name is called when production struct_name is exited.
func (s *BaseBigQueryParserListener) ExitStruct_name(ctx *Struct_nameContext) {}

// EnterTable_name is called when production table_name is entered.
func (s *BaseBigQueryParserListener) EnterTable_name(ctx *Table_nameContext) {}

// ExitTable_name is called when production table_name is exited.
func (s *BaseBigQueryParserListener) ExitTable_name(ctx *Table_nameContext) {}

// EnterTable_expr is called when production table_expr is entered.
func (s *BaseBigQueryParserListener) EnterTable_expr(ctx *Table_exprContext) {}

// ExitTable_expr is called when production table_expr is exited.
func (s *BaseBigQueryParserListener) ExitTable_expr(ctx *Table_exprContext) {}

// EnterNumber is called when production number is entered.
func (s *BaseBigQueryParserListener) EnterNumber(ctx *NumberContext) {}

// ExitNumber is called when production number is exited.
func (s *BaseBigQueryParserListener) ExitNumber(ctx *NumberContext) {}

// EnterInteger_type is called when production integer_type is entered.
func (s *BaseBigQueryParserListener) EnterInteger_type(ctx *Integer_typeContext) {}

// ExitInteger_type is called when production integer_type is exited.
func (s *BaseBigQueryParserListener) ExitInteger_type(ctx *Integer_typeContext) {}

// EnterFloat_type is called when production float_type is entered.
func (s *BaseBigQueryParserListener) EnterFloat_type(ctx *Float_typeContext) {}

// ExitFloat_type is called when production float_type is exited.
func (s *BaseBigQueryParserListener) ExitFloat_type(ctx *Float_typeContext) {}

// EnterString is called when production string is entered.
func (s *BaseBigQueryParserListener) EnterString(ctx *StringContext) {}

// ExitString is called when production string is exited.
func (s *BaseBigQueryParserListener) ExitString(ctx *StringContext) {}

// EnterQuoted_string is called when production quoted_string is entered.
func (s *BaseBigQueryParserListener) EnterQuoted_string(ctx *Quoted_stringContext) {}

// ExitQuoted_string is called when production quoted_string is exited.
func (s *BaseBigQueryParserListener) ExitQuoted_string(ctx *Quoted_stringContext) {}

// EnterTriple_quoted_string is called when production triple_quoted_string is entered.
func (s *BaseBigQueryParserListener) EnterTriple_quoted_string(ctx *Triple_quoted_stringContext) {}

// ExitTriple_quoted_string is called when production triple_quoted_string is exited.
func (s *BaseBigQueryParserListener) ExitTriple_quoted_string(ctx *Triple_quoted_stringContext) {}

// EnterRaw_string is called when production raw_string is entered.
func (s *BaseBigQueryParserListener) EnterRaw_string(ctx *Raw_stringContext) {}

// ExitRaw_string is called when production raw_string is exited.
func (s *BaseBigQueryParserListener) ExitRaw_string(ctx *Raw_stringContext) {}

// EnterByte_string is called when production byte_string is entered.
func (s *BaseBigQueryParserListener) EnterByte_string(ctx *Byte_stringContext) {}

// ExitByte_string is called when production byte_string is exited.
func (s *BaseBigQueryParserListener) ExitByte_string(ctx *Byte_stringContext) {}

// EnterRaw_byte_string is called when production raw_byte_string is entered.
func (s *BaseBigQueryParserListener) EnterRaw_byte_string(ctx *Raw_byte_stringContext) {}

// ExitRaw_byte_string is called when production raw_byte_string is exited.
func (s *BaseBigQueryParserListener) ExitRaw_byte_string(ctx *Raw_byte_stringContext) {}

// EnterSpecial_string is called when production special_string is entered.
func (s *BaseBigQueryParserListener) EnterSpecial_string(ctx *Special_stringContext) {}

// ExitSpecial_string is called when production special_string is exited.
func (s *BaseBigQueryParserListener) ExitSpecial_string(ctx *Special_stringContext) {}

// EnterKeyword is called when production keyword is entered.
func (s *BaseBigQueryParserListener) EnterKeyword(ctx *KeywordContext) {}

// ExitKeyword is called when production keyword is exited.
func (s *BaseBigQueryParserListener) ExitKeyword(ctx *KeywordContext) {}
