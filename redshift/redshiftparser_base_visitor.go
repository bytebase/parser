// Code generated from RedshiftParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package redshift // RedshiftParser
import "github.com/antlr4-go/antlr/v4"

type BaseRedshiftParserVisitor struct {
	*antlr.BaseParseTreeVisitor
}

func (v *BaseRedshiftParserVisitor) VisitRoot(ctx *RootContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitPlsqlroot(ctx *PlsqlrootContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitStmtblock(ctx *StmtblockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitStmtmulti(ctx *StmtmultiContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitStmt(ctx *StmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitPlsqlconsolecommand(ctx *PlsqlconsolecommandContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCallstmt(ctx *CallstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCreaterolestmt(ctx *CreaterolestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_with(ctx *Opt_withContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOptrolelist(ctx *OptrolelistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAlteroptrolelist(ctx *AlteroptrolelistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAlteroptroleelem(ctx *AlteroptroleelemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCreateoptroleelem(ctx *CreateoptroleelemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCreateuserstmt(ctx *CreateuserstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAlterrolestmt(ctx *AlterrolestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAlterroleaction(ctx *AlterroleactionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_in_database(ctx *Opt_in_databaseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAlterrolesetstmt(ctx *AlterrolesetstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAlterschemastmt(ctx *AlterschemastmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitDroprolestmt(ctx *DroprolestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitDropuserstmt(ctx *DropuserstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitDropgroupstmt(ctx *DropgroupstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCreategroupstmt(ctx *CreategroupstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAltergroupstmt(ctx *AltergroupstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAdd_drop(ctx *Add_dropContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCreateschemastmt(ctx *CreateschemastmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_auth_clause(ctx *Opt_auth_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_quota(ctx *Opt_quotaContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOptschemaeltlist(ctx *OptschemaeltlistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitSchema_stmt(ctx *Schema_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitVariablesetstmt(ctx *VariablesetstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitSet_rest(ctx *Set_restContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitGeneric_set(ctx *Generic_setContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitSet_rest_more(ctx *Set_rest_moreContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitVar_name(ctx *Var_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitVar_list(ctx *Var_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitVar_value(ctx *Var_valueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitIso_level(ctx *Iso_levelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_boolean_or_string(ctx *Opt_boolean_or_stringContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitZone_value(ctx *Zone_valueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_encoding(ctx *Opt_encodingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitNonreservedword_or_sconst(ctx *Nonreservedword_or_sconstContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitVariableresetstmt(ctx *VariableresetstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitReset_rest(ctx *Reset_restContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitGeneric_reset(ctx *Generic_resetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitSetresetclause(ctx *SetresetclauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitFunctionsetresetclause(ctx *FunctionsetresetclauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitVariableshowstmt(ctx *VariableshowstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitConstraintssetstmt(ctx *ConstraintssetstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitConstraints_set_list(ctx *Constraints_set_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitConstraints_set_mode(ctx *Constraints_set_modeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCheckpointstmt(ctx *CheckpointstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitDiscardstmt(ctx *DiscardstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAltertablestmt(ctx *AltertablestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAlter_table_cmds(ctx *Alter_table_cmdsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitTable_constraint(ctx *Table_constraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitPartition_cmd(ctx *Partition_cmdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitIndex_partition_cmd(ctx *Index_partition_cmdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAlter_table_cmd(ctx *Alter_table_cmdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAlter_column_default(ctx *Alter_column_defaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_drop_behavior(ctx *Opt_drop_behaviorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_collate_clause(ctx *Opt_collate_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAlter_using(ctx *Alter_usingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitReplica_identity(ctx *Replica_identityContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitReloptions(ctx *ReloptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_reloptions(ctx *Opt_reloptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitReloption_list(ctx *Reloption_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitReloption_elem(ctx *Reloption_elemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAlter_identity_column_option_list(ctx *Alter_identity_column_option_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAlter_identity_column_option(ctx *Alter_identity_column_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitPartitionboundspec(ctx *PartitionboundspecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitHash_partbound_elem(ctx *Hash_partbound_elemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitHash_partbound(ctx *Hash_partboundContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAltercompositetypestmt(ctx *AltercompositetypestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAlter_type_cmds(ctx *Alter_type_cmdsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAlter_type_cmd(ctx *Alter_type_cmdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCloseportalstmt(ctx *CloseportalstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCopystmt(ctx *CopystmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitRedshift_copy_authorization(ctx *Redshift_copy_authorizationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitRedshift_copy_format(ctx *Redshift_copy_formatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitRedshift_copy_parameter(ctx *Redshift_copy_parameterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCopy_param_name(ctx *Copy_param_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCopy_param_value(ctx *Copy_param_valueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCopy_from(ctx *Copy_fromContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_program(ctx *Opt_programContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCopy_file_name(ctx *Copy_file_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCopy_options(ctx *Copy_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCopy_opt_list(ctx *Copy_opt_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCopy_opt_item(ctx *Copy_opt_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_binary(ctx *Opt_binaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCopy_delimiter(ctx *Copy_delimiterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_using(ctx *Opt_usingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCopy_generic_opt_list(ctx *Copy_generic_opt_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCopy_generic_opt_elem(ctx *Copy_generic_opt_elemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCopy_generic_opt_arg(ctx *Copy_generic_opt_argContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCopy_generic_opt_arg_list(ctx *Copy_generic_opt_arg_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCopy_generic_opt_arg_list_item(ctx *Copy_generic_opt_arg_list_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCreatestmt(ctx *CreatestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_table_attributes(ctx *Opt_table_attributesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpttemp(ctx *OpttempContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpttableelementlist(ctx *OpttableelementlistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpttypedtableelementlist(ctx *OpttypedtableelementlistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitTableelementlist(ctx *TableelementlistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitTypedtableelementlist(ctx *TypedtableelementlistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitTableelement(ctx *TableelementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitTypedtableelement(ctx *TypedtableelementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitColumnDef(ctx *ColumnDefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitRs_colattributes(ctx *Rs_colattributesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitRs_colattribute(ctx *Rs_colattributeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitColumnOptions(ctx *ColumnOptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitColquallist(ctx *ColquallistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitColconstraint(ctx *ColconstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitColconstraintelem(ctx *ColconstraintelemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_unique_null_treatment(ctx *Opt_unique_null_treatmentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitGenerated_when(ctx *Generated_whenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitConstraintattr(ctx *ConstraintattrContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitTablelikeclause(ctx *TablelikeclauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitTablelikeoptionlist(ctx *TablelikeoptionlistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitTablelikeoption(ctx *TablelikeoptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitTableconstraint(ctx *TableconstraintContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitConstraintelem(ctx *ConstraintelemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_no_inherit(ctx *Opt_no_inheritContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_column_list(ctx *Opt_column_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitColumnlist(ctx *ColumnlistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitColumnElem(ctx *ColumnElemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_c_include(ctx *Opt_c_includeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitKey_match(ctx *Key_matchContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitExclusionconstraintlist(ctx *ExclusionconstraintlistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitExclusionconstraintelem(ctx *ExclusionconstraintelemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitExclusionwhereclause(ctx *ExclusionwhereclauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitKey_actions(ctx *Key_actionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitKey_update(ctx *Key_updateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitKey_delete(ctx *Key_deleteContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitKey_action(ctx *Key_actionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOptinherit(ctx *OptinheritContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOptpartitionspec(ctx *OptpartitionspecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitPartitionspec(ctx *PartitionspecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitPart_params(ctx *Part_paramsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitPart_elem(ctx *Part_elemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitTable_access_method_clause(ctx *Table_access_method_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOptwith(ctx *OptwithContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOncommitoption(ctx *OncommitoptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpttablespace(ctx *OpttablespaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOptredshifttableoptions(ctx *OptredshifttableoptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitRedshifttableoption(ctx *RedshifttableoptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitSortkeyclause(ctx *SortkeyclauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitSortkeyclausetype(ctx *SortkeyclausetypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOptconstablespace(ctx *OptconstablespaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitExistingindex(ctx *ExistingindexContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCreatestatsstmt(ctx *CreatestatsstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAlterstatsstmt(ctx *AlterstatsstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCreateasstmt(ctx *CreateasstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCreate_as_target(ctx *Create_as_targetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_backup_clause_table_attributes(ctx *Opt_backup_clause_table_attributesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitTable_attributes(ctx *Table_attributesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_backup_clause(ctx *Opt_backup_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_with_data(ctx *Opt_with_dataContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCreatematviewstmt(ctx *CreatematviewstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_auto_refresh(ctx *Opt_auto_refreshContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitRefreshmatviewstmt(ctx *RefreshmatviewstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCreateseqstmt(ctx *CreateseqstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAlterseqstmt(ctx *AlterseqstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOptseqoptlist(ctx *OptseqoptlistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOptparenthesizedseqoptlist(ctx *OptparenthesizedseqoptlistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitSeqoptlist(ctx *SeqoptlistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitSeqoptelem(ctx *SeqoptelemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_by(ctx *Opt_byContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitNumericonly(ctx *NumericonlyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitNumericonly_list(ctx *Numericonly_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCreateplangstmt(ctx *CreateplangstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_trusted(ctx *Opt_trustedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitHandler_name(ctx *Handler_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_inline_handler(ctx *Opt_inline_handlerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitValidator_clause(ctx *Validator_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_validator(ctx *Opt_validatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_procedural(ctx *Opt_proceduralContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCreatetablespacestmt(ctx *CreatetablespacestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpttablespaceowner(ctx *OpttablespaceownerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitDroptablespacestmt(ctx *DroptablespacestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCreateextensionstmt(ctx *CreateextensionstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCreate_extension_opt_list(ctx *Create_extension_opt_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCreate_extension_opt_item(ctx *Create_extension_opt_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAlterextensionstmt(ctx *AlterextensionstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAlter_extension_opt_list(ctx *Alter_extension_opt_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAlter_extension_opt_item(ctx *Alter_extension_opt_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAlterextensioncontentsstmt(ctx *AlterextensioncontentsstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCreatefdwstmt(ctx *CreatefdwstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitFdw_option(ctx *Fdw_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitFdw_options(ctx *Fdw_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_fdw_options(ctx *Opt_fdw_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAlterfdwstmt(ctx *AlterfdwstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCreate_generic_options(ctx *Create_generic_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitGeneric_option_list(ctx *Generic_option_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAlter_generic_options(ctx *Alter_generic_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAlter_generic_option_list(ctx *Alter_generic_option_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAlter_generic_option_elem(ctx *Alter_generic_option_elemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitGeneric_option_elem(ctx *Generic_option_elemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitGeneric_option_name(ctx *Generic_option_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitGeneric_option_arg(ctx *Generic_option_argContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCreateforeignserverstmt(ctx *CreateforeignserverstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_type(ctx *Opt_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitForeign_server_version(ctx *Foreign_server_versionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_foreign_server_version(ctx *Opt_foreign_server_versionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAlterforeignserverstmt(ctx *AlterforeignserverstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCreateforeigntablestmt(ctx *CreateforeigntablestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitImportforeignschemastmt(ctx *ImportforeignschemastmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitImport_qualification_type(ctx *Import_qualification_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitImport_qualification(ctx *Import_qualificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCreateusermappingstmt(ctx *CreateusermappingstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAuth_ident(ctx *Auth_identContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitDropusermappingstmt(ctx *DropusermappingstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAlterusermappingstmt(ctx *AlterusermappingstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCreatepolicystmt(ctx *CreatepolicystmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAlterpolicystmt(ctx *AlterpolicystmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitRowsecurityoptionalexpr(ctx *RowsecurityoptionalexprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitRowsecurityoptionalwithcheck(ctx *RowsecurityoptionalwithcheckContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitRowsecuritydefaulttorole(ctx *RowsecuritydefaulttoroleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitRowsecurityoptionaltorole(ctx *RowsecurityoptionaltoroleContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitRowsecuritydefaultpermissive(ctx *RowsecuritydefaultpermissiveContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitRowsecuritydefaultforcmd(ctx *RowsecuritydefaultforcmdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitRow_security_cmd(ctx *Row_security_cmdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCreateamstmt(ctx *CreateamstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAm_type(ctx *Am_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCreatetrigstmt(ctx *CreatetrigstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitTriggeractiontime(ctx *TriggeractiontimeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitTriggerevents(ctx *TriggereventsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitTriggeroneevent(ctx *TriggeroneeventContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitTriggerreferencing(ctx *TriggerreferencingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitTriggertransitions(ctx *TriggertransitionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitTriggertransition(ctx *TriggertransitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitTransitionoldornew(ctx *TransitionoldornewContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitTransitionrowortable(ctx *TransitionrowortableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitTransitionrelname(ctx *TransitionrelnameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitTriggerforspec(ctx *TriggerforspecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitTriggerforopteach(ctx *TriggerforopteachContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitTriggerfortype(ctx *TriggerfortypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitTriggerwhen(ctx *TriggerwhenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitFunction_or_procedure(ctx *Function_or_procedureContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitTriggerfuncargs(ctx *TriggerfuncargsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitTriggerfuncarg(ctx *TriggerfuncargContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOptconstrfromtable(ctx *OptconstrfromtableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitConstraintattributespec(ctx *ConstraintattributespecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitConstraintattributeElem(ctx *ConstraintattributeElemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCreateeventtrigstmt(ctx *CreateeventtrigstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitEvent_trigger_when_list(ctx *Event_trigger_when_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitEvent_trigger_when_item(ctx *Event_trigger_when_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitEvent_trigger_value_list(ctx *Event_trigger_value_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAltereventtrigstmt(ctx *AltereventtrigstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitEnable_trigger(ctx *Enable_triggerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCreateassertionstmt(ctx *CreateassertionstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitDefinestmt(ctx *DefinestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitDefinition(ctx *DefinitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitDef_list(ctx *Def_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitDef_elem(ctx *Def_elemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitDef_arg(ctx *Def_argContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOld_aggr_definition(ctx *Old_aggr_definitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOld_aggr_list(ctx *Old_aggr_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOld_aggr_elem(ctx *Old_aggr_elemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_enum_val_list(ctx *Opt_enum_val_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitEnum_val_list(ctx *Enum_val_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAlterenumstmt(ctx *AlterenumstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_if_not_exists(ctx *Opt_if_not_existsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCreateopclassstmt(ctx *CreateopclassstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpclass_item_list(ctx *Opclass_item_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpclass_item(ctx *Opclass_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_default(ctx *Opt_defaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_opfamily(ctx *Opt_opfamilyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpclass_purpose(ctx *Opclass_purposeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_recheck(ctx *Opt_recheckContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCreateopfamilystmt(ctx *CreateopfamilystmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAlteropfamilystmt(ctx *AlteropfamilystmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpclass_drop_list(ctx *Opclass_drop_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpclass_drop(ctx *Opclass_dropContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitDropopclassstmt(ctx *DropopclassstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitDropopfamilystmt(ctx *DropopfamilystmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitDropownedstmt(ctx *DropownedstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitReassignownedstmt(ctx *ReassignownedstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitDropstmt(ctx *DropstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitObject_type_any_name(ctx *Object_type_any_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitObject_type_name(ctx *Object_type_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitDrop_type_name(ctx *Drop_type_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitObject_type_name_on_any_name(ctx *Object_type_name_on_any_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAny_name_list(ctx *Any_name_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAny_name(ctx *Any_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAttrs(ctx *AttrsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitType_name_list(ctx *Type_name_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitTruncatestmt(ctx *TruncatestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_restart_seqs(ctx *Opt_restart_seqsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCommentstmt(ctx *CommentstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitComment_text(ctx *Comment_textContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitSeclabelstmt(ctx *SeclabelstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_provider(ctx *Opt_providerContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitSecurity_label(ctx *Security_labelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitFetchstmt(ctx *FetchstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitFetch_args(ctx *Fetch_argsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitGrantstmt(ctx *GrantstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitGrant_permissions_for_rls_lookup_tables(ctx *Grant_permissions_for_rls_lookup_tablesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitGrant_explain_permissions_for_row_level_security_policy_filters(ctx *Grant_explain_permissions_for_row_level_security_policy_filtersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitGrant_machine_learning_permissions(ctx *Grant_machine_learning_permissionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitGrant_role_permissions(ctx *Grant_role_permissionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitGrant_role_permission_target_list(ctx *Grant_role_permission_target_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitGrant_role_permission_target_list_item(ctx *Grant_role_permission_target_list_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitSystem_permissions(ctx *System_permissionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitSystem_permissions_item(ctx *System_permissions_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_with_admin_option(ctx *Opt_with_admin_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitGrant_scoped_permissions(ctx *Grant_scoped_permissionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitGrant_scoped_schemas_permissions(ctx *Grant_scoped_schemas_permissionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitGrant_scoped_tables_permissions(ctx *Grant_scoped_tables_permissionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitGrant_scoped_functions_permissions(ctx *Grant_scoped_functions_permissionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitGrant_scoped_procedures_permissions(ctx *Grant_scoped_procedures_permissionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitGrant_scoped_languages_permissions(ctx *Grant_scoped_languages_permissionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitGrant_scoped_copy_jobs_permissions(ctx *Grant_scoped_copy_jobs_permissionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitGrantee_list_without_public(ctx *Grantee_list_without_publicContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitGrantee_without_public(ctx *Grantee_without_publicContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitGrant_datashare_permissions(ctx *Grant_datashare_permissionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitGrant_spectrum_integration_permissions(ctx *Grant_spectrum_integration_permissionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitGrant_spectrum_integration_external_schema_permissions(ctx *Grant_spectrum_integration_external_schema_permissionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitSpectrum_integration_external_schema_permission_list(ctx *Spectrum_integration_external_schema_permission_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitSpectrum_integration_external_schema_permission(ctx *Spectrum_integration_external_schema_permissionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitGrant_spectrum_integration_external_table_permissions(ctx *Grant_spectrum_integration_external_table_permissionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitSpectrum_integration_external_table_permission(ctx *Spectrum_integration_external_table_permissionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitSpectrum_integration_external_table_permission_list(ctx *Spectrum_integration_external_table_permission_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitGrant_spectrum_integration_extenral_column_permissions(ctx *Grant_spectrum_integration_extenral_column_permissionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitIamrolelist_or_public(ctx *Iamrolelist_or_publicContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitIamrolelist(ctx *IamrolelistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitGrant_assume_role_permissions(ctx *Grant_assume_role_permissionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitGrant_assume_role_for_list(ctx *Grant_assume_role_for_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitGrant_assume_role_for_item(ctx *Grant_assume_role_for_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitGrant_assume_role_target(ctx *Grant_assume_role_targetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitGrant_column_level_permissions(ctx *Grant_column_level_permissionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitColumn_privilege_target(ctx *Column_privilege_targetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitColumn_privilege_list(ctx *Column_privilege_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitColumn_all_privilege(ctx *Column_all_privilegeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitColumn_select_update_privilege(ctx *Column_select_update_privilegeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCommon_grant(ctx *Common_grantContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCopy_job_privilege_list(ctx *Copy_job_privilege_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCopy_job_privilege(ctx *Copy_job_privilegeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCopy_job_target(ctx *Copy_job_targetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCopy_job_name(ctx *Copy_job_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitLanguage_privilege_list(ctx *Language_privilege_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitGrant_language_target(ctx *Grant_language_targetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitGrant_procedure_target(ctx *Grant_procedure_targetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitProcedure_privilege_list(ctx *Procedure_privilege_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitProcedure_privilege(ctx *Procedure_privilegeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitFunction_privilege_list(ctx *Function_privilege_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitFunction_privilege(ctx *Function_privilegeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitGrant_function_target(ctx *Grant_function_targetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitGrant_schema_target(ctx *Grant_schema_targetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitRevoke_schema_target(ctx *Revoke_schema_targetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitSchema_privilege_list(ctx *Schema_privilege_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitSchema_privilege(ctx *Schema_privilegeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitDatabase_privilege_list(ctx *Database_privilege_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitDatabase_privilege(ctx *Database_privilegeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitGrant_database_target(ctx *Grant_database_targetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitGrant_table_target(ctx *Grant_table_targetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitRevoke_table_target(ctx *Revoke_table_targetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAll_tables_in_schema_list(ctx *All_tables_in_schema_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAll_privileges(ctx *All_privilegesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitGrantee_list(ctx *Grantee_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitGrantee(ctx *GranteeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_with_grant_option(ctx *Opt_with_grant_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitTable_privilege(ctx *Table_privilegeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitTable_privilege_list(ctx *Table_privilege_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitRevokestmt(ctx *RevokestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitRevoke_permissions_for_rls_lookup_tables(ctx *Revoke_permissions_for_rls_lookup_tablesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitRevoke_explain_permissions_for_row_level_security_policy_filters(ctx *Revoke_explain_permissions_for_row_level_security_policy_filtersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitRevoke_machine_learning_permissions(ctx *Revoke_machine_learning_permissionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitRevoke_role_permissions(ctx *Revoke_role_permissionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitRevoke_scoped_permissions(ctx *Revoke_scoped_permissionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitRevoke_scoped_schemas_permissions(ctx *Revoke_scoped_schemas_permissionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitRevoke_scoped_tables_permissions(ctx *Revoke_scoped_tables_permissionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitRevoke_scoped_functions_permissions(ctx *Revoke_scoped_functions_permissionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitRevoke_scoped_procedures_permissions(ctx *Revoke_scoped_procedures_permissionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitRevoke_scoped_languages_permissions(ctx *Revoke_scoped_languages_permissionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitRevoke_scoped_copy_jobs_permissions(ctx *Revoke_scoped_copy_jobs_permissionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitRevoke_datashare_permissions(ctx *Revoke_datashare_permissionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitRevoke_spectrum_integration_permissions(ctx *Revoke_spectrum_integration_permissionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitRevoke_spectrum_integration_external_schema_permissions(ctx *Revoke_spectrum_integration_external_schema_permissionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitRevoke_spectrum_integration_external_table_permissions(ctx *Revoke_spectrum_integration_external_table_permissionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitRevoke_spectrum_integration_extenral_column_permissions(ctx *Revoke_spectrum_integration_extenral_column_permissionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitRevoke_assume_role_permissions(ctx *Revoke_assume_role_permissionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitRevoke_column_level_permissions(ctx *Revoke_column_level_permissionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCommon_revoke(ctx *Common_revokeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitPrivileges(ctx *PrivilegesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitPrivilege_list(ctx *Privilege_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitPrivilege(ctx *PrivilegeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitPrivilege_target(ctx *Privilege_targetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitParameter_name_list(ctx *Parameter_name_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitParameter_name(ctx *Parameter_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_grant_grant_option(ctx *Opt_grant_grant_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitGrantrolestmt(ctx *GrantrolestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitRevokerolestmt(ctx *RevokerolestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_grant_admin_option(ctx *Opt_grant_admin_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_granted_by(ctx *Opt_granted_byContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAlterdefaultprivilegesstmt(ctx *AlterdefaultprivilegesstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitDefacloptionlist(ctx *DefacloptionlistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitDefacloption(ctx *DefacloptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitDefaclaction(ctx *DefaclactionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitDefacl_privilege_target(ctx *Defacl_privilege_targetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitIndexstmt(ctx *IndexstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_unique(ctx *Opt_uniqueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_concurrently(ctx *Opt_concurrentlyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_index_name(ctx *Opt_index_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAccess_method_clause(ctx *Access_method_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitIndex_params(ctx *Index_paramsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitIndex_elem_options(ctx *Index_elem_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitIndex_elem(ctx *Index_elemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_include(ctx *Opt_includeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitIndex_including_params(ctx *Index_including_paramsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_collate(ctx *Opt_collateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_class(ctx *Opt_classContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_asc_desc(ctx *Opt_asc_descContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_nulls_order(ctx *Opt_nulls_orderContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCreatefunctionstmt(ctx *CreatefunctionstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCreateprocedurestmt(ctx *CreateprocedurestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_nonatomic(ctx *Opt_nonatomicContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_or_replace(ctx *Opt_or_replaceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitFunc_py_args_or_sql_args(ctx *Func_py_args_or_sql_argsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitFunc_py_args_or_sql_args_list(ctx *Func_py_args_or_sql_args_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitFunc_args(ctx *Func_argsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitFunc_args_list(ctx *Func_args_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitFunction_with_argtypes_list(ctx *Function_with_argtypes_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitFunction_with_argtypes(ctx *Function_with_argtypesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitFunc_args_with_defaults(ctx *Func_args_with_defaultsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitFunc_args_with_defaults_list(ctx *Func_args_with_defaults_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitFunc_arg(ctx *Func_argContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitArg_class(ctx *Arg_classContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitParam_name(ctx *Param_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitFunc_return(ctx *Func_returnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitFunc_type(ctx *Func_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitFunc_arg_with_default(ctx *Func_arg_with_defaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAggr_arg(ctx *Aggr_argContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAggr_args(ctx *Aggr_argsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAggr_args_list(ctx *Aggr_args_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAggregate_with_argtypes(ctx *Aggregate_with_argtypesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAggregate_with_argtypes_list(ctx *Aggregate_with_argtypes_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCreatefunc_opt_list(ctx *Createfunc_opt_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCommon_func_opt_item(ctx *Common_func_opt_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCreatefunc_opt_item(ctx *Createfunc_opt_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitFunc_as(ctx *Func_asContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitTransform_type_list(ctx *Transform_type_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_definition(ctx *Opt_definitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitTable_func_column(ctx *Table_func_columnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitTable_func_column_list(ctx *Table_func_column_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAlterfunctionstmt(ctx *AlterfunctionstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAlterprocedurestmt(ctx *AlterprocedurestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAlterfunc_opt_list(ctx *Alterfunc_opt_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_restrict(ctx *Opt_restrictContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitRemovefuncstmt(ctx *RemovefuncstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitRemoveaggrstmt(ctx *RemoveaggrstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitRemoveoperstmt(ctx *RemoveoperstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOper_argtypes(ctx *Oper_argtypesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAny_operator(ctx *Any_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOperator_with_argtypes_list(ctx *Operator_with_argtypes_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOperator_with_argtypes(ctx *Operator_with_argtypesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitDostmt(ctx *DostmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitDostmt_opt_list(ctx *Dostmt_opt_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitDostmt_opt_item(ctx *Dostmt_opt_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCreatecaststmt(ctx *CreatecaststmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCast_context(ctx *Cast_contextContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitDropcaststmt(ctx *DropcaststmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_if_exists(ctx *Opt_if_existsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCreatetransformstmt(ctx *CreatetransformstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitTransform_element_list(ctx *Transform_element_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitDroptransformstmt(ctx *DroptransformstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitReindexstmt(ctx *ReindexstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitReindex_target_type(ctx *Reindex_target_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitReindex_target_multitable(ctx *Reindex_target_multitableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitReindex_option_list(ctx *Reindex_option_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitReindex_option_elem(ctx *Reindex_option_elemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAltertblspcstmt(ctx *AltertblspcstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitRenamestmt(ctx *RenamestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_column(ctx *Opt_columnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_set_data(ctx *Opt_set_dataContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAlterobjectdependsstmt(ctx *AlterobjectdependsstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_no(ctx *Opt_noContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAlterobjectschemastmt(ctx *AlterobjectschemastmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAlteroperatorstmt(ctx *AlteroperatorstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOperator_def_list(ctx *Operator_def_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOperator_def_elem(ctx *Operator_def_elemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOperator_def_arg(ctx *Operator_def_argContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAltertypestmt(ctx *AltertypestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAlterownerstmt(ctx *AlterownerstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCreatepublicationstmt(ctx *CreatepublicationstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitPub_obj_list(ctx *Pub_obj_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitPublication_obj_spec(ctx *Publication_obj_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_where_clause(ctx *Opt_where_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAlterpublicationstmt(ctx *AlterpublicationstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCreatesubscriptionstmt(ctx *CreatesubscriptionstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitPublication_name_list(ctx *Publication_name_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitPublication_name_item(ctx *Publication_name_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAltersubscriptionstmt(ctx *AltersubscriptionstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitDropsubscriptionstmt(ctx *DropsubscriptionstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitRulestmt(ctx *RulestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitRuleactionlist(ctx *RuleactionlistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitRuleactionmulti(ctx *RuleactionmultiContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitRuleactionstmt(ctx *RuleactionstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitRuleactionstmtOrEmpty(ctx *RuleactionstmtOrEmptyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitEvent(ctx *EventContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_instead(ctx *Opt_insteadContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitNotifystmt(ctx *NotifystmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitNotify_payload(ctx *Notify_payloadContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitListenstmt(ctx *ListenstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitUnlistenstmt(ctx *UnlistenstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitTransactionstmt(ctx *TransactionstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_transaction(ctx *Opt_transactionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitTransaction_mode_item(ctx *Transaction_mode_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitTransaction_mode_list(ctx *Transaction_mode_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitTransaction_mode_list_or_empty(ctx *Transaction_mode_list_or_emptyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_transaction_chain(ctx *Opt_transaction_chainContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitViewstmt(ctx *ViewstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitWith_no_schema_binding(ctx *With_no_schema_bindingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_check_option(ctx *Opt_check_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitLoadstmt(ctx *LoadstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAlterdatasharestmt(ctx *AlterdatasharestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAlterdatashare_action(ctx *Alterdatashare_actionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAlterdatashare_add_drop(ctx *Alterdatashare_add_dropContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAlterdatashare_objects(ctx *Alterdatashare_objectsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitDatashare_table_list(ctx *Datashare_table_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitDatashare_table_name(ctx *Datashare_table_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitTable_name(ctx *Table_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitTemporary_table_name(ctx *Temporary_table_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitDatashare_function_list(ctx *Datashare_function_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitDatashare_function(ctx *Datashare_functionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitDatashare_function_name(ctx *Datashare_function_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCreatedatasharestmt(ctx *CreatedatasharestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCreatedatashareoptions(ctx *CreatedatashareoptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCreatedatashareoption(ctx *CreatedatashareoptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitSetpublicaccessibleoption(ctx *SetpublicaccessibleoptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitManagedbyoption(ctx *ManagedbyoptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitDescdatasharestmt(ctx *DescdatasharestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitDropdatasharestmt(ctx *DropdatasharestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAlterexternalschemastmt(ctx *AlterexternalschemastmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAltexternalschemaopts(ctx *AltexternalschemaoptsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAlterexternalviewstmt(ctx *AlterexternalviewstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCreateexternalschemastmt(ctx *CreateexternalschemastmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitFromdatacatalogclause(ctx *FromdatacatalogclauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitDropschemastmt(ctx *DropschemastmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitImplicitdatacatalogclause(ctx *ImplicitdatacatalogclauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitFromhivemetastoreclause(ctx *FromhivemetastoreclauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitFrompostgresclause(ctx *FrompostgresclauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitFrommysqlclause(ctx *FrommysqlclauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitFromkinesisclause(ctx *FromkinesisclauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitFromkafkaclause(ctx *FromkafkaclauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitFrommskclause(ctx *FrommskclauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitFromredshiftclause(ctx *FromredshiftclauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitIamrolevalue(ctx *IamrolevalueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCatalogrolevalue(ctx *CatalogrolevalueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAuthenticationvalue(ctx *AuthenticationvalueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCreateexternalfunctionstmt(ctx *CreateexternalfunctionstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitExternal_func_params(ctx *External_func_paramsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitParamlist(ctx *ParamlistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitParam_spec(ctx *Param_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCreateexternalmodelstmt(ctx *CreateexternalmodelstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCreateexternaltablestmt(ctx *CreateexternaltablestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitExtern_column_list(ctx *Extern_column_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitExtern_column_def(ctx *Extern_column_defContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitExtern_typename(ctx *Extern_typenameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitExtern_table_format(ctx *Extern_table_formatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitRow_format_spec(ctx *Row_format_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitSerde_properties_list(ctx *Serde_properties_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitSerde_property(ctx *Serde_propertyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitExternal_format_spec(ctx *External_format_specContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitTable_properties_list(ctx *Table_properties_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitTable_property(ctx *Table_propertyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCreateexternalviewstmt(ctx *CreateexternalviewstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitDropexternalviewstmt(ctx *DropexternalviewstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAlteridentityproviderstmt(ctx *AlteridentityproviderstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAlteridprovideropts(ctx *AlteridprovideroptsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAltermaskingpolicystmt(ctx *AltermaskingpolicystmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAltermaterializedviewstmt(ctx *AltermaterializedviewstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAltmaskingpolicyopts(ctx *AltmaskingpolicyoptsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAltmaskingpolicyargs(ctx *AltmaskingpolicyargsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAltmaskingpolicyarg(ctx *AltmaskingpolicyargContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAlterrlspolicystmt(ctx *AlterrlspolicystmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAttachmaskingpolicystmt(ctx *AttachmaskingpolicystmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAttachpolicycollist(ctx *AttachpolicycollistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAttachpolicycolumn(ctx *AttachpolicycolumnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAttachpolicytargets(ctx *AttachpolicytargetsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAttachpolicytarget(ctx *AttachpolicytargetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAttachrlspolicystmt(ctx *AttachrlspolicystmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitTable_name_list(ctx *Table_name_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCreateidentityproviderstmt(ctx *CreateidentityproviderstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCreateidprovideropts(ctx *CreateidprovideroptsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitGroupfilter(ctx *GroupfilterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCreatelibrarystmt(ctx *CreatelibrarystmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCreatelibraryopts(ctx *CreatelibraryoptsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCreatemaskingpolicystmt(ctx *CreatemaskingpolicystmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitInputcolumnlist(ctx *InputcolumnlistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitInputcolumn(ctx *InputcolumnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitMaskingexpression(ctx *MaskingexpressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCreatemodelstmt(ctx *CreatemodelstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCreatemodelfromclause(ctx *CreatemodelfromclauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitIamrolespec(ctx *IamrolespecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitSagemakerspec(ctx *SagemakerspecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitModeltypespec(ctx *ModeltypespecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitProblemtypespec(ctx *ProblemtypespecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitProblemtype(ctx *ProblemtypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitObjectivespec(ctx *ObjectivespecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitHyperparametersspec(ctx *HyperparametersspecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitHyperparameterslist(ctx *HyperparameterslistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitHyperparameteritem(ctx *HyperparameteritemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitSettingsclause(ctx *SettingsclauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitSettingsitem(ctx *SettingsitemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitDatatypelist(ctx *DatatypelistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitDatatype(ctx *DatatypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCreaterlspolicystmt(ctx *CreaterlspolicystmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitDescidentityproviderstmt(ctx *DescidentityproviderstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitDetachmaskingpolicystmt(ctx *DetachmaskingpolicystmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitDetachrlspolicystmt(ctx *DetachrlspolicystmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitRole_or_user_or_public_list(ctx *Role_or_user_or_public_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitRole_or_user_or_public(ctx *Role_or_user_or_publicContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitRlspolicyname(ctx *RlspolicynameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitDropidentityproviderstmt(ctx *DropidentityproviderstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitDroplibrarystmt(ctx *DroplibrarystmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitDropmaskingpolicystmt(ctx *DropmaskingpolicystmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitDropmodelstmt(ctx *DropmodelstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitDroprlspolicystmt(ctx *DroprlspolicystmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAltertableappendstmt(ctx *AltertableappendstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAppendoptions(ctx *AppendoptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAlteruserstmt(ctx *AlteruserstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAlteruseropts(ctx *AlteruseroptsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAnalyzecompressionstmt(ctx *AnalyzecompressionstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCancelstmt(ctx *CancelstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitClosestmt(ctx *ClosestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitInsertexternaltablestmt(ctx *InsertexternaltablestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitSelect_or_values(ctx *Select_or_valuesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitSelectintostmt(ctx *SelectintostmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitSetsessionauthorizationstmt(ctx *SetsessionauthorizationstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitSetsessioncharacteristicsstmt(ctx *SetsessioncharacteristicsstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitShowcolumnsstmt(ctx *ShowcolumnsstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitShowdatabasesstmt(ctx *ShowdatabasesstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitShowdbsopts(ctx *ShowdbsoptsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitShowdatasharesstmt(ctx *ShowdatasharesstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitShowexternaltablestmt(ctx *ShowexternaltablestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitShowgrantsstmt(ctx *ShowgrantsstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitGrantobject(ctx *GrantobjectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitGrantprincipal(ctx *GrantprincipalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitShowmodelstmt(ctx *ShowmodelstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitShowprocedurestmt(ctx *ShowprocedurestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitShowschemasstmt(ctx *ShowschemasstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitShowtablestmt(ctx *ShowtablestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitShowtablesstmt(ctx *ShowtablesstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitShowviewstmt(ctx *ShowviewstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitUnloadstmt(ctx *UnloadstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitIamroleclause(ctx *IamroleclauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitUnloadoptions(ctx *UnloadoptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitFormatoption(ctx *FormatoptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitPartitionbyoption(ctx *PartitionbyoptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitManifestoption(ctx *ManifestoptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitHeaderoption(ctx *HeaderoptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitDelimiteroption(ctx *DelimiteroptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitFixedwidthoption(ctx *FixedwidthoptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitEncryptedoption(ctx *EncryptedoptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitKmskeyoption(ctx *KmskeyoptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCompressionoption(ctx *CompressionoptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAddquotesoption(ctx *AddquotesoptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitNullasoption(ctx *NullasoptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitEscapeoption(ctx *EscapeoptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAllowoverwriteoption(ctx *AllowoverwriteoptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCleanpathoption(ctx *CleanpathoptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitParalleloption(ctx *ParalleloptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitMaxfilesizeoption(ctx *MaxfilesizeoptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitRowgroupsizeoption(ctx *RowgroupsizeoptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitSizeunit(ctx *SizeunitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitRegionoption(ctx *RegionoptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitExtensionoption(ctx *ExtensionoptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitUsestmt(ctx *UsestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCreatedbstmt(ctx *CreatedbstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCreatedb_opt_list(ctx *Createdb_opt_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCreatedb_opt_items(ctx *Createdb_opt_itemsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCreatedb_opt_item(ctx *Createdb_opt_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCreatedb_opt_name(ctx *Createdb_opt_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_equal(ctx *Opt_equalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAlterdatabasestmt(ctx *AlterdatabasestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAlterdatabasesetstmt(ctx *AlterdatabasesetstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitDropdbstmt(ctx *DropdbstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitDrop_option_list(ctx *Drop_option_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitDrop_option(ctx *Drop_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAltercollationstmt(ctx *AltercollationstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAltersystemstmt(ctx *AltersystemstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCreatedomainstmt(ctx *CreatedomainstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAlterdomainstmt(ctx *AlterdomainstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_as(ctx *Opt_asContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAltertsdictionarystmt(ctx *AltertsdictionarystmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAltertsconfigurationstmt(ctx *AltertsconfigurationstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAny_with(ctx *Any_withContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCreateconversionstmt(ctx *CreateconversionstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitClusterstmt(ctx *ClusterstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCluster_index_specification(ctx *Cluster_index_specificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitVacuumstmt(ctx *VacuumstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitVacuum_option(ctx *Vacuum_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAnalyzestmt(ctx *AnalyzestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitVac_analyze_option_list(ctx *Vac_analyze_option_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAnalyze_keyword(ctx *Analyze_keywordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitVac_analyze_option_elem(ctx *Vac_analyze_option_elemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitVac_analyze_option_name(ctx *Vac_analyze_option_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitVac_analyze_option_arg(ctx *Vac_analyze_option_argContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_analyze(ctx *Opt_analyzeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_verbose(ctx *Opt_verboseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_full(ctx *Opt_fullContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_freeze(ctx *Opt_freezeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_name_list(ctx *Opt_name_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitVacuum_relation(ctx *Vacuum_relationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitVacuum_relation_list(ctx *Vacuum_relation_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_vacuum_relation_list(ctx *Opt_vacuum_relation_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitExplainstmt(ctx *ExplainstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitExplainablestmt(ctx *ExplainablestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitExplain_option_list(ctx *Explain_option_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitExplain_option_elem(ctx *Explain_option_elemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitExplain_option_name(ctx *Explain_option_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitExplain_option_arg(ctx *Explain_option_argContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitPreparestmt(ctx *PreparestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitPrep_type_clause(ctx *Prep_type_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitPreparablestmt(ctx *PreparablestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitExecutestmt(ctx *ExecutestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitExecute_param_clause(ctx *Execute_param_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitDeallocatestmt(ctx *DeallocatestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitInsertstmt(ctx *InsertstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitInsert_target(ctx *Insert_targetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitInsert_rest(ctx *Insert_restContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOverride_kind(ctx *Override_kindContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitInsert_column_list(ctx *Insert_column_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitInsert_column_item(ctx *Insert_column_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_on_conflict(ctx *Opt_on_conflictContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_conf_expr(ctx *Opt_conf_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitReturning_clause(ctx *Returning_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitMergestmt(ctx *MergestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitMerge_when_clause(ctx *Merge_when_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitMerge_insert_clause(ctx *Merge_insert_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitMerge_update_clause(ctx *Merge_update_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitMerge_delete_clause(ctx *Merge_delete_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitDeletestmt(ctx *DeletestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitUsing_clause(ctx *Using_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitLockstmt(ctx *LockstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_lock(ctx *Opt_lockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitLock_type(ctx *Lock_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_nowait(ctx *Opt_nowaitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_nowait_or_skip(ctx *Opt_nowait_or_skipContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitUpdatestmt(ctx *UpdatestmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitSet_clause_list(ctx *Set_clause_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitSet_clause(ctx *Set_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitSet_target(ctx *Set_targetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitSet_target_list(ctx *Set_target_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitDeclarecursorstmt(ctx *DeclarecursorstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCursor_name(ctx *Cursor_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCursor_options(ctx *Cursor_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_hold(ctx *Opt_holdContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitSelectstmt(ctx *SelectstmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitSelect_with_parens(ctx *Select_with_parensContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitSelect_no_parens(ctx *Select_no_parensContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitSelect_clause(ctx *Select_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitSimple_select_intersect(ctx *Simple_select_intersectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitSimple_select_pramary(ctx *Simple_select_pramaryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitExclude_clause(ctx *Exclude_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitQualify_clause(ctx *Qualify_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitStart_with_clause(ctx *Start_with_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitWith_clause(ctx *With_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCte_list(ctx *Cte_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCommon_table_expr(ctx *Common_table_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_materialized(ctx *Opt_materializedContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_with_clause(ctx *Opt_with_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitInto_clause(ctx *Into_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_top_clause(ctx *Opt_top_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_strict(ctx *Opt_strictContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpttempTableName(ctx *OpttempTableNameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_table(ctx *Opt_tableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAll_or_distinct(ctx *All_or_distinctContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitDistinct_clause(ctx *Distinct_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_all_clause(ctx *Opt_all_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_sort_clause(ctx *Opt_sort_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitSort_clause(ctx *Sort_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitSortby_list(ctx *Sortby_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitSortby(ctx *SortbyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitSelect_limit(ctx *Select_limitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_select_limit(ctx *Opt_select_limitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitLimit_clause(ctx *Limit_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOffset_clause(ctx *Offset_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitSelect_limit_value(ctx *Select_limit_valueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitSelect_offset_value(ctx *Select_offset_valueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitSelect_fetch_first_value(ctx *Select_fetch_first_valueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitI_or_f_const(ctx *I_or_f_constContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitRow_or_rows(ctx *Row_or_rowsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitFirst_or_next(ctx *First_or_nextContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitGroup_clause(ctx *Group_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitGroup_by_list(ctx *Group_by_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitGroup_by_item(ctx *Group_by_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitEmpty_grouping_set(ctx *Empty_grouping_setContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitRollup_clause(ctx *Rollup_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCube_clause(ctx *Cube_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitGrouping_sets_clause(ctx *Grouping_sets_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitHaving_clause(ctx *Having_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitFor_locking_clause(ctx *For_locking_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_for_locking_clause(ctx *Opt_for_locking_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitFor_locking_items(ctx *For_locking_itemsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitFor_locking_item(ctx *For_locking_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitFor_locking_strength(ctx *For_locking_strengthContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitLocked_rels_list(ctx *Locked_rels_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitValues_clause(ctx *Values_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitFrom_clause(ctx *From_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitFrom_list(ctx *From_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitTable_ref(ctx *Table_refContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitJoined_table(ctx *Joined_tableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAlias_clause(ctx *Alias_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_alias_clause(ctx *Opt_alias_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitTable_alias_clause(ctx *Table_alias_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitFunc_alias_clause(ctx *Func_alias_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitJoin_type(ctx *Join_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitJoin_qual(ctx *Join_qualContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitRelation_expr(ctx *Relation_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitRelation_expr_list(ctx *Relation_expr_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitRelation_expr_opt_alias(ctx *Relation_expr_opt_aliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitTablesample_clause(ctx *Tablesample_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_repeatable_clause(ctx *Opt_repeatable_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitFunc_table(ctx *Func_tableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitRowsfrom_item(ctx *Rowsfrom_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitRowsfrom_list(ctx *Rowsfrom_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_col_def_list(ctx *Opt_col_def_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_ordinality(ctx *Opt_ordinalityContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitWhere_clause(ctx *Where_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitWhere_or_current_clause(ctx *Where_or_current_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpttablefuncelementlist(ctx *OpttablefuncelementlistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitTablefuncelementlist(ctx *TablefuncelementlistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitTablefuncelement(ctx *TablefuncelementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitXmltable(ctx *XmltableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitXmltable_column_list(ctx *Xmltable_column_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitXmltable_column_el(ctx *Xmltable_column_elContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitXmltable_column_option_list(ctx *Xmltable_column_option_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitXmltable_column_option_el(ctx *Xmltable_column_option_elContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitXml_namespace_list(ctx *Xml_namespace_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitXml_namespace_el(ctx *Xml_namespace_elContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitTypename(ctx *TypenameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_array_bounds(ctx *Opt_array_boundsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitSimpletypename(ctx *SimpletypenameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitVarbyte(ctx *VarbyteContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitJson_type(ctx *Json_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitConsttypename(ctx *ConsttypenameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitGenerictype(ctx *GenerictypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_type_modifiers(ctx *Opt_type_modifiersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitNumeric(ctx *NumericContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_float(ctx *Opt_floatContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitBit(ctx *BitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitConstbit(ctx *ConstbitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitBitwithlength(ctx *BitwithlengthContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitBitwithoutlength(ctx *BitwithoutlengthContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCharacter(ctx *CharacterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitConstcharacter(ctx *ConstcharacterContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCharacter_c(ctx *Character_cContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_varying(ctx *Opt_varyingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitConstdatetime(ctx *ConstdatetimeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitConstinterval(ctx *ConstintervalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_timezone(ctx *Opt_timezoneContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_interval(ctx *Opt_intervalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitInterval_second(ctx *Interval_secondContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_escape(ctx *Opt_escapeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitA_expr(ctx *A_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitA_expr_qual(ctx *A_expr_qualContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitA_expr_lessless(ctx *A_expr_lesslessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitA_expr_or(ctx *A_expr_orContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitA_expr_and(ctx *A_expr_andContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitA_expr_between(ctx *A_expr_betweenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitA_expr_in(ctx *A_expr_inContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitA_expr_unary_not(ctx *A_expr_unary_notContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitA_expr_isnull(ctx *A_expr_isnullContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitA_expr_is_not(ctx *A_expr_is_notContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitA_expr_compare(ctx *A_expr_compareContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitA_expr_prior_or_level(ctx *A_expr_prior_or_levelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitA_expr_like(ctx *A_expr_likeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitA_expr_qual_op(ctx *A_expr_qual_opContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitA_expr_unary_qualop(ctx *A_expr_unary_qualopContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitA_expr_add(ctx *A_expr_addContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitA_expr_mul(ctx *A_expr_mulContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitA_expr_caret(ctx *A_expr_caretContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitA_expr_unary_sign(ctx *A_expr_unary_signContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitA_expr_at_time_zone(ctx *A_expr_at_time_zoneContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitA_expr_collate(ctx *A_expr_collateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitA_expr_typecast(ctx *A_expr_typecastContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitB_expr(ctx *B_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitC_expr_exists(ctx *C_expr_existsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitC_expr_expr(ctx *C_expr_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitC_expr_case(ctx *C_expr_caseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitPlsqlvariablename(ctx *PlsqlvariablenameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitFunc_application(ctx *Func_applicationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitFunc_expr(ctx *Func_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitFunc_expr_windowless(ctx *Func_expr_windowlessContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitFunc_expr_common_subexpr(ctx *Func_expr_common_subexprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitXml_root_version(ctx *Xml_root_versionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_xml_root_standalone(ctx *Opt_xml_root_standaloneContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitXml_attributes(ctx *Xml_attributesContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitXml_attribute_list(ctx *Xml_attribute_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitXml_attribute_el(ctx *Xml_attribute_elContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitDocument_or_content(ctx *Document_or_contentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitXml_whitespace_option(ctx *Xml_whitespace_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitXmlexists_argument(ctx *Xmlexists_argumentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitXml_passing_mech(ctx *Xml_passing_mechContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitWithin_group_clause(ctx *Within_group_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitFilter_clause(ctx *Filter_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitWindow_clause(ctx *Window_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitWindow_definition_list(ctx *Window_definition_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitWindow_definition(ctx *Window_definitionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOver_clause(ctx *Over_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitWindow_specification(ctx *Window_specificationContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_existing_window_name(ctx *Opt_existing_window_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_partition_clause(ctx *Opt_partition_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_frame_clause(ctx *Opt_frame_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitFrame_extent(ctx *Frame_extentContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitFrame_bound(ctx *Frame_boundContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_window_exclusion_clause(ctx *Opt_window_exclusion_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitRow(ctx *RowContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitExplicit_row(ctx *Explicit_rowContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitImplicit_row(ctx *Implicit_rowContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitSub_type(ctx *Sub_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAll_op(ctx *All_opContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitMathop(ctx *MathopContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitQual_op(ctx *Qual_opContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitQual_all_op(ctx *Qual_all_opContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitSubquery_Op(ctx *Subquery_OpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitExpr_list(ctx *Expr_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitFunc_arg_list(ctx *Func_arg_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitFunc_arg_expr(ctx *Func_arg_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitType_list(ctx *Type_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitArray_expr(ctx *Array_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitArray_expr_list(ctx *Array_expr_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitExtract_list(ctx *Extract_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitExtract_arg(ctx *Extract_argContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitUnicode_normal_form(ctx *Unicode_normal_formContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOverlay_list(ctx *Overlay_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitPosition_list(ctx *Position_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitSubstr_list(ctx *Substr_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitTrim_list(ctx *Trim_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitIn_expr_select(ctx *In_expr_selectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitIn_expr_list(ctx *In_expr_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCase_expr(ctx *Case_exprContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitWhen_clause_list(ctx *When_clause_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitWhen_clause(ctx *When_clauseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCase_default(ctx *Case_defaultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCase_arg(ctx *Case_argContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitColumnref(ctx *ColumnrefContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitIndirection_el(ctx *Indirection_elContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_slice_bound(ctx *Opt_slice_boundContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitIndirection(ctx *IndirectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_indirection(ctx *Opt_indirectionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_target_list(ctx *Opt_target_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitTarget_list(ctx *Target_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitTarget_label(ctx *Target_labelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitTarget_star(ctx *Target_starContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitTarget_alias(ctx *Target_aliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitQualified_name_list(ctx *Qualified_name_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitQualified_name(ctx *Qualified_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitName_list(ctx *Name_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitName(ctx *NameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAttr_name(ctx *Attr_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitFile_name(ctx *File_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitFunc_name(ctx *Func_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAexprconst(ctx *AexprconstContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitXconst(ctx *XconstContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitBconst(ctx *BconstContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitFconst(ctx *FconstContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitIconst(ctx *IconstContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitSconst(ctx *SconstContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAnysconst(ctx *AnysconstContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_uescape(ctx *Opt_uescapeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitSignediconst(ctx *SignediconstContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitRoleid(ctx *RoleidContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitRolespec(ctx *RolespecContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitRole_list(ctx *Role_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitColid(ctx *ColidContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitTable_alias(ctx *Table_aliasContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitType_function_name(ctx *Type_function_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitNonreservedword(ctx *NonreservedwordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCollabel(ctx *CollabelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitIdentifier(ctx *IdentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitPlsqlidentifier(ctx *PlsqlidentifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitUnreserved_keyword(ctx *Unreserved_keywordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCol_name_keyword(ctx *Col_name_keywordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitType_func_name_keyword(ctx *Type_func_name_keywordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitReserved_keyword(ctx *Reserved_keywordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitBuiltin_function_name(ctx *Builtin_function_nameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitPl_function(ctx *Pl_functionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitComp_options(ctx *Comp_optionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitComp_option(ctx *Comp_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitSharp(ctx *SharpContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOption_value(ctx *Option_valueContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_semi(ctx *Opt_semiContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitPl_block(ctx *Pl_blockContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitDecl_sect(ctx *Decl_sectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitDecl_start(ctx *Decl_startContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitDecl_stmts(ctx *Decl_stmtsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitLabel_decl(ctx *Label_declContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitDecl_stmt(ctx *Decl_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitDecl_statement(ctx *Decl_statementContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_scrollable(ctx *Opt_scrollableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitDecl_cursor_query(ctx *Decl_cursor_queryContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitDecl_cursor_args(ctx *Decl_cursor_argsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitDecl_cursor_arglist(ctx *Decl_cursor_arglistContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitDecl_cursor_arg(ctx *Decl_cursor_argContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitDecl_is_for(ctx *Decl_is_forContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitDecl_aliasitem(ctx *Decl_aliasitemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitDecl_varname(ctx *Decl_varnameContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitDecl_const(ctx *Decl_constContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitDecl_datatype(ctx *Decl_datatypeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitDecl_collate(ctx *Decl_collateContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitDecl_notnull(ctx *Decl_notnullContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitDecl_defval(ctx *Decl_defvalContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitDecl_defkey(ctx *Decl_defkeyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAssign_operator(ctx *Assign_operatorContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitProc_sect(ctx *Proc_sectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitProc_stmt(ctx *Proc_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitStmt_perform(ctx *Stmt_performContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitStmt_call(ctx *Stmt_callContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_expr_list(ctx *Opt_expr_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitStmt_assign(ctx *Stmt_assignContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitStmt_getdiag(ctx *Stmt_getdiagContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitGetdiag_area_opt(ctx *Getdiag_area_optContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitGetdiag_list(ctx *Getdiag_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitGetdiag_list_item(ctx *Getdiag_list_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitGetdiag_item(ctx *Getdiag_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitGetdiag_target(ctx *Getdiag_targetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAssign_var(ctx *Assign_varContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitStmt_if(ctx *Stmt_ifContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitStmt_elsifs(ctx *Stmt_elsifsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitStmt_else(ctx *Stmt_elseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitStmt_case(ctx *Stmt_caseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_expr_until_when(ctx *Opt_expr_until_whenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCase_when_list(ctx *Case_when_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCase_when(ctx *Case_whenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_case_else(ctx *Opt_case_elseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitStmt_loop(ctx *Stmt_loopContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitStmt_while(ctx *Stmt_whileContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitStmt_for(ctx *Stmt_forContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitFor_control(ctx *For_controlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_for_using_expression(ctx *Opt_for_using_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_cursor_parameters(ctx *Opt_cursor_parametersContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_reverse(ctx *Opt_reverseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_by_expression(ctx *Opt_by_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitFor_variable(ctx *For_variableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitStmt_foreach_a(ctx *Stmt_foreach_aContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitForeach_slice(ctx *Foreach_sliceContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitStmt_exit(ctx *Stmt_exitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitExit_type(ctx *Exit_typeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitStmt_return(ctx *Stmt_returnContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_return_result(ctx *Opt_return_resultContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitStmt_raise(ctx *Stmt_raiseContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_stmt_raise_level(ctx *Opt_stmt_raise_levelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_raise_list(ctx *Opt_raise_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_raise_using(ctx *Opt_raise_usingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_raise_using_elem(ctx *Opt_raise_using_elemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_raise_using_elem_list(ctx *Opt_raise_using_elem_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitStmt_assert(ctx *Stmt_assertContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_stmt_assert_message(ctx *Opt_stmt_assert_messageContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitLoop_body(ctx *Loop_bodyContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitStmt_execsql(ctx *Stmt_execsqlContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitStmt_dynexecute(ctx *Stmt_dynexecuteContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_execute_using(ctx *Opt_execute_usingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_execute_using_list(ctx *Opt_execute_using_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_execute_into(ctx *Opt_execute_intoContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitStmt_open(ctx *Stmt_openContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_open_bound_list_item(ctx *Opt_open_bound_list_itemContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_open_bound_list(ctx *Opt_open_bound_listContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_open_using(ctx *Opt_open_usingContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_scroll_option(ctx *Opt_scroll_optionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_scroll_option_no(ctx *Opt_scroll_option_noContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitStmt_fetch(ctx *Stmt_fetchContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitInto_target(ctx *Into_targetContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_cursor_from(ctx *Opt_cursor_fromContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_fetch_direction(ctx *Opt_fetch_directionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitStmt_move(ctx *Stmt_moveContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitStmt_close(ctx *Stmt_closeContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitStmt_null(ctx *Stmt_nullContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitStmt_commit(ctx *Stmt_commitContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitStmt_rollback(ctx *Stmt_rollbackContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitPlsql_opt_transaction_chain(ctx *Plsql_opt_transaction_chainContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitStmt_set(ctx *Stmt_setContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitCursor_variable(ctx *Cursor_variableContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitException_sect(ctx *Exception_sectContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitProc_exceptions(ctx *Proc_exceptionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitProc_exception(ctx *Proc_exceptionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitProc_conditions(ctx *Proc_conditionsContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitProc_condition(ctx *Proc_conditionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_block_label(ctx *Opt_block_labelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_loop_label(ctx *Opt_loop_labelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_label(ctx *Opt_labelContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_exitcond(ctx *Opt_exitcondContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitAny_identifier(ctx *Any_identifierContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitPlsql_unreserved_keyword(ctx *Plsql_unreserved_keywordContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitSql_expression(ctx *Sql_expressionContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitExpr_until_then(ctx *Expr_until_thenContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitExpr_until_semi(ctx *Expr_until_semiContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitExpr_until_rightbracket(ctx *Expr_until_rightbracketContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitExpr_until_loop(ctx *Expr_until_loopContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitMake_execsql_stmt(ctx *Make_execsql_stmtContext) interface{} {
	return v.VisitChildren(ctx)
}

func (v *BaseRedshiftParserVisitor) VisitOpt_returning_clause_into(ctx *Opt_returning_clause_intoContext) interface{} {
	return v.VisitChildren(ctx)
}
