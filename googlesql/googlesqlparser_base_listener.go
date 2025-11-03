// Code generated from GoogleSQLParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package googlesql // GoogleSQLParser
import "github.com/antlr4-go/antlr/v4"

// BaseGoogleSQLParserListener is a complete listener for a parse tree produced by GoogleSQLParser.
type BaseGoogleSQLParserListener struct{}

var _ GoogleSQLParserListener = &BaseGoogleSQLParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseGoogleSQLParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseGoogleSQLParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseGoogleSQLParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseGoogleSQLParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterRoot is called when production root is entered.
func (s *BaseGoogleSQLParserListener) EnterRoot(ctx *RootContext) {}

// ExitRoot is called when production root is exited.
func (s *BaseGoogleSQLParserListener) ExitRoot(ctx *RootContext) {}

// EnterStmts is called when production stmts is entered.
func (s *BaseGoogleSQLParserListener) EnterStmts(ctx *StmtsContext) {}

// ExitStmts is called when production stmts is exited.
func (s *BaseGoogleSQLParserListener) ExitStmts(ctx *StmtsContext) {}

// EnterUnterminated_sql_statement is called when production unterminated_sql_statement is entered.
func (s *BaseGoogleSQLParserListener) EnterUnterminated_sql_statement(ctx *Unterminated_sql_statementContext) {
}

// ExitUnterminated_sql_statement is called when production unterminated_sql_statement is exited.
func (s *BaseGoogleSQLParserListener) ExitUnterminated_sql_statement(ctx *Unterminated_sql_statementContext) {
}

// EnterSql_statement_body is called when production sql_statement_body is entered.
func (s *BaseGoogleSQLParserListener) EnterSql_statement_body(ctx *Sql_statement_bodyContext) {}

// ExitSql_statement_body is called when production sql_statement_body is exited.
func (s *BaseGoogleSQLParserListener) ExitSql_statement_body(ctx *Sql_statement_bodyContext) {}

// EnterGql_statement is called when production gql_statement is entered.
func (s *BaseGoogleSQLParserListener) EnterGql_statement(ctx *Gql_statementContext) {}

// ExitGql_statement is called when production gql_statement is exited.
func (s *BaseGoogleSQLParserListener) ExitGql_statement(ctx *Gql_statementContext) {}

// EnterGraph_operation_block is called when production graph_operation_block is entered.
func (s *BaseGoogleSQLParserListener) EnterGraph_operation_block(ctx *Graph_operation_blockContext) {}

// ExitGraph_operation_block is called when production graph_operation_block is exited.
func (s *BaseGoogleSQLParserListener) ExitGraph_operation_block(ctx *Graph_operation_blockContext) {}

// EnterGraph_composite_query_block is called when production graph_composite_query_block is entered.
func (s *BaseGoogleSQLParserListener) EnterGraph_composite_query_block(ctx *Graph_composite_query_blockContext) {
}

// ExitGraph_composite_query_block is called when production graph_composite_query_block is exited.
func (s *BaseGoogleSQLParserListener) ExitGraph_composite_query_block(ctx *Graph_composite_query_blockContext) {
}

// EnterGraph_composite_query_prefix is called when production graph_composite_query_prefix is entered.
func (s *BaseGoogleSQLParserListener) EnterGraph_composite_query_prefix(ctx *Graph_composite_query_prefixContext) {
}

// ExitGraph_composite_query_prefix is called when production graph_composite_query_prefix is exited.
func (s *BaseGoogleSQLParserListener) ExitGraph_composite_query_prefix(ctx *Graph_composite_query_prefixContext) {
}

// EnterGraph_set_operation_metadata is called when production graph_set_operation_metadata is entered.
func (s *BaseGoogleSQLParserListener) EnterGraph_set_operation_metadata(ctx *Graph_set_operation_metadataContext) {
}

// ExitGraph_set_operation_metadata is called when production graph_set_operation_metadata is exited.
func (s *BaseGoogleSQLParserListener) ExitGraph_set_operation_metadata(ctx *Graph_set_operation_metadataContext) {
}

// EnterGraph_linear_query_operation is called when production graph_linear_query_operation is entered.
func (s *BaseGoogleSQLParserListener) EnterGraph_linear_query_operation(ctx *Graph_linear_query_operationContext) {
}

// ExitGraph_linear_query_operation is called when production graph_linear_query_operation is exited.
func (s *BaseGoogleSQLParserListener) ExitGraph_linear_query_operation(ctx *Graph_linear_query_operationContext) {
}

// EnterGraph_linear_operator_list is called when production graph_linear_operator_list is entered.
func (s *BaseGoogleSQLParserListener) EnterGraph_linear_operator_list(ctx *Graph_linear_operator_listContext) {
}

// ExitGraph_linear_operator_list is called when production graph_linear_operator_list is exited.
func (s *BaseGoogleSQLParserListener) ExitGraph_linear_operator_list(ctx *Graph_linear_operator_listContext) {
}

// EnterGraph_linear_operator is called when production graph_linear_operator is entered.
func (s *BaseGoogleSQLParserListener) EnterGraph_linear_operator(ctx *Graph_linear_operatorContext) {}

// ExitGraph_linear_operator is called when production graph_linear_operator is exited.
func (s *BaseGoogleSQLParserListener) ExitGraph_linear_operator(ctx *Graph_linear_operatorContext) {}

// EnterGraph_sample_clause is called when production graph_sample_clause is entered.
func (s *BaseGoogleSQLParserListener) EnterGraph_sample_clause(ctx *Graph_sample_clauseContext) {}

// ExitGraph_sample_clause is called when production graph_sample_clause is exited.
func (s *BaseGoogleSQLParserListener) ExitGraph_sample_clause(ctx *Graph_sample_clauseContext) {}

// EnterOpt_graph_sample_clause_suffix is called when production opt_graph_sample_clause_suffix is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_graph_sample_clause_suffix(ctx *Opt_graph_sample_clause_suffixContext) {
}

// ExitOpt_graph_sample_clause_suffix is called when production opt_graph_sample_clause_suffix is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_graph_sample_clause_suffix(ctx *Opt_graph_sample_clause_suffixContext) {
}

// EnterGraph_for_operator is called when production graph_for_operator is entered.
func (s *BaseGoogleSQLParserListener) EnterGraph_for_operator(ctx *Graph_for_operatorContext) {}

// ExitGraph_for_operator is called when production graph_for_operator is exited.
func (s *BaseGoogleSQLParserListener) ExitGraph_for_operator(ctx *Graph_for_operatorContext) {}

// EnterOpt_with_offset_and_alias_with_required_as is called when production opt_with_offset_and_alias_with_required_as is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_with_offset_and_alias_with_required_as(ctx *Opt_with_offset_and_alias_with_required_asContext) {
}

// ExitOpt_with_offset_and_alias_with_required_as is called when production opt_with_offset_and_alias_with_required_as is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_with_offset_and_alias_with_required_as(ctx *Opt_with_offset_and_alias_with_required_asContext) {
}

// EnterGraph_with_operator is called when production graph_with_operator is entered.
func (s *BaseGoogleSQLParserListener) EnterGraph_with_operator(ctx *Graph_with_operatorContext) {}

// ExitGraph_with_operator is called when production graph_with_operator is exited.
func (s *BaseGoogleSQLParserListener) ExitGraph_with_operator(ctx *Graph_with_operatorContext) {}

// EnterGraph_page_operator is called when production graph_page_operator is entered.
func (s *BaseGoogleSQLParserListener) EnterGraph_page_operator(ctx *Graph_page_operatorContext) {}

// ExitGraph_page_operator is called when production graph_page_operator is exited.
func (s *BaseGoogleSQLParserListener) ExitGraph_page_operator(ctx *Graph_page_operatorContext) {}

// EnterGraph_order_by_operator is called when production graph_order_by_operator is entered.
func (s *BaseGoogleSQLParserListener) EnterGraph_order_by_operator(ctx *Graph_order_by_operatorContext) {
}

// ExitGraph_order_by_operator is called when production graph_order_by_operator is exited.
func (s *BaseGoogleSQLParserListener) ExitGraph_order_by_operator(ctx *Graph_order_by_operatorContext) {
}

// EnterGraph_filter_operator is called when production graph_filter_operator is entered.
func (s *BaseGoogleSQLParserListener) EnterGraph_filter_operator(ctx *Graph_filter_operatorContext) {}

// ExitGraph_filter_operator is called when production graph_filter_operator is exited.
func (s *BaseGoogleSQLParserListener) ExitGraph_filter_operator(ctx *Graph_filter_operatorContext) {}

// EnterGraph_let_operator is called when production graph_let_operator is entered.
func (s *BaseGoogleSQLParserListener) EnterGraph_let_operator(ctx *Graph_let_operatorContext) {}

// ExitGraph_let_operator is called when production graph_let_operator is exited.
func (s *BaseGoogleSQLParserListener) ExitGraph_let_operator(ctx *Graph_let_operatorContext) {}

// EnterGraph_let_variable_definition_list is called when production graph_let_variable_definition_list is entered.
func (s *BaseGoogleSQLParserListener) EnterGraph_let_variable_definition_list(ctx *Graph_let_variable_definition_listContext) {
}

// ExitGraph_let_variable_definition_list is called when production graph_let_variable_definition_list is exited.
func (s *BaseGoogleSQLParserListener) ExitGraph_let_variable_definition_list(ctx *Graph_let_variable_definition_listContext) {
}

// EnterGraph_let_variable_definition is called when production graph_let_variable_definition is entered.
func (s *BaseGoogleSQLParserListener) EnterGraph_let_variable_definition(ctx *Graph_let_variable_definitionContext) {
}

// ExitGraph_let_variable_definition is called when production graph_let_variable_definition is exited.
func (s *BaseGoogleSQLParserListener) ExitGraph_let_variable_definition(ctx *Graph_let_variable_definitionContext) {
}

// EnterGraph_optional_match_operator is called when production graph_optional_match_operator is entered.
func (s *BaseGoogleSQLParserListener) EnterGraph_optional_match_operator(ctx *Graph_optional_match_operatorContext) {
}

// ExitGraph_optional_match_operator is called when production graph_optional_match_operator is exited.
func (s *BaseGoogleSQLParserListener) ExitGraph_optional_match_operator(ctx *Graph_optional_match_operatorContext) {
}

// EnterGraph_match_operator is called when production graph_match_operator is entered.
func (s *BaseGoogleSQLParserListener) EnterGraph_match_operator(ctx *Graph_match_operatorContext) {}

// ExitGraph_match_operator is called when production graph_match_operator is exited.
func (s *BaseGoogleSQLParserListener) ExitGraph_match_operator(ctx *Graph_match_operatorContext) {}

// EnterGraph_pattern is called when production graph_pattern is entered.
func (s *BaseGoogleSQLParserListener) EnterGraph_pattern(ctx *Graph_patternContext) {}

// ExitGraph_pattern is called when production graph_pattern is exited.
func (s *BaseGoogleSQLParserListener) ExitGraph_pattern(ctx *Graph_patternContext) {}

// EnterGraph_path_pattern_list is called when production graph_path_pattern_list is entered.
func (s *BaseGoogleSQLParserListener) EnterGraph_path_pattern_list(ctx *Graph_path_pattern_listContext) {
}

// ExitGraph_path_pattern_list is called when production graph_path_pattern_list is exited.
func (s *BaseGoogleSQLParserListener) ExitGraph_path_pattern_list(ctx *Graph_path_pattern_listContext) {
}

// EnterGraph_path_pattern is called when production graph_path_pattern is entered.
func (s *BaseGoogleSQLParserListener) EnterGraph_path_pattern(ctx *Graph_path_patternContext) {}

// ExitGraph_path_pattern is called when production graph_path_pattern is exited.
func (s *BaseGoogleSQLParserListener) ExitGraph_path_pattern(ctx *Graph_path_patternContext) {}

// EnterGraph_path_pattern_expr is called when production graph_path_pattern_expr is entered.
func (s *BaseGoogleSQLParserListener) EnterGraph_path_pattern_expr(ctx *Graph_path_pattern_exprContext) {
}

// ExitGraph_path_pattern_expr is called when production graph_path_pattern_expr is exited.
func (s *BaseGoogleSQLParserListener) ExitGraph_path_pattern_expr(ctx *Graph_path_pattern_exprContext) {
}

// EnterGraph_path_factor is called when production graph_path_factor is entered.
func (s *BaseGoogleSQLParserListener) EnterGraph_path_factor(ctx *Graph_path_factorContext) {}

// ExitGraph_path_factor is called when production graph_path_factor is exited.
func (s *BaseGoogleSQLParserListener) ExitGraph_path_factor(ctx *Graph_path_factorContext) {}

// EnterGraph_quantified_path_primary is called when production graph_quantified_path_primary is entered.
func (s *BaseGoogleSQLParserListener) EnterGraph_quantified_path_primary(ctx *Graph_quantified_path_primaryContext) {
}

// ExitGraph_quantified_path_primary is called when production graph_quantified_path_primary is exited.
func (s *BaseGoogleSQLParserListener) ExitGraph_quantified_path_primary(ctx *Graph_quantified_path_primaryContext) {
}

// EnterGraph_path_primary is called when production graph_path_primary is entered.
func (s *BaseGoogleSQLParserListener) EnterGraph_path_primary(ctx *Graph_path_primaryContext) {}

// ExitGraph_path_primary is called when production graph_path_primary is exited.
func (s *BaseGoogleSQLParserListener) ExitGraph_path_primary(ctx *Graph_path_primaryContext) {}

// EnterGraph_parenthesized_path_pattern is called when production graph_parenthesized_path_pattern is entered.
func (s *BaseGoogleSQLParserListener) EnterGraph_parenthesized_path_pattern(ctx *Graph_parenthesized_path_patternContext) {
}

// ExitGraph_parenthesized_path_pattern is called when production graph_parenthesized_path_pattern is exited.
func (s *BaseGoogleSQLParserListener) ExitGraph_parenthesized_path_pattern(ctx *Graph_parenthesized_path_patternContext) {
}

// EnterGraph_element_pattern is called when production graph_element_pattern is entered.
func (s *BaseGoogleSQLParserListener) EnterGraph_element_pattern(ctx *Graph_element_patternContext) {}

// ExitGraph_element_pattern is called when production graph_element_pattern is exited.
func (s *BaseGoogleSQLParserListener) ExitGraph_element_pattern(ctx *Graph_element_patternContext) {}

// EnterGraph_edge_pattern is called when production graph_edge_pattern is entered.
func (s *BaseGoogleSQLParserListener) EnterGraph_edge_pattern(ctx *Graph_edge_patternContext) {}

// ExitGraph_edge_pattern is called when production graph_edge_pattern is exited.
func (s *BaseGoogleSQLParserListener) ExitGraph_edge_pattern(ctx *Graph_edge_patternContext) {}

// EnterGraph_node_pattern is called when production graph_node_pattern is entered.
func (s *BaseGoogleSQLParserListener) EnterGraph_node_pattern(ctx *Graph_node_patternContext) {}

// ExitGraph_node_pattern is called when production graph_node_pattern is exited.
func (s *BaseGoogleSQLParserListener) ExitGraph_node_pattern(ctx *Graph_node_patternContext) {}

// EnterGraph_element_pattern_filler is called when production graph_element_pattern_filler is entered.
func (s *BaseGoogleSQLParserListener) EnterGraph_element_pattern_filler(ctx *Graph_element_pattern_fillerContext) {
}

// ExitGraph_element_pattern_filler is called when production graph_element_pattern_filler is exited.
func (s *BaseGoogleSQLParserListener) ExitGraph_element_pattern_filler(ctx *Graph_element_pattern_fillerContext) {
}

// EnterGraph_property_specification is called when production graph_property_specification is entered.
func (s *BaseGoogleSQLParserListener) EnterGraph_property_specification(ctx *Graph_property_specificationContext) {
}

// ExitGraph_property_specification is called when production graph_property_specification is exited.
func (s *BaseGoogleSQLParserListener) ExitGraph_property_specification(ctx *Graph_property_specificationContext) {
}

// EnterGraph_property_name_and_value is called when production graph_property_name_and_value is entered.
func (s *BaseGoogleSQLParserListener) EnterGraph_property_name_and_value(ctx *Graph_property_name_and_valueContext) {
}

// ExitGraph_property_name_and_value is called when production graph_property_name_and_value is exited.
func (s *BaseGoogleSQLParserListener) ExitGraph_property_name_and_value(ctx *Graph_property_name_and_valueContext) {
}

// EnterOpt_is_label_expression is called when production opt_is_label_expression is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_is_label_expression(ctx *Opt_is_label_expressionContext) {
}

// ExitOpt_is_label_expression is called when production opt_is_label_expression is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_is_label_expression(ctx *Opt_is_label_expressionContext) {
}

// EnterLabel_expression is called when production label_expression is entered.
func (s *BaseGoogleSQLParserListener) EnterLabel_expression(ctx *Label_expressionContext) {}

// ExitLabel_expression is called when production label_expression is exited.
func (s *BaseGoogleSQLParserListener) ExitLabel_expression(ctx *Label_expressionContext) {}

// EnterLabel_primary is called when production label_primary is entered.
func (s *BaseGoogleSQLParserListener) EnterLabel_primary(ctx *Label_primaryContext) {}

// ExitLabel_primary is called when production label_primary is exited.
func (s *BaseGoogleSQLParserListener) ExitLabel_primary(ctx *Label_primaryContext) {}

// EnterParenthesized_label_expression is called when production parenthesized_label_expression is entered.
func (s *BaseGoogleSQLParserListener) EnterParenthesized_label_expression(ctx *Parenthesized_label_expressionContext) {
}

// ExitParenthesized_label_expression is called when production parenthesized_label_expression is exited.
func (s *BaseGoogleSQLParserListener) ExitParenthesized_label_expression(ctx *Parenthesized_label_expressionContext) {
}

// EnterOpt_graph_element_identifier is called when production opt_graph_element_identifier is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_graph_element_identifier(ctx *Opt_graph_element_identifierContext) {
}

// ExitOpt_graph_element_identifier is called when production opt_graph_element_identifier is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_graph_element_identifier(ctx *Opt_graph_element_identifierContext) {
}

// EnterOpt_graph_path_mode_prefix is called when production opt_graph_path_mode_prefix is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_graph_path_mode_prefix(ctx *Opt_graph_path_mode_prefixContext) {
}

// ExitOpt_graph_path_mode_prefix is called when production opt_graph_path_mode_prefix is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_graph_path_mode_prefix(ctx *Opt_graph_path_mode_prefixContext) {
}

// EnterPath_or_paths is called when production path_or_paths is entered.
func (s *BaseGoogleSQLParserListener) EnterPath_or_paths(ctx *Path_or_pathsContext) {}

// ExitPath_or_paths is called when production path_or_paths is exited.
func (s *BaseGoogleSQLParserListener) ExitPath_or_paths(ctx *Path_or_pathsContext) {}

// EnterOpt_graph_path_mode is called when production opt_graph_path_mode is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_graph_path_mode(ctx *Opt_graph_path_modeContext) {}

// ExitOpt_graph_path_mode is called when production opt_graph_path_mode is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_graph_path_mode(ctx *Opt_graph_path_modeContext) {}

// EnterOpt_graph_search_prefix is called when production opt_graph_search_prefix is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_graph_search_prefix(ctx *Opt_graph_search_prefixContext) {
}

// ExitOpt_graph_search_prefix is called when production opt_graph_search_prefix is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_graph_search_prefix(ctx *Opt_graph_search_prefixContext) {
}

// EnterOpt_path_variable_assignment is called when production opt_path_variable_assignment is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_path_variable_assignment(ctx *Opt_path_variable_assignmentContext) {
}

// ExitOpt_path_variable_assignment is called when production opt_path_variable_assignment is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_path_variable_assignment(ctx *Opt_path_variable_assignmentContext) {
}

// EnterGraph_identifier is called when production graph_identifier is entered.
func (s *BaseGoogleSQLParserListener) EnterGraph_identifier(ctx *Graph_identifierContext) {}

// ExitGraph_identifier is called when production graph_identifier is exited.
func (s *BaseGoogleSQLParserListener) ExitGraph_identifier(ctx *Graph_identifierContext) {}

// EnterGraph_return_operator is called when production graph_return_operator is entered.
func (s *BaseGoogleSQLParserListener) EnterGraph_return_operator(ctx *Graph_return_operatorContext) {}

// ExitGraph_return_operator is called when production graph_return_operator is exited.
func (s *BaseGoogleSQLParserListener) ExitGraph_return_operator(ctx *Graph_return_operatorContext) {}

// EnterGraph_page_clause is called when production graph_page_clause is entered.
func (s *BaseGoogleSQLParserListener) EnterGraph_page_clause(ctx *Graph_page_clauseContext) {}

// ExitGraph_page_clause is called when production graph_page_clause is exited.
func (s *BaseGoogleSQLParserListener) ExitGraph_page_clause(ctx *Graph_page_clauseContext) {}

// EnterGraph_order_by_clause is called when production graph_order_by_clause is entered.
func (s *BaseGoogleSQLParserListener) EnterGraph_order_by_clause(ctx *Graph_order_by_clauseContext) {}

// ExitGraph_order_by_clause is called when production graph_order_by_clause is exited.
func (s *BaseGoogleSQLParserListener) ExitGraph_order_by_clause(ctx *Graph_order_by_clauseContext) {}

// EnterGraph_ordering_expression is called when production graph_ordering_expression is entered.
func (s *BaseGoogleSQLParserListener) EnterGraph_ordering_expression(ctx *Graph_ordering_expressionContext) {
}

// ExitGraph_ordering_expression is called when production graph_ordering_expression is exited.
func (s *BaseGoogleSQLParserListener) ExitGraph_ordering_expression(ctx *Graph_ordering_expressionContext) {
}

// EnterOpt_graph_asc_or_desc is called when production opt_graph_asc_or_desc is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_graph_asc_or_desc(ctx *Opt_graph_asc_or_descContext) {}

// ExitOpt_graph_asc_or_desc is called when production opt_graph_asc_or_desc is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_graph_asc_or_desc(ctx *Opt_graph_asc_or_descContext) {}

// EnterGraph_return_item_list is called when production graph_return_item_list is entered.
func (s *BaseGoogleSQLParserListener) EnterGraph_return_item_list(ctx *Graph_return_item_listContext) {
}

// ExitGraph_return_item_list is called when production graph_return_item_list is exited.
func (s *BaseGoogleSQLParserListener) ExitGraph_return_item_list(ctx *Graph_return_item_listContext) {
}

// EnterGraph_return_item is called when production graph_return_item is entered.
func (s *BaseGoogleSQLParserListener) EnterGraph_return_item(ctx *Graph_return_itemContext) {}

// ExitGraph_return_item is called when production graph_return_item is exited.
func (s *BaseGoogleSQLParserListener) ExitGraph_return_item(ctx *Graph_return_itemContext) {}

// EnterUndrop_statement is called when production undrop_statement is entered.
func (s *BaseGoogleSQLParserListener) EnterUndrop_statement(ctx *Undrop_statementContext) {}

// ExitUndrop_statement is called when production undrop_statement is exited.
func (s *BaseGoogleSQLParserListener) ExitUndrop_statement(ctx *Undrop_statementContext) {}

// EnterModule_statement is called when production module_statement is entered.
func (s *BaseGoogleSQLParserListener) EnterModule_statement(ctx *Module_statementContext) {}

// ExitModule_statement is called when production module_statement is exited.
func (s *BaseGoogleSQLParserListener) ExitModule_statement(ctx *Module_statementContext) {}

// EnterImport_statement is called when production import_statement is entered.
func (s *BaseGoogleSQLParserListener) EnterImport_statement(ctx *Import_statementContext) {}

// ExitImport_statement is called when production import_statement is exited.
func (s *BaseGoogleSQLParserListener) ExitImport_statement(ctx *Import_statementContext) {}

// EnterOpt_as_or_into_alias is called when production opt_as_or_into_alias is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_as_or_into_alias(ctx *Opt_as_or_into_aliasContext) {}

// ExitOpt_as_or_into_alias is called when production opt_as_or_into_alias is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_as_or_into_alias(ctx *Opt_as_or_into_aliasContext) {}

// EnterPath_expression_or_string is called when production path_expression_or_string is entered.
func (s *BaseGoogleSQLParserListener) EnterPath_expression_or_string(ctx *Path_expression_or_stringContext) {
}

// ExitPath_expression_or_string is called when production path_expression_or_string is exited.
func (s *BaseGoogleSQLParserListener) ExitPath_expression_or_string(ctx *Path_expression_or_stringContext) {
}

// EnterImport_type is called when production import_type is entered.
func (s *BaseGoogleSQLParserListener) EnterImport_type(ctx *Import_typeContext) {}

// ExitImport_type is called when production import_type is exited.
func (s *BaseGoogleSQLParserListener) ExitImport_type(ctx *Import_typeContext) {}

// EnterCall_statement is called when production call_statement is entered.
func (s *BaseGoogleSQLParserListener) EnterCall_statement(ctx *Call_statementContext) {}

// ExitCall_statement is called when production call_statement is exited.
func (s *BaseGoogleSQLParserListener) ExitCall_statement(ctx *Call_statementContext) {}

// EnterDrop_statement is called when production drop_statement is entered.
func (s *BaseGoogleSQLParserListener) EnterDrop_statement(ctx *Drop_statementContext) {}

// ExitDrop_statement is called when production drop_statement is exited.
func (s *BaseGoogleSQLParserListener) ExitDrop_statement(ctx *Drop_statementContext) {}

// EnterOpt_drop_mode is called when production opt_drop_mode is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_drop_mode(ctx *Opt_drop_modeContext) {}

// ExitOpt_drop_mode is called when production opt_drop_mode is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_drop_mode(ctx *Opt_drop_modeContext) {}

// EnterDrop_all_row_access_policies_statement is called when production drop_all_row_access_policies_statement is entered.
func (s *BaseGoogleSQLParserListener) EnterDrop_all_row_access_policies_statement(ctx *Drop_all_row_access_policies_statementContext) {
}

// ExitDrop_all_row_access_policies_statement is called when production drop_all_row_access_policies_statement is exited.
func (s *BaseGoogleSQLParserListener) ExitDrop_all_row_access_policies_statement(ctx *Drop_all_row_access_policies_statementContext) {
}

// EnterShow_statement is called when production show_statement is entered.
func (s *BaseGoogleSQLParserListener) EnterShow_statement(ctx *Show_statementContext) {}

// ExitShow_statement is called when production show_statement is exited.
func (s *BaseGoogleSQLParserListener) ExitShow_statement(ctx *Show_statementContext) {}

// EnterOpt_like_string_literal is called when production opt_like_string_literal is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_like_string_literal(ctx *Opt_like_string_literalContext) {
}

// ExitOpt_like_string_literal is called when production opt_like_string_literal is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_like_string_literal(ctx *Opt_like_string_literalContext) {
}

// EnterShow_target is called when production show_target is entered.
func (s *BaseGoogleSQLParserListener) EnterShow_target(ctx *Show_targetContext) {}

// ExitShow_target is called when production show_target is exited.
func (s *BaseGoogleSQLParserListener) ExitShow_target(ctx *Show_targetContext) {}

// EnterRename_statement is called when production rename_statement is entered.
func (s *BaseGoogleSQLParserListener) EnterRename_statement(ctx *Rename_statementContext) {}

// ExitRename_statement is called when production rename_statement is exited.
func (s *BaseGoogleSQLParserListener) ExitRename_statement(ctx *Rename_statementContext) {}

// EnterRevoke_statement is called when production revoke_statement is entered.
func (s *BaseGoogleSQLParserListener) EnterRevoke_statement(ctx *Revoke_statementContext) {}

// ExitRevoke_statement is called when production revoke_statement is exited.
func (s *BaseGoogleSQLParserListener) ExitRevoke_statement(ctx *Revoke_statementContext) {}

// EnterGrant_statement is called when production grant_statement is entered.
func (s *BaseGoogleSQLParserListener) EnterGrant_statement(ctx *Grant_statementContext) {}

// ExitGrant_statement is called when production grant_statement is exited.
func (s *BaseGoogleSQLParserListener) ExitGrant_statement(ctx *Grant_statementContext) {}

// EnterPrivileges is called when production privileges is entered.
func (s *BaseGoogleSQLParserListener) EnterPrivileges(ctx *PrivilegesContext) {}

// ExitPrivileges is called when production privileges is exited.
func (s *BaseGoogleSQLParserListener) ExitPrivileges(ctx *PrivilegesContext) {}

// EnterExport_metadata_statement is called when production export_metadata_statement is entered.
func (s *BaseGoogleSQLParserListener) EnterExport_metadata_statement(ctx *Export_metadata_statementContext) {
}

// ExitExport_metadata_statement is called when production export_metadata_statement is exited.
func (s *BaseGoogleSQLParserListener) ExitExport_metadata_statement(ctx *Export_metadata_statementContext) {
}

// EnterExport_model_statement is called when production export_model_statement is entered.
func (s *BaseGoogleSQLParserListener) EnterExport_model_statement(ctx *Export_model_statementContext) {
}

// ExitExport_model_statement is called when production export_model_statement is exited.
func (s *BaseGoogleSQLParserListener) ExitExport_model_statement(ctx *Export_model_statementContext) {
}

// EnterExport_data_statement is called when production export_data_statement is entered.
func (s *BaseGoogleSQLParserListener) EnterExport_data_statement(ctx *Export_data_statementContext) {}

// ExitExport_data_statement is called when production export_data_statement is exited.
func (s *BaseGoogleSQLParserListener) ExitExport_data_statement(ctx *Export_data_statementContext) {}

// EnterExport_data_no_query is called when production export_data_no_query is entered.
func (s *BaseGoogleSQLParserListener) EnterExport_data_no_query(ctx *Export_data_no_queryContext) {}

// ExitExport_data_no_query is called when production export_data_no_query is exited.
func (s *BaseGoogleSQLParserListener) ExitExport_data_no_query(ctx *Export_data_no_queryContext) {}

// EnterExplain_statement is called when production explain_statement is entered.
func (s *BaseGoogleSQLParserListener) EnterExplain_statement(ctx *Explain_statementContext) {}

// ExitExplain_statement is called when production explain_statement is exited.
func (s *BaseGoogleSQLParserListener) ExitExplain_statement(ctx *Explain_statementContext) {}

// EnterExecute_immediate is called when production execute_immediate is entered.
func (s *BaseGoogleSQLParserListener) EnterExecute_immediate(ctx *Execute_immediateContext) {}

// ExitExecute_immediate is called when production execute_immediate is exited.
func (s *BaseGoogleSQLParserListener) ExitExecute_immediate(ctx *Execute_immediateContext) {}

// EnterOpt_execute_into_clause is called when production opt_execute_into_clause is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_execute_into_clause(ctx *Opt_execute_into_clauseContext) {
}

// ExitOpt_execute_into_clause is called when production opt_execute_into_clause is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_execute_into_clause(ctx *Opt_execute_into_clauseContext) {
}

// EnterOpt_execute_using_clause is called when production opt_execute_using_clause is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_execute_using_clause(ctx *Opt_execute_using_clauseContext) {
}

// ExitOpt_execute_using_clause is called when production opt_execute_using_clause is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_execute_using_clause(ctx *Opt_execute_using_clauseContext) {
}

// EnterExecute_using_argument_list is called when production execute_using_argument_list is entered.
func (s *BaseGoogleSQLParserListener) EnterExecute_using_argument_list(ctx *Execute_using_argument_listContext) {
}

// ExitExecute_using_argument_list is called when production execute_using_argument_list is exited.
func (s *BaseGoogleSQLParserListener) ExitExecute_using_argument_list(ctx *Execute_using_argument_listContext) {
}

// EnterExecute_using_argument is called when production execute_using_argument is entered.
func (s *BaseGoogleSQLParserListener) EnterExecute_using_argument(ctx *Execute_using_argumentContext) {
}

// ExitExecute_using_argument is called when production execute_using_argument is exited.
func (s *BaseGoogleSQLParserListener) ExitExecute_using_argument(ctx *Execute_using_argumentContext) {
}

// EnterDescribe_statement is called when production describe_statement is entered.
func (s *BaseGoogleSQLParserListener) EnterDescribe_statement(ctx *Describe_statementContext) {}

// ExitDescribe_statement is called when production describe_statement is exited.
func (s *BaseGoogleSQLParserListener) ExitDescribe_statement(ctx *Describe_statementContext) {}

// EnterDescribe_info is called when production describe_info is entered.
func (s *BaseGoogleSQLParserListener) EnterDescribe_info(ctx *Describe_infoContext) {}

// ExitDescribe_info is called when production describe_info is exited.
func (s *BaseGoogleSQLParserListener) ExitDescribe_info(ctx *Describe_infoContext) {}

// EnterOpt_from_path_expression is called when production opt_from_path_expression is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_from_path_expression(ctx *Opt_from_path_expressionContext) {
}

// ExitOpt_from_path_expression is called when production opt_from_path_expression is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_from_path_expression(ctx *Opt_from_path_expressionContext) {
}

// EnterDescribe_keyword is called when production describe_keyword is entered.
func (s *BaseGoogleSQLParserListener) EnterDescribe_keyword(ctx *Describe_keywordContext) {}

// ExitDescribe_keyword is called when production describe_keyword is exited.
func (s *BaseGoogleSQLParserListener) ExitDescribe_keyword(ctx *Describe_keywordContext) {}

// EnterDefine_table_statement is called when production define_table_statement is entered.
func (s *BaseGoogleSQLParserListener) EnterDefine_table_statement(ctx *Define_table_statementContext) {
}

// ExitDefine_table_statement is called when production define_table_statement is exited.
func (s *BaseGoogleSQLParserListener) ExitDefine_table_statement(ctx *Define_table_statementContext) {
}

// EnterCreate_entity_statement is called when production create_entity_statement is entered.
func (s *BaseGoogleSQLParserListener) EnterCreate_entity_statement(ctx *Create_entity_statementContext) {
}

// ExitCreate_entity_statement is called when production create_entity_statement is exited.
func (s *BaseGoogleSQLParserListener) ExitCreate_entity_statement(ctx *Create_entity_statementContext) {
}

// EnterOpt_generic_entity_body is called when production opt_generic_entity_body is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_generic_entity_body(ctx *Opt_generic_entity_bodyContext) {
}

// ExitOpt_generic_entity_body is called when production opt_generic_entity_body is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_generic_entity_body(ctx *Opt_generic_entity_bodyContext) {
}

// EnterCreate_view_statement is called when production create_view_statement is entered.
func (s *BaseGoogleSQLParserListener) EnterCreate_view_statement(ctx *Create_view_statementContext) {}

// ExitCreate_view_statement is called when production create_view_statement is exited.
func (s *BaseGoogleSQLParserListener) ExitCreate_view_statement(ctx *Create_view_statementContext) {}

// EnterQuery_or_replica_source is called when production query_or_replica_source is entered.
func (s *BaseGoogleSQLParserListener) EnterQuery_or_replica_source(ctx *Query_or_replica_sourceContext) {
}

// ExitQuery_or_replica_source is called when production query_or_replica_source is exited.
func (s *BaseGoogleSQLParserListener) ExitQuery_or_replica_source(ctx *Query_or_replica_sourceContext) {
}

// EnterColumn_with_options_list is called when production column_with_options_list is entered.
func (s *BaseGoogleSQLParserListener) EnterColumn_with_options_list(ctx *Column_with_options_listContext) {
}

// ExitColumn_with_options_list is called when production column_with_options_list is exited.
func (s *BaseGoogleSQLParserListener) ExitColumn_with_options_list(ctx *Column_with_options_listContext) {
}

// EnterColumn_with_options is called when production column_with_options is entered.
func (s *BaseGoogleSQLParserListener) EnterColumn_with_options(ctx *Column_with_optionsContext) {}

// ExitColumn_with_options is called when production column_with_options is exited.
func (s *BaseGoogleSQLParserListener) ExitColumn_with_options(ctx *Column_with_optionsContext) {}

// EnterCreate_table_statement is called when production create_table_statement is entered.
func (s *BaseGoogleSQLParserListener) EnterCreate_table_statement(ctx *Create_table_statementContext) {
}

// ExitCreate_table_statement is called when production create_table_statement is exited.
func (s *BaseGoogleSQLParserListener) ExitCreate_table_statement(ctx *Create_table_statementContext) {
}

// EnterOpt_ttl_clause is called when production opt_ttl_clause is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_ttl_clause(ctx *Opt_ttl_clauseContext) {}

// ExitOpt_ttl_clause is called when production opt_ttl_clause is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_ttl_clause(ctx *Opt_ttl_clauseContext) {}

// EnterOpt_copy_table is called when production opt_copy_table is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_copy_table(ctx *Opt_copy_tableContext) {}

// ExitOpt_copy_table is called when production opt_copy_table is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_copy_table(ctx *Opt_copy_tableContext) {}

// EnterCopy_data_source is called when production copy_data_source is entered.
func (s *BaseGoogleSQLParserListener) EnterCopy_data_source(ctx *Copy_data_sourceContext) {}

// ExitCopy_data_source is called when production copy_data_source is exited.
func (s *BaseGoogleSQLParserListener) ExitCopy_data_source(ctx *Copy_data_sourceContext) {}

// EnterOpt_clone_table is called when production opt_clone_table is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_clone_table(ctx *Opt_clone_tableContext) {}

// ExitOpt_clone_table is called when production opt_clone_table is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_clone_table(ctx *Opt_clone_tableContext) {}

// EnterOpt_spanner_table_options is called when production opt_spanner_table_options is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_spanner_table_options(ctx *Opt_spanner_table_optionsContext) {
}

// ExitOpt_spanner_table_options is called when production opt_spanner_table_options is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_spanner_table_options(ctx *Opt_spanner_table_optionsContext) {
}

// EnterOpt_spanner_interleave_in_parent_clause is called when production opt_spanner_interleave_in_parent_clause is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_spanner_interleave_in_parent_clause(ctx *Opt_spanner_interleave_in_parent_clauseContext) {
}

// ExitOpt_spanner_interleave_in_parent_clause is called when production opt_spanner_interleave_in_parent_clause is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_spanner_interleave_in_parent_clause(ctx *Opt_spanner_interleave_in_parent_clauseContext) {
}

// EnterSpanner_primary_key is called when production spanner_primary_key is entered.
func (s *BaseGoogleSQLParserListener) EnterSpanner_primary_key(ctx *Spanner_primary_keyContext) {}

// ExitSpanner_primary_key is called when production spanner_primary_key is exited.
func (s *BaseGoogleSQLParserListener) ExitSpanner_primary_key(ctx *Spanner_primary_keyContext) {}

// EnterCreate_table_function_statement is called when production create_table_function_statement is entered.
func (s *BaseGoogleSQLParserListener) EnterCreate_table_function_statement(ctx *Create_table_function_statementContext) {
}

// ExitCreate_table_function_statement is called when production create_table_function_statement is exited.
func (s *BaseGoogleSQLParserListener) ExitCreate_table_function_statement(ctx *Create_table_function_statementContext) {
}

// EnterOpt_as_query_or_string is called when production opt_as_query_or_string is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_as_query_or_string(ctx *Opt_as_query_or_stringContext) {
}

// ExitOpt_as_query_or_string is called when production opt_as_query_or_string is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_as_query_or_string(ctx *Opt_as_query_or_stringContext) {
}

// EnterUnordered_language_options is called when production unordered_language_options is entered.
func (s *BaseGoogleSQLParserListener) EnterUnordered_language_options(ctx *Unordered_language_optionsContext) {
}

// ExitUnordered_language_options is called when production unordered_language_options is exited.
func (s *BaseGoogleSQLParserListener) ExitUnordered_language_options(ctx *Unordered_language_optionsContext) {
}

// EnterOpt_function_parameters is called when production opt_function_parameters is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_function_parameters(ctx *Opt_function_parametersContext) {
}

// ExitOpt_function_parameters is called when production opt_function_parameters is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_function_parameters(ctx *Opt_function_parametersContext) {
}

// EnterCreate_snapshot_statement is called when production create_snapshot_statement is entered.
func (s *BaseGoogleSQLParserListener) EnterCreate_snapshot_statement(ctx *Create_snapshot_statementContext) {
}

// ExitCreate_snapshot_statement is called when production create_snapshot_statement is exited.
func (s *BaseGoogleSQLParserListener) ExitCreate_snapshot_statement(ctx *Create_snapshot_statementContext) {
}

// EnterCreate_external_schema_statement is called when production create_external_schema_statement is entered.
func (s *BaseGoogleSQLParserListener) EnterCreate_external_schema_statement(ctx *Create_external_schema_statementContext) {
}

// ExitCreate_external_schema_statement is called when production create_external_schema_statement is exited.
func (s *BaseGoogleSQLParserListener) ExitCreate_external_schema_statement(ctx *Create_external_schema_statementContext) {
}

// EnterCreate_schema_statement is called when production create_schema_statement is entered.
func (s *BaseGoogleSQLParserListener) EnterCreate_schema_statement(ctx *Create_schema_statementContext) {
}

// ExitCreate_schema_statement is called when production create_schema_statement is exited.
func (s *BaseGoogleSQLParserListener) ExitCreate_schema_statement(ctx *Create_schema_statementContext) {
}

// EnterCreate_property_graph_statement is called when production create_property_graph_statement is entered.
func (s *BaseGoogleSQLParserListener) EnterCreate_property_graph_statement(ctx *Create_property_graph_statementContext) {
}

// ExitCreate_property_graph_statement is called when production create_property_graph_statement is exited.
func (s *BaseGoogleSQLParserListener) ExitCreate_property_graph_statement(ctx *Create_property_graph_statementContext) {
}

// EnterOpt_edge_table_clause is called when production opt_edge_table_clause is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_edge_table_clause(ctx *Opt_edge_table_clauseContext) {}

// ExitOpt_edge_table_clause is called when production opt_edge_table_clause is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_edge_table_clause(ctx *Opt_edge_table_clauseContext) {}

// EnterElement_table_list is called when production element_table_list is entered.
func (s *BaseGoogleSQLParserListener) EnterElement_table_list(ctx *Element_table_listContext) {}

// ExitElement_table_list is called when production element_table_list is exited.
func (s *BaseGoogleSQLParserListener) ExitElement_table_list(ctx *Element_table_listContext) {}

// EnterElement_table_definition is called when production element_table_definition is entered.
func (s *BaseGoogleSQLParserListener) EnterElement_table_definition(ctx *Element_table_definitionContext) {
}

// ExitElement_table_definition is called when production element_table_definition is exited.
func (s *BaseGoogleSQLParserListener) ExitElement_table_definition(ctx *Element_table_definitionContext) {
}

// EnterOpt_label_and_properties_clause is called when production opt_label_and_properties_clause is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_label_and_properties_clause(ctx *Opt_label_and_properties_clauseContext) {
}

// ExitOpt_label_and_properties_clause is called when production opt_label_and_properties_clause is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_label_and_properties_clause(ctx *Opt_label_and_properties_clauseContext) {
}

// EnterLabel_and_properties_list is called when production label_and_properties_list is entered.
func (s *BaseGoogleSQLParserListener) EnterLabel_and_properties_list(ctx *Label_and_properties_listContext) {
}

// ExitLabel_and_properties_list is called when production label_and_properties_list is exited.
func (s *BaseGoogleSQLParserListener) ExitLabel_and_properties_list(ctx *Label_and_properties_listContext) {
}

// EnterLabel_and_properties is called when production label_and_properties is entered.
func (s *BaseGoogleSQLParserListener) EnterLabel_and_properties(ctx *Label_and_propertiesContext) {}

// ExitLabel_and_properties is called when production label_and_properties is exited.
func (s *BaseGoogleSQLParserListener) ExitLabel_and_properties(ctx *Label_and_propertiesContext) {}

// EnterProperties_clause is called when production properties_clause is entered.
func (s *BaseGoogleSQLParserListener) EnterProperties_clause(ctx *Properties_clauseContext) {}

// ExitProperties_clause is called when production properties_clause is exited.
func (s *BaseGoogleSQLParserListener) ExitProperties_clause(ctx *Properties_clauseContext) {}

// EnterDerived_property_list is called when production derived_property_list is entered.
func (s *BaseGoogleSQLParserListener) EnterDerived_property_list(ctx *Derived_property_listContext) {}

// ExitDerived_property_list is called when production derived_property_list is exited.
func (s *BaseGoogleSQLParserListener) ExitDerived_property_list(ctx *Derived_property_listContext) {}

// EnterDerived_property is called when production derived_property is entered.
func (s *BaseGoogleSQLParserListener) EnterDerived_property(ctx *Derived_propertyContext) {}

// ExitDerived_property is called when production derived_property is exited.
func (s *BaseGoogleSQLParserListener) ExitDerived_property(ctx *Derived_propertyContext) {}

// EnterOpt_except_column_list is called when production opt_except_column_list is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_except_column_list(ctx *Opt_except_column_listContext) {
}

// ExitOpt_except_column_list is called when production opt_except_column_list is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_except_column_list(ctx *Opt_except_column_listContext) {
}

// EnterProperties_all_columns is called when production properties_all_columns is entered.
func (s *BaseGoogleSQLParserListener) EnterProperties_all_columns(ctx *Properties_all_columnsContext) {
}

// ExitProperties_all_columns is called when production properties_all_columns is exited.
func (s *BaseGoogleSQLParserListener) ExitProperties_all_columns(ctx *Properties_all_columnsContext) {
}

// EnterOpt_dest_node_table_clause is called when production opt_dest_node_table_clause is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_dest_node_table_clause(ctx *Opt_dest_node_table_clauseContext) {
}

// ExitOpt_dest_node_table_clause is called when production opt_dest_node_table_clause is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_dest_node_table_clause(ctx *Opt_dest_node_table_clauseContext) {
}

// EnterOpt_source_node_table_clause is called when production opt_source_node_table_clause is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_source_node_table_clause(ctx *Opt_source_node_table_clauseContext) {
}

// ExitOpt_source_node_table_clause is called when production opt_source_node_table_clause is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_source_node_table_clause(ctx *Opt_source_node_table_clauseContext) {
}

// EnterOpt_key_clause is called when production opt_key_clause is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_key_clause(ctx *Opt_key_clauseContext) {}

// ExitOpt_key_clause is called when production opt_key_clause is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_key_clause(ctx *Opt_key_clauseContext) {}

// EnterCreate_model_statement is called when production create_model_statement is entered.
func (s *BaseGoogleSQLParserListener) EnterCreate_model_statement(ctx *Create_model_statementContext) {
}

// ExitCreate_model_statement is called when production create_model_statement is exited.
func (s *BaseGoogleSQLParserListener) ExitCreate_model_statement(ctx *Create_model_statementContext) {
}

// EnterOpt_input_output_clause is called when production opt_input_output_clause is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_input_output_clause(ctx *Opt_input_output_clauseContext) {
}

// ExitOpt_input_output_clause is called when production opt_input_output_clause is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_input_output_clause(ctx *Opt_input_output_clauseContext) {
}

// EnterOpt_transform_clause is called when production opt_transform_clause is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_transform_clause(ctx *Opt_transform_clauseContext) {}

// ExitOpt_transform_clause is called when production opt_transform_clause is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_transform_clause(ctx *Opt_transform_clauseContext) {}

// EnterOpt_as_query_or_aliased_query_list is called when production opt_as_query_or_aliased_query_list is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_as_query_or_aliased_query_list(ctx *Opt_as_query_or_aliased_query_listContext) {
}

// ExitOpt_as_query_or_aliased_query_list is called when production opt_as_query_or_aliased_query_list is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_as_query_or_aliased_query_list(ctx *Opt_as_query_or_aliased_query_listContext) {
}

// EnterAliased_query_list is called when production aliased_query_list is entered.
func (s *BaseGoogleSQLParserListener) EnterAliased_query_list(ctx *Aliased_query_listContext) {}

// ExitAliased_query_list is called when production aliased_query_list is exited.
func (s *BaseGoogleSQLParserListener) ExitAliased_query_list(ctx *Aliased_query_listContext) {}

// EnterAs_query is called when production as_query is entered.
func (s *BaseGoogleSQLParserListener) EnterAs_query(ctx *As_queryContext) {}

// ExitAs_query is called when production as_query is exited.
func (s *BaseGoogleSQLParserListener) ExitAs_query(ctx *As_queryContext) {}

// EnterCreate_external_table_function_statement is called when production create_external_table_function_statement is entered.
func (s *BaseGoogleSQLParserListener) EnterCreate_external_table_function_statement(ctx *Create_external_table_function_statementContext) {
}

// ExitCreate_external_table_function_statement is called when production create_external_table_function_statement is exited.
func (s *BaseGoogleSQLParserListener) ExitCreate_external_table_function_statement(ctx *Create_external_table_function_statementContext) {
}

// EnterCreate_external_table_statement is called when production create_external_table_statement is entered.
func (s *BaseGoogleSQLParserListener) EnterCreate_external_table_statement(ctx *Create_external_table_statementContext) {
}

// ExitCreate_external_table_statement is called when production create_external_table_statement is exited.
func (s *BaseGoogleSQLParserListener) ExitCreate_external_table_statement(ctx *Create_external_table_statementContext) {
}

// EnterOpt_default_collate_clause is called when production opt_default_collate_clause is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_default_collate_clause(ctx *Opt_default_collate_clauseContext) {
}

// ExitOpt_default_collate_clause is called when production opt_default_collate_clause is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_default_collate_clause(ctx *Opt_default_collate_clauseContext) {
}

// EnterOpt_like_path_expression is called when production opt_like_path_expression is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_like_path_expression(ctx *Opt_like_path_expressionContext) {
}

// ExitOpt_like_path_expression is called when production opt_like_path_expression is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_like_path_expression(ctx *Opt_like_path_expressionContext) {
}

// EnterCreate_row_access_policy_statement is called when production create_row_access_policy_statement is entered.
func (s *BaseGoogleSQLParserListener) EnterCreate_row_access_policy_statement(ctx *Create_row_access_policy_statementContext) {
}

// ExitCreate_row_access_policy_statement is called when production create_row_access_policy_statement is exited.
func (s *BaseGoogleSQLParserListener) ExitCreate_row_access_policy_statement(ctx *Create_row_access_policy_statementContext) {
}

// EnterFilter_using_clause is called when production filter_using_clause is entered.
func (s *BaseGoogleSQLParserListener) EnterFilter_using_clause(ctx *Filter_using_clauseContext) {}

// ExitFilter_using_clause is called when production filter_using_clause is exited.
func (s *BaseGoogleSQLParserListener) ExitFilter_using_clause(ctx *Filter_using_clauseContext) {}

// EnterCreate_row_access_policy_grant_to_clause is called when production create_row_access_policy_grant_to_clause is entered.
func (s *BaseGoogleSQLParserListener) EnterCreate_row_access_policy_grant_to_clause(ctx *Create_row_access_policy_grant_to_clauseContext) {
}

// ExitCreate_row_access_policy_grant_to_clause is called when production create_row_access_policy_grant_to_clause is exited.
func (s *BaseGoogleSQLParserListener) ExitCreate_row_access_policy_grant_to_clause(ctx *Create_row_access_policy_grant_to_clauseContext) {
}

// EnterCreate_privilege_restriction_statement is called when production create_privilege_restriction_statement is entered.
func (s *BaseGoogleSQLParserListener) EnterCreate_privilege_restriction_statement(ctx *Create_privilege_restriction_statementContext) {
}

// ExitCreate_privilege_restriction_statement is called when production create_privilege_restriction_statement is exited.
func (s *BaseGoogleSQLParserListener) ExitCreate_privilege_restriction_statement(ctx *Create_privilege_restriction_statementContext) {
}

// EnterRestrict_to_clause is called when production restrict_to_clause is entered.
func (s *BaseGoogleSQLParserListener) EnterRestrict_to_clause(ctx *Restrict_to_clauseContext) {}

// ExitRestrict_to_clause is called when production restrict_to_clause is exited.
func (s *BaseGoogleSQLParserListener) ExitRestrict_to_clause(ctx *Restrict_to_clauseContext) {}

// EnterPossibly_empty_grantee_list is called when production possibly_empty_grantee_list is entered.
func (s *BaseGoogleSQLParserListener) EnterPossibly_empty_grantee_list(ctx *Possibly_empty_grantee_listContext) {
}

// ExitPossibly_empty_grantee_list is called when production possibly_empty_grantee_list is exited.
func (s *BaseGoogleSQLParserListener) ExitPossibly_empty_grantee_list(ctx *Possibly_empty_grantee_listContext) {
}

// EnterCreate_index_statement is called when production create_index_statement is entered.
func (s *BaseGoogleSQLParserListener) EnterCreate_index_statement(ctx *Create_index_statementContext) {
}

// ExitCreate_index_statement is called when production create_index_statement is exited.
func (s *BaseGoogleSQLParserListener) ExitCreate_index_statement(ctx *Create_index_statementContext) {
}

// EnterOpt_create_index_statement_suffix is called when production opt_create_index_statement_suffix is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_create_index_statement_suffix(ctx *Opt_create_index_statement_suffixContext) {
}

// ExitOpt_create_index_statement_suffix is called when production opt_create_index_statement_suffix is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_create_index_statement_suffix(ctx *Opt_create_index_statement_suffixContext) {
}

// EnterSpanner_index_interleave_clause is called when production spanner_index_interleave_clause is entered.
func (s *BaseGoogleSQLParserListener) EnterSpanner_index_interleave_clause(ctx *Spanner_index_interleave_clauseContext) {
}

// ExitSpanner_index_interleave_clause is called when production spanner_index_interleave_clause is exited.
func (s *BaseGoogleSQLParserListener) ExitSpanner_index_interleave_clause(ctx *Spanner_index_interleave_clauseContext) {
}

// EnterIndex_storing_list is called when production index_storing_list is entered.
func (s *BaseGoogleSQLParserListener) EnterIndex_storing_list(ctx *Index_storing_listContext) {}

// ExitIndex_storing_list is called when production index_storing_list is exited.
func (s *BaseGoogleSQLParserListener) ExitIndex_storing_list(ctx *Index_storing_listContext) {}

// EnterIndex_storing_expression_list is called when production index_storing_expression_list is entered.
func (s *BaseGoogleSQLParserListener) EnterIndex_storing_expression_list(ctx *Index_storing_expression_listContext) {
}

// ExitIndex_storing_expression_list is called when production index_storing_expression_list is exited.
func (s *BaseGoogleSQLParserListener) ExitIndex_storing_expression_list(ctx *Index_storing_expression_listContext) {
}

// EnterIndex_order_by_and_options is called when production index_order_by_and_options is entered.
func (s *BaseGoogleSQLParserListener) EnterIndex_order_by_and_options(ctx *Index_order_by_and_optionsContext) {
}

// ExitIndex_order_by_and_options is called when production index_order_by_and_options is exited.
func (s *BaseGoogleSQLParserListener) ExitIndex_order_by_and_options(ctx *Index_order_by_and_optionsContext) {
}

// EnterIndex_all_columns is called when production index_all_columns is entered.
func (s *BaseGoogleSQLParserListener) EnterIndex_all_columns(ctx *Index_all_columnsContext) {}

// ExitIndex_all_columns is called when production index_all_columns is exited.
func (s *BaseGoogleSQLParserListener) ExitIndex_all_columns(ctx *Index_all_columnsContext) {}

// EnterOpt_with_column_options is called when production opt_with_column_options is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_with_column_options(ctx *Opt_with_column_optionsContext) {
}

// ExitOpt_with_column_options is called when production opt_with_column_options is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_with_column_options(ctx *Opt_with_column_optionsContext) {
}

// EnterAll_column_column_options is called when production all_column_column_options is entered.
func (s *BaseGoogleSQLParserListener) EnterAll_column_column_options(ctx *All_column_column_optionsContext) {
}

// ExitAll_column_column_options is called when production all_column_column_options is exited.
func (s *BaseGoogleSQLParserListener) ExitAll_column_column_options(ctx *All_column_column_optionsContext) {
}

// EnterColumn_ordering_and_options_expr is called when production column_ordering_and_options_expr is entered.
func (s *BaseGoogleSQLParserListener) EnterColumn_ordering_and_options_expr(ctx *Column_ordering_and_options_exprContext) {
}

// ExitColumn_ordering_and_options_expr is called when production column_ordering_and_options_expr is exited.
func (s *BaseGoogleSQLParserListener) ExitColumn_ordering_and_options_expr(ctx *Column_ordering_and_options_exprContext) {
}

// EnterIndex_unnest_expression_list is called when production index_unnest_expression_list is entered.
func (s *BaseGoogleSQLParserListener) EnterIndex_unnest_expression_list(ctx *Index_unnest_expression_listContext) {
}

// ExitIndex_unnest_expression_list is called when production index_unnest_expression_list is exited.
func (s *BaseGoogleSQLParserListener) ExitIndex_unnest_expression_list(ctx *Index_unnest_expression_listContext) {
}

// EnterUnnest_expression_with_opt_alias_and_offset is called when production unnest_expression_with_opt_alias_and_offset is entered.
func (s *BaseGoogleSQLParserListener) EnterUnnest_expression_with_opt_alias_and_offset(ctx *Unnest_expression_with_opt_alias_and_offsetContext) {
}

// ExitUnnest_expression_with_opt_alias_and_offset is called when production unnest_expression_with_opt_alias_and_offset is exited.
func (s *BaseGoogleSQLParserListener) ExitUnnest_expression_with_opt_alias_and_offset(ctx *Unnest_expression_with_opt_alias_and_offsetContext) {
}

// EnterOn_path_expression is called when production on_path_expression is entered.
func (s *BaseGoogleSQLParserListener) EnterOn_path_expression(ctx *On_path_expressionContext) {}

// ExitOn_path_expression is called when production on_path_expression is exited.
func (s *BaseGoogleSQLParserListener) ExitOn_path_expression(ctx *On_path_expressionContext) {}

// EnterIndex_type is called when production index_type is entered.
func (s *BaseGoogleSQLParserListener) EnterIndex_type(ctx *Index_typeContext) {}

// ExitIndex_type is called when production index_type is exited.
func (s *BaseGoogleSQLParserListener) ExitIndex_type(ctx *Index_typeContext) {}

// EnterOpt_spanner_null_filtered is called when production opt_spanner_null_filtered is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_spanner_null_filtered(ctx *Opt_spanner_null_filteredContext) {
}

// ExitOpt_spanner_null_filtered is called when production opt_spanner_null_filtered is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_spanner_null_filtered(ctx *Opt_spanner_null_filteredContext) {
}

// EnterCreate_procedure_statement is called when production create_procedure_statement is entered.
func (s *BaseGoogleSQLParserListener) EnterCreate_procedure_statement(ctx *Create_procedure_statementContext) {
}

// ExitCreate_procedure_statement is called when production create_procedure_statement is exited.
func (s *BaseGoogleSQLParserListener) ExitCreate_procedure_statement(ctx *Create_procedure_statementContext) {
}

// EnterBegin_end_block_or_language_as_code is called when production begin_end_block_or_language_as_code is entered.
func (s *BaseGoogleSQLParserListener) EnterBegin_end_block_or_language_as_code(ctx *Begin_end_block_or_language_as_codeContext) {
}

// ExitBegin_end_block_or_language_as_code is called when production begin_end_block_or_language_as_code is exited.
func (s *BaseGoogleSQLParserListener) ExitBegin_end_block_or_language_as_code(ctx *Begin_end_block_or_language_as_codeContext) {
}

// EnterBegin_end_block is called when production begin_end_block is entered.
func (s *BaseGoogleSQLParserListener) EnterBegin_end_block(ctx *Begin_end_blockContext) {}

// ExitBegin_end_block is called when production begin_end_block is exited.
func (s *BaseGoogleSQLParserListener) ExitBegin_end_block(ctx *Begin_end_blockContext) {}

// EnterOpt_exception_handler is called when production opt_exception_handler is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_exception_handler(ctx *Opt_exception_handlerContext) {}

// ExitOpt_exception_handler is called when production opt_exception_handler is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_exception_handler(ctx *Opt_exception_handlerContext) {}

// EnterStatement_list is called when production statement_list is entered.
func (s *BaseGoogleSQLParserListener) EnterStatement_list(ctx *Statement_listContext) {}

// ExitStatement_list is called when production statement_list is exited.
func (s *BaseGoogleSQLParserListener) ExitStatement_list(ctx *Statement_listContext) {}

// EnterUnterminated_non_empty_statement_list is called when production unterminated_non_empty_statement_list is entered.
func (s *BaseGoogleSQLParserListener) EnterUnterminated_non_empty_statement_list(ctx *Unterminated_non_empty_statement_listContext) {
}

// ExitUnterminated_non_empty_statement_list is called when production unterminated_non_empty_statement_list is exited.
func (s *BaseGoogleSQLParserListener) ExitUnterminated_non_empty_statement_list(ctx *Unterminated_non_empty_statement_listContext) {
}

// EnterUnterminated_statement is called when production unterminated_statement is entered.
func (s *BaseGoogleSQLParserListener) EnterUnterminated_statement(ctx *Unterminated_statementContext) {
}

// ExitUnterminated_statement is called when production unterminated_statement is exited.
func (s *BaseGoogleSQLParserListener) ExitUnterminated_statement(ctx *Unterminated_statementContext) {
}

// EnterUnterminated_script_statement is called when production unterminated_script_statement is entered.
func (s *BaseGoogleSQLParserListener) EnterUnterminated_script_statement(ctx *Unterminated_script_statementContext) {
}

// ExitUnterminated_script_statement is called when production unterminated_script_statement is exited.
func (s *BaseGoogleSQLParserListener) ExitUnterminated_script_statement(ctx *Unterminated_script_statementContext) {
}

// EnterLabel is called when production label is entered.
func (s *BaseGoogleSQLParserListener) EnterLabel(ctx *LabelContext) {}

// ExitLabel is called when production label is exited.
func (s *BaseGoogleSQLParserListener) ExitLabel(ctx *LabelContext) {}

// EnterUnterminated_unlabeled_script_statement is called when production unterminated_unlabeled_script_statement is entered.
func (s *BaseGoogleSQLParserListener) EnterUnterminated_unlabeled_script_statement(ctx *Unterminated_unlabeled_script_statementContext) {
}

// ExitUnterminated_unlabeled_script_statement is called when production unterminated_unlabeled_script_statement is exited.
func (s *BaseGoogleSQLParserListener) ExitUnterminated_unlabeled_script_statement(ctx *Unterminated_unlabeled_script_statementContext) {
}

// EnterFor_in_statement is called when production for_in_statement is entered.
func (s *BaseGoogleSQLParserListener) EnterFor_in_statement(ctx *For_in_statementContext) {}

// ExitFor_in_statement is called when production for_in_statement is exited.
func (s *BaseGoogleSQLParserListener) ExitFor_in_statement(ctx *For_in_statementContext) {}

// EnterRepeat_statement is called when production repeat_statement is entered.
func (s *BaseGoogleSQLParserListener) EnterRepeat_statement(ctx *Repeat_statementContext) {}

// ExitRepeat_statement is called when production repeat_statement is exited.
func (s *BaseGoogleSQLParserListener) ExitRepeat_statement(ctx *Repeat_statementContext) {}

// EnterUntil_clause is called when production until_clause is entered.
func (s *BaseGoogleSQLParserListener) EnterUntil_clause(ctx *Until_clauseContext) {}

// ExitUntil_clause is called when production until_clause is exited.
func (s *BaseGoogleSQLParserListener) ExitUntil_clause(ctx *Until_clauseContext) {}

// EnterLoop_statement is called when production loop_statement is entered.
func (s *BaseGoogleSQLParserListener) EnterLoop_statement(ctx *Loop_statementContext) {}

// ExitLoop_statement is called when production loop_statement is exited.
func (s *BaseGoogleSQLParserListener) ExitLoop_statement(ctx *Loop_statementContext) {}

// EnterWhile_statement is called when production while_statement is entered.
func (s *BaseGoogleSQLParserListener) EnterWhile_statement(ctx *While_statementContext) {}

// ExitWhile_statement is called when production while_statement is exited.
func (s *BaseGoogleSQLParserListener) ExitWhile_statement(ctx *While_statementContext) {}

// EnterRaise_statement is called when production raise_statement is entered.
func (s *BaseGoogleSQLParserListener) EnterRaise_statement(ctx *Raise_statementContext) {}

// ExitRaise_statement is called when production raise_statement is exited.
func (s *BaseGoogleSQLParserListener) ExitRaise_statement(ctx *Raise_statementContext) {}

// EnterReturn_statement is called when production return_statement is entered.
func (s *BaseGoogleSQLParserListener) EnterReturn_statement(ctx *Return_statementContext) {}

// ExitReturn_statement is called when production return_statement is exited.
func (s *BaseGoogleSQLParserListener) ExitReturn_statement(ctx *Return_statementContext) {}

// EnterContinue_statement is called when production continue_statement is entered.
func (s *BaseGoogleSQLParserListener) EnterContinue_statement(ctx *Continue_statementContext) {}

// ExitContinue_statement is called when production continue_statement is exited.
func (s *BaseGoogleSQLParserListener) ExitContinue_statement(ctx *Continue_statementContext) {}

// EnterVariable_declaration is called when production variable_declaration is entered.
func (s *BaseGoogleSQLParserListener) EnterVariable_declaration(ctx *Variable_declarationContext) {}

// ExitVariable_declaration is called when production variable_declaration is exited.
func (s *BaseGoogleSQLParserListener) ExitVariable_declaration(ctx *Variable_declarationContext) {}

// EnterBreak_statement is called when production break_statement is entered.
func (s *BaseGoogleSQLParserListener) EnterBreak_statement(ctx *Break_statementContext) {}

// ExitBreak_statement is called when production break_statement is exited.
func (s *BaseGoogleSQLParserListener) ExitBreak_statement(ctx *Break_statementContext) {}

// EnterCase_statement is called when production case_statement is entered.
func (s *BaseGoogleSQLParserListener) EnterCase_statement(ctx *Case_statementContext) {}

// ExitCase_statement is called when production case_statement is exited.
func (s *BaseGoogleSQLParserListener) ExitCase_statement(ctx *Case_statementContext) {}

// EnterWhen_then_clauses is called when production when_then_clauses is entered.
func (s *BaseGoogleSQLParserListener) EnterWhen_then_clauses(ctx *When_then_clausesContext) {}

// ExitWhen_then_clauses is called when production when_then_clauses is exited.
func (s *BaseGoogleSQLParserListener) ExitWhen_then_clauses(ctx *When_then_clausesContext) {}

// EnterIf_statement is called when production if_statement is entered.
func (s *BaseGoogleSQLParserListener) EnterIf_statement(ctx *If_statementContext) {}

// ExitIf_statement is called when production if_statement is exited.
func (s *BaseGoogleSQLParserListener) ExitIf_statement(ctx *If_statementContext) {}

// EnterElseif_clauses is called when production elseif_clauses is entered.
func (s *BaseGoogleSQLParserListener) EnterElseif_clauses(ctx *Elseif_clausesContext) {}

// ExitElseif_clauses is called when production elseif_clauses is exited.
func (s *BaseGoogleSQLParserListener) ExitElseif_clauses(ctx *Elseif_clausesContext) {}

// EnterOpt_else is called when production opt_else is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_else(ctx *Opt_elseContext) {}

// ExitOpt_else is called when production opt_else is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_else(ctx *Opt_elseContext) {}

// EnterOpt_as_code is called when production opt_as_code is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_as_code(ctx *Opt_as_codeContext) {}

// ExitOpt_as_code is called when production opt_as_code is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_as_code(ctx *Opt_as_codeContext) {}

// EnterOpt_external_security_clause is called when production opt_external_security_clause is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_external_security_clause(ctx *Opt_external_security_clauseContext) {
}

// ExitOpt_external_security_clause is called when production opt_external_security_clause is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_external_security_clause(ctx *Opt_external_security_clauseContext) {
}

// EnterExternal_security_clause_kind is called when production external_security_clause_kind is entered.
func (s *BaseGoogleSQLParserListener) EnterExternal_security_clause_kind(ctx *External_security_clause_kindContext) {
}

// ExitExternal_security_clause_kind is called when production external_security_clause_kind is exited.
func (s *BaseGoogleSQLParserListener) ExitExternal_security_clause_kind(ctx *External_security_clause_kindContext) {
}

// EnterProcedure_parameters is called when production procedure_parameters is entered.
func (s *BaseGoogleSQLParserListener) EnterProcedure_parameters(ctx *Procedure_parametersContext) {}

// ExitProcedure_parameters is called when production procedure_parameters is exited.
func (s *BaseGoogleSQLParserListener) ExitProcedure_parameters(ctx *Procedure_parametersContext) {}

// EnterProcedure_parameter is called when production procedure_parameter is entered.
func (s *BaseGoogleSQLParserListener) EnterProcedure_parameter(ctx *Procedure_parameterContext) {}

// ExitProcedure_parameter is called when production procedure_parameter is exited.
func (s *BaseGoogleSQLParserListener) ExitProcedure_parameter(ctx *Procedure_parameterContext) {}

// EnterProcedure_parameter_termination is called when production procedure_parameter_termination is entered.
func (s *BaseGoogleSQLParserListener) EnterProcedure_parameter_termination(ctx *Procedure_parameter_terminationContext) {
}

// ExitProcedure_parameter_termination is called when production procedure_parameter_termination is exited.
func (s *BaseGoogleSQLParserListener) ExitProcedure_parameter_termination(ctx *Procedure_parameter_terminationContext) {
}

// EnterOpt_procedure_parameter_mode is called when production opt_procedure_parameter_mode is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_procedure_parameter_mode(ctx *Opt_procedure_parameter_modeContext) {
}

// ExitOpt_procedure_parameter_mode is called when production opt_procedure_parameter_mode is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_procedure_parameter_mode(ctx *Opt_procedure_parameter_modeContext) {
}

// EnterCreate_function_statement is called when production create_function_statement is entered.
func (s *BaseGoogleSQLParserListener) EnterCreate_function_statement(ctx *Create_function_statementContext) {
}

// ExitCreate_function_statement is called when production create_function_statement is exited.
func (s *BaseGoogleSQLParserListener) ExitCreate_function_statement(ctx *Create_function_statementContext) {
}

// EnterOpt_determinism_level is called when production opt_determinism_level is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_determinism_level(ctx *Opt_determinism_levelContext) {}

// ExitOpt_determinism_level is called when production opt_determinism_level is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_determinism_level(ctx *Opt_determinism_levelContext) {}

// EnterOpt_sql_security_clause is called when production opt_sql_security_clause is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_sql_security_clause(ctx *Opt_sql_security_clauseContext) {
}

// ExitOpt_sql_security_clause is called when production opt_sql_security_clause is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_sql_security_clause(ctx *Opt_sql_security_clauseContext) {
}

// EnterSql_security_clause_kind is called when production sql_security_clause_kind is entered.
func (s *BaseGoogleSQLParserListener) EnterSql_security_clause_kind(ctx *Sql_security_clause_kindContext) {
}

// ExitSql_security_clause_kind is called when production sql_security_clause_kind is exited.
func (s *BaseGoogleSQLParserListener) ExitSql_security_clause_kind(ctx *Sql_security_clause_kindContext) {
}

// EnterAs_sql_function_body_or_string is called when production as_sql_function_body_or_string is entered.
func (s *BaseGoogleSQLParserListener) EnterAs_sql_function_body_or_string(ctx *As_sql_function_body_or_stringContext) {
}

// ExitAs_sql_function_body_or_string is called when production as_sql_function_body_or_string is exited.
func (s *BaseGoogleSQLParserListener) ExitAs_sql_function_body_or_string(ctx *As_sql_function_body_or_stringContext) {
}

// EnterSql_function_body is called when production sql_function_body is entered.
func (s *BaseGoogleSQLParserListener) EnterSql_function_body(ctx *Sql_function_bodyContext) {}

// ExitSql_function_body is called when production sql_function_body is exited.
func (s *BaseGoogleSQLParserListener) ExitSql_function_body(ctx *Sql_function_bodyContext) {}

// EnterUnordered_options_body is called when production unordered_options_body is entered.
func (s *BaseGoogleSQLParserListener) EnterUnordered_options_body(ctx *Unordered_options_bodyContext) {
}

// ExitUnordered_options_body is called when production unordered_options_body is exited.
func (s *BaseGoogleSQLParserListener) ExitUnordered_options_body(ctx *Unordered_options_bodyContext) {
}

// EnterOpt_language_or_remote_with_connection is called when production opt_language_or_remote_with_connection is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_language_or_remote_with_connection(ctx *Opt_language_or_remote_with_connectionContext) {
}

// ExitOpt_language_or_remote_with_connection is called when production opt_language_or_remote_with_connection is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_language_or_remote_with_connection(ctx *Opt_language_or_remote_with_connectionContext) {
}

// EnterLanguage is called when production language is entered.
func (s *BaseGoogleSQLParserListener) EnterLanguage(ctx *LanguageContext) {}

// ExitLanguage is called when production language is exited.
func (s *BaseGoogleSQLParserListener) ExitLanguage(ctx *LanguageContext) {}

// EnterRemote_with_connection_clause is called when production remote_with_connection_clause is entered.
func (s *BaseGoogleSQLParserListener) EnterRemote_with_connection_clause(ctx *Remote_with_connection_clauseContext) {
}

// ExitRemote_with_connection_clause is called when production remote_with_connection_clause is exited.
func (s *BaseGoogleSQLParserListener) ExitRemote_with_connection_clause(ctx *Remote_with_connection_clauseContext) {
}

// EnterWith_connection_clause is called when production with_connection_clause is entered.
func (s *BaseGoogleSQLParserListener) EnterWith_connection_clause(ctx *With_connection_clauseContext) {
}

// ExitWith_connection_clause is called when production with_connection_clause is exited.
func (s *BaseGoogleSQLParserListener) ExitWith_connection_clause(ctx *With_connection_clauseContext) {
}

// EnterOpt_function_returns is called when production opt_function_returns is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_function_returns(ctx *Opt_function_returnsContext) {}

// ExitOpt_function_returns is called when production opt_function_returns is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_function_returns(ctx *Opt_function_returnsContext) {}

// EnterOpt_returns is called when production opt_returns is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_returns(ctx *Opt_returnsContext) {}

// ExitOpt_returns is called when production opt_returns is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_returns(ctx *Opt_returnsContext) {}

// EnterFunction_declaration is called when production function_declaration is entered.
func (s *BaseGoogleSQLParserListener) EnterFunction_declaration(ctx *Function_declarationContext) {}

// ExitFunction_declaration is called when production function_declaration is exited.
func (s *BaseGoogleSQLParserListener) ExitFunction_declaration(ctx *Function_declarationContext) {}

// EnterFunction_parameters is called when production function_parameters is entered.
func (s *BaseGoogleSQLParserListener) EnterFunction_parameters(ctx *Function_parametersContext) {}

// ExitFunction_parameters is called when production function_parameters is exited.
func (s *BaseGoogleSQLParserListener) ExitFunction_parameters(ctx *Function_parametersContext) {}

// EnterFunction_parameter is called when production function_parameter is entered.
func (s *BaseGoogleSQLParserListener) EnterFunction_parameter(ctx *Function_parameterContext) {}

// ExitFunction_parameter is called when production function_parameter is exited.
func (s *BaseGoogleSQLParserListener) ExitFunction_parameter(ctx *Function_parameterContext) {}

// EnterOpt_not_aggregate is called when production opt_not_aggregate is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_not_aggregate(ctx *Opt_not_aggregateContext) {}

// ExitOpt_not_aggregate is called when production opt_not_aggregate is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_not_aggregate(ctx *Opt_not_aggregateContext) {}

// EnterOpt_default_expression is called when production opt_default_expression is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_default_expression(ctx *Opt_default_expressionContext) {
}

// ExitOpt_default_expression is called when production opt_default_expression is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_default_expression(ctx *Opt_default_expressionContext) {
}

// EnterType_or_tvf_schema is called when production type_or_tvf_schema is entered.
func (s *BaseGoogleSQLParserListener) EnterType_or_tvf_schema(ctx *Type_or_tvf_schemaContext) {}

// ExitType_or_tvf_schema is called when production type_or_tvf_schema is exited.
func (s *BaseGoogleSQLParserListener) ExitType_or_tvf_schema(ctx *Type_or_tvf_schemaContext) {}

// EnterTvf_schema is called when production tvf_schema is entered.
func (s *BaseGoogleSQLParserListener) EnterTvf_schema(ctx *Tvf_schemaContext) {}

// ExitTvf_schema is called when production tvf_schema is exited.
func (s *BaseGoogleSQLParserListener) ExitTvf_schema(ctx *Tvf_schemaContext) {}

// EnterTvf_schema_column is called when production tvf_schema_column is entered.
func (s *BaseGoogleSQLParserListener) EnterTvf_schema_column(ctx *Tvf_schema_columnContext) {}

// ExitTvf_schema_column is called when production tvf_schema_column is exited.
func (s *BaseGoogleSQLParserListener) ExitTvf_schema_column(ctx *Tvf_schema_columnContext) {}

// EnterTemplated_parameter_type is called when production templated_parameter_type is entered.
func (s *BaseGoogleSQLParserListener) EnterTemplated_parameter_type(ctx *Templated_parameter_typeContext) {
}

// ExitTemplated_parameter_type is called when production templated_parameter_type is exited.
func (s *BaseGoogleSQLParserListener) ExitTemplated_parameter_type(ctx *Templated_parameter_typeContext) {
}

// EnterTemplated_parameter_kind is called when production templated_parameter_kind is entered.
func (s *BaseGoogleSQLParserListener) EnterTemplated_parameter_kind(ctx *Templated_parameter_kindContext) {
}

// ExitTemplated_parameter_kind is called when production templated_parameter_kind is exited.
func (s *BaseGoogleSQLParserListener) ExitTemplated_parameter_kind(ctx *Templated_parameter_kindContext) {
}

// EnterOpt_aggregate is called when production opt_aggregate is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_aggregate(ctx *Opt_aggregateContext) {}

// ExitOpt_aggregate is called when production opt_aggregate is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_aggregate(ctx *Opt_aggregateContext) {}

// EnterCreate_database_statement is called when production create_database_statement is entered.
func (s *BaseGoogleSQLParserListener) EnterCreate_database_statement(ctx *Create_database_statementContext) {
}

// ExitCreate_database_statement is called when production create_database_statement is exited.
func (s *BaseGoogleSQLParserListener) ExitCreate_database_statement(ctx *Create_database_statementContext) {
}

// EnterCreate_connection_statement is called when production create_connection_statement is entered.
func (s *BaseGoogleSQLParserListener) EnterCreate_connection_statement(ctx *Create_connection_statementContext) {
}

// ExitCreate_connection_statement is called when production create_connection_statement is exited.
func (s *BaseGoogleSQLParserListener) ExitCreate_connection_statement(ctx *Create_connection_statementContext) {
}

// EnterCreate_constant_statement is called when production create_constant_statement is entered.
func (s *BaseGoogleSQLParserListener) EnterCreate_constant_statement(ctx *Create_constant_statementContext) {
}

// ExitCreate_constant_statement is called when production create_constant_statement is exited.
func (s *BaseGoogleSQLParserListener) ExitCreate_constant_statement(ctx *Create_constant_statementContext) {
}

// EnterOpt_or_replace is called when production opt_or_replace is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_or_replace(ctx *Opt_or_replaceContext) {}

// ExitOpt_or_replace is called when production opt_or_replace is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_or_replace(ctx *Opt_or_replaceContext) {}

// EnterOpt_create_scope is called when production opt_create_scope is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_create_scope(ctx *Opt_create_scopeContext) {}

// ExitOpt_create_scope is called when production opt_create_scope is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_create_scope(ctx *Opt_create_scopeContext) {}

// EnterRun_batch_statement is called when production run_batch_statement is entered.
func (s *BaseGoogleSQLParserListener) EnterRun_batch_statement(ctx *Run_batch_statementContext) {}

// ExitRun_batch_statement is called when production run_batch_statement is exited.
func (s *BaseGoogleSQLParserListener) ExitRun_batch_statement(ctx *Run_batch_statementContext) {}

// EnterAbort_batch_statement is called when production abort_batch_statement is entered.
func (s *BaseGoogleSQLParserListener) EnterAbort_batch_statement(ctx *Abort_batch_statementContext) {}

// ExitAbort_batch_statement is called when production abort_batch_statement is exited.
func (s *BaseGoogleSQLParserListener) ExitAbort_batch_statement(ctx *Abort_batch_statementContext) {}

// EnterStart_batch_statement is called when production start_batch_statement is entered.
func (s *BaseGoogleSQLParserListener) EnterStart_batch_statement(ctx *Start_batch_statementContext) {}

// ExitStart_batch_statement is called when production start_batch_statement is exited.
func (s *BaseGoogleSQLParserListener) ExitStart_batch_statement(ctx *Start_batch_statementContext) {}

// EnterRollback_statement is called when production rollback_statement is entered.
func (s *BaseGoogleSQLParserListener) EnterRollback_statement(ctx *Rollback_statementContext) {}

// ExitRollback_statement is called when production rollback_statement is exited.
func (s *BaseGoogleSQLParserListener) ExitRollback_statement(ctx *Rollback_statementContext) {}

// EnterCommit_statement is called when production commit_statement is entered.
func (s *BaseGoogleSQLParserListener) EnterCommit_statement(ctx *Commit_statementContext) {}

// ExitCommit_statement is called when production commit_statement is exited.
func (s *BaseGoogleSQLParserListener) ExitCommit_statement(ctx *Commit_statementContext) {}

// EnterSet_statement is called when production set_statement is entered.
func (s *BaseGoogleSQLParserListener) EnterSet_statement(ctx *Set_statementContext) {}

// ExitSet_statement is called when production set_statement is exited.
func (s *BaseGoogleSQLParserListener) ExitSet_statement(ctx *Set_statementContext) {}

// EnterIdentifier_list is called when production identifier_list is entered.
func (s *BaseGoogleSQLParserListener) EnterIdentifier_list(ctx *Identifier_listContext) {}

// ExitIdentifier_list is called when production identifier_list is exited.
func (s *BaseGoogleSQLParserListener) ExitIdentifier_list(ctx *Identifier_listContext) {}

// EnterBegin_statement is called when production begin_statement is entered.
func (s *BaseGoogleSQLParserListener) EnterBegin_statement(ctx *Begin_statementContext) {}

// ExitBegin_statement is called when production begin_statement is exited.
func (s *BaseGoogleSQLParserListener) ExitBegin_statement(ctx *Begin_statementContext) {}

// EnterBegin_transaction_keywords is called when production begin_transaction_keywords is entered.
func (s *BaseGoogleSQLParserListener) EnterBegin_transaction_keywords(ctx *Begin_transaction_keywordsContext) {
}

// ExitBegin_transaction_keywords is called when production begin_transaction_keywords is exited.
func (s *BaseGoogleSQLParserListener) ExitBegin_transaction_keywords(ctx *Begin_transaction_keywordsContext) {
}

// EnterTransaction_mode_list is called when production transaction_mode_list is entered.
func (s *BaseGoogleSQLParserListener) EnterTransaction_mode_list(ctx *Transaction_mode_listContext) {}

// ExitTransaction_mode_list is called when production transaction_mode_list is exited.
func (s *BaseGoogleSQLParserListener) ExitTransaction_mode_list(ctx *Transaction_mode_listContext) {}

// EnterTransaction_mode is called when production transaction_mode is entered.
func (s *BaseGoogleSQLParserListener) EnterTransaction_mode(ctx *Transaction_modeContext) {}

// ExitTransaction_mode is called when production transaction_mode is exited.
func (s *BaseGoogleSQLParserListener) ExitTransaction_mode(ctx *Transaction_modeContext) {}

// EnterTruncate_statement is called when production truncate_statement is entered.
func (s *BaseGoogleSQLParserListener) EnterTruncate_statement(ctx *Truncate_statementContext) {}

// ExitTruncate_statement is called when production truncate_statement is exited.
func (s *BaseGoogleSQLParserListener) ExitTruncate_statement(ctx *Truncate_statementContext) {}

// EnterMerge_statement is called when production merge_statement is entered.
func (s *BaseGoogleSQLParserListener) EnterMerge_statement(ctx *Merge_statementContext) {}

// ExitMerge_statement is called when production merge_statement is exited.
func (s *BaseGoogleSQLParserListener) ExitMerge_statement(ctx *Merge_statementContext) {}

// EnterMerge_source is called when production merge_source is entered.
func (s *BaseGoogleSQLParserListener) EnterMerge_source(ctx *Merge_sourceContext) {}

// ExitMerge_source is called when production merge_source is exited.
func (s *BaseGoogleSQLParserListener) ExitMerge_source(ctx *Merge_sourceContext) {}

// EnterMerge_when_clause is called when production merge_when_clause is entered.
func (s *BaseGoogleSQLParserListener) EnterMerge_when_clause(ctx *Merge_when_clauseContext) {}

// ExitMerge_when_clause is called when production merge_when_clause is exited.
func (s *BaseGoogleSQLParserListener) ExitMerge_when_clause(ctx *Merge_when_clauseContext) {}

// EnterMerge_action is called when production merge_action is entered.
func (s *BaseGoogleSQLParserListener) EnterMerge_action(ctx *Merge_actionContext) {}

// ExitMerge_action is called when production merge_action is exited.
func (s *BaseGoogleSQLParserListener) ExitMerge_action(ctx *Merge_actionContext) {}

// EnterMerge_insert_value_list_or_source_row is called when production merge_insert_value_list_or_source_row is entered.
func (s *BaseGoogleSQLParserListener) EnterMerge_insert_value_list_or_source_row(ctx *Merge_insert_value_list_or_source_rowContext) {
}

// ExitMerge_insert_value_list_or_source_row is called when production merge_insert_value_list_or_source_row is exited.
func (s *BaseGoogleSQLParserListener) ExitMerge_insert_value_list_or_source_row(ctx *Merge_insert_value_list_or_source_rowContext) {
}

// EnterBy_target is called when production by_target is entered.
func (s *BaseGoogleSQLParserListener) EnterBy_target(ctx *By_targetContext) {}

// ExitBy_target is called when production by_target is exited.
func (s *BaseGoogleSQLParserListener) ExitBy_target(ctx *By_targetContext) {}

// EnterOpt_and_expression is called when production opt_and_expression is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_and_expression(ctx *Opt_and_expressionContext) {}

// ExitOpt_and_expression is called when production opt_and_expression is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_and_expression(ctx *Opt_and_expressionContext) {}

// EnterStatement_level_hint is called when production statement_level_hint is entered.
func (s *BaseGoogleSQLParserListener) EnterStatement_level_hint(ctx *Statement_level_hintContext) {}

// ExitStatement_level_hint is called when production statement_level_hint is exited.
func (s *BaseGoogleSQLParserListener) ExitStatement_level_hint(ctx *Statement_level_hintContext) {}

// EnterQuery_statement is called when production query_statement is entered.
func (s *BaseGoogleSQLParserListener) EnterQuery_statement(ctx *Query_statementContext) {}

// ExitQuery_statement is called when production query_statement is exited.
func (s *BaseGoogleSQLParserListener) ExitQuery_statement(ctx *Query_statementContext) {}

// EnterDml_statement is called when production dml_statement is entered.
func (s *BaseGoogleSQLParserListener) EnterDml_statement(ctx *Dml_statementContext) {}

// ExitDml_statement is called when production dml_statement is exited.
func (s *BaseGoogleSQLParserListener) ExitDml_statement(ctx *Dml_statementContext) {}

// EnterUpdate_statement is called when production update_statement is entered.
func (s *BaseGoogleSQLParserListener) EnterUpdate_statement(ctx *Update_statementContext) {}

// ExitUpdate_statement is called when production update_statement is exited.
func (s *BaseGoogleSQLParserListener) ExitUpdate_statement(ctx *Update_statementContext) {}

// EnterDelete_statement is called when production delete_statement is entered.
func (s *BaseGoogleSQLParserListener) EnterDelete_statement(ctx *Delete_statementContext) {}

// ExitDelete_statement is called when production delete_statement is exited.
func (s *BaseGoogleSQLParserListener) ExitDelete_statement(ctx *Delete_statementContext) {}

// EnterInsert_statement is called when production insert_statement is entered.
func (s *BaseGoogleSQLParserListener) EnterInsert_statement(ctx *Insert_statementContext) {}

// ExitInsert_statement is called when production insert_statement is exited.
func (s *BaseGoogleSQLParserListener) ExitInsert_statement(ctx *Insert_statementContext) {}

// EnterOn_conflict_clause is called when production on_conflict_clause is entered.
func (s *BaseGoogleSQLParserListener) EnterOn_conflict_clause(ctx *On_conflict_clauseContext) {}

// ExitOn_conflict_clause is called when production on_conflict_clause is exited.
func (s *BaseGoogleSQLParserListener) ExitOn_conflict_clause(ctx *On_conflict_clauseContext) {}

// EnterOpt_where_expression is called when production opt_where_expression is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_where_expression(ctx *Opt_where_expressionContext) {}

// ExitOpt_where_expression is called when production opt_where_expression is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_where_expression(ctx *Opt_where_expressionContext) {}

// EnterOpt_conflict_target is called when production opt_conflict_target is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_conflict_target(ctx *Opt_conflict_targetContext) {}

// ExitOpt_conflict_target is called when production opt_conflict_target is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_conflict_target(ctx *Opt_conflict_targetContext) {}

// EnterUpdate_item_list is called when production update_item_list is entered.
func (s *BaseGoogleSQLParserListener) EnterUpdate_item_list(ctx *Update_item_listContext) {}

// ExitUpdate_item_list is called when production update_item_list is exited.
func (s *BaseGoogleSQLParserListener) ExitUpdate_item_list(ctx *Update_item_listContext) {}

// EnterUpdate_item is called when production update_item is entered.
func (s *BaseGoogleSQLParserListener) EnterUpdate_item(ctx *Update_itemContext) {}

// ExitUpdate_item is called when production update_item is exited.
func (s *BaseGoogleSQLParserListener) ExitUpdate_item(ctx *Update_itemContext) {}

// EnterUpdate_set_value is called when production update_set_value is entered.
func (s *BaseGoogleSQLParserListener) EnterUpdate_set_value(ctx *Update_set_valueContext) {}

// ExitUpdate_set_value is called when production update_set_value is exited.
func (s *BaseGoogleSQLParserListener) ExitUpdate_set_value(ctx *Update_set_valueContext) {}

// EnterNested_dml_statement is called when production nested_dml_statement is entered.
func (s *BaseGoogleSQLParserListener) EnterNested_dml_statement(ctx *Nested_dml_statementContext) {}

// ExitNested_dml_statement is called when production nested_dml_statement is exited.
func (s *BaseGoogleSQLParserListener) ExitNested_dml_statement(ctx *Nested_dml_statementContext) {}

// EnterInsert_values_list_or_table_clause is called when production insert_values_list_or_table_clause is entered.
func (s *BaseGoogleSQLParserListener) EnterInsert_values_list_or_table_clause(ctx *Insert_values_list_or_table_clauseContext) {
}

// ExitInsert_values_list_or_table_clause is called when production insert_values_list_or_table_clause is exited.
func (s *BaseGoogleSQLParserListener) ExitInsert_values_list_or_table_clause(ctx *Insert_values_list_or_table_clauseContext) {
}

// EnterTable_clause_unreversed is called when production table_clause_unreversed is entered.
func (s *BaseGoogleSQLParserListener) EnterTable_clause_unreversed(ctx *Table_clause_unreversedContext) {
}

// ExitTable_clause_unreversed is called when production table_clause_unreversed is exited.
func (s *BaseGoogleSQLParserListener) ExitTable_clause_unreversed(ctx *Table_clause_unreversedContext) {
}

// EnterTable_clause_no_keyword is called when production table_clause_no_keyword is entered.
func (s *BaseGoogleSQLParserListener) EnterTable_clause_no_keyword(ctx *Table_clause_no_keywordContext) {
}

// ExitTable_clause_no_keyword is called when production table_clause_no_keyword is exited.
func (s *BaseGoogleSQLParserListener) ExitTable_clause_no_keyword(ctx *Table_clause_no_keywordContext) {
}

// EnterOpt_returning_clause is called when production opt_returning_clause is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_returning_clause(ctx *Opt_returning_clauseContext) {}

// ExitOpt_returning_clause is called when production opt_returning_clause is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_returning_clause(ctx *Opt_returning_clauseContext) {}

// EnterOpt_assert_rows_modified is called when production opt_assert_rows_modified is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_assert_rows_modified(ctx *Opt_assert_rows_modifiedContext) {
}

// ExitOpt_assert_rows_modified is called when production opt_assert_rows_modified is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_assert_rows_modified(ctx *Opt_assert_rows_modifiedContext) {
}

// EnterInsert_values_or_query is called when production insert_values_or_query is entered.
func (s *BaseGoogleSQLParserListener) EnterInsert_values_or_query(ctx *Insert_values_or_queryContext) {
}

// ExitInsert_values_or_query is called when production insert_values_or_query is exited.
func (s *BaseGoogleSQLParserListener) ExitInsert_values_or_query(ctx *Insert_values_or_queryContext) {
}

// EnterInsert_values_list is called when production insert_values_list is entered.
func (s *BaseGoogleSQLParserListener) EnterInsert_values_list(ctx *Insert_values_listContext) {}

// ExitInsert_values_list is called when production insert_values_list is exited.
func (s *BaseGoogleSQLParserListener) ExitInsert_values_list(ctx *Insert_values_listContext) {}

// EnterInsert_values_row is called when production insert_values_row is entered.
func (s *BaseGoogleSQLParserListener) EnterInsert_values_row(ctx *Insert_values_rowContext) {}

// ExitInsert_values_row is called when production insert_values_row is exited.
func (s *BaseGoogleSQLParserListener) ExitInsert_values_row(ctx *Insert_values_rowContext) {}

// EnterExpression_or_default is called when production expression_or_default is entered.
func (s *BaseGoogleSQLParserListener) EnterExpression_or_default(ctx *Expression_or_defaultContext) {}

// ExitExpression_or_default is called when production expression_or_default is exited.
func (s *BaseGoogleSQLParserListener) ExitExpression_or_default(ctx *Expression_or_defaultContext) {}

// EnterInsert_statement_prefix is called when production insert_statement_prefix is entered.
func (s *BaseGoogleSQLParserListener) EnterInsert_statement_prefix(ctx *Insert_statement_prefixContext) {
}

// ExitInsert_statement_prefix is called when production insert_statement_prefix is exited.
func (s *BaseGoogleSQLParserListener) ExitInsert_statement_prefix(ctx *Insert_statement_prefixContext) {
}

// EnterMaybe_dashed_generalized_path_expression is called when production maybe_dashed_generalized_path_expression is entered.
func (s *BaseGoogleSQLParserListener) EnterMaybe_dashed_generalized_path_expression(ctx *Maybe_dashed_generalized_path_expressionContext) {
}

// ExitMaybe_dashed_generalized_path_expression is called when production maybe_dashed_generalized_path_expression is exited.
func (s *BaseGoogleSQLParserListener) ExitMaybe_dashed_generalized_path_expression(ctx *Maybe_dashed_generalized_path_expressionContext) {
}

// EnterOpt_into is called when production opt_into is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_into(ctx *Opt_intoContext) {}

// ExitOpt_into is called when production opt_into is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_into(ctx *Opt_intoContext) {}

// EnterOpt_or_ignore_replace_update is called when production opt_or_ignore_replace_update is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_or_ignore_replace_update(ctx *Opt_or_ignore_replace_updateContext) {
}

// ExitOpt_or_ignore_replace_update is called when production opt_or_ignore_replace_update is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_or_ignore_replace_update(ctx *Opt_or_ignore_replace_updateContext) {
}

// EnterAlter_statement is called when production alter_statement is entered.
func (s *BaseGoogleSQLParserListener) EnterAlter_statement(ctx *Alter_statementContext) {}

// ExitAlter_statement is called when production alter_statement is exited.
func (s *BaseGoogleSQLParserListener) ExitAlter_statement(ctx *Alter_statementContext) {}

// EnterAnalyze_statement is called when production analyze_statement is entered.
func (s *BaseGoogleSQLParserListener) EnterAnalyze_statement(ctx *Analyze_statementContext) {}

// ExitAnalyze_statement is called when production analyze_statement is exited.
func (s *BaseGoogleSQLParserListener) ExitAnalyze_statement(ctx *Analyze_statementContext) {}

// EnterAssert_statement is called when production assert_statement is entered.
func (s *BaseGoogleSQLParserListener) EnterAssert_statement(ctx *Assert_statementContext) {}

// ExitAssert_statement is called when production assert_statement is exited.
func (s *BaseGoogleSQLParserListener) ExitAssert_statement(ctx *Assert_statementContext) {}

// EnterAux_load_data_statement is called when production aux_load_data_statement is entered.
func (s *BaseGoogleSQLParserListener) EnterAux_load_data_statement(ctx *Aux_load_data_statementContext) {
}

// ExitAux_load_data_statement is called when production aux_load_data_statement is exited.
func (s *BaseGoogleSQLParserListener) ExitAux_load_data_statement(ctx *Aux_load_data_statementContext) {
}

// EnterClone_data_statement is called when production clone_data_statement is entered.
func (s *BaseGoogleSQLParserListener) EnterClone_data_statement(ctx *Clone_data_statementContext) {}

// ExitClone_data_statement is called when production clone_data_statement is exited.
func (s *BaseGoogleSQLParserListener) ExitClone_data_statement(ctx *Clone_data_statementContext) {}

// EnterClone_data_source_list is called when production clone_data_source_list is entered.
func (s *BaseGoogleSQLParserListener) EnterClone_data_source_list(ctx *Clone_data_source_listContext) {
}

// ExitClone_data_source_list is called when production clone_data_source_list is exited.
func (s *BaseGoogleSQLParserListener) ExitClone_data_source_list(ctx *Clone_data_source_listContext) {
}

// EnterClone_data_source is called when production clone_data_source is entered.
func (s *BaseGoogleSQLParserListener) EnterClone_data_source(ctx *Clone_data_sourceContext) {}

// ExitClone_data_source is called when production clone_data_source is exited.
func (s *BaseGoogleSQLParserListener) ExitClone_data_source(ctx *Clone_data_sourceContext) {}

// EnterOpt_external_table_with_clauses is called when production opt_external_table_with_clauses is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_external_table_with_clauses(ctx *Opt_external_table_with_clausesContext) {
}

// ExitOpt_external_table_with_clauses is called when production opt_external_table_with_clauses is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_external_table_with_clauses(ctx *Opt_external_table_with_clausesContext) {
}

// EnterWith_partition_columns_clause is called when production with_partition_columns_clause is entered.
func (s *BaseGoogleSQLParserListener) EnterWith_partition_columns_clause(ctx *With_partition_columns_clauseContext) {
}

// ExitWith_partition_columns_clause is called when production with_partition_columns_clause is exited.
func (s *BaseGoogleSQLParserListener) ExitWith_partition_columns_clause(ctx *With_partition_columns_clauseContext) {
}

// EnterAux_load_data_from_files_options_list is called when production aux_load_data_from_files_options_list is entered.
func (s *BaseGoogleSQLParserListener) EnterAux_load_data_from_files_options_list(ctx *Aux_load_data_from_files_options_listContext) {
}

// ExitAux_load_data_from_files_options_list is called when production aux_load_data_from_files_options_list is exited.
func (s *BaseGoogleSQLParserListener) ExitAux_load_data_from_files_options_list(ctx *Aux_load_data_from_files_options_listContext) {
}

// EnterCluster_by_clause_prefix_no_hint is called when production cluster_by_clause_prefix_no_hint is entered.
func (s *BaseGoogleSQLParserListener) EnterCluster_by_clause_prefix_no_hint(ctx *Cluster_by_clause_prefix_no_hintContext) {
}

// ExitCluster_by_clause_prefix_no_hint is called when production cluster_by_clause_prefix_no_hint is exited.
func (s *BaseGoogleSQLParserListener) ExitCluster_by_clause_prefix_no_hint(ctx *Cluster_by_clause_prefix_no_hintContext) {
}

// EnterLoad_data_partitions_clause is called when production load_data_partitions_clause is entered.
func (s *BaseGoogleSQLParserListener) EnterLoad_data_partitions_clause(ctx *Load_data_partitions_clauseContext) {
}

// ExitLoad_data_partitions_clause is called when production load_data_partitions_clause is exited.
func (s *BaseGoogleSQLParserListener) ExitLoad_data_partitions_clause(ctx *Load_data_partitions_clauseContext) {
}

// EnterMaybe_dashed_path_expression_with_scope is called when production maybe_dashed_path_expression_with_scope is entered.
func (s *BaseGoogleSQLParserListener) EnterMaybe_dashed_path_expression_with_scope(ctx *Maybe_dashed_path_expression_with_scopeContext) {
}

// ExitMaybe_dashed_path_expression_with_scope is called when production maybe_dashed_path_expression_with_scope is exited.
func (s *BaseGoogleSQLParserListener) ExitMaybe_dashed_path_expression_with_scope(ctx *Maybe_dashed_path_expression_with_scopeContext) {
}

// EnterTable_element_list is called when production table_element_list is entered.
func (s *BaseGoogleSQLParserListener) EnterTable_element_list(ctx *Table_element_listContext) {}

// ExitTable_element_list is called when production table_element_list is exited.
func (s *BaseGoogleSQLParserListener) ExitTable_element_list(ctx *Table_element_listContext) {}

// EnterTable_element is called when production table_element is entered.
func (s *BaseGoogleSQLParserListener) EnterTable_element(ctx *Table_elementContext) {}

// ExitTable_element is called when production table_element is exited.
func (s *BaseGoogleSQLParserListener) ExitTable_element(ctx *Table_elementContext) {}

// EnterTable_constraint_definition is called when production table_constraint_definition is entered.
func (s *BaseGoogleSQLParserListener) EnterTable_constraint_definition(ctx *Table_constraint_definitionContext) {
}

// ExitTable_constraint_definition is called when production table_constraint_definition is exited.
func (s *BaseGoogleSQLParserListener) ExitTable_constraint_definition(ctx *Table_constraint_definitionContext) {
}

// EnterAppend_or_overwrite is called when production append_or_overwrite is entered.
func (s *BaseGoogleSQLParserListener) EnterAppend_or_overwrite(ctx *Append_or_overwriteContext) {}

// ExitAppend_or_overwrite is called when production append_or_overwrite is exited.
func (s *BaseGoogleSQLParserListener) ExitAppend_or_overwrite(ctx *Append_or_overwriteContext) {}

// EnterOpt_description is called when production opt_description is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_description(ctx *Opt_descriptionContext) {}

// ExitOpt_description is called when production opt_description is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_description(ctx *Opt_descriptionContext) {}

// EnterTable_and_column_info_list is called when production table_and_column_info_list is entered.
func (s *BaseGoogleSQLParserListener) EnterTable_and_column_info_list(ctx *Table_and_column_info_listContext) {
}

// ExitTable_and_column_info_list is called when production table_and_column_info_list is exited.
func (s *BaseGoogleSQLParserListener) ExitTable_and_column_info_list(ctx *Table_and_column_info_listContext) {
}

// EnterTable_and_column_info is called when production table_and_column_info is entered.
func (s *BaseGoogleSQLParserListener) EnterTable_and_column_info(ctx *Table_and_column_infoContext) {}

// ExitTable_and_column_info is called when production table_and_column_info is exited.
func (s *BaseGoogleSQLParserListener) ExitTable_and_column_info(ctx *Table_and_column_infoContext) {}

// EnterRow_access_policy_alter_action_list is called when production row_access_policy_alter_action_list is entered.
func (s *BaseGoogleSQLParserListener) EnterRow_access_policy_alter_action_list(ctx *Row_access_policy_alter_action_listContext) {
}

// ExitRow_access_policy_alter_action_list is called when production row_access_policy_alter_action_list is exited.
func (s *BaseGoogleSQLParserListener) ExitRow_access_policy_alter_action_list(ctx *Row_access_policy_alter_action_listContext) {
}

// EnterRow_access_policy_alter_action is called when production row_access_policy_alter_action is entered.
func (s *BaseGoogleSQLParserListener) EnterRow_access_policy_alter_action(ctx *Row_access_policy_alter_actionContext) {
}

// ExitRow_access_policy_alter_action is called when production row_access_policy_alter_action is exited.
func (s *BaseGoogleSQLParserListener) ExitRow_access_policy_alter_action(ctx *Row_access_policy_alter_actionContext) {
}

// EnterGrant_to_clause is called when production grant_to_clause is entered.
func (s *BaseGoogleSQLParserListener) EnterGrant_to_clause(ctx *Grant_to_clauseContext) {}

// ExitGrant_to_clause is called when production grant_to_clause is exited.
func (s *BaseGoogleSQLParserListener) ExitGrant_to_clause(ctx *Grant_to_clauseContext) {}

// EnterGrantee_list is called when production grantee_list is entered.
func (s *BaseGoogleSQLParserListener) EnterGrantee_list(ctx *Grantee_listContext) {}

// ExitGrantee_list is called when production grantee_list is exited.
func (s *BaseGoogleSQLParserListener) ExitGrantee_list(ctx *Grantee_listContext) {}

// EnterPrivilege_list is called when production privilege_list is entered.
func (s *BaseGoogleSQLParserListener) EnterPrivilege_list(ctx *Privilege_listContext) {}

// ExitPrivilege_list is called when production privilege_list is exited.
func (s *BaseGoogleSQLParserListener) ExitPrivilege_list(ctx *Privilege_listContext) {}

// EnterPrivilege is called when production privilege is entered.
func (s *BaseGoogleSQLParserListener) EnterPrivilege(ctx *PrivilegeContext) {}

// ExitPrivilege is called when production privilege is exited.
func (s *BaseGoogleSQLParserListener) ExitPrivilege(ctx *PrivilegeContext) {}

// EnterPath_expression_list_with_parens is called when production path_expression_list_with_parens is entered.
func (s *BaseGoogleSQLParserListener) EnterPath_expression_list_with_parens(ctx *Path_expression_list_with_parensContext) {
}

// ExitPath_expression_list_with_parens is called when production path_expression_list_with_parens is exited.
func (s *BaseGoogleSQLParserListener) ExitPath_expression_list_with_parens(ctx *Path_expression_list_with_parensContext) {
}

// EnterPrivilege_name is called when production privilege_name is entered.
func (s *BaseGoogleSQLParserListener) EnterPrivilege_name(ctx *Privilege_nameContext) {}

// ExitPrivilege_name is called when production privilege_name is exited.
func (s *BaseGoogleSQLParserListener) ExitPrivilege_name(ctx *Privilege_nameContext) {}

// EnterGeneric_entity_type is called when production generic_entity_type is entered.
func (s *BaseGoogleSQLParserListener) EnterGeneric_entity_type(ctx *Generic_entity_typeContext) {}

// ExitGeneric_entity_type is called when production generic_entity_type is exited.
func (s *BaseGoogleSQLParserListener) ExitGeneric_entity_type(ctx *Generic_entity_typeContext) {}

// EnterGeneric_entity_type_unchecked is called when production generic_entity_type_unchecked is entered.
func (s *BaseGoogleSQLParserListener) EnterGeneric_entity_type_unchecked(ctx *Generic_entity_type_uncheckedContext) {
}

// ExitGeneric_entity_type_unchecked is called when production generic_entity_type_unchecked is exited.
func (s *BaseGoogleSQLParserListener) ExitGeneric_entity_type_unchecked(ctx *Generic_entity_type_uncheckedContext) {
}

// EnterSchema_object_kind is called when production schema_object_kind is entered.
func (s *BaseGoogleSQLParserListener) EnterSchema_object_kind(ctx *Schema_object_kindContext) {}

// ExitSchema_object_kind is called when production schema_object_kind is exited.
func (s *BaseGoogleSQLParserListener) ExitSchema_object_kind(ctx *Schema_object_kindContext) {}

// EnterAlter_action_list is called when production alter_action_list is entered.
func (s *BaseGoogleSQLParserListener) EnterAlter_action_list(ctx *Alter_action_listContext) {}

// ExitAlter_action_list is called when production alter_action_list is exited.
func (s *BaseGoogleSQLParserListener) ExitAlter_action_list(ctx *Alter_action_listContext) {}

// EnterAlter_action is called when production alter_action is entered.
func (s *BaseGoogleSQLParserListener) EnterAlter_action(ctx *Alter_actionContext) {}

// ExitAlter_action is called when production alter_action is exited.
func (s *BaseGoogleSQLParserListener) ExitAlter_action(ctx *Alter_actionContext) {}

// EnterSpanner_set_on_delete_action is called when production spanner_set_on_delete_action is entered.
func (s *BaseGoogleSQLParserListener) EnterSpanner_set_on_delete_action(ctx *Spanner_set_on_delete_actionContext) {
}

// ExitSpanner_set_on_delete_action is called when production spanner_set_on_delete_action is exited.
func (s *BaseGoogleSQLParserListener) ExitSpanner_set_on_delete_action(ctx *Spanner_set_on_delete_actionContext) {
}

// EnterSpanner_alter_column_action is called when production spanner_alter_column_action is entered.
func (s *BaseGoogleSQLParserListener) EnterSpanner_alter_column_action(ctx *Spanner_alter_column_actionContext) {
}

// ExitSpanner_alter_column_action is called when production spanner_alter_column_action is exited.
func (s *BaseGoogleSQLParserListener) ExitSpanner_alter_column_action(ctx *Spanner_alter_column_actionContext) {
}

// EnterSpanner_generated_or_default is called when production spanner_generated_or_default is entered.
func (s *BaseGoogleSQLParserListener) EnterSpanner_generated_or_default(ctx *Spanner_generated_or_defaultContext) {
}

// ExitSpanner_generated_or_default is called when production spanner_generated_or_default is exited.
func (s *BaseGoogleSQLParserListener) ExitSpanner_generated_or_default(ctx *Spanner_generated_or_defaultContext) {
}

// EnterGeneric_sub_entity_type is called when production generic_sub_entity_type is entered.
func (s *BaseGoogleSQLParserListener) EnterGeneric_sub_entity_type(ctx *Generic_sub_entity_typeContext) {
}

// ExitGeneric_sub_entity_type is called when production generic_sub_entity_type is exited.
func (s *BaseGoogleSQLParserListener) ExitGeneric_sub_entity_type(ctx *Generic_sub_entity_typeContext) {
}

// EnterSub_entity_type_identifier is called when production sub_entity_type_identifier is entered.
func (s *BaseGoogleSQLParserListener) EnterSub_entity_type_identifier(ctx *Sub_entity_type_identifierContext) {
}

// ExitSub_entity_type_identifier is called when production sub_entity_type_identifier is exited.
func (s *BaseGoogleSQLParserListener) ExitSub_entity_type_identifier(ctx *Sub_entity_type_identifierContext) {
}

// EnterFill_using_expression is called when production fill_using_expression is entered.
func (s *BaseGoogleSQLParserListener) EnterFill_using_expression(ctx *Fill_using_expressionContext) {}

// ExitFill_using_expression is called when production fill_using_expression is exited.
func (s *BaseGoogleSQLParserListener) ExitFill_using_expression(ctx *Fill_using_expressionContext) {}

// EnterColumn_position is called when production column_position is entered.
func (s *BaseGoogleSQLParserListener) EnterColumn_position(ctx *Column_positionContext) {}

// ExitColumn_position is called when production column_position is exited.
func (s *BaseGoogleSQLParserListener) ExitColumn_position(ctx *Column_positionContext) {}

// EnterTable_column_definition is called when production table_column_definition is entered.
func (s *BaseGoogleSQLParserListener) EnterTable_column_definition(ctx *Table_column_definitionContext) {
}

// ExitTable_column_definition is called when production table_column_definition is exited.
func (s *BaseGoogleSQLParserListener) ExitTable_column_definition(ctx *Table_column_definitionContext) {
}

// EnterColumn_attributes is called when production column_attributes is entered.
func (s *BaseGoogleSQLParserListener) EnterColumn_attributes(ctx *Column_attributesContext) {}

// ExitColumn_attributes is called when production column_attributes is exited.
func (s *BaseGoogleSQLParserListener) ExitColumn_attributes(ctx *Column_attributesContext) {}

// EnterColumn_attribute is called when production column_attribute is entered.
func (s *BaseGoogleSQLParserListener) EnterColumn_attribute(ctx *Column_attributeContext) {}

// ExitColumn_attribute is called when production column_attribute is exited.
func (s *BaseGoogleSQLParserListener) ExitColumn_attribute(ctx *Column_attributeContext) {}

// EnterPrimary_key_column_attribute is called when production primary_key_column_attribute is entered.
func (s *BaseGoogleSQLParserListener) EnterPrimary_key_column_attribute(ctx *Primary_key_column_attributeContext) {
}

// ExitPrimary_key_column_attribute is called when production primary_key_column_attribute is exited.
func (s *BaseGoogleSQLParserListener) ExitPrimary_key_column_attribute(ctx *Primary_key_column_attributeContext) {
}

// EnterForeign_key_column_attribute is called when production foreign_key_column_attribute is entered.
func (s *BaseGoogleSQLParserListener) EnterForeign_key_column_attribute(ctx *Foreign_key_column_attributeContext) {
}

// ExitForeign_key_column_attribute is called when production foreign_key_column_attribute is exited.
func (s *BaseGoogleSQLParserListener) ExitForeign_key_column_attribute(ctx *Foreign_key_column_attributeContext) {
}

// EnterHidden_column_attribute is called when production hidden_column_attribute is entered.
func (s *BaseGoogleSQLParserListener) EnterHidden_column_attribute(ctx *Hidden_column_attributeContext) {
}

// ExitHidden_column_attribute is called when production hidden_column_attribute is exited.
func (s *BaseGoogleSQLParserListener) ExitHidden_column_attribute(ctx *Hidden_column_attributeContext) {
}

// EnterOpt_constraint_identity is called when production opt_constraint_identity is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_constraint_identity(ctx *Opt_constraint_identityContext) {
}

// ExitOpt_constraint_identity is called when production opt_constraint_identity is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_constraint_identity(ctx *Opt_constraint_identityContext) {
}

// EnterTable_column_schema is called when production table_column_schema is entered.
func (s *BaseGoogleSQLParserListener) EnterTable_column_schema(ctx *Table_column_schemaContext) {}

// ExitTable_column_schema is called when production table_column_schema is exited.
func (s *BaseGoogleSQLParserListener) ExitTable_column_schema(ctx *Table_column_schemaContext) {}

// EnterOpt_column_info is called when production opt_column_info is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_column_info(ctx *Opt_column_infoContext) {}

// ExitOpt_column_info is called when production opt_column_info is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_column_info(ctx *Opt_column_infoContext) {}

// EnterInvalid_generated_column is called when production invalid_generated_column is entered.
func (s *BaseGoogleSQLParserListener) EnterInvalid_generated_column(ctx *Invalid_generated_columnContext) {
}

// ExitInvalid_generated_column is called when production invalid_generated_column is exited.
func (s *BaseGoogleSQLParserListener) ExitInvalid_generated_column(ctx *Invalid_generated_columnContext) {
}

// EnterInvalid_default_column is called when production invalid_default_column is entered.
func (s *BaseGoogleSQLParserListener) EnterInvalid_default_column(ctx *Invalid_default_columnContext) {
}

// ExitInvalid_default_column is called when production invalid_default_column is exited.
func (s *BaseGoogleSQLParserListener) ExitInvalid_default_column(ctx *Invalid_default_columnContext) {
}

// EnterDefault_column_info is called when production default_column_info is entered.
func (s *BaseGoogleSQLParserListener) EnterDefault_column_info(ctx *Default_column_infoContext) {}

// ExitDefault_column_info is called when production default_column_info is exited.
func (s *BaseGoogleSQLParserListener) ExitDefault_column_info(ctx *Default_column_infoContext) {}

// EnterGenerated_column_info is called when production generated_column_info is entered.
func (s *BaseGoogleSQLParserListener) EnterGenerated_column_info(ctx *Generated_column_infoContext) {}

// ExitGenerated_column_info is called when production generated_column_info is exited.
func (s *BaseGoogleSQLParserListener) ExitGenerated_column_info(ctx *Generated_column_infoContext) {}

// EnterIdentity_column_info is called when production identity_column_info is entered.
func (s *BaseGoogleSQLParserListener) EnterIdentity_column_info(ctx *Identity_column_infoContext) {}

// ExitIdentity_column_info is called when production identity_column_info is exited.
func (s *BaseGoogleSQLParserListener) ExitIdentity_column_info(ctx *Identity_column_infoContext) {}

// EnterOpt_start_with is called when production opt_start_with is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_start_with(ctx *Opt_start_withContext) {}

// ExitOpt_start_with is called when production opt_start_with is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_start_with(ctx *Opt_start_withContext) {}

// EnterOpt_increment_by is called when production opt_increment_by is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_increment_by(ctx *Opt_increment_byContext) {}

// ExitOpt_increment_by is called when production opt_increment_by is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_increment_by(ctx *Opt_increment_byContext) {}

// EnterOpt_maxvalue is called when production opt_maxvalue is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_maxvalue(ctx *Opt_maxvalueContext) {}

// ExitOpt_maxvalue is called when production opt_maxvalue is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_maxvalue(ctx *Opt_maxvalueContext) {}

// EnterOpt_minvalue is called when production opt_minvalue is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_minvalue(ctx *Opt_minvalueContext) {}

// ExitOpt_minvalue is called when production opt_minvalue is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_minvalue(ctx *Opt_minvalueContext) {}

// EnterOpt_cycle is called when production opt_cycle is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_cycle(ctx *Opt_cycleContext) {}

// ExitOpt_cycle is called when production opt_cycle is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_cycle(ctx *Opt_cycleContext) {}

// EnterSigned_numeric_literal is called when production signed_numeric_literal is entered.
func (s *BaseGoogleSQLParserListener) EnterSigned_numeric_literal(ctx *Signed_numeric_literalContext) {
}

// ExitSigned_numeric_literal is called when production signed_numeric_literal is exited.
func (s *BaseGoogleSQLParserListener) ExitSigned_numeric_literal(ctx *Signed_numeric_literalContext) {
}

// EnterStored_mode is called when production stored_mode is entered.
func (s *BaseGoogleSQLParserListener) EnterStored_mode(ctx *Stored_modeContext) {}

// ExitStored_mode is called when production stored_mode is exited.
func (s *BaseGoogleSQLParserListener) ExitStored_mode(ctx *Stored_modeContext) {}

// EnterGenerated_mode is called when production generated_mode is entered.
func (s *BaseGoogleSQLParserListener) EnterGenerated_mode(ctx *Generated_modeContext) {}

// ExitGenerated_mode is called when production generated_mode is exited.
func (s *BaseGoogleSQLParserListener) ExitGenerated_mode(ctx *Generated_modeContext) {}

// EnterColumn_schema_inner is called when production column_schema_inner is entered.
func (s *BaseGoogleSQLParserListener) EnterColumn_schema_inner(ctx *Column_schema_innerContext) {}

// ExitColumn_schema_inner is called when production column_schema_inner is exited.
func (s *BaseGoogleSQLParserListener) ExitColumn_schema_inner(ctx *Column_schema_innerContext) {}

// EnterRaw_column_schema_inner is called when production raw_column_schema_inner is entered.
func (s *BaseGoogleSQLParserListener) EnterRaw_column_schema_inner(ctx *Raw_column_schema_innerContext) {
}

// ExitRaw_column_schema_inner is called when production raw_column_schema_inner is exited.
func (s *BaseGoogleSQLParserListener) ExitRaw_column_schema_inner(ctx *Raw_column_schema_innerContext) {
}

// EnterRange_column_schema_inner is called when production range_column_schema_inner is entered.
func (s *BaseGoogleSQLParserListener) EnterRange_column_schema_inner(ctx *Range_column_schema_innerContext) {
}

// ExitRange_column_schema_inner is called when production range_column_schema_inner is exited.
func (s *BaseGoogleSQLParserListener) ExitRange_column_schema_inner(ctx *Range_column_schema_innerContext) {
}

// EnterStruct_column_schema_inner is called when production struct_column_schema_inner is entered.
func (s *BaseGoogleSQLParserListener) EnterStruct_column_schema_inner(ctx *Struct_column_schema_innerContext) {
}

// ExitStruct_column_schema_inner is called when production struct_column_schema_inner is exited.
func (s *BaseGoogleSQLParserListener) ExitStruct_column_schema_inner(ctx *Struct_column_schema_innerContext) {
}

// EnterStruct_column_field is called when production struct_column_field is entered.
func (s *BaseGoogleSQLParserListener) EnterStruct_column_field(ctx *Struct_column_fieldContext) {}

// ExitStruct_column_field is called when production struct_column_field is exited.
func (s *BaseGoogleSQLParserListener) ExitStruct_column_field(ctx *Struct_column_fieldContext) {}

// EnterSimple_column_schema_inner is called when production simple_column_schema_inner is entered.
func (s *BaseGoogleSQLParserListener) EnterSimple_column_schema_inner(ctx *Simple_column_schema_innerContext) {
}

// ExitSimple_column_schema_inner is called when production simple_column_schema_inner is exited.
func (s *BaseGoogleSQLParserListener) ExitSimple_column_schema_inner(ctx *Simple_column_schema_innerContext) {
}

// EnterArray_column_schema_inner is called when production array_column_schema_inner is entered.
func (s *BaseGoogleSQLParserListener) EnterArray_column_schema_inner(ctx *Array_column_schema_innerContext) {
}

// ExitArray_column_schema_inner is called when production array_column_schema_inner is exited.
func (s *BaseGoogleSQLParserListener) ExitArray_column_schema_inner(ctx *Array_column_schema_innerContext) {
}

// EnterField_schema is called when production field_schema is entered.
func (s *BaseGoogleSQLParserListener) EnterField_schema(ctx *Field_schemaContext) {}

// ExitField_schema is called when production field_schema is exited.
func (s *BaseGoogleSQLParserListener) ExitField_schema(ctx *Field_schemaContext) {}

// EnterOpt_field_attributes is called when production opt_field_attributes is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_field_attributes(ctx *Opt_field_attributesContext) {}

// ExitOpt_field_attributes is called when production opt_field_attributes is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_field_attributes(ctx *Opt_field_attributesContext) {}

// EnterNot_null_column_attribute is called when production not_null_column_attribute is entered.
func (s *BaseGoogleSQLParserListener) EnterNot_null_column_attribute(ctx *Not_null_column_attributeContext) {
}

// ExitNot_null_column_attribute is called when production not_null_column_attribute is exited.
func (s *BaseGoogleSQLParserListener) ExitNot_null_column_attribute(ctx *Not_null_column_attributeContext) {
}

// EnterPrimary_key_or_table_constraint_spec is called when production primary_key_or_table_constraint_spec is entered.
func (s *BaseGoogleSQLParserListener) EnterPrimary_key_or_table_constraint_spec(ctx *Primary_key_or_table_constraint_specContext) {
}

// ExitPrimary_key_or_table_constraint_spec is called when production primary_key_or_table_constraint_spec is exited.
func (s *BaseGoogleSQLParserListener) ExitPrimary_key_or_table_constraint_spec(ctx *Primary_key_or_table_constraint_specContext) {
}

// EnterOpt_if_not_exists is called when production opt_if_not_exists is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_if_not_exists(ctx *Opt_if_not_existsContext) {}

// ExitOpt_if_not_exists is called when production opt_if_not_exists is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_if_not_exists(ctx *Opt_if_not_existsContext) {}

// EnterPrimary_key_spec is called when production primary_key_spec is entered.
func (s *BaseGoogleSQLParserListener) EnterPrimary_key_spec(ctx *Primary_key_specContext) {}

// ExitPrimary_key_spec is called when production primary_key_spec is exited.
func (s *BaseGoogleSQLParserListener) ExitPrimary_key_spec(ctx *Primary_key_specContext) {}

// EnterPrimary_key_element_list is called when production primary_key_element_list is entered.
func (s *BaseGoogleSQLParserListener) EnterPrimary_key_element_list(ctx *Primary_key_element_listContext) {
}

// ExitPrimary_key_element_list is called when production primary_key_element_list is exited.
func (s *BaseGoogleSQLParserListener) ExitPrimary_key_element_list(ctx *Primary_key_element_listContext) {
}

// EnterPrimary_key_element is called when production primary_key_element is entered.
func (s *BaseGoogleSQLParserListener) EnterPrimary_key_element(ctx *Primary_key_elementContext) {}

// ExitPrimary_key_element is called when production primary_key_element is exited.
func (s *BaseGoogleSQLParserListener) ExitPrimary_key_element(ctx *Primary_key_elementContext) {}

// EnterTable_constraint_spec is called when production table_constraint_spec is entered.
func (s *BaseGoogleSQLParserListener) EnterTable_constraint_spec(ctx *Table_constraint_specContext) {}

// ExitTable_constraint_spec is called when production table_constraint_spec is exited.
func (s *BaseGoogleSQLParserListener) ExitTable_constraint_spec(ctx *Table_constraint_specContext) {}

// EnterForeign_key_reference is called when production foreign_key_reference is entered.
func (s *BaseGoogleSQLParserListener) EnterForeign_key_reference(ctx *Foreign_key_referenceContext) {}

// ExitForeign_key_reference is called when production foreign_key_reference is exited.
func (s *BaseGoogleSQLParserListener) ExitForeign_key_reference(ctx *Foreign_key_referenceContext) {}

// EnterOpt_foreign_key_action is called when production opt_foreign_key_action is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_foreign_key_action(ctx *Opt_foreign_key_actionContext) {
}

// ExitOpt_foreign_key_action is called when production opt_foreign_key_action is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_foreign_key_action(ctx *Opt_foreign_key_actionContext) {
}

// EnterForeign_key_on_update is called when production foreign_key_on_update is entered.
func (s *BaseGoogleSQLParserListener) EnterForeign_key_on_update(ctx *Foreign_key_on_updateContext) {}

// ExitForeign_key_on_update is called when production foreign_key_on_update is exited.
func (s *BaseGoogleSQLParserListener) ExitForeign_key_on_update(ctx *Foreign_key_on_updateContext) {}

// EnterForeign_key_on_delete is called when production foreign_key_on_delete is entered.
func (s *BaseGoogleSQLParserListener) EnterForeign_key_on_delete(ctx *Foreign_key_on_deleteContext) {}

// ExitForeign_key_on_delete is called when production foreign_key_on_delete is exited.
func (s *BaseGoogleSQLParserListener) ExitForeign_key_on_delete(ctx *Foreign_key_on_deleteContext) {}

// EnterForeign_key_action is called when production foreign_key_action is entered.
func (s *BaseGoogleSQLParserListener) EnterForeign_key_action(ctx *Foreign_key_actionContext) {}

// ExitForeign_key_action is called when production foreign_key_action is exited.
func (s *BaseGoogleSQLParserListener) ExitForeign_key_action(ctx *Foreign_key_actionContext) {}

// EnterOpt_foreign_key_match is called when production opt_foreign_key_match is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_foreign_key_match(ctx *Opt_foreign_key_matchContext) {}

// ExitOpt_foreign_key_match is called when production opt_foreign_key_match is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_foreign_key_match(ctx *Opt_foreign_key_matchContext) {}

// EnterForeign_key_match_mode is called when production foreign_key_match_mode is entered.
func (s *BaseGoogleSQLParserListener) EnterForeign_key_match_mode(ctx *Foreign_key_match_modeContext) {
}

// ExitForeign_key_match_mode is called when production foreign_key_match_mode is exited.
func (s *BaseGoogleSQLParserListener) ExitForeign_key_match_mode(ctx *Foreign_key_match_modeContext) {
}

// EnterColumn_list is called when production column_list is entered.
func (s *BaseGoogleSQLParserListener) EnterColumn_list(ctx *Column_listContext) {}

// ExitColumn_list is called when production column_list is exited.
func (s *BaseGoogleSQLParserListener) ExitColumn_list(ctx *Column_listContext) {}

// EnterOpt_options_list is called when production opt_options_list is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_options_list(ctx *Opt_options_listContext) {}

// ExitOpt_options_list is called when production opt_options_list is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_options_list(ctx *Opt_options_listContext) {}

// EnterConstraint_enforcement is called when production constraint_enforcement is entered.
func (s *BaseGoogleSQLParserListener) EnterConstraint_enforcement(ctx *Constraint_enforcementContext) {
}

// ExitConstraint_enforcement is called when production constraint_enforcement is exited.
func (s *BaseGoogleSQLParserListener) ExitConstraint_enforcement(ctx *Constraint_enforcementContext) {
}

// EnterGeneric_entity_body is called when production generic_entity_body is entered.
func (s *BaseGoogleSQLParserListener) EnterGeneric_entity_body(ctx *Generic_entity_bodyContext) {}

// ExitGeneric_entity_body is called when production generic_entity_body is exited.
func (s *BaseGoogleSQLParserListener) ExitGeneric_entity_body(ctx *Generic_entity_bodyContext) {}

// EnterOpt_if_exists is called when production opt_if_exists is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_if_exists(ctx *Opt_if_existsContext) {}

// ExitOpt_if_exists is called when production opt_if_exists is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_if_exists(ctx *Opt_if_existsContext) {}

// EnterTable_or_table_function is called when production table_or_table_function is entered.
func (s *BaseGoogleSQLParserListener) EnterTable_or_table_function(ctx *Table_or_table_functionContext) {
}

// ExitTable_or_table_function is called when production table_or_table_function is exited.
func (s *BaseGoogleSQLParserListener) ExitTable_or_table_function(ctx *Table_or_table_functionContext) {
}

// EnterQuery is called when production query is entered.
func (s *BaseGoogleSQLParserListener) EnterQuery(ctx *QueryContext) {}

// ExitQuery is called when production query is exited.
func (s *BaseGoogleSQLParserListener) ExitQuery(ctx *QueryContext) {}

// EnterQuery_without_pipe_operators is called when production query_without_pipe_operators is entered.
func (s *BaseGoogleSQLParserListener) EnterQuery_without_pipe_operators(ctx *Query_without_pipe_operatorsContext) {
}

// ExitQuery_without_pipe_operators is called when production query_without_pipe_operators is exited.
func (s *BaseGoogleSQLParserListener) ExitQuery_without_pipe_operators(ctx *Query_without_pipe_operatorsContext) {
}

// EnterBad_keyword_after_from_query is called when production bad_keyword_after_from_query is entered.
func (s *BaseGoogleSQLParserListener) EnterBad_keyword_after_from_query(ctx *Bad_keyword_after_from_queryContext) {
}

// ExitBad_keyword_after_from_query is called when production bad_keyword_after_from_query is exited.
func (s *BaseGoogleSQLParserListener) ExitBad_keyword_after_from_query(ctx *Bad_keyword_after_from_queryContext) {
}

// EnterBad_keyword_after_from_query_allows_parens is called when production bad_keyword_after_from_query_allows_parens is entered.
func (s *BaseGoogleSQLParserListener) EnterBad_keyword_after_from_query_allows_parens(ctx *Bad_keyword_after_from_query_allows_parensContext) {
}

// ExitBad_keyword_after_from_query_allows_parens is called when production bad_keyword_after_from_query_allows_parens is exited.
func (s *BaseGoogleSQLParserListener) ExitBad_keyword_after_from_query_allows_parens(ctx *Bad_keyword_after_from_query_allows_parensContext) {
}

// EnterWith_clause_with_trailing_comma is called when production with_clause_with_trailing_comma is entered.
func (s *BaseGoogleSQLParserListener) EnterWith_clause_with_trailing_comma(ctx *With_clause_with_trailing_commaContext) {
}

// ExitWith_clause_with_trailing_comma is called when production with_clause_with_trailing_comma is exited.
func (s *BaseGoogleSQLParserListener) ExitWith_clause_with_trailing_comma(ctx *With_clause_with_trailing_commaContext) {
}

// EnterSelect_or_from_keyword is called when production select_or_from_keyword is entered.
func (s *BaseGoogleSQLParserListener) EnterSelect_or_from_keyword(ctx *Select_or_from_keywordContext) {
}

// ExitSelect_or_from_keyword is called when production select_or_from_keyword is exited.
func (s *BaseGoogleSQLParserListener) ExitSelect_or_from_keyword(ctx *Select_or_from_keywordContext) {
}

// EnterQuery_primary_or_set_operation is called when production query_primary_or_set_operation is entered.
func (s *BaseGoogleSQLParserListener) EnterQuery_primary_or_set_operation(ctx *Query_primary_or_set_operationContext) {
}

// ExitQuery_primary_or_set_operation is called when production query_primary_or_set_operation is exited.
func (s *BaseGoogleSQLParserListener) ExitQuery_primary_or_set_operation(ctx *Query_primary_or_set_operationContext) {
}

// EnterQuery_set_operation is called when production query_set_operation is entered.
func (s *BaseGoogleSQLParserListener) EnterQuery_set_operation(ctx *Query_set_operationContext) {}

// ExitQuery_set_operation is called when production query_set_operation is exited.
func (s *BaseGoogleSQLParserListener) ExitQuery_set_operation(ctx *Query_set_operationContext) {}

// EnterQuery_set_operation_prefix is called when production query_set_operation_prefix is entered.
func (s *BaseGoogleSQLParserListener) EnterQuery_set_operation_prefix(ctx *Query_set_operation_prefixContext) {
}

// ExitQuery_set_operation_prefix is called when production query_set_operation_prefix is exited.
func (s *BaseGoogleSQLParserListener) ExitQuery_set_operation_prefix(ctx *Query_set_operation_prefixContext) {
}

// EnterQuery_set_operation_item is called when production query_set_operation_item is entered.
func (s *BaseGoogleSQLParserListener) EnterQuery_set_operation_item(ctx *Query_set_operation_itemContext) {
}

// ExitQuery_set_operation_item is called when production query_set_operation_item is exited.
func (s *BaseGoogleSQLParserListener) ExitQuery_set_operation_item(ctx *Query_set_operation_itemContext) {
}

// EnterQuery_primary is called when production query_primary is entered.
func (s *BaseGoogleSQLParserListener) EnterQuery_primary(ctx *Query_primaryContext) {}

// ExitQuery_primary is called when production query_primary is exited.
func (s *BaseGoogleSQLParserListener) ExitQuery_primary(ctx *Query_primaryContext) {}

// EnterSet_operation_metadata is called when production set_operation_metadata is entered.
func (s *BaseGoogleSQLParserListener) EnterSet_operation_metadata(ctx *Set_operation_metadataContext) {
}

// ExitSet_operation_metadata is called when production set_operation_metadata is exited.
func (s *BaseGoogleSQLParserListener) ExitSet_operation_metadata(ctx *Set_operation_metadataContext) {
}

// EnterOpt_column_match_suffix is called when production opt_column_match_suffix is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_column_match_suffix(ctx *Opt_column_match_suffixContext) {
}

// ExitOpt_column_match_suffix is called when production opt_column_match_suffix is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_column_match_suffix(ctx *Opt_column_match_suffixContext) {
}

// EnterOpt_strict is called when production opt_strict is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_strict(ctx *Opt_strictContext) {}

// ExitOpt_strict is called when production opt_strict is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_strict(ctx *Opt_strictContext) {}

// EnterAll_or_distinct is called when production all_or_distinct is entered.
func (s *BaseGoogleSQLParserListener) EnterAll_or_distinct(ctx *All_or_distinctContext) {}

// ExitAll_or_distinct is called when production all_or_distinct is exited.
func (s *BaseGoogleSQLParserListener) ExitAll_or_distinct(ctx *All_or_distinctContext) {}

// EnterQuery_set_operation_type is called when production query_set_operation_type is entered.
func (s *BaseGoogleSQLParserListener) EnterQuery_set_operation_type(ctx *Query_set_operation_typeContext) {
}

// ExitQuery_set_operation_type is called when production query_set_operation_type is exited.
func (s *BaseGoogleSQLParserListener) ExitQuery_set_operation_type(ctx *Query_set_operation_typeContext) {
}

// EnterOpt_corresponding_outer_mode is called when production opt_corresponding_outer_mode is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_corresponding_outer_mode(ctx *Opt_corresponding_outer_modeContext) {
}

// ExitOpt_corresponding_outer_mode is called when production opt_corresponding_outer_mode is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_corresponding_outer_mode(ctx *Opt_corresponding_outer_modeContext) {
}

// EnterOpt_outer is called when production opt_outer is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_outer(ctx *Opt_outerContext) {}

// ExitOpt_outer is called when production opt_outer is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_outer(ctx *Opt_outerContext) {}

// EnterWith_clause is called when production with_clause is entered.
func (s *BaseGoogleSQLParserListener) EnterWith_clause(ctx *With_clauseContext) {}

// ExitWith_clause is called when production with_clause is exited.
func (s *BaseGoogleSQLParserListener) ExitWith_clause(ctx *With_clauseContext) {}

// EnterAliased_query is called when production aliased_query is entered.
func (s *BaseGoogleSQLParserListener) EnterAliased_query(ctx *Aliased_queryContext) {}

// ExitAliased_query is called when production aliased_query is exited.
func (s *BaseGoogleSQLParserListener) ExitAliased_query(ctx *Aliased_queryContext) {}

// EnterOpt_aliased_query_modifiers is called when production opt_aliased_query_modifiers is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_aliased_query_modifiers(ctx *Opt_aliased_query_modifiersContext) {
}

// ExitOpt_aliased_query_modifiers is called when production opt_aliased_query_modifiers is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_aliased_query_modifiers(ctx *Opt_aliased_query_modifiersContext) {
}

// EnterRecursion_depth_modifier is called when production recursion_depth_modifier is entered.
func (s *BaseGoogleSQLParserListener) EnterRecursion_depth_modifier(ctx *Recursion_depth_modifierContext) {
}

// ExitRecursion_depth_modifier is called when production recursion_depth_modifier is exited.
func (s *BaseGoogleSQLParserListener) ExitRecursion_depth_modifier(ctx *Recursion_depth_modifierContext) {
}

// EnterPossibly_unbounded_int_literal_or_parameter is called when production possibly_unbounded_int_literal_or_parameter is entered.
func (s *BaseGoogleSQLParserListener) EnterPossibly_unbounded_int_literal_or_parameter(ctx *Possibly_unbounded_int_literal_or_parameterContext) {
}

// ExitPossibly_unbounded_int_literal_or_parameter is called when production possibly_unbounded_int_literal_or_parameter is exited.
func (s *BaseGoogleSQLParserListener) ExitPossibly_unbounded_int_literal_or_parameter(ctx *Possibly_unbounded_int_literal_or_parameterContext) {
}

// EnterInt_literal_or_parameter is called when production int_literal_or_parameter is entered.
func (s *BaseGoogleSQLParserListener) EnterInt_literal_or_parameter(ctx *Int_literal_or_parameterContext) {
}

// ExitInt_literal_or_parameter is called when production int_literal_or_parameter is exited.
func (s *BaseGoogleSQLParserListener) ExitInt_literal_or_parameter(ctx *Int_literal_or_parameterContext) {
}

// EnterOrder_by_clause is called when production order_by_clause is entered.
func (s *BaseGoogleSQLParserListener) EnterOrder_by_clause(ctx *Order_by_clauseContext) {}

// ExitOrder_by_clause is called when production order_by_clause is exited.
func (s *BaseGoogleSQLParserListener) ExitOrder_by_clause(ctx *Order_by_clauseContext) {}

// EnterOrder_by_clause_prefix is called when production order_by_clause_prefix is entered.
func (s *BaseGoogleSQLParserListener) EnterOrder_by_clause_prefix(ctx *Order_by_clause_prefixContext) {
}

// ExitOrder_by_clause_prefix is called when production order_by_clause_prefix is exited.
func (s *BaseGoogleSQLParserListener) ExitOrder_by_clause_prefix(ctx *Order_by_clause_prefixContext) {
}

// EnterOrdering_expression is called when production ordering_expression is entered.
func (s *BaseGoogleSQLParserListener) EnterOrdering_expression(ctx *Ordering_expressionContext) {}

// ExitOrdering_expression is called when production ordering_expression is exited.
func (s *BaseGoogleSQLParserListener) ExitOrdering_expression(ctx *Ordering_expressionContext) {}

// EnterSelect is called when production select is entered.
func (s *BaseGoogleSQLParserListener) EnterSelect(ctx *SelectContext) {}

// ExitSelect is called when production select is exited.
func (s *BaseGoogleSQLParserListener) ExitSelect(ctx *SelectContext) {}

// EnterOpt_clauses_following_from is called when production opt_clauses_following_from is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_clauses_following_from(ctx *Opt_clauses_following_fromContext) {
}

// ExitOpt_clauses_following_from is called when production opt_clauses_following_from is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_clauses_following_from(ctx *Opt_clauses_following_fromContext) {
}

// EnterOpt_clauses_following_where is called when production opt_clauses_following_where is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_clauses_following_where(ctx *Opt_clauses_following_whereContext) {
}

// ExitOpt_clauses_following_where is called when production opt_clauses_following_where is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_clauses_following_where(ctx *Opt_clauses_following_whereContext) {
}

// EnterOpt_clauses_following_group_by is called when production opt_clauses_following_group_by is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_clauses_following_group_by(ctx *Opt_clauses_following_group_byContext) {
}

// ExitOpt_clauses_following_group_by is called when production opt_clauses_following_group_by is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_clauses_following_group_by(ctx *Opt_clauses_following_group_byContext) {
}

// EnterWindow_clause is called when production window_clause is entered.
func (s *BaseGoogleSQLParserListener) EnterWindow_clause(ctx *Window_clauseContext) {}

// ExitWindow_clause is called when production window_clause is exited.
func (s *BaseGoogleSQLParserListener) ExitWindow_clause(ctx *Window_clauseContext) {}

// EnterWindow_clause_prefix is called when production window_clause_prefix is entered.
func (s *BaseGoogleSQLParserListener) EnterWindow_clause_prefix(ctx *Window_clause_prefixContext) {}

// ExitWindow_clause_prefix is called when production window_clause_prefix is exited.
func (s *BaseGoogleSQLParserListener) ExitWindow_clause_prefix(ctx *Window_clause_prefixContext) {}

// EnterWindow_definition is called when production window_definition is entered.
func (s *BaseGoogleSQLParserListener) EnterWindow_definition(ctx *Window_definitionContext) {}

// ExitWindow_definition is called when production window_definition is exited.
func (s *BaseGoogleSQLParserListener) ExitWindow_definition(ctx *Window_definitionContext) {}

// EnterWhere_clause is called when production where_clause is entered.
func (s *BaseGoogleSQLParserListener) EnterWhere_clause(ctx *Where_clauseContext) {}

// ExitWhere_clause is called when production where_clause is exited.
func (s *BaseGoogleSQLParserListener) ExitWhere_clause(ctx *Where_clauseContext) {}

// EnterHaving_clause is called when production having_clause is entered.
func (s *BaseGoogleSQLParserListener) EnterHaving_clause(ctx *Having_clauseContext) {}

// ExitHaving_clause is called when production having_clause is exited.
func (s *BaseGoogleSQLParserListener) ExitHaving_clause(ctx *Having_clauseContext) {}

// EnterGroup_by_clause is called when production group_by_clause is entered.
func (s *BaseGoogleSQLParserListener) EnterGroup_by_clause(ctx *Group_by_clauseContext) {}

// ExitGroup_by_clause is called when production group_by_clause is exited.
func (s *BaseGoogleSQLParserListener) ExitGroup_by_clause(ctx *Group_by_clauseContext) {}

// EnterGroup_by_all is called when production group_by_all is entered.
func (s *BaseGoogleSQLParserListener) EnterGroup_by_all(ctx *Group_by_allContext) {}

// ExitGroup_by_all is called when production group_by_all is exited.
func (s *BaseGoogleSQLParserListener) ExitGroup_by_all(ctx *Group_by_allContext) {}

// EnterSelect_clause is called when production select_clause is entered.
func (s *BaseGoogleSQLParserListener) EnterSelect_clause(ctx *Select_clauseContext) {}

// ExitSelect_clause is called when production select_clause is exited.
func (s *BaseGoogleSQLParserListener) ExitSelect_clause(ctx *Select_clauseContext) {}

// EnterOpt_select_as_clause is called when production opt_select_as_clause is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_select_as_clause(ctx *Opt_select_as_clauseContext) {}

// ExitOpt_select_as_clause is called when production opt_select_as_clause is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_select_as_clause(ctx *Opt_select_as_clauseContext) {}

// EnterOpt_select_with is called when production opt_select_with is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_select_with(ctx *Opt_select_withContext) {}

// ExitOpt_select_with is called when production opt_select_with is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_select_with(ctx *Opt_select_withContext) {}

// EnterFrom_clause is called when production from_clause is entered.
func (s *BaseGoogleSQLParserListener) EnterFrom_clause(ctx *From_clauseContext) {}

// ExitFrom_clause is called when production from_clause is exited.
func (s *BaseGoogleSQLParserListener) ExitFrom_clause(ctx *From_clauseContext) {}

// EnterFrom_clause_contents is called when production from_clause_contents is entered.
func (s *BaseGoogleSQLParserListener) EnterFrom_clause_contents(ctx *From_clause_contentsContext) {}

// ExitFrom_clause_contents is called when production from_clause_contents is exited.
func (s *BaseGoogleSQLParserListener) ExitFrom_clause_contents(ctx *From_clause_contentsContext) {}

// EnterFrom_clause_contents_suffix is called when production from_clause_contents_suffix is entered.
func (s *BaseGoogleSQLParserListener) EnterFrom_clause_contents_suffix(ctx *From_clause_contents_suffixContext) {
}

// ExitFrom_clause_contents_suffix is called when production from_clause_contents_suffix is exited.
func (s *BaseGoogleSQLParserListener) ExitFrom_clause_contents_suffix(ctx *From_clause_contents_suffixContext) {
}

// EnterTable_primary is called when production table_primary is entered.
func (s *BaseGoogleSQLParserListener) EnterTable_primary(ctx *Table_primaryContext) {}

// ExitTable_primary is called when production table_primary is exited.
func (s *BaseGoogleSQLParserListener) ExitTable_primary(ctx *Table_primaryContext) {}

// EnterTvf_with_suffixes is called when production tvf_with_suffixes is entered.
func (s *BaseGoogleSQLParserListener) EnterTvf_with_suffixes(ctx *Tvf_with_suffixesContext) {}

// ExitTvf_with_suffixes is called when production tvf_with_suffixes is exited.
func (s *BaseGoogleSQLParserListener) ExitTvf_with_suffixes(ctx *Tvf_with_suffixesContext) {}

// EnterPivot_or_unpivot_clause_and_aliases is called when production pivot_or_unpivot_clause_and_aliases is entered.
func (s *BaseGoogleSQLParserListener) EnterPivot_or_unpivot_clause_and_aliases(ctx *Pivot_or_unpivot_clause_and_aliasesContext) {
}

// ExitPivot_or_unpivot_clause_and_aliases is called when production pivot_or_unpivot_clause_and_aliases is exited.
func (s *BaseGoogleSQLParserListener) ExitPivot_or_unpivot_clause_and_aliases(ctx *Pivot_or_unpivot_clause_and_aliasesContext) {
}

// EnterAs_alias is called when production as_alias is entered.
func (s *BaseGoogleSQLParserListener) EnterAs_alias(ctx *As_aliasContext) {}

// ExitAs_alias is called when production as_alias is exited.
func (s *BaseGoogleSQLParserListener) ExitAs_alias(ctx *As_aliasContext) {}

// EnterSample_clause is called when production sample_clause is entered.
func (s *BaseGoogleSQLParserListener) EnterSample_clause(ctx *Sample_clauseContext) {}

// ExitSample_clause is called when production sample_clause is exited.
func (s *BaseGoogleSQLParserListener) ExitSample_clause(ctx *Sample_clauseContext) {}

// EnterOpt_sample_clause_suffix is called when production opt_sample_clause_suffix is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_sample_clause_suffix(ctx *Opt_sample_clause_suffixContext) {
}

// ExitOpt_sample_clause_suffix is called when production opt_sample_clause_suffix is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_sample_clause_suffix(ctx *Opt_sample_clause_suffixContext) {
}

// EnterRepeatable_clause is called when production repeatable_clause is entered.
func (s *BaseGoogleSQLParserListener) EnterRepeatable_clause(ctx *Repeatable_clauseContext) {}

// ExitRepeatable_clause is called when production repeatable_clause is exited.
func (s *BaseGoogleSQLParserListener) ExitRepeatable_clause(ctx *Repeatable_clauseContext) {}

// EnterPossibly_cast_int_literal_or_parameter is called when production possibly_cast_int_literal_or_parameter is entered.
func (s *BaseGoogleSQLParserListener) EnterPossibly_cast_int_literal_or_parameter(ctx *Possibly_cast_int_literal_or_parameterContext) {
}

// ExitPossibly_cast_int_literal_or_parameter is called when production possibly_cast_int_literal_or_parameter is exited.
func (s *BaseGoogleSQLParserListener) ExitPossibly_cast_int_literal_or_parameter(ctx *Possibly_cast_int_literal_or_parameterContext) {
}

// EnterCast_int_literal_or_parameter is called when production cast_int_literal_or_parameter is entered.
func (s *BaseGoogleSQLParserListener) EnterCast_int_literal_or_parameter(ctx *Cast_int_literal_or_parameterContext) {
}

// ExitCast_int_literal_or_parameter is called when production cast_int_literal_or_parameter is exited.
func (s *BaseGoogleSQLParserListener) ExitCast_int_literal_or_parameter(ctx *Cast_int_literal_or_parameterContext) {
}

// EnterSample_size is called when production sample_size is entered.
func (s *BaseGoogleSQLParserListener) EnterSample_size(ctx *Sample_sizeContext) {}

// ExitSample_size is called when production sample_size is exited.
func (s *BaseGoogleSQLParserListener) ExitSample_size(ctx *Sample_sizeContext) {}

// EnterSample_size_value is called when production sample_size_value is entered.
func (s *BaseGoogleSQLParserListener) EnterSample_size_value(ctx *Sample_size_valueContext) {}

// ExitSample_size_value is called when production sample_size_value is exited.
func (s *BaseGoogleSQLParserListener) ExitSample_size_value(ctx *Sample_size_valueContext) {}

// EnterSample_size_unit is called when production sample_size_unit is entered.
func (s *BaseGoogleSQLParserListener) EnterSample_size_unit(ctx *Sample_size_unitContext) {}

// ExitSample_size_unit is called when production sample_size_unit is exited.
func (s *BaseGoogleSQLParserListener) ExitSample_size_unit(ctx *Sample_size_unitContext) {}

// EnterPartition_by_clause_prefix_no_hint is called when production partition_by_clause_prefix_no_hint is entered.
func (s *BaseGoogleSQLParserListener) EnterPartition_by_clause_prefix_no_hint(ctx *Partition_by_clause_prefix_no_hintContext) {
}

// ExitPartition_by_clause_prefix_no_hint is called when production partition_by_clause_prefix_no_hint is exited.
func (s *BaseGoogleSQLParserListener) ExitPartition_by_clause_prefix_no_hint(ctx *Partition_by_clause_prefix_no_hintContext) {
}

// EnterMatch_recognize_clause is called when production match_recognize_clause is entered.
func (s *BaseGoogleSQLParserListener) EnterMatch_recognize_clause(ctx *Match_recognize_clauseContext) {
}

// ExitMatch_recognize_clause is called when production match_recognize_clause is exited.
func (s *BaseGoogleSQLParserListener) ExitMatch_recognize_clause(ctx *Match_recognize_clauseContext) {
}

// EnterRow_pattern_expr is called when production row_pattern_expr is entered.
func (s *BaseGoogleSQLParserListener) EnterRow_pattern_expr(ctx *Row_pattern_exprContext) {}

// ExitRow_pattern_expr is called when production row_pattern_expr is exited.
func (s *BaseGoogleSQLParserListener) ExitRow_pattern_expr(ctx *Row_pattern_exprContext) {}

// EnterRow_pattern_concatenation is called when production row_pattern_concatenation is entered.
func (s *BaseGoogleSQLParserListener) EnterRow_pattern_concatenation(ctx *Row_pattern_concatenationContext) {
}

// ExitRow_pattern_concatenation is called when production row_pattern_concatenation is exited.
func (s *BaseGoogleSQLParserListener) ExitRow_pattern_concatenation(ctx *Row_pattern_concatenationContext) {
}

// EnterRow_pattern_factor is called when production row_pattern_factor is entered.
func (s *BaseGoogleSQLParserListener) EnterRow_pattern_factor(ctx *Row_pattern_factorContext) {}

// ExitRow_pattern_factor is called when production row_pattern_factor is exited.
func (s *BaseGoogleSQLParserListener) ExitRow_pattern_factor(ctx *Row_pattern_factorContext) {}

// EnterSelect_list_prefix_with_as_aliases is called when production select_list_prefix_with_as_aliases is entered.
func (s *BaseGoogleSQLParserListener) EnterSelect_list_prefix_with_as_aliases(ctx *Select_list_prefix_with_as_aliasesContext) {
}

// ExitSelect_list_prefix_with_as_aliases is called when production select_list_prefix_with_as_aliases is exited.
func (s *BaseGoogleSQLParserListener) ExitSelect_list_prefix_with_as_aliases(ctx *Select_list_prefix_with_as_aliasesContext) {
}

// EnterSelect_column_expr_with_as_alias is called when production select_column_expr_with_as_alias is entered.
func (s *BaseGoogleSQLParserListener) EnterSelect_column_expr_with_as_alias(ctx *Select_column_expr_with_as_aliasContext) {
}

// ExitSelect_column_expr_with_as_alias is called when production select_column_expr_with_as_alias is exited.
func (s *BaseGoogleSQLParserListener) ExitSelect_column_expr_with_as_alias(ctx *Select_column_expr_with_as_aliasContext) {
}

// EnterTable_subquery is called when production table_subquery is entered.
func (s *BaseGoogleSQLParserListener) EnterTable_subquery(ctx *Table_subqueryContext) {}

// ExitTable_subquery is called when production table_subquery is exited.
func (s *BaseGoogleSQLParserListener) ExitTable_subquery(ctx *Table_subqueryContext) {}

// EnterJoin is called when production join is entered.
func (s *BaseGoogleSQLParserListener) EnterJoin(ctx *JoinContext) {}

// ExitJoin is called when production join is exited.
func (s *BaseGoogleSQLParserListener) ExitJoin(ctx *JoinContext) {}

// EnterJoin_item is called when production join_item is entered.
func (s *BaseGoogleSQLParserListener) EnterJoin_item(ctx *Join_itemContext) {}

// ExitJoin_item is called when production join_item is exited.
func (s *BaseGoogleSQLParserListener) ExitJoin_item(ctx *Join_itemContext) {}

// EnterOn_or_using_clause_list is called when production on_or_using_clause_list is entered.
func (s *BaseGoogleSQLParserListener) EnterOn_or_using_clause_list(ctx *On_or_using_clause_listContext) {
}

// ExitOn_or_using_clause_list is called when production on_or_using_clause_list is exited.
func (s *BaseGoogleSQLParserListener) ExitOn_or_using_clause_list(ctx *On_or_using_clause_listContext) {
}

// EnterOn_or_using_clause is called when production on_or_using_clause is entered.
func (s *BaseGoogleSQLParserListener) EnterOn_or_using_clause(ctx *On_or_using_clauseContext) {}

// ExitOn_or_using_clause is called when production on_or_using_clause is exited.
func (s *BaseGoogleSQLParserListener) ExitOn_or_using_clause(ctx *On_or_using_clauseContext) {}

// EnterUsing_clause is called when production using_clause is entered.
func (s *BaseGoogleSQLParserListener) EnterUsing_clause(ctx *Using_clauseContext) {}

// ExitUsing_clause is called when production using_clause is exited.
func (s *BaseGoogleSQLParserListener) ExitUsing_clause(ctx *Using_clauseContext) {}

// EnterJoin_hint is called when production join_hint is entered.
func (s *BaseGoogleSQLParserListener) EnterJoin_hint(ctx *Join_hintContext) {}

// ExitJoin_hint is called when production join_hint is exited.
func (s *BaseGoogleSQLParserListener) ExitJoin_hint(ctx *Join_hintContext) {}

// EnterTable_path_expression is called when production table_path_expression is entered.
func (s *BaseGoogleSQLParserListener) EnterTable_path_expression(ctx *Table_path_expressionContext) {}

// ExitTable_path_expression is called when production table_path_expression is exited.
func (s *BaseGoogleSQLParserListener) ExitTable_path_expression(ctx *Table_path_expressionContext) {}

// EnterOpt_at_system_time is called when production opt_at_system_time is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_at_system_time(ctx *Opt_at_system_timeContext) {}

// ExitOpt_at_system_time is called when production opt_at_system_time is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_at_system_time(ctx *Opt_at_system_timeContext) {}

// EnterOpt_with_offset_and_alias is called when production opt_with_offset_and_alias is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_with_offset_and_alias(ctx *Opt_with_offset_and_aliasContext) {
}

// ExitOpt_with_offset_and_alias is called when production opt_with_offset_and_alias is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_with_offset_and_alias(ctx *Opt_with_offset_and_aliasContext) {
}

// EnterOpt_pivot_or_unpivot_clause_and_alias is called when production opt_pivot_or_unpivot_clause_and_alias is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_pivot_or_unpivot_clause_and_alias(ctx *Opt_pivot_or_unpivot_clause_and_aliasContext) {
}

// ExitOpt_pivot_or_unpivot_clause_and_alias is called when production opt_pivot_or_unpivot_clause_and_alias is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_pivot_or_unpivot_clause_and_alias(ctx *Opt_pivot_or_unpivot_clause_and_aliasContext) {
}

// EnterTable_path_expression_base is called when production table_path_expression_base is entered.
func (s *BaseGoogleSQLParserListener) EnterTable_path_expression_base(ctx *Table_path_expression_baseContext) {
}

// ExitTable_path_expression_base is called when production table_path_expression_base is exited.
func (s *BaseGoogleSQLParserListener) ExitTable_path_expression_base(ctx *Table_path_expression_baseContext) {
}

// EnterMaybe_slashed_or_dashed_path_expression is called when production maybe_slashed_or_dashed_path_expression is entered.
func (s *BaseGoogleSQLParserListener) EnterMaybe_slashed_or_dashed_path_expression(ctx *Maybe_slashed_or_dashed_path_expressionContext) {
}

// ExitMaybe_slashed_or_dashed_path_expression is called when production maybe_slashed_or_dashed_path_expression is exited.
func (s *BaseGoogleSQLParserListener) ExitMaybe_slashed_or_dashed_path_expression(ctx *Maybe_slashed_or_dashed_path_expressionContext) {
}

// EnterMaybe_dashed_path_expression is called when production maybe_dashed_path_expression is entered.
func (s *BaseGoogleSQLParserListener) EnterMaybe_dashed_path_expression(ctx *Maybe_dashed_path_expressionContext) {
}

// ExitMaybe_dashed_path_expression is called when production maybe_dashed_path_expression is exited.
func (s *BaseGoogleSQLParserListener) ExitMaybe_dashed_path_expression(ctx *Maybe_dashed_path_expressionContext) {
}

// EnterDashed_path_expression is called when production dashed_path_expression is entered.
func (s *BaseGoogleSQLParserListener) EnterDashed_path_expression(ctx *Dashed_path_expressionContext) {
}

// ExitDashed_path_expression is called when production dashed_path_expression is exited.
func (s *BaseGoogleSQLParserListener) ExitDashed_path_expression(ctx *Dashed_path_expressionContext) {
}

// EnterDashed_identifier is called when production dashed_identifier is entered.
func (s *BaseGoogleSQLParserListener) EnterDashed_identifier(ctx *Dashed_identifierContext) {}

// ExitDashed_identifier is called when production dashed_identifier is exited.
func (s *BaseGoogleSQLParserListener) ExitDashed_identifier(ctx *Dashed_identifierContext) {}

// EnterSlashed_identifier is called when production slashed_identifier is entered.
func (s *BaseGoogleSQLParserListener) EnterSlashed_identifier(ctx *Slashed_identifierContext) {}

// ExitSlashed_identifier is called when production slashed_identifier is exited.
func (s *BaseGoogleSQLParserListener) ExitSlashed_identifier(ctx *Slashed_identifierContext) {}

// EnterIdentifier_or_integer is called when production identifier_or_integer is entered.
func (s *BaseGoogleSQLParserListener) EnterIdentifier_or_integer(ctx *Identifier_or_integerContext) {}

// ExitIdentifier_or_integer is called when production identifier_or_integer is exited.
func (s *BaseGoogleSQLParserListener) ExitIdentifier_or_integer(ctx *Identifier_or_integerContext) {}

// EnterSlashed_identifier_separator is called when production slashed_identifier_separator is entered.
func (s *BaseGoogleSQLParserListener) EnterSlashed_identifier_separator(ctx *Slashed_identifier_separatorContext) {
}

// ExitSlashed_identifier_separator is called when production slashed_identifier_separator is exited.
func (s *BaseGoogleSQLParserListener) ExitSlashed_identifier_separator(ctx *Slashed_identifier_separatorContext) {
}

// EnterSlashed_path_expression is called when production slashed_path_expression is entered.
func (s *BaseGoogleSQLParserListener) EnterSlashed_path_expression(ctx *Slashed_path_expressionContext) {
}

// ExitSlashed_path_expression is called when production slashed_path_expression is exited.
func (s *BaseGoogleSQLParserListener) ExitSlashed_path_expression(ctx *Slashed_path_expressionContext) {
}

// EnterUnnest_expression is called when production unnest_expression is entered.
func (s *BaseGoogleSQLParserListener) EnterUnnest_expression(ctx *Unnest_expressionContext) {}

// ExitUnnest_expression is called when production unnest_expression is exited.
func (s *BaseGoogleSQLParserListener) ExitUnnest_expression(ctx *Unnest_expressionContext) {}

// EnterUnnest_expression_prefix is called when production unnest_expression_prefix is entered.
func (s *BaseGoogleSQLParserListener) EnterUnnest_expression_prefix(ctx *Unnest_expression_prefixContext) {
}

// ExitUnnest_expression_prefix is called when production unnest_expression_prefix is exited.
func (s *BaseGoogleSQLParserListener) ExitUnnest_expression_prefix(ctx *Unnest_expression_prefixContext) {
}

// EnterOpt_array_zip_mode is called when production opt_array_zip_mode is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_array_zip_mode(ctx *Opt_array_zip_modeContext) {}

// ExitOpt_array_zip_mode is called when production opt_array_zip_mode is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_array_zip_mode(ctx *Opt_array_zip_modeContext) {}

// EnterExpression_with_opt_alias is called when production expression_with_opt_alias is entered.
func (s *BaseGoogleSQLParserListener) EnterExpression_with_opt_alias(ctx *Expression_with_opt_aliasContext) {
}

// ExitExpression_with_opt_alias is called when production expression_with_opt_alias is exited.
func (s *BaseGoogleSQLParserListener) ExitExpression_with_opt_alias(ctx *Expression_with_opt_aliasContext) {
}

// EnterTvf_prefix is called when production tvf_prefix is entered.
func (s *BaseGoogleSQLParserListener) EnterTvf_prefix(ctx *Tvf_prefixContext) {}

// ExitTvf_prefix is called when production tvf_prefix is exited.
func (s *BaseGoogleSQLParserListener) ExitTvf_prefix(ctx *Tvf_prefixContext) {}

// EnterTvf_argument is called when production tvf_argument is entered.
func (s *BaseGoogleSQLParserListener) EnterTvf_argument(ctx *Tvf_argumentContext) {}

// ExitTvf_argument is called when production tvf_argument is exited.
func (s *BaseGoogleSQLParserListener) ExitTvf_argument(ctx *Tvf_argumentContext) {}

// EnterConnection_clause is called when production connection_clause is entered.
func (s *BaseGoogleSQLParserListener) EnterConnection_clause(ctx *Connection_clauseContext) {}

// ExitConnection_clause is called when production connection_clause is exited.
func (s *BaseGoogleSQLParserListener) ExitConnection_clause(ctx *Connection_clauseContext) {}

// EnterPath_expression_or_default is called when production path_expression_or_default is entered.
func (s *BaseGoogleSQLParserListener) EnterPath_expression_or_default(ctx *Path_expression_or_defaultContext) {
}

// ExitPath_expression_or_default is called when production path_expression_or_default is exited.
func (s *BaseGoogleSQLParserListener) ExitPath_expression_or_default(ctx *Path_expression_or_defaultContext) {
}

// EnterDescriptor_argument is called when production descriptor_argument is entered.
func (s *BaseGoogleSQLParserListener) EnterDescriptor_argument(ctx *Descriptor_argumentContext) {}

// ExitDescriptor_argument is called when production descriptor_argument is exited.
func (s *BaseGoogleSQLParserListener) ExitDescriptor_argument(ctx *Descriptor_argumentContext) {}

// EnterDescriptor_column_list is called when production descriptor_column_list is entered.
func (s *BaseGoogleSQLParserListener) EnterDescriptor_column_list(ctx *Descriptor_column_listContext) {
}

// ExitDescriptor_column_list is called when production descriptor_column_list is exited.
func (s *BaseGoogleSQLParserListener) ExitDescriptor_column_list(ctx *Descriptor_column_listContext) {
}

// EnterDescriptor_column is called when production descriptor_column is entered.
func (s *BaseGoogleSQLParserListener) EnterDescriptor_column(ctx *Descriptor_columnContext) {}

// ExitDescriptor_column is called when production descriptor_column is exited.
func (s *BaseGoogleSQLParserListener) ExitDescriptor_column(ctx *Descriptor_columnContext) {}

// EnterTable_clause is called when production table_clause is entered.
func (s *BaseGoogleSQLParserListener) EnterTable_clause(ctx *Table_clauseContext) {}

// ExitTable_clause is called when production table_clause is exited.
func (s *BaseGoogleSQLParserListener) ExitTable_clause(ctx *Table_clauseContext) {}

// EnterModel_clause is called when production model_clause is entered.
func (s *BaseGoogleSQLParserListener) EnterModel_clause(ctx *Model_clauseContext) {}

// ExitModel_clause is called when production model_clause is exited.
func (s *BaseGoogleSQLParserListener) ExitModel_clause(ctx *Model_clauseContext) {}

// EnterQualify_clause_nonreserved is called when production qualify_clause_nonreserved is entered.
func (s *BaseGoogleSQLParserListener) EnterQualify_clause_nonreserved(ctx *Qualify_clause_nonreservedContext) {
}

// ExitQualify_clause_nonreserved is called when production qualify_clause_nonreserved is exited.
func (s *BaseGoogleSQLParserListener) ExitQualify_clause_nonreserved(ctx *Qualify_clause_nonreservedContext) {
}

// EnterUnpivot_clause is called when production unpivot_clause is entered.
func (s *BaseGoogleSQLParserListener) EnterUnpivot_clause(ctx *Unpivot_clauseContext) {}

// ExitUnpivot_clause is called when production unpivot_clause is exited.
func (s *BaseGoogleSQLParserListener) ExitUnpivot_clause(ctx *Unpivot_clauseContext) {}

// EnterUnpivot_in_item_list is called when production unpivot_in_item_list is entered.
func (s *BaseGoogleSQLParserListener) EnterUnpivot_in_item_list(ctx *Unpivot_in_item_listContext) {}

// ExitUnpivot_in_item_list is called when production unpivot_in_item_list is exited.
func (s *BaseGoogleSQLParserListener) ExitUnpivot_in_item_list(ctx *Unpivot_in_item_listContext) {}

// EnterUnpivot_in_item_list_prefix is called when production unpivot_in_item_list_prefix is entered.
func (s *BaseGoogleSQLParserListener) EnterUnpivot_in_item_list_prefix(ctx *Unpivot_in_item_list_prefixContext) {
}

// ExitUnpivot_in_item_list_prefix is called when production unpivot_in_item_list_prefix is exited.
func (s *BaseGoogleSQLParserListener) ExitUnpivot_in_item_list_prefix(ctx *Unpivot_in_item_list_prefixContext) {
}

// EnterUnpivot_in_item is called when production unpivot_in_item is entered.
func (s *BaseGoogleSQLParserListener) EnterUnpivot_in_item(ctx *Unpivot_in_itemContext) {}

// ExitUnpivot_in_item is called when production unpivot_in_item is exited.
func (s *BaseGoogleSQLParserListener) ExitUnpivot_in_item(ctx *Unpivot_in_itemContext) {}

// EnterOpt_as_string_or_integer is called when production opt_as_string_or_integer is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_as_string_or_integer(ctx *Opt_as_string_or_integerContext) {
}

// ExitOpt_as_string_or_integer is called when production opt_as_string_or_integer is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_as_string_or_integer(ctx *Opt_as_string_or_integerContext) {
}

// EnterPath_expression_list_with_opt_parens is called when production path_expression_list_with_opt_parens is entered.
func (s *BaseGoogleSQLParserListener) EnterPath_expression_list_with_opt_parens(ctx *Path_expression_list_with_opt_parensContext) {
}

// ExitPath_expression_list_with_opt_parens is called when production path_expression_list_with_opt_parens is exited.
func (s *BaseGoogleSQLParserListener) ExitPath_expression_list_with_opt_parens(ctx *Path_expression_list_with_opt_parensContext) {
}

// EnterPath_expression_list is called when production path_expression_list is entered.
func (s *BaseGoogleSQLParserListener) EnterPath_expression_list(ctx *Path_expression_listContext) {}

// ExitPath_expression_list is called when production path_expression_list is exited.
func (s *BaseGoogleSQLParserListener) ExitPath_expression_list(ctx *Path_expression_listContext) {}

// EnterUnpivot_nulls_filter is called when production unpivot_nulls_filter is entered.
func (s *BaseGoogleSQLParserListener) EnterUnpivot_nulls_filter(ctx *Unpivot_nulls_filterContext) {}

// ExitUnpivot_nulls_filter is called when production unpivot_nulls_filter is exited.
func (s *BaseGoogleSQLParserListener) ExitUnpivot_nulls_filter(ctx *Unpivot_nulls_filterContext) {}

// EnterPivot_clause is called when production pivot_clause is entered.
func (s *BaseGoogleSQLParserListener) EnterPivot_clause(ctx *Pivot_clauseContext) {}

// ExitPivot_clause is called when production pivot_clause is exited.
func (s *BaseGoogleSQLParserListener) ExitPivot_clause(ctx *Pivot_clauseContext) {}

// EnterPivot_expression_list is called when production pivot_expression_list is entered.
func (s *BaseGoogleSQLParserListener) EnterPivot_expression_list(ctx *Pivot_expression_listContext) {}

// ExitPivot_expression_list is called when production pivot_expression_list is exited.
func (s *BaseGoogleSQLParserListener) ExitPivot_expression_list(ctx *Pivot_expression_listContext) {}

// EnterPivot_expression is called when production pivot_expression is entered.
func (s *BaseGoogleSQLParserListener) EnterPivot_expression(ctx *Pivot_expressionContext) {}

// ExitPivot_expression is called when production pivot_expression is exited.
func (s *BaseGoogleSQLParserListener) ExitPivot_expression(ctx *Pivot_expressionContext) {}

// EnterPivot_value_list is called when production pivot_value_list is entered.
func (s *BaseGoogleSQLParserListener) EnterPivot_value_list(ctx *Pivot_value_listContext) {}

// ExitPivot_value_list is called when production pivot_value_list is exited.
func (s *BaseGoogleSQLParserListener) ExitPivot_value_list(ctx *Pivot_value_listContext) {}

// EnterPivot_value is called when production pivot_value is entered.
func (s *BaseGoogleSQLParserListener) EnterPivot_value(ctx *Pivot_valueContext) {}

// ExitPivot_value is called when production pivot_value is exited.
func (s *BaseGoogleSQLParserListener) ExitPivot_value(ctx *Pivot_valueContext) {}

// EnterTvf_prefix_no_args is called when production tvf_prefix_no_args is entered.
func (s *BaseGoogleSQLParserListener) EnterTvf_prefix_no_args(ctx *Tvf_prefix_no_argsContext) {}

// ExitTvf_prefix_no_args is called when production tvf_prefix_no_args is exited.
func (s *BaseGoogleSQLParserListener) ExitTvf_prefix_no_args(ctx *Tvf_prefix_no_argsContext) {}

// EnterJoin_type is called when production join_type is entered.
func (s *BaseGoogleSQLParserListener) EnterJoin_type(ctx *Join_typeContext) {}

// ExitJoin_type is called when production join_type is exited.
func (s *BaseGoogleSQLParserListener) ExitJoin_type(ctx *Join_typeContext) {}

// EnterOpt_natural is called when production opt_natural is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_natural(ctx *Opt_naturalContext) {}

// ExitOpt_natural is called when production opt_natural is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_natural(ctx *Opt_naturalContext) {}

// EnterOn_clause is called when production on_clause is entered.
func (s *BaseGoogleSQLParserListener) EnterOn_clause(ctx *On_clauseContext) {}

// ExitOn_clause is called when production on_clause is exited.
func (s *BaseGoogleSQLParserListener) ExitOn_clause(ctx *On_clauseContext) {}

// EnterSelect_list is called when production select_list is entered.
func (s *BaseGoogleSQLParserListener) EnterSelect_list(ctx *Select_listContext) {}

// ExitSelect_list is called when production select_list is exited.
func (s *BaseGoogleSQLParserListener) ExitSelect_list(ctx *Select_listContext) {}

// EnterSelect_list_item is called when production select_list_item is entered.
func (s *BaseGoogleSQLParserListener) EnterSelect_list_item(ctx *Select_list_itemContext) {}

// ExitSelect_list_item is called when production select_list_item is exited.
func (s *BaseGoogleSQLParserListener) ExitSelect_list_item(ctx *Select_list_itemContext) {}

// EnterSelect_column_star is called when production select_column_star is entered.
func (s *BaseGoogleSQLParserListener) EnterSelect_column_star(ctx *Select_column_starContext) {}

// ExitSelect_column_star is called when production select_column_star is exited.
func (s *BaseGoogleSQLParserListener) ExitSelect_column_star(ctx *Select_column_starContext) {}

// EnterSelect_column_expr is called when production select_column_expr is entered.
func (s *BaseGoogleSQLParserListener) EnterSelect_column_expr(ctx *Select_column_exprContext) {}

// ExitSelect_column_expr is called when production select_column_expr is exited.
func (s *BaseGoogleSQLParserListener) ExitSelect_column_expr(ctx *Select_column_exprContext) {}

// EnterSelect_column_dot_star is called when production select_column_dot_star is entered.
func (s *BaseGoogleSQLParserListener) EnterSelect_column_dot_star(ctx *Select_column_dot_starContext) {
}

// ExitSelect_column_dot_star is called when production select_column_dot_star is exited.
func (s *BaseGoogleSQLParserListener) ExitSelect_column_dot_star(ctx *Select_column_dot_starContext) {
}

// EnterStar_modifiers is called when production star_modifiers is entered.
func (s *BaseGoogleSQLParserListener) EnterStar_modifiers(ctx *Star_modifiersContext) {}

// ExitStar_modifiers is called when production star_modifiers is exited.
func (s *BaseGoogleSQLParserListener) ExitStar_modifiers(ctx *Star_modifiersContext) {}

// EnterStar_except_list is called when production star_except_list is entered.
func (s *BaseGoogleSQLParserListener) EnterStar_except_list(ctx *Star_except_listContext) {}

// ExitStar_except_list is called when production star_except_list is exited.
func (s *BaseGoogleSQLParserListener) ExitStar_except_list(ctx *Star_except_listContext) {}

// EnterStar_replace_list is called when production star_replace_list is entered.
func (s *BaseGoogleSQLParserListener) EnterStar_replace_list(ctx *Star_replace_listContext) {}

// ExitStar_replace_list is called when production star_replace_list is exited.
func (s *BaseGoogleSQLParserListener) ExitStar_replace_list(ctx *Star_replace_listContext) {}

// EnterStar_replace_item is called when production star_replace_item is entered.
func (s *BaseGoogleSQLParserListener) EnterStar_replace_item(ctx *Star_replace_itemContext) {}

// ExitStar_replace_item is called when production star_replace_item is exited.
func (s *BaseGoogleSQLParserListener) ExitStar_replace_item(ctx *Star_replace_itemContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseGoogleSQLParserListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseGoogleSQLParserListener) ExitExpression(ctx *ExpressionContext) {}

// EnterExpression_higher_prec_than_and is called when production expression_higher_prec_than_and is entered.
func (s *BaseGoogleSQLParserListener) EnterExpression_higher_prec_than_and(ctx *Expression_higher_prec_than_andContext) {
}

// ExitExpression_higher_prec_than_and is called when production expression_higher_prec_than_and is exited.
func (s *BaseGoogleSQLParserListener) ExitExpression_higher_prec_than_and(ctx *Expression_higher_prec_than_andContext) {
}

// EnterExpression_maybe_parenthesized_not_a_query is called when production expression_maybe_parenthesized_not_a_query is entered.
func (s *BaseGoogleSQLParserListener) EnterExpression_maybe_parenthesized_not_a_query(ctx *Expression_maybe_parenthesized_not_a_queryContext) {
}

// ExitExpression_maybe_parenthesized_not_a_query is called when production expression_maybe_parenthesized_not_a_query is exited.
func (s *BaseGoogleSQLParserListener) ExitExpression_maybe_parenthesized_not_a_query(ctx *Expression_maybe_parenthesized_not_a_queryContext) {
}

// EnterParenthesized_in_rhs is called when production parenthesized_in_rhs is entered.
func (s *BaseGoogleSQLParserListener) EnterParenthesized_in_rhs(ctx *Parenthesized_in_rhsContext) {}

// ExitParenthesized_in_rhs is called when production parenthesized_in_rhs is exited.
func (s *BaseGoogleSQLParserListener) ExitParenthesized_in_rhs(ctx *Parenthesized_in_rhsContext) {}

// EnterUnary_operator is called when production unary_operator is entered.
func (s *BaseGoogleSQLParserListener) EnterUnary_operator(ctx *Unary_operatorContext) {}

// ExitUnary_operator is called when production unary_operator is exited.
func (s *BaseGoogleSQLParserListener) ExitUnary_operator(ctx *Unary_operatorContext) {}

// EnterComparative_operator is called when production comparative_operator is entered.
func (s *BaseGoogleSQLParserListener) EnterComparative_operator(ctx *Comparative_operatorContext) {}

// ExitComparative_operator is called when production comparative_operator is exited.
func (s *BaseGoogleSQLParserListener) ExitComparative_operator(ctx *Comparative_operatorContext) {}

// EnterShift_operator is called when production shift_operator is entered.
func (s *BaseGoogleSQLParserListener) EnterShift_operator(ctx *Shift_operatorContext) {}

// ExitShift_operator is called when production shift_operator is exited.
func (s *BaseGoogleSQLParserListener) ExitShift_operator(ctx *Shift_operatorContext) {}

// EnterAdditive_operator is called when production additive_operator is entered.
func (s *BaseGoogleSQLParserListener) EnterAdditive_operator(ctx *Additive_operatorContext) {}

// ExitAdditive_operator is called when production additive_operator is exited.
func (s *BaseGoogleSQLParserListener) ExitAdditive_operator(ctx *Additive_operatorContext) {}

// EnterMultiplicative_operator is called when production multiplicative_operator is entered.
func (s *BaseGoogleSQLParserListener) EnterMultiplicative_operator(ctx *Multiplicative_operatorContext) {
}

// ExitMultiplicative_operator is called when production multiplicative_operator is exited.
func (s *BaseGoogleSQLParserListener) ExitMultiplicative_operator(ctx *Multiplicative_operatorContext) {
}

// EnterIs_operator is called when production is_operator is entered.
func (s *BaseGoogleSQLParserListener) EnterIs_operator(ctx *Is_operatorContext) {}

// ExitIs_operator is called when production is_operator is exited.
func (s *BaseGoogleSQLParserListener) ExitIs_operator(ctx *Is_operatorContext) {}

// EnterBetween_operator is called when production between_operator is entered.
func (s *BaseGoogleSQLParserListener) EnterBetween_operator(ctx *Between_operatorContext) {}

// ExitBetween_operator is called when production between_operator is exited.
func (s *BaseGoogleSQLParserListener) ExitBetween_operator(ctx *Between_operatorContext) {}

// EnterIn_operator is called when production in_operator is entered.
func (s *BaseGoogleSQLParserListener) EnterIn_operator(ctx *In_operatorContext) {}

// ExitIn_operator is called when production in_operator is exited.
func (s *BaseGoogleSQLParserListener) ExitIn_operator(ctx *In_operatorContext) {}

// EnterDistinct_operator is called when production distinct_operator is entered.
func (s *BaseGoogleSQLParserListener) EnterDistinct_operator(ctx *Distinct_operatorContext) {}

// ExitDistinct_operator is called when production distinct_operator is exited.
func (s *BaseGoogleSQLParserListener) ExitDistinct_operator(ctx *Distinct_operatorContext) {}

// EnterParenthesized_query is called when production parenthesized_query is entered.
func (s *BaseGoogleSQLParserListener) EnterParenthesized_query(ctx *Parenthesized_queryContext) {}

// ExitParenthesized_query is called when production parenthesized_query is exited.
func (s *BaseGoogleSQLParserListener) ExitParenthesized_query(ctx *Parenthesized_queryContext) {}

// EnterParenthesized_expression_not_a_query is called when production parenthesized_expression_not_a_query is entered.
func (s *BaseGoogleSQLParserListener) EnterParenthesized_expression_not_a_query(ctx *Parenthesized_expression_not_a_queryContext) {
}

// ExitParenthesized_expression_not_a_query is called when production parenthesized_expression_not_a_query is exited.
func (s *BaseGoogleSQLParserListener) ExitParenthesized_expression_not_a_query(ctx *Parenthesized_expression_not_a_queryContext) {
}

// EnterParenthesized_anysomeall_list_in_rhs is called when production parenthesized_anysomeall_list_in_rhs is entered.
func (s *BaseGoogleSQLParserListener) EnterParenthesized_anysomeall_list_in_rhs(ctx *Parenthesized_anysomeall_list_in_rhsContext) {
}

// ExitParenthesized_anysomeall_list_in_rhs is called when production parenthesized_anysomeall_list_in_rhs is exited.
func (s *BaseGoogleSQLParserListener) ExitParenthesized_anysomeall_list_in_rhs(ctx *Parenthesized_anysomeall_list_in_rhsContext) {
}

// EnterAnd_expression is called when production and_expression is entered.
func (s *BaseGoogleSQLParserListener) EnterAnd_expression(ctx *And_expressionContext) {}

// ExitAnd_expression is called when production and_expression is exited.
func (s *BaseGoogleSQLParserListener) ExitAnd_expression(ctx *And_expressionContext) {}

// EnterIn_list_two_or_more_prefix is called when production in_list_two_or_more_prefix is entered.
func (s *BaseGoogleSQLParserListener) EnterIn_list_two_or_more_prefix(ctx *In_list_two_or_more_prefixContext) {
}

// ExitIn_list_two_or_more_prefix is called when production in_list_two_or_more_prefix is exited.
func (s *BaseGoogleSQLParserListener) ExitIn_list_two_or_more_prefix(ctx *In_list_two_or_more_prefixContext) {
}

// EnterAny_some_all is called when production any_some_all is entered.
func (s *BaseGoogleSQLParserListener) EnterAny_some_all(ctx *Any_some_allContext) {}

// ExitAny_some_all is called when production any_some_all is exited.
func (s *BaseGoogleSQLParserListener) ExitAny_some_all(ctx *Any_some_allContext) {}

// EnterLike_operator is called when production like_operator is entered.
func (s *BaseGoogleSQLParserListener) EnterLike_operator(ctx *Like_operatorContext) {}

// ExitLike_operator is called when production like_operator is exited.
func (s *BaseGoogleSQLParserListener) ExitLike_operator(ctx *Like_operatorContext) {}

// EnterExpression_subquery_with_keyword is called when production expression_subquery_with_keyword is entered.
func (s *BaseGoogleSQLParserListener) EnterExpression_subquery_with_keyword(ctx *Expression_subquery_with_keywordContext) {
}

// ExitExpression_subquery_with_keyword is called when production expression_subquery_with_keyword is exited.
func (s *BaseGoogleSQLParserListener) ExitExpression_subquery_with_keyword(ctx *Expression_subquery_with_keywordContext) {
}

// EnterStruct_constructor is called when production struct_constructor is entered.
func (s *BaseGoogleSQLParserListener) EnterStruct_constructor(ctx *Struct_constructorContext) {}

// ExitStruct_constructor is called when production struct_constructor is exited.
func (s *BaseGoogleSQLParserListener) ExitStruct_constructor(ctx *Struct_constructorContext) {}

// EnterStruct_constructor_prefix_with_keyword is called when production struct_constructor_prefix_with_keyword is entered.
func (s *BaseGoogleSQLParserListener) EnterStruct_constructor_prefix_with_keyword(ctx *Struct_constructor_prefix_with_keywordContext) {
}

// ExitStruct_constructor_prefix_with_keyword is called when production struct_constructor_prefix_with_keyword is exited.
func (s *BaseGoogleSQLParserListener) ExitStruct_constructor_prefix_with_keyword(ctx *Struct_constructor_prefix_with_keywordContext) {
}

// EnterStruct_constructor_arg is called when production struct_constructor_arg is entered.
func (s *BaseGoogleSQLParserListener) EnterStruct_constructor_arg(ctx *Struct_constructor_argContext) {
}

// ExitStruct_constructor_arg is called when production struct_constructor_arg is exited.
func (s *BaseGoogleSQLParserListener) ExitStruct_constructor_arg(ctx *Struct_constructor_argContext) {
}

// EnterStruct_constructor_prefix_without_keyword is called when production struct_constructor_prefix_without_keyword is entered.
func (s *BaseGoogleSQLParserListener) EnterStruct_constructor_prefix_without_keyword(ctx *Struct_constructor_prefix_without_keywordContext) {
}

// ExitStruct_constructor_prefix_without_keyword is called when production struct_constructor_prefix_without_keyword is exited.
func (s *BaseGoogleSQLParserListener) ExitStruct_constructor_prefix_without_keyword(ctx *Struct_constructor_prefix_without_keywordContext) {
}

// EnterStruct_constructor_prefix_with_keyword_no_arg is called when production struct_constructor_prefix_with_keyword_no_arg is entered.
func (s *BaseGoogleSQLParserListener) EnterStruct_constructor_prefix_with_keyword_no_arg(ctx *Struct_constructor_prefix_with_keyword_no_argContext) {
}

// ExitStruct_constructor_prefix_with_keyword_no_arg is called when production struct_constructor_prefix_with_keyword_no_arg is exited.
func (s *BaseGoogleSQLParserListener) ExitStruct_constructor_prefix_with_keyword_no_arg(ctx *Struct_constructor_prefix_with_keyword_no_argContext) {
}

// EnterInterval_expression is called when production interval_expression is entered.
func (s *BaseGoogleSQLParserListener) EnterInterval_expression(ctx *Interval_expressionContext) {}

// ExitInterval_expression is called when production interval_expression is exited.
func (s *BaseGoogleSQLParserListener) ExitInterval_expression(ctx *Interval_expressionContext) {}

// EnterFunction_call_expression_with_clauses is called when production function_call_expression_with_clauses is entered.
func (s *BaseGoogleSQLParserListener) EnterFunction_call_expression_with_clauses(ctx *Function_call_expression_with_clausesContext) {
}

// ExitFunction_call_expression_with_clauses is called when production function_call_expression_with_clauses is exited.
func (s *BaseGoogleSQLParserListener) ExitFunction_call_expression_with_clauses(ctx *Function_call_expression_with_clausesContext) {
}

// EnterFunction_call_expression_with_clauses_suffix is called when production function_call_expression_with_clauses_suffix is entered.
func (s *BaseGoogleSQLParserListener) EnterFunction_call_expression_with_clauses_suffix(ctx *Function_call_expression_with_clauses_suffixContext) {
}

// ExitFunction_call_expression_with_clauses_suffix is called when production function_call_expression_with_clauses_suffix is exited.
func (s *BaseGoogleSQLParserListener) ExitFunction_call_expression_with_clauses_suffix(ctx *Function_call_expression_with_clauses_suffixContext) {
}

// EnterOver_clause is called when production over_clause is entered.
func (s *BaseGoogleSQLParserListener) EnterOver_clause(ctx *Over_clauseContext) {}

// ExitOver_clause is called when production over_clause is exited.
func (s *BaseGoogleSQLParserListener) ExitOver_clause(ctx *Over_clauseContext) {}

// EnterWindow_specification is called when production window_specification is entered.
func (s *BaseGoogleSQLParserListener) EnterWindow_specification(ctx *Window_specificationContext) {}

// ExitWindow_specification is called when production window_specification is exited.
func (s *BaseGoogleSQLParserListener) ExitWindow_specification(ctx *Window_specificationContext) {}

// EnterOpt_window_frame_clause is called when production opt_window_frame_clause is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_window_frame_clause(ctx *Opt_window_frame_clauseContext) {
}

// ExitOpt_window_frame_clause is called when production opt_window_frame_clause is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_window_frame_clause(ctx *Opt_window_frame_clauseContext) {
}

// EnterWindow_frame_bound is called when production window_frame_bound is entered.
func (s *BaseGoogleSQLParserListener) EnterWindow_frame_bound(ctx *Window_frame_boundContext) {}

// ExitWindow_frame_bound is called when production window_frame_bound is exited.
func (s *BaseGoogleSQLParserListener) ExitWindow_frame_bound(ctx *Window_frame_boundContext) {}

// EnterPreceding_or_following is called when production preceding_or_following is entered.
func (s *BaseGoogleSQLParserListener) EnterPreceding_or_following(ctx *Preceding_or_followingContext) {
}

// ExitPreceding_or_following is called when production preceding_or_following is exited.
func (s *BaseGoogleSQLParserListener) ExitPreceding_or_following(ctx *Preceding_or_followingContext) {
}

// EnterFrame_unit is called when production frame_unit is entered.
func (s *BaseGoogleSQLParserListener) EnterFrame_unit(ctx *Frame_unitContext) {}

// ExitFrame_unit is called when production frame_unit is exited.
func (s *BaseGoogleSQLParserListener) ExitFrame_unit(ctx *Frame_unitContext) {}

// EnterPartition_by_clause is called when production partition_by_clause is entered.
func (s *BaseGoogleSQLParserListener) EnterPartition_by_clause(ctx *Partition_by_clauseContext) {}

// ExitPartition_by_clause is called when production partition_by_clause is exited.
func (s *BaseGoogleSQLParserListener) ExitPartition_by_clause(ctx *Partition_by_clauseContext) {}

// EnterPartition_by_clause_prefix is called when production partition_by_clause_prefix is entered.
func (s *BaseGoogleSQLParserListener) EnterPartition_by_clause_prefix(ctx *Partition_by_clause_prefixContext) {
}

// ExitPartition_by_clause_prefix is called when production partition_by_clause_prefix is exited.
func (s *BaseGoogleSQLParserListener) ExitPartition_by_clause_prefix(ctx *Partition_by_clause_prefixContext) {
}

// EnterWith_group_rows is called when production with_group_rows is entered.
func (s *BaseGoogleSQLParserListener) EnterWith_group_rows(ctx *With_group_rowsContext) {}

// ExitWith_group_rows is called when production with_group_rows is exited.
func (s *BaseGoogleSQLParserListener) ExitWith_group_rows(ctx *With_group_rowsContext) {}

// EnterWith_report_modifier is called when production with_report_modifier is entered.
func (s *BaseGoogleSQLParserListener) EnterWith_report_modifier(ctx *With_report_modifierContext) {}

// ExitWith_report_modifier is called when production with_report_modifier is exited.
func (s *BaseGoogleSQLParserListener) ExitWith_report_modifier(ctx *With_report_modifierContext) {}

// EnterClamped_between_modifier is called when production clamped_between_modifier is entered.
func (s *BaseGoogleSQLParserListener) EnterClamped_between_modifier(ctx *Clamped_between_modifierContext) {
}

// ExitClamped_between_modifier is called when production clamped_between_modifier is exited.
func (s *BaseGoogleSQLParserListener) ExitClamped_between_modifier(ctx *Clamped_between_modifierContext) {
}

// EnterWith_report_format is called when production with_report_format is entered.
func (s *BaseGoogleSQLParserListener) EnterWith_report_format(ctx *With_report_formatContext) {}

// ExitWith_report_format is called when production with_report_format is exited.
func (s *BaseGoogleSQLParserListener) ExitWith_report_format(ctx *With_report_formatContext) {}

// EnterOptions_list is called when production options_list is entered.
func (s *BaseGoogleSQLParserListener) EnterOptions_list(ctx *Options_listContext) {}

// ExitOptions_list is called when production options_list is exited.
func (s *BaseGoogleSQLParserListener) ExitOptions_list(ctx *Options_listContext) {}

// EnterOptions_list_prefix is called when production options_list_prefix is entered.
func (s *BaseGoogleSQLParserListener) EnterOptions_list_prefix(ctx *Options_list_prefixContext) {}

// ExitOptions_list_prefix is called when production options_list_prefix is exited.
func (s *BaseGoogleSQLParserListener) ExitOptions_list_prefix(ctx *Options_list_prefixContext) {}

// EnterOptions_entry is called when production options_entry is entered.
func (s *BaseGoogleSQLParserListener) EnterOptions_entry(ctx *Options_entryContext) {}

// ExitOptions_entry is called when production options_entry is exited.
func (s *BaseGoogleSQLParserListener) ExitOptions_entry(ctx *Options_entryContext) {}

// EnterExpression_or_proto is called when production expression_or_proto is entered.
func (s *BaseGoogleSQLParserListener) EnterExpression_or_proto(ctx *Expression_or_protoContext) {}

// ExitExpression_or_proto is called when production expression_or_proto is exited.
func (s *BaseGoogleSQLParserListener) ExitExpression_or_proto(ctx *Expression_or_protoContext) {}

// EnterOptions_assignment_operator is called when production options_assignment_operator is entered.
func (s *BaseGoogleSQLParserListener) EnterOptions_assignment_operator(ctx *Options_assignment_operatorContext) {
}

// ExitOptions_assignment_operator is called when production options_assignment_operator is exited.
func (s *BaseGoogleSQLParserListener) ExitOptions_assignment_operator(ctx *Options_assignment_operatorContext) {
}

// EnterOpt_null_handling_modifier is called when production opt_null_handling_modifier is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_null_handling_modifier(ctx *Opt_null_handling_modifierContext) {
}

// ExitOpt_null_handling_modifier is called when production opt_null_handling_modifier is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_null_handling_modifier(ctx *Opt_null_handling_modifierContext) {
}

// EnterFunction_call_argument is called when production function_call_argument is entered.
func (s *BaseGoogleSQLParserListener) EnterFunction_call_argument(ctx *Function_call_argumentContext) {
}

// ExitFunction_call_argument is called when production function_call_argument is exited.
func (s *BaseGoogleSQLParserListener) ExitFunction_call_argument(ctx *Function_call_argumentContext) {
}

// EnterSequence_arg is called when production sequence_arg is entered.
func (s *BaseGoogleSQLParserListener) EnterSequence_arg(ctx *Sequence_argContext) {}

// ExitSequence_arg is called when production sequence_arg is exited.
func (s *BaseGoogleSQLParserListener) ExitSequence_arg(ctx *Sequence_argContext) {}

// EnterNamed_argument is called when production named_argument is entered.
func (s *BaseGoogleSQLParserListener) EnterNamed_argument(ctx *Named_argumentContext) {}

// ExitNamed_argument is called when production named_argument is exited.
func (s *BaseGoogleSQLParserListener) ExitNamed_argument(ctx *Named_argumentContext) {}

// EnterLambda_argument is called when production lambda_argument is entered.
func (s *BaseGoogleSQLParserListener) EnterLambda_argument(ctx *Lambda_argumentContext) {}

// ExitLambda_argument is called when production lambda_argument is exited.
func (s *BaseGoogleSQLParserListener) ExitLambda_argument(ctx *Lambda_argumentContext) {}

// EnterLambda_argument_list is called when production lambda_argument_list is entered.
func (s *BaseGoogleSQLParserListener) EnterLambda_argument_list(ctx *Lambda_argument_listContext) {}

// ExitLambda_argument_list is called when production lambda_argument_list is exited.
func (s *BaseGoogleSQLParserListener) ExitLambda_argument_list(ctx *Lambda_argument_listContext) {}

// EnterLimit_offset_clause is called when production limit_offset_clause is entered.
func (s *BaseGoogleSQLParserListener) EnterLimit_offset_clause(ctx *Limit_offset_clauseContext) {}

// ExitLimit_offset_clause is called when production limit_offset_clause is exited.
func (s *BaseGoogleSQLParserListener) ExitLimit_offset_clause(ctx *Limit_offset_clauseContext) {}

// EnterOpt_having_or_group_by_modifier is called when production opt_having_or_group_by_modifier is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_having_or_group_by_modifier(ctx *Opt_having_or_group_by_modifierContext) {
}

// ExitOpt_having_or_group_by_modifier is called when production opt_having_or_group_by_modifier is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_having_or_group_by_modifier(ctx *Opt_having_or_group_by_modifierContext) {
}

// EnterGroup_by_clause_prefix is called when production group_by_clause_prefix is entered.
func (s *BaseGoogleSQLParserListener) EnterGroup_by_clause_prefix(ctx *Group_by_clause_prefixContext) {
}

// ExitGroup_by_clause_prefix is called when production group_by_clause_prefix is exited.
func (s *BaseGoogleSQLParserListener) ExitGroup_by_clause_prefix(ctx *Group_by_clause_prefixContext) {
}

// EnterGroup_by_preamble is called when production group_by_preamble is entered.
func (s *BaseGoogleSQLParserListener) EnterGroup_by_preamble(ctx *Group_by_preambleContext) {}

// ExitGroup_by_preamble is called when production group_by_preamble is exited.
func (s *BaseGoogleSQLParserListener) ExitGroup_by_preamble(ctx *Group_by_preambleContext) {}

// EnterOpt_and_order is called when production opt_and_order is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_and_order(ctx *Opt_and_orderContext) {}

// ExitOpt_and_order is called when production opt_and_order is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_and_order(ctx *Opt_and_orderContext) {}

// EnterHint is called when production hint is entered.
func (s *BaseGoogleSQLParserListener) EnterHint(ctx *HintContext) {}

// ExitHint is called when production hint is exited.
func (s *BaseGoogleSQLParserListener) ExitHint(ctx *HintContext) {}

// EnterHint_with_body is called when production hint_with_body is entered.
func (s *BaseGoogleSQLParserListener) EnterHint_with_body(ctx *Hint_with_bodyContext) {}

// ExitHint_with_body is called when production hint_with_body is exited.
func (s *BaseGoogleSQLParserListener) ExitHint_with_body(ctx *Hint_with_bodyContext) {}

// EnterHint_with_body_prefix is called when production hint_with_body_prefix is entered.
func (s *BaseGoogleSQLParserListener) EnterHint_with_body_prefix(ctx *Hint_with_body_prefixContext) {}

// ExitHint_with_body_prefix is called when production hint_with_body_prefix is exited.
func (s *BaseGoogleSQLParserListener) ExitHint_with_body_prefix(ctx *Hint_with_body_prefixContext) {}

// EnterHint_entry is called when production hint_entry is entered.
func (s *BaseGoogleSQLParserListener) EnterHint_entry(ctx *Hint_entryContext) {}

// ExitHint_entry is called when production hint_entry is exited.
func (s *BaseGoogleSQLParserListener) ExitHint_entry(ctx *Hint_entryContext) {}

// EnterIdentifier_in_hints is called when production identifier_in_hints is entered.
func (s *BaseGoogleSQLParserListener) EnterIdentifier_in_hints(ctx *Identifier_in_hintsContext) {}

// ExitIdentifier_in_hints is called when production identifier_in_hints is exited.
func (s *BaseGoogleSQLParserListener) ExitIdentifier_in_hints(ctx *Identifier_in_hintsContext) {}

// EnterExtra_identifier_in_hints_name is called when production extra_identifier_in_hints_name is entered.
func (s *BaseGoogleSQLParserListener) EnterExtra_identifier_in_hints_name(ctx *Extra_identifier_in_hints_nameContext) {
}

// ExitExtra_identifier_in_hints_name is called when production extra_identifier_in_hints_name is exited.
func (s *BaseGoogleSQLParserListener) ExitExtra_identifier_in_hints_name(ctx *Extra_identifier_in_hints_nameContext) {
}

// EnterGrouping_item is called when production grouping_item is entered.
func (s *BaseGoogleSQLParserListener) EnterGrouping_item(ctx *Grouping_itemContext) {}

// ExitGrouping_item is called when production grouping_item is exited.
func (s *BaseGoogleSQLParserListener) ExitGrouping_item(ctx *Grouping_itemContext) {}

// EnterGrouping_set_list is called when production grouping_set_list is entered.
func (s *BaseGoogleSQLParserListener) EnterGrouping_set_list(ctx *Grouping_set_listContext) {}

// ExitGrouping_set_list is called when production grouping_set_list is exited.
func (s *BaseGoogleSQLParserListener) ExitGrouping_set_list(ctx *Grouping_set_listContext) {}

// EnterGrouping_set is called when production grouping_set is entered.
func (s *BaseGoogleSQLParserListener) EnterGrouping_set(ctx *Grouping_setContext) {}

// ExitGrouping_set is called when production grouping_set is exited.
func (s *BaseGoogleSQLParserListener) ExitGrouping_set(ctx *Grouping_setContext) {}

// EnterCube_list is called when production cube_list is entered.
func (s *BaseGoogleSQLParserListener) EnterCube_list(ctx *Cube_listContext) {}

// ExitCube_list is called when production cube_list is exited.
func (s *BaseGoogleSQLParserListener) ExitCube_list(ctx *Cube_listContext) {}

// EnterRollup_list is called when production rollup_list is entered.
func (s *BaseGoogleSQLParserListener) EnterRollup_list(ctx *Rollup_listContext) {}

// ExitRollup_list is called when production rollup_list is exited.
func (s *BaseGoogleSQLParserListener) ExitRollup_list(ctx *Rollup_listContext) {}

// EnterOpt_as_alias_with_required_as is called when production opt_as_alias_with_required_as is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_as_alias_with_required_as(ctx *Opt_as_alias_with_required_asContext) {
}

// ExitOpt_as_alias_with_required_as is called when production opt_as_alias_with_required_as is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_as_alias_with_required_as(ctx *Opt_as_alias_with_required_asContext) {
}

// EnterOpt_grouping_item_order is called when production opt_grouping_item_order is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_grouping_item_order(ctx *Opt_grouping_item_orderContext) {
}

// ExitOpt_grouping_item_order is called when production opt_grouping_item_order is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_grouping_item_order(ctx *Opt_grouping_item_orderContext) {
}

// EnterOpt_selection_item_order is called when production opt_selection_item_order is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_selection_item_order(ctx *Opt_selection_item_orderContext) {
}

// ExitOpt_selection_item_order is called when production opt_selection_item_order is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_selection_item_order(ctx *Opt_selection_item_orderContext) {
}

// EnterAsc_or_desc is called when production asc_or_desc is entered.
func (s *BaseGoogleSQLParserListener) EnterAsc_or_desc(ctx *Asc_or_descContext) {}

// ExitAsc_or_desc is called when production asc_or_desc is exited.
func (s *BaseGoogleSQLParserListener) ExitAsc_or_desc(ctx *Asc_or_descContext) {}

// EnterNull_order is called when production null_order is entered.
func (s *BaseGoogleSQLParserListener) EnterNull_order(ctx *Null_orderContext) {}

// ExitNull_order is called when production null_order is exited.
func (s *BaseGoogleSQLParserListener) ExitNull_order(ctx *Null_orderContext) {}

// EnterFunction_name_from_keyword is called when production function_name_from_keyword is entered.
func (s *BaseGoogleSQLParserListener) EnterFunction_name_from_keyword(ctx *Function_name_from_keywordContext) {
}

// ExitFunction_name_from_keyword is called when production function_name_from_keyword is exited.
func (s *BaseGoogleSQLParserListener) ExitFunction_name_from_keyword(ctx *Function_name_from_keywordContext) {
}

// EnterReplace_fields_expression is called when production replace_fields_expression is entered.
func (s *BaseGoogleSQLParserListener) EnterReplace_fields_expression(ctx *Replace_fields_expressionContext) {
}

// ExitReplace_fields_expression is called when production replace_fields_expression is exited.
func (s *BaseGoogleSQLParserListener) ExitReplace_fields_expression(ctx *Replace_fields_expressionContext) {
}

// EnterReplace_fields_prefix is called when production replace_fields_prefix is entered.
func (s *BaseGoogleSQLParserListener) EnterReplace_fields_prefix(ctx *Replace_fields_prefixContext) {}

// ExitReplace_fields_prefix is called when production replace_fields_prefix is exited.
func (s *BaseGoogleSQLParserListener) ExitReplace_fields_prefix(ctx *Replace_fields_prefixContext) {}

// EnterReplace_fields_arg is called when production replace_fields_arg is entered.
func (s *BaseGoogleSQLParserListener) EnterReplace_fields_arg(ctx *Replace_fields_argContext) {}

// ExitReplace_fields_arg is called when production replace_fields_arg is exited.
func (s *BaseGoogleSQLParserListener) ExitReplace_fields_arg(ctx *Replace_fields_argContext) {}

// EnterGeneralized_path_expression is called when production generalized_path_expression is entered.
func (s *BaseGoogleSQLParserListener) EnterGeneralized_path_expression(ctx *Generalized_path_expressionContext) {
}

// ExitGeneralized_path_expression is called when production generalized_path_expression is exited.
func (s *BaseGoogleSQLParserListener) ExitGeneralized_path_expression(ctx *Generalized_path_expressionContext) {
}

// EnterGeneralized_extension_path is called when production generalized_extension_path is entered.
func (s *BaseGoogleSQLParserListener) EnterGeneralized_extension_path(ctx *Generalized_extension_pathContext) {
}

// ExitGeneralized_extension_path is called when production generalized_extension_path is exited.
func (s *BaseGoogleSQLParserListener) ExitGeneralized_extension_path(ctx *Generalized_extension_pathContext) {
}

// EnterWith_expression is called when production with_expression is entered.
func (s *BaseGoogleSQLParserListener) EnterWith_expression(ctx *With_expressionContext) {}

// ExitWith_expression is called when production with_expression is exited.
func (s *BaseGoogleSQLParserListener) ExitWith_expression(ctx *With_expressionContext) {}

// EnterWith_expression_variable_prefix is called when production with_expression_variable_prefix is entered.
func (s *BaseGoogleSQLParserListener) EnterWith_expression_variable_prefix(ctx *With_expression_variable_prefixContext) {
}

// ExitWith_expression_variable_prefix is called when production with_expression_variable_prefix is exited.
func (s *BaseGoogleSQLParserListener) ExitWith_expression_variable_prefix(ctx *With_expression_variable_prefixContext) {
}

// EnterWith_expression_variable is called when production with_expression_variable is entered.
func (s *BaseGoogleSQLParserListener) EnterWith_expression_variable(ctx *With_expression_variableContext) {
}

// ExitWith_expression_variable is called when production with_expression_variable is exited.
func (s *BaseGoogleSQLParserListener) ExitWith_expression_variable(ctx *With_expression_variableContext) {
}

// EnterExtract_expression is called when production extract_expression is entered.
func (s *BaseGoogleSQLParserListener) EnterExtract_expression(ctx *Extract_expressionContext) {}

// ExitExtract_expression is called when production extract_expression is exited.
func (s *BaseGoogleSQLParserListener) ExitExtract_expression(ctx *Extract_expressionContext) {}

// EnterExtract_expression_base is called when production extract_expression_base is entered.
func (s *BaseGoogleSQLParserListener) EnterExtract_expression_base(ctx *Extract_expression_baseContext) {
}

// ExitExtract_expression_base is called when production extract_expression_base is exited.
func (s *BaseGoogleSQLParserListener) ExitExtract_expression_base(ctx *Extract_expression_baseContext) {
}

// EnterOpt_format is called when production opt_format is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_format(ctx *Opt_formatContext) {}

// ExitOpt_format is called when production opt_format is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_format(ctx *Opt_formatContext) {}

// EnterOpt_at_time_zone is called when production opt_at_time_zone is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_at_time_zone(ctx *Opt_at_time_zoneContext) {}

// ExitOpt_at_time_zone is called when production opt_at_time_zone is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_at_time_zone(ctx *Opt_at_time_zoneContext) {}

// EnterCast_expression is called when production cast_expression is entered.
func (s *BaseGoogleSQLParserListener) EnterCast_expression(ctx *Cast_expressionContext) {}

// ExitCast_expression is called when production cast_expression is exited.
func (s *BaseGoogleSQLParserListener) ExitCast_expression(ctx *Cast_expressionContext) {}

// EnterCase_expression is called when production case_expression is entered.
func (s *BaseGoogleSQLParserListener) EnterCase_expression(ctx *Case_expressionContext) {}

// ExitCase_expression is called when production case_expression is exited.
func (s *BaseGoogleSQLParserListener) ExitCase_expression(ctx *Case_expressionContext) {}

// EnterCase_expression_prefix is called when production case_expression_prefix is entered.
func (s *BaseGoogleSQLParserListener) EnterCase_expression_prefix(ctx *Case_expression_prefixContext) {
}

// ExitCase_expression_prefix is called when production case_expression_prefix is exited.
func (s *BaseGoogleSQLParserListener) ExitCase_expression_prefix(ctx *Case_expression_prefixContext) {
}

// EnterCase_value_expression_prefix is called when production case_value_expression_prefix is entered.
func (s *BaseGoogleSQLParserListener) EnterCase_value_expression_prefix(ctx *Case_value_expression_prefixContext) {
}

// ExitCase_value_expression_prefix is called when production case_value_expression_prefix is exited.
func (s *BaseGoogleSQLParserListener) ExitCase_value_expression_prefix(ctx *Case_value_expression_prefixContext) {
}

// EnterCase_no_value_expression_prefix is called when production case_no_value_expression_prefix is entered.
func (s *BaseGoogleSQLParserListener) EnterCase_no_value_expression_prefix(ctx *Case_no_value_expression_prefixContext) {
}

// ExitCase_no_value_expression_prefix is called when production case_no_value_expression_prefix is exited.
func (s *BaseGoogleSQLParserListener) ExitCase_no_value_expression_prefix(ctx *Case_no_value_expression_prefixContext) {
}

// EnterStruct_braced_constructor is called when production struct_braced_constructor is entered.
func (s *BaseGoogleSQLParserListener) EnterStruct_braced_constructor(ctx *Struct_braced_constructorContext) {
}

// ExitStruct_braced_constructor is called when production struct_braced_constructor is exited.
func (s *BaseGoogleSQLParserListener) ExitStruct_braced_constructor(ctx *Struct_braced_constructorContext) {
}

// EnterBraced_new_constructor is called when production braced_new_constructor is entered.
func (s *BaseGoogleSQLParserListener) EnterBraced_new_constructor(ctx *Braced_new_constructorContext) {
}

// ExitBraced_new_constructor is called when production braced_new_constructor is exited.
func (s *BaseGoogleSQLParserListener) ExitBraced_new_constructor(ctx *Braced_new_constructorContext) {
}

// EnterBraced_constructor is called when production braced_constructor is entered.
func (s *BaseGoogleSQLParserListener) EnterBraced_constructor(ctx *Braced_constructorContext) {}

// ExitBraced_constructor is called when production braced_constructor is exited.
func (s *BaseGoogleSQLParserListener) ExitBraced_constructor(ctx *Braced_constructorContext) {}

// EnterBraced_constructor_start is called when production braced_constructor_start is entered.
func (s *BaseGoogleSQLParserListener) EnterBraced_constructor_start(ctx *Braced_constructor_startContext) {
}

// ExitBraced_constructor_start is called when production braced_constructor_start is exited.
func (s *BaseGoogleSQLParserListener) ExitBraced_constructor_start(ctx *Braced_constructor_startContext) {
}

// EnterBraced_constructor_prefix is called when production braced_constructor_prefix is entered.
func (s *BaseGoogleSQLParserListener) EnterBraced_constructor_prefix(ctx *Braced_constructor_prefixContext) {
}

// ExitBraced_constructor_prefix is called when production braced_constructor_prefix is exited.
func (s *BaseGoogleSQLParserListener) ExitBraced_constructor_prefix(ctx *Braced_constructor_prefixContext) {
}

// EnterBraced_constructor_field is called when production braced_constructor_field is entered.
func (s *BaseGoogleSQLParserListener) EnterBraced_constructor_field(ctx *Braced_constructor_fieldContext) {
}

// ExitBraced_constructor_field is called when production braced_constructor_field is exited.
func (s *BaseGoogleSQLParserListener) ExitBraced_constructor_field(ctx *Braced_constructor_fieldContext) {
}

// EnterBraced_constructor_lhs is called when production braced_constructor_lhs is entered.
func (s *BaseGoogleSQLParserListener) EnterBraced_constructor_lhs(ctx *Braced_constructor_lhsContext) {
}

// ExitBraced_constructor_lhs is called when production braced_constructor_lhs is exited.
func (s *BaseGoogleSQLParserListener) ExitBraced_constructor_lhs(ctx *Braced_constructor_lhsContext) {
}

// EnterBraced_constructor_field_value is called when production braced_constructor_field_value is entered.
func (s *BaseGoogleSQLParserListener) EnterBraced_constructor_field_value(ctx *Braced_constructor_field_valueContext) {
}

// ExitBraced_constructor_field_value is called when production braced_constructor_field_value is exited.
func (s *BaseGoogleSQLParserListener) ExitBraced_constructor_field_value(ctx *Braced_constructor_field_valueContext) {
}

// EnterBraced_constructor_extension is called when production braced_constructor_extension is entered.
func (s *BaseGoogleSQLParserListener) EnterBraced_constructor_extension(ctx *Braced_constructor_extensionContext) {
}

// ExitBraced_constructor_extension is called when production braced_constructor_extension is exited.
func (s *BaseGoogleSQLParserListener) ExitBraced_constructor_extension(ctx *Braced_constructor_extensionContext) {
}

// EnterNew_constructor is called when production new_constructor is entered.
func (s *BaseGoogleSQLParserListener) EnterNew_constructor(ctx *New_constructorContext) {}

// ExitNew_constructor is called when production new_constructor is exited.
func (s *BaseGoogleSQLParserListener) ExitNew_constructor(ctx *New_constructorContext) {}

// EnterNew_constructor_prefix is called when production new_constructor_prefix is entered.
func (s *BaseGoogleSQLParserListener) EnterNew_constructor_prefix(ctx *New_constructor_prefixContext) {
}

// ExitNew_constructor_prefix is called when production new_constructor_prefix is exited.
func (s *BaseGoogleSQLParserListener) ExitNew_constructor_prefix(ctx *New_constructor_prefixContext) {
}

// EnterNew_constructor_prefix_no_arg is called when production new_constructor_prefix_no_arg is entered.
func (s *BaseGoogleSQLParserListener) EnterNew_constructor_prefix_no_arg(ctx *New_constructor_prefix_no_argContext) {
}

// ExitNew_constructor_prefix_no_arg is called when production new_constructor_prefix_no_arg is exited.
func (s *BaseGoogleSQLParserListener) ExitNew_constructor_prefix_no_arg(ctx *New_constructor_prefix_no_argContext) {
}

// EnterNew_constructor_arg is called when production new_constructor_arg is entered.
func (s *BaseGoogleSQLParserListener) EnterNew_constructor_arg(ctx *New_constructor_argContext) {}

// ExitNew_constructor_arg is called when production new_constructor_arg is exited.
func (s *BaseGoogleSQLParserListener) ExitNew_constructor_arg(ctx *New_constructor_argContext) {}

// EnterArray_constructor is called when production array_constructor is entered.
func (s *BaseGoogleSQLParserListener) EnterArray_constructor(ctx *Array_constructorContext) {}

// ExitArray_constructor is called when production array_constructor is exited.
func (s *BaseGoogleSQLParserListener) ExitArray_constructor(ctx *Array_constructorContext) {}

// EnterArray_constructor_prefix is called when production array_constructor_prefix is entered.
func (s *BaseGoogleSQLParserListener) EnterArray_constructor_prefix(ctx *Array_constructor_prefixContext) {
}

// ExitArray_constructor_prefix is called when production array_constructor_prefix is exited.
func (s *BaseGoogleSQLParserListener) ExitArray_constructor_prefix(ctx *Array_constructor_prefixContext) {
}

// EnterArray_constructor_prefix_no_expressions is called when production array_constructor_prefix_no_expressions is entered.
func (s *BaseGoogleSQLParserListener) EnterArray_constructor_prefix_no_expressions(ctx *Array_constructor_prefix_no_expressionsContext) {
}

// ExitArray_constructor_prefix_no_expressions is called when production array_constructor_prefix_no_expressions is exited.
func (s *BaseGoogleSQLParserListener) ExitArray_constructor_prefix_no_expressions(ctx *Array_constructor_prefix_no_expressionsContext) {
}

// EnterRange_literal is called when production range_literal is entered.
func (s *BaseGoogleSQLParserListener) EnterRange_literal(ctx *Range_literalContext) {}

// ExitRange_literal is called when production range_literal is exited.
func (s *BaseGoogleSQLParserListener) ExitRange_literal(ctx *Range_literalContext) {}

// EnterRange_type is called when production range_type is entered.
func (s *BaseGoogleSQLParserListener) EnterRange_type(ctx *Range_typeContext) {}

// ExitRange_type is called when production range_type is exited.
func (s *BaseGoogleSQLParserListener) ExitRange_type(ctx *Range_typeContext) {}

// EnterType is called when production type is entered.
func (s *BaseGoogleSQLParserListener) EnterType(ctx *TypeContext) {}

// ExitType is called when production type is exited.
func (s *BaseGoogleSQLParserListener) ExitType(ctx *TypeContext) {}

// EnterCollate_clause is called when production collate_clause is entered.
func (s *BaseGoogleSQLParserListener) EnterCollate_clause(ctx *Collate_clauseContext) {}

// ExitCollate_clause is called when production collate_clause is exited.
func (s *BaseGoogleSQLParserListener) ExitCollate_clause(ctx *Collate_clauseContext) {}

// EnterString_literal_or_parameter is called when production string_literal_or_parameter is entered.
func (s *BaseGoogleSQLParserListener) EnterString_literal_or_parameter(ctx *String_literal_or_parameterContext) {
}

// ExitString_literal_or_parameter is called when production string_literal_or_parameter is exited.
func (s *BaseGoogleSQLParserListener) ExitString_literal_or_parameter(ctx *String_literal_or_parameterContext) {
}

// EnterSystem_variable_expression is called when production system_variable_expression is entered.
func (s *BaseGoogleSQLParserListener) EnterSystem_variable_expression(ctx *System_variable_expressionContext) {
}

// ExitSystem_variable_expression is called when production system_variable_expression is exited.
func (s *BaseGoogleSQLParserListener) ExitSystem_variable_expression(ctx *System_variable_expressionContext) {
}

// EnterParameter_expression is called when production parameter_expression is entered.
func (s *BaseGoogleSQLParserListener) EnterParameter_expression(ctx *Parameter_expressionContext) {}

// ExitParameter_expression is called when production parameter_expression is exited.
func (s *BaseGoogleSQLParserListener) ExitParameter_expression(ctx *Parameter_expressionContext) {}

// EnterNamed_parameter_expression is called when production named_parameter_expression is entered.
func (s *BaseGoogleSQLParserListener) EnterNamed_parameter_expression(ctx *Named_parameter_expressionContext) {
}

// ExitNamed_parameter_expression is called when production named_parameter_expression is exited.
func (s *BaseGoogleSQLParserListener) ExitNamed_parameter_expression(ctx *Named_parameter_expressionContext) {
}

// EnterOpt_type_parameters is called when production opt_type_parameters is entered.
func (s *BaseGoogleSQLParserListener) EnterOpt_type_parameters(ctx *Opt_type_parametersContext) {}

// ExitOpt_type_parameters is called when production opt_type_parameters is exited.
func (s *BaseGoogleSQLParserListener) ExitOpt_type_parameters(ctx *Opt_type_parametersContext) {}

// EnterType_parameters_prefix is called when production type_parameters_prefix is entered.
func (s *BaseGoogleSQLParserListener) EnterType_parameters_prefix(ctx *Type_parameters_prefixContext) {
}

// ExitType_parameters_prefix is called when production type_parameters_prefix is exited.
func (s *BaseGoogleSQLParserListener) ExitType_parameters_prefix(ctx *Type_parameters_prefixContext) {
}

// EnterType_parameter is called when production type_parameter is entered.
func (s *BaseGoogleSQLParserListener) EnterType_parameter(ctx *Type_parameterContext) {}

// ExitType_parameter is called when production type_parameter is exited.
func (s *BaseGoogleSQLParserListener) ExitType_parameter(ctx *Type_parameterContext) {}

// EnterRaw_type is called when production raw_type is entered.
func (s *BaseGoogleSQLParserListener) EnterRaw_type(ctx *Raw_typeContext) {}

// ExitRaw_type is called when production raw_type is exited.
func (s *BaseGoogleSQLParserListener) ExitRaw_type(ctx *Raw_typeContext) {}

// EnterMap_type is called when production map_type is entered.
func (s *BaseGoogleSQLParserListener) EnterMap_type(ctx *Map_typeContext) {}

// ExitMap_type is called when production map_type is exited.
func (s *BaseGoogleSQLParserListener) ExitMap_type(ctx *Map_typeContext) {}

// EnterFunction_type is called when production function_type is entered.
func (s *BaseGoogleSQLParserListener) EnterFunction_type(ctx *Function_typeContext) {}

// ExitFunction_type is called when production function_type is exited.
func (s *BaseGoogleSQLParserListener) ExitFunction_type(ctx *Function_typeContext) {}

// EnterFunction_type_prefix is called when production function_type_prefix is entered.
func (s *BaseGoogleSQLParserListener) EnterFunction_type_prefix(ctx *Function_type_prefixContext) {}

// ExitFunction_type_prefix is called when production function_type_prefix is exited.
func (s *BaseGoogleSQLParserListener) ExitFunction_type_prefix(ctx *Function_type_prefixContext) {}

// EnterType_name is called when production type_name is entered.
func (s *BaseGoogleSQLParserListener) EnterType_name(ctx *Type_nameContext) {}

// ExitType_name is called when production type_name is exited.
func (s *BaseGoogleSQLParserListener) ExitType_name(ctx *Type_nameContext) {}

// EnterPath_expression is called when production path_expression is entered.
func (s *BaseGoogleSQLParserListener) EnterPath_expression(ctx *Path_expressionContext) {}

// ExitPath_expression is called when production path_expression is exited.
func (s *BaseGoogleSQLParserListener) ExitPath_expression(ctx *Path_expressionContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseGoogleSQLParserListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseGoogleSQLParserListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterKeyword_as_identifier is called when production keyword_as_identifier is entered.
func (s *BaseGoogleSQLParserListener) EnterKeyword_as_identifier(ctx *Keyword_as_identifierContext) {}

// ExitKeyword_as_identifier is called when production keyword_as_identifier is exited.
func (s *BaseGoogleSQLParserListener) ExitKeyword_as_identifier(ctx *Keyword_as_identifierContext) {}

// EnterCommon_keyword_as_identifier is called when production common_keyword_as_identifier is entered.
func (s *BaseGoogleSQLParserListener) EnterCommon_keyword_as_identifier(ctx *Common_keyword_as_identifierContext) {
}

// ExitCommon_keyword_as_identifier is called when production common_keyword_as_identifier is exited.
func (s *BaseGoogleSQLParserListener) ExitCommon_keyword_as_identifier(ctx *Common_keyword_as_identifierContext) {
}

// EnterToken_identifier is called when production token_identifier is entered.
func (s *BaseGoogleSQLParserListener) EnterToken_identifier(ctx *Token_identifierContext) {}

// ExitToken_identifier is called when production token_identifier is exited.
func (s *BaseGoogleSQLParserListener) ExitToken_identifier(ctx *Token_identifierContext) {}

// EnterStruct_type is called when production struct_type is entered.
func (s *BaseGoogleSQLParserListener) EnterStruct_type(ctx *Struct_typeContext) {}

// ExitStruct_type is called when production struct_type is exited.
func (s *BaseGoogleSQLParserListener) ExitStruct_type(ctx *Struct_typeContext) {}

// EnterStruct_type_prefix is called when production struct_type_prefix is entered.
func (s *BaseGoogleSQLParserListener) EnterStruct_type_prefix(ctx *Struct_type_prefixContext) {}

// ExitStruct_type_prefix is called when production struct_type_prefix is exited.
func (s *BaseGoogleSQLParserListener) ExitStruct_type_prefix(ctx *Struct_type_prefixContext) {}

// EnterStruct_field is called when production struct_field is entered.
func (s *BaseGoogleSQLParserListener) EnterStruct_field(ctx *Struct_fieldContext) {}

// ExitStruct_field is called when production struct_field is exited.
func (s *BaseGoogleSQLParserListener) ExitStruct_field(ctx *Struct_fieldContext) {}

// EnterArray_type is called when production array_type is entered.
func (s *BaseGoogleSQLParserListener) EnterArray_type(ctx *Array_typeContext) {}

// ExitArray_type is called when production array_type is exited.
func (s *BaseGoogleSQLParserListener) ExitArray_type(ctx *Array_typeContext) {}

// EnterTemplate_type_open is called when production template_type_open is entered.
func (s *BaseGoogleSQLParserListener) EnterTemplate_type_open(ctx *Template_type_openContext) {}

// ExitTemplate_type_open is called when production template_type_open is exited.
func (s *BaseGoogleSQLParserListener) ExitTemplate_type_open(ctx *Template_type_openContext) {}

// EnterTemplate_type_close is called when production template_type_close is entered.
func (s *BaseGoogleSQLParserListener) EnterTemplate_type_close(ctx *Template_type_closeContext) {}

// ExitTemplate_type_close is called when production template_type_close is exited.
func (s *BaseGoogleSQLParserListener) ExitTemplate_type_close(ctx *Template_type_closeContext) {}

// EnterDate_or_time_literal is called when production date_or_time_literal is entered.
func (s *BaseGoogleSQLParserListener) EnterDate_or_time_literal(ctx *Date_or_time_literalContext) {}

// ExitDate_or_time_literal is called when production date_or_time_literal is exited.
func (s *BaseGoogleSQLParserListener) ExitDate_or_time_literal(ctx *Date_or_time_literalContext) {}

// EnterDate_or_time_literal_kind is called when production date_or_time_literal_kind is entered.
func (s *BaseGoogleSQLParserListener) EnterDate_or_time_literal_kind(ctx *Date_or_time_literal_kindContext) {
}

// ExitDate_or_time_literal_kind is called when production date_or_time_literal_kind is exited.
func (s *BaseGoogleSQLParserListener) ExitDate_or_time_literal_kind(ctx *Date_or_time_literal_kindContext) {
}

// EnterFloating_point_literal is called when production floating_point_literal is entered.
func (s *BaseGoogleSQLParserListener) EnterFloating_point_literal(ctx *Floating_point_literalContext) {
}

// ExitFloating_point_literal is called when production floating_point_literal is exited.
func (s *BaseGoogleSQLParserListener) ExitFloating_point_literal(ctx *Floating_point_literalContext) {
}

// EnterJson_literal is called when production json_literal is entered.
func (s *BaseGoogleSQLParserListener) EnterJson_literal(ctx *Json_literalContext) {}

// ExitJson_literal is called when production json_literal is exited.
func (s *BaseGoogleSQLParserListener) ExitJson_literal(ctx *Json_literalContext) {}

// EnterBignumeric_literal is called when production bignumeric_literal is entered.
func (s *BaseGoogleSQLParserListener) EnterBignumeric_literal(ctx *Bignumeric_literalContext) {}

// ExitBignumeric_literal is called when production bignumeric_literal is exited.
func (s *BaseGoogleSQLParserListener) ExitBignumeric_literal(ctx *Bignumeric_literalContext) {}

// EnterBignumeric_literal_prefix is called when production bignumeric_literal_prefix is entered.
func (s *BaseGoogleSQLParserListener) EnterBignumeric_literal_prefix(ctx *Bignumeric_literal_prefixContext) {
}

// ExitBignumeric_literal_prefix is called when production bignumeric_literal_prefix is exited.
func (s *BaseGoogleSQLParserListener) ExitBignumeric_literal_prefix(ctx *Bignumeric_literal_prefixContext) {
}

// EnterNumeric_literal is called when production numeric_literal is entered.
func (s *BaseGoogleSQLParserListener) EnterNumeric_literal(ctx *Numeric_literalContext) {}

// ExitNumeric_literal is called when production numeric_literal is exited.
func (s *BaseGoogleSQLParserListener) ExitNumeric_literal(ctx *Numeric_literalContext) {}

// EnterNumeric_literal_prefix is called when production numeric_literal_prefix is entered.
func (s *BaseGoogleSQLParserListener) EnterNumeric_literal_prefix(ctx *Numeric_literal_prefixContext) {
}

// ExitNumeric_literal_prefix is called when production numeric_literal_prefix is exited.
func (s *BaseGoogleSQLParserListener) ExitNumeric_literal_prefix(ctx *Numeric_literal_prefixContext) {
}

// EnterInteger_literal is called when production integer_literal is entered.
func (s *BaseGoogleSQLParserListener) EnterInteger_literal(ctx *Integer_literalContext) {}

// ExitInteger_literal is called when production integer_literal is exited.
func (s *BaseGoogleSQLParserListener) ExitInteger_literal(ctx *Integer_literalContext) {}

// EnterBytes_literal is called when production bytes_literal is entered.
func (s *BaseGoogleSQLParserListener) EnterBytes_literal(ctx *Bytes_literalContext) {}

// ExitBytes_literal is called when production bytes_literal is exited.
func (s *BaseGoogleSQLParserListener) ExitBytes_literal(ctx *Bytes_literalContext) {}

// EnterNull_literal is called when production null_literal is entered.
func (s *BaseGoogleSQLParserListener) EnterNull_literal(ctx *Null_literalContext) {}

// ExitNull_literal is called when production null_literal is exited.
func (s *BaseGoogleSQLParserListener) ExitNull_literal(ctx *Null_literalContext) {}

// EnterBoolean_literal is called when production boolean_literal is entered.
func (s *BaseGoogleSQLParserListener) EnterBoolean_literal(ctx *Boolean_literalContext) {}

// ExitBoolean_literal is called when production boolean_literal is exited.
func (s *BaseGoogleSQLParserListener) ExitBoolean_literal(ctx *Boolean_literalContext) {}

// EnterString_literal is called when production string_literal is entered.
func (s *BaseGoogleSQLParserListener) EnterString_literal(ctx *String_literalContext) {}

// ExitString_literal is called when production string_literal is exited.
func (s *BaseGoogleSQLParserListener) ExitString_literal(ctx *String_literalContext) {}

// EnterString_literal_component is called when production string_literal_component is entered.
func (s *BaseGoogleSQLParserListener) EnterString_literal_component(ctx *String_literal_componentContext) {
}

// ExitString_literal_component is called when production string_literal_component is exited.
func (s *BaseGoogleSQLParserListener) ExitString_literal_component(ctx *String_literal_componentContext) {
}

// EnterBytes_literal_component is called when production bytes_literal_component is entered.
func (s *BaseGoogleSQLParserListener) EnterBytes_literal_component(ctx *Bytes_literal_componentContext) {
}

// ExitBytes_literal_component is called when production bytes_literal_component is exited.
func (s *BaseGoogleSQLParserListener) ExitBytes_literal_component(ctx *Bytes_literal_componentContext) {
}
