// Code generated from BigQueryParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package bq // BigQueryParser
import "github.com/antlr4-go/antlr/v4"

// BigQueryParserListener is a complete listener for a parse tree produced by BigQueryParser.
type BigQueryParserListener interface {
	antlr.ParseTreeListener

	// EnterRoot is called when entering the root production.
	EnterRoot(c *RootContext)

	// EnterStmtmulti is called when entering the stmtmulti production.
	EnterStmtmulti(c *StmtmultiContext)

	// EnterStmt is called when entering the stmt production.
	EnterStmt(c *StmtContext)

	// EnterQuery_statement is called when entering the query_statement production.
	EnterQuery_statement(c *Query_statementContext)

	// EnterQuery_expr is called when entering the query_expr production.
	EnterQuery_expr(c *Query_exprContext)

	// EnterSelect_statement is called when entering the select_statement production.
	EnterSelect_statement(c *Select_statementContext)

	// EnterFrom_statement is called when entering the from_statement production.
	EnterFrom_statement(c *From_statementContext)

	// EnterFrom_item is called when entering the from_item production.
	EnterFrom_item(c *From_itemContext)

	// EnterWhere_statement is called when entering the where_statement production.
	EnterWhere_statement(c *Where_statementContext)

	// EnterGroup_statement is called when entering the group_statement production.
	EnterGroup_statement(c *Group_statementContext)

	// EnterHaving_statement is called when entering the having_statement production.
	EnterHaving_statement(c *Having_statementContext)

	// EnterWindow_statement is called when entering the window_statement production.
	EnterWindow_statement(c *Window_statementContext)

	// EnterOrder_clause is called when entering the order_clause production.
	EnterOrder_clause(c *Order_clauseContext)

	// EnterLimit_clause is called when entering the limit_clause production.
	EnterLimit_clause(c *Limit_clauseContext)

	// EnterUnary_operator is called when entering the unary_operator production.
	EnterUnary_operator(c *Unary_operatorContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterCast_expr is called when entering the cast_expr production.
	EnterCast_expr(c *Cast_exprContext)

	// EnterColumn_expr is called when entering the column_expr production.
	EnterColumn_expr(c *Column_exprContext)

	// EnterExcept_statement is called when entering the except_statement production.
	EnterExcept_statement(c *Except_statementContext)

	// EnterReplace_statement is called when entering the replace_statement production.
	EnterReplace_statement(c *Replace_statementContext)

	// EnterJoin_type is called when entering the join_type production.
	EnterJoin_type(c *Join_typeContext)

	// EnterOn_clause is called when entering the on_clause production.
	EnterOn_clause(c *On_clauseContext)

	// EnterSet_op is called when entering the set_op production.
	EnterSet_op(c *Set_opContext)

	// EnterUsing_clause is called when entering the using_clause production.
	EnterUsing_clause(c *Using_clauseContext)

	// EnterField_path is called when entering the field_path production.
	EnterField_path(c *Field_pathContext)

	// EnterSstruct is called when entering the sstruct production.
	EnterSstruct(c *SstructContext)

	// EnterArray_expr is called when entering the array_expr production.
	EnterArray_expr(c *Array_exprContext)

	// EnterArray_path is called when entering the array_path production.
	EnterArray_path(c *Array_pathContext)

	// EnterBool_expression is called when entering the bool_expression production.
	EnterBool_expression(c *Bool_expressionContext)

	// EnterWindow_name is called when entering the window_name production.
	EnterWindow_name(c *Window_nameContext)

	// EnterWindow_definition is called when entering the window_definition production.
	EnterWindow_definition(c *Window_definitionContext)

	// EnterCount is called when entering the count production.
	EnterCount(c *CountContext)

	// EnterSkip_rows is called when entering the skip_rows production.
	EnterSkip_rows(c *Skip_rowsContext)

	// EnterWith_statement is called when entering the with_statement production.
	EnterWith_statement(c *With_statementContext)

	// EnterName is called when entering the name production.
	EnterName(c *NameContext)

	// EnterDot_name is called when entering the dot_name production.
	EnterDot_name(c *Dot_nameContext)

	// EnterAlias_name is called when entering the alias_name production.
	EnterAlias_name(c *Alias_nameContext)

	// EnterArray_name is called when entering the array_name production.
	EnterArray_name(c *Array_nameContext)

	// EnterColumn_name is called when entering the column_name production.
	EnterColumn_name(c *Column_nameContext)

	// EnterCte_name is called when entering the cte_name production.
	EnterCte_name(c *Cte_nameContext)

	// EnterDataset_name is called when entering the dataset_name production.
	EnterDataset_name(c *Dataset_nameContext)

	// EnterDatatype_name is called when entering the datatype_name production.
	EnterDatatype_name(c *Datatype_nameContext)

	// EnterFunction_name is called when entering the function_name production.
	EnterFunction_name(c *Function_nameContext)

	// EnterJoin_name is called when entering the join_name production.
	EnterJoin_name(c *Join_nameContext)

	// EnterMember_name is called when entering the member_name production.
	EnterMember_name(c *Member_nameContext)

	// EnterProject_name is called when entering the project_name production.
	EnterProject_name(c *Project_nameContext)

	// EnterStruct_name is called when entering the struct_name production.
	EnterStruct_name(c *Struct_nameContext)

	// EnterTable_name is called when entering the table_name production.
	EnterTable_name(c *Table_nameContext)

	// EnterTable_expr is called when entering the table_expr production.
	EnterTable_expr(c *Table_exprContext)

	// EnterNumber is called when entering the number production.
	EnterNumber(c *NumberContext)

	// EnterInteger_type is called when entering the integer_type production.
	EnterInteger_type(c *Integer_typeContext)

	// EnterFloat_type is called when entering the float_type production.
	EnterFloat_type(c *Float_typeContext)

	// EnterString is called when entering the string production.
	EnterString(c *StringContext)

	// EnterQuoted_string is called when entering the quoted_string production.
	EnterQuoted_string(c *Quoted_stringContext)

	// EnterTriple_quoted_string is called when entering the triple_quoted_string production.
	EnterTriple_quoted_string(c *Triple_quoted_stringContext)

	// EnterRaw_string is called when entering the raw_string production.
	EnterRaw_string(c *Raw_stringContext)

	// EnterByte_string is called when entering the byte_string production.
	EnterByte_string(c *Byte_stringContext)

	// EnterRaw_byte_string is called when entering the raw_byte_string production.
	EnterRaw_byte_string(c *Raw_byte_stringContext)

	// EnterSpecial_string is called when entering the special_string production.
	EnterSpecial_string(c *Special_stringContext)

	// EnterKeyword is called when entering the keyword production.
	EnterKeyword(c *KeywordContext)

	// ExitRoot is called when exiting the root production.
	ExitRoot(c *RootContext)

	// ExitStmtmulti is called when exiting the stmtmulti production.
	ExitStmtmulti(c *StmtmultiContext)

	// ExitStmt is called when exiting the stmt production.
	ExitStmt(c *StmtContext)

	// ExitQuery_statement is called when exiting the query_statement production.
	ExitQuery_statement(c *Query_statementContext)

	// ExitQuery_expr is called when exiting the query_expr production.
	ExitQuery_expr(c *Query_exprContext)

	// ExitSelect_statement is called when exiting the select_statement production.
	ExitSelect_statement(c *Select_statementContext)

	// ExitFrom_statement is called when exiting the from_statement production.
	ExitFrom_statement(c *From_statementContext)

	// ExitFrom_item is called when exiting the from_item production.
	ExitFrom_item(c *From_itemContext)

	// ExitWhere_statement is called when exiting the where_statement production.
	ExitWhere_statement(c *Where_statementContext)

	// ExitGroup_statement is called when exiting the group_statement production.
	ExitGroup_statement(c *Group_statementContext)

	// ExitHaving_statement is called when exiting the having_statement production.
	ExitHaving_statement(c *Having_statementContext)

	// ExitWindow_statement is called when exiting the window_statement production.
	ExitWindow_statement(c *Window_statementContext)

	// ExitOrder_clause is called when exiting the order_clause production.
	ExitOrder_clause(c *Order_clauseContext)

	// ExitLimit_clause is called when exiting the limit_clause production.
	ExitLimit_clause(c *Limit_clauseContext)

	// ExitUnary_operator is called when exiting the unary_operator production.
	ExitUnary_operator(c *Unary_operatorContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitCast_expr is called when exiting the cast_expr production.
	ExitCast_expr(c *Cast_exprContext)

	// ExitColumn_expr is called when exiting the column_expr production.
	ExitColumn_expr(c *Column_exprContext)

	// ExitExcept_statement is called when exiting the except_statement production.
	ExitExcept_statement(c *Except_statementContext)

	// ExitReplace_statement is called when exiting the replace_statement production.
	ExitReplace_statement(c *Replace_statementContext)

	// ExitJoin_type is called when exiting the join_type production.
	ExitJoin_type(c *Join_typeContext)

	// ExitOn_clause is called when exiting the on_clause production.
	ExitOn_clause(c *On_clauseContext)

	// ExitSet_op is called when exiting the set_op production.
	ExitSet_op(c *Set_opContext)

	// ExitUsing_clause is called when exiting the using_clause production.
	ExitUsing_clause(c *Using_clauseContext)

	// ExitField_path is called when exiting the field_path production.
	ExitField_path(c *Field_pathContext)

	// ExitSstruct is called when exiting the sstruct production.
	ExitSstruct(c *SstructContext)

	// ExitArray_expr is called when exiting the array_expr production.
	ExitArray_expr(c *Array_exprContext)

	// ExitArray_path is called when exiting the array_path production.
	ExitArray_path(c *Array_pathContext)

	// ExitBool_expression is called when exiting the bool_expression production.
	ExitBool_expression(c *Bool_expressionContext)

	// ExitWindow_name is called when exiting the window_name production.
	ExitWindow_name(c *Window_nameContext)

	// ExitWindow_definition is called when exiting the window_definition production.
	ExitWindow_definition(c *Window_definitionContext)

	// ExitCount is called when exiting the count production.
	ExitCount(c *CountContext)

	// ExitSkip_rows is called when exiting the skip_rows production.
	ExitSkip_rows(c *Skip_rowsContext)

	// ExitWith_statement is called when exiting the with_statement production.
	ExitWith_statement(c *With_statementContext)

	// ExitName is called when exiting the name production.
	ExitName(c *NameContext)

	// ExitDot_name is called when exiting the dot_name production.
	ExitDot_name(c *Dot_nameContext)

	// ExitAlias_name is called when exiting the alias_name production.
	ExitAlias_name(c *Alias_nameContext)

	// ExitArray_name is called when exiting the array_name production.
	ExitArray_name(c *Array_nameContext)

	// ExitColumn_name is called when exiting the column_name production.
	ExitColumn_name(c *Column_nameContext)

	// ExitCte_name is called when exiting the cte_name production.
	ExitCte_name(c *Cte_nameContext)

	// ExitDataset_name is called when exiting the dataset_name production.
	ExitDataset_name(c *Dataset_nameContext)

	// ExitDatatype_name is called when exiting the datatype_name production.
	ExitDatatype_name(c *Datatype_nameContext)

	// ExitFunction_name is called when exiting the function_name production.
	ExitFunction_name(c *Function_nameContext)

	// ExitJoin_name is called when exiting the join_name production.
	ExitJoin_name(c *Join_nameContext)

	// ExitMember_name is called when exiting the member_name production.
	ExitMember_name(c *Member_nameContext)

	// ExitProject_name is called when exiting the project_name production.
	ExitProject_name(c *Project_nameContext)

	// ExitStruct_name is called when exiting the struct_name production.
	ExitStruct_name(c *Struct_nameContext)

	// ExitTable_name is called when exiting the table_name production.
	ExitTable_name(c *Table_nameContext)

	// ExitTable_expr is called when exiting the table_expr production.
	ExitTable_expr(c *Table_exprContext)

	// ExitNumber is called when exiting the number production.
	ExitNumber(c *NumberContext)

	// ExitInteger_type is called when exiting the integer_type production.
	ExitInteger_type(c *Integer_typeContext)

	// ExitFloat_type is called when exiting the float_type production.
	ExitFloat_type(c *Float_typeContext)

	// ExitString is called when exiting the string production.
	ExitString(c *StringContext)

	// ExitQuoted_string is called when exiting the quoted_string production.
	ExitQuoted_string(c *Quoted_stringContext)

	// ExitTriple_quoted_string is called when exiting the triple_quoted_string production.
	ExitTriple_quoted_string(c *Triple_quoted_stringContext)

	// ExitRaw_string is called when exiting the raw_string production.
	ExitRaw_string(c *Raw_stringContext)

	// ExitByte_string is called when exiting the byte_string production.
	ExitByte_string(c *Byte_stringContext)

	// ExitRaw_byte_string is called when exiting the raw_byte_string production.
	ExitRaw_byte_string(c *Raw_byte_stringContext)

	// ExitSpecial_string is called when exiting the special_string production.
	ExitSpecial_string(c *Special_stringContext)

	// ExitKeyword is called when exiting the keyword production.
	ExitKeyword(c *KeywordContext)
}
