// Code generated from GoogleSQLParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package googlesql // GoogleSQLParser
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by GoogleSQLParser.
type GoogleSQLParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by GoogleSQLParser#root.
	VisitRoot(ctx *RootContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#stmts.
	VisitStmts(ctx *StmtsContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#unterminated_sql_statement.
	VisitUnterminated_sql_statement(ctx *Unterminated_sql_statementContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#sql_statement_body.
	VisitSql_statement_body(ctx *Sql_statement_bodyContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#gql_statement.
	VisitGql_statement(ctx *Gql_statementContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#graph_operation_block.
	VisitGraph_operation_block(ctx *Graph_operation_blockContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#graph_composite_query_block.
	VisitGraph_composite_query_block(ctx *Graph_composite_query_blockContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#graph_composite_query_prefix.
	VisitGraph_composite_query_prefix(ctx *Graph_composite_query_prefixContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#graph_set_operation_metadata.
	VisitGraph_set_operation_metadata(ctx *Graph_set_operation_metadataContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#graph_linear_query_operation.
	VisitGraph_linear_query_operation(ctx *Graph_linear_query_operationContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#graph_linear_operator_list.
	VisitGraph_linear_operator_list(ctx *Graph_linear_operator_listContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#graph_linear_operator.
	VisitGraph_linear_operator(ctx *Graph_linear_operatorContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#graph_sample_clause.
	VisitGraph_sample_clause(ctx *Graph_sample_clauseContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_graph_sample_clause_suffix.
	VisitOpt_graph_sample_clause_suffix(ctx *Opt_graph_sample_clause_suffixContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#graph_for_operator.
	VisitGraph_for_operator(ctx *Graph_for_operatorContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_with_offset_and_alias_with_required_as.
	VisitOpt_with_offset_and_alias_with_required_as(ctx *Opt_with_offset_and_alias_with_required_asContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#graph_with_operator.
	VisitGraph_with_operator(ctx *Graph_with_operatorContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#graph_page_operator.
	VisitGraph_page_operator(ctx *Graph_page_operatorContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#graph_order_by_operator.
	VisitGraph_order_by_operator(ctx *Graph_order_by_operatorContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#graph_filter_operator.
	VisitGraph_filter_operator(ctx *Graph_filter_operatorContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#graph_let_operator.
	VisitGraph_let_operator(ctx *Graph_let_operatorContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#graph_let_variable_definition_list.
	VisitGraph_let_variable_definition_list(ctx *Graph_let_variable_definition_listContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#graph_let_variable_definition.
	VisitGraph_let_variable_definition(ctx *Graph_let_variable_definitionContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#graph_optional_match_operator.
	VisitGraph_optional_match_operator(ctx *Graph_optional_match_operatorContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#graph_match_operator.
	VisitGraph_match_operator(ctx *Graph_match_operatorContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#graph_pattern.
	VisitGraph_pattern(ctx *Graph_patternContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#graph_path_pattern_list.
	VisitGraph_path_pattern_list(ctx *Graph_path_pattern_listContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#graph_path_pattern.
	VisitGraph_path_pattern(ctx *Graph_path_patternContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#graph_path_pattern_expr.
	VisitGraph_path_pattern_expr(ctx *Graph_path_pattern_exprContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#graph_path_factor.
	VisitGraph_path_factor(ctx *Graph_path_factorContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#graph_quantified_path_primary.
	VisitGraph_quantified_path_primary(ctx *Graph_quantified_path_primaryContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#graph_path_primary.
	VisitGraph_path_primary(ctx *Graph_path_primaryContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#graph_parenthesized_path_pattern.
	VisitGraph_parenthesized_path_pattern(ctx *Graph_parenthesized_path_patternContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#graph_element_pattern.
	VisitGraph_element_pattern(ctx *Graph_element_patternContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#graph_edge_pattern.
	VisitGraph_edge_pattern(ctx *Graph_edge_patternContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#graph_node_pattern.
	VisitGraph_node_pattern(ctx *Graph_node_patternContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#graph_element_pattern_filler.
	VisitGraph_element_pattern_filler(ctx *Graph_element_pattern_fillerContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#graph_property_specification.
	VisitGraph_property_specification(ctx *Graph_property_specificationContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#graph_property_name_and_value.
	VisitGraph_property_name_and_value(ctx *Graph_property_name_and_valueContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_is_label_expression.
	VisitOpt_is_label_expression(ctx *Opt_is_label_expressionContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#label_expression.
	VisitLabel_expression(ctx *Label_expressionContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#label_primary.
	VisitLabel_primary(ctx *Label_primaryContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#parenthesized_label_expression.
	VisitParenthesized_label_expression(ctx *Parenthesized_label_expressionContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_graph_element_identifier.
	VisitOpt_graph_element_identifier(ctx *Opt_graph_element_identifierContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_graph_path_mode_prefix.
	VisitOpt_graph_path_mode_prefix(ctx *Opt_graph_path_mode_prefixContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#path_or_paths.
	VisitPath_or_paths(ctx *Path_or_pathsContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_graph_path_mode.
	VisitOpt_graph_path_mode(ctx *Opt_graph_path_modeContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_graph_search_prefix.
	VisitOpt_graph_search_prefix(ctx *Opt_graph_search_prefixContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_path_variable_assignment.
	VisitOpt_path_variable_assignment(ctx *Opt_path_variable_assignmentContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#graph_identifier.
	VisitGraph_identifier(ctx *Graph_identifierContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#graph_return_operator.
	VisitGraph_return_operator(ctx *Graph_return_operatorContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#graph_page_clause.
	VisitGraph_page_clause(ctx *Graph_page_clauseContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#graph_order_by_clause.
	VisitGraph_order_by_clause(ctx *Graph_order_by_clauseContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#graph_ordering_expression.
	VisitGraph_ordering_expression(ctx *Graph_ordering_expressionContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_graph_asc_or_desc.
	VisitOpt_graph_asc_or_desc(ctx *Opt_graph_asc_or_descContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#graph_return_item_list.
	VisitGraph_return_item_list(ctx *Graph_return_item_listContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#graph_return_item.
	VisitGraph_return_item(ctx *Graph_return_itemContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#undrop_statement.
	VisitUndrop_statement(ctx *Undrop_statementContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#module_statement.
	VisitModule_statement(ctx *Module_statementContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#import_statement.
	VisitImport_statement(ctx *Import_statementContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_as_or_into_alias.
	VisitOpt_as_or_into_alias(ctx *Opt_as_or_into_aliasContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#path_expression_or_string.
	VisitPath_expression_or_string(ctx *Path_expression_or_stringContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#import_type.
	VisitImport_type(ctx *Import_typeContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#call_statement.
	VisitCall_statement(ctx *Call_statementContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#drop_statement.
	VisitDrop_statement(ctx *Drop_statementContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_drop_mode.
	VisitOpt_drop_mode(ctx *Opt_drop_modeContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#drop_all_row_access_policies_statement.
	VisitDrop_all_row_access_policies_statement(ctx *Drop_all_row_access_policies_statementContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#show_statement.
	VisitShow_statement(ctx *Show_statementContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_like_string_literal.
	VisitOpt_like_string_literal(ctx *Opt_like_string_literalContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#show_target.
	VisitShow_target(ctx *Show_targetContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#rename_statement.
	VisitRename_statement(ctx *Rename_statementContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#revoke_statement.
	VisitRevoke_statement(ctx *Revoke_statementContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#grant_statement.
	VisitGrant_statement(ctx *Grant_statementContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#privileges.
	VisitPrivileges(ctx *PrivilegesContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#export_metadata_statement.
	VisitExport_metadata_statement(ctx *Export_metadata_statementContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#export_model_statement.
	VisitExport_model_statement(ctx *Export_model_statementContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#export_data_statement.
	VisitExport_data_statement(ctx *Export_data_statementContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#export_data_no_query.
	VisitExport_data_no_query(ctx *Export_data_no_queryContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#explain_statement.
	VisitExplain_statement(ctx *Explain_statementContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#execute_immediate.
	VisitExecute_immediate(ctx *Execute_immediateContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_execute_into_clause.
	VisitOpt_execute_into_clause(ctx *Opt_execute_into_clauseContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_execute_using_clause.
	VisitOpt_execute_using_clause(ctx *Opt_execute_using_clauseContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#execute_using_argument_list.
	VisitExecute_using_argument_list(ctx *Execute_using_argument_listContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#execute_using_argument.
	VisitExecute_using_argument(ctx *Execute_using_argumentContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#describe_statement.
	VisitDescribe_statement(ctx *Describe_statementContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#describe_info.
	VisitDescribe_info(ctx *Describe_infoContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_from_path_expression.
	VisitOpt_from_path_expression(ctx *Opt_from_path_expressionContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#describe_keyword.
	VisitDescribe_keyword(ctx *Describe_keywordContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#define_table_statement.
	VisitDefine_table_statement(ctx *Define_table_statementContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#create_entity_statement.
	VisitCreate_entity_statement(ctx *Create_entity_statementContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_generic_entity_body.
	VisitOpt_generic_entity_body(ctx *Opt_generic_entity_bodyContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#create_view_statement.
	VisitCreate_view_statement(ctx *Create_view_statementContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#query_or_replica_source.
	VisitQuery_or_replica_source(ctx *Query_or_replica_sourceContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#column_with_options_list.
	VisitColumn_with_options_list(ctx *Column_with_options_listContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#column_with_options.
	VisitColumn_with_options(ctx *Column_with_optionsContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#create_table_statement.
	VisitCreate_table_statement(ctx *Create_table_statementContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_ttl_clause.
	VisitOpt_ttl_clause(ctx *Opt_ttl_clauseContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_copy_table.
	VisitOpt_copy_table(ctx *Opt_copy_tableContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#copy_data_source.
	VisitCopy_data_source(ctx *Copy_data_sourceContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_clone_table.
	VisitOpt_clone_table(ctx *Opt_clone_tableContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_spanner_table_options.
	VisitOpt_spanner_table_options(ctx *Opt_spanner_table_optionsContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_spanner_interleave_in_parent_clause.
	VisitOpt_spanner_interleave_in_parent_clause(ctx *Opt_spanner_interleave_in_parent_clauseContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#spanner_primary_key.
	VisitSpanner_primary_key(ctx *Spanner_primary_keyContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#create_table_function_statement.
	VisitCreate_table_function_statement(ctx *Create_table_function_statementContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_as_query_or_string.
	VisitOpt_as_query_or_string(ctx *Opt_as_query_or_stringContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#unordered_language_options.
	VisitUnordered_language_options(ctx *Unordered_language_optionsContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_function_parameters.
	VisitOpt_function_parameters(ctx *Opt_function_parametersContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#create_snapshot_statement.
	VisitCreate_snapshot_statement(ctx *Create_snapshot_statementContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#create_external_schema_statement.
	VisitCreate_external_schema_statement(ctx *Create_external_schema_statementContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#create_schema_statement.
	VisitCreate_schema_statement(ctx *Create_schema_statementContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#create_property_graph_statement.
	VisitCreate_property_graph_statement(ctx *Create_property_graph_statementContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_edge_table_clause.
	VisitOpt_edge_table_clause(ctx *Opt_edge_table_clauseContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#element_table_list.
	VisitElement_table_list(ctx *Element_table_listContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#element_table_definition.
	VisitElement_table_definition(ctx *Element_table_definitionContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_label_and_properties_clause.
	VisitOpt_label_and_properties_clause(ctx *Opt_label_and_properties_clauseContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#label_and_properties_list.
	VisitLabel_and_properties_list(ctx *Label_and_properties_listContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#label_and_properties.
	VisitLabel_and_properties(ctx *Label_and_propertiesContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#properties_clause.
	VisitProperties_clause(ctx *Properties_clauseContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#derived_property_list.
	VisitDerived_property_list(ctx *Derived_property_listContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#derived_property.
	VisitDerived_property(ctx *Derived_propertyContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_except_column_list.
	VisitOpt_except_column_list(ctx *Opt_except_column_listContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#properties_all_columns.
	VisitProperties_all_columns(ctx *Properties_all_columnsContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_dest_node_table_clause.
	VisitOpt_dest_node_table_clause(ctx *Opt_dest_node_table_clauseContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_source_node_table_clause.
	VisitOpt_source_node_table_clause(ctx *Opt_source_node_table_clauseContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_key_clause.
	VisitOpt_key_clause(ctx *Opt_key_clauseContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#create_model_statement.
	VisitCreate_model_statement(ctx *Create_model_statementContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_input_output_clause.
	VisitOpt_input_output_clause(ctx *Opt_input_output_clauseContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_transform_clause.
	VisitOpt_transform_clause(ctx *Opt_transform_clauseContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_as_query_or_aliased_query_list.
	VisitOpt_as_query_or_aliased_query_list(ctx *Opt_as_query_or_aliased_query_listContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#aliased_query_list.
	VisitAliased_query_list(ctx *Aliased_query_listContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#as_query.
	VisitAs_query(ctx *As_queryContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#create_external_table_function_statement.
	VisitCreate_external_table_function_statement(ctx *Create_external_table_function_statementContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#create_external_table_statement.
	VisitCreate_external_table_statement(ctx *Create_external_table_statementContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_default_collate_clause.
	VisitOpt_default_collate_clause(ctx *Opt_default_collate_clauseContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_like_path_expression.
	VisitOpt_like_path_expression(ctx *Opt_like_path_expressionContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#create_row_access_policy_statement.
	VisitCreate_row_access_policy_statement(ctx *Create_row_access_policy_statementContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#filter_using_clause.
	VisitFilter_using_clause(ctx *Filter_using_clauseContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#create_row_access_policy_grant_to_clause.
	VisitCreate_row_access_policy_grant_to_clause(ctx *Create_row_access_policy_grant_to_clauseContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#create_privilege_restriction_statement.
	VisitCreate_privilege_restriction_statement(ctx *Create_privilege_restriction_statementContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#restrict_to_clause.
	VisitRestrict_to_clause(ctx *Restrict_to_clauseContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#possibly_empty_grantee_list.
	VisitPossibly_empty_grantee_list(ctx *Possibly_empty_grantee_listContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#create_index_statement.
	VisitCreate_index_statement(ctx *Create_index_statementContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_create_index_statement_suffix.
	VisitOpt_create_index_statement_suffix(ctx *Opt_create_index_statement_suffixContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#spanner_index_interleave_clause.
	VisitSpanner_index_interleave_clause(ctx *Spanner_index_interleave_clauseContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#index_storing_list.
	VisitIndex_storing_list(ctx *Index_storing_listContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#index_storing_expression_list.
	VisitIndex_storing_expression_list(ctx *Index_storing_expression_listContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#index_order_by_and_options.
	VisitIndex_order_by_and_options(ctx *Index_order_by_and_optionsContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#index_all_columns.
	VisitIndex_all_columns(ctx *Index_all_columnsContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_with_column_options.
	VisitOpt_with_column_options(ctx *Opt_with_column_optionsContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#all_column_column_options.
	VisitAll_column_column_options(ctx *All_column_column_optionsContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#column_ordering_and_options_expr.
	VisitColumn_ordering_and_options_expr(ctx *Column_ordering_and_options_exprContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#index_unnest_expression_list.
	VisitIndex_unnest_expression_list(ctx *Index_unnest_expression_listContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#unnest_expression_with_opt_alias_and_offset.
	VisitUnnest_expression_with_opt_alias_and_offset(ctx *Unnest_expression_with_opt_alias_and_offsetContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#on_path_expression.
	VisitOn_path_expression(ctx *On_path_expressionContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#index_type.
	VisitIndex_type(ctx *Index_typeContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_spanner_null_filtered.
	VisitOpt_spanner_null_filtered(ctx *Opt_spanner_null_filteredContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#create_procedure_statement.
	VisitCreate_procedure_statement(ctx *Create_procedure_statementContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#begin_end_block_or_language_as_code.
	VisitBegin_end_block_or_language_as_code(ctx *Begin_end_block_or_language_as_codeContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#begin_end_block.
	VisitBegin_end_block(ctx *Begin_end_blockContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_exception_handler.
	VisitOpt_exception_handler(ctx *Opt_exception_handlerContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#statement_list.
	VisitStatement_list(ctx *Statement_listContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#unterminated_non_empty_statement_list.
	VisitUnterminated_non_empty_statement_list(ctx *Unterminated_non_empty_statement_listContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#unterminated_statement.
	VisitUnterminated_statement(ctx *Unterminated_statementContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#unterminated_script_statement.
	VisitUnterminated_script_statement(ctx *Unterminated_script_statementContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#label.
	VisitLabel(ctx *LabelContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#unterminated_unlabeled_script_statement.
	VisitUnterminated_unlabeled_script_statement(ctx *Unterminated_unlabeled_script_statementContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#for_in_statement.
	VisitFor_in_statement(ctx *For_in_statementContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#repeat_statement.
	VisitRepeat_statement(ctx *Repeat_statementContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#until_clause.
	VisitUntil_clause(ctx *Until_clauseContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#loop_statement.
	VisitLoop_statement(ctx *Loop_statementContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#while_statement.
	VisitWhile_statement(ctx *While_statementContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#raise_statement.
	VisitRaise_statement(ctx *Raise_statementContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#return_statement.
	VisitReturn_statement(ctx *Return_statementContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#continue_statement.
	VisitContinue_statement(ctx *Continue_statementContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#variable_declaration.
	VisitVariable_declaration(ctx *Variable_declarationContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#break_statement.
	VisitBreak_statement(ctx *Break_statementContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#case_statement.
	VisitCase_statement(ctx *Case_statementContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#when_then_clauses.
	VisitWhen_then_clauses(ctx *When_then_clausesContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#if_statement.
	VisitIf_statement(ctx *If_statementContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#elseif_clauses.
	VisitElseif_clauses(ctx *Elseif_clausesContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_else.
	VisitOpt_else(ctx *Opt_elseContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_as_code.
	VisitOpt_as_code(ctx *Opt_as_codeContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_external_security_clause.
	VisitOpt_external_security_clause(ctx *Opt_external_security_clauseContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#external_security_clause_kind.
	VisitExternal_security_clause_kind(ctx *External_security_clause_kindContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#procedure_parameters.
	VisitProcedure_parameters(ctx *Procedure_parametersContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#procedure_parameter.
	VisitProcedure_parameter(ctx *Procedure_parameterContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#procedure_parameter_termination.
	VisitProcedure_parameter_termination(ctx *Procedure_parameter_terminationContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_procedure_parameter_mode.
	VisitOpt_procedure_parameter_mode(ctx *Opt_procedure_parameter_modeContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#create_function_statement.
	VisitCreate_function_statement(ctx *Create_function_statementContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_determinism_level.
	VisitOpt_determinism_level(ctx *Opt_determinism_levelContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_sql_security_clause.
	VisitOpt_sql_security_clause(ctx *Opt_sql_security_clauseContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#sql_security_clause_kind.
	VisitSql_security_clause_kind(ctx *Sql_security_clause_kindContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#as_sql_function_body_or_string.
	VisitAs_sql_function_body_or_string(ctx *As_sql_function_body_or_stringContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#sql_function_body.
	VisitSql_function_body(ctx *Sql_function_bodyContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#unordered_options_body.
	VisitUnordered_options_body(ctx *Unordered_options_bodyContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_language_or_remote_with_connection.
	VisitOpt_language_or_remote_with_connection(ctx *Opt_language_or_remote_with_connectionContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#language.
	VisitLanguage(ctx *LanguageContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#remote_with_connection_clause.
	VisitRemote_with_connection_clause(ctx *Remote_with_connection_clauseContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#with_connection_clause.
	VisitWith_connection_clause(ctx *With_connection_clauseContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_function_returns.
	VisitOpt_function_returns(ctx *Opt_function_returnsContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_returns.
	VisitOpt_returns(ctx *Opt_returnsContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#function_declaration.
	VisitFunction_declaration(ctx *Function_declarationContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#function_parameters.
	VisitFunction_parameters(ctx *Function_parametersContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#function_parameter.
	VisitFunction_parameter(ctx *Function_parameterContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_not_aggregate.
	VisitOpt_not_aggregate(ctx *Opt_not_aggregateContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_default_expression.
	VisitOpt_default_expression(ctx *Opt_default_expressionContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#type_or_tvf_schema.
	VisitType_or_tvf_schema(ctx *Type_or_tvf_schemaContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#tvf_schema.
	VisitTvf_schema(ctx *Tvf_schemaContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#tvf_schema_column.
	VisitTvf_schema_column(ctx *Tvf_schema_columnContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#templated_parameter_type.
	VisitTemplated_parameter_type(ctx *Templated_parameter_typeContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#templated_parameter_kind.
	VisitTemplated_parameter_kind(ctx *Templated_parameter_kindContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_aggregate.
	VisitOpt_aggregate(ctx *Opt_aggregateContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#create_database_statement.
	VisitCreate_database_statement(ctx *Create_database_statementContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#create_connection_statement.
	VisitCreate_connection_statement(ctx *Create_connection_statementContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#create_constant_statement.
	VisitCreate_constant_statement(ctx *Create_constant_statementContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_or_replace.
	VisitOpt_or_replace(ctx *Opt_or_replaceContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_create_scope.
	VisitOpt_create_scope(ctx *Opt_create_scopeContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#run_batch_statement.
	VisitRun_batch_statement(ctx *Run_batch_statementContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#abort_batch_statement.
	VisitAbort_batch_statement(ctx *Abort_batch_statementContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#start_batch_statement.
	VisitStart_batch_statement(ctx *Start_batch_statementContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#rollback_statement.
	VisitRollback_statement(ctx *Rollback_statementContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#commit_statement.
	VisitCommit_statement(ctx *Commit_statementContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#set_statement.
	VisitSet_statement(ctx *Set_statementContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#identifier_list.
	VisitIdentifier_list(ctx *Identifier_listContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#begin_statement.
	VisitBegin_statement(ctx *Begin_statementContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#begin_transaction_keywords.
	VisitBegin_transaction_keywords(ctx *Begin_transaction_keywordsContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#transaction_mode_list.
	VisitTransaction_mode_list(ctx *Transaction_mode_listContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#transaction_mode.
	VisitTransaction_mode(ctx *Transaction_modeContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#truncate_statement.
	VisitTruncate_statement(ctx *Truncate_statementContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#merge_statement.
	VisitMerge_statement(ctx *Merge_statementContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#merge_source.
	VisitMerge_source(ctx *Merge_sourceContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#merge_when_clause.
	VisitMerge_when_clause(ctx *Merge_when_clauseContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#merge_action.
	VisitMerge_action(ctx *Merge_actionContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#merge_insert_value_list_or_source_row.
	VisitMerge_insert_value_list_or_source_row(ctx *Merge_insert_value_list_or_source_rowContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#by_target.
	VisitBy_target(ctx *By_targetContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_and_expression.
	VisitOpt_and_expression(ctx *Opt_and_expressionContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#statement_level_hint.
	VisitStatement_level_hint(ctx *Statement_level_hintContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#query_statement.
	VisitQuery_statement(ctx *Query_statementContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#dml_statement.
	VisitDml_statement(ctx *Dml_statementContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#update_statement.
	VisitUpdate_statement(ctx *Update_statementContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#delete_statement.
	VisitDelete_statement(ctx *Delete_statementContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#insert_statement.
	VisitInsert_statement(ctx *Insert_statementContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#on_conflict_clause.
	VisitOn_conflict_clause(ctx *On_conflict_clauseContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_where_expression.
	VisitOpt_where_expression(ctx *Opt_where_expressionContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_conflict_target.
	VisitOpt_conflict_target(ctx *Opt_conflict_targetContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#update_item_list.
	VisitUpdate_item_list(ctx *Update_item_listContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#update_item.
	VisitUpdate_item(ctx *Update_itemContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#update_set_value.
	VisitUpdate_set_value(ctx *Update_set_valueContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#nested_dml_statement.
	VisitNested_dml_statement(ctx *Nested_dml_statementContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#insert_values_list_or_table_clause.
	VisitInsert_values_list_or_table_clause(ctx *Insert_values_list_or_table_clauseContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#table_clause_unreversed.
	VisitTable_clause_unreversed(ctx *Table_clause_unreversedContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#table_clause_no_keyword.
	VisitTable_clause_no_keyword(ctx *Table_clause_no_keywordContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_returning_clause.
	VisitOpt_returning_clause(ctx *Opt_returning_clauseContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_assert_rows_modified.
	VisitOpt_assert_rows_modified(ctx *Opt_assert_rows_modifiedContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#insert_values_or_query.
	VisitInsert_values_or_query(ctx *Insert_values_or_queryContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#insert_values_list.
	VisitInsert_values_list(ctx *Insert_values_listContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#insert_values_row.
	VisitInsert_values_row(ctx *Insert_values_rowContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#expression_or_default.
	VisitExpression_or_default(ctx *Expression_or_defaultContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#insert_statement_prefix.
	VisitInsert_statement_prefix(ctx *Insert_statement_prefixContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#maybe_dashed_generalized_path_expression.
	VisitMaybe_dashed_generalized_path_expression(ctx *Maybe_dashed_generalized_path_expressionContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_into.
	VisitOpt_into(ctx *Opt_intoContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_or_ignore_replace_update.
	VisitOpt_or_ignore_replace_update(ctx *Opt_or_ignore_replace_updateContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#alter_statement.
	VisitAlter_statement(ctx *Alter_statementContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#analyze_statement.
	VisitAnalyze_statement(ctx *Analyze_statementContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#assert_statement.
	VisitAssert_statement(ctx *Assert_statementContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#aux_load_data_statement.
	VisitAux_load_data_statement(ctx *Aux_load_data_statementContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#clone_data_statement.
	VisitClone_data_statement(ctx *Clone_data_statementContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#clone_data_source_list.
	VisitClone_data_source_list(ctx *Clone_data_source_listContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#clone_data_source.
	VisitClone_data_source(ctx *Clone_data_sourceContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_external_table_with_clauses.
	VisitOpt_external_table_with_clauses(ctx *Opt_external_table_with_clausesContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#with_partition_columns_clause.
	VisitWith_partition_columns_clause(ctx *With_partition_columns_clauseContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#aux_load_data_from_files_options_list.
	VisitAux_load_data_from_files_options_list(ctx *Aux_load_data_from_files_options_listContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#cluster_by_clause_prefix_no_hint.
	VisitCluster_by_clause_prefix_no_hint(ctx *Cluster_by_clause_prefix_no_hintContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#load_data_partitions_clause.
	VisitLoad_data_partitions_clause(ctx *Load_data_partitions_clauseContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#maybe_dashed_path_expression_with_scope.
	VisitMaybe_dashed_path_expression_with_scope(ctx *Maybe_dashed_path_expression_with_scopeContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#table_element_list.
	VisitTable_element_list(ctx *Table_element_listContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#table_element.
	VisitTable_element(ctx *Table_elementContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#table_constraint_definition.
	VisitTable_constraint_definition(ctx *Table_constraint_definitionContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#append_or_overwrite.
	VisitAppend_or_overwrite(ctx *Append_or_overwriteContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_description.
	VisitOpt_description(ctx *Opt_descriptionContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#table_and_column_info_list.
	VisitTable_and_column_info_list(ctx *Table_and_column_info_listContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#table_and_column_info.
	VisitTable_and_column_info(ctx *Table_and_column_infoContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#row_access_policy_alter_action_list.
	VisitRow_access_policy_alter_action_list(ctx *Row_access_policy_alter_action_listContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#row_access_policy_alter_action.
	VisitRow_access_policy_alter_action(ctx *Row_access_policy_alter_actionContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#grant_to_clause.
	VisitGrant_to_clause(ctx *Grant_to_clauseContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#grantee_list.
	VisitGrantee_list(ctx *Grantee_listContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#privilege_list.
	VisitPrivilege_list(ctx *Privilege_listContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#privilege.
	VisitPrivilege(ctx *PrivilegeContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#path_expression_list_with_parens.
	VisitPath_expression_list_with_parens(ctx *Path_expression_list_with_parensContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#privilege_name.
	VisitPrivilege_name(ctx *Privilege_nameContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#generic_entity_type.
	VisitGeneric_entity_type(ctx *Generic_entity_typeContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#generic_entity_type_unchecked.
	VisitGeneric_entity_type_unchecked(ctx *Generic_entity_type_uncheckedContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#schema_object_kind.
	VisitSchema_object_kind(ctx *Schema_object_kindContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#alter_action_list.
	VisitAlter_action_list(ctx *Alter_action_listContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#alter_action.
	VisitAlter_action(ctx *Alter_actionContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#spanner_set_on_delete_action.
	VisitSpanner_set_on_delete_action(ctx *Spanner_set_on_delete_actionContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#spanner_alter_column_action.
	VisitSpanner_alter_column_action(ctx *Spanner_alter_column_actionContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#spanner_generated_or_default.
	VisitSpanner_generated_or_default(ctx *Spanner_generated_or_defaultContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#generic_sub_entity_type.
	VisitGeneric_sub_entity_type(ctx *Generic_sub_entity_typeContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#sub_entity_type_identifier.
	VisitSub_entity_type_identifier(ctx *Sub_entity_type_identifierContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#fill_using_expression.
	VisitFill_using_expression(ctx *Fill_using_expressionContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#column_position.
	VisitColumn_position(ctx *Column_positionContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#table_column_definition.
	VisitTable_column_definition(ctx *Table_column_definitionContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#column_attributes.
	VisitColumn_attributes(ctx *Column_attributesContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#column_attribute.
	VisitColumn_attribute(ctx *Column_attributeContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#primary_key_column_attribute.
	VisitPrimary_key_column_attribute(ctx *Primary_key_column_attributeContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#foreign_key_column_attribute.
	VisitForeign_key_column_attribute(ctx *Foreign_key_column_attributeContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#hidden_column_attribute.
	VisitHidden_column_attribute(ctx *Hidden_column_attributeContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_constraint_identity.
	VisitOpt_constraint_identity(ctx *Opt_constraint_identityContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#table_column_schema.
	VisitTable_column_schema(ctx *Table_column_schemaContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_column_info.
	VisitOpt_column_info(ctx *Opt_column_infoContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#invalid_generated_column.
	VisitInvalid_generated_column(ctx *Invalid_generated_columnContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#invalid_default_column.
	VisitInvalid_default_column(ctx *Invalid_default_columnContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#default_column_info.
	VisitDefault_column_info(ctx *Default_column_infoContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#generated_column_info.
	VisitGenerated_column_info(ctx *Generated_column_infoContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#identity_column_info.
	VisitIdentity_column_info(ctx *Identity_column_infoContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_start_with.
	VisitOpt_start_with(ctx *Opt_start_withContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_increment_by.
	VisitOpt_increment_by(ctx *Opt_increment_byContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_maxvalue.
	VisitOpt_maxvalue(ctx *Opt_maxvalueContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_minvalue.
	VisitOpt_minvalue(ctx *Opt_minvalueContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_cycle.
	VisitOpt_cycle(ctx *Opt_cycleContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#signed_numeric_literal.
	VisitSigned_numeric_literal(ctx *Signed_numeric_literalContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#stored_mode.
	VisitStored_mode(ctx *Stored_modeContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#generated_mode.
	VisitGenerated_mode(ctx *Generated_modeContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#column_schema_inner.
	VisitColumn_schema_inner(ctx *Column_schema_innerContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#raw_column_schema_inner.
	VisitRaw_column_schema_inner(ctx *Raw_column_schema_innerContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#range_column_schema_inner.
	VisitRange_column_schema_inner(ctx *Range_column_schema_innerContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#struct_column_schema_inner.
	VisitStruct_column_schema_inner(ctx *Struct_column_schema_innerContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#struct_column_field.
	VisitStruct_column_field(ctx *Struct_column_fieldContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#simple_column_schema_inner.
	VisitSimple_column_schema_inner(ctx *Simple_column_schema_innerContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#array_column_schema_inner.
	VisitArray_column_schema_inner(ctx *Array_column_schema_innerContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#field_schema.
	VisitField_schema(ctx *Field_schemaContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_field_attributes.
	VisitOpt_field_attributes(ctx *Opt_field_attributesContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#not_null_column_attribute.
	VisitNot_null_column_attribute(ctx *Not_null_column_attributeContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#primary_key_or_table_constraint_spec.
	VisitPrimary_key_or_table_constraint_spec(ctx *Primary_key_or_table_constraint_specContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_if_not_exists.
	VisitOpt_if_not_exists(ctx *Opt_if_not_existsContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#primary_key_spec.
	VisitPrimary_key_spec(ctx *Primary_key_specContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#primary_key_element_list.
	VisitPrimary_key_element_list(ctx *Primary_key_element_listContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#primary_key_element.
	VisitPrimary_key_element(ctx *Primary_key_elementContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#table_constraint_spec.
	VisitTable_constraint_spec(ctx *Table_constraint_specContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#foreign_key_reference.
	VisitForeign_key_reference(ctx *Foreign_key_referenceContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_foreign_key_action.
	VisitOpt_foreign_key_action(ctx *Opt_foreign_key_actionContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#foreign_key_on_update.
	VisitForeign_key_on_update(ctx *Foreign_key_on_updateContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#foreign_key_on_delete.
	VisitForeign_key_on_delete(ctx *Foreign_key_on_deleteContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#foreign_key_action.
	VisitForeign_key_action(ctx *Foreign_key_actionContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_foreign_key_match.
	VisitOpt_foreign_key_match(ctx *Opt_foreign_key_matchContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#foreign_key_match_mode.
	VisitForeign_key_match_mode(ctx *Foreign_key_match_modeContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#column_list.
	VisitColumn_list(ctx *Column_listContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_options_list.
	VisitOpt_options_list(ctx *Opt_options_listContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#constraint_enforcement.
	VisitConstraint_enforcement(ctx *Constraint_enforcementContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#generic_entity_body.
	VisitGeneric_entity_body(ctx *Generic_entity_bodyContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_if_exists.
	VisitOpt_if_exists(ctx *Opt_if_existsContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#table_or_table_function.
	VisitTable_or_table_function(ctx *Table_or_table_functionContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#query.
	VisitQuery(ctx *QueryContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#query_without_pipe_operators.
	VisitQuery_without_pipe_operators(ctx *Query_without_pipe_operatorsContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#bad_keyword_after_from_query.
	VisitBad_keyword_after_from_query(ctx *Bad_keyword_after_from_queryContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#bad_keyword_after_from_query_allows_parens.
	VisitBad_keyword_after_from_query_allows_parens(ctx *Bad_keyword_after_from_query_allows_parensContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#with_clause_with_trailing_comma.
	VisitWith_clause_with_trailing_comma(ctx *With_clause_with_trailing_commaContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#select_or_from_keyword.
	VisitSelect_or_from_keyword(ctx *Select_or_from_keywordContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#query_primary_or_set_operation.
	VisitQuery_primary_or_set_operation(ctx *Query_primary_or_set_operationContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#query_set_operation.
	VisitQuery_set_operation(ctx *Query_set_operationContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#query_set_operation_prefix.
	VisitQuery_set_operation_prefix(ctx *Query_set_operation_prefixContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#query_set_operation_item.
	VisitQuery_set_operation_item(ctx *Query_set_operation_itemContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#query_primary.
	VisitQuery_primary(ctx *Query_primaryContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#set_operation_metadata.
	VisitSet_operation_metadata(ctx *Set_operation_metadataContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_column_match_suffix.
	VisitOpt_column_match_suffix(ctx *Opt_column_match_suffixContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_strict.
	VisitOpt_strict(ctx *Opt_strictContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#all_or_distinct.
	VisitAll_or_distinct(ctx *All_or_distinctContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#query_set_operation_type.
	VisitQuery_set_operation_type(ctx *Query_set_operation_typeContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_corresponding_outer_mode.
	VisitOpt_corresponding_outer_mode(ctx *Opt_corresponding_outer_modeContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_outer.
	VisitOpt_outer(ctx *Opt_outerContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#with_clause.
	VisitWith_clause(ctx *With_clauseContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#aliased_query.
	VisitAliased_query(ctx *Aliased_queryContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_aliased_query_modifiers.
	VisitOpt_aliased_query_modifiers(ctx *Opt_aliased_query_modifiersContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#recursion_depth_modifier.
	VisitRecursion_depth_modifier(ctx *Recursion_depth_modifierContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#possibly_unbounded_int_literal_or_parameter.
	VisitPossibly_unbounded_int_literal_or_parameter(ctx *Possibly_unbounded_int_literal_or_parameterContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#int_literal_or_parameter.
	VisitInt_literal_or_parameter(ctx *Int_literal_or_parameterContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#order_by_clause.
	VisitOrder_by_clause(ctx *Order_by_clauseContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#order_by_clause_prefix.
	VisitOrder_by_clause_prefix(ctx *Order_by_clause_prefixContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#ordering_expression.
	VisitOrdering_expression(ctx *Ordering_expressionContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#select.
	VisitSelect(ctx *SelectContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_clauses_following_from.
	VisitOpt_clauses_following_from(ctx *Opt_clauses_following_fromContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_clauses_following_where.
	VisitOpt_clauses_following_where(ctx *Opt_clauses_following_whereContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_clauses_following_group_by.
	VisitOpt_clauses_following_group_by(ctx *Opt_clauses_following_group_byContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#window_clause.
	VisitWindow_clause(ctx *Window_clauseContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#window_clause_prefix.
	VisitWindow_clause_prefix(ctx *Window_clause_prefixContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#window_definition.
	VisitWindow_definition(ctx *Window_definitionContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#where_clause.
	VisitWhere_clause(ctx *Where_clauseContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#having_clause.
	VisitHaving_clause(ctx *Having_clauseContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#group_by_clause.
	VisitGroup_by_clause(ctx *Group_by_clauseContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#group_by_all.
	VisitGroup_by_all(ctx *Group_by_allContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#select_clause.
	VisitSelect_clause(ctx *Select_clauseContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_select_as_clause.
	VisitOpt_select_as_clause(ctx *Opt_select_as_clauseContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_select_with.
	VisitOpt_select_with(ctx *Opt_select_withContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#from_clause.
	VisitFrom_clause(ctx *From_clauseContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#from_clause_contents.
	VisitFrom_clause_contents(ctx *From_clause_contentsContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#from_clause_contents_suffix.
	VisitFrom_clause_contents_suffix(ctx *From_clause_contents_suffixContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#table_primary.
	VisitTable_primary(ctx *Table_primaryContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#tvf_with_suffixes.
	VisitTvf_with_suffixes(ctx *Tvf_with_suffixesContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#pivot_or_unpivot_clause_and_aliases.
	VisitPivot_or_unpivot_clause_and_aliases(ctx *Pivot_or_unpivot_clause_and_aliasesContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#as_alias.
	VisitAs_alias(ctx *As_aliasContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#sample_clause.
	VisitSample_clause(ctx *Sample_clauseContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_sample_clause_suffix.
	VisitOpt_sample_clause_suffix(ctx *Opt_sample_clause_suffixContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#repeatable_clause.
	VisitRepeatable_clause(ctx *Repeatable_clauseContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#possibly_cast_int_literal_or_parameter.
	VisitPossibly_cast_int_literal_or_parameter(ctx *Possibly_cast_int_literal_or_parameterContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#cast_int_literal_or_parameter.
	VisitCast_int_literal_or_parameter(ctx *Cast_int_literal_or_parameterContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#sample_size.
	VisitSample_size(ctx *Sample_sizeContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#sample_size_value.
	VisitSample_size_value(ctx *Sample_size_valueContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#sample_size_unit.
	VisitSample_size_unit(ctx *Sample_size_unitContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#partition_by_clause_prefix_no_hint.
	VisitPartition_by_clause_prefix_no_hint(ctx *Partition_by_clause_prefix_no_hintContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#match_recognize_clause.
	VisitMatch_recognize_clause(ctx *Match_recognize_clauseContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#row_pattern_expr.
	VisitRow_pattern_expr(ctx *Row_pattern_exprContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#row_pattern_concatenation.
	VisitRow_pattern_concatenation(ctx *Row_pattern_concatenationContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#row_pattern_factor.
	VisitRow_pattern_factor(ctx *Row_pattern_factorContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#select_list_prefix_with_as_aliases.
	VisitSelect_list_prefix_with_as_aliases(ctx *Select_list_prefix_with_as_aliasesContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#select_column_expr_with_as_alias.
	VisitSelect_column_expr_with_as_alias(ctx *Select_column_expr_with_as_aliasContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#table_subquery.
	VisitTable_subquery(ctx *Table_subqueryContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#join.
	VisitJoin(ctx *JoinContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#join_item.
	VisitJoin_item(ctx *Join_itemContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#on_or_using_clause_list.
	VisitOn_or_using_clause_list(ctx *On_or_using_clause_listContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#on_or_using_clause.
	VisitOn_or_using_clause(ctx *On_or_using_clauseContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#using_clause.
	VisitUsing_clause(ctx *Using_clauseContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#join_hint.
	VisitJoin_hint(ctx *Join_hintContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#table_path_expression.
	VisitTable_path_expression(ctx *Table_path_expressionContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_at_system_time.
	VisitOpt_at_system_time(ctx *Opt_at_system_timeContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_with_offset_and_alias.
	VisitOpt_with_offset_and_alias(ctx *Opt_with_offset_and_aliasContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_pivot_or_unpivot_clause_and_alias.
	VisitOpt_pivot_or_unpivot_clause_and_alias(ctx *Opt_pivot_or_unpivot_clause_and_aliasContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#table_path_expression_base.
	VisitTable_path_expression_base(ctx *Table_path_expression_baseContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#maybe_slashed_or_dashed_path_expression.
	VisitMaybe_slashed_or_dashed_path_expression(ctx *Maybe_slashed_or_dashed_path_expressionContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#maybe_dashed_path_expression.
	VisitMaybe_dashed_path_expression(ctx *Maybe_dashed_path_expressionContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#dashed_path_expression.
	VisitDashed_path_expression(ctx *Dashed_path_expressionContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#dashed_identifier.
	VisitDashed_identifier(ctx *Dashed_identifierContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#slashed_identifier.
	VisitSlashed_identifier(ctx *Slashed_identifierContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#identifier_or_integer.
	VisitIdentifier_or_integer(ctx *Identifier_or_integerContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#slashed_identifier_separator.
	VisitSlashed_identifier_separator(ctx *Slashed_identifier_separatorContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#slashed_path_expression.
	VisitSlashed_path_expression(ctx *Slashed_path_expressionContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#unnest_expression.
	VisitUnnest_expression(ctx *Unnest_expressionContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#unnest_expression_prefix.
	VisitUnnest_expression_prefix(ctx *Unnest_expression_prefixContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_array_zip_mode.
	VisitOpt_array_zip_mode(ctx *Opt_array_zip_modeContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#expression_with_opt_alias.
	VisitExpression_with_opt_alias(ctx *Expression_with_opt_aliasContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#tvf_prefix.
	VisitTvf_prefix(ctx *Tvf_prefixContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#tvf_argument.
	VisitTvf_argument(ctx *Tvf_argumentContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#connection_clause.
	VisitConnection_clause(ctx *Connection_clauseContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#path_expression_or_default.
	VisitPath_expression_or_default(ctx *Path_expression_or_defaultContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#descriptor_argument.
	VisitDescriptor_argument(ctx *Descriptor_argumentContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#descriptor_column_list.
	VisitDescriptor_column_list(ctx *Descriptor_column_listContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#descriptor_column.
	VisitDescriptor_column(ctx *Descriptor_columnContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#table_clause.
	VisitTable_clause(ctx *Table_clauseContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#model_clause.
	VisitModel_clause(ctx *Model_clauseContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#qualify_clause_nonreserved.
	VisitQualify_clause_nonreserved(ctx *Qualify_clause_nonreservedContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#unpivot_clause.
	VisitUnpivot_clause(ctx *Unpivot_clauseContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#unpivot_in_item_list.
	VisitUnpivot_in_item_list(ctx *Unpivot_in_item_listContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#unpivot_in_item_list_prefix.
	VisitUnpivot_in_item_list_prefix(ctx *Unpivot_in_item_list_prefixContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#unpivot_in_item.
	VisitUnpivot_in_item(ctx *Unpivot_in_itemContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_as_string_or_integer.
	VisitOpt_as_string_or_integer(ctx *Opt_as_string_or_integerContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#path_expression_list_with_opt_parens.
	VisitPath_expression_list_with_opt_parens(ctx *Path_expression_list_with_opt_parensContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#path_expression_list.
	VisitPath_expression_list(ctx *Path_expression_listContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#unpivot_nulls_filter.
	VisitUnpivot_nulls_filter(ctx *Unpivot_nulls_filterContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#pivot_clause.
	VisitPivot_clause(ctx *Pivot_clauseContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#pivot_expression_list.
	VisitPivot_expression_list(ctx *Pivot_expression_listContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#pivot_expression.
	VisitPivot_expression(ctx *Pivot_expressionContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#pivot_value_list.
	VisitPivot_value_list(ctx *Pivot_value_listContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#pivot_value.
	VisitPivot_value(ctx *Pivot_valueContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#tvf_prefix_no_args.
	VisitTvf_prefix_no_args(ctx *Tvf_prefix_no_argsContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#join_type.
	VisitJoin_type(ctx *Join_typeContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_natural.
	VisitOpt_natural(ctx *Opt_naturalContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#on_clause.
	VisitOn_clause(ctx *On_clauseContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#select_list.
	VisitSelect_list(ctx *Select_listContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#select_list_item.
	VisitSelect_list_item(ctx *Select_list_itemContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#select_column_star.
	VisitSelect_column_star(ctx *Select_column_starContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#select_column_expr.
	VisitSelect_column_expr(ctx *Select_column_exprContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#select_column_dot_star.
	VisitSelect_column_dot_star(ctx *Select_column_dot_starContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#star_modifiers.
	VisitStar_modifiers(ctx *Star_modifiersContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#star_except_list.
	VisitStar_except_list(ctx *Star_except_listContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#star_replace_list.
	VisitStar_replace_list(ctx *Star_replace_listContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#star_replace_item.
	VisitStar_replace_item(ctx *Star_replace_itemContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#expression.
	VisitExpression(ctx *ExpressionContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#expression_higher_prec_than_and.
	VisitExpression_higher_prec_than_and(ctx *Expression_higher_prec_than_andContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#expression_maybe_parenthesized_not_a_query.
	VisitExpression_maybe_parenthesized_not_a_query(ctx *Expression_maybe_parenthesized_not_a_queryContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#parenthesized_in_rhs.
	VisitParenthesized_in_rhs(ctx *Parenthesized_in_rhsContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#unary_operator.
	VisitUnary_operator(ctx *Unary_operatorContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#comparative_operator.
	VisitComparative_operator(ctx *Comparative_operatorContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#shift_operator.
	VisitShift_operator(ctx *Shift_operatorContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#additive_operator.
	VisitAdditive_operator(ctx *Additive_operatorContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#multiplicative_operator.
	VisitMultiplicative_operator(ctx *Multiplicative_operatorContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#is_operator.
	VisitIs_operator(ctx *Is_operatorContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#between_operator.
	VisitBetween_operator(ctx *Between_operatorContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#in_operator.
	VisitIn_operator(ctx *In_operatorContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#distinct_operator.
	VisitDistinct_operator(ctx *Distinct_operatorContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#parenthesized_query.
	VisitParenthesized_query(ctx *Parenthesized_queryContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#parenthesized_expression_not_a_query.
	VisitParenthesized_expression_not_a_query(ctx *Parenthesized_expression_not_a_queryContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#parenthesized_anysomeall_list_in_rhs.
	VisitParenthesized_anysomeall_list_in_rhs(ctx *Parenthesized_anysomeall_list_in_rhsContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#and_expression.
	VisitAnd_expression(ctx *And_expressionContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#in_list_two_or_more_prefix.
	VisitIn_list_two_or_more_prefix(ctx *In_list_two_or_more_prefixContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#any_some_all.
	VisitAny_some_all(ctx *Any_some_allContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#like_operator.
	VisitLike_operator(ctx *Like_operatorContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#expression_subquery_with_keyword.
	VisitExpression_subquery_with_keyword(ctx *Expression_subquery_with_keywordContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#struct_constructor.
	VisitStruct_constructor(ctx *Struct_constructorContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#struct_constructor_prefix_with_keyword.
	VisitStruct_constructor_prefix_with_keyword(ctx *Struct_constructor_prefix_with_keywordContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#struct_constructor_arg.
	VisitStruct_constructor_arg(ctx *Struct_constructor_argContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#struct_constructor_prefix_without_keyword.
	VisitStruct_constructor_prefix_without_keyword(ctx *Struct_constructor_prefix_without_keywordContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#struct_constructor_prefix_with_keyword_no_arg.
	VisitStruct_constructor_prefix_with_keyword_no_arg(ctx *Struct_constructor_prefix_with_keyword_no_argContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#interval_expression.
	VisitInterval_expression(ctx *Interval_expressionContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#function_call_expression_with_clauses.
	VisitFunction_call_expression_with_clauses(ctx *Function_call_expression_with_clausesContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#function_call_expression_with_clauses_suffix.
	VisitFunction_call_expression_with_clauses_suffix(ctx *Function_call_expression_with_clauses_suffixContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#over_clause.
	VisitOver_clause(ctx *Over_clauseContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#window_specification.
	VisitWindow_specification(ctx *Window_specificationContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_window_frame_clause.
	VisitOpt_window_frame_clause(ctx *Opt_window_frame_clauseContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#window_frame_bound.
	VisitWindow_frame_bound(ctx *Window_frame_boundContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#preceding_or_following.
	VisitPreceding_or_following(ctx *Preceding_or_followingContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#frame_unit.
	VisitFrame_unit(ctx *Frame_unitContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#partition_by_clause.
	VisitPartition_by_clause(ctx *Partition_by_clauseContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#partition_by_clause_prefix.
	VisitPartition_by_clause_prefix(ctx *Partition_by_clause_prefixContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#with_group_rows.
	VisitWith_group_rows(ctx *With_group_rowsContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#with_report_modifier.
	VisitWith_report_modifier(ctx *With_report_modifierContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#clamped_between_modifier.
	VisitClamped_between_modifier(ctx *Clamped_between_modifierContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#with_report_format.
	VisitWith_report_format(ctx *With_report_formatContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#options_list.
	VisitOptions_list(ctx *Options_listContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#options_list_prefix.
	VisitOptions_list_prefix(ctx *Options_list_prefixContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#options_entry.
	VisitOptions_entry(ctx *Options_entryContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#expression_or_proto.
	VisitExpression_or_proto(ctx *Expression_or_protoContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#options_assignment_operator.
	VisitOptions_assignment_operator(ctx *Options_assignment_operatorContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_null_handling_modifier.
	VisitOpt_null_handling_modifier(ctx *Opt_null_handling_modifierContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#function_call_argument.
	VisitFunction_call_argument(ctx *Function_call_argumentContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#sequence_arg.
	VisitSequence_arg(ctx *Sequence_argContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#named_argument.
	VisitNamed_argument(ctx *Named_argumentContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#lambda_argument.
	VisitLambda_argument(ctx *Lambda_argumentContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#lambda_argument_list.
	VisitLambda_argument_list(ctx *Lambda_argument_listContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#limit_offset_clause.
	VisitLimit_offset_clause(ctx *Limit_offset_clauseContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_having_or_group_by_modifier.
	VisitOpt_having_or_group_by_modifier(ctx *Opt_having_or_group_by_modifierContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#group_by_clause_prefix.
	VisitGroup_by_clause_prefix(ctx *Group_by_clause_prefixContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#group_by_preamble.
	VisitGroup_by_preamble(ctx *Group_by_preambleContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_and_order.
	VisitOpt_and_order(ctx *Opt_and_orderContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#hint.
	VisitHint(ctx *HintContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#hint_with_body.
	VisitHint_with_body(ctx *Hint_with_bodyContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#hint_with_body_prefix.
	VisitHint_with_body_prefix(ctx *Hint_with_body_prefixContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#hint_entry.
	VisitHint_entry(ctx *Hint_entryContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#identifier_in_hints.
	VisitIdentifier_in_hints(ctx *Identifier_in_hintsContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#extra_identifier_in_hints_name.
	VisitExtra_identifier_in_hints_name(ctx *Extra_identifier_in_hints_nameContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#grouping_item.
	VisitGrouping_item(ctx *Grouping_itemContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#grouping_set_list.
	VisitGrouping_set_list(ctx *Grouping_set_listContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#grouping_set.
	VisitGrouping_set(ctx *Grouping_setContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#cube_list.
	VisitCube_list(ctx *Cube_listContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#rollup_list.
	VisitRollup_list(ctx *Rollup_listContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_as_alias_with_required_as.
	VisitOpt_as_alias_with_required_as(ctx *Opt_as_alias_with_required_asContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_grouping_item_order.
	VisitOpt_grouping_item_order(ctx *Opt_grouping_item_orderContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_selection_item_order.
	VisitOpt_selection_item_order(ctx *Opt_selection_item_orderContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#asc_or_desc.
	VisitAsc_or_desc(ctx *Asc_or_descContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#null_order.
	VisitNull_order(ctx *Null_orderContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#function_name_from_keyword.
	VisitFunction_name_from_keyword(ctx *Function_name_from_keywordContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#replace_fields_expression.
	VisitReplace_fields_expression(ctx *Replace_fields_expressionContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#replace_fields_prefix.
	VisitReplace_fields_prefix(ctx *Replace_fields_prefixContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#replace_fields_arg.
	VisitReplace_fields_arg(ctx *Replace_fields_argContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#generalized_path_expression.
	VisitGeneralized_path_expression(ctx *Generalized_path_expressionContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#generalized_extension_path.
	VisitGeneralized_extension_path(ctx *Generalized_extension_pathContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#with_expression.
	VisitWith_expression(ctx *With_expressionContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#with_expression_variable_prefix.
	VisitWith_expression_variable_prefix(ctx *With_expression_variable_prefixContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#with_expression_variable.
	VisitWith_expression_variable(ctx *With_expression_variableContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#extract_expression.
	VisitExtract_expression(ctx *Extract_expressionContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#extract_expression_base.
	VisitExtract_expression_base(ctx *Extract_expression_baseContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_format.
	VisitOpt_format(ctx *Opt_formatContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_at_time_zone.
	VisitOpt_at_time_zone(ctx *Opt_at_time_zoneContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#cast_expression.
	VisitCast_expression(ctx *Cast_expressionContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#case_expression.
	VisitCase_expression(ctx *Case_expressionContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#case_expression_prefix.
	VisitCase_expression_prefix(ctx *Case_expression_prefixContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#case_value_expression_prefix.
	VisitCase_value_expression_prefix(ctx *Case_value_expression_prefixContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#case_no_value_expression_prefix.
	VisitCase_no_value_expression_prefix(ctx *Case_no_value_expression_prefixContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#struct_braced_constructor.
	VisitStruct_braced_constructor(ctx *Struct_braced_constructorContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#braced_new_constructor.
	VisitBraced_new_constructor(ctx *Braced_new_constructorContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#braced_constructor.
	VisitBraced_constructor(ctx *Braced_constructorContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#braced_constructor_start.
	VisitBraced_constructor_start(ctx *Braced_constructor_startContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#braced_constructor_prefix.
	VisitBraced_constructor_prefix(ctx *Braced_constructor_prefixContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#braced_constructor_field.
	VisitBraced_constructor_field(ctx *Braced_constructor_fieldContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#braced_constructor_lhs.
	VisitBraced_constructor_lhs(ctx *Braced_constructor_lhsContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#braced_constructor_field_value.
	VisitBraced_constructor_field_value(ctx *Braced_constructor_field_valueContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#braced_constructor_extension.
	VisitBraced_constructor_extension(ctx *Braced_constructor_extensionContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#new_constructor.
	VisitNew_constructor(ctx *New_constructorContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#new_constructor_prefix.
	VisitNew_constructor_prefix(ctx *New_constructor_prefixContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#new_constructor_prefix_no_arg.
	VisitNew_constructor_prefix_no_arg(ctx *New_constructor_prefix_no_argContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#new_constructor_arg.
	VisitNew_constructor_arg(ctx *New_constructor_argContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#array_constructor.
	VisitArray_constructor(ctx *Array_constructorContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#array_constructor_prefix.
	VisitArray_constructor_prefix(ctx *Array_constructor_prefixContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#array_constructor_prefix_no_expressions.
	VisitArray_constructor_prefix_no_expressions(ctx *Array_constructor_prefix_no_expressionsContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#range_literal.
	VisitRange_literal(ctx *Range_literalContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#range_type.
	VisitRange_type(ctx *Range_typeContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#type.
	VisitType(ctx *TypeContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#collate_clause.
	VisitCollate_clause(ctx *Collate_clauseContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#string_literal_or_parameter.
	VisitString_literal_or_parameter(ctx *String_literal_or_parameterContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#system_variable_expression.
	VisitSystem_variable_expression(ctx *System_variable_expressionContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#parameter_expression.
	VisitParameter_expression(ctx *Parameter_expressionContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#named_parameter_expression.
	VisitNamed_parameter_expression(ctx *Named_parameter_expressionContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#opt_type_parameters.
	VisitOpt_type_parameters(ctx *Opt_type_parametersContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#type_parameters_prefix.
	VisitType_parameters_prefix(ctx *Type_parameters_prefixContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#type_parameter.
	VisitType_parameter(ctx *Type_parameterContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#raw_type.
	VisitRaw_type(ctx *Raw_typeContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#map_type.
	VisitMap_type(ctx *Map_typeContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#function_type.
	VisitFunction_type(ctx *Function_typeContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#function_type_prefix.
	VisitFunction_type_prefix(ctx *Function_type_prefixContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#type_name.
	VisitType_name(ctx *Type_nameContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#path_expression.
	VisitPath_expression(ctx *Path_expressionContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#identifier.
	VisitIdentifier(ctx *IdentifierContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#keyword_as_identifier.
	VisitKeyword_as_identifier(ctx *Keyword_as_identifierContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#common_keyword_as_identifier.
	VisitCommon_keyword_as_identifier(ctx *Common_keyword_as_identifierContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#token_identifier.
	VisitToken_identifier(ctx *Token_identifierContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#struct_type.
	VisitStruct_type(ctx *Struct_typeContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#struct_type_prefix.
	VisitStruct_type_prefix(ctx *Struct_type_prefixContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#struct_field.
	VisitStruct_field(ctx *Struct_fieldContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#array_type.
	VisitArray_type(ctx *Array_typeContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#template_type_open.
	VisitTemplate_type_open(ctx *Template_type_openContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#template_type_close.
	VisitTemplate_type_close(ctx *Template_type_closeContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#date_or_time_literal.
	VisitDate_or_time_literal(ctx *Date_or_time_literalContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#date_or_time_literal_kind.
	VisitDate_or_time_literal_kind(ctx *Date_or_time_literal_kindContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#floating_point_literal.
	VisitFloating_point_literal(ctx *Floating_point_literalContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#json_literal.
	VisitJson_literal(ctx *Json_literalContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#bignumeric_literal.
	VisitBignumeric_literal(ctx *Bignumeric_literalContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#bignumeric_literal_prefix.
	VisitBignumeric_literal_prefix(ctx *Bignumeric_literal_prefixContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#numeric_literal.
	VisitNumeric_literal(ctx *Numeric_literalContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#numeric_literal_prefix.
	VisitNumeric_literal_prefix(ctx *Numeric_literal_prefixContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#integer_literal.
	VisitInteger_literal(ctx *Integer_literalContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#bytes_literal.
	VisitBytes_literal(ctx *Bytes_literalContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#null_literal.
	VisitNull_literal(ctx *Null_literalContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#boolean_literal.
	VisitBoolean_literal(ctx *Boolean_literalContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#string_literal.
	VisitString_literal(ctx *String_literalContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#string_literal_component.
	VisitString_literal_component(ctx *String_literal_componentContext) interface{}

	// Visit a parse tree produced by GoogleSQLParser#bytes_literal_component.
	VisitBytes_literal_component(ctx *Bytes_literal_componentContext) interface{}
}
