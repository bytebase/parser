// Code generated from RedshiftParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package redshift // RedshiftParser
import "github.com/antlr4-go/antlr/v4"

// A complete Visitor for a parse tree produced by RedshiftParser.
type RedshiftParserVisitor interface {
	antlr.ParseTreeVisitor

	// Visit a parse tree produced by RedshiftParser#root.
	VisitRoot(ctx *RootContext) interface{}

	// Visit a parse tree produced by RedshiftParser#plsqlroot.
	VisitPlsqlroot(ctx *PlsqlrootContext) interface{}

	// Visit a parse tree produced by RedshiftParser#stmtblock.
	VisitStmtblock(ctx *StmtblockContext) interface{}

	// Visit a parse tree produced by RedshiftParser#stmtmulti.
	VisitStmtmulti(ctx *StmtmultiContext) interface{}

	// Visit a parse tree produced by RedshiftParser#stmt.
	VisitStmt(ctx *StmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#plsqlconsolecommand.
	VisitPlsqlconsolecommand(ctx *PlsqlconsolecommandContext) interface{}

	// Visit a parse tree produced by RedshiftParser#callstmt.
	VisitCallstmt(ctx *CallstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#createrolestmt.
	VisitCreaterolestmt(ctx *CreaterolestmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_with.
	VisitOpt_with(ctx *Opt_withContext) interface{}

	// Visit a parse tree produced by RedshiftParser#optrolelist.
	VisitOptrolelist(ctx *OptrolelistContext) interface{}

	// Visit a parse tree produced by RedshiftParser#alteroptrolelist.
	VisitAlteroptrolelist(ctx *AlteroptrolelistContext) interface{}

	// Visit a parse tree produced by RedshiftParser#alteroptroleelem.
	VisitAlteroptroleelem(ctx *AlteroptroleelemContext) interface{}

	// Visit a parse tree produced by RedshiftParser#createoptroleelem.
	VisitCreateoptroleelem(ctx *CreateoptroleelemContext) interface{}

	// Visit a parse tree produced by RedshiftParser#createuserstmt.
	VisitCreateuserstmt(ctx *CreateuserstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#alterrolestmt.
	VisitAlterrolestmt(ctx *AlterrolestmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#alterroleaction.
	VisitAlterroleaction(ctx *AlterroleactionContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_in_database.
	VisitOpt_in_database(ctx *Opt_in_databaseContext) interface{}

	// Visit a parse tree produced by RedshiftParser#alterrolesetstmt.
	VisitAlterrolesetstmt(ctx *AlterrolesetstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#alterschemastmt.
	VisitAlterschemastmt(ctx *AlterschemastmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#droprolestmt.
	VisitDroprolestmt(ctx *DroprolestmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#dropuserstmt.
	VisitDropuserstmt(ctx *DropuserstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#dropgroupstmt.
	VisitDropgroupstmt(ctx *DropgroupstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#creategroupstmt.
	VisitCreategroupstmt(ctx *CreategroupstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#altergroupstmt.
	VisitAltergroupstmt(ctx *AltergroupstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#add_drop.
	VisitAdd_drop(ctx *Add_dropContext) interface{}

	// Visit a parse tree produced by RedshiftParser#createschemastmt.
	VisitCreateschemastmt(ctx *CreateschemastmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_auth_clause.
	VisitOpt_auth_clause(ctx *Opt_auth_clauseContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_quota.
	VisitOpt_quota(ctx *Opt_quotaContext) interface{}

	// Visit a parse tree produced by RedshiftParser#optschemaeltlist.
	VisitOptschemaeltlist(ctx *OptschemaeltlistContext) interface{}

	// Visit a parse tree produced by RedshiftParser#schema_stmt.
	VisitSchema_stmt(ctx *Schema_stmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#variablesetstmt.
	VisitVariablesetstmt(ctx *VariablesetstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#set_rest.
	VisitSet_rest(ctx *Set_restContext) interface{}

	// Visit a parse tree produced by RedshiftParser#generic_set.
	VisitGeneric_set(ctx *Generic_setContext) interface{}

	// Visit a parse tree produced by RedshiftParser#set_rest_more.
	VisitSet_rest_more(ctx *Set_rest_moreContext) interface{}

	// Visit a parse tree produced by RedshiftParser#var_name.
	VisitVar_name(ctx *Var_nameContext) interface{}

	// Visit a parse tree produced by RedshiftParser#var_list.
	VisitVar_list(ctx *Var_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#var_value.
	VisitVar_value(ctx *Var_valueContext) interface{}

	// Visit a parse tree produced by RedshiftParser#iso_level.
	VisitIso_level(ctx *Iso_levelContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_boolean_or_string.
	VisitOpt_boolean_or_string(ctx *Opt_boolean_or_stringContext) interface{}

	// Visit a parse tree produced by RedshiftParser#zone_value.
	VisitZone_value(ctx *Zone_valueContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_encoding.
	VisitOpt_encoding(ctx *Opt_encodingContext) interface{}

	// Visit a parse tree produced by RedshiftParser#nonreservedword_or_sconst.
	VisitNonreservedword_or_sconst(ctx *Nonreservedword_or_sconstContext) interface{}

	// Visit a parse tree produced by RedshiftParser#variableresetstmt.
	VisitVariableresetstmt(ctx *VariableresetstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#reset_rest.
	VisitReset_rest(ctx *Reset_restContext) interface{}

	// Visit a parse tree produced by RedshiftParser#generic_reset.
	VisitGeneric_reset(ctx *Generic_resetContext) interface{}

	// Visit a parse tree produced by RedshiftParser#setresetclause.
	VisitSetresetclause(ctx *SetresetclauseContext) interface{}

	// Visit a parse tree produced by RedshiftParser#functionsetresetclause.
	VisitFunctionsetresetclause(ctx *FunctionsetresetclauseContext) interface{}

	// Visit a parse tree produced by RedshiftParser#variableshowstmt.
	VisitVariableshowstmt(ctx *VariableshowstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#constraintssetstmt.
	VisitConstraintssetstmt(ctx *ConstraintssetstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#constraints_set_list.
	VisitConstraints_set_list(ctx *Constraints_set_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#constraints_set_mode.
	VisitConstraints_set_mode(ctx *Constraints_set_modeContext) interface{}

	// Visit a parse tree produced by RedshiftParser#checkpointstmt.
	VisitCheckpointstmt(ctx *CheckpointstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#discardstmt.
	VisitDiscardstmt(ctx *DiscardstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#altertablestmt.
	VisitAltertablestmt(ctx *AltertablestmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#alter_table_cmds.
	VisitAlter_table_cmds(ctx *Alter_table_cmdsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#table_constraint.
	VisitTable_constraint(ctx *Table_constraintContext) interface{}

	// Visit a parse tree produced by RedshiftParser#partition_cmd.
	VisitPartition_cmd(ctx *Partition_cmdContext) interface{}

	// Visit a parse tree produced by RedshiftParser#index_partition_cmd.
	VisitIndex_partition_cmd(ctx *Index_partition_cmdContext) interface{}

	// Visit a parse tree produced by RedshiftParser#alter_table_cmd.
	VisitAlter_table_cmd(ctx *Alter_table_cmdContext) interface{}

	// Visit a parse tree produced by RedshiftParser#alter_column_default.
	VisitAlter_column_default(ctx *Alter_column_defaultContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_drop_behavior.
	VisitOpt_drop_behavior(ctx *Opt_drop_behaviorContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_collate_clause.
	VisitOpt_collate_clause(ctx *Opt_collate_clauseContext) interface{}

	// Visit a parse tree produced by RedshiftParser#alter_using.
	VisitAlter_using(ctx *Alter_usingContext) interface{}

	// Visit a parse tree produced by RedshiftParser#replica_identity.
	VisitReplica_identity(ctx *Replica_identityContext) interface{}

	// Visit a parse tree produced by RedshiftParser#reloptions.
	VisitReloptions(ctx *ReloptionsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_reloptions.
	VisitOpt_reloptions(ctx *Opt_reloptionsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#reloption_list.
	VisitReloption_list(ctx *Reloption_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#reloption_elem.
	VisitReloption_elem(ctx *Reloption_elemContext) interface{}

	// Visit a parse tree produced by RedshiftParser#alter_identity_column_option_list.
	VisitAlter_identity_column_option_list(ctx *Alter_identity_column_option_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#alter_identity_column_option.
	VisitAlter_identity_column_option(ctx *Alter_identity_column_optionContext) interface{}

	// Visit a parse tree produced by RedshiftParser#partitionboundspec.
	VisitPartitionboundspec(ctx *PartitionboundspecContext) interface{}

	// Visit a parse tree produced by RedshiftParser#hash_partbound_elem.
	VisitHash_partbound_elem(ctx *Hash_partbound_elemContext) interface{}

	// Visit a parse tree produced by RedshiftParser#hash_partbound.
	VisitHash_partbound(ctx *Hash_partboundContext) interface{}

	// Visit a parse tree produced by RedshiftParser#altercompositetypestmt.
	VisitAltercompositetypestmt(ctx *AltercompositetypestmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#alter_type_cmds.
	VisitAlter_type_cmds(ctx *Alter_type_cmdsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#alter_type_cmd.
	VisitAlter_type_cmd(ctx *Alter_type_cmdContext) interface{}

	// Visit a parse tree produced by RedshiftParser#closeportalstmt.
	VisitCloseportalstmt(ctx *CloseportalstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#copystmt.
	VisitCopystmt(ctx *CopystmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#redshift_copy_authorization.
	VisitRedshift_copy_authorization(ctx *Redshift_copy_authorizationContext) interface{}

	// Visit a parse tree produced by RedshiftParser#redshift_copy_format.
	VisitRedshift_copy_format(ctx *Redshift_copy_formatContext) interface{}

	// Visit a parse tree produced by RedshiftParser#redshift_copy_parameter.
	VisitRedshift_copy_parameter(ctx *Redshift_copy_parameterContext) interface{}

	// Visit a parse tree produced by RedshiftParser#copy_param_name.
	VisitCopy_param_name(ctx *Copy_param_nameContext) interface{}

	// Visit a parse tree produced by RedshiftParser#copy_param_value.
	VisitCopy_param_value(ctx *Copy_param_valueContext) interface{}

	// Visit a parse tree produced by RedshiftParser#copy_from.
	VisitCopy_from(ctx *Copy_fromContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_program.
	VisitOpt_program(ctx *Opt_programContext) interface{}

	// Visit a parse tree produced by RedshiftParser#copy_file_name.
	VisitCopy_file_name(ctx *Copy_file_nameContext) interface{}

	// Visit a parse tree produced by RedshiftParser#copy_options.
	VisitCopy_options(ctx *Copy_optionsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#copy_opt_list.
	VisitCopy_opt_list(ctx *Copy_opt_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#copy_opt_item.
	VisitCopy_opt_item(ctx *Copy_opt_itemContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_binary.
	VisitOpt_binary(ctx *Opt_binaryContext) interface{}

	// Visit a parse tree produced by RedshiftParser#copy_delimiter.
	VisitCopy_delimiter(ctx *Copy_delimiterContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_using.
	VisitOpt_using(ctx *Opt_usingContext) interface{}

	// Visit a parse tree produced by RedshiftParser#copy_generic_opt_list.
	VisitCopy_generic_opt_list(ctx *Copy_generic_opt_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#copy_generic_opt_elem.
	VisitCopy_generic_opt_elem(ctx *Copy_generic_opt_elemContext) interface{}

	// Visit a parse tree produced by RedshiftParser#copy_generic_opt_arg.
	VisitCopy_generic_opt_arg(ctx *Copy_generic_opt_argContext) interface{}

	// Visit a parse tree produced by RedshiftParser#copy_generic_opt_arg_list.
	VisitCopy_generic_opt_arg_list(ctx *Copy_generic_opt_arg_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#copy_generic_opt_arg_list_item.
	VisitCopy_generic_opt_arg_list_item(ctx *Copy_generic_opt_arg_list_itemContext) interface{}

	// Visit a parse tree produced by RedshiftParser#createstmt.
	VisitCreatestmt(ctx *CreatestmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_table_attributes.
	VisitOpt_table_attributes(ctx *Opt_table_attributesContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opttemp.
	VisitOpttemp(ctx *OpttempContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opttableelementlist.
	VisitOpttableelementlist(ctx *OpttableelementlistContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opttypedtableelementlist.
	VisitOpttypedtableelementlist(ctx *OpttypedtableelementlistContext) interface{}

	// Visit a parse tree produced by RedshiftParser#tableelementlist.
	VisitTableelementlist(ctx *TableelementlistContext) interface{}

	// Visit a parse tree produced by RedshiftParser#typedtableelementlist.
	VisitTypedtableelementlist(ctx *TypedtableelementlistContext) interface{}

	// Visit a parse tree produced by RedshiftParser#tableelement.
	VisitTableelement(ctx *TableelementContext) interface{}

	// Visit a parse tree produced by RedshiftParser#typedtableelement.
	VisitTypedtableelement(ctx *TypedtableelementContext) interface{}

	// Visit a parse tree produced by RedshiftParser#columnDef.
	VisitColumnDef(ctx *ColumnDefContext) interface{}

	// Visit a parse tree produced by RedshiftParser#rs_colattributes.
	VisitRs_colattributes(ctx *Rs_colattributesContext) interface{}

	// Visit a parse tree produced by RedshiftParser#rs_colattribute.
	VisitRs_colattribute(ctx *Rs_colattributeContext) interface{}

	// Visit a parse tree produced by RedshiftParser#columnOptions.
	VisitColumnOptions(ctx *ColumnOptionsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#colquallist.
	VisitColquallist(ctx *ColquallistContext) interface{}

	// Visit a parse tree produced by RedshiftParser#colconstraint.
	VisitColconstraint(ctx *ColconstraintContext) interface{}

	// Visit a parse tree produced by RedshiftParser#colconstraintelem.
	VisitColconstraintelem(ctx *ColconstraintelemContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_unique_null_treatment.
	VisitOpt_unique_null_treatment(ctx *Opt_unique_null_treatmentContext) interface{}

	// Visit a parse tree produced by RedshiftParser#generated_when.
	VisitGenerated_when(ctx *Generated_whenContext) interface{}

	// Visit a parse tree produced by RedshiftParser#constraintattr.
	VisitConstraintattr(ctx *ConstraintattrContext) interface{}

	// Visit a parse tree produced by RedshiftParser#tablelikeclause.
	VisitTablelikeclause(ctx *TablelikeclauseContext) interface{}

	// Visit a parse tree produced by RedshiftParser#tablelikeoptionlist.
	VisitTablelikeoptionlist(ctx *TablelikeoptionlistContext) interface{}

	// Visit a parse tree produced by RedshiftParser#tablelikeoption.
	VisitTablelikeoption(ctx *TablelikeoptionContext) interface{}

	// Visit a parse tree produced by RedshiftParser#tableconstraint.
	VisitTableconstraint(ctx *TableconstraintContext) interface{}

	// Visit a parse tree produced by RedshiftParser#constraintelem.
	VisitConstraintelem(ctx *ConstraintelemContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_no_inherit.
	VisitOpt_no_inherit(ctx *Opt_no_inheritContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_column_list.
	VisitOpt_column_list(ctx *Opt_column_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#columnlist.
	VisitColumnlist(ctx *ColumnlistContext) interface{}

	// Visit a parse tree produced by RedshiftParser#columnElem.
	VisitColumnElem(ctx *ColumnElemContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_c_include.
	VisitOpt_c_include(ctx *Opt_c_includeContext) interface{}

	// Visit a parse tree produced by RedshiftParser#key_match.
	VisitKey_match(ctx *Key_matchContext) interface{}

	// Visit a parse tree produced by RedshiftParser#exclusionconstraintlist.
	VisitExclusionconstraintlist(ctx *ExclusionconstraintlistContext) interface{}

	// Visit a parse tree produced by RedshiftParser#exclusionconstraintelem.
	VisitExclusionconstraintelem(ctx *ExclusionconstraintelemContext) interface{}

	// Visit a parse tree produced by RedshiftParser#exclusionwhereclause.
	VisitExclusionwhereclause(ctx *ExclusionwhereclauseContext) interface{}

	// Visit a parse tree produced by RedshiftParser#key_actions.
	VisitKey_actions(ctx *Key_actionsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#key_update.
	VisitKey_update(ctx *Key_updateContext) interface{}

	// Visit a parse tree produced by RedshiftParser#key_delete.
	VisitKey_delete(ctx *Key_deleteContext) interface{}

	// Visit a parse tree produced by RedshiftParser#key_action.
	VisitKey_action(ctx *Key_actionContext) interface{}

	// Visit a parse tree produced by RedshiftParser#optinherit.
	VisitOptinherit(ctx *OptinheritContext) interface{}

	// Visit a parse tree produced by RedshiftParser#optpartitionspec.
	VisitOptpartitionspec(ctx *OptpartitionspecContext) interface{}

	// Visit a parse tree produced by RedshiftParser#partitionspec.
	VisitPartitionspec(ctx *PartitionspecContext) interface{}

	// Visit a parse tree produced by RedshiftParser#part_params.
	VisitPart_params(ctx *Part_paramsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#part_elem.
	VisitPart_elem(ctx *Part_elemContext) interface{}

	// Visit a parse tree produced by RedshiftParser#table_access_method_clause.
	VisitTable_access_method_clause(ctx *Table_access_method_clauseContext) interface{}

	// Visit a parse tree produced by RedshiftParser#optwith.
	VisitOptwith(ctx *OptwithContext) interface{}

	// Visit a parse tree produced by RedshiftParser#oncommitoption.
	VisitOncommitoption(ctx *OncommitoptionContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opttablespace.
	VisitOpttablespace(ctx *OpttablespaceContext) interface{}

	// Visit a parse tree produced by RedshiftParser#optredshifttableoptions.
	VisitOptredshifttableoptions(ctx *OptredshifttableoptionsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#redshifttableoption.
	VisitRedshifttableoption(ctx *RedshifttableoptionContext) interface{}

	// Visit a parse tree produced by RedshiftParser#sortkeyclause.
	VisitSortkeyclause(ctx *SortkeyclauseContext) interface{}

	// Visit a parse tree produced by RedshiftParser#sortkeyclausetype.
	VisitSortkeyclausetype(ctx *SortkeyclausetypeContext) interface{}

	// Visit a parse tree produced by RedshiftParser#optconstablespace.
	VisitOptconstablespace(ctx *OptconstablespaceContext) interface{}

	// Visit a parse tree produced by RedshiftParser#existingindex.
	VisitExistingindex(ctx *ExistingindexContext) interface{}

	// Visit a parse tree produced by RedshiftParser#createstatsstmt.
	VisitCreatestatsstmt(ctx *CreatestatsstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#alterstatsstmt.
	VisitAlterstatsstmt(ctx *AlterstatsstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#createasstmt.
	VisitCreateasstmt(ctx *CreateasstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#create_as_target.
	VisitCreate_as_target(ctx *Create_as_targetContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_backup_clause_table_attributes.
	VisitOpt_backup_clause_table_attributes(ctx *Opt_backup_clause_table_attributesContext) interface{}

	// Visit a parse tree produced by RedshiftParser#table_attributes.
	VisitTable_attributes(ctx *Table_attributesContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_backup_clause.
	VisitOpt_backup_clause(ctx *Opt_backup_clauseContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_with_data.
	VisitOpt_with_data(ctx *Opt_with_dataContext) interface{}

	// Visit a parse tree produced by RedshiftParser#creatematviewstmt.
	VisitCreatematviewstmt(ctx *CreatematviewstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_auto_refresh.
	VisitOpt_auto_refresh(ctx *Opt_auto_refreshContext) interface{}

	// Visit a parse tree produced by RedshiftParser#refreshmatviewstmt.
	VisitRefreshmatviewstmt(ctx *RefreshmatviewstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#createseqstmt.
	VisitCreateseqstmt(ctx *CreateseqstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#alterseqstmt.
	VisitAlterseqstmt(ctx *AlterseqstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#optseqoptlist.
	VisitOptseqoptlist(ctx *OptseqoptlistContext) interface{}

	// Visit a parse tree produced by RedshiftParser#optparenthesizedseqoptlist.
	VisitOptparenthesizedseqoptlist(ctx *OptparenthesizedseqoptlistContext) interface{}

	// Visit a parse tree produced by RedshiftParser#seqoptlist.
	VisitSeqoptlist(ctx *SeqoptlistContext) interface{}

	// Visit a parse tree produced by RedshiftParser#seqoptelem.
	VisitSeqoptelem(ctx *SeqoptelemContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_by.
	VisitOpt_by(ctx *Opt_byContext) interface{}

	// Visit a parse tree produced by RedshiftParser#numericonly.
	VisitNumericonly(ctx *NumericonlyContext) interface{}

	// Visit a parse tree produced by RedshiftParser#numericonly_list.
	VisitNumericonly_list(ctx *Numericonly_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#createplangstmt.
	VisitCreateplangstmt(ctx *CreateplangstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_trusted.
	VisitOpt_trusted(ctx *Opt_trustedContext) interface{}

	// Visit a parse tree produced by RedshiftParser#handler_name.
	VisitHandler_name(ctx *Handler_nameContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_inline_handler.
	VisitOpt_inline_handler(ctx *Opt_inline_handlerContext) interface{}

	// Visit a parse tree produced by RedshiftParser#validator_clause.
	VisitValidator_clause(ctx *Validator_clauseContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_validator.
	VisitOpt_validator(ctx *Opt_validatorContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_procedural.
	VisitOpt_procedural(ctx *Opt_proceduralContext) interface{}

	// Visit a parse tree produced by RedshiftParser#createtablespacestmt.
	VisitCreatetablespacestmt(ctx *CreatetablespacestmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opttablespaceowner.
	VisitOpttablespaceowner(ctx *OpttablespaceownerContext) interface{}

	// Visit a parse tree produced by RedshiftParser#droptablespacestmt.
	VisitDroptablespacestmt(ctx *DroptablespacestmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#createextensionstmt.
	VisitCreateextensionstmt(ctx *CreateextensionstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#create_extension_opt_list.
	VisitCreate_extension_opt_list(ctx *Create_extension_opt_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#create_extension_opt_item.
	VisitCreate_extension_opt_item(ctx *Create_extension_opt_itemContext) interface{}

	// Visit a parse tree produced by RedshiftParser#alterextensionstmt.
	VisitAlterextensionstmt(ctx *AlterextensionstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#alter_extension_opt_list.
	VisitAlter_extension_opt_list(ctx *Alter_extension_opt_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#alter_extension_opt_item.
	VisitAlter_extension_opt_item(ctx *Alter_extension_opt_itemContext) interface{}

	// Visit a parse tree produced by RedshiftParser#alterextensioncontentsstmt.
	VisitAlterextensioncontentsstmt(ctx *AlterextensioncontentsstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#createfdwstmt.
	VisitCreatefdwstmt(ctx *CreatefdwstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#fdw_option.
	VisitFdw_option(ctx *Fdw_optionContext) interface{}

	// Visit a parse tree produced by RedshiftParser#fdw_options.
	VisitFdw_options(ctx *Fdw_optionsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_fdw_options.
	VisitOpt_fdw_options(ctx *Opt_fdw_optionsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#alterfdwstmt.
	VisitAlterfdwstmt(ctx *AlterfdwstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#create_generic_options.
	VisitCreate_generic_options(ctx *Create_generic_optionsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#generic_option_list.
	VisitGeneric_option_list(ctx *Generic_option_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#alter_generic_options.
	VisitAlter_generic_options(ctx *Alter_generic_optionsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#alter_generic_option_list.
	VisitAlter_generic_option_list(ctx *Alter_generic_option_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#alter_generic_option_elem.
	VisitAlter_generic_option_elem(ctx *Alter_generic_option_elemContext) interface{}

	// Visit a parse tree produced by RedshiftParser#generic_option_elem.
	VisitGeneric_option_elem(ctx *Generic_option_elemContext) interface{}

	// Visit a parse tree produced by RedshiftParser#generic_option_name.
	VisitGeneric_option_name(ctx *Generic_option_nameContext) interface{}

	// Visit a parse tree produced by RedshiftParser#generic_option_arg.
	VisitGeneric_option_arg(ctx *Generic_option_argContext) interface{}

	// Visit a parse tree produced by RedshiftParser#createforeignserverstmt.
	VisitCreateforeignserverstmt(ctx *CreateforeignserverstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_type.
	VisitOpt_type(ctx *Opt_typeContext) interface{}

	// Visit a parse tree produced by RedshiftParser#foreign_server_version.
	VisitForeign_server_version(ctx *Foreign_server_versionContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_foreign_server_version.
	VisitOpt_foreign_server_version(ctx *Opt_foreign_server_versionContext) interface{}

	// Visit a parse tree produced by RedshiftParser#alterforeignserverstmt.
	VisitAlterforeignserverstmt(ctx *AlterforeignserverstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#createforeigntablestmt.
	VisitCreateforeigntablestmt(ctx *CreateforeigntablestmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#importforeignschemastmt.
	VisitImportforeignschemastmt(ctx *ImportforeignschemastmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#import_qualification_type.
	VisitImport_qualification_type(ctx *Import_qualification_typeContext) interface{}

	// Visit a parse tree produced by RedshiftParser#import_qualification.
	VisitImport_qualification(ctx *Import_qualificationContext) interface{}

	// Visit a parse tree produced by RedshiftParser#createusermappingstmt.
	VisitCreateusermappingstmt(ctx *CreateusermappingstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#auth_ident.
	VisitAuth_ident(ctx *Auth_identContext) interface{}

	// Visit a parse tree produced by RedshiftParser#dropusermappingstmt.
	VisitDropusermappingstmt(ctx *DropusermappingstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#alterusermappingstmt.
	VisitAlterusermappingstmt(ctx *AlterusermappingstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#createpolicystmt.
	VisitCreatepolicystmt(ctx *CreatepolicystmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#alterpolicystmt.
	VisitAlterpolicystmt(ctx *AlterpolicystmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#rowsecurityoptionalexpr.
	VisitRowsecurityoptionalexpr(ctx *RowsecurityoptionalexprContext) interface{}

	// Visit a parse tree produced by RedshiftParser#rowsecurityoptionalwithcheck.
	VisitRowsecurityoptionalwithcheck(ctx *RowsecurityoptionalwithcheckContext) interface{}

	// Visit a parse tree produced by RedshiftParser#rowsecuritydefaulttorole.
	VisitRowsecuritydefaulttorole(ctx *RowsecuritydefaulttoroleContext) interface{}

	// Visit a parse tree produced by RedshiftParser#rowsecurityoptionaltorole.
	VisitRowsecurityoptionaltorole(ctx *RowsecurityoptionaltoroleContext) interface{}

	// Visit a parse tree produced by RedshiftParser#rowsecuritydefaultpermissive.
	VisitRowsecuritydefaultpermissive(ctx *RowsecuritydefaultpermissiveContext) interface{}

	// Visit a parse tree produced by RedshiftParser#rowsecuritydefaultforcmd.
	VisitRowsecuritydefaultforcmd(ctx *RowsecuritydefaultforcmdContext) interface{}

	// Visit a parse tree produced by RedshiftParser#row_security_cmd.
	VisitRow_security_cmd(ctx *Row_security_cmdContext) interface{}

	// Visit a parse tree produced by RedshiftParser#createamstmt.
	VisitCreateamstmt(ctx *CreateamstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#am_type.
	VisitAm_type(ctx *Am_typeContext) interface{}

	// Visit a parse tree produced by RedshiftParser#createtrigstmt.
	VisitCreatetrigstmt(ctx *CreatetrigstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#triggeractiontime.
	VisitTriggeractiontime(ctx *TriggeractiontimeContext) interface{}

	// Visit a parse tree produced by RedshiftParser#triggerevents.
	VisitTriggerevents(ctx *TriggereventsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#triggeroneevent.
	VisitTriggeroneevent(ctx *TriggeroneeventContext) interface{}

	// Visit a parse tree produced by RedshiftParser#triggerreferencing.
	VisitTriggerreferencing(ctx *TriggerreferencingContext) interface{}

	// Visit a parse tree produced by RedshiftParser#triggertransitions.
	VisitTriggertransitions(ctx *TriggertransitionsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#triggertransition.
	VisitTriggertransition(ctx *TriggertransitionContext) interface{}

	// Visit a parse tree produced by RedshiftParser#transitionoldornew.
	VisitTransitionoldornew(ctx *TransitionoldornewContext) interface{}

	// Visit a parse tree produced by RedshiftParser#transitionrowortable.
	VisitTransitionrowortable(ctx *TransitionrowortableContext) interface{}

	// Visit a parse tree produced by RedshiftParser#transitionrelname.
	VisitTransitionrelname(ctx *TransitionrelnameContext) interface{}

	// Visit a parse tree produced by RedshiftParser#triggerforspec.
	VisitTriggerforspec(ctx *TriggerforspecContext) interface{}

	// Visit a parse tree produced by RedshiftParser#triggerforopteach.
	VisitTriggerforopteach(ctx *TriggerforopteachContext) interface{}

	// Visit a parse tree produced by RedshiftParser#triggerfortype.
	VisitTriggerfortype(ctx *TriggerfortypeContext) interface{}

	// Visit a parse tree produced by RedshiftParser#triggerwhen.
	VisitTriggerwhen(ctx *TriggerwhenContext) interface{}

	// Visit a parse tree produced by RedshiftParser#function_or_procedure.
	VisitFunction_or_procedure(ctx *Function_or_procedureContext) interface{}

	// Visit a parse tree produced by RedshiftParser#triggerfuncargs.
	VisitTriggerfuncargs(ctx *TriggerfuncargsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#triggerfuncarg.
	VisitTriggerfuncarg(ctx *TriggerfuncargContext) interface{}

	// Visit a parse tree produced by RedshiftParser#optconstrfromtable.
	VisitOptconstrfromtable(ctx *OptconstrfromtableContext) interface{}

	// Visit a parse tree produced by RedshiftParser#constraintattributespec.
	VisitConstraintattributespec(ctx *ConstraintattributespecContext) interface{}

	// Visit a parse tree produced by RedshiftParser#constraintattributeElem.
	VisitConstraintattributeElem(ctx *ConstraintattributeElemContext) interface{}

	// Visit a parse tree produced by RedshiftParser#createeventtrigstmt.
	VisitCreateeventtrigstmt(ctx *CreateeventtrigstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#event_trigger_when_list.
	VisitEvent_trigger_when_list(ctx *Event_trigger_when_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#event_trigger_when_item.
	VisitEvent_trigger_when_item(ctx *Event_trigger_when_itemContext) interface{}

	// Visit a parse tree produced by RedshiftParser#event_trigger_value_list.
	VisitEvent_trigger_value_list(ctx *Event_trigger_value_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#altereventtrigstmt.
	VisitAltereventtrigstmt(ctx *AltereventtrigstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#enable_trigger.
	VisitEnable_trigger(ctx *Enable_triggerContext) interface{}

	// Visit a parse tree produced by RedshiftParser#createassertionstmt.
	VisitCreateassertionstmt(ctx *CreateassertionstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#definestmt.
	VisitDefinestmt(ctx *DefinestmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#definition.
	VisitDefinition(ctx *DefinitionContext) interface{}

	// Visit a parse tree produced by RedshiftParser#def_list.
	VisitDef_list(ctx *Def_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#def_elem.
	VisitDef_elem(ctx *Def_elemContext) interface{}

	// Visit a parse tree produced by RedshiftParser#def_arg.
	VisitDef_arg(ctx *Def_argContext) interface{}

	// Visit a parse tree produced by RedshiftParser#old_aggr_definition.
	VisitOld_aggr_definition(ctx *Old_aggr_definitionContext) interface{}

	// Visit a parse tree produced by RedshiftParser#old_aggr_list.
	VisitOld_aggr_list(ctx *Old_aggr_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#old_aggr_elem.
	VisitOld_aggr_elem(ctx *Old_aggr_elemContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_enum_val_list.
	VisitOpt_enum_val_list(ctx *Opt_enum_val_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#enum_val_list.
	VisitEnum_val_list(ctx *Enum_val_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#alterenumstmt.
	VisitAlterenumstmt(ctx *AlterenumstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_if_not_exists.
	VisitOpt_if_not_exists(ctx *Opt_if_not_existsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#createopclassstmt.
	VisitCreateopclassstmt(ctx *CreateopclassstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opclass_item_list.
	VisitOpclass_item_list(ctx *Opclass_item_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opclass_item.
	VisitOpclass_item(ctx *Opclass_itemContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_default.
	VisitOpt_default(ctx *Opt_defaultContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_opfamily.
	VisitOpt_opfamily(ctx *Opt_opfamilyContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opclass_purpose.
	VisitOpclass_purpose(ctx *Opclass_purposeContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_recheck.
	VisitOpt_recheck(ctx *Opt_recheckContext) interface{}

	// Visit a parse tree produced by RedshiftParser#createopfamilystmt.
	VisitCreateopfamilystmt(ctx *CreateopfamilystmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#alteropfamilystmt.
	VisitAlteropfamilystmt(ctx *AlteropfamilystmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opclass_drop_list.
	VisitOpclass_drop_list(ctx *Opclass_drop_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opclass_drop.
	VisitOpclass_drop(ctx *Opclass_dropContext) interface{}

	// Visit a parse tree produced by RedshiftParser#dropopclassstmt.
	VisitDropopclassstmt(ctx *DropopclassstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#dropopfamilystmt.
	VisitDropopfamilystmt(ctx *DropopfamilystmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#dropownedstmt.
	VisitDropownedstmt(ctx *DropownedstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#reassignownedstmt.
	VisitReassignownedstmt(ctx *ReassignownedstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#dropstmt.
	VisitDropstmt(ctx *DropstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#object_type_any_name.
	VisitObject_type_any_name(ctx *Object_type_any_nameContext) interface{}

	// Visit a parse tree produced by RedshiftParser#object_type_name.
	VisitObject_type_name(ctx *Object_type_nameContext) interface{}

	// Visit a parse tree produced by RedshiftParser#drop_type_name.
	VisitDrop_type_name(ctx *Drop_type_nameContext) interface{}

	// Visit a parse tree produced by RedshiftParser#object_type_name_on_any_name.
	VisitObject_type_name_on_any_name(ctx *Object_type_name_on_any_nameContext) interface{}

	// Visit a parse tree produced by RedshiftParser#any_name_list.
	VisitAny_name_list(ctx *Any_name_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#any_name.
	VisitAny_name(ctx *Any_nameContext) interface{}

	// Visit a parse tree produced by RedshiftParser#attrs.
	VisitAttrs(ctx *AttrsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#type_name_list.
	VisitType_name_list(ctx *Type_name_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#truncatestmt.
	VisitTruncatestmt(ctx *TruncatestmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_restart_seqs.
	VisitOpt_restart_seqs(ctx *Opt_restart_seqsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#commentstmt.
	VisitCommentstmt(ctx *CommentstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#comment_text.
	VisitComment_text(ctx *Comment_textContext) interface{}

	// Visit a parse tree produced by RedshiftParser#seclabelstmt.
	VisitSeclabelstmt(ctx *SeclabelstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_provider.
	VisitOpt_provider(ctx *Opt_providerContext) interface{}

	// Visit a parse tree produced by RedshiftParser#security_label.
	VisitSecurity_label(ctx *Security_labelContext) interface{}

	// Visit a parse tree produced by RedshiftParser#fetchstmt.
	VisitFetchstmt(ctx *FetchstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#fetch_args.
	VisitFetch_args(ctx *Fetch_argsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#grantstmt.
	VisitGrantstmt(ctx *GrantstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#grant_permissions_for_rls_lookup_tables.
	VisitGrant_permissions_for_rls_lookup_tables(ctx *Grant_permissions_for_rls_lookup_tablesContext) interface{}

	// Visit a parse tree produced by RedshiftParser#grant_explain_permissions_for_row_level_security_policy_filters.
	VisitGrant_explain_permissions_for_row_level_security_policy_filters(ctx *Grant_explain_permissions_for_row_level_security_policy_filtersContext) interface{}

	// Visit a parse tree produced by RedshiftParser#grant_machine_learning_permissions.
	VisitGrant_machine_learning_permissions(ctx *Grant_machine_learning_permissionsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#grant_role_permissions.
	VisitGrant_role_permissions(ctx *Grant_role_permissionsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#grant_role_permission_target_list.
	VisitGrant_role_permission_target_list(ctx *Grant_role_permission_target_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#grant_role_permission_target_list_item.
	VisitGrant_role_permission_target_list_item(ctx *Grant_role_permission_target_list_itemContext) interface{}

	// Visit a parse tree produced by RedshiftParser#system_permissions.
	VisitSystem_permissions(ctx *System_permissionsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#system_permissions_item.
	VisitSystem_permissions_item(ctx *System_permissions_itemContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_with_admin_option.
	VisitOpt_with_admin_option(ctx *Opt_with_admin_optionContext) interface{}

	// Visit a parse tree produced by RedshiftParser#grant_scoped_permissions.
	VisitGrant_scoped_permissions(ctx *Grant_scoped_permissionsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#grant_scoped_schemas_permissions.
	VisitGrant_scoped_schemas_permissions(ctx *Grant_scoped_schemas_permissionsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#grant_scoped_tables_permissions.
	VisitGrant_scoped_tables_permissions(ctx *Grant_scoped_tables_permissionsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#grant_scoped_functions_permissions.
	VisitGrant_scoped_functions_permissions(ctx *Grant_scoped_functions_permissionsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#grant_scoped_procedures_permissions.
	VisitGrant_scoped_procedures_permissions(ctx *Grant_scoped_procedures_permissionsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#grant_scoped_languages_permissions.
	VisitGrant_scoped_languages_permissions(ctx *Grant_scoped_languages_permissionsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#grant_scoped_copy_jobs_permissions.
	VisitGrant_scoped_copy_jobs_permissions(ctx *Grant_scoped_copy_jobs_permissionsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#grantee_list_without_public.
	VisitGrantee_list_without_public(ctx *Grantee_list_without_publicContext) interface{}

	// Visit a parse tree produced by RedshiftParser#grantee_without_public.
	VisitGrantee_without_public(ctx *Grantee_without_publicContext) interface{}

	// Visit a parse tree produced by RedshiftParser#grant_datashare_permissions.
	VisitGrant_datashare_permissions(ctx *Grant_datashare_permissionsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#grant_spectrum_integration_permissions.
	VisitGrant_spectrum_integration_permissions(ctx *Grant_spectrum_integration_permissionsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#grant_spectrum_integration_external_schema_permissions.
	VisitGrant_spectrum_integration_external_schema_permissions(ctx *Grant_spectrum_integration_external_schema_permissionsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#spectrum_integration_external_schema_permission_list.
	VisitSpectrum_integration_external_schema_permission_list(ctx *Spectrum_integration_external_schema_permission_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#spectrum_integration_external_schema_permission.
	VisitSpectrum_integration_external_schema_permission(ctx *Spectrum_integration_external_schema_permissionContext) interface{}

	// Visit a parse tree produced by RedshiftParser#grant_spectrum_integration_external_table_permissions.
	VisitGrant_spectrum_integration_external_table_permissions(ctx *Grant_spectrum_integration_external_table_permissionsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#spectrum_integration_external_table_permission.
	VisitSpectrum_integration_external_table_permission(ctx *Spectrum_integration_external_table_permissionContext) interface{}

	// Visit a parse tree produced by RedshiftParser#spectrum_integration_external_table_permission_list.
	VisitSpectrum_integration_external_table_permission_list(ctx *Spectrum_integration_external_table_permission_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#grant_spectrum_integration_extenral_column_permissions.
	VisitGrant_spectrum_integration_extenral_column_permissions(ctx *Grant_spectrum_integration_extenral_column_permissionsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#iamrolelist_or_public.
	VisitIamrolelist_or_public(ctx *Iamrolelist_or_publicContext) interface{}

	// Visit a parse tree produced by RedshiftParser#iamrolelist.
	VisitIamrolelist(ctx *IamrolelistContext) interface{}

	// Visit a parse tree produced by RedshiftParser#grant_assume_role_permissions.
	VisitGrant_assume_role_permissions(ctx *Grant_assume_role_permissionsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#grant_assume_role_for_list.
	VisitGrant_assume_role_for_list(ctx *Grant_assume_role_for_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#grant_assume_role_for_item.
	VisitGrant_assume_role_for_item(ctx *Grant_assume_role_for_itemContext) interface{}

	// Visit a parse tree produced by RedshiftParser#grant_assume_role_target.
	VisitGrant_assume_role_target(ctx *Grant_assume_role_targetContext) interface{}

	// Visit a parse tree produced by RedshiftParser#grant_column_level_permissions.
	VisitGrant_column_level_permissions(ctx *Grant_column_level_permissionsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#column_privilege_target.
	VisitColumn_privilege_target(ctx *Column_privilege_targetContext) interface{}

	// Visit a parse tree produced by RedshiftParser#column_privilege_list.
	VisitColumn_privilege_list(ctx *Column_privilege_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#column_all_privilege.
	VisitColumn_all_privilege(ctx *Column_all_privilegeContext) interface{}

	// Visit a parse tree produced by RedshiftParser#column_select_update_privilege.
	VisitColumn_select_update_privilege(ctx *Column_select_update_privilegeContext) interface{}

	// Visit a parse tree produced by RedshiftParser#common_grant.
	VisitCommon_grant(ctx *Common_grantContext) interface{}

	// Visit a parse tree produced by RedshiftParser#copy_job_privilege_list.
	VisitCopy_job_privilege_list(ctx *Copy_job_privilege_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#copy_job_privilege.
	VisitCopy_job_privilege(ctx *Copy_job_privilegeContext) interface{}

	// Visit a parse tree produced by RedshiftParser#copy_job_target.
	VisitCopy_job_target(ctx *Copy_job_targetContext) interface{}

	// Visit a parse tree produced by RedshiftParser#copy_job_name.
	VisitCopy_job_name(ctx *Copy_job_nameContext) interface{}

	// Visit a parse tree produced by RedshiftParser#language_privilege_list.
	VisitLanguage_privilege_list(ctx *Language_privilege_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#grant_language_target.
	VisitGrant_language_target(ctx *Grant_language_targetContext) interface{}

	// Visit a parse tree produced by RedshiftParser#grant_procedure_target.
	VisitGrant_procedure_target(ctx *Grant_procedure_targetContext) interface{}

	// Visit a parse tree produced by RedshiftParser#procedure_privilege_list.
	VisitProcedure_privilege_list(ctx *Procedure_privilege_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#procedure_privilege.
	VisitProcedure_privilege(ctx *Procedure_privilegeContext) interface{}

	// Visit a parse tree produced by RedshiftParser#function_privilege_list.
	VisitFunction_privilege_list(ctx *Function_privilege_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#function_privilege.
	VisitFunction_privilege(ctx *Function_privilegeContext) interface{}

	// Visit a parse tree produced by RedshiftParser#grant_function_target.
	VisitGrant_function_target(ctx *Grant_function_targetContext) interface{}

	// Visit a parse tree produced by RedshiftParser#grant_schema_target.
	VisitGrant_schema_target(ctx *Grant_schema_targetContext) interface{}

	// Visit a parse tree produced by RedshiftParser#revoke_schema_target.
	VisitRevoke_schema_target(ctx *Revoke_schema_targetContext) interface{}

	// Visit a parse tree produced by RedshiftParser#schema_privilege_list.
	VisitSchema_privilege_list(ctx *Schema_privilege_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#schema_privilege.
	VisitSchema_privilege(ctx *Schema_privilegeContext) interface{}

	// Visit a parse tree produced by RedshiftParser#database_privilege_list.
	VisitDatabase_privilege_list(ctx *Database_privilege_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#database_privilege.
	VisitDatabase_privilege(ctx *Database_privilegeContext) interface{}

	// Visit a parse tree produced by RedshiftParser#grant_database_target.
	VisitGrant_database_target(ctx *Grant_database_targetContext) interface{}

	// Visit a parse tree produced by RedshiftParser#grant_table_target.
	VisitGrant_table_target(ctx *Grant_table_targetContext) interface{}

	// Visit a parse tree produced by RedshiftParser#revoke_table_target.
	VisitRevoke_table_target(ctx *Revoke_table_targetContext) interface{}

	// Visit a parse tree produced by RedshiftParser#all_tables_in_schema_list.
	VisitAll_tables_in_schema_list(ctx *All_tables_in_schema_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#all_privileges.
	VisitAll_privileges(ctx *All_privilegesContext) interface{}

	// Visit a parse tree produced by RedshiftParser#grantee_list.
	VisitGrantee_list(ctx *Grantee_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#grantee.
	VisitGrantee(ctx *GranteeContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_with_grant_option.
	VisitOpt_with_grant_option(ctx *Opt_with_grant_optionContext) interface{}

	// Visit a parse tree produced by RedshiftParser#table_privilege.
	VisitTable_privilege(ctx *Table_privilegeContext) interface{}

	// Visit a parse tree produced by RedshiftParser#table_privilege_list.
	VisitTable_privilege_list(ctx *Table_privilege_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#revokestmt.
	VisitRevokestmt(ctx *RevokestmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#revoke_permissions_for_rls_lookup_tables.
	VisitRevoke_permissions_for_rls_lookup_tables(ctx *Revoke_permissions_for_rls_lookup_tablesContext) interface{}

	// Visit a parse tree produced by RedshiftParser#revoke_explain_permissions_for_row_level_security_policy_filters.
	VisitRevoke_explain_permissions_for_row_level_security_policy_filters(ctx *Revoke_explain_permissions_for_row_level_security_policy_filtersContext) interface{}

	// Visit a parse tree produced by RedshiftParser#revoke_machine_learning_permissions.
	VisitRevoke_machine_learning_permissions(ctx *Revoke_machine_learning_permissionsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#revoke_role_permissions.
	VisitRevoke_role_permissions(ctx *Revoke_role_permissionsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#revoke_scoped_permissions.
	VisitRevoke_scoped_permissions(ctx *Revoke_scoped_permissionsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#revoke_scoped_schemas_permissions.
	VisitRevoke_scoped_schemas_permissions(ctx *Revoke_scoped_schemas_permissionsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#revoke_scoped_tables_permissions.
	VisitRevoke_scoped_tables_permissions(ctx *Revoke_scoped_tables_permissionsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#revoke_scoped_functions_permissions.
	VisitRevoke_scoped_functions_permissions(ctx *Revoke_scoped_functions_permissionsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#revoke_scoped_procedures_permissions.
	VisitRevoke_scoped_procedures_permissions(ctx *Revoke_scoped_procedures_permissionsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#revoke_scoped_languages_permissions.
	VisitRevoke_scoped_languages_permissions(ctx *Revoke_scoped_languages_permissionsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#revoke_scoped_copy_jobs_permissions.
	VisitRevoke_scoped_copy_jobs_permissions(ctx *Revoke_scoped_copy_jobs_permissionsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#revoke_datashare_permissions.
	VisitRevoke_datashare_permissions(ctx *Revoke_datashare_permissionsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#revoke_spectrum_integration_permissions.
	VisitRevoke_spectrum_integration_permissions(ctx *Revoke_spectrum_integration_permissionsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#revoke_spectrum_integration_external_schema_permissions.
	VisitRevoke_spectrum_integration_external_schema_permissions(ctx *Revoke_spectrum_integration_external_schema_permissionsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#revoke_spectrum_integration_external_table_permissions.
	VisitRevoke_spectrum_integration_external_table_permissions(ctx *Revoke_spectrum_integration_external_table_permissionsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#revoke_spectrum_integration_extenral_column_permissions.
	VisitRevoke_spectrum_integration_extenral_column_permissions(ctx *Revoke_spectrum_integration_extenral_column_permissionsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#revoke_assume_role_permissions.
	VisitRevoke_assume_role_permissions(ctx *Revoke_assume_role_permissionsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#revoke_column_level_permissions.
	VisitRevoke_column_level_permissions(ctx *Revoke_column_level_permissionsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#common_revoke.
	VisitCommon_revoke(ctx *Common_revokeContext) interface{}

	// Visit a parse tree produced by RedshiftParser#privileges.
	VisitPrivileges(ctx *PrivilegesContext) interface{}

	// Visit a parse tree produced by RedshiftParser#privilege_list.
	VisitPrivilege_list(ctx *Privilege_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#privilege.
	VisitPrivilege(ctx *PrivilegeContext) interface{}

	// Visit a parse tree produced by RedshiftParser#privilege_target.
	VisitPrivilege_target(ctx *Privilege_targetContext) interface{}

	// Visit a parse tree produced by RedshiftParser#parameter_name_list.
	VisitParameter_name_list(ctx *Parameter_name_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#parameter_name.
	VisitParameter_name(ctx *Parameter_nameContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_grant_grant_option.
	VisitOpt_grant_grant_option(ctx *Opt_grant_grant_optionContext) interface{}

	// Visit a parse tree produced by RedshiftParser#grantrolestmt.
	VisitGrantrolestmt(ctx *GrantrolestmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#revokerolestmt.
	VisitRevokerolestmt(ctx *RevokerolestmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_grant_admin_option.
	VisitOpt_grant_admin_option(ctx *Opt_grant_admin_optionContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_granted_by.
	VisitOpt_granted_by(ctx *Opt_granted_byContext) interface{}

	// Visit a parse tree produced by RedshiftParser#alterdefaultprivilegesstmt.
	VisitAlterdefaultprivilegesstmt(ctx *AlterdefaultprivilegesstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#defacloptionlist.
	VisitDefacloptionlist(ctx *DefacloptionlistContext) interface{}

	// Visit a parse tree produced by RedshiftParser#defacloption.
	VisitDefacloption(ctx *DefacloptionContext) interface{}

	// Visit a parse tree produced by RedshiftParser#defaclaction.
	VisitDefaclaction(ctx *DefaclactionContext) interface{}

	// Visit a parse tree produced by RedshiftParser#defacl_privilege_target.
	VisitDefacl_privilege_target(ctx *Defacl_privilege_targetContext) interface{}

	// Visit a parse tree produced by RedshiftParser#indexstmt.
	VisitIndexstmt(ctx *IndexstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_unique.
	VisitOpt_unique(ctx *Opt_uniqueContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_concurrently.
	VisitOpt_concurrently(ctx *Opt_concurrentlyContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_index_name.
	VisitOpt_index_name(ctx *Opt_index_nameContext) interface{}

	// Visit a parse tree produced by RedshiftParser#access_method_clause.
	VisitAccess_method_clause(ctx *Access_method_clauseContext) interface{}

	// Visit a parse tree produced by RedshiftParser#index_params.
	VisitIndex_params(ctx *Index_paramsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#index_elem_options.
	VisitIndex_elem_options(ctx *Index_elem_optionsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#index_elem.
	VisitIndex_elem(ctx *Index_elemContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_include.
	VisitOpt_include(ctx *Opt_includeContext) interface{}

	// Visit a parse tree produced by RedshiftParser#index_including_params.
	VisitIndex_including_params(ctx *Index_including_paramsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_collate.
	VisitOpt_collate(ctx *Opt_collateContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_class.
	VisitOpt_class(ctx *Opt_classContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_asc_desc.
	VisitOpt_asc_desc(ctx *Opt_asc_descContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_nulls_order.
	VisitOpt_nulls_order(ctx *Opt_nulls_orderContext) interface{}

	// Visit a parse tree produced by RedshiftParser#createfunctionstmt.
	VisitCreatefunctionstmt(ctx *CreatefunctionstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#createprocedurestmt.
	VisitCreateprocedurestmt(ctx *CreateprocedurestmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_nonatomic.
	VisitOpt_nonatomic(ctx *Opt_nonatomicContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_or_replace.
	VisitOpt_or_replace(ctx *Opt_or_replaceContext) interface{}

	// Visit a parse tree produced by RedshiftParser#func_py_args_or_sql_args.
	VisitFunc_py_args_or_sql_args(ctx *Func_py_args_or_sql_argsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#func_py_args_or_sql_args_list.
	VisitFunc_py_args_or_sql_args_list(ctx *Func_py_args_or_sql_args_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#func_args.
	VisitFunc_args(ctx *Func_argsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#func_args_list.
	VisitFunc_args_list(ctx *Func_args_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#function_with_argtypes_list.
	VisitFunction_with_argtypes_list(ctx *Function_with_argtypes_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#function_with_argtypes.
	VisitFunction_with_argtypes(ctx *Function_with_argtypesContext) interface{}

	// Visit a parse tree produced by RedshiftParser#func_args_with_defaults.
	VisitFunc_args_with_defaults(ctx *Func_args_with_defaultsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#func_args_with_defaults_list.
	VisitFunc_args_with_defaults_list(ctx *Func_args_with_defaults_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#func_arg.
	VisitFunc_arg(ctx *Func_argContext) interface{}

	// Visit a parse tree produced by RedshiftParser#arg_class.
	VisitArg_class(ctx *Arg_classContext) interface{}

	// Visit a parse tree produced by RedshiftParser#param_name.
	VisitParam_name(ctx *Param_nameContext) interface{}

	// Visit a parse tree produced by RedshiftParser#func_return.
	VisitFunc_return(ctx *Func_returnContext) interface{}

	// Visit a parse tree produced by RedshiftParser#func_type.
	VisitFunc_type(ctx *Func_typeContext) interface{}

	// Visit a parse tree produced by RedshiftParser#func_arg_with_default.
	VisitFunc_arg_with_default(ctx *Func_arg_with_defaultContext) interface{}

	// Visit a parse tree produced by RedshiftParser#aggr_arg.
	VisitAggr_arg(ctx *Aggr_argContext) interface{}

	// Visit a parse tree produced by RedshiftParser#aggr_args.
	VisitAggr_args(ctx *Aggr_argsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#aggr_args_list.
	VisitAggr_args_list(ctx *Aggr_args_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#aggregate_with_argtypes.
	VisitAggregate_with_argtypes(ctx *Aggregate_with_argtypesContext) interface{}

	// Visit a parse tree produced by RedshiftParser#aggregate_with_argtypes_list.
	VisitAggregate_with_argtypes_list(ctx *Aggregate_with_argtypes_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#createfunc_opt_list.
	VisitCreatefunc_opt_list(ctx *Createfunc_opt_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#common_func_opt_item.
	VisitCommon_func_opt_item(ctx *Common_func_opt_itemContext) interface{}

	// Visit a parse tree produced by RedshiftParser#createfunc_opt_item.
	VisitCreatefunc_opt_item(ctx *Createfunc_opt_itemContext) interface{}

	// Visit a parse tree produced by RedshiftParser#func_as.
	VisitFunc_as(ctx *Func_asContext) interface{}

	// Visit a parse tree produced by RedshiftParser#transform_type_list.
	VisitTransform_type_list(ctx *Transform_type_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_definition.
	VisitOpt_definition(ctx *Opt_definitionContext) interface{}

	// Visit a parse tree produced by RedshiftParser#table_func_column.
	VisitTable_func_column(ctx *Table_func_columnContext) interface{}

	// Visit a parse tree produced by RedshiftParser#table_func_column_list.
	VisitTable_func_column_list(ctx *Table_func_column_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#alterfunctionstmt.
	VisitAlterfunctionstmt(ctx *AlterfunctionstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#alterprocedurestmt.
	VisitAlterprocedurestmt(ctx *AlterprocedurestmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#alterfunc_opt_list.
	VisitAlterfunc_opt_list(ctx *Alterfunc_opt_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_restrict.
	VisitOpt_restrict(ctx *Opt_restrictContext) interface{}

	// Visit a parse tree produced by RedshiftParser#removefuncstmt.
	VisitRemovefuncstmt(ctx *RemovefuncstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#removeaggrstmt.
	VisitRemoveaggrstmt(ctx *RemoveaggrstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#removeoperstmt.
	VisitRemoveoperstmt(ctx *RemoveoperstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#oper_argtypes.
	VisitOper_argtypes(ctx *Oper_argtypesContext) interface{}

	// Visit a parse tree produced by RedshiftParser#any_operator.
	VisitAny_operator(ctx *Any_operatorContext) interface{}

	// Visit a parse tree produced by RedshiftParser#operator_with_argtypes_list.
	VisitOperator_with_argtypes_list(ctx *Operator_with_argtypes_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#operator_with_argtypes.
	VisitOperator_with_argtypes(ctx *Operator_with_argtypesContext) interface{}

	// Visit a parse tree produced by RedshiftParser#dostmt.
	VisitDostmt(ctx *DostmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#dostmt_opt_list.
	VisitDostmt_opt_list(ctx *Dostmt_opt_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#dostmt_opt_item.
	VisitDostmt_opt_item(ctx *Dostmt_opt_itemContext) interface{}

	// Visit a parse tree produced by RedshiftParser#createcaststmt.
	VisitCreatecaststmt(ctx *CreatecaststmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#cast_context.
	VisitCast_context(ctx *Cast_contextContext) interface{}

	// Visit a parse tree produced by RedshiftParser#dropcaststmt.
	VisitDropcaststmt(ctx *DropcaststmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_if_exists.
	VisitOpt_if_exists(ctx *Opt_if_existsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#createtransformstmt.
	VisitCreatetransformstmt(ctx *CreatetransformstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#transform_element_list.
	VisitTransform_element_list(ctx *Transform_element_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#droptransformstmt.
	VisitDroptransformstmt(ctx *DroptransformstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#reindexstmt.
	VisitReindexstmt(ctx *ReindexstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#reindex_target_type.
	VisitReindex_target_type(ctx *Reindex_target_typeContext) interface{}

	// Visit a parse tree produced by RedshiftParser#reindex_target_multitable.
	VisitReindex_target_multitable(ctx *Reindex_target_multitableContext) interface{}

	// Visit a parse tree produced by RedshiftParser#reindex_option_list.
	VisitReindex_option_list(ctx *Reindex_option_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#reindex_option_elem.
	VisitReindex_option_elem(ctx *Reindex_option_elemContext) interface{}

	// Visit a parse tree produced by RedshiftParser#altertblspcstmt.
	VisitAltertblspcstmt(ctx *AltertblspcstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#renamestmt.
	VisitRenamestmt(ctx *RenamestmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_column.
	VisitOpt_column(ctx *Opt_columnContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_set_data.
	VisitOpt_set_data(ctx *Opt_set_dataContext) interface{}

	// Visit a parse tree produced by RedshiftParser#alterobjectdependsstmt.
	VisitAlterobjectdependsstmt(ctx *AlterobjectdependsstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_no.
	VisitOpt_no(ctx *Opt_noContext) interface{}

	// Visit a parse tree produced by RedshiftParser#alterobjectschemastmt.
	VisitAlterobjectschemastmt(ctx *AlterobjectschemastmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#alteroperatorstmt.
	VisitAlteroperatorstmt(ctx *AlteroperatorstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#operator_def_list.
	VisitOperator_def_list(ctx *Operator_def_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#operator_def_elem.
	VisitOperator_def_elem(ctx *Operator_def_elemContext) interface{}

	// Visit a parse tree produced by RedshiftParser#operator_def_arg.
	VisitOperator_def_arg(ctx *Operator_def_argContext) interface{}

	// Visit a parse tree produced by RedshiftParser#altertypestmt.
	VisitAltertypestmt(ctx *AltertypestmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#alterownerstmt.
	VisitAlterownerstmt(ctx *AlterownerstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#createpublicationstmt.
	VisitCreatepublicationstmt(ctx *CreatepublicationstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#pub_obj_list.
	VisitPub_obj_list(ctx *Pub_obj_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#publication_obj_spec.
	VisitPublication_obj_spec(ctx *Publication_obj_specContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_where_clause.
	VisitOpt_where_clause(ctx *Opt_where_clauseContext) interface{}

	// Visit a parse tree produced by RedshiftParser#alterpublicationstmt.
	VisitAlterpublicationstmt(ctx *AlterpublicationstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#createsubscriptionstmt.
	VisitCreatesubscriptionstmt(ctx *CreatesubscriptionstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#publication_name_list.
	VisitPublication_name_list(ctx *Publication_name_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#publication_name_item.
	VisitPublication_name_item(ctx *Publication_name_itemContext) interface{}

	// Visit a parse tree produced by RedshiftParser#altersubscriptionstmt.
	VisitAltersubscriptionstmt(ctx *AltersubscriptionstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#dropsubscriptionstmt.
	VisitDropsubscriptionstmt(ctx *DropsubscriptionstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#rulestmt.
	VisitRulestmt(ctx *RulestmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#ruleactionlist.
	VisitRuleactionlist(ctx *RuleactionlistContext) interface{}

	// Visit a parse tree produced by RedshiftParser#ruleactionmulti.
	VisitRuleactionmulti(ctx *RuleactionmultiContext) interface{}

	// Visit a parse tree produced by RedshiftParser#ruleactionstmt.
	VisitRuleactionstmt(ctx *RuleactionstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#ruleactionstmtOrEmpty.
	VisitRuleactionstmtOrEmpty(ctx *RuleactionstmtOrEmptyContext) interface{}

	// Visit a parse tree produced by RedshiftParser#event.
	VisitEvent(ctx *EventContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_instead.
	VisitOpt_instead(ctx *Opt_insteadContext) interface{}

	// Visit a parse tree produced by RedshiftParser#notifystmt.
	VisitNotifystmt(ctx *NotifystmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#notify_payload.
	VisitNotify_payload(ctx *Notify_payloadContext) interface{}

	// Visit a parse tree produced by RedshiftParser#listenstmt.
	VisitListenstmt(ctx *ListenstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#unlistenstmt.
	VisitUnlistenstmt(ctx *UnlistenstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#transactionstmt.
	VisitTransactionstmt(ctx *TransactionstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_transaction.
	VisitOpt_transaction(ctx *Opt_transactionContext) interface{}

	// Visit a parse tree produced by RedshiftParser#transaction_mode_item.
	VisitTransaction_mode_item(ctx *Transaction_mode_itemContext) interface{}

	// Visit a parse tree produced by RedshiftParser#transaction_mode_list.
	VisitTransaction_mode_list(ctx *Transaction_mode_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#transaction_mode_list_or_empty.
	VisitTransaction_mode_list_or_empty(ctx *Transaction_mode_list_or_emptyContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_transaction_chain.
	VisitOpt_transaction_chain(ctx *Opt_transaction_chainContext) interface{}

	// Visit a parse tree produced by RedshiftParser#viewstmt.
	VisitViewstmt(ctx *ViewstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#with_no_schema_binding.
	VisitWith_no_schema_binding(ctx *With_no_schema_bindingContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_check_option.
	VisitOpt_check_option(ctx *Opt_check_optionContext) interface{}

	// Visit a parse tree produced by RedshiftParser#loadstmt.
	VisitLoadstmt(ctx *LoadstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#alterdatasharestmt.
	VisitAlterdatasharestmt(ctx *AlterdatasharestmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#alterdatashare_action.
	VisitAlterdatashare_action(ctx *Alterdatashare_actionContext) interface{}

	// Visit a parse tree produced by RedshiftParser#alterdatashare_add_drop.
	VisitAlterdatashare_add_drop(ctx *Alterdatashare_add_dropContext) interface{}

	// Visit a parse tree produced by RedshiftParser#alterdatashare_objects.
	VisitAlterdatashare_objects(ctx *Alterdatashare_objectsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#datashare_table_list.
	VisitDatashare_table_list(ctx *Datashare_table_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#datashare_table_name.
	VisitDatashare_table_name(ctx *Datashare_table_nameContext) interface{}

	// Visit a parse tree produced by RedshiftParser#table_name.
	VisitTable_name(ctx *Table_nameContext) interface{}

	// Visit a parse tree produced by RedshiftParser#temporary_table_name.
	VisitTemporary_table_name(ctx *Temporary_table_nameContext) interface{}

	// Visit a parse tree produced by RedshiftParser#datashare_function_list.
	VisitDatashare_function_list(ctx *Datashare_function_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#datashare_function.
	VisitDatashare_function(ctx *Datashare_functionContext) interface{}

	// Visit a parse tree produced by RedshiftParser#datashare_function_name.
	VisitDatashare_function_name(ctx *Datashare_function_nameContext) interface{}

	// Visit a parse tree produced by RedshiftParser#createdatasharestmt.
	VisitCreatedatasharestmt(ctx *CreatedatasharestmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#createdatashareoptions.
	VisitCreatedatashareoptions(ctx *CreatedatashareoptionsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#createdatashareoption.
	VisitCreatedatashareoption(ctx *CreatedatashareoptionContext) interface{}

	// Visit a parse tree produced by RedshiftParser#setpublicaccessibleoption.
	VisitSetpublicaccessibleoption(ctx *SetpublicaccessibleoptionContext) interface{}

	// Visit a parse tree produced by RedshiftParser#managedbyoption.
	VisitManagedbyoption(ctx *ManagedbyoptionContext) interface{}

	// Visit a parse tree produced by RedshiftParser#descdatasharestmt.
	VisitDescdatasharestmt(ctx *DescdatasharestmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#dropdatasharestmt.
	VisitDropdatasharestmt(ctx *DropdatasharestmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#alterexternalschemastmt.
	VisitAlterexternalschemastmt(ctx *AlterexternalschemastmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#altexternalschemaopts.
	VisitAltexternalschemaopts(ctx *AltexternalschemaoptsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#alterexternalviewstmt.
	VisitAlterexternalviewstmt(ctx *AlterexternalviewstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#createexternalschemastmt.
	VisitCreateexternalschemastmt(ctx *CreateexternalschemastmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#fromdatacatalogclause.
	VisitFromdatacatalogclause(ctx *FromdatacatalogclauseContext) interface{}

	// Visit a parse tree produced by RedshiftParser#dropschemastmt.
	VisitDropschemastmt(ctx *DropschemastmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#implicitdatacatalogclause.
	VisitImplicitdatacatalogclause(ctx *ImplicitdatacatalogclauseContext) interface{}

	// Visit a parse tree produced by RedshiftParser#fromhivemetastoreclause.
	VisitFromhivemetastoreclause(ctx *FromhivemetastoreclauseContext) interface{}

	// Visit a parse tree produced by RedshiftParser#frompostgresclause.
	VisitFrompostgresclause(ctx *FrompostgresclauseContext) interface{}

	// Visit a parse tree produced by RedshiftParser#frommysqlclause.
	VisitFrommysqlclause(ctx *FrommysqlclauseContext) interface{}

	// Visit a parse tree produced by RedshiftParser#fromkinesisclause.
	VisitFromkinesisclause(ctx *FromkinesisclauseContext) interface{}

	// Visit a parse tree produced by RedshiftParser#fromkafkaclause.
	VisitFromkafkaclause(ctx *FromkafkaclauseContext) interface{}

	// Visit a parse tree produced by RedshiftParser#frommskclause.
	VisitFrommskclause(ctx *FrommskclauseContext) interface{}

	// Visit a parse tree produced by RedshiftParser#fromredshiftclause.
	VisitFromredshiftclause(ctx *FromredshiftclauseContext) interface{}

	// Visit a parse tree produced by RedshiftParser#iamrolevalue.
	VisitIamrolevalue(ctx *IamrolevalueContext) interface{}

	// Visit a parse tree produced by RedshiftParser#catalogrolevalue.
	VisitCatalogrolevalue(ctx *CatalogrolevalueContext) interface{}

	// Visit a parse tree produced by RedshiftParser#authenticationvalue.
	VisitAuthenticationvalue(ctx *AuthenticationvalueContext) interface{}

	// Visit a parse tree produced by RedshiftParser#createexternalfunctionstmt.
	VisitCreateexternalfunctionstmt(ctx *CreateexternalfunctionstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#external_func_params.
	VisitExternal_func_params(ctx *External_func_paramsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#paramlist.
	VisitParamlist(ctx *ParamlistContext) interface{}

	// Visit a parse tree produced by RedshiftParser#param_spec.
	VisitParam_spec(ctx *Param_specContext) interface{}

	// Visit a parse tree produced by RedshiftParser#createexternalmodelstmt.
	VisitCreateexternalmodelstmt(ctx *CreateexternalmodelstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#createexternaltablestmt.
	VisitCreateexternaltablestmt(ctx *CreateexternaltablestmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#extern_column_list.
	VisitExtern_column_list(ctx *Extern_column_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#extern_column_def.
	VisitExtern_column_def(ctx *Extern_column_defContext) interface{}

	// Visit a parse tree produced by RedshiftParser#extern_typename.
	VisitExtern_typename(ctx *Extern_typenameContext) interface{}

	// Visit a parse tree produced by RedshiftParser#extern_table_format.
	VisitExtern_table_format(ctx *Extern_table_formatContext) interface{}

	// Visit a parse tree produced by RedshiftParser#row_format_spec.
	VisitRow_format_spec(ctx *Row_format_specContext) interface{}

	// Visit a parse tree produced by RedshiftParser#serde_properties_list.
	VisitSerde_properties_list(ctx *Serde_properties_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#serde_property.
	VisitSerde_property(ctx *Serde_propertyContext) interface{}

	// Visit a parse tree produced by RedshiftParser#external_format_spec.
	VisitExternal_format_spec(ctx *External_format_specContext) interface{}

	// Visit a parse tree produced by RedshiftParser#table_properties_list.
	VisitTable_properties_list(ctx *Table_properties_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#table_property.
	VisitTable_property(ctx *Table_propertyContext) interface{}

	// Visit a parse tree produced by RedshiftParser#createexternalviewstmt.
	VisitCreateexternalviewstmt(ctx *CreateexternalviewstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#dropexternalviewstmt.
	VisitDropexternalviewstmt(ctx *DropexternalviewstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#alteridentityproviderstmt.
	VisitAlteridentityproviderstmt(ctx *AlteridentityproviderstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#alteridprovideropts.
	VisitAlteridprovideropts(ctx *AlteridprovideroptsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#altermaskingpolicystmt.
	VisitAltermaskingpolicystmt(ctx *AltermaskingpolicystmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#altermaterializedviewstmt.
	VisitAltermaterializedviewstmt(ctx *AltermaterializedviewstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#altmaskingpolicyopts.
	VisitAltmaskingpolicyopts(ctx *AltmaskingpolicyoptsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#altmaskingpolicyargs.
	VisitAltmaskingpolicyargs(ctx *AltmaskingpolicyargsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#altmaskingpolicyarg.
	VisitAltmaskingpolicyarg(ctx *AltmaskingpolicyargContext) interface{}

	// Visit a parse tree produced by RedshiftParser#alterrlspolicystmt.
	VisitAlterrlspolicystmt(ctx *AlterrlspolicystmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#attachmaskingpolicystmt.
	VisitAttachmaskingpolicystmt(ctx *AttachmaskingpolicystmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#attachpolicycollist.
	VisitAttachpolicycollist(ctx *AttachpolicycollistContext) interface{}

	// Visit a parse tree produced by RedshiftParser#attachpolicycolumn.
	VisitAttachpolicycolumn(ctx *AttachpolicycolumnContext) interface{}

	// Visit a parse tree produced by RedshiftParser#attachpolicytargets.
	VisitAttachpolicytargets(ctx *AttachpolicytargetsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#attachpolicytarget.
	VisitAttachpolicytarget(ctx *AttachpolicytargetContext) interface{}

	// Visit a parse tree produced by RedshiftParser#attachrlspolicystmt.
	VisitAttachrlspolicystmt(ctx *AttachrlspolicystmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#table_name_list.
	VisitTable_name_list(ctx *Table_name_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#createidentityproviderstmt.
	VisitCreateidentityproviderstmt(ctx *CreateidentityproviderstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#createidprovideropts.
	VisitCreateidprovideropts(ctx *CreateidprovideroptsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#groupfilter.
	VisitGroupfilter(ctx *GroupfilterContext) interface{}

	// Visit a parse tree produced by RedshiftParser#createlibrarystmt.
	VisitCreatelibrarystmt(ctx *CreatelibrarystmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#createlibraryopts.
	VisitCreatelibraryopts(ctx *CreatelibraryoptsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#createmaskingpolicystmt.
	VisitCreatemaskingpolicystmt(ctx *CreatemaskingpolicystmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#inputcolumnlist.
	VisitInputcolumnlist(ctx *InputcolumnlistContext) interface{}

	// Visit a parse tree produced by RedshiftParser#inputcolumn.
	VisitInputcolumn(ctx *InputcolumnContext) interface{}

	// Visit a parse tree produced by RedshiftParser#maskingexpression.
	VisitMaskingexpression(ctx *MaskingexpressionContext) interface{}

	// Visit a parse tree produced by RedshiftParser#createmodelstmt.
	VisitCreatemodelstmt(ctx *CreatemodelstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#createmodelfromclause.
	VisitCreatemodelfromclause(ctx *CreatemodelfromclauseContext) interface{}

	// Visit a parse tree produced by RedshiftParser#iamrolespec.
	VisitIamrolespec(ctx *IamrolespecContext) interface{}

	// Visit a parse tree produced by RedshiftParser#sagemakerspec.
	VisitSagemakerspec(ctx *SagemakerspecContext) interface{}

	// Visit a parse tree produced by RedshiftParser#modeltypespec.
	VisitModeltypespec(ctx *ModeltypespecContext) interface{}

	// Visit a parse tree produced by RedshiftParser#problemtypespec.
	VisitProblemtypespec(ctx *ProblemtypespecContext) interface{}

	// Visit a parse tree produced by RedshiftParser#problemtype.
	VisitProblemtype(ctx *ProblemtypeContext) interface{}

	// Visit a parse tree produced by RedshiftParser#objectivespec.
	VisitObjectivespec(ctx *ObjectivespecContext) interface{}

	// Visit a parse tree produced by RedshiftParser#hyperparametersspec.
	VisitHyperparametersspec(ctx *HyperparametersspecContext) interface{}

	// Visit a parse tree produced by RedshiftParser#hyperparameterslist.
	VisitHyperparameterslist(ctx *HyperparameterslistContext) interface{}

	// Visit a parse tree produced by RedshiftParser#hyperparameteritem.
	VisitHyperparameteritem(ctx *HyperparameteritemContext) interface{}

	// Visit a parse tree produced by RedshiftParser#settingsclause.
	VisitSettingsclause(ctx *SettingsclauseContext) interface{}

	// Visit a parse tree produced by RedshiftParser#settingsitem.
	VisitSettingsitem(ctx *SettingsitemContext) interface{}

	// Visit a parse tree produced by RedshiftParser#datatypelist.
	VisitDatatypelist(ctx *DatatypelistContext) interface{}

	// Visit a parse tree produced by RedshiftParser#datatype.
	VisitDatatype(ctx *DatatypeContext) interface{}

	// Visit a parse tree produced by RedshiftParser#createrlspolicystmt.
	VisitCreaterlspolicystmt(ctx *CreaterlspolicystmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#descidentityproviderstmt.
	VisitDescidentityproviderstmt(ctx *DescidentityproviderstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#detachmaskingpolicystmt.
	VisitDetachmaskingpolicystmt(ctx *DetachmaskingpolicystmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#detachrlspolicystmt.
	VisitDetachrlspolicystmt(ctx *DetachrlspolicystmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#role_or_user_or_public_list.
	VisitRole_or_user_or_public_list(ctx *Role_or_user_or_public_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#role_or_user_or_public.
	VisitRole_or_user_or_public(ctx *Role_or_user_or_publicContext) interface{}

	// Visit a parse tree produced by RedshiftParser#rlspolicyname.
	VisitRlspolicyname(ctx *RlspolicynameContext) interface{}

	// Visit a parse tree produced by RedshiftParser#dropidentityproviderstmt.
	VisitDropidentityproviderstmt(ctx *DropidentityproviderstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#droplibrarystmt.
	VisitDroplibrarystmt(ctx *DroplibrarystmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#dropmaskingpolicystmt.
	VisitDropmaskingpolicystmt(ctx *DropmaskingpolicystmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#dropmodelstmt.
	VisitDropmodelstmt(ctx *DropmodelstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#droprlspolicystmt.
	VisitDroprlspolicystmt(ctx *DroprlspolicystmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#altertableappendstmt.
	VisitAltertableappendstmt(ctx *AltertableappendstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#appendoptions.
	VisitAppendoptions(ctx *AppendoptionsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#alteruserstmt.
	VisitAlteruserstmt(ctx *AlteruserstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#alteruseropts.
	VisitAlteruseropts(ctx *AlteruseroptsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#analyzecompressionstmt.
	VisitAnalyzecompressionstmt(ctx *AnalyzecompressionstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#cancelstmt.
	VisitCancelstmt(ctx *CancelstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#closestmt.
	VisitClosestmt(ctx *ClosestmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#insertexternaltablestmt.
	VisitInsertexternaltablestmt(ctx *InsertexternaltablestmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#select_or_values.
	VisitSelect_or_values(ctx *Select_or_valuesContext) interface{}

	// Visit a parse tree produced by RedshiftParser#selectintostmt.
	VisitSelectintostmt(ctx *SelectintostmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#setsessionauthorizationstmt.
	VisitSetsessionauthorizationstmt(ctx *SetsessionauthorizationstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#setsessioncharacteristicsstmt.
	VisitSetsessioncharacteristicsstmt(ctx *SetsessioncharacteristicsstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#showcolumnsstmt.
	VisitShowcolumnsstmt(ctx *ShowcolumnsstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#showdatabasesstmt.
	VisitShowdatabasesstmt(ctx *ShowdatabasesstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#showdbsopts.
	VisitShowdbsopts(ctx *ShowdbsoptsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#showdatasharesstmt.
	VisitShowdatasharesstmt(ctx *ShowdatasharesstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#showexternaltablestmt.
	VisitShowexternaltablestmt(ctx *ShowexternaltablestmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#showgrantsstmt.
	VisitShowgrantsstmt(ctx *ShowgrantsstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#grantobject.
	VisitGrantobject(ctx *GrantobjectContext) interface{}

	// Visit a parse tree produced by RedshiftParser#grantprincipal.
	VisitGrantprincipal(ctx *GrantprincipalContext) interface{}

	// Visit a parse tree produced by RedshiftParser#showmodelstmt.
	VisitShowmodelstmt(ctx *ShowmodelstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#showprocedurestmt.
	VisitShowprocedurestmt(ctx *ShowprocedurestmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#showschemasstmt.
	VisitShowschemasstmt(ctx *ShowschemasstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#showtablestmt.
	VisitShowtablestmt(ctx *ShowtablestmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#showtablesstmt.
	VisitShowtablesstmt(ctx *ShowtablesstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#showviewstmt.
	VisitShowviewstmt(ctx *ShowviewstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#unloadstmt.
	VisitUnloadstmt(ctx *UnloadstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#iamroleclause.
	VisitIamroleclause(ctx *IamroleclauseContext) interface{}

	// Visit a parse tree produced by RedshiftParser#unloadoptions.
	VisitUnloadoptions(ctx *UnloadoptionsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#formatoption.
	VisitFormatoption(ctx *FormatoptionContext) interface{}

	// Visit a parse tree produced by RedshiftParser#partitionbyoption.
	VisitPartitionbyoption(ctx *PartitionbyoptionContext) interface{}

	// Visit a parse tree produced by RedshiftParser#manifestoption.
	VisitManifestoption(ctx *ManifestoptionContext) interface{}

	// Visit a parse tree produced by RedshiftParser#headeroption.
	VisitHeaderoption(ctx *HeaderoptionContext) interface{}

	// Visit a parse tree produced by RedshiftParser#delimiteroption.
	VisitDelimiteroption(ctx *DelimiteroptionContext) interface{}

	// Visit a parse tree produced by RedshiftParser#fixedwidthoption.
	VisitFixedwidthoption(ctx *FixedwidthoptionContext) interface{}

	// Visit a parse tree produced by RedshiftParser#encryptedoption.
	VisitEncryptedoption(ctx *EncryptedoptionContext) interface{}

	// Visit a parse tree produced by RedshiftParser#kmskeyoption.
	VisitKmskeyoption(ctx *KmskeyoptionContext) interface{}

	// Visit a parse tree produced by RedshiftParser#compressionoption.
	VisitCompressionoption(ctx *CompressionoptionContext) interface{}

	// Visit a parse tree produced by RedshiftParser#addquotesoption.
	VisitAddquotesoption(ctx *AddquotesoptionContext) interface{}

	// Visit a parse tree produced by RedshiftParser#nullasoption.
	VisitNullasoption(ctx *NullasoptionContext) interface{}

	// Visit a parse tree produced by RedshiftParser#escapeoption.
	VisitEscapeoption(ctx *EscapeoptionContext) interface{}

	// Visit a parse tree produced by RedshiftParser#allowoverwriteoption.
	VisitAllowoverwriteoption(ctx *AllowoverwriteoptionContext) interface{}

	// Visit a parse tree produced by RedshiftParser#cleanpathoption.
	VisitCleanpathoption(ctx *CleanpathoptionContext) interface{}

	// Visit a parse tree produced by RedshiftParser#paralleloption.
	VisitParalleloption(ctx *ParalleloptionContext) interface{}

	// Visit a parse tree produced by RedshiftParser#maxfilesizeoption.
	VisitMaxfilesizeoption(ctx *MaxfilesizeoptionContext) interface{}

	// Visit a parse tree produced by RedshiftParser#rowgroupsizeoption.
	VisitRowgroupsizeoption(ctx *RowgroupsizeoptionContext) interface{}

	// Visit a parse tree produced by RedshiftParser#sizeunit.
	VisitSizeunit(ctx *SizeunitContext) interface{}

	// Visit a parse tree produced by RedshiftParser#regionoption.
	VisitRegionoption(ctx *RegionoptionContext) interface{}

	// Visit a parse tree produced by RedshiftParser#extensionoption.
	VisitExtensionoption(ctx *ExtensionoptionContext) interface{}

	// Visit a parse tree produced by RedshiftParser#usestmt.
	VisitUsestmt(ctx *UsestmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#createdbstmt.
	VisitCreatedbstmt(ctx *CreatedbstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#createdb_opt_list.
	VisitCreatedb_opt_list(ctx *Createdb_opt_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#createdb_opt_items.
	VisitCreatedb_opt_items(ctx *Createdb_opt_itemsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#createdb_opt_item.
	VisitCreatedb_opt_item(ctx *Createdb_opt_itemContext) interface{}

	// Visit a parse tree produced by RedshiftParser#createdb_opt_name.
	VisitCreatedb_opt_name(ctx *Createdb_opt_nameContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_equal.
	VisitOpt_equal(ctx *Opt_equalContext) interface{}

	// Visit a parse tree produced by RedshiftParser#alterdatabasestmt.
	VisitAlterdatabasestmt(ctx *AlterdatabasestmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#alterdatabasesetstmt.
	VisitAlterdatabasesetstmt(ctx *AlterdatabasesetstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#dropdbstmt.
	VisitDropdbstmt(ctx *DropdbstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#drop_option_list.
	VisitDrop_option_list(ctx *Drop_option_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#drop_option.
	VisitDrop_option(ctx *Drop_optionContext) interface{}

	// Visit a parse tree produced by RedshiftParser#altercollationstmt.
	VisitAltercollationstmt(ctx *AltercollationstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#altersystemstmt.
	VisitAltersystemstmt(ctx *AltersystemstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#createdomainstmt.
	VisitCreatedomainstmt(ctx *CreatedomainstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#alterdomainstmt.
	VisitAlterdomainstmt(ctx *AlterdomainstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_as.
	VisitOpt_as(ctx *Opt_asContext) interface{}

	// Visit a parse tree produced by RedshiftParser#altertsdictionarystmt.
	VisitAltertsdictionarystmt(ctx *AltertsdictionarystmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#altertsconfigurationstmt.
	VisitAltertsconfigurationstmt(ctx *AltertsconfigurationstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#any_with.
	VisitAny_with(ctx *Any_withContext) interface{}

	// Visit a parse tree produced by RedshiftParser#createconversionstmt.
	VisitCreateconversionstmt(ctx *CreateconversionstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#clusterstmt.
	VisitClusterstmt(ctx *ClusterstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#cluster_index_specification.
	VisitCluster_index_specification(ctx *Cluster_index_specificationContext) interface{}

	// Visit a parse tree produced by RedshiftParser#vacuumstmt.
	VisitVacuumstmt(ctx *VacuumstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#vacuum_option.
	VisitVacuum_option(ctx *Vacuum_optionContext) interface{}

	// Visit a parse tree produced by RedshiftParser#analyzestmt.
	VisitAnalyzestmt(ctx *AnalyzestmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#vac_analyze_option_list.
	VisitVac_analyze_option_list(ctx *Vac_analyze_option_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#analyze_keyword.
	VisitAnalyze_keyword(ctx *Analyze_keywordContext) interface{}

	// Visit a parse tree produced by RedshiftParser#vac_analyze_option_elem.
	VisitVac_analyze_option_elem(ctx *Vac_analyze_option_elemContext) interface{}

	// Visit a parse tree produced by RedshiftParser#vac_analyze_option_name.
	VisitVac_analyze_option_name(ctx *Vac_analyze_option_nameContext) interface{}

	// Visit a parse tree produced by RedshiftParser#vac_analyze_option_arg.
	VisitVac_analyze_option_arg(ctx *Vac_analyze_option_argContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_analyze.
	VisitOpt_analyze(ctx *Opt_analyzeContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_verbose.
	VisitOpt_verbose(ctx *Opt_verboseContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_full.
	VisitOpt_full(ctx *Opt_fullContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_freeze.
	VisitOpt_freeze(ctx *Opt_freezeContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_name_list.
	VisitOpt_name_list(ctx *Opt_name_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#vacuum_relation.
	VisitVacuum_relation(ctx *Vacuum_relationContext) interface{}

	// Visit a parse tree produced by RedshiftParser#vacuum_relation_list.
	VisitVacuum_relation_list(ctx *Vacuum_relation_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_vacuum_relation_list.
	VisitOpt_vacuum_relation_list(ctx *Opt_vacuum_relation_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#explainstmt.
	VisitExplainstmt(ctx *ExplainstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#explainablestmt.
	VisitExplainablestmt(ctx *ExplainablestmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#explain_option_list.
	VisitExplain_option_list(ctx *Explain_option_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#explain_option_elem.
	VisitExplain_option_elem(ctx *Explain_option_elemContext) interface{}

	// Visit a parse tree produced by RedshiftParser#explain_option_name.
	VisitExplain_option_name(ctx *Explain_option_nameContext) interface{}

	// Visit a parse tree produced by RedshiftParser#explain_option_arg.
	VisitExplain_option_arg(ctx *Explain_option_argContext) interface{}

	// Visit a parse tree produced by RedshiftParser#preparestmt.
	VisitPreparestmt(ctx *PreparestmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#prep_type_clause.
	VisitPrep_type_clause(ctx *Prep_type_clauseContext) interface{}

	// Visit a parse tree produced by RedshiftParser#preparablestmt.
	VisitPreparablestmt(ctx *PreparablestmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#executestmt.
	VisitExecutestmt(ctx *ExecutestmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#execute_param_clause.
	VisitExecute_param_clause(ctx *Execute_param_clauseContext) interface{}

	// Visit a parse tree produced by RedshiftParser#deallocatestmt.
	VisitDeallocatestmt(ctx *DeallocatestmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#insertstmt.
	VisitInsertstmt(ctx *InsertstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#insert_target.
	VisitInsert_target(ctx *Insert_targetContext) interface{}

	// Visit a parse tree produced by RedshiftParser#insert_rest.
	VisitInsert_rest(ctx *Insert_restContext) interface{}

	// Visit a parse tree produced by RedshiftParser#override_kind.
	VisitOverride_kind(ctx *Override_kindContext) interface{}

	// Visit a parse tree produced by RedshiftParser#insert_column_list.
	VisitInsert_column_list(ctx *Insert_column_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#insert_column_item.
	VisitInsert_column_item(ctx *Insert_column_itemContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_on_conflict.
	VisitOpt_on_conflict(ctx *Opt_on_conflictContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_conf_expr.
	VisitOpt_conf_expr(ctx *Opt_conf_exprContext) interface{}

	// Visit a parse tree produced by RedshiftParser#returning_clause.
	VisitReturning_clause(ctx *Returning_clauseContext) interface{}

	// Visit a parse tree produced by RedshiftParser#mergestmt.
	VisitMergestmt(ctx *MergestmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#merge_when_clause.
	VisitMerge_when_clause(ctx *Merge_when_clauseContext) interface{}

	// Visit a parse tree produced by RedshiftParser#merge_insert_clause.
	VisitMerge_insert_clause(ctx *Merge_insert_clauseContext) interface{}

	// Visit a parse tree produced by RedshiftParser#merge_update_clause.
	VisitMerge_update_clause(ctx *Merge_update_clauseContext) interface{}

	// Visit a parse tree produced by RedshiftParser#merge_delete_clause.
	VisitMerge_delete_clause(ctx *Merge_delete_clauseContext) interface{}

	// Visit a parse tree produced by RedshiftParser#deletestmt.
	VisitDeletestmt(ctx *DeletestmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#using_clause.
	VisitUsing_clause(ctx *Using_clauseContext) interface{}

	// Visit a parse tree produced by RedshiftParser#lockstmt.
	VisitLockstmt(ctx *LockstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_lock.
	VisitOpt_lock(ctx *Opt_lockContext) interface{}

	// Visit a parse tree produced by RedshiftParser#lock_type.
	VisitLock_type(ctx *Lock_typeContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_nowait.
	VisitOpt_nowait(ctx *Opt_nowaitContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_nowait_or_skip.
	VisitOpt_nowait_or_skip(ctx *Opt_nowait_or_skipContext) interface{}

	// Visit a parse tree produced by RedshiftParser#updatestmt.
	VisitUpdatestmt(ctx *UpdatestmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#set_clause_list.
	VisitSet_clause_list(ctx *Set_clause_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#set_clause.
	VisitSet_clause(ctx *Set_clauseContext) interface{}

	// Visit a parse tree produced by RedshiftParser#set_target.
	VisitSet_target(ctx *Set_targetContext) interface{}

	// Visit a parse tree produced by RedshiftParser#set_target_list.
	VisitSet_target_list(ctx *Set_target_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#declarecursorstmt.
	VisitDeclarecursorstmt(ctx *DeclarecursorstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#cursor_name.
	VisitCursor_name(ctx *Cursor_nameContext) interface{}

	// Visit a parse tree produced by RedshiftParser#cursor_options.
	VisitCursor_options(ctx *Cursor_optionsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_hold.
	VisitOpt_hold(ctx *Opt_holdContext) interface{}

	// Visit a parse tree produced by RedshiftParser#selectstmt.
	VisitSelectstmt(ctx *SelectstmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#select_with_parens.
	VisitSelect_with_parens(ctx *Select_with_parensContext) interface{}

	// Visit a parse tree produced by RedshiftParser#select_no_parens.
	VisitSelect_no_parens(ctx *Select_no_parensContext) interface{}

	// Visit a parse tree produced by RedshiftParser#select_clause.
	VisitSelect_clause(ctx *Select_clauseContext) interface{}

	// Visit a parse tree produced by RedshiftParser#simple_select_intersect.
	VisitSimple_select_intersect(ctx *Simple_select_intersectContext) interface{}

	// Visit a parse tree produced by RedshiftParser#simple_select_pramary.
	VisitSimple_select_pramary(ctx *Simple_select_pramaryContext) interface{}

	// Visit a parse tree produced by RedshiftParser#exclude_clause.
	VisitExclude_clause(ctx *Exclude_clauseContext) interface{}

	// Visit a parse tree produced by RedshiftParser#qualify_clause.
	VisitQualify_clause(ctx *Qualify_clauseContext) interface{}

	// Visit a parse tree produced by RedshiftParser#start_with_clause.
	VisitStart_with_clause(ctx *Start_with_clauseContext) interface{}

	// Visit a parse tree produced by RedshiftParser#with_clause.
	VisitWith_clause(ctx *With_clauseContext) interface{}

	// Visit a parse tree produced by RedshiftParser#cte_list.
	VisitCte_list(ctx *Cte_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#common_table_expr.
	VisitCommon_table_expr(ctx *Common_table_exprContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_materialized.
	VisitOpt_materialized(ctx *Opt_materializedContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_with_clause.
	VisitOpt_with_clause(ctx *Opt_with_clauseContext) interface{}

	// Visit a parse tree produced by RedshiftParser#into_clause.
	VisitInto_clause(ctx *Into_clauseContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_top_clause.
	VisitOpt_top_clause(ctx *Opt_top_clauseContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_strict.
	VisitOpt_strict(ctx *Opt_strictContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opttempTableName.
	VisitOpttempTableName(ctx *OpttempTableNameContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_table.
	VisitOpt_table(ctx *Opt_tableContext) interface{}

	// Visit a parse tree produced by RedshiftParser#all_or_distinct.
	VisitAll_or_distinct(ctx *All_or_distinctContext) interface{}

	// Visit a parse tree produced by RedshiftParser#distinct_clause.
	VisitDistinct_clause(ctx *Distinct_clauseContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_all_clause.
	VisitOpt_all_clause(ctx *Opt_all_clauseContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_sort_clause.
	VisitOpt_sort_clause(ctx *Opt_sort_clauseContext) interface{}

	// Visit a parse tree produced by RedshiftParser#sort_clause.
	VisitSort_clause(ctx *Sort_clauseContext) interface{}

	// Visit a parse tree produced by RedshiftParser#sortby_list.
	VisitSortby_list(ctx *Sortby_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#sortby.
	VisitSortby(ctx *SortbyContext) interface{}

	// Visit a parse tree produced by RedshiftParser#select_limit.
	VisitSelect_limit(ctx *Select_limitContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_select_limit.
	VisitOpt_select_limit(ctx *Opt_select_limitContext) interface{}

	// Visit a parse tree produced by RedshiftParser#limit_clause.
	VisitLimit_clause(ctx *Limit_clauseContext) interface{}

	// Visit a parse tree produced by RedshiftParser#offset_clause.
	VisitOffset_clause(ctx *Offset_clauseContext) interface{}

	// Visit a parse tree produced by RedshiftParser#select_limit_value.
	VisitSelect_limit_value(ctx *Select_limit_valueContext) interface{}

	// Visit a parse tree produced by RedshiftParser#select_offset_value.
	VisitSelect_offset_value(ctx *Select_offset_valueContext) interface{}

	// Visit a parse tree produced by RedshiftParser#select_fetch_first_value.
	VisitSelect_fetch_first_value(ctx *Select_fetch_first_valueContext) interface{}

	// Visit a parse tree produced by RedshiftParser#i_or_f_const.
	VisitI_or_f_const(ctx *I_or_f_constContext) interface{}

	// Visit a parse tree produced by RedshiftParser#row_or_rows.
	VisitRow_or_rows(ctx *Row_or_rowsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#first_or_next.
	VisitFirst_or_next(ctx *First_or_nextContext) interface{}

	// Visit a parse tree produced by RedshiftParser#group_clause.
	VisitGroup_clause(ctx *Group_clauseContext) interface{}

	// Visit a parse tree produced by RedshiftParser#group_by_list.
	VisitGroup_by_list(ctx *Group_by_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#group_by_item.
	VisitGroup_by_item(ctx *Group_by_itemContext) interface{}

	// Visit a parse tree produced by RedshiftParser#empty_grouping_set.
	VisitEmpty_grouping_set(ctx *Empty_grouping_setContext) interface{}

	// Visit a parse tree produced by RedshiftParser#rollup_clause.
	VisitRollup_clause(ctx *Rollup_clauseContext) interface{}

	// Visit a parse tree produced by RedshiftParser#cube_clause.
	VisitCube_clause(ctx *Cube_clauseContext) interface{}

	// Visit a parse tree produced by RedshiftParser#grouping_sets_clause.
	VisitGrouping_sets_clause(ctx *Grouping_sets_clauseContext) interface{}

	// Visit a parse tree produced by RedshiftParser#having_clause.
	VisitHaving_clause(ctx *Having_clauseContext) interface{}

	// Visit a parse tree produced by RedshiftParser#for_locking_clause.
	VisitFor_locking_clause(ctx *For_locking_clauseContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_for_locking_clause.
	VisitOpt_for_locking_clause(ctx *Opt_for_locking_clauseContext) interface{}

	// Visit a parse tree produced by RedshiftParser#for_locking_items.
	VisitFor_locking_items(ctx *For_locking_itemsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#for_locking_item.
	VisitFor_locking_item(ctx *For_locking_itemContext) interface{}

	// Visit a parse tree produced by RedshiftParser#for_locking_strength.
	VisitFor_locking_strength(ctx *For_locking_strengthContext) interface{}

	// Visit a parse tree produced by RedshiftParser#locked_rels_list.
	VisitLocked_rels_list(ctx *Locked_rels_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#values_clause.
	VisitValues_clause(ctx *Values_clauseContext) interface{}

	// Visit a parse tree produced by RedshiftParser#from_clause.
	VisitFrom_clause(ctx *From_clauseContext) interface{}

	// Visit a parse tree produced by RedshiftParser#from_list.
	VisitFrom_list(ctx *From_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#table_ref.
	VisitTable_ref(ctx *Table_refContext) interface{}

	// Visit a parse tree produced by RedshiftParser#joined_table.
	VisitJoined_table(ctx *Joined_tableContext) interface{}

	// Visit a parse tree produced by RedshiftParser#alias_clause.
	VisitAlias_clause(ctx *Alias_clauseContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_alias_clause.
	VisitOpt_alias_clause(ctx *Opt_alias_clauseContext) interface{}

	// Visit a parse tree produced by RedshiftParser#table_alias_clause.
	VisitTable_alias_clause(ctx *Table_alias_clauseContext) interface{}

	// Visit a parse tree produced by RedshiftParser#func_alias_clause.
	VisitFunc_alias_clause(ctx *Func_alias_clauseContext) interface{}

	// Visit a parse tree produced by RedshiftParser#join_type.
	VisitJoin_type(ctx *Join_typeContext) interface{}

	// Visit a parse tree produced by RedshiftParser#join_qual.
	VisitJoin_qual(ctx *Join_qualContext) interface{}

	// Visit a parse tree produced by RedshiftParser#relation_expr.
	VisitRelation_expr(ctx *Relation_exprContext) interface{}

	// Visit a parse tree produced by RedshiftParser#relation_expr_list.
	VisitRelation_expr_list(ctx *Relation_expr_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#relation_expr_opt_alias.
	VisitRelation_expr_opt_alias(ctx *Relation_expr_opt_aliasContext) interface{}

	// Visit a parse tree produced by RedshiftParser#tablesample_clause.
	VisitTablesample_clause(ctx *Tablesample_clauseContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_repeatable_clause.
	VisitOpt_repeatable_clause(ctx *Opt_repeatable_clauseContext) interface{}

	// Visit a parse tree produced by RedshiftParser#func_table.
	VisitFunc_table(ctx *Func_tableContext) interface{}

	// Visit a parse tree produced by RedshiftParser#rowsfrom_item.
	VisitRowsfrom_item(ctx *Rowsfrom_itemContext) interface{}

	// Visit a parse tree produced by RedshiftParser#rowsfrom_list.
	VisitRowsfrom_list(ctx *Rowsfrom_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_col_def_list.
	VisitOpt_col_def_list(ctx *Opt_col_def_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_ordinality.
	VisitOpt_ordinality(ctx *Opt_ordinalityContext) interface{}

	// Visit a parse tree produced by RedshiftParser#where_clause.
	VisitWhere_clause(ctx *Where_clauseContext) interface{}

	// Visit a parse tree produced by RedshiftParser#where_or_current_clause.
	VisitWhere_or_current_clause(ctx *Where_or_current_clauseContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opttablefuncelementlist.
	VisitOpttablefuncelementlist(ctx *OpttablefuncelementlistContext) interface{}

	// Visit a parse tree produced by RedshiftParser#tablefuncelementlist.
	VisitTablefuncelementlist(ctx *TablefuncelementlistContext) interface{}

	// Visit a parse tree produced by RedshiftParser#tablefuncelement.
	VisitTablefuncelement(ctx *TablefuncelementContext) interface{}

	// Visit a parse tree produced by RedshiftParser#xmltable.
	VisitXmltable(ctx *XmltableContext) interface{}

	// Visit a parse tree produced by RedshiftParser#xmltable_column_list.
	VisitXmltable_column_list(ctx *Xmltable_column_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#xmltable_column_el.
	VisitXmltable_column_el(ctx *Xmltable_column_elContext) interface{}

	// Visit a parse tree produced by RedshiftParser#xmltable_column_option_list.
	VisitXmltable_column_option_list(ctx *Xmltable_column_option_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#xmltable_column_option_el.
	VisitXmltable_column_option_el(ctx *Xmltable_column_option_elContext) interface{}

	// Visit a parse tree produced by RedshiftParser#xml_namespace_list.
	VisitXml_namespace_list(ctx *Xml_namespace_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#xml_namespace_el.
	VisitXml_namespace_el(ctx *Xml_namespace_elContext) interface{}

	// Visit a parse tree produced by RedshiftParser#typename.
	VisitTypename(ctx *TypenameContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_array_bounds.
	VisitOpt_array_bounds(ctx *Opt_array_boundsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#simpletypename.
	VisitSimpletypename(ctx *SimpletypenameContext) interface{}

	// Visit a parse tree produced by RedshiftParser#varbyte.
	VisitVarbyte(ctx *VarbyteContext) interface{}

	// Visit a parse tree produced by RedshiftParser#json_type.
	VisitJson_type(ctx *Json_typeContext) interface{}

	// Visit a parse tree produced by RedshiftParser#consttypename.
	VisitConsttypename(ctx *ConsttypenameContext) interface{}

	// Visit a parse tree produced by RedshiftParser#generictype.
	VisitGenerictype(ctx *GenerictypeContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_type_modifiers.
	VisitOpt_type_modifiers(ctx *Opt_type_modifiersContext) interface{}

	// Visit a parse tree produced by RedshiftParser#numeric.
	VisitNumeric(ctx *NumericContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_float.
	VisitOpt_float(ctx *Opt_floatContext) interface{}

	// Visit a parse tree produced by RedshiftParser#bit.
	VisitBit(ctx *BitContext) interface{}

	// Visit a parse tree produced by RedshiftParser#constbit.
	VisitConstbit(ctx *ConstbitContext) interface{}

	// Visit a parse tree produced by RedshiftParser#bitwithlength.
	VisitBitwithlength(ctx *BitwithlengthContext) interface{}

	// Visit a parse tree produced by RedshiftParser#bitwithoutlength.
	VisitBitwithoutlength(ctx *BitwithoutlengthContext) interface{}

	// Visit a parse tree produced by RedshiftParser#character.
	VisitCharacter(ctx *CharacterContext) interface{}

	// Visit a parse tree produced by RedshiftParser#constcharacter.
	VisitConstcharacter(ctx *ConstcharacterContext) interface{}

	// Visit a parse tree produced by RedshiftParser#character_c.
	VisitCharacter_c(ctx *Character_cContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_varying.
	VisitOpt_varying(ctx *Opt_varyingContext) interface{}

	// Visit a parse tree produced by RedshiftParser#constdatetime.
	VisitConstdatetime(ctx *ConstdatetimeContext) interface{}

	// Visit a parse tree produced by RedshiftParser#constinterval.
	VisitConstinterval(ctx *ConstintervalContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_timezone.
	VisitOpt_timezone(ctx *Opt_timezoneContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_interval.
	VisitOpt_interval(ctx *Opt_intervalContext) interface{}

	// Visit a parse tree produced by RedshiftParser#interval_second.
	VisitInterval_second(ctx *Interval_secondContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_escape.
	VisitOpt_escape(ctx *Opt_escapeContext) interface{}

	// Visit a parse tree produced by RedshiftParser#a_expr.
	VisitA_expr(ctx *A_exprContext) interface{}

	// Visit a parse tree produced by RedshiftParser#a_expr_qual.
	VisitA_expr_qual(ctx *A_expr_qualContext) interface{}

	// Visit a parse tree produced by RedshiftParser#a_expr_lessless.
	VisitA_expr_lessless(ctx *A_expr_lesslessContext) interface{}

	// Visit a parse tree produced by RedshiftParser#a_expr_or.
	VisitA_expr_or(ctx *A_expr_orContext) interface{}

	// Visit a parse tree produced by RedshiftParser#a_expr_and.
	VisitA_expr_and(ctx *A_expr_andContext) interface{}

	// Visit a parse tree produced by RedshiftParser#a_expr_between.
	VisitA_expr_between(ctx *A_expr_betweenContext) interface{}

	// Visit a parse tree produced by RedshiftParser#a_expr_in.
	VisitA_expr_in(ctx *A_expr_inContext) interface{}

	// Visit a parse tree produced by RedshiftParser#a_expr_unary_not.
	VisitA_expr_unary_not(ctx *A_expr_unary_notContext) interface{}

	// Visit a parse tree produced by RedshiftParser#a_expr_isnull.
	VisitA_expr_isnull(ctx *A_expr_isnullContext) interface{}

	// Visit a parse tree produced by RedshiftParser#a_expr_is_not.
	VisitA_expr_is_not(ctx *A_expr_is_notContext) interface{}

	// Visit a parse tree produced by RedshiftParser#a_expr_compare.
	VisitA_expr_compare(ctx *A_expr_compareContext) interface{}

	// Visit a parse tree produced by RedshiftParser#a_expr_prior_or_level.
	VisitA_expr_prior_or_level(ctx *A_expr_prior_or_levelContext) interface{}

	// Visit a parse tree produced by RedshiftParser#a_expr_like.
	VisitA_expr_like(ctx *A_expr_likeContext) interface{}

	// Visit a parse tree produced by RedshiftParser#a_expr_qual_op.
	VisitA_expr_qual_op(ctx *A_expr_qual_opContext) interface{}

	// Visit a parse tree produced by RedshiftParser#a_expr_unary_qualop.
	VisitA_expr_unary_qualop(ctx *A_expr_unary_qualopContext) interface{}

	// Visit a parse tree produced by RedshiftParser#a_expr_add.
	VisitA_expr_add(ctx *A_expr_addContext) interface{}

	// Visit a parse tree produced by RedshiftParser#a_expr_mul.
	VisitA_expr_mul(ctx *A_expr_mulContext) interface{}

	// Visit a parse tree produced by RedshiftParser#a_expr_caret.
	VisitA_expr_caret(ctx *A_expr_caretContext) interface{}

	// Visit a parse tree produced by RedshiftParser#a_expr_unary_sign.
	VisitA_expr_unary_sign(ctx *A_expr_unary_signContext) interface{}

	// Visit a parse tree produced by RedshiftParser#a_expr_at_time_zone.
	VisitA_expr_at_time_zone(ctx *A_expr_at_time_zoneContext) interface{}

	// Visit a parse tree produced by RedshiftParser#a_expr_collate.
	VisitA_expr_collate(ctx *A_expr_collateContext) interface{}

	// Visit a parse tree produced by RedshiftParser#a_expr_typecast.
	VisitA_expr_typecast(ctx *A_expr_typecastContext) interface{}

	// Visit a parse tree produced by RedshiftParser#b_expr.
	VisitB_expr(ctx *B_exprContext) interface{}

	// Visit a parse tree produced by RedshiftParser#c_expr_exists.
	VisitC_expr_exists(ctx *C_expr_existsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#c_expr_expr.
	VisitC_expr_expr(ctx *C_expr_exprContext) interface{}

	// Visit a parse tree produced by RedshiftParser#c_expr_case.
	VisitC_expr_case(ctx *C_expr_caseContext) interface{}

	// Visit a parse tree produced by RedshiftParser#plsqlvariablename.
	VisitPlsqlvariablename(ctx *PlsqlvariablenameContext) interface{}

	// Visit a parse tree produced by RedshiftParser#func_application.
	VisitFunc_application(ctx *Func_applicationContext) interface{}

	// Visit a parse tree produced by RedshiftParser#func_expr.
	VisitFunc_expr(ctx *Func_exprContext) interface{}

	// Visit a parse tree produced by RedshiftParser#func_expr_windowless.
	VisitFunc_expr_windowless(ctx *Func_expr_windowlessContext) interface{}

	// Visit a parse tree produced by RedshiftParser#func_expr_common_subexpr.
	VisitFunc_expr_common_subexpr(ctx *Func_expr_common_subexprContext) interface{}

	// Visit a parse tree produced by RedshiftParser#xml_root_version.
	VisitXml_root_version(ctx *Xml_root_versionContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_xml_root_standalone.
	VisitOpt_xml_root_standalone(ctx *Opt_xml_root_standaloneContext) interface{}

	// Visit a parse tree produced by RedshiftParser#xml_attributes.
	VisitXml_attributes(ctx *Xml_attributesContext) interface{}

	// Visit a parse tree produced by RedshiftParser#xml_attribute_list.
	VisitXml_attribute_list(ctx *Xml_attribute_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#xml_attribute_el.
	VisitXml_attribute_el(ctx *Xml_attribute_elContext) interface{}

	// Visit a parse tree produced by RedshiftParser#document_or_content.
	VisitDocument_or_content(ctx *Document_or_contentContext) interface{}

	// Visit a parse tree produced by RedshiftParser#xml_whitespace_option.
	VisitXml_whitespace_option(ctx *Xml_whitespace_optionContext) interface{}

	// Visit a parse tree produced by RedshiftParser#xmlexists_argument.
	VisitXmlexists_argument(ctx *Xmlexists_argumentContext) interface{}

	// Visit a parse tree produced by RedshiftParser#xml_passing_mech.
	VisitXml_passing_mech(ctx *Xml_passing_mechContext) interface{}

	// Visit a parse tree produced by RedshiftParser#within_group_clause.
	VisitWithin_group_clause(ctx *Within_group_clauseContext) interface{}

	// Visit a parse tree produced by RedshiftParser#filter_clause.
	VisitFilter_clause(ctx *Filter_clauseContext) interface{}

	// Visit a parse tree produced by RedshiftParser#window_clause.
	VisitWindow_clause(ctx *Window_clauseContext) interface{}

	// Visit a parse tree produced by RedshiftParser#window_definition_list.
	VisitWindow_definition_list(ctx *Window_definition_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#window_definition.
	VisitWindow_definition(ctx *Window_definitionContext) interface{}

	// Visit a parse tree produced by RedshiftParser#over_clause.
	VisitOver_clause(ctx *Over_clauseContext) interface{}

	// Visit a parse tree produced by RedshiftParser#window_specification.
	VisitWindow_specification(ctx *Window_specificationContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_existing_window_name.
	VisitOpt_existing_window_name(ctx *Opt_existing_window_nameContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_partition_clause.
	VisitOpt_partition_clause(ctx *Opt_partition_clauseContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_frame_clause.
	VisitOpt_frame_clause(ctx *Opt_frame_clauseContext) interface{}

	// Visit a parse tree produced by RedshiftParser#frame_extent.
	VisitFrame_extent(ctx *Frame_extentContext) interface{}

	// Visit a parse tree produced by RedshiftParser#frame_bound.
	VisitFrame_bound(ctx *Frame_boundContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_window_exclusion_clause.
	VisitOpt_window_exclusion_clause(ctx *Opt_window_exclusion_clauseContext) interface{}

	// Visit a parse tree produced by RedshiftParser#row.
	VisitRow(ctx *RowContext) interface{}

	// Visit a parse tree produced by RedshiftParser#explicit_row.
	VisitExplicit_row(ctx *Explicit_rowContext) interface{}

	// Visit a parse tree produced by RedshiftParser#implicit_row.
	VisitImplicit_row(ctx *Implicit_rowContext) interface{}

	// Visit a parse tree produced by RedshiftParser#sub_type.
	VisitSub_type(ctx *Sub_typeContext) interface{}

	// Visit a parse tree produced by RedshiftParser#all_op.
	VisitAll_op(ctx *All_opContext) interface{}

	// Visit a parse tree produced by RedshiftParser#mathop.
	VisitMathop(ctx *MathopContext) interface{}

	// Visit a parse tree produced by RedshiftParser#qual_op.
	VisitQual_op(ctx *Qual_opContext) interface{}

	// Visit a parse tree produced by RedshiftParser#qual_all_op.
	VisitQual_all_op(ctx *Qual_all_opContext) interface{}

	// Visit a parse tree produced by RedshiftParser#subquery_Op.
	VisitSubquery_Op(ctx *Subquery_OpContext) interface{}

	// Visit a parse tree produced by RedshiftParser#expr_list.
	VisitExpr_list(ctx *Expr_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#func_arg_list.
	VisitFunc_arg_list(ctx *Func_arg_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#func_arg_expr.
	VisitFunc_arg_expr(ctx *Func_arg_exprContext) interface{}

	// Visit a parse tree produced by RedshiftParser#type_list.
	VisitType_list(ctx *Type_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#array_expr.
	VisitArray_expr(ctx *Array_exprContext) interface{}

	// Visit a parse tree produced by RedshiftParser#array_expr_list.
	VisitArray_expr_list(ctx *Array_expr_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#extract_list.
	VisitExtract_list(ctx *Extract_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#extract_arg.
	VisitExtract_arg(ctx *Extract_argContext) interface{}

	// Visit a parse tree produced by RedshiftParser#unicode_normal_form.
	VisitUnicode_normal_form(ctx *Unicode_normal_formContext) interface{}

	// Visit a parse tree produced by RedshiftParser#overlay_list.
	VisitOverlay_list(ctx *Overlay_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#position_list.
	VisitPosition_list(ctx *Position_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#substr_list.
	VisitSubstr_list(ctx *Substr_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#trim_list.
	VisitTrim_list(ctx *Trim_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#in_expr_select.
	VisitIn_expr_select(ctx *In_expr_selectContext) interface{}

	// Visit a parse tree produced by RedshiftParser#in_expr_list.
	VisitIn_expr_list(ctx *In_expr_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#case_expr.
	VisitCase_expr(ctx *Case_exprContext) interface{}

	// Visit a parse tree produced by RedshiftParser#when_clause_list.
	VisitWhen_clause_list(ctx *When_clause_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#when_clause.
	VisitWhen_clause(ctx *When_clauseContext) interface{}

	// Visit a parse tree produced by RedshiftParser#case_default.
	VisitCase_default(ctx *Case_defaultContext) interface{}

	// Visit a parse tree produced by RedshiftParser#case_arg.
	VisitCase_arg(ctx *Case_argContext) interface{}

	// Visit a parse tree produced by RedshiftParser#columnref.
	VisitColumnref(ctx *ColumnrefContext) interface{}

	// Visit a parse tree produced by RedshiftParser#indirection_el.
	VisitIndirection_el(ctx *Indirection_elContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_slice_bound.
	VisitOpt_slice_bound(ctx *Opt_slice_boundContext) interface{}

	// Visit a parse tree produced by RedshiftParser#indirection.
	VisitIndirection(ctx *IndirectionContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_indirection.
	VisitOpt_indirection(ctx *Opt_indirectionContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_target_list.
	VisitOpt_target_list(ctx *Opt_target_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#target_list.
	VisitTarget_list(ctx *Target_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#target_label.
	VisitTarget_label(ctx *Target_labelContext) interface{}

	// Visit a parse tree produced by RedshiftParser#target_star.
	VisitTarget_star(ctx *Target_starContext) interface{}

	// Visit a parse tree produced by RedshiftParser#target_alias.
	VisitTarget_alias(ctx *Target_aliasContext) interface{}

	// Visit a parse tree produced by RedshiftParser#qualified_name_list.
	VisitQualified_name_list(ctx *Qualified_name_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#qualified_name.
	VisitQualified_name(ctx *Qualified_nameContext) interface{}

	// Visit a parse tree produced by RedshiftParser#name_list.
	VisitName_list(ctx *Name_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#name.
	VisitName(ctx *NameContext) interface{}

	// Visit a parse tree produced by RedshiftParser#attr_name.
	VisitAttr_name(ctx *Attr_nameContext) interface{}

	// Visit a parse tree produced by RedshiftParser#file_name.
	VisitFile_name(ctx *File_nameContext) interface{}

	// Visit a parse tree produced by RedshiftParser#func_name.
	VisitFunc_name(ctx *Func_nameContext) interface{}

	// Visit a parse tree produced by RedshiftParser#aexprconst.
	VisitAexprconst(ctx *AexprconstContext) interface{}

	// Visit a parse tree produced by RedshiftParser#xconst.
	VisitXconst(ctx *XconstContext) interface{}

	// Visit a parse tree produced by RedshiftParser#bconst.
	VisitBconst(ctx *BconstContext) interface{}

	// Visit a parse tree produced by RedshiftParser#fconst.
	VisitFconst(ctx *FconstContext) interface{}

	// Visit a parse tree produced by RedshiftParser#iconst.
	VisitIconst(ctx *IconstContext) interface{}

	// Visit a parse tree produced by RedshiftParser#sconst.
	VisitSconst(ctx *SconstContext) interface{}

	// Visit a parse tree produced by RedshiftParser#anysconst.
	VisitAnysconst(ctx *AnysconstContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_uescape.
	VisitOpt_uescape(ctx *Opt_uescapeContext) interface{}

	// Visit a parse tree produced by RedshiftParser#signediconst.
	VisitSignediconst(ctx *SignediconstContext) interface{}

	// Visit a parse tree produced by RedshiftParser#roleid.
	VisitRoleid(ctx *RoleidContext) interface{}

	// Visit a parse tree produced by RedshiftParser#rolespec.
	VisitRolespec(ctx *RolespecContext) interface{}

	// Visit a parse tree produced by RedshiftParser#role_list.
	VisitRole_list(ctx *Role_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#colid.
	VisitColid(ctx *ColidContext) interface{}

	// Visit a parse tree produced by RedshiftParser#table_alias.
	VisitTable_alias(ctx *Table_aliasContext) interface{}

	// Visit a parse tree produced by RedshiftParser#type_function_name.
	VisitType_function_name(ctx *Type_function_nameContext) interface{}

	// Visit a parse tree produced by RedshiftParser#nonreservedword.
	VisitNonreservedword(ctx *NonreservedwordContext) interface{}

	// Visit a parse tree produced by RedshiftParser#collabel.
	VisitCollabel(ctx *CollabelContext) interface{}

	// Visit a parse tree produced by RedshiftParser#identifier.
	VisitIdentifier(ctx *IdentifierContext) interface{}

	// Visit a parse tree produced by RedshiftParser#plsqlidentifier.
	VisitPlsqlidentifier(ctx *PlsqlidentifierContext) interface{}

	// Visit a parse tree produced by RedshiftParser#unreserved_keyword.
	VisitUnreserved_keyword(ctx *Unreserved_keywordContext) interface{}

	// Visit a parse tree produced by RedshiftParser#col_name_keyword.
	VisitCol_name_keyword(ctx *Col_name_keywordContext) interface{}

	// Visit a parse tree produced by RedshiftParser#type_func_name_keyword.
	VisitType_func_name_keyword(ctx *Type_func_name_keywordContext) interface{}

	// Visit a parse tree produced by RedshiftParser#reserved_keyword.
	VisitReserved_keyword(ctx *Reserved_keywordContext) interface{}

	// Visit a parse tree produced by RedshiftParser#builtin_function_name.
	VisitBuiltin_function_name(ctx *Builtin_function_nameContext) interface{}

	// Visit a parse tree produced by RedshiftParser#pl_function.
	VisitPl_function(ctx *Pl_functionContext) interface{}

	// Visit a parse tree produced by RedshiftParser#comp_options.
	VisitComp_options(ctx *Comp_optionsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#comp_option.
	VisitComp_option(ctx *Comp_optionContext) interface{}

	// Visit a parse tree produced by RedshiftParser#sharp.
	VisitSharp(ctx *SharpContext) interface{}

	// Visit a parse tree produced by RedshiftParser#option_value.
	VisitOption_value(ctx *Option_valueContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_semi.
	VisitOpt_semi(ctx *Opt_semiContext) interface{}

	// Visit a parse tree produced by RedshiftParser#pl_block.
	VisitPl_block(ctx *Pl_blockContext) interface{}

	// Visit a parse tree produced by RedshiftParser#decl_sect.
	VisitDecl_sect(ctx *Decl_sectContext) interface{}

	// Visit a parse tree produced by RedshiftParser#decl_start.
	VisitDecl_start(ctx *Decl_startContext) interface{}

	// Visit a parse tree produced by RedshiftParser#decl_stmts.
	VisitDecl_stmts(ctx *Decl_stmtsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#label_decl.
	VisitLabel_decl(ctx *Label_declContext) interface{}

	// Visit a parse tree produced by RedshiftParser#decl_stmt.
	VisitDecl_stmt(ctx *Decl_stmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#decl_statement.
	VisitDecl_statement(ctx *Decl_statementContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_scrollable.
	VisitOpt_scrollable(ctx *Opt_scrollableContext) interface{}

	// Visit a parse tree produced by RedshiftParser#decl_cursor_query.
	VisitDecl_cursor_query(ctx *Decl_cursor_queryContext) interface{}

	// Visit a parse tree produced by RedshiftParser#decl_cursor_args.
	VisitDecl_cursor_args(ctx *Decl_cursor_argsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#decl_cursor_arglist.
	VisitDecl_cursor_arglist(ctx *Decl_cursor_arglistContext) interface{}

	// Visit a parse tree produced by RedshiftParser#decl_cursor_arg.
	VisitDecl_cursor_arg(ctx *Decl_cursor_argContext) interface{}

	// Visit a parse tree produced by RedshiftParser#decl_is_for.
	VisitDecl_is_for(ctx *Decl_is_forContext) interface{}

	// Visit a parse tree produced by RedshiftParser#decl_aliasitem.
	VisitDecl_aliasitem(ctx *Decl_aliasitemContext) interface{}

	// Visit a parse tree produced by RedshiftParser#decl_varname.
	VisitDecl_varname(ctx *Decl_varnameContext) interface{}

	// Visit a parse tree produced by RedshiftParser#decl_const.
	VisitDecl_const(ctx *Decl_constContext) interface{}

	// Visit a parse tree produced by RedshiftParser#decl_datatype.
	VisitDecl_datatype(ctx *Decl_datatypeContext) interface{}

	// Visit a parse tree produced by RedshiftParser#decl_collate.
	VisitDecl_collate(ctx *Decl_collateContext) interface{}

	// Visit a parse tree produced by RedshiftParser#decl_notnull.
	VisitDecl_notnull(ctx *Decl_notnullContext) interface{}

	// Visit a parse tree produced by RedshiftParser#decl_defval.
	VisitDecl_defval(ctx *Decl_defvalContext) interface{}

	// Visit a parse tree produced by RedshiftParser#decl_defkey.
	VisitDecl_defkey(ctx *Decl_defkeyContext) interface{}

	// Visit a parse tree produced by RedshiftParser#assign_operator.
	VisitAssign_operator(ctx *Assign_operatorContext) interface{}

	// Visit a parse tree produced by RedshiftParser#proc_sect.
	VisitProc_sect(ctx *Proc_sectContext) interface{}

	// Visit a parse tree produced by RedshiftParser#proc_stmt.
	VisitProc_stmt(ctx *Proc_stmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#stmt_perform.
	VisitStmt_perform(ctx *Stmt_performContext) interface{}

	// Visit a parse tree produced by RedshiftParser#stmt_call.
	VisitStmt_call(ctx *Stmt_callContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_expr_list.
	VisitOpt_expr_list(ctx *Opt_expr_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#stmt_assign.
	VisitStmt_assign(ctx *Stmt_assignContext) interface{}

	// Visit a parse tree produced by RedshiftParser#stmt_getdiag.
	VisitStmt_getdiag(ctx *Stmt_getdiagContext) interface{}

	// Visit a parse tree produced by RedshiftParser#getdiag_area_opt.
	VisitGetdiag_area_opt(ctx *Getdiag_area_optContext) interface{}

	// Visit a parse tree produced by RedshiftParser#getdiag_list.
	VisitGetdiag_list(ctx *Getdiag_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#getdiag_list_item.
	VisitGetdiag_list_item(ctx *Getdiag_list_itemContext) interface{}

	// Visit a parse tree produced by RedshiftParser#getdiag_item.
	VisitGetdiag_item(ctx *Getdiag_itemContext) interface{}

	// Visit a parse tree produced by RedshiftParser#getdiag_target.
	VisitGetdiag_target(ctx *Getdiag_targetContext) interface{}

	// Visit a parse tree produced by RedshiftParser#assign_var.
	VisitAssign_var(ctx *Assign_varContext) interface{}

	// Visit a parse tree produced by RedshiftParser#stmt_if.
	VisitStmt_if(ctx *Stmt_ifContext) interface{}

	// Visit a parse tree produced by RedshiftParser#stmt_elsifs.
	VisitStmt_elsifs(ctx *Stmt_elsifsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#stmt_else.
	VisitStmt_else(ctx *Stmt_elseContext) interface{}

	// Visit a parse tree produced by RedshiftParser#stmt_case.
	VisitStmt_case(ctx *Stmt_caseContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_expr_until_when.
	VisitOpt_expr_until_when(ctx *Opt_expr_until_whenContext) interface{}

	// Visit a parse tree produced by RedshiftParser#case_when_list.
	VisitCase_when_list(ctx *Case_when_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#case_when.
	VisitCase_when(ctx *Case_whenContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_case_else.
	VisitOpt_case_else(ctx *Opt_case_elseContext) interface{}

	// Visit a parse tree produced by RedshiftParser#stmt_loop.
	VisitStmt_loop(ctx *Stmt_loopContext) interface{}

	// Visit a parse tree produced by RedshiftParser#stmt_while.
	VisitStmt_while(ctx *Stmt_whileContext) interface{}

	// Visit a parse tree produced by RedshiftParser#stmt_for.
	VisitStmt_for(ctx *Stmt_forContext) interface{}

	// Visit a parse tree produced by RedshiftParser#for_control.
	VisitFor_control(ctx *For_controlContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_for_using_expression.
	VisitOpt_for_using_expression(ctx *Opt_for_using_expressionContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_cursor_parameters.
	VisitOpt_cursor_parameters(ctx *Opt_cursor_parametersContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_reverse.
	VisitOpt_reverse(ctx *Opt_reverseContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_by_expression.
	VisitOpt_by_expression(ctx *Opt_by_expressionContext) interface{}

	// Visit a parse tree produced by RedshiftParser#for_variable.
	VisitFor_variable(ctx *For_variableContext) interface{}

	// Visit a parse tree produced by RedshiftParser#stmt_foreach_a.
	VisitStmt_foreach_a(ctx *Stmt_foreach_aContext) interface{}

	// Visit a parse tree produced by RedshiftParser#foreach_slice.
	VisitForeach_slice(ctx *Foreach_sliceContext) interface{}

	// Visit a parse tree produced by RedshiftParser#stmt_exit.
	VisitStmt_exit(ctx *Stmt_exitContext) interface{}

	// Visit a parse tree produced by RedshiftParser#exit_type.
	VisitExit_type(ctx *Exit_typeContext) interface{}

	// Visit a parse tree produced by RedshiftParser#stmt_return.
	VisitStmt_return(ctx *Stmt_returnContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_return_result.
	VisitOpt_return_result(ctx *Opt_return_resultContext) interface{}

	// Visit a parse tree produced by RedshiftParser#stmt_raise.
	VisitStmt_raise(ctx *Stmt_raiseContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_stmt_raise_level.
	VisitOpt_stmt_raise_level(ctx *Opt_stmt_raise_levelContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_raise_list.
	VisitOpt_raise_list(ctx *Opt_raise_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_raise_using.
	VisitOpt_raise_using(ctx *Opt_raise_usingContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_raise_using_elem.
	VisitOpt_raise_using_elem(ctx *Opt_raise_using_elemContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_raise_using_elem_list.
	VisitOpt_raise_using_elem_list(ctx *Opt_raise_using_elem_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#stmt_assert.
	VisitStmt_assert(ctx *Stmt_assertContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_stmt_assert_message.
	VisitOpt_stmt_assert_message(ctx *Opt_stmt_assert_messageContext) interface{}

	// Visit a parse tree produced by RedshiftParser#loop_body.
	VisitLoop_body(ctx *Loop_bodyContext) interface{}

	// Visit a parse tree produced by RedshiftParser#stmt_execsql.
	VisitStmt_execsql(ctx *Stmt_execsqlContext) interface{}

	// Visit a parse tree produced by RedshiftParser#stmt_dynexecute.
	VisitStmt_dynexecute(ctx *Stmt_dynexecuteContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_execute_using.
	VisitOpt_execute_using(ctx *Opt_execute_usingContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_execute_using_list.
	VisitOpt_execute_using_list(ctx *Opt_execute_using_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_execute_into.
	VisitOpt_execute_into(ctx *Opt_execute_intoContext) interface{}

	// Visit a parse tree produced by RedshiftParser#stmt_open.
	VisitStmt_open(ctx *Stmt_openContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_open_bound_list_item.
	VisitOpt_open_bound_list_item(ctx *Opt_open_bound_list_itemContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_open_bound_list.
	VisitOpt_open_bound_list(ctx *Opt_open_bound_listContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_open_using.
	VisitOpt_open_using(ctx *Opt_open_usingContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_scroll_option.
	VisitOpt_scroll_option(ctx *Opt_scroll_optionContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_scroll_option_no.
	VisitOpt_scroll_option_no(ctx *Opt_scroll_option_noContext) interface{}

	// Visit a parse tree produced by RedshiftParser#stmt_fetch.
	VisitStmt_fetch(ctx *Stmt_fetchContext) interface{}

	// Visit a parse tree produced by RedshiftParser#into_target.
	VisitInto_target(ctx *Into_targetContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_cursor_from.
	VisitOpt_cursor_from(ctx *Opt_cursor_fromContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_fetch_direction.
	VisitOpt_fetch_direction(ctx *Opt_fetch_directionContext) interface{}

	// Visit a parse tree produced by RedshiftParser#stmt_move.
	VisitStmt_move(ctx *Stmt_moveContext) interface{}

	// Visit a parse tree produced by RedshiftParser#stmt_close.
	VisitStmt_close(ctx *Stmt_closeContext) interface{}

	// Visit a parse tree produced by RedshiftParser#stmt_null.
	VisitStmt_null(ctx *Stmt_nullContext) interface{}

	// Visit a parse tree produced by RedshiftParser#stmt_commit.
	VisitStmt_commit(ctx *Stmt_commitContext) interface{}

	// Visit a parse tree produced by RedshiftParser#stmt_rollback.
	VisitStmt_rollback(ctx *Stmt_rollbackContext) interface{}

	// Visit a parse tree produced by RedshiftParser#plsql_opt_transaction_chain.
	VisitPlsql_opt_transaction_chain(ctx *Plsql_opt_transaction_chainContext) interface{}

	// Visit a parse tree produced by RedshiftParser#stmt_set.
	VisitStmt_set(ctx *Stmt_setContext) interface{}

	// Visit a parse tree produced by RedshiftParser#cursor_variable.
	VisitCursor_variable(ctx *Cursor_variableContext) interface{}

	// Visit a parse tree produced by RedshiftParser#exception_sect.
	VisitException_sect(ctx *Exception_sectContext) interface{}

	// Visit a parse tree produced by RedshiftParser#proc_exceptions.
	VisitProc_exceptions(ctx *Proc_exceptionsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#proc_exception.
	VisitProc_exception(ctx *Proc_exceptionContext) interface{}

	// Visit a parse tree produced by RedshiftParser#proc_conditions.
	VisitProc_conditions(ctx *Proc_conditionsContext) interface{}

	// Visit a parse tree produced by RedshiftParser#proc_condition.
	VisitProc_condition(ctx *Proc_conditionContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_block_label.
	VisitOpt_block_label(ctx *Opt_block_labelContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_loop_label.
	VisitOpt_loop_label(ctx *Opt_loop_labelContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_label.
	VisitOpt_label(ctx *Opt_labelContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_exitcond.
	VisitOpt_exitcond(ctx *Opt_exitcondContext) interface{}

	// Visit a parse tree produced by RedshiftParser#any_identifier.
	VisitAny_identifier(ctx *Any_identifierContext) interface{}

	// Visit a parse tree produced by RedshiftParser#plsql_unreserved_keyword.
	VisitPlsql_unreserved_keyword(ctx *Plsql_unreserved_keywordContext) interface{}

	// Visit a parse tree produced by RedshiftParser#sql_expression.
	VisitSql_expression(ctx *Sql_expressionContext) interface{}

	// Visit a parse tree produced by RedshiftParser#expr_until_then.
	VisitExpr_until_then(ctx *Expr_until_thenContext) interface{}

	// Visit a parse tree produced by RedshiftParser#expr_until_semi.
	VisitExpr_until_semi(ctx *Expr_until_semiContext) interface{}

	// Visit a parse tree produced by RedshiftParser#expr_until_rightbracket.
	VisitExpr_until_rightbracket(ctx *Expr_until_rightbracketContext) interface{}

	// Visit a parse tree produced by RedshiftParser#expr_until_loop.
	VisitExpr_until_loop(ctx *Expr_until_loopContext) interface{}

	// Visit a parse tree produced by RedshiftParser#make_execsql_stmt.
	VisitMake_execsql_stmt(ctx *Make_execsql_stmtContext) interface{}

	// Visit a parse tree produced by RedshiftParser#opt_returning_clause_into.
	VisitOpt_returning_clause_into(ctx *Opt_returning_clause_intoContext) interface{}
}
