// Code generated from GoogleSQLParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package googlesql // GoogleSQLParser
import "github.com/antlr4-go/antlr/v4"

type BaseGoogleSQLParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseGoogleSQLParserVisitor) VisitRoot(ctx *RootContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitStmts(ctx *StmtsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitUnterminated_sql_statement(ctx *Unterminated_sql_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitSql_statement_body(ctx *Sql_statement_bodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitGql_statement(ctx *Gql_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitGraph_operation_block(ctx *Graph_operation_blockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitGraph_composite_query_block(ctx *Graph_composite_query_blockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitGraph_composite_query_prefix(ctx *Graph_composite_query_prefixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitGraph_set_operation_metadata(ctx *Graph_set_operation_metadataContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitGraph_linear_query_operation(ctx *Graph_linear_query_operationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitGraph_linear_operator_list(ctx *Graph_linear_operator_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitGraph_linear_operator(ctx *Graph_linear_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitGraph_sample_clause(ctx *Graph_sample_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_graph_sample_clause_suffix(ctx *Opt_graph_sample_clause_suffixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitGraph_for_operator(ctx *Graph_for_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_with_offset_and_alias_with_required_as(ctx *Opt_with_offset_and_alias_with_required_asContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitGraph_with_operator(ctx *Graph_with_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitGraph_page_operator(ctx *Graph_page_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitGraph_order_by_operator(ctx *Graph_order_by_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitGraph_filter_operator(ctx *Graph_filter_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitGraph_let_operator(ctx *Graph_let_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitGraph_let_variable_definition_list(ctx *Graph_let_variable_definition_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitGraph_let_variable_definition(ctx *Graph_let_variable_definitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitGraph_optional_match_operator(ctx *Graph_optional_match_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitGraph_match_operator(ctx *Graph_match_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitGraph_pattern(ctx *Graph_patternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitGraph_path_pattern_list(ctx *Graph_path_pattern_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitGraph_path_pattern(ctx *Graph_path_patternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitGraph_path_pattern_expr(ctx *Graph_path_pattern_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitGraph_path_factor(ctx *Graph_path_factorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitGraph_quantified_path_primary(ctx *Graph_quantified_path_primaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitGraph_path_primary(ctx *Graph_path_primaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitGraph_parenthesized_path_pattern(ctx *Graph_parenthesized_path_patternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitGraph_element_pattern(ctx *Graph_element_patternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitGraph_edge_pattern(ctx *Graph_edge_patternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitGraph_node_pattern(ctx *Graph_node_patternContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitGraph_element_pattern_filler(ctx *Graph_element_pattern_fillerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitGraph_property_specification(ctx *Graph_property_specificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitGraph_property_name_and_value(ctx *Graph_property_name_and_valueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_is_label_expression(ctx *Opt_is_label_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitLabel_expression(ctx *Label_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitLabel_primary(ctx *Label_primaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitParenthesized_label_expression(ctx *Parenthesized_label_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_graph_element_identifier(ctx *Opt_graph_element_identifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_graph_path_mode_prefix(ctx *Opt_graph_path_mode_prefixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitPath_or_paths(ctx *Path_or_pathsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_graph_path_mode(ctx *Opt_graph_path_modeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_graph_search_prefix(ctx *Opt_graph_search_prefixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_path_variable_assignment(ctx *Opt_path_variable_assignmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitGraph_identifier(ctx *Graph_identifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitGraph_return_operator(ctx *Graph_return_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitGraph_page_clause(ctx *Graph_page_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitGraph_order_by_clause(ctx *Graph_order_by_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitGraph_ordering_expression(ctx *Graph_ordering_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_graph_asc_or_desc(ctx *Opt_graph_asc_or_descContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitGraph_return_item_list(ctx *Graph_return_item_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitGraph_return_item(ctx *Graph_return_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitUndrop_statement(ctx *Undrop_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitModule_statement(ctx *Module_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitImport_statement(ctx *Import_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_as_or_into_alias(ctx *Opt_as_or_into_aliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitPath_expression_or_string(ctx *Path_expression_or_stringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitImport_type(ctx *Import_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitCall_statement(ctx *Call_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitDrop_statement(ctx *Drop_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_drop_mode(ctx *Opt_drop_modeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitDrop_all_row_access_policies_statement(ctx *Drop_all_row_access_policies_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitShow_statement(ctx *Show_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_like_string_literal(ctx *Opt_like_string_literalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitShow_target(ctx *Show_targetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitRename_statement(ctx *Rename_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitRevoke_statement(ctx *Revoke_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitGrant_statement(ctx *Grant_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitPrivileges(ctx *PrivilegesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitExport_metadata_statement(ctx *Export_metadata_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitExport_model_statement(ctx *Export_model_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitExport_data_statement(ctx *Export_data_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitExport_data_no_query(ctx *Export_data_no_queryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitExplain_statement(ctx *Explain_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitExecute_immediate(ctx *Execute_immediateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_execute_into_clause(ctx *Opt_execute_into_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_execute_using_clause(ctx *Opt_execute_using_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitExecute_using_argument_list(ctx *Execute_using_argument_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitExecute_using_argument(ctx *Execute_using_argumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitDescribe_statement(ctx *Describe_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitDescribe_info(ctx *Describe_infoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_from_path_expression(ctx *Opt_from_path_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitDescribe_keyword(ctx *Describe_keywordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitDefine_table_statement(ctx *Define_table_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitCreate_entity_statement(ctx *Create_entity_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_generic_entity_body(ctx *Opt_generic_entity_bodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitCreate_view_statement(ctx *Create_view_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitQuery_or_replica_source(ctx *Query_or_replica_sourceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitColumn_with_options_list(ctx *Column_with_options_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitColumn_with_options(ctx *Column_with_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitCreate_table_statement(ctx *Create_table_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_ttl_clause(ctx *Opt_ttl_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_copy_table(ctx *Opt_copy_tableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitCopy_data_source(ctx *Copy_data_sourceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_clone_table(ctx *Opt_clone_tableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_spanner_table_options(ctx *Opt_spanner_table_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_spanner_interleave_in_parent_clause(ctx *Opt_spanner_interleave_in_parent_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitSpanner_primary_key(ctx *Spanner_primary_keyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitCreate_table_function_statement(ctx *Create_table_function_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_as_query_or_string(ctx *Opt_as_query_or_stringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitUnordered_language_options(ctx *Unordered_language_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_function_parameters(ctx *Opt_function_parametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitCreate_snapshot_statement(ctx *Create_snapshot_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitCreate_external_schema_statement(ctx *Create_external_schema_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitCreate_schema_statement(ctx *Create_schema_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitCreate_property_graph_statement(ctx *Create_property_graph_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_edge_table_clause(ctx *Opt_edge_table_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitElement_table_list(ctx *Element_table_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitElement_table_definition(ctx *Element_table_definitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_label_and_properties_clause(ctx *Opt_label_and_properties_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitLabel_and_properties_list(ctx *Label_and_properties_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitLabel_and_properties(ctx *Label_and_propertiesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitProperties_clause(ctx *Properties_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitDerived_property_list(ctx *Derived_property_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitDerived_property(ctx *Derived_propertyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_except_column_list(ctx *Opt_except_column_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitProperties_all_columns(ctx *Properties_all_columnsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_dest_node_table_clause(ctx *Opt_dest_node_table_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_source_node_table_clause(ctx *Opt_source_node_table_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_key_clause(ctx *Opt_key_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitCreate_model_statement(ctx *Create_model_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_input_output_clause(ctx *Opt_input_output_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_transform_clause(ctx *Opt_transform_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_as_query_or_aliased_query_list(ctx *Opt_as_query_or_aliased_query_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitAliased_query_list(ctx *Aliased_query_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitAs_query(ctx *As_queryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitCreate_external_table_function_statement(ctx *Create_external_table_function_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitCreate_external_table_statement(ctx *Create_external_table_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_default_collate_clause(ctx *Opt_default_collate_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_like_path_expression(ctx *Opt_like_path_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitCreate_row_access_policy_statement(ctx *Create_row_access_policy_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitFilter_using_clause(ctx *Filter_using_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitCreate_row_access_policy_grant_to_clause(ctx *Create_row_access_policy_grant_to_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitCreate_privilege_restriction_statement(ctx *Create_privilege_restriction_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitRestrict_to_clause(ctx *Restrict_to_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitPossibly_empty_grantee_list(ctx *Possibly_empty_grantee_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitCreate_index_statement(ctx *Create_index_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_create_index_statement_suffix(ctx *Opt_create_index_statement_suffixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitSpanner_index_interleave_clause(ctx *Spanner_index_interleave_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitIndex_storing_list(ctx *Index_storing_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitIndex_storing_expression_list(ctx *Index_storing_expression_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitIndex_order_by_and_options(ctx *Index_order_by_and_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitIndex_all_columns(ctx *Index_all_columnsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_with_column_options(ctx *Opt_with_column_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitAll_column_column_options(ctx *All_column_column_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitColumn_ordering_and_options_expr(ctx *Column_ordering_and_options_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitIndex_unnest_expression_list(ctx *Index_unnest_expression_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitUnnest_expression_with_opt_alias_and_offset(ctx *Unnest_expression_with_opt_alias_and_offsetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOn_path_expression(ctx *On_path_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitIndex_type(ctx *Index_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_spanner_null_filtered(ctx *Opt_spanner_null_filteredContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitCreate_procedure_statement(ctx *Create_procedure_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitBegin_end_block_or_language_as_code(ctx *Begin_end_block_or_language_as_codeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitBegin_end_block(ctx *Begin_end_blockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_exception_handler(ctx *Opt_exception_handlerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitStatement_list(ctx *Statement_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitUnterminated_non_empty_statement_list(ctx *Unterminated_non_empty_statement_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitUnterminated_statement(ctx *Unterminated_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitUnterminated_script_statement(ctx *Unterminated_script_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitLabel(ctx *LabelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitUnterminated_unlabeled_script_statement(ctx *Unterminated_unlabeled_script_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitFor_in_statement(ctx *For_in_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitRepeat_statement(ctx *Repeat_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitUntil_clause(ctx *Until_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitLoop_statement(ctx *Loop_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitWhile_statement(ctx *While_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitRaise_statement(ctx *Raise_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitReturn_statement(ctx *Return_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitContinue_statement(ctx *Continue_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitVariable_declaration(ctx *Variable_declarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitBreak_statement(ctx *Break_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitCase_statement(ctx *Case_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitWhen_then_clauses(ctx *When_then_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitIf_statement(ctx *If_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitElseif_clauses(ctx *Elseif_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_else(ctx *Opt_elseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_as_code(ctx *Opt_as_codeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_external_security_clause(ctx *Opt_external_security_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitExternal_security_clause_kind(ctx *External_security_clause_kindContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitProcedure_parameters(ctx *Procedure_parametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitProcedure_parameter(ctx *Procedure_parameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitProcedure_parameter_termination(ctx *Procedure_parameter_terminationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_procedure_parameter_mode(ctx *Opt_procedure_parameter_modeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitCreate_function_statement(ctx *Create_function_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_determinism_level(ctx *Opt_determinism_levelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_sql_security_clause(ctx *Opt_sql_security_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitSql_security_clause_kind(ctx *Sql_security_clause_kindContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitAs_sql_function_body_or_string(ctx *As_sql_function_body_or_stringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitSql_function_body(ctx *Sql_function_bodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitUnordered_options_body(ctx *Unordered_options_bodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_language_or_remote_with_connection(ctx *Opt_language_or_remote_with_connectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitLanguage(ctx *LanguageContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitRemote_with_connection_clause(ctx *Remote_with_connection_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitWith_connection_clause(ctx *With_connection_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_function_returns(ctx *Opt_function_returnsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_returns(ctx *Opt_returnsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitFunction_declaration(ctx *Function_declarationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitFunction_parameters(ctx *Function_parametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitFunction_parameter(ctx *Function_parameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_not_aggregate(ctx *Opt_not_aggregateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_default_expression(ctx *Opt_default_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitType_or_tvf_schema(ctx *Type_or_tvf_schemaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitTvf_schema(ctx *Tvf_schemaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitTvf_schema_column(ctx *Tvf_schema_columnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitTemplated_parameter_type(ctx *Templated_parameter_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitTemplated_parameter_kind(ctx *Templated_parameter_kindContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_aggregate(ctx *Opt_aggregateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitCreate_database_statement(ctx *Create_database_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitCreate_connection_statement(ctx *Create_connection_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitCreate_constant_statement(ctx *Create_constant_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_or_replace(ctx *Opt_or_replaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_create_scope(ctx *Opt_create_scopeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitRun_batch_statement(ctx *Run_batch_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitAbort_batch_statement(ctx *Abort_batch_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitStart_batch_statement(ctx *Start_batch_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitRollback_statement(ctx *Rollback_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitCommit_statement(ctx *Commit_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitSet_statement(ctx *Set_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitIdentifier_list(ctx *Identifier_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitBegin_statement(ctx *Begin_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitBegin_transaction_keywords(ctx *Begin_transaction_keywordsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitTransaction_mode_list(ctx *Transaction_mode_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitTransaction_mode(ctx *Transaction_modeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitTruncate_statement(ctx *Truncate_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitMerge_statement(ctx *Merge_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitMerge_source(ctx *Merge_sourceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitMerge_when_clause(ctx *Merge_when_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitMerge_action(ctx *Merge_actionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitMerge_insert_value_list_or_source_row(ctx *Merge_insert_value_list_or_source_rowContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitBy_target(ctx *By_targetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_and_expression(ctx *Opt_and_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitStatement_level_hint(ctx *Statement_level_hintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitQuery_statement(ctx *Query_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitDml_statement(ctx *Dml_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitUpdate_statement(ctx *Update_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitDelete_statement(ctx *Delete_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitInsert_statement(ctx *Insert_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOn_conflict_clause(ctx *On_conflict_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_where_expression(ctx *Opt_where_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_conflict_target(ctx *Opt_conflict_targetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitUpdate_item_list(ctx *Update_item_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitUpdate_item(ctx *Update_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitUpdate_set_value(ctx *Update_set_valueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitNested_dml_statement(ctx *Nested_dml_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitInsert_values_list_or_table_clause(ctx *Insert_values_list_or_table_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitTable_clause_unreversed(ctx *Table_clause_unreversedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitTable_clause_no_keyword(ctx *Table_clause_no_keywordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_returning_clause(ctx *Opt_returning_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_assert_rows_modified(ctx *Opt_assert_rows_modifiedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitInsert_values_or_query(ctx *Insert_values_or_queryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitInsert_values_list(ctx *Insert_values_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitInsert_values_row(ctx *Insert_values_rowContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitExpression_or_default(ctx *Expression_or_defaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitInsert_statement_prefix(ctx *Insert_statement_prefixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitMaybe_dashed_generalized_path_expression(ctx *Maybe_dashed_generalized_path_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_into(ctx *Opt_intoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_or_ignore_replace_update(ctx *Opt_or_ignore_replace_updateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitAlter_statement(ctx *Alter_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitAnalyze_statement(ctx *Analyze_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitAssert_statement(ctx *Assert_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitAux_load_data_statement(ctx *Aux_load_data_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitClone_data_statement(ctx *Clone_data_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitClone_data_source_list(ctx *Clone_data_source_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitClone_data_source(ctx *Clone_data_sourceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_external_table_with_clauses(ctx *Opt_external_table_with_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitWith_partition_columns_clause(ctx *With_partition_columns_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitAux_load_data_from_files_options_list(ctx *Aux_load_data_from_files_options_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitCluster_by_clause_prefix_no_hint(ctx *Cluster_by_clause_prefix_no_hintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitLoad_data_partitions_clause(ctx *Load_data_partitions_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitMaybe_dashed_path_expression_with_scope(ctx *Maybe_dashed_path_expression_with_scopeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitTable_element_list(ctx *Table_element_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitTable_element(ctx *Table_elementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitTable_constraint_definition(ctx *Table_constraint_definitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitAppend_or_overwrite(ctx *Append_or_overwriteContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_description(ctx *Opt_descriptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitTable_and_column_info_list(ctx *Table_and_column_info_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitTable_and_column_info(ctx *Table_and_column_infoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitRow_access_policy_alter_action_list(ctx *Row_access_policy_alter_action_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitRow_access_policy_alter_action(ctx *Row_access_policy_alter_actionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitGrant_to_clause(ctx *Grant_to_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitGrantee_list(ctx *Grantee_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitPrivilege_list(ctx *Privilege_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitPrivilege(ctx *PrivilegeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitPath_expression_list_with_parens(ctx *Path_expression_list_with_parensContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitPrivilege_name(ctx *Privilege_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitGeneric_entity_type(ctx *Generic_entity_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitGeneric_entity_type_unchecked(ctx *Generic_entity_type_uncheckedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitSchema_object_kind(ctx *Schema_object_kindContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitAlter_action_list(ctx *Alter_action_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitAlter_action(ctx *Alter_actionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitSpanner_set_on_delete_action(ctx *Spanner_set_on_delete_actionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitSpanner_alter_column_action(ctx *Spanner_alter_column_actionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitSpanner_generated_or_default(ctx *Spanner_generated_or_defaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitGeneric_sub_entity_type(ctx *Generic_sub_entity_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitSub_entity_type_identifier(ctx *Sub_entity_type_identifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitFill_using_expression(ctx *Fill_using_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitColumn_position(ctx *Column_positionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitTable_column_definition(ctx *Table_column_definitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitColumn_attributes(ctx *Column_attributesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitColumn_attribute(ctx *Column_attributeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitPrimary_key_column_attribute(ctx *Primary_key_column_attributeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitForeign_key_column_attribute(ctx *Foreign_key_column_attributeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitHidden_column_attribute(ctx *Hidden_column_attributeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_constraint_identity(ctx *Opt_constraint_identityContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitTable_column_schema(ctx *Table_column_schemaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_column_info(ctx *Opt_column_infoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitInvalid_generated_column(ctx *Invalid_generated_columnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitInvalid_default_column(ctx *Invalid_default_columnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitDefault_column_info(ctx *Default_column_infoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitGenerated_column_info(ctx *Generated_column_infoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitIdentity_column_info(ctx *Identity_column_infoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_start_with(ctx *Opt_start_withContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_increment_by(ctx *Opt_increment_byContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_maxvalue(ctx *Opt_maxvalueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_minvalue(ctx *Opt_minvalueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_cycle(ctx *Opt_cycleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitSigned_numeric_literal(ctx *Signed_numeric_literalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitStored_mode(ctx *Stored_modeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitGenerated_mode(ctx *Generated_modeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitColumn_schema_inner(ctx *Column_schema_innerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitRaw_column_schema_inner(ctx *Raw_column_schema_innerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitRange_column_schema_inner(ctx *Range_column_schema_innerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitStruct_column_schema_inner(ctx *Struct_column_schema_innerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitStruct_column_field(ctx *Struct_column_fieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitSimple_column_schema_inner(ctx *Simple_column_schema_innerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitArray_column_schema_inner(ctx *Array_column_schema_innerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitField_schema(ctx *Field_schemaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_field_attributes(ctx *Opt_field_attributesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitNot_null_column_attribute(ctx *Not_null_column_attributeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitPrimary_key_or_table_constraint_spec(ctx *Primary_key_or_table_constraint_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_if_not_exists(ctx *Opt_if_not_existsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitPrimary_key_spec(ctx *Primary_key_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitPrimary_key_element_list(ctx *Primary_key_element_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitPrimary_key_element(ctx *Primary_key_elementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitTable_constraint_spec(ctx *Table_constraint_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitForeign_key_reference(ctx *Foreign_key_referenceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_foreign_key_action(ctx *Opt_foreign_key_actionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitForeign_key_on_update(ctx *Foreign_key_on_updateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitForeign_key_on_delete(ctx *Foreign_key_on_deleteContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitForeign_key_action(ctx *Foreign_key_actionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_foreign_key_match(ctx *Opt_foreign_key_matchContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitForeign_key_match_mode(ctx *Foreign_key_match_modeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitColumn_list(ctx *Column_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_options_list(ctx *Opt_options_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitConstraint_enforcement(ctx *Constraint_enforcementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitGeneric_entity_body(ctx *Generic_entity_bodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_if_exists(ctx *Opt_if_existsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitTable_or_table_function(ctx *Table_or_table_functionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitQuery(ctx *QueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitQuery_without_pipe_operators(ctx *Query_without_pipe_operatorsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitBad_keyword_after_from_query(ctx *Bad_keyword_after_from_queryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitBad_keyword_after_from_query_allows_parens(ctx *Bad_keyword_after_from_query_allows_parensContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitWith_clause_with_trailing_comma(ctx *With_clause_with_trailing_commaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitSelect_or_from_keyword(ctx *Select_or_from_keywordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitQuery_primary_or_set_operation(ctx *Query_primary_or_set_operationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitQuery_set_operation(ctx *Query_set_operationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitQuery_set_operation_prefix(ctx *Query_set_operation_prefixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitQuery_set_operation_item(ctx *Query_set_operation_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitQuery_primary(ctx *Query_primaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitSet_operation_metadata(ctx *Set_operation_metadataContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_column_match_suffix(ctx *Opt_column_match_suffixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_strict(ctx *Opt_strictContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitAll_or_distinct(ctx *All_or_distinctContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitQuery_set_operation_type(ctx *Query_set_operation_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_corresponding_outer_mode(ctx *Opt_corresponding_outer_modeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_outer(ctx *Opt_outerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitWith_clause(ctx *With_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitAliased_query(ctx *Aliased_queryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_aliased_query_modifiers(ctx *Opt_aliased_query_modifiersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitRecursion_depth_modifier(ctx *Recursion_depth_modifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitPossibly_unbounded_int_literal_or_parameter(ctx *Possibly_unbounded_int_literal_or_parameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitInt_literal_or_parameter(ctx *Int_literal_or_parameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOrder_by_clause(ctx *Order_by_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOrder_by_clause_prefix(ctx *Order_by_clause_prefixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOrdering_expression(ctx *Ordering_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitSelect(ctx *SelectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_clauses_following_from(ctx *Opt_clauses_following_fromContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_clauses_following_where(ctx *Opt_clauses_following_whereContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_clauses_following_group_by(ctx *Opt_clauses_following_group_byContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitWindow_clause(ctx *Window_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitWindow_clause_prefix(ctx *Window_clause_prefixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitWindow_definition(ctx *Window_definitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitWhere_clause(ctx *Where_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitHaving_clause(ctx *Having_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitGroup_by_clause(ctx *Group_by_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitGroup_by_all(ctx *Group_by_allContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitSelect_clause(ctx *Select_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_select_as_clause(ctx *Opt_select_as_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_select_with(ctx *Opt_select_withContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitFrom_clause(ctx *From_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitFrom_clause_contents(ctx *From_clause_contentsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitFrom_clause_contents_suffix(ctx *From_clause_contents_suffixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitTable_primary(ctx *Table_primaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitTvf_with_suffixes(ctx *Tvf_with_suffixesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitPivot_or_unpivot_clause_and_aliases(ctx *Pivot_or_unpivot_clause_and_aliasesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitAs_alias(ctx *As_aliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitSample_clause(ctx *Sample_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_sample_clause_suffix(ctx *Opt_sample_clause_suffixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitRepeatable_clause(ctx *Repeatable_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitPossibly_cast_int_literal_or_parameter(ctx *Possibly_cast_int_literal_or_parameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitCast_int_literal_or_parameter(ctx *Cast_int_literal_or_parameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitSample_size(ctx *Sample_sizeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitSample_size_value(ctx *Sample_size_valueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitSample_size_unit(ctx *Sample_size_unitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitPartition_by_clause_prefix_no_hint(ctx *Partition_by_clause_prefix_no_hintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitMatch_recognize_clause(ctx *Match_recognize_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitRow_pattern_expr(ctx *Row_pattern_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitRow_pattern_concatenation(ctx *Row_pattern_concatenationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitRow_pattern_factor(ctx *Row_pattern_factorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitSelect_list_prefix_with_as_aliases(ctx *Select_list_prefix_with_as_aliasesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitSelect_column_expr_with_as_alias(ctx *Select_column_expr_with_as_aliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitTable_subquery(ctx *Table_subqueryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitJoin(ctx *JoinContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitJoin_item(ctx *Join_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOn_or_using_clause_list(ctx *On_or_using_clause_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOn_or_using_clause(ctx *On_or_using_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitUsing_clause(ctx *Using_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitJoin_hint(ctx *Join_hintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitTable_path_expression(ctx *Table_path_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_at_system_time(ctx *Opt_at_system_timeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_with_offset_and_alias(ctx *Opt_with_offset_and_aliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_pivot_or_unpivot_clause_and_alias(ctx *Opt_pivot_or_unpivot_clause_and_aliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitTable_path_expression_base(ctx *Table_path_expression_baseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitMaybe_slashed_or_dashed_path_expression(ctx *Maybe_slashed_or_dashed_path_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitMaybe_dashed_path_expression(ctx *Maybe_dashed_path_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitDashed_path_expression(ctx *Dashed_path_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitDashed_identifier(ctx *Dashed_identifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitSlashed_identifier(ctx *Slashed_identifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitIdentifier_or_integer(ctx *Identifier_or_integerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitSlashed_identifier_separator(ctx *Slashed_identifier_separatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitSlashed_path_expression(ctx *Slashed_path_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitUnnest_expression(ctx *Unnest_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitUnnest_expression_prefix(ctx *Unnest_expression_prefixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_array_zip_mode(ctx *Opt_array_zip_modeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitExpression_with_opt_alias(ctx *Expression_with_opt_aliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitTvf_prefix(ctx *Tvf_prefixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitTvf_argument(ctx *Tvf_argumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitConnection_clause(ctx *Connection_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitPath_expression_or_default(ctx *Path_expression_or_defaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitDescriptor_argument(ctx *Descriptor_argumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitDescriptor_column_list(ctx *Descriptor_column_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitDescriptor_column(ctx *Descriptor_columnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitTable_clause(ctx *Table_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitModel_clause(ctx *Model_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitQualify_clause_nonreserved(ctx *Qualify_clause_nonreservedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitUnpivot_clause(ctx *Unpivot_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitUnpivot_in_item_list(ctx *Unpivot_in_item_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitUnpivot_in_item_list_prefix(ctx *Unpivot_in_item_list_prefixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitUnpivot_in_item(ctx *Unpivot_in_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_as_string_or_integer(ctx *Opt_as_string_or_integerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitPath_expression_list_with_opt_parens(ctx *Path_expression_list_with_opt_parensContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitPath_expression_list(ctx *Path_expression_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitUnpivot_nulls_filter(ctx *Unpivot_nulls_filterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitPivot_clause(ctx *Pivot_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitPivot_expression_list(ctx *Pivot_expression_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitPivot_expression(ctx *Pivot_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitPivot_value_list(ctx *Pivot_value_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitPivot_value(ctx *Pivot_valueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitTvf_prefix_no_args(ctx *Tvf_prefix_no_argsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitJoin_type(ctx *Join_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_natural(ctx *Opt_naturalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOn_clause(ctx *On_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitSelect_list(ctx *Select_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitSelect_list_item(ctx *Select_list_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitSelect_column_star(ctx *Select_column_starContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitSelect_column_expr(ctx *Select_column_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitSelect_column_dot_star(ctx *Select_column_dot_starContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitStar_modifiers(ctx *Star_modifiersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitStar_except_list(ctx *Star_except_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitStar_replace_list(ctx *Star_replace_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitStar_replace_item(ctx *Star_replace_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitExpression(ctx *ExpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitExpression_higher_prec_than_and(ctx *Expression_higher_prec_than_andContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitExpression_maybe_parenthesized_not_a_query(ctx *Expression_maybe_parenthesized_not_a_queryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitParenthesized_in_rhs(ctx *Parenthesized_in_rhsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitUnary_operator(ctx *Unary_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitComparative_operator(ctx *Comparative_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitShift_operator(ctx *Shift_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitAdditive_operator(ctx *Additive_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitMultiplicative_operator(ctx *Multiplicative_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitIs_operator(ctx *Is_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitBetween_operator(ctx *Between_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitIn_operator(ctx *In_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitDistinct_operator(ctx *Distinct_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitParenthesized_query(ctx *Parenthesized_queryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitParenthesized_expression_not_a_query(ctx *Parenthesized_expression_not_a_queryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitParenthesized_anysomeall_list_in_rhs(ctx *Parenthesized_anysomeall_list_in_rhsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitAnd_expression(ctx *And_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitIn_list_two_or_more_prefix(ctx *In_list_two_or_more_prefixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitAny_some_all(ctx *Any_some_allContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitLike_operator(ctx *Like_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitExpression_subquery_with_keyword(ctx *Expression_subquery_with_keywordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitStruct_constructor(ctx *Struct_constructorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitStruct_constructor_prefix_with_keyword(ctx *Struct_constructor_prefix_with_keywordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitStruct_constructor_arg(ctx *Struct_constructor_argContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitStruct_constructor_prefix_without_keyword(ctx *Struct_constructor_prefix_without_keywordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitStruct_constructor_prefix_with_keyword_no_arg(ctx *Struct_constructor_prefix_with_keyword_no_argContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitInterval_expression(ctx *Interval_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitFunction_call_expression_with_clauses(ctx *Function_call_expression_with_clausesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitFunction_call_expression_with_clauses_suffix(ctx *Function_call_expression_with_clauses_suffixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOver_clause(ctx *Over_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitWindow_specification(ctx *Window_specificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_window_frame_clause(ctx *Opt_window_frame_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitWindow_frame_bound(ctx *Window_frame_boundContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitPreceding_or_following(ctx *Preceding_or_followingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitFrame_unit(ctx *Frame_unitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitPartition_by_clause(ctx *Partition_by_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitPartition_by_clause_prefix(ctx *Partition_by_clause_prefixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitWith_group_rows(ctx *With_group_rowsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitWith_report_modifier(ctx *With_report_modifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitClamped_between_modifier(ctx *Clamped_between_modifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitWith_report_format(ctx *With_report_formatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOptions_list(ctx *Options_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOptions_list_prefix(ctx *Options_list_prefixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOptions_entry(ctx *Options_entryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitExpression_or_proto(ctx *Expression_or_protoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOptions_assignment_operator(ctx *Options_assignment_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_null_handling_modifier(ctx *Opt_null_handling_modifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitFunction_call_argument(ctx *Function_call_argumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitSequence_arg(ctx *Sequence_argContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitNamed_argument(ctx *Named_argumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitLambda_argument(ctx *Lambda_argumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitLambda_argument_list(ctx *Lambda_argument_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitLimit_offset_clause(ctx *Limit_offset_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_having_or_group_by_modifier(ctx *Opt_having_or_group_by_modifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitGroup_by_clause_prefix(ctx *Group_by_clause_prefixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitGroup_by_preamble(ctx *Group_by_preambleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_and_order(ctx *Opt_and_orderContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitHint(ctx *HintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitHint_with_body(ctx *Hint_with_bodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitHint_with_body_prefix(ctx *Hint_with_body_prefixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitHint_entry(ctx *Hint_entryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitIdentifier_in_hints(ctx *Identifier_in_hintsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitExtra_identifier_in_hints_name(ctx *Extra_identifier_in_hints_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitGrouping_item(ctx *Grouping_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitGrouping_set_list(ctx *Grouping_set_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitGrouping_set(ctx *Grouping_setContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitCube_list(ctx *Cube_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitRollup_list(ctx *Rollup_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_as_alias_with_required_as(ctx *Opt_as_alias_with_required_asContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_grouping_item_order(ctx *Opt_grouping_item_orderContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_selection_item_order(ctx *Opt_selection_item_orderContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitAsc_or_desc(ctx *Asc_or_descContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitNull_order(ctx *Null_orderContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitFunction_name_from_keyword(ctx *Function_name_from_keywordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitReplace_fields_expression(ctx *Replace_fields_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitReplace_fields_prefix(ctx *Replace_fields_prefixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitReplace_fields_arg(ctx *Replace_fields_argContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitGeneralized_path_expression(ctx *Generalized_path_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitGeneralized_extension_path(ctx *Generalized_extension_pathContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitWith_expression(ctx *With_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitWith_expression_variable_prefix(ctx *With_expression_variable_prefixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitWith_expression_variable(ctx *With_expression_variableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitExtract_expression(ctx *Extract_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitExtract_expression_base(ctx *Extract_expression_baseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_format(ctx *Opt_formatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_at_time_zone(ctx *Opt_at_time_zoneContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitCast_expression(ctx *Cast_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitCase_expression(ctx *Case_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitCase_expression_prefix(ctx *Case_expression_prefixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitCase_value_expression_prefix(ctx *Case_value_expression_prefixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitCase_no_value_expression_prefix(ctx *Case_no_value_expression_prefixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitStruct_braced_constructor(ctx *Struct_braced_constructorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitBraced_new_constructor(ctx *Braced_new_constructorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitBraced_constructor(ctx *Braced_constructorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitBraced_constructor_start(ctx *Braced_constructor_startContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitBraced_constructor_prefix(ctx *Braced_constructor_prefixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitBraced_constructor_field(ctx *Braced_constructor_fieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitBraced_constructor_lhs(ctx *Braced_constructor_lhsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitBraced_constructor_field_value(ctx *Braced_constructor_field_valueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitBraced_constructor_extension(ctx *Braced_constructor_extensionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitNew_constructor(ctx *New_constructorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitNew_constructor_prefix(ctx *New_constructor_prefixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitNew_constructor_prefix_no_arg(ctx *New_constructor_prefix_no_argContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitNew_constructor_arg(ctx *New_constructor_argContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitArray_constructor(ctx *Array_constructorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitArray_constructor_prefix(ctx *Array_constructor_prefixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitArray_constructor_prefix_no_expressions(ctx *Array_constructor_prefix_no_expressionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitRange_literal(ctx *Range_literalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitRange_type(ctx *Range_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitType(ctx *TypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitCollate_clause(ctx *Collate_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitString_literal_or_parameter(ctx *String_literal_or_parameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitSystem_variable_expression(ctx *System_variable_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitParameter_expression(ctx *Parameter_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitNamed_parameter_expression(ctx *Named_parameter_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitOpt_type_parameters(ctx *Opt_type_parametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitType_parameters_prefix(ctx *Type_parameters_prefixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitType_parameter(ctx *Type_parameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitRaw_type(ctx *Raw_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitMap_type(ctx *Map_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitFunction_type(ctx *Function_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitFunction_type_prefix(ctx *Function_type_prefixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitType_name(ctx *Type_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitPath_expression(ctx *Path_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitIdentifier(ctx *IdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitKeyword_as_identifier(ctx *Keyword_as_identifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitCommon_keyword_as_identifier(ctx *Common_keyword_as_identifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitToken_identifier(ctx *Token_identifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitStruct_type(ctx *Struct_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitStruct_type_prefix(ctx *Struct_type_prefixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitStruct_field(ctx *Struct_fieldContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitArray_type(ctx *Array_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitTemplate_type_open(ctx *Template_type_openContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitTemplate_type_close(ctx *Template_type_closeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitDate_or_time_literal(ctx *Date_or_time_literalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitDate_or_time_literal_kind(ctx *Date_or_time_literal_kindContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitFloating_point_literal(ctx *Floating_point_literalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitJson_literal(ctx *Json_literalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitBignumeric_literal(ctx *Bignumeric_literalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitBignumeric_literal_prefix(ctx *Bignumeric_literal_prefixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitNumeric_literal(ctx *Numeric_literalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitNumeric_literal_prefix(ctx *Numeric_literal_prefixContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitInteger_literal(ctx *Integer_literalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitBytes_literal(ctx *Bytes_literalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitNull_literal(ctx *Null_literalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitBoolean_literal(ctx *Boolean_literalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitString_literal(ctx *String_literalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitString_literal_component(ctx *String_literal_componentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseGoogleSQLParserVisitor) VisitBytes_literal_component(ctx *Bytes_literal_componentContext) interface{} {
	return v.VisitChildren(ctx)
}
