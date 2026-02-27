// Code generated from SnowflakeParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package snowflake // SnowflakeParser
import "github.com/antlr4-go/antlr/v4"

// SnowflakeParserListener is a complete listener for a parse tree produced by SnowflakeParser.
type SnowflakeParserListener interface {
	antlr.ParseTreeListener

	// EnterSnowflake_file is called when entering the snowflake_file production.
	EnterSnowflake_file(c *Snowflake_fileContext)

	// EnterBatch is called when entering the batch production.
	EnterBatch(c *BatchContext)

	// EnterSql_command is called when entering the sql_command production.
	EnterSql_command(c *Sql_commandContext)

	// EnterDdl_command is called when entering the ddl_command production.
	EnterDdl_command(c *Ddl_commandContext)

	// EnterDml_command is called when entering the dml_command production.
	EnterDml_command(c *Dml_commandContext)

	// EnterInsert_statement is called when entering the insert_statement production.
	EnterInsert_statement(c *Insert_statementContext)

	// EnterInsert_multi_table_statement is called when entering the insert_multi_table_statement production.
	EnterInsert_multi_table_statement(c *Insert_multi_table_statementContext)

	// EnterInto_clause2 is called when entering the into_clause2 production.
	EnterInto_clause2(c *Into_clause2Context)

	// EnterValues_list is called when entering the values_list production.
	EnterValues_list(c *Values_listContext)

	// EnterValue_item is called when entering the value_item production.
	EnterValue_item(c *Value_itemContext)

	// EnterMerge_statement is called when entering the merge_statement production.
	EnterMerge_statement(c *Merge_statementContext)

	// EnterMerge_matches is called when entering the merge_matches production.
	EnterMerge_matches(c *Merge_matchesContext)

	// EnterMerge_cond is called when entering the merge_cond production.
	EnterMerge_cond(c *Merge_condContext)

	// EnterMerge_update_delete is called when entering the merge_update_delete production.
	EnterMerge_update_delete(c *Merge_update_deleteContext)

	// EnterMerge_insert is called when entering the merge_insert production.
	EnterMerge_insert(c *Merge_insertContext)

	// EnterUpdate_statement is called when entering the update_statement production.
	EnterUpdate_statement(c *Update_statementContext)

	// EnterTable_or_query is called when entering the table_or_query production.
	EnterTable_or_query(c *Table_or_queryContext)

	// EnterDelete_statement is called when entering the delete_statement production.
	EnterDelete_statement(c *Delete_statementContext)

	// EnterValues_builder is called when entering the values_builder production.
	EnterValues_builder(c *Values_builderContext)

	// EnterOther_command is called when entering the other_command production.
	EnterOther_command(c *Other_commandContext)

	// EnterBegin_txn is called when entering the begin_txn production.
	EnterBegin_txn(c *Begin_txnContext)

	// EnterCopy_into_table is called when entering the copy_into_table production.
	EnterCopy_into_table(c *Copy_into_tableContext)

	// EnterExternal_location is called when entering the external_location production.
	EnterExternal_location(c *External_locationContext)

	// EnterFiles is called when entering the files production.
	EnterFiles(c *FilesContext)

	// EnterFile_format is called when entering the file_format production.
	EnterFile_format(c *File_formatContext)

	// EnterFormat_name is called when entering the format_name production.
	EnterFormat_name(c *Format_nameContext)

	// EnterFormat_type is called when entering the format_type production.
	EnterFormat_type(c *Format_typeContext)

	// EnterStage_file_format is called when entering the stage_file_format production.
	EnterStage_file_format(c *Stage_file_formatContext)

	// EnterCopy_into_location is called when entering the copy_into_location production.
	EnterCopy_into_location(c *Copy_into_locationContext)

	// EnterComment is called when entering the comment production.
	EnterComment(c *CommentContext)

	// EnterFunction_signature is called when entering the function_signature production.
	EnterFunction_signature(c *Function_signatureContext)

	// EnterCommit is called when entering the commit production.
	EnterCommit(c *CommitContext)

	// EnterExecute_immediate is called when entering the execute_immediate production.
	EnterExecute_immediate(c *Execute_immediateContext)

	// EnterExecute_task is called when entering the execute_task production.
	EnterExecute_task(c *Execute_taskContext)

	// EnterRetry_last is called when entering the retry_last production.
	EnterRetry_last(c *Retry_lastContext)

	// EnterExplain is called when entering the explain production.
	EnterExplain(c *ExplainContext)

	// EnterParallel is called when entering the parallel production.
	EnterParallel(c *ParallelContext)

	// EnterGet_dml is called when entering the get_dml production.
	EnterGet_dml(c *Get_dmlContext)

	// EnterGrant_ownership is called when entering the grant_ownership production.
	EnterGrant_ownership(c *Grant_ownershipContext)

	// EnterGrant_to_role is called when entering the grant_to_role production.
	EnterGrant_to_role(c *Grant_to_roleContext)

	// EnterGlobal_privileges is called when entering the global_privileges production.
	EnterGlobal_privileges(c *Global_privilegesContext)

	// EnterGlobal_privilege is called when entering the global_privilege production.
	EnterGlobal_privilege(c *Global_privilegeContext)

	// EnterAccount_object_privileges is called when entering the account_object_privileges production.
	EnterAccount_object_privileges(c *Account_object_privilegesContext)

	// EnterAccount_object_privilege is called when entering the account_object_privilege production.
	EnterAccount_object_privilege(c *Account_object_privilegeContext)

	// EnterSchema_privileges is called when entering the schema_privileges production.
	EnterSchema_privileges(c *Schema_privilegesContext)

	// EnterSchema_privilege is called when entering the schema_privilege production.
	EnterSchema_privilege(c *Schema_privilegeContext)

	// EnterSchema_object_privileges is called when entering the schema_object_privileges production.
	EnterSchema_object_privileges(c *Schema_object_privilegesContext)

	// EnterSchema_object_privilege is called when entering the schema_object_privilege production.
	EnterSchema_object_privilege(c *Schema_object_privilegeContext)

	// EnterGrant_to_share is called when entering the grant_to_share production.
	EnterGrant_to_share(c *Grant_to_shareContext)

	// EnterObject_privilege is called when entering the object_privilege production.
	EnterObject_privilege(c *Object_privilegeContext)

	// EnterGrant_role is called when entering the grant_role production.
	EnterGrant_role(c *Grant_roleContext)

	// EnterRole_name is called when entering the role_name production.
	EnterRole_name(c *Role_nameContext)

	// EnterSystem_defined_role is called when entering the system_defined_role production.
	EnterSystem_defined_role(c *System_defined_roleContext)

	// EnterList is called when entering the list production.
	EnterList(c *ListContext)

	// EnterUser_stage is called when entering the user_stage production.
	EnterUser_stage(c *User_stageContext)

	// EnterTable_stage is called when entering the table_stage production.
	EnterTable_stage(c *Table_stageContext)

	// EnterNamed_stage is called when entering the named_stage production.
	EnterNamed_stage(c *Named_stageContext)

	// EnterStage_path is called when entering the stage_path production.
	EnterStage_path(c *Stage_pathContext)

	// EnterPut is called when entering the put production.
	EnterPut(c *PutContext)

	// EnterRemove is called when entering the remove production.
	EnterRemove(c *RemoveContext)

	// EnterRevoke_from_role is called when entering the revoke_from_role production.
	EnterRevoke_from_role(c *Revoke_from_roleContext)

	// EnterRevoke_from_share is called when entering the revoke_from_share production.
	EnterRevoke_from_share(c *Revoke_from_shareContext)

	// EnterRevoke_role is called when entering the revoke_role production.
	EnterRevoke_role(c *Revoke_roleContext)

	// EnterRollback is called when entering the rollback production.
	EnterRollback(c *RollbackContext)

	// EnterSet is called when entering the set production.
	EnterSet(c *SetContext)

	// EnterTruncate_materialized_view is called when entering the truncate_materialized_view production.
	EnterTruncate_materialized_view(c *Truncate_materialized_viewContext)

	// EnterTruncate_table is called when entering the truncate_table production.
	EnterTruncate_table(c *Truncate_tableContext)

	// EnterUnset is called when entering the unset production.
	EnterUnset(c *UnsetContext)

	// EnterAlter_command is called when entering the alter_command production.
	EnterAlter_command(c *Alter_commandContext)

	// EnterAccount_params is called when entering the account_params production.
	EnterAccount_params(c *Account_paramsContext)

	// EnterObject_params is called when entering the object_params production.
	EnterObject_params(c *Object_paramsContext)

	// EnterDefault_ddl_collation is called when entering the default_ddl_collation production.
	EnterDefault_ddl_collation(c *Default_ddl_collationContext)

	// EnterObject_properties is called when entering the object_properties production.
	EnterObject_properties(c *Object_propertiesContext)

	// EnterSession_params is called when entering the session_params production.
	EnterSession_params(c *Session_paramsContext)

	// EnterAlter_account is called when entering the alter_account production.
	EnterAlter_account(c *Alter_accountContext)

	// EnterEnabled_true_false is called when entering the enabled_true_false production.
	EnterEnabled_true_false(c *Enabled_true_falseContext)

	// EnterAlter_alert is called when entering the alter_alert production.
	EnterAlter_alert(c *Alter_alertContext)

	// EnterResume_suspend is called when entering the resume_suspend production.
	EnterResume_suspend(c *Resume_suspendContext)

	// EnterAlert_set_clause is called when entering the alert_set_clause production.
	EnterAlert_set_clause(c *Alert_set_clauseContext)

	// EnterAlert_unset_clause is called when entering the alert_unset_clause production.
	EnterAlert_unset_clause(c *Alert_unset_clauseContext)

	// EnterAlter_api_integration is called when entering the alter_api_integration production.
	EnterAlter_api_integration(c *Alter_api_integrationContext)

	// EnterApi_integration_property is called when entering the api_integration_property production.
	EnterApi_integration_property(c *Api_integration_propertyContext)

	// EnterAlter_connection is called when entering the alter_connection production.
	EnterAlter_connection(c *Alter_connectionContext)

	// EnterAlter_database is called when entering the alter_database production.
	EnterAlter_database(c *Alter_databaseContext)

	// EnterDatabase_property is called when entering the database_property production.
	EnterDatabase_property(c *Database_propertyContext)

	// EnterAccount_id_list is called when entering the account_id_list production.
	EnterAccount_id_list(c *Account_id_listContext)

	// EnterAlter_dataset is called when entering the alter_dataset production.
	EnterAlter_dataset(c *Alter_datasetContext)

	// EnterAlter_dynamic_table is called when entering the alter_dynamic_table production.
	EnterAlter_dynamic_table(c *Alter_dynamic_tableContext)

	// EnterId_list is called when entering the id_list production.
	EnterId_list(c *Id_listContext)

	// EnterAlter_external_table is called when entering the alter_external_table production.
	EnterAlter_external_table(c *Alter_external_tableContext)

	// EnterIgnore_edition_check is called when entering the ignore_edition_check production.
	EnterIgnore_edition_check(c *Ignore_edition_checkContext)

	// EnterReplication_schedule is called when entering the replication_schedule production.
	EnterReplication_schedule(c *Replication_scheduleContext)

	// EnterDb_name_list is called when entering the db_name_list production.
	EnterDb_name_list(c *Db_name_listContext)

	// EnterShare_name_list is called when entering the share_name_list production.
	EnterShare_name_list(c *Share_name_listContext)

	// EnterFull_acct_list is called when entering the full_acct_list production.
	EnterFull_acct_list(c *Full_acct_listContext)

	// EnterAlter_failover_group is called when entering the alter_failover_group production.
	EnterAlter_failover_group(c *Alter_failover_groupContext)

	// EnterAlter_file_format is called when entering the alter_file_format production.
	EnterAlter_file_format(c *Alter_file_formatContext)

	// EnterAlter_function is called when entering the alter_function production.
	EnterAlter_function(c *Alter_functionContext)

	// EnterAlter_function_signature is called when entering the alter_function_signature production.
	EnterAlter_function_signature(c *Alter_function_signatureContext)

	// EnterData_type_list is called when entering the data_type_list production.
	EnterData_type_list(c *Data_type_listContext)

	// EnterAlter_git_repository is called when entering the alter_git_repository production.
	EnterAlter_git_repository(c *Alter_git_repositoryContext)

	// EnterAlter_git_set_opts is called when entering the alter_git_set_opts production.
	EnterAlter_git_set_opts(c *Alter_git_set_optsContext)

	// EnterAlter_git_unset_opts is called when entering the alter_git_unset_opts production.
	EnterAlter_git_unset_opts(c *Alter_git_unset_optsContext)

	// EnterAlter_masking_policy is called when entering the alter_masking_policy production.
	EnterAlter_masking_policy(c *Alter_masking_policyContext)

	// EnterAlter_materialized_view is called when entering the alter_materialized_view production.
	EnterAlter_materialized_view(c *Alter_materialized_viewContext)

	// EnterAlter_network_policy is called when entering the alter_network_policy production.
	EnterAlter_network_policy(c *Alter_network_policyContext)

	// EnterAlter_notification_integration is called when entering the alter_notification_integration production.
	EnterAlter_notification_integration(c *Alter_notification_integrationContext)

	// EnterAlter_pipe is called when entering the alter_pipe production.
	EnterAlter_pipe(c *Alter_pipeContext)

	// EnterAlter_procedure is called when entering the alter_procedure production.
	EnterAlter_procedure(c *Alter_procedureContext)

	// EnterAlter_replication_group is called when entering the alter_replication_group production.
	EnterAlter_replication_group(c *Alter_replication_groupContext)

	// EnterCredit_quota is called when entering the credit_quota production.
	EnterCredit_quota(c *Credit_quotaContext)

	// EnterFrequency is called when entering the frequency production.
	EnterFrequency(c *FrequencyContext)

	// EnterNotify_users is called when entering the notify_users production.
	EnterNotify_users(c *Notify_usersContext)

	// EnterTriggerDefinition is called when entering the triggerDefinition production.
	EnterTriggerDefinition(c *TriggerDefinitionContext)

	// EnterAlter_resource_monitor is called when entering the alter_resource_monitor production.
	EnterAlter_resource_monitor(c *Alter_resource_monitorContext)

	// EnterAlter_role is called when entering the alter_role production.
	EnterAlter_role(c *Alter_roleContext)

	// EnterAlter_row_access_policy is called when entering the alter_row_access_policy production.
	EnterAlter_row_access_policy(c *Alter_row_access_policyContext)

	// EnterAlter_schema is called when entering the alter_schema production.
	EnterAlter_schema(c *Alter_schemaContext)

	// EnterSchema_property is called when entering the schema_property production.
	EnterSchema_property(c *Schema_propertyContext)

	// EnterAlter_sequence is called when entering the alter_sequence production.
	EnterAlter_sequence(c *Alter_sequenceContext)

	// EnterAlter_secret is called when entering the alter_secret production.
	EnterAlter_secret(c *Alter_secretContext)

	// EnterSecret_opts is called when entering the secret_opts production.
	EnterSecret_opts(c *Secret_optsContext)

	// EnterSecret_set_opts is called when entering the secret_set_opts production.
	EnterSecret_set_opts(c *Secret_set_optsContext)

	// EnterSecret_oauth_client_creds_opts is called when entering the secret_oauth_client_creds_opts production.
	EnterSecret_oauth_client_creds_opts(c *Secret_oauth_client_creds_optsContext)

	// EnterSecret_oauth_auth_code_opts is called when entering the secret_oauth_auth_code_opts production.
	EnterSecret_oauth_auth_code_opts(c *Secret_oauth_auth_code_optsContext)

	// EnterSecret_api_auth_opts is called when entering the secret_api_auth_opts production.
	EnterSecret_api_auth_opts(c *Secret_api_auth_optsContext)

	// EnterSecret_basic_auth_opts is called when entering the secret_basic_auth_opts production.
	EnterSecret_basic_auth_opts(c *Secret_basic_auth_optsContext)

	// EnterSecret_generic_string_opts is called when entering the secret_generic_string_opts production.
	EnterSecret_generic_string_opts(c *Secret_generic_string_optsContext)

	// EnterAlter_security_integration_external_oauth is called when entering the alter_security_integration_external_oauth production.
	EnterAlter_security_integration_external_oauth(c *Alter_security_integration_external_oauthContext)

	// EnterSecurity_integration_external_oauth_property is called when entering the security_integration_external_oauth_property production.
	EnterSecurity_integration_external_oauth_property(c *Security_integration_external_oauth_propertyContext)

	// EnterAlter_security_integration_snowflake_oauth is called when entering the alter_security_integration_snowflake_oauth production.
	EnterAlter_security_integration_snowflake_oauth(c *Alter_security_integration_snowflake_oauthContext)

	// EnterSecurity_integration_snowflake_oauth_property is called when entering the security_integration_snowflake_oauth_property production.
	EnterSecurity_integration_snowflake_oauth_property(c *Security_integration_snowflake_oauth_propertyContext)

	// EnterAlter_security_integration_saml2 is called when entering the alter_security_integration_saml2 production.
	EnterAlter_security_integration_saml2(c *Alter_security_integration_saml2Context)

	// EnterAlter_security_integration_scim is called when entering the alter_security_integration_scim production.
	EnterAlter_security_integration_scim(c *Alter_security_integration_scimContext)

	// EnterSecurity_integration_scim_property is called when entering the security_integration_scim_property production.
	EnterSecurity_integration_scim_property(c *Security_integration_scim_propertyContext)

	// EnterAlter_session is called when entering the alter_session production.
	EnterAlter_session(c *Alter_sessionContext)

	// EnterAlter_session_policy is called when entering the alter_session_policy production.
	EnterAlter_session_policy(c *Alter_session_policyContext)

	// EnterAlter_password_policy is called when entering the alter_password_policy production.
	EnterAlter_password_policy(c *Alter_password_policyContext)

	// EnterAlter_share is called when entering the alter_share production.
	EnterAlter_share(c *Alter_shareContext)

	// EnterAlter_storage_integration is called when entering the alter_storage_integration production.
	EnterAlter_storage_integration(c *Alter_storage_integrationContext)

	// EnterAlter_stream is called when entering the alter_stream production.
	EnterAlter_stream(c *Alter_streamContext)

	// EnterAlter_table is called when entering the alter_table production.
	EnterAlter_table(c *Alter_tableContext)

	// EnterRls_operations is called when entering the rls_operations production.
	EnterRls_operations(c *Rls_operationsContext)

	// EnterClustering_action is called when entering the clustering_action production.
	EnterClustering_action(c *Clustering_actionContext)

	// EnterTable_column_action is called when entering the table_column_action production.
	EnterTable_column_action(c *Table_column_actionContext)

	// EnterAlter_column_clause is called when entering the alter_column_clause production.
	EnterAlter_column_clause(c *Alter_column_clauseContext)

	// EnterInline_constraint is called when entering the inline_constraint production.
	EnterInline_constraint(c *Inline_constraintContext)

	// EnterEnforced_not_enforced is called when entering the enforced_not_enforced production.
	EnterEnforced_not_enforced(c *Enforced_not_enforcedContext)

	// EnterDeferrable_not_deferrable is called when entering the deferrable_not_deferrable production.
	EnterDeferrable_not_deferrable(c *Deferrable_not_deferrableContext)

	// EnterInitially_deferred_or_immediate is called when entering the initially_deferred_or_immediate production.
	EnterInitially_deferred_or_immediate(c *Initially_deferred_or_immediateContext)

	// EnterCommon_constraint_properties is called when entering the common_constraint_properties production.
	EnterCommon_constraint_properties(c *Common_constraint_propertiesContext)

	// EnterOn_update is called when entering the on_update production.
	EnterOn_update(c *On_updateContext)

	// EnterOn_delete is called when entering the on_delete production.
	EnterOn_delete(c *On_deleteContext)

	// EnterForeign_key_match is called when entering the foreign_key_match production.
	EnterForeign_key_match(c *Foreign_key_matchContext)

	// EnterOn_action is called when entering the on_action production.
	EnterOn_action(c *On_actionContext)

	// EnterConstraint_properties is called when entering the constraint_properties production.
	EnterConstraint_properties(c *Constraint_propertiesContext)

	// EnterExt_table_column_action is called when entering the ext_table_column_action production.
	EnterExt_table_column_action(c *Ext_table_column_actionContext)

	// EnterConstraint_action is called when entering the constraint_action production.
	EnterConstraint_action(c *Constraint_actionContext)

	// EnterSearch_optimization_action is called when entering the search_optimization_action production.
	EnterSearch_optimization_action(c *Search_optimization_actionContext)

	// EnterSearch_method_with_target is called when entering the search_method_with_target production.
	EnterSearch_method_with_target(c *Search_method_with_targetContext)

	// EnterAlter_table_alter_column is called when entering the alter_table_alter_column production.
	EnterAlter_table_alter_column(c *Alter_table_alter_columnContext)

	// EnterAlter_column_decl_list is called when entering the alter_column_decl_list production.
	EnterAlter_column_decl_list(c *Alter_column_decl_listContext)

	// EnterAlter_column_decl is called when entering the alter_column_decl production.
	EnterAlter_column_decl(c *Alter_column_declContext)

	// EnterAlter_column_opts is called when entering the alter_column_opts production.
	EnterAlter_column_opts(c *Alter_column_optsContext)

	// EnterColumn_set_tags is called when entering the column_set_tags production.
	EnterColumn_set_tags(c *Column_set_tagsContext)

	// EnterColumn_unset_tags is called when entering the column_unset_tags production.
	EnterColumn_unset_tags(c *Column_unset_tagsContext)

	// EnterAlter_tag is called when entering the alter_tag production.
	EnterAlter_tag(c *Alter_tagContext)

	// EnterAlter_task is called when entering the alter_task production.
	EnterAlter_task(c *Alter_taskContext)

	// EnterAlter_user is called when entering the alter_user production.
	EnterAlter_user(c *Alter_userContext)

	// EnterAlter_view is called when entering the alter_view production.
	EnterAlter_view(c *Alter_viewContext)

	// EnterAlter_modify is called when entering the alter_modify production.
	EnterAlter_modify(c *Alter_modifyContext)

	// EnterAlter_warehouse is called when entering the alter_warehouse production.
	EnterAlter_warehouse(c *Alter_warehouseContext)

	// EnterAlter_connection_opts is called when entering the alter_connection_opts production.
	EnterAlter_connection_opts(c *Alter_connection_optsContext)

	// EnterAlter_user_opts is called when entering the alter_user_opts production.
	EnterAlter_user_opts(c *Alter_user_optsContext)

	// EnterAlter_tag_opts is called when entering the alter_tag_opts production.
	EnterAlter_tag_opts(c *Alter_tag_optsContext)

	// EnterAlter_network_policy_opts is called when entering the alter_network_policy_opts production.
	EnterAlter_network_policy_opts(c *Alter_network_policy_optsContext)

	// EnterAlter_warehouse_opts is called when entering the alter_warehouse_opts production.
	EnterAlter_warehouse_opts(c *Alter_warehouse_optsContext)

	// EnterAlter_account_opts is called when entering the alter_account_opts production.
	EnterAlter_account_opts(c *Alter_account_optsContext)

	// EnterSet_tags is called when entering the set_tags production.
	EnterSet_tags(c *Set_tagsContext)

	// EnterTag_decl_list is called when entering the tag_decl_list production.
	EnterTag_decl_list(c *Tag_decl_listContext)

	// EnterUnset_tags is called when entering the unset_tags production.
	EnterUnset_tags(c *Unset_tagsContext)

	// EnterTag_list is called when entering the tag_list production.
	EnterTag_list(c *Tag_listContext)

	// EnterCreate_command is called when entering the create_command production.
	EnterCreate_command(c *Create_commandContext)

	// EnterCreate_account is called when entering the create_account production.
	EnterCreate_account(c *Create_accountContext)

	// EnterCreate_alert is called when entering the create_alert production.
	EnterCreate_alert(c *Create_alertContext)

	// EnterAlert_condition is called when entering the alert_condition production.
	EnterAlert_condition(c *Alert_conditionContext)

	// EnterAlert_action is called when entering the alert_action production.
	EnterAlert_action(c *Alert_actionContext)

	// EnterCreate_api_integration is called when entering the create_api_integration production.
	EnterCreate_api_integration(c *Create_api_integrationContext)

	// EnterCreate_object_clone is called when entering the create_object_clone production.
	EnterCreate_object_clone(c *Create_object_cloneContext)

	// EnterCreate_connection is called when entering the create_connection production.
	EnterCreate_connection(c *Create_connectionContext)

	// EnterCreate_database is called when entering the create_database production.
	EnterCreate_database(c *Create_databaseContext)

	// EnterClone_at_before is called when entering the clone_at_before production.
	EnterClone_at_before(c *Clone_at_beforeContext)

	// EnterAt_before1 is called when entering the at_before1 production.
	EnterAt_before1(c *At_before1Context)

	// EnterHeader_decl is called when entering the header_decl production.
	EnterHeader_decl(c *Header_declContext)

	// EnterCompression_type is called when entering the compression_type production.
	EnterCompression_type(c *Compression_typeContext)

	// EnterCompression is called when entering the compression production.
	EnterCompression(c *CompressionContext)

	// EnterCreate_dataset is called when entering the create_dataset production.
	EnterCreate_dataset(c *Create_datasetContext)

	// EnterCreate_dynamic_table is called when entering the create_dynamic_table production.
	EnterCreate_dynamic_table(c *Create_dynamic_tableContext)

	// EnterDynamic_table_params is called when entering the dynamic_table_params production.
	EnterDynamic_table_params(c *Dynamic_table_paramsContext)

	// EnterDynamic_table_settable_params is called when entering the dynamic_table_settable_params production.
	EnterDynamic_table_settable_params(c *Dynamic_table_settable_paramsContext)

	// EnterDynamic_table_unsettable_params is called when entering the dynamic_table_unsettable_params production.
	EnterDynamic_table_unsettable_params(c *Dynamic_table_unsettable_paramsContext)

	// EnterData_retention_params is called when entering the data_retention_params production.
	EnterData_retention_params(c *Data_retention_paramsContext)

	// EnterSet_data_retention_params is called when entering the set_data_retention_params production.
	EnterSet_data_retention_params(c *Set_data_retention_paramsContext)

	// EnterCreate_event_table is called when entering the create_event_table production.
	EnterCreate_event_table(c *Create_event_tableContext)

	// EnterCreate_external_function is called when entering the create_external_function production.
	EnterCreate_external_function(c *Create_external_functionContext)

	// EnterCreate_external_table is called when entering the create_external_table production.
	EnterCreate_external_table(c *Create_external_tableContext)

	// EnterExternal_table_column_decl is called when entering the external_table_column_decl production.
	EnterExternal_table_column_decl(c *External_table_column_declContext)

	// EnterExternal_table_column_decl_list is called when entering the external_table_column_decl_list production.
	EnterExternal_table_column_decl_list(c *External_table_column_decl_listContext)

	// EnterFull_acct is called when entering the full_acct production.
	EnterFull_acct(c *Full_acctContext)

	// EnterIntegration_type_name is called when entering the integration_type_name production.
	EnterIntegration_type_name(c *Integration_type_nameContext)

	// EnterCreate_failover_group is called when entering the create_failover_group production.
	EnterCreate_failover_group(c *Create_failover_groupContext)

	// EnterType_fileformat is called when entering the type_fileformat production.
	EnterType_fileformat(c *Type_fileformatContext)

	// EnterCreate_file_format is called when entering the create_file_format production.
	EnterCreate_file_format(c *Create_file_formatContext)

	// EnterArg_decl is called when entering the arg_decl production.
	EnterArg_decl(c *Arg_declContext)

	// EnterArg_default_value_clause is called when entering the arg_default_value_clause production.
	EnterArg_default_value_clause(c *Arg_default_value_clauseContext)

	// EnterCol_decl is called when entering the col_decl production.
	EnterCol_decl(c *Col_declContext)

	// EnterVirtual_column_decl is called when entering the virtual_column_decl production.
	EnterVirtual_column_decl(c *Virtual_column_declContext)

	// EnterFunction_definition is called when entering the function_definition production.
	EnterFunction_definition(c *Function_definitionContext)

	// EnterCreate_function is called when entering the create_function production.
	EnterCreate_function(c *Create_functionContext)

	// EnterCreate_git_repository is called when entering the create_git_repository production.
	EnterCreate_git_repository(c *Create_git_repositoryContext)

	// EnterCreate_git_opts is called when entering the create_git_opts production.
	EnterCreate_git_opts(c *Create_git_optsContext)

	// EnterCreate_managed_account is called when entering the create_managed_account production.
	EnterCreate_managed_account(c *Create_managed_accountContext)

	// EnterCreate_masking_policy is called when entering the create_masking_policy production.
	EnterCreate_masking_policy(c *Create_masking_policyContext)

	// EnterTag_decl is called when entering the tag_decl production.
	EnterTag_decl(c *Tag_declContext)

	// EnterColumn_list_in_parentheses is called when entering the column_list_in_parentheses production.
	EnterColumn_list_in_parentheses(c *Column_list_in_parenthesesContext)

	// EnterCreate_materialized_view is called when entering the create_materialized_view production.
	EnterCreate_materialized_view(c *Create_materialized_viewContext)

	// EnterCreate_network_policy is called when entering the create_network_policy production.
	EnterCreate_network_policy(c *Create_network_policyContext)

	// EnterCloud_provider_params_auto is called when entering the cloud_provider_params_auto production.
	EnterCloud_provider_params_auto(c *Cloud_provider_params_autoContext)

	// EnterCloud_provider_params_push is called when entering the cloud_provider_params_push production.
	EnterCloud_provider_params_push(c *Cloud_provider_params_pushContext)

	// EnterCreate_notification_integration is called when entering the create_notification_integration production.
	EnterCreate_notification_integration(c *Create_notification_integrationContext)

	// EnterCreate_pipe is called when entering the create_pipe production.
	EnterCreate_pipe(c *Create_pipeContext)

	// EnterCaller_owner is called when entering the caller_owner production.
	EnterCaller_owner(c *Caller_ownerContext)

	// EnterExecuta_as is called when entering the executa_as production.
	EnterExecuta_as(c *Executa_asContext)

	// EnterProcedure_definition is called when entering the procedure_definition production.
	EnterProcedure_definition(c *Procedure_definitionContext)

	// EnterNot_null is called when entering the not_null production.
	EnterNot_null(c *Not_nullContext)

	// EnterCreate_procedure is called when entering the create_procedure production.
	EnterCreate_procedure(c *Create_procedureContext)

	// EnterCreate_replication_group is called when entering the create_replication_group production.
	EnterCreate_replication_group(c *Create_replication_groupContext)

	// EnterCreate_resource_monitor is called when entering the create_resource_monitor production.
	EnterCreate_resource_monitor(c *Create_resource_monitorContext)

	// EnterCreate_role is called when entering the create_role production.
	EnterCreate_role(c *Create_roleContext)

	// EnterCreate_row_access_policy is called when entering the create_row_access_policy production.
	EnterCreate_row_access_policy(c *Create_row_access_policyContext)

	// EnterCreate_schema is called when entering the create_schema production.
	EnterCreate_schema(c *Create_schemaContext)

	// EnterCreate_secret is called when entering the create_secret production.
	EnterCreate_secret(c *Create_secretContext)

	// EnterCreate_security_integration_external_oauth is called when entering the create_security_integration_external_oauth production.
	EnterCreate_security_integration_external_oauth(c *Create_security_integration_external_oauthContext)

	// EnterImplicit_none is called when entering the implicit_none production.
	EnterImplicit_none(c *Implicit_noneContext)

	// EnterCreate_security_integration_snowflake_oauth is called when entering the create_security_integration_snowflake_oauth production.
	EnterCreate_security_integration_snowflake_oauth(c *Create_security_integration_snowflake_oauthContext)

	// EnterCreate_security_integration_saml2 is called when entering the create_security_integration_saml2 production.
	EnterCreate_security_integration_saml2(c *Create_security_integration_saml2Context)

	// EnterCreate_security_integration_scim is called when entering the create_security_integration_scim production.
	EnterCreate_security_integration_scim(c *Create_security_integration_scimContext)

	// EnterNetwork_policy is called when entering the network_policy production.
	EnterNetwork_policy(c *Network_policyContext)

	// EnterPartner_application is called when entering the partner_application production.
	EnterPartner_application(c *Partner_applicationContext)

	// EnterStart_with is called when entering the start_with production.
	EnterStart_with(c *Start_withContext)

	// EnterIncrement_by is called when entering the increment_by production.
	EnterIncrement_by(c *Increment_byContext)

	// EnterCreate_semantic_view is called when entering the create_semantic_view production.
	EnterCreate_semantic_view(c *Create_semantic_viewContext)

	// EnterLogical_table is called when entering the logical_table production.
	EnterLogical_table(c *Logical_tableContext)

	// EnterRelationship_def is called when entering the relationship_def production.
	EnterRelationship_def(c *Relationship_defContext)

	// EnterWith_synonyms_clause is called when entering the with_synonyms_clause production.
	EnterWith_synonyms_clause(c *With_synonyms_clauseContext)

	// EnterSemantic_expression_list is called when entering the semantic_expression_list production.
	EnterSemantic_expression_list(c *Semantic_expression_listContext)

	// EnterSemantic_expression is called when entering the semantic_expression production.
	EnterSemantic_expression(c *Semantic_expressionContext)

	// EnterWith_extension_clause is called when entering the with_extension_clause production.
	EnterWith_extension_clause(c *With_extension_clauseContext)

	// EnterCreate_sequence is called when entering the create_sequence production.
	EnterCreate_sequence(c *Create_sequenceContext)

	// EnterCreate_session_policy is called when entering the create_session_policy production.
	EnterCreate_session_policy(c *Create_session_policyContext)

	// EnterSession_policy_params is called when entering the session_policy_params production.
	EnterSession_policy_params(c *Session_policy_paramsContext)

	// EnterSession_policy_param_name is called when entering the session_policy_param_name production.
	EnterSession_policy_param_name(c *Session_policy_param_nameContext)

	// EnterCreate_password_policy is called when entering the create_password_policy production.
	EnterCreate_password_policy(c *Create_password_policyContext)

	// EnterPassword_policy_params is called when entering the password_policy_params production.
	EnterPassword_policy_params(c *Password_policy_paramsContext)

	// EnterPassword_policy_param_name is called when entering the password_policy_param_name production.
	EnterPassword_policy_param_name(c *Password_policy_param_nameContext)

	// EnterCreate_share is called when entering the create_share production.
	EnterCreate_share(c *Create_shareContext)

	// EnterCharacter is called when entering the character production.
	EnterCharacter(c *CharacterContext)

	// EnterFormat_type_options is called when entering the format_type_options production.
	EnterFormat_type_options(c *Format_type_optionsContext)

	// EnterCopy_options is called when entering the copy_options production.
	EnterCopy_options(c *Copy_optionsContext)

	// EnterStage_encryption_opts_internal is called when entering the stage_encryption_opts_internal production.
	EnterStage_encryption_opts_internal(c *Stage_encryption_opts_internalContext)

	// EnterStage_type is called when entering the stage_type production.
	EnterStage_type(c *Stage_typeContext)

	// EnterStage_master_key is called when entering the stage_master_key production.
	EnterStage_master_key(c *Stage_master_keyContext)

	// EnterStage_kms_key is called when entering the stage_kms_key production.
	EnterStage_kms_key(c *Stage_kms_keyContext)

	// EnterStage_encryption_opts_aws is called when entering the stage_encryption_opts_aws production.
	EnterStage_encryption_opts_aws(c *Stage_encryption_opts_awsContext)

	// EnterAws_token is called when entering the aws_token production.
	EnterAws_token(c *Aws_tokenContext)

	// EnterAws_key_id is called when entering the aws_key_id production.
	EnterAws_key_id(c *Aws_key_idContext)

	// EnterAws_secret_key is called when entering the aws_secret_key production.
	EnterAws_secret_key(c *Aws_secret_keyContext)

	// EnterAws_role is called when entering the aws_role production.
	EnterAws_role(c *Aws_roleContext)

	// EnterAzure_encryption_value is called when entering the azure_encryption_value production.
	EnterAzure_encryption_value(c *Azure_encryption_valueContext)

	// EnterStage_encryption_opts_az is called when entering the stage_encryption_opts_az production.
	EnterStage_encryption_opts_az(c *Stage_encryption_opts_azContext)

	// EnterStorage_integration_eq_id is called when entering the storage_integration_eq_id production.
	EnterStorage_integration_eq_id(c *Storage_integration_eq_idContext)

	// EnterAz_credential_or_storage_integration is called when entering the az_credential_or_storage_integration production.
	EnterAz_credential_or_storage_integration(c *Az_credential_or_storage_integrationContext)

	// EnterGcp_encryption_value is called when entering the gcp_encryption_value production.
	EnterGcp_encryption_value(c *Gcp_encryption_valueContext)

	// EnterStage_encryption_opts_gcp is called when entering the stage_encryption_opts_gcp production.
	EnterStage_encryption_opts_gcp(c *Stage_encryption_opts_gcpContext)

	// EnterAws_credential_or_storage_integration is called when entering the aws_credential_or_storage_integration production.
	EnterAws_credential_or_storage_integration(c *Aws_credential_or_storage_integrationContext)

	// EnterExternal_stage_params is called when entering the external_stage_params production.
	EnterExternal_stage_params(c *External_stage_paramsContext)

	// EnterTrue_false is called when entering the true_false production.
	EnterTrue_false(c *True_falseContext)

	// EnterEnable is called when entering the enable production.
	EnterEnable(c *EnableContext)

	// EnterRefresh_on_create is called when entering the refresh_on_create production.
	EnterRefresh_on_create(c *Refresh_on_createContext)

	// EnterAuto_refresh is called when entering the auto_refresh production.
	EnterAuto_refresh(c *Auto_refreshContext)

	// EnterNotification_integration is called when entering the notification_integration production.
	EnterNotification_integration(c *Notification_integrationContext)

	// EnterDirectory_table_internal_params is called when entering the directory_table_internal_params production.
	EnterDirectory_table_internal_params(c *Directory_table_internal_paramsContext)

	// EnterDirectory_table_external_params is called when entering the directory_table_external_params production.
	EnterDirectory_table_external_params(c *Directory_table_external_paramsContext)

	// EnterCreate_stage is called when entering the create_stage production.
	EnterCreate_stage(c *Create_stageContext)

	// EnterAlter_semantic_view is called when entering the alter_semantic_view production.
	EnterAlter_semantic_view(c *Alter_semantic_viewContext)

	// EnterAlter_stage is called when entering the alter_stage production.
	EnterAlter_stage(c *Alter_stageContext)

	// EnterDrop_stage is called when entering the drop_stage production.
	EnterDrop_stage(c *Drop_stageContext)

	// EnterDescribe_stage is called when entering the describe_stage production.
	EnterDescribe_stage(c *Describe_stageContext)

	// EnterShow_stages is called when entering the show_stages production.
	EnterShow_stages(c *Show_stagesContext)

	// EnterCloud_provider_params is called when entering the cloud_provider_params production.
	EnterCloud_provider_params(c *Cloud_provider_paramsContext)

	// EnterCloud_provider_params2 is called when entering the cloud_provider_params2 production.
	EnterCloud_provider_params2(c *Cloud_provider_params2Context)

	// EnterCloud_provider_params3 is called when entering the cloud_provider_params3 production.
	EnterCloud_provider_params3(c *Cloud_provider_params3Context)

	// EnterCreate_storage_integration is called when entering the create_storage_integration production.
	EnterCreate_storage_integration(c *Create_storage_integrationContext)

	// EnterCopy_grants is called when entering the copy_grants production.
	EnterCopy_grants(c *Copy_grantsContext)

	// EnterAppend_only is called when entering the append_only production.
	EnterAppend_only(c *Append_onlyContext)

	// EnterInsert_only is called when entering the insert_only production.
	EnterInsert_only(c *Insert_onlyContext)

	// EnterShow_initial_rows is called when entering the show_initial_rows production.
	EnterShow_initial_rows(c *Show_initial_rowsContext)

	// EnterStream_time is called when entering the stream_time production.
	EnterStream_time(c *Stream_timeContext)

	// EnterCreate_stream is called when entering the create_stream production.
	EnterCreate_stream(c *Create_streamContext)

	// EnterTemporary is called when entering the temporary production.
	EnterTemporary(c *TemporaryContext)

	// EnterTable_type is called when entering the table_type production.
	EnterTable_type(c *Table_typeContext)

	// EnterWith_tags is called when entering the with_tags production.
	EnterWith_tags(c *With_tagsContext)

	// EnterWith_row_access_policy is called when entering the with_row_access_policy production.
	EnterWith_row_access_policy(c *With_row_access_policyContext)

	// EnterCluster_by is called when entering the cluster_by production.
	EnterCluster_by(c *Cluster_byContext)

	// EnterChange_tracking is called when entering the change_tracking production.
	EnterChange_tracking(c *Change_trackingContext)

	// EnterWith_masking_policy is called when entering the with_masking_policy production.
	EnterWith_masking_policy(c *With_masking_policyContext)

	// EnterCollate is called when entering the collate production.
	EnterCollate(c *CollateContext)

	// EnterOrder_noorder is called when entering the order_noorder production.
	EnterOrder_noorder(c *Order_noorderContext)

	// EnterDefault_value is called when entering the default_value production.
	EnterDefault_value(c *Default_valueContext)

	// EnterForeign_key is called when entering the foreign_key production.
	EnterForeign_key(c *Foreign_keyContext)

	// EnterPrimary_key is called when entering the primary_key production.
	EnterPrimary_key(c *Primary_keyContext)

	// EnterOut_of_line_constraint is called when entering the out_of_line_constraint production.
	EnterOut_of_line_constraint(c *Out_of_line_constraintContext)

	// EnterFull_col_decl is called when entering the full_col_decl production.
	EnterFull_col_decl(c *Full_col_declContext)

	// EnterMaterialized_col_decl is called when entering the materialized_col_decl production.
	EnterMaterialized_col_decl(c *Materialized_col_declContext)

	// EnterMaterialized_col_decl_list is called when entering the materialized_col_decl_list production.
	EnterMaterialized_col_decl_list(c *Materialized_col_decl_listContext)

	// EnterColumn_decl_item is called when entering the column_decl_item production.
	EnterColumn_decl_item(c *Column_decl_itemContext)

	// EnterColumn_decl_item_list is called when entering the column_decl_item_list production.
	EnterColumn_decl_item_list(c *Column_decl_item_listContext)

	// EnterCreate_table is called when entering the create_table production.
	EnterCreate_table(c *Create_tableContext)

	// EnterColumn_decl_item_list_paren is called when entering the column_decl_item_list_paren production.
	EnterColumn_decl_item_list_paren(c *Column_decl_item_list_parenContext)

	// EnterCreate_table_clause is called when entering the create_table_clause production.
	EnterCreate_table_clause(c *Create_table_clauseContext)

	// EnterCreate_table_as_select is called when entering the create_table_as_select production.
	EnterCreate_table_as_select(c *Create_table_as_selectContext)

	// EnterCreate_table_like is called when entering the create_table_like production.
	EnterCreate_table_like(c *Create_table_likeContext)

	// EnterCreate_tag is called when entering the create_tag production.
	EnterCreate_tag(c *Create_tagContext)

	// EnterTag_allowed_values is called when entering the tag_allowed_values production.
	EnterTag_allowed_values(c *Tag_allowed_valuesContext)

	// EnterSession_parameter is called when entering the session_parameter production.
	EnterSession_parameter(c *Session_parameterContext)

	// EnterSession_parameter_list is called when entering the session_parameter_list production.
	EnterSession_parameter_list(c *Session_parameter_listContext)

	// EnterSession_params_list is called when entering the session_params_list production.
	EnterSession_params_list(c *Session_params_listContext)

	// EnterCreate_task is called when entering the create_task production.
	EnterCreate_task(c *Create_taskContext)

	// EnterTask_parameters is called when entering the task_parameters production.
	EnterTask_parameters(c *Task_parametersContext)

	// EnterTask_compute is called when entering the task_compute production.
	EnterTask_compute(c *Task_computeContext)

	// EnterTask_schedule is called when entering the task_schedule production.
	EnterTask_schedule(c *Task_scheduleContext)

	// EnterTask_timeout is called when entering the task_timeout production.
	EnterTask_timeout(c *Task_timeoutContext)

	// EnterTask_suspend_after_failure_number is called when entering the task_suspend_after_failure_number production.
	EnterTask_suspend_after_failure_number(c *Task_suspend_after_failure_numberContext)

	// EnterTask_error_integration is called when entering the task_error_integration production.
	EnterTask_error_integration(c *Task_error_integrationContext)

	// EnterTask_overlap is called when entering the task_overlap production.
	EnterTask_overlap(c *Task_overlapContext)

	// EnterSql is called when entering the sql production.
	EnterSql(c *SqlContext)

	// EnterTask_sql is called when entering the task_sql production.
	EnterTask_sql(c *Task_sqlContext)

	// EnterTask_scripting_block is called when entering the task_scripting_block production.
	EnterTask_scripting_block(c *Task_scripting_blockContext)

	// EnterTask_scripting_declaration_list is called when entering the task_scripting_declaration_list production.
	EnterTask_scripting_declaration_list(c *Task_scripting_declaration_listContext)

	// EnterTask_scripting_declaration is called when entering the task_scripting_declaration production.
	EnterTask_scripting_declaration(c *Task_scripting_declarationContext)

	// EnterTask_scripting_statement_list is called when entering the task_scripting_statement_list production.
	EnterTask_scripting_statement_list(c *Task_scripting_statement_listContext)

	// EnterTask_scripting_statement is called when entering the task_scripting_statement production.
	EnterTask_scripting_statement(c *Task_scripting_statementContext)

	// EnterTask_scripting_assignment is called when entering the task_scripting_assignment production.
	EnterTask_scripting_assignment(c *Task_scripting_assignmentContext)

	// EnterTask_scripting_return is called when entering the task_scripting_return production.
	EnterTask_scripting_return(c *Task_scripting_returnContext)

	// EnterCall is called when entering the call production.
	EnterCall(c *CallContext)

	// EnterCreate_user is called when entering the create_user production.
	EnterCreate_user(c *Create_userContext)

	// EnterView_col is called when entering the view_col production.
	EnterView_col(c *View_colContext)

	// EnterCreate_view is called when entering the create_view production.
	EnterCreate_view(c *Create_viewContext)

	// EnterCreate_warehouse is called when entering the create_warehouse production.
	EnterCreate_warehouse(c *Create_warehouseContext)

	// EnterWh_common_size is called when entering the wh_common_size production.
	EnterWh_common_size(c *Wh_common_sizeContext)

	// EnterWh_extra_size is called when entering the wh_extra_size production.
	EnterWh_extra_size(c *Wh_extra_sizeContext)

	// EnterWh_properties is called when entering the wh_properties production.
	EnterWh_properties(c *Wh_propertiesContext)

	// EnterWh_params is called when entering the wh_params production.
	EnterWh_params(c *Wh_paramsContext)

	// EnterTrigger_definition is called when entering the trigger_definition production.
	EnterTrigger_definition(c *Trigger_definitionContext)

	// EnterObject_type_name is called when entering the object_type_name production.
	EnterObject_type_name(c *Object_type_nameContext)

	// EnterObject_type_plural is called when entering the object_type_plural production.
	EnterObject_type_plural(c *Object_type_pluralContext)

	// EnterDrop_command is called when entering the drop_command production.
	EnterDrop_command(c *Drop_commandContext)

	// EnterDrop_object is called when entering the drop_object production.
	EnterDrop_object(c *Drop_objectContext)

	// EnterDrop_alert is called when entering the drop_alert production.
	EnterDrop_alert(c *Drop_alertContext)

	// EnterDrop_connection is called when entering the drop_connection production.
	EnterDrop_connection(c *Drop_connectionContext)

	// EnterDrop_database is called when entering the drop_database production.
	EnterDrop_database(c *Drop_databaseContext)

	// EnterDrop_dynamic_table is called when entering the drop_dynamic_table production.
	EnterDrop_dynamic_table(c *Drop_dynamic_tableContext)

	// EnterDrop_external_table is called when entering the drop_external_table production.
	EnterDrop_external_table(c *Drop_external_tableContext)

	// EnterDrop_failover_group is called when entering the drop_failover_group production.
	EnterDrop_failover_group(c *Drop_failover_groupContext)

	// EnterDrop_file_format is called when entering the drop_file_format production.
	EnterDrop_file_format(c *Drop_file_formatContext)

	// EnterDrop_function is called when entering the drop_function production.
	EnterDrop_function(c *Drop_functionContext)

	// EnterDrop_git_repository is called when entering the drop_git_repository production.
	EnterDrop_git_repository(c *Drop_git_repositoryContext)

	// EnterDrop_integration is called when entering the drop_integration production.
	EnterDrop_integration(c *Drop_integrationContext)

	// EnterDrop_managed_account is called when entering the drop_managed_account production.
	EnterDrop_managed_account(c *Drop_managed_accountContext)

	// EnterDrop_masking_policy is called when entering the drop_masking_policy production.
	EnterDrop_masking_policy(c *Drop_masking_policyContext)

	// EnterDrop_materialized_view is called when entering the drop_materialized_view production.
	EnterDrop_materialized_view(c *Drop_materialized_viewContext)

	// EnterDrop_network_policy is called when entering the drop_network_policy production.
	EnterDrop_network_policy(c *Drop_network_policyContext)

	// EnterDrop_pipe is called when entering the drop_pipe production.
	EnterDrop_pipe(c *Drop_pipeContext)

	// EnterDrop_procedure is called when entering the drop_procedure production.
	EnterDrop_procedure(c *Drop_procedureContext)

	// EnterDrop_replication_group is called when entering the drop_replication_group production.
	EnterDrop_replication_group(c *Drop_replication_groupContext)

	// EnterDrop_resource_monitor is called when entering the drop_resource_monitor production.
	EnterDrop_resource_monitor(c *Drop_resource_monitorContext)

	// EnterDrop_role is called when entering the drop_role production.
	EnterDrop_role(c *Drop_roleContext)

	// EnterDrop_row_access_policy is called when entering the drop_row_access_policy production.
	EnterDrop_row_access_policy(c *Drop_row_access_policyContext)

	// EnterDrop_schema is called when entering the drop_schema production.
	EnterDrop_schema(c *Drop_schemaContext)

	// EnterDrop_secret is called when entering the drop_secret production.
	EnterDrop_secret(c *Drop_secretContext)

	// EnterDrop_semantic_view is called when entering the drop_semantic_view production.
	EnterDrop_semantic_view(c *Drop_semantic_viewContext)

	// EnterDrop_sequence is called when entering the drop_sequence production.
	EnterDrop_sequence(c *Drop_sequenceContext)

	// EnterDrop_session_policy is called when entering the drop_session_policy production.
	EnterDrop_session_policy(c *Drop_session_policyContext)

	// EnterDrop_password_policy is called when entering the drop_password_policy production.
	EnterDrop_password_policy(c *Drop_password_policyContext)

	// EnterDrop_share is called when entering the drop_share production.
	EnterDrop_share(c *Drop_shareContext)

	// EnterDrop_stream is called when entering the drop_stream production.
	EnterDrop_stream(c *Drop_streamContext)

	// EnterDrop_table is called when entering the drop_table production.
	EnterDrop_table(c *Drop_tableContext)

	// EnterDrop_tag is called when entering the drop_tag production.
	EnterDrop_tag(c *Drop_tagContext)

	// EnterDrop_task is called when entering the drop_task production.
	EnterDrop_task(c *Drop_taskContext)

	// EnterDrop_user is called when entering the drop_user production.
	EnterDrop_user(c *Drop_userContext)

	// EnterDrop_view is called when entering the drop_view production.
	EnterDrop_view(c *Drop_viewContext)

	// EnterDrop_warehouse is called when entering the drop_warehouse production.
	EnterDrop_warehouse(c *Drop_warehouseContext)

	// EnterCascade_restrict is called when entering the cascade_restrict production.
	EnterCascade_restrict(c *Cascade_restrictContext)

	// EnterArg_types is called when entering the arg_types production.
	EnterArg_types(c *Arg_typesContext)

	// EnterUndrop_command is called when entering the undrop_command production.
	EnterUndrop_command(c *Undrop_commandContext)

	// EnterUndrop_database is called when entering the undrop_database production.
	EnterUndrop_database(c *Undrop_databaseContext)

	// EnterUndrop_dynamic_table is called when entering the undrop_dynamic_table production.
	EnterUndrop_dynamic_table(c *Undrop_dynamic_tableContext)

	// EnterUndrop_schema is called when entering the undrop_schema production.
	EnterUndrop_schema(c *Undrop_schemaContext)

	// EnterUndrop_table is called when entering the undrop_table production.
	EnterUndrop_table(c *Undrop_tableContext)

	// EnterUndrop_tag is called when entering the undrop_tag production.
	EnterUndrop_tag(c *Undrop_tagContext)

	// EnterUse_command is called when entering the use_command production.
	EnterUse_command(c *Use_commandContext)

	// EnterUse_database is called when entering the use_database production.
	EnterUse_database(c *Use_databaseContext)

	// EnterUse_role is called when entering the use_role production.
	EnterUse_role(c *Use_roleContext)

	// EnterUse_schema is called when entering the use_schema production.
	EnterUse_schema(c *Use_schemaContext)

	// EnterUse_secondary_roles is called when entering the use_secondary_roles production.
	EnterUse_secondary_roles(c *Use_secondary_rolesContext)

	// EnterUse_warehouse is called when entering the use_warehouse production.
	EnterUse_warehouse(c *Use_warehouseContext)

	// EnterComment_clause is called when entering the comment_clause production.
	EnterComment_clause(c *Comment_clauseContext)

	// EnterInline_comment_clause is called when entering the inline_comment_clause production.
	EnterInline_comment_clause(c *Inline_comment_clauseContext)

	// EnterIf_suspended is called when entering the if_suspended production.
	EnterIf_suspended(c *If_suspendedContext)

	// EnterIf_exists is called when entering the if_exists production.
	EnterIf_exists(c *If_existsContext)

	// EnterIf_not_exists is called when entering the if_not_exists production.
	EnterIf_not_exists(c *If_not_existsContext)

	// EnterOr_replace is called when entering the or_replace production.
	EnterOr_replace(c *Or_replaceContext)

	// EnterOr_alter is called when entering the or_alter production.
	EnterOr_alter(c *Or_alterContext)

	// EnterDescribe is called when entering the describe production.
	EnterDescribe(c *DescribeContext)

	// EnterDescribe_command is called when entering the describe_command production.
	EnterDescribe_command(c *Describe_commandContext)

	// EnterDescribe_alert is called when entering the describe_alert production.
	EnterDescribe_alert(c *Describe_alertContext)

	// EnterDescribe_database is called when entering the describe_database production.
	EnterDescribe_database(c *Describe_databaseContext)

	// EnterDescribe_dynamic_table is called when entering the describe_dynamic_table production.
	EnterDescribe_dynamic_table(c *Describe_dynamic_tableContext)

	// EnterDescribe_event_table is called when entering the describe_event_table production.
	EnterDescribe_event_table(c *Describe_event_tableContext)

	// EnterDescribe_external_table is called when entering the describe_external_table production.
	EnterDescribe_external_table(c *Describe_external_tableContext)

	// EnterDescribe_file_format is called when entering the describe_file_format production.
	EnterDescribe_file_format(c *Describe_file_formatContext)

	// EnterDescribe_function is called when entering the describe_function production.
	EnterDescribe_function(c *Describe_functionContext)

	// EnterDescribe_git_repository is called when entering the describe_git_repository production.
	EnterDescribe_git_repository(c *Describe_git_repositoryContext)

	// EnterDescribe_integration is called when entering the describe_integration production.
	EnterDescribe_integration(c *Describe_integrationContext)

	// EnterDescribe_masking_policy is called when entering the describe_masking_policy production.
	EnterDescribe_masking_policy(c *Describe_masking_policyContext)

	// EnterDescribe_materialized_view is called when entering the describe_materialized_view production.
	EnterDescribe_materialized_view(c *Describe_materialized_viewContext)

	// EnterDescribe_network_policy is called when entering the describe_network_policy production.
	EnterDescribe_network_policy(c *Describe_network_policyContext)

	// EnterDescribe_pipe is called when entering the describe_pipe production.
	EnterDescribe_pipe(c *Describe_pipeContext)

	// EnterDescribe_procedure is called when entering the describe_procedure production.
	EnterDescribe_procedure(c *Describe_procedureContext)

	// EnterDescribe_result is called when entering the describe_result production.
	EnterDescribe_result(c *Describe_resultContext)

	// EnterDescribe_row_access_policy is called when entering the describe_row_access_policy production.
	EnterDescribe_row_access_policy(c *Describe_row_access_policyContext)

	// EnterDescribe_schema is called when entering the describe_schema production.
	EnterDescribe_schema(c *Describe_schemaContext)

	// EnterDescribe_search_optimization is called when entering the describe_search_optimization production.
	EnterDescribe_search_optimization(c *Describe_search_optimizationContext)

	// EnterDescribe_semantic_view is called when entering the describe_semantic_view production.
	EnterDescribe_semantic_view(c *Describe_semantic_viewContext)

	// EnterDescribe_sequence is called when entering the describe_sequence production.
	EnterDescribe_sequence(c *Describe_sequenceContext)

	// EnterDescribe_session_policy is called when entering the describe_session_policy production.
	EnterDescribe_session_policy(c *Describe_session_policyContext)

	// EnterDescribe_password_policy is called when entering the describe_password_policy production.
	EnterDescribe_password_policy(c *Describe_password_policyContext)

	// EnterDescribe_share is called when entering the describe_share production.
	EnterDescribe_share(c *Describe_shareContext)

	// EnterDescribe_stream is called when entering the describe_stream production.
	EnterDescribe_stream(c *Describe_streamContext)

	// EnterDescribe_table is called when entering the describe_table production.
	EnterDescribe_table(c *Describe_tableContext)

	// EnterDescribe_task is called when entering the describe_task production.
	EnterDescribe_task(c *Describe_taskContext)

	// EnterDescribe_transaction is called when entering the describe_transaction production.
	EnterDescribe_transaction(c *Describe_transactionContext)

	// EnterDescribe_user is called when entering the describe_user production.
	EnterDescribe_user(c *Describe_userContext)

	// EnterDescribe_view is called when entering the describe_view production.
	EnterDescribe_view(c *Describe_viewContext)

	// EnterDescribe_warehouse is called when entering the describe_warehouse production.
	EnterDescribe_warehouse(c *Describe_warehouseContext)

	// EnterShow_command is called when entering the show_command production.
	EnterShow_command(c *Show_commandContext)

	// EnterShow_alerts is called when entering the show_alerts production.
	EnterShow_alerts(c *Show_alertsContext)

	// EnterShow_channels is called when entering the show_channels production.
	EnterShow_channels(c *Show_channelsContext)

	// EnterShow_columns is called when entering the show_columns production.
	EnterShow_columns(c *Show_columnsContext)

	// EnterShow_connections is called when entering the show_connections production.
	EnterShow_connections(c *Show_connectionsContext)

	// EnterStarts_with is called when entering the starts_with production.
	EnterStarts_with(c *Starts_withContext)

	// EnterLimit_rows is called when entering the limit_rows production.
	EnterLimit_rows(c *Limit_rowsContext)

	// EnterLimit_rows_2 is called when entering the limit_rows_2 production.
	EnterLimit_rows_2(c *Limit_rows_2Context)

	// EnterShow_databases is called when entering the show_databases production.
	EnterShow_databases(c *Show_databasesContext)

	// EnterShow_databases_in_failover_group is called when entering the show_databases_in_failover_group production.
	EnterShow_databases_in_failover_group(c *Show_databases_in_failover_groupContext)

	// EnterShow_databases_in_replication_group is called when entering the show_databases_in_replication_group production.
	EnterShow_databases_in_replication_group(c *Show_databases_in_replication_groupContext)

	// EnterShow_datasets is called when entering the show_datasets production.
	EnterShow_datasets(c *Show_datasetsContext)

	// EnterShow_delegated_authorizations is called when entering the show_delegated_authorizations production.
	EnterShow_delegated_authorizations(c *Show_delegated_authorizationsContext)

	// EnterShow_dynamic_tables is called when entering the show_dynamic_tables production.
	EnterShow_dynamic_tables(c *Show_dynamic_tablesContext)

	// EnterShow_event_tables is called when entering the show_event_tables production.
	EnterShow_event_tables(c *Show_event_tablesContext)

	// EnterShow_external_functions is called when entering the show_external_functions production.
	EnterShow_external_functions(c *Show_external_functionsContext)

	// EnterShow_external_tables is called when entering the show_external_tables production.
	EnterShow_external_tables(c *Show_external_tablesContext)

	// EnterShow_failover_groups is called when entering the show_failover_groups production.
	EnterShow_failover_groups(c *Show_failover_groupsContext)

	// EnterShow_file_formats is called when entering the show_file_formats production.
	EnterShow_file_formats(c *Show_file_formatsContext)

	// EnterShow_functions is called when entering the show_functions production.
	EnterShow_functions(c *Show_functionsContext)

	// EnterShow_git_branches is called when entering the show_git_branches production.
	EnterShow_git_branches(c *Show_git_branchesContext)

	// EnterShow_git_repositories is called when entering the show_git_repositories production.
	EnterShow_git_repositories(c *Show_git_repositoriesContext)

	// EnterShow_git_tags is called when entering the show_git_tags production.
	EnterShow_git_tags(c *Show_git_tagsContext)

	// EnterShow_global_accounts is called when entering the show_global_accounts production.
	EnterShow_global_accounts(c *Show_global_accountsContext)

	// EnterShow_grants is called when entering the show_grants production.
	EnterShow_grants(c *Show_grantsContext)

	// EnterShow_grants_opts is called when entering the show_grants_opts production.
	EnterShow_grants_opts(c *Show_grants_optsContext)

	// EnterShow_integrations is called when entering the show_integrations production.
	EnterShow_integrations(c *Show_integrationsContext)

	// EnterShow_locks is called when entering the show_locks production.
	EnterShow_locks(c *Show_locksContext)

	// EnterShow_managed_accounts is called when entering the show_managed_accounts production.
	EnterShow_managed_accounts(c *Show_managed_accountsContext)

	// EnterShow_masking_policies is called when entering the show_masking_policies production.
	EnterShow_masking_policies(c *Show_masking_policiesContext)

	// EnterIn_obj is called when entering the in_obj production.
	EnterIn_obj(c *In_objContext)

	// EnterIn_obj_2 is called when entering the in_obj_2 production.
	EnterIn_obj_2(c *In_obj_2Context)

	// EnterIn_obj_3 is called when entering the in_obj_3 production.
	EnterIn_obj_3(c *In_obj_3Context)

	// EnterShow_materialized_views is called when entering the show_materialized_views production.
	EnterShow_materialized_views(c *Show_materialized_viewsContext)

	// EnterShow_network_policies is called when entering the show_network_policies production.
	EnterShow_network_policies(c *Show_network_policiesContext)

	// EnterShow_objects is called when entering the show_objects production.
	EnterShow_objects(c *Show_objectsContext)

	// EnterShow_organization_accounts is called when entering the show_organization_accounts production.
	EnterShow_organization_accounts(c *Show_organization_accountsContext)

	// EnterIn_for is called when entering the in_for production.
	EnterIn_for(c *In_forContext)

	// EnterShow_parameters is called when entering the show_parameters production.
	EnterShow_parameters(c *Show_parametersContext)

	// EnterShow_pipes is called when entering the show_pipes production.
	EnterShow_pipes(c *Show_pipesContext)

	// EnterShow_primary_keys is called when entering the show_primary_keys production.
	EnterShow_primary_keys(c *Show_primary_keysContext)

	// EnterShow_procedures is called when entering the show_procedures production.
	EnterShow_procedures(c *Show_proceduresContext)

	// EnterShow_regions is called when entering the show_regions production.
	EnterShow_regions(c *Show_regionsContext)

	// EnterShow_replication_accounts is called when entering the show_replication_accounts production.
	EnterShow_replication_accounts(c *Show_replication_accountsContext)

	// EnterShow_replication_databases is called when entering the show_replication_databases production.
	EnterShow_replication_databases(c *Show_replication_databasesContext)

	// EnterShow_replication_groups is called when entering the show_replication_groups production.
	EnterShow_replication_groups(c *Show_replication_groupsContext)

	// EnterShow_resource_monitors is called when entering the show_resource_monitors production.
	EnterShow_resource_monitors(c *Show_resource_monitorsContext)

	// EnterShow_roles is called when entering the show_roles production.
	EnterShow_roles(c *Show_rolesContext)

	// EnterShow_row_access_policies is called when entering the show_row_access_policies production.
	EnterShow_row_access_policies(c *Show_row_access_policiesContext)

	// EnterShow_schemas is called when entering the show_schemas production.
	EnterShow_schemas(c *Show_schemasContext)

	// EnterShow_secrets is called when entering the show_secrets production.
	EnterShow_secrets(c *Show_secretsContext)

	// EnterShow_semantic_views is called when entering the show_semantic_views production.
	EnterShow_semantic_views(c *Show_semantic_viewsContext)

	// EnterShow_semantic_dimensions is called when entering the show_semantic_dimensions production.
	EnterShow_semantic_dimensions(c *Show_semantic_dimensionsContext)

	// EnterShow_semantic_dimensions_for_metric is called when entering the show_semantic_dimensions_for_metric production.
	EnterShow_semantic_dimensions_for_metric(c *Show_semantic_dimensions_for_metricContext)

	// EnterShow_semantic_facts is called when entering the show_semantic_facts production.
	EnterShow_semantic_facts(c *Show_semantic_factsContext)

	// EnterShow_semantic_metrics is called when entering the show_semantic_metrics production.
	EnterShow_semantic_metrics(c *Show_semantic_metricsContext)

	// EnterShow_sequences is called when entering the show_sequences production.
	EnterShow_sequences(c *Show_sequencesContext)

	// EnterShow_session_policies is called when entering the show_session_policies production.
	EnterShow_session_policies(c *Show_session_policiesContext)

	// EnterShow_password_policies is called when entering the show_password_policies production.
	EnterShow_password_policies(c *Show_password_policiesContext)

	// EnterShow_shares is called when entering the show_shares production.
	EnterShow_shares(c *Show_sharesContext)

	// EnterShow_shares_in_failover_group is called when entering the show_shares_in_failover_group production.
	EnterShow_shares_in_failover_group(c *Show_shares_in_failover_groupContext)

	// EnterShow_shares_in_replication_group is called when entering the show_shares_in_replication_group production.
	EnterShow_shares_in_replication_group(c *Show_shares_in_replication_groupContext)

	// EnterShow_streams is called when entering the show_streams production.
	EnterShow_streams(c *Show_streamsContext)

	// EnterShow_tables is called when entering the show_tables production.
	EnterShow_tables(c *Show_tablesContext)

	// EnterShow_tags is called when entering the show_tags production.
	EnterShow_tags(c *Show_tagsContext)

	// EnterShow_tasks is called when entering the show_tasks production.
	EnterShow_tasks(c *Show_tasksContext)

	// EnterShow_transactions is called when entering the show_transactions production.
	EnterShow_transactions(c *Show_transactionsContext)

	// EnterShow_user_functions is called when entering the show_user_functions production.
	EnterShow_user_functions(c *Show_user_functionsContext)

	// EnterShow_users is called when entering the show_users production.
	EnterShow_users(c *Show_usersContext)

	// EnterShow_variables is called when entering the show_variables production.
	EnterShow_variables(c *Show_variablesContext)

	// EnterShow_versions_in_dataset is called when entering the show_versions_in_dataset production.
	EnterShow_versions_in_dataset(c *Show_versions_in_datasetContext)

	// EnterShow_views is called when entering the show_views production.
	EnterShow_views(c *Show_viewsContext)

	// EnterShow_warehouses is called when entering the show_warehouses production.
	EnterShow_warehouses(c *Show_warehousesContext)

	// EnterLike_pattern is called when entering the like_pattern production.
	EnterLike_pattern(c *Like_patternContext)

	// EnterAccount_identifier is called when entering the account_identifier production.
	EnterAccount_identifier(c *Account_identifierContext)

	// EnterSchema_name is called when entering the schema_name production.
	EnterSchema_name(c *Schema_nameContext)

	// EnterObject_type is called when entering the object_type production.
	EnterObject_type(c *Object_typeContext)

	// EnterObject_type_list is called when entering the object_type_list production.
	EnterObject_type_list(c *Object_type_listContext)

	// EnterTag_value is called when entering the tag_value production.
	EnterTag_value(c *Tag_valueContext)

	// EnterArg_data_type is called when entering the arg_data_type production.
	EnterArg_data_type(c *Arg_data_typeContext)

	// EnterArg_name is called when entering the arg_name production.
	EnterArg_name(c *Arg_nameContext)

	// EnterParam_name is called when entering the param_name production.
	EnterParam_name(c *Param_nameContext)

	// EnterRegion_group_id is called when entering the region_group_id production.
	EnterRegion_group_id(c *Region_group_idContext)

	// EnterSnowflake_region_id is called when entering the snowflake_region_id production.
	EnterSnowflake_region_id(c *Snowflake_region_idContext)

	// EnterString is called when entering the string production.
	EnterString(c *StringContext)

	// EnterString_list is called when entering the string_list production.
	EnterString_list(c *String_listContext)

	// EnterId_fn is called when entering the id_fn production.
	EnterId_fn(c *Id_fnContext)

	// EnterId_ is called when entering the id_ production.
	EnterId_(c *Id_Context)

	// EnterKeyword is called when entering the keyword production.
	EnterKeyword(c *KeywordContext)

	// EnterNon_reserved_words is called when entering the non_reserved_words production.
	EnterNon_reserved_words(c *Non_reserved_wordsContext)

	// EnterBuiltin_function is called when entering the builtin_function production.
	EnterBuiltin_function(c *Builtin_functionContext)

	// EnterUnary_or_binary_builtin_function is called when entering the unary_or_binary_builtin_function production.
	EnterUnary_or_binary_builtin_function(c *Unary_or_binary_builtin_functionContext)

	// EnterBinary_builtin_function is called when entering the binary_builtin_function production.
	EnterBinary_builtin_function(c *Binary_builtin_functionContext)

	// EnterBinary_or_ternary_builtin_function is called when entering the binary_or_ternary_builtin_function production.
	EnterBinary_or_ternary_builtin_function(c *Binary_or_ternary_builtin_functionContext)

	// EnterTernary_builtin_function is called when entering the ternary_builtin_function production.
	EnterTernary_builtin_function(c *Ternary_builtin_functionContext)

	// EnterList_function is called when entering the list_function production.
	EnterList_function(c *List_functionContext)

	// EnterPattern is called when entering the pattern production.
	EnterPattern(c *PatternContext)

	// EnterColumn_name is called when entering the column_name production.
	EnterColumn_name(c *Column_nameContext)

	// EnterColumn_list is called when entering the column_list production.
	EnterColumn_list(c *Column_listContext)

	// EnterAliased_column_list is called when entering the aliased_column_list production.
	EnterAliased_column_list(c *Aliased_column_listContext)

	// EnterColumn_list_with_comment is called when entering the column_list_with_comment production.
	EnterColumn_list_with_comment(c *Column_list_with_commentContext)

	// EnterObject_name is called when entering the object_name production.
	EnterObject_name(c *Object_nameContext)

	// EnterObject_name_or_identifier is called when entering the object_name_or_identifier production.
	EnterObject_name_or_identifier(c *Object_name_or_identifierContext)

	// EnterNum is called when entering the num production.
	EnterNum(c *NumContext)

	// EnterExpr_list is called when entering the expr_list production.
	EnterExpr_list(c *Expr_listContext)

	// EnterExpr_list_sorted is called when entering the expr_list_sorted production.
	EnterExpr_list_sorted(c *Expr_list_sortedContext)

	// EnterExpr is called when entering the expr production.
	EnterExpr(c *ExprContext)

	// EnterLambda_params is called when entering the lambda_params production.
	EnterLambda_params(c *Lambda_paramsContext)

	// EnterIff_expr is called when entering the iff_expr production.
	EnterIff_expr(c *Iff_exprContext)

	// EnterTrim_expression is called when entering the trim_expression production.
	EnterTrim_expression(c *Trim_expressionContext)

	// EnterTry_cast_expr is called when entering the try_cast_expr production.
	EnterTry_cast_expr(c *Try_cast_exprContext)

	// EnterCast_expr is called when entering the cast_expr production.
	EnterCast_expr(c *Cast_exprContext)

	// EnterJson_literal is called when entering the json_literal production.
	EnterJson_literal(c *Json_literalContext)

	// EnterKv_pair is called when entering the kv_pair production.
	EnterKv_pair(c *Kv_pairContext)

	// EnterValue is called when entering the value production.
	EnterValue(c *ValueContext)

	// EnterArr_literal is called when entering the arr_literal production.
	EnterArr_literal(c *Arr_literalContext)

	// EnterData_type_size is called when entering the data_type_size production.
	EnterData_type_size(c *Data_type_sizeContext)

	// EnterData_type is called when entering the data_type production.
	EnterData_type(c *Data_typeContext)

	// EnterVector_element_type is called when entering the vector_element_type production.
	EnterVector_element_type(c *Vector_element_typeContext)

	// EnterPrimitive_expression is called when entering the primitive_expression production.
	EnterPrimitive_expression(c *Primitive_expressionContext)

	// EnterOrder_by_expr is called when entering the order_by_expr production.
	EnterOrder_by_expr(c *Order_by_exprContext)

	// EnterAsc_desc is called when entering the asc_desc production.
	EnterAsc_desc(c *Asc_descContext)

	// EnterOver_clause is called when entering the over_clause production.
	EnterOver_clause(c *Over_clauseContext)

	// EnterFunction_call is called when entering the function_call production.
	EnterFunction_call(c *Function_callContext)

	// EnterParam_assoc_list is called when entering the param_assoc_list production.
	EnterParam_assoc_list(c *Param_assoc_listContext)

	// EnterParam_assoc is called when entering the param_assoc production.
	EnterParam_assoc(c *Param_assocContext)

	// EnterIgnore_or_repect_nulls is called when entering the ignore_or_repect_nulls production.
	EnterIgnore_or_repect_nulls(c *Ignore_or_repect_nullsContext)

	// EnterRanking_windowed_function is called when entering the ranking_windowed_function production.
	EnterRanking_windowed_function(c *Ranking_windowed_functionContext)

	// EnterAggregate_function is called when entering the aggregate_function production.
	EnterAggregate_function(c *Aggregate_functionContext)

	// EnterLiteral is called when entering the literal production.
	EnterLiteral(c *LiteralContext)

	// EnterSign is called when entering the sign production.
	EnterSign(c *SignContext)

	// EnterFull_column_name is called when entering the full_column_name production.
	EnterFull_column_name(c *Full_column_nameContext)

	// EnterBracket_expression is called when entering the bracket_expression production.
	EnterBracket_expression(c *Bracket_expressionContext)

	// EnterCase_expression is called when entering the case_expression production.
	EnterCase_expression(c *Case_expressionContext)

	// EnterSwitch_search_condition_section is called when entering the switch_search_condition_section production.
	EnterSwitch_search_condition_section(c *Switch_search_condition_sectionContext)

	// EnterSwitch_section is called when entering the switch_section production.
	EnterSwitch_section(c *Switch_sectionContext)

	// EnterQuery_statement is called when entering the query_statement production.
	EnterQuery_statement(c *Query_statementContext)

	// EnterWith_expression is called when entering the with_expression production.
	EnterWith_expression(c *With_expressionContext)

	// EnterCommon_table_expression is called when entering the common_table_expression production.
	EnterCommon_table_expression(c *Common_table_expressionContext)

	// EnterSelect_statement is called when entering the select_statement production.
	EnterSelect_statement(c *Select_statementContext)

	// EnterSet_operators is called when entering the set_operators production.
	EnterSet_operators(c *Set_operatorsContext)

	// EnterBy_name is called when entering the by_name production.
	EnterBy_name(c *By_nameContext)

	// EnterSelect_statement_in_parentheses is called when entering the select_statement_in_parentheses production.
	EnterSelect_statement_in_parentheses(c *Select_statement_in_parenthesesContext)

	// EnterSelect_optional_clauses is called when entering the select_optional_clauses production.
	EnterSelect_optional_clauses(c *Select_optional_clausesContext)

	// EnterSelect_clause is called when entering the select_clause production.
	EnterSelect_clause(c *Select_clauseContext)

	// EnterSelect_top_clause is called when entering the select_top_clause production.
	EnterSelect_top_clause(c *Select_top_clauseContext)

	// EnterSelect_list_no_top is called when entering the select_list_no_top production.
	EnterSelect_list_no_top(c *Select_list_no_topContext)

	// EnterSelect_list_top is called when entering the select_list_top production.
	EnterSelect_list_top(c *Select_list_topContext)

	// EnterSelect_list is called when entering the select_list production.
	EnterSelect_list(c *Select_listContext)

	// EnterSelect_list_elem is called when entering the select_list_elem production.
	EnterSelect_list_elem(c *Select_list_elemContext)

	// EnterColumn_elem_star is called when entering the column_elem_star production.
	EnterColumn_elem_star(c *Column_elem_starContext)

	// EnterColumn_elem is called when entering the column_elem production.
	EnterColumn_elem(c *Column_elemContext)

	// EnterObject_name_or_alias is called when entering the object_name_or_alias production.
	EnterObject_name_or_alias(c *Object_name_or_aliasContext)

	// EnterExclude_clause is called when entering the exclude_clause production.
	EnterExclude_clause(c *Exclude_clauseContext)

	// EnterAs_alias is called when entering the as_alias production.
	EnterAs_alias(c *As_aliasContext)

	// EnterExpression_elem is called when entering the expression_elem production.
	EnterExpression_elem(c *Expression_elemContext)

	// EnterColumn_position is called when entering the column_position production.
	EnterColumn_position(c *Column_positionContext)

	// EnterAll_distinct is called when entering the all_distinct production.
	EnterAll_distinct(c *All_distinctContext)

	// EnterTop_clause is called when entering the top_clause production.
	EnterTop_clause(c *Top_clauseContext)

	// EnterInto_clause is called when entering the into_clause production.
	EnterInto_clause(c *Into_clauseContext)

	// EnterVar_list is called when entering the var_list production.
	EnterVar_list(c *Var_listContext)

	// EnterVar is called when entering the var production.
	EnterVar(c *VarContext)

	// EnterFrom_clause is called when entering the from_clause production.
	EnterFrom_clause(c *From_clauseContext)

	// EnterTable_sources is called when entering the table_sources production.
	EnterTable_sources(c *Table_sourcesContext)

	// EnterTable_source is called when entering the table_source production.
	EnterTable_source(c *Table_sourceContext)

	// EnterTable_source_item_joined is called when entering the table_source_item_joined production.
	EnterTable_source_item_joined(c *Table_source_item_joinedContext)

	// EnterObject_ref is called when entering the object_ref production.
	EnterObject_ref(c *Object_refContext)

	// EnterFlatten_table_option is called when entering the flatten_table_option production.
	EnterFlatten_table_option(c *Flatten_table_optionContext)

	// EnterFlatten_table is called when entering the flatten_table production.
	EnterFlatten_table(c *Flatten_tableContext)

	// EnterSplited_table is called when entering the splited_table production.
	EnterSplited_table(c *Splited_tableContext)

	// EnterPrior_list is called when entering the prior_list production.
	EnterPrior_list(c *Prior_listContext)

	// EnterPrior_item is called when entering the prior_item production.
	EnterPrior_item(c *Prior_itemContext)

	// EnterOuter_join is called when entering the outer_join production.
	EnterOuter_join(c *Outer_joinContext)

	// EnterJoin_type is called when entering the join_type production.
	EnterJoin_type(c *Join_typeContext)

	// EnterJoin_clause is called when entering the join_clause production.
	EnterJoin_clause(c *Join_clauseContext)

	// EnterOn_using_clause is called when entering the on_using_clause production.
	EnterOn_using_clause(c *On_using_clauseContext)

	// EnterAt_before is called when entering the at_before production.
	EnterAt_before(c *At_beforeContext)

	// EnterEnd is called when entering the end production.
	EnterEnd(c *EndContext)

	// EnterChanges is called when entering the changes production.
	EnterChanges(c *ChangesContext)

	// EnterDefault_append_only is called when entering the default_append_only production.
	EnterDefault_append_only(c *Default_append_onlyContext)

	// EnterPartition_by is called when entering the partition_by production.
	EnterPartition_by(c *Partition_byContext)

	// EnterAlias is called when entering the alias production.
	EnterAlias(c *AliasContext)

	// EnterExpr_alias_list is called when entering the expr_alias_list production.
	EnterExpr_alias_list(c *Expr_alias_listContext)

	// EnterMeasures is called when entering the measures production.
	EnterMeasures(c *MeasuresContext)

	// EnterMatch_opts is called when entering the match_opts production.
	EnterMatch_opts(c *Match_optsContext)

	// EnterRow_match is called when entering the row_match production.
	EnterRow_match(c *Row_matchContext)

	// EnterFirst_last is called when entering the first_last production.
	EnterFirst_last(c *First_lastContext)

	// EnterSymbol is called when entering the symbol production.
	EnterSymbol(c *SymbolContext)

	// EnterAfter_match is called when entering the after_match production.
	EnterAfter_match(c *After_matchContext)

	// EnterSymbol_list is called when entering the symbol_list production.
	EnterSymbol_list(c *Symbol_listContext)

	// EnterDefine is called when entering the define production.
	EnterDefine(c *DefineContext)

	// EnterMatch_recognize is called when entering the match_recognize production.
	EnterMatch_recognize(c *Match_recognizeContext)

	// EnterPivot_unpivot is called when entering the pivot_unpivot production.
	EnterPivot_unpivot(c *Pivot_unpivotContext)

	// EnterInclude_exclude is called when entering the include_exclude production.
	EnterInclude_exclude(c *Include_excludeContext)

	// EnterPivot_in_clause is called when entering the pivot_in_clause production.
	EnterPivot_in_clause(c *Pivot_in_clauseContext)

	// EnterDefault_on_null is called when entering the default_on_null production.
	EnterDefault_on_null(c *Default_on_nullContext)

	// EnterColumn_alias_list_in_brackets is called when entering the column_alias_list_in_brackets production.
	EnterColumn_alias_list_in_brackets(c *Column_alias_list_in_bracketsContext)

	// EnterExpr_list_in_parentheses is called when entering the expr_list_in_parentheses production.
	EnterExpr_list_in_parentheses(c *Expr_list_in_parenthesesContext)

	// EnterValues_table is called when entering the values_table production.
	EnterValues_table(c *Values_tableContext)

	// EnterValues_table_body is called when entering the values_table_body production.
	EnterValues_table_body(c *Values_table_bodyContext)

	// EnterSample_method is called when entering the sample_method production.
	EnterSample_method(c *Sample_methodContext)

	// EnterRepeatable_seed is called when entering the repeatable_seed production.
	EnterRepeatable_seed(c *Repeatable_seedContext)

	// EnterSample_opts is called when entering the sample_opts production.
	EnterSample_opts(c *Sample_optsContext)

	// EnterSample is called when entering the sample production.
	EnterSample(c *SampleContext)

	// EnterSearch_condition is called when entering the search_condition production.
	EnterSearch_condition(c *Search_conditionContext)

	// EnterComparison_operator is called when entering the comparison_operator production.
	EnterComparison_operator(c *Comparison_operatorContext)

	// EnterNull_not_null is called when entering the null_not_null production.
	EnterNull_not_null(c *Null_not_nullContext)

	// EnterNot_distinct_from is called when entering the not_distinct_from production.
	EnterNot_distinct_from(c *Not_distinct_fromContext)

	// EnterSubquery is called when entering the subquery production.
	EnterSubquery(c *SubqueryContext)

	// EnterPredicate is called when entering the predicate production.
	EnterPredicate(c *PredicateContext)

	// EnterWhere_clause is called when entering the where_clause production.
	EnterWhere_clause(c *Where_clauseContext)

	// EnterGroup_by_elem is called when entering the group_by_elem production.
	EnterGroup_by_elem(c *Group_by_elemContext)

	// EnterGroup_by_list is called when entering the group_by_list production.
	EnterGroup_by_list(c *Group_by_listContext)

	// EnterGroup_by_clause is called when entering the group_by_clause production.
	EnterGroup_by_clause(c *Group_by_clauseContext)

	// EnterHaving_clause is called when entering the having_clause production.
	EnterHaving_clause(c *Having_clauseContext)

	// EnterQualify_clause is called when entering the qualify_clause production.
	EnterQualify_clause(c *Qualify_clauseContext)

	// EnterOrder_item is called when entering the order_item production.
	EnterOrder_item(c *Order_itemContext)

	// EnterOrder_by_clause is called when entering the order_by_clause production.
	EnterOrder_by_clause(c *Order_by_clauseContext)

	// EnterRow_rows is called when entering the row_rows production.
	EnterRow_rows(c *Row_rowsContext)

	// EnterFirst_next is called when entering the first_next production.
	EnterFirst_next(c *First_nextContext)

	// EnterLimit_clause is called when entering the limit_clause production.
	EnterLimit_clause(c *Limit_clauseContext)

	// EnterRound_mode is called when entering the round_mode production.
	EnterRound_mode(c *Round_modeContext)

	// EnterRound_expr is called when entering the round_expr production.
	EnterRound_expr(c *Round_exprContext)

	// EnterSupplement_non_reserved_words is called when entering the supplement_non_reserved_words production.
	EnterSupplement_non_reserved_words(c *Supplement_non_reserved_wordsContext)

	// ExitSnowflake_file is called when exiting the snowflake_file production.
	ExitSnowflake_file(c *Snowflake_fileContext)

	// ExitBatch is called when exiting the batch production.
	ExitBatch(c *BatchContext)

	// ExitSql_command is called when exiting the sql_command production.
	ExitSql_command(c *Sql_commandContext)

	// ExitDdl_command is called when exiting the ddl_command production.
	ExitDdl_command(c *Ddl_commandContext)

	// ExitDml_command is called when exiting the dml_command production.
	ExitDml_command(c *Dml_commandContext)

	// ExitInsert_statement is called when exiting the insert_statement production.
	ExitInsert_statement(c *Insert_statementContext)

	// ExitInsert_multi_table_statement is called when exiting the insert_multi_table_statement production.
	ExitInsert_multi_table_statement(c *Insert_multi_table_statementContext)

	// ExitInto_clause2 is called when exiting the into_clause2 production.
	ExitInto_clause2(c *Into_clause2Context)

	// ExitValues_list is called when exiting the values_list production.
	ExitValues_list(c *Values_listContext)

	// ExitValue_item is called when exiting the value_item production.
	ExitValue_item(c *Value_itemContext)

	// ExitMerge_statement is called when exiting the merge_statement production.
	ExitMerge_statement(c *Merge_statementContext)

	// ExitMerge_matches is called when exiting the merge_matches production.
	ExitMerge_matches(c *Merge_matchesContext)

	// ExitMerge_cond is called when exiting the merge_cond production.
	ExitMerge_cond(c *Merge_condContext)

	// ExitMerge_update_delete is called when exiting the merge_update_delete production.
	ExitMerge_update_delete(c *Merge_update_deleteContext)

	// ExitMerge_insert is called when exiting the merge_insert production.
	ExitMerge_insert(c *Merge_insertContext)

	// ExitUpdate_statement is called when exiting the update_statement production.
	ExitUpdate_statement(c *Update_statementContext)

	// ExitTable_or_query is called when exiting the table_or_query production.
	ExitTable_or_query(c *Table_or_queryContext)

	// ExitDelete_statement is called when exiting the delete_statement production.
	ExitDelete_statement(c *Delete_statementContext)

	// ExitValues_builder is called when exiting the values_builder production.
	ExitValues_builder(c *Values_builderContext)

	// ExitOther_command is called when exiting the other_command production.
	ExitOther_command(c *Other_commandContext)

	// ExitBegin_txn is called when exiting the begin_txn production.
	ExitBegin_txn(c *Begin_txnContext)

	// ExitCopy_into_table is called when exiting the copy_into_table production.
	ExitCopy_into_table(c *Copy_into_tableContext)

	// ExitExternal_location is called when exiting the external_location production.
	ExitExternal_location(c *External_locationContext)

	// ExitFiles is called when exiting the files production.
	ExitFiles(c *FilesContext)

	// ExitFile_format is called when exiting the file_format production.
	ExitFile_format(c *File_formatContext)

	// ExitFormat_name is called when exiting the format_name production.
	ExitFormat_name(c *Format_nameContext)

	// ExitFormat_type is called when exiting the format_type production.
	ExitFormat_type(c *Format_typeContext)

	// ExitStage_file_format is called when exiting the stage_file_format production.
	ExitStage_file_format(c *Stage_file_formatContext)

	// ExitCopy_into_location is called when exiting the copy_into_location production.
	ExitCopy_into_location(c *Copy_into_locationContext)

	// ExitComment is called when exiting the comment production.
	ExitComment(c *CommentContext)

	// ExitFunction_signature is called when exiting the function_signature production.
	ExitFunction_signature(c *Function_signatureContext)

	// ExitCommit is called when exiting the commit production.
	ExitCommit(c *CommitContext)

	// ExitExecute_immediate is called when exiting the execute_immediate production.
	ExitExecute_immediate(c *Execute_immediateContext)

	// ExitExecute_task is called when exiting the execute_task production.
	ExitExecute_task(c *Execute_taskContext)

	// ExitRetry_last is called when exiting the retry_last production.
	ExitRetry_last(c *Retry_lastContext)

	// ExitExplain is called when exiting the explain production.
	ExitExplain(c *ExplainContext)

	// ExitParallel is called when exiting the parallel production.
	ExitParallel(c *ParallelContext)

	// ExitGet_dml is called when exiting the get_dml production.
	ExitGet_dml(c *Get_dmlContext)

	// ExitGrant_ownership is called when exiting the grant_ownership production.
	ExitGrant_ownership(c *Grant_ownershipContext)

	// ExitGrant_to_role is called when exiting the grant_to_role production.
	ExitGrant_to_role(c *Grant_to_roleContext)

	// ExitGlobal_privileges is called when exiting the global_privileges production.
	ExitGlobal_privileges(c *Global_privilegesContext)

	// ExitGlobal_privilege is called when exiting the global_privilege production.
	ExitGlobal_privilege(c *Global_privilegeContext)

	// ExitAccount_object_privileges is called when exiting the account_object_privileges production.
	ExitAccount_object_privileges(c *Account_object_privilegesContext)

	// ExitAccount_object_privilege is called when exiting the account_object_privilege production.
	ExitAccount_object_privilege(c *Account_object_privilegeContext)

	// ExitSchema_privileges is called when exiting the schema_privileges production.
	ExitSchema_privileges(c *Schema_privilegesContext)

	// ExitSchema_privilege is called when exiting the schema_privilege production.
	ExitSchema_privilege(c *Schema_privilegeContext)

	// ExitSchema_object_privileges is called when exiting the schema_object_privileges production.
	ExitSchema_object_privileges(c *Schema_object_privilegesContext)

	// ExitSchema_object_privilege is called when exiting the schema_object_privilege production.
	ExitSchema_object_privilege(c *Schema_object_privilegeContext)

	// ExitGrant_to_share is called when exiting the grant_to_share production.
	ExitGrant_to_share(c *Grant_to_shareContext)

	// ExitObject_privilege is called when exiting the object_privilege production.
	ExitObject_privilege(c *Object_privilegeContext)

	// ExitGrant_role is called when exiting the grant_role production.
	ExitGrant_role(c *Grant_roleContext)

	// ExitRole_name is called when exiting the role_name production.
	ExitRole_name(c *Role_nameContext)

	// ExitSystem_defined_role is called when exiting the system_defined_role production.
	ExitSystem_defined_role(c *System_defined_roleContext)

	// ExitList is called when exiting the list production.
	ExitList(c *ListContext)

	// ExitUser_stage is called when exiting the user_stage production.
	ExitUser_stage(c *User_stageContext)

	// ExitTable_stage is called when exiting the table_stage production.
	ExitTable_stage(c *Table_stageContext)

	// ExitNamed_stage is called when exiting the named_stage production.
	ExitNamed_stage(c *Named_stageContext)

	// ExitStage_path is called when exiting the stage_path production.
	ExitStage_path(c *Stage_pathContext)

	// ExitPut is called when exiting the put production.
	ExitPut(c *PutContext)

	// ExitRemove is called when exiting the remove production.
	ExitRemove(c *RemoveContext)

	// ExitRevoke_from_role is called when exiting the revoke_from_role production.
	ExitRevoke_from_role(c *Revoke_from_roleContext)

	// ExitRevoke_from_share is called when exiting the revoke_from_share production.
	ExitRevoke_from_share(c *Revoke_from_shareContext)

	// ExitRevoke_role is called when exiting the revoke_role production.
	ExitRevoke_role(c *Revoke_roleContext)

	// ExitRollback is called when exiting the rollback production.
	ExitRollback(c *RollbackContext)

	// ExitSet is called when exiting the set production.
	ExitSet(c *SetContext)

	// ExitTruncate_materialized_view is called when exiting the truncate_materialized_view production.
	ExitTruncate_materialized_view(c *Truncate_materialized_viewContext)

	// ExitTruncate_table is called when exiting the truncate_table production.
	ExitTruncate_table(c *Truncate_tableContext)

	// ExitUnset is called when exiting the unset production.
	ExitUnset(c *UnsetContext)

	// ExitAlter_command is called when exiting the alter_command production.
	ExitAlter_command(c *Alter_commandContext)

	// ExitAccount_params is called when exiting the account_params production.
	ExitAccount_params(c *Account_paramsContext)

	// ExitObject_params is called when exiting the object_params production.
	ExitObject_params(c *Object_paramsContext)

	// ExitDefault_ddl_collation is called when exiting the default_ddl_collation production.
	ExitDefault_ddl_collation(c *Default_ddl_collationContext)

	// ExitObject_properties is called when exiting the object_properties production.
	ExitObject_properties(c *Object_propertiesContext)

	// ExitSession_params is called when exiting the session_params production.
	ExitSession_params(c *Session_paramsContext)

	// ExitAlter_account is called when exiting the alter_account production.
	ExitAlter_account(c *Alter_accountContext)

	// ExitEnabled_true_false is called when exiting the enabled_true_false production.
	ExitEnabled_true_false(c *Enabled_true_falseContext)

	// ExitAlter_alert is called when exiting the alter_alert production.
	ExitAlter_alert(c *Alter_alertContext)

	// ExitResume_suspend is called when exiting the resume_suspend production.
	ExitResume_suspend(c *Resume_suspendContext)

	// ExitAlert_set_clause is called when exiting the alert_set_clause production.
	ExitAlert_set_clause(c *Alert_set_clauseContext)

	// ExitAlert_unset_clause is called when exiting the alert_unset_clause production.
	ExitAlert_unset_clause(c *Alert_unset_clauseContext)

	// ExitAlter_api_integration is called when exiting the alter_api_integration production.
	ExitAlter_api_integration(c *Alter_api_integrationContext)

	// ExitApi_integration_property is called when exiting the api_integration_property production.
	ExitApi_integration_property(c *Api_integration_propertyContext)

	// ExitAlter_connection is called when exiting the alter_connection production.
	ExitAlter_connection(c *Alter_connectionContext)

	// ExitAlter_database is called when exiting the alter_database production.
	ExitAlter_database(c *Alter_databaseContext)

	// ExitDatabase_property is called when exiting the database_property production.
	ExitDatabase_property(c *Database_propertyContext)

	// ExitAccount_id_list is called when exiting the account_id_list production.
	ExitAccount_id_list(c *Account_id_listContext)

	// ExitAlter_dataset is called when exiting the alter_dataset production.
	ExitAlter_dataset(c *Alter_datasetContext)

	// ExitAlter_dynamic_table is called when exiting the alter_dynamic_table production.
	ExitAlter_dynamic_table(c *Alter_dynamic_tableContext)

	// ExitId_list is called when exiting the id_list production.
	ExitId_list(c *Id_listContext)

	// ExitAlter_external_table is called when exiting the alter_external_table production.
	ExitAlter_external_table(c *Alter_external_tableContext)

	// ExitIgnore_edition_check is called when exiting the ignore_edition_check production.
	ExitIgnore_edition_check(c *Ignore_edition_checkContext)

	// ExitReplication_schedule is called when exiting the replication_schedule production.
	ExitReplication_schedule(c *Replication_scheduleContext)

	// ExitDb_name_list is called when exiting the db_name_list production.
	ExitDb_name_list(c *Db_name_listContext)

	// ExitShare_name_list is called when exiting the share_name_list production.
	ExitShare_name_list(c *Share_name_listContext)

	// ExitFull_acct_list is called when exiting the full_acct_list production.
	ExitFull_acct_list(c *Full_acct_listContext)

	// ExitAlter_failover_group is called when exiting the alter_failover_group production.
	ExitAlter_failover_group(c *Alter_failover_groupContext)

	// ExitAlter_file_format is called when exiting the alter_file_format production.
	ExitAlter_file_format(c *Alter_file_formatContext)

	// ExitAlter_function is called when exiting the alter_function production.
	ExitAlter_function(c *Alter_functionContext)

	// ExitAlter_function_signature is called when exiting the alter_function_signature production.
	ExitAlter_function_signature(c *Alter_function_signatureContext)

	// ExitData_type_list is called when exiting the data_type_list production.
	ExitData_type_list(c *Data_type_listContext)

	// ExitAlter_git_repository is called when exiting the alter_git_repository production.
	ExitAlter_git_repository(c *Alter_git_repositoryContext)

	// ExitAlter_git_set_opts is called when exiting the alter_git_set_opts production.
	ExitAlter_git_set_opts(c *Alter_git_set_optsContext)

	// ExitAlter_git_unset_opts is called when exiting the alter_git_unset_opts production.
	ExitAlter_git_unset_opts(c *Alter_git_unset_optsContext)

	// ExitAlter_masking_policy is called when exiting the alter_masking_policy production.
	ExitAlter_masking_policy(c *Alter_masking_policyContext)

	// ExitAlter_materialized_view is called when exiting the alter_materialized_view production.
	ExitAlter_materialized_view(c *Alter_materialized_viewContext)

	// ExitAlter_network_policy is called when exiting the alter_network_policy production.
	ExitAlter_network_policy(c *Alter_network_policyContext)

	// ExitAlter_notification_integration is called when exiting the alter_notification_integration production.
	ExitAlter_notification_integration(c *Alter_notification_integrationContext)

	// ExitAlter_pipe is called when exiting the alter_pipe production.
	ExitAlter_pipe(c *Alter_pipeContext)

	// ExitAlter_procedure is called when exiting the alter_procedure production.
	ExitAlter_procedure(c *Alter_procedureContext)

	// ExitAlter_replication_group is called when exiting the alter_replication_group production.
	ExitAlter_replication_group(c *Alter_replication_groupContext)

	// ExitCredit_quota is called when exiting the credit_quota production.
	ExitCredit_quota(c *Credit_quotaContext)

	// ExitFrequency is called when exiting the frequency production.
	ExitFrequency(c *FrequencyContext)

	// ExitNotify_users is called when exiting the notify_users production.
	ExitNotify_users(c *Notify_usersContext)

	// ExitTriggerDefinition is called when exiting the triggerDefinition production.
	ExitTriggerDefinition(c *TriggerDefinitionContext)

	// ExitAlter_resource_monitor is called when exiting the alter_resource_monitor production.
	ExitAlter_resource_monitor(c *Alter_resource_monitorContext)

	// ExitAlter_role is called when exiting the alter_role production.
	ExitAlter_role(c *Alter_roleContext)

	// ExitAlter_row_access_policy is called when exiting the alter_row_access_policy production.
	ExitAlter_row_access_policy(c *Alter_row_access_policyContext)

	// ExitAlter_schema is called when exiting the alter_schema production.
	ExitAlter_schema(c *Alter_schemaContext)

	// ExitSchema_property is called when exiting the schema_property production.
	ExitSchema_property(c *Schema_propertyContext)

	// ExitAlter_sequence is called when exiting the alter_sequence production.
	ExitAlter_sequence(c *Alter_sequenceContext)

	// ExitAlter_secret is called when exiting the alter_secret production.
	ExitAlter_secret(c *Alter_secretContext)

	// ExitSecret_opts is called when exiting the secret_opts production.
	ExitSecret_opts(c *Secret_optsContext)

	// ExitSecret_set_opts is called when exiting the secret_set_opts production.
	ExitSecret_set_opts(c *Secret_set_optsContext)

	// ExitSecret_oauth_client_creds_opts is called when exiting the secret_oauth_client_creds_opts production.
	ExitSecret_oauth_client_creds_opts(c *Secret_oauth_client_creds_optsContext)

	// ExitSecret_oauth_auth_code_opts is called when exiting the secret_oauth_auth_code_opts production.
	ExitSecret_oauth_auth_code_opts(c *Secret_oauth_auth_code_optsContext)

	// ExitSecret_api_auth_opts is called when exiting the secret_api_auth_opts production.
	ExitSecret_api_auth_opts(c *Secret_api_auth_optsContext)

	// ExitSecret_basic_auth_opts is called when exiting the secret_basic_auth_opts production.
	ExitSecret_basic_auth_opts(c *Secret_basic_auth_optsContext)

	// ExitSecret_generic_string_opts is called when exiting the secret_generic_string_opts production.
	ExitSecret_generic_string_opts(c *Secret_generic_string_optsContext)

	// ExitAlter_security_integration_external_oauth is called when exiting the alter_security_integration_external_oauth production.
	ExitAlter_security_integration_external_oauth(c *Alter_security_integration_external_oauthContext)

	// ExitSecurity_integration_external_oauth_property is called when exiting the security_integration_external_oauth_property production.
	ExitSecurity_integration_external_oauth_property(c *Security_integration_external_oauth_propertyContext)

	// ExitAlter_security_integration_snowflake_oauth is called when exiting the alter_security_integration_snowflake_oauth production.
	ExitAlter_security_integration_snowflake_oauth(c *Alter_security_integration_snowflake_oauthContext)

	// ExitSecurity_integration_snowflake_oauth_property is called when exiting the security_integration_snowflake_oauth_property production.
	ExitSecurity_integration_snowflake_oauth_property(c *Security_integration_snowflake_oauth_propertyContext)

	// ExitAlter_security_integration_saml2 is called when exiting the alter_security_integration_saml2 production.
	ExitAlter_security_integration_saml2(c *Alter_security_integration_saml2Context)

	// ExitAlter_security_integration_scim is called when exiting the alter_security_integration_scim production.
	ExitAlter_security_integration_scim(c *Alter_security_integration_scimContext)

	// ExitSecurity_integration_scim_property is called when exiting the security_integration_scim_property production.
	ExitSecurity_integration_scim_property(c *Security_integration_scim_propertyContext)

	// ExitAlter_session is called when exiting the alter_session production.
	ExitAlter_session(c *Alter_sessionContext)

	// ExitAlter_session_policy is called when exiting the alter_session_policy production.
	ExitAlter_session_policy(c *Alter_session_policyContext)

	// ExitAlter_password_policy is called when exiting the alter_password_policy production.
	ExitAlter_password_policy(c *Alter_password_policyContext)

	// ExitAlter_share is called when exiting the alter_share production.
	ExitAlter_share(c *Alter_shareContext)

	// ExitAlter_storage_integration is called when exiting the alter_storage_integration production.
	ExitAlter_storage_integration(c *Alter_storage_integrationContext)

	// ExitAlter_stream is called when exiting the alter_stream production.
	ExitAlter_stream(c *Alter_streamContext)

	// ExitAlter_table is called when exiting the alter_table production.
	ExitAlter_table(c *Alter_tableContext)

	// ExitRls_operations is called when exiting the rls_operations production.
	ExitRls_operations(c *Rls_operationsContext)

	// ExitClustering_action is called when exiting the clustering_action production.
	ExitClustering_action(c *Clustering_actionContext)

	// ExitTable_column_action is called when exiting the table_column_action production.
	ExitTable_column_action(c *Table_column_actionContext)

	// ExitAlter_column_clause is called when exiting the alter_column_clause production.
	ExitAlter_column_clause(c *Alter_column_clauseContext)

	// ExitInline_constraint is called when exiting the inline_constraint production.
	ExitInline_constraint(c *Inline_constraintContext)

	// ExitEnforced_not_enforced is called when exiting the enforced_not_enforced production.
	ExitEnforced_not_enforced(c *Enforced_not_enforcedContext)

	// ExitDeferrable_not_deferrable is called when exiting the deferrable_not_deferrable production.
	ExitDeferrable_not_deferrable(c *Deferrable_not_deferrableContext)

	// ExitInitially_deferred_or_immediate is called when exiting the initially_deferred_or_immediate production.
	ExitInitially_deferred_or_immediate(c *Initially_deferred_or_immediateContext)

	// ExitCommon_constraint_properties is called when exiting the common_constraint_properties production.
	ExitCommon_constraint_properties(c *Common_constraint_propertiesContext)

	// ExitOn_update is called when exiting the on_update production.
	ExitOn_update(c *On_updateContext)

	// ExitOn_delete is called when exiting the on_delete production.
	ExitOn_delete(c *On_deleteContext)

	// ExitForeign_key_match is called when exiting the foreign_key_match production.
	ExitForeign_key_match(c *Foreign_key_matchContext)

	// ExitOn_action is called when exiting the on_action production.
	ExitOn_action(c *On_actionContext)

	// ExitConstraint_properties is called when exiting the constraint_properties production.
	ExitConstraint_properties(c *Constraint_propertiesContext)

	// ExitExt_table_column_action is called when exiting the ext_table_column_action production.
	ExitExt_table_column_action(c *Ext_table_column_actionContext)

	// ExitConstraint_action is called when exiting the constraint_action production.
	ExitConstraint_action(c *Constraint_actionContext)

	// ExitSearch_optimization_action is called when exiting the search_optimization_action production.
	ExitSearch_optimization_action(c *Search_optimization_actionContext)

	// ExitSearch_method_with_target is called when exiting the search_method_with_target production.
	ExitSearch_method_with_target(c *Search_method_with_targetContext)

	// ExitAlter_table_alter_column is called when exiting the alter_table_alter_column production.
	ExitAlter_table_alter_column(c *Alter_table_alter_columnContext)

	// ExitAlter_column_decl_list is called when exiting the alter_column_decl_list production.
	ExitAlter_column_decl_list(c *Alter_column_decl_listContext)

	// ExitAlter_column_decl is called when exiting the alter_column_decl production.
	ExitAlter_column_decl(c *Alter_column_declContext)

	// ExitAlter_column_opts is called when exiting the alter_column_opts production.
	ExitAlter_column_opts(c *Alter_column_optsContext)

	// ExitColumn_set_tags is called when exiting the column_set_tags production.
	ExitColumn_set_tags(c *Column_set_tagsContext)

	// ExitColumn_unset_tags is called when exiting the column_unset_tags production.
	ExitColumn_unset_tags(c *Column_unset_tagsContext)

	// ExitAlter_tag is called when exiting the alter_tag production.
	ExitAlter_tag(c *Alter_tagContext)

	// ExitAlter_task is called when exiting the alter_task production.
	ExitAlter_task(c *Alter_taskContext)

	// ExitAlter_user is called when exiting the alter_user production.
	ExitAlter_user(c *Alter_userContext)

	// ExitAlter_view is called when exiting the alter_view production.
	ExitAlter_view(c *Alter_viewContext)

	// ExitAlter_modify is called when exiting the alter_modify production.
	ExitAlter_modify(c *Alter_modifyContext)

	// ExitAlter_warehouse is called when exiting the alter_warehouse production.
	ExitAlter_warehouse(c *Alter_warehouseContext)

	// ExitAlter_connection_opts is called when exiting the alter_connection_opts production.
	ExitAlter_connection_opts(c *Alter_connection_optsContext)

	// ExitAlter_user_opts is called when exiting the alter_user_opts production.
	ExitAlter_user_opts(c *Alter_user_optsContext)

	// ExitAlter_tag_opts is called when exiting the alter_tag_opts production.
	ExitAlter_tag_opts(c *Alter_tag_optsContext)

	// ExitAlter_network_policy_opts is called when exiting the alter_network_policy_opts production.
	ExitAlter_network_policy_opts(c *Alter_network_policy_optsContext)

	// ExitAlter_warehouse_opts is called when exiting the alter_warehouse_opts production.
	ExitAlter_warehouse_opts(c *Alter_warehouse_optsContext)

	// ExitAlter_account_opts is called when exiting the alter_account_opts production.
	ExitAlter_account_opts(c *Alter_account_optsContext)

	// ExitSet_tags is called when exiting the set_tags production.
	ExitSet_tags(c *Set_tagsContext)

	// ExitTag_decl_list is called when exiting the tag_decl_list production.
	ExitTag_decl_list(c *Tag_decl_listContext)

	// ExitUnset_tags is called when exiting the unset_tags production.
	ExitUnset_tags(c *Unset_tagsContext)

	// ExitTag_list is called when exiting the tag_list production.
	ExitTag_list(c *Tag_listContext)

	// ExitCreate_command is called when exiting the create_command production.
	ExitCreate_command(c *Create_commandContext)

	// ExitCreate_account is called when exiting the create_account production.
	ExitCreate_account(c *Create_accountContext)

	// ExitCreate_alert is called when exiting the create_alert production.
	ExitCreate_alert(c *Create_alertContext)

	// ExitAlert_condition is called when exiting the alert_condition production.
	ExitAlert_condition(c *Alert_conditionContext)

	// ExitAlert_action is called when exiting the alert_action production.
	ExitAlert_action(c *Alert_actionContext)

	// ExitCreate_api_integration is called when exiting the create_api_integration production.
	ExitCreate_api_integration(c *Create_api_integrationContext)

	// ExitCreate_object_clone is called when exiting the create_object_clone production.
	ExitCreate_object_clone(c *Create_object_cloneContext)

	// ExitCreate_connection is called when exiting the create_connection production.
	ExitCreate_connection(c *Create_connectionContext)

	// ExitCreate_database is called when exiting the create_database production.
	ExitCreate_database(c *Create_databaseContext)

	// ExitClone_at_before is called when exiting the clone_at_before production.
	ExitClone_at_before(c *Clone_at_beforeContext)

	// ExitAt_before1 is called when exiting the at_before1 production.
	ExitAt_before1(c *At_before1Context)

	// ExitHeader_decl is called when exiting the header_decl production.
	ExitHeader_decl(c *Header_declContext)

	// ExitCompression_type is called when exiting the compression_type production.
	ExitCompression_type(c *Compression_typeContext)

	// ExitCompression is called when exiting the compression production.
	ExitCompression(c *CompressionContext)

	// ExitCreate_dataset is called when exiting the create_dataset production.
	ExitCreate_dataset(c *Create_datasetContext)

	// ExitCreate_dynamic_table is called when exiting the create_dynamic_table production.
	ExitCreate_dynamic_table(c *Create_dynamic_tableContext)

	// ExitDynamic_table_params is called when exiting the dynamic_table_params production.
	ExitDynamic_table_params(c *Dynamic_table_paramsContext)

	// ExitDynamic_table_settable_params is called when exiting the dynamic_table_settable_params production.
	ExitDynamic_table_settable_params(c *Dynamic_table_settable_paramsContext)

	// ExitDynamic_table_unsettable_params is called when exiting the dynamic_table_unsettable_params production.
	ExitDynamic_table_unsettable_params(c *Dynamic_table_unsettable_paramsContext)

	// ExitData_retention_params is called when exiting the data_retention_params production.
	ExitData_retention_params(c *Data_retention_paramsContext)

	// ExitSet_data_retention_params is called when exiting the set_data_retention_params production.
	ExitSet_data_retention_params(c *Set_data_retention_paramsContext)

	// ExitCreate_event_table is called when exiting the create_event_table production.
	ExitCreate_event_table(c *Create_event_tableContext)

	// ExitCreate_external_function is called when exiting the create_external_function production.
	ExitCreate_external_function(c *Create_external_functionContext)

	// ExitCreate_external_table is called when exiting the create_external_table production.
	ExitCreate_external_table(c *Create_external_tableContext)

	// ExitExternal_table_column_decl is called when exiting the external_table_column_decl production.
	ExitExternal_table_column_decl(c *External_table_column_declContext)

	// ExitExternal_table_column_decl_list is called when exiting the external_table_column_decl_list production.
	ExitExternal_table_column_decl_list(c *External_table_column_decl_listContext)

	// ExitFull_acct is called when exiting the full_acct production.
	ExitFull_acct(c *Full_acctContext)

	// ExitIntegration_type_name is called when exiting the integration_type_name production.
	ExitIntegration_type_name(c *Integration_type_nameContext)

	// ExitCreate_failover_group is called when exiting the create_failover_group production.
	ExitCreate_failover_group(c *Create_failover_groupContext)

	// ExitType_fileformat is called when exiting the type_fileformat production.
	ExitType_fileformat(c *Type_fileformatContext)

	// ExitCreate_file_format is called when exiting the create_file_format production.
	ExitCreate_file_format(c *Create_file_formatContext)

	// ExitArg_decl is called when exiting the arg_decl production.
	ExitArg_decl(c *Arg_declContext)

	// ExitArg_default_value_clause is called when exiting the arg_default_value_clause production.
	ExitArg_default_value_clause(c *Arg_default_value_clauseContext)

	// ExitCol_decl is called when exiting the col_decl production.
	ExitCol_decl(c *Col_declContext)

	// ExitVirtual_column_decl is called when exiting the virtual_column_decl production.
	ExitVirtual_column_decl(c *Virtual_column_declContext)

	// ExitFunction_definition is called when exiting the function_definition production.
	ExitFunction_definition(c *Function_definitionContext)

	// ExitCreate_function is called when exiting the create_function production.
	ExitCreate_function(c *Create_functionContext)

	// ExitCreate_git_repository is called when exiting the create_git_repository production.
	ExitCreate_git_repository(c *Create_git_repositoryContext)

	// ExitCreate_git_opts is called when exiting the create_git_opts production.
	ExitCreate_git_opts(c *Create_git_optsContext)

	// ExitCreate_managed_account is called when exiting the create_managed_account production.
	ExitCreate_managed_account(c *Create_managed_accountContext)

	// ExitCreate_masking_policy is called when exiting the create_masking_policy production.
	ExitCreate_masking_policy(c *Create_masking_policyContext)

	// ExitTag_decl is called when exiting the tag_decl production.
	ExitTag_decl(c *Tag_declContext)

	// ExitColumn_list_in_parentheses is called when exiting the column_list_in_parentheses production.
	ExitColumn_list_in_parentheses(c *Column_list_in_parenthesesContext)

	// ExitCreate_materialized_view is called when exiting the create_materialized_view production.
	ExitCreate_materialized_view(c *Create_materialized_viewContext)

	// ExitCreate_network_policy is called when exiting the create_network_policy production.
	ExitCreate_network_policy(c *Create_network_policyContext)

	// ExitCloud_provider_params_auto is called when exiting the cloud_provider_params_auto production.
	ExitCloud_provider_params_auto(c *Cloud_provider_params_autoContext)

	// ExitCloud_provider_params_push is called when exiting the cloud_provider_params_push production.
	ExitCloud_provider_params_push(c *Cloud_provider_params_pushContext)

	// ExitCreate_notification_integration is called when exiting the create_notification_integration production.
	ExitCreate_notification_integration(c *Create_notification_integrationContext)

	// ExitCreate_pipe is called when exiting the create_pipe production.
	ExitCreate_pipe(c *Create_pipeContext)

	// ExitCaller_owner is called when exiting the caller_owner production.
	ExitCaller_owner(c *Caller_ownerContext)

	// ExitExecuta_as is called when exiting the executa_as production.
	ExitExecuta_as(c *Executa_asContext)

	// ExitProcedure_definition is called when exiting the procedure_definition production.
	ExitProcedure_definition(c *Procedure_definitionContext)

	// ExitNot_null is called when exiting the not_null production.
	ExitNot_null(c *Not_nullContext)

	// ExitCreate_procedure is called when exiting the create_procedure production.
	ExitCreate_procedure(c *Create_procedureContext)

	// ExitCreate_replication_group is called when exiting the create_replication_group production.
	ExitCreate_replication_group(c *Create_replication_groupContext)

	// ExitCreate_resource_monitor is called when exiting the create_resource_monitor production.
	ExitCreate_resource_monitor(c *Create_resource_monitorContext)

	// ExitCreate_role is called when exiting the create_role production.
	ExitCreate_role(c *Create_roleContext)

	// ExitCreate_row_access_policy is called when exiting the create_row_access_policy production.
	ExitCreate_row_access_policy(c *Create_row_access_policyContext)

	// ExitCreate_schema is called when exiting the create_schema production.
	ExitCreate_schema(c *Create_schemaContext)

	// ExitCreate_secret is called when exiting the create_secret production.
	ExitCreate_secret(c *Create_secretContext)

	// ExitCreate_security_integration_external_oauth is called when exiting the create_security_integration_external_oauth production.
	ExitCreate_security_integration_external_oauth(c *Create_security_integration_external_oauthContext)

	// ExitImplicit_none is called when exiting the implicit_none production.
	ExitImplicit_none(c *Implicit_noneContext)

	// ExitCreate_security_integration_snowflake_oauth is called when exiting the create_security_integration_snowflake_oauth production.
	ExitCreate_security_integration_snowflake_oauth(c *Create_security_integration_snowflake_oauthContext)

	// ExitCreate_security_integration_saml2 is called when exiting the create_security_integration_saml2 production.
	ExitCreate_security_integration_saml2(c *Create_security_integration_saml2Context)

	// ExitCreate_security_integration_scim is called when exiting the create_security_integration_scim production.
	ExitCreate_security_integration_scim(c *Create_security_integration_scimContext)

	// ExitNetwork_policy is called when exiting the network_policy production.
	ExitNetwork_policy(c *Network_policyContext)

	// ExitPartner_application is called when exiting the partner_application production.
	ExitPartner_application(c *Partner_applicationContext)

	// ExitStart_with is called when exiting the start_with production.
	ExitStart_with(c *Start_withContext)

	// ExitIncrement_by is called when exiting the increment_by production.
	ExitIncrement_by(c *Increment_byContext)

	// ExitCreate_semantic_view is called when exiting the create_semantic_view production.
	ExitCreate_semantic_view(c *Create_semantic_viewContext)

	// ExitLogical_table is called when exiting the logical_table production.
	ExitLogical_table(c *Logical_tableContext)

	// ExitRelationship_def is called when exiting the relationship_def production.
	ExitRelationship_def(c *Relationship_defContext)

	// ExitWith_synonyms_clause is called when exiting the with_synonyms_clause production.
	ExitWith_synonyms_clause(c *With_synonyms_clauseContext)

	// ExitSemantic_expression_list is called when exiting the semantic_expression_list production.
	ExitSemantic_expression_list(c *Semantic_expression_listContext)

	// ExitSemantic_expression is called when exiting the semantic_expression production.
	ExitSemantic_expression(c *Semantic_expressionContext)

	// ExitWith_extension_clause is called when exiting the with_extension_clause production.
	ExitWith_extension_clause(c *With_extension_clauseContext)

	// ExitCreate_sequence is called when exiting the create_sequence production.
	ExitCreate_sequence(c *Create_sequenceContext)

	// ExitCreate_session_policy is called when exiting the create_session_policy production.
	ExitCreate_session_policy(c *Create_session_policyContext)

	// ExitSession_policy_params is called when exiting the session_policy_params production.
	ExitSession_policy_params(c *Session_policy_paramsContext)

	// ExitSession_policy_param_name is called when exiting the session_policy_param_name production.
	ExitSession_policy_param_name(c *Session_policy_param_nameContext)

	// ExitCreate_password_policy is called when exiting the create_password_policy production.
	ExitCreate_password_policy(c *Create_password_policyContext)

	// ExitPassword_policy_params is called when exiting the password_policy_params production.
	ExitPassword_policy_params(c *Password_policy_paramsContext)

	// ExitPassword_policy_param_name is called when exiting the password_policy_param_name production.
	ExitPassword_policy_param_name(c *Password_policy_param_nameContext)

	// ExitCreate_share is called when exiting the create_share production.
	ExitCreate_share(c *Create_shareContext)

	// ExitCharacter is called when exiting the character production.
	ExitCharacter(c *CharacterContext)

	// ExitFormat_type_options is called when exiting the format_type_options production.
	ExitFormat_type_options(c *Format_type_optionsContext)

	// ExitCopy_options is called when exiting the copy_options production.
	ExitCopy_options(c *Copy_optionsContext)

	// ExitStage_encryption_opts_internal is called when exiting the stage_encryption_opts_internal production.
	ExitStage_encryption_opts_internal(c *Stage_encryption_opts_internalContext)

	// ExitStage_type is called when exiting the stage_type production.
	ExitStage_type(c *Stage_typeContext)

	// ExitStage_master_key is called when exiting the stage_master_key production.
	ExitStage_master_key(c *Stage_master_keyContext)

	// ExitStage_kms_key is called when exiting the stage_kms_key production.
	ExitStage_kms_key(c *Stage_kms_keyContext)

	// ExitStage_encryption_opts_aws is called when exiting the stage_encryption_opts_aws production.
	ExitStage_encryption_opts_aws(c *Stage_encryption_opts_awsContext)

	// ExitAws_token is called when exiting the aws_token production.
	ExitAws_token(c *Aws_tokenContext)

	// ExitAws_key_id is called when exiting the aws_key_id production.
	ExitAws_key_id(c *Aws_key_idContext)

	// ExitAws_secret_key is called when exiting the aws_secret_key production.
	ExitAws_secret_key(c *Aws_secret_keyContext)

	// ExitAws_role is called when exiting the aws_role production.
	ExitAws_role(c *Aws_roleContext)

	// ExitAzure_encryption_value is called when exiting the azure_encryption_value production.
	ExitAzure_encryption_value(c *Azure_encryption_valueContext)

	// ExitStage_encryption_opts_az is called when exiting the stage_encryption_opts_az production.
	ExitStage_encryption_opts_az(c *Stage_encryption_opts_azContext)

	// ExitStorage_integration_eq_id is called when exiting the storage_integration_eq_id production.
	ExitStorage_integration_eq_id(c *Storage_integration_eq_idContext)

	// ExitAz_credential_or_storage_integration is called when exiting the az_credential_or_storage_integration production.
	ExitAz_credential_or_storage_integration(c *Az_credential_or_storage_integrationContext)

	// ExitGcp_encryption_value is called when exiting the gcp_encryption_value production.
	ExitGcp_encryption_value(c *Gcp_encryption_valueContext)

	// ExitStage_encryption_opts_gcp is called when exiting the stage_encryption_opts_gcp production.
	ExitStage_encryption_opts_gcp(c *Stage_encryption_opts_gcpContext)

	// ExitAws_credential_or_storage_integration is called when exiting the aws_credential_or_storage_integration production.
	ExitAws_credential_or_storage_integration(c *Aws_credential_or_storage_integrationContext)

	// ExitExternal_stage_params is called when exiting the external_stage_params production.
	ExitExternal_stage_params(c *External_stage_paramsContext)

	// ExitTrue_false is called when exiting the true_false production.
	ExitTrue_false(c *True_falseContext)

	// ExitEnable is called when exiting the enable production.
	ExitEnable(c *EnableContext)

	// ExitRefresh_on_create is called when exiting the refresh_on_create production.
	ExitRefresh_on_create(c *Refresh_on_createContext)

	// ExitAuto_refresh is called when exiting the auto_refresh production.
	ExitAuto_refresh(c *Auto_refreshContext)

	// ExitNotification_integration is called when exiting the notification_integration production.
	ExitNotification_integration(c *Notification_integrationContext)

	// ExitDirectory_table_internal_params is called when exiting the directory_table_internal_params production.
	ExitDirectory_table_internal_params(c *Directory_table_internal_paramsContext)

	// ExitDirectory_table_external_params is called when exiting the directory_table_external_params production.
	ExitDirectory_table_external_params(c *Directory_table_external_paramsContext)

	// ExitCreate_stage is called when exiting the create_stage production.
	ExitCreate_stage(c *Create_stageContext)

	// ExitAlter_semantic_view is called when exiting the alter_semantic_view production.
	ExitAlter_semantic_view(c *Alter_semantic_viewContext)

	// ExitAlter_stage is called when exiting the alter_stage production.
	ExitAlter_stage(c *Alter_stageContext)

	// ExitDrop_stage is called when exiting the drop_stage production.
	ExitDrop_stage(c *Drop_stageContext)

	// ExitDescribe_stage is called when exiting the describe_stage production.
	ExitDescribe_stage(c *Describe_stageContext)

	// ExitShow_stages is called when exiting the show_stages production.
	ExitShow_stages(c *Show_stagesContext)

	// ExitCloud_provider_params is called when exiting the cloud_provider_params production.
	ExitCloud_provider_params(c *Cloud_provider_paramsContext)

	// ExitCloud_provider_params2 is called when exiting the cloud_provider_params2 production.
	ExitCloud_provider_params2(c *Cloud_provider_params2Context)

	// ExitCloud_provider_params3 is called when exiting the cloud_provider_params3 production.
	ExitCloud_provider_params3(c *Cloud_provider_params3Context)

	// ExitCreate_storage_integration is called when exiting the create_storage_integration production.
	ExitCreate_storage_integration(c *Create_storage_integrationContext)

	// ExitCopy_grants is called when exiting the copy_grants production.
	ExitCopy_grants(c *Copy_grantsContext)

	// ExitAppend_only is called when exiting the append_only production.
	ExitAppend_only(c *Append_onlyContext)

	// ExitInsert_only is called when exiting the insert_only production.
	ExitInsert_only(c *Insert_onlyContext)

	// ExitShow_initial_rows is called when exiting the show_initial_rows production.
	ExitShow_initial_rows(c *Show_initial_rowsContext)

	// ExitStream_time is called when exiting the stream_time production.
	ExitStream_time(c *Stream_timeContext)

	// ExitCreate_stream is called when exiting the create_stream production.
	ExitCreate_stream(c *Create_streamContext)

	// ExitTemporary is called when exiting the temporary production.
	ExitTemporary(c *TemporaryContext)

	// ExitTable_type is called when exiting the table_type production.
	ExitTable_type(c *Table_typeContext)

	// ExitWith_tags is called when exiting the with_tags production.
	ExitWith_tags(c *With_tagsContext)

	// ExitWith_row_access_policy is called when exiting the with_row_access_policy production.
	ExitWith_row_access_policy(c *With_row_access_policyContext)

	// ExitCluster_by is called when exiting the cluster_by production.
	ExitCluster_by(c *Cluster_byContext)

	// ExitChange_tracking is called when exiting the change_tracking production.
	ExitChange_tracking(c *Change_trackingContext)

	// ExitWith_masking_policy is called when exiting the with_masking_policy production.
	ExitWith_masking_policy(c *With_masking_policyContext)

	// ExitCollate is called when exiting the collate production.
	ExitCollate(c *CollateContext)

	// ExitOrder_noorder is called when exiting the order_noorder production.
	ExitOrder_noorder(c *Order_noorderContext)

	// ExitDefault_value is called when exiting the default_value production.
	ExitDefault_value(c *Default_valueContext)

	// ExitForeign_key is called when exiting the foreign_key production.
	ExitForeign_key(c *Foreign_keyContext)

	// ExitPrimary_key is called when exiting the primary_key production.
	ExitPrimary_key(c *Primary_keyContext)

	// ExitOut_of_line_constraint is called when exiting the out_of_line_constraint production.
	ExitOut_of_line_constraint(c *Out_of_line_constraintContext)

	// ExitFull_col_decl is called when exiting the full_col_decl production.
	ExitFull_col_decl(c *Full_col_declContext)

	// ExitMaterialized_col_decl is called when exiting the materialized_col_decl production.
	ExitMaterialized_col_decl(c *Materialized_col_declContext)

	// ExitMaterialized_col_decl_list is called when exiting the materialized_col_decl_list production.
	ExitMaterialized_col_decl_list(c *Materialized_col_decl_listContext)

	// ExitColumn_decl_item is called when exiting the column_decl_item production.
	ExitColumn_decl_item(c *Column_decl_itemContext)

	// ExitColumn_decl_item_list is called when exiting the column_decl_item_list production.
	ExitColumn_decl_item_list(c *Column_decl_item_listContext)

	// ExitCreate_table is called when exiting the create_table production.
	ExitCreate_table(c *Create_tableContext)

	// ExitColumn_decl_item_list_paren is called when exiting the column_decl_item_list_paren production.
	ExitColumn_decl_item_list_paren(c *Column_decl_item_list_parenContext)

	// ExitCreate_table_clause is called when exiting the create_table_clause production.
	ExitCreate_table_clause(c *Create_table_clauseContext)

	// ExitCreate_table_as_select is called when exiting the create_table_as_select production.
	ExitCreate_table_as_select(c *Create_table_as_selectContext)

	// ExitCreate_table_like is called when exiting the create_table_like production.
	ExitCreate_table_like(c *Create_table_likeContext)

	// ExitCreate_tag is called when exiting the create_tag production.
	ExitCreate_tag(c *Create_tagContext)

	// ExitTag_allowed_values is called when exiting the tag_allowed_values production.
	ExitTag_allowed_values(c *Tag_allowed_valuesContext)

	// ExitSession_parameter is called when exiting the session_parameter production.
	ExitSession_parameter(c *Session_parameterContext)

	// ExitSession_parameter_list is called when exiting the session_parameter_list production.
	ExitSession_parameter_list(c *Session_parameter_listContext)

	// ExitSession_params_list is called when exiting the session_params_list production.
	ExitSession_params_list(c *Session_params_listContext)

	// ExitCreate_task is called when exiting the create_task production.
	ExitCreate_task(c *Create_taskContext)

	// ExitTask_parameters is called when exiting the task_parameters production.
	ExitTask_parameters(c *Task_parametersContext)

	// ExitTask_compute is called when exiting the task_compute production.
	ExitTask_compute(c *Task_computeContext)

	// ExitTask_schedule is called when exiting the task_schedule production.
	ExitTask_schedule(c *Task_scheduleContext)

	// ExitTask_timeout is called when exiting the task_timeout production.
	ExitTask_timeout(c *Task_timeoutContext)

	// ExitTask_suspend_after_failure_number is called when exiting the task_suspend_after_failure_number production.
	ExitTask_suspend_after_failure_number(c *Task_suspend_after_failure_numberContext)

	// ExitTask_error_integration is called when exiting the task_error_integration production.
	ExitTask_error_integration(c *Task_error_integrationContext)

	// ExitTask_overlap is called when exiting the task_overlap production.
	ExitTask_overlap(c *Task_overlapContext)

	// ExitSql is called when exiting the sql production.
	ExitSql(c *SqlContext)

	// ExitTask_sql is called when exiting the task_sql production.
	ExitTask_sql(c *Task_sqlContext)

	// ExitTask_scripting_block is called when exiting the task_scripting_block production.
	ExitTask_scripting_block(c *Task_scripting_blockContext)

	// ExitTask_scripting_declaration_list is called when exiting the task_scripting_declaration_list production.
	ExitTask_scripting_declaration_list(c *Task_scripting_declaration_listContext)

	// ExitTask_scripting_declaration is called when exiting the task_scripting_declaration production.
	ExitTask_scripting_declaration(c *Task_scripting_declarationContext)

	// ExitTask_scripting_statement_list is called when exiting the task_scripting_statement_list production.
	ExitTask_scripting_statement_list(c *Task_scripting_statement_listContext)

	// ExitTask_scripting_statement is called when exiting the task_scripting_statement production.
	ExitTask_scripting_statement(c *Task_scripting_statementContext)

	// ExitTask_scripting_assignment is called when exiting the task_scripting_assignment production.
	ExitTask_scripting_assignment(c *Task_scripting_assignmentContext)

	// ExitTask_scripting_return is called when exiting the task_scripting_return production.
	ExitTask_scripting_return(c *Task_scripting_returnContext)

	// ExitCall is called when exiting the call production.
	ExitCall(c *CallContext)

	// ExitCreate_user is called when exiting the create_user production.
	ExitCreate_user(c *Create_userContext)

	// ExitView_col is called when exiting the view_col production.
	ExitView_col(c *View_colContext)

	// ExitCreate_view is called when exiting the create_view production.
	ExitCreate_view(c *Create_viewContext)

	// ExitCreate_warehouse is called when exiting the create_warehouse production.
	ExitCreate_warehouse(c *Create_warehouseContext)

	// ExitWh_common_size is called when exiting the wh_common_size production.
	ExitWh_common_size(c *Wh_common_sizeContext)

	// ExitWh_extra_size is called when exiting the wh_extra_size production.
	ExitWh_extra_size(c *Wh_extra_sizeContext)

	// ExitWh_properties is called when exiting the wh_properties production.
	ExitWh_properties(c *Wh_propertiesContext)

	// ExitWh_params is called when exiting the wh_params production.
	ExitWh_params(c *Wh_paramsContext)

	// ExitTrigger_definition is called when exiting the trigger_definition production.
	ExitTrigger_definition(c *Trigger_definitionContext)

	// ExitObject_type_name is called when exiting the object_type_name production.
	ExitObject_type_name(c *Object_type_nameContext)

	// ExitObject_type_plural is called when exiting the object_type_plural production.
	ExitObject_type_plural(c *Object_type_pluralContext)

	// ExitDrop_command is called when exiting the drop_command production.
	ExitDrop_command(c *Drop_commandContext)

	// ExitDrop_object is called when exiting the drop_object production.
	ExitDrop_object(c *Drop_objectContext)

	// ExitDrop_alert is called when exiting the drop_alert production.
	ExitDrop_alert(c *Drop_alertContext)

	// ExitDrop_connection is called when exiting the drop_connection production.
	ExitDrop_connection(c *Drop_connectionContext)

	// ExitDrop_database is called when exiting the drop_database production.
	ExitDrop_database(c *Drop_databaseContext)

	// ExitDrop_dynamic_table is called when exiting the drop_dynamic_table production.
	ExitDrop_dynamic_table(c *Drop_dynamic_tableContext)

	// ExitDrop_external_table is called when exiting the drop_external_table production.
	ExitDrop_external_table(c *Drop_external_tableContext)

	// ExitDrop_failover_group is called when exiting the drop_failover_group production.
	ExitDrop_failover_group(c *Drop_failover_groupContext)

	// ExitDrop_file_format is called when exiting the drop_file_format production.
	ExitDrop_file_format(c *Drop_file_formatContext)

	// ExitDrop_function is called when exiting the drop_function production.
	ExitDrop_function(c *Drop_functionContext)

	// ExitDrop_git_repository is called when exiting the drop_git_repository production.
	ExitDrop_git_repository(c *Drop_git_repositoryContext)

	// ExitDrop_integration is called when exiting the drop_integration production.
	ExitDrop_integration(c *Drop_integrationContext)

	// ExitDrop_managed_account is called when exiting the drop_managed_account production.
	ExitDrop_managed_account(c *Drop_managed_accountContext)

	// ExitDrop_masking_policy is called when exiting the drop_masking_policy production.
	ExitDrop_masking_policy(c *Drop_masking_policyContext)

	// ExitDrop_materialized_view is called when exiting the drop_materialized_view production.
	ExitDrop_materialized_view(c *Drop_materialized_viewContext)

	// ExitDrop_network_policy is called when exiting the drop_network_policy production.
	ExitDrop_network_policy(c *Drop_network_policyContext)

	// ExitDrop_pipe is called when exiting the drop_pipe production.
	ExitDrop_pipe(c *Drop_pipeContext)

	// ExitDrop_procedure is called when exiting the drop_procedure production.
	ExitDrop_procedure(c *Drop_procedureContext)

	// ExitDrop_replication_group is called when exiting the drop_replication_group production.
	ExitDrop_replication_group(c *Drop_replication_groupContext)

	// ExitDrop_resource_monitor is called when exiting the drop_resource_monitor production.
	ExitDrop_resource_monitor(c *Drop_resource_monitorContext)

	// ExitDrop_role is called when exiting the drop_role production.
	ExitDrop_role(c *Drop_roleContext)

	// ExitDrop_row_access_policy is called when exiting the drop_row_access_policy production.
	ExitDrop_row_access_policy(c *Drop_row_access_policyContext)

	// ExitDrop_schema is called when exiting the drop_schema production.
	ExitDrop_schema(c *Drop_schemaContext)

	// ExitDrop_secret is called when exiting the drop_secret production.
	ExitDrop_secret(c *Drop_secretContext)

	// ExitDrop_semantic_view is called when exiting the drop_semantic_view production.
	ExitDrop_semantic_view(c *Drop_semantic_viewContext)

	// ExitDrop_sequence is called when exiting the drop_sequence production.
	ExitDrop_sequence(c *Drop_sequenceContext)

	// ExitDrop_session_policy is called when exiting the drop_session_policy production.
	ExitDrop_session_policy(c *Drop_session_policyContext)

	// ExitDrop_password_policy is called when exiting the drop_password_policy production.
	ExitDrop_password_policy(c *Drop_password_policyContext)

	// ExitDrop_share is called when exiting the drop_share production.
	ExitDrop_share(c *Drop_shareContext)

	// ExitDrop_stream is called when exiting the drop_stream production.
	ExitDrop_stream(c *Drop_streamContext)

	// ExitDrop_table is called when exiting the drop_table production.
	ExitDrop_table(c *Drop_tableContext)

	// ExitDrop_tag is called when exiting the drop_tag production.
	ExitDrop_tag(c *Drop_tagContext)

	// ExitDrop_task is called when exiting the drop_task production.
	ExitDrop_task(c *Drop_taskContext)

	// ExitDrop_user is called when exiting the drop_user production.
	ExitDrop_user(c *Drop_userContext)

	// ExitDrop_view is called when exiting the drop_view production.
	ExitDrop_view(c *Drop_viewContext)

	// ExitDrop_warehouse is called when exiting the drop_warehouse production.
	ExitDrop_warehouse(c *Drop_warehouseContext)

	// ExitCascade_restrict is called when exiting the cascade_restrict production.
	ExitCascade_restrict(c *Cascade_restrictContext)

	// ExitArg_types is called when exiting the arg_types production.
	ExitArg_types(c *Arg_typesContext)

	// ExitUndrop_command is called when exiting the undrop_command production.
	ExitUndrop_command(c *Undrop_commandContext)

	// ExitUndrop_database is called when exiting the undrop_database production.
	ExitUndrop_database(c *Undrop_databaseContext)

	// ExitUndrop_dynamic_table is called when exiting the undrop_dynamic_table production.
	ExitUndrop_dynamic_table(c *Undrop_dynamic_tableContext)

	// ExitUndrop_schema is called when exiting the undrop_schema production.
	ExitUndrop_schema(c *Undrop_schemaContext)

	// ExitUndrop_table is called when exiting the undrop_table production.
	ExitUndrop_table(c *Undrop_tableContext)

	// ExitUndrop_tag is called when exiting the undrop_tag production.
	ExitUndrop_tag(c *Undrop_tagContext)

	// ExitUse_command is called when exiting the use_command production.
	ExitUse_command(c *Use_commandContext)

	// ExitUse_database is called when exiting the use_database production.
	ExitUse_database(c *Use_databaseContext)

	// ExitUse_role is called when exiting the use_role production.
	ExitUse_role(c *Use_roleContext)

	// ExitUse_schema is called when exiting the use_schema production.
	ExitUse_schema(c *Use_schemaContext)

	// ExitUse_secondary_roles is called when exiting the use_secondary_roles production.
	ExitUse_secondary_roles(c *Use_secondary_rolesContext)

	// ExitUse_warehouse is called when exiting the use_warehouse production.
	ExitUse_warehouse(c *Use_warehouseContext)

	// ExitComment_clause is called when exiting the comment_clause production.
	ExitComment_clause(c *Comment_clauseContext)

	// ExitInline_comment_clause is called when exiting the inline_comment_clause production.
	ExitInline_comment_clause(c *Inline_comment_clauseContext)

	// ExitIf_suspended is called when exiting the if_suspended production.
	ExitIf_suspended(c *If_suspendedContext)

	// ExitIf_exists is called when exiting the if_exists production.
	ExitIf_exists(c *If_existsContext)

	// ExitIf_not_exists is called when exiting the if_not_exists production.
	ExitIf_not_exists(c *If_not_existsContext)

	// ExitOr_replace is called when exiting the or_replace production.
	ExitOr_replace(c *Or_replaceContext)

	// ExitOr_alter is called when exiting the or_alter production.
	ExitOr_alter(c *Or_alterContext)

	// ExitDescribe is called when exiting the describe production.
	ExitDescribe(c *DescribeContext)

	// ExitDescribe_command is called when exiting the describe_command production.
	ExitDescribe_command(c *Describe_commandContext)

	// ExitDescribe_alert is called when exiting the describe_alert production.
	ExitDescribe_alert(c *Describe_alertContext)

	// ExitDescribe_database is called when exiting the describe_database production.
	ExitDescribe_database(c *Describe_databaseContext)

	// ExitDescribe_dynamic_table is called when exiting the describe_dynamic_table production.
	ExitDescribe_dynamic_table(c *Describe_dynamic_tableContext)

	// ExitDescribe_event_table is called when exiting the describe_event_table production.
	ExitDescribe_event_table(c *Describe_event_tableContext)

	// ExitDescribe_external_table is called when exiting the describe_external_table production.
	ExitDescribe_external_table(c *Describe_external_tableContext)

	// ExitDescribe_file_format is called when exiting the describe_file_format production.
	ExitDescribe_file_format(c *Describe_file_formatContext)

	// ExitDescribe_function is called when exiting the describe_function production.
	ExitDescribe_function(c *Describe_functionContext)

	// ExitDescribe_git_repository is called when exiting the describe_git_repository production.
	ExitDescribe_git_repository(c *Describe_git_repositoryContext)

	// ExitDescribe_integration is called when exiting the describe_integration production.
	ExitDescribe_integration(c *Describe_integrationContext)

	// ExitDescribe_masking_policy is called when exiting the describe_masking_policy production.
	ExitDescribe_masking_policy(c *Describe_masking_policyContext)

	// ExitDescribe_materialized_view is called when exiting the describe_materialized_view production.
	ExitDescribe_materialized_view(c *Describe_materialized_viewContext)

	// ExitDescribe_network_policy is called when exiting the describe_network_policy production.
	ExitDescribe_network_policy(c *Describe_network_policyContext)

	// ExitDescribe_pipe is called when exiting the describe_pipe production.
	ExitDescribe_pipe(c *Describe_pipeContext)

	// ExitDescribe_procedure is called when exiting the describe_procedure production.
	ExitDescribe_procedure(c *Describe_procedureContext)

	// ExitDescribe_result is called when exiting the describe_result production.
	ExitDescribe_result(c *Describe_resultContext)

	// ExitDescribe_row_access_policy is called when exiting the describe_row_access_policy production.
	ExitDescribe_row_access_policy(c *Describe_row_access_policyContext)

	// ExitDescribe_schema is called when exiting the describe_schema production.
	ExitDescribe_schema(c *Describe_schemaContext)

	// ExitDescribe_search_optimization is called when exiting the describe_search_optimization production.
	ExitDescribe_search_optimization(c *Describe_search_optimizationContext)

	// ExitDescribe_semantic_view is called when exiting the describe_semantic_view production.
	ExitDescribe_semantic_view(c *Describe_semantic_viewContext)

	// ExitDescribe_sequence is called when exiting the describe_sequence production.
	ExitDescribe_sequence(c *Describe_sequenceContext)

	// ExitDescribe_session_policy is called when exiting the describe_session_policy production.
	ExitDescribe_session_policy(c *Describe_session_policyContext)

	// ExitDescribe_password_policy is called when exiting the describe_password_policy production.
	ExitDescribe_password_policy(c *Describe_password_policyContext)

	// ExitDescribe_share is called when exiting the describe_share production.
	ExitDescribe_share(c *Describe_shareContext)

	// ExitDescribe_stream is called when exiting the describe_stream production.
	ExitDescribe_stream(c *Describe_streamContext)

	// ExitDescribe_table is called when exiting the describe_table production.
	ExitDescribe_table(c *Describe_tableContext)

	// ExitDescribe_task is called when exiting the describe_task production.
	ExitDescribe_task(c *Describe_taskContext)

	// ExitDescribe_transaction is called when exiting the describe_transaction production.
	ExitDescribe_transaction(c *Describe_transactionContext)

	// ExitDescribe_user is called when exiting the describe_user production.
	ExitDescribe_user(c *Describe_userContext)

	// ExitDescribe_view is called when exiting the describe_view production.
	ExitDescribe_view(c *Describe_viewContext)

	// ExitDescribe_warehouse is called when exiting the describe_warehouse production.
	ExitDescribe_warehouse(c *Describe_warehouseContext)

	// ExitShow_command is called when exiting the show_command production.
	ExitShow_command(c *Show_commandContext)

	// ExitShow_alerts is called when exiting the show_alerts production.
	ExitShow_alerts(c *Show_alertsContext)

	// ExitShow_channels is called when exiting the show_channels production.
	ExitShow_channels(c *Show_channelsContext)

	// ExitShow_columns is called when exiting the show_columns production.
	ExitShow_columns(c *Show_columnsContext)

	// ExitShow_connections is called when exiting the show_connections production.
	ExitShow_connections(c *Show_connectionsContext)

	// ExitStarts_with is called when exiting the starts_with production.
	ExitStarts_with(c *Starts_withContext)

	// ExitLimit_rows is called when exiting the limit_rows production.
	ExitLimit_rows(c *Limit_rowsContext)

	// ExitLimit_rows_2 is called when exiting the limit_rows_2 production.
	ExitLimit_rows_2(c *Limit_rows_2Context)

	// ExitShow_databases is called when exiting the show_databases production.
	ExitShow_databases(c *Show_databasesContext)

	// ExitShow_databases_in_failover_group is called when exiting the show_databases_in_failover_group production.
	ExitShow_databases_in_failover_group(c *Show_databases_in_failover_groupContext)

	// ExitShow_databases_in_replication_group is called when exiting the show_databases_in_replication_group production.
	ExitShow_databases_in_replication_group(c *Show_databases_in_replication_groupContext)

	// ExitShow_datasets is called when exiting the show_datasets production.
	ExitShow_datasets(c *Show_datasetsContext)

	// ExitShow_delegated_authorizations is called when exiting the show_delegated_authorizations production.
	ExitShow_delegated_authorizations(c *Show_delegated_authorizationsContext)

	// ExitShow_dynamic_tables is called when exiting the show_dynamic_tables production.
	ExitShow_dynamic_tables(c *Show_dynamic_tablesContext)

	// ExitShow_event_tables is called when exiting the show_event_tables production.
	ExitShow_event_tables(c *Show_event_tablesContext)

	// ExitShow_external_functions is called when exiting the show_external_functions production.
	ExitShow_external_functions(c *Show_external_functionsContext)

	// ExitShow_external_tables is called when exiting the show_external_tables production.
	ExitShow_external_tables(c *Show_external_tablesContext)

	// ExitShow_failover_groups is called when exiting the show_failover_groups production.
	ExitShow_failover_groups(c *Show_failover_groupsContext)

	// ExitShow_file_formats is called when exiting the show_file_formats production.
	ExitShow_file_formats(c *Show_file_formatsContext)

	// ExitShow_functions is called when exiting the show_functions production.
	ExitShow_functions(c *Show_functionsContext)

	// ExitShow_git_branches is called when exiting the show_git_branches production.
	ExitShow_git_branches(c *Show_git_branchesContext)

	// ExitShow_git_repositories is called when exiting the show_git_repositories production.
	ExitShow_git_repositories(c *Show_git_repositoriesContext)

	// ExitShow_git_tags is called when exiting the show_git_tags production.
	ExitShow_git_tags(c *Show_git_tagsContext)

	// ExitShow_global_accounts is called when exiting the show_global_accounts production.
	ExitShow_global_accounts(c *Show_global_accountsContext)

	// ExitShow_grants is called when exiting the show_grants production.
	ExitShow_grants(c *Show_grantsContext)

	// ExitShow_grants_opts is called when exiting the show_grants_opts production.
	ExitShow_grants_opts(c *Show_grants_optsContext)

	// ExitShow_integrations is called when exiting the show_integrations production.
	ExitShow_integrations(c *Show_integrationsContext)

	// ExitShow_locks is called when exiting the show_locks production.
	ExitShow_locks(c *Show_locksContext)

	// ExitShow_managed_accounts is called when exiting the show_managed_accounts production.
	ExitShow_managed_accounts(c *Show_managed_accountsContext)

	// ExitShow_masking_policies is called when exiting the show_masking_policies production.
	ExitShow_masking_policies(c *Show_masking_policiesContext)

	// ExitIn_obj is called when exiting the in_obj production.
	ExitIn_obj(c *In_objContext)

	// ExitIn_obj_2 is called when exiting the in_obj_2 production.
	ExitIn_obj_2(c *In_obj_2Context)

	// ExitIn_obj_3 is called when exiting the in_obj_3 production.
	ExitIn_obj_3(c *In_obj_3Context)

	// ExitShow_materialized_views is called when exiting the show_materialized_views production.
	ExitShow_materialized_views(c *Show_materialized_viewsContext)

	// ExitShow_network_policies is called when exiting the show_network_policies production.
	ExitShow_network_policies(c *Show_network_policiesContext)

	// ExitShow_objects is called when exiting the show_objects production.
	ExitShow_objects(c *Show_objectsContext)

	// ExitShow_organization_accounts is called when exiting the show_organization_accounts production.
	ExitShow_organization_accounts(c *Show_organization_accountsContext)

	// ExitIn_for is called when exiting the in_for production.
	ExitIn_for(c *In_forContext)

	// ExitShow_parameters is called when exiting the show_parameters production.
	ExitShow_parameters(c *Show_parametersContext)

	// ExitShow_pipes is called when exiting the show_pipes production.
	ExitShow_pipes(c *Show_pipesContext)

	// ExitShow_primary_keys is called when exiting the show_primary_keys production.
	ExitShow_primary_keys(c *Show_primary_keysContext)

	// ExitShow_procedures is called when exiting the show_procedures production.
	ExitShow_procedures(c *Show_proceduresContext)

	// ExitShow_regions is called when exiting the show_regions production.
	ExitShow_regions(c *Show_regionsContext)

	// ExitShow_replication_accounts is called when exiting the show_replication_accounts production.
	ExitShow_replication_accounts(c *Show_replication_accountsContext)

	// ExitShow_replication_databases is called when exiting the show_replication_databases production.
	ExitShow_replication_databases(c *Show_replication_databasesContext)

	// ExitShow_replication_groups is called when exiting the show_replication_groups production.
	ExitShow_replication_groups(c *Show_replication_groupsContext)

	// ExitShow_resource_monitors is called when exiting the show_resource_monitors production.
	ExitShow_resource_monitors(c *Show_resource_monitorsContext)

	// ExitShow_roles is called when exiting the show_roles production.
	ExitShow_roles(c *Show_rolesContext)

	// ExitShow_row_access_policies is called when exiting the show_row_access_policies production.
	ExitShow_row_access_policies(c *Show_row_access_policiesContext)

	// ExitShow_schemas is called when exiting the show_schemas production.
	ExitShow_schemas(c *Show_schemasContext)

	// ExitShow_secrets is called when exiting the show_secrets production.
	ExitShow_secrets(c *Show_secretsContext)

	// ExitShow_semantic_views is called when exiting the show_semantic_views production.
	ExitShow_semantic_views(c *Show_semantic_viewsContext)

	// ExitShow_semantic_dimensions is called when exiting the show_semantic_dimensions production.
	ExitShow_semantic_dimensions(c *Show_semantic_dimensionsContext)

	// ExitShow_semantic_dimensions_for_metric is called when exiting the show_semantic_dimensions_for_metric production.
	ExitShow_semantic_dimensions_for_metric(c *Show_semantic_dimensions_for_metricContext)

	// ExitShow_semantic_facts is called when exiting the show_semantic_facts production.
	ExitShow_semantic_facts(c *Show_semantic_factsContext)

	// ExitShow_semantic_metrics is called when exiting the show_semantic_metrics production.
	ExitShow_semantic_metrics(c *Show_semantic_metricsContext)

	// ExitShow_sequences is called when exiting the show_sequences production.
	ExitShow_sequences(c *Show_sequencesContext)

	// ExitShow_session_policies is called when exiting the show_session_policies production.
	ExitShow_session_policies(c *Show_session_policiesContext)

	// ExitShow_password_policies is called when exiting the show_password_policies production.
	ExitShow_password_policies(c *Show_password_policiesContext)

	// ExitShow_shares is called when exiting the show_shares production.
	ExitShow_shares(c *Show_sharesContext)

	// ExitShow_shares_in_failover_group is called when exiting the show_shares_in_failover_group production.
	ExitShow_shares_in_failover_group(c *Show_shares_in_failover_groupContext)

	// ExitShow_shares_in_replication_group is called when exiting the show_shares_in_replication_group production.
	ExitShow_shares_in_replication_group(c *Show_shares_in_replication_groupContext)

	// ExitShow_streams is called when exiting the show_streams production.
	ExitShow_streams(c *Show_streamsContext)

	// ExitShow_tables is called when exiting the show_tables production.
	ExitShow_tables(c *Show_tablesContext)

	// ExitShow_tags is called when exiting the show_tags production.
	ExitShow_tags(c *Show_tagsContext)

	// ExitShow_tasks is called when exiting the show_tasks production.
	ExitShow_tasks(c *Show_tasksContext)

	// ExitShow_transactions is called when exiting the show_transactions production.
	ExitShow_transactions(c *Show_transactionsContext)

	// ExitShow_user_functions is called when exiting the show_user_functions production.
	ExitShow_user_functions(c *Show_user_functionsContext)

	// ExitShow_users is called when exiting the show_users production.
	ExitShow_users(c *Show_usersContext)

	// ExitShow_variables is called when exiting the show_variables production.
	ExitShow_variables(c *Show_variablesContext)

	// ExitShow_versions_in_dataset is called when exiting the show_versions_in_dataset production.
	ExitShow_versions_in_dataset(c *Show_versions_in_datasetContext)

	// ExitShow_views is called when exiting the show_views production.
	ExitShow_views(c *Show_viewsContext)

	// ExitShow_warehouses is called when exiting the show_warehouses production.
	ExitShow_warehouses(c *Show_warehousesContext)

	// ExitLike_pattern is called when exiting the like_pattern production.
	ExitLike_pattern(c *Like_patternContext)

	// ExitAccount_identifier is called when exiting the account_identifier production.
	ExitAccount_identifier(c *Account_identifierContext)

	// ExitSchema_name is called when exiting the schema_name production.
	ExitSchema_name(c *Schema_nameContext)

	// ExitObject_type is called when exiting the object_type production.
	ExitObject_type(c *Object_typeContext)

	// ExitObject_type_list is called when exiting the object_type_list production.
	ExitObject_type_list(c *Object_type_listContext)

	// ExitTag_value is called when exiting the tag_value production.
	ExitTag_value(c *Tag_valueContext)

	// ExitArg_data_type is called when exiting the arg_data_type production.
	ExitArg_data_type(c *Arg_data_typeContext)

	// ExitArg_name is called when exiting the arg_name production.
	ExitArg_name(c *Arg_nameContext)

	// ExitParam_name is called when exiting the param_name production.
	ExitParam_name(c *Param_nameContext)

	// ExitRegion_group_id is called when exiting the region_group_id production.
	ExitRegion_group_id(c *Region_group_idContext)

	// ExitSnowflake_region_id is called when exiting the snowflake_region_id production.
	ExitSnowflake_region_id(c *Snowflake_region_idContext)

	// ExitString is called when exiting the string production.
	ExitString(c *StringContext)

	// ExitString_list is called when exiting the string_list production.
	ExitString_list(c *String_listContext)

	// ExitId_fn is called when exiting the id_fn production.
	ExitId_fn(c *Id_fnContext)

	// ExitId_ is called when exiting the id_ production.
	ExitId_(c *Id_Context)

	// ExitKeyword is called when exiting the keyword production.
	ExitKeyword(c *KeywordContext)

	// ExitNon_reserved_words is called when exiting the non_reserved_words production.
	ExitNon_reserved_words(c *Non_reserved_wordsContext)

	// ExitBuiltin_function is called when exiting the builtin_function production.
	ExitBuiltin_function(c *Builtin_functionContext)

	// ExitUnary_or_binary_builtin_function is called when exiting the unary_or_binary_builtin_function production.
	ExitUnary_or_binary_builtin_function(c *Unary_or_binary_builtin_functionContext)

	// ExitBinary_builtin_function is called when exiting the binary_builtin_function production.
	ExitBinary_builtin_function(c *Binary_builtin_functionContext)

	// ExitBinary_or_ternary_builtin_function is called when exiting the binary_or_ternary_builtin_function production.
	ExitBinary_or_ternary_builtin_function(c *Binary_or_ternary_builtin_functionContext)

	// ExitTernary_builtin_function is called when exiting the ternary_builtin_function production.
	ExitTernary_builtin_function(c *Ternary_builtin_functionContext)

	// ExitList_function is called when exiting the list_function production.
	ExitList_function(c *List_functionContext)

	// ExitPattern is called when exiting the pattern production.
	ExitPattern(c *PatternContext)

	// ExitColumn_name is called when exiting the column_name production.
	ExitColumn_name(c *Column_nameContext)

	// ExitColumn_list is called when exiting the column_list production.
	ExitColumn_list(c *Column_listContext)

	// ExitAliased_column_list is called when exiting the aliased_column_list production.
	ExitAliased_column_list(c *Aliased_column_listContext)

	// ExitColumn_list_with_comment is called when exiting the column_list_with_comment production.
	ExitColumn_list_with_comment(c *Column_list_with_commentContext)

	// ExitObject_name is called when exiting the object_name production.
	ExitObject_name(c *Object_nameContext)

	// ExitObject_name_or_identifier is called when exiting the object_name_or_identifier production.
	ExitObject_name_or_identifier(c *Object_name_or_identifierContext)

	// ExitNum is called when exiting the num production.
	ExitNum(c *NumContext)

	// ExitExpr_list is called when exiting the expr_list production.
	ExitExpr_list(c *Expr_listContext)

	// ExitExpr_list_sorted is called when exiting the expr_list_sorted production.
	ExitExpr_list_sorted(c *Expr_list_sortedContext)

	// ExitExpr is called when exiting the expr production.
	ExitExpr(c *ExprContext)

	// ExitLambda_params is called when exiting the lambda_params production.
	ExitLambda_params(c *Lambda_paramsContext)

	// ExitIff_expr is called when exiting the iff_expr production.
	ExitIff_expr(c *Iff_exprContext)

	// ExitTrim_expression is called when exiting the trim_expression production.
	ExitTrim_expression(c *Trim_expressionContext)

	// ExitTry_cast_expr is called when exiting the try_cast_expr production.
	ExitTry_cast_expr(c *Try_cast_exprContext)

	// ExitCast_expr is called when exiting the cast_expr production.
	ExitCast_expr(c *Cast_exprContext)

	// ExitJson_literal is called when exiting the json_literal production.
	ExitJson_literal(c *Json_literalContext)

	// ExitKv_pair is called when exiting the kv_pair production.
	ExitKv_pair(c *Kv_pairContext)

	// ExitValue is called when exiting the value production.
	ExitValue(c *ValueContext)

	// ExitArr_literal is called when exiting the arr_literal production.
	ExitArr_literal(c *Arr_literalContext)

	// ExitData_type_size is called when exiting the data_type_size production.
	ExitData_type_size(c *Data_type_sizeContext)

	// ExitData_type is called when exiting the data_type production.
	ExitData_type(c *Data_typeContext)

	// ExitVector_element_type is called when exiting the vector_element_type production.
	ExitVector_element_type(c *Vector_element_typeContext)

	// ExitPrimitive_expression is called when exiting the primitive_expression production.
	ExitPrimitive_expression(c *Primitive_expressionContext)

	// ExitOrder_by_expr is called when exiting the order_by_expr production.
	ExitOrder_by_expr(c *Order_by_exprContext)

	// ExitAsc_desc is called when exiting the asc_desc production.
	ExitAsc_desc(c *Asc_descContext)

	// ExitOver_clause is called when exiting the over_clause production.
	ExitOver_clause(c *Over_clauseContext)

	// ExitFunction_call is called when exiting the function_call production.
	ExitFunction_call(c *Function_callContext)

	// ExitParam_assoc_list is called when exiting the param_assoc_list production.
	ExitParam_assoc_list(c *Param_assoc_listContext)

	// ExitParam_assoc is called when exiting the param_assoc production.
	ExitParam_assoc(c *Param_assocContext)

	// ExitIgnore_or_repect_nulls is called when exiting the ignore_or_repect_nulls production.
	ExitIgnore_or_repect_nulls(c *Ignore_or_repect_nullsContext)

	// ExitRanking_windowed_function is called when exiting the ranking_windowed_function production.
	ExitRanking_windowed_function(c *Ranking_windowed_functionContext)

	// ExitAggregate_function is called when exiting the aggregate_function production.
	ExitAggregate_function(c *Aggregate_functionContext)

	// ExitLiteral is called when exiting the literal production.
	ExitLiteral(c *LiteralContext)

	// ExitSign is called when exiting the sign production.
	ExitSign(c *SignContext)

	// ExitFull_column_name is called when exiting the full_column_name production.
	ExitFull_column_name(c *Full_column_nameContext)

	// ExitBracket_expression is called when exiting the bracket_expression production.
	ExitBracket_expression(c *Bracket_expressionContext)

	// ExitCase_expression is called when exiting the case_expression production.
	ExitCase_expression(c *Case_expressionContext)

	// ExitSwitch_search_condition_section is called when exiting the switch_search_condition_section production.
	ExitSwitch_search_condition_section(c *Switch_search_condition_sectionContext)

	// ExitSwitch_section is called when exiting the switch_section production.
	ExitSwitch_section(c *Switch_sectionContext)

	// ExitQuery_statement is called when exiting the query_statement production.
	ExitQuery_statement(c *Query_statementContext)

	// ExitWith_expression is called when exiting the with_expression production.
	ExitWith_expression(c *With_expressionContext)

	// ExitCommon_table_expression is called when exiting the common_table_expression production.
	ExitCommon_table_expression(c *Common_table_expressionContext)

	// ExitSelect_statement is called when exiting the select_statement production.
	ExitSelect_statement(c *Select_statementContext)

	// ExitSet_operators is called when exiting the set_operators production.
	ExitSet_operators(c *Set_operatorsContext)

	// ExitBy_name is called when exiting the by_name production.
	ExitBy_name(c *By_nameContext)

	// ExitSelect_statement_in_parentheses is called when exiting the select_statement_in_parentheses production.
	ExitSelect_statement_in_parentheses(c *Select_statement_in_parenthesesContext)

	// ExitSelect_optional_clauses is called when exiting the select_optional_clauses production.
	ExitSelect_optional_clauses(c *Select_optional_clausesContext)

	// ExitSelect_clause is called when exiting the select_clause production.
	ExitSelect_clause(c *Select_clauseContext)

	// ExitSelect_top_clause is called when exiting the select_top_clause production.
	ExitSelect_top_clause(c *Select_top_clauseContext)

	// ExitSelect_list_no_top is called when exiting the select_list_no_top production.
	ExitSelect_list_no_top(c *Select_list_no_topContext)

	// ExitSelect_list_top is called when exiting the select_list_top production.
	ExitSelect_list_top(c *Select_list_topContext)

	// ExitSelect_list is called when exiting the select_list production.
	ExitSelect_list(c *Select_listContext)

	// ExitSelect_list_elem is called when exiting the select_list_elem production.
	ExitSelect_list_elem(c *Select_list_elemContext)

	// ExitColumn_elem_star is called when exiting the column_elem_star production.
	ExitColumn_elem_star(c *Column_elem_starContext)

	// ExitColumn_elem is called when exiting the column_elem production.
	ExitColumn_elem(c *Column_elemContext)

	// ExitObject_name_or_alias is called when exiting the object_name_or_alias production.
	ExitObject_name_or_alias(c *Object_name_or_aliasContext)

	// ExitExclude_clause is called when exiting the exclude_clause production.
	ExitExclude_clause(c *Exclude_clauseContext)

	// ExitAs_alias is called when exiting the as_alias production.
	ExitAs_alias(c *As_aliasContext)

	// ExitExpression_elem is called when exiting the expression_elem production.
	ExitExpression_elem(c *Expression_elemContext)

	// ExitColumn_position is called when exiting the column_position production.
	ExitColumn_position(c *Column_positionContext)

	// ExitAll_distinct is called when exiting the all_distinct production.
	ExitAll_distinct(c *All_distinctContext)

	// ExitTop_clause is called when exiting the top_clause production.
	ExitTop_clause(c *Top_clauseContext)

	// ExitInto_clause is called when exiting the into_clause production.
	ExitInto_clause(c *Into_clauseContext)

	// ExitVar_list is called when exiting the var_list production.
	ExitVar_list(c *Var_listContext)

	// ExitVar is called when exiting the var production.
	ExitVar(c *VarContext)

	// ExitFrom_clause is called when exiting the from_clause production.
	ExitFrom_clause(c *From_clauseContext)

	// ExitTable_sources is called when exiting the table_sources production.
	ExitTable_sources(c *Table_sourcesContext)

	// ExitTable_source is called when exiting the table_source production.
	ExitTable_source(c *Table_sourceContext)

	// ExitTable_source_item_joined is called when exiting the table_source_item_joined production.
	ExitTable_source_item_joined(c *Table_source_item_joinedContext)

	// ExitObject_ref is called when exiting the object_ref production.
	ExitObject_ref(c *Object_refContext)

	// ExitFlatten_table_option is called when exiting the flatten_table_option production.
	ExitFlatten_table_option(c *Flatten_table_optionContext)

	// ExitFlatten_table is called when exiting the flatten_table production.
	ExitFlatten_table(c *Flatten_tableContext)

	// ExitSplited_table is called when exiting the splited_table production.
	ExitSplited_table(c *Splited_tableContext)

	// ExitPrior_list is called when exiting the prior_list production.
	ExitPrior_list(c *Prior_listContext)

	// ExitPrior_item is called when exiting the prior_item production.
	ExitPrior_item(c *Prior_itemContext)

	// ExitOuter_join is called when exiting the outer_join production.
	ExitOuter_join(c *Outer_joinContext)

	// ExitJoin_type is called when exiting the join_type production.
	ExitJoin_type(c *Join_typeContext)

	// ExitJoin_clause is called when exiting the join_clause production.
	ExitJoin_clause(c *Join_clauseContext)

	// ExitOn_using_clause is called when exiting the on_using_clause production.
	ExitOn_using_clause(c *On_using_clauseContext)

	// ExitAt_before is called when exiting the at_before production.
	ExitAt_before(c *At_beforeContext)

	// ExitEnd is called when exiting the end production.
	ExitEnd(c *EndContext)

	// ExitChanges is called when exiting the changes production.
	ExitChanges(c *ChangesContext)

	// ExitDefault_append_only is called when exiting the default_append_only production.
	ExitDefault_append_only(c *Default_append_onlyContext)

	// ExitPartition_by is called when exiting the partition_by production.
	ExitPartition_by(c *Partition_byContext)

	// ExitAlias is called when exiting the alias production.
	ExitAlias(c *AliasContext)

	// ExitExpr_alias_list is called when exiting the expr_alias_list production.
	ExitExpr_alias_list(c *Expr_alias_listContext)

	// ExitMeasures is called when exiting the measures production.
	ExitMeasures(c *MeasuresContext)

	// ExitMatch_opts is called when exiting the match_opts production.
	ExitMatch_opts(c *Match_optsContext)

	// ExitRow_match is called when exiting the row_match production.
	ExitRow_match(c *Row_matchContext)

	// ExitFirst_last is called when exiting the first_last production.
	ExitFirst_last(c *First_lastContext)

	// ExitSymbol is called when exiting the symbol production.
	ExitSymbol(c *SymbolContext)

	// ExitAfter_match is called when exiting the after_match production.
	ExitAfter_match(c *After_matchContext)

	// ExitSymbol_list is called when exiting the symbol_list production.
	ExitSymbol_list(c *Symbol_listContext)

	// ExitDefine is called when exiting the define production.
	ExitDefine(c *DefineContext)

	// ExitMatch_recognize is called when exiting the match_recognize production.
	ExitMatch_recognize(c *Match_recognizeContext)

	// ExitPivot_unpivot is called when exiting the pivot_unpivot production.
	ExitPivot_unpivot(c *Pivot_unpivotContext)

	// ExitInclude_exclude is called when exiting the include_exclude production.
	ExitInclude_exclude(c *Include_excludeContext)

	// ExitPivot_in_clause is called when exiting the pivot_in_clause production.
	ExitPivot_in_clause(c *Pivot_in_clauseContext)

	// ExitDefault_on_null is called when exiting the default_on_null production.
	ExitDefault_on_null(c *Default_on_nullContext)

	// ExitColumn_alias_list_in_brackets is called when exiting the column_alias_list_in_brackets production.
	ExitColumn_alias_list_in_brackets(c *Column_alias_list_in_bracketsContext)

	// ExitExpr_list_in_parentheses is called when exiting the expr_list_in_parentheses production.
	ExitExpr_list_in_parentheses(c *Expr_list_in_parenthesesContext)

	// ExitValues_table is called when exiting the values_table production.
	ExitValues_table(c *Values_tableContext)

	// ExitValues_table_body is called when exiting the values_table_body production.
	ExitValues_table_body(c *Values_table_bodyContext)

	// ExitSample_method is called when exiting the sample_method production.
	ExitSample_method(c *Sample_methodContext)

	// ExitRepeatable_seed is called when exiting the repeatable_seed production.
	ExitRepeatable_seed(c *Repeatable_seedContext)

	// ExitSample_opts is called when exiting the sample_opts production.
	ExitSample_opts(c *Sample_optsContext)

	// ExitSample is called when exiting the sample production.
	ExitSample(c *SampleContext)

	// ExitSearch_condition is called when exiting the search_condition production.
	ExitSearch_condition(c *Search_conditionContext)

	// ExitComparison_operator is called when exiting the comparison_operator production.
	ExitComparison_operator(c *Comparison_operatorContext)

	// ExitNull_not_null is called when exiting the null_not_null production.
	ExitNull_not_null(c *Null_not_nullContext)

	// ExitNot_distinct_from is called when exiting the not_distinct_from production.
	ExitNot_distinct_from(c *Not_distinct_fromContext)

	// ExitSubquery is called when exiting the subquery production.
	ExitSubquery(c *SubqueryContext)

	// ExitPredicate is called when exiting the predicate production.
	ExitPredicate(c *PredicateContext)

	// ExitWhere_clause is called when exiting the where_clause production.
	ExitWhere_clause(c *Where_clauseContext)

	// ExitGroup_by_elem is called when exiting the group_by_elem production.
	ExitGroup_by_elem(c *Group_by_elemContext)

	// ExitGroup_by_list is called when exiting the group_by_list production.
	ExitGroup_by_list(c *Group_by_listContext)

	// ExitGroup_by_clause is called when exiting the group_by_clause production.
	ExitGroup_by_clause(c *Group_by_clauseContext)

	// ExitHaving_clause is called when exiting the having_clause production.
	ExitHaving_clause(c *Having_clauseContext)

	// ExitQualify_clause is called when exiting the qualify_clause production.
	ExitQualify_clause(c *Qualify_clauseContext)

	// ExitOrder_item is called when exiting the order_item production.
	ExitOrder_item(c *Order_itemContext)

	// ExitOrder_by_clause is called when exiting the order_by_clause production.
	ExitOrder_by_clause(c *Order_by_clauseContext)

	// ExitRow_rows is called when exiting the row_rows production.
	ExitRow_rows(c *Row_rowsContext)

	// ExitFirst_next is called when exiting the first_next production.
	ExitFirst_next(c *First_nextContext)

	// ExitLimit_clause is called when exiting the limit_clause production.
	ExitLimit_clause(c *Limit_clauseContext)

	// ExitRound_mode is called when exiting the round_mode production.
	ExitRound_mode(c *Round_modeContext)

	// ExitRound_expr is called when exiting the round_expr production.
	ExitRound_expr(c *Round_exprContext)

	// ExitSupplement_non_reserved_words is called when exiting the supplement_non_reserved_words production.
	ExitSupplement_non_reserved_words(c *Supplement_non_reserved_wordsContext)
}
