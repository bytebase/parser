// Code generated from GoogleSQLParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package googlesql // GoogleSQLParser
import "github.com/antlr4-go/antlr/v4"

// GoogleSQLParserListener is a complete listener for a parse tree produced by GoogleSQLParser.
type GoogleSQLParserListener interface {
	antlr.ParseTreeListener

	// EnterRoot is called when entering the root production.
	EnterRoot(c *RootContext)

	// EnterStmts is called when entering the stmts production.
	EnterStmts(c *StmtsContext)

	// EnterUnterminated_sql_statement is called when entering the unterminated_sql_statement production.
	EnterUnterminated_sql_statement(c *Unterminated_sql_statementContext)

	// EnterSql_statement_body is called when entering the sql_statement_body production.
	EnterSql_statement_body(c *Sql_statement_bodyContext)

	// EnterGql_statement is called when entering the gql_statement production.
	EnterGql_statement(c *Gql_statementContext)

	// EnterGraph_operation_block is called when entering the graph_operation_block production.
	EnterGraph_operation_block(c *Graph_operation_blockContext)

	// EnterGraph_composite_query_block is called when entering the graph_composite_query_block production.
	EnterGraph_composite_query_block(c *Graph_composite_query_blockContext)

	// EnterGraph_composite_query_prefix is called when entering the graph_composite_query_prefix production.
	EnterGraph_composite_query_prefix(c *Graph_composite_query_prefixContext)

	// EnterGraph_set_operation_metadata is called when entering the graph_set_operation_metadata production.
	EnterGraph_set_operation_metadata(c *Graph_set_operation_metadataContext)

	// EnterGraph_linear_query_operation is called when entering the graph_linear_query_operation production.
	EnterGraph_linear_query_operation(c *Graph_linear_query_operationContext)

	// EnterGraph_linear_operator_list is called when entering the graph_linear_operator_list production.
	EnterGraph_linear_operator_list(c *Graph_linear_operator_listContext)

	// EnterGraph_linear_operator is called when entering the graph_linear_operator production.
	EnterGraph_linear_operator(c *Graph_linear_operatorContext)

	// EnterGraph_sample_clause is called when entering the graph_sample_clause production.
	EnterGraph_sample_clause(c *Graph_sample_clauseContext)

	// EnterOpt_graph_sample_clause_suffix is called when entering the opt_graph_sample_clause_suffix production.
	EnterOpt_graph_sample_clause_suffix(c *Opt_graph_sample_clause_suffixContext)

	// EnterGraph_for_operator is called when entering the graph_for_operator production.
	EnterGraph_for_operator(c *Graph_for_operatorContext)

	// EnterOpt_with_offset_and_alias_with_required_as is called when entering the opt_with_offset_and_alias_with_required_as production.
	EnterOpt_with_offset_and_alias_with_required_as(c *Opt_with_offset_and_alias_with_required_asContext)

	// EnterGraph_with_operator is called when entering the graph_with_operator production.
	EnterGraph_with_operator(c *Graph_with_operatorContext)

	// EnterGraph_page_operator is called when entering the graph_page_operator production.
	EnterGraph_page_operator(c *Graph_page_operatorContext)

	// EnterGraph_order_by_operator is called when entering the graph_order_by_operator production.
	EnterGraph_order_by_operator(c *Graph_order_by_operatorContext)

	// EnterGraph_filter_operator is called when entering the graph_filter_operator production.
	EnterGraph_filter_operator(c *Graph_filter_operatorContext)

	// EnterGraph_let_operator is called when entering the graph_let_operator production.
	EnterGraph_let_operator(c *Graph_let_operatorContext)

	// EnterGraph_let_variable_definition_list is called when entering the graph_let_variable_definition_list production.
	EnterGraph_let_variable_definition_list(c *Graph_let_variable_definition_listContext)

	// EnterGraph_let_variable_definition is called when entering the graph_let_variable_definition production.
	EnterGraph_let_variable_definition(c *Graph_let_variable_definitionContext)

	// EnterGraph_optional_match_operator is called when entering the graph_optional_match_operator production.
	EnterGraph_optional_match_operator(c *Graph_optional_match_operatorContext)

	// EnterGraph_match_operator is called when entering the graph_match_operator production.
	EnterGraph_match_operator(c *Graph_match_operatorContext)

	// EnterGraph_pattern is called when entering the graph_pattern production.
	EnterGraph_pattern(c *Graph_patternContext)

	// EnterGraph_path_pattern_list is called when entering the graph_path_pattern_list production.
	EnterGraph_path_pattern_list(c *Graph_path_pattern_listContext)

	// EnterGraph_path_pattern is called when entering the graph_path_pattern production.
	EnterGraph_path_pattern(c *Graph_path_patternContext)

	// EnterGraph_path_pattern_expr is called when entering the graph_path_pattern_expr production.
	EnterGraph_path_pattern_expr(c *Graph_path_pattern_exprContext)

	// EnterGraph_path_factor is called when entering the graph_path_factor production.
	EnterGraph_path_factor(c *Graph_path_factorContext)

	// EnterGraph_quantified_path_primary is called when entering the graph_quantified_path_primary production.
	EnterGraph_quantified_path_primary(c *Graph_quantified_path_primaryContext)

	// EnterGraph_path_primary is called when entering the graph_path_primary production.
	EnterGraph_path_primary(c *Graph_path_primaryContext)

	// EnterGraph_parenthesized_path_pattern is called when entering the graph_parenthesized_path_pattern production.
	EnterGraph_parenthesized_path_pattern(c *Graph_parenthesized_path_patternContext)

	// EnterGraph_element_pattern is called when entering the graph_element_pattern production.
	EnterGraph_element_pattern(c *Graph_element_patternContext)

	// EnterGraph_edge_pattern is called when entering the graph_edge_pattern production.
	EnterGraph_edge_pattern(c *Graph_edge_patternContext)

	// EnterGraph_node_pattern is called when entering the graph_node_pattern production.
	EnterGraph_node_pattern(c *Graph_node_patternContext)

	// EnterGraph_element_pattern_filler is called when entering the graph_element_pattern_filler production.
	EnterGraph_element_pattern_filler(c *Graph_element_pattern_fillerContext)

	// EnterGraph_property_specification is called when entering the graph_property_specification production.
	EnterGraph_property_specification(c *Graph_property_specificationContext)

	// EnterGraph_property_name_and_value is called when entering the graph_property_name_and_value production.
	EnterGraph_property_name_and_value(c *Graph_property_name_and_valueContext)

	// EnterOpt_is_label_expression is called when entering the opt_is_label_expression production.
	EnterOpt_is_label_expression(c *Opt_is_label_expressionContext)

	// EnterLabel_expression is called when entering the label_expression production.
	EnterLabel_expression(c *Label_expressionContext)

	// EnterLabel_primary is called when entering the label_primary production.
	EnterLabel_primary(c *Label_primaryContext)

	// EnterParenthesized_label_expression is called when entering the parenthesized_label_expression production.
	EnterParenthesized_label_expression(c *Parenthesized_label_expressionContext)

	// EnterOpt_graph_element_identifier is called when entering the opt_graph_element_identifier production.
	EnterOpt_graph_element_identifier(c *Opt_graph_element_identifierContext)

	// EnterOpt_graph_path_mode_prefix is called when entering the opt_graph_path_mode_prefix production.
	EnterOpt_graph_path_mode_prefix(c *Opt_graph_path_mode_prefixContext)

	// EnterPath_or_paths is called when entering the path_or_paths production.
	EnterPath_or_paths(c *Path_or_pathsContext)

	// EnterOpt_graph_path_mode is called when entering the opt_graph_path_mode production.
	EnterOpt_graph_path_mode(c *Opt_graph_path_modeContext)

	// EnterOpt_graph_search_prefix is called when entering the opt_graph_search_prefix production.
	EnterOpt_graph_search_prefix(c *Opt_graph_search_prefixContext)

	// EnterOpt_path_variable_assignment is called when entering the opt_path_variable_assignment production.
	EnterOpt_path_variable_assignment(c *Opt_path_variable_assignmentContext)

	// EnterGraph_identifier is called when entering the graph_identifier production.
	EnterGraph_identifier(c *Graph_identifierContext)

	// EnterGraph_return_operator is called when entering the graph_return_operator production.
	EnterGraph_return_operator(c *Graph_return_operatorContext)

	// EnterGraph_page_clause is called when entering the graph_page_clause production.
	EnterGraph_page_clause(c *Graph_page_clauseContext)

	// EnterGraph_order_by_clause is called when entering the graph_order_by_clause production.
	EnterGraph_order_by_clause(c *Graph_order_by_clauseContext)

	// EnterGraph_ordering_expression is called when entering the graph_ordering_expression production.
	EnterGraph_ordering_expression(c *Graph_ordering_expressionContext)

	// EnterOpt_graph_asc_or_desc is called when entering the opt_graph_asc_or_desc production.
	EnterOpt_graph_asc_or_desc(c *Opt_graph_asc_or_descContext)

	// EnterGraph_return_item_list is called when entering the graph_return_item_list production.
	EnterGraph_return_item_list(c *Graph_return_item_listContext)

	// EnterGraph_return_item is called when entering the graph_return_item production.
	EnterGraph_return_item(c *Graph_return_itemContext)

	// EnterUndrop_statement is called when entering the undrop_statement production.
	EnterUndrop_statement(c *Undrop_statementContext)

	// EnterModule_statement is called when entering the module_statement production.
	EnterModule_statement(c *Module_statementContext)

	// EnterImport_statement is called when entering the import_statement production.
	EnterImport_statement(c *Import_statementContext)

	// EnterOpt_as_or_into_alias is called when entering the opt_as_or_into_alias production.
	EnterOpt_as_or_into_alias(c *Opt_as_or_into_aliasContext)

	// EnterPath_expression_or_string is called when entering the path_expression_or_string production.
	EnterPath_expression_or_string(c *Path_expression_or_stringContext)

	// EnterImport_type is called when entering the import_type production.
	EnterImport_type(c *Import_typeContext)

	// EnterCall_statement is called when entering the call_statement production.
	EnterCall_statement(c *Call_statementContext)

	// EnterDrop_statement is called when entering the drop_statement production.
	EnterDrop_statement(c *Drop_statementContext)

	// EnterOpt_drop_mode is called when entering the opt_drop_mode production.
	EnterOpt_drop_mode(c *Opt_drop_modeContext)

	// EnterDrop_all_row_access_policies_statement is called when entering the drop_all_row_access_policies_statement production.
	EnterDrop_all_row_access_policies_statement(c *Drop_all_row_access_policies_statementContext)

	// EnterShow_statement is called when entering the show_statement production.
	EnterShow_statement(c *Show_statementContext)

	// EnterOpt_like_string_literal is called when entering the opt_like_string_literal production.
	EnterOpt_like_string_literal(c *Opt_like_string_literalContext)

	// EnterShow_target is called when entering the show_target production.
	EnterShow_target(c *Show_targetContext)

	// EnterRename_statement is called when entering the rename_statement production.
	EnterRename_statement(c *Rename_statementContext)

	// EnterRevoke_statement is called when entering the revoke_statement production.
	EnterRevoke_statement(c *Revoke_statementContext)

	// EnterGrant_statement is called when entering the grant_statement production.
	EnterGrant_statement(c *Grant_statementContext)

	// EnterPrivileges is called when entering the privileges production.
	EnterPrivileges(c *PrivilegesContext)

	// EnterExport_metadata_statement is called when entering the export_metadata_statement production.
	EnterExport_metadata_statement(c *Export_metadata_statementContext)

	// EnterExport_model_statement is called when entering the export_model_statement production.
	EnterExport_model_statement(c *Export_model_statementContext)

	// EnterExport_data_statement is called when entering the export_data_statement production.
	EnterExport_data_statement(c *Export_data_statementContext)

	// EnterExport_data_no_query is called when entering the export_data_no_query production.
	EnterExport_data_no_query(c *Export_data_no_queryContext)

	// EnterExplain_statement is called when entering the explain_statement production.
	EnterExplain_statement(c *Explain_statementContext)

	// EnterExecute_immediate is called when entering the execute_immediate production.
	EnterExecute_immediate(c *Execute_immediateContext)

	// EnterOpt_execute_into_clause is called when entering the opt_execute_into_clause production.
	EnterOpt_execute_into_clause(c *Opt_execute_into_clauseContext)

	// EnterOpt_execute_using_clause is called when entering the opt_execute_using_clause production.
	EnterOpt_execute_using_clause(c *Opt_execute_using_clauseContext)

	// EnterExecute_using_argument_list is called when entering the execute_using_argument_list production.
	EnterExecute_using_argument_list(c *Execute_using_argument_listContext)

	// EnterExecute_using_argument is called when entering the execute_using_argument production.
	EnterExecute_using_argument(c *Execute_using_argumentContext)

	// EnterDescribe_statement is called when entering the describe_statement production.
	EnterDescribe_statement(c *Describe_statementContext)

	// EnterDescribe_info is called when entering the describe_info production.
	EnterDescribe_info(c *Describe_infoContext)

	// EnterOpt_from_path_expression is called when entering the opt_from_path_expression production.
	EnterOpt_from_path_expression(c *Opt_from_path_expressionContext)

	// EnterDescribe_keyword is called when entering the describe_keyword production.
	EnterDescribe_keyword(c *Describe_keywordContext)

	// EnterDefine_table_statement is called when entering the define_table_statement production.
	EnterDefine_table_statement(c *Define_table_statementContext)

	// EnterCreate_entity_statement is called when entering the create_entity_statement production.
	EnterCreate_entity_statement(c *Create_entity_statementContext)

	// EnterOpt_generic_entity_body is called when entering the opt_generic_entity_body production.
	EnterOpt_generic_entity_body(c *Opt_generic_entity_bodyContext)

	// EnterCreate_view_statement is called when entering the create_view_statement production.
	EnterCreate_view_statement(c *Create_view_statementContext)

	// EnterQuery_or_replica_source is called when entering the query_or_replica_source production.
	EnterQuery_or_replica_source(c *Query_or_replica_sourceContext)

	// EnterColumn_with_options_list is called when entering the column_with_options_list production.
	EnterColumn_with_options_list(c *Column_with_options_listContext)

	// EnterColumn_with_options is called when entering the column_with_options production.
	EnterColumn_with_options(c *Column_with_optionsContext)

	// EnterCreate_table_statement is called when entering the create_table_statement production.
	EnterCreate_table_statement(c *Create_table_statementContext)

	// EnterOpt_ttl_clause is called when entering the opt_ttl_clause production.
	EnterOpt_ttl_clause(c *Opt_ttl_clauseContext)

	// EnterOpt_copy_table is called when entering the opt_copy_table production.
	EnterOpt_copy_table(c *Opt_copy_tableContext)

	// EnterCopy_data_source is called when entering the copy_data_source production.
	EnterCopy_data_source(c *Copy_data_sourceContext)

	// EnterOpt_clone_table is called when entering the opt_clone_table production.
	EnterOpt_clone_table(c *Opt_clone_tableContext)

	// EnterOpt_spanner_table_options is called when entering the opt_spanner_table_options production.
	EnterOpt_spanner_table_options(c *Opt_spanner_table_optionsContext)

	// EnterOpt_spanner_interleave_in_parent_clause is called when entering the opt_spanner_interleave_in_parent_clause production.
	EnterOpt_spanner_interleave_in_parent_clause(c *Opt_spanner_interleave_in_parent_clauseContext)

	// EnterSpanner_primary_key is called when entering the spanner_primary_key production.
	EnterSpanner_primary_key(c *Spanner_primary_keyContext)

	// EnterCreate_table_function_statement is called when entering the create_table_function_statement production.
	EnterCreate_table_function_statement(c *Create_table_function_statementContext)

	// EnterOpt_as_query_or_string is called when entering the opt_as_query_or_string production.
	EnterOpt_as_query_or_string(c *Opt_as_query_or_stringContext)

	// EnterUnordered_language_options is called when entering the unordered_language_options production.
	EnterUnordered_language_options(c *Unordered_language_optionsContext)

	// EnterOpt_function_parameters is called when entering the opt_function_parameters production.
	EnterOpt_function_parameters(c *Opt_function_parametersContext)

	// EnterCreate_snapshot_statement is called when entering the create_snapshot_statement production.
	EnterCreate_snapshot_statement(c *Create_snapshot_statementContext)

	// EnterCreate_external_schema_statement is called when entering the create_external_schema_statement production.
	EnterCreate_external_schema_statement(c *Create_external_schema_statementContext)

	// EnterCreate_schema_statement is called when entering the create_schema_statement production.
	EnterCreate_schema_statement(c *Create_schema_statementContext)

	// EnterCreate_property_graph_statement is called when entering the create_property_graph_statement production.
	EnterCreate_property_graph_statement(c *Create_property_graph_statementContext)

	// EnterOpt_edge_table_clause is called when entering the opt_edge_table_clause production.
	EnterOpt_edge_table_clause(c *Opt_edge_table_clauseContext)

	// EnterElement_table_list is called when entering the element_table_list production.
	EnterElement_table_list(c *Element_table_listContext)

	// EnterElement_table_definition is called when entering the element_table_definition production.
	EnterElement_table_definition(c *Element_table_definitionContext)

	// EnterOpt_label_and_properties_clause is called when entering the opt_label_and_properties_clause production.
	EnterOpt_label_and_properties_clause(c *Opt_label_and_properties_clauseContext)

	// EnterLabel_and_properties_list is called when entering the label_and_properties_list production.
	EnterLabel_and_properties_list(c *Label_and_properties_listContext)

	// EnterLabel_and_properties is called when entering the label_and_properties production.
	EnterLabel_and_properties(c *Label_and_propertiesContext)

	// EnterProperties_clause is called when entering the properties_clause production.
	EnterProperties_clause(c *Properties_clauseContext)

	// EnterDerived_property_list is called when entering the derived_property_list production.
	EnterDerived_property_list(c *Derived_property_listContext)

	// EnterDerived_property is called when entering the derived_property production.
	EnterDerived_property(c *Derived_propertyContext)

	// EnterOpt_except_column_list is called when entering the opt_except_column_list production.
	EnterOpt_except_column_list(c *Opt_except_column_listContext)

	// EnterProperties_all_columns is called when entering the properties_all_columns production.
	EnterProperties_all_columns(c *Properties_all_columnsContext)

	// EnterOpt_dest_node_table_clause is called when entering the opt_dest_node_table_clause production.
	EnterOpt_dest_node_table_clause(c *Opt_dest_node_table_clauseContext)

	// EnterOpt_source_node_table_clause is called when entering the opt_source_node_table_clause production.
	EnterOpt_source_node_table_clause(c *Opt_source_node_table_clauseContext)

	// EnterOpt_key_clause is called when entering the opt_key_clause production.
	EnterOpt_key_clause(c *Opt_key_clauseContext)

	// EnterCreate_model_statement is called when entering the create_model_statement production.
	EnterCreate_model_statement(c *Create_model_statementContext)

	// EnterOpt_input_output_clause is called when entering the opt_input_output_clause production.
	EnterOpt_input_output_clause(c *Opt_input_output_clauseContext)

	// EnterOpt_transform_clause is called when entering the opt_transform_clause production.
	EnterOpt_transform_clause(c *Opt_transform_clauseContext)

	// EnterOpt_as_query_or_aliased_query_list is called when entering the opt_as_query_or_aliased_query_list production.
	EnterOpt_as_query_or_aliased_query_list(c *Opt_as_query_or_aliased_query_listContext)

	// EnterAliased_query_list is called when entering the aliased_query_list production.
	EnterAliased_query_list(c *Aliased_query_listContext)

	// EnterAs_query is called when entering the as_query production.
	EnterAs_query(c *As_queryContext)

	// EnterCreate_external_table_function_statement is called when entering the create_external_table_function_statement production.
	EnterCreate_external_table_function_statement(c *Create_external_table_function_statementContext)

	// EnterCreate_external_table_statement is called when entering the create_external_table_statement production.
	EnterCreate_external_table_statement(c *Create_external_table_statementContext)

	// EnterOpt_default_collate_clause is called when entering the opt_default_collate_clause production.
	EnterOpt_default_collate_clause(c *Opt_default_collate_clauseContext)

	// EnterOpt_like_path_expression is called when entering the opt_like_path_expression production.
	EnterOpt_like_path_expression(c *Opt_like_path_expressionContext)

	// EnterCreate_row_access_policy_statement is called when entering the create_row_access_policy_statement production.
	EnterCreate_row_access_policy_statement(c *Create_row_access_policy_statementContext)

	// EnterFilter_using_clause is called when entering the filter_using_clause production.
	EnterFilter_using_clause(c *Filter_using_clauseContext)

	// EnterCreate_row_access_policy_grant_to_clause is called when entering the create_row_access_policy_grant_to_clause production.
	EnterCreate_row_access_policy_grant_to_clause(c *Create_row_access_policy_grant_to_clauseContext)

	// EnterCreate_privilege_restriction_statement is called when entering the create_privilege_restriction_statement production.
	EnterCreate_privilege_restriction_statement(c *Create_privilege_restriction_statementContext)

	// EnterRestrict_to_clause is called when entering the restrict_to_clause production.
	EnterRestrict_to_clause(c *Restrict_to_clauseContext)

	// EnterPossibly_empty_grantee_list is called when entering the possibly_empty_grantee_list production.
	EnterPossibly_empty_grantee_list(c *Possibly_empty_grantee_listContext)

	// EnterCreate_index_statement is called when entering the create_index_statement production.
	EnterCreate_index_statement(c *Create_index_statementContext)

	// EnterOpt_create_index_statement_suffix is called when entering the opt_create_index_statement_suffix production.
	EnterOpt_create_index_statement_suffix(c *Opt_create_index_statement_suffixContext)

	// EnterSpanner_index_interleave_clause is called when entering the spanner_index_interleave_clause production.
	EnterSpanner_index_interleave_clause(c *Spanner_index_interleave_clauseContext)

	// EnterIndex_storing_list is called when entering the index_storing_list production.
	EnterIndex_storing_list(c *Index_storing_listContext)

	// EnterIndex_storing_expression_list is called when entering the index_storing_expression_list production.
	EnterIndex_storing_expression_list(c *Index_storing_expression_listContext)

	// EnterIndex_order_by_and_options is called when entering the index_order_by_and_options production.
	EnterIndex_order_by_and_options(c *Index_order_by_and_optionsContext)

	// EnterIndex_all_columns is called when entering the index_all_columns production.
	EnterIndex_all_columns(c *Index_all_columnsContext)

	// EnterOpt_with_column_options is called when entering the opt_with_column_options production.
	EnterOpt_with_column_options(c *Opt_with_column_optionsContext)

	// EnterAll_column_column_options is called when entering the all_column_column_options production.
	EnterAll_column_column_options(c *All_column_column_optionsContext)

	// EnterColumn_ordering_and_options_expr is called when entering the column_ordering_and_options_expr production.
	EnterColumn_ordering_and_options_expr(c *Column_ordering_and_options_exprContext)

	// EnterIndex_unnest_expression_list is called when entering the index_unnest_expression_list production.
	EnterIndex_unnest_expression_list(c *Index_unnest_expression_listContext)

	// EnterUnnest_expression_with_opt_alias_and_offset is called when entering the unnest_expression_with_opt_alias_and_offset production.
	EnterUnnest_expression_with_opt_alias_and_offset(c *Unnest_expression_with_opt_alias_and_offsetContext)

	// EnterOn_path_expression is called when entering the on_path_expression production.
	EnterOn_path_expression(c *On_path_expressionContext)

	// EnterIndex_type is called when entering the index_type production.
	EnterIndex_type(c *Index_typeContext)

	// EnterOpt_spanner_null_filtered is called when entering the opt_spanner_null_filtered production.
	EnterOpt_spanner_null_filtered(c *Opt_spanner_null_filteredContext)

	// EnterCreate_procedure_statement is called when entering the create_procedure_statement production.
	EnterCreate_procedure_statement(c *Create_procedure_statementContext)

	// EnterBegin_end_block_or_language_as_code is called when entering the begin_end_block_or_language_as_code production.
	EnterBegin_end_block_or_language_as_code(c *Begin_end_block_or_language_as_codeContext)

	// EnterBegin_end_block is called when entering the begin_end_block production.
	EnterBegin_end_block(c *Begin_end_blockContext)

	// EnterOpt_exception_handler is called when entering the opt_exception_handler production.
	EnterOpt_exception_handler(c *Opt_exception_handlerContext)

	// EnterStatement_list is called when entering the statement_list production.
	EnterStatement_list(c *Statement_listContext)

	// EnterUnterminated_non_empty_statement_list is called when entering the unterminated_non_empty_statement_list production.
	EnterUnterminated_non_empty_statement_list(c *Unterminated_non_empty_statement_listContext)

	// EnterUnterminated_statement is called when entering the unterminated_statement production.
	EnterUnterminated_statement(c *Unterminated_statementContext)

	// EnterUnterminated_script_statement is called when entering the unterminated_script_statement production.
	EnterUnterminated_script_statement(c *Unterminated_script_statementContext)

	// EnterLabel is called when entering the label production.
	EnterLabel(c *LabelContext)

	// EnterUnterminated_unlabeled_script_statement is called when entering the unterminated_unlabeled_script_statement production.
	EnterUnterminated_unlabeled_script_statement(c *Unterminated_unlabeled_script_statementContext)

	// EnterFor_in_statement is called when entering the for_in_statement production.
	EnterFor_in_statement(c *For_in_statementContext)

	// EnterRepeat_statement is called when entering the repeat_statement production.
	EnterRepeat_statement(c *Repeat_statementContext)

	// EnterUntil_clause is called when entering the until_clause production.
	EnterUntil_clause(c *Until_clauseContext)

	// EnterLoop_statement is called when entering the loop_statement production.
	EnterLoop_statement(c *Loop_statementContext)

	// EnterWhile_statement is called when entering the while_statement production.
	EnterWhile_statement(c *While_statementContext)

	// EnterRaise_statement is called when entering the raise_statement production.
	EnterRaise_statement(c *Raise_statementContext)

	// EnterReturn_statement is called when entering the return_statement production.
	EnterReturn_statement(c *Return_statementContext)

	// EnterContinue_statement is called when entering the continue_statement production.
	EnterContinue_statement(c *Continue_statementContext)

	// EnterVariable_declaration is called when entering the variable_declaration production.
	EnterVariable_declaration(c *Variable_declarationContext)

	// EnterBreak_statement is called when entering the break_statement production.
	EnterBreak_statement(c *Break_statementContext)

	// EnterCase_statement is called when entering the case_statement production.
	EnterCase_statement(c *Case_statementContext)

	// EnterWhen_then_clauses is called when entering the when_then_clauses production.
	EnterWhen_then_clauses(c *When_then_clausesContext)

	// EnterIf_statement is called when entering the if_statement production.
	EnterIf_statement(c *If_statementContext)

	// EnterElseif_clauses is called when entering the elseif_clauses production.
	EnterElseif_clauses(c *Elseif_clausesContext)

	// EnterOpt_else is called when entering the opt_else production.
	EnterOpt_else(c *Opt_elseContext)

	// EnterOpt_as_code is called when entering the opt_as_code production.
	EnterOpt_as_code(c *Opt_as_codeContext)

	// EnterOpt_external_security_clause is called when entering the opt_external_security_clause production.
	EnterOpt_external_security_clause(c *Opt_external_security_clauseContext)

	// EnterExternal_security_clause_kind is called when entering the external_security_clause_kind production.
	EnterExternal_security_clause_kind(c *External_security_clause_kindContext)

	// EnterProcedure_parameters is called when entering the procedure_parameters production.
	EnterProcedure_parameters(c *Procedure_parametersContext)

	// EnterProcedure_parameter is called when entering the procedure_parameter production.
	EnterProcedure_parameter(c *Procedure_parameterContext)

	// EnterProcedure_parameter_termination is called when entering the procedure_parameter_termination production.
	EnterProcedure_parameter_termination(c *Procedure_parameter_terminationContext)

	// EnterOpt_procedure_parameter_mode is called when entering the opt_procedure_parameter_mode production.
	EnterOpt_procedure_parameter_mode(c *Opt_procedure_parameter_modeContext)

	// EnterCreate_function_statement is called when entering the create_function_statement production.
	EnterCreate_function_statement(c *Create_function_statementContext)

	// EnterOpt_determinism_level is called when entering the opt_determinism_level production.
	EnterOpt_determinism_level(c *Opt_determinism_levelContext)

	// EnterOpt_sql_security_clause is called when entering the opt_sql_security_clause production.
	EnterOpt_sql_security_clause(c *Opt_sql_security_clauseContext)

	// EnterSql_security_clause_kind is called when entering the sql_security_clause_kind production.
	EnterSql_security_clause_kind(c *Sql_security_clause_kindContext)

	// EnterAs_sql_function_body_or_string is called when entering the as_sql_function_body_or_string production.
	EnterAs_sql_function_body_or_string(c *As_sql_function_body_or_stringContext)

	// EnterSql_function_body is called when entering the sql_function_body production.
	EnterSql_function_body(c *Sql_function_bodyContext)

	// EnterUnordered_options_body is called when entering the unordered_options_body production.
	EnterUnordered_options_body(c *Unordered_options_bodyContext)

	// EnterOpt_language_or_remote_with_connection is called when entering the opt_language_or_remote_with_connection production.
	EnterOpt_language_or_remote_with_connection(c *Opt_language_or_remote_with_connectionContext)

	// EnterLanguage is called when entering the language production.
	EnterLanguage(c *LanguageContext)

	// EnterRemote_with_connection_clause is called when entering the remote_with_connection_clause production.
	EnterRemote_with_connection_clause(c *Remote_with_connection_clauseContext)

	// EnterWith_connection_clause is called when entering the with_connection_clause production.
	EnterWith_connection_clause(c *With_connection_clauseContext)

	// EnterOpt_function_returns is called when entering the opt_function_returns production.
	EnterOpt_function_returns(c *Opt_function_returnsContext)

	// EnterOpt_returns is called when entering the opt_returns production.
	EnterOpt_returns(c *Opt_returnsContext)

	// EnterFunction_declaration is called when entering the function_declaration production.
	EnterFunction_declaration(c *Function_declarationContext)

	// EnterFunction_parameters is called when entering the function_parameters production.
	EnterFunction_parameters(c *Function_parametersContext)

	// EnterFunction_parameter is called when entering the function_parameter production.
	EnterFunction_parameter(c *Function_parameterContext)

	// EnterOpt_not_aggregate is called when entering the opt_not_aggregate production.
	EnterOpt_not_aggregate(c *Opt_not_aggregateContext)

	// EnterOpt_default_expression is called when entering the opt_default_expression production.
	EnterOpt_default_expression(c *Opt_default_expressionContext)

	// EnterType_or_tvf_schema is called when entering the type_or_tvf_schema production.
	EnterType_or_tvf_schema(c *Type_or_tvf_schemaContext)

	// EnterTvf_schema is called when entering the tvf_schema production.
	EnterTvf_schema(c *Tvf_schemaContext)

	// EnterTvf_schema_column is called when entering the tvf_schema_column production.
	EnterTvf_schema_column(c *Tvf_schema_columnContext)

	// EnterTemplated_parameter_type is called when entering the templated_parameter_type production.
	EnterTemplated_parameter_type(c *Templated_parameter_typeContext)

	// EnterTemplated_parameter_kind is called when entering the templated_parameter_kind production.
	EnterTemplated_parameter_kind(c *Templated_parameter_kindContext)

	// EnterOpt_aggregate is called when entering the opt_aggregate production.
	EnterOpt_aggregate(c *Opt_aggregateContext)

	// EnterCreate_database_statement is called when entering the create_database_statement production.
	EnterCreate_database_statement(c *Create_database_statementContext)

	// EnterCreate_connection_statement is called when entering the create_connection_statement production.
	EnterCreate_connection_statement(c *Create_connection_statementContext)

	// EnterCreate_constant_statement is called when entering the create_constant_statement production.
	EnterCreate_constant_statement(c *Create_constant_statementContext)

	// EnterOpt_or_replace is called when entering the opt_or_replace production.
	EnterOpt_or_replace(c *Opt_or_replaceContext)

	// EnterOpt_create_scope is called when entering the opt_create_scope production.
	EnterOpt_create_scope(c *Opt_create_scopeContext)

	// EnterRun_batch_statement is called when entering the run_batch_statement production.
	EnterRun_batch_statement(c *Run_batch_statementContext)

	// EnterAbort_batch_statement is called when entering the abort_batch_statement production.
	EnterAbort_batch_statement(c *Abort_batch_statementContext)

	// EnterStart_batch_statement is called when entering the start_batch_statement production.
	EnterStart_batch_statement(c *Start_batch_statementContext)

	// EnterRollback_statement is called when entering the rollback_statement production.
	EnterRollback_statement(c *Rollback_statementContext)

	// EnterCommit_statement is called when entering the commit_statement production.
	EnterCommit_statement(c *Commit_statementContext)

	// EnterSet_statement is called when entering the set_statement production.
	EnterSet_statement(c *Set_statementContext)

	// EnterIdentifier_list is called when entering the identifier_list production.
	EnterIdentifier_list(c *Identifier_listContext)

	// EnterBegin_statement is called when entering the begin_statement production.
	EnterBegin_statement(c *Begin_statementContext)

	// EnterBegin_transaction_keywords is called when entering the begin_transaction_keywords production.
	EnterBegin_transaction_keywords(c *Begin_transaction_keywordsContext)

	// EnterTransaction_mode_list is called when entering the transaction_mode_list production.
	EnterTransaction_mode_list(c *Transaction_mode_listContext)

	// EnterTransaction_mode is called when entering the transaction_mode production.
	EnterTransaction_mode(c *Transaction_modeContext)

	// EnterTruncate_statement is called when entering the truncate_statement production.
	EnterTruncate_statement(c *Truncate_statementContext)

	// EnterMerge_statement is called when entering the merge_statement production.
	EnterMerge_statement(c *Merge_statementContext)

	// EnterMerge_source is called when entering the merge_source production.
	EnterMerge_source(c *Merge_sourceContext)

	// EnterMerge_when_clause is called when entering the merge_when_clause production.
	EnterMerge_when_clause(c *Merge_when_clauseContext)

	// EnterMerge_action is called when entering the merge_action production.
	EnterMerge_action(c *Merge_actionContext)

	// EnterMerge_insert_value_list_or_source_row is called when entering the merge_insert_value_list_or_source_row production.
	EnterMerge_insert_value_list_or_source_row(c *Merge_insert_value_list_or_source_rowContext)

	// EnterBy_target is called when entering the by_target production.
	EnterBy_target(c *By_targetContext)

	// EnterOpt_and_expression is called when entering the opt_and_expression production.
	EnterOpt_and_expression(c *Opt_and_expressionContext)

	// EnterStatement_level_hint is called when entering the statement_level_hint production.
	EnterStatement_level_hint(c *Statement_level_hintContext)

	// EnterQuery_statement is called when entering the query_statement production.
	EnterQuery_statement(c *Query_statementContext)

	// EnterDml_statement is called when entering the dml_statement production.
	EnterDml_statement(c *Dml_statementContext)

	// EnterUpdate_statement is called when entering the update_statement production.
	EnterUpdate_statement(c *Update_statementContext)

	// EnterDelete_statement is called when entering the delete_statement production.
	EnterDelete_statement(c *Delete_statementContext)

	// EnterInsert_statement is called when entering the insert_statement production.
	EnterInsert_statement(c *Insert_statementContext)

	// EnterOn_conflict_clause is called when entering the on_conflict_clause production.
	EnterOn_conflict_clause(c *On_conflict_clauseContext)

	// EnterOpt_where_expression is called when entering the opt_where_expression production.
	EnterOpt_where_expression(c *Opt_where_expressionContext)

	// EnterOpt_conflict_target is called when entering the opt_conflict_target production.
	EnterOpt_conflict_target(c *Opt_conflict_targetContext)

	// EnterUpdate_item_list is called when entering the update_item_list production.
	EnterUpdate_item_list(c *Update_item_listContext)

	// EnterUpdate_item is called when entering the update_item production.
	EnterUpdate_item(c *Update_itemContext)

	// EnterUpdate_set_value is called when entering the update_set_value production.
	EnterUpdate_set_value(c *Update_set_valueContext)

	// EnterNested_dml_statement is called when entering the nested_dml_statement production.
	EnterNested_dml_statement(c *Nested_dml_statementContext)

	// EnterInsert_values_list_or_table_clause is called when entering the insert_values_list_or_table_clause production.
	EnterInsert_values_list_or_table_clause(c *Insert_values_list_or_table_clauseContext)

	// EnterTable_clause_unreversed is called when entering the table_clause_unreversed production.
	EnterTable_clause_unreversed(c *Table_clause_unreversedContext)

	// EnterTable_clause_no_keyword is called when entering the table_clause_no_keyword production.
	EnterTable_clause_no_keyword(c *Table_clause_no_keywordContext)

	// EnterOpt_returning_clause is called when entering the opt_returning_clause production.
	EnterOpt_returning_clause(c *Opt_returning_clauseContext)

	// EnterOpt_assert_rows_modified is called when entering the opt_assert_rows_modified production.
	EnterOpt_assert_rows_modified(c *Opt_assert_rows_modifiedContext)

	// EnterInsert_values_or_query is called when entering the insert_values_or_query production.
	EnterInsert_values_or_query(c *Insert_values_or_queryContext)

	// EnterInsert_values_list is called when entering the insert_values_list production.
	EnterInsert_values_list(c *Insert_values_listContext)

	// EnterInsert_values_row is called when entering the insert_values_row production.
	EnterInsert_values_row(c *Insert_values_rowContext)

	// EnterExpression_or_default is called when entering the expression_or_default production.
	EnterExpression_or_default(c *Expression_or_defaultContext)

	// EnterInsert_statement_prefix is called when entering the insert_statement_prefix production.
	EnterInsert_statement_prefix(c *Insert_statement_prefixContext)

	// EnterMaybe_dashed_generalized_path_expression is called when entering the maybe_dashed_generalized_path_expression production.
	EnterMaybe_dashed_generalized_path_expression(c *Maybe_dashed_generalized_path_expressionContext)

	// EnterOpt_into is called when entering the opt_into production.
	EnterOpt_into(c *Opt_intoContext)

	// EnterOpt_or_ignore_replace_update is called when entering the opt_or_ignore_replace_update production.
	EnterOpt_or_ignore_replace_update(c *Opt_or_ignore_replace_updateContext)

	// EnterAlter_statement is called when entering the alter_statement production.
	EnterAlter_statement(c *Alter_statementContext)

	// EnterAnalyze_statement is called when entering the analyze_statement production.
	EnterAnalyze_statement(c *Analyze_statementContext)

	// EnterAssert_statement is called when entering the assert_statement production.
	EnterAssert_statement(c *Assert_statementContext)

	// EnterAux_load_data_statement is called when entering the aux_load_data_statement production.
	EnterAux_load_data_statement(c *Aux_load_data_statementContext)

	// EnterClone_data_statement is called when entering the clone_data_statement production.
	EnterClone_data_statement(c *Clone_data_statementContext)

	// EnterClone_data_source_list is called when entering the clone_data_source_list production.
	EnterClone_data_source_list(c *Clone_data_source_listContext)

	// EnterClone_data_source is called when entering the clone_data_source production.
	EnterClone_data_source(c *Clone_data_sourceContext)

	// EnterOpt_external_table_with_clauses is called when entering the opt_external_table_with_clauses production.
	EnterOpt_external_table_with_clauses(c *Opt_external_table_with_clausesContext)

	// EnterWith_partition_columns_clause is called when entering the with_partition_columns_clause production.
	EnterWith_partition_columns_clause(c *With_partition_columns_clauseContext)

	// EnterAux_load_data_from_files_options_list is called when entering the aux_load_data_from_files_options_list production.
	EnterAux_load_data_from_files_options_list(c *Aux_load_data_from_files_options_listContext)

	// EnterCluster_by_clause_prefix_no_hint is called when entering the cluster_by_clause_prefix_no_hint production.
	EnterCluster_by_clause_prefix_no_hint(c *Cluster_by_clause_prefix_no_hintContext)

	// EnterLoad_data_partitions_clause is called when entering the load_data_partitions_clause production.
	EnterLoad_data_partitions_clause(c *Load_data_partitions_clauseContext)

	// EnterMaybe_dashed_path_expression_with_scope is called when entering the maybe_dashed_path_expression_with_scope production.
	EnterMaybe_dashed_path_expression_with_scope(c *Maybe_dashed_path_expression_with_scopeContext)

	// EnterTable_element_list is called when entering the table_element_list production.
	EnterTable_element_list(c *Table_element_listContext)

	// EnterTable_element is called when entering the table_element production.
	EnterTable_element(c *Table_elementContext)

	// EnterTable_constraint_definition is called when entering the table_constraint_definition production.
	EnterTable_constraint_definition(c *Table_constraint_definitionContext)

	// EnterAppend_or_overwrite is called when entering the append_or_overwrite production.
	EnterAppend_or_overwrite(c *Append_or_overwriteContext)

	// EnterOpt_description is called when entering the opt_description production.
	EnterOpt_description(c *Opt_descriptionContext)

	// EnterTable_and_column_info_list is called when entering the table_and_column_info_list production.
	EnterTable_and_column_info_list(c *Table_and_column_info_listContext)

	// EnterTable_and_column_info is called when entering the table_and_column_info production.
	EnterTable_and_column_info(c *Table_and_column_infoContext)

	// EnterRow_access_policy_alter_action_list is called when entering the row_access_policy_alter_action_list production.
	EnterRow_access_policy_alter_action_list(c *Row_access_policy_alter_action_listContext)

	// EnterRow_access_policy_alter_action is called when entering the row_access_policy_alter_action production.
	EnterRow_access_policy_alter_action(c *Row_access_policy_alter_actionContext)

	// EnterGrant_to_clause is called when entering the grant_to_clause production.
	EnterGrant_to_clause(c *Grant_to_clauseContext)

	// EnterGrantee_list is called when entering the grantee_list production.
	EnterGrantee_list(c *Grantee_listContext)

	// EnterPrivilege_list is called when entering the privilege_list production.
	EnterPrivilege_list(c *Privilege_listContext)

	// EnterPrivilege is called when entering the privilege production.
	EnterPrivilege(c *PrivilegeContext)

	// EnterPath_expression_list_with_parens is called when entering the path_expression_list_with_parens production.
	EnterPath_expression_list_with_parens(c *Path_expression_list_with_parensContext)

	// EnterPrivilege_name is called when entering the privilege_name production.
	EnterPrivilege_name(c *Privilege_nameContext)

	// EnterGeneric_entity_type is called when entering the generic_entity_type production.
	EnterGeneric_entity_type(c *Generic_entity_typeContext)

	// EnterGeneric_entity_type_unchecked is called when entering the generic_entity_type_unchecked production.
	EnterGeneric_entity_type_unchecked(c *Generic_entity_type_uncheckedContext)

	// EnterSchema_object_kind is called when entering the schema_object_kind production.
	EnterSchema_object_kind(c *Schema_object_kindContext)

	// EnterAlter_action_list is called when entering the alter_action_list production.
	EnterAlter_action_list(c *Alter_action_listContext)

	// EnterAlter_action is called when entering the alter_action production.
	EnterAlter_action(c *Alter_actionContext)

	// EnterSpanner_set_on_delete_action is called when entering the spanner_set_on_delete_action production.
	EnterSpanner_set_on_delete_action(c *Spanner_set_on_delete_actionContext)

	// EnterSpanner_alter_column_action is called when entering the spanner_alter_column_action production.
	EnterSpanner_alter_column_action(c *Spanner_alter_column_actionContext)

	// EnterSpanner_generated_or_default is called when entering the spanner_generated_or_default production.
	EnterSpanner_generated_or_default(c *Spanner_generated_or_defaultContext)

	// EnterGeneric_sub_entity_type is called when entering the generic_sub_entity_type production.
	EnterGeneric_sub_entity_type(c *Generic_sub_entity_typeContext)

	// EnterSub_entity_type_identifier is called when entering the sub_entity_type_identifier production.
	EnterSub_entity_type_identifier(c *Sub_entity_type_identifierContext)

	// EnterFill_using_expression is called when entering the fill_using_expression production.
	EnterFill_using_expression(c *Fill_using_expressionContext)

	// EnterColumn_position is called when entering the column_position production.
	EnterColumn_position(c *Column_positionContext)

	// EnterTable_column_definition is called when entering the table_column_definition production.
	EnterTable_column_definition(c *Table_column_definitionContext)

	// EnterColumn_attributes is called when entering the column_attributes production.
	EnterColumn_attributes(c *Column_attributesContext)

	// EnterColumn_attribute is called when entering the column_attribute production.
	EnterColumn_attribute(c *Column_attributeContext)

	// EnterPrimary_key_column_attribute is called when entering the primary_key_column_attribute production.
	EnterPrimary_key_column_attribute(c *Primary_key_column_attributeContext)

	// EnterForeign_key_column_attribute is called when entering the foreign_key_column_attribute production.
	EnterForeign_key_column_attribute(c *Foreign_key_column_attributeContext)

	// EnterHidden_column_attribute is called when entering the hidden_column_attribute production.
	EnterHidden_column_attribute(c *Hidden_column_attributeContext)

	// EnterOpt_constraint_identity is called when entering the opt_constraint_identity production.
	EnterOpt_constraint_identity(c *Opt_constraint_identityContext)

	// EnterTable_column_schema is called when entering the table_column_schema production.
	EnterTable_column_schema(c *Table_column_schemaContext)

	// EnterOpt_column_info is called when entering the opt_column_info production.
	EnterOpt_column_info(c *Opt_column_infoContext)

	// EnterInvalid_generated_column is called when entering the invalid_generated_column production.
	EnterInvalid_generated_column(c *Invalid_generated_columnContext)

	// EnterInvalid_default_column is called when entering the invalid_default_column production.
	EnterInvalid_default_column(c *Invalid_default_columnContext)

	// EnterDefault_column_info is called when entering the default_column_info production.
	EnterDefault_column_info(c *Default_column_infoContext)

	// EnterGenerated_column_info is called when entering the generated_column_info production.
	EnterGenerated_column_info(c *Generated_column_infoContext)

	// EnterIdentity_column_info is called when entering the identity_column_info production.
	EnterIdentity_column_info(c *Identity_column_infoContext)

	// EnterOpt_start_with is called when entering the opt_start_with production.
	EnterOpt_start_with(c *Opt_start_withContext)

	// EnterOpt_increment_by is called when entering the opt_increment_by production.
	EnterOpt_increment_by(c *Opt_increment_byContext)

	// EnterOpt_maxvalue is called when entering the opt_maxvalue production.
	EnterOpt_maxvalue(c *Opt_maxvalueContext)

	// EnterOpt_minvalue is called when entering the opt_minvalue production.
	EnterOpt_minvalue(c *Opt_minvalueContext)

	// EnterOpt_cycle is called when entering the opt_cycle production.
	EnterOpt_cycle(c *Opt_cycleContext)

	// EnterSigned_numeric_literal is called when entering the signed_numeric_literal production.
	EnterSigned_numeric_literal(c *Signed_numeric_literalContext)

	// EnterStored_mode is called when entering the stored_mode production.
	EnterStored_mode(c *Stored_modeContext)

	// EnterGenerated_mode is called when entering the generated_mode production.
	EnterGenerated_mode(c *Generated_modeContext)

	// EnterColumn_schema_inner is called when entering the column_schema_inner production.
	EnterColumn_schema_inner(c *Column_schema_innerContext)

	// EnterRaw_column_schema_inner is called when entering the raw_column_schema_inner production.
	EnterRaw_column_schema_inner(c *Raw_column_schema_innerContext)

	// EnterRange_column_schema_inner is called when entering the range_column_schema_inner production.
	EnterRange_column_schema_inner(c *Range_column_schema_innerContext)

	// EnterStruct_column_schema_inner is called when entering the struct_column_schema_inner production.
	EnterStruct_column_schema_inner(c *Struct_column_schema_innerContext)

	// EnterStruct_column_field is called when entering the struct_column_field production.
	EnterStruct_column_field(c *Struct_column_fieldContext)

	// EnterSimple_column_schema_inner is called when entering the simple_column_schema_inner production.
	EnterSimple_column_schema_inner(c *Simple_column_schema_innerContext)

	// EnterArray_column_schema_inner is called when entering the array_column_schema_inner production.
	EnterArray_column_schema_inner(c *Array_column_schema_innerContext)

	// EnterField_schema is called when entering the field_schema production.
	EnterField_schema(c *Field_schemaContext)

	// EnterOpt_field_attributes is called when entering the opt_field_attributes production.
	EnterOpt_field_attributes(c *Opt_field_attributesContext)

	// EnterNot_null_column_attribute is called when entering the not_null_column_attribute production.
	EnterNot_null_column_attribute(c *Not_null_column_attributeContext)

	// EnterPrimary_key_or_table_constraint_spec is called when entering the primary_key_or_table_constraint_spec production.
	EnterPrimary_key_or_table_constraint_spec(c *Primary_key_or_table_constraint_specContext)

	// EnterOpt_if_not_exists is called when entering the opt_if_not_exists production.
	EnterOpt_if_not_exists(c *Opt_if_not_existsContext)

	// EnterPrimary_key_spec is called when entering the primary_key_spec production.
	EnterPrimary_key_spec(c *Primary_key_specContext)

	// EnterPrimary_key_element_list is called when entering the primary_key_element_list production.
	EnterPrimary_key_element_list(c *Primary_key_element_listContext)

	// EnterPrimary_key_element is called when entering the primary_key_element production.
	EnterPrimary_key_element(c *Primary_key_elementContext)

	// EnterTable_constraint_spec is called when entering the table_constraint_spec production.
	EnterTable_constraint_spec(c *Table_constraint_specContext)

	// EnterForeign_key_reference is called when entering the foreign_key_reference production.
	EnterForeign_key_reference(c *Foreign_key_referenceContext)

	// EnterOpt_foreign_key_action is called when entering the opt_foreign_key_action production.
	EnterOpt_foreign_key_action(c *Opt_foreign_key_actionContext)

	// EnterForeign_key_on_update is called when entering the foreign_key_on_update production.
	EnterForeign_key_on_update(c *Foreign_key_on_updateContext)

	// EnterForeign_key_on_delete is called when entering the foreign_key_on_delete production.
	EnterForeign_key_on_delete(c *Foreign_key_on_deleteContext)

	// EnterForeign_key_action is called when entering the foreign_key_action production.
	EnterForeign_key_action(c *Foreign_key_actionContext)

	// EnterOpt_foreign_key_match is called when entering the opt_foreign_key_match production.
	EnterOpt_foreign_key_match(c *Opt_foreign_key_matchContext)

	// EnterForeign_key_match_mode is called when entering the foreign_key_match_mode production.
	EnterForeign_key_match_mode(c *Foreign_key_match_modeContext)

	// EnterColumn_list is called when entering the column_list production.
	EnterColumn_list(c *Column_listContext)

	// EnterOpt_options_list is called when entering the opt_options_list production.
	EnterOpt_options_list(c *Opt_options_listContext)

	// EnterConstraint_enforcement is called when entering the constraint_enforcement production.
	EnterConstraint_enforcement(c *Constraint_enforcementContext)

	// EnterGeneric_entity_body is called when entering the generic_entity_body production.
	EnterGeneric_entity_body(c *Generic_entity_bodyContext)

	// EnterOpt_if_exists is called when entering the opt_if_exists production.
	EnterOpt_if_exists(c *Opt_if_existsContext)

	// EnterTable_or_table_function is called when entering the table_or_table_function production.
	EnterTable_or_table_function(c *Table_or_table_functionContext)

	// EnterQuery is called when entering the query production.
	EnterQuery(c *QueryContext)

	// EnterQuery_without_pipe_operators is called when entering the query_without_pipe_operators production.
	EnterQuery_without_pipe_operators(c *Query_without_pipe_operatorsContext)

	// EnterBad_keyword_after_from_query is called when entering the bad_keyword_after_from_query production.
	EnterBad_keyword_after_from_query(c *Bad_keyword_after_from_queryContext)

	// EnterBad_keyword_after_from_query_allows_parens is called when entering the bad_keyword_after_from_query_allows_parens production.
	EnterBad_keyword_after_from_query_allows_parens(c *Bad_keyword_after_from_query_allows_parensContext)

	// EnterWith_clause_with_trailing_comma is called when entering the with_clause_with_trailing_comma production.
	EnterWith_clause_with_trailing_comma(c *With_clause_with_trailing_commaContext)

	// EnterSelect_or_from_keyword is called when entering the select_or_from_keyword production.
	EnterSelect_or_from_keyword(c *Select_or_from_keywordContext)

	// EnterQuery_primary_or_set_operation is called when entering the query_primary_or_set_operation production.
	EnterQuery_primary_or_set_operation(c *Query_primary_or_set_operationContext)

	// EnterQuery_set_operation is called when entering the query_set_operation production.
	EnterQuery_set_operation(c *Query_set_operationContext)

	// EnterQuery_set_operation_prefix is called when entering the query_set_operation_prefix production.
	EnterQuery_set_operation_prefix(c *Query_set_operation_prefixContext)

	// EnterQuery_set_operation_item is called when entering the query_set_operation_item production.
	EnterQuery_set_operation_item(c *Query_set_operation_itemContext)

	// EnterQuery_primary is called when entering the query_primary production.
	EnterQuery_primary(c *Query_primaryContext)

	// EnterSet_operation_metadata is called when entering the set_operation_metadata production.
	EnterSet_operation_metadata(c *Set_operation_metadataContext)

	// EnterOpt_column_match_suffix is called when entering the opt_column_match_suffix production.
	EnterOpt_column_match_suffix(c *Opt_column_match_suffixContext)

	// EnterOpt_strict is called when entering the opt_strict production.
	EnterOpt_strict(c *Opt_strictContext)

	// EnterAll_or_distinct is called when entering the all_or_distinct production.
	EnterAll_or_distinct(c *All_or_distinctContext)

	// EnterQuery_set_operation_type is called when entering the query_set_operation_type production.
	EnterQuery_set_operation_type(c *Query_set_operation_typeContext)

	// EnterOpt_corresponding_outer_mode is called when entering the opt_corresponding_outer_mode production.
	EnterOpt_corresponding_outer_mode(c *Opt_corresponding_outer_modeContext)

	// EnterOpt_outer is called when entering the opt_outer production.
	EnterOpt_outer(c *Opt_outerContext)

	// EnterWith_clause is called when entering the with_clause production.
	EnterWith_clause(c *With_clauseContext)

	// EnterAliased_query is called when entering the aliased_query production.
	EnterAliased_query(c *Aliased_queryContext)

	// EnterOpt_aliased_query_modifiers is called when entering the opt_aliased_query_modifiers production.
	EnterOpt_aliased_query_modifiers(c *Opt_aliased_query_modifiersContext)

	// EnterRecursion_depth_modifier is called when entering the recursion_depth_modifier production.
	EnterRecursion_depth_modifier(c *Recursion_depth_modifierContext)

	// EnterPossibly_unbounded_int_literal_or_parameter is called when entering the possibly_unbounded_int_literal_or_parameter production.
	EnterPossibly_unbounded_int_literal_or_parameter(c *Possibly_unbounded_int_literal_or_parameterContext)

	// EnterInt_literal_or_parameter is called when entering the int_literal_or_parameter production.
	EnterInt_literal_or_parameter(c *Int_literal_or_parameterContext)

	// EnterOrder_by_clause is called when entering the order_by_clause production.
	EnterOrder_by_clause(c *Order_by_clauseContext)

	// EnterOrder_by_clause_prefix is called when entering the order_by_clause_prefix production.
	EnterOrder_by_clause_prefix(c *Order_by_clause_prefixContext)

	// EnterOrdering_expression is called when entering the ordering_expression production.
	EnterOrdering_expression(c *Ordering_expressionContext)

	// EnterSelect is called when entering the select production.
	EnterSelect(c *SelectContext)

	// EnterOpt_clauses_following_from is called when entering the opt_clauses_following_from production.
	EnterOpt_clauses_following_from(c *Opt_clauses_following_fromContext)

	// EnterOpt_clauses_following_where is called when entering the opt_clauses_following_where production.
	EnterOpt_clauses_following_where(c *Opt_clauses_following_whereContext)

	// EnterOpt_clauses_following_group_by is called when entering the opt_clauses_following_group_by production.
	EnterOpt_clauses_following_group_by(c *Opt_clauses_following_group_byContext)

	// EnterWindow_clause is called when entering the window_clause production.
	EnterWindow_clause(c *Window_clauseContext)

	// EnterWindow_clause_prefix is called when entering the window_clause_prefix production.
	EnterWindow_clause_prefix(c *Window_clause_prefixContext)

	// EnterWindow_definition is called when entering the window_definition production.
	EnterWindow_definition(c *Window_definitionContext)

	// EnterWhere_clause is called when entering the where_clause production.
	EnterWhere_clause(c *Where_clauseContext)

	// EnterHaving_clause is called when entering the having_clause production.
	EnterHaving_clause(c *Having_clauseContext)

	// EnterGroup_by_clause is called when entering the group_by_clause production.
	EnterGroup_by_clause(c *Group_by_clauseContext)

	// EnterGroup_by_all is called when entering the group_by_all production.
	EnterGroup_by_all(c *Group_by_allContext)

	// EnterSelect_clause is called when entering the select_clause production.
	EnterSelect_clause(c *Select_clauseContext)

	// EnterOpt_select_as_clause is called when entering the opt_select_as_clause production.
	EnterOpt_select_as_clause(c *Opt_select_as_clauseContext)

	// EnterOpt_select_with is called when entering the opt_select_with production.
	EnterOpt_select_with(c *Opt_select_withContext)

	// EnterFrom_clause is called when entering the from_clause production.
	EnterFrom_clause(c *From_clauseContext)

	// EnterFrom_clause_contents is called when entering the from_clause_contents production.
	EnterFrom_clause_contents(c *From_clause_contentsContext)

	// EnterFrom_clause_contents_suffix is called when entering the from_clause_contents_suffix production.
	EnterFrom_clause_contents_suffix(c *From_clause_contents_suffixContext)

	// EnterTable_primary is called when entering the table_primary production.
	EnterTable_primary(c *Table_primaryContext)

	// EnterTvf_with_suffixes is called when entering the tvf_with_suffixes production.
	EnterTvf_with_suffixes(c *Tvf_with_suffixesContext)

	// EnterPivot_or_unpivot_clause_and_aliases is called when entering the pivot_or_unpivot_clause_and_aliases production.
	EnterPivot_or_unpivot_clause_and_aliases(c *Pivot_or_unpivot_clause_and_aliasesContext)

	// EnterAs_alias is called when entering the as_alias production.
	EnterAs_alias(c *As_aliasContext)

	// EnterSample_clause is called when entering the sample_clause production.
	EnterSample_clause(c *Sample_clauseContext)

	// EnterOpt_sample_clause_suffix is called when entering the opt_sample_clause_suffix production.
	EnterOpt_sample_clause_suffix(c *Opt_sample_clause_suffixContext)

	// EnterRepeatable_clause is called when entering the repeatable_clause production.
	EnterRepeatable_clause(c *Repeatable_clauseContext)

	// EnterPossibly_cast_int_literal_or_parameter is called when entering the possibly_cast_int_literal_or_parameter production.
	EnterPossibly_cast_int_literal_or_parameter(c *Possibly_cast_int_literal_or_parameterContext)

	// EnterCast_int_literal_or_parameter is called when entering the cast_int_literal_or_parameter production.
	EnterCast_int_literal_or_parameter(c *Cast_int_literal_or_parameterContext)

	// EnterSample_size is called when entering the sample_size production.
	EnterSample_size(c *Sample_sizeContext)

	// EnterSample_size_value is called when entering the sample_size_value production.
	EnterSample_size_value(c *Sample_size_valueContext)

	// EnterSample_size_unit is called when entering the sample_size_unit production.
	EnterSample_size_unit(c *Sample_size_unitContext)

	// EnterPartition_by_clause_prefix_no_hint is called when entering the partition_by_clause_prefix_no_hint production.
	EnterPartition_by_clause_prefix_no_hint(c *Partition_by_clause_prefix_no_hintContext)

	// EnterMatch_recognize_clause is called when entering the match_recognize_clause production.
	EnterMatch_recognize_clause(c *Match_recognize_clauseContext)

	// EnterRow_pattern_expr is called when entering the row_pattern_expr production.
	EnterRow_pattern_expr(c *Row_pattern_exprContext)

	// EnterRow_pattern_concatenation is called when entering the row_pattern_concatenation production.
	EnterRow_pattern_concatenation(c *Row_pattern_concatenationContext)

	// EnterRow_pattern_factor is called when entering the row_pattern_factor production.
	EnterRow_pattern_factor(c *Row_pattern_factorContext)

	// EnterSelect_list_prefix_with_as_aliases is called when entering the select_list_prefix_with_as_aliases production.
	EnterSelect_list_prefix_with_as_aliases(c *Select_list_prefix_with_as_aliasesContext)

	// EnterSelect_column_expr_with_as_alias is called when entering the select_column_expr_with_as_alias production.
	EnterSelect_column_expr_with_as_alias(c *Select_column_expr_with_as_aliasContext)

	// EnterTable_subquery is called when entering the table_subquery production.
	EnterTable_subquery(c *Table_subqueryContext)

	// EnterJoin is called when entering the join production.
	EnterJoin(c *JoinContext)

	// EnterJoin_item is called when entering the join_item production.
	EnterJoin_item(c *Join_itemContext)

	// EnterOn_or_using_clause_list is called when entering the on_or_using_clause_list production.
	EnterOn_or_using_clause_list(c *On_or_using_clause_listContext)

	// EnterOn_or_using_clause is called when entering the on_or_using_clause production.
	EnterOn_or_using_clause(c *On_or_using_clauseContext)

	// EnterUsing_clause is called when entering the using_clause production.
	EnterUsing_clause(c *Using_clauseContext)

	// EnterJoin_hint is called when entering the join_hint production.
	EnterJoin_hint(c *Join_hintContext)

	// EnterTable_path_expression is called when entering the table_path_expression production.
	EnterTable_path_expression(c *Table_path_expressionContext)

	// EnterOpt_at_system_time is called when entering the opt_at_system_time production.
	EnterOpt_at_system_time(c *Opt_at_system_timeContext)

	// EnterOpt_with_offset_and_alias is called when entering the opt_with_offset_and_alias production.
	EnterOpt_with_offset_and_alias(c *Opt_with_offset_and_aliasContext)

	// EnterOpt_pivot_or_unpivot_clause_and_alias is called when entering the opt_pivot_or_unpivot_clause_and_alias production.
	EnterOpt_pivot_or_unpivot_clause_and_alias(c *Opt_pivot_or_unpivot_clause_and_aliasContext)

	// EnterTable_path_expression_base is called when entering the table_path_expression_base production.
	EnterTable_path_expression_base(c *Table_path_expression_baseContext)

	// EnterMaybe_slashed_or_dashed_path_expression is called when entering the maybe_slashed_or_dashed_path_expression production.
	EnterMaybe_slashed_or_dashed_path_expression(c *Maybe_slashed_or_dashed_path_expressionContext)

	// EnterMaybe_dashed_path_expression is called when entering the maybe_dashed_path_expression production.
	EnterMaybe_dashed_path_expression(c *Maybe_dashed_path_expressionContext)

	// EnterDashed_path_expression is called when entering the dashed_path_expression production.
	EnterDashed_path_expression(c *Dashed_path_expressionContext)

	// EnterDashed_identifier is called when entering the dashed_identifier production.
	EnterDashed_identifier(c *Dashed_identifierContext)

	// EnterSlashed_identifier is called when entering the slashed_identifier production.
	EnterSlashed_identifier(c *Slashed_identifierContext)

	// EnterIdentifier_or_integer is called when entering the identifier_or_integer production.
	EnterIdentifier_or_integer(c *Identifier_or_integerContext)

	// EnterSlashed_identifier_separator is called when entering the slashed_identifier_separator production.
	EnterSlashed_identifier_separator(c *Slashed_identifier_separatorContext)

	// EnterSlashed_path_expression is called when entering the slashed_path_expression production.
	EnterSlashed_path_expression(c *Slashed_path_expressionContext)

	// EnterUnnest_expression is called when entering the unnest_expression production.
	EnterUnnest_expression(c *Unnest_expressionContext)

	// EnterUnnest_expression_prefix is called when entering the unnest_expression_prefix production.
	EnterUnnest_expression_prefix(c *Unnest_expression_prefixContext)

	// EnterOpt_array_zip_mode is called when entering the opt_array_zip_mode production.
	EnterOpt_array_zip_mode(c *Opt_array_zip_modeContext)

	// EnterExpression_with_opt_alias is called when entering the expression_with_opt_alias production.
	EnterExpression_with_opt_alias(c *Expression_with_opt_aliasContext)

	// EnterTvf_prefix is called when entering the tvf_prefix production.
	EnterTvf_prefix(c *Tvf_prefixContext)

	// EnterTvf_argument is called when entering the tvf_argument production.
	EnterTvf_argument(c *Tvf_argumentContext)

	// EnterConnection_clause is called when entering the connection_clause production.
	EnterConnection_clause(c *Connection_clauseContext)

	// EnterPath_expression_or_default is called when entering the path_expression_or_default production.
	EnterPath_expression_or_default(c *Path_expression_or_defaultContext)

	// EnterDescriptor_argument is called when entering the descriptor_argument production.
	EnterDescriptor_argument(c *Descriptor_argumentContext)

	// EnterDescriptor_column_list is called when entering the descriptor_column_list production.
	EnterDescriptor_column_list(c *Descriptor_column_listContext)

	// EnterDescriptor_column is called when entering the descriptor_column production.
	EnterDescriptor_column(c *Descriptor_columnContext)

	// EnterTable_clause is called when entering the table_clause production.
	EnterTable_clause(c *Table_clauseContext)

	// EnterModel_clause is called when entering the model_clause production.
	EnterModel_clause(c *Model_clauseContext)

	// EnterQualify_clause_nonreserved is called when entering the qualify_clause_nonreserved production.
	EnterQualify_clause_nonreserved(c *Qualify_clause_nonreservedContext)

	// EnterUnpivot_clause is called when entering the unpivot_clause production.
	EnterUnpivot_clause(c *Unpivot_clauseContext)

	// EnterUnpivot_in_item_list is called when entering the unpivot_in_item_list production.
	EnterUnpivot_in_item_list(c *Unpivot_in_item_listContext)

	// EnterUnpivot_in_item_list_prefix is called when entering the unpivot_in_item_list_prefix production.
	EnterUnpivot_in_item_list_prefix(c *Unpivot_in_item_list_prefixContext)

	// EnterUnpivot_in_item is called when entering the unpivot_in_item production.
	EnterUnpivot_in_item(c *Unpivot_in_itemContext)

	// EnterOpt_as_string_or_integer is called when entering the opt_as_string_or_integer production.
	EnterOpt_as_string_or_integer(c *Opt_as_string_or_integerContext)

	// EnterPath_expression_list_with_opt_parens is called when entering the path_expression_list_with_opt_parens production.
	EnterPath_expression_list_with_opt_parens(c *Path_expression_list_with_opt_parensContext)

	// EnterPath_expression_list is called when entering the path_expression_list production.
	EnterPath_expression_list(c *Path_expression_listContext)

	// EnterUnpivot_nulls_filter is called when entering the unpivot_nulls_filter production.
	EnterUnpivot_nulls_filter(c *Unpivot_nulls_filterContext)

	// EnterPivot_clause is called when entering the pivot_clause production.
	EnterPivot_clause(c *Pivot_clauseContext)

	// EnterPivot_expression_list is called when entering the pivot_expression_list production.
	EnterPivot_expression_list(c *Pivot_expression_listContext)

	// EnterPivot_expression is called when entering the pivot_expression production.
	EnterPivot_expression(c *Pivot_expressionContext)

	// EnterPivot_value_list is called when entering the pivot_value_list production.
	EnterPivot_value_list(c *Pivot_value_listContext)

	// EnterPivot_value is called when entering the pivot_value production.
	EnterPivot_value(c *Pivot_valueContext)

	// EnterTvf_prefix_no_args is called when entering the tvf_prefix_no_args production.
	EnterTvf_prefix_no_args(c *Tvf_prefix_no_argsContext)

	// EnterJoin_type is called when entering the join_type production.
	EnterJoin_type(c *Join_typeContext)

	// EnterOpt_natural is called when entering the opt_natural production.
	EnterOpt_natural(c *Opt_naturalContext)

	// EnterOn_clause is called when entering the on_clause production.
	EnterOn_clause(c *On_clauseContext)

	// EnterSelect_list is called when entering the select_list production.
	EnterSelect_list(c *Select_listContext)

	// EnterSelect_list_item is called when entering the select_list_item production.
	EnterSelect_list_item(c *Select_list_itemContext)

	// EnterSelect_column_star is called when entering the select_column_star production.
	EnterSelect_column_star(c *Select_column_starContext)

	// EnterSelect_column_expr is called when entering the select_column_expr production.
	EnterSelect_column_expr(c *Select_column_exprContext)

	// EnterSelect_column_dot_star is called when entering the select_column_dot_star production.
	EnterSelect_column_dot_star(c *Select_column_dot_starContext)

	// EnterStar_modifiers is called when entering the star_modifiers production.
	EnterStar_modifiers(c *Star_modifiersContext)

	// EnterStar_except_list is called when entering the star_except_list production.
	EnterStar_except_list(c *Star_except_listContext)

	// EnterStar_replace_list is called when entering the star_replace_list production.
	EnterStar_replace_list(c *Star_replace_listContext)

	// EnterStar_replace_item is called when entering the star_replace_item production.
	EnterStar_replace_item(c *Star_replace_itemContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterExpression_higher_prec_than_and is called when entering the expression_higher_prec_than_and production.
	EnterExpression_higher_prec_than_and(c *Expression_higher_prec_than_andContext)

	// EnterExpression_maybe_parenthesized_not_a_query is called when entering the expression_maybe_parenthesized_not_a_query production.
	EnterExpression_maybe_parenthesized_not_a_query(c *Expression_maybe_parenthesized_not_a_queryContext)

	// EnterParenthesized_in_rhs is called when entering the parenthesized_in_rhs production.
	EnterParenthesized_in_rhs(c *Parenthesized_in_rhsContext)

	// EnterUnary_operator is called when entering the unary_operator production.
	EnterUnary_operator(c *Unary_operatorContext)

	// EnterComparative_operator is called when entering the comparative_operator production.
	EnterComparative_operator(c *Comparative_operatorContext)

	// EnterShift_operator is called when entering the shift_operator production.
	EnterShift_operator(c *Shift_operatorContext)

	// EnterAdditive_operator is called when entering the additive_operator production.
	EnterAdditive_operator(c *Additive_operatorContext)

	// EnterMultiplicative_operator is called when entering the multiplicative_operator production.
	EnterMultiplicative_operator(c *Multiplicative_operatorContext)

	// EnterIs_operator is called when entering the is_operator production.
	EnterIs_operator(c *Is_operatorContext)

	// EnterBetween_operator is called when entering the between_operator production.
	EnterBetween_operator(c *Between_operatorContext)

	// EnterIn_operator is called when entering the in_operator production.
	EnterIn_operator(c *In_operatorContext)

	// EnterDistinct_operator is called when entering the distinct_operator production.
	EnterDistinct_operator(c *Distinct_operatorContext)

	// EnterParenthesized_query is called when entering the parenthesized_query production.
	EnterParenthesized_query(c *Parenthesized_queryContext)

	// EnterParenthesized_expression_not_a_query is called when entering the parenthesized_expression_not_a_query production.
	EnterParenthesized_expression_not_a_query(c *Parenthesized_expression_not_a_queryContext)

	// EnterParenthesized_anysomeall_list_in_rhs is called when entering the parenthesized_anysomeall_list_in_rhs production.
	EnterParenthesized_anysomeall_list_in_rhs(c *Parenthesized_anysomeall_list_in_rhsContext)

	// EnterAnd_expression is called when entering the and_expression production.
	EnterAnd_expression(c *And_expressionContext)

	// EnterIn_list_two_or_more_prefix is called when entering the in_list_two_or_more_prefix production.
	EnterIn_list_two_or_more_prefix(c *In_list_two_or_more_prefixContext)

	// EnterAny_some_all is called when entering the any_some_all production.
	EnterAny_some_all(c *Any_some_allContext)

	// EnterLike_operator is called when entering the like_operator production.
	EnterLike_operator(c *Like_operatorContext)

	// EnterExpression_subquery_with_keyword is called when entering the expression_subquery_with_keyword production.
	EnterExpression_subquery_with_keyword(c *Expression_subquery_with_keywordContext)

	// EnterStruct_constructor is called when entering the struct_constructor production.
	EnterStruct_constructor(c *Struct_constructorContext)

	// EnterStruct_constructor_prefix_with_keyword is called when entering the struct_constructor_prefix_with_keyword production.
	EnterStruct_constructor_prefix_with_keyword(c *Struct_constructor_prefix_with_keywordContext)

	// EnterStruct_constructor_arg is called when entering the struct_constructor_arg production.
	EnterStruct_constructor_arg(c *Struct_constructor_argContext)

	// EnterStruct_constructor_prefix_without_keyword is called when entering the struct_constructor_prefix_without_keyword production.
	EnterStruct_constructor_prefix_without_keyword(c *Struct_constructor_prefix_without_keywordContext)

	// EnterStruct_constructor_prefix_with_keyword_no_arg is called when entering the struct_constructor_prefix_with_keyword_no_arg production.
	EnterStruct_constructor_prefix_with_keyword_no_arg(c *Struct_constructor_prefix_with_keyword_no_argContext)

	// EnterInterval_expression is called when entering the interval_expression production.
	EnterInterval_expression(c *Interval_expressionContext)

	// EnterFunction_call_expression_with_clauses is called when entering the function_call_expression_with_clauses production.
	EnterFunction_call_expression_with_clauses(c *Function_call_expression_with_clausesContext)

	// EnterFunction_call_expression_with_clauses_suffix is called when entering the function_call_expression_with_clauses_suffix production.
	EnterFunction_call_expression_with_clauses_suffix(c *Function_call_expression_with_clauses_suffixContext)

	// EnterOver_clause is called when entering the over_clause production.
	EnterOver_clause(c *Over_clauseContext)

	// EnterWindow_specification is called when entering the window_specification production.
	EnterWindow_specification(c *Window_specificationContext)

	// EnterOpt_window_frame_clause is called when entering the opt_window_frame_clause production.
	EnterOpt_window_frame_clause(c *Opt_window_frame_clauseContext)

	// EnterWindow_frame_bound is called when entering the window_frame_bound production.
	EnterWindow_frame_bound(c *Window_frame_boundContext)

	// EnterPreceding_or_following is called when entering the preceding_or_following production.
	EnterPreceding_or_following(c *Preceding_or_followingContext)

	// EnterFrame_unit is called when entering the frame_unit production.
	EnterFrame_unit(c *Frame_unitContext)

	// EnterPartition_by_clause is called when entering the partition_by_clause production.
	EnterPartition_by_clause(c *Partition_by_clauseContext)

	// EnterPartition_by_clause_prefix is called when entering the partition_by_clause_prefix production.
	EnterPartition_by_clause_prefix(c *Partition_by_clause_prefixContext)

	// EnterWith_group_rows is called when entering the with_group_rows production.
	EnterWith_group_rows(c *With_group_rowsContext)

	// EnterWith_report_modifier is called when entering the with_report_modifier production.
	EnterWith_report_modifier(c *With_report_modifierContext)

	// EnterClamped_between_modifier is called when entering the clamped_between_modifier production.
	EnterClamped_between_modifier(c *Clamped_between_modifierContext)

	// EnterWith_report_format is called when entering the with_report_format production.
	EnterWith_report_format(c *With_report_formatContext)

	// EnterOptions_list is called when entering the options_list production.
	EnterOptions_list(c *Options_listContext)

	// EnterOptions_list_prefix is called when entering the options_list_prefix production.
	EnterOptions_list_prefix(c *Options_list_prefixContext)

	// EnterOptions_entry is called when entering the options_entry production.
	EnterOptions_entry(c *Options_entryContext)

	// EnterExpression_or_proto is called when entering the expression_or_proto production.
	EnterExpression_or_proto(c *Expression_or_protoContext)

	// EnterOptions_assignment_operator is called when entering the options_assignment_operator production.
	EnterOptions_assignment_operator(c *Options_assignment_operatorContext)

	// EnterOpt_null_handling_modifier is called when entering the opt_null_handling_modifier production.
	EnterOpt_null_handling_modifier(c *Opt_null_handling_modifierContext)

	// EnterFunction_call_argument is called when entering the function_call_argument production.
	EnterFunction_call_argument(c *Function_call_argumentContext)

	// EnterSequence_arg is called when entering the sequence_arg production.
	EnterSequence_arg(c *Sequence_argContext)

	// EnterNamed_argument is called when entering the named_argument production.
	EnterNamed_argument(c *Named_argumentContext)

	// EnterLambda_argument is called when entering the lambda_argument production.
	EnterLambda_argument(c *Lambda_argumentContext)

	// EnterLambda_argument_list is called when entering the lambda_argument_list production.
	EnterLambda_argument_list(c *Lambda_argument_listContext)

	// EnterLimit_offset_clause is called when entering the limit_offset_clause production.
	EnterLimit_offset_clause(c *Limit_offset_clauseContext)

	// EnterOpt_having_or_group_by_modifier is called when entering the opt_having_or_group_by_modifier production.
	EnterOpt_having_or_group_by_modifier(c *Opt_having_or_group_by_modifierContext)

	// EnterGroup_by_clause_prefix is called when entering the group_by_clause_prefix production.
	EnterGroup_by_clause_prefix(c *Group_by_clause_prefixContext)

	// EnterGroup_by_preamble is called when entering the group_by_preamble production.
	EnterGroup_by_preamble(c *Group_by_preambleContext)

	// EnterOpt_and_order is called when entering the opt_and_order production.
	EnterOpt_and_order(c *Opt_and_orderContext)

	// EnterHint is called when entering the hint production.
	EnterHint(c *HintContext)

	// EnterHint_with_body is called when entering the hint_with_body production.
	EnterHint_with_body(c *Hint_with_bodyContext)

	// EnterHint_with_body_prefix is called when entering the hint_with_body_prefix production.
	EnterHint_with_body_prefix(c *Hint_with_body_prefixContext)

	// EnterHint_entry is called when entering the hint_entry production.
	EnterHint_entry(c *Hint_entryContext)

	// EnterIdentifier_in_hints is called when entering the identifier_in_hints production.
	EnterIdentifier_in_hints(c *Identifier_in_hintsContext)

	// EnterExtra_identifier_in_hints_name is called when entering the extra_identifier_in_hints_name production.
	EnterExtra_identifier_in_hints_name(c *Extra_identifier_in_hints_nameContext)

	// EnterGrouping_item is called when entering the grouping_item production.
	EnterGrouping_item(c *Grouping_itemContext)

	// EnterGrouping_set_list is called when entering the grouping_set_list production.
	EnterGrouping_set_list(c *Grouping_set_listContext)

	// EnterGrouping_set is called when entering the grouping_set production.
	EnterGrouping_set(c *Grouping_setContext)

	// EnterCube_list is called when entering the cube_list production.
	EnterCube_list(c *Cube_listContext)

	// EnterRollup_list is called when entering the rollup_list production.
	EnterRollup_list(c *Rollup_listContext)

	// EnterOpt_as_alias_with_required_as is called when entering the opt_as_alias_with_required_as production.
	EnterOpt_as_alias_with_required_as(c *Opt_as_alias_with_required_asContext)

	// EnterOpt_grouping_item_order is called when entering the opt_grouping_item_order production.
	EnterOpt_grouping_item_order(c *Opt_grouping_item_orderContext)

	// EnterOpt_selection_item_order is called when entering the opt_selection_item_order production.
	EnterOpt_selection_item_order(c *Opt_selection_item_orderContext)

	// EnterAsc_or_desc is called when entering the asc_or_desc production.
	EnterAsc_or_desc(c *Asc_or_descContext)

	// EnterNull_order is called when entering the null_order production.
	EnterNull_order(c *Null_orderContext)

	// EnterFunction_name_from_keyword is called when entering the function_name_from_keyword production.
	EnterFunction_name_from_keyword(c *Function_name_from_keywordContext)

	// EnterReplace_fields_expression is called when entering the replace_fields_expression production.
	EnterReplace_fields_expression(c *Replace_fields_expressionContext)

	// EnterReplace_fields_prefix is called when entering the replace_fields_prefix production.
	EnterReplace_fields_prefix(c *Replace_fields_prefixContext)

	// EnterReplace_fields_arg is called when entering the replace_fields_arg production.
	EnterReplace_fields_arg(c *Replace_fields_argContext)

	// EnterGeneralized_path_expression is called when entering the generalized_path_expression production.
	EnterGeneralized_path_expression(c *Generalized_path_expressionContext)

	// EnterGeneralized_extension_path is called when entering the generalized_extension_path production.
	EnterGeneralized_extension_path(c *Generalized_extension_pathContext)

	// EnterWith_expression is called when entering the with_expression production.
	EnterWith_expression(c *With_expressionContext)

	// EnterWith_expression_variable_prefix is called when entering the with_expression_variable_prefix production.
	EnterWith_expression_variable_prefix(c *With_expression_variable_prefixContext)

	// EnterWith_expression_variable is called when entering the with_expression_variable production.
	EnterWith_expression_variable(c *With_expression_variableContext)

	// EnterExtract_expression is called when entering the extract_expression production.
	EnterExtract_expression(c *Extract_expressionContext)

	// EnterExtract_expression_base is called when entering the extract_expression_base production.
	EnterExtract_expression_base(c *Extract_expression_baseContext)

	// EnterOpt_format is called when entering the opt_format production.
	EnterOpt_format(c *Opt_formatContext)

	// EnterOpt_at_time_zone is called when entering the opt_at_time_zone production.
	EnterOpt_at_time_zone(c *Opt_at_time_zoneContext)

	// EnterCast_expression is called when entering the cast_expression production.
	EnterCast_expression(c *Cast_expressionContext)

	// EnterCase_expression is called when entering the case_expression production.
	EnterCase_expression(c *Case_expressionContext)

	// EnterCase_expression_prefix is called when entering the case_expression_prefix production.
	EnterCase_expression_prefix(c *Case_expression_prefixContext)

	// EnterCase_value_expression_prefix is called when entering the case_value_expression_prefix production.
	EnterCase_value_expression_prefix(c *Case_value_expression_prefixContext)

	// EnterCase_no_value_expression_prefix is called when entering the case_no_value_expression_prefix production.
	EnterCase_no_value_expression_prefix(c *Case_no_value_expression_prefixContext)

	// EnterStruct_braced_constructor is called when entering the struct_braced_constructor production.
	EnterStruct_braced_constructor(c *Struct_braced_constructorContext)

	// EnterBraced_new_constructor is called when entering the braced_new_constructor production.
	EnterBraced_new_constructor(c *Braced_new_constructorContext)

	// EnterBraced_constructor is called when entering the braced_constructor production.
	EnterBraced_constructor(c *Braced_constructorContext)

	// EnterBraced_constructor_start is called when entering the braced_constructor_start production.
	EnterBraced_constructor_start(c *Braced_constructor_startContext)

	// EnterBraced_constructor_prefix is called when entering the braced_constructor_prefix production.
	EnterBraced_constructor_prefix(c *Braced_constructor_prefixContext)

	// EnterBraced_constructor_field is called when entering the braced_constructor_field production.
	EnterBraced_constructor_field(c *Braced_constructor_fieldContext)

	// EnterBraced_constructor_lhs is called when entering the braced_constructor_lhs production.
	EnterBraced_constructor_lhs(c *Braced_constructor_lhsContext)

	// EnterBraced_constructor_field_value is called when entering the braced_constructor_field_value production.
	EnterBraced_constructor_field_value(c *Braced_constructor_field_valueContext)

	// EnterBraced_constructor_extension is called when entering the braced_constructor_extension production.
	EnterBraced_constructor_extension(c *Braced_constructor_extensionContext)

	// EnterNew_constructor is called when entering the new_constructor production.
	EnterNew_constructor(c *New_constructorContext)

	// EnterNew_constructor_prefix is called when entering the new_constructor_prefix production.
	EnterNew_constructor_prefix(c *New_constructor_prefixContext)

	// EnterNew_constructor_prefix_no_arg is called when entering the new_constructor_prefix_no_arg production.
	EnterNew_constructor_prefix_no_arg(c *New_constructor_prefix_no_argContext)

	// EnterNew_constructor_arg is called when entering the new_constructor_arg production.
	EnterNew_constructor_arg(c *New_constructor_argContext)

	// EnterArray_constructor is called when entering the array_constructor production.
	EnterArray_constructor(c *Array_constructorContext)

	// EnterArray_constructor_prefix is called when entering the array_constructor_prefix production.
	EnterArray_constructor_prefix(c *Array_constructor_prefixContext)

	// EnterArray_constructor_prefix_no_expressions is called when entering the array_constructor_prefix_no_expressions production.
	EnterArray_constructor_prefix_no_expressions(c *Array_constructor_prefix_no_expressionsContext)

	// EnterRange_literal is called when entering the range_literal production.
	EnterRange_literal(c *Range_literalContext)

	// EnterRange_type is called when entering the range_type production.
	EnterRange_type(c *Range_typeContext)

	// EnterType is called when entering the type production.
	EnterType(c *TypeContext)

	// EnterCollate_clause is called when entering the collate_clause production.
	EnterCollate_clause(c *Collate_clauseContext)

	// EnterString_literal_or_parameter is called when entering the string_literal_or_parameter production.
	EnterString_literal_or_parameter(c *String_literal_or_parameterContext)

	// EnterSystem_variable_expression is called when entering the system_variable_expression production.
	EnterSystem_variable_expression(c *System_variable_expressionContext)

	// EnterParameter_expression is called when entering the parameter_expression production.
	EnterParameter_expression(c *Parameter_expressionContext)

	// EnterNamed_parameter_expression is called when entering the named_parameter_expression production.
	EnterNamed_parameter_expression(c *Named_parameter_expressionContext)

	// EnterOpt_type_parameters is called when entering the opt_type_parameters production.
	EnterOpt_type_parameters(c *Opt_type_parametersContext)

	// EnterType_parameters_prefix is called when entering the type_parameters_prefix production.
	EnterType_parameters_prefix(c *Type_parameters_prefixContext)

	// EnterType_parameter is called when entering the type_parameter production.
	EnterType_parameter(c *Type_parameterContext)

	// EnterRaw_type is called when entering the raw_type production.
	EnterRaw_type(c *Raw_typeContext)

	// EnterMap_type is called when entering the map_type production.
	EnterMap_type(c *Map_typeContext)

	// EnterFunction_type is called when entering the function_type production.
	EnterFunction_type(c *Function_typeContext)

	// EnterFunction_type_prefix is called when entering the function_type_prefix production.
	EnterFunction_type_prefix(c *Function_type_prefixContext)

	// EnterType_name is called when entering the type_name production.
	EnterType_name(c *Type_nameContext)

	// EnterPath_expression is called when entering the path_expression production.
	EnterPath_expression(c *Path_expressionContext)

	// EnterIdentifier is called when entering the identifier production.
	EnterIdentifier(c *IdentifierContext)

	// EnterKeyword_as_identifier is called when entering the keyword_as_identifier production.
	EnterKeyword_as_identifier(c *Keyword_as_identifierContext)

	// EnterCommon_keyword_as_identifier is called when entering the common_keyword_as_identifier production.
	EnterCommon_keyword_as_identifier(c *Common_keyword_as_identifierContext)

	// EnterToken_identifier is called when entering the token_identifier production.
	EnterToken_identifier(c *Token_identifierContext)

	// EnterStruct_type is called when entering the struct_type production.
	EnterStruct_type(c *Struct_typeContext)

	// EnterStruct_type_prefix is called when entering the struct_type_prefix production.
	EnterStruct_type_prefix(c *Struct_type_prefixContext)

	// EnterStruct_field is called when entering the struct_field production.
	EnterStruct_field(c *Struct_fieldContext)

	// EnterArray_type is called when entering the array_type production.
	EnterArray_type(c *Array_typeContext)

	// EnterTemplate_type_open is called when entering the template_type_open production.
	EnterTemplate_type_open(c *Template_type_openContext)

	// EnterTemplate_type_close is called when entering the template_type_close production.
	EnterTemplate_type_close(c *Template_type_closeContext)

	// EnterDate_or_time_literal is called when entering the date_or_time_literal production.
	EnterDate_or_time_literal(c *Date_or_time_literalContext)

	// EnterDate_or_time_literal_kind is called when entering the date_or_time_literal_kind production.
	EnterDate_or_time_literal_kind(c *Date_or_time_literal_kindContext)

	// EnterFloating_point_literal is called when entering the floating_point_literal production.
	EnterFloating_point_literal(c *Floating_point_literalContext)

	// EnterJson_literal is called when entering the json_literal production.
	EnterJson_literal(c *Json_literalContext)

	// EnterBignumeric_literal is called when entering the bignumeric_literal production.
	EnterBignumeric_literal(c *Bignumeric_literalContext)

	// EnterBignumeric_literal_prefix is called when entering the bignumeric_literal_prefix production.
	EnterBignumeric_literal_prefix(c *Bignumeric_literal_prefixContext)

	// EnterNumeric_literal is called when entering the numeric_literal production.
	EnterNumeric_literal(c *Numeric_literalContext)

	// EnterNumeric_literal_prefix is called when entering the numeric_literal_prefix production.
	EnterNumeric_literal_prefix(c *Numeric_literal_prefixContext)

	// EnterInteger_literal is called when entering the integer_literal production.
	EnterInteger_literal(c *Integer_literalContext)

	// EnterBytes_literal is called when entering the bytes_literal production.
	EnterBytes_literal(c *Bytes_literalContext)

	// EnterNull_literal is called when entering the null_literal production.
	EnterNull_literal(c *Null_literalContext)

	// EnterBoolean_literal is called when entering the boolean_literal production.
	EnterBoolean_literal(c *Boolean_literalContext)

	// EnterString_literal is called when entering the string_literal production.
	EnterString_literal(c *String_literalContext)

	// EnterString_literal_component is called when entering the string_literal_component production.
	EnterString_literal_component(c *String_literal_componentContext)

	// EnterBytes_literal_component is called when entering the bytes_literal_component production.
	EnterBytes_literal_component(c *Bytes_literal_componentContext)

	// ExitRoot is called when exiting the root production.
	ExitRoot(c *RootContext)

	// ExitStmts is called when exiting the stmts production.
	ExitStmts(c *StmtsContext)

	// ExitUnterminated_sql_statement is called when exiting the unterminated_sql_statement production.
	ExitUnterminated_sql_statement(c *Unterminated_sql_statementContext)

	// ExitSql_statement_body is called when exiting the sql_statement_body production.
	ExitSql_statement_body(c *Sql_statement_bodyContext)

	// ExitGql_statement is called when exiting the gql_statement production.
	ExitGql_statement(c *Gql_statementContext)

	// ExitGraph_operation_block is called when exiting the graph_operation_block production.
	ExitGraph_operation_block(c *Graph_operation_blockContext)

	// ExitGraph_composite_query_block is called when exiting the graph_composite_query_block production.
	ExitGraph_composite_query_block(c *Graph_composite_query_blockContext)

	// ExitGraph_composite_query_prefix is called when exiting the graph_composite_query_prefix production.
	ExitGraph_composite_query_prefix(c *Graph_composite_query_prefixContext)

	// ExitGraph_set_operation_metadata is called when exiting the graph_set_operation_metadata production.
	ExitGraph_set_operation_metadata(c *Graph_set_operation_metadataContext)

	// ExitGraph_linear_query_operation is called when exiting the graph_linear_query_operation production.
	ExitGraph_linear_query_operation(c *Graph_linear_query_operationContext)

	// ExitGraph_linear_operator_list is called when exiting the graph_linear_operator_list production.
	ExitGraph_linear_operator_list(c *Graph_linear_operator_listContext)

	// ExitGraph_linear_operator is called when exiting the graph_linear_operator production.
	ExitGraph_linear_operator(c *Graph_linear_operatorContext)

	// ExitGraph_sample_clause is called when exiting the graph_sample_clause production.
	ExitGraph_sample_clause(c *Graph_sample_clauseContext)

	// ExitOpt_graph_sample_clause_suffix is called when exiting the opt_graph_sample_clause_suffix production.
	ExitOpt_graph_sample_clause_suffix(c *Opt_graph_sample_clause_suffixContext)

	// ExitGraph_for_operator is called when exiting the graph_for_operator production.
	ExitGraph_for_operator(c *Graph_for_operatorContext)

	// ExitOpt_with_offset_and_alias_with_required_as is called when exiting the opt_with_offset_and_alias_with_required_as production.
	ExitOpt_with_offset_and_alias_with_required_as(c *Opt_with_offset_and_alias_with_required_asContext)

	// ExitGraph_with_operator is called when exiting the graph_with_operator production.
	ExitGraph_with_operator(c *Graph_with_operatorContext)

	// ExitGraph_page_operator is called when exiting the graph_page_operator production.
	ExitGraph_page_operator(c *Graph_page_operatorContext)

	// ExitGraph_order_by_operator is called when exiting the graph_order_by_operator production.
	ExitGraph_order_by_operator(c *Graph_order_by_operatorContext)

	// ExitGraph_filter_operator is called when exiting the graph_filter_operator production.
	ExitGraph_filter_operator(c *Graph_filter_operatorContext)

	// ExitGraph_let_operator is called when exiting the graph_let_operator production.
	ExitGraph_let_operator(c *Graph_let_operatorContext)

	// ExitGraph_let_variable_definition_list is called when exiting the graph_let_variable_definition_list production.
	ExitGraph_let_variable_definition_list(c *Graph_let_variable_definition_listContext)

	// ExitGraph_let_variable_definition is called when exiting the graph_let_variable_definition production.
	ExitGraph_let_variable_definition(c *Graph_let_variable_definitionContext)

	// ExitGraph_optional_match_operator is called when exiting the graph_optional_match_operator production.
	ExitGraph_optional_match_operator(c *Graph_optional_match_operatorContext)

	// ExitGraph_match_operator is called when exiting the graph_match_operator production.
	ExitGraph_match_operator(c *Graph_match_operatorContext)

	// ExitGraph_pattern is called when exiting the graph_pattern production.
	ExitGraph_pattern(c *Graph_patternContext)

	// ExitGraph_path_pattern_list is called when exiting the graph_path_pattern_list production.
	ExitGraph_path_pattern_list(c *Graph_path_pattern_listContext)

	// ExitGraph_path_pattern is called when exiting the graph_path_pattern production.
	ExitGraph_path_pattern(c *Graph_path_patternContext)

	// ExitGraph_path_pattern_expr is called when exiting the graph_path_pattern_expr production.
	ExitGraph_path_pattern_expr(c *Graph_path_pattern_exprContext)

	// ExitGraph_path_factor is called when exiting the graph_path_factor production.
	ExitGraph_path_factor(c *Graph_path_factorContext)

	// ExitGraph_quantified_path_primary is called when exiting the graph_quantified_path_primary production.
	ExitGraph_quantified_path_primary(c *Graph_quantified_path_primaryContext)

	// ExitGraph_path_primary is called when exiting the graph_path_primary production.
	ExitGraph_path_primary(c *Graph_path_primaryContext)

	// ExitGraph_parenthesized_path_pattern is called when exiting the graph_parenthesized_path_pattern production.
	ExitGraph_parenthesized_path_pattern(c *Graph_parenthesized_path_patternContext)

	// ExitGraph_element_pattern is called when exiting the graph_element_pattern production.
	ExitGraph_element_pattern(c *Graph_element_patternContext)

	// ExitGraph_edge_pattern is called when exiting the graph_edge_pattern production.
	ExitGraph_edge_pattern(c *Graph_edge_patternContext)

	// ExitGraph_node_pattern is called when exiting the graph_node_pattern production.
	ExitGraph_node_pattern(c *Graph_node_patternContext)

	// ExitGraph_element_pattern_filler is called when exiting the graph_element_pattern_filler production.
	ExitGraph_element_pattern_filler(c *Graph_element_pattern_fillerContext)

	// ExitGraph_property_specification is called when exiting the graph_property_specification production.
	ExitGraph_property_specification(c *Graph_property_specificationContext)

	// ExitGraph_property_name_and_value is called when exiting the graph_property_name_and_value production.
	ExitGraph_property_name_and_value(c *Graph_property_name_and_valueContext)

	// ExitOpt_is_label_expression is called when exiting the opt_is_label_expression production.
	ExitOpt_is_label_expression(c *Opt_is_label_expressionContext)

	// ExitLabel_expression is called when exiting the label_expression production.
	ExitLabel_expression(c *Label_expressionContext)

	// ExitLabel_primary is called when exiting the label_primary production.
	ExitLabel_primary(c *Label_primaryContext)

	// ExitParenthesized_label_expression is called when exiting the parenthesized_label_expression production.
	ExitParenthesized_label_expression(c *Parenthesized_label_expressionContext)

	// ExitOpt_graph_element_identifier is called when exiting the opt_graph_element_identifier production.
	ExitOpt_graph_element_identifier(c *Opt_graph_element_identifierContext)

	// ExitOpt_graph_path_mode_prefix is called when exiting the opt_graph_path_mode_prefix production.
	ExitOpt_graph_path_mode_prefix(c *Opt_graph_path_mode_prefixContext)

	// ExitPath_or_paths is called when exiting the path_or_paths production.
	ExitPath_or_paths(c *Path_or_pathsContext)

	// ExitOpt_graph_path_mode is called when exiting the opt_graph_path_mode production.
	ExitOpt_graph_path_mode(c *Opt_graph_path_modeContext)

	// ExitOpt_graph_search_prefix is called when exiting the opt_graph_search_prefix production.
	ExitOpt_graph_search_prefix(c *Opt_graph_search_prefixContext)

	// ExitOpt_path_variable_assignment is called when exiting the opt_path_variable_assignment production.
	ExitOpt_path_variable_assignment(c *Opt_path_variable_assignmentContext)

	// ExitGraph_identifier is called when exiting the graph_identifier production.
	ExitGraph_identifier(c *Graph_identifierContext)

	// ExitGraph_return_operator is called when exiting the graph_return_operator production.
	ExitGraph_return_operator(c *Graph_return_operatorContext)

	// ExitGraph_page_clause is called when exiting the graph_page_clause production.
	ExitGraph_page_clause(c *Graph_page_clauseContext)

	// ExitGraph_order_by_clause is called when exiting the graph_order_by_clause production.
	ExitGraph_order_by_clause(c *Graph_order_by_clauseContext)

	// ExitGraph_ordering_expression is called when exiting the graph_ordering_expression production.
	ExitGraph_ordering_expression(c *Graph_ordering_expressionContext)

	// ExitOpt_graph_asc_or_desc is called when exiting the opt_graph_asc_or_desc production.
	ExitOpt_graph_asc_or_desc(c *Opt_graph_asc_or_descContext)

	// ExitGraph_return_item_list is called when exiting the graph_return_item_list production.
	ExitGraph_return_item_list(c *Graph_return_item_listContext)

	// ExitGraph_return_item is called when exiting the graph_return_item production.
	ExitGraph_return_item(c *Graph_return_itemContext)

	// ExitUndrop_statement is called when exiting the undrop_statement production.
	ExitUndrop_statement(c *Undrop_statementContext)

	// ExitModule_statement is called when exiting the module_statement production.
	ExitModule_statement(c *Module_statementContext)

	// ExitImport_statement is called when exiting the import_statement production.
	ExitImport_statement(c *Import_statementContext)

	// ExitOpt_as_or_into_alias is called when exiting the opt_as_or_into_alias production.
	ExitOpt_as_or_into_alias(c *Opt_as_or_into_aliasContext)

	// ExitPath_expression_or_string is called when exiting the path_expression_or_string production.
	ExitPath_expression_or_string(c *Path_expression_or_stringContext)

	// ExitImport_type is called when exiting the import_type production.
	ExitImport_type(c *Import_typeContext)

	// ExitCall_statement is called when exiting the call_statement production.
	ExitCall_statement(c *Call_statementContext)

	// ExitDrop_statement is called when exiting the drop_statement production.
	ExitDrop_statement(c *Drop_statementContext)

	// ExitOpt_drop_mode is called when exiting the opt_drop_mode production.
	ExitOpt_drop_mode(c *Opt_drop_modeContext)

	// ExitDrop_all_row_access_policies_statement is called when exiting the drop_all_row_access_policies_statement production.
	ExitDrop_all_row_access_policies_statement(c *Drop_all_row_access_policies_statementContext)

	// ExitShow_statement is called when exiting the show_statement production.
	ExitShow_statement(c *Show_statementContext)

	// ExitOpt_like_string_literal is called when exiting the opt_like_string_literal production.
	ExitOpt_like_string_literal(c *Opt_like_string_literalContext)

	// ExitShow_target is called when exiting the show_target production.
	ExitShow_target(c *Show_targetContext)

	// ExitRename_statement is called when exiting the rename_statement production.
	ExitRename_statement(c *Rename_statementContext)

	// ExitRevoke_statement is called when exiting the revoke_statement production.
	ExitRevoke_statement(c *Revoke_statementContext)

	// ExitGrant_statement is called when exiting the grant_statement production.
	ExitGrant_statement(c *Grant_statementContext)

	// ExitPrivileges is called when exiting the privileges production.
	ExitPrivileges(c *PrivilegesContext)

	// ExitExport_metadata_statement is called when exiting the export_metadata_statement production.
	ExitExport_metadata_statement(c *Export_metadata_statementContext)

	// ExitExport_model_statement is called when exiting the export_model_statement production.
	ExitExport_model_statement(c *Export_model_statementContext)

	// ExitExport_data_statement is called when exiting the export_data_statement production.
	ExitExport_data_statement(c *Export_data_statementContext)

	// ExitExport_data_no_query is called when exiting the export_data_no_query production.
	ExitExport_data_no_query(c *Export_data_no_queryContext)

	// ExitExplain_statement is called when exiting the explain_statement production.
	ExitExplain_statement(c *Explain_statementContext)

	// ExitExecute_immediate is called when exiting the execute_immediate production.
	ExitExecute_immediate(c *Execute_immediateContext)

	// ExitOpt_execute_into_clause is called when exiting the opt_execute_into_clause production.
	ExitOpt_execute_into_clause(c *Opt_execute_into_clauseContext)

	// ExitOpt_execute_using_clause is called when exiting the opt_execute_using_clause production.
	ExitOpt_execute_using_clause(c *Opt_execute_using_clauseContext)

	// ExitExecute_using_argument_list is called when exiting the execute_using_argument_list production.
	ExitExecute_using_argument_list(c *Execute_using_argument_listContext)

	// ExitExecute_using_argument is called when exiting the execute_using_argument production.
	ExitExecute_using_argument(c *Execute_using_argumentContext)

	// ExitDescribe_statement is called when exiting the describe_statement production.
	ExitDescribe_statement(c *Describe_statementContext)

	// ExitDescribe_info is called when exiting the describe_info production.
	ExitDescribe_info(c *Describe_infoContext)

	// ExitOpt_from_path_expression is called when exiting the opt_from_path_expression production.
	ExitOpt_from_path_expression(c *Opt_from_path_expressionContext)

	// ExitDescribe_keyword is called when exiting the describe_keyword production.
	ExitDescribe_keyword(c *Describe_keywordContext)

	// ExitDefine_table_statement is called when exiting the define_table_statement production.
	ExitDefine_table_statement(c *Define_table_statementContext)

	// ExitCreate_entity_statement is called when exiting the create_entity_statement production.
	ExitCreate_entity_statement(c *Create_entity_statementContext)

	// ExitOpt_generic_entity_body is called when exiting the opt_generic_entity_body production.
	ExitOpt_generic_entity_body(c *Opt_generic_entity_bodyContext)

	// ExitCreate_view_statement is called when exiting the create_view_statement production.
	ExitCreate_view_statement(c *Create_view_statementContext)

	// ExitQuery_or_replica_source is called when exiting the query_or_replica_source production.
	ExitQuery_or_replica_source(c *Query_or_replica_sourceContext)

	// ExitColumn_with_options_list is called when exiting the column_with_options_list production.
	ExitColumn_with_options_list(c *Column_with_options_listContext)

	// ExitColumn_with_options is called when exiting the column_with_options production.
	ExitColumn_with_options(c *Column_with_optionsContext)

	// ExitCreate_table_statement is called when exiting the create_table_statement production.
	ExitCreate_table_statement(c *Create_table_statementContext)

	// ExitOpt_ttl_clause is called when exiting the opt_ttl_clause production.
	ExitOpt_ttl_clause(c *Opt_ttl_clauseContext)

	// ExitOpt_copy_table is called when exiting the opt_copy_table production.
	ExitOpt_copy_table(c *Opt_copy_tableContext)

	// ExitCopy_data_source is called when exiting the copy_data_source production.
	ExitCopy_data_source(c *Copy_data_sourceContext)

	// ExitOpt_clone_table is called when exiting the opt_clone_table production.
	ExitOpt_clone_table(c *Opt_clone_tableContext)

	// ExitOpt_spanner_table_options is called when exiting the opt_spanner_table_options production.
	ExitOpt_spanner_table_options(c *Opt_spanner_table_optionsContext)

	// ExitOpt_spanner_interleave_in_parent_clause is called when exiting the opt_spanner_interleave_in_parent_clause production.
	ExitOpt_spanner_interleave_in_parent_clause(c *Opt_spanner_interleave_in_parent_clauseContext)

	// ExitSpanner_primary_key is called when exiting the spanner_primary_key production.
	ExitSpanner_primary_key(c *Spanner_primary_keyContext)

	// ExitCreate_table_function_statement is called when exiting the create_table_function_statement production.
	ExitCreate_table_function_statement(c *Create_table_function_statementContext)

	// ExitOpt_as_query_or_string is called when exiting the opt_as_query_or_string production.
	ExitOpt_as_query_or_string(c *Opt_as_query_or_stringContext)

	// ExitUnordered_language_options is called when exiting the unordered_language_options production.
	ExitUnordered_language_options(c *Unordered_language_optionsContext)

	// ExitOpt_function_parameters is called when exiting the opt_function_parameters production.
	ExitOpt_function_parameters(c *Opt_function_parametersContext)

	// ExitCreate_snapshot_statement is called when exiting the create_snapshot_statement production.
	ExitCreate_snapshot_statement(c *Create_snapshot_statementContext)

	// ExitCreate_external_schema_statement is called when exiting the create_external_schema_statement production.
	ExitCreate_external_schema_statement(c *Create_external_schema_statementContext)

	// ExitCreate_schema_statement is called when exiting the create_schema_statement production.
	ExitCreate_schema_statement(c *Create_schema_statementContext)

	// ExitCreate_property_graph_statement is called when exiting the create_property_graph_statement production.
	ExitCreate_property_graph_statement(c *Create_property_graph_statementContext)

	// ExitOpt_edge_table_clause is called when exiting the opt_edge_table_clause production.
	ExitOpt_edge_table_clause(c *Opt_edge_table_clauseContext)

	// ExitElement_table_list is called when exiting the element_table_list production.
	ExitElement_table_list(c *Element_table_listContext)

	// ExitElement_table_definition is called when exiting the element_table_definition production.
	ExitElement_table_definition(c *Element_table_definitionContext)

	// ExitOpt_label_and_properties_clause is called when exiting the opt_label_and_properties_clause production.
	ExitOpt_label_and_properties_clause(c *Opt_label_and_properties_clauseContext)

	// ExitLabel_and_properties_list is called when exiting the label_and_properties_list production.
	ExitLabel_and_properties_list(c *Label_and_properties_listContext)

	// ExitLabel_and_properties is called when exiting the label_and_properties production.
	ExitLabel_and_properties(c *Label_and_propertiesContext)

	// ExitProperties_clause is called when exiting the properties_clause production.
	ExitProperties_clause(c *Properties_clauseContext)

	// ExitDerived_property_list is called when exiting the derived_property_list production.
	ExitDerived_property_list(c *Derived_property_listContext)

	// ExitDerived_property is called when exiting the derived_property production.
	ExitDerived_property(c *Derived_propertyContext)

	// ExitOpt_except_column_list is called when exiting the opt_except_column_list production.
	ExitOpt_except_column_list(c *Opt_except_column_listContext)

	// ExitProperties_all_columns is called when exiting the properties_all_columns production.
	ExitProperties_all_columns(c *Properties_all_columnsContext)

	// ExitOpt_dest_node_table_clause is called when exiting the opt_dest_node_table_clause production.
	ExitOpt_dest_node_table_clause(c *Opt_dest_node_table_clauseContext)

	// ExitOpt_source_node_table_clause is called when exiting the opt_source_node_table_clause production.
	ExitOpt_source_node_table_clause(c *Opt_source_node_table_clauseContext)

	// ExitOpt_key_clause is called when exiting the opt_key_clause production.
	ExitOpt_key_clause(c *Opt_key_clauseContext)

	// ExitCreate_model_statement is called when exiting the create_model_statement production.
	ExitCreate_model_statement(c *Create_model_statementContext)

	// ExitOpt_input_output_clause is called when exiting the opt_input_output_clause production.
	ExitOpt_input_output_clause(c *Opt_input_output_clauseContext)

	// ExitOpt_transform_clause is called when exiting the opt_transform_clause production.
	ExitOpt_transform_clause(c *Opt_transform_clauseContext)

	// ExitOpt_as_query_or_aliased_query_list is called when exiting the opt_as_query_or_aliased_query_list production.
	ExitOpt_as_query_or_aliased_query_list(c *Opt_as_query_or_aliased_query_listContext)

	// ExitAliased_query_list is called when exiting the aliased_query_list production.
	ExitAliased_query_list(c *Aliased_query_listContext)

	// ExitAs_query is called when exiting the as_query production.
	ExitAs_query(c *As_queryContext)

	// ExitCreate_external_table_function_statement is called when exiting the create_external_table_function_statement production.
	ExitCreate_external_table_function_statement(c *Create_external_table_function_statementContext)

	// ExitCreate_external_table_statement is called when exiting the create_external_table_statement production.
	ExitCreate_external_table_statement(c *Create_external_table_statementContext)

	// ExitOpt_default_collate_clause is called when exiting the opt_default_collate_clause production.
	ExitOpt_default_collate_clause(c *Opt_default_collate_clauseContext)

	// ExitOpt_like_path_expression is called when exiting the opt_like_path_expression production.
	ExitOpt_like_path_expression(c *Opt_like_path_expressionContext)

	// ExitCreate_row_access_policy_statement is called when exiting the create_row_access_policy_statement production.
	ExitCreate_row_access_policy_statement(c *Create_row_access_policy_statementContext)

	// ExitFilter_using_clause is called when exiting the filter_using_clause production.
	ExitFilter_using_clause(c *Filter_using_clauseContext)

	// ExitCreate_row_access_policy_grant_to_clause is called when exiting the create_row_access_policy_grant_to_clause production.
	ExitCreate_row_access_policy_grant_to_clause(c *Create_row_access_policy_grant_to_clauseContext)

	// ExitCreate_privilege_restriction_statement is called when exiting the create_privilege_restriction_statement production.
	ExitCreate_privilege_restriction_statement(c *Create_privilege_restriction_statementContext)

	// ExitRestrict_to_clause is called when exiting the restrict_to_clause production.
	ExitRestrict_to_clause(c *Restrict_to_clauseContext)

	// ExitPossibly_empty_grantee_list is called when exiting the possibly_empty_grantee_list production.
	ExitPossibly_empty_grantee_list(c *Possibly_empty_grantee_listContext)

	// ExitCreate_index_statement is called when exiting the create_index_statement production.
	ExitCreate_index_statement(c *Create_index_statementContext)

	// ExitOpt_create_index_statement_suffix is called when exiting the opt_create_index_statement_suffix production.
	ExitOpt_create_index_statement_suffix(c *Opt_create_index_statement_suffixContext)

	// ExitSpanner_index_interleave_clause is called when exiting the spanner_index_interleave_clause production.
	ExitSpanner_index_interleave_clause(c *Spanner_index_interleave_clauseContext)

	// ExitIndex_storing_list is called when exiting the index_storing_list production.
	ExitIndex_storing_list(c *Index_storing_listContext)

	// ExitIndex_storing_expression_list is called when exiting the index_storing_expression_list production.
	ExitIndex_storing_expression_list(c *Index_storing_expression_listContext)

	// ExitIndex_order_by_and_options is called when exiting the index_order_by_and_options production.
	ExitIndex_order_by_and_options(c *Index_order_by_and_optionsContext)

	// ExitIndex_all_columns is called when exiting the index_all_columns production.
	ExitIndex_all_columns(c *Index_all_columnsContext)

	// ExitOpt_with_column_options is called when exiting the opt_with_column_options production.
	ExitOpt_with_column_options(c *Opt_with_column_optionsContext)

	// ExitAll_column_column_options is called when exiting the all_column_column_options production.
	ExitAll_column_column_options(c *All_column_column_optionsContext)

	// ExitColumn_ordering_and_options_expr is called when exiting the column_ordering_and_options_expr production.
	ExitColumn_ordering_and_options_expr(c *Column_ordering_and_options_exprContext)

	// ExitIndex_unnest_expression_list is called when exiting the index_unnest_expression_list production.
	ExitIndex_unnest_expression_list(c *Index_unnest_expression_listContext)

	// ExitUnnest_expression_with_opt_alias_and_offset is called when exiting the unnest_expression_with_opt_alias_and_offset production.
	ExitUnnest_expression_with_opt_alias_and_offset(c *Unnest_expression_with_opt_alias_and_offsetContext)

	// ExitOn_path_expression is called when exiting the on_path_expression production.
	ExitOn_path_expression(c *On_path_expressionContext)

	// ExitIndex_type is called when exiting the index_type production.
	ExitIndex_type(c *Index_typeContext)

	// ExitOpt_spanner_null_filtered is called when exiting the opt_spanner_null_filtered production.
	ExitOpt_spanner_null_filtered(c *Opt_spanner_null_filteredContext)

	// ExitCreate_procedure_statement is called when exiting the create_procedure_statement production.
	ExitCreate_procedure_statement(c *Create_procedure_statementContext)

	// ExitBegin_end_block_or_language_as_code is called when exiting the begin_end_block_or_language_as_code production.
	ExitBegin_end_block_or_language_as_code(c *Begin_end_block_or_language_as_codeContext)

	// ExitBegin_end_block is called when exiting the begin_end_block production.
	ExitBegin_end_block(c *Begin_end_blockContext)

	// ExitOpt_exception_handler is called when exiting the opt_exception_handler production.
	ExitOpt_exception_handler(c *Opt_exception_handlerContext)

	// ExitStatement_list is called when exiting the statement_list production.
	ExitStatement_list(c *Statement_listContext)

	// ExitUnterminated_non_empty_statement_list is called when exiting the unterminated_non_empty_statement_list production.
	ExitUnterminated_non_empty_statement_list(c *Unterminated_non_empty_statement_listContext)

	// ExitUnterminated_statement is called when exiting the unterminated_statement production.
	ExitUnterminated_statement(c *Unterminated_statementContext)

	// ExitUnterminated_script_statement is called when exiting the unterminated_script_statement production.
	ExitUnterminated_script_statement(c *Unterminated_script_statementContext)

	// ExitLabel is called when exiting the label production.
	ExitLabel(c *LabelContext)

	// ExitUnterminated_unlabeled_script_statement is called when exiting the unterminated_unlabeled_script_statement production.
	ExitUnterminated_unlabeled_script_statement(c *Unterminated_unlabeled_script_statementContext)

	// ExitFor_in_statement is called when exiting the for_in_statement production.
	ExitFor_in_statement(c *For_in_statementContext)

	// ExitRepeat_statement is called when exiting the repeat_statement production.
	ExitRepeat_statement(c *Repeat_statementContext)

	// ExitUntil_clause is called when exiting the until_clause production.
	ExitUntil_clause(c *Until_clauseContext)

	// ExitLoop_statement is called when exiting the loop_statement production.
	ExitLoop_statement(c *Loop_statementContext)

	// ExitWhile_statement is called when exiting the while_statement production.
	ExitWhile_statement(c *While_statementContext)

	// ExitRaise_statement is called when exiting the raise_statement production.
	ExitRaise_statement(c *Raise_statementContext)

	// ExitReturn_statement is called when exiting the return_statement production.
	ExitReturn_statement(c *Return_statementContext)

	// ExitContinue_statement is called when exiting the continue_statement production.
	ExitContinue_statement(c *Continue_statementContext)

	// ExitVariable_declaration is called when exiting the variable_declaration production.
	ExitVariable_declaration(c *Variable_declarationContext)

	// ExitBreak_statement is called when exiting the break_statement production.
	ExitBreak_statement(c *Break_statementContext)

	// ExitCase_statement is called when exiting the case_statement production.
	ExitCase_statement(c *Case_statementContext)

	// ExitWhen_then_clauses is called when exiting the when_then_clauses production.
	ExitWhen_then_clauses(c *When_then_clausesContext)

	// ExitIf_statement is called when exiting the if_statement production.
	ExitIf_statement(c *If_statementContext)

	// ExitElseif_clauses is called when exiting the elseif_clauses production.
	ExitElseif_clauses(c *Elseif_clausesContext)

	// ExitOpt_else is called when exiting the opt_else production.
	ExitOpt_else(c *Opt_elseContext)

	// ExitOpt_as_code is called when exiting the opt_as_code production.
	ExitOpt_as_code(c *Opt_as_codeContext)

	// ExitOpt_external_security_clause is called when exiting the opt_external_security_clause production.
	ExitOpt_external_security_clause(c *Opt_external_security_clauseContext)

	// ExitExternal_security_clause_kind is called when exiting the external_security_clause_kind production.
	ExitExternal_security_clause_kind(c *External_security_clause_kindContext)

	// ExitProcedure_parameters is called when exiting the procedure_parameters production.
	ExitProcedure_parameters(c *Procedure_parametersContext)

	// ExitProcedure_parameter is called when exiting the procedure_parameter production.
	ExitProcedure_parameter(c *Procedure_parameterContext)

	// ExitProcedure_parameter_termination is called when exiting the procedure_parameter_termination production.
	ExitProcedure_parameter_termination(c *Procedure_parameter_terminationContext)

	// ExitOpt_procedure_parameter_mode is called when exiting the opt_procedure_parameter_mode production.
	ExitOpt_procedure_parameter_mode(c *Opt_procedure_parameter_modeContext)

	// ExitCreate_function_statement is called when exiting the create_function_statement production.
	ExitCreate_function_statement(c *Create_function_statementContext)

	// ExitOpt_determinism_level is called when exiting the opt_determinism_level production.
	ExitOpt_determinism_level(c *Opt_determinism_levelContext)

	// ExitOpt_sql_security_clause is called when exiting the opt_sql_security_clause production.
	ExitOpt_sql_security_clause(c *Opt_sql_security_clauseContext)

	// ExitSql_security_clause_kind is called when exiting the sql_security_clause_kind production.
	ExitSql_security_clause_kind(c *Sql_security_clause_kindContext)

	// ExitAs_sql_function_body_or_string is called when exiting the as_sql_function_body_or_string production.
	ExitAs_sql_function_body_or_string(c *As_sql_function_body_or_stringContext)

	// ExitSql_function_body is called when exiting the sql_function_body production.
	ExitSql_function_body(c *Sql_function_bodyContext)

	// ExitUnordered_options_body is called when exiting the unordered_options_body production.
	ExitUnordered_options_body(c *Unordered_options_bodyContext)

	// ExitOpt_language_or_remote_with_connection is called when exiting the opt_language_or_remote_with_connection production.
	ExitOpt_language_or_remote_with_connection(c *Opt_language_or_remote_with_connectionContext)

	// ExitLanguage is called when exiting the language production.
	ExitLanguage(c *LanguageContext)

	// ExitRemote_with_connection_clause is called when exiting the remote_with_connection_clause production.
	ExitRemote_with_connection_clause(c *Remote_with_connection_clauseContext)

	// ExitWith_connection_clause is called when exiting the with_connection_clause production.
	ExitWith_connection_clause(c *With_connection_clauseContext)

	// ExitOpt_function_returns is called when exiting the opt_function_returns production.
	ExitOpt_function_returns(c *Opt_function_returnsContext)

	// ExitOpt_returns is called when exiting the opt_returns production.
	ExitOpt_returns(c *Opt_returnsContext)

	// ExitFunction_declaration is called when exiting the function_declaration production.
	ExitFunction_declaration(c *Function_declarationContext)

	// ExitFunction_parameters is called when exiting the function_parameters production.
	ExitFunction_parameters(c *Function_parametersContext)

	// ExitFunction_parameter is called when exiting the function_parameter production.
	ExitFunction_parameter(c *Function_parameterContext)

	// ExitOpt_not_aggregate is called when exiting the opt_not_aggregate production.
	ExitOpt_not_aggregate(c *Opt_not_aggregateContext)

	// ExitOpt_default_expression is called when exiting the opt_default_expression production.
	ExitOpt_default_expression(c *Opt_default_expressionContext)

	// ExitType_or_tvf_schema is called when exiting the type_or_tvf_schema production.
	ExitType_or_tvf_schema(c *Type_or_tvf_schemaContext)

	// ExitTvf_schema is called when exiting the tvf_schema production.
	ExitTvf_schema(c *Tvf_schemaContext)

	// ExitTvf_schema_column is called when exiting the tvf_schema_column production.
	ExitTvf_schema_column(c *Tvf_schema_columnContext)

	// ExitTemplated_parameter_type is called when exiting the templated_parameter_type production.
	ExitTemplated_parameter_type(c *Templated_parameter_typeContext)

	// ExitTemplated_parameter_kind is called when exiting the templated_parameter_kind production.
	ExitTemplated_parameter_kind(c *Templated_parameter_kindContext)

	// ExitOpt_aggregate is called when exiting the opt_aggregate production.
	ExitOpt_aggregate(c *Opt_aggregateContext)

	// ExitCreate_database_statement is called when exiting the create_database_statement production.
	ExitCreate_database_statement(c *Create_database_statementContext)

	// ExitCreate_connection_statement is called when exiting the create_connection_statement production.
	ExitCreate_connection_statement(c *Create_connection_statementContext)

	// ExitCreate_constant_statement is called when exiting the create_constant_statement production.
	ExitCreate_constant_statement(c *Create_constant_statementContext)

	// ExitOpt_or_replace is called when exiting the opt_or_replace production.
	ExitOpt_or_replace(c *Opt_or_replaceContext)

	// ExitOpt_create_scope is called when exiting the opt_create_scope production.
	ExitOpt_create_scope(c *Opt_create_scopeContext)

	// ExitRun_batch_statement is called when exiting the run_batch_statement production.
	ExitRun_batch_statement(c *Run_batch_statementContext)

	// ExitAbort_batch_statement is called when exiting the abort_batch_statement production.
	ExitAbort_batch_statement(c *Abort_batch_statementContext)

	// ExitStart_batch_statement is called when exiting the start_batch_statement production.
	ExitStart_batch_statement(c *Start_batch_statementContext)

	// ExitRollback_statement is called when exiting the rollback_statement production.
	ExitRollback_statement(c *Rollback_statementContext)

	// ExitCommit_statement is called when exiting the commit_statement production.
	ExitCommit_statement(c *Commit_statementContext)

	// ExitSet_statement is called when exiting the set_statement production.
	ExitSet_statement(c *Set_statementContext)

	// ExitIdentifier_list is called when exiting the identifier_list production.
	ExitIdentifier_list(c *Identifier_listContext)

	// ExitBegin_statement is called when exiting the begin_statement production.
	ExitBegin_statement(c *Begin_statementContext)

	// ExitBegin_transaction_keywords is called when exiting the begin_transaction_keywords production.
	ExitBegin_transaction_keywords(c *Begin_transaction_keywordsContext)

	// ExitTransaction_mode_list is called when exiting the transaction_mode_list production.
	ExitTransaction_mode_list(c *Transaction_mode_listContext)

	// ExitTransaction_mode is called when exiting the transaction_mode production.
	ExitTransaction_mode(c *Transaction_modeContext)

	// ExitTruncate_statement is called when exiting the truncate_statement production.
	ExitTruncate_statement(c *Truncate_statementContext)

	// ExitMerge_statement is called when exiting the merge_statement production.
	ExitMerge_statement(c *Merge_statementContext)

	// ExitMerge_source is called when exiting the merge_source production.
	ExitMerge_source(c *Merge_sourceContext)

	// ExitMerge_when_clause is called when exiting the merge_when_clause production.
	ExitMerge_when_clause(c *Merge_when_clauseContext)

	// ExitMerge_action is called when exiting the merge_action production.
	ExitMerge_action(c *Merge_actionContext)

	// ExitMerge_insert_value_list_or_source_row is called when exiting the merge_insert_value_list_or_source_row production.
	ExitMerge_insert_value_list_or_source_row(c *Merge_insert_value_list_or_source_rowContext)

	// ExitBy_target is called when exiting the by_target production.
	ExitBy_target(c *By_targetContext)

	// ExitOpt_and_expression is called when exiting the opt_and_expression production.
	ExitOpt_and_expression(c *Opt_and_expressionContext)

	// ExitStatement_level_hint is called when exiting the statement_level_hint production.
	ExitStatement_level_hint(c *Statement_level_hintContext)

	// ExitQuery_statement is called when exiting the query_statement production.
	ExitQuery_statement(c *Query_statementContext)

	// ExitDml_statement is called when exiting the dml_statement production.
	ExitDml_statement(c *Dml_statementContext)

	// ExitUpdate_statement is called when exiting the update_statement production.
	ExitUpdate_statement(c *Update_statementContext)

	// ExitDelete_statement is called when exiting the delete_statement production.
	ExitDelete_statement(c *Delete_statementContext)

	// ExitInsert_statement is called when exiting the insert_statement production.
	ExitInsert_statement(c *Insert_statementContext)

	// ExitOn_conflict_clause is called when exiting the on_conflict_clause production.
	ExitOn_conflict_clause(c *On_conflict_clauseContext)

	// ExitOpt_where_expression is called when exiting the opt_where_expression production.
	ExitOpt_where_expression(c *Opt_where_expressionContext)

	// ExitOpt_conflict_target is called when exiting the opt_conflict_target production.
	ExitOpt_conflict_target(c *Opt_conflict_targetContext)

	// ExitUpdate_item_list is called when exiting the update_item_list production.
	ExitUpdate_item_list(c *Update_item_listContext)

	// ExitUpdate_item is called when exiting the update_item production.
	ExitUpdate_item(c *Update_itemContext)

	// ExitUpdate_set_value is called when exiting the update_set_value production.
	ExitUpdate_set_value(c *Update_set_valueContext)

	// ExitNested_dml_statement is called when exiting the nested_dml_statement production.
	ExitNested_dml_statement(c *Nested_dml_statementContext)

	// ExitInsert_values_list_or_table_clause is called when exiting the insert_values_list_or_table_clause production.
	ExitInsert_values_list_or_table_clause(c *Insert_values_list_or_table_clauseContext)

	// ExitTable_clause_unreversed is called when exiting the table_clause_unreversed production.
	ExitTable_clause_unreversed(c *Table_clause_unreversedContext)

	// ExitTable_clause_no_keyword is called when exiting the table_clause_no_keyword production.
	ExitTable_clause_no_keyword(c *Table_clause_no_keywordContext)

	// ExitOpt_returning_clause is called when exiting the opt_returning_clause production.
	ExitOpt_returning_clause(c *Opt_returning_clauseContext)

	// ExitOpt_assert_rows_modified is called when exiting the opt_assert_rows_modified production.
	ExitOpt_assert_rows_modified(c *Opt_assert_rows_modifiedContext)

	// ExitInsert_values_or_query is called when exiting the insert_values_or_query production.
	ExitInsert_values_or_query(c *Insert_values_or_queryContext)

	// ExitInsert_values_list is called when exiting the insert_values_list production.
	ExitInsert_values_list(c *Insert_values_listContext)

	// ExitInsert_values_row is called when exiting the insert_values_row production.
	ExitInsert_values_row(c *Insert_values_rowContext)

	// ExitExpression_or_default is called when exiting the expression_or_default production.
	ExitExpression_or_default(c *Expression_or_defaultContext)

	// ExitInsert_statement_prefix is called when exiting the insert_statement_prefix production.
	ExitInsert_statement_prefix(c *Insert_statement_prefixContext)

	// ExitMaybe_dashed_generalized_path_expression is called when exiting the maybe_dashed_generalized_path_expression production.
	ExitMaybe_dashed_generalized_path_expression(c *Maybe_dashed_generalized_path_expressionContext)

	// ExitOpt_into is called when exiting the opt_into production.
	ExitOpt_into(c *Opt_intoContext)

	// ExitOpt_or_ignore_replace_update is called when exiting the opt_or_ignore_replace_update production.
	ExitOpt_or_ignore_replace_update(c *Opt_or_ignore_replace_updateContext)

	// ExitAlter_statement is called when exiting the alter_statement production.
	ExitAlter_statement(c *Alter_statementContext)

	// ExitAnalyze_statement is called when exiting the analyze_statement production.
	ExitAnalyze_statement(c *Analyze_statementContext)

	// ExitAssert_statement is called when exiting the assert_statement production.
	ExitAssert_statement(c *Assert_statementContext)

	// ExitAux_load_data_statement is called when exiting the aux_load_data_statement production.
	ExitAux_load_data_statement(c *Aux_load_data_statementContext)

	// ExitClone_data_statement is called when exiting the clone_data_statement production.
	ExitClone_data_statement(c *Clone_data_statementContext)

	// ExitClone_data_source_list is called when exiting the clone_data_source_list production.
	ExitClone_data_source_list(c *Clone_data_source_listContext)

	// ExitClone_data_source is called when exiting the clone_data_source production.
	ExitClone_data_source(c *Clone_data_sourceContext)

	// ExitOpt_external_table_with_clauses is called when exiting the opt_external_table_with_clauses production.
	ExitOpt_external_table_with_clauses(c *Opt_external_table_with_clausesContext)

	// ExitWith_partition_columns_clause is called when exiting the with_partition_columns_clause production.
	ExitWith_partition_columns_clause(c *With_partition_columns_clauseContext)

	// ExitAux_load_data_from_files_options_list is called when exiting the aux_load_data_from_files_options_list production.
	ExitAux_load_data_from_files_options_list(c *Aux_load_data_from_files_options_listContext)

	// ExitCluster_by_clause_prefix_no_hint is called when exiting the cluster_by_clause_prefix_no_hint production.
	ExitCluster_by_clause_prefix_no_hint(c *Cluster_by_clause_prefix_no_hintContext)

	// ExitLoad_data_partitions_clause is called when exiting the load_data_partitions_clause production.
	ExitLoad_data_partitions_clause(c *Load_data_partitions_clauseContext)

	// ExitMaybe_dashed_path_expression_with_scope is called when exiting the maybe_dashed_path_expression_with_scope production.
	ExitMaybe_dashed_path_expression_with_scope(c *Maybe_dashed_path_expression_with_scopeContext)

	// ExitTable_element_list is called when exiting the table_element_list production.
	ExitTable_element_list(c *Table_element_listContext)

	// ExitTable_element is called when exiting the table_element production.
	ExitTable_element(c *Table_elementContext)

	// ExitTable_constraint_definition is called when exiting the table_constraint_definition production.
	ExitTable_constraint_definition(c *Table_constraint_definitionContext)

	// ExitAppend_or_overwrite is called when exiting the append_or_overwrite production.
	ExitAppend_or_overwrite(c *Append_or_overwriteContext)

	// ExitOpt_description is called when exiting the opt_description production.
	ExitOpt_description(c *Opt_descriptionContext)

	// ExitTable_and_column_info_list is called when exiting the table_and_column_info_list production.
	ExitTable_and_column_info_list(c *Table_and_column_info_listContext)

	// ExitTable_and_column_info is called when exiting the table_and_column_info production.
	ExitTable_and_column_info(c *Table_and_column_infoContext)

	// ExitRow_access_policy_alter_action_list is called when exiting the row_access_policy_alter_action_list production.
	ExitRow_access_policy_alter_action_list(c *Row_access_policy_alter_action_listContext)

	// ExitRow_access_policy_alter_action is called when exiting the row_access_policy_alter_action production.
	ExitRow_access_policy_alter_action(c *Row_access_policy_alter_actionContext)

	// ExitGrant_to_clause is called when exiting the grant_to_clause production.
	ExitGrant_to_clause(c *Grant_to_clauseContext)

	// ExitGrantee_list is called when exiting the grantee_list production.
	ExitGrantee_list(c *Grantee_listContext)

	// ExitPrivilege_list is called when exiting the privilege_list production.
	ExitPrivilege_list(c *Privilege_listContext)

	// ExitPrivilege is called when exiting the privilege production.
	ExitPrivilege(c *PrivilegeContext)

	// ExitPath_expression_list_with_parens is called when exiting the path_expression_list_with_parens production.
	ExitPath_expression_list_with_parens(c *Path_expression_list_with_parensContext)

	// ExitPrivilege_name is called when exiting the privilege_name production.
	ExitPrivilege_name(c *Privilege_nameContext)

	// ExitGeneric_entity_type is called when exiting the generic_entity_type production.
	ExitGeneric_entity_type(c *Generic_entity_typeContext)

	// ExitGeneric_entity_type_unchecked is called when exiting the generic_entity_type_unchecked production.
	ExitGeneric_entity_type_unchecked(c *Generic_entity_type_uncheckedContext)

	// ExitSchema_object_kind is called when exiting the schema_object_kind production.
	ExitSchema_object_kind(c *Schema_object_kindContext)

	// ExitAlter_action_list is called when exiting the alter_action_list production.
	ExitAlter_action_list(c *Alter_action_listContext)

	// ExitAlter_action is called when exiting the alter_action production.
	ExitAlter_action(c *Alter_actionContext)

	// ExitSpanner_set_on_delete_action is called when exiting the spanner_set_on_delete_action production.
	ExitSpanner_set_on_delete_action(c *Spanner_set_on_delete_actionContext)

	// ExitSpanner_alter_column_action is called when exiting the spanner_alter_column_action production.
	ExitSpanner_alter_column_action(c *Spanner_alter_column_actionContext)

	// ExitSpanner_generated_or_default is called when exiting the spanner_generated_or_default production.
	ExitSpanner_generated_or_default(c *Spanner_generated_or_defaultContext)

	// ExitGeneric_sub_entity_type is called when exiting the generic_sub_entity_type production.
	ExitGeneric_sub_entity_type(c *Generic_sub_entity_typeContext)

	// ExitSub_entity_type_identifier is called when exiting the sub_entity_type_identifier production.
	ExitSub_entity_type_identifier(c *Sub_entity_type_identifierContext)

	// ExitFill_using_expression is called when exiting the fill_using_expression production.
	ExitFill_using_expression(c *Fill_using_expressionContext)

	// ExitColumn_position is called when exiting the column_position production.
	ExitColumn_position(c *Column_positionContext)

	// ExitTable_column_definition is called when exiting the table_column_definition production.
	ExitTable_column_definition(c *Table_column_definitionContext)

	// ExitColumn_attributes is called when exiting the column_attributes production.
	ExitColumn_attributes(c *Column_attributesContext)

	// ExitColumn_attribute is called when exiting the column_attribute production.
	ExitColumn_attribute(c *Column_attributeContext)

	// ExitPrimary_key_column_attribute is called when exiting the primary_key_column_attribute production.
	ExitPrimary_key_column_attribute(c *Primary_key_column_attributeContext)

	// ExitForeign_key_column_attribute is called when exiting the foreign_key_column_attribute production.
	ExitForeign_key_column_attribute(c *Foreign_key_column_attributeContext)

	// ExitHidden_column_attribute is called when exiting the hidden_column_attribute production.
	ExitHidden_column_attribute(c *Hidden_column_attributeContext)

	// ExitOpt_constraint_identity is called when exiting the opt_constraint_identity production.
	ExitOpt_constraint_identity(c *Opt_constraint_identityContext)

	// ExitTable_column_schema is called when exiting the table_column_schema production.
	ExitTable_column_schema(c *Table_column_schemaContext)

	// ExitOpt_column_info is called when exiting the opt_column_info production.
	ExitOpt_column_info(c *Opt_column_infoContext)

	// ExitInvalid_generated_column is called when exiting the invalid_generated_column production.
	ExitInvalid_generated_column(c *Invalid_generated_columnContext)

	// ExitInvalid_default_column is called when exiting the invalid_default_column production.
	ExitInvalid_default_column(c *Invalid_default_columnContext)

	// ExitDefault_column_info is called when exiting the default_column_info production.
	ExitDefault_column_info(c *Default_column_infoContext)

	// ExitGenerated_column_info is called when exiting the generated_column_info production.
	ExitGenerated_column_info(c *Generated_column_infoContext)

	// ExitIdentity_column_info is called when exiting the identity_column_info production.
	ExitIdentity_column_info(c *Identity_column_infoContext)

	// ExitOpt_start_with is called when exiting the opt_start_with production.
	ExitOpt_start_with(c *Opt_start_withContext)

	// ExitOpt_increment_by is called when exiting the opt_increment_by production.
	ExitOpt_increment_by(c *Opt_increment_byContext)

	// ExitOpt_maxvalue is called when exiting the opt_maxvalue production.
	ExitOpt_maxvalue(c *Opt_maxvalueContext)

	// ExitOpt_minvalue is called when exiting the opt_minvalue production.
	ExitOpt_minvalue(c *Opt_minvalueContext)

	// ExitOpt_cycle is called when exiting the opt_cycle production.
	ExitOpt_cycle(c *Opt_cycleContext)

	// ExitSigned_numeric_literal is called when exiting the signed_numeric_literal production.
	ExitSigned_numeric_literal(c *Signed_numeric_literalContext)

	// ExitStored_mode is called when exiting the stored_mode production.
	ExitStored_mode(c *Stored_modeContext)

	// ExitGenerated_mode is called when exiting the generated_mode production.
	ExitGenerated_mode(c *Generated_modeContext)

	// ExitColumn_schema_inner is called when exiting the column_schema_inner production.
	ExitColumn_schema_inner(c *Column_schema_innerContext)

	// ExitRaw_column_schema_inner is called when exiting the raw_column_schema_inner production.
	ExitRaw_column_schema_inner(c *Raw_column_schema_innerContext)

	// ExitRange_column_schema_inner is called when exiting the range_column_schema_inner production.
	ExitRange_column_schema_inner(c *Range_column_schema_innerContext)

	// ExitStruct_column_schema_inner is called when exiting the struct_column_schema_inner production.
	ExitStruct_column_schema_inner(c *Struct_column_schema_innerContext)

	// ExitStruct_column_field is called when exiting the struct_column_field production.
	ExitStruct_column_field(c *Struct_column_fieldContext)

	// ExitSimple_column_schema_inner is called when exiting the simple_column_schema_inner production.
	ExitSimple_column_schema_inner(c *Simple_column_schema_innerContext)

	// ExitArray_column_schema_inner is called when exiting the array_column_schema_inner production.
	ExitArray_column_schema_inner(c *Array_column_schema_innerContext)

	// ExitField_schema is called when exiting the field_schema production.
	ExitField_schema(c *Field_schemaContext)

	// ExitOpt_field_attributes is called when exiting the opt_field_attributes production.
	ExitOpt_field_attributes(c *Opt_field_attributesContext)

	// ExitNot_null_column_attribute is called when exiting the not_null_column_attribute production.
	ExitNot_null_column_attribute(c *Not_null_column_attributeContext)

	// ExitPrimary_key_or_table_constraint_spec is called when exiting the primary_key_or_table_constraint_spec production.
	ExitPrimary_key_or_table_constraint_spec(c *Primary_key_or_table_constraint_specContext)

	// ExitOpt_if_not_exists is called when exiting the opt_if_not_exists production.
	ExitOpt_if_not_exists(c *Opt_if_not_existsContext)

	// ExitPrimary_key_spec is called when exiting the primary_key_spec production.
	ExitPrimary_key_spec(c *Primary_key_specContext)

	// ExitPrimary_key_element_list is called when exiting the primary_key_element_list production.
	ExitPrimary_key_element_list(c *Primary_key_element_listContext)

	// ExitPrimary_key_element is called when exiting the primary_key_element production.
	ExitPrimary_key_element(c *Primary_key_elementContext)

	// ExitTable_constraint_spec is called when exiting the table_constraint_spec production.
	ExitTable_constraint_spec(c *Table_constraint_specContext)

	// ExitForeign_key_reference is called when exiting the foreign_key_reference production.
	ExitForeign_key_reference(c *Foreign_key_referenceContext)

	// ExitOpt_foreign_key_action is called when exiting the opt_foreign_key_action production.
	ExitOpt_foreign_key_action(c *Opt_foreign_key_actionContext)

	// ExitForeign_key_on_update is called when exiting the foreign_key_on_update production.
	ExitForeign_key_on_update(c *Foreign_key_on_updateContext)

	// ExitForeign_key_on_delete is called when exiting the foreign_key_on_delete production.
	ExitForeign_key_on_delete(c *Foreign_key_on_deleteContext)

	// ExitForeign_key_action is called when exiting the foreign_key_action production.
	ExitForeign_key_action(c *Foreign_key_actionContext)

	// ExitOpt_foreign_key_match is called when exiting the opt_foreign_key_match production.
	ExitOpt_foreign_key_match(c *Opt_foreign_key_matchContext)

	// ExitForeign_key_match_mode is called when exiting the foreign_key_match_mode production.
	ExitForeign_key_match_mode(c *Foreign_key_match_modeContext)

	// ExitColumn_list is called when exiting the column_list production.
	ExitColumn_list(c *Column_listContext)

	// ExitOpt_options_list is called when exiting the opt_options_list production.
	ExitOpt_options_list(c *Opt_options_listContext)

	// ExitConstraint_enforcement is called when exiting the constraint_enforcement production.
	ExitConstraint_enforcement(c *Constraint_enforcementContext)

	// ExitGeneric_entity_body is called when exiting the generic_entity_body production.
	ExitGeneric_entity_body(c *Generic_entity_bodyContext)

	// ExitOpt_if_exists is called when exiting the opt_if_exists production.
	ExitOpt_if_exists(c *Opt_if_existsContext)

	// ExitTable_or_table_function is called when exiting the table_or_table_function production.
	ExitTable_or_table_function(c *Table_or_table_functionContext)

	// ExitQuery is called when exiting the query production.
	ExitQuery(c *QueryContext)

	// ExitQuery_without_pipe_operators is called when exiting the query_without_pipe_operators production.
	ExitQuery_without_pipe_operators(c *Query_without_pipe_operatorsContext)

	// ExitBad_keyword_after_from_query is called when exiting the bad_keyword_after_from_query production.
	ExitBad_keyword_after_from_query(c *Bad_keyword_after_from_queryContext)

	// ExitBad_keyword_after_from_query_allows_parens is called when exiting the bad_keyword_after_from_query_allows_parens production.
	ExitBad_keyword_after_from_query_allows_parens(c *Bad_keyword_after_from_query_allows_parensContext)

	// ExitWith_clause_with_trailing_comma is called when exiting the with_clause_with_trailing_comma production.
	ExitWith_clause_with_trailing_comma(c *With_clause_with_trailing_commaContext)

	// ExitSelect_or_from_keyword is called when exiting the select_or_from_keyword production.
	ExitSelect_or_from_keyword(c *Select_or_from_keywordContext)

	// ExitQuery_primary_or_set_operation is called when exiting the query_primary_or_set_operation production.
	ExitQuery_primary_or_set_operation(c *Query_primary_or_set_operationContext)

	// ExitQuery_set_operation is called when exiting the query_set_operation production.
	ExitQuery_set_operation(c *Query_set_operationContext)

	// ExitQuery_set_operation_prefix is called when exiting the query_set_operation_prefix production.
	ExitQuery_set_operation_prefix(c *Query_set_operation_prefixContext)

	// ExitQuery_set_operation_item is called when exiting the query_set_operation_item production.
	ExitQuery_set_operation_item(c *Query_set_operation_itemContext)

	// ExitQuery_primary is called when exiting the query_primary production.
	ExitQuery_primary(c *Query_primaryContext)

	// ExitSet_operation_metadata is called when exiting the set_operation_metadata production.
	ExitSet_operation_metadata(c *Set_operation_metadataContext)

	// ExitOpt_column_match_suffix is called when exiting the opt_column_match_suffix production.
	ExitOpt_column_match_suffix(c *Opt_column_match_suffixContext)

	// ExitOpt_strict is called when exiting the opt_strict production.
	ExitOpt_strict(c *Opt_strictContext)

	// ExitAll_or_distinct is called when exiting the all_or_distinct production.
	ExitAll_or_distinct(c *All_or_distinctContext)

	// ExitQuery_set_operation_type is called when exiting the query_set_operation_type production.
	ExitQuery_set_operation_type(c *Query_set_operation_typeContext)

	// ExitOpt_corresponding_outer_mode is called when exiting the opt_corresponding_outer_mode production.
	ExitOpt_corresponding_outer_mode(c *Opt_corresponding_outer_modeContext)

	// ExitOpt_outer is called when exiting the opt_outer production.
	ExitOpt_outer(c *Opt_outerContext)

	// ExitWith_clause is called when exiting the with_clause production.
	ExitWith_clause(c *With_clauseContext)

	// ExitAliased_query is called when exiting the aliased_query production.
	ExitAliased_query(c *Aliased_queryContext)

	// ExitOpt_aliased_query_modifiers is called when exiting the opt_aliased_query_modifiers production.
	ExitOpt_aliased_query_modifiers(c *Opt_aliased_query_modifiersContext)

	// ExitRecursion_depth_modifier is called when exiting the recursion_depth_modifier production.
	ExitRecursion_depth_modifier(c *Recursion_depth_modifierContext)

	// ExitPossibly_unbounded_int_literal_or_parameter is called when exiting the possibly_unbounded_int_literal_or_parameter production.
	ExitPossibly_unbounded_int_literal_or_parameter(c *Possibly_unbounded_int_literal_or_parameterContext)

	// ExitInt_literal_or_parameter is called when exiting the int_literal_or_parameter production.
	ExitInt_literal_or_parameter(c *Int_literal_or_parameterContext)

	// ExitOrder_by_clause is called when exiting the order_by_clause production.
	ExitOrder_by_clause(c *Order_by_clauseContext)

	// ExitOrder_by_clause_prefix is called when exiting the order_by_clause_prefix production.
	ExitOrder_by_clause_prefix(c *Order_by_clause_prefixContext)

	// ExitOrdering_expression is called when exiting the ordering_expression production.
	ExitOrdering_expression(c *Ordering_expressionContext)

	// ExitSelect is called when exiting the select production.
	ExitSelect(c *SelectContext)

	// ExitOpt_clauses_following_from is called when exiting the opt_clauses_following_from production.
	ExitOpt_clauses_following_from(c *Opt_clauses_following_fromContext)

	// ExitOpt_clauses_following_where is called when exiting the opt_clauses_following_where production.
	ExitOpt_clauses_following_where(c *Opt_clauses_following_whereContext)

	// ExitOpt_clauses_following_group_by is called when exiting the opt_clauses_following_group_by production.
	ExitOpt_clauses_following_group_by(c *Opt_clauses_following_group_byContext)

	// ExitWindow_clause is called when exiting the window_clause production.
	ExitWindow_clause(c *Window_clauseContext)

	// ExitWindow_clause_prefix is called when exiting the window_clause_prefix production.
	ExitWindow_clause_prefix(c *Window_clause_prefixContext)

	// ExitWindow_definition is called when exiting the window_definition production.
	ExitWindow_definition(c *Window_definitionContext)

	// ExitWhere_clause is called when exiting the where_clause production.
	ExitWhere_clause(c *Where_clauseContext)

	// ExitHaving_clause is called when exiting the having_clause production.
	ExitHaving_clause(c *Having_clauseContext)

	// ExitGroup_by_clause is called when exiting the group_by_clause production.
	ExitGroup_by_clause(c *Group_by_clauseContext)

	// ExitGroup_by_all is called when exiting the group_by_all production.
	ExitGroup_by_all(c *Group_by_allContext)

	// ExitSelect_clause is called when exiting the select_clause production.
	ExitSelect_clause(c *Select_clauseContext)

	// ExitOpt_select_as_clause is called when exiting the opt_select_as_clause production.
	ExitOpt_select_as_clause(c *Opt_select_as_clauseContext)

	// ExitOpt_select_with is called when exiting the opt_select_with production.
	ExitOpt_select_with(c *Opt_select_withContext)

	// ExitFrom_clause is called when exiting the from_clause production.
	ExitFrom_clause(c *From_clauseContext)

	// ExitFrom_clause_contents is called when exiting the from_clause_contents production.
	ExitFrom_clause_contents(c *From_clause_contentsContext)

	// ExitFrom_clause_contents_suffix is called when exiting the from_clause_contents_suffix production.
	ExitFrom_clause_contents_suffix(c *From_clause_contents_suffixContext)

	// ExitTable_primary is called when exiting the table_primary production.
	ExitTable_primary(c *Table_primaryContext)

	// ExitTvf_with_suffixes is called when exiting the tvf_with_suffixes production.
	ExitTvf_with_suffixes(c *Tvf_with_suffixesContext)

	// ExitPivot_or_unpivot_clause_and_aliases is called when exiting the pivot_or_unpivot_clause_and_aliases production.
	ExitPivot_or_unpivot_clause_and_aliases(c *Pivot_or_unpivot_clause_and_aliasesContext)

	// ExitAs_alias is called when exiting the as_alias production.
	ExitAs_alias(c *As_aliasContext)

	// ExitSample_clause is called when exiting the sample_clause production.
	ExitSample_clause(c *Sample_clauseContext)

	// ExitOpt_sample_clause_suffix is called when exiting the opt_sample_clause_suffix production.
	ExitOpt_sample_clause_suffix(c *Opt_sample_clause_suffixContext)

	// ExitRepeatable_clause is called when exiting the repeatable_clause production.
	ExitRepeatable_clause(c *Repeatable_clauseContext)

	// ExitPossibly_cast_int_literal_or_parameter is called when exiting the possibly_cast_int_literal_or_parameter production.
	ExitPossibly_cast_int_literal_or_parameter(c *Possibly_cast_int_literal_or_parameterContext)

	// ExitCast_int_literal_or_parameter is called when exiting the cast_int_literal_or_parameter production.
	ExitCast_int_literal_or_parameter(c *Cast_int_literal_or_parameterContext)

	// ExitSample_size is called when exiting the sample_size production.
	ExitSample_size(c *Sample_sizeContext)

	// ExitSample_size_value is called when exiting the sample_size_value production.
	ExitSample_size_value(c *Sample_size_valueContext)

	// ExitSample_size_unit is called when exiting the sample_size_unit production.
	ExitSample_size_unit(c *Sample_size_unitContext)

	// ExitPartition_by_clause_prefix_no_hint is called when exiting the partition_by_clause_prefix_no_hint production.
	ExitPartition_by_clause_prefix_no_hint(c *Partition_by_clause_prefix_no_hintContext)

	// ExitMatch_recognize_clause is called when exiting the match_recognize_clause production.
	ExitMatch_recognize_clause(c *Match_recognize_clauseContext)

	// ExitRow_pattern_expr is called when exiting the row_pattern_expr production.
	ExitRow_pattern_expr(c *Row_pattern_exprContext)

	// ExitRow_pattern_concatenation is called when exiting the row_pattern_concatenation production.
	ExitRow_pattern_concatenation(c *Row_pattern_concatenationContext)

	// ExitRow_pattern_factor is called when exiting the row_pattern_factor production.
	ExitRow_pattern_factor(c *Row_pattern_factorContext)

	// ExitSelect_list_prefix_with_as_aliases is called when exiting the select_list_prefix_with_as_aliases production.
	ExitSelect_list_prefix_with_as_aliases(c *Select_list_prefix_with_as_aliasesContext)

	// ExitSelect_column_expr_with_as_alias is called when exiting the select_column_expr_with_as_alias production.
	ExitSelect_column_expr_with_as_alias(c *Select_column_expr_with_as_aliasContext)

	// ExitTable_subquery is called when exiting the table_subquery production.
	ExitTable_subquery(c *Table_subqueryContext)

	// ExitJoin is called when exiting the join production.
	ExitJoin(c *JoinContext)

	// ExitJoin_item is called when exiting the join_item production.
	ExitJoin_item(c *Join_itemContext)

	// ExitOn_or_using_clause_list is called when exiting the on_or_using_clause_list production.
	ExitOn_or_using_clause_list(c *On_or_using_clause_listContext)

	// ExitOn_or_using_clause is called when exiting the on_or_using_clause production.
	ExitOn_or_using_clause(c *On_or_using_clauseContext)

	// ExitUsing_clause is called when exiting the using_clause production.
	ExitUsing_clause(c *Using_clauseContext)

	// ExitJoin_hint is called when exiting the join_hint production.
	ExitJoin_hint(c *Join_hintContext)

	// ExitTable_path_expression is called when exiting the table_path_expression production.
	ExitTable_path_expression(c *Table_path_expressionContext)

	// ExitOpt_at_system_time is called when exiting the opt_at_system_time production.
	ExitOpt_at_system_time(c *Opt_at_system_timeContext)

	// ExitOpt_with_offset_and_alias is called when exiting the opt_with_offset_and_alias production.
	ExitOpt_with_offset_and_alias(c *Opt_with_offset_and_aliasContext)

	// ExitOpt_pivot_or_unpivot_clause_and_alias is called when exiting the opt_pivot_or_unpivot_clause_and_alias production.
	ExitOpt_pivot_or_unpivot_clause_and_alias(c *Opt_pivot_or_unpivot_clause_and_aliasContext)

	// ExitTable_path_expression_base is called when exiting the table_path_expression_base production.
	ExitTable_path_expression_base(c *Table_path_expression_baseContext)

	// ExitMaybe_slashed_or_dashed_path_expression is called when exiting the maybe_slashed_or_dashed_path_expression production.
	ExitMaybe_slashed_or_dashed_path_expression(c *Maybe_slashed_or_dashed_path_expressionContext)

	// ExitMaybe_dashed_path_expression is called when exiting the maybe_dashed_path_expression production.
	ExitMaybe_dashed_path_expression(c *Maybe_dashed_path_expressionContext)

	// ExitDashed_path_expression is called when exiting the dashed_path_expression production.
	ExitDashed_path_expression(c *Dashed_path_expressionContext)

	// ExitDashed_identifier is called when exiting the dashed_identifier production.
	ExitDashed_identifier(c *Dashed_identifierContext)

	// ExitSlashed_identifier is called when exiting the slashed_identifier production.
	ExitSlashed_identifier(c *Slashed_identifierContext)

	// ExitIdentifier_or_integer is called when exiting the identifier_or_integer production.
	ExitIdentifier_or_integer(c *Identifier_or_integerContext)

	// ExitSlashed_identifier_separator is called when exiting the slashed_identifier_separator production.
	ExitSlashed_identifier_separator(c *Slashed_identifier_separatorContext)

	// ExitSlashed_path_expression is called when exiting the slashed_path_expression production.
	ExitSlashed_path_expression(c *Slashed_path_expressionContext)

	// ExitUnnest_expression is called when exiting the unnest_expression production.
	ExitUnnest_expression(c *Unnest_expressionContext)

	// ExitUnnest_expression_prefix is called when exiting the unnest_expression_prefix production.
	ExitUnnest_expression_prefix(c *Unnest_expression_prefixContext)

	// ExitOpt_array_zip_mode is called when exiting the opt_array_zip_mode production.
	ExitOpt_array_zip_mode(c *Opt_array_zip_modeContext)

	// ExitExpression_with_opt_alias is called when exiting the expression_with_opt_alias production.
	ExitExpression_with_opt_alias(c *Expression_with_opt_aliasContext)

	// ExitTvf_prefix is called when exiting the tvf_prefix production.
	ExitTvf_prefix(c *Tvf_prefixContext)

	// ExitTvf_argument is called when exiting the tvf_argument production.
	ExitTvf_argument(c *Tvf_argumentContext)

	// ExitConnection_clause is called when exiting the connection_clause production.
	ExitConnection_clause(c *Connection_clauseContext)

	// ExitPath_expression_or_default is called when exiting the path_expression_or_default production.
	ExitPath_expression_or_default(c *Path_expression_or_defaultContext)

	// ExitDescriptor_argument is called when exiting the descriptor_argument production.
	ExitDescriptor_argument(c *Descriptor_argumentContext)

	// ExitDescriptor_column_list is called when exiting the descriptor_column_list production.
	ExitDescriptor_column_list(c *Descriptor_column_listContext)

	// ExitDescriptor_column is called when exiting the descriptor_column production.
	ExitDescriptor_column(c *Descriptor_columnContext)

	// ExitTable_clause is called when exiting the table_clause production.
	ExitTable_clause(c *Table_clauseContext)

	// ExitModel_clause is called when exiting the model_clause production.
	ExitModel_clause(c *Model_clauseContext)

	// ExitQualify_clause_nonreserved is called when exiting the qualify_clause_nonreserved production.
	ExitQualify_clause_nonreserved(c *Qualify_clause_nonreservedContext)

	// ExitUnpivot_clause is called when exiting the unpivot_clause production.
	ExitUnpivot_clause(c *Unpivot_clauseContext)

	// ExitUnpivot_in_item_list is called when exiting the unpivot_in_item_list production.
	ExitUnpivot_in_item_list(c *Unpivot_in_item_listContext)

	// ExitUnpivot_in_item_list_prefix is called when exiting the unpivot_in_item_list_prefix production.
	ExitUnpivot_in_item_list_prefix(c *Unpivot_in_item_list_prefixContext)

	// ExitUnpivot_in_item is called when exiting the unpivot_in_item production.
	ExitUnpivot_in_item(c *Unpivot_in_itemContext)

	// ExitOpt_as_string_or_integer is called when exiting the opt_as_string_or_integer production.
	ExitOpt_as_string_or_integer(c *Opt_as_string_or_integerContext)

	// ExitPath_expression_list_with_opt_parens is called when exiting the path_expression_list_with_opt_parens production.
	ExitPath_expression_list_with_opt_parens(c *Path_expression_list_with_opt_parensContext)

	// ExitPath_expression_list is called when exiting the path_expression_list production.
	ExitPath_expression_list(c *Path_expression_listContext)

	// ExitUnpivot_nulls_filter is called when exiting the unpivot_nulls_filter production.
	ExitUnpivot_nulls_filter(c *Unpivot_nulls_filterContext)

	// ExitPivot_clause is called when exiting the pivot_clause production.
	ExitPivot_clause(c *Pivot_clauseContext)

	// ExitPivot_expression_list is called when exiting the pivot_expression_list production.
	ExitPivot_expression_list(c *Pivot_expression_listContext)

	// ExitPivot_expression is called when exiting the pivot_expression production.
	ExitPivot_expression(c *Pivot_expressionContext)

	// ExitPivot_value_list is called when exiting the pivot_value_list production.
	ExitPivot_value_list(c *Pivot_value_listContext)

	// ExitPivot_value is called when exiting the pivot_value production.
	ExitPivot_value(c *Pivot_valueContext)

	// ExitTvf_prefix_no_args is called when exiting the tvf_prefix_no_args production.
	ExitTvf_prefix_no_args(c *Tvf_prefix_no_argsContext)

	// ExitJoin_type is called when exiting the join_type production.
	ExitJoin_type(c *Join_typeContext)

	// ExitOpt_natural is called when exiting the opt_natural production.
	ExitOpt_natural(c *Opt_naturalContext)

	// ExitOn_clause is called when exiting the on_clause production.
	ExitOn_clause(c *On_clauseContext)

	// ExitSelect_list is called when exiting the select_list production.
	ExitSelect_list(c *Select_listContext)

	// ExitSelect_list_item is called when exiting the select_list_item production.
	ExitSelect_list_item(c *Select_list_itemContext)

	// ExitSelect_column_star is called when exiting the select_column_star production.
	ExitSelect_column_star(c *Select_column_starContext)

	// ExitSelect_column_expr is called when exiting the select_column_expr production.
	ExitSelect_column_expr(c *Select_column_exprContext)

	// ExitSelect_column_dot_star is called when exiting the select_column_dot_star production.
	ExitSelect_column_dot_star(c *Select_column_dot_starContext)

	// ExitStar_modifiers is called when exiting the star_modifiers production.
	ExitStar_modifiers(c *Star_modifiersContext)

	// ExitStar_except_list is called when exiting the star_except_list production.
	ExitStar_except_list(c *Star_except_listContext)

	// ExitStar_replace_list is called when exiting the star_replace_list production.
	ExitStar_replace_list(c *Star_replace_listContext)

	// ExitStar_replace_item is called when exiting the star_replace_item production.
	ExitStar_replace_item(c *Star_replace_itemContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitExpression_higher_prec_than_and is called when exiting the expression_higher_prec_than_and production.
	ExitExpression_higher_prec_than_and(c *Expression_higher_prec_than_andContext)

	// ExitExpression_maybe_parenthesized_not_a_query is called when exiting the expression_maybe_parenthesized_not_a_query production.
	ExitExpression_maybe_parenthesized_not_a_query(c *Expression_maybe_parenthesized_not_a_queryContext)

	// ExitParenthesized_in_rhs is called when exiting the parenthesized_in_rhs production.
	ExitParenthesized_in_rhs(c *Parenthesized_in_rhsContext)

	// ExitUnary_operator is called when exiting the unary_operator production.
	ExitUnary_operator(c *Unary_operatorContext)

	// ExitComparative_operator is called when exiting the comparative_operator production.
	ExitComparative_operator(c *Comparative_operatorContext)

	// ExitShift_operator is called when exiting the shift_operator production.
	ExitShift_operator(c *Shift_operatorContext)

	// ExitAdditive_operator is called when exiting the additive_operator production.
	ExitAdditive_operator(c *Additive_operatorContext)

	// ExitMultiplicative_operator is called when exiting the multiplicative_operator production.
	ExitMultiplicative_operator(c *Multiplicative_operatorContext)

	// ExitIs_operator is called when exiting the is_operator production.
	ExitIs_operator(c *Is_operatorContext)

	// ExitBetween_operator is called when exiting the between_operator production.
	ExitBetween_operator(c *Between_operatorContext)

	// ExitIn_operator is called when exiting the in_operator production.
	ExitIn_operator(c *In_operatorContext)

	// ExitDistinct_operator is called when exiting the distinct_operator production.
	ExitDistinct_operator(c *Distinct_operatorContext)

	// ExitParenthesized_query is called when exiting the parenthesized_query production.
	ExitParenthesized_query(c *Parenthesized_queryContext)

	// ExitParenthesized_expression_not_a_query is called when exiting the parenthesized_expression_not_a_query production.
	ExitParenthesized_expression_not_a_query(c *Parenthesized_expression_not_a_queryContext)

	// ExitParenthesized_anysomeall_list_in_rhs is called when exiting the parenthesized_anysomeall_list_in_rhs production.
	ExitParenthesized_anysomeall_list_in_rhs(c *Parenthesized_anysomeall_list_in_rhsContext)

	// ExitAnd_expression is called when exiting the and_expression production.
	ExitAnd_expression(c *And_expressionContext)

	// ExitIn_list_two_or_more_prefix is called when exiting the in_list_two_or_more_prefix production.
	ExitIn_list_two_or_more_prefix(c *In_list_two_or_more_prefixContext)

	// ExitAny_some_all is called when exiting the any_some_all production.
	ExitAny_some_all(c *Any_some_allContext)

	// ExitLike_operator is called when exiting the like_operator production.
	ExitLike_operator(c *Like_operatorContext)

	// ExitExpression_subquery_with_keyword is called when exiting the expression_subquery_with_keyword production.
	ExitExpression_subquery_with_keyword(c *Expression_subquery_with_keywordContext)

	// ExitStruct_constructor is called when exiting the struct_constructor production.
	ExitStruct_constructor(c *Struct_constructorContext)

	// ExitStruct_constructor_prefix_with_keyword is called when exiting the struct_constructor_prefix_with_keyword production.
	ExitStruct_constructor_prefix_with_keyword(c *Struct_constructor_prefix_with_keywordContext)

	// ExitStruct_constructor_arg is called when exiting the struct_constructor_arg production.
	ExitStruct_constructor_arg(c *Struct_constructor_argContext)

	// ExitStruct_constructor_prefix_without_keyword is called when exiting the struct_constructor_prefix_without_keyword production.
	ExitStruct_constructor_prefix_without_keyword(c *Struct_constructor_prefix_without_keywordContext)

	// ExitStruct_constructor_prefix_with_keyword_no_arg is called when exiting the struct_constructor_prefix_with_keyword_no_arg production.
	ExitStruct_constructor_prefix_with_keyword_no_arg(c *Struct_constructor_prefix_with_keyword_no_argContext)

	// ExitInterval_expression is called when exiting the interval_expression production.
	ExitInterval_expression(c *Interval_expressionContext)

	// ExitFunction_call_expression_with_clauses is called when exiting the function_call_expression_with_clauses production.
	ExitFunction_call_expression_with_clauses(c *Function_call_expression_with_clausesContext)

	// ExitFunction_call_expression_with_clauses_suffix is called when exiting the function_call_expression_with_clauses_suffix production.
	ExitFunction_call_expression_with_clauses_suffix(c *Function_call_expression_with_clauses_suffixContext)

	// ExitOver_clause is called when exiting the over_clause production.
	ExitOver_clause(c *Over_clauseContext)

	// ExitWindow_specification is called when exiting the window_specification production.
	ExitWindow_specification(c *Window_specificationContext)

	// ExitOpt_window_frame_clause is called when exiting the opt_window_frame_clause production.
	ExitOpt_window_frame_clause(c *Opt_window_frame_clauseContext)

	// ExitWindow_frame_bound is called when exiting the window_frame_bound production.
	ExitWindow_frame_bound(c *Window_frame_boundContext)

	// ExitPreceding_or_following is called when exiting the preceding_or_following production.
	ExitPreceding_or_following(c *Preceding_or_followingContext)

	// ExitFrame_unit is called when exiting the frame_unit production.
	ExitFrame_unit(c *Frame_unitContext)

	// ExitPartition_by_clause is called when exiting the partition_by_clause production.
	ExitPartition_by_clause(c *Partition_by_clauseContext)

	// ExitPartition_by_clause_prefix is called when exiting the partition_by_clause_prefix production.
	ExitPartition_by_clause_prefix(c *Partition_by_clause_prefixContext)

	// ExitWith_group_rows is called when exiting the with_group_rows production.
	ExitWith_group_rows(c *With_group_rowsContext)

	// ExitWith_report_modifier is called when exiting the with_report_modifier production.
	ExitWith_report_modifier(c *With_report_modifierContext)

	// ExitClamped_between_modifier is called when exiting the clamped_between_modifier production.
	ExitClamped_between_modifier(c *Clamped_between_modifierContext)

	// ExitWith_report_format is called when exiting the with_report_format production.
	ExitWith_report_format(c *With_report_formatContext)

	// ExitOptions_list is called when exiting the options_list production.
	ExitOptions_list(c *Options_listContext)

	// ExitOptions_list_prefix is called when exiting the options_list_prefix production.
	ExitOptions_list_prefix(c *Options_list_prefixContext)

	// ExitOptions_entry is called when exiting the options_entry production.
	ExitOptions_entry(c *Options_entryContext)

	// ExitExpression_or_proto is called when exiting the expression_or_proto production.
	ExitExpression_or_proto(c *Expression_or_protoContext)

	// ExitOptions_assignment_operator is called when exiting the options_assignment_operator production.
	ExitOptions_assignment_operator(c *Options_assignment_operatorContext)

	// ExitOpt_null_handling_modifier is called when exiting the opt_null_handling_modifier production.
	ExitOpt_null_handling_modifier(c *Opt_null_handling_modifierContext)

	// ExitFunction_call_argument is called when exiting the function_call_argument production.
	ExitFunction_call_argument(c *Function_call_argumentContext)

	// ExitSequence_arg is called when exiting the sequence_arg production.
	ExitSequence_arg(c *Sequence_argContext)

	// ExitNamed_argument is called when exiting the named_argument production.
	ExitNamed_argument(c *Named_argumentContext)

	// ExitLambda_argument is called when exiting the lambda_argument production.
	ExitLambda_argument(c *Lambda_argumentContext)

	// ExitLambda_argument_list is called when exiting the lambda_argument_list production.
	ExitLambda_argument_list(c *Lambda_argument_listContext)

	// ExitLimit_offset_clause is called when exiting the limit_offset_clause production.
	ExitLimit_offset_clause(c *Limit_offset_clauseContext)

	// ExitOpt_having_or_group_by_modifier is called when exiting the opt_having_or_group_by_modifier production.
	ExitOpt_having_or_group_by_modifier(c *Opt_having_or_group_by_modifierContext)

	// ExitGroup_by_clause_prefix is called when exiting the group_by_clause_prefix production.
	ExitGroup_by_clause_prefix(c *Group_by_clause_prefixContext)

	// ExitGroup_by_preamble is called when exiting the group_by_preamble production.
	ExitGroup_by_preamble(c *Group_by_preambleContext)

	// ExitOpt_and_order is called when exiting the opt_and_order production.
	ExitOpt_and_order(c *Opt_and_orderContext)

	// ExitHint is called when exiting the hint production.
	ExitHint(c *HintContext)

	// ExitHint_with_body is called when exiting the hint_with_body production.
	ExitHint_with_body(c *Hint_with_bodyContext)

	// ExitHint_with_body_prefix is called when exiting the hint_with_body_prefix production.
	ExitHint_with_body_prefix(c *Hint_with_body_prefixContext)

	// ExitHint_entry is called when exiting the hint_entry production.
	ExitHint_entry(c *Hint_entryContext)

	// ExitIdentifier_in_hints is called when exiting the identifier_in_hints production.
	ExitIdentifier_in_hints(c *Identifier_in_hintsContext)

	// ExitExtra_identifier_in_hints_name is called when exiting the extra_identifier_in_hints_name production.
	ExitExtra_identifier_in_hints_name(c *Extra_identifier_in_hints_nameContext)

	// ExitGrouping_item is called when exiting the grouping_item production.
	ExitGrouping_item(c *Grouping_itemContext)

	// ExitGrouping_set_list is called when exiting the grouping_set_list production.
	ExitGrouping_set_list(c *Grouping_set_listContext)

	// ExitGrouping_set is called when exiting the grouping_set production.
	ExitGrouping_set(c *Grouping_setContext)

	// ExitCube_list is called when exiting the cube_list production.
	ExitCube_list(c *Cube_listContext)

	// ExitRollup_list is called when exiting the rollup_list production.
	ExitRollup_list(c *Rollup_listContext)

	// ExitOpt_as_alias_with_required_as is called when exiting the opt_as_alias_with_required_as production.
	ExitOpt_as_alias_with_required_as(c *Opt_as_alias_with_required_asContext)

	// ExitOpt_grouping_item_order is called when exiting the opt_grouping_item_order production.
	ExitOpt_grouping_item_order(c *Opt_grouping_item_orderContext)

	// ExitOpt_selection_item_order is called when exiting the opt_selection_item_order production.
	ExitOpt_selection_item_order(c *Opt_selection_item_orderContext)

	// ExitAsc_or_desc is called when exiting the asc_or_desc production.
	ExitAsc_or_desc(c *Asc_or_descContext)

	// ExitNull_order is called when exiting the null_order production.
	ExitNull_order(c *Null_orderContext)

	// ExitFunction_name_from_keyword is called when exiting the function_name_from_keyword production.
	ExitFunction_name_from_keyword(c *Function_name_from_keywordContext)

	// ExitReplace_fields_expression is called when exiting the replace_fields_expression production.
	ExitReplace_fields_expression(c *Replace_fields_expressionContext)

	// ExitReplace_fields_prefix is called when exiting the replace_fields_prefix production.
	ExitReplace_fields_prefix(c *Replace_fields_prefixContext)

	// ExitReplace_fields_arg is called when exiting the replace_fields_arg production.
	ExitReplace_fields_arg(c *Replace_fields_argContext)

	// ExitGeneralized_path_expression is called when exiting the generalized_path_expression production.
	ExitGeneralized_path_expression(c *Generalized_path_expressionContext)

	// ExitGeneralized_extension_path is called when exiting the generalized_extension_path production.
	ExitGeneralized_extension_path(c *Generalized_extension_pathContext)

	// ExitWith_expression is called when exiting the with_expression production.
	ExitWith_expression(c *With_expressionContext)

	// ExitWith_expression_variable_prefix is called when exiting the with_expression_variable_prefix production.
	ExitWith_expression_variable_prefix(c *With_expression_variable_prefixContext)

	// ExitWith_expression_variable is called when exiting the with_expression_variable production.
	ExitWith_expression_variable(c *With_expression_variableContext)

	// ExitExtract_expression is called when exiting the extract_expression production.
	ExitExtract_expression(c *Extract_expressionContext)

	// ExitExtract_expression_base is called when exiting the extract_expression_base production.
	ExitExtract_expression_base(c *Extract_expression_baseContext)

	// ExitOpt_format is called when exiting the opt_format production.
	ExitOpt_format(c *Opt_formatContext)

	// ExitOpt_at_time_zone is called when exiting the opt_at_time_zone production.
	ExitOpt_at_time_zone(c *Opt_at_time_zoneContext)

	// ExitCast_expression is called when exiting the cast_expression production.
	ExitCast_expression(c *Cast_expressionContext)

	// ExitCase_expression is called when exiting the case_expression production.
	ExitCase_expression(c *Case_expressionContext)

	// ExitCase_expression_prefix is called when exiting the case_expression_prefix production.
	ExitCase_expression_prefix(c *Case_expression_prefixContext)

	// ExitCase_value_expression_prefix is called when exiting the case_value_expression_prefix production.
	ExitCase_value_expression_prefix(c *Case_value_expression_prefixContext)

	// ExitCase_no_value_expression_prefix is called when exiting the case_no_value_expression_prefix production.
	ExitCase_no_value_expression_prefix(c *Case_no_value_expression_prefixContext)

	// ExitStruct_braced_constructor is called when exiting the struct_braced_constructor production.
	ExitStruct_braced_constructor(c *Struct_braced_constructorContext)

	// ExitBraced_new_constructor is called when exiting the braced_new_constructor production.
	ExitBraced_new_constructor(c *Braced_new_constructorContext)

	// ExitBraced_constructor is called when exiting the braced_constructor production.
	ExitBraced_constructor(c *Braced_constructorContext)

	// ExitBraced_constructor_start is called when exiting the braced_constructor_start production.
	ExitBraced_constructor_start(c *Braced_constructor_startContext)

	// ExitBraced_constructor_prefix is called when exiting the braced_constructor_prefix production.
	ExitBraced_constructor_prefix(c *Braced_constructor_prefixContext)

	// ExitBraced_constructor_field is called when exiting the braced_constructor_field production.
	ExitBraced_constructor_field(c *Braced_constructor_fieldContext)

	// ExitBraced_constructor_lhs is called when exiting the braced_constructor_lhs production.
	ExitBraced_constructor_lhs(c *Braced_constructor_lhsContext)

	// ExitBraced_constructor_field_value is called when exiting the braced_constructor_field_value production.
	ExitBraced_constructor_field_value(c *Braced_constructor_field_valueContext)

	// ExitBraced_constructor_extension is called when exiting the braced_constructor_extension production.
	ExitBraced_constructor_extension(c *Braced_constructor_extensionContext)

	// ExitNew_constructor is called when exiting the new_constructor production.
	ExitNew_constructor(c *New_constructorContext)

	// ExitNew_constructor_prefix is called when exiting the new_constructor_prefix production.
	ExitNew_constructor_prefix(c *New_constructor_prefixContext)

	// ExitNew_constructor_prefix_no_arg is called when exiting the new_constructor_prefix_no_arg production.
	ExitNew_constructor_prefix_no_arg(c *New_constructor_prefix_no_argContext)

	// ExitNew_constructor_arg is called when exiting the new_constructor_arg production.
	ExitNew_constructor_arg(c *New_constructor_argContext)

	// ExitArray_constructor is called when exiting the array_constructor production.
	ExitArray_constructor(c *Array_constructorContext)

	// ExitArray_constructor_prefix is called when exiting the array_constructor_prefix production.
	ExitArray_constructor_prefix(c *Array_constructor_prefixContext)

	// ExitArray_constructor_prefix_no_expressions is called when exiting the array_constructor_prefix_no_expressions production.
	ExitArray_constructor_prefix_no_expressions(c *Array_constructor_prefix_no_expressionsContext)

	// ExitRange_literal is called when exiting the range_literal production.
	ExitRange_literal(c *Range_literalContext)

	// ExitRange_type is called when exiting the range_type production.
	ExitRange_type(c *Range_typeContext)

	// ExitType is called when exiting the type production.
	ExitType(c *TypeContext)

	// ExitCollate_clause is called when exiting the collate_clause production.
	ExitCollate_clause(c *Collate_clauseContext)

	// ExitString_literal_or_parameter is called when exiting the string_literal_or_parameter production.
	ExitString_literal_or_parameter(c *String_literal_or_parameterContext)

	// ExitSystem_variable_expression is called when exiting the system_variable_expression production.
	ExitSystem_variable_expression(c *System_variable_expressionContext)

	// ExitParameter_expression is called when exiting the parameter_expression production.
	ExitParameter_expression(c *Parameter_expressionContext)

	// ExitNamed_parameter_expression is called when exiting the named_parameter_expression production.
	ExitNamed_parameter_expression(c *Named_parameter_expressionContext)

	// ExitOpt_type_parameters is called when exiting the opt_type_parameters production.
	ExitOpt_type_parameters(c *Opt_type_parametersContext)

	// ExitType_parameters_prefix is called when exiting the type_parameters_prefix production.
	ExitType_parameters_prefix(c *Type_parameters_prefixContext)

	// ExitType_parameter is called when exiting the type_parameter production.
	ExitType_parameter(c *Type_parameterContext)

	// ExitRaw_type is called when exiting the raw_type production.
	ExitRaw_type(c *Raw_typeContext)

	// ExitMap_type is called when exiting the map_type production.
	ExitMap_type(c *Map_typeContext)

	// ExitFunction_type is called when exiting the function_type production.
	ExitFunction_type(c *Function_typeContext)

	// ExitFunction_type_prefix is called when exiting the function_type_prefix production.
	ExitFunction_type_prefix(c *Function_type_prefixContext)

	// ExitType_name is called when exiting the type_name production.
	ExitType_name(c *Type_nameContext)

	// ExitPath_expression is called when exiting the path_expression production.
	ExitPath_expression(c *Path_expressionContext)

	// ExitIdentifier is called when exiting the identifier production.
	ExitIdentifier(c *IdentifierContext)

	// ExitKeyword_as_identifier is called when exiting the keyword_as_identifier production.
	ExitKeyword_as_identifier(c *Keyword_as_identifierContext)

	// ExitCommon_keyword_as_identifier is called when exiting the common_keyword_as_identifier production.
	ExitCommon_keyword_as_identifier(c *Common_keyword_as_identifierContext)

	// ExitToken_identifier is called when exiting the token_identifier production.
	ExitToken_identifier(c *Token_identifierContext)

	// ExitStruct_type is called when exiting the struct_type production.
	ExitStruct_type(c *Struct_typeContext)

	// ExitStruct_type_prefix is called when exiting the struct_type_prefix production.
	ExitStruct_type_prefix(c *Struct_type_prefixContext)

	// ExitStruct_field is called when exiting the struct_field production.
	ExitStruct_field(c *Struct_fieldContext)

	// ExitArray_type is called when exiting the array_type production.
	ExitArray_type(c *Array_typeContext)

	// ExitTemplate_type_open is called when exiting the template_type_open production.
	ExitTemplate_type_open(c *Template_type_openContext)

	// ExitTemplate_type_close is called when exiting the template_type_close production.
	ExitTemplate_type_close(c *Template_type_closeContext)

	// ExitDate_or_time_literal is called when exiting the date_or_time_literal production.
	ExitDate_or_time_literal(c *Date_or_time_literalContext)

	// ExitDate_or_time_literal_kind is called when exiting the date_or_time_literal_kind production.
	ExitDate_or_time_literal_kind(c *Date_or_time_literal_kindContext)

	// ExitFloating_point_literal is called when exiting the floating_point_literal production.
	ExitFloating_point_literal(c *Floating_point_literalContext)

	// ExitJson_literal is called when exiting the json_literal production.
	ExitJson_literal(c *Json_literalContext)

	// ExitBignumeric_literal is called when exiting the bignumeric_literal production.
	ExitBignumeric_literal(c *Bignumeric_literalContext)

	// ExitBignumeric_literal_prefix is called when exiting the bignumeric_literal_prefix production.
	ExitBignumeric_literal_prefix(c *Bignumeric_literal_prefixContext)

	// ExitNumeric_literal is called when exiting the numeric_literal production.
	ExitNumeric_literal(c *Numeric_literalContext)

	// ExitNumeric_literal_prefix is called when exiting the numeric_literal_prefix production.
	ExitNumeric_literal_prefix(c *Numeric_literal_prefixContext)

	// ExitInteger_literal is called when exiting the integer_literal production.
	ExitInteger_literal(c *Integer_literalContext)

	// ExitBytes_literal is called when exiting the bytes_literal production.
	ExitBytes_literal(c *Bytes_literalContext)

	// ExitNull_literal is called when exiting the null_literal production.
	ExitNull_literal(c *Null_literalContext)

	// ExitBoolean_literal is called when exiting the boolean_literal production.
	ExitBoolean_literal(c *Boolean_literalContext)

	// ExitString_literal is called when exiting the string_literal production.
	ExitString_literal(c *String_literalContext)

	// ExitString_literal_component is called when exiting the string_literal_component production.
	ExitString_literal_component(c *String_literal_componentContext)

	// ExitBytes_literal_component is called when exiting the bytes_literal_component production.
	ExitBytes_literal_component(c *Bytes_literal_componentContext)
}
