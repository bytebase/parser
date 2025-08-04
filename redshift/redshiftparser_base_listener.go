// Code generated from RedshiftParser.g4 by ANTLR 4.13.2. DO NOT EDIT.

package redshift // RedshiftParser
import "github.com/antlr4-go/antlr/v4"

// BaseRedshiftParserListener is a complete listener for a parse tree produced by RedshiftParser.
type BaseRedshiftParserListener struct{}

var _ RedshiftParserListener = &BaseRedshiftParserListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseRedshiftParserListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseRedshiftParserListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseRedshiftParserListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseRedshiftParserListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterRoot is called when production root is entered.
func (s *BaseRedshiftParserListener) EnterRoot(ctx *RootContext) {}

// ExitRoot is called when production root is exited.
func (s *BaseRedshiftParserListener) ExitRoot(ctx *RootContext) {}

// EnterPlsqlroot is called when production plsqlroot is entered.
func (s *BaseRedshiftParserListener) EnterPlsqlroot(ctx *PlsqlrootContext) {}

// ExitPlsqlroot is called when production plsqlroot is exited.
func (s *BaseRedshiftParserListener) ExitPlsqlroot(ctx *PlsqlrootContext) {}

// EnterStmtblock is called when production stmtblock is entered.
func (s *BaseRedshiftParserListener) EnterStmtblock(ctx *StmtblockContext) {}

// ExitStmtblock is called when production stmtblock is exited.
func (s *BaseRedshiftParserListener) ExitStmtblock(ctx *StmtblockContext) {}

// EnterStmtmulti is called when production stmtmulti is entered.
func (s *BaseRedshiftParserListener) EnterStmtmulti(ctx *StmtmultiContext) {}

// ExitStmtmulti is called when production stmtmulti is exited.
func (s *BaseRedshiftParserListener) ExitStmtmulti(ctx *StmtmultiContext) {}

// EnterStmt is called when production stmt is entered.
func (s *BaseRedshiftParserListener) EnterStmt(ctx *StmtContext) {}

// ExitStmt is called when production stmt is exited.
func (s *BaseRedshiftParserListener) ExitStmt(ctx *StmtContext) {}

// EnterPlsqlconsolecommand is called when production plsqlconsolecommand is entered.
func (s *BaseRedshiftParserListener) EnterPlsqlconsolecommand(ctx *PlsqlconsolecommandContext) {}

// ExitPlsqlconsolecommand is called when production plsqlconsolecommand is exited.
func (s *BaseRedshiftParserListener) ExitPlsqlconsolecommand(ctx *PlsqlconsolecommandContext) {}

// EnterCallstmt is called when production callstmt is entered.
func (s *BaseRedshiftParserListener) EnterCallstmt(ctx *CallstmtContext) {}

// ExitCallstmt is called when production callstmt is exited.
func (s *BaseRedshiftParserListener) ExitCallstmt(ctx *CallstmtContext) {}

// EnterCreaterolestmt is called when production createrolestmt is entered.
func (s *BaseRedshiftParserListener) EnterCreaterolestmt(ctx *CreaterolestmtContext) {}

// ExitCreaterolestmt is called when production createrolestmt is exited.
func (s *BaseRedshiftParserListener) ExitCreaterolestmt(ctx *CreaterolestmtContext) {}

// EnterOpt_with is called when production opt_with is entered.
func (s *BaseRedshiftParserListener) EnterOpt_with(ctx *Opt_withContext) {}

// ExitOpt_with is called when production opt_with is exited.
func (s *BaseRedshiftParserListener) ExitOpt_with(ctx *Opt_withContext) {}

// EnterOptrolelist is called when production optrolelist is entered.
func (s *BaseRedshiftParserListener) EnterOptrolelist(ctx *OptrolelistContext) {}

// ExitOptrolelist is called when production optrolelist is exited.
func (s *BaseRedshiftParserListener) ExitOptrolelist(ctx *OptrolelistContext) {}

// EnterAlteroptrolelist is called when production alteroptrolelist is entered.
func (s *BaseRedshiftParserListener) EnterAlteroptrolelist(ctx *AlteroptrolelistContext) {}

// ExitAlteroptrolelist is called when production alteroptrolelist is exited.
func (s *BaseRedshiftParserListener) ExitAlteroptrolelist(ctx *AlteroptrolelistContext) {}

// EnterAlteroptroleelem is called when production alteroptroleelem is entered.
func (s *BaseRedshiftParserListener) EnterAlteroptroleelem(ctx *AlteroptroleelemContext) {}

// ExitAlteroptroleelem is called when production alteroptroleelem is exited.
func (s *BaseRedshiftParserListener) ExitAlteroptroleelem(ctx *AlteroptroleelemContext) {}

// EnterCreateoptroleelem is called when production createoptroleelem is entered.
func (s *BaseRedshiftParserListener) EnterCreateoptroleelem(ctx *CreateoptroleelemContext) {}

// ExitCreateoptroleelem is called when production createoptroleelem is exited.
func (s *BaseRedshiftParserListener) ExitCreateoptroleelem(ctx *CreateoptroleelemContext) {}

// EnterCreateuserstmt is called when production createuserstmt is entered.
func (s *BaseRedshiftParserListener) EnterCreateuserstmt(ctx *CreateuserstmtContext) {}

// ExitCreateuserstmt is called when production createuserstmt is exited.
func (s *BaseRedshiftParserListener) ExitCreateuserstmt(ctx *CreateuserstmtContext) {}

// EnterAlterrolestmt is called when production alterrolestmt is entered.
func (s *BaseRedshiftParserListener) EnterAlterrolestmt(ctx *AlterrolestmtContext) {}

// ExitAlterrolestmt is called when production alterrolestmt is exited.
func (s *BaseRedshiftParserListener) ExitAlterrolestmt(ctx *AlterrolestmtContext) {}

// EnterAlterroleaction is called when production alterroleaction is entered.
func (s *BaseRedshiftParserListener) EnterAlterroleaction(ctx *AlterroleactionContext) {}

// ExitAlterroleaction is called when production alterroleaction is exited.
func (s *BaseRedshiftParserListener) ExitAlterroleaction(ctx *AlterroleactionContext) {}

// EnterOpt_in_database is called when production opt_in_database is entered.
func (s *BaseRedshiftParserListener) EnterOpt_in_database(ctx *Opt_in_databaseContext) {}

// ExitOpt_in_database is called when production opt_in_database is exited.
func (s *BaseRedshiftParserListener) ExitOpt_in_database(ctx *Opt_in_databaseContext) {}

// EnterAlterrolesetstmt is called when production alterrolesetstmt is entered.
func (s *BaseRedshiftParserListener) EnterAlterrolesetstmt(ctx *AlterrolesetstmtContext) {}

// ExitAlterrolesetstmt is called when production alterrolesetstmt is exited.
func (s *BaseRedshiftParserListener) ExitAlterrolesetstmt(ctx *AlterrolesetstmtContext) {}

// EnterAlterschemastmt is called when production alterschemastmt is entered.
func (s *BaseRedshiftParserListener) EnterAlterschemastmt(ctx *AlterschemastmtContext) {}

// ExitAlterschemastmt is called when production alterschemastmt is exited.
func (s *BaseRedshiftParserListener) ExitAlterschemastmt(ctx *AlterschemastmtContext) {}

// EnterDroprolestmt is called when production droprolestmt is entered.
func (s *BaseRedshiftParserListener) EnterDroprolestmt(ctx *DroprolestmtContext) {}

// ExitDroprolestmt is called when production droprolestmt is exited.
func (s *BaseRedshiftParserListener) ExitDroprolestmt(ctx *DroprolestmtContext) {}

// EnterDropuserstmt is called when production dropuserstmt is entered.
func (s *BaseRedshiftParserListener) EnterDropuserstmt(ctx *DropuserstmtContext) {}

// ExitDropuserstmt is called when production dropuserstmt is exited.
func (s *BaseRedshiftParserListener) ExitDropuserstmt(ctx *DropuserstmtContext) {}

// EnterDropgroupstmt is called when production dropgroupstmt is entered.
func (s *BaseRedshiftParserListener) EnterDropgroupstmt(ctx *DropgroupstmtContext) {}

// ExitDropgroupstmt is called when production dropgroupstmt is exited.
func (s *BaseRedshiftParserListener) ExitDropgroupstmt(ctx *DropgroupstmtContext) {}

// EnterCreategroupstmt is called when production creategroupstmt is entered.
func (s *BaseRedshiftParserListener) EnterCreategroupstmt(ctx *CreategroupstmtContext) {}

// ExitCreategroupstmt is called when production creategroupstmt is exited.
func (s *BaseRedshiftParserListener) ExitCreategroupstmt(ctx *CreategroupstmtContext) {}

// EnterAltergroupstmt is called when production altergroupstmt is entered.
func (s *BaseRedshiftParserListener) EnterAltergroupstmt(ctx *AltergroupstmtContext) {}

// ExitAltergroupstmt is called when production altergroupstmt is exited.
func (s *BaseRedshiftParserListener) ExitAltergroupstmt(ctx *AltergroupstmtContext) {}

// EnterAdd_drop is called when production add_drop is entered.
func (s *BaseRedshiftParserListener) EnterAdd_drop(ctx *Add_dropContext) {}

// ExitAdd_drop is called when production add_drop is exited.
func (s *BaseRedshiftParserListener) ExitAdd_drop(ctx *Add_dropContext) {}

// EnterCreateschemastmt is called when production createschemastmt is entered.
func (s *BaseRedshiftParserListener) EnterCreateschemastmt(ctx *CreateschemastmtContext) {}

// ExitCreateschemastmt is called when production createschemastmt is exited.
func (s *BaseRedshiftParserListener) ExitCreateschemastmt(ctx *CreateschemastmtContext) {}

// EnterOpt_auth_clause is called when production opt_auth_clause is entered.
func (s *BaseRedshiftParserListener) EnterOpt_auth_clause(ctx *Opt_auth_clauseContext) {}

// ExitOpt_auth_clause is called when production opt_auth_clause is exited.
func (s *BaseRedshiftParserListener) ExitOpt_auth_clause(ctx *Opt_auth_clauseContext) {}

// EnterOpt_quota is called when production opt_quota is entered.
func (s *BaseRedshiftParserListener) EnterOpt_quota(ctx *Opt_quotaContext) {}

// ExitOpt_quota is called when production opt_quota is exited.
func (s *BaseRedshiftParserListener) ExitOpt_quota(ctx *Opt_quotaContext) {}

// EnterOptschemaeltlist is called when production optschemaeltlist is entered.
func (s *BaseRedshiftParserListener) EnterOptschemaeltlist(ctx *OptschemaeltlistContext) {}

// ExitOptschemaeltlist is called when production optschemaeltlist is exited.
func (s *BaseRedshiftParserListener) ExitOptschemaeltlist(ctx *OptschemaeltlistContext) {}

// EnterSchema_stmt is called when production schema_stmt is entered.
func (s *BaseRedshiftParserListener) EnterSchema_stmt(ctx *Schema_stmtContext) {}

// ExitSchema_stmt is called when production schema_stmt is exited.
func (s *BaseRedshiftParserListener) ExitSchema_stmt(ctx *Schema_stmtContext) {}

// EnterVariablesetstmt is called when production variablesetstmt is entered.
func (s *BaseRedshiftParserListener) EnterVariablesetstmt(ctx *VariablesetstmtContext) {}

// ExitVariablesetstmt is called when production variablesetstmt is exited.
func (s *BaseRedshiftParserListener) ExitVariablesetstmt(ctx *VariablesetstmtContext) {}

// EnterSet_rest is called when production set_rest is entered.
func (s *BaseRedshiftParserListener) EnterSet_rest(ctx *Set_restContext) {}

// ExitSet_rest is called when production set_rest is exited.
func (s *BaseRedshiftParserListener) ExitSet_rest(ctx *Set_restContext) {}

// EnterGeneric_set is called when production generic_set is entered.
func (s *BaseRedshiftParserListener) EnterGeneric_set(ctx *Generic_setContext) {}

// ExitGeneric_set is called when production generic_set is exited.
func (s *BaseRedshiftParserListener) ExitGeneric_set(ctx *Generic_setContext) {}

// EnterSet_rest_more is called when production set_rest_more is entered.
func (s *BaseRedshiftParserListener) EnterSet_rest_more(ctx *Set_rest_moreContext) {}

// ExitSet_rest_more is called when production set_rest_more is exited.
func (s *BaseRedshiftParserListener) ExitSet_rest_more(ctx *Set_rest_moreContext) {}

// EnterVar_name is called when production var_name is entered.
func (s *BaseRedshiftParserListener) EnterVar_name(ctx *Var_nameContext) {}

// ExitVar_name is called when production var_name is exited.
func (s *BaseRedshiftParserListener) ExitVar_name(ctx *Var_nameContext) {}

// EnterVar_list is called when production var_list is entered.
func (s *BaseRedshiftParserListener) EnterVar_list(ctx *Var_listContext) {}

// ExitVar_list is called when production var_list is exited.
func (s *BaseRedshiftParserListener) ExitVar_list(ctx *Var_listContext) {}

// EnterVar_value is called when production var_value is entered.
func (s *BaseRedshiftParserListener) EnterVar_value(ctx *Var_valueContext) {}

// ExitVar_value is called when production var_value is exited.
func (s *BaseRedshiftParserListener) ExitVar_value(ctx *Var_valueContext) {}

// EnterIso_level is called when production iso_level is entered.
func (s *BaseRedshiftParserListener) EnterIso_level(ctx *Iso_levelContext) {}

// ExitIso_level is called when production iso_level is exited.
func (s *BaseRedshiftParserListener) ExitIso_level(ctx *Iso_levelContext) {}

// EnterOpt_boolean_or_string is called when production opt_boolean_or_string is entered.
func (s *BaseRedshiftParserListener) EnterOpt_boolean_or_string(ctx *Opt_boolean_or_stringContext) {}

// ExitOpt_boolean_or_string is called when production opt_boolean_or_string is exited.
func (s *BaseRedshiftParserListener) ExitOpt_boolean_or_string(ctx *Opt_boolean_or_stringContext) {}

// EnterZone_value is called when production zone_value is entered.
func (s *BaseRedshiftParserListener) EnterZone_value(ctx *Zone_valueContext) {}

// ExitZone_value is called when production zone_value is exited.
func (s *BaseRedshiftParserListener) ExitZone_value(ctx *Zone_valueContext) {}

// EnterOpt_encoding is called when production opt_encoding is entered.
func (s *BaseRedshiftParserListener) EnterOpt_encoding(ctx *Opt_encodingContext) {}

// ExitOpt_encoding is called when production opt_encoding is exited.
func (s *BaseRedshiftParserListener) ExitOpt_encoding(ctx *Opt_encodingContext) {}

// EnterNonreservedword_or_sconst is called when production nonreservedword_or_sconst is entered.
func (s *BaseRedshiftParserListener) EnterNonreservedword_or_sconst(ctx *Nonreservedword_or_sconstContext) {
}

// ExitNonreservedword_or_sconst is called when production nonreservedword_or_sconst is exited.
func (s *BaseRedshiftParserListener) ExitNonreservedword_or_sconst(ctx *Nonreservedword_or_sconstContext) {
}

// EnterVariableresetstmt is called when production variableresetstmt is entered.
func (s *BaseRedshiftParserListener) EnterVariableresetstmt(ctx *VariableresetstmtContext) {}

// ExitVariableresetstmt is called when production variableresetstmt is exited.
func (s *BaseRedshiftParserListener) ExitVariableresetstmt(ctx *VariableresetstmtContext) {}

// EnterReset_rest is called when production reset_rest is entered.
func (s *BaseRedshiftParserListener) EnterReset_rest(ctx *Reset_restContext) {}

// ExitReset_rest is called when production reset_rest is exited.
func (s *BaseRedshiftParserListener) ExitReset_rest(ctx *Reset_restContext) {}

// EnterGeneric_reset is called when production generic_reset is entered.
func (s *BaseRedshiftParserListener) EnterGeneric_reset(ctx *Generic_resetContext) {}

// ExitGeneric_reset is called when production generic_reset is exited.
func (s *BaseRedshiftParserListener) ExitGeneric_reset(ctx *Generic_resetContext) {}

// EnterSetresetclause is called when production setresetclause is entered.
func (s *BaseRedshiftParserListener) EnterSetresetclause(ctx *SetresetclauseContext) {}

// ExitSetresetclause is called when production setresetclause is exited.
func (s *BaseRedshiftParserListener) ExitSetresetclause(ctx *SetresetclauseContext) {}

// EnterFunctionsetresetclause is called when production functionsetresetclause is entered.
func (s *BaseRedshiftParserListener) EnterFunctionsetresetclause(ctx *FunctionsetresetclauseContext) {
}

// ExitFunctionsetresetclause is called when production functionsetresetclause is exited.
func (s *BaseRedshiftParserListener) ExitFunctionsetresetclause(ctx *FunctionsetresetclauseContext) {}

// EnterVariableshowstmt is called when production variableshowstmt is entered.
func (s *BaseRedshiftParserListener) EnterVariableshowstmt(ctx *VariableshowstmtContext) {}

// ExitVariableshowstmt is called when production variableshowstmt is exited.
func (s *BaseRedshiftParserListener) ExitVariableshowstmt(ctx *VariableshowstmtContext) {}

// EnterConstraintssetstmt is called when production constraintssetstmt is entered.
func (s *BaseRedshiftParserListener) EnterConstraintssetstmt(ctx *ConstraintssetstmtContext) {}

// ExitConstraintssetstmt is called when production constraintssetstmt is exited.
func (s *BaseRedshiftParserListener) ExitConstraintssetstmt(ctx *ConstraintssetstmtContext) {}

// EnterConstraints_set_list is called when production constraints_set_list is entered.
func (s *BaseRedshiftParserListener) EnterConstraints_set_list(ctx *Constraints_set_listContext) {}

// ExitConstraints_set_list is called when production constraints_set_list is exited.
func (s *BaseRedshiftParserListener) ExitConstraints_set_list(ctx *Constraints_set_listContext) {}

// EnterConstraints_set_mode is called when production constraints_set_mode is entered.
func (s *BaseRedshiftParserListener) EnterConstraints_set_mode(ctx *Constraints_set_modeContext) {}

// ExitConstraints_set_mode is called when production constraints_set_mode is exited.
func (s *BaseRedshiftParserListener) ExitConstraints_set_mode(ctx *Constraints_set_modeContext) {}

// EnterCheckpointstmt is called when production checkpointstmt is entered.
func (s *BaseRedshiftParserListener) EnterCheckpointstmt(ctx *CheckpointstmtContext) {}

// ExitCheckpointstmt is called when production checkpointstmt is exited.
func (s *BaseRedshiftParserListener) ExitCheckpointstmt(ctx *CheckpointstmtContext) {}

// EnterDiscardstmt is called when production discardstmt is entered.
func (s *BaseRedshiftParserListener) EnterDiscardstmt(ctx *DiscardstmtContext) {}

// ExitDiscardstmt is called when production discardstmt is exited.
func (s *BaseRedshiftParserListener) ExitDiscardstmt(ctx *DiscardstmtContext) {}

// EnterAltertablestmt is called when production altertablestmt is entered.
func (s *BaseRedshiftParserListener) EnterAltertablestmt(ctx *AltertablestmtContext) {}

// ExitAltertablestmt is called when production altertablestmt is exited.
func (s *BaseRedshiftParserListener) ExitAltertablestmt(ctx *AltertablestmtContext) {}

// EnterAlter_table_cmds is called when production alter_table_cmds is entered.
func (s *BaseRedshiftParserListener) EnterAlter_table_cmds(ctx *Alter_table_cmdsContext) {}

// ExitAlter_table_cmds is called when production alter_table_cmds is exited.
func (s *BaseRedshiftParserListener) ExitAlter_table_cmds(ctx *Alter_table_cmdsContext) {}

// EnterTable_constraint is called when production table_constraint is entered.
func (s *BaseRedshiftParserListener) EnterTable_constraint(ctx *Table_constraintContext) {}

// ExitTable_constraint is called when production table_constraint is exited.
func (s *BaseRedshiftParserListener) ExitTable_constraint(ctx *Table_constraintContext) {}

// EnterPartition_cmd is called when production partition_cmd is entered.
func (s *BaseRedshiftParserListener) EnterPartition_cmd(ctx *Partition_cmdContext) {}

// ExitPartition_cmd is called when production partition_cmd is exited.
func (s *BaseRedshiftParserListener) ExitPartition_cmd(ctx *Partition_cmdContext) {}

// EnterIndex_partition_cmd is called when production index_partition_cmd is entered.
func (s *BaseRedshiftParserListener) EnterIndex_partition_cmd(ctx *Index_partition_cmdContext) {}

// ExitIndex_partition_cmd is called when production index_partition_cmd is exited.
func (s *BaseRedshiftParserListener) ExitIndex_partition_cmd(ctx *Index_partition_cmdContext) {}

// EnterAlter_table_cmd is called when production alter_table_cmd is entered.
func (s *BaseRedshiftParserListener) EnterAlter_table_cmd(ctx *Alter_table_cmdContext) {}

// ExitAlter_table_cmd is called when production alter_table_cmd is exited.
func (s *BaseRedshiftParserListener) ExitAlter_table_cmd(ctx *Alter_table_cmdContext) {}

// EnterAlter_column_default is called when production alter_column_default is entered.
func (s *BaseRedshiftParserListener) EnterAlter_column_default(ctx *Alter_column_defaultContext) {}

// ExitAlter_column_default is called when production alter_column_default is exited.
func (s *BaseRedshiftParserListener) ExitAlter_column_default(ctx *Alter_column_defaultContext) {}

// EnterOpt_drop_behavior is called when production opt_drop_behavior is entered.
func (s *BaseRedshiftParserListener) EnterOpt_drop_behavior(ctx *Opt_drop_behaviorContext) {}

// ExitOpt_drop_behavior is called when production opt_drop_behavior is exited.
func (s *BaseRedshiftParserListener) ExitOpt_drop_behavior(ctx *Opt_drop_behaviorContext) {}

// EnterOpt_collate_clause is called when production opt_collate_clause is entered.
func (s *BaseRedshiftParserListener) EnterOpt_collate_clause(ctx *Opt_collate_clauseContext) {}

// ExitOpt_collate_clause is called when production opt_collate_clause is exited.
func (s *BaseRedshiftParserListener) ExitOpt_collate_clause(ctx *Opt_collate_clauseContext) {}

// EnterAlter_using is called when production alter_using is entered.
func (s *BaseRedshiftParserListener) EnterAlter_using(ctx *Alter_usingContext) {}

// ExitAlter_using is called when production alter_using is exited.
func (s *BaseRedshiftParserListener) ExitAlter_using(ctx *Alter_usingContext) {}

// EnterReplica_identity is called when production replica_identity is entered.
func (s *BaseRedshiftParserListener) EnterReplica_identity(ctx *Replica_identityContext) {}

// ExitReplica_identity is called when production replica_identity is exited.
func (s *BaseRedshiftParserListener) ExitReplica_identity(ctx *Replica_identityContext) {}

// EnterReloptions is called when production reloptions is entered.
func (s *BaseRedshiftParserListener) EnterReloptions(ctx *ReloptionsContext) {}

// ExitReloptions is called when production reloptions is exited.
func (s *BaseRedshiftParserListener) ExitReloptions(ctx *ReloptionsContext) {}

// EnterOpt_reloptions is called when production opt_reloptions is entered.
func (s *BaseRedshiftParserListener) EnterOpt_reloptions(ctx *Opt_reloptionsContext) {}

// ExitOpt_reloptions is called when production opt_reloptions is exited.
func (s *BaseRedshiftParserListener) ExitOpt_reloptions(ctx *Opt_reloptionsContext) {}

// EnterReloption_list is called when production reloption_list is entered.
func (s *BaseRedshiftParserListener) EnterReloption_list(ctx *Reloption_listContext) {}

// ExitReloption_list is called when production reloption_list is exited.
func (s *BaseRedshiftParserListener) ExitReloption_list(ctx *Reloption_listContext) {}

// EnterReloption_elem is called when production reloption_elem is entered.
func (s *BaseRedshiftParserListener) EnterReloption_elem(ctx *Reloption_elemContext) {}

// ExitReloption_elem is called when production reloption_elem is exited.
func (s *BaseRedshiftParserListener) ExitReloption_elem(ctx *Reloption_elemContext) {}

// EnterAlter_identity_column_option_list is called when production alter_identity_column_option_list is entered.
func (s *BaseRedshiftParserListener) EnterAlter_identity_column_option_list(ctx *Alter_identity_column_option_listContext) {
}

// ExitAlter_identity_column_option_list is called when production alter_identity_column_option_list is exited.
func (s *BaseRedshiftParserListener) ExitAlter_identity_column_option_list(ctx *Alter_identity_column_option_listContext) {
}

// EnterAlter_identity_column_option is called when production alter_identity_column_option is entered.
func (s *BaseRedshiftParserListener) EnterAlter_identity_column_option(ctx *Alter_identity_column_optionContext) {
}

// ExitAlter_identity_column_option is called when production alter_identity_column_option is exited.
func (s *BaseRedshiftParserListener) ExitAlter_identity_column_option(ctx *Alter_identity_column_optionContext) {
}

// EnterPartitionboundspec is called when production partitionboundspec is entered.
func (s *BaseRedshiftParserListener) EnterPartitionboundspec(ctx *PartitionboundspecContext) {}

// ExitPartitionboundspec is called when production partitionboundspec is exited.
func (s *BaseRedshiftParserListener) ExitPartitionboundspec(ctx *PartitionboundspecContext) {}

// EnterHash_partbound_elem is called when production hash_partbound_elem is entered.
func (s *BaseRedshiftParserListener) EnterHash_partbound_elem(ctx *Hash_partbound_elemContext) {}

// ExitHash_partbound_elem is called when production hash_partbound_elem is exited.
func (s *BaseRedshiftParserListener) ExitHash_partbound_elem(ctx *Hash_partbound_elemContext) {}

// EnterHash_partbound is called when production hash_partbound is entered.
func (s *BaseRedshiftParserListener) EnterHash_partbound(ctx *Hash_partboundContext) {}

// ExitHash_partbound is called when production hash_partbound is exited.
func (s *BaseRedshiftParserListener) ExitHash_partbound(ctx *Hash_partboundContext) {}

// EnterAltercompositetypestmt is called when production altercompositetypestmt is entered.
func (s *BaseRedshiftParserListener) EnterAltercompositetypestmt(ctx *AltercompositetypestmtContext) {
}

// ExitAltercompositetypestmt is called when production altercompositetypestmt is exited.
func (s *BaseRedshiftParserListener) ExitAltercompositetypestmt(ctx *AltercompositetypestmtContext) {}

// EnterAlter_type_cmds is called when production alter_type_cmds is entered.
func (s *BaseRedshiftParserListener) EnterAlter_type_cmds(ctx *Alter_type_cmdsContext) {}

// ExitAlter_type_cmds is called when production alter_type_cmds is exited.
func (s *BaseRedshiftParserListener) ExitAlter_type_cmds(ctx *Alter_type_cmdsContext) {}

// EnterAlter_type_cmd is called when production alter_type_cmd is entered.
func (s *BaseRedshiftParserListener) EnterAlter_type_cmd(ctx *Alter_type_cmdContext) {}

// ExitAlter_type_cmd is called when production alter_type_cmd is exited.
func (s *BaseRedshiftParserListener) ExitAlter_type_cmd(ctx *Alter_type_cmdContext) {}

// EnterCloseportalstmt is called when production closeportalstmt is entered.
func (s *BaseRedshiftParserListener) EnterCloseportalstmt(ctx *CloseportalstmtContext) {}

// ExitCloseportalstmt is called when production closeportalstmt is exited.
func (s *BaseRedshiftParserListener) ExitCloseportalstmt(ctx *CloseportalstmtContext) {}

// EnterCopystmt is called when production copystmt is entered.
func (s *BaseRedshiftParserListener) EnterCopystmt(ctx *CopystmtContext) {}

// ExitCopystmt is called when production copystmt is exited.
func (s *BaseRedshiftParserListener) ExitCopystmt(ctx *CopystmtContext) {}

// EnterRedshift_copy_authorization is called when production redshift_copy_authorization is entered.
func (s *BaseRedshiftParserListener) EnterRedshift_copy_authorization(ctx *Redshift_copy_authorizationContext) {
}

// ExitRedshift_copy_authorization is called when production redshift_copy_authorization is exited.
func (s *BaseRedshiftParserListener) ExitRedshift_copy_authorization(ctx *Redshift_copy_authorizationContext) {
}

// EnterRedshift_copy_format is called when production redshift_copy_format is entered.
func (s *BaseRedshiftParserListener) EnterRedshift_copy_format(ctx *Redshift_copy_formatContext) {}

// ExitRedshift_copy_format is called when production redshift_copy_format is exited.
func (s *BaseRedshiftParserListener) ExitRedshift_copy_format(ctx *Redshift_copy_formatContext) {}

// EnterRedshift_copy_parameter is called when production redshift_copy_parameter is entered.
func (s *BaseRedshiftParserListener) EnterRedshift_copy_parameter(ctx *Redshift_copy_parameterContext) {
}

// ExitRedshift_copy_parameter is called when production redshift_copy_parameter is exited.
func (s *BaseRedshiftParserListener) ExitRedshift_copy_parameter(ctx *Redshift_copy_parameterContext) {
}

// EnterCopy_param_name is called when production copy_param_name is entered.
func (s *BaseRedshiftParserListener) EnterCopy_param_name(ctx *Copy_param_nameContext) {}

// ExitCopy_param_name is called when production copy_param_name is exited.
func (s *BaseRedshiftParserListener) ExitCopy_param_name(ctx *Copy_param_nameContext) {}

// EnterCopy_param_value is called when production copy_param_value is entered.
func (s *BaseRedshiftParserListener) EnterCopy_param_value(ctx *Copy_param_valueContext) {}

// ExitCopy_param_value is called when production copy_param_value is exited.
func (s *BaseRedshiftParserListener) ExitCopy_param_value(ctx *Copy_param_valueContext) {}

// EnterCopy_from is called when production copy_from is entered.
func (s *BaseRedshiftParserListener) EnterCopy_from(ctx *Copy_fromContext) {}

// ExitCopy_from is called when production copy_from is exited.
func (s *BaseRedshiftParserListener) ExitCopy_from(ctx *Copy_fromContext) {}

// EnterOpt_program is called when production opt_program is entered.
func (s *BaseRedshiftParserListener) EnterOpt_program(ctx *Opt_programContext) {}

// ExitOpt_program is called when production opt_program is exited.
func (s *BaseRedshiftParserListener) ExitOpt_program(ctx *Opt_programContext) {}

// EnterCopy_file_name is called when production copy_file_name is entered.
func (s *BaseRedshiftParserListener) EnterCopy_file_name(ctx *Copy_file_nameContext) {}

// ExitCopy_file_name is called when production copy_file_name is exited.
func (s *BaseRedshiftParserListener) ExitCopy_file_name(ctx *Copy_file_nameContext) {}

// EnterCopy_options is called when production copy_options is entered.
func (s *BaseRedshiftParserListener) EnterCopy_options(ctx *Copy_optionsContext) {}

// ExitCopy_options is called when production copy_options is exited.
func (s *BaseRedshiftParserListener) ExitCopy_options(ctx *Copy_optionsContext) {}

// EnterCopy_opt_list is called when production copy_opt_list is entered.
func (s *BaseRedshiftParserListener) EnterCopy_opt_list(ctx *Copy_opt_listContext) {}

// ExitCopy_opt_list is called when production copy_opt_list is exited.
func (s *BaseRedshiftParserListener) ExitCopy_opt_list(ctx *Copy_opt_listContext) {}

// EnterCopy_opt_item is called when production copy_opt_item is entered.
func (s *BaseRedshiftParserListener) EnterCopy_opt_item(ctx *Copy_opt_itemContext) {}

// ExitCopy_opt_item is called when production copy_opt_item is exited.
func (s *BaseRedshiftParserListener) ExitCopy_opt_item(ctx *Copy_opt_itemContext) {}

// EnterOpt_binary is called when production opt_binary is entered.
func (s *BaseRedshiftParserListener) EnterOpt_binary(ctx *Opt_binaryContext) {}

// ExitOpt_binary is called when production opt_binary is exited.
func (s *BaseRedshiftParserListener) ExitOpt_binary(ctx *Opt_binaryContext) {}

// EnterCopy_delimiter is called when production copy_delimiter is entered.
func (s *BaseRedshiftParserListener) EnterCopy_delimiter(ctx *Copy_delimiterContext) {}

// ExitCopy_delimiter is called when production copy_delimiter is exited.
func (s *BaseRedshiftParserListener) ExitCopy_delimiter(ctx *Copy_delimiterContext) {}

// EnterOpt_using is called when production opt_using is entered.
func (s *BaseRedshiftParserListener) EnterOpt_using(ctx *Opt_usingContext) {}

// ExitOpt_using is called when production opt_using is exited.
func (s *BaseRedshiftParserListener) ExitOpt_using(ctx *Opt_usingContext) {}

// EnterCopy_generic_opt_list is called when production copy_generic_opt_list is entered.
func (s *BaseRedshiftParserListener) EnterCopy_generic_opt_list(ctx *Copy_generic_opt_listContext) {}

// ExitCopy_generic_opt_list is called when production copy_generic_opt_list is exited.
func (s *BaseRedshiftParserListener) ExitCopy_generic_opt_list(ctx *Copy_generic_opt_listContext) {}

// EnterCopy_generic_opt_elem is called when production copy_generic_opt_elem is entered.
func (s *BaseRedshiftParserListener) EnterCopy_generic_opt_elem(ctx *Copy_generic_opt_elemContext) {}

// ExitCopy_generic_opt_elem is called when production copy_generic_opt_elem is exited.
func (s *BaseRedshiftParserListener) ExitCopy_generic_opt_elem(ctx *Copy_generic_opt_elemContext) {}

// EnterCopy_generic_opt_arg is called when production copy_generic_opt_arg is entered.
func (s *BaseRedshiftParserListener) EnterCopy_generic_opt_arg(ctx *Copy_generic_opt_argContext) {}

// ExitCopy_generic_opt_arg is called when production copy_generic_opt_arg is exited.
func (s *BaseRedshiftParserListener) ExitCopy_generic_opt_arg(ctx *Copy_generic_opt_argContext) {}

// EnterCopy_generic_opt_arg_list is called when production copy_generic_opt_arg_list is entered.
func (s *BaseRedshiftParserListener) EnterCopy_generic_opt_arg_list(ctx *Copy_generic_opt_arg_listContext) {
}

// ExitCopy_generic_opt_arg_list is called when production copy_generic_opt_arg_list is exited.
func (s *BaseRedshiftParserListener) ExitCopy_generic_opt_arg_list(ctx *Copy_generic_opt_arg_listContext) {
}

// EnterCopy_generic_opt_arg_list_item is called when production copy_generic_opt_arg_list_item is entered.
func (s *BaseRedshiftParserListener) EnterCopy_generic_opt_arg_list_item(ctx *Copy_generic_opt_arg_list_itemContext) {
}

// ExitCopy_generic_opt_arg_list_item is called when production copy_generic_opt_arg_list_item is exited.
func (s *BaseRedshiftParserListener) ExitCopy_generic_opt_arg_list_item(ctx *Copy_generic_opt_arg_list_itemContext) {
}

// EnterCreatestmt is called when production createstmt is entered.
func (s *BaseRedshiftParserListener) EnterCreatestmt(ctx *CreatestmtContext) {}

// ExitCreatestmt is called when production createstmt is exited.
func (s *BaseRedshiftParserListener) ExitCreatestmt(ctx *CreatestmtContext) {}

// EnterOpt_table_attributes is called when production opt_table_attributes is entered.
func (s *BaseRedshiftParserListener) EnterOpt_table_attributes(ctx *Opt_table_attributesContext) {}

// ExitOpt_table_attributes is called when production opt_table_attributes is exited.
func (s *BaseRedshiftParserListener) ExitOpt_table_attributes(ctx *Opt_table_attributesContext) {}

// EnterOpttemp is called when production opttemp is entered.
func (s *BaseRedshiftParserListener) EnterOpttemp(ctx *OpttempContext) {}

// ExitOpttemp is called when production opttemp is exited.
func (s *BaseRedshiftParserListener) ExitOpttemp(ctx *OpttempContext) {}

// EnterOpttableelementlist is called when production opttableelementlist is entered.
func (s *BaseRedshiftParserListener) EnterOpttableelementlist(ctx *OpttableelementlistContext) {}

// ExitOpttableelementlist is called when production opttableelementlist is exited.
func (s *BaseRedshiftParserListener) ExitOpttableelementlist(ctx *OpttableelementlistContext) {}

// EnterOpttypedtableelementlist is called when production opttypedtableelementlist is entered.
func (s *BaseRedshiftParserListener) EnterOpttypedtableelementlist(ctx *OpttypedtableelementlistContext) {
}

// ExitOpttypedtableelementlist is called when production opttypedtableelementlist is exited.
func (s *BaseRedshiftParserListener) ExitOpttypedtableelementlist(ctx *OpttypedtableelementlistContext) {
}

// EnterTableelementlist is called when production tableelementlist is entered.
func (s *BaseRedshiftParserListener) EnterTableelementlist(ctx *TableelementlistContext) {}

// ExitTableelementlist is called when production tableelementlist is exited.
func (s *BaseRedshiftParserListener) ExitTableelementlist(ctx *TableelementlistContext) {}

// EnterTypedtableelementlist is called when production typedtableelementlist is entered.
func (s *BaseRedshiftParserListener) EnterTypedtableelementlist(ctx *TypedtableelementlistContext) {}

// ExitTypedtableelementlist is called when production typedtableelementlist is exited.
func (s *BaseRedshiftParserListener) ExitTypedtableelementlist(ctx *TypedtableelementlistContext) {}

// EnterTableelement is called when production tableelement is entered.
func (s *BaseRedshiftParserListener) EnterTableelement(ctx *TableelementContext) {}

// ExitTableelement is called when production tableelement is exited.
func (s *BaseRedshiftParserListener) ExitTableelement(ctx *TableelementContext) {}

// EnterTypedtableelement is called when production typedtableelement is entered.
func (s *BaseRedshiftParserListener) EnterTypedtableelement(ctx *TypedtableelementContext) {}

// ExitTypedtableelement is called when production typedtableelement is exited.
func (s *BaseRedshiftParserListener) ExitTypedtableelement(ctx *TypedtableelementContext) {}

// EnterColumnDef is called when production columnDef is entered.
func (s *BaseRedshiftParserListener) EnterColumnDef(ctx *ColumnDefContext) {}

// ExitColumnDef is called when production columnDef is exited.
func (s *BaseRedshiftParserListener) ExitColumnDef(ctx *ColumnDefContext) {}

// EnterRs_colattributes is called when production rs_colattributes is entered.
func (s *BaseRedshiftParserListener) EnterRs_colattributes(ctx *Rs_colattributesContext) {}

// ExitRs_colattributes is called when production rs_colattributes is exited.
func (s *BaseRedshiftParserListener) ExitRs_colattributes(ctx *Rs_colattributesContext) {}

// EnterRs_colattribute is called when production rs_colattribute is entered.
func (s *BaseRedshiftParserListener) EnterRs_colattribute(ctx *Rs_colattributeContext) {}

// ExitRs_colattribute is called when production rs_colattribute is exited.
func (s *BaseRedshiftParserListener) ExitRs_colattribute(ctx *Rs_colattributeContext) {}

// EnterColumnOptions is called when production columnOptions is entered.
func (s *BaseRedshiftParserListener) EnterColumnOptions(ctx *ColumnOptionsContext) {}

// ExitColumnOptions is called when production columnOptions is exited.
func (s *BaseRedshiftParserListener) ExitColumnOptions(ctx *ColumnOptionsContext) {}

// EnterColquallist is called when production colquallist is entered.
func (s *BaseRedshiftParserListener) EnterColquallist(ctx *ColquallistContext) {}

// ExitColquallist is called when production colquallist is exited.
func (s *BaseRedshiftParserListener) ExitColquallist(ctx *ColquallistContext) {}

// EnterColconstraint is called when production colconstraint is entered.
func (s *BaseRedshiftParserListener) EnterColconstraint(ctx *ColconstraintContext) {}

// ExitColconstraint is called when production colconstraint is exited.
func (s *BaseRedshiftParserListener) ExitColconstraint(ctx *ColconstraintContext) {}

// EnterColconstraintelem is called when production colconstraintelem is entered.
func (s *BaseRedshiftParserListener) EnterColconstraintelem(ctx *ColconstraintelemContext) {}

// ExitColconstraintelem is called when production colconstraintelem is exited.
func (s *BaseRedshiftParserListener) ExitColconstraintelem(ctx *ColconstraintelemContext) {}

// EnterOpt_unique_null_treatment is called when production opt_unique_null_treatment is entered.
func (s *BaseRedshiftParserListener) EnterOpt_unique_null_treatment(ctx *Opt_unique_null_treatmentContext) {
}

// ExitOpt_unique_null_treatment is called when production opt_unique_null_treatment is exited.
func (s *BaseRedshiftParserListener) ExitOpt_unique_null_treatment(ctx *Opt_unique_null_treatmentContext) {
}

// EnterGenerated_when is called when production generated_when is entered.
func (s *BaseRedshiftParserListener) EnterGenerated_when(ctx *Generated_whenContext) {}

// ExitGenerated_when is called when production generated_when is exited.
func (s *BaseRedshiftParserListener) ExitGenerated_when(ctx *Generated_whenContext) {}

// EnterConstraintattr is called when production constraintattr is entered.
func (s *BaseRedshiftParserListener) EnterConstraintattr(ctx *ConstraintattrContext) {}

// ExitConstraintattr is called when production constraintattr is exited.
func (s *BaseRedshiftParserListener) ExitConstraintattr(ctx *ConstraintattrContext) {}

// EnterTablelikeclause is called when production tablelikeclause is entered.
func (s *BaseRedshiftParserListener) EnterTablelikeclause(ctx *TablelikeclauseContext) {}

// ExitTablelikeclause is called when production tablelikeclause is exited.
func (s *BaseRedshiftParserListener) ExitTablelikeclause(ctx *TablelikeclauseContext) {}

// EnterTablelikeoptionlist is called when production tablelikeoptionlist is entered.
func (s *BaseRedshiftParserListener) EnterTablelikeoptionlist(ctx *TablelikeoptionlistContext) {}

// ExitTablelikeoptionlist is called when production tablelikeoptionlist is exited.
func (s *BaseRedshiftParserListener) ExitTablelikeoptionlist(ctx *TablelikeoptionlistContext) {}

// EnterTablelikeoption is called when production tablelikeoption is entered.
func (s *BaseRedshiftParserListener) EnterTablelikeoption(ctx *TablelikeoptionContext) {}

// ExitTablelikeoption is called when production tablelikeoption is exited.
func (s *BaseRedshiftParserListener) ExitTablelikeoption(ctx *TablelikeoptionContext) {}

// EnterTableconstraint is called when production tableconstraint is entered.
func (s *BaseRedshiftParserListener) EnterTableconstraint(ctx *TableconstraintContext) {}

// ExitTableconstraint is called when production tableconstraint is exited.
func (s *BaseRedshiftParserListener) ExitTableconstraint(ctx *TableconstraintContext) {}

// EnterConstraintelem is called when production constraintelem is entered.
func (s *BaseRedshiftParserListener) EnterConstraintelem(ctx *ConstraintelemContext) {}

// ExitConstraintelem is called when production constraintelem is exited.
func (s *BaseRedshiftParserListener) ExitConstraintelem(ctx *ConstraintelemContext) {}

// EnterOpt_no_inherit is called when production opt_no_inherit is entered.
func (s *BaseRedshiftParserListener) EnterOpt_no_inherit(ctx *Opt_no_inheritContext) {}

// ExitOpt_no_inherit is called when production opt_no_inherit is exited.
func (s *BaseRedshiftParserListener) ExitOpt_no_inherit(ctx *Opt_no_inheritContext) {}

// EnterOpt_column_list is called when production opt_column_list is entered.
func (s *BaseRedshiftParserListener) EnterOpt_column_list(ctx *Opt_column_listContext) {}

// ExitOpt_column_list is called when production opt_column_list is exited.
func (s *BaseRedshiftParserListener) ExitOpt_column_list(ctx *Opt_column_listContext) {}

// EnterColumnlist is called when production columnlist is entered.
func (s *BaseRedshiftParserListener) EnterColumnlist(ctx *ColumnlistContext) {}

// ExitColumnlist is called when production columnlist is exited.
func (s *BaseRedshiftParserListener) ExitColumnlist(ctx *ColumnlistContext) {}

// EnterColumnElem is called when production columnElem is entered.
func (s *BaseRedshiftParserListener) EnterColumnElem(ctx *ColumnElemContext) {}

// ExitColumnElem is called when production columnElem is exited.
func (s *BaseRedshiftParserListener) ExitColumnElem(ctx *ColumnElemContext) {}

// EnterOpt_c_include is called when production opt_c_include is entered.
func (s *BaseRedshiftParserListener) EnterOpt_c_include(ctx *Opt_c_includeContext) {}

// ExitOpt_c_include is called when production opt_c_include is exited.
func (s *BaseRedshiftParserListener) ExitOpt_c_include(ctx *Opt_c_includeContext) {}

// EnterKey_match is called when production key_match is entered.
func (s *BaseRedshiftParserListener) EnterKey_match(ctx *Key_matchContext) {}

// ExitKey_match is called when production key_match is exited.
func (s *BaseRedshiftParserListener) ExitKey_match(ctx *Key_matchContext) {}

// EnterExclusionconstraintlist is called when production exclusionconstraintlist is entered.
func (s *BaseRedshiftParserListener) EnterExclusionconstraintlist(ctx *ExclusionconstraintlistContext) {
}

// ExitExclusionconstraintlist is called when production exclusionconstraintlist is exited.
func (s *BaseRedshiftParserListener) ExitExclusionconstraintlist(ctx *ExclusionconstraintlistContext) {
}

// EnterExclusionconstraintelem is called when production exclusionconstraintelem is entered.
func (s *BaseRedshiftParserListener) EnterExclusionconstraintelem(ctx *ExclusionconstraintelemContext) {
}

// ExitExclusionconstraintelem is called when production exclusionconstraintelem is exited.
func (s *BaseRedshiftParserListener) ExitExclusionconstraintelem(ctx *ExclusionconstraintelemContext) {
}

// EnterExclusionwhereclause is called when production exclusionwhereclause is entered.
func (s *BaseRedshiftParserListener) EnterExclusionwhereclause(ctx *ExclusionwhereclauseContext) {}

// ExitExclusionwhereclause is called when production exclusionwhereclause is exited.
func (s *BaseRedshiftParserListener) ExitExclusionwhereclause(ctx *ExclusionwhereclauseContext) {}

// EnterKey_actions is called when production key_actions is entered.
func (s *BaseRedshiftParserListener) EnterKey_actions(ctx *Key_actionsContext) {}

// ExitKey_actions is called when production key_actions is exited.
func (s *BaseRedshiftParserListener) ExitKey_actions(ctx *Key_actionsContext) {}

// EnterKey_update is called when production key_update is entered.
func (s *BaseRedshiftParserListener) EnterKey_update(ctx *Key_updateContext) {}

// ExitKey_update is called when production key_update is exited.
func (s *BaseRedshiftParserListener) ExitKey_update(ctx *Key_updateContext) {}

// EnterKey_delete is called when production key_delete is entered.
func (s *BaseRedshiftParserListener) EnterKey_delete(ctx *Key_deleteContext) {}

// ExitKey_delete is called when production key_delete is exited.
func (s *BaseRedshiftParserListener) ExitKey_delete(ctx *Key_deleteContext) {}

// EnterKey_action is called when production key_action is entered.
func (s *BaseRedshiftParserListener) EnterKey_action(ctx *Key_actionContext) {}

// ExitKey_action is called when production key_action is exited.
func (s *BaseRedshiftParserListener) ExitKey_action(ctx *Key_actionContext) {}

// EnterOptinherit is called when production optinherit is entered.
func (s *BaseRedshiftParserListener) EnterOptinherit(ctx *OptinheritContext) {}

// ExitOptinherit is called when production optinherit is exited.
func (s *BaseRedshiftParserListener) ExitOptinherit(ctx *OptinheritContext) {}

// EnterOptpartitionspec is called when production optpartitionspec is entered.
func (s *BaseRedshiftParserListener) EnterOptpartitionspec(ctx *OptpartitionspecContext) {}

// ExitOptpartitionspec is called when production optpartitionspec is exited.
func (s *BaseRedshiftParserListener) ExitOptpartitionspec(ctx *OptpartitionspecContext) {}

// EnterPartitionspec is called when production partitionspec is entered.
func (s *BaseRedshiftParserListener) EnterPartitionspec(ctx *PartitionspecContext) {}

// ExitPartitionspec is called when production partitionspec is exited.
func (s *BaseRedshiftParserListener) ExitPartitionspec(ctx *PartitionspecContext) {}

// EnterPart_params is called when production part_params is entered.
func (s *BaseRedshiftParserListener) EnterPart_params(ctx *Part_paramsContext) {}

// ExitPart_params is called when production part_params is exited.
func (s *BaseRedshiftParserListener) ExitPart_params(ctx *Part_paramsContext) {}

// EnterPart_elem is called when production part_elem is entered.
func (s *BaseRedshiftParserListener) EnterPart_elem(ctx *Part_elemContext) {}

// ExitPart_elem is called when production part_elem is exited.
func (s *BaseRedshiftParserListener) ExitPart_elem(ctx *Part_elemContext) {}

// EnterTable_access_method_clause is called when production table_access_method_clause is entered.
func (s *BaseRedshiftParserListener) EnterTable_access_method_clause(ctx *Table_access_method_clauseContext) {
}

// ExitTable_access_method_clause is called when production table_access_method_clause is exited.
func (s *BaseRedshiftParserListener) ExitTable_access_method_clause(ctx *Table_access_method_clauseContext) {
}

// EnterOptwith is called when production optwith is entered.
func (s *BaseRedshiftParserListener) EnterOptwith(ctx *OptwithContext) {}

// ExitOptwith is called when production optwith is exited.
func (s *BaseRedshiftParserListener) ExitOptwith(ctx *OptwithContext) {}

// EnterOncommitoption is called when production oncommitoption is entered.
func (s *BaseRedshiftParserListener) EnterOncommitoption(ctx *OncommitoptionContext) {}

// ExitOncommitoption is called when production oncommitoption is exited.
func (s *BaseRedshiftParserListener) ExitOncommitoption(ctx *OncommitoptionContext) {}

// EnterOpttablespace is called when production opttablespace is entered.
func (s *BaseRedshiftParserListener) EnterOpttablespace(ctx *OpttablespaceContext) {}

// ExitOpttablespace is called when production opttablespace is exited.
func (s *BaseRedshiftParserListener) ExitOpttablespace(ctx *OpttablespaceContext) {}

// EnterOptredshifttableoptions is called when production optredshifttableoptions is entered.
func (s *BaseRedshiftParserListener) EnterOptredshifttableoptions(ctx *OptredshifttableoptionsContext) {
}

// ExitOptredshifttableoptions is called when production optredshifttableoptions is exited.
func (s *BaseRedshiftParserListener) ExitOptredshifttableoptions(ctx *OptredshifttableoptionsContext) {
}

// EnterRedshifttableoption is called when production redshifttableoption is entered.
func (s *BaseRedshiftParserListener) EnterRedshifttableoption(ctx *RedshifttableoptionContext) {}

// ExitRedshifttableoption is called when production redshifttableoption is exited.
func (s *BaseRedshiftParserListener) ExitRedshifttableoption(ctx *RedshifttableoptionContext) {}

// EnterSortkeyclause is called when production sortkeyclause is entered.
func (s *BaseRedshiftParserListener) EnterSortkeyclause(ctx *SortkeyclauseContext) {}

// ExitSortkeyclause is called when production sortkeyclause is exited.
func (s *BaseRedshiftParserListener) ExitSortkeyclause(ctx *SortkeyclauseContext) {}

// EnterSortkeyclausetype is called when production sortkeyclausetype is entered.
func (s *BaseRedshiftParserListener) EnterSortkeyclausetype(ctx *SortkeyclausetypeContext) {}

// ExitSortkeyclausetype is called when production sortkeyclausetype is exited.
func (s *BaseRedshiftParserListener) ExitSortkeyclausetype(ctx *SortkeyclausetypeContext) {}

// EnterOptconstablespace is called when production optconstablespace is entered.
func (s *BaseRedshiftParserListener) EnterOptconstablespace(ctx *OptconstablespaceContext) {}

// ExitOptconstablespace is called when production optconstablespace is exited.
func (s *BaseRedshiftParserListener) ExitOptconstablespace(ctx *OptconstablespaceContext) {}

// EnterExistingindex is called when production existingindex is entered.
func (s *BaseRedshiftParserListener) EnterExistingindex(ctx *ExistingindexContext) {}

// ExitExistingindex is called when production existingindex is exited.
func (s *BaseRedshiftParserListener) ExitExistingindex(ctx *ExistingindexContext) {}

// EnterCreatestatsstmt is called when production createstatsstmt is entered.
func (s *BaseRedshiftParserListener) EnterCreatestatsstmt(ctx *CreatestatsstmtContext) {}

// ExitCreatestatsstmt is called when production createstatsstmt is exited.
func (s *BaseRedshiftParserListener) ExitCreatestatsstmt(ctx *CreatestatsstmtContext) {}

// EnterAlterstatsstmt is called when production alterstatsstmt is entered.
func (s *BaseRedshiftParserListener) EnterAlterstatsstmt(ctx *AlterstatsstmtContext) {}

// ExitAlterstatsstmt is called when production alterstatsstmt is exited.
func (s *BaseRedshiftParserListener) ExitAlterstatsstmt(ctx *AlterstatsstmtContext) {}

// EnterCreateasstmt is called when production createasstmt is entered.
func (s *BaseRedshiftParserListener) EnterCreateasstmt(ctx *CreateasstmtContext) {}

// ExitCreateasstmt is called when production createasstmt is exited.
func (s *BaseRedshiftParserListener) ExitCreateasstmt(ctx *CreateasstmtContext) {}

// EnterCreate_as_target is called when production create_as_target is entered.
func (s *BaseRedshiftParserListener) EnterCreate_as_target(ctx *Create_as_targetContext) {}

// ExitCreate_as_target is called when production create_as_target is exited.
func (s *BaseRedshiftParserListener) ExitCreate_as_target(ctx *Create_as_targetContext) {}

// EnterOpt_backup_clause_table_attributes is called when production opt_backup_clause_table_attributes is entered.
func (s *BaseRedshiftParserListener) EnterOpt_backup_clause_table_attributes(ctx *Opt_backup_clause_table_attributesContext) {
}

// ExitOpt_backup_clause_table_attributes is called when production opt_backup_clause_table_attributes is exited.
func (s *BaseRedshiftParserListener) ExitOpt_backup_clause_table_attributes(ctx *Opt_backup_clause_table_attributesContext) {
}

// EnterTable_attributes is called when production table_attributes is entered.
func (s *BaseRedshiftParserListener) EnterTable_attributes(ctx *Table_attributesContext) {}

// ExitTable_attributes is called when production table_attributes is exited.
func (s *BaseRedshiftParserListener) ExitTable_attributes(ctx *Table_attributesContext) {}

// EnterOpt_backup_clause is called when production opt_backup_clause is entered.
func (s *BaseRedshiftParserListener) EnterOpt_backup_clause(ctx *Opt_backup_clauseContext) {}

// ExitOpt_backup_clause is called when production opt_backup_clause is exited.
func (s *BaseRedshiftParserListener) ExitOpt_backup_clause(ctx *Opt_backup_clauseContext) {}

// EnterOpt_with_data is called when production opt_with_data is entered.
func (s *BaseRedshiftParserListener) EnterOpt_with_data(ctx *Opt_with_dataContext) {}

// ExitOpt_with_data is called when production opt_with_data is exited.
func (s *BaseRedshiftParserListener) ExitOpt_with_data(ctx *Opt_with_dataContext) {}

// EnterCreatematviewstmt is called when production creatematviewstmt is entered.
func (s *BaseRedshiftParserListener) EnterCreatematviewstmt(ctx *CreatematviewstmtContext) {}

// ExitCreatematviewstmt is called when production creatematviewstmt is exited.
func (s *BaseRedshiftParserListener) ExitCreatematviewstmt(ctx *CreatematviewstmtContext) {}

// EnterOpt_auto_refresh is called when production opt_auto_refresh is entered.
func (s *BaseRedshiftParserListener) EnterOpt_auto_refresh(ctx *Opt_auto_refreshContext) {}

// ExitOpt_auto_refresh is called when production opt_auto_refresh is exited.
func (s *BaseRedshiftParserListener) ExitOpt_auto_refresh(ctx *Opt_auto_refreshContext) {}

// EnterRefreshmatviewstmt is called when production refreshmatviewstmt is entered.
func (s *BaseRedshiftParserListener) EnterRefreshmatviewstmt(ctx *RefreshmatviewstmtContext) {}

// ExitRefreshmatviewstmt is called when production refreshmatviewstmt is exited.
func (s *BaseRedshiftParserListener) ExitRefreshmatviewstmt(ctx *RefreshmatviewstmtContext) {}

// EnterCreateseqstmt is called when production createseqstmt is entered.
func (s *BaseRedshiftParserListener) EnterCreateseqstmt(ctx *CreateseqstmtContext) {}

// ExitCreateseqstmt is called when production createseqstmt is exited.
func (s *BaseRedshiftParserListener) ExitCreateseqstmt(ctx *CreateseqstmtContext) {}

// EnterAlterseqstmt is called when production alterseqstmt is entered.
func (s *BaseRedshiftParserListener) EnterAlterseqstmt(ctx *AlterseqstmtContext) {}

// ExitAlterseqstmt is called when production alterseqstmt is exited.
func (s *BaseRedshiftParserListener) ExitAlterseqstmt(ctx *AlterseqstmtContext) {}

// EnterOptseqoptlist is called when production optseqoptlist is entered.
func (s *BaseRedshiftParserListener) EnterOptseqoptlist(ctx *OptseqoptlistContext) {}

// ExitOptseqoptlist is called when production optseqoptlist is exited.
func (s *BaseRedshiftParserListener) ExitOptseqoptlist(ctx *OptseqoptlistContext) {}

// EnterOptparenthesizedseqoptlist is called when production optparenthesizedseqoptlist is entered.
func (s *BaseRedshiftParserListener) EnterOptparenthesizedseqoptlist(ctx *OptparenthesizedseqoptlistContext) {
}

// ExitOptparenthesizedseqoptlist is called when production optparenthesizedseqoptlist is exited.
func (s *BaseRedshiftParserListener) ExitOptparenthesizedseqoptlist(ctx *OptparenthesizedseqoptlistContext) {
}

// EnterSeqoptlist is called when production seqoptlist is entered.
func (s *BaseRedshiftParserListener) EnterSeqoptlist(ctx *SeqoptlistContext) {}

// ExitSeqoptlist is called when production seqoptlist is exited.
func (s *BaseRedshiftParserListener) ExitSeqoptlist(ctx *SeqoptlistContext) {}

// EnterSeqoptelem is called when production seqoptelem is entered.
func (s *BaseRedshiftParserListener) EnterSeqoptelem(ctx *SeqoptelemContext) {}

// ExitSeqoptelem is called when production seqoptelem is exited.
func (s *BaseRedshiftParserListener) ExitSeqoptelem(ctx *SeqoptelemContext) {}

// EnterOpt_by is called when production opt_by is entered.
func (s *BaseRedshiftParserListener) EnterOpt_by(ctx *Opt_byContext) {}

// ExitOpt_by is called when production opt_by is exited.
func (s *BaseRedshiftParserListener) ExitOpt_by(ctx *Opt_byContext) {}

// EnterNumericonly is called when production numericonly is entered.
func (s *BaseRedshiftParserListener) EnterNumericonly(ctx *NumericonlyContext) {}

// ExitNumericonly is called when production numericonly is exited.
func (s *BaseRedshiftParserListener) ExitNumericonly(ctx *NumericonlyContext) {}

// EnterNumericonly_list is called when production numericonly_list is entered.
func (s *BaseRedshiftParserListener) EnterNumericonly_list(ctx *Numericonly_listContext) {}

// ExitNumericonly_list is called when production numericonly_list is exited.
func (s *BaseRedshiftParserListener) ExitNumericonly_list(ctx *Numericonly_listContext) {}

// EnterCreateplangstmt is called when production createplangstmt is entered.
func (s *BaseRedshiftParserListener) EnterCreateplangstmt(ctx *CreateplangstmtContext) {}

// ExitCreateplangstmt is called when production createplangstmt is exited.
func (s *BaseRedshiftParserListener) ExitCreateplangstmt(ctx *CreateplangstmtContext) {}

// EnterOpt_trusted is called when production opt_trusted is entered.
func (s *BaseRedshiftParserListener) EnterOpt_trusted(ctx *Opt_trustedContext) {}

// ExitOpt_trusted is called when production opt_trusted is exited.
func (s *BaseRedshiftParserListener) ExitOpt_trusted(ctx *Opt_trustedContext) {}

// EnterHandler_name is called when production handler_name is entered.
func (s *BaseRedshiftParserListener) EnterHandler_name(ctx *Handler_nameContext) {}

// ExitHandler_name is called when production handler_name is exited.
func (s *BaseRedshiftParserListener) ExitHandler_name(ctx *Handler_nameContext) {}

// EnterOpt_inline_handler is called when production opt_inline_handler is entered.
func (s *BaseRedshiftParserListener) EnterOpt_inline_handler(ctx *Opt_inline_handlerContext) {}

// ExitOpt_inline_handler is called when production opt_inline_handler is exited.
func (s *BaseRedshiftParserListener) ExitOpt_inline_handler(ctx *Opt_inline_handlerContext) {}

// EnterValidator_clause is called when production validator_clause is entered.
func (s *BaseRedshiftParserListener) EnterValidator_clause(ctx *Validator_clauseContext) {}

// ExitValidator_clause is called when production validator_clause is exited.
func (s *BaseRedshiftParserListener) ExitValidator_clause(ctx *Validator_clauseContext) {}

// EnterOpt_validator is called when production opt_validator is entered.
func (s *BaseRedshiftParserListener) EnterOpt_validator(ctx *Opt_validatorContext) {}

// ExitOpt_validator is called when production opt_validator is exited.
func (s *BaseRedshiftParserListener) ExitOpt_validator(ctx *Opt_validatorContext) {}

// EnterOpt_procedural is called when production opt_procedural is entered.
func (s *BaseRedshiftParserListener) EnterOpt_procedural(ctx *Opt_proceduralContext) {}

// ExitOpt_procedural is called when production opt_procedural is exited.
func (s *BaseRedshiftParserListener) ExitOpt_procedural(ctx *Opt_proceduralContext) {}

// EnterCreatetablespacestmt is called when production createtablespacestmt is entered.
func (s *BaseRedshiftParserListener) EnterCreatetablespacestmt(ctx *CreatetablespacestmtContext) {}

// ExitCreatetablespacestmt is called when production createtablespacestmt is exited.
func (s *BaseRedshiftParserListener) ExitCreatetablespacestmt(ctx *CreatetablespacestmtContext) {}

// EnterOpttablespaceowner is called when production opttablespaceowner is entered.
func (s *BaseRedshiftParserListener) EnterOpttablespaceowner(ctx *OpttablespaceownerContext) {}

// ExitOpttablespaceowner is called when production opttablespaceowner is exited.
func (s *BaseRedshiftParserListener) ExitOpttablespaceowner(ctx *OpttablespaceownerContext) {}

// EnterDroptablespacestmt is called when production droptablespacestmt is entered.
func (s *BaseRedshiftParserListener) EnterDroptablespacestmt(ctx *DroptablespacestmtContext) {}

// ExitDroptablespacestmt is called when production droptablespacestmt is exited.
func (s *BaseRedshiftParserListener) ExitDroptablespacestmt(ctx *DroptablespacestmtContext) {}

// EnterCreateextensionstmt is called when production createextensionstmt is entered.
func (s *BaseRedshiftParserListener) EnterCreateextensionstmt(ctx *CreateextensionstmtContext) {}

// ExitCreateextensionstmt is called when production createextensionstmt is exited.
func (s *BaseRedshiftParserListener) ExitCreateextensionstmt(ctx *CreateextensionstmtContext) {}

// EnterCreate_extension_opt_list is called when production create_extension_opt_list is entered.
func (s *BaseRedshiftParserListener) EnterCreate_extension_opt_list(ctx *Create_extension_opt_listContext) {
}

// ExitCreate_extension_opt_list is called when production create_extension_opt_list is exited.
func (s *BaseRedshiftParserListener) ExitCreate_extension_opt_list(ctx *Create_extension_opt_listContext) {
}

// EnterCreate_extension_opt_item is called when production create_extension_opt_item is entered.
func (s *BaseRedshiftParserListener) EnterCreate_extension_opt_item(ctx *Create_extension_opt_itemContext) {
}

// ExitCreate_extension_opt_item is called when production create_extension_opt_item is exited.
func (s *BaseRedshiftParserListener) ExitCreate_extension_opt_item(ctx *Create_extension_opt_itemContext) {
}

// EnterAlterextensionstmt is called when production alterextensionstmt is entered.
func (s *BaseRedshiftParserListener) EnterAlterextensionstmt(ctx *AlterextensionstmtContext) {}

// ExitAlterextensionstmt is called when production alterextensionstmt is exited.
func (s *BaseRedshiftParserListener) ExitAlterextensionstmt(ctx *AlterextensionstmtContext) {}

// EnterAlter_extension_opt_list is called when production alter_extension_opt_list is entered.
func (s *BaseRedshiftParserListener) EnterAlter_extension_opt_list(ctx *Alter_extension_opt_listContext) {
}

// ExitAlter_extension_opt_list is called when production alter_extension_opt_list is exited.
func (s *BaseRedshiftParserListener) ExitAlter_extension_opt_list(ctx *Alter_extension_opt_listContext) {
}

// EnterAlter_extension_opt_item is called when production alter_extension_opt_item is entered.
func (s *BaseRedshiftParserListener) EnterAlter_extension_opt_item(ctx *Alter_extension_opt_itemContext) {
}

// ExitAlter_extension_opt_item is called when production alter_extension_opt_item is exited.
func (s *BaseRedshiftParserListener) ExitAlter_extension_opt_item(ctx *Alter_extension_opt_itemContext) {
}

// EnterAlterextensioncontentsstmt is called when production alterextensioncontentsstmt is entered.
func (s *BaseRedshiftParserListener) EnterAlterextensioncontentsstmt(ctx *AlterextensioncontentsstmtContext) {
}

// ExitAlterextensioncontentsstmt is called when production alterextensioncontentsstmt is exited.
func (s *BaseRedshiftParserListener) ExitAlterextensioncontentsstmt(ctx *AlterextensioncontentsstmtContext) {
}

// EnterCreatefdwstmt is called when production createfdwstmt is entered.
func (s *BaseRedshiftParserListener) EnterCreatefdwstmt(ctx *CreatefdwstmtContext) {}

// ExitCreatefdwstmt is called when production createfdwstmt is exited.
func (s *BaseRedshiftParserListener) ExitCreatefdwstmt(ctx *CreatefdwstmtContext) {}

// EnterFdw_option is called when production fdw_option is entered.
func (s *BaseRedshiftParserListener) EnterFdw_option(ctx *Fdw_optionContext) {}

// ExitFdw_option is called when production fdw_option is exited.
func (s *BaseRedshiftParserListener) ExitFdw_option(ctx *Fdw_optionContext) {}

// EnterFdw_options is called when production fdw_options is entered.
func (s *BaseRedshiftParserListener) EnterFdw_options(ctx *Fdw_optionsContext) {}

// ExitFdw_options is called when production fdw_options is exited.
func (s *BaseRedshiftParserListener) ExitFdw_options(ctx *Fdw_optionsContext) {}

// EnterOpt_fdw_options is called when production opt_fdw_options is entered.
func (s *BaseRedshiftParserListener) EnterOpt_fdw_options(ctx *Opt_fdw_optionsContext) {}

// ExitOpt_fdw_options is called when production opt_fdw_options is exited.
func (s *BaseRedshiftParserListener) ExitOpt_fdw_options(ctx *Opt_fdw_optionsContext) {}

// EnterAlterfdwstmt is called when production alterfdwstmt is entered.
func (s *BaseRedshiftParserListener) EnterAlterfdwstmt(ctx *AlterfdwstmtContext) {}

// ExitAlterfdwstmt is called when production alterfdwstmt is exited.
func (s *BaseRedshiftParserListener) ExitAlterfdwstmt(ctx *AlterfdwstmtContext) {}

// EnterCreate_generic_options is called when production create_generic_options is entered.
func (s *BaseRedshiftParserListener) EnterCreate_generic_options(ctx *Create_generic_optionsContext) {
}

// ExitCreate_generic_options is called when production create_generic_options is exited.
func (s *BaseRedshiftParserListener) ExitCreate_generic_options(ctx *Create_generic_optionsContext) {}

// EnterGeneric_option_list is called when production generic_option_list is entered.
func (s *BaseRedshiftParserListener) EnterGeneric_option_list(ctx *Generic_option_listContext) {}

// ExitGeneric_option_list is called when production generic_option_list is exited.
func (s *BaseRedshiftParserListener) ExitGeneric_option_list(ctx *Generic_option_listContext) {}

// EnterAlter_generic_options is called when production alter_generic_options is entered.
func (s *BaseRedshiftParserListener) EnterAlter_generic_options(ctx *Alter_generic_optionsContext) {}

// ExitAlter_generic_options is called when production alter_generic_options is exited.
func (s *BaseRedshiftParserListener) ExitAlter_generic_options(ctx *Alter_generic_optionsContext) {}

// EnterAlter_generic_option_list is called when production alter_generic_option_list is entered.
func (s *BaseRedshiftParserListener) EnterAlter_generic_option_list(ctx *Alter_generic_option_listContext) {
}

// ExitAlter_generic_option_list is called when production alter_generic_option_list is exited.
func (s *BaseRedshiftParserListener) ExitAlter_generic_option_list(ctx *Alter_generic_option_listContext) {
}

// EnterAlter_generic_option_elem is called when production alter_generic_option_elem is entered.
func (s *BaseRedshiftParserListener) EnterAlter_generic_option_elem(ctx *Alter_generic_option_elemContext) {
}

// ExitAlter_generic_option_elem is called when production alter_generic_option_elem is exited.
func (s *BaseRedshiftParserListener) ExitAlter_generic_option_elem(ctx *Alter_generic_option_elemContext) {
}

// EnterGeneric_option_elem is called when production generic_option_elem is entered.
func (s *BaseRedshiftParserListener) EnterGeneric_option_elem(ctx *Generic_option_elemContext) {}

// ExitGeneric_option_elem is called when production generic_option_elem is exited.
func (s *BaseRedshiftParserListener) ExitGeneric_option_elem(ctx *Generic_option_elemContext) {}

// EnterGeneric_option_name is called when production generic_option_name is entered.
func (s *BaseRedshiftParserListener) EnterGeneric_option_name(ctx *Generic_option_nameContext) {}

// ExitGeneric_option_name is called when production generic_option_name is exited.
func (s *BaseRedshiftParserListener) ExitGeneric_option_name(ctx *Generic_option_nameContext) {}

// EnterGeneric_option_arg is called when production generic_option_arg is entered.
func (s *BaseRedshiftParserListener) EnterGeneric_option_arg(ctx *Generic_option_argContext) {}

// ExitGeneric_option_arg is called when production generic_option_arg is exited.
func (s *BaseRedshiftParserListener) ExitGeneric_option_arg(ctx *Generic_option_argContext) {}

// EnterCreateforeignserverstmt is called when production createforeignserverstmt is entered.
func (s *BaseRedshiftParserListener) EnterCreateforeignserverstmt(ctx *CreateforeignserverstmtContext) {
}

// ExitCreateforeignserverstmt is called when production createforeignserverstmt is exited.
func (s *BaseRedshiftParserListener) ExitCreateforeignserverstmt(ctx *CreateforeignserverstmtContext) {
}

// EnterOpt_type is called when production opt_type is entered.
func (s *BaseRedshiftParserListener) EnterOpt_type(ctx *Opt_typeContext) {}

// ExitOpt_type is called when production opt_type is exited.
func (s *BaseRedshiftParserListener) ExitOpt_type(ctx *Opt_typeContext) {}

// EnterForeign_server_version is called when production foreign_server_version is entered.
func (s *BaseRedshiftParserListener) EnterForeign_server_version(ctx *Foreign_server_versionContext) {
}

// ExitForeign_server_version is called when production foreign_server_version is exited.
func (s *BaseRedshiftParserListener) ExitForeign_server_version(ctx *Foreign_server_versionContext) {}

// EnterOpt_foreign_server_version is called when production opt_foreign_server_version is entered.
func (s *BaseRedshiftParserListener) EnterOpt_foreign_server_version(ctx *Opt_foreign_server_versionContext) {
}

// ExitOpt_foreign_server_version is called when production opt_foreign_server_version is exited.
func (s *BaseRedshiftParserListener) ExitOpt_foreign_server_version(ctx *Opt_foreign_server_versionContext) {
}

// EnterAlterforeignserverstmt is called when production alterforeignserverstmt is entered.
func (s *BaseRedshiftParserListener) EnterAlterforeignserverstmt(ctx *AlterforeignserverstmtContext) {
}

// ExitAlterforeignserverstmt is called when production alterforeignserverstmt is exited.
func (s *BaseRedshiftParserListener) ExitAlterforeignserverstmt(ctx *AlterforeignserverstmtContext) {}

// EnterCreateforeigntablestmt is called when production createforeigntablestmt is entered.
func (s *BaseRedshiftParserListener) EnterCreateforeigntablestmt(ctx *CreateforeigntablestmtContext) {
}

// ExitCreateforeigntablestmt is called when production createforeigntablestmt is exited.
func (s *BaseRedshiftParserListener) ExitCreateforeigntablestmt(ctx *CreateforeigntablestmtContext) {}

// EnterImportforeignschemastmt is called when production importforeignschemastmt is entered.
func (s *BaseRedshiftParserListener) EnterImportforeignschemastmt(ctx *ImportforeignschemastmtContext) {
}

// ExitImportforeignschemastmt is called when production importforeignschemastmt is exited.
func (s *BaseRedshiftParserListener) ExitImportforeignschemastmt(ctx *ImportforeignschemastmtContext) {
}

// EnterImport_qualification_type is called when production import_qualification_type is entered.
func (s *BaseRedshiftParserListener) EnterImport_qualification_type(ctx *Import_qualification_typeContext) {
}

// ExitImport_qualification_type is called when production import_qualification_type is exited.
func (s *BaseRedshiftParserListener) ExitImport_qualification_type(ctx *Import_qualification_typeContext) {
}

// EnterImport_qualification is called when production import_qualification is entered.
func (s *BaseRedshiftParserListener) EnterImport_qualification(ctx *Import_qualificationContext) {}

// ExitImport_qualification is called when production import_qualification is exited.
func (s *BaseRedshiftParserListener) ExitImport_qualification(ctx *Import_qualificationContext) {}

// EnterCreateusermappingstmt is called when production createusermappingstmt is entered.
func (s *BaseRedshiftParserListener) EnterCreateusermappingstmt(ctx *CreateusermappingstmtContext) {}

// ExitCreateusermappingstmt is called when production createusermappingstmt is exited.
func (s *BaseRedshiftParserListener) ExitCreateusermappingstmt(ctx *CreateusermappingstmtContext) {}

// EnterAuth_ident is called when production auth_ident is entered.
func (s *BaseRedshiftParserListener) EnterAuth_ident(ctx *Auth_identContext) {}

// ExitAuth_ident is called when production auth_ident is exited.
func (s *BaseRedshiftParserListener) ExitAuth_ident(ctx *Auth_identContext) {}

// EnterDropusermappingstmt is called when production dropusermappingstmt is entered.
func (s *BaseRedshiftParserListener) EnterDropusermappingstmt(ctx *DropusermappingstmtContext) {}

// ExitDropusermappingstmt is called when production dropusermappingstmt is exited.
func (s *BaseRedshiftParserListener) ExitDropusermappingstmt(ctx *DropusermappingstmtContext) {}

// EnterAlterusermappingstmt is called when production alterusermappingstmt is entered.
func (s *BaseRedshiftParserListener) EnterAlterusermappingstmt(ctx *AlterusermappingstmtContext) {}

// ExitAlterusermappingstmt is called when production alterusermappingstmt is exited.
func (s *BaseRedshiftParserListener) ExitAlterusermappingstmt(ctx *AlterusermappingstmtContext) {}

// EnterCreatepolicystmt is called when production createpolicystmt is entered.
func (s *BaseRedshiftParserListener) EnterCreatepolicystmt(ctx *CreatepolicystmtContext) {}

// ExitCreatepolicystmt is called when production createpolicystmt is exited.
func (s *BaseRedshiftParserListener) ExitCreatepolicystmt(ctx *CreatepolicystmtContext) {}

// EnterAlterpolicystmt is called when production alterpolicystmt is entered.
func (s *BaseRedshiftParserListener) EnterAlterpolicystmt(ctx *AlterpolicystmtContext) {}

// ExitAlterpolicystmt is called when production alterpolicystmt is exited.
func (s *BaseRedshiftParserListener) ExitAlterpolicystmt(ctx *AlterpolicystmtContext) {}

// EnterRowsecurityoptionalexpr is called when production rowsecurityoptionalexpr is entered.
func (s *BaseRedshiftParserListener) EnterRowsecurityoptionalexpr(ctx *RowsecurityoptionalexprContext) {
}

// ExitRowsecurityoptionalexpr is called when production rowsecurityoptionalexpr is exited.
func (s *BaseRedshiftParserListener) ExitRowsecurityoptionalexpr(ctx *RowsecurityoptionalexprContext) {
}

// EnterRowsecurityoptionalwithcheck is called when production rowsecurityoptionalwithcheck is entered.
func (s *BaseRedshiftParserListener) EnterRowsecurityoptionalwithcheck(ctx *RowsecurityoptionalwithcheckContext) {
}

// ExitRowsecurityoptionalwithcheck is called when production rowsecurityoptionalwithcheck is exited.
func (s *BaseRedshiftParserListener) ExitRowsecurityoptionalwithcheck(ctx *RowsecurityoptionalwithcheckContext) {
}

// EnterRowsecuritydefaulttorole is called when production rowsecuritydefaulttorole is entered.
func (s *BaseRedshiftParserListener) EnterRowsecuritydefaulttorole(ctx *RowsecuritydefaulttoroleContext) {
}

// ExitRowsecuritydefaulttorole is called when production rowsecuritydefaulttorole is exited.
func (s *BaseRedshiftParserListener) ExitRowsecuritydefaulttorole(ctx *RowsecuritydefaulttoroleContext) {
}

// EnterRowsecurityoptionaltorole is called when production rowsecurityoptionaltorole is entered.
func (s *BaseRedshiftParserListener) EnterRowsecurityoptionaltorole(ctx *RowsecurityoptionaltoroleContext) {
}

// ExitRowsecurityoptionaltorole is called when production rowsecurityoptionaltorole is exited.
func (s *BaseRedshiftParserListener) ExitRowsecurityoptionaltorole(ctx *RowsecurityoptionaltoroleContext) {
}

// EnterRowsecuritydefaultpermissive is called when production rowsecuritydefaultpermissive is entered.
func (s *BaseRedshiftParserListener) EnterRowsecuritydefaultpermissive(ctx *RowsecuritydefaultpermissiveContext) {
}

// ExitRowsecuritydefaultpermissive is called when production rowsecuritydefaultpermissive is exited.
func (s *BaseRedshiftParserListener) ExitRowsecuritydefaultpermissive(ctx *RowsecuritydefaultpermissiveContext) {
}

// EnterRowsecuritydefaultforcmd is called when production rowsecuritydefaultforcmd is entered.
func (s *BaseRedshiftParserListener) EnterRowsecuritydefaultforcmd(ctx *RowsecuritydefaultforcmdContext) {
}

// ExitRowsecuritydefaultforcmd is called when production rowsecuritydefaultforcmd is exited.
func (s *BaseRedshiftParserListener) ExitRowsecuritydefaultforcmd(ctx *RowsecuritydefaultforcmdContext) {
}

// EnterRow_security_cmd is called when production row_security_cmd is entered.
func (s *BaseRedshiftParserListener) EnterRow_security_cmd(ctx *Row_security_cmdContext) {}

// ExitRow_security_cmd is called when production row_security_cmd is exited.
func (s *BaseRedshiftParserListener) ExitRow_security_cmd(ctx *Row_security_cmdContext) {}

// EnterCreateamstmt is called when production createamstmt is entered.
func (s *BaseRedshiftParserListener) EnterCreateamstmt(ctx *CreateamstmtContext) {}

// ExitCreateamstmt is called when production createamstmt is exited.
func (s *BaseRedshiftParserListener) ExitCreateamstmt(ctx *CreateamstmtContext) {}

// EnterAm_type is called when production am_type is entered.
func (s *BaseRedshiftParserListener) EnterAm_type(ctx *Am_typeContext) {}

// ExitAm_type is called when production am_type is exited.
func (s *BaseRedshiftParserListener) ExitAm_type(ctx *Am_typeContext) {}

// EnterCreatetrigstmt is called when production createtrigstmt is entered.
func (s *BaseRedshiftParserListener) EnterCreatetrigstmt(ctx *CreatetrigstmtContext) {}

// ExitCreatetrigstmt is called when production createtrigstmt is exited.
func (s *BaseRedshiftParserListener) ExitCreatetrigstmt(ctx *CreatetrigstmtContext) {}

// EnterTriggeractiontime is called when production triggeractiontime is entered.
func (s *BaseRedshiftParserListener) EnterTriggeractiontime(ctx *TriggeractiontimeContext) {}

// ExitTriggeractiontime is called when production triggeractiontime is exited.
func (s *BaseRedshiftParserListener) ExitTriggeractiontime(ctx *TriggeractiontimeContext) {}

// EnterTriggerevents is called when production triggerevents is entered.
func (s *BaseRedshiftParserListener) EnterTriggerevents(ctx *TriggereventsContext) {}

// ExitTriggerevents is called when production triggerevents is exited.
func (s *BaseRedshiftParserListener) ExitTriggerevents(ctx *TriggereventsContext) {}

// EnterTriggeroneevent is called when production triggeroneevent is entered.
func (s *BaseRedshiftParserListener) EnterTriggeroneevent(ctx *TriggeroneeventContext) {}

// ExitTriggeroneevent is called when production triggeroneevent is exited.
func (s *BaseRedshiftParserListener) ExitTriggeroneevent(ctx *TriggeroneeventContext) {}

// EnterTriggerreferencing is called when production triggerreferencing is entered.
func (s *BaseRedshiftParserListener) EnterTriggerreferencing(ctx *TriggerreferencingContext) {}

// ExitTriggerreferencing is called when production triggerreferencing is exited.
func (s *BaseRedshiftParserListener) ExitTriggerreferencing(ctx *TriggerreferencingContext) {}

// EnterTriggertransitions is called when production triggertransitions is entered.
func (s *BaseRedshiftParserListener) EnterTriggertransitions(ctx *TriggertransitionsContext) {}

// ExitTriggertransitions is called when production triggertransitions is exited.
func (s *BaseRedshiftParserListener) ExitTriggertransitions(ctx *TriggertransitionsContext) {}

// EnterTriggertransition is called when production triggertransition is entered.
func (s *BaseRedshiftParserListener) EnterTriggertransition(ctx *TriggertransitionContext) {}

// ExitTriggertransition is called when production triggertransition is exited.
func (s *BaseRedshiftParserListener) ExitTriggertransition(ctx *TriggertransitionContext) {}

// EnterTransitionoldornew is called when production transitionoldornew is entered.
func (s *BaseRedshiftParserListener) EnterTransitionoldornew(ctx *TransitionoldornewContext) {}

// ExitTransitionoldornew is called when production transitionoldornew is exited.
func (s *BaseRedshiftParserListener) ExitTransitionoldornew(ctx *TransitionoldornewContext) {}

// EnterTransitionrowortable is called when production transitionrowortable is entered.
func (s *BaseRedshiftParserListener) EnterTransitionrowortable(ctx *TransitionrowortableContext) {}

// ExitTransitionrowortable is called when production transitionrowortable is exited.
func (s *BaseRedshiftParserListener) ExitTransitionrowortable(ctx *TransitionrowortableContext) {}

// EnterTransitionrelname is called when production transitionrelname is entered.
func (s *BaseRedshiftParserListener) EnterTransitionrelname(ctx *TransitionrelnameContext) {}

// ExitTransitionrelname is called when production transitionrelname is exited.
func (s *BaseRedshiftParserListener) ExitTransitionrelname(ctx *TransitionrelnameContext) {}

// EnterTriggerforspec is called when production triggerforspec is entered.
func (s *BaseRedshiftParserListener) EnterTriggerforspec(ctx *TriggerforspecContext) {}

// ExitTriggerforspec is called when production triggerforspec is exited.
func (s *BaseRedshiftParserListener) ExitTriggerforspec(ctx *TriggerforspecContext) {}

// EnterTriggerforopteach is called when production triggerforopteach is entered.
func (s *BaseRedshiftParserListener) EnterTriggerforopteach(ctx *TriggerforopteachContext) {}

// ExitTriggerforopteach is called when production triggerforopteach is exited.
func (s *BaseRedshiftParserListener) ExitTriggerforopteach(ctx *TriggerforopteachContext) {}

// EnterTriggerfortype is called when production triggerfortype is entered.
func (s *BaseRedshiftParserListener) EnterTriggerfortype(ctx *TriggerfortypeContext) {}

// ExitTriggerfortype is called when production triggerfortype is exited.
func (s *BaseRedshiftParserListener) ExitTriggerfortype(ctx *TriggerfortypeContext) {}

// EnterTriggerwhen is called when production triggerwhen is entered.
func (s *BaseRedshiftParserListener) EnterTriggerwhen(ctx *TriggerwhenContext) {}

// ExitTriggerwhen is called when production triggerwhen is exited.
func (s *BaseRedshiftParserListener) ExitTriggerwhen(ctx *TriggerwhenContext) {}

// EnterFunction_or_procedure is called when production function_or_procedure is entered.
func (s *BaseRedshiftParserListener) EnterFunction_or_procedure(ctx *Function_or_procedureContext) {}

// ExitFunction_or_procedure is called when production function_or_procedure is exited.
func (s *BaseRedshiftParserListener) ExitFunction_or_procedure(ctx *Function_or_procedureContext) {}

// EnterTriggerfuncargs is called when production triggerfuncargs is entered.
func (s *BaseRedshiftParserListener) EnterTriggerfuncargs(ctx *TriggerfuncargsContext) {}

// ExitTriggerfuncargs is called when production triggerfuncargs is exited.
func (s *BaseRedshiftParserListener) ExitTriggerfuncargs(ctx *TriggerfuncargsContext) {}

// EnterTriggerfuncarg is called when production triggerfuncarg is entered.
func (s *BaseRedshiftParserListener) EnterTriggerfuncarg(ctx *TriggerfuncargContext) {}

// ExitTriggerfuncarg is called when production triggerfuncarg is exited.
func (s *BaseRedshiftParserListener) ExitTriggerfuncarg(ctx *TriggerfuncargContext) {}

// EnterOptconstrfromtable is called when production optconstrfromtable is entered.
func (s *BaseRedshiftParserListener) EnterOptconstrfromtable(ctx *OptconstrfromtableContext) {}

// ExitOptconstrfromtable is called when production optconstrfromtable is exited.
func (s *BaseRedshiftParserListener) ExitOptconstrfromtable(ctx *OptconstrfromtableContext) {}

// EnterConstraintattributespec is called when production constraintattributespec is entered.
func (s *BaseRedshiftParserListener) EnterConstraintattributespec(ctx *ConstraintattributespecContext) {
}

// ExitConstraintattributespec is called when production constraintattributespec is exited.
func (s *BaseRedshiftParserListener) ExitConstraintattributespec(ctx *ConstraintattributespecContext) {
}

// EnterConstraintattributeElem is called when production constraintattributeElem is entered.
func (s *BaseRedshiftParserListener) EnterConstraintattributeElem(ctx *ConstraintattributeElemContext) {
}

// ExitConstraintattributeElem is called when production constraintattributeElem is exited.
func (s *BaseRedshiftParserListener) ExitConstraintattributeElem(ctx *ConstraintattributeElemContext) {
}

// EnterCreateeventtrigstmt is called when production createeventtrigstmt is entered.
func (s *BaseRedshiftParserListener) EnterCreateeventtrigstmt(ctx *CreateeventtrigstmtContext) {}

// ExitCreateeventtrigstmt is called when production createeventtrigstmt is exited.
func (s *BaseRedshiftParserListener) ExitCreateeventtrigstmt(ctx *CreateeventtrigstmtContext) {}

// EnterEvent_trigger_when_list is called when production event_trigger_when_list is entered.
func (s *BaseRedshiftParserListener) EnterEvent_trigger_when_list(ctx *Event_trigger_when_listContext) {
}

// ExitEvent_trigger_when_list is called when production event_trigger_when_list is exited.
func (s *BaseRedshiftParserListener) ExitEvent_trigger_when_list(ctx *Event_trigger_when_listContext) {
}

// EnterEvent_trigger_when_item is called when production event_trigger_when_item is entered.
func (s *BaseRedshiftParserListener) EnterEvent_trigger_when_item(ctx *Event_trigger_when_itemContext) {
}

// ExitEvent_trigger_when_item is called when production event_trigger_when_item is exited.
func (s *BaseRedshiftParserListener) ExitEvent_trigger_when_item(ctx *Event_trigger_when_itemContext) {
}

// EnterEvent_trigger_value_list is called when production event_trigger_value_list is entered.
func (s *BaseRedshiftParserListener) EnterEvent_trigger_value_list(ctx *Event_trigger_value_listContext) {
}

// ExitEvent_trigger_value_list is called when production event_trigger_value_list is exited.
func (s *BaseRedshiftParserListener) ExitEvent_trigger_value_list(ctx *Event_trigger_value_listContext) {
}

// EnterAltereventtrigstmt is called when production altereventtrigstmt is entered.
func (s *BaseRedshiftParserListener) EnterAltereventtrigstmt(ctx *AltereventtrigstmtContext) {}

// ExitAltereventtrigstmt is called when production altereventtrigstmt is exited.
func (s *BaseRedshiftParserListener) ExitAltereventtrigstmt(ctx *AltereventtrigstmtContext) {}

// EnterEnable_trigger is called when production enable_trigger is entered.
func (s *BaseRedshiftParserListener) EnterEnable_trigger(ctx *Enable_triggerContext) {}

// ExitEnable_trigger is called when production enable_trigger is exited.
func (s *BaseRedshiftParserListener) ExitEnable_trigger(ctx *Enable_triggerContext) {}

// EnterCreateassertionstmt is called when production createassertionstmt is entered.
func (s *BaseRedshiftParserListener) EnterCreateassertionstmt(ctx *CreateassertionstmtContext) {}

// ExitCreateassertionstmt is called when production createassertionstmt is exited.
func (s *BaseRedshiftParserListener) ExitCreateassertionstmt(ctx *CreateassertionstmtContext) {}

// EnterDefinestmt is called when production definestmt is entered.
func (s *BaseRedshiftParserListener) EnterDefinestmt(ctx *DefinestmtContext) {}

// ExitDefinestmt is called when production definestmt is exited.
func (s *BaseRedshiftParserListener) ExitDefinestmt(ctx *DefinestmtContext) {}

// EnterDefinition is called when production definition is entered.
func (s *BaseRedshiftParserListener) EnterDefinition(ctx *DefinitionContext) {}

// ExitDefinition is called when production definition is exited.
func (s *BaseRedshiftParserListener) ExitDefinition(ctx *DefinitionContext) {}

// EnterDef_list is called when production def_list is entered.
func (s *BaseRedshiftParserListener) EnterDef_list(ctx *Def_listContext) {}

// ExitDef_list is called when production def_list is exited.
func (s *BaseRedshiftParserListener) ExitDef_list(ctx *Def_listContext) {}

// EnterDef_elem is called when production def_elem is entered.
func (s *BaseRedshiftParserListener) EnterDef_elem(ctx *Def_elemContext) {}

// ExitDef_elem is called when production def_elem is exited.
func (s *BaseRedshiftParserListener) ExitDef_elem(ctx *Def_elemContext) {}

// EnterDef_arg is called when production def_arg is entered.
func (s *BaseRedshiftParserListener) EnterDef_arg(ctx *Def_argContext) {}

// ExitDef_arg is called when production def_arg is exited.
func (s *BaseRedshiftParserListener) ExitDef_arg(ctx *Def_argContext) {}

// EnterOld_aggr_definition is called when production old_aggr_definition is entered.
func (s *BaseRedshiftParserListener) EnterOld_aggr_definition(ctx *Old_aggr_definitionContext) {}

// ExitOld_aggr_definition is called when production old_aggr_definition is exited.
func (s *BaseRedshiftParserListener) ExitOld_aggr_definition(ctx *Old_aggr_definitionContext) {}

// EnterOld_aggr_list is called when production old_aggr_list is entered.
func (s *BaseRedshiftParserListener) EnterOld_aggr_list(ctx *Old_aggr_listContext) {}

// ExitOld_aggr_list is called when production old_aggr_list is exited.
func (s *BaseRedshiftParserListener) ExitOld_aggr_list(ctx *Old_aggr_listContext) {}

// EnterOld_aggr_elem is called when production old_aggr_elem is entered.
func (s *BaseRedshiftParserListener) EnterOld_aggr_elem(ctx *Old_aggr_elemContext) {}

// ExitOld_aggr_elem is called when production old_aggr_elem is exited.
func (s *BaseRedshiftParserListener) ExitOld_aggr_elem(ctx *Old_aggr_elemContext) {}

// EnterOpt_enum_val_list is called when production opt_enum_val_list is entered.
func (s *BaseRedshiftParserListener) EnterOpt_enum_val_list(ctx *Opt_enum_val_listContext) {}

// ExitOpt_enum_val_list is called when production opt_enum_val_list is exited.
func (s *BaseRedshiftParserListener) ExitOpt_enum_val_list(ctx *Opt_enum_val_listContext) {}

// EnterEnum_val_list is called when production enum_val_list is entered.
func (s *BaseRedshiftParserListener) EnterEnum_val_list(ctx *Enum_val_listContext) {}

// ExitEnum_val_list is called when production enum_val_list is exited.
func (s *BaseRedshiftParserListener) ExitEnum_val_list(ctx *Enum_val_listContext) {}

// EnterAlterenumstmt is called when production alterenumstmt is entered.
func (s *BaseRedshiftParserListener) EnterAlterenumstmt(ctx *AlterenumstmtContext) {}

// ExitAlterenumstmt is called when production alterenumstmt is exited.
func (s *BaseRedshiftParserListener) ExitAlterenumstmt(ctx *AlterenumstmtContext) {}

// EnterOpt_if_not_exists is called when production opt_if_not_exists is entered.
func (s *BaseRedshiftParserListener) EnterOpt_if_not_exists(ctx *Opt_if_not_existsContext) {}

// ExitOpt_if_not_exists is called when production opt_if_not_exists is exited.
func (s *BaseRedshiftParserListener) ExitOpt_if_not_exists(ctx *Opt_if_not_existsContext) {}

// EnterCreateopclassstmt is called when production createopclassstmt is entered.
func (s *BaseRedshiftParserListener) EnterCreateopclassstmt(ctx *CreateopclassstmtContext) {}

// ExitCreateopclassstmt is called when production createopclassstmt is exited.
func (s *BaseRedshiftParserListener) ExitCreateopclassstmt(ctx *CreateopclassstmtContext) {}

// EnterOpclass_item_list is called when production opclass_item_list is entered.
func (s *BaseRedshiftParserListener) EnterOpclass_item_list(ctx *Opclass_item_listContext) {}

// ExitOpclass_item_list is called when production opclass_item_list is exited.
func (s *BaseRedshiftParserListener) ExitOpclass_item_list(ctx *Opclass_item_listContext) {}

// EnterOpclass_item is called when production opclass_item is entered.
func (s *BaseRedshiftParserListener) EnterOpclass_item(ctx *Opclass_itemContext) {}

// ExitOpclass_item is called when production opclass_item is exited.
func (s *BaseRedshiftParserListener) ExitOpclass_item(ctx *Opclass_itemContext) {}

// EnterOpt_default is called when production opt_default is entered.
func (s *BaseRedshiftParserListener) EnterOpt_default(ctx *Opt_defaultContext) {}

// ExitOpt_default is called when production opt_default is exited.
func (s *BaseRedshiftParserListener) ExitOpt_default(ctx *Opt_defaultContext) {}

// EnterOpt_opfamily is called when production opt_opfamily is entered.
func (s *BaseRedshiftParserListener) EnterOpt_opfamily(ctx *Opt_opfamilyContext) {}

// ExitOpt_opfamily is called when production opt_opfamily is exited.
func (s *BaseRedshiftParserListener) ExitOpt_opfamily(ctx *Opt_opfamilyContext) {}

// EnterOpclass_purpose is called when production opclass_purpose is entered.
func (s *BaseRedshiftParserListener) EnterOpclass_purpose(ctx *Opclass_purposeContext) {}

// ExitOpclass_purpose is called when production opclass_purpose is exited.
func (s *BaseRedshiftParserListener) ExitOpclass_purpose(ctx *Opclass_purposeContext) {}

// EnterOpt_recheck is called when production opt_recheck is entered.
func (s *BaseRedshiftParserListener) EnterOpt_recheck(ctx *Opt_recheckContext) {}

// ExitOpt_recheck is called when production opt_recheck is exited.
func (s *BaseRedshiftParserListener) ExitOpt_recheck(ctx *Opt_recheckContext) {}

// EnterCreateopfamilystmt is called when production createopfamilystmt is entered.
func (s *BaseRedshiftParserListener) EnterCreateopfamilystmt(ctx *CreateopfamilystmtContext) {}

// ExitCreateopfamilystmt is called when production createopfamilystmt is exited.
func (s *BaseRedshiftParserListener) ExitCreateopfamilystmt(ctx *CreateopfamilystmtContext) {}

// EnterAlteropfamilystmt is called when production alteropfamilystmt is entered.
func (s *BaseRedshiftParserListener) EnterAlteropfamilystmt(ctx *AlteropfamilystmtContext) {}

// ExitAlteropfamilystmt is called when production alteropfamilystmt is exited.
func (s *BaseRedshiftParserListener) ExitAlteropfamilystmt(ctx *AlteropfamilystmtContext) {}

// EnterOpclass_drop_list is called when production opclass_drop_list is entered.
func (s *BaseRedshiftParserListener) EnterOpclass_drop_list(ctx *Opclass_drop_listContext) {}

// ExitOpclass_drop_list is called when production opclass_drop_list is exited.
func (s *BaseRedshiftParserListener) ExitOpclass_drop_list(ctx *Opclass_drop_listContext) {}

// EnterOpclass_drop is called when production opclass_drop is entered.
func (s *BaseRedshiftParserListener) EnterOpclass_drop(ctx *Opclass_dropContext) {}

// ExitOpclass_drop is called when production opclass_drop is exited.
func (s *BaseRedshiftParserListener) ExitOpclass_drop(ctx *Opclass_dropContext) {}

// EnterDropopclassstmt is called when production dropopclassstmt is entered.
func (s *BaseRedshiftParserListener) EnterDropopclassstmt(ctx *DropopclassstmtContext) {}

// ExitDropopclassstmt is called when production dropopclassstmt is exited.
func (s *BaseRedshiftParserListener) ExitDropopclassstmt(ctx *DropopclassstmtContext) {}

// EnterDropopfamilystmt is called when production dropopfamilystmt is entered.
func (s *BaseRedshiftParserListener) EnterDropopfamilystmt(ctx *DropopfamilystmtContext) {}

// ExitDropopfamilystmt is called when production dropopfamilystmt is exited.
func (s *BaseRedshiftParserListener) ExitDropopfamilystmt(ctx *DropopfamilystmtContext) {}

// EnterDropownedstmt is called when production dropownedstmt is entered.
func (s *BaseRedshiftParserListener) EnterDropownedstmt(ctx *DropownedstmtContext) {}

// ExitDropownedstmt is called when production dropownedstmt is exited.
func (s *BaseRedshiftParserListener) ExitDropownedstmt(ctx *DropownedstmtContext) {}

// EnterReassignownedstmt is called when production reassignownedstmt is entered.
func (s *BaseRedshiftParserListener) EnterReassignownedstmt(ctx *ReassignownedstmtContext) {}

// ExitReassignownedstmt is called when production reassignownedstmt is exited.
func (s *BaseRedshiftParserListener) ExitReassignownedstmt(ctx *ReassignownedstmtContext) {}

// EnterDropstmt is called when production dropstmt is entered.
func (s *BaseRedshiftParserListener) EnterDropstmt(ctx *DropstmtContext) {}

// ExitDropstmt is called when production dropstmt is exited.
func (s *BaseRedshiftParserListener) ExitDropstmt(ctx *DropstmtContext) {}

// EnterObject_type_any_name is called when production object_type_any_name is entered.
func (s *BaseRedshiftParserListener) EnterObject_type_any_name(ctx *Object_type_any_nameContext) {}

// ExitObject_type_any_name is called when production object_type_any_name is exited.
func (s *BaseRedshiftParserListener) ExitObject_type_any_name(ctx *Object_type_any_nameContext) {}

// EnterObject_type_name is called when production object_type_name is entered.
func (s *BaseRedshiftParserListener) EnterObject_type_name(ctx *Object_type_nameContext) {}

// ExitObject_type_name is called when production object_type_name is exited.
func (s *BaseRedshiftParserListener) ExitObject_type_name(ctx *Object_type_nameContext) {}

// EnterDrop_type_name is called when production drop_type_name is entered.
func (s *BaseRedshiftParserListener) EnterDrop_type_name(ctx *Drop_type_nameContext) {}

// ExitDrop_type_name is called when production drop_type_name is exited.
func (s *BaseRedshiftParserListener) ExitDrop_type_name(ctx *Drop_type_nameContext) {}

// EnterObject_type_name_on_any_name is called when production object_type_name_on_any_name is entered.
func (s *BaseRedshiftParserListener) EnterObject_type_name_on_any_name(ctx *Object_type_name_on_any_nameContext) {
}

// ExitObject_type_name_on_any_name is called when production object_type_name_on_any_name is exited.
func (s *BaseRedshiftParserListener) ExitObject_type_name_on_any_name(ctx *Object_type_name_on_any_nameContext) {
}

// EnterAny_name_list is called when production any_name_list is entered.
func (s *BaseRedshiftParserListener) EnterAny_name_list(ctx *Any_name_listContext) {}

// ExitAny_name_list is called when production any_name_list is exited.
func (s *BaseRedshiftParserListener) ExitAny_name_list(ctx *Any_name_listContext) {}

// EnterAny_name is called when production any_name is entered.
func (s *BaseRedshiftParserListener) EnterAny_name(ctx *Any_nameContext) {}

// ExitAny_name is called when production any_name is exited.
func (s *BaseRedshiftParserListener) ExitAny_name(ctx *Any_nameContext) {}

// EnterAttrs is called when production attrs is entered.
func (s *BaseRedshiftParserListener) EnterAttrs(ctx *AttrsContext) {}

// ExitAttrs is called when production attrs is exited.
func (s *BaseRedshiftParserListener) ExitAttrs(ctx *AttrsContext) {}

// EnterType_name_list is called when production type_name_list is entered.
func (s *BaseRedshiftParserListener) EnterType_name_list(ctx *Type_name_listContext) {}

// ExitType_name_list is called when production type_name_list is exited.
func (s *BaseRedshiftParserListener) ExitType_name_list(ctx *Type_name_listContext) {}

// EnterTruncatestmt is called when production truncatestmt is entered.
func (s *BaseRedshiftParserListener) EnterTruncatestmt(ctx *TruncatestmtContext) {}

// ExitTruncatestmt is called when production truncatestmt is exited.
func (s *BaseRedshiftParserListener) ExitTruncatestmt(ctx *TruncatestmtContext) {}

// EnterOpt_restart_seqs is called when production opt_restart_seqs is entered.
func (s *BaseRedshiftParserListener) EnterOpt_restart_seqs(ctx *Opt_restart_seqsContext) {}

// ExitOpt_restart_seqs is called when production opt_restart_seqs is exited.
func (s *BaseRedshiftParserListener) ExitOpt_restart_seqs(ctx *Opt_restart_seqsContext) {}

// EnterCommentstmt is called when production commentstmt is entered.
func (s *BaseRedshiftParserListener) EnterCommentstmt(ctx *CommentstmtContext) {}

// ExitCommentstmt is called when production commentstmt is exited.
func (s *BaseRedshiftParserListener) ExitCommentstmt(ctx *CommentstmtContext) {}

// EnterComment_text is called when production comment_text is entered.
func (s *BaseRedshiftParserListener) EnterComment_text(ctx *Comment_textContext) {}

// ExitComment_text is called when production comment_text is exited.
func (s *BaseRedshiftParserListener) ExitComment_text(ctx *Comment_textContext) {}

// EnterSeclabelstmt is called when production seclabelstmt is entered.
func (s *BaseRedshiftParserListener) EnterSeclabelstmt(ctx *SeclabelstmtContext) {}

// ExitSeclabelstmt is called when production seclabelstmt is exited.
func (s *BaseRedshiftParserListener) ExitSeclabelstmt(ctx *SeclabelstmtContext) {}

// EnterOpt_provider is called when production opt_provider is entered.
func (s *BaseRedshiftParserListener) EnterOpt_provider(ctx *Opt_providerContext) {}

// ExitOpt_provider is called when production opt_provider is exited.
func (s *BaseRedshiftParserListener) ExitOpt_provider(ctx *Opt_providerContext) {}

// EnterSecurity_label is called when production security_label is entered.
func (s *BaseRedshiftParserListener) EnterSecurity_label(ctx *Security_labelContext) {}

// ExitSecurity_label is called when production security_label is exited.
func (s *BaseRedshiftParserListener) ExitSecurity_label(ctx *Security_labelContext) {}

// EnterFetchstmt is called when production fetchstmt is entered.
func (s *BaseRedshiftParserListener) EnterFetchstmt(ctx *FetchstmtContext) {}

// ExitFetchstmt is called when production fetchstmt is exited.
func (s *BaseRedshiftParserListener) ExitFetchstmt(ctx *FetchstmtContext) {}

// EnterFetch_args is called when production fetch_args is entered.
func (s *BaseRedshiftParserListener) EnterFetch_args(ctx *Fetch_argsContext) {}

// ExitFetch_args is called when production fetch_args is exited.
func (s *BaseRedshiftParserListener) ExitFetch_args(ctx *Fetch_argsContext) {}

// EnterGrantstmt is called when production grantstmt is entered.
func (s *BaseRedshiftParserListener) EnterGrantstmt(ctx *GrantstmtContext) {}

// ExitGrantstmt is called when production grantstmt is exited.
func (s *BaseRedshiftParserListener) ExitGrantstmt(ctx *GrantstmtContext) {}

// EnterGrant_permissions_for_rls_lookup_tables is called when production grant_permissions_for_rls_lookup_tables is entered.
func (s *BaseRedshiftParserListener) EnterGrant_permissions_for_rls_lookup_tables(ctx *Grant_permissions_for_rls_lookup_tablesContext) {
}

// ExitGrant_permissions_for_rls_lookup_tables is called when production grant_permissions_for_rls_lookup_tables is exited.
func (s *BaseRedshiftParserListener) ExitGrant_permissions_for_rls_lookup_tables(ctx *Grant_permissions_for_rls_lookup_tablesContext) {
}

// EnterGrant_explain_permissions_for_row_level_security_policy_filters is called when production grant_explain_permissions_for_row_level_security_policy_filters is entered.
func (s *BaseRedshiftParserListener) EnterGrant_explain_permissions_for_row_level_security_policy_filters(ctx *Grant_explain_permissions_for_row_level_security_policy_filtersContext) {
}

// ExitGrant_explain_permissions_for_row_level_security_policy_filters is called when production grant_explain_permissions_for_row_level_security_policy_filters is exited.
func (s *BaseRedshiftParserListener) ExitGrant_explain_permissions_for_row_level_security_policy_filters(ctx *Grant_explain_permissions_for_row_level_security_policy_filtersContext) {
}

// EnterGrant_machine_learning_permissions is called when production grant_machine_learning_permissions is entered.
func (s *BaseRedshiftParserListener) EnterGrant_machine_learning_permissions(ctx *Grant_machine_learning_permissionsContext) {
}

// ExitGrant_machine_learning_permissions is called when production grant_machine_learning_permissions is exited.
func (s *BaseRedshiftParserListener) ExitGrant_machine_learning_permissions(ctx *Grant_machine_learning_permissionsContext) {
}

// EnterGrant_role_permissions is called when production grant_role_permissions is entered.
func (s *BaseRedshiftParserListener) EnterGrant_role_permissions(ctx *Grant_role_permissionsContext) {
}

// ExitGrant_role_permissions is called when production grant_role_permissions is exited.
func (s *BaseRedshiftParserListener) ExitGrant_role_permissions(ctx *Grant_role_permissionsContext) {}

// EnterGrant_role_permission_target_list is called when production grant_role_permission_target_list is entered.
func (s *BaseRedshiftParserListener) EnterGrant_role_permission_target_list(ctx *Grant_role_permission_target_listContext) {
}

// ExitGrant_role_permission_target_list is called when production grant_role_permission_target_list is exited.
func (s *BaseRedshiftParserListener) ExitGrant_role_permission_target_list(ctx *Grant_role_permission_target_listContext) {
}

// EnterGrant_role_permission_target_list_item is called when production grant_role_permission_target_list_item is entered.
func (s *BaseRedshiftParserListener) EnterGrant_role_permission_target_list_item(ctx *Grant_role_permission_target_list_itemContext) {
}

// ExitGrant_role_permission_target_list_item is called when production grant_role_permission_target_list_item is exited.
func (s *BaseRedshiftParserListener) ExitGrant_role_permission_target_list_item(ctx *Grant_role_permission_target_list_itemContext) {
}

// EnterSystem_permissions is called when production system_permissions is entered.
func (s *BaseRedshiftParserListener) EnterSystem_permissions(ctx *System_permissionsContext) {}

// ExitSystem_permissions is called when production system_permissions is exited.
func (s *BaseRedshiftParserListener) ExitSystem_permissions(ctx *System_permissionsContext) {}

// EnterSystem_permissions_item is called when production system_permissions_item is entered.
func (s *BaseRedshiftParserListener) EnterSystem_permissions_item(ctx *System_permissions_itemContext) {
}

// ExitSystem_permissions_item is called when production system_permissions_item is exited.
func (s *BaseRedshiftParserListener) ExitSystem_permissions_item(ctx *System_permissions_itemContext) {
}

// EnterOpt_with_admin_option is called when production opt_with_admin_option is entered.
func (s *BaseRedshiftParserListener) EnterOpt_with_admin_option(ctx *Opt_with_admin_optionContext) {}

// ExitOpt_with_admin_option is called when production opt_with_admin_option is exited.
func (s *BaseRedshiftParserListener) ExitOpt_with_admin_option(ctx *Opt_with_admin_optionContext) {}

// EnterGrant_scoped_permissions is called when production grant_scoped_permissions is entered.
func (s *BaseRedshiftParserListener) EnterGrant_scoped_permissions(ctx *Grant_scoped_permissionsContext) {
}

// ExitGrant_scoped_permissions is called when production grant_scoped_permissions is exited.
func (s *BaseRedshiftParserListener) ExitGrant_scoped_permissions(ctx *Grant_scoped_permissionsContext) {
}

// EnterGrant_scoped_schemas_permissions is called when production grant_scoped_schemas_permissions is entered.
func (s *BaseRedshiftParserListener) EnterGrant_scoped_schemas_permissions(ctx *Grant_scoped_schemas_permissionsContext) {
}

// ExitGrant_scoped_schemas_permissions is called when production grant_scoped_schemas_permissions is exited.
func (s *BaseRedshiftParserListener) ExitGrant_scoped_schemas_permissions(ctx *Grant_scoped_schemas_permissionsContext) {
}

// EnterGrant_scoped_tables_permissions is called when production grant_scoped_tables_permissions is entered.
func (s *BaseRedshiftParserListener) EnterGrant_scoped_tables_permissions(ctx *Grant_scoped_tables_permissionsContext) {
}

// ExitGrant_scoped_tables_permissions is called when production grant_scoped_tables_permissions is exited.
func (s *BaseRedshiftParserListener) ExitGrant_scoped_tables_permissions(ctx *Grant_scoped_tables_permissionsContext) {
}

// EnterGrant_scoped_functions_permissions is called when production grant_scoped_functions_permissions is entered.
func (s *BaseRedshiftParserListener) EnterGrant_scoped_functions_permissions(ctx *Grant_scoped_functions_permissionsContext) {
}

// ExitGrant_scoped_functions_permissions is called when production grant_scoped_functions_permissions is exited.
func (s *BaseRedshiftParserListener) ExitGrant_scoped_functions_permissions(ctx *Grant_scoped_functions_permissionsContext) {
}

// EnterGrant_scoped_procedures_permissions is called when production grant_scoped_procedures_permissions is entered.
func (s *BaseRedshiftParserListener) EnterGrant_scoped_procedures_permissions(ctx *Grant_scoped_procedures_permissionsContext) {
}

// ExitGrant_scoped_procedures_permissions is called when production grant_scoped_procedures_permissions is exited.
func (s *BaseRedshiftParserListener) ExitGrant_scoped_procedures_permissions(ctx *Grant_scoped_procedures_permissionsContext) {
}

// EnterGrant_scoped_languages_permissions is called when production grant_scoped_languages_permissions is entered.
func (s *BaseRedshiftParserListener) EnterGrant_scoped_languages_permissions(ctx *Grant_scoped_languages_permissionsContext) {
}

// ExitGrant_scoped_languages_permissions is called when production grant_scoped_languages_permissions is exited.
func (s *BaseRedshiftParserListener) ExitGrant_scoped_languages_permissions(ctx *Grant_scoped_languages_permissionsContext) {
}

// EnterGrant_scoped_copy_jobs_permissions is called when production grant_scoped_copy_jobs_permissions is entered.
func (s *BaseRedshiftParserListener) EnterGrant_scoped_copy_jobs_permissions(ctx *Grant_scoped_copy_jobs_permissionsContext) {
}

// ExitGrant_scoped_copy_jobs_permissions is called when production grant_scoped_copy_jobs_permissions is exited.
func (s *BaseRedshiftParserListener) ExitGrant_scoped_copy_jobs_permissions(ctx *Grant_scoped_copy_jobs_permissionsContext) {
}

// EnterGrantee_list_without_public is called when production grantee_list_without_public is entered.
func (s *BaseRedshiftParserListener) EnterGrantee_list_without_public(ctx *Grantee_list_without_publicContext) {
}

// ExitGrantee_list_without_public is called when production grantee_list_without_public is exited.
func (s *BaseRedshiftParserListener) ExitGrantee_list_without_public(ctx *Grantee_list_without_publicContext) {
}

// EnterGrantee_without_public is called when production grantee_without_public is entered.
func (s *BaseRedshiftParserListener) EnterGrantee_without_public(ctx *Grantee_without_publicContext) {
}

// ExitGrantee_without_public is called when production grantee_without_public is exited.
func (s *BaseRedshiftParserListener) ExitGrantee_without_public(ctx *Grantee_without_publicContext) {}

// EnterGrant_datashare_permissions is called when production grant_datashare_permissions is entered.
func (s *BaseRedshiftParserListener) EnterGrant_datashare_permissions(ctx *Grant_datashare_permissionsContext) {
}

// ExitGrant_datashare_permissions is called when production grant_datashare_permissions is exited.
func (s *BaseRedshiftParserListener) ExitGrant_datashare_permissions(ctx *Grant_datashare_permissionsContext) {
}

// EnterGrant_spectrum_integration_permissions is called when production grant_spectrum_integration_permissions is entered.
func (s *BaseRedshiftParserListener) EnterGrant_spectrum_integration_permissions(ctx *Grant_spectrum_integration_permissionsContext) {
}

// ExitGrant_spectrum_integration_permissions is called when production grant_spectrum_integration_permissions is exited.
func (s *BaseRedshiftParserListener) ExitGrant_spectrum_integration_permissions(ctx *Grant_spectrum_integration_permissionsContext) {
}

// EnterGrant_spectrum_integration_external_schema_permissions is called when production grant_spectrum_integration_external_schema_permissions is entered.
func (s *BaseRedshiftParserListener) EnterGrant_spectrum_integration_external_schema_permissions(ctx *Grant_spectrum_integration_external_schema_permissionsContext) {
}

// ExitGrant_spectrum_integration_external_schema_permissions is called when production grant_spectrum_integration_external_schema_permissions is exited.
func (s *BaseRedshiftParserListener) ExitGrant_spectrum_integration_external_schema_permissions(ctx *Grant_spectrum_integration_external_schema_permissionsContext) {
}

// EnterSpectrum_integration_external_schema_permission_list is called when production spectrum_integration_external_schema_permission_list is entered.
func (s *BaseRedshiftParserListener) EnterSpectrum_integration_external_schema_permission_list(ctx *Spectrum_integration_external_schema_permission_listContext) {
}

// ExitSpectrum_integration_external_schema_permission_list is called when production spectrum_integration_external_schema_permission_list is exited.
func (s *BaseRedshiftParserListener) ExitSpectrum_integration_external_schema_permission_list(ctx *Spectrum_integration_external_schema_permission_listContext) {
}

// EnterSpectrum_integration_external_schema_permission is called when production spectrum_integration_external_schema_permission is entered.
func (s *BaseRedshiftParserListener) EnterSpectrum_integration_external_schema_permission(ctx *Spectrum_integration_external_schema_permissionContext) {
}

// ExitSpectrum_integration_external_schema_permission is called when production spectrum_integration_external_schema_permission is exited.
func (s *BaseRedshiftParserListener) ExitSpectrum_integration_external_schema_permission(ctx *Spectrum_integration_external_schema_permissionContext) {
}

// EnterGrant_spectrum_integration_external_table_permissions is called when production grant_spectrum_integration_external_table_permissions is entered.
func (s *BaseRedshiftParserListener) EnterGrant_spectrum_integration_external_table_permissions(ctx *Grant_spectrum_integration_external_table_permissionsContext) {
}

// ExitGrant_spectrum_integration_external_table_permissions is called when production grant_spectrum_integration_external_table_permissions is exited.
func (s *BaseRedshiftParserListener) ExitGrant_spectrum_integration_external_table_permissions(ctx *Grant_spectrum_integration_external_table_permissionsContext) {
}

// EnterSpectrum_integration_external_table_permission is called when production spectrum_integration_external_table_permission is entered.
func (s *BaseRedshiftParserListener) EnterSpectrum_integration_external_table_permission(ctx *Spectrum_integration_external_table_permissionContext) {
}

// ExitSpectrum_integration_external_table_permission is called when production spectrum_integration_external_table_permission is exited.
func (s *BaseRedshiftParserListener) ExitSpectrum_integration_external_table_permission(ctx *Spectrum_integration_external_table_permissionContext) {
}

// EnterSpectrum_integration_external_table_permission_list is called when production spectrum_integration_external_table_permission_list is entered.
func (s *BaseRedshiftParserListener) EnterSpectrum_integration_external_table_permission_list(ctx *Spectrum_integration_external_table_permission_listContext) {
}

// ExitSpectrum_integration_external_table_permission_list is called when production spectrum_integration_external_table_permission_list is exited.
func (s *BaseRedshiftParserListener) ExitSpectrum_integration_external_table_permission_list(ctx *Spectrum_integration_external_table_permission_listContext) {
}

// EnterGrant_spectrum_integration_extenral_column_permissions is called when production grant_spectrum_integration_extenral_column_permissions is entered.
func (s *BaseRedshiftParserListener) EnterGrant_spectrum_integration_extenral_column_permissions(ctx *Grant_spectrum_integration_extenral_column_permissionsContext) {
}

// ExitGrant_spectrum_integration_extenral_column_permissions is called when production grant_spectrum_integration_extenral_column_permissions is exited.
func (s *BaseRedshiftParserListener) ExitGrant_spectrum_integration_extenral_column_permissions(ctx *Grant_spectrum_integration_extenral_column_permissionsContext) {
}

// EnterIamrolelist_or_public is called when production iamrolelist_or_public is entered.
func (s *BaseRedshiftParserListener) EnterIamrolelist_or_public(ctx *Iamrolelist_or_publicContext) {}

// ExitIamrolelist_or_public is called when production iamrolelist_or_public is exited.
func (s *BaseRedshiftParserListener) ExitIamrolelist_or_public(ctx *Iamrolelist_or_publicContext) {}

// EnterIamrolelist is called when production iamrolelist is entered.
func (s *BaseRedshiftParserListener) EnterIamrolelist(ctx *IamrolelistContext) {}

// ExitIamrolelist is called when production iamrolelist is exited.
func (s *BaseRedshiftParserListener) ExitIamrolelist(ctx *IamrolelistContext) {}

// EnterGrant_assume_role_permissions is called when production grant_assume_role_permissions is entered.
func (s *BaseRedshiftParserListener) EnterGrant_assume_role_permissions(ctx *Grant_assume_role_permissionsContext) {
}

// ExitGrant_assume_role_permissions is called when production grant_assume_role_permissions is exited.
func (s *BaseRedshiftParserListener) ExitGrant_assume_role_permissions(ctx *Grant_assume_role_permissionsContext) {
}

// EnterGrant_assume_role_for_list is called when production grant_assume_role_for_list is entered.
func (s *BaseRedshiftParserListener) EnterGrant_assume_role_for_list(ctx *Grant_assume_role_for_listContext) {
}

// ExitGrant_assume_role_for_list is called when production grant_assume_role_for_list is exited.
func (s *BaseRedshiftParserListener) ExitGrant_assume_role_for_list(ctx *Grant_assume_role_for_listContext) {
}

// EnterGrant_assume_role_for_item is called when production grant_assume_role_for_item is entered.
func (s *BaseRedshiftParserListener) EnterGrant_assume_role_for_item(ctx *Grant_assume_role_for_itemContext) {
}

// ExitGrant_assume_role_for_item is called when production grant_assume_role_for_item is exited.
func (s *BaseRedshiftParserListener) ExitGrant_assume_role_for_item(ctx *Grant_assume_role_for_itemContext) {
}

// EnterGrant_assume_role_target is called when production grant_assume_role_target is entered.
func (s *BaseRedshiftParserListener) EnterGrant_assume_role_target(ctx *Grant_assume_role_targetContext) {
}

// ExitGrant_assume_role_target is called when production grant_assume_role_target is exited.
func (s *BaseRedshiftParserListener) ExitGrant_assume_role_target(ctx *Grant_assume_role_targetContext) {
}

// EnterGrant_column_level_permissions is called when production grant_column_level_permissions is entered.
func (s *BaseRedshiftParserListener) EnterGrant_column_level_permissions(ctx *Grant_column_level_permissionsContext) {
}

// ExitGrant_column_level_permissions is called when production grant_column_level_permissions is exited.
func (s *BaseRedshiftParserListener) ExitGrant_column_level_permissions(ctx *Grant_column_level_permissionsContext) {
}

// EnterColumn_privilege_target is called when production column_privilege_target is entered.
func (s *BaseRedshiftParserListener) EnterColumn_privilege_target(ctx *Column_privilege_targetContext) {
}

// ExitColumn_privilege_target is called when production column_privilege_target is exited.
func (s *BaseRedshiftParserListener) ExitColumn_privilege_target(ctx *Column_privilege_targetContext) {
}

// EnterColumn_privilege_list is called when production column_privilege_list is entered.
func (s *BaseRedshiftParserListener) EnterColumn_privilege_list(ctx *Column_privilege_listContext) {}

// ExitColumn_privilege_list is called when production column_privilege_list is exited.
func (s *BaseRedshiftParserListener) ExitColumn_privilege_list(ctx *Column_privilege_listContext) {}

// EnterColumn_all_privilege is called when production column_all_privilege is entered.
func (s *BaseRedshiftParserListener) EnterColumn_all_privilege(ctx *Column_all_privilegeContext) {}

// ExitColumn_all_privilege is called when production column_all_privilege is exited.
func (s *BaseRedshiftParserListener) ExitColumn_all_privilege(ctx *Column_all_privilegeContext) {}

// EnterColumn_select_update_privilege is called when production column_select_update_privilege is entered.
func (s *BaseRedshiftParserListener) EnterColumn_select_update_privilege(ctx *Column_select_update_privilegeContext) {
}

// ExitColumn_select_update_privilege is called when production column_select_update_privilege is exited.
func (s *BaseRedshiftParserListener) ExitColumn_select_update_privilege(ctx *Column_select_update_privilegeContext) {
}

// EnterCommon_grant is called when production common_grant is entered.
func (s *BaseRedshiftParserListener) EnterCommon_grant(ctx *Common_grantContext) {}

// ExitCommon_grant is called when production common_grant is exited.
func (s *BaseRedshiftParserListener) ExitCommon_grant(ctx *Common_grantContext) {}

// EnterCopy_job_privilege_list is called when production copy_job_privilege_list is entered.
func (s *BaseRedshiftParserListener) EnterCopy_job_privilege_list(ctx *Copy_job_privilege_listContext) {
}

// ExitCopy_job_privilege_list is called when production copy_job_privilege_list is exited.
func (s *BaseRedshiftParserListener) ExitCopy_job_privilege_list(ctx *Copy_job_privilege_listContext) {
}

// EnterCopy_job_privilege is called when production copy_job_privilege is entered.
func (s *BaseRedshiftParserListener) EnterCopy_job_privilege(ctx *Copy_job_privilegeContext) {}

// ExitCopy_job_privilege is called when production copy_job_privilege is exited.
func (s *BaseRedshiftParserListener) ExitCopy_job_privilege(ctx *Copy_job_privilegeContext) {}

// EnterCopy_job_target is called when production copy_job_target is entered.
func (s *BaseRedshiftParserListener) EnterCopy_job_target(ctx *Copy_job_targetContext) {}

// ExitCopy_job_target is called when production copy_job_target is exited.
func (s *BaseRedshiftParserListener) ExitCopy_job_target(ctx *Copy_job_targetContext) {}

// EnterCopy_job_name is called when production copy_job_name is entered.
func (s *BaseRedshiftParserListener) EnterCopy_job_name(ctx *Copy_job_nameContext) {}

// ExitCopy_job_name is called when production copy_job_name is exited.
func (s *BaseRedshiftParserListener) ExitCopy_job_name(ctx *Copy_job_nameContext) {}

// EnterLanguage_privilege_list is called when production language_privilege_list is entered.
func (s *BaseRedshiftParserListener) EnterLanguage_privilege_list(ctx *Language_privilege_listContext) {
}

// ExitLanguage_privilege_list is called when production language_privilege_list is exited.
func (s *BaseRedshiftParserListener) ExitLanguage_privilege_list(ctx *Language_privilege_listContext) {
}

// EnterGrant_language_target is called when production grant_language_target is entered.
func (s *BaseRedshiftParserListener) EnterGrant_language_target(ctx *Grant_language_targetContext) {}

// ExitGrant_language_target is called when production grant_language_target is exited.
func (s *BaseRedshiftParserListener) ExitGrant_language_target(ctx *Grant_language_targetContext) {}

// EnterGrant_procedure_target is called when production grant_procedure_target is entered.
func (s *BaseRedshiftParserListener) EnterGrant_procedure_target(ctx *Grant_procedure_targetContext) {
}

// ExitGrant_procedure_target is called when production grant_procedure_target is exited.
func (s *BaseRedshiftParserListener) ExitGrant_procedure_target(ctx *Grant_procedure_targetContext) {}

// EnterProcedure_privilege_list is called when production procedure_privilege_list is entered.
func (s *BaseRedshiftParserListener) EnterProcedure_privilege_list(ctx *Procedure_privilege_listContext) {
}

// ExitProcedure_privilege_list is called when production procedure_privilege_list is exited.
func (s *BaseRedshiftParserListener) ExitProcedure_privilege_list(ctx *Procedure_privilege_listContext) {
}

// EnterProcedure_privilege is called when production procedure_privilege is entered.
func (s *BaseRedshiftParserListener) EnterProcedure_privilege(ctx *Procedure_privilegeContext) {}

// ExitProcedure_privilege is called when production procedure_privilege is exited.
func (s *BaseRedshiftParserListener) ExitProcedure_privilege(ctx *Procedure_privilegeContext) {}

// EnterFunction_privilege_list is called when production function_privilege_list is entered.
func (s *BaseRedshiftParserListener) EnterFunction_privilege_list(ctx *Function_privilege_listContext) {
}

// ExitFunction_privilege_list is called when production function_privilege_list is exited.
func (s *BaseRedshiftParserListener) ExitFunction_privilege_list(ctx *Function_privilege_listContext) {
}

// EnterFunction_privilege is called when production function_privilege is entered.
func (s *BaseRedshiftParserListener) EnterFunction_privilege(ctx *Function_privilegeContext) {}

// ExitFunction_privilege is called when production function_privilege is exited.
func (s *BaseRedshiftParserListener) ExitFunction_privilege(ctx *Function_privilegeContext) {}

// EnterGrant_function_target is called when production grant_function_target is entered.
func (s *BaseRedshiftParserListener) EnterGrant_function_target(ctx *Grant_function_targetContext) {}

// ExitGrant_function_target is called when production grant_function_target is exited.
func (s *BaseRedshiftParserListener) ExitGrant_function_target(ctx *Grant_function_targetContext) {}

// EnterGrant_schema_target is called when production grant_schema_target is entered.
func (s *BaseRedshiftParserListener) EnterGrant_schema_target(ctx *Grant_schema_targetContext) {}

// ExitGrant_schema_target is called when production grant_schema_target is exited.
func (s *BaseRedshiftParserListener) ExitGrant_schema_target(ctx *Grant_schema_targetContext) {}

// EnterRevoke_schema_target is called when production revoke_schema_target is entered.
func (s *BaseRedshiftParserListener) EnterRevoke_schema_target(ctx *Revoke_schema_targetContext) {}

// ExitRevoke_schema_target is called when production revoke_schema_target is exited.
func (s *BaseRedshiftParserListener) ExitRevoke_schema_target(ctx *Revoke_schema_targetContext) {}

// EnterSchema_privilege_list is called when production schema_privilege_list is entered.
func (s *BaseRedshiftParserListener) EnterSchema_privilege_list(ctx *Schema_privilege_listContext) {}

// ExitSchema_privilege_list is called when production schema_privilege_list is exited.
func (s *BaseRedshiftParserListener) ExitSchema_privilege_list(ctx *Schema_privilege_listContext) {}

// EnterSchema_privilege is called when production schema_privilege is entered.
func (s *BaseRedshiftParserListener) EnterSchema_privilege(ctx *Schema_privilegeContext) {}

// ExitSchema_privilege is called when production schema_privilege is exited.
func (s *BaseRedshiftParserListener) ExitSchema_privilege(ctx *Schema_privilegeContext) {}

// EnterDatabase_privilege_list is called when production database_privilege_list is entered.
func (s *BaseRedshiftParserListener) EnterDatabase_privilege_list(ctx *Database_privilege_listContext) {
}

// ExitDatabase_privilege_list is called when production database_privilege_list is exited.
func (s *BaseRedshiftParserListener) ExitDatabase_privilege_list(ctx *Database_privilege_listContext) {
}

// EnterDatabase_privilege is called when production database_privilege is entered.
func (s *BaseRedshiftParserListener) EnterDatabase_privilege(ctx *Database_privilegeContext) {}

// ExitDatabase_privilege is called when production database_privilege is exited.
func (s *BaseRedshiftParserListener) ExitDatabase_privilege(ctx *Database_privilegeContext) {}

// EnterGrant_database_target is called when production grant_database_target is entered.
func (s *BaseRedshiftParserListener) EnterGrant_database_target(ctx *Grant_database_targetContext) {}

// ExitGrant_database_target is called when production grant_database_target is exited.
func (s *BaseRedshiftParserListener) ExitGrant_database_target(ctx *Grant_database_targetContext) {}

// EnterGrant_table_target is called when production grant_table_target is entered.
func (s *BaseRedshiftParserListener) EnterGrant_table_target(ctx *Grant_table_targetContext) {}

// ExitGrant_table_target is called when production grant_table_target is exited.
func (s *BaseRedshiftParserListener) ExitGrant_table_target(ctx *Grant_table_targetContext) {}

// EnterRevoke_table_target is called when production revoke_table_target is entered.
func (s *BaseRedshiftParserListener) EnterRevoke_table_target(ctx *Revoke_table_targetContext) {}

// ExitRevoke_table_target is called when production revoke_table_target is exited.
func (s *BaseRedshiftParserListener) ExitRevoke_table_target(ctx *Revoke_table_targetContext) {}

// EnterAll_tables_in_schema_list is called when production all_tables_in_schema_list is entered.
func (s *BaseRedshiftParserListener) EnterAll_tables_in_schema_list(ctx *All_tables_in_schema_listContext) {
}

// ExitAll_tables_in_schema_list is called when production all_tables_in_schema_list is exited.
func (s *BaseRedshiftParserListener) ExitAll_tables_in_schema_list(ctx *All_tables_in_schema_listContext) {
}

// EnterAll_privileges is called when production all_privileges is entered.
func (s *BaseRedshiftParserListener) EnterAll_privileges(ctx *All_privilegesContext) {}

// ExitAll_privileges is called when production all_privileges is exited.
func (s *BaseRedshiftParserListener) ExitAll_privileges(ctx *All_privilegesContext) {}

// EnterGrantee_list is called when production grantee_list is entered.
func (s *BaseRedshiftParserListener) EnterGrantee_list(ctx *Grantee_listContext) {}

// ExitGrantee_list is called when production grantee_list is exited.
func (s *BaseRedshiftParserListener) ExitGrantee_list(ctx *Grantee_listContext) {}

// EnterGrantee is called when production grantee is entered.
func (s *BaseRedshiftParserListener) EnterGrantee(ctx *GranteeContext) {}

// ExitGrantee is called when production grantee is exited.
func (s *BaseRedshiftParserListener) ExitGrantee(ctx *GranteeContext) {}

// EnterOpt_with_grant_option is called when production opt_with_grant_option is entered.
func (s *BaseRedshiftParserListener) EnterOpt_with_grant_option(ctx *Opt_with_grant_optionContext) {}

// ExitOpt_with_grant_option is called when production opt_with_grant_option is exited.
func (s *BaseRedshiftParserListener) ExitOpt_with_grant_option(ctx *Opt_with_grant_optionContext) {}

// EnterTable_privilege is called when production table_privilege is entered.
func (s *BaseRedshiftParserListener) EnterTable_privilege(ctx *Table_privilegeContext) {}

// ExitTable_privilege is called when production table_privilege is exited.
func (s *BaseRedshiftParserListener) ExitTable_privilege(ctx *Table_privilegeContext) {}

// EnterTable_privilege_list is called when production table_privilege_list is entered.
func (s *BaseRedshiftParserListener) EnterTable_privilege_list(ctx *Table_privilege_listContext) {}

// ExitTable_privilege_list is called when production table_privilege_list is exited.
func (s *BaseRedshiftParserListener) ExitTable_privilege_list(ctx *Table_privilege_listContext) {}

// EnterRevokestmt is called when production revokestmt is entered.
func (s *BaseRedshiftParserListener) EnterRevokestmt(ctx *RevokestmtContext) {}

// ExitRevokestmt is called when production revokestmt is exited.
func (s *BaseRedshiftParserListener) ExitRevokestmt(ctx *RevokestmtContext) {}

// EnterRevoke_permissions_for_rls_lookup_tables is called when production revoke_permissions_for_rls_lookup_tables is entered.
func (s *BaseRedshiftParserListener) EnterRevoke_permissions_for_rls_lookup_tables(ctx *Revoke_permissions_for_rls_lookup_tablesContext) {
}

// ExitRevoke_permissions_for_rls_lookup_tables is called when production revoke_permissions_for_rls_lookup_tables is exited.
func (s *BaseRedshiftParserListener) ExitRevoke_permissions_for_rls_lookup_tables(ctx *Revoke_permissions_for_rls_lookup_tablesContext) {
}

// EnterRevoke_explain_permissions_for_row_level_security_policy_filters is called when production revoke_explain_permissions_for_row_level_security_policy_filters is entered.
func (s *BaseRedshiftParserListener) EnterRevoke_explain_permissions_for_row_level_security_policy_filters(ctx *Revoke_explain_permissions_for_row_level_security_policy_filtersContext) {
}

// ExitRevoke_explain_permissions_for_row_level_security_policy_filters is called when production revoke_explain_permissions_for_row_level_security_policy_filters is exited.
func (s *BaseRedshiftParserListener) ExitRevoke_explain_permissions_for_row_level_security_policy_filters(ctx *Revoke_explain_permissions_for_row_level_security_policy_filtersContext) {
}

// EnterRevoke_machine_learning_permissions is called when production revoke_machine_learning_permissions is entered.
func (s *BaseRedshiftParserListener) EnterRevoke_machine_learning_permissions(ctx *Revoke_machine_learning_permissionsContext) {
}

// ExitRevoke_machine_learning_permissions is called when production revoke_machine_learning_permissions is exited.
func (s *BaseRedshiftParserListener) ExitRevoke_machine_learning_permissions(ctx *Revoke_machine_learning_permissionsContext) {
}

// EnterRevoke_role_permissions is called when production revoke_role_permissions is entered.
func (s *BaseRedshiftParserListener) EnterRevoke_role_permissions(ctx *Revoke_role_permissionsContext) {
}

// ExitRevoke_role_permissions is called when production revoke_role_permissions is exited.
func (s *BaseRedshiftParserListener) ExitRevoke_role_permissions(ctx *Revoke_role_permissionsContext) {
}

// EnterRevoke_scoped_permissions is called when production revoke_scoped_permissions is entered.
func (s *BaseRedshiftParserListener) EnterRevoke_scoped_permissions(ctx *Revoke_scoped_permissionsContext) {
}

// ExitRevoke_scoped_permissions is called when production revoke_scoped_permissions is exited.
func (s *BaseRedshiftParserListener) ExitRevoke_scoped_permissions(ctx *Revoke_scoped_permissionsContext) {
}

// EnterRevoke_scoped_schemas_permissions is called when production revoke_scoped_schemas_permissions is entered.
func (s *BaseRedshiftParserListener) EnterRevoke_scoped_schemas_permissions(ctx *Revoke_scoped_schemas_permissionsContext) {
}

// ExitRevoke_scoped_schemas_permissions is called when production revoke_scoped_schemas_permissions is exited.
func (s *BaseRedshiftParserListener) ExitRevoke_scoped_schemas_permissions(ctx *Revoke_scoped_schemas_permissionsContext) {
}

// EnterRevoke_scoped_tables_permissions is called when production revoke_scoped_tables_permissions is entered.
func (s *BaseRedshiftParserListener) EnterRevoke_scoped_tables_permissions(ctx *Revoke_scoped_tables_permissionsContext) {
}

// ExitRevoke_scoped_tables_permissions is called when production revoke_scoped_tables_permissions is exited.
func (s *BaseRedshiftParserListener) ExitRevoke_scoped_tables_permissions(ctx *Revoke_scoped_tables_permissionsContext) {
}

// EnterRevoke_scoped_functions_permissions is called when production revoke_scoped_functions_permissions is entered.
func (s *BaseRedshiftParserListener) EnterRevoke_scoped_functions_permissions(ctx *Revoke_scoped_functions_permissionsContext) {
}

// ExitRevoke_scoped_functions_permissions is called when production revoke_scoped_functions_permissions is exited.
func (s *BaseRedshiftParserListener) ExitRevoke_scoped_functions_permissions(ctx *Revoke_scoped_functions_permissionsContext) {
}

// EnterRevoke_scoped_procedures_permissions is called when production revoke_scoped_procedures_permissions is entered.
func (s *BaseRedshiftParserListener) EnterRevoke_scoped_procedures_permissions(ctx *Revoke_scoped_procedures_permissionsContext) {
}

// ExitRevoke_scoped_procedures_permissions is called when production revoke_scoped_procedures_permissions is exited.
func (s *BaseRedshiftParserListener) ExitRevoke_scoped_procedures_permissions(ctx *Revoke_scoped_procedures_permissionsContext) {
}

// EnterRevoke_scoped_languages_permissions is called when production revoke_scoped_languages_permissions is entered.
func (s *BaseRedshiftParserListener) EnterRevoke_scoped_languages_permissions(ctx *Revoke_scoped_languages_permissionsContext) {
}

// ExitRevoke_scoped_languages_permissions is called when production revoke_scoped_languages_permissions is exited.
func (s *BaseRedshiftParserListener) ExitRevoke_scoped_languages_permissions(ctx *Revoke_scoped_languages_permissionsContext) {
}

// EnterRevoke_scoped_copy_jobs_permissions is called when production revoke_scoped_copy_jobs_permissions is entered.
func (s *BaseRedshiftParserListener) EnterRevoke_scoped_copy_jobs_permissions(ctx *Revoke_scoped_copy_jobs_permissionsContext) {
}

// ExitRevoke_scoped_copy_jobs_permissions is called when production revoke_scoped_copy_jobs_permissions is exited.
func (s *BaseRedshiftParserListener) ExitRevoke_scoped_copy_jobs_permissions(ctx *Revoke_scoped_copy_jobs_permissionsContext) {
}

// EnterRevoke_datashare_permissions is called when production revoke_datashare_permissions is entered.
func (s *BaseRedshiftParserListener) EnterRevoke_datashare_permissions(ctx *Revoke_datashare_permissionsContext) {
}

// ExitRevoke_datashare_permissions is called when production revoke_datashare_permissions is exited.
func (s *BaseRedshiftParserListener) ExitRevoke_datashare_permissions(ctx *Revoke_datashare_permissionsContext) {
}

// EnterRevoke_spectrum_integration_permissions is called when production revoke_spectrum_integration_permissions is entered.
func (s *BaseRedshiftParserListener) EnterRevoke_spectrum_integration_permissions(ctx *Revoke_spectrum_integration_permissionsContext) {
}

// ExitRevoke_spectrum_integration_permissions is called when production revoke_spectrum_integration_permissions is exited.
func (s *BaseRedshiftParserListener) ExitRevoke_spectrum_integration_permissions(ctx *Revoke_spectrum_integration_permissionsContext) {
}

// EnterRevoke_spectrum_integration_external_schema_permissions is called when production revoke_spectrum_integration_external_schema_permissions is entered.
func (s *BaseRedshiftParserListener) EnterRevoke_spectrum_integration_external_schema_permissions(ctx *Revoke_spectrum_integration_external_schema_permissionsContext) {
}

// ExitRevoke_spectrum_integration_external_schema_permissions is called when production revoke_spectrum_integration_external_schema_permissions is exited.
func (s *BaseRedshiftParserListener) ExitRevoke_spectrum_integration_external_schema_permissions(ctx *Revoke_spectrum_integration_external_schema_permissionsContext) {
}

// EnterRevoke_spectrum_integration_external_table_permissions is called when production revoke_spectrum_integration_external_table_permissions is entered.
func (s *BaseRedshiftParserListener) EnterRevoke_spectrum_integration_external_table_permissions(ctx *Revoke_spectrum_integration_external_table_permissionsContext) {
}

// ExitRevoke_spectrum_integration_external_table_permissions is called when production revoke_spectrum_integration_external_table_permissions is exited.
func (s *BaseRedshiftParserListener) ExitRevoke_spectrum_integration_external_table_permissions(ctx *Revoke_spectrum_integration_external_table_permissionsContext) {
}

// EnterRevoke_spectrum_integration_extenral_column_permissions is called when production revoke_spectrum_integration_extenral_column_permissions is entered.
func (s *BaseRedshiftParserListener) EnterRevoke_spectrum_integration_extenral_column_permissions(ctx *Revoke_spectrum_integration_extenral_column_permissionsContext) {
}

// ExitRevoke_spectrum_integration_extenral_column_permissions is called when production revoke_spectrum_integration_extenral_column_permissions is exited.
func (s *BaseRedshiftParserListener) ExitRevoke_spectrum_integration_extenral_column_permissions(ctx *Revoke_spectrum_integration_extenral_column_permissionsContext) {
}

// EnterRevoke_assume_role_permissions is called when production revoke_assume_role_permissions is entered.
func (s *BaseRedshiftParserListener) EnterRevoke_assume_role_permissions(ctx *Revoke_assume_role_permissionsContext) {
}

// ExitRevoke_assume_role_permissions is called when production revoke_assume_role_permissions is exited.
func (s *BaseRedshiftParserListener) ExitRevoke_assume_role_permissions(ctx *Revoke_assume_role_permissionsContext) {
}

// EnterRevoke_column_level_permissions is called when production revoke_column_level_permissions is entered.
func (s *BaseRedshiftParserListener) EnterRevoke_column_level_permissions(ctx *Revoke_column_level_permissionsContext) {
}

// ExitRevoke_column_level_permissions is called when production revoke_column_level_permissions is exited.
func (s *BaseRedshiftParserListener) ExitRevoke_column_level_permissions(ctx *Revoke_column_level_permissionsContext) {
}

// EnterCommon_revoke is called when production common_revoke is entered.
func (s *BaseRedshiftParserListener) EnterCommon_revoke(ctx *Common_revokeContext) {}

// ExitCommon_revoke is called when production common_revoke is exited.
func (s *BaseRedshiftParserListener) ExitCommon_revoke(ctx *Common_revokeContext) {}

// EnterPrivileges is called when production privileges is entered.
func (s *BaseRedshiftParserListener) EnterPrivileges(ctx *PrivilegesContext) {}

// ExitPrivileges is called when production privileges is exited.
func (s *BaseRedshiftParserListener) ExitPrivileges(ctx *PrivilegesContext) {}

// EnterPrivilege_list is called when production privilege_list is entered.
func (s *BaseRedshiftParserListener) EnterPrivilege_list(ctx *Privilege_listContext) {}

// ExitPrivilege_list is called when production privilege_list is exited.
func (s *BaseRedshiftParserListener) ExitPrivilege_list(ctx *Privilege_listContext) {}

// EnterPrivilege is called when production privilege is entered.
func (s *BaseRedshiftParserListener) EnterPrivilege(ctx *PrivilegeContext) {}

// ExitPrivilege is called when production privilege is exited.
func (s *BaseRedshiftParserListener) ExitPrivilege(ctx *PrivilegeContext) {}

// EnterPrivilege_target is called when production privilege_target is entered.
func (s *BaseRedshiftParserListener) EnterPrivilege_target(ctx *Privilege_targetContext) {}

// ExitPrivilege_target is called when production privilege_target is exited.
func (s *BaseRedshiftParserListener) ExitPrivilege_target(ctx *Privilege_targetContext) {}

// EnterParameter_name_list is called when production parameter_name_list is entered.
func (s *BaseRedshiftParserListener) EnterParameter_name_list(ctx *Parameter_name_listContext) {}

// ExitParameter_name_list is called when production parameter_name_list is exited.
func (s *BaseRedshiftParserListener) ExitParameter_name_list(ctx *Parameter_name_listContext) {}

// EnterParameter_name is called when production parameter_name is entered.
func (s *BaseRedshiftParserListener) EnterParameter_name(ctx *Parameter_nameContext) {}

// ExitParameter_name is called when production parameter_name is exited.
func (s *BaseRedshiftParserListener) ExitParameter_name(ctx *Parameter_nameContext) {}

// EnterOpt_grant_grant_option is called when production opt_grant_grant_option is entered.
func (s *BaseRedshiftParserListener) EnterOpt_grant_grant_option(ctx *Opt_grant_grant_optionContext) {
}

// ExitOpt_grant_grant_option is called when production opt_grant_grant_option is exited.
func (s *BaseRedshiftParserListener) ExitOpt_grant_grant_option(ctx *Opt_grant_grant_optionContext) {}

// EnterGrantrolestmt is called when production grantrolestmt is entered.
func (s *BaseRedshiftParserListener) EnterGrantrolestmt(ctx *GrantrolestmtContext) {}

// ExitGrantrolestmt is called when production grantrolestmt is exited.
func (s *BaseRedshiftParserListener) ExitGrantrolestmt(ctx *GrantrolestmtContext) {}

// EnterRevokerolestmt is called when production revokerolestmt is entered.
func (s *BaseRedshiftParserListener) EnterRevokerolestmt(ctx *RevokerolestmtContext) {}

// ExitRevokerolestmt is called when production revokerolestmt is exited.
func (s *BaseRedshiftParserListener) ExitRevokerolestmt(ctx *RevokerolestmtContext) {}

// EnterOpt_grant_admin_option is called when production opt_grant_admin_option is entered.
func (s *BaseRedshiftParserListener) EnterOpt_grant_admin_option(ctx *Opt_grant_admin_optionContext) {
}

// ExitOpt_grant_admin_option is called when production opt_grant_admin_option is exited.
func (s *BaseRedshiftParserListener) ExitOpt_grant_admin_option(ctx *Opt_grant_admin_optionContext) {}

// EnterOpt_granted_by is called when production opt_granted_by is entered.
func (s *BaseRedshiftParserListener) EnterOpt_granted_by(ctx *Opt_granted_byContext) {}

// ExitOpt_granted_by is called when production opt_granted_by is exited.
func (s *BaseRedshiftParserListener) ExitOpt_granted_by(ctx *Opt_granted_byContext) {}

// EnterAlterdefaultprivilegesstmt is called when production alterdefaultprivilegesstmt is entered.
func (s *BaseRedshiftParserListener) EnterAlterdefaultprivilegesstmt(ctx *AlterdefaultprivilegesstmtContext) {
}

// ExitAlterdefaultprivilegesstmt is called when production alterdefaultprivilegesstmt is exited.
func (s *BaseRedshiftParserListener) ExitAlterdefaultprivilegesstmt(ctx *AlterdefaultprivilegesstmtContext) {
}

// EnterDefacloptionlist is called when production defacloptionlist is entered.
func (s *BaseRedshiftParserListener) EnterDefacloptionlist(ctx *DefacloptionlistContext) {}

// ExitDefacloptionlist is called when production defacloptionlist is exited.
func (s *BaseRedshiftParserListener) ExitDefacloptionlist(ctx *DefacloptionlistContext) {}

// EnterDefacloption is called when production defacloption is entered.
func (s *BaseRedshiftParserListener) EnterDefacloption(ctx *DefacloptionContext) {}

// ExitDefacloption is called when production defacloption is exited.
func (s *BaseRedshiftParserListener) ExitDefacloption(ctx *DefacloptionContext) {}

// EnterDefaclaction is called when production defaclaction is entered.
func (s *BaseRedshiftParserListener) EnterDefaclaction(ctx *DefaclactionContext) {}

// ExitDefaclaction is called when production defaclaction is exited.
func (s *BaseRedshiftParserListener) ExitDefaclaction(ctx *DefaclactionContext) {}

// EnterDefacl_privilege_target is called when production defacl_privilege_target is entered.
func (s *BaseRedshiftParserListener) EnterDefacl_privilege_target(ctx *Defacl_privilege_targetContext) {
}

// ExitDefacl_privilege_target is called when production defacl_privilege_target is exited.
func (s *BaseRedshiftParserListener) ExitDefacl_privilege_target(ctx *Defacl_privilege_targetContext) {
}

// EnterIndexstmt is called when production indexstmt is entered.
func (s *BaseRedshiftParserListener) EnterIndexstmt(ctx *IndexstmtContext) {}

// ExitIndexstmt is called when production indexstmt is exited.
func (s *BaseRedshiftParserListener) ExitIndexstmt(ctx *IndexstmtContext) {}

// EnterOpt_unique is called when production opt_unique is entered.
func (s *BaseRedshiftParserListener) EnterOpt_unique(ctx *Opt_uniqueContext) {}

// ExitOpt_unique is called when production opt_unique is exited.
func (s *BaseRedshiftParserListener) ExitOpt_unique(ctx *Opt_uniqueContext) {}

// EnterOpt_concurrently is called when production opt_concurrently is entered.
func (s *BaseRedshiftParserListener) EnterOpt_concurrently(ctx *Opt_concurrentlyContext) {}

// ExitOpt_concurrently is called when production opt_concurrently is exited.
func (s *BaseRedshiftParserListener) ExitOpt_concurrently(ctx *Opt_concurrentlyContext) {}

// EnterOpt_index_name is called when production opt_index_name is entered.
func (s *BaseRedshiftParserListener) EnterOpt_index_name(ctx *Opt_index_nameContext) {}

// ExitOpt_index_name is called when production opt_index_name is exited.
func (s *BaseRedshiftParserListener) ExitOpt_index_name(ctx *Opt_index_nameContext) {}

// EnterAccess_method_clause is called when production access_method_clause is entered.
func (s *BaseRedshiftParserListener) EnterAccess_method_clause(ctx *Access_method_clauseContext) {}

// ExitAccess_method_clause is called when production access_method_clause is exited.
func (s *BaseRedshiftParserListener) ExitAccess_method_clause(ctx *Access_method_clauseContext) {}

// EnterIndex_params is called when production index_params is entered.
func (s *BaseRedshiftParserListener) EnterIndex_params(ctx *Index_paramsContext) {}

// ExitIndex_params is called when production index_params is exited.
func (s *BaseRedshiftParserListener) ExitIndex_params(ctx *Index_paramsContext) {}

// EnterIndex_elem_options is called when production index_elem_options is entered.
func (s *BaseRedshiftParserListener) EnterIndex_elem_options(ctx *Index_elem_optionsContext) {}

// ExitIndex_elem_options is called when production index_elem_options is exited.
func (s *BaseRedshiftParserListener) ExitIndex_elem_options(ctx *Index_elem_optionsContext) {}

// EnterIndex_elem is called when production index_elem is entered.
func (s *BaseRedshiftParserListener) EnterIndex_elem(ctx *Index_elemContext) {}

// ExitIndex_elem is called when production index_elem is exited.
func (s *BaseRedshiftParserListener) ExitIndex_elem(ctx *Index_elemContext) {}

// EnterOpt_include is called when production opt_include is entered.
func (s *BaseRedshiftParserListener) EnterOpt_include(ctx *Opt_includeContext) {}

// ExitOpt_include is called when production opt_include is exited.
func (s *BaseRedshiftParserListener) ExitOpt_include(ctx *Opt_includeContext) {}

// EnterIndex_including_params is called when production index_including_params is entered.
func (s *BaseRedshiftParserListener) EnterIndex_including_params(ctx *Index_including_paramsContext) {
}

// ExitIndex_including_params is called when production index_including_params is exited.
func (s *BaseRedshiftParserListener) ExitIndex_including_params(ctx *Index_including_paramsContext) {}

// EnterOpt_collate is called when production opt_collate is entered.
func (s *BaseRedshiftParserListener) EnterOpt_collate(ctx *Opt_collateContext) {}

// ExitOpt_collate is called when production opt_collate is exited.
func (s *BaseRedshiftParserListener) ExitOpt_collate(ctx *Opt_collateContext) {}

// EnterOpt_class is called when production opt_class is entered.
func (s *BaseRedshiftParserListener) EnterOpt_class(ctx *Opt_classContext) {}

// ExitOpt_class is called when production opt_class is exited.
func (s *BaseRedshiftParserListener) ExitOpt_class(ctx *Opt_classContext) {}

// EnterOpt_asc_desc is called when production opt_asc_desc is entered.
func (s *BaseRedshiftParserListener) EnterOpt_asc_desc(ctx *Opt_asc_descContext) {}

// ExitOpt_asc_desc is called when production opt_asc_desc is exited.
func (s *BaseRedshiftParserListener) ExitOpt_asc_desc(ctx *Opt_asc_descContext) {}

// EnterOpt_nulls_order is called when production opt_nulls_order is entered.
func (s *BaseRedshiftParserListener) EnterOpt_nulls_order(ctx *Opt_nulls_orderContext) {}

// ExitOpt_nulls_order is called when production opt_nulls_order is exited.
func (s *BaseRedshiftParserListener) ExitOpt_nulls_order(ctx *Opt_nulls_orderContext) {}

// EnterCreatefunctionstmt is called when production createfunctionstmt is entered.
func (s *BaseRedshiftParserListener) EnterCreatefunctionstmt(ctx *CreatefunctionstmtContext) {}

// ExitCreatefunctionstmt is called when production createfunctionstmt is exited.
func (s *BaseRedshiftParserListener) ExitCreatefunctionstmt(ctx *CreatefunctionstmtContext) {}

// EnterCreateprocedurestmt is called when production createprocedurestmt is entered.
func (s *BaseRedshiftParserListener) EnterCreateprocedurestmt(ctx *CreateprocedurestmtContext) {}

// ExitCreateprocedurestmt is called when production createprocedurestmt is exited.
func (s *BaseRedshiftParserListener) ExitCreateprocedurestmt(ctx *CreateprocedurestmtContext) {}

// EnterOpt_nonatomic is called when production opt_nonatomic is entered.
func (s *BaseRedshiftParserListener) EnterOpt_nonatomic(ctx *Opt_nonatomicContext) {}

// ExitOpt_nonatomic is called when production opt_nonatomic is exited.
func (s *BaseRedshiftParserListener) ExitOpt_nonatomic(ctx *Opt_nonatomicContext) {}

// EnterOpt_or_replace is called when production opt_or_replace is entered.
func (s *BaseRedshiftParserListener) EnterOpt_or_replace(ctx *Opt_or_replaceContext) {}

// ExitOpt_or_replace is called when production opt_or_replace is exited.
func (s *BaseRedshiftParserListener) ExitOpt_or_replace(ctx *Opt_or_replaceContext) {}

// EnterFunc_py_args_or_sql_args is called when production func_py_args_or_sql_args is entered.
func (s *BaseRedshiftParserListener) EnterFunc_py_args_or_sql_args(ctx *Func_py_args_or_sql_argsContext) {
}

// ExitFunc_py_args_or_sql_args is called when production func_py_args_or_sql_args is exited.
func (s *BaseRedshiftParserListener) ExitFunc_py_args_or_sql_args(ctx *Func_py_args_or_sql_argsContext) {
}

// EnterFunc_py_args_or_sql_args_list is called when production func_py_args_or_sql_args_list is entered.
func (s *BaseRedshiftParserListener) EnterFunc_py_args_or_sql_args_list(ctx *Func_py_args_or_sql_args_listContext) {
}

// ExitFunc_py_args_or_sql_args_list is called when production func_py_args_or_sql_args_list is exited.
func (s *BaseRedshiftParserListener) ExitFunc_py_args_or_sql_args_list(ctx *Func_py_args_or_sql_args_listContext) {
}

// EnterFunc_args is called when production func_args is entered.
func (s *BaseRedshiftParserListener) EnterFunc_args(ctx *Func_argsContext) {}

// ExitFunc_args is called when production func_args is exited.
func (s *BaseRedshiftParserListener) ExitFunc_args(ctx *Func_argsContext) {}

// EnterFunc_args_list is called when production func_args_list is entered.
func (s *BaseRedshiftParserListener) EnterFunc_args_list(ctx *Func_args_listContext) {}

// ExitFunc_args_list is called when production func_args_list is exited.
func (s *BaseRedshiftParserListener) ExitFunc_args_list(ctx *Func_args_listContext) {}

// EnterFunction_with_argtypes_list is called when production function_with_argtypes_list is entered.
func (s *BaseRedshiftParserListener) EnterFunction_with_argtypes_list(ctx *Function_with_argtypes_listContext) {
}

// ExitFunction_with_argtypes_list is called when production function_with_argtypes_list is exited.
func (s *BaseRedshiftParserListener) ExitFunction_with_argtypes_list(ctx *Function_with_argtypes_listContext) {
}

// EnterFunction_with_argtypes is called when production function_with_argtypes is entered.
func (s *BaseRedshiftParserListener) EnterFunction_with_argtypes(ctx *Function_with_argtypesContext) {
}

// ExitFunction_with_argtypes is called when production function_with_argtypes is exited.
func (s *BaseRedshiftParserListener) ExitFunction_with_argtypes(ctx *Function_with_argtypesContext) {}

// EnterFunc_args_with_defaults is called when production func_args_with_defaults is entered.
func (s *BaseRedshiftParserListener) EnterFunc_args_with_defaults(ctx *Func_args_with_defaultsContext) {
}

// ExitFunc_args_with_defaults is called when production func_args_with_defaults is exited.
func (s *BaseRedshiftParserListener) ExitFunc_args_with_defaults(ctx *Func_args_with_defaultsContext) {
}

// EnterFunc_args_with_defaults_list is called when production func_args_with_defaults_list is entered.
func (s *BaseRedshiftParserListener) EnterFunc_args_with_defaults_list(ctx *Func_args_with_defaults_listContext) {
}

// ExitFunc_args_with_defaults_list is called when production func_args_with_defaults_list is exited.
func (s *BaseRedshiftParserListener) ExitFunc_args_with_defaults_list(ctx *Func_args_with_defaults_listContext) {
}

// EnterFunc_arg is called when production func_arg is entered.
func (s *BaseRedshiftParserListener) EnterFunc_arg(ctx *Func_argContext) {}

// ExitFunc_arg is called when production func_arg is exited.
func (s *BaseRedshiftParserListener) ExitFunc_arg(ctx *Func_argContext) {}

// EnterArg_class is called when production arg_class is entered.
func (s *BaseRedshiftParserListener) EnterArg_class(ctx *Arg_classContext) {}

// ExitArg_class is called when production arg_class is exited.
func (s *BaseRedshiftParserListener) ExitArg_class(ctx *Arg_classContext) {}

// EnterParam_name is called when production param_name is entered.
func (s *BaseRedshiftParserListener) EnterParam_name(ctx *Param_nameContext) {}

// ExitParam_name is called when production param_name is exited.
func (s *BaseRedshiftParserListener) ExitParam_name(ctx *Param_nameContext) {}

// EnterFunc_return is called when production func_return is entered.
func (s *BaseRedshiftParserListener) EnterFunc_return(ctx *Func_returnContext) {}

// ExitFunc_return is called when production func_return is exited.
func (s *BaseRedshiftParserListener) ExitFunc_return(ctx *Func_returnContext) {}

// EnterFunc_type is called when production func_type is entered.
func (s *BaseRedshiftParserListener) EnterFunc_type(ctx *Func_typeContext) {}

// ExitFunc_type is called when production func_type is exited.
func (s *BaseRedshiftParserListener) ExitFunc_type(ctx *Func_typeContext) {}

// EnterFunc_arg_with_default is called when production func_arg_with_default is entered.
func (s *BaseRedshiftParserListener) EnterFunc_arg_with_default(ctx *Func_arg_with_defaultContext) {}

// ExitFunc_arg_with_default is called when production func_arg_with_default is exited.
func (s *BaseRedshiftParserListener) ExitFunc_arg_with_default(ctx *Func_arg_with_defaultContext) {}

// EnterAggr_arg is called when production aggr_arg is entered.
func (s *BaseRedshiftParserListener) EnterAggr_arg(ctx *Aggr_argContext) {}

// ExitAggr_arg is called when production aggr_arg is exited.
func (s *BaseRedshiftParserListener) ExitAggr_arg(ctx *Aggr_argContext) {}

// EnterAggr_args is called when production aggr_args is entered.
func (s *BaseRedshiftParserListener) EnterAggr_args(ctx *Aggr_argsContext) {}

// ExitAggr_args is called when production aggr_args is exited.
func (s *BaseRedshiftParserListener) ExitAggr_args(ctx *Aggr_argsContext) {}

// EnterAggr_args_list is called when production aggr_args_list is entered.
func (s *BaseRedshiftParserListener) EnterAggr_args_list(ctx *Aggr_args_listContext) {}

// ExitAggr_args_list is called when production aggr_args_list is exited.
func (s *BaseRedshiftParserListener) ExitAggr_args_list(ctx *Aggr_args_listContext) {}

// EnterAggregate_with_argtypes is called when production aggregate_with_argtypes is entered.
func (s *BaseRedshiftParserListener) EnterAggregate_with_argtypes(ctx *Aggregate_with_argtypesContext) {
}

// ExitAggregate_with_argtypes is called when production aggregate_with_argtypes is exited.
func (s *BaseRedshiftParserListener) ExitAggregate_with_argtypes(ctx *Aggregate_with_argtypesContext) {
}

// EnterAggregate_with_argtypes_list is called when production aggregate_with_argtypes_list is entered.
func (s *BaseRedshiftParserListener) EnterAggregate_with_argtypes_list(ctx *Aggregate_with_argtypes_listContext) {
}

// ExitAggregate_with_argtypes_list is called when production aggregate_with_argtypes_list is exited.
func (s *BaseRedshiftParserListener) ExitAggregate_with_argtypes_list(ctx *Aggregate_with_argtypes_listContext) {
}

// EnterCreatefunc_opt_list is called when production createfunc_opt_list is entered.
func (s *BaseRedshiftParserListener) EnterCreatefunc_opt_list(ctx *Createfunc_opt_listContext) {}

// ExitCreatefunc_opt_list is called when production createfunc_opt_list is exited.
func (s *BaseRedshiftParserListener) ExitCreatefunc_opt_list(ctx *Createfunc_opt_listContext) {}

// EnterCommon_func_opt_item is called when production common_func_opt_item is entered.
func (s *BaseRedshiftParserListener) EnterCommon_func_opt_item(ctx *Common_func_opt_itemContext) {}

// ExitCommon_func_opt_item is called when production common_func_opt_item is exited.
func (s *BaseRedshiftParserListener) ExitCommon_func_opt_item(ctx *Common_func_opt_itemContext) {}

// EnterCreatefunc_opt_item is called when production createfunc_opt_item is entered.
func (s *BaseRedshiftParserListener) EnterCreatefunc_opt_item(ctx *Createfunc_opt_itemContext) {}

// ExitCreatefunc_opt_item is called when production createfunc_opt_item is exited.
func (s *BaseRedshiftParserListener) ExitCreatefunc_opt_item(ctx *Createfunc_opt_itemContext) {}

// EnterFunc_as is called when production func_as is entered.
func (s *BaseRedshiftParserListener) EnterFunc_as(ctx *Func_asContext) {}

// ExitFunc_as is called when production func_as is exited.
func (s *BaseRedshiftParserListener) ExitFunc_as(ctx *Func_asContext) {}

// EnterTransform_type_list is called when production transform_type_list is entered.
func (s *BaseRedshiftParserListener) EnterTransform_type_list(ctx *Transform_type_listContext) {}

// ExitTransform_type_list is called when production transform_type_list is exited.
func (s *BaseRedshiftParserListener) ExitTransform_type_list(ctx *Transform_type_listContext) {}

// EnterOpt_definition is called when production opt_definition is entered.
func (s *BaseRedshiftParserListener) EnterOpt_definition(ctx *Opt_definitionContext) {}

// ExitOpt_definition is called when production opt_definition is exited.
func (s *BaseRedshiftParserListener) ExitOpt_definition(ctx *Opt_definitionContext) {}

// EnterTable_func_column is called when production table_func_column is entered.
func (s *BaseRedshiftParserListener) EnterTable_func_column(ctx *Table_func_columnContext) {}

// ExitTable_func_column is called when production table_func_column is exited.
func (s *BaseRedshiftParserListener) ExitTable_func_column(ctx *Table_func_columnContext) {}

// EnterTable_func_column_list is called when production table_func_column_list is entered.
func (s *BaseRedshiftParserListener) EnterTable_func_column_list(ctx *Table_func_column_listContext) {
}

// ExitTable_func_column_list is called when production table_func_column_list is exited.
func (s *BaseRedshiftParserListener) ExitTable_func_column_list(ctx *Table_func_column_listContext) {}

// EnterAlterfunctionstmt is called when production alterfunctionstmt is entered.
func (s *BaseRedshiftParserListener) EnterAlterfunctionstmt(ctx *AlterfunctionstmtContext) {}

// ExitAlterfunctionstmt is called when production alterfunctionstmt is exited.
func (s *BaseRedshiftParserListener) ExitAlterfunctionstmt(ctx *AlterfunctionstmtContext) {}

// EnterAlterprocedurestmt is called when production alterprocedurestmt is entered.
func (s *BaseRedshiftParserListener) EnterAlterprocedurestmt(ctx *AlterprocedurestmtContext) {}

// ExitAlterprocedurestmt is called when production alterprocedurestmt is exited.
func (s *BaseRedshiftParserListener) ExitAlterprocedurestmt(ctx *AlterprocedurestmtContext) {}

// EnterAlterfunc_opt_list is called when production alterfunc_opt_list is entered.
func (s *BaseRedshiftParserListener) EnterAlterfunc_opt_list(ctx *Alterfunc_opt_listContext) {}

// ExitAlterfunc_opt_list is called when production alterfunc_opt_list is exited.
func (s *BaseRedshiftParserListener) ExitAlterfunc_opt_list(ctx *Alterfunc_opt_listContext) {}

// EnterOpt_restrict is called when production opt_restrict is entered.
func (s *BaseRedshiftParserListener) EnterOpt_restrict(ctx *Opt_restrictContext) {}

// ExitOpt_restrict is called when production opt_restrict is exited.
func (s *BaseRedshiftParserListener) ExitOpt_restrict(ctx *Opt_restrictContext) {}

// EnterRemovefuncstmt is called when production removefuncstmt is entered.
func (s *BaseRedshiftParserListener) EnterRemovefuncstmt(ctx *RemovefuncstmtContext) {}

// ExitRemovefuncstmt is called when production removefuncstmt is exited.
func (s *BaseRedshiftParserListener) ExitRemovefuncstmt(ctx *RemovefuncstmtContext) {}

// EnterRemoveaggrstmt is called when production removeaggrstmt is entered.
func (s *BaseRedshiftParserListener) EnterRemoveaggrstmt(ctx *RemoveaggrstmtContext) {}

// ExitRemoveaggrstmt is called when production removeaggrstmt is exited.
func (s *BaseRedshiftParserListener) ExitRemoveaggrstmt(ctx *RemoveaggrstmtContext) {}

// EnterRemoveoperstmt is called when production removeoperstmt is entered.
func (s *BaseRedshiftParserListener) EnterRemoveoperstmt(ctx *RemoveoperstmtContext) {}

// ExitRemoveoperstmt is called when production removeoperstmt is exited.
func (s *BaseRedshiftParserListener) ExitRemoveoperstmt(ctx *RemoveoperstmtContext) {}

// EnterOper_argtypes is called when production oper_argtypes is entered.
func (s *BaseRedshiftParserListener) EnterOper_argtypes(ctx *Oper_argtypesContext) {}

// ExitOper_argtypes is called when production oper_argtypes is exited.
func (s *BaseRedshiftParserListener) ExitOper_argtypes(ctx *Oper_argtypesContext) {}

// EnterAny_operator is called when production any_operator is entered.
func (s *BaseRedshiftParserListener) EnterAny_operator(ctx *Any_operatorContext) {}

// ExitAny_operator is called when production any_operator is exited.
func (s *BaseRedshiftParserListener) ExitAny_operator(ctx *Any_operatorContext) {}

// EnterOperator_with_argtypes_list is called when production operator_with_argtypes_list is entered.
func (s *BaseRedshiftParserListener) EnterOperator_with_argtypes_list(ctx *Operator_with_argtypes_listContext) {
}

// ExitOperator_with_argtypes_list is called when production operator_with_argtypes_list is exited.
func (s *BaseRedshiftParserListener) ExitOperator_with_argtypes_list(ctx *Operator_with_argtypes_listContext) {
}

// EnterOperator_with_argtypes is called when production operator_with_argtypes is entered.
func (s *BaseRedshiftParserListener) EnterOperator_with_argtypes(ctx *Operator_with_argtypesContext) {
}

// ExitOperator_with_argtypes is called when production operator_with_argtypes is exited.
func (s *BaseRedshiftParserListener) ExitOperator_with_argtypes(ctx *Operator_with_argtypesContext) {}

// EnterDostmt is called when production dostmt is entered.
func (s *BaseRedshiftParserListener) EnterDostmt(ctx *DostmtContext) {}

// ExitDostmt is called when production dostmt is exited.
func (s *BaseRedshiftParserListener) ExitDostmt(ctx *DostmtContext) {}

// EnterDostmt_opt_list is called when production dostmt_opt_list is entered.
func (s *BaseRedshiftParserListener) EnterDostmt_opt_list(ctx *Dostmt_opt_listContext) {}

// ExitDostmt_opt_list is called when production dostmt_opt_list is exited.
func (s *BaseRedshiftParserListener) ExitDostmt_opt_list(ctx *Dostmt_opt_listContext) {}

// EnterDostmt_opt_item is called when production dostmt_opt_item is entered.
func (s *BaseRedshiftParserListener) EnterDostmt_opt_item(ctx *Dostmt_opt_itemContext) {}

// ExitDostmt_opt_item is called when production dostmt_opt_item is exited.
func (s *BaseRedshiftParserListener) ExitDostmt_opt_item(ctx *Dostmt_opt_itemContext) {}

// EnterCreatecaststmt is called when production createcaststmt is entered.
func (s *BaseRedshiftParserListener) EnterCreatecaststmt(ctx *CreatecaststmtContext) {}

// ExitCreatecaststmt is called when production createcaststmt is exited.
func (s *BaseRedshiftParserListener) ExitCreatecaststmt(ctx *CreatecaststmtContext) {}

// EnterCast_context is called when production cast_context is entered.
func (s *BaseRedshiftParserListener) EnterCast_context(ctx *Cast_contextContext) {}

// ExitCast_context is called when production cast_context is exited.
func (s *BaseRedshiftParserListener) ExitCast_context(ctx *Cast_contextContext) {}

// EnterDropcaststmt is called when production dropcaststmt is entered.
func (s *BaseRedshiftParserListener) EnterDropcaststmt(ctx *DropcaststmtContext) {}

// ExitDropcaststmt is called when production dropcaststmt is exited.
func (s *BaseRedshiftParserListener) ExitDropcaststmt(ctx *DropcaststmtContext) {}

// EnterOpt_if_exists is called when production opt_if_exists is entered.
func (s *BaseRedshiftParserListener) EnterOpt_if_exists(ctx *Opt_if_existsContext) {}

// ExitOpt_if_exists is called when production opt_if_exists is exited.
func (s *BaseRedshiftParserListener) ExitOpt_if_exists(ctx *Opt_if_existsContext) {}

// EnterCreatetransformstmt is called when production createtransformstmt is entered.
func (s *BaseRedshiftParserListener) EnterCreatetransformstmt(ctx *CreatetransformstmtContext) {}

// ExitCreatetransformstmt is called when production createtransformstmt is exited.
func (s *BaseRedshiftParserListener) ExitCreatetransformstmt(ctx *CreatetransformstmtContext) {}

// EnterTransform_element_list is called when production transform_element_list is entered.
func (s *BaseRedshiftParserListener) EnterTransform_element_list(ctx *Transform_element_listContext) {
}

// ExitTransform_element_list is called when production transform_element_list is exited.
func (s *BaseRedshiftParserListener) ExitTransform_element_list(ctx *Transform_element_listContext) {}

// EnterDroptransformstmt is called when production droptransformstmt is entered.
func (s *BaseRedshiftParserListener) EnterDroptransformstmt(ctx *DroptransformstmtContext) {}

// ExitDroptransformstmt is called when production droptransformstmt is exited.
func (s *BaseRedshiftParserListener) ExitDroptransformstmt(ctx *DroptransformstmtContext) {}

// EnterReindexstmt is called when production reindexstmt is entered.
func (s *BaseRedshiftParserListener) EnterReindexstmt(ctx *ReindexstmtContext) {}

// ExitReindexstmt is called when production reindexstmt is exited.
func (s *BaseRedshiftParserListener) ExitReindexstmt(ctx *ReindexstmtContext) {}

// EnterReindex_target_type is called when production reindex_target_type is entered.
func (s *BaseRedshiftParserListener) EnterReindex_target_type(ctx *Reindex_target_typeContext) {}

// ExitReindex_target_type is called when production reindex_target_type is exited.
func (s *BaseRedshiftParserListener) ExitReindex_target_type(ctx *Reindex_target_typeContext) {}

// EnterReindex_target_multitable is called when production reindex_target_multitable is entered.
func (s *BaseRedshiftParserListener) EnterReindex_target_multitable(ctx *Reindex_target_multitableContext) {
}

// ExitReindex_target_multitable is called when production reindex_target_multitable is exited.
func (s *BaseRedshiftParserListener) ExitReindex_target_multitable(ctx *Reindex_target_multitableContext) {
}

// EnterReindex_option_list is called when production reindex_option_list is entered.
func (s *BaseRedshiftParserListener) EnterReindex_option_list(ctx *Reindex_option_listContext) {}

// ExitReindex_option_list is called when production reindex_option_list is exited.
func (s *BaseRedshiftParserListener) ExitReindex_option_list(ctx *Reindex_option_listContext) {}

// EnterReindex_option_elem is called when production reindex_option_elem is entered.
func (s *BaseRedshiftParserListener) EnterReindex_option_elem(ctx *Reindex_option_elemContext) {}

// ExitReindex_option_elem is called when production reindex_option_elem is exited.
func (s *BaseRedshiftParserListener) ExitReindex_option_elem(ctx *Reindex_option_elemContext) {}

// EnterAltertblspcstmt is called when production altertblspcstmt is entered.
func (s *BaseRedshiftParserListener) EnterAltertblspcstmt(ctx *AltertblspcstmtContext) {}

// ExitAltertblspcstmt is called when production altertblspcstmt is exited.
func (s *BaseRedshiftParserListener) ExitAltertblspcstmt(ctx *AltertblspcstmtContext) {}

// EnterRenamestmt is called when production renamestmt is entered.
func (s *BaseRedshiftParserListener) EnterRenamestmt(ctx *RenamestmtContext) {}

// ExitRenamestmt is called when production renamestmt is exited.
func (s *BaseRedshiftParserListener) ExitRenamestmt(ctx *RenamestmtContext) {}

// EnterOpt_column is called when production opt_column is entered.
func (s *BaseRedshiftParserListener) EnterOpt_column(ctx *Opt_columnContext) {}

// ExitOpt_column is called when production opt_column is exited.
func (s *BaseRedshiftParserListener) ExitOpt_column(ctx *Opt_columnContext) {}

// EnterOpt_set_data is called when production opt_set_data is entered.
func (s *BaseRedshiftParserListener) EnterOpt_set_data(ctx *Opt_set_dataContext) {}

// ExitOpt_set_data is called when production opt_set_data is exited.
func (s *BaseRedshiftParserListener) ExitOpt_set_data(ctx *Opt_set_dataContext) {}

// EnterAlterobjectdependsstmt is called when production alterobjectdependsstmt is entered.
func (s *BaseRedshiftParserListener) EnterAlterobjectdependsstmt(ctx *AlterobjectdependsstmtContext) {
}

// ExitAlterobjectdependsstmt is called when production alterobjectdependsstmt is exited.
func (s *BaseRedshiftParserListener) ExitAlterobjectdependsstmt(ctx *AlterobjectdependsstmtContext) {}

// EnterOpt_no is called when production opt_no is entered.
func (s *BaseRedshiftParserListener) EnterOpt_no(ctx *Opt_noContext) {}

// ExitOpt_no is called when production opt_no is exited.
func (s *BaseRedshiftParserListener) ExitOpt_no(ctx *Opt_noContext) {}

// EnterAlterobjectschemastmt is called when production alterobjectschemastmt is entered.
func (s *BaseRedshiftParserListener) EnterAlterobjectschemastmt(ctx *AlterobjectschemastmtContext) {}

// ExitAlterobjectschemastmt is called when production alterobjectschemastmt is exited.
func (s *BaseRedshiftParserListener) ExitAlterobjectschemastmt(ctx *AlterobjectschemastmtContext) {}

// EnterAlteroperatorstmt is called when production alteroperatorstmt is entered.
func (s *BaseRedshiftParserListener) EnterAlteroperatorstmt(ctx *AlteroperatorstmtContext) {}

// ExitAlteroperatorstmt is called when production alteroperatorstmt is exited.
func (s *BaseRedshiftParserListener) ExitAlteroperatorstmt(ctx *AlteroperatorstmtContext) {}

// EnterOperator_def_list is called when production operator_def_list is entered.
func (s *BaseRedshiftParserListener) EnterOperator_def_list(ctx *Operator_def_listContext) {}

// ExitOperator_def_list is called when production operator_def_list is exited.
func (s *BaseRedshiftParserListener) ExitOperator_def_list(ctx *Operator_def_listContext) {}

// EnterOperator_def_elem is called when production operator_def_elem is entered.
func (s *BaseRedshiftParserListener) EnterOperator_def_elem(ctx *Operator_def_elemContext) {}

// ExitOperator_def_elem is called when production operator_def_elem is exited.
func (s *BaseRedshiftParserListener) ExitOperator_def_elem(ctx *Operator_def_elemContext) {}

// EnterOperator_def_arg is called when production operator_def_arg is entered.
func (s *BaseRedshiftParserListener) EnterOperator_def_arg(ctx *Operator_def_argContext) {}

// ExitOperator_def_arg is called when production operator_def_arg is exited.
func (s *BaseRedshiftParserListener) ExitOperator_def_arg(ctx *Operator_def_argContext) {}

// EnterAltertypestmt is called when production altertypestmt is entered.
func (s *BaseRedshiftParserListener) EnterAltertypestmt(ctx *AltertypestmtContext) {}

// ExitAltertypestmt is called when production altertypestmt is exited.
func (s *BaseRedshiftParserListener) ExitAltertypestmt(ctx *AltertypestmtContext) {}

// EnterAlterownerstmt is called when production alterownerstmt is entered.
func (s *BaseRedshiftParserListener) EnterAlterownerstmt(ctx *AlterownerstmtContext) {}

// ExitAlterownerstmt is called when production alterownerstmt is exited.
func (s *BaseRedshiftParserListener) ExitAlterownerstmt(ctx *AlterownerstmtContext) {}

// EnterCreatepublicationstmt is called when production createpublicationstmt is entered.
func (s *BaseRedshiftParserListener) EnterCreatepublicationstmt(ctx *CreatepublicationstmtContext) {}

// ExitCreatepublicationstmt is called when production createpublicationstmt is exited.
func (s *BaseRedshiftParserListener) ExitCreatepublicationstmt(ctx *CreatepublicationstmtContext) {}

// EnterPub_obj_list is called when production pub_obj_list is entered.
func (s *BaseRedshiftParserListener) EnterPub_obj_list(ctx *Pub_obj_listContext) {}

// ExitPub_obj_list is called when production pub_obj_list is exited.
func (s *BaseRedshiftParserListener) ExitPub_obj_list(ctx *Pub_obj_listContext) {}

// EnterPublication_obj_spec is called when production publication_obj_spec is entered.
func (s *BaseRedshiftParserListener) EnterPublication_obj_spec(ctx *Publication_obj_specContext) {}

// ExitPublication_obj_spec is called when production publication_obj_spec is exited.
func (s *BaseRedshiftParserListener) ExitPublication_obj_spec(ctx *Publication_obj_specContext) {}

// EnterOpt_where_clause is called when production opt_where_clause is entered.
func (s *BaseRedshiftParserListener) EnterOpt_where_clause(ctx *Opt_where_clauseContext) {}

// ExitOpt_where_clause is called when production opt_where_clause is exited.
func (s *BaseRedshiftParserListener) ExitOpt_where_clause(ctx *Opt_where_clauseContext) {}

// EnterAlterpublicationstmt is called when production alterpublicationstmt is entered.
func (s *BaseRedshiftParserListener) EnterAlterpublicationstmt(ctx *AlterpublicationstmtContext) {}

// ExitAlterpublicationstmt is called when production alterpublicationstmt is exited.
func (s *BaseRedshiftParserListener) ExitAlterpublicationstmt(ctx *AlterpublicationstmtContext) {}

// EnterCreatesubscriptionstmt is called when production createsubscriptionstmt is entered.
func (s *BaseRedshiftParserListener) EnterCreatesubscriptionstmt(ctx *CreatesubscriptionstmtContext) {
}

// ExitCreatesubscriptionstmt is called when production createsubscriptionstmt is exited.
func (s *BaseRedshiftParserListener) ExitCreatesubscriptionstmt(ctx *CreatesubscriptionstmtContext) {}

// EnterPublication_name_list is called when production publication_name_list is entered.
func (s *BaseRedshiftParserListener) EnterPublication_name_list(ctx *Publication_name_listContext) {}

// ExitPublication_name_list is called when production publication_name_list is exited.
func (s *BaseRedshiftParserListener) ExitPublication_name_list(ctx *Publication_name_listContext) {}

// EnterPublication_name_item is called when production publication_name_item is entered.
func (s *BaseRedshiftParserListener) EnterPublication_name_item(ctx *Publication_name_itemContext) {}

// ExitPublication_name_item is called when production publication_name_item is exited.
func (s *BaseRedshiftParserListener) ExitPublication_name_item(ctx *Publication_name_itemContext) {}

// EnterAltersubscriptionstmt is called when production altersubscriptionstmt is entered.
func (s *BaseRedshiftParserListener) EnterAltersubscriptionstmt(ctx *AltersubscriptionstmtContext) {}

// ExitAltersubscriptionstmt is called when production altersubscriptionstmt is exited.
func (s *BaseRedshiftParserListener) ExitAltersubscriptionstmt(ctx *AltersubscriptionstmtContext) {}

// EnterDropsubscriptionstmt is called when production dropsubscriptionstmt is entered.
func (s *BaseRedshiftParserListener) EnterDropsubscriptionstmt(ctx *DropsubscriptionstmtContext) {}

// ExitDropsubscriptionstmt is called when production dropsubscriptionstmt is exited.
func (s *BaseRedshiftParserListener) ExitDropsubscriptionstmt(ctx *DropsubscriptionstmtContext) {}

// EnterRulestmt is called when production rulestmt is entered.
func (s *BaseRedshiftParserListener) EnterRulestmt(ctx *RulestmtContext) {}

// ExitRulestmt is called when production rulestmt is exited.
func (s *BaseRedshiftParserListener) ExitRulestmt(ctx *RulestmtContext) {}

// EnterRuleactionlist is called when production ruleactionlist is entered.
func (s *BaseRedshiftParserListener) EnterRuleactionlist(ctx *RuleactionlistContext) {}

// ExitRuleactionlist is called when production ruleactionlist is exited.
func (s *BaseRedshiftParserListener) ExitRuleactionlist(ctx *RuleactionlistContext) {}

// EnterRuleactionmulti is called when production ruleactionmulti is entered.
func (s *BaseRedshiftParserListener) EnterRuleactionmulti(ctx *RuleactionmultiContext) {}

// ExitRuleactionmulti is called when production ruleactionmulti is exited.
func (s *BaseRedshiftParserListener) ExitRuleactionmulti(ctx *RuleactionmultiContext) {}

// EnterRuleactionstmt is called when production ruleactionstmt is entered.
func (s *BaseRedshiftParserListener) EnterRuleactionstmt(ctx *RuleactionstmtContext) {}

// ExitRuleactionstmt is called when production ruleactionstmt is exited.
func (s *BaseRedshiftParserListener) ExitRuleactionstmt(ctx *RuleactionstmtContext) {}

// EnterRuleactionstmtOrEmpty is called when production ruleactionstmtOrEmpty is entered.
func (s *BaseRedshiftParserListener) EnterRuleactionstmtOrEmpty(ctx *RuleactionstmtOrEmptyContext) {}

// ExitRuleactionstmtOrEmpty is called when production ruleactionstmtOrEmpty is exited.
func (s *BaseRedshiftParserListener) ExitRuleactionstmtOrEmpty(ctx *RuleactionstmtOrEmptyContext) {}

// EnterEvent is called when production event is entered.
func (s *BaseRedshiftParserListener) EnterEvent(ctx *EventContext) {}

// ExitEvent is called when production event is exited.
func (s *BaseRedshiftParserListener) ExitEvent(ctx *EventContext) {}

// EnterOpt_instead is called when production opt_instead is entered.
func (s *BaseRedshiftParserListener) EnterOpt_instead(ctx *Opt_insteadContext) {}

// ExitOpt_instead is called when production opt_instead is exited.
func (s *BaseRedshiftParserListener) ExitOpt_instead(ctx *Opt_insteadContext) {}

// EnterNotifystmt is called when production notifystmt is entered.
func (s *BaseRedshiftParserListener) EnterNotifystmt(ctx *NotifystmtContext) {}

// ExitNotifystmt is called when production notifystmt is exited.
func (s *BaseRedshiftParserListener) ExitNotifystmt(ctx *NotifystmtContext) {}

// EnterNotify_payload is called when production notify_payload is entered.
func (s *BaseRedshiftParserListener) EnterNotify_payload(ctx *Notify_payloadContext) {}

// ExitNotify_payload is called when production notify_payload is exited.
func (s *BaseRedshiftParserListener) ExitNotify_payload(ctx *Notify_payloadContext) {}

// EnterListenstmt is called when production listenstmt is entered.
func (s *BaseRedshiftParserListener) EnterListenstmt(ctx *ListenstmtContext) {}

// ExitListenstmt is called when production listenstmt is exited.
func (s *BaseRedshiftParserListener) ExitListenstmt(ctx *ListenstmtContext) {}

// EnterUnlistenstmt is called when production unlistenstmt is entered.
func (s *BaseRedshiftParserListener) EnterUnlistenstmt(ctx *UnlistenstmtContext) {}

// ExitUnlistenstmt is called when production unlistenstmt is exited.
func (s *BaseRedshiftParserListener) ExitUnlistenstmt(ctx *UnlistenstmtContext) {}

// EnterTransactionstmt is called when production transactionstmt is entered.
func (s *BaseRedshiftParserListener) EnterTransactionstmt(ctx *TransactionstmtContext) {}

// ExitTransactionstmt is called when production transactionstmt is exited.
func (s *BaseRedshiftParserListener) ExitTransactionstmt(ctx *TransactionstmtContext) {}

// EnterOpt_transaction is called when production opt_transaction is entered.
func (s *BaseRedshiftParserListener) EnterOpt_transaction(ctx *Opt_transactionContext) {}

// ExitOpt_transaction is called when production opt_transaction is exited.
func (s *BaseRedshiftParserListener) ExitOpt_transaction(ctx *Opt_transactionContext) {}

// EnterTransaction_mode_item is called when production transaction_mode_item is entered.
func (s *BaseRedshiftParserListener) EnterTransaction_mode_item(ctx *Transaction_mode_itemContext) {}

// ExitTransaction_mode_item is called when production transaction_mode_item is exited.
func (s *BaseRedshiftParserListener) ExitTransaction_mode_item(ctx *Transaction_mode_itemContext) {}

// EnterTransaction_mode_list is called when production transaction_mode_list is entered.
func (s *BaseRedshiftParserListener) EnterTransaction_mode_list(ctx *Transaction_mode_listContext) {}

// ExitTransaction_mode_list is called when production transaction_mode_list is exited.
func (s *BaseRedshiftParserListener) ExitTransaction_mode_list(ctx *Transaction_mode_listContext) {}

// EnterTransaction_mode_list_or_empty is called when production transaction_mode_list_or_empty is entered.
func (s *BaseRedshiftParserListener) EnterTransaction_mode_list_or_empty(ctx *Transaction_mode_list_or_emptyContext) {
}

// ExitTransaction_mode_list_or_empty is called when production transaction_mode_list_or_empty is exited.
func (s *BaseRedshiftParserListener) ExitTransaction_mode_list_or_empty(ctx *Transaction_mode_list_or_emptyContext) {
}

// EnterOpt_transaction_chain is called when production opt_transaction_chain is entered.
func (s *BaseRedshiftParserListener) EnterOpt_transaction_chain(ctx *Opt_transaction_chainContext) {}

// ExitOpt_transaction_chain is called when production opt_transaction_chain is exited.
func (s *BaseRedshiftParserListener) ExitOpt_transaction_chain(ctx *Opt_transaction_chainContext) {}

// EnterViewstmt is called when production viewstmt is entered.
func (s *BaseRedshiftParserListener) EnterViewstmt(ctx *ViewstmtContext) {}

// ExitViewstmt is called when production viewstmt is exited.
func (s *BaseRedshiftParserListener) ExitViewstmt(ctx *ViewstmtContext) {}

// EnterWith_no_schema_binding is called when production with_no_schema_binding is entered.
func (s *BaseRedshiftParserListener) EnterWith_no_schema_binding(ctx *With_no_schema_bindingContext) {
}

// ExitWith_no_schema_binding is called when production with_no_schema_binding is exited.
func (s *BaseRedshiftParserListener) ExitWith_no_schema_binding(ctx *With_no_schema_bindingContext) {}

// EnterOpt_check_option is called when production opt_check_option is entered.
func (s *BaseRedshiftParserListener) EnterOpt_check_option(ctx *Opt_check_optionContext) {}

// ExitOpt_check_option is called when production opt_check_option is exited.
func (s *BaseRedshiftParserListener) ExitOpt_check_option(ctx *Opt_check_optionContext) {}

// EnterLoadstmt is called when production loadstmt is entered.
func (s *BaseRedshiftParserListener) EnterLoadstmt(ctx *LoadstmtContext) {}

// ExitLoadstmt is called when production loadstmt is exited.
func (s *BaseRedshiftParserListener) ExitLoadstmt(ctx *LoadstmtContext) {}

// EnterAlterdatasharestmt is called when production alterdatasharestmt is entered.
func (s *BaseRedshiftParserListener) EnterAlterdatasharestmt(ctx *AlterdatasharestmtContext) {}

// ExitAlterdatasharestmt is called when production alterdatasharestmt is exited.
func (s *BaseRedshiftParserListener) ExitAlterdatasharestmt(ctx *AlterdatasharestmtContext) {}

// EnterAlterdatashare_action is called when production alterdatashare_action is entered.
func (s *BaseRedshiftParserListener) EnterAlterdatashare_action(ctx *Alterdatashare_actionContext) {}

// ExitAlterdatashare_action is called when production alterdatashare_action is exited.
func (s *BaseRedshiftParserListener) ExitAlterdatashare_action(ctx *Alterdatashare_actionContext) {}

// EnterAlterdatashare_add_drop is called when production alterdatashare_add_drop is entered.
func (s *BaseRedshiftParserListener) EnterAlterdatashare_add_drop(ctx *Alterdatashare_add_dropContext) {
}

// ExitAlterdatashare_add_drop is called when production alterdatashare_add_drop is exited.
func (s *BaseRedshiftParserListener) ExitAlterdatashare_add_drop(ctx *Alterdatashare_add_dropContext) {
}

// EnterAlterdatashare_objects is called when production alterdatashare_objects is entered.
func (s *BaseRedshiftParserListener) EnterAlterdatashare_objects(ctx *Alterdatashare_objectsContext) {
}

// ExitAlterdatashare_objects is called when production alterdatashare_objects is exited.
func (s *BaseRedshiftParserListener) ExitAlterdatashare_objects(ctx *Alterdatashare_objectsContext) {}

// EnterDatashare_table_list is called when production datashare_table_list is entered.
func (s *BaseRedshiftParserListener) EnterDatashare_table_list(ctx *Datashare_table_listContext) {}

// ExitDatashare_table_list is called when production datashare_table_list is exited.
func (s *BaseRedshiftParserListener) ExitDatashare_table_list(ctx *Datashare_table_listContext) {}

// EnterDatashare_table_name is called when production datashare_table_name is entered.
func (s *BaseRedshiftParserListener) EnterDatashare_table_name(ctx *Datashare_table_nameContext) {}

// ExitDatashare_table_name is called when production datashare_table_name is exited.
func (s *BaseRedshiftParserListener) ExitDatashare_table_name(ctx *Datashare_table_nameContext) {}

// EnterTable_name is called when production table_name is entered.
func (s *BaseRedshiftParserListener) EnterTable_name(ctx *Table_nameContext) {}

// ExitTable_name is called when production table_name is exited.
func (s *BaseRedshiftParserListener) ExitTable_name(ctx *Table_nameContext) {}

// EnterTemporary_table_name is called when production temporary_table_name is entered.
func (s *BaseRedshiftParserListener) EnterTemporary_table_name(ctx *Temporary_table_nameContext) {}

// ExitTemporary_table_name is called when production temporary_table_name is exited.
func (s *BaseRedshiftParserListener) ExitTemporary_table_name(ctx *Temporary_table_nameContext) {}

// EnterDatashare_function_list is called when production datashare_function_list is entered.
func (s *BaseRedshiftParserListener) EnterDatashare_function_list(ctx *Datashare_function_listContext) {
}

// ExitDatashare_function_list is called when production datashare_function_list is exited.
func (s *BaseRedshiftParserListener) ExitDatashare_function_list(ctx *Datashare_function_listContext) {
}

// EnterDatashare_function is called when production datashare_function is entered.
func (s *BaseRedshiftParserListener) EnterDatashare_function(ctx *Datashare_functionContext) {}

// ExitDatashare_function is called when production datashare_function is exited.
func (s *BaseRedshiftParserListener) ExitDatashare_function(ctx *Datashare_functionContext) {}

// EnterDatashare_function_name is called when production datashare_function_name is entered.
func (s *BaseRedshiftParserListener) EnterDatashare_function_name(ctx *Datashare_function_nameContext) {
}

// ExitDatashare_function_name is called when production datashare_function_name is exited.
func (s *BaseRedshiftParserListener) ExitDatashare_function_name(ctx *Datashare_function_nameContext) {
}

// EnterCreatedatasharestmt is called when production createdatasharestmt is entered.
func (s *BaseRedshiftParserListener) EnterCreatedatasharestmt(ctx *CreatedatasharestmtContext) {}

// ExitCreatedatasharestmt is called when production createdatasharestmt is exited.
func (s *BaseRedshiftParserListener) ExitCreatedatasharestmt(ctx *CreatedatasharestmtContext) {}

// EnterCreatedatashareoptions is called when production createdatashareoptions is entered.
func (s *BaseRedshiftParserListener) EnterCreatedatashareoptions(ctx *CreatedatashareoptionsContext) {
}

// ExitCreatedatashareoptions is called when production createdatashareoptions is exited.
func (s *BaseRedshiftParserListener) ExitCreatedatashareoptions(ctx *CreatedatashareoptionsContext) {}

// EnterCreatedatashareoption is called when production createdatashareoption is entered.
func (s *BaseRedshiftParserListener) EnterCreatedatashareoption(ctx *CreatedatashareoptionContext) {}

// ExitCreatedatashareoption is called when production createdatashareoption is exited.
func (s *BaseRedshiftParserListener) ExitCreatedatashareoption(ctx *CreatedatashareoptionContext) {}

// EnterSetpublicaccessibleoption is called when production setpublicaccessibleoption is entered.
func (s *BaseRedshiftParserListener) EnterSetpublicaccessibleoption(ctx *SetpublicaccessibleoptionContext) {
}

// ExitSetpublicaccessibleoption is called when production setpublicaccessibleoption is exited.
func (s *BaseRedshiftParserListener) ExitSetpublicaccessibleoption(ctx *SetpublicaccessibleoptionContext) {
}

// EnterManagedbyoption is called when production managedbyoption is entered.
func (s *BaseRedshiftParserListener) EnterManagedbyoption(ctx *ManagedbyoptionContext) {}

// ExitManagedbyoption is called when production managedbyoption is exited.
func (s *BaseRedshiftParserListener) ExitManagedbyoption(ctx *ManagedbyoptionContext) {}

// EnterDescdatasharestmt is called when production descdatasharestmt is entered.
func (s *BaseRedshiftParserListener) EnterDescdatasharestmt(ctx *DescdatasharestmtContext) {}

// ExitDescdatasharestmt is called when production descdatasharestmt is exited.
func (s *BaseRedshiftParserListener) ExitDescdatasharestmt(ctx *DescdatasharestmtContext) {}

// EnterDropdatasharestmt is called when production dropdatasharestmt is entered.
func (s *BaseRedshiftParserListener) EnterDropdatasharestmt(ctx *DropdatasharestmtContext) {}

// ExitDropdatasharestmt is called when production dropdatasharestmt is exited.
func (s *BaseRedshiftParserListener) ExitDropdatasharestmt(ctx *DropdatasharestmtContext) {}

// EnterAlterexternalschemastmt is called when production alterexternalschemastmt is entered.
func (s *BaseRedshiftParserListener) EnterAlterexternalschemastmt(ctx *AlterexternalschemastmtContext) {
}

// ExitAlterexternalschemastmt is called when production alterexternalschemastmt is exited.
func (s *BaseRedshiftParserListener) ExitAlterexternalschemastmt(ctx *AlterexternalschemastmtContext) {
}

// EnterAltexternalschemaopts is called when production altexternalschemaopts is entered.
func (s *BaseRedshiftParserListener) EnterAltexternalschemaopts(ctx *AltexternalschemaoptsContext) {}

// ExitAltexternalschemaopts is called when production altexternalschemaopts is exited.
func (s *BaseRedshiftParserListener) ExitAltexternalschemaopts(ctx *AltexternalschemaoptsContext) {}

// EnterAlterexternalviewstmt is called when production alterexternalviewstmt is entered.
func (s *BaseRedshiftParserListener) EnterAlterexternalviewstmt(ctx *AlterexternalviewstmtContext) {}

// ExitAlterexternalviewstmt is called when production alterexternalviewstmt is exited.
func (s *BaseRedshiftParserListener) ExitAlterexternalviewstmt(ctx *AlterexternalviewstmtContext) {}

// EnterCreateexternalschemastmt is called when production createexternalschemastmt is entered.
func (s *BaseRedshiftParserListener) EnterCreateexternalschemastmt(ctx *CreateexternalschemastmtContext) {
}

// ExitCreateexternalschemastmt is called when production createexternalschemastmt is exited.
func (s *BaseRedshiftParserListener) ExitCreateexternalschemastmt(ctx *CreateexternalschemastmtContext) {
}

// EnterFromdatacatalogclause is called when production fromdatacatalogclause is entered.
func (s *BaseRedshiftParserListener) EnterFromdatacatalogclause(ctx *FromdatacatalogclauseContext) {}

// ExitFromdatacatalogclause is called when production fromdatacatalogclause is exited.
func (s *BaseRedshiftParserListener) ExitFromdatacatalogclause(ctx *FromdatacatalogclauseContext) {}

// EnterDropschemastmt is called when production dropschemastmt is entered.
func (s *BaseRedshiftParserListener) EnterDropschemastmt(ctx *DropschemastmtContext) {}

// ExitDropschemastmt is called when production dropschemastmt is exited.
func (s *BaseRedshiftParserListener) ExitDropschemastmt(ctx *DropschemastmtContext) {}

// EnterImplicitdatacatalogclause is called when production implicitdatacatalogclause is entered.
func (s *BaseRedshiftParserListener) EnterImplicitdatacatalogclause(ctx *ImplicitdatacatalogclauseContext) {
}

// ExitImplicitdatacatalogclause is called when production implicitdatacatalogclause is exited.
func (s *BaseRedshiftParserListener) ExitImplicitdatacatalogclause(ctx *ImplicitdatacatalogclauseContext) {
}

// EnterFromhivemetastoreclause is called when production fromhivemetastoreclause is entered.
func (s *BaseRedshiftParserListener) EnterFromhivemetastoreclause(ctx *FromhivemetastoreclauseContext) {
}

// ExitFromhivemetastoreclause is called when production fromhivemetastoreclause is exited.
func (s *BaseRedshiftParserListener) ExitFromhivemetastoreclause(ctx *FromhivemetastoreclauseContext) {
}

// EnterFrompostgresclause is called when production frompostgresclause is entered.
func (s *BaseRedshiftParserListener) EnterFrompostgresclause(ctx *FrompostgresclauseContext) {}

// ExitFrompostgresclause is called when production frompostgresclause is exited.
func (s *BaseRedshiftParserListener) ExitFrompostgresclause(ctx *FrompostgresclauseContext) {}

// EnterFrommysqlclause is called when production frommysqlclause is entered.
func (s *BaseRedshiftParserListener) EnterFrommysqlclause(ctx *FrommysqlclauseContext) {}

// ExitFrommysqlclause is called when production frommysqlclause is exited.
func (s *BaseRedshiftParserListener) ExitFrommysqlclause(ctx *FrommysqlclauseContext) {}

// EnterFromkinesisclause is called when production fromkinesisclause is entered.
func (s *BaseRedshiftParserListener) EnterFromkinesisclause(ctx *FromkinesisclauseContext) {}

// ExitFromkinesisclause is called when production fromkinesisclause is exited.
func (s *BaseRedshiftParserListener) ExitFromkinesisclause(ctx *FromkinesisclauseContext) {}

// EnterFromkafkaclause is called when production fromkafkaclause is entered.
func (s *BaseRedshiftParserListener) EnterFromkafkaclause(ctx *FromkafkaclauseContext) {}

// ExitFromkafkaclause is called when production fromkafkaclause is exited.
func (s *BaseRedshiftParserListener) ExitFromkafkaclause(ctx *FromkafkaclauseContext) {}

// EnterFrommskclause is called when production frommskclause is entered.
func (s *BaseRedshiftParserListener) EnterFrommskclause(ctx *FrommskclauseContext) {}

// ExitFrommskclause is called when production frommskclause is exited.
func (s *BaseRedshiftParserListener) ExitFrommskclause(ctx *FrommskclauseContext) {}

// EnterFromredshiftclause is called when production fromredshiftclause is entered.
func (s *BaseRedshiftParserListener) EnterFromredshiftclause(ctx *FromredshiftclauseContext) {}

// ExitFromredshiftclause is called when production fromredshiftclause is exited.
func (s *BaseRedshiftParserListener) ExitFromredshiftclause(ctx *FromredshiftclauseContext) {}

// EnterIamrolevalue is called when production iamrolevalue is entered.
func (s *BaseRedshiftParserListener) EnterIamrolevalue(ctx *IamrolevalueContext) {}

// ExitIamrolevalue is called when production iamrolevalue is exited.
func (s *BaseRedshiftParserListener) ExitIamrolevalue(ctx *IamrolevalueContext) {}

// EnterCatalogrolevalue is called when production catalogrolevalue is entered.
func (s *BaseRedshiftParserListener) EnterCatalogrolevalue(ctx *CatalogrolevalueContext) {}

// ExitCatalogrolevalue is called when production catalogrolevalue is exited.
func (s *BaseRedshiftParserListener) ExitCatalogrolevalue(ctx *CatalogrolevalueContext) {}

// EnterAuthenticationvalue is called when production authenticationvalue is entered.
func (s *BaseRedshiftParserListener) EnterAuthenticationvalue(ctx *AuthenticationvalueContext) {}

// ExitAuthenticationvalue is called when production authenticationvalue is exited.
func (s *BaseRedshiftParserListener) ExitAuthenticationvalue(ctx *AuthenticationvalueContext) {}

// EnterCreateexternalfunctionstmt is called when production createexternalfunctionstmt is entered.
func (s *BaseRedshiftParserListener) EnterCreateexternalfunctionstmt(ctx *CreateexternalfunctionstmtContext) {
}

// ExitCreateexternalfunctionstmt is called when production createexternalfunctionstmt is exited.
func (s *BaseRedshiftParserListener) ExitCreateexternalfunctionstmt(ctx *CreateexternalfunctionstmtContext) {
}

// EnterExternal_func_params is called when production external_func_params is entered.
func (s *BaseRedshiftParserListener) EnterExternal_func_params(ctx *External_func_paramsContext) {}

// ExitExternal_func_params is called when production external_func_params is exited.
func (s *BaseRedshiftParserListener) ExitExternal_func_params(ctx *External_func_paramsContext) {}

// EnterParamlist is called when production paramlist is entered.
func (s *BaseRedshiftParserListener) EnterParamlist(ctx *ParamlistContext) {}

// ExitParamlist is called when production paramlist is exited.
func (s *BaseRedshiftParserListener) ExitParamlist(ctx *ParamlistContext) {}

// EnterParam_spec is called when production param_spec is entered.
func (s *BaseRedshiftParserListener) EnterParam_spec(ctx *Param_specContext) {}

// ExitParam_spec is called when production param_spec is exited.
func (s *BaseRedshiftParserListener) ExitParam_spec(ctx *Param_specContext) {}

// EnterCreateexternalmodelstmt is called when production createexternalmodelstmt is entered.
func (s *BaseRedshiftParserListener) EnterCreateexternalmodelstmt(ctx *CreateexternalmodelstmtContext) {
}

// ExitCreateexternalmodelstmt is called when production createexternalmodelstmt is exited.
func (s *BaseRedshiftParserListener) ExitCreateexternalmodelstmt(ctx *CreateexternalmodelstmtContext) {
}

// EnterCreateexternaltablestmt is called when production createexternaltablestmt is entered.
func (s *BaseRedshiftParserListener) EnterCreateexternaltablestmt(ctx *CreateexternaltablestmtContext) {
}

// ExitCreateexternaltablestmt is called when production createexternaltablestmt is exited.
func (s *BaseRedshiftParserListener) ExitCreateexternaltablestmt(ctx *CreateexternaltablestmtContext) {
}

// EnterExtern_column_list is called when production extern_column_list is entered.
func (s *BaseRedshiftParserListener) EnterExtern_column_list(ctx *Extern_column_listContext) {}

// ExitExtern_column_list is called when production extern_column_list is exited.
func (s *BaseRedshiftParserListener) ExitExtern_column_list(ctx *Extern_column_listContext) {}

// EnterExtern_column_def is called when production extern_column_def is entered.
func (s *BaseRedshiftParserListener) EnterExtern_column_def(ctx *Extern_column_defContext) {}

// ExitExtern_column_def is called when production extern_column_def is exited.
func (s *BaseRedshiftParserListener) ExitExtern_column_def(ctx *Extern_column_defContext) {}

// EnterExtern_typename is called when production extern_typename is entered.
func (s *BaseRedshiftParserListener) EnterExtern_typename(ctx *Extern_typenameContext) {}

// ExitExtern_typename is called when production extern_typename is exited.
func (s *BaseRedshiftParserListener) ExitExtern_typename(ctx *Extern_typenameContext) {}

// EnterExtern_table_format is called when production extern_table_format is entered.
func (s *BaseRedshiftParserListener) EnterExtern_table_format(ctx *Extern_table_formatContext) {}

// ExitExtern_table_format is called when production extern_table_format is exited.
func (s *BaseRedshiftParserListener) ExitExtern_table_format(ctx *Extern_table_formatContext) {}

// EnterRow_format_spec is called when production row_format_spec is entered.
func (s *BaseRedshiftParserListener) EnterRow_format_spec(ctx *Row_format_specContext) {}

// ExitRow_format_spec is called when production row_format_spec is exited.
func (s *BaseRedshiftParserListener) ExitRow_format_spec(ctx *Row_format_specContext) {}

// EnterSerde_properties_list is called when production serde_properties_list is entered.
func (s *BaseRedshiftParserListener) EnterSerde_properties_list(ctx *Serde_properties_listContext) {}

// ExitSerde_properties_list is called when production serde_properties_list is exited.
func (s *BaseRedshiftParserListener) ExitSerde_properties_list(ctx *Serde_properties_listContext) {}

// EnterSerde_property is called when production serde_property is entered.
func (s *BaseRedshiftParserListener) EnterSerde_property(ctx *Serde_propertyContext) {}

// ExitSerde_property is called when production serde_property is exited.
func (s *BaseRedshiftParserListener) ExitSerde_property(ctx *Serde_propertyContext) {}

// EnterExternal_format_spec is called when production external_format_spec is entered.
func (s *BaseRedshiftParserListener) EnterExternal_format_spec(ctx *External_format_specContext) {}

// ExitExternal_format_spec is called when production external_format_spec is exited.
func (s *BaseRedshiftParserListener) ExitExternal_format_spec(ctx *External_format_specContext) {}

// EnterTable_properties_list is called when production table_properties_list is entered.
func (s *BaseRedshiftParserListener) EnterTable_properties_list(ctx *Table_properties_listContext) {}

// ExitTable_properties_list is called when production table_properties_list is exited.
func (s *BaseRedshiftParserListener) ExitTable_properties_list(ctx *Table_properties_listContext) {}

// EnterTable_property is called when production table_property is entered.
func (s *BaseRedshiftParserListener) EnterTable_property(ctx *Table_propertyContext) {}

// ExitTable_property is called when production table_property is exited.
func (s *BaseRedshiftParserListener) ExitTable_property(ctx *Table_propertyContext) {}

// EnterCreateexternalviewstmt is called when production createexternalviewstmt is entered.
func (s *BaseRedshiftParserListener) EnterCreateexternalviewstmt(ctx *CreateexternalviewstmtContext) {
}

// ExitCreateexternalviewstmt is called when production createexternalviewstmt is exited.
func (s *BaseRedshiftParserListener) ExitCreateexternalviewstmt(ctx *CreateexternalviewstmtContext) {}

// EnterDropexternalviewstmt is called when production dropexternalviewstmt is entered.
func (s *BaseRedshiftParserListener) EnterDropexternalviewstmt(ctx *DropexternalviewstmtContext) {}

// ExitDropexternalviewstmt is called when production dropexternalviewstmt is exited.
func (s *BaseRedshiftParserListener) ExitDropexternalviewstmt(ctx *DropexternalviewstmtContext) {}

// EnterAlteridentityproviderstmt is called when production alteridentityproviderstmt is entered.
func (s *BaseRedshiftParserListener) EnterAlteridentityproviderstmt(ctx *AlteridentityproviderstmtContext) {
}

// ExitAlteridentityproviderstmt is called when production alteridentityproviderstmt is exited.
func (s *BaseRedshiftParserListener) ExitAlteridentityproviderstmt(ctx *AlteridentityproviderstmtContext) {
}

// EnterAlteridprovideropts is called when production alteridprovideropts is entered.
func (s *BaseRedshiftParserListener) EnterAlteridprovideropts(ctx *AlteridprovideroptsContext) {}

// ExitAlteridprovideropts is called when production alteridprovideropts is exited.
func (s *BaseRedshiftParserListener) ExitAlteridprovideropts(ctx *AlteridprovideroptsContext) {}

// EnterAltermaskingpolicystmt is called when production altermaskingpolicystmt is entered.
func (s *BaseRedshiftParserListener) EnterAltermaskingpolicystmt(ctx *AltermaskingpolicystmtContext) {
}

// ExitAltermaskingpolicystmt is called when production altermaskingpolicystmt is exited.
func (s *BaseRedshiftParserListener) ExitAltermaskingpolicystmt(ctx *AltermaskingpolicystmtContext) {}

// EnterAltermaterializedviewstmt is called when production altermaterializedviewstmt is entered.
func (s *BaseRedshiftParserListener) EnterAltermaterializedviewstmt(ctx *AltermaterializedviewstmtContext) {
}

// ExitAltermaterializedviewstmt is called when production altermaterializedviewstmt is exited.
func (s *BaseRedshiftParserListener) ExitAltermaterializedviewstmt(ctx *AltermaterializedviewstmtContext) {
}

// EnterAltmaskingpolicyopts is called when production altmaskingpolicyopts is entered.
func (s *BaseRedshiftParserListener) EnterAltmaskingpolicyopts(ctx *AltmaskingpolicyoptsContext) {}

// ExitAltmaskingpolicyopts is called when production altmaskingpolicyopts is exited.
func (s *BaseRedshiftParserListener) ExitAltmaskingpolicyopts(ctx *AltmaskingpolicyoptsContext) {}

// EnterAltmaskingpolicyargs is called when production altmaskingpolicyargs is entered.
func (s *BaseRedshiftParserListener) EnterAltmaskingpolicyargs(ctx *AltmaskingpolicyargsContext) {}

// ExitAltmaskingpolicyargs is called when production altmaskingpolicyargs is exited.
func (s *BaseRedshiftParserListener) ExitAltmaskingpolicyargs(ctx *AltmaskingpolicyargsContext) {}

// EnterAltmaskingpolicyarg is called when production altmaskingpolicyarg is entered.
func (s *BaseRedshiftParserListener) EnterAltmaskingpolicyarg(ctx *AltmaskingpolicyargContext) {}

// ExitAltmaskingpolicyarg is called when production altmaskingpolicyarg is exited.
func (s *BaseRedshiftParserListener) ExitAltmaskingpolicyarg(ctx *AltmaskingpolicyargContext) {}

// EnterAlterrlspolicystmt is called when production alterrlspolicystmt is entered.
func (s *BaseRedshiftParserListener) EnterAlterrlspolicystmt(ctx *AlterrlspolicystmtContext) {}

// ExitAlterrlspolicystmt is called when production alterrlspolicystmt is exited.
func (s *BaseRedshiftParserListener) ExitAlterrlspolicystmt(ctx *AlterrlspolicystmtContext) {}

// EnterAttachmaskingpolicystmt is called when production attachmaskingpolicystmt is entered.
func (s *BaseRedshiftParserListener) EnterAttachmaskingpolicystmt(ctx *AttachmaskingpolicystmtContext) {
}

// ExitAttachmaskingpolicystmt is called when production attachmaskingpolicystmt is exited.
func (s *BaseRedshiftParserListener) ExitAttachmaskingpolicystmt(ctx *AttachmaskingpolicystmtContext) {
}

// EnterAttachpolicycollist is called when production attachpolicycollist is entered.
func (s *BaseRedshiftParserListener) EnterAttachpolicycollist(ctx *AttachpolicycollistContext) {}

// ExitAttachpolicycollist is called when production attachpolicycollist is exited.
func (s *BaseRedshiftParserListener) ExitAttachpolicycollist(ctx *AttachpolicycollistContext) {}

// EnterAttachpolicycolumn is called when production attachpolicycolumn is entered.
func (s *BaseRedshiftParserListener) EnterAttachpolicycolumn(ctx *AttachpolicycolumnContext) {}

// ExitAttachpolicycolumn is called when production attachpolicycolumn is exited.
func (s *BaseRedshiftParserListener) ExitAttachpolicycolumn(ctx *AttachpolicycolumnContext) {}

// EnterAttachpolicytargets is called when production attachpolicytargets is entered.
func (s *BaseRedshiftParserListener) EnterAttachpolicytargets(ctx *AttachpolicytargetsContext) {}

// ExitAttachpolicytargets is called when production attachpolicytargets is exited.
func (s *BaseRedshiftParserListener) ExitAttachpolicytargets(ctx *AttachpolicytargetsContext) {}

// EnterAttachpolicytarget is called when production attachpolicytarget is entered.
func (s *BaseRedshiftParserListener) EnterAttachpolicytarget(ctx *AttachpolicytargetContext) {}

// ExitAttachpolicytarget is called when production attachpolicytarget is exited.
func (s *BaseRedshiftParserListener) ExitAttachpolicytarget(ctx *AttachpolicytargetContext) {}

// EnterAttachrlspolicystmt is called when production attachrlspolicystmt is entered.
func (s *BaseRedshiftParserListener) EnterAttachrlspolicystmt(ctx *AttachrlspolicystmtContext) {}

// ExitAttachrlspolicystmt is called when production attachrlspolicystmt is exited.
func (s *BaseRedshiftParserListener) ExitAttachrlspolicystmt(ctx *AttachrlspolicystmtContext) {}

// EnterTable_name_list is called when production table_name_list is entered.
func (s *BaseRedshiftParserListener) EnterTable_name_list(ctx *Table_name_listContext) {}

// ExitTable_name_list is called when production table_name_list is exited.
func (s *BaseRedshiftParserListener) ExitTable_name_list(ctx *Table_name_listContext) {}

// EnterCreateidentityproviderstmt is called when production createidentityproviderstmt is entered.
func (s *BaseRedshiftParserListener) EnterCreateidentityproviderstmt(ctx *CreateidentityproviderstmtContext) {
}

// ExitCreateidentityproviderstmt is called when production createidentityproviderstmt is exited.
func (s *BaseRedshiftParserListener) ExitCreateidentityproviderstmt(ctx *CreateidentityproviderstmtContext) {
}

// EnterCreateidprovideropts is called when production createidprovideropts is entered.
func (s *BaseRedshiftParserListener) EnterCreateidprovideropts(ctx *CreateidprovideroptsContext) {}

// ExitCreateidprovideropts is called when production createidprovideropts is exited.
func (s *BaseRedshiftParserListener) ExitCreateidprovideropts(ctx *CreateidprovideroptsContext) {}

// EnterGroupfilter is called when production groupfilter is entered.
func (s *BaseRedshiftParserListener) EnterGroupfilter(ctx *GroupfilterContext) {}

// ExitGroupfilter is called when production groupfilter is exited.
func (s *BaseRedshiftParserListener) ExitGroupfilter(ctx *GroupfilterContext) {}

// EnterCreatelibrarystmt is called when production createlibrarystmt is entered.
func (s *BaseRedshiftParserListener) EnterCreatelibrarystmt(ctx *CreatelibrarystmtContext) {}

// ExitCreatelibrarystmt is called when production createlibrarystmt is exited.
func (s *BaseRedshiftParserListener) ExitCreatelibrarystmt(ctx *CreatelibrarystmtContext) {}

// EnterCreatelibraryopts is called when production createlibraryopts is entered.
func (s *BaseRedshiftParserListener) EnterCreatelibraryopts(ctx *CreatelibraryoptsContext) {}

// ExitCreatelibraryopts is called when production createlibraryopts is exited.
func (s *BaseRedshiftParserListener) ExitCreatelibraryopts(ctx *CreatelibraryoptsContext) {}

// EnterCreatemaskingpolicystmt is called when production createmaskingpolicystmt is entered.
func (s *BaseRedshiftParserListener) EnterCreatemaskingpolicystmt(ctx *CreatemaskingpolicystmtContext) {
}

// ExitCreatemaskingpolicystmt is called when production createmaskingpolicystmt is exited.
func (s *BaseRedshiftParserListener) ExitCreatemaskingpolicystmt(ctx *CreatemaskingpolicystmtContext) {
}

// EnterInputcolumnlist is called when production inputcolumnlist is entered.
func (s *BaseRedshiftParserListener) EnterInputcolumnlist(ctx *InputcolumnlistContext) {}

// ExitInputcolumnlist is called when production inputcolumnlist is exited.
func (s *BaseRedshiftParserListener) ExitInputcolumnlist(ctx *InputcolumnlistContext) {}

// EnterInputcolumn is called when production inputcolumn is entered.
func (s *BaseRedshiftParserListener) EnterInputcolumn(ctx *InputcolumnContext) {}

// ExitInputcolumn is called when production inputcolumn is exited.
func (s *BaseRedshiftParserListener) ExitInputcolumn(ctx *InputcolumnContext) {}

// EnterMaskingexpression is called when production maskingexpression is entered.
func (s *BaseRedshiftParserListener) EnterMaskingexpression(ctx *MaskingexpressionContext) {}

// ExitMaskingexpression is called when production maskingexpression is exited.
func (s *BaseRedshiftParserListener) ExitMaskingexpression(ctx *MaskingexpressionContext) {}

// EnterCreatemodelstmt is called when production createmodelstmt is entered.
func (s *BaseRedshiftParserListener) EnterCreatemodelstmt(ctx *CreatemodelstmtContext) {}

// ExitCreatemodelstmt is called when production createmodelstmt is exited.
func (s *BaseRedshiftParserListener) ExitCreatemodelstmt(ctx *CreatemodelstmtContext) {}

// EnterCreatemodelfromclause is called when production createmodelfromclause is entered.
func (s *BaseRedshiftParserListener) EnterCreatemodelfromclause(ctx *CreatemodelfromclauseContext) {}

// ExitCreatemodelfromclause is called when production createmodelfromclause is exited.
func (s *BaseRedshiftParserListener) ExitCreatemodelfromclause(ctx *CreatemodelfromclauseContext) {}

// EnterIamrolespec is called when production iamrolespec is entered.
func (s *BaseRedshiftParserListener) EnterIamrolespec(ctx *IamrolespecContext) {}

// ExitIamrolespec is called when production iamrolespec is exited.
func (s *BaseRedshiftParserListener) ExitIamrolespec(ctx *IamrolespecContext) {}

// EnterSagemakerspec is called when production sagemakerspec is entered.
func (s *BaseRedshiftParserListener) EnterSagemakerspec(ctx *SagemakerspecContext) {}

// ExitSagemakerspec is called when production sagemakerspec is exited.
func (s *BaseRedshiftParserListener) ExitSagemakerspec(ctx *SagemakerspecContext) {}

// EnterModeltypespec is called when production modeltypespec is entered.
func (s *BaseRedshiftParserListener) EnterModeltypespec(ctx *ModeltypespecContext) {}

// ExitModeltypespec is called when production modeltypespec is exited.
func (s *BaseRedshiftParserListener) ExitModeltypespec(ctx *ModeltypespecContext) {}

// EnterProblemtypespec is called when production problemtypespec is entered.
func (s *BaseRedshiftParserListener) EnterProblemtypespec(ctx *ProblemtypespecContext) {}

// ExitProblemtypespec is called when production problemtypespec is exited.
func (s *BaseRedshiftParserListener) ExitProblemtypespec(ctx *ProblemtypespecContext) {}

// EnterProblemtype is called when production problemtype is entered.
func (s *BaseRedshiftParserListener) EnterProblemtype(ctx *ProblemtypeContext) {}

// ExitProblemtype is called when production problemtype is exited.
func (s *BaseRedshiftParserListener) ExitProblemtype(ctx *ProblemtypeContext) {}

// EnterObjectivespec is called when production objectivespec is entered.
func (s *BaseRedshiftParserListener) EnterObjectivespec(ctx *ObjectivespecContext) {}

// ExitObjectivespec is called when production objectivespec is exited.
func (s *BaseRedshiftParserListener) ExitObjectivespec(ctx *ObjectivespecContext) {}

// EnterHyperparametersspec is called when production hyperparametersspec is entered.
func (s *BaseRedshiftParserListener) EnterHyperparametersspec(ctx *HyperparametersspecContext) {}

// ExitHyperparametersspec is called when production hyperparametersspec is exited.
func (s *BaseRedshiftParserListener) ExitHyperparametersspec(ctx *HyperparametersspecContext) {}

// EnterHyperparameterslist is called when production hyperparameterslist is entered.
func (s *BaseRedshiftParserListener) EnterHyperparameterslist(ctx *HyperparameterslistContext) {}

// ExitHyperparameterslist is called when production hyperparameterslist is exited.
func (s *BaseRedshiftParserListener) ExitHyperparameterslist(ctx *HyperparameterslistContext) {}

// EnterHyperparameteritem is called when production hyperparameteritem is entered.
func (s *BaseRedshiftParserListener) EnterHyperparameteritem(ctx *HyperparameteritemContext) {}

// ExitHyperparameteritem is called when production hyperparameteritem is exited.
func (s *BaseRedshiftParserListener) ExitHyperparameteritem(ctx *HyperparameteritemContext) {}

// EnterSettingsclause is called when production settingsclause is entered.
func (s *BaseRedshiftParserListener) EnterSettingsclause(ctx *SettingsclauseContext) {}

// ExitSettingsclause is called when production settingsclause is exited.
func (s *BaseRedshiftParserListener) ExitSettingsclause(ctx *SettingsclauseContext) {}

// EnterSettingsitem is called when production settingsitem is entered.
func (s *BaseRedshiftParserListener) EnterSettingsitem(ctx *SettingsitemContext) {}

// ExitSettingsitem is called when production settingsitem is exited.
func (s *BaseRedshiftParserListener) ExitSettingsitem(ctx *SettingsitemContext) {}

// EnterDatatypelist is called when production datatypelist is entered.
func (s *BaseRedshiftParserListener) EnterDatatypelist(ctx *DatatypelistContext) {}

// ExitDatatypelist is called when production datatypelist is exited.
func (s *BaseRedshiftParserListener) ExitDatatypelist(ctx *DatatypelistContext) {}

// EnterDatatype is called when production datatype is entered.
func (s *BaseRedshiftParserListener) EnterDatatype(ctx *DatatypeContext) {}

// ExitDatatype is called when production datatype is exited.
func (s *BaseRedshiftParserListener) ExitDatatype(ctx *DatatypeContext) {}

// EnterCreaterlspolicystmt is called when production createrlspolicystmt is entered.
func (s *BaseRedshiftParserListener) EnterCreaterlspolicystmt(ctx *CreaterlspolicystmtContext) {}

// ExitCreaterlspolicystmt is called when production createrlspolicystmt is exited.
func (s *BaseRedshiftParserListener) ExitCreaterlspolicystmt(ctx *CreaterlspolicystmtContext) {}

// EnterDescidentityproviderstmt is called when production descidentityproviderstmt is entered.
func (s *BaseRedshiftParserListener) EnterDescidentityproviderstmt(ctx *DescidentityproviderstmtContext) {
}

// ExitDescidentityproviderstmt is called when production descidentityproviderstmt is exited.
func (s *BaseRedshiftParserListener) ExitDescidentityproviderstmt(ctx *DescidentityproviderstmtContext) {
}

// EnterDetachmaskingpolicystmt is called when production detachmaskingpolicystmt is entered.
func (s *BaseRedshiftParserListener) EnterDetachmaskingpolicystmt(ctx *DetachmaskingpolicystmtContext) {
}

// ExitDetachmaskingpolicystmt is called when production detachmaskingpolicystmt is exited.
func (s *BaseRedshiftParserListener) ExitDetachmaskingpolicystmt(ctx *DetachmaskingpolicystmtContext) {
}

// EnterDetachrlspolicystmt is called when production detachrlspolicystmt is entered.
func (s *BaseRedshiftParserListener) EnterDetachrlspolicystmt(ctx *DetachrlspolicystmtContext) {}

// ExitDetachrlspolicystmt is called when production detachrlspolicystmt is exited.
func (s *BaseRedshiftParserListener) ExitDetachrlspolicystmt(ctx *DetachrlspolicystmtContext) {}

// EnterRole_or_user_or_public_list is called when production role_or_user_or_public_list is entered.
func (s *BaseRedshiftParserListener) EnterRole_or_user_or_public_list(ctx *Role_or_user_or_public_listContext) {
}

// ExitRole_or_user_or_public_list is called when production role_or_user_or_public_list is exited.
func (s *BaseRedshiftParserListener) ExitRole_or_user_or_public_list(ctx *Role_or_user_or_public_listContext) {
}

// EnterRole_or_user_or_public is called when production role_or_user_or_public is entered.
func (s *BaseRedshiftParserListener) EnterRole_or_user_or_public(ctx *Role_or_user_or_publicContext) {
}

// ExitRole_or_user_or_public is called when production role_or_user_or_public is exited.
func (s *BaseRedshiftParserListener) ExitRole_or_user_or_public(ctx *Role_or_user_or_publicContext) {}

// EnterRlspolicyname is called when production rlspolicyname is entered.
func (s *BaseRedshiftParserListener) EnterRlspolicyname(ctx *RlspolicynameContext) {}

// ExitRlspolicyname is called when production rlspolicyname is exited.
func (s *BaseRedshiftParserListener) ExitRlspolicyname(ctx *RlspolicynameContext) {}

// EnterDropidentityproviderstmt is called when production dropidentityproviderstmt is entered.
func (s *BaseRedshiftParserListener) EnterDropidentityproviderstmt(ctx *DropidentityproviderstmtContext) {
}

// ExitDropidentityproviderstmt is called when production dropidentityproviderstmt is exited.
func (s *BaseRedshiftParserListener) ExitDropidentityproviderstmt(ctx *DropidentityproviderstmtContext) {
}

// EnterDroplibrarystmt is called when production droplibrarystmt is entered.
func (s *BaseRedshiftParserListener) EnterDroplibrarystmt(ctx *DroplibrarystmtContext) {}

// ExitDroplibrarystmt is called when production droplibrarystmt is exited.
func (s *BaseRedshiftParserListener) ExitDroplibrarystmt(ctx *DroplibrarystmtContext) {}

// EnterDropmaskingpolicystmt is called when production dropmaskingpolicystmt is entered.
func (s *BaseRedshiftParserListener) EnterDropmaskingpolicystmt(ctx *DropmaskingpolicystmtContext) {}

// ExitDropmaskingpolicystmt is called when production dropmaskingpolicystmt is exited.
func (s *BaseRedshiftParserListener) ExitDropmaskingpolicystmt(ctx *DropmaskingpolicystmtContext) {}

// EnterDropmodelstmt is called when production dropmodelstmt is entered.
func (s *BaseRedshiftParserListener) EnterDropmodelstmt(ctx *DropmodelstmtContext) {}

// ExitDropmodelstmt is called when production dropmodelstmt is exited.
func (s *BaseRedshiftParserListener) ExitDropmodelstmt(ctx *DropmodelstmtContext) {}

// EnterDroprlspolicystmt is called when production droprlspolicystmt is entered.
func (s *BaseRedshiftParserListener) EnterDroprlspolicystmt(ctx *DroprlspolicystmtContext) {}

// ExitDroprlspolicystmt is called when production droprlspolicystmt is exited.
func (s *BaseRedshiftParserListener) ExitDroprlspolicystmt(ctx *DroprlspolicystmtContext) {}

// EnterAltertableappendstmt is called when production altertableappendstmt is entered.
func (s *BaseRedshiftParserListener) EnterAltertableappendstmt(ctx *AltertableappendstmtContext) {}

// ExitAltertableappendstmt is called when production altertableappendstmt is exited.
func (s *BaseRedshiftParserListener) ExitAltertableappendstmt(ctx *AltertableappendstmtContext) {}

// EnterAppendoptions is called when production appendoptions is entered.
func (s *BaseRedshiftParserListener) EnterAppendoptions(ctx *AppendoptionsContext) {}

// ExitAppendoptions is called when production appendoptions is exited.
func (s *BaseRedshiftParserListener) ExitAppendoptions(ctx *AppendoptionsContext) {}

// EnterAlteruserstmt is called when production alteruserstmt is entered.
func (s *BaseRedshiftParserListener) EnterAlteruserstmt(ctx *AlteruserstmtContext) {}

// ExitAlteruserstmt is called when production alteruserstmt is exited.
func (s *BaseRedshiftParserListener) ExitAlteruserstmt(ctx *AlteruserstmtContext) {}

// EnterAlteruseropts is called when production alteruseropts is entered.
func (s *BaseRedshiftParserListener) EnterAlteruseropts(ctx *AlteruseroptsContext) {}

// ExitAlteruseropts is called when production alteruseropts is exited.
func (s *BaseRedshiftParserListener) ExitAlteruseropts(ctx *AlteruseroptsContext) {}

// EnterAnalyzecompressionstmt is called when production analyzecompressionstmt is entered.
func (s *BaseRedshiftParserListener) EnterAnalyzecompressionstmt(ctx *AnalyzecompressionstmtContext) {
}

// ExitAnalyzecompressionstmt is called when production analyzecompressionstmt is exited.
func (s *BaseRedshiftParserListener) ExitAnalyzecompressionstmt(ctx *AnalyzecompressionstmtContext) {}

// EnterCancelstmt is called when production cancelstmt is entered.
func (s *BaseRedshiftParserListener) EnterCancelstmt(ctx *CancelstmtContext) {}

// ExitCancelstmt is called when production cancelstmt is exited.
func (s *BaseRedshiftParserListener) ExitCancelstmt(ctx *CancelstmtContext) {}

// EnterClosestmt is called when production closestmt is entered.
func (s *BaseRedshiftParserListener) EnterClosestmt(ctx *ClosestmtContext) {}

// ExitClosestmt is called when production closestmt is exited.
func (s *BaseRedshiftParserListener) ExitClosestmt(ctx *ClosestmtContext) {}

// EnterInsertexternaltablestmt is called when production insertexternaltablestmt is entered.
func (s *BaseRedshiftParserListener) EnterInsertexternaltablestmt(ctx *InsertexternaltablestmtContext) {
}

// ExitInsertexternaltablestmt is called when production insertexternaltablestmt is exited.
func (s *BaseRedshiftParserListener) ExitInsertexternaltablestmt(ctx *InsertexternaltablestmtContext) {
}

// EnterSelect_or_values is called when production select_or_values is entered.
func (s *BaseRedshiftParserListener) EnterSelect_or_values(ctx *Select_or_valuesContext) {}

// ExitSelect_or_values is called when production select_or_values is exited.
func (s *BaseRedshiftParserListener) ExitSelect_or_values(ctx *Select_or_valuesContext) {}

// EnterSelectintostmt is called when production selectintostmt is entered.
func (s *BaseRedshiftParserListener) EnterSelectintostmt(ctx *SelectintostmtContext) {}

// ExitSelectintostmt is called when production selectintostmt is exited.
func (s *BaseRedshiftParserListener) ExitSelectintostmt(ctx *SelectintostmtContext) {}

// EnterSetsessionauthorizationstmt is called when production setsessionauthorizationstmt is entered.
func (s *BaseRedshiftParserListener) EnterSetsessionauthorizationstmt(ctx *SetsessionauthorizationstmtContext) {
}

// ExitSetsessionauthorizationstmt is called when production setsessionauthorizationstmt is exited.
func (s *BaseRedshiftParserListener) ExitSetsessionauthorizationstmt(ctx *SetsessionauthorizationstmtContext) {
}

// EnterSetsessioncharacteristicsstmt is called when production setsessioncharacteristicsstmt is entered.
func (s *BaseRedshiftParserListener) EnterSetsessioncharacteristicsstmt(ctx *SetsessioncharacteristicsstmtContext) {
}

// ExitSetsessioncharacteristicsstmt is called when production setsessioncharacteristicsstmt is exited.
func (s *BaseRedshiftParserListener) ExitSetsessioncharacteristicsstmt(ctx *SetsessioncharacteristicsstmtContext) {
}

// EnterShowcolumnsstmt is called when production showcolumnsstmt is entered.
func (s *BaseRedshiftParserListener) EnterShowcolumnsstmt(ctx *ShowcolumnsstmtContext) {}

// ExitShowcolumnsstmt is called when production showcolumnsstmt is exited.
func (s *BaseRedshiftParserListener) ExitShowcolumnsstmt(ctx *ShowcolumnsstmtContext) {}

// EnterShowdatabasesstmt is called when production showdatabasesstmt is entered.
func (s *BaseRedshiftParserListener) EnterShowdatabasesstmt(ctx *ShowdatabasesstmtContext) {}

// ExitShowdatabasesstmt is called when production showdatabasesstmt is exited.
func (s *BaseRedshiftParserListener) ExitShowdatabasesstmt(ctx *ShowdatabasesstmtContext) {}

// EnterShowdbsopts is called when production showdbsopts is entered.
func (s *BaseRedshiftParserListener) EnterShowdbsopts(ctx *ShowdbsoptsContext) {}

// ExitShowdbsopts is called when production showdbsopts is exited.
func (s *BaseRedshiftParserListener) ExitShowdbsopts(ctx *ShowdbsoptsContext) {}

// EnterShowdatasharesstmt is called when production showdatasharesstmt is entered.
func (s *BaseRedshiftParserListener) EnterShowdatasharesstmt(ctx *ShowdatasharesstmtContext) {}

// ExitShowdatasharesstmt is called when production showdatasharesstmt is exited.
func (s *BaseRedshiftParserListener) ExitShowdatasharesstmt(ctx *ShowdatasharesstmtContext) {}

// EnterShowexternaltablestmt is called when production showexternaltablestmt is entered.
func (s *BaseRedshiftParserListener) EnterShowexternaltablestmt(ctx *ShowexternaltablestmtContext) {}

// ExitShowexternaltablestmt is called when production showexternaltablestmt is exited.
func (s *BaseRedshiftParserListener) ExitShowexternaltablestmt(ctx *ShowexternaltablestmtContext) {}

// EnterShowgrantsstmt is called when production showgrantsstmt is entered.
func (s *BaseRedshiftParserListener) EnterShowgrantsstmt(ctx *ShowgrantsstmtContext) {}

// ExitShowgrantsstmt is called when production showgrantsstmt is exited.
func (s *BaseRedshiftParserListener) ExitShowgrantsstmt(ctx *ShowgrantsstmtContext) {}

// EnterGrantobject is called when production grantobject is entered.
func (s *BaseRedshiftParserListener) EnterGrantobject(ctx *GrantobjectContext) {}

// ExitGrantobject is called when production grantobject is exited.
func (s *BaseRedshiftParserListener) ExitGrantobject(ctx *GrantobjectContext) {}

// EnterGrantprincipal is called when production grantprincipal is entered.
func (s *BaseRedshiftParserListener) EnterGrantprincipal(ctx *GrantprincipalContext) {}

// ExitGrantprincipal is called when production grantprincipal is exited.
func (s *BaseRedshiftParserListener) ExitGrantprincipal(ctx *GrantprincipalContext) {}

// EnterShowmodelstmt is called when production showmodelstmt is entered.
func (s *BaseRedshiftParserListener) EnterShowmodelstmt(ctx *ShowmodelstmtContext) {}

// ExitShowmodelstmt is called when production showmodelstmt is exited.
func (s *BaseRedshiftParserListener) ExitShowmodelstmt(ctx *ShowmodelstmtContext) {}

// EnterShowprocedurestmt is called when production showprocedurestmt is entered.
func (s *BaseRedshiftParserListener) EnterShowprocedurestmt(ctx *ShowprocedurestmtContext) {}

// ExitShowprocedurestmt is called when production showprocedurestmt is exited.
func (s *BaseRedshiftParserListener) ExitShowprocedurestmt(ctx *ShowprocedurestmtContext) {}

// EnterShowschemasstmt is called when production showschemasstmt is entered.
func (s *BaseRedshiftParserListener) EnterShowschemasstmt(ctx *ShowschemasstmtContext) {}

// ExitShowschemasstmt is called when production showschemasstmt is exited.
func (s *BaseRedshiftParserListener) ExitShowschemasstmt(ctx *ShowschemasstmtContext) {}

// EnterShowtablestmt is called when production showtablestmt is entered.
func (s *BaseRedshiftParserListener) EnterShowtablestmt(ctx *ShowtablestmtContext) {}

// ExitShowtablestmt is called when production showtablestmt is exited.
func (s *BaseRedshiftParserListener) ExitShowtablestmt(ctx *ShowtablestmtContext) {}

// EnterShowtablesstmt is called when production showtablesstmt is entered.
func (s *BaseRedshiftParserListener) EnterShowtablesstmt(ctx *ShowtablesstmtContext) {}

// ExitShowtablesstmt is called when production showtablesstmt is exited.
func (s *BaseRedshiftParserListener) ExitShowtablesstmt(ctx *ShowtablesstmtContext) {}

// EnterShowviewstmt is called when production showviewstmt is entered.
func (s *BaseRedshiftParserListener) EnterShowviewstmt(ctx *ShowviewstmtContext) {}

// ExitShowviewstmt is called when production showviewstmt is exited.
func (s *BaseRedshiftParserListener) ExitShowviewstmt(ctx *ShowviewstmtContext) {}

// EnterUnloadstmt is called when production unloadstmt is entered.
func (s *BaseRedshiftParserListener) EnterUnloadstmt(ctx *UnloadstmtContext) {}

// ExitUnloadstmt is called when production unloadstmt is exited.
func (s *BaseRedshiftParserListener) ExitUnloadstmt(ctx *UnloadstmtContext) {}

// EnterIamroleclause is called when production iamroleclause is entered.
func (s *BaseRedshiftParserListener) EnterIamroleclause(ctx *IamroleclauseContext) {}

// ExitIamroleclause is called when production iamroleclause is exited.
func (s *BaseRedshiftParserListener) ExitIamroleclause(ctx *IamroleclauseContext) {}

// EnterUnloadoptions is called when production unloadoptions is entered.
func (s *BaseRedshiftParserListener) EnterUnloadoptions(ctx *UnloadoptionsContext) {}

// ExitUnloadoptions is called when production unloadoptions is exited.
func (s *BaseRedshiftParserListener) ExitUnloadoptions(ctx *UnloadoptionsContext) {}

// EnterFormatoption is called when production formatoption is entered.
func (s *BaseRedshiftParserListener) EnterFormatoption(ctx *FormatoptionContext) {}

// ExitFormatoption is called when production formatoption is exited.
func (s *BaseRedshiftParserListener) ExitFormatoption(ctx *FormatoptionContext) {}

// EnterPartitionbyoption is called when production partitionbyoption is entered.
func (s *BaseRedshiftParserListener) EnterPartitionbyoption(ctx *PartitionbyoptionContext) {}

// ExitPartitionbyoption is called when production partitionbyoption is exited.
func (s *BaseRedshiftParserListener) ExitPartitionbyoption(ctx *PartitionbyoptionContext) {}

// EnterManifestoption is called when production manifestoption is entered.
func (s *BaseRedshiftParserListener) EnterManifestoption(ctx *ManifestoptionContext) {}

// ExitManifestoption is called when production manifestoption is exited.
func (s *BaseRedshiftParserListener) ExitManifestoption(ctx *ManifestoptionContext) {}

// EnterHeaderoption is called when production headeroption is entered.
func (s *BaseRedshiftParserListener) EnterHeaderoption(ctx *HeaderoptionContext) {}

// ExitHeaderoption is called when production headeroption is exited.
func (s *BaseRedshiftParserListener) ExitHeaderoption(ctx *HeaderoptionContext) {}

// EnterDelimiteroption is called when production delimiteroption is entered.
func (s *BaseRedshiftParserListener) EnterDelimiteroption(ctx *DelimiteroptionContext) {}

// ExitDelimiteroption is called when production delimiteroption is exited.
func (s *BaseRedshiftParserListener) ExitDelimiteroption(ctx *DelimiteroptionContext) {}

// EnterFixedwidthoption is called when production fixedwidthoption is entered.
func (s *BaseRedshiftParserListener) EnterFixedwidthoption(ctx *FixedwidthoptionContext) {}

// ExitFixedwidthoption is called when production fixedwidthoption is exited.
func (s *BaseRedshiftParserListener) ExitFixedwidthoption(ctx *FixedwidthoptionContext) {}

// EnterEncryptedoption is called when production encryptedoption is entered.
func (s *BaseRedshiftParserListener) EnterEncryptedoption(ctx *EncryptedoptionContext) {}

// ExitEncryptedoption is called when production encryptedoption is exited.
func (s *BaseRedshiftParserListener) ExitEncryptedoption(ctx *EncryptedoptionContext) {}

// EnterKmskeyoption is called when production kmskeyoption is entered.
func (s *BaseRedshiftParserListener) EnterKmskeyoption(ctx *KmskeyoptionContext) {}

// ExitKmskeyoption is called when production kmskeyoption is exited.
func (s *BaseRedshiftParserListener) ExitKmskeyoption(ctx *KmskeyoptionContext) {}

// EnterCompressionoption is called when production compressionoption is entered.
func (s *BaseRedshiftParserListener) EnterCompressionoption(ctx *CompressionoptionContext) {}

// ExitCompressionoption is called when production compressionoption is exited.
func (s *BaseRedshiftParserListener) ExitCompressionoption(ctx *CompressionoptionContext) {}

// EnterAddquotesoption is called when production addquotesoption is entered.
func (s *BaseRedshiftParserListener) EnterAddquotesoption(ctx *AddquotesoptionContext) {}

// ExitAddquotesoption is called when production addquotesoption is exited.
func (s *BaseRedshiftParserListener) ExitAddquotesoption(ctx *AddquotesoptionContext) {}

// EnterNullasoption is called when production nullasoption is entered.
func (s *BaseRedshiftParserListener) EnterNullasoption(ctx *NullasoptionContext) {}

// ExitNullasoption is called when production nullasoption is exited.
func (s *BaseRedshiftParserListener) ExitNullasoption(ctx *NullasoptionContext) {}

// EnterEscapeoption is called when production escapeoption is entered.
func (s *BaseRedshiftParserListener) EnterEscapeoption(ctx *EscapeoptionContext) {}

// ExitEscapeoption is called when production escapeoption is exited.
func (s *BaseRedshiftParserListener) ExitEscapeoption(ctx *EscapeoptionContext) {}

// EnterAllowoverwriteoption is called when production allowoverwriteoption is entered.
func (s *BaseRedshiftParserListener) EnterAllowoverwriteoption(ctx *AllowoverwriteoptionContext) {}

// ExitAllowoverwriteoption is called when production allowoverwriteoption is exited.
func (s *BaseRedshiftParserListener) ExitAllowoverwriteoption(ctx *AllowoverwriteoptionContext) {}

// EnterCleanpathoption is called when production cleanpathoption is entered.
func (s *BaseRedshiftParserListener) EnterCleanpathoption(ctx *CleanpathoptionContext) {}

// ExitCleanpathoption is called when production cleanpathoption is exited.
func (s *BaseRedshiftParserListener) ExitCleanpathoption(ctx *CleanpathoptionContext) {}

// EnterParalleloption is called when production paralleloption is entered.
func (s *BaseRedshiftParserListener) EnterParalleloption(ctx *ParalleloptionContext) {}

// ExitParalleloption is called when production paralleloption is exited.
func (s *BaseRedshiftParserListener) ExitParalleloption(ctx *ParalleloptionContext) {}

// EnterMaxfilesizeoption is called when production maxfilesizeoption is entered.
func (s *BaseRedshiftParserListener) EnterMaxfilesizeoption(ctx *MaxfilesizeoptionContext) {}

// ExitMaxfilesizeoption is called when production maxfilesizeoption is exited.
func (s *BaseRedshiftParserListener) ExitMaxfilesizeoption(ctx *MaxfilesizeoptionContext) {}

// EnterRowgroupsizeoption is called when production rowgroupsizeoption is entered.
func (s *BaseRedshiftParserListener) EnterRowgroupsizeoption(ctx *RowgroupsizeoptionContext) {}

// ExitRowgroupsizeoption is called when production rowgroupsizeoption is exited.
func (s *BaseRedshiftParserListener) ExitRowgroupsizeoption(ctx *RowgroupsizeoptionContext) {}

// EnterSizeunit is called when production sizeunit is entered.
func (s *BaseRedshiftParserListener) EnterSizeunit(ctx *SizeunitContext) {}

// ExitSizeunit is called when production sizeunit is exited.
func (s *BaseRedshiftParserListener) ExitSizeunit(ctx *SizeunitContext) {}

// EnterRegionoption is called when production regionoption is entered.
func (s *BaseRedshiftParserListener) EnterRegionoption(ctx *RegionoptionContext) {}

// ExitRegionoption is called when production regionoption is exited.
func (s *BaseRedshiftParserListener) ExitRegionoption(ctx *RegionoptionContext) {}

// EnterExtensionoption is called when production extensionoption is entered.
func (s *BaseRedshiftParserListener) EnterExtensionoption(ctx *ExtensionoptionContext) {}

// ExitExtensionoption is called when production extensionoption is exited.
func (s *BaseRedshiftParserListener) ExitExtensionoption(ctx *ExtensionoptionContext) {}

// EnterUsestmt is called when production usestmt is entered.
func (s *BaseRedshiftParserListener) EnterUsestmt(ctx *UsestmtContext) {}

// ExitUsestmt is called when production usestmt is exited.
func (s *BaseRedshiftParserListener) ExitUsestmt(ctx *UsestmtContext) {}

// EnterCreatedbstmt is called when production createdbstmt is entered.
func (s *BaseRedshiftParserListener) EnterCreatedbstmt(ctx *CreatedbstmtContext) {}

// ExitCreatedbstmt is called when production createdbstmt is exited.
func (s *BaseRedshiftParserListener) ExitCreatedbstmt(ctx *CreatedbstmtContext) {}

// EnterCreatedb_opt_list is called when production createdb_opt_list is entered.
func (s *BaseRedshiftParserListener) EnterCreatedb_opt_list(ctx *Createdb_opt_listContext) {}

// ExitCreatedb_opt_list is called when production createdb_opt_list is exited.
func (s *BaseRedshiftParserListener) ExitCreatedb_opt_list(ctx *Createdb_opt_listContext) {}

// EnterCreatedb_opt_items is called when production createdb_opt_items is entered.
func (s *BaseRedshiftParserListener) EnterCreatedb_opt_items(ctx *Createdb_opt_itemsContext) {}

// ExitCreatedb_opt_items is called when production createdb_opt_items is exited.
func (s *BaseRedshiftParserListener) ExitCreatedb_opt_items(ctx *Createdb_opt_itemsContext) {}

// EnterCreatedb_opt_item is called when production createdb_opt_item is entered.
func (s *BaseRedshiftParserListener) EnterCreatedb_opt_item(ctx *Createdb_opt_itemContext) {}

// ExitCreatedb_opt_item is called when production createdb_opt_item is exited.
func (s *BaseRedshiftParserListener) ExitCreatedb_opt_item(ctx *Createdb_opt_itemContext) {}

// EnterCreatedb_opt_name is called when production createdb_opt_name is entered.
func (s *BaseRedshiftParserListener) EnterCreatedb_opt_name(ctx *Createdb_opt_nameContext) {}

// ExitCreatedb_opt_name is called when production createdb_opt_name is exited.
func (s *BaseRedshiftParserListener) ExitCreatedb_opt_name(ctx *Createdb_opt_nameContext) {}

// EnterOpt_equal is called when production opt_equal is entered.
func (s *BaseRedshiftParserListener) EnterOpt_equal(ctx *Opt_equalContext) {}

// ExitOpt_equal is called when production opt_equal is exited.
func (s *BaseRedshiftParserListener) ExitOpt_equal(ctx *Opt_equalContext) {}

// EnterAlterdatabasestmt is called when production alterdatabasestmt is entered.
func (s *BaseRedshiftParserListener) EnterAlterdatabasestmt(ctx *AlterdatabasestmtContext) {}

// ExitAlterdatabasestmt is called when production alterdatabasestmt is exited.
func (s *BaseRedshiftParserListener) ExitAlterdatabasestmt(ctx *AlterdatabasestmtContext) {}

// EnterAlterdatabasesetstmt is called when production alterdatabasesetstmt is entered.
func (s *BaseRedshiftParserListener) EnterAlterdatabasesetstmt(ctx *AlterdatabasesetstmtContext) {}

// ExitAlterdatabasesetstmt is called when production alterdatabasesetstmt is exited.
func (s *BaseRedshiftParserListener) ExitAlterdatabasesetstmt(ctx *AlterdatabasesetstmtContext) {}

// EnterDropdbstmt is called when production dropdbstmt is entered.
func (s *BaseRedshiftParserListener) EnterDropdbstmt(ctx *DropdbstmtContext) {}

// ExitDropdbstmt is called when production dropdbstmt is exited.
func (s *BaseRedshiftParserListener) ExitDropdbstmt(ctx *DropdbstmtContext) {}

// EnterDrop_option_list is called when production drop_option_list is entered.
func (s *BaseRedshiftParserListener) EnterDrop_option_list(ctx *Drop_option_listContext) {}

// ExitDrop_option_list is called when production drop_option_list is exited.
func (s *BaseRedshiftParserListener) ExitDrop_option_list(ctx *Drop_option_listContext) {}

// EnterDrop_option is called when production drop_option is entered.
func (s *BaseRedshiftParserListener) EnterDrop_option(ctx *Drop_optionContext) {}

// ExitDrop_option is called when production drop_option is exited.
func (s *BaseRedshiftParserListener) ExitDrop_option(ctx *Drop_optionContext) {}

// EnterAltercollationstmt is called when production altercollationstmt is entered.
func (s *BaseRedshiftParserListener) EnterAltercollationstmt(ctx *AltercollationstmtContext) {}

// ExitAltercollationstmt is called when production altercollationstmt is exited.
func (s *BaseRedshiftParserListener) ExitAltercollationstmt(ctx *AltercollationstmtContext) {}

// EnterAltersystemstmt is called when production altersystemstmt is entered.
func (s *BaseRedshiftParserListener) EnterAltersystemstmt(ctx *AltersystemstmtContext) {}

// ExitAltersystemstmt is called when production altersystemstmt is exited.
func (s *BaseRedshiftParserListener) ExitAltersystemstmt(ctx *AltersystemstmtContext) {}

// EnterCreatedomainstmt is called when production createdomainstmt is entered.
func (s *BaseRedshiftParserListener) EnterCreatedomainstmt(ctx *CreatedomainstmtContext) {}

// ExitCreatedomainstmt is called when production createdomainstmt is exited.
func (s *BaseRedshiftParserListener) ExitCreatedomainstmt(ctx *CreatedomainstmtContext) {}

// EnterAlterdomainstmt is called when production alterdomainstmt is entered.
func (s *BaseRedshiftParserListener) EnterAlterdomainstmt(ctx *AlterdomainstmtContext) {}

// ExitAlterdomainstmt is called when production alterdomainstmt is exited.
func (s *BaseRedshiftParserListener) ExitAlterdomainstmt(ctx *AlterdomainstmtContext) {}

// EnterOpt_as is called when production opt_as is entered.
func (s *BaseRedshiftParserListener) EnterOpt_as(ctx *Opt_asContext) {}

// ExitOpt_as is called when production opt_as is exited.
func (s *BaseRedshiftParserListener) ExitOpt_as(ctx *Opt_asContext) {}

// EnterAltertsdictionarystmt is called when production altertsdictionarystmt is entered.
func (s *BaseRedshiftParserListener) EnterAltertsdictionarystmt(ctx *AltertsdictionarystmtContext) {}

// ExitAltertsdictionarystmt is called when production altertsdictionarystmt is exited.
func (s *BaseRedshiftParserListener) ExitAltertsdictionarystmt(ctx *AltertsdictionarystmtContext) {}

// EnterAltertsconfigurationstmt is called when production altertsconfigurationstmt is entered.
func (s *BaseRedshiftParserListener) EnterAltertsconfigurationstmt(ctx *AltertsconfigurationstmtContext) {
}

// ExitAltertsconfigurationstmt is called when production altertsconfigurationstmt is exited.
func (s *BaseRedshiftParserListener) ExitAltertsconfigurationstmt(ctx *AltertsconfigurationstmtContext) {
}

// EnterAny_with is called when production any_with is entered.
func (s *BaseRedshiftParserListener) EnterAny_with(ctx *Any_withContext) {}

// ExitAny_with is called when production any_with is exited.
func (s *BaseRedshiftParserListener) ExitAny_with(ctx *Any_withContext) {}

// EnterCreateconversionstmt is called when production createconversionstmt is entered.
func (s *BaseRedshiftParserListener) EnterCreateconversionstmt(ctx *CreateconversionstmtContext) {}

// ExitCreateconversionstmt is called when production createconversionstmt is exited.
func (s *BaseRedshiftParserListener) ExitCreateconversionstmt(ctx *CreateconversionstmtContext) {}

// EnterClusterstmt is called when production clusterstmt is entered.
func (s *BaseRedshiftParserListener) EnterClusterstmt(ctx *ClusterstmtContext) {}

// ExitClusterstmt is called when production clusterstmt is exited.
func (s *BaseRedshiftParserListener) ExitClusterstmt(ctx *ClusterstmtContext) {}

// EnterCluster_index_specification is called when production cluster_index_specification is entered.
func (s *BaseRedshiftParserListener) EnterCluster_index_specification(ctx *Cluster_index_specificationContext) {
}

// ExitCluster_index_specification is called when production cluster_index_specification is exited.
func (s *BaseRedshiftParserListener) ExitCluster_index_specification(ctx *Cluster_index_specificationContext) {
}

// EnterVacuumstmt is called when production vacuumstmt is entered.
func (s *BaseRedshiftParserListener) EnterVacuumstmt(ctx *VacuumstmtContext) {}

// ExitVacuumstmt is called when production vacuumstmt is exited.
func (s *BaseRedshiftParserListener) ExitVacuumstmt(ctx *VacuumstmtContext) {}

// EnterVacuum_option is called when production vacuum_option is entered.
func (s *BaseRedshiftParserListener) EnterVacuum_option(ctx *Vacuum_optionContext) {}

// ExitVacuum_option is called when production vacuum_option is exited.
func (s *BaseRedshiftParserListener) ExitVacuum_option(ctx *Vacuum_optionContext) {}

// EnterAnalyzestmt is called when production analyzestmt is entered.
func (s *BaseRedshiftParserListener) EnterAnalyzestmt(ctx *AnalyzestmtContext) {}

// ExitAnalyzestmt is called when production analyzestmt is exited.
func (s *BaseRedshiftParserListener) ExitAnalyzestmt(ctx *AnalyzestmtContext) {}

// EnterVac_analyze_option_list is called when production vac_analyze_option_list is entered.
func (s *BaseRedshiftParserListener) EnterVac_analyze_option_list(ctx *Vac_analyze_option_listContext) {
}

// ExitVac_analyze_option_list is called when production vac_analyze_option_list is exited.
func (s *BaseRedshiftParserListener) ExitVac_analyze_option_list(ctx *Vac_analyze_option_listContext) {
}

// EnterAnalyze_keyword is called when production analyze_keyword is entered.
func (s *BaseRedshiftParserListener) EnterAnalyze_keyword(ctx *Analyze_keywordContext) {}

// ExitAnalyze_keyword is called when production analyze_keyword is exited.
func (s *BaseRedshiftParserListener) ExitAnalyze_keyword(ctx *Analyze_keywordContext) {}

// EnterVac_analyze_option_elem is called when production vac_analyze_option_elem is entered.
func (s *BaseRedshiftParserListener) EnterVac_analyze_option_elem(ctx *Vac_analyze_option_elemContext) {
}

// ExitVac_analyze_option_elem is called when production vac_analyze_option_elem is exited.
func (s *BaseRedshiftParserListener) ExitVac_analyze_option_elem(ctx *Vac_analyze_option_elemContext) {
}

// EnterVac_analyze_option_name is called when production vac_analyze_option_name is entered.
func (s *BaseRedshiftParserListener) EnterVac_analyze_option_name(ctx *Vac_analyze_option_nameContext) {
}

// ExitVac_analyze_option_name is called when production vac_analyze_option_name is exited.
func (s *BaseRedshiftParserListener) ExitVac_analyze_option_name(ctx *Vac_analyze_option_nameContext) {
}

// EnterVac_analyze_option_arg is called when production vac_analyze_option_arg is entered.
func (s *BaseRedshiftParserListener) EnterVac_analyze_option_arg(ctx *Vac_analyze_option_argContext) {
}

// ExitVac_analyze_option_arg is called when production vac_analyze_option_arg is exited.
func (s *BaseRedshiftParserListener) ExitVac_analyze_option_arg(ctx *Vac_analyze_option_argContext) {}

// EnterOpt_analyze is called when production opt_analyze is entered.
func (s *BaseRedshiftParserListener) EnterOpt_analyze(ctx *Opt_analyzeContext) {}

// ExitOpt_analyze is called when production opt_analyze is exited.
func (s *BaseRedshiftParserListener) ExitOpt_analyze(ctx *Opt_analyzeContext) {}

// EnterOpt_verbose is called when production opt_verbose is entered.
func (s *BaseRedshiftParserListener) EnterOpt_verbose(ctx *Opt_verboseContext) {}

// ExitOpt_verbose is called when production opt_verbose is exited.
func (s *BaseRedshiftParserListener) ExitOpt_verbose(ctx *Opt_verboseContext) {}

// EnterOpt_full is called when production opt_full is entered.
func (s *BaseRedshiftParserListener) EnterOpt_full(ctx *Opt_fullContext) {}

// ExitOpt_full is called when production opt_full is exited.
func (s *BaseRedshiftParserListener) ExitOpt_full(ctx *Opt_fullContext) {}

// EnterOpt_freeze is called when production opt_freeze is entered.
func (s *BaseRedshiftParserListener) EnterOpt_freeze(ctx *Opt_freezeContext) {}

// ExitOpt_freeze is called when production opt_freeze is exited.
func (s *BaseRedshiftParserListener) ExitOpt_freeze(ctx *Opt_freezeContext) {}

// EnterOpt_name_list is called when production opt_name_list is entered.
func (s *BaseRedshiftParserListener) EnterOpt_name_list(ctx *Opt_name_listContext) {}

// ExitOpt_name_list is called when production opt_name_list is exited.
func (s *BaseRedshiftParserListener) ExitOpt_name_list(ctx *Opt_name_listContext) {}

// EnterVacuum_relation is called when production vacuum_relation is entered.
func (s *BaseRedshiftParserListener) EnterVacuum_relation(ctx *Vacuum_relationContext) {}

// ExitVacuum_relation is called when production vacuum_relation is exited.
func (s *BaseRedshiftParserListener) ExitVacuum_relation(ctx *Vacuum_relationContext) {}

// EnterVacuum_relation_list is called when production vacuum_relation_list is entered.
func (s *BaseRedshiftParserListener) EnterVacuum_relation_list(ctx *Vacuum_relation_listContext) {}

// ExitVacuum_relation_list is called when production vacuum_relation_list is exited.
func (s *BaseRedshiftParserListener) ExitVacuum_relation_list(ctx *Vacuum_relation_listContext) {}

// EnterOpt_vacuum_relation_list is called when production opt_vacuum_relation_list is entered.
func (s *BaseRedshiftParserListener) EnterOpt_vacuum_relation_list(ctx *Opt_vacuum_relation_listContext) {
}

// ExitOpt_vacuum_relation_list is called when production opt_vacuum_relation_list is exited.
func (s *BaseRedshiftParserListener) ExitOpt_vacuum_relation_list(ctx *Opt_vacuum_relation_listContext) {
}

// EnterExplainstmt is called when production explainstmt is entered.
func (s *BaseRedshiftParserListener) EnterExplainstmt(ctx *ExplainstmtContext) {}

// ExitExplainstmt is called when production explainstmt is exited.
func (s *BaseRedshiftParserListener) ExitExplainstmt(ctx *ExplainstmtContext) {}

// EnterExplainablestmt is called when production explainablestmt is entered.
func (s *BaseRedshiftParserListener) EnterExplainablestmt(ctx *ExplainablestmtContext) {}

// ExitExplainablestmt is called when production explainablestmt is exited.
func (s *BaseRedshiftParserListener) ExitExplainablestmt(ctx *ExplainablestmtContext) {}

// EnterExplain_option_list is called when production explain_option_list is entered.
func (s *BaseRedshiftParserListener) EnterExplain_option_list(ctx *Explain_option_listContext) {}

// ExitExplain_option_list is called when production explain_option_list is exited.
func (s *BaseRedshiftParserListener) ExitExplain_option_list(ctx *Explain_option_listContext) {}

// EnterExplain_option_elem is called when production explain_option_elem is entered.
func (s *BaseRedshiftParserListener) EnterExplain_option_elem(ctx *Explain_option_elemContext) {}

// ExitExplain_option_elem is called when production explain_option_elem is exited.
func (s *BaseRedshiftParserListener) ExitExplain_option_elem(ctx *Explain_option_elemContext) {}

// EnterExplain_option_name is called when production explain_option_name is entered.
func (s *BaseRedshiftParserListener) EnterExplain_option_name(ctx *Explain_option_nameContext) {}

// ExitExplain_option_name is called when production explain_option_name is exited.
func (s *BaseRedshiftParserListener) ExitExplain_option_name(ctx *Explain_option_nameContext) {}

// EnterExplain_option_arg is called when production explain_option_arg is entered.
func (s *BaseRedshiftParserListener) EnterExplain_option_arg(ctx *Explain_option_argContext) {}

// ExitExplain_option_arg is called when production explain_option_arg is exited.
func (s *BaseRedshiftParserListener) ExitExplain_option_arg(ctx *Explain_option_argContext) {}

// EnterPreparestmt is called when production preparestmt is entered.
func (s *BaseRedshiftParserListener) EnterPreparestmt(ctx *PreparestmtContext) {}

// ExitPreparestmt is called when production preparestmt is exited.
func (s *BaseRedshiftParserListener) ExitPreparestmt(ctx *PreparestmtContext) {}

// EnterPrep_type_clause is called when production prep_type_clause is entered.
func (s *BaseRedshiftParserListener) EnterPrep_type_clause(ctx *Prep_type_clauseContext) {}

// ExitPrep_type_clause is called when production prep_type_clause is exited.
func (s *BaseRedshiftParserListener) ExitPrep_type_clause(ctx *Prep_type_clauseContext) {}

// EnterPreparablestmt is called when production preparablestmt is entered.
func (s *BaseRedshiftParserListener) EnterPreparablestmt(ctx *PreparablestmtContext) {}

// ExitPreparablestmt is called when production preparablestmt is exited.
func (s *BaseRedshiftParserListener) ExitPreparablestmt(ctx *PreparablestmtContext) {}

// EnterExecutestmt is called when production executestmt is entered.
func (s *BaseRedshiftParserListener) EnterExecutestmt(ctx *ExecutestmtContext) {}

// ExitExecutestmt is called when production executestmt is exited.
func (s *BaseRedshiftParserListener) ExitExecutestmt(ctx *ExecutestmtContext) {}

// EnterExecute_param_clause is called when production execute_param_clause is entered.
func (s *BaseRedshiftParserListener) EnterExecute_param_clause(ctx *Execute_param_clauseContext) {}

// ExitExecute_param_clause is called when production execute_param_clause is exited.
func (s *BaseRedshiftParserListener) ExitExecute_param_clause(ctx *Execute_param_clauseContext) {}

// EnterDeallocatestmt is called when production deallocatestmt is entered.
func (s *BaseRedshiftParserListener) EnterDeallocatestmt(ctx *DeallocatestmtContext) {}

// ExitDeallocatestmt is called when production deallocatestmt is exited.
func (s *BaseRedshiftParserListener) ExitDeallocatestmt(ctx *DeallocatestmtContext) {}

// EnterInsertstmt is called when production insertstmt is entered.
func (s *BaseRedshiftParserListener) EnterInsertstmt(ctx *InsertstmtContext) {}

// ExitInsertstmt is called when production insertstmt is exited.
func (s *BaseRedshiftParserListener) ExitInsertstmt(ctx *InsertstmtContext) {}

// EnterInsert_target is called when production insert_target is entered.
func (s *BaseRedshiftParserListener) EnterInsert_target(ctx *Insert_targetContext) {}

// ExitInsert_target is called when production insert_target is exited.
func (s *BaseRedshiftParserListener) ExitInsert_target(ctx *Insert_targetContext) {}

// EnterInsert_rest is called when production insert_rest is entered.
func (s *BaseRedshiftParserListener) EnterInsert_rest(ctx *Insert_restContext) {}

// ExitInsert_rest is called when production insert_rest is exited.
func (s *BaseRedshiftParserListener) ExitInsert_rest(ctx *Insert_restContext) {}

// EnterOverride_kind is called when production override_kind is entered.
func (s *BaseRedshiftParserListener) EnterOverride_kind(ctx *Override_kindContext) {}

// ExitOverride_kind is called when production override_kind is exited.
func (s *BaseRedshiftParserListener) ExitOverride_kind(ctx *Override_kindContext) {}

// EnterInsert_column_list is called when production insert_column_list is entered.
func (s *BaseRedshiftParserListener) EnterInsert_column_list(ctx *Insert_column_listContext) {}

// ExitInsert_column_list is called when production insert_column_list is exited.
func (s *BaseRedshiftParserListener) ExitInsert_column_list(ctx *Insert_column_listContext) {}

// EnterInsert_column_item is called when production insert_column_item is entered.
func (s *BaseRedshiftParserListener) EnterInsert_column_item(ctx *Insert_column_itemContext) {}

// ExitInsert_column_item is called when production insert_column_item is exited.
func (s *BaseRedshiftParserListener) ExitInsert_column_item(ctx *Insert_column_itemContext) {}

// EnterOpt_on_conflict is called when production opt_on_conflict is entered.
func (s *BaseRedshiftParserListener) EnterOpt_on_conflict(ctx *Opt_on_conflictContext) {}

// ExitOpt_on_conflict is called when production opt_on_conflict is exited.
func (s *BaseRedshiftParserListener) ExitOpt_on_conflict(ctx *Opt_on_conflictContext) {}

// EnterOpt_conf_expr is called when production opt_conf_expr is entered.
func (s *BaseRedshiftParserListener) EnterOpt_conf_expr(ctx *Opt_conf_exprContext) {}

// ExitOpt_conf_expr is called when production opt_conf_expr is exited.
func (s *BaseRedshiftParserListener) ExitOpt_conf_expr(ctx *Opt_conf_exprContext) {}

// EnterReturning_clause is called when production returning_clause is entered.
func (s *BaseRedshiftParserListener) EnterReturning_clause(ctx *Returning_clauseContext) {}

// ExitReturning_clause is called when production returning_clause is exited.
func (s *BaseRedshiftParserListener) ExitReturning_clause(ctx *Returning_clauseContext) {}

// EnterMergestmt is called when production mergestmt is entered.
func (s *BaseRedshiftParserListener) EnterMergestmt(ctx *MergestmtContext) {}

// ExitMergestmt is called when production mergestmt is exited.
func (s *BaseRedshiftParserListener) ExitMergestmt(ctx *MergestmtContext) {}

// EnterMerge_when_clause is called when production merge_when_clause is entered.
func (s *BaseRedshiftParserListener) EnterMerge_when_clause(ctx *Merge_when_clauseContext) {}

// ExitMerge_when_clause is called when production merge_when_clause is exited.
func (s *BaseRedshiftParserListener) ExitMerge_when_clause(ctx *Merge_when_clauseContext) {}

// EnterMerge_insert_clause is called when production merge_insert_clause is entered.
func (s *BaseRedshiftParserListener) EnterMerge_insert_clause(ctx *Merge_insert_clauseContext) {}

// ExitMerge_insert_clause is called when production merge_insert_clause is exited.
func (s *BaseRedshiftParserListener) ExitMerge_insert_clause(ctx *Merge_insert_clauseContext) {}

// EnterMerge_update_clause is called when production merge_update_clause is entered.
func (s *BaseRedshiftParserListener) EnterMerge_update_clause(ctx *Merge_update_clauseContext) {}

// ExitMerge_update_clause is called when production merge_update_clause is exited.
func (s *BaseRedshiftParserListener) ExitMerge_update_clause(ctx *Merge_update_clauseContext) {}

// EnterMerge_delete_clause is called when production merge_delete_clause is entered.
func (s *BaseRedshiftParserListener) EnterMerge_delete_clause(ctx *Merge_delete_clauseContext) {}

// ExitMerge_delete_clause is called when production merge_delete_clause is exited.
func (s *BaseRedshiftParserListener) ExitMerge_delete_clause(ctx *Merge_delete_clauseContext) {}

// EnterDeletestmt is called when production deletestmt is entered.
func (s *BaseRedshiftParserListener) EnterDeletestmt(ctx *DeletestmtContext) {}

// ExitDeletestmt is called when production deletestmt is exited.
func (s *BaseRedshiftParserListener) ExitDeletestmt(ctx *DeletestmtContext) {}

// EnterUsing_clause is called when production using_clause is entered.
func (s *BaseRedshiftParserListener) EnterUsing_clause(ctx *Using_clauseContext) {}

// ExitUsing_clause is called when production using_clause is exited.
func (s *BaseRedshiftParserListener) ExitUsing_clause(ctx *Using_clauseContext) {}

// EnterLockstmt is called when production lockstmt is entered.
func (s *BaseRedshiftParserListener) EnterLockstmt(ctx *LockstmtContext) {}

// ExitLockstmt is called when production lockstmt is exited.
func (s *BaseRedshiftParserListener) ExitLockstmt(ctx *LockstmtContext) {}

// EnterOpt_lock is called when production opt_lock is entered.
func (s *BaseRedshiftParserListener) EnterOpt_lock(ctx *Opt_lockContext) {}

// ExitOpt_lock is called when production opt_lock is exited.
func (s *BaseRedshiftParserListener) ExitOpt_lock(ctx *Opt_lockContext) {}

// EnterLock_type is called when production lock_type is entered.
func (s *BaseRedshiftParserListener) EnterLock_type(ctx *Lock_typeContext) {}

// ExitLock_type is called when production lock_type is exited.
func (s *BaseRedshiftParserListener) ExitLock_type(ctx *Lock_typeContext) {}

// EnterOpt_nowait is called when production opt_nowait is entered.
func (s *BaseRedshiftParserListener) EnterOpt_nowait(ctx *Opt_nowaitContext) {}

// ExitOpt_nowait is called when production opt_nowait is exited.
func (s *BaseRedshiftParserListener) ExitOpt_nowait(ctx *Opt_nowaitContext) {}

// EnterOpt_nowait_or_skip is called when production opt_nowait_or_skip is entered.
func (s *BaseRedshiftParserListener) EnterOpt_nowait_or_skip(ctx *Opt_nowait_or_skipContext) {}

// ExitOpt_nowait_or_skip is called when production opt_nowait_or_skip is exited.
func (s *BaseRedshiftParserListener) ExitOpt_nowait_or_skip(ctx *Opt_nowait_or_skipContext) {}

// EnterUpdatestmt is called when production updatestmt is entered.
func (s *BaseRedshiftParserListener) EnterUpdatestmt(ctx *UpdatestmtContext) {}

// ExitUpdatestmt is called when production updatestmt is exited.
func (s *BaseRedshiftParserListener) ExitUpdatestmt(ctx *UpdatestmtContext) {}

// EnterSet_clause_list is called when production set_clause_list is entered.
func (s *BaseRedshiftParserListener) EnterSet_clause_list(ctx *Set_clause_listContext) {}

// ExitSet_clause_list is called when production set_clause_list is exited.
func (s *BaseRedshiftParserListener) ExitSet_clause_list(ctx *Set_clause_listContext) {}

// EnterSet_clause is called when production set_clause is entered.
func (s *BaseRedshiftParserListener) EnterSet_clause(ctx *Set_clauseContext) {}

// ExitSet_clause is called when production set_clause is exited.
func (s *BaseRedshiftParserListener) ExitSet_clause(ctx *Set_clauseContext) {}

// EnterSet_target is called when production set_target is entered.
func (s *BaseRedshiftParserListener) EnterSet_target(ctx *Set_targetContext) {}

// ExitSet_target is called when production set_target is exited.
func (s *BaseRedshiftParserListener) ExitSet_target(ctx *Set_targetContext) {}

// EnterSet_target_list is called when production set_target_list is entered.
func (s *BaseRedshiftParserListener) EnterSet_target_list(ctx *Set_target_listContext) {}

// ExitSet_target_list is called when production set_target_list is exited.
func (s *BaseRedshiftParserListener) ExitSet_target_list(ctx *Set_target_listContext) {}

// EnterDeclarecursorstmt is called when production declarecursorstmt is entered.
func (s *BaseRedshiftParserListener) EnterDeclarecursorstmt(ctx *DeclarecursorstmtContext) {}

// ExitDeclarecursorstmt is called when production declarecursorstmt is exited.
func (s *BaseRedshiftParserListener) ExitDeclarecursorstmt(ctx *DeclarecursorstmtContext) {}

// EnterCursor_name is called when production cursor_name is entered.
func (s *BaseRedshiftParserListener) EnterCursor_name(ctx *Cursor_nameContext) {}

// ExitCursor_name is called when production cursor_name is exited.
func (s *BaseRedshiftParserListener) ExitCursor_name(ctx *Cursor_nameContext) {}

// EnterCursor_options is called when production cursor_options is entered.
func (s *BaseRedshiftParserListener) EnterCursor_options(ctx *Cursor_optionsContext) {}

// ExitCursor_options is called when production cursor_options is exited.
func (s *BaseRedshiftParserListener) ExitCursor_options(ctx *Cursor_optionsContext) {}

// EnterOpt_hold is called when production opt_hold is entered.
func (s *BaseRedshiftParserListener) EnterOpt_hold(ctx *Opt_holdContext) {}

// ExitOpt_hold is called when production opt_hold is exited.
func (s *BaseRedshiftParserListener) ExitOpt_hold(ctx *Opt_holdContext) {}

// EnterSelectstmt is called when production selectstmt is entered.
func (s *BaseRedshiftParserListener) EnterSelectstmt(ctx *SelectstmtContext) {}

// ExitSelectstmt is called when production selectstmt is exited.
func (s *BaseRedshiftParserListener) ExitSelectstmt(ctx *SelectstmtContext) {}

// EnterSelect_with_parens is called when production select_with_parens is entered.
func (s *BaseRedshiftParserListener) EnterSelect_with_parens(ctx *Select_with_parensContext) {}

// ExitSelect_with_parens is called when production select_with_parens is exited.
func (s *BaseRedshiftParserListener) ExitSelect_with_parens(ctx *Select_with_parensContext) {}

// EnterSelect_no_parens is called when production select_no_parens is entered.
func (s *BaseRedshiftParserListener) EnterSelect_no_parens(ctx *Select_no_parensContext) {}

// ExitSelect_no_parens is called when production select_no_parens is exited.
func (s *BaseRedshiftParserListener) ExitSelect_no_parens(ctx *Select_no_parensContext) {}

// EnterSelect_clause is called when production select_clause is entered.
func (s *BaseRedshiftParserListener) EnterSelect_clause(ctx *Select_clauseContext) {}

// ExitSelect_clause is called when production select_clause is exited.
func (s *BaseRedshiftParserListener) ExitSelect_clause(ctx *Select_clauseContext) {}

// EnterSimple_select_intersect is called when production simple_select_intersect is entered.
func (s *BaseRedshiftParserListener) EnterSimple_select_intersect(ctx *Simple_select_intersectContext) {
}

// ExitSimple_select_intersect is called when production simple_select_intersect is exited.
func (s *BaseRedshiftParserListener) ExitSimple_select_intersect(ctx *Simple_select_intersectContext) {
}

// EnterSimple_select_pramary is called when production simple_select_pramary is entered.
func (s *BaseRedshiftParserListener) EnterSimple_select_pramary(ctx *Simple_select_pramaryContext) {}

// ExitSimple_select_pramary is called when production simple_select_pramary is exited.
func (s *BaseRedshiftParserListener) ExitSimple_select_pramary(ctx *Simple_select_pramaryContext) {}

// EnterExclude_clause is called when production exclude_clause is entered.
func (s *BaseRedshiftParserListener) EnterExclude_clause(ctx *Exclude_clauseContext) {}

// ExitExclude_clause is called when production exclude_clause is exited.
func (s *BaseRedshiftParserListener) ExitExclude_clause(ctx *Exclude_clauseContext) {}

// EnterQualify_clause is called when production qualify_clause is entered.
func (s *BaseRedshiftParserListener) EnterQualify_clause(ctx *Qualify_clauseContext) {}

// ExitQualify_clause is called when production qualify_clause is exited.
func (s *BaseRedshiftParserListener) ExitQualify_clause(ctx *Qualify_clauseContext) {}

// EnterStart_with_clause is called when production start_with_clause is entered.
func (s *BaseRedshiftParserListener) EnterStart_with_clause(ctx *Start_with_clauseContext) {}

// ExitStart_with_clause is called when production start_with_clause is exited.
func (s *BaseRedshiftParserListener) ExitStart_with_clause(ctx *Start_with_clauseContext) {}

// EnterWith_clause is called when production with_clause is entered.
func (s *BaseRedshiftParserListener) EnterWith_clause(ctx *With_clauseContext) {}

// ExitWith_clause is called when production with_clause is exited.
func (s *BaseRedshiftParserListener) ExitWith_clause(ctx *With_clauseContext) {}

// EnterCte_list is called when production cte_list is entered.
func (s *BaseRedshiftParserListener) EnterCte_list(ctx *Cte_listContext) {}

// ExitCte_list is called when production cte_list is exited.
func (s *BaseRedshiftParserListener) ExitCte_list(ctx *Cte_listContext) {}

// EnterCommon_table_expr is called when production common_table_expr is entered.
func (s *BaseRedshiftParserListener) EnterCommon_table_expr(ctx *Common_table_exprContext) {}

// ExitCommon_table_expr is called when production common_table_expr is exited.
func (s *BaseRedshiftParserListener) ExitCommon_table_expr(ctx *Common_table_exprContext) {}

// EnterOpt_materialized is called when production opt_materialized is entered.
func (s *BaseRedshiftParserListener) EnterOpt_materialized(ctx *Opt_materializedContext) {}

// ExitOpt_materialized is called when production opt_materialized is exited.
func (s *BaseRedshiftParserListener) ExitOpt_materialized(ctx *Opt_materializedContext) {}

// EnterOpt_with_clause is called when production opt_with_clause is entered.
func (s *BaseRedshiftParserListener) EnterOpt_with_clause(ctx *Opt_with_clauseContext) {}

// ExitOpt_with_clause is called when production opt_with_clause is exited.
func (s *BaseRedshiftParserListener) ExitOpt_with_clause(ctx *Opt_with_clauseContext) {}

// EnterInto_clause is called when production into_clause is entered.
func (s *BaseRedshiftParserListener) EnterInto_clause(ctx *Into_clauseContext) {}

// ExitInto_clause is called when production into_clause is exited.
func (s *BaseRedshiftParserListener) ExitInto_clause(ctx *Into_clauseContext) {}

// EnterOpt_top_clause is called when production opt_top_clause is entered.
func (s *BaseRedshiftParserListener) EnterOpt_top_clause(ctx *Opt_top_clauseContext) {}

// ExitOpt_top_clause is called when production opt_top_clause is exited.
func (s *BaseRedshiftParserListener) ExitOpt_top_clause(ctx *Opt_top_clauseContext) {}

// EnterOpt_strict is called when production opt_strict is entered.
func (s *BaseRedshiftParserListener) EnterOpt_strict(ctx *Opt_strictContext) {}

// ExitOpt_strict is called when production opt_strict is exited.
func (s *BaseRedshiftParserListener) ExitOpt_strict(ctx *Opt_strictContext) {}

// EnterOpttempTableName is called when production opttempTableName is entered.
func (s *BaseRedshiftParserListener) EnterOpttempTableName(ctx *OpttempTableNameContext) {}

// ExitOpttempTableName is called when production opttempTableName is exited.
func (s *BaseRedshiftParserListener) ExitOpttempTableName(ctx *OpttempTableNameContext) {}

// EnterOpt_table is called when production opt_table is entered.
func (s *BaseRedshiftParserListener) EnterOpt_table(ctx *Opt_tableContext) {}

// ExitOpt_table is called when production opt_table is exited.
func (s *BaseRedshiftParserListener) ExitOpt_table(ctx *Opt_tableContext) {}

// EnterAll_or_distinct is called when production all_or_distinct is entered.
func (s *BaseRedshiftParserListener) EnterAll_or_distinct(ctx *All_or_distinctContext) {}

// ExitAll_or_distinct is called when production all_or_distinct is exited.
func (s *BaseRedshiftParserListener) ExitAll_or_distinct(ctx *All_or_distinctContext) {}

// EnterDistinct_clause is called when production distinct_clause is entered.
func (s *BaseRedshiftParserListener) EnterDistinct_clause(ctx *Distinct_clauseContext) {}

// ExitDistinct_clause is called when production distinct_clause is exited.
func (s *BaseRedshiftParserListener) ExitDistinct_clause(ctx *Distinct_clauseContext) {}

// EnterOpt_all_clause is called when production opt_all_clause is entered.
func (s *BaseRedshiftParserListener) EnterOpt_all_clause(ctx *Opt_all_clauseContext) {}

// ExitOpt_all_clause is called when production opt_all_clause is exited.
func (s *BaseRedshiftParserListener) ExitOpt_all_clause(ctx *Opt_all_clauseContext) {}

// EnterOpt_sort_clause is called when production opt_sort_clause is entered.
func (s *BaseRedshiftParserListener) EnterOpt_sort_clause(ctx *Opt_sort_clauseContext) {}

// ExitOpt_sort_clause is called when production opt_sort_clause is exited.
func (s *BaseRedshiftParserListener) ExitOpt_sort_clause(ctx *Opt_sort_clauseContext) {}

// EnterSort_clause is called when production sort_clause is entered.
func (s *BaseRedshiftParserListener) EnterSort_clause(ctx *Sort_clauseContext) {}

// ExitSort_clause is called when production sort_clause is exited.
func (s *BaseRedshiftParserListener) ExitSort_clause(ctx *Sort_clauseContext) {}

// EnterSortby_list is called when production sortby_list is entered.
func (s *BaseRedshiftParserListener) EnterSortby_list(ctx *Sortby_listContext) {}

// ExitSortby_list is called when production sortby_list is exited.
func (s *BaseRedshiftParserListener) ExitSortby_list(ctx *Sortby_listContext) {}

// EnterSortby is called when production sortby is entered.
func (s *BaseRedshiftParserListener) EnterSortby(ctx *SortbyContext) {}

// ExitSortby is called when production sortby is exited.
func (s *BaseRedshiftParserListener) ExitSortby(ctx *SortbyContext) {}

// EnterSelect_limit is called when production select_limit is entered.
func (s *BaseRedshiftParserListener) EnterSelect_limit(ctx *Select_limitContext) {}

// ExitSelect_limit is called when production select_limit is exited.
func (s *BaseRedshiftParserListener) ExitSelect_limit(ctx *Select_limitContext) {}

// EnterOpt_select_limit is called when production opt_select_limit is entered.
func (s *BaseRedshiftParserListener) EnterOpt_select_limit(ctx *Opt_select_limitContext) {}

// ExitOpt_select_limit is called when production opt_select_limit is exited.
func (s *BaseRedshiftParserListener) ExitOpt_select_limit(ctx *Opt_select_limitContext) {}

// EnterLimit_clause is called when production limit_clause is entered.
func (s *BaseRedshiftParserListener) EnterLimit_clause(ctx *Limit_clauseContext) {}

// ExitLimit_clause is called when production limit_clause is exited.
func (s *BaseRedshiftParserListener) ExitLimit_clause(ctx *Limit_clauseContext) {}

// EnterOffset_clause is called when production offset_clause is entered.
func (s *BaseRedshiftParserListener) EnterOffset_clause(ctx *Offset_clauseContext) {}

// ExitOffset_clause is called when production offset_clause is exited.
func (s *BaseRedshiftParserListener) ExitOffset_clause(ctx *Offset_clauseContext) {}

// EnterSelect_limit_value is called when production select_limit_value is entered.
func (s *BaseRedshiftParserListener) EnterSelect_limit_value(ctx *Select_limit_valueContext) {}

// ExitSelect_limit_value is called when production select_limit_value is exited.
func (s *BaseRedshiftParserListener) ExitSelect_limit_value(ctx *Select_limit_valueContext) {}

// EnterSelect_offset_value is called when production select_offset_value is entered.
func (s *BaseRedshiftParserListener) EnterSelect_offset_value(ctx *Select_offset_valueContext) {}

// ExitSelect_offset_value is called when production select_offset_value is exited.
func (s *BaseRedshiftParserListener) ExitSelect_offset_value(ctx *Select_offset_valueContext) {}

// EnterSelect_fetch_first_value is called when production select_fetch_first_value is entered.
func (s *BaseRedshiftParserListener) EnterSelect_fetch_first_value(ctx *Select_fetch_first_valueContext) {
}

// ExitSelect_fetch_first_value is called when production select_fetch_first_value is exited.
func (s *BaseRedshiftParserListener) ExitSelect_fetch_first_value(ctx *Select_fetch_first_valueContext) {
}

// EnterI_or_f_const is called when production i_or_f_const is entered.
func (s *BaseRedshiftParserListener) EnterI_or_f_const(ctx *I_or_f_constContext) {}

// ExitI_or_f_const is called when production i_or_f_const is exited.
func (s *BaseRedshiftParserListener) ExitI_or_f_const(ctx *I_or_f_constContext) {}

// EnterRow_or_rows is called when production row_or_rows is entered.
func (s *BaseRedshiftParserListener) EnterRow_or_rows(ctx *Row_or_rowsContext) {}

// ExitRow_or_rows is called when production row_or_rows is exited.
func (s *BaseRedshiftParserListener) ExitRow_or_rows(ctx *Row_or_rowsContext) {}

// EnterFirst_or_next is called when production first_or_next is entered.
func (s *BaseRedshiftParserListener) EnterFirst_or_next(ctx *First_or_nextContext) {}

// ExitFirst_or_next is called when production first_or_next is exited.
func (s *BaseRedshiftParserListener) ExitFirst_or_next(ctx *First_or_nextContext) {}

// EnterGroup_clause is called when production group_clause is entered.
func (s *BaseRedshiftParserListener) EnterGroup_clause(ctx *Group_clauseContext) {}

// ExitGroup_clause is called when production group_clause is exited.
func (s *BaseRedshiftParserListener) ExitGroup_clause(ctx *Group_clauseContext) {}

// EnterGroup_by_list is called when production group_by_list is entered.
func (s *BaseRedshiftParserListener) EnterGroup_by_list(ctx *Group_by_listContext) {}

// ExitGroup_by_list is called when production group_by_list is exited.
func (s *BaseRedshiftParserListener) ExitGroup_by_list(ctx *Group_by_listContext) {}

// EnterGroup_by_item is called when production group_by_item is entered.
func (s *BaseRedshiftParserListener) EnterGroup_by_item(ctx *Group_by_itemContext) {}

// ExitGroup_by_item is called when production group_by_item is exited.
func (s *BaseRedshiftParserListener) ExitGroup_by_item(ctx *Group_by_itemContext) {}

// EnterEmpty_grouping_set is called when production empty_grouping_set is entered.
func (s *BaseRedshiftParserListener) EnterEmpty_grouping_set(ctx *Empty_grouping_setContext) {}

// ExitEmpty_grouping_set is called when production empty_grouping_set is exited.
func (s *BaseRedshiftParserListener) ExitEmpty_grouping_set(ctx *Empty_grouping_setContext) {}

// EnterRollup_clause is called when production rollup_clause is entered.
func (s *BaseRedshiftParserListener) EnterRollup_clause(ctx *Rollup_clauseContext) {}

// ExitRollup_clause is called when production rollup_clause is exited.
func (s *BaseRedshiftParserListener) ExitRollup_clause(ctx *Rollup_clauseContext) {}

// EnterCube_clause is called when production cube_clause is entered.
func (s *BaseRedshiftParserListener) EnterCube_clause(ctx *Cube_clauseContext) {}

// ExitCube_clause is called when production cube_clause is exited.
func (s *BaseRedshiftParserListener) ExitCube_clause(ctx *Cube_clauseContext) {}

// EnterGrouping_sets_clause is called when production grouping_sets_clause is entered.
func (s *BaseRedshiftParserListener) EnterGrouping_sets_clause(ctx *Grouping_sets_clauseContext) {}

// ExitGrouping_sets_clause is called when production grouping_sets_clause is exited.
func (s *BaseRedshiftParserListener) ExitGrouping_sets_clause(ctx *Grouping_sets_clauseContext) {}

// EnterHaving_clause is called when production having_clause is entered.
func (s *BaseRedshiftParserListener) EnterHaving_clause(ctx *Having_clauseContext) {}

// ExitHaving_clause is called when production having_clause is exited.
func (s *BaseRedshiftParserListener) ExitHaving_clause(ctx *Having_clauseContext) {}

// EnterFor_locking_clause is called when production for_locking_clause is entered.
func (s *BaseRedshiftParserListener) EnterFor_locking_clause(ctx *For_locking_clauseContext) {}

// ExitFor_locking_clause is called when production for_locking_clause is exited.
func (s *BaseRedshiftParserListener) ExitFor_locking_clause(ctx *For_locking_clauseContext) {}

// EnterOpt_for_locking_clause is called when production opt_for_locking_clause is entered.
func (s *BaseRedshiftParserListener) EnterOpt_for_locking_clause(ctx *Opt_for_locking_clauseContext) {
}

// ExitOpt_for_locking_clause is called when production opt_for_locking_clause is exited.
func (s *BaseRedshiftParserListener) ExitOpt_for_locking_clause(ctx *Opt_for_locking_clauseContext) {}

// EnterFor_locking_items is called when production for_locking_items is entered.
func (s *BaseRedshiftParserListener) EnterFor_locking_items(ctx *For_locking_itemsContext) {}

// ExitFor_locking_items is called when production for_locking_items is exited.
func (s *BaseRedshiftParserListener) ExitFor_locking_items(ctx *For_locking_itemsContext) {}

// EnterFor_locking_item is called when production for_locking_item is entered.
func (s *BaseRedshiftParserListener) EnterFor_locking_item(ctx *For_locking_itemContext) {}

// ExitFor_locking_item is called when production for_locking_item is exited.
func (s *BaseRedshiftParserListener) ExitFor_locking_item(ctx *For_locking_itemContext) {}

// EnterFor_locking_strength is called when production for_locking_strength is entered.
func (s *BaseRedshiftParserListener) EnterFor_locking_strength(ctx *For_locking_strengthContext) {}

// ExitFor_locking_strength is called when production for_locking_strength is exited.
func (s *BaseRedshiftParserListener) ExitFor_locking_strength(ctx *For_locking_strengthContext) {}

// EnterLocked_rels_list is called when production locked_rels_list is entered.
func (s *BaseRedshiftParserListener) EnterLocked_rels_list(ctx *Locked_rels_listContext) {}

// ExitLocked_rels_list is called when production locked_rels_list is exited.
func (s *BaseRedshiftParserListener) ExitLocked_rels_list(ctx *Locked_rels_listContext) {}

// EnterValues_clause is called when production values_clause is entered.
func (s *BaseRedshiftParserListener) EnterValues_clause(ctx *Values_clauseContext) {}

// ExitValues_clause is called when production values_clause is exited.
func (s *BaseRedshiftParserListener) ExitValues_clause(ctx *Values_clauseContext) {}

// EnterFrom_clause is called when production from_clause is entered.
func (s *BaseRedshiftParserListener) EnterFrom_clause(ctx *From_clauseContext) {}

// ExitFrom_clause is called when production from_clause is exited.
func (s *BaseRedshiftParserListener) ExitFrom_clause(ctx *From_clauseContext) {}

// EnterFrom_list is called when production from_list is entered.
func (s *BaseRedshiftParserListener) EnterFrom_list(ctx *From_listContext) {}

// ExitFrom_list is called when production from_list is exited.
func (s *BaseRedshiftParserListener) ExitFrom_list(ctx *From_listContext) {}

// EnterTable_ref is called when production table_ref is entered.
func (s *BaseRedshiftParserListener) EnterTable_ref(ctx *Table_refContext) {}

// ExitTable_ref is called when production table_ref is exited.
func (s *BaseRedshiftParserListener) ExitTable_ref(ctx *Table_refContext) {}

// EnterJoined_table is called when production joined_table is entered.
func (s *BaseRedshiftParserListener) EnterJoined_table(ctx *Joined_tableContext) {}

// ExitJoined_table is called when production joined_table is exited.
func (s *BaseRedshiftParserListener) ExitJoined_table(ctx *Joined_tableContext) {}

// EnterAlias_clause is called when production alias_clause is entered.
func (s *BaseRedshiftParserListener) EnterAlias_clause(ctx *Alias_clauseContext) {}

// ExitAlias_clause is called when production alias_clause is exited.
func (s *BaseRedshiftParserListener) ExitAlias_clause(ctx *Alias_clauseContext) {}

// EnterOpt_alias_clause is called when production opt_alias_clause is entered.
func (s *BaseRedshiftParserListener) EnterOpt_alias_clause(ctx *Opt_alias_clauseContext) {}

// ExitOpt_alias_clause is called when production opt_alias_clause is exited.
func (s *BaseRedshiftParserListener) ExitOpt_alias_clause(ctx *Opt_alias_clauseContext) {}

// EnterTable_alias_clause is called when production table_alias_clause is entered.
func (s *BaseRedshiftParserListener) EnterTable_alias_clause(ctx *Table_alias_clauseContext) {}

// ExitTable_alias_clause is called when production table_alias_clause is exited.
func (s *BaseRedshiftParserListener) ExitTable_alias_clause(ctx *Table_alias_clauseContext) {}

// EnterFunc_alias_clause is called when production func_alias_clause is entered.
func (s *BaseRedshiftParserListener) EnterFunc_alias_clause(ctx *Func_alias_clauseContext) {}

// ExitFunc_alias_clause is called when production func_alias_clause is exited.
func (s *BaseRedshiftParserListener) ExitFunc_alias_clause(ctx *Func_alias_clauseContext) {}

// EnterJoin_type is called when production join_type is entered.
func (s *BaseRedshiftParserListener) EnterJoin_type(ctx *Join_typeContext) {}

// ExitJoin_type is called when production join_type is exited.
func (s *BaseRedshiftParserListener) ExitJoin_type(ctx *Join_typeContext) {}

// EnterJoin_qual is called when production join_qual is entered.
func (s *BaseRedshiftParserListener) EnterJoin_qual(ctx *Join_qualContext) {}

// ExitJoin_qual is called when production join_qual is exited.
func (s *BaseRedshiftParserListener) ExitJoin_qual(ctx *Join_qualContext) {}

// EnterRelation_expr is called when production relation_expr is entered.
func (s *BaseRedshiftParserListener) EnterRelation_expr(ctx *Relation_exprContext) {}

// ExitRelation_expr is called when production relation_expr is exited.
func (s *BaseRedshiftParserListener) ExitRelation_expr(ctx *Relation_exprContext) {}

// EnterRelation_expr_list is called when production relation_expr_list is entered.
func (s *BaseRedshiftParserListener) EnterRelation_expr_list(ctx *Relation_expr_listContext) {}

// ExitRelation_expr_list is called when production relation_expr_list is exited.
func (s *BaseRedshiftParserListener) ExitRelation_expr_list(ctx *Relation_expr_listContext) {}

// EnterRelation_expr_opt_alias is called when production relation_expr_opt_alias is entered.
func (s *BaseRedshiftParserListener) EnterRelation_expr_opt_alias(ctx *Relation_expr_opt_aliasContext) {
}

// ExitRelation_expr_opt_alias is called when production relation_expr_opt_alias is exited.
func (s *BaseRedshiftParserListener) ExitRelation_expr_opt_alias(ctx *Relation_expr_opt_aliasContext) {
}

// EnterTablesample_clause is called when production tablesample_clause is entered.
func (s *BaseRedshiftParserListener) EnterTablesample_clause(ctx *Tablesample_clauseContext) {}

// ExitTablesample_clause is called when production tablesample_clause is exited.
func (s *BaseRedshiftParserListener) ExitTablesample_clause(ctx *Tablesample_clauseContext) {}

// EnterOpt_repeatable_clause is called when production opt_repeatable_clause is entered.
func (s *BaseRedshiftParserListener) EnterOpt_repeatable_clause(ctx *Opt_repeatable_clauseContext) {}

// ExitOpt_repeatable_clause is called when production opt_repeatable_clause is exited.
func (s *BaseRedshiftParserListener) ExitOpt_repeatable_clause(ctx *Opt_repeatable_clauseContext) {}

// EnterFunc_table is called when production func_table is entered.
func (s *BaseRedshiftParserListener) EnterFunc_table(ctx *Func_tableContext) {}

// ExitFunc_table is called when production func_table is exited.
func (s *BaseRedshiftParserListener) ExitFunc_table(ctx *Func_tableContext) {}

// EnterRowsfrom_item is called when production rowsfrom_item is entered.
func (s *BaseRedshiftParserListener) EnterRowsfrom_item(ctx *Rowsfrom_itemContext) {}

// ExitRowsfrom_item is called when production rowsfrom_item is exited.
func (s *BaseRedshiftParserListener) ExitRowsfrom_item(ctx *Rowsfrom_itemContext) {}

// EnterRowsfrom_list is called when production rowsfrom_list is entered.
func (s *BaseRedshiftParserListener) EnterRowsfrom_list(ctx *Rowsfrom_listContext) {}

// ExitRowsfrom_list is called when production rowsfrom_list is exited.
func (s *BaseRedshiftParserListener) ExitRowsfrom_list(ctx *Rowsfrom_listContext) {}

// EnterOpt_col_def_list is called when production opt_col_def_list is entered.
func (s *BaseRedshiftParserListener) EnterOpt_col_def_list(ctx *Opt_col_def_listContext) {}

// ExitOpt_col_def_list is called when production opt_col_def_list is exited.
func (s *BaseRedshiftParserListener) ExitOpt_col_def_list(ctx *Opt_col_def_listContext) {}

// EnterOpt_ordinality is called when production opt_ordinality is entered.
func (s *BaseRedshiftParserListener) EnterOpt_ordinality(ctx *Opt_ordinalityContext) {}

// ExitOpt_ordinality is called when production opt_ordinality is exited.
func (s *BaseRedshiftParserListener) ExitOpt_ordinality(ctx *Opt_ordinalityContext) {}

// EnterWhere_clause is called when production where_clause is entered.
func (s *BaseRedshiftParserListener) EnterWhere_clause(ctx *Where_clauseContext) {}

// ExitWhere_clause is called when production where_clause is exited.
func (s *BaseRedshiftParserListener) ExitWhere_clause(ctx *Where_clauseContext) {}

// EnterWhere_or_current_clause is called when production where_or_current_clause is entered.
func (s *BaseRedshiftParserListener) EnterWhere_or_current_clause(ctx *Where_or_current_clauseContext) {
}

// ExitWhere_or_current_clause is called when production where_or_current_clause is exited.
func (s *BaseRedshiftParserListener) ExitWhere_or_current_clause(ctx *Where_or_current_clauseContext) {
}

// EnterOpttablefuncelementlist is called when production opttablefuncelementlist is entered.
func (s *BaseRedshiftParserListener) EnterOpttablefuncelementlist(ctx *OpttablefuncelementlistContext) {
}

// ExitOpttablefuncelementlist is called when production opttablefuncelementlist is exited.
func (s *BaseRedshiftParserListener) ExitOpttablefuncelementlist(ctx *OpttablefuncelementlistContext) {
}

// EnterTablefuncelementlist is called when production tablefuncelementlist is entered.
func (s *BaseRedshiftParserListener) EnterTablefuncelementlist(ctx *TablefuncelementlistContext) {}

// ExitTablefuncelementlist is called when production tablefuncelementlist is exited.
func (s *BaseRedshiftParserListener) ExitTablefuncelementlist(ctx *TablefuncelementlistContext) {}

// EnterTablefuncelement is called when production tablefuncelement is entered.
func (s *BaseRedshiftParserListener) EnterTablefuncelement(ctx *TablefuncelementContext) {}

// ExitTablefuncelement is called when production tablefuncelement is exited.
func (s *BaseRedshiftParserListener) ExitTablefuncelement(ctx *TablefuncelementContext) {}

// EnterXmltable is called when production xmltable is entered.
func (s *BaseRedshiftParserListener) EnterXmltable(ctx *XmltableContext) {}

// ExitXmltable is called when production xmltable is exited.
func (s *BaseRedshiftParserListener) ExitXmltable(ctx *XmltableContext) {}

// EnterXmltable_column_list is called when production xmltable_column_list is entered.
func (s *BaseRedshiftParserListener) EnterXmltable_column_list(ctx *Xmltable_column_listContext) {}

// ExitXmltable_column_list is called when production xmltable_column_list is exited.
func (s *BaseRedshiftParserListener) ExitXmltable_column_list(ctx *Xmltable_column_listContext) {}

// EnterXmltable_column_el is called when production xmltable_column_el is entered.
func (s *BaseRedshiftParserListener) EnterXmltable_column_el(ctx *Xmltable_column_elContext) {}

// ExitXmltable_column_el is called when production xmltable_column_el is exited.
func (s *BaseRedshiftParserListener) ExitXmltable_column_el(ctx *Xmltable_column_elContext) {}

// EnterXmltable_column_option_list is called when production xmltable_column_option_list is entered.
func (s *BaseRedshiftParserListener) EnterXmltable_column_option_list(ctx *Xmltable_column_option_listContext) {
}

// ExitXmltable_column_option_list is called when production xmltable_column_option_list is exited.
func (s *BaseRedshiftParserListener) ExitXmltable_column_option_list(ctx *Xmltable_column_option_listContext) {
}

// EnterXmltable_column_option_el is called when production xmltable_column_option_el is entered.
func (s *BaseRedshiftParserListener) EnterXmltable_column_option_el(ctx *Xmltable_column_option_elContext) {
}

// ExitXmltable_column_option_el is called when production xmltable_column_option_el is exited.
func (s *BaseRedshiftParserListener) ExitXmltable_column_option_el(ctx *Xmltable_column_option_elContext) {
}

// EnterXml_namespace_list is called when production xml_namespace_list is entered.
func (s *BaseRedshiftParserListener) EnterXml_namespace_list(ctx *Xml_namespace_listContext) {}

// ExitXml_namespace_list is called when production xml_namespace_list is exited.
func (s *BaseRedshiftParserListener) ExitXml_namespace_list(ctx *Xml_namespace_listContext) {}

// EnterXml_namespace_el is called when production xml_namespace_el is entered.
func (s *BaseRedshiftParserListener) EnterXml_namespace_el(ctx *Xml_namespace_elContext) {}

// ExitXml_namespace_el is called when production xml_namespace_el is exited.
func (s *BaseRedshiftParserListener) ExitXml_namespace_el(ctx *Xml_namespace_elContext) {}

// EnterTypename is called when production typename is entered.
func (s *BaseRedshiftParserListener) EnterTypename(ctx *TypenameContext) {}

// ExitTypename is called when production typename is exited.
func (s *BaseRedshiftParserListener) ExitTypename(ctx *TypenameContext) {}

// EnterOpt_array_bounds is called when production opt_array_bounds is entered.
func (s *BaseRedshiftParserListener) EnterOpt_array_bounds(ctx *Opt_array_boundsContext) {}

// ExitOpt_array_bounds is called when production opt_array_bounds is exited.
func (s *BaseRedshiftParserListener) ExitOpt_array_bounds(ctx *Opt_array_boundsContext) {}

// EnterSimpletypename is called when production simpletypename is entered.
func (s *BaseRedshiftParserListener) EnterSimpletypename(ctx *SimpletypenameContext) {}

// ExitSimpletypename is called when production simpletypename is exited.
func (s *BaseRedshiftParserListener) ExitSimpletypename(ctx *SimpletypenameContext) {}

// EnterVarbyte is called when production varbyte is entered.
func (s *BaseRedshiftParserListener) EnterVarbyte(ctx *VarbyteContext) {}

// ExitVarbyte is called when production varbyte is exited.
func (s *BaseRedshiftParserListener) ExitVarbyte(ctx *VarbyteContext) {}

// EnterJson_type is called when production json_type is entered.
func (s *BaseRedshiftParserListener) EnterJson_type(ctx *Json_typeContext) {}

// ExitJson_type is called when production json_type is exited.
func (s *BaseRedshiftParserListener) ExitJson_type(ctx *Json_typeContext) {}

// EnterConsttypename is called when production consttypename is entered.
func (s *BaseRedshiftParserListener) EnterConsttypename(ctx *ConsttypenameContext) {}

// ExitConsttypename is called when production consttypename is exited.
func (s *BaseRedshiftParserListener) ExitConsttypename(ctx *ConsttypenameContext) {}

// EnterGenerictype is called when production generictype is entered.
func (s *BaseRedshiftParserListener) EnterGenerictype(ctx *GenerictypeContext) {}

// ExitGenerictype is called when production generictype is exited.
func (s *BaseRedshiftParserListener) ExitGenerictype(ctx *GenerictypeContext) {}

// EnterOpt_type_modifiers is called when production opt_type_modifiers is entered.
func (s *BaseRedshiftParserListener) EnterOpt_type_modifiers(ctx *Opt_type_modifiersContext) {}

// ExitOpt_type_modifiers is called when production opt_type_modifiers is exited.
func (s *BaseRedshiftParserListener) ExitOpt_type_modifiers(ctx *Opt_type_modifiersContext) {}

// EnterNumeric is called when production numeric is entered.
func (s *BaseRedshiftParserListener) EnterNumeric(ctx *NumericContext) {}

// ExitNumeric is called when production numeric is exited.
func (s *BaseRedshiftParserListener) ExitNumeric(ctx *NumericContext) {}

// EnterOpt_float is called when production opt_float is entered.
func (s *BaseRedshiftParserListener) EnterOpt_float(ctx *Opt_floatContext) {}

// ExitOpt_float is called when production opt_float is exited.
func (s *BaseRedshiftParserListener) ExitOpt_float(ctx *Opt_floatContext) {}

// EnterBit is called when production bit is entered.
func (s *BaseRedshiftParserListener) EnterBit(ctx *BitContext) {}

// ExitBit is called when production bit is exited.
func (s *BaseRedshiftParserListener) ExitBit(ctx *BitContext) {}

// EnterConstbit is called when production constbit is entered.
func (s *BaseRedshiftParserListener) EnterConstbit(ctx *ConstbitContext) {}

// ExitConstbit is called when production constbit is exited.
func (s *BaseRedshiftParserListener) ExitConstbit(ctx *ConstbitContext) {}

// EnterBitwithlength is called when production bitwithlength is entered.
func (s *BaseRedshiftParserListener) EnterBitwithlength(ctx *BitwithlengthContext) {}

// ExitBitwithlength is called when production bitwithlength is exited.
func (s *BaseRedshiftParserListener) ExitBitwithlength(ctx *BitwithlengthContext) {}

// EnterBitwithoutlength is called when production bitwithoutlength is entered.
func (s *BaseRedshiftParserListener) EnterBitwithoutlength(ctx *BitwithoutlengthContext) {}

// ExitBitwithoutlength is called when production bitwithoutlength is exited.
func (s *BaseRedshiftParserListener) ExitBitwithoutlength(ctx *BitwithoutlengthContext) {}

// EnterCharacter is called when production character is entered.
func (s *BaseRedshiftParserListener) EnterCharacter(ctx *CharacterContext) {}

// ExitCharacter is called when production character is exited.
func (s *BaseRedshiftParserListener) ExitCharacter(ctx *CharacterContext) {}

// EnterConstcharacter is called when production constcharacter is entered.
func (s *BaseRedshiftParserListener) EnterConstcharacter(ctx *ConstcharacterContext) {}

// ExitConstcharacter is called when production constcharacter is exited.
func (s *BaseRedshiftParserListener) ExitConstcharacter(ctx *ConstcharacterContext) {}

// EnterCharacter_c is called when production character_c is entered.
func (s *BaseRedshiftParserListener) EnterCharacter_c(ctx *Character_cContext) {}

// ExitCharacter_c is called when production character_c is exited.
func (s *BaseRedshiftParserListener) ExitCharacter_c(ctx *Character_cContext) {}

// EnterOpt_varying is called when production opt_varying is entered.
func (s *BaseRedshiftParserListener) EnterOpt_varying(ctx *Opt_varyingContext) {}

// ExitOpt_varying is called when production opt_varying is exited.
func (s *BaseRedshiftParserListener) ExitOpt_varying(ctx *Opt_varyingContext) {}

// EnterConstdatetime is called when production constdatetime is entered.
func (s *BaseRedshiftParserListener) EnterConstdatetime(ctx *ConstdatetimeContext) {}

// ExitConstdatetime is called when production constdatetime is exited.
func (s *BaseRedshiftParserListener) ExitConstdatetime(ctx *ConstdatetimeContext) {}

// EnterConstinterval is called when production constinterval is entered.
func (s *BaseRedshiftParserListener) EnterConstinterval(ctx *ConstintervalContext) {}

// ExitConstinterval is called when production constinterval is exited.
func (s *BaseRedshiftParserListener) ExitConstinterval(ctx *ConstintervalContext) {}

// EnterOpt_timezone is called when production opt_timezone is entered.
func (s *BaseRedshiftParserListener) EnterOpt_timezone(ctx *Opt_timezoneContext) {}

// ExitOpt_timezone is called when production opt_timezone is exited.
func (s *BaseRedshiftParserListener) ExitOpt_timezone(ctx *Opt_timezoneContext) {}

// EnterOpt_interval is called when production opt_interval is entered.
func (s *BaseRedshiftParserListener) EnterOpt_interval(ctx *Opt_intervalContext) {}

// ExitOpt_interval is called when production opt_interval is exited.
func (s *BaseRedshiftParserListener) ExitOpt_interval(ctx *Opt_intervalContext) {}

// EnterInterval_second is called when production interval_second is entered.
func (s *BaseRedshiftParserListener) EnterInterval_second(ctx *Interval_secondContext) {}

// ExitInterval_second is called when production interval_second is exited.
func (s *BaseRedshiftParserListener) ExitInterval_second(ctx *Interval_secondContext) {}

// EnterOpt_escape is called when production opt_escape is entered.
func (s *BaseRedshiftParserListener) EnterOpt_escape(ctx *Opt_escapeContext) {}

// ExitOpt_escape is called when production opt_escape is exited.
func (s *BaseRedshiftParserListener) ExitOpt_escape(ctx *Opt_escapeContext) {}

// EnterA_expr is called when production a_expr is entered.
func (s *BaseRedshiftParserListener) EnterA_expr(ctx *A_exprContext) {}

// ExitA_expr is called when production a_expr is exited.
func (s *BaseRedshiftParserListener) ExitA_expr(ctx *A_exprContext) {}

// EnterA_expr_qual is called when production a_expr_qual is entered.
func (s *BaseRedshiftParserListener) EnterA_expr_qual(ctx *A_expr_qualContext) {}

// ExitA_expr_qual is called when production a_expr_qual is exited.
func (s *BaseRedshiftParserListener) ExitA_expr_qual(ctx *A_expr_qualContext) {}

// EnterA_expr_lessless is called when production a_expr_lessless is entered.
func (s *BaseRedshiftParserListener) EnterA_expr_lessless(ctx *A_expr_lesslessContext) {}

// ExitA_expr_lessless is called when production a_expr_lessless is exited.
func (s *BaseRedshiftParserListener) ExitA_expr_lessless(ctx *A_expr_lesslessContext) {}

// EnterA_expr_or is called when production a_expr_or is entered.
func (s *BaseRedshiftParserListener) EnterA_expr_or(ctx *A_expr_orContext) {}

// ExitA_expr_or is called when production a_expr_or is exited.
func (s *BaseRedshiftParserListener) ExitA_expr_or(ctx *A_expr_orContext) {}

// EnterA_expr_and is called when production a_expr_and is entered.
func (s *BaseRedshiftParserListener) EnterA_expr_and(ctx *A_expr_andContext) {}

// ExitA_expr_and is called when production a_expr_and is exited.
func (s *BaseRedshiftParserListener) ExitA_expr_and(ctx *A_expr_andContext) {}

// EnterA_expr_between is called when production a_expr_between is entered.
func (s *BaseRedshiftParserListener) EnterA_expr_between(ctx *A_expr_betweenContext) {}

// ExitA_expr_between is called when production a_expr_between is exited.
func (s *BaseRedshiftParserListener) ExitA_expr_between(ctx *A_expr_betweenContext) {}

// EnterA_expr_in is called when production a_expr_in is entered.
func (s *BaseRedshiftParserListener) EnterA_expr_in(ctx *A_expr_inContext) {}

// ExitA_expr_in is called when production a_expr_in is exited.
func (s *BaseRedshiftParserListener) ExitA_expr_in(ctx *A_expr_inContext) {}

// EnterA_expr_unary_not is called when production a_expr_unary_not is entered.
func (s *BaseRedshiftParserListener) EnterA_expr_unary_not(ctx *A_expr_unary_notContext) {}

// ExitA_expr_unary_not is called when production a_expr_unary_not is exited.
func (s *BaseRedshiftParserListener) ExitA_expr_unary_not(ctx *A_expr_unary_notContext) {}

// EnterA_expr_isnull is called when production a_expr_isnull is entered.
func (s *BaseRedshiftParserListener) EnterA_expr_isnull(ctx *A_expr_isnullContext) {}

// ExitA_expr_isnull is called when production a_expr_isnull is exited.
func (s *BaseRedshiftParserListener) ExitA_expr_isnull(ctx *A_expr_isnullContext) {}

// EnterA_expr_is_not is called when production a_expr_is_not is entered.
func (s *BaseRedshiftParserListener) EnterA_expr_is_not(ctx *A_expr_is_notContext) {}

// ExitA_expr_is_not is called when production a_expr_is_not is exited.
func (s *BaseRedshiftParserListener) ExitA_expr_is_not(ctx *A_expr_is_notContext) {}

// EnterA_expr_compare is called when production a_expr_compare is entered.
func (s *BaseRedshiftParserListener) EnterA_expr_compare(ctx *A_expr_compareContext) {}

// ExitA_expr_compare is called when production a_expr_compare is exited.
func (s *BaseRedshiftParserListener) ExitA_expr_compare(ctx *A_expr_compareContext) {}

// EnterA_expr_prior_or_level is called when production a_expr_prior_or_level is entered.
func (s *BaseRedshiftParserListener) EnterA_expr_prior_or_level(ctx *A_expr_prior_or_levelContext) {}

// ExitA_expr_prior_or_level is called when production a_expr_prior_or_level is exited.
func (s *BaseRedshiftParserListener) ExitA_expr_prior_or_level(ctx *A_expr_prior_or_levelContext) {}

// EnterA_expr_like is called when production a_expr_like is entered.
func (s *BaseRedshiftParserListener) EnterA_expr_like(ctx *A_expr_likeContext) {}

// ExitA_expr_like is called when production a_expr_like is exited.
func (s *BaseRedshiftParserListener) ExitA_expr_like(ctx *A_expr_likeContext) {}

// EnterA_expr_qual_op is called when production a_expr_qual_op is entered.
func (s *BaseRedshiftParserListener) EnterA_expr_qual_op(ctx *A_expr_qual_opContext) {}

// ExitA_expr_qual_op is called when production a_expr_qual_op is exited.
func (s *BaseRedshiftParserListener) ExitA_expr_qual_op(ctx *A_expr_qual_opContext) {}

// EnterA_expr_unary_qualop is called when production a_expr_unary_qualop is entered.
func (s *BaseRedshiftParserListener) EnterA_expr_unary_qualop(ctx *A_expr_unary_qualopContext) {}

// ExitA_expr_unary_qualop is called when production a_expr_unary_qualop is exited.
func (s *BaseRedshiftParserListener) ExitA_expr_unary_qualop(ctx *A_expr_unary_qualopContext) {}

// EnterA_expr_add is called when production a_expr_add is entered.
func (s *BaseRedshiftParserListener) EnterA_expr_add(ctx *A_expr_addContext) {}

// ExitA_expr_add is called when production a_expr_add is exited.
func (s *BaseRedshiftParserListener) ExitA_expr_add(ctx *A_expr_addContext) {}

// EnterA_expr_mul is called when production a_expr_mul is entered.
func (s *BaseRedshiftParserListener) EnterA_expr_mul(ctx *A_expr_mulContext) {}

// ExitA_expr_mul is called when production a_expr_mul is exited.
func (s *BaseRedshiftParserListener) ExitA_expr_mul(ctx *A_expr_mulContext) {}

// EnterA_expr_caret is called when production a_expr_caret is entered.
func (s *BaseRedshiftParserListener) EnterA_expr_caret(ctx *A_expr_caretContext) {}

// ExitA_expr_caret is called when production a_expr_caret is exited.
func (s *BaseRedshiftParserListener) ExitA_expr_caret(ctx *A_expr_caretContext) {}

// EnterA_expr_unary_sign is called when production a_expr_unary_sign is entered.
func (s *BaseRedshiftParserListener) EnterA_expr_unary_sign(ctx *A_expr_unary_signContext) {}

// ExitA_expr_unary_sign is called when production a_expr_unary_sign is exited.
func (s *BaseRedshiftParserListener) ExitA_expr_unary_sign(ctx *A_expr_unary_signContext) {}

// EnterA_expr_at_time_zone is called when production a_expr_at_time_zone is entered.
func (s *BaseRedshiftParserListener) EnterA_expr_at_time_zone(ctx *A_expr_at_time_zoneContext) {}

// ExitA_expr_at_time_zone is called when production a_expr_at_time_zone is exited.
func (s *BaseRedshiftParserListener) ExitA_expr_at_time_zone(ctx *A_expr_at_time_zoneContext) {}

// EnterA_expr_collate is called when production a_expr_collate is entered.
func (s *BaseRedshiftParserListener) EnterA_expr_collate(ctx *A_expr_collateContext) {}

// ExitA_expr_collate is called when production a_expr_collate is exited.
func (s *BaseRedshiftParserListener) ExitA_expr_collate(ctx *A_expr_collateContext) {}

// EnterA_expr_typecast is called when production a_expr_typecast is entered.
func (s *BaseRedshiftParserListener) EnterA_expr_typecast(ctx *A_expr_typecastContext) {}

// ExitA_expr_typecast is called when production a_expr_typecast is exited.
func (s *BaseRedshiftParserListener) ExitA_expr_typecast(ctx *A_expr_typecastContext) {}

// EnterB_expr is called when production b_expr is entered.
func (s *BaseRedshiftParserListener) EnterB_expr(ctx *B_exprContext) {}

// ExitB_expr is called when production b_expr is exited.
func (s *BaseRedshiftParserListener) ExitB_expr(ctx *B_exprContext) {}

// EnterC_expr_exists is called when production c_expr_exists is entered.
func (s *BaseRedshiftParserListener) EnterC_expr_exists(ctx *C_expr_existsContext) {}

// ExitC_expr_exists is called when production c_expr_exists is exited.
func (s *BaseRedshiftParserListener) ExitC_expr_exists(ctx *C_expr_existsContext) {}

// EnterC_expr_expr is called when production c_expr_expr is entered.
func (s *BaseRedshiftParserListener) EnterC_expr_expr(ctx *C_expr_exprContext) {}

// ExitC_expr_expr is called when production c_expr_expr is exited.
func (s *BaseRedshiftParserListener) ExitC_expr_expr(ctx *C_expr_exprContext) {}

// EnterC_expr_case is called when production c_expr_case is entered.
func (s *BaseRedshiftParserListener) EnterC_expr_case(ctx *C_expr_caseContext) {}

// ExitC_expr_case is called when production c_expr_case is exited.
func (s *BaseRedshiftParserListener) ExitC_expr_case(ctx *C_expr_caseContext) {}

// EnterPlsqlvariablename is called when production plsqlvariablename is entered.
func (s *BaseRedshiftParserListener) EnterPlsqlvariablename(ctx *PlsqlvariablenameContext) {}

// ExitPlsqlvariablename is called when production plsqlvariablename is exited.
func (s *BaseRedshiftParserListener) ExitPlsqlvariablename(ctx *PlsqlvariablenameContext) {}

// EnterFunc_application is called when production func_application is entered.
func (s *BaseRedshiftParserListener) EnterFunc_application(ctx *Func_applicationContext) {}

// ExitFunc_application is called when production func_application is exited.
func (s *BaseRedshiftParserListener) ExitFunc_application(ctx *Func_applicationContext) {}

// EnterFunc_expr is called when production func_expr is entered.
func (s *BaseRedshiftParserListener) EnterFunc_expr(ctx *Func_exprContext) {}

// ExitFunc_expr is called when production func_expr is exited.
func (s *BaseRedshiftParserListener) ExitFunc_expr(ctx *Func_exprContext) {}

// EnterFunc_expr_windowless is called when production func_expr_windowless is entered.
func (s *BaseRedshiftParserListener) EnterFunc_expr_windowless(ctx *Func_expr_windowlessContext) {}

// ExitFunc_expr_windowless is called when production func_expr_windowless is exited.
func (s *BaseRedshiftParserListener) ExitFunc_expr_windowless(ctx *Func_expr_windowlessContext) {}

// EnterFunc_expr_common_subexpr is called when production func_expr_common_subexpr is entered.
func (s *BaseRedshiftParserListener) EnterFunc_expr_common_subexpr(ctx *Func_expr_common_subexprContext) {
}

// ExitFunc_expr_common_subexpr is called when production func_expr_common_subexpr is exited.
func (s *BaseRedshiftParserListener) ExitFunc_expr_common_subexpr(ctx *Func_expr_common_subexprContext) {
}

// EnterXml_root_version is called when production xml_root_version is entered.
func (s *BaseRedshiftParserListener) EnterXml_root_version(ctx *Xml_root_versionContext) {}

// ExitXml_root_version is called when production xml_root_version is exited.
func (s *BaseRedshiftParserListener) ExitXml_root_version(ctx *Xml_root_versionContext) {}

// EnterOpt_xml_root_standalone is called when production opt_xml_root_standalone is entered.
func (s *BaseRedshiftParserListener) EnterOpt_xml_root_standalone(ctx *Opt_xml_root_standaloneContext) {
}

// ExitOpt_xml_root_standalone is called when production opt_xml_root_standalone is exited.
func (s *BaseRedshiftParserListener) ExitOpt_xml_root_standalone(ctx *Opt_xml_root_standaloneContext) {
}

// EnterXml_attributes is called when production xml_attributes is entered.
func (s *BaseRedshiftParserListener) EnterXml_attributes(ctx *Xml_attributesContext) {}

// ExitXml_attributes is called when production xml_attributes is exited.
func (s *BaseRedshiftParserListener) ExitXml_attributes(ctx *Xml_attributesContext) {}

// EnterXml_attribute_list is called when production xml_attribute_list is entered.
func (s *BaseRedshiftParserListener) EnterXml_attribute_list(ctx *Xml_attribute_listContext) {}

// ExitXml_attribute_list is called when production xml_attribute_list is exited.
func (s *BaseRedshiftParserListener) ExitXml_attribute_list(ctx *Xml_attribute_listContext) {}

// EnterXml_attribute_el is called when production xml_attribute_el is entered.
func (s *BaseRedshiftParserListener) EnterXml_attribute_el(ctx *Xml_attribute_elContext) {}

// ExitXml_attribute_el is called when production xml_attribute_el is exited.
func (s *BaseRedshiftParserListener) ExitXml_attribute_el(ctx *Xml_attribute_elContext) {}

// EnterDocument_or_content is called when production document_or_content is entered.
func (s *BaseRedshiftParserListener) EnterDocument_or_content(ctx *Document_or_contentContext) {}

// ExitDocument_or_content is called when production document_or_content is exited.
func (s *BaseRedshiftParserListener) ExitDocument_or_content(ctx *Document_or_contentContext) {}

// EnterXml_whitespace_option is called when production xml_whitespace_option is entered.
func (s *BaseRedshiftParserListener) EnterXml_whitespace_option(ctx *Xml_whitespace_optionContext) {}

// ExitXml_whitespace_option is called when production xml_whitespace_option is exited.
func (s *BaseRedshiftParserListener) ExitXml_whitespace_option(ctx *Xml_whitespace_optionContext) {}

// EnterXmlexists_argument is called when production xmlexists_argument is entered.
func (s *BaseRedshiftParserListener) EnterXmlexists_argument(ctx *Xmlexists_argumentContext) {}

// ExitXmlexists_argument is called when production xmlexists_argument is exited.
func (s *BaseRedshiftParserListener) ExitXmlexists_argument(ctx *Xmlexists_argumentContext) {}

// EnterXml_passing_mech is called when production xml_passing_mech is entered.
func (s *BaseRedshiftParserListener) EnterXml_passing_mech(ctx *Xml_passing_mechContext) {}

// ExitXml_passing_mech is called when production xml_passing_mech is exited.
func (s *BaseRedshiftParserListener) ExitXml_passing_mech(ctx *Xml_passing_mechContext) {}

// EnterWithin_group_clause is called when production within_group_clause is entered.
func (s *BaseRedshiftParserListener) EnterWithin_group_clause(ctx *Within_group_clauseContext) {}

// ExitWithin_group_clause is called when production within_group_clause is exited.
func (s *BaseRedshiftParserListener) ExitWithin_group_clause(ctx *Within_group_clauseContext) {}

// EnterFilter_clause is called when production filter_clause is entered.
func (s *BaseRedshiftParserListener) EnterFilter_clause(ctx *Filter_clauseContext) {}

// ExitFilter_clause is called when production filter_clause is exited.
func (s *BaseRedshiftParserListener) ExitFilter_clause(ctx *Filter_clauseContext) {}

// EnterWindow_clause is called when production window_clause is entered.
func (s *BaseRedshiftParserListener) EnterWindow_clause(ctx *Window_clauseContext) {}

// ExitWindow_clause is called when production window_clause is exited.
func (s *BaseRedshiftParserListener) ExitWindow_clause(ctx *Window_clauseContext) {}

// EnterWindow_definition_list is called when production window_definition_list is entered.
func (s *BaseRedshiftParserListener) EnterWindow_definition_list(ctx *Window_definition_listContext) {
}

// ExitWindow_definition_list is called when production window_definition_list is exited.
func (s *BaseRedshiftParserListener) ExitWindow_definition_list(ctx *Window_definition_listContext) {}

// EnterWindow_definition is called when production window_definition is entered.
func (s *BaseRedshiftParserListener) EnterWindow_definition(ctx *Window_definitionContext) {}

// ExitWindow_definition is called when production window_definition is exited.
func (s *BaseRedshiftParserListener) ExitWindow_definition(ctx *Window_definitionContext) {}

// EnterOver_clause is called when production over_clause is entered.
func (s *BaseRedshiftParserListener) EnterOver_clause(ctx *Over_clauseContext) {}

// ExitOver_clause is called when production over_clause is exited.
func (s *BaseRedshiftParserListener) ExitOver_clause(ctx *Over_clauseContext) {}

// EnterWindow_specification is called when production window_specification is entered.
func (s *BaseRedshiftParserListener) EnterWindow_specification(ctx *Window_specificationContext) {}

// ExitWindow_specification is called when production window_specification is exited.
func (s *BaseRedshiftParserListener) ExitWindow_specification(ctx *Window_specificationContext) {}

// EnterOpt_existing_window_name is called when production opt_existing_window_name is entered.
func (s *BaseRedshiftParserListener) EnterOpt_existing_window_name(ctx *Opt_existing_window_nameContext) {
}

// ExitOpt_existing_window_name is called when production opt_existing_window_name is exited.
func (s *BaseRedshiftParserListener) ExitOpt_existing_window_name(ctx *Opt_existing_window_nameContext) {
}

// EnterOpt_partition_clause is called when production opt_partition_clause is entered.
func (s *BaseRedshiftParserListener) EnterOpt_partition_clause(ctx *Opt_partition_clauseContext) {}

// ExitOpt_partition_clause is called when production opt_partition_clause is exited.
func (s *BaseRedshiftParserListener) ExitOpt_partition_clause(ctx *Opt_partition_clauseContext) {}

// EnterOpt_frame_clause is called when production opt_frame_clause is entered.
func (s *BaseRedshiftParserListener) EnterOpt_frame_clause(ctx *Opt_frame_clauseContext) {}

// ExitOpt_frame_clause is called when production opt_frame_clause is exited.
func (s *BaseRedshiftParserListener) ExitOpt_frame_clause(ctx *Opt_frame_clauseContext) {}

// EnterFrame_extent is called when production frame_extent is entered.
func (s *BaseRedshiftParserListener) EnterFrame_extent(ctx *Frame_extentContext) {}

// ExitFrame_extent is called when production frame_extent is exited.
func (s *BaseRedshiftParserListener) ExitFrame_extent(ctx *Frame_extentContext) {}

// EnterFrame_bound is called when production frame_bound is entered.
func (s *BaseRedshiftParserListener) EnterFrame_bound(ctx *Frame_boundContext) {}

// ExitFrame_bound is called when production frame_bound is exited.
func (s *BaseRedshiftParserListener) ExitFrame_bound(ctx *Frame_boundContext) {}

// EnterOpt_window_exclusion_clause is called when production opt_window_exclusion_clause is entered.
func (s *BaseRedshiftParserListener) EnterOpt_window_exclusion_clause(ctx *Opt_window_exclusion_clauseContext) {
}

// ExitOpt_window_exclusion_clause is called when production opt_window_exclusion_clause is exited.
func (s *BaseRedshiftParserListener) ExitOpt_window_exclusion_clause(ctx *Opt_window_exclusion_clauseContext) {
}

// EnterRow is called when production row is entered.
func (s *BaseRedshiftParserListener) EnterRow(ctx *RowContext) {}

// ExitRow is called when production row is exited.
func (s *BaseRedshiftParserListener) ExitRow(ctx *RowContext) {}

// EnterExplicit_row is called when production explicit_row is entered.
func (s *BaseRedshiftParserListener) EnterExplicit_row(ctx *Explicit_rowContext) {}

// ExitExplicit_row is called when production explicit_row is exited.
func (s *BaseRedshiftParserListener) ExitExplicit_row(ctx *Explicit_rowContext) {}

// EnterImplicit_row is called when production implicit_row is entered.
func (s *BaseRedshiftParserListener) EnterImplicit_row(ctx *Implicit_rowContext) {}

// ExitImplicit_row is called when production implicit_row is exited.
func (s *BaseRedshiftParserListener) ExitImplicit_row(ctx *Implicit_rowContext) {}

// EnterSub_type is called when production sub_type is entered.
func (s *BaseRedshiftParserListener) EnterSub_type(ctx *Sub_typeContext) {}

// ExitSub_type is called when production sub_type is exited.
func (s *BaseRedshiftParserListener) ExitSub_type(ctx *Sub_typeContext) {}

// EnterAll_op is called when production all_op is entered.
func (s *BaseRedshiftParserListener) EnterAll_op(ctx *All_opContext) {}

// ExitAll_op is called when production all_op is exited.
func (s *BaseRedshiftParserListener) ExitAll_op(ctx *All_opContext) {}

// EnterMathop is called when production mathop is entered.
func (s *BaseRedshiftParserListener) EnterMathop(ctx *MathopContext) {}

// ExitMathop is called when production mathop is exited.
func (s *BaseRedshiftParserListener) ExitMathop(ctx *MathopContext) {}

// EnterQual_op is called when production qual_op is entered.
func (s *BaseRedshiftParserListener) EnterQual_op(ctx *Qual_opContext) {}

// ExitQual_op is called when production qual_op is exited.
func (s *BaseRedshiftParserListener) ExitQual_op(ctx *Qual_opContext) {}

// EnterQual_all_op is called when production qual_all_op is entered.
func (s *BaseRedshiftParserListener) EnterQual_all_op(ctx *Qual_all_opContext) {}

// ExitQual_all_op is called when production qual_all_op is exited.
func (s *BaseRedshiftParserListener) ExitQual_all_op(ctx *Qual_all_opContext) {}

// EnterSubquery_Op is called when production subquery_Op is entered.
func (s *BaseRedshiftParserListener) EnterSubquery_Op(ctx *Subquery_OpContext) {}

// ExitSubquery_Op is called when production subquery_Op is exited.
func (s *BaseRedshiftParserListener) ExitSubquery_Op(ctx *Subquery_OpContext) {}

// EnterExpr_list is called when production expr_list is entered.
func (s *BaseRedshiftParserListener) EnterExpr_list(ctx *Expr_listContext) {}

// ExitExpr_list is called when production expr_list is exited.
func (s *BaseRedshiftParserListener) ExitExpr_list(ctx *Expr_listContext) {}

// EnterFunc_arg_list is called when production func_arg_list is entered.
func (s *BaseRedshiftParserListener) EnterFunc_arg_list(ctx *Func_arg_listContext) {}

// ExitFunc_arg_list is called when production func_arg_list is exited.
func (s *BaseRedshiftParserListener) ExitFunc_arg_list(ctx *Func_arg_listContext) {}

// EnterFunc_arg_expr is called when production func_arg_expr is entered.
func (s *BaseRedshiftParserListener) EnterFunc_arg_expr(ctx *Func_arg_exprContext) {}

// ExitFunc_arg_expr is called when production func_arg_expr is exited.
func (s *BaseRedshiftParserListener) ExitFunc_arg_expr(ctx *Func_arg_exprContext) {}

// EnterType_list is called when production type_list is entered.
func (s *BaseRedshiftParserListener) EnterType_list(ctx *Type_listContext) {}

// ExitType_list is called when production type_list is exited.
func (s *BaseRedshiftParserListener) ExitType_list(ctx *Type_listContext) {}

// EnterArray_expr is called when production array_expr is entered.
func (s *BaseRedshiftParserListener) EnterArray_expr(ctx *Array_exprContext) {}

// ExitArray_expr is called when production array_expr is exited.
func (s *BaseRedshiftParserListener) ExitArray_expr(ctx *Array_exprContext) {}

// EnterArray_expr_list is called when production array_expr_list is entered.
func (s *BaseRedshiftParserListener) EnterArray_expr_list(ctx *Array_expr_listContext) {}

// ExitArray_expr_list is called when production array_expr_list is exited.
func (s *BaseRedshiftParserListener) ExitArray_expr_list(ctx *Array_expr_listContext) {}

// EnterExtract_list is called when production extract_list is entered.
func (s *BaseRedshiftParserListener) EnterExtract_list(ctx *Extract_listContext) {}

// ExitExtract_list is called when production extract_list is exited.
func (s *BaseRedshiftParserListener) ExitExtract_list(ctx *Extract_listContext) {}

// EnterExtract_arg is called when production extract_arg is entered.
func (s *BaseRedshiftParserListener) EnterExtract_arg(ctx *Extract_argContext) {}

// ExitExtract_arg is called when production extract_arg is exited.
func (s *BaseRedshiftParserListener) ExitExtract_arg(ctx *Extract_argContext) {}

// EnterUnicode_normal_form is called when production unicode_normal_form is entered.
func (s *BaseRedshiftParserListener) EnterUnicode_normal_form(ctx *Unicode_normal_formContext) {}

// ExitUnicode_normal_form is called when production unicode_normal_form is exited.
func (s *BaseRedshiftParserListener) ExitUnicode_normal_form(ctx *Unicode_normal_formContext) {}

// EnterOverlay_list is called when production overlay_list is entered.
func (s *BaseRedshiftParserListener) EnterOverlay_list(ctx *Overlay_listContext) {}

// ExitOverlay_list is called when production overlay_list is exited.
func (s *BaseRedshiftParserListener) ExitOverlay_list(ctx *Overlay_listContext) {}

// EnterPosition_list is called when production position_list is entered.
func (s *BaseRedshiftParserListener) EnterPosition_list(ctx *Position_listContext) {}

// ExitPosition_list is called when production position_list is exited.
func (s *BaseRedshiftParserListener) ExitPosition_list(ctx *Position_listContext) {}

// EnterSubstr_list is called when production substr_list is entered.
func (s *BaseRedshiftParserListener) EnterSubstr_list(ctx *Substr_listContext) {}

// ExitSubstr_list is called when production substr_list is exited.
func (s *BaseRedshiftParserListener) ExitSubstr_list(ctx *Substr_listContext) {}

// EnterTrim_list is called when production trim_list is entered.
func (s *BaseRedshiftParserListener) EnterTrim_list(ctx *Trim_listContext) {}

// ExitTrim_list is called when production trim_list is exited.
func (s *BaseRedshiftParserListener) ExitTrim_list(ctx *Trim_listContext) {}

// EnterIn_expr_select is called when production in_expr_select is entered.
func (s *BaseRedshiftParserListener) EnterIn_expr_select(ctx *In_expr_selectContext) {}

// ExitIn_expr_select is called when production in_expr_select is exited.
func (s *BaseRedshiftParserListener) ExitIn_expr_select(ctx *In_expr_selectContext) {}

// EnterIn_expr_list is called when production in_expr_list is entered.
func (s *BaseRedshiftParserListener) EnterIn_expr_list(ctx *In_expr_listContext) {}

// ExitIn_expr_list is called when production in_expr_list is exited.
func (s *BaseRedshiftParserListener) ExitIn_expr_list(ctx *In_expr_listContext) {}

// EnterCase_expr is called when production case_expr is entered.
func (s *BaseRedshiftParserListener) EnterCase_expr(ctx *Case_exprContext) {}

// ExitCase_expr is called when production case_expr is exited.
func (s *BaseRedshiftParserListener) ExitCase_expr(ctx *Case_exprContext) {}

// EnterWhen_clause_list is called when production when_clause_list is entered.
func (s *BaseRedshiftParserListener) EnterWhen_clause_list(ctx *When_clause_listContext) {}

// ExitWhen_clause_list is called when production when_clause_list is exited.
func (s *BaseRedshiftParserListener) ExitWhen_clause_list(ctx *When_clause_listContext) {}

// EnterWhen_clause is called when production when_clause is entered.
func (s *BaseRedshiftParserListener) EnterWhen_clause(ctx *When_clauseContext) {}

// ExitWhen_clause is called when production when_clause is exited.
func (s *BaseRedshiftParserListener) ExitWhen_clause(ctx *When_clauseContext) {}

// EnterCase_default is called when production case_default is entered.
func (s *BaseRedshiftParserListener) EnterCase_default(ctx *Case_defaultContext) {}

// ExitCase_default is called when production case_default is exited.
func (s *BaseRedshiftParserListener) ExitCase_default(ctx *Case_defaultContext) {}

// EnterCase_arg is called when production case_arg is entered.
func (s *BaseRedshiftParserListener) EnterCase_arg(ctx *Case_argContext) {}

// ExitCase_arg is called when production case_arg is exited.
func (s *BaseRedshiftParserListener) ExitCase_arg(ctx *Case_argContext) {}

// EnterColumnref is called when production columnref is entered.
func (s *BaseRedshiftParserListener) EnterColumnref(ctx *ColumnrefContext) {}

// ExitColumnref is called when production columnref is exited.
func (s *BaseRedshiftParserListener) ExitColumnref(ctx *ColumnrefContext) {}

// EnterIndirection_el is called when production indirection_el is entered.
func (s *BaseRedshiftParserListener) EnterIndirection_el(ctx *Indirection_elContext) {}

// ExitIndirection_el is called when production indirection_el is exited.
func (s *BaseRedshiftParserListener) ExitIndirection_el(ctx *Indirection_elContext) {}

// EnterOpt_slice_bound is called when production opt_slice_bound is entered.
func (s *BaseRedshiftParserListener) EnterOpt_slice_bound(ctx *Opt_slice_boundContext) {}

// ExitOpt_slice_bound is called when production opt_slice_bound is exited.
func (s *BaseRedshiftParserListener) ExitOpt_slice_bound(ctx *Opt_slice_boundContext) {}

// EnterIndirection is called when production indirection is entered.
func (s *BaseRedshiftParserListener) EnterIndirection(ctx *IndirectionContext) {}

// ExitIndirection is called when production indirection is exited.
func (s *BaseRedshiftParserListener) ExitIndirection(ctx *IndirectionContext) {}

// EnterOpt_indirection is called when production opt_indirection is entered.
func (s *BaseRedshiftParserListener) EnterOpt_indirection(ctx *Opt_indirectionContext) {}

// ExitOpt_indirection is called when production opt_indirection is exited.
func (s *BaseRedshiftParserListener) ExitOpt_indirection(ctx *Opt_indirectionContext) {}

// EnterOpt_target_list is called when production opt_target_list is entered.
func (s *BaseRedshiftParserListener) EnterOpt_target_list(ctx *Opt_target_listContext) {}

// ExitOpt_target_list is called when production opt_target_list is exited.
func (s *BaseRedshiftParserListener) ExitOpt_target_list(ctx *Opt_target_listContext) {}

// EnterTarget_list is called when production target_list is entered.
func (s *BaseRedshiftParserListener) EnterTarget_list(ctx *Target_listContext) {}

// ExitTarget_list is called when production target_list is exited.
func (s *BaseRedshiftParserListener) ExitTarget_list(ctx *Target_listContext) {}

// EnterTarget_label is called when production target_label is entered.
func (s *BaseRedshiftParserListener) EnterTarget_label(ctx *Target_labelContext) {}

// ExitTarget_label is called when production target_label is exited.
func (s *BaseRedshiftParserListener) ExitTarget_label(ctx *Target_labelContext) {}

// EnterTarget_star is called when production target_star is entered.
func (s *BaseRedshiftParserListener) EnterTarget_star(ctx *Target_starContext) {}

// ExitTarget_star is called when production target_star is exited.
func (s *BaseRedshiftParserListener) ExitTarget_star(ctx *Target_starContext) {}

// EnterTarget_alias is called when production target_alias is entered.
func (s *BaseRedshiftParserListener) EnterTarget_alias(ctx *Target_aliasContext) {}

// ExitTarget_alias is called when production target_alias is exited.
func (s *BaseRedshiftParserListener) ExitTarget_alias(ctx *Target_aliasContext) {}

// EnterQualified_name_list is called when production qualified_name_list is entered.
func (s *BaseRedshiftParserListener) EnterQualified_name_list(ctx *Qualified_name_listContext) {}

// ExitQualified_name_list is called when production qualified_name_list is exited.
func (s *BaseRedshiftParserListener) ExitQualified_name_list(ctx *Qualified_name_listContext) {}

// EnterQualified_name is called when production qualified_name is entered.
func (s *BaseRedshiftParserListener) EnterQualified_name(ctx *Qualified_nameContext) {}

// ExitQualified_name is called when production qualified_name is exited.
func (s *BaseRedshiftParserListener) ExitQualified_name(ctx *Qualified_nameContext) {}

// EnterName_list is called when production name_list is entered.
func (s *BaseRedshiftParserListener) EnterName_list(ctx *Name_listContext) {}

// ExitName_list is called when production name_list is exited.
func (s *BaseRedshiftParserListener) ExitName_list(ctx *Name_listContext) {}

// EnterName is called when production name is entered.
func (s *BaseRedshiftParserListener) EnterName(ctx *NameContext) {}

// ExitName is called when production name is exited.
func (s *BaseRedshiftParserListener) ExitName(ctx *NameContext) {}

// EnterAttr_name is called when production attr_name is entered.
func (s *BaseRedshiftParserListener) EnterAttr_name(ctx *Attr_nameContext) {}

// ExitAttr_name is called when production attr_name is exited.
func (s *BaseRedshiftParserListener) ExitAttr_name(ctx *Attr_nameContext) {}

// EnterFile_name is called when production file_name is entered.
func (s *BaseRedshiftParserListener) EnterFile_name(ctx *File_nameContext) {}

// ExitFile_name is called when production file_name is exited.
func (s *BaseRedshiftParserListener) ExitFile_name(ctx *File_nameContext) {}

// EnterFunc_name is called when production func_name is entered.
func (s *BaseRedshiftParserListener) EnterFunc_name(ctx *Func_nameContext) {}

// ExitFunc_name is called when production func_name is exited.
func (s *BaseRedshiftParserListener) ExitFunc_name(ctx *Func_nameContext) {}

// EnterAexprconst is called when production aexprconst is entered.
func (s *BaseRedshiftParserListener) EnterAexprconst(ctx *AexprconstContext) {}

// ExitAexprconst is called when production aexprconst is exited.
func (s *BaseRedshiftParserListener) ExitAexprconst(ctx *AexprconstContext) {}

// EnterXconst is called when production xconst is entered.
func (s *BaseRedshiftParserListener) EnterXconst(ctx *XconstContext) {}

// ExitXconst is called when production xconst is exited.
func (s *BaseRedshiftParserListener) ExitXconst(ctx *XconstContext) {}

// EnterBconst is called when production bconst is entered.
func (s *BaseRedshiftParserListener) EnterBconst(ctx *BconstContext) {}

// ExitBconst is called when production bconst is exited.
func (s *BaseRedshiftParserListener) ExitBconst(ctx *BconstContext) {}

// EnterFconst is called when production fconst is entered.
func (s *BaseRedshiftParserListener) EnterFconst(ctx *FconstContext) {}

// ExitFconst is called when production fconst is exited.
func (s *BaseRedshiftParserListener) ExitFconst(ctx *FconstContext) {}

// EnterIconst is called when production iconst is entered.
func (s *BaseRedshiftParserListener) EnterIconst(ctx *IconstContext) {}

// ExitIconst is called when production iconst is exited.
func (s *BaseRedshiftParserListener) ExitIconst(ctx *IconstContext) {}

// EnterSconst is called when production sconst is entered.
func (s *BaseRedshiftParserListener) EnterSconst(ctx *SconstContext) {}

// ExitSconst is called when production sconst is exited.
func (s *BaseRedshiftParserListener) ExitSconst(ctx *SconstContext) {}

// EnterAnysconst is called when production anysconst is entered.
func (s *BaseRedshiftParserListener) EnterAnysconst(ctx *AnysconstContext) {}

// ExitAnysconst is called when production anysconst is exited.
func (s *BaseRedshiftParserListener) ExitAnysconst(ctx *AnysconstContext) {}

// EnterOpt_uescape is called when production opt_uescape is entered.
func (s *BaseRedshiftParserListener) EnterOpt_uescape(ctx *Opt_uescapeContext) {}

// ExitOpt_uescape is called when production opt_uescape is exited.
func (s *BaseRedshiftParserListener) ExitOpt_uescape(ctx *Opt_uescapeContext) {}

// EnterSignediconst is called when production signediconst is entered.
func (s *BaseRedshiftParserListener) EnterSignediconst(ctx *SignediconstContext) {}

// ExitSignediconst is called when production signediconst is exited.
func (s *BaseRedshiftParserListener) ExitSignediconst(ctx *SignediconstContext) {}

// EnterRoleid is called when production roleid is entered.
func (s *BaseRedshiftParserListener) EnterRoleid(ctx *RoleidContext) {}

// ExitRoleid is called when production roleid is exited.
func (s *BaseRedshiftParserListener) ExitRoleid(ctx *RoleidContext) {}

// EnterRolespec is called when production rolespec is entered.
func (s *BaseRedshiftParserListener) EnterRolespec(ctx *RolespecContext) {}

// ExitRolespec is called when production rolespec is exited.
func (s *BaseRedshiftParserListener) ExitRolespec(ctx *RolespecContext) {}

// EnterRole_list is called when production role_list is entered.
func (s *BaseRedshiftParserListener) EnterRole_list(ctx *Role_listContext) {}

// ExitRole_list is called when production role_list is exited.
func (s *BaseRedshiftParserListener) ExitRole_list(ctx *Role_listContext) {}

// EnterColid is called when production colid is entered.
func (s *BaseRedshiftParserListener) EnterColid(ctx *ColidContext) {}

// ExitColid is called when production colid is exited.
func (s *BaseRedshiftParserListener) ExitColid(ctx *ColidContext) {}

// EnterTable_alias is called when production table_alias is entered.
func (s *BaseRedshiftParserListener) EnterTable_alias(ctx *Table_aliasContext) {}

// ExitTable_alias is called when production table_alias is exited.
func (s *BaseRedshiftParserListener) ExitTable_alias(ctx *Table_aliasContext) {}

// EnterType_function_name is called when production type_function_name is entered.
func (s *BaseRedshiftParserListener) EnterType_function_name(ctx *Type_function_nameContext) {}

// ExitType_function_name is called when production type_function_name is exited.
func (s *BaseRedshiftParserListener) ExitType_function_name(ctx *Type_function_nameContext) {}

// EnterNonreservedword is called when production nonreservedword is entered.
func (s *BaseRedshiftParserListener) EnterNonreservedword(ctx *NonreservedwordContext) {}

// ExitNonreservedword is called when production nonreservedword is exited.
func (s *BaseRedshiftParserListener) ExitNonreservedword(ctx *NonreservedwordContext) {}

// EnterCollabel is called when production collabel is entered.
func (s *BaseRedshiftParserListener) EnterCollabel(ctx *CollabelContext) {}

// ExitCollabel is called when production collabel is exited.
func (s *BaseRedshiftParserListener) ExitCollabel(ctx *CollabelContext) {}

// EnterIdentifier is called when production identifier is entered.
func (s *BaseRedshiftParserListener) EnterIdentifier(ctx *IdentifierContext) {}

// ExitIdentifier is called when production identifier is exited.
func (s *BaseRedshiftParserListener) ExitIdentifier(ctx *IdentifierContext) {}

// EnterPlsqlidentifier is called when production plsqlidentifier is entered.
func (s *BaseRedshiftParserListener) EnterPlsqlidentifier(ctx *PlsqlidentifierContext) {}

// ExitPlsqlidentifier is called when production plsqlidentifier is exited.
func (s *BaseRedshiftParserListener) ExitPlsqlidentifier(ctx *PlsqlidentifierContext) {}

// EnterUnreserved_keyword is called when production unreserved_keyword is entered.
func (s *BaseRedshiftParserListener) EnterUnreserved_keyword(ctx *Unreserved_keywordContext) {}

// ExitUnreserved_keyword is called when production unreserved_keyword is exited.
func (s *BaseRedshiftParserListener) ExitUnreserved_keyword(ctx *Unreserved_keywordContext) {}

// EnterCol_name_keyword is called when production col_name_keyword is entered.
func (s *BaseRedshiftParserListener) EnterCol_name_keyword(ctx *Col_name_keywordContext) {}

// ExitCol_name_keyword is called when production col_name_keyword is exited.
func (s *BaseRedshiftParserListener) ExitCol_name_keyword(ctx *Col_name_keywordContext) {}

// EnterType_func_name_keyword is called when production type_func_name_keyword is entered.
func (s *BaseRedshiftParserListener) EnterType_func_name_keyword(ctx *Type_func_name_keywordContext) {
}

// ExitType_func_name_keyword is called when production type_func_name_keyword is exited.
func (s *BaseRedshiftParserListener) ExitType_func_name_keyword(ctx *Type_func_name_keywordContext) {}

// EnterReserved_keyword is called when production reserved_keyword is entered.
func (s *BaseRedshiftParserListener) EnterReserved_keyword(ctx *Reserved_keywordContext) {}

// ExitReserved_keyword is called when production reserved_keyword is exited.
func (s *BaseRedshiftParserListener) ExitReserved_keyword(ctx *Reserved_keywordContext) {}

// EnterBuiltin_function_name is called when production builtin_function_name is entered.
func (s *BaseRedshiftParserListener) EnterBuiltin_function_name(ctx *Builtin_function_nameContext) {}

// ExitBuiltin_function_name is called when production builtin_function_name is exited.
func (s *BaseRedshiftParserListener) ExitBuiltin_function_name(ctx *Builtin_function_nameContext) {}

// EnterPl_function is called when production pl_function is entered.
func (s *BaseRedshiftParserListener) EnterPl_function(ctx *Pl_functionContext) {}

// ExitPl_function is called when production pl_function is exited.
func (s *BaseRedshiftParserListener) ExitPl_function(ctx *Pl_functionContext) {}

// EnterComp_options is called when production comp_options is entered.
func (s *BaseRedshiftParserListener) EnterComp_options(ctx *Comp_optionsContext) {}

// ExitComp_options is called when production comp_options is exited.
func (s *BaseRedshiftParserListener) ExitComp_options(ctx *Comp_optionsContext) {}

// EnterComp_option is called when production comp_option is entered.
func (s *BaseRedshiftParserListener) EnterComp_option(ctx *Comp_optionContext) {}

// ExitComp_option is called when production comp_option is exited.
func (s *BaseRedshiftParserListener) ExitComp_option(ctx *Comp_optionContext) {}

// EnterSharp is called when production sharp is entered.
func (s *BaseRedshiftParserListener) EnterSharp(ctx *SharpContext) {}

// ExitSharp is called when production sharp is exited.
func (s *BaseRedshiftParserListener) ExitSharp(ctx *SharpContext) {}

// EnterOption_value is called when production option_value is entered.
func (s *BaseRedshiftParserListener) EnterOption_value(ctx *Option_valueContext) {}

// ExitOption_value is called when production option_value is exited.
func (s *BaseRedshiftParserListener) ExitOption_value(ctx *Option_valueContext) {}

// EnterOpt_semi is called when production opt_semi is entered.
func (s *BaseRedshiftParserListener) EnterOpt_semi(ctx *Opt_semiContext) {}

// ExitOpt_semi is called when production opt_semi is exited.
func (s *BaseRedshiftParserListener) ExitOpt_semi(ctx *Opt_semiContext) {}

// EnterPl_block is called when production pl_block is entered.
func (s *BaseRedshiftParserListener) EnterPl_block(ctx *Pl_blockContext) {}

// ExitPl_block is called when production pl_block is exited.
func (s *BaseRedshiftParserListener) ExitPl_block(ctx *Pl_blockContext) {}

// EnterDecl_sect is called when production decl_sect is entered.
func (s *BaseRedshiftParserListener) EnterDecl_sect(ctx *Decl_sectContext) {}

// ExitDecl_sect is called when production decl_sect is exited.
func (s *BaseRedshiftParserListener) ExitDecl_sect(ctx *Decl_sectContext) {}

// EnterDecl_start is called when production decl_start is entered.
func (s *BaseRedshiftParserListener) EnterDecl_start(ctx *Decl_startContext) {}

// ExitDecl_start is called when production decl_start is exited.
func (s *BaseRedshiftParserListener) ExitDecl_start(ctx *Decl_startContext) {}

// EnterDecl_stmts is called when production decl_stmts is entered.
func (s *BaseRedshiftParserListener) EnterDecl_stmts(ctx *Decl_stmtsContext) {}

// ExitDecl_stmts is called when production decl_stmts is exited.
func (s *BaseRedshiftParserListener) ExitDecl_stmts(ctx *Decl_stmtsContext) {}

// EnterLabel_decl is called when production label_decl is entered.
func (s *BaseRedshiftParserListener) EnterLabel_decl(ctx *Label_declContext) {}

// ExitLabel_decl is called when production label_decl is exited.
func (s *BaseRedshiftParserListener) ExitLabel_decl(ctx *Label_declContext) {}

// EnterDecl_stmt is called when production decl_stmt is entered.
func (s *BaseRedshiftParserListener) EnterDecl_stmt(ctx *Decl_stmtContext) {}

// ExitDecl_stmt is called when production decl_stmt is exited.
func (s *BaseRedshiftParserListener) ExitDecl_stmt(ctx *Decl_stmtContext) {}

// EnterDecl_statement is called when production decl_statement is entered.
func (s *BaseRedshiftParserListener) EnterDecl_statement(ctx *Decl_statementContext) {}

// ExitDecl_statement is called when production decl_statement is exited.
func (s *BaseRedshiftParserListener) ExitDecl_statement(ctx *Decl_statementContext) {}

// EnterOpt_scrollable is called when production opt_scrollable is entered.
func (s *BaseRedshiftParserListener) EnterOpt_scrollable(ctx *Opt_scrollableContext) {}

// ExitOpt_scrollable is called when production opt_scrollable is exited.
func (s *BaseRedshiftParserListener) ExitOpt_scrollable(ctx *Opt_scrollableContext) {}

// EnterDecl_cursor_query is called when production decl_cursor_query is entered.
func (s *BaseRedshiftParserListener) EnterDecl_cursor_query(ctx *Decl_cursor_queryContext) {}

// ExitDecl_cursor_query is called when production decl_cursor_query is exited.
func (s *BaseRedshiftParserListener) ExitDecl_cursor_query(ctx *Decl_cursor_queryContext) {}

// EnterDecl_cursor_args is called when production decl_cursor_args is entered.
func (s *BaseRedshiftParserListener) EnterDecl_cursor_args(ctx *Decl_cursor_argsContext) {}

// ExitDecl_cursor_args is called when production decl_cursor_args is exited.
func (s *BaseRedshiftParserListener) ExitDecl_cursor_args(ctx *Decl_cursor_argsContext) {}

// EnterDecl_cursor_arglist is called when production decl_cursor_arglist is entered.
func (s *BaseRedshiftParserListener) EnterDecl_cursor_arglist(ctx *Decl_cursor_arglistContext) {}

// ExitDecl_cursor_arglist is called when production decl_cursor_arglist is exited.
func (s *BaseRedshiftParserListener) ExitDecl_cursor_arglist(ctx *Decl_cursor_arglistContext) {}

// EnterDecl_cursor_arg is called when production decl_cursor_arg is entered.
func (s *BaseRedshiftParserListener) EnterDecl_cursor_arg(ctx *Decl_cursor_argContext) {}

// ExitDecl_cursor_arg is called when production decl_cursor_arg is exited.
func (s *BaseRedshiftParserListener) ExitDecl_cursor_arg(ctx *Decl_cursor_argContext) {}

// EnterDecl_is_for is called when production decl_is_for is entered.
func (s *BaseRedshiftParserListener) EnterDecl_is_for(ctx *Decl_is_forContext) {}

// ExitDecl_is_for is called when production decl_is_for is exited.
func (s *BaseRedshiftParserListener) ExitDecl_is_for(ctx *Decl_is_forContext) {}

// EnterDecl_aliasitem is called when production decl_aliasitem is entered.
func (s *BaseRedshiftParserListener) EnterDecl_aliasitem(ctx *Decl_aliasitemContext) {}

// ExitDecl_aliasitem is called when production decl_aliasitem is exited.
func (s *BaseRedshiftParserListener) ExitDecl_aliasitem(ctx *Decl_aliasitemContext) {}

// EnterDecl_varname is called when production decl_varname is entered.
func (s *BaseRedshiftParserListener) EnterDecl_varname(ctx *Decl_varnameContext) {}

// ExitDecl_varname is called when production decl_varname is exited.
func (s *BaseRedshiftParserListener) ExitDecl_varname(ctx *Decl_varnameContext) {}

// EnterDecl_const is called when production decl_const is entered.
func (s *BaseRedshiftParserListener) EnterDecl_const(ctx *Decl_constContext) {}

// ExitDecl_const is called when production decl_const is exited.
func (s *BaseRedshiftParserListener) ExitDecl_const(ctx *Decl_constContext) {}

// EnterDecl_datatype is called when production decl_datatype is entered.
func (s *BaseRedshiftParserListener) EnterDecl_datatype(ctx *Decl_datatypeContext) {}

// ExitDecl_datatype is called when production decl_datatype is exited.
func (s *BaseRedshiftParserListener) ExitDecl_datatype(ctx *Decl_datatypeContext) {}

// EnterDecl_collate is called when production decl_collate is entered.
func (s *BaseRedshiftParserListener) EnterDecl_collate(ctx *Decl_collateContext) {}

// ExitDecl_collate is called when production decl_collate is exited.
func (s *BaseRedshiftParserListener) ExitDecl_collate(ctx *Decl_collateContext) {}

// EnterDecl_notnull is called when production decl_notnull is entered.
func (s *BaseRedshiftParserListener) EnterDecl_notnull(ctx *Decl_notnullContext) {}

// ExitDecl_notnull is called when production decl_notnull is exited.
func (s *BaseRedshiftParserListener) ExitDecl_notnull(ctx *Decl_notnullContext) {}

// EnterDecl_defval is called when production decl_defval is entered.
func (s *BaseRedshiftParserListener) EnterDecl_defval(ctx *Decl_defvalContext) {}

// ExitDecl_defval is called when production decl_defval is exited.
func (s *BaseRedshiftParserListener) ExitDecl_defval(ctx *Decl_defvalContext) {}

// EnterDecl_defkey is called when production decl_defkey is entered.
func (s *BaseRedshiftParserListener) EnterDecl_defkey(ctx *Decl_defkeyContext) {}

// ExitDecl_defkey is called when production decl_defkey is exited.
func (s *BaseRedshiftParserListener) ExitDecl_defkey(ctx *Decl_defkeyContext) {}

// EnterAssign_operator is called when production assign_operator is entered.
func (s *BaseRedshiftParserListener) EnterAssign_operator(ctx *Assign_operatorContext) {}

// ExitAssign_operator is called when production assign_operator is exited.
func (s *BaseRedshiftParserListener) ExitAssign_operator(ctx *Assign_operatorContext) {}

// EnterProc_sect is called when production proc_sect is entered.
func (s *BaseRedshiftParserListener) EnterProc_sect(ctx *Proc_sectContext) {}

// ExitProc_sect is called when production proc_sect is exited.
func (s *BaseRedshiftParserListener) ExitProc_sect(ctx *Proc_sectContext) {}

// EnterProc_stmt is called when production proc_stmt is entered.
func (s *BaseRedshiftParserListener) EnterProc_stmt(ctx *Proc_stmtContext) {}

// ExitProc_stmt is called when production proc_stmt is exited.
func (s *BaseRedshiftParserListener) ExitProc_stmt(ctx *Proc_stmtContext) {}

// EnterStmt_perform is called when production stmt_perform is entered.
func (s *BaseRedshiftParserListener) EnterStmt_perform(ctx *Stmt_performContext) {}

// ExitStmt_perform is called when production stmt_perform is exited.
func (s *BaseRedshiftParserListener) ExitStmt_perform(ctx *Stmt_performContext) {}

// EnterStmt_call is called when production stmt_call is entered.
func (s *BaseRedshiftParserListener) EnterStmt_call(ctx *Stmt_callContext) {}

// ExitStmt_call is called when production stmt_call is exited.
func (s *BaseRedshiftParserListener) ExitStmt_call(ctx *Stmt_callContext) {}

// EnterOpt_expr_list is called when production opt_expr_list is entered.
func (s *BaseRedshiftParserListener) EnterOpt_expr_list(ctx *Opt_expr_listContext) {}

// ExitOpt_expr_list is called when production opt_expr_list is exited.
func (s *BaseRedshiftParserListener) ExitOpt_expr_list(ctx *Opt_expr_listContext) {}

// EnterStmt_assign is called when production stmt_assign is entered.
func (s *BaseRedshiftParserListener) EnterStmt_assign(ctx *Stmt_assignContext) {}

// ExitStmt_assign is called when production stmt_assign is exited.
func (s *BaseRedshiftParserListener) ExitStmt_assign(ctx *Stmt_assignContext) {}

// EnterStmt_getdiag is called when production stmt_getdiag is entered.
func (s *BaseRedshiftParserListener) EnterStmt_getdiag(ctx *Stmt_getdiagContext) {}

// ExitStmt_getdiag is called when production stmt_getdiag is exited.
func (s *BaseRedshiftParserListener) ExitStmt_getdiag(ctx *Stmt_getdiagContext) {}

// EnterGetdiag_area_opt is called when production getdiag_area_opt is entered.
func (s *BaseRedshiftParserListener) EnterGetdiag_area_opt(ctx *Getdiag_area_optContext) {}

// ExitGetdiag_area_opt is called when production getdiag_area_opt is exited.
func (s *BaseRedshiftParserListener) ExitGetdiag_area_opt(ctx *Getdiag_area_optContext) {}

// EnterGetdiag_list is called when production getdiag_list is entered.
func (s *BaseRedshiftParserListener) EnterGetdiag_list(ctx *Getdiag_listContext) {}

// ExitGetdiag_list is called when production getdiag_list is exited.
func (s *BaseRedshiftParserListener) ExitGetdiag_list(ctx *Getdiag_listContext) {}

// EnterGetdiag_list_item is called when production getdiag_list_item is entered.
func (s *BaseRedshiftParserListener) EnterGetdiag_list_item(ctx *Getdiag_list_itemContext) {}

// ExitGetdiag_list_item is called when production getdiag_list_item is exited.
func (s *BaseRedshiftParserListener) ExitGetdiag_list_item(ctx *Getdiag_list_itemContext) {}

// EnterGetdiag_item is called when production getdiag_item is entered.
func (s *BaseRedshiftParserListener) EnterGetdiag_item(ctx *Getdiag_itemContext) {}

// ExitGetdiag_item is called when production getdiag_item is exited.
func (s *BaseRedshiftParserListener) ExitGetdiag_item(ctx *Getdiag_itemContext) {}

// EnterGetdiag_target is called when production getdiag_target is entered.
func (s *BaseRedshiftParserListener) EnterGetdiag_target(ctx *Getdiag_targetContext) {}

// ExitGetdiag_target is called when production getdiag_target is exited.
func (s *BaseRedshiftParserListener) ExitGetdiag_target(ctx *Getdiag_targetContext) {}

// EnterAssign_var is called when production assign_var is entered.
func (s *BaseRedshiftParserListener) EnterAssign_var(ctx *Assign_varContext) {}

// ExitAssign_var is called when production assign_var is exited.
func (s *BaseRedshiftParserListener) ExitAssign_var(ctx *Assign_varContext) {}

// EnterStmt_if is called when production stmt_if is entered.
func (s *BaseRedshiftParserListener) EnterStmt_if(ctx *Stmt_ifContext) {}

// ExitStmt_if is called when production stmt_if is exited.
func (s *BaseRedshiftParserListener) ExitStmt_if(ctx *Stmt_ifContext) {}

// EnterStmt_elsifs is called when production stmt_elsifs is entered.
func (s *BaseRedshiftParserListener) EnterStmt_elsifs(ctx *Stmt_elsifsContext) {}

// ExitStmt_elsifs is called when production stmt_elsifs is exited.
func (s *BaseRedshiftParserListener) ExitStmt_elsifs(ctx *Stmt_elsifsContext) {}

// EnterStmt_else is called when production stmt_else is entered.
func (s *BaseRedshiftParserListener) EnterStmt_else(ctx *Stmt_elseContext) {}

// ExitStmt_else is called when production stmt_else is exited.
func (s *BaseRedshiftParserListener) ExitStmt_else(ctx *Stmt_elseContext) {}

// EnterStmt_case is called when production stmt_case is entered.
func (s *BaseRedshiftParserListener) EnterStmt_case(ctx *Stmt_caseContext) {}

// ExitStmt_case is called when production stmt_case is exited.
func (s *BaseRedshiftParserListener) ExitStmt_case(ctx *Stmt_caseContext) {}

// EnterOpt_expr_until_when is called when production opt_expr_until_when is entered.
func (s *BaseRedshiftParserListener) EnterOpt_expr_until_when(ctx *Opt_expr_until_whenContext) {}

// ExitOpt_expr_until_when is called when production opt_expr_until_when is exited.
func (s *BaseRedshiftParserListener) ExitOpt_expr_until_when(ctx *Opt_expr_until_whenContext) {}

// EnterCase_when_list is called when production case_when_list is entered.
func (s *BaseRedshiftParserListener) EnterCase_when_list(ctx *Case_when_listContext) {}

// ExitCase_when_list is called when production case_when_list is exited.
func (s *BaseRedshiftParserListener) ExitCase_when_list(ctx *Case_when_listContext) {}

// EnterCase_when is called when production case_when is entered.
func (s *BaseRedshiftParserListener) EnterCase_when(ctx *Case_whenContext) {}

// ExitCase_when is called when production case_when is exited.
func (s *BaseRedshiftParserListener) ExitCase_when(ctx *Case_whenContext) {}

// EnterOpt_case_else is called when production opt_case_else is entered.
func (s *BaseRedshiftParserListener) EnterOpt_case_else(ctx *Opt_case_elseContext) {}

// ExitOpt_case_else is called when production opt_case_else is exited.
func (s *BaseRedshiftParserListener) ExitOpt_case_else(ctx *Opt_case_elseContext) {}

// EnterStmt_loop is called when production stmt_loop is entered.
func (s *BaseRedshiftParserListener) EnterStmt_loop(ctx *Stmt_loopContext) {}

// ExitStmt_loop is called when production stmt_loop is exited.
func (s *BaseRedshiftParserListener) ExitStmt_loop(ctx *Stmt_loopContext) {}

// EnterStmt_while is called when production stmt_while is entered.
func (s *BaseRedshiftParserListener) EnterStmt_while(ctx *Stmt_whileContext) {}

// ExitStmt_while is called when production stmt_while is exited.
func (s *BaseRedshiftParserListener) ExitStmt_while(ctx *Stmt_whileContext) {}

// EnterStmt_for is called when production stmt_for is entered.
func (s *BaseRedshiftParserListener) EnterStmt_for(ctx *Stmt_forContext) {}

// ExitStmt_for is called when production stmt_for is exited.
func (s *BaseRedshiftParserListener) ExitStmt_for(ctx *Stmt_forContext) {}

// EnterFor_control is called when production for_control is entered.
func (s *BaseRedshiftParserListener) EnterFor_control(ctx *For_controlContext) {}

// ExitFor_control is called when production for_control is exited.
func (s *BaseRedshiftParserListener) ExitFor_control(ctx *For_controlContext) {}

// EnterOpt_for_using_expression is called when production opt_for_using_expression is entered.
func (s *BaseRedshiftParserListener) EnterOpt_for_using_expression(ctx *Opt_for_using_expressionContext) {
}

// ExitOpt_for_using_expression is called when production opt_for_using_expression is exited.
func (s *BaseRedshiftParserListener) ExitOpt_for_using_expression(ctx *Opt_for_using_expressionContext) {
}

// EnterOpt_cursor_parameters is called when production opt_cursor_parameters is entered.
func (s *BaseRedshiftParserListener) EnterOpt_cursor_parameters(ctx *Opt_cursor_parametersContext) {}

// ExitOpt_cursor_parameters is called when production opt_cursor_parameters is exited.
func (s *BaseRedshiftParserListener) ExitOpt_cursor_parameters(ctx *Opt_cursor_parametersContext) {}

// EnterOpt_reverse is called when production opt_reverse is entered.
func (s *BaseRedshiftParserListener) EnterOpt_reverse(ctx *Opt_reverseContext) {}

// ExitOpt_reverse is called when production opt_reverse is exited.
func (s *BaseRedshiftParserListener) ExitOpt_reverse(ctx *Opt_reverseContext) {}

// EnterOpt_by_expression is called when production opt_by_expression is entered.
func (s *BaseRedshiftParserListener) EnterOpt_by_expression(ctx *Opt_by_expressionContext) {}

// ExitOpt_by_expression is called when production opt_by_expression is exited.
func (s *BaseRedshiftParserListener) ExitOpt_by_expression(ctx *Opt_by_expressionContext) {}

// EnterFor_variable is called when production for_variable is entered.
func (s *BaseRedshiftParserListener) EnterFor_variable(ctx *For_variableContext) {}

// ExitFor_variable is called when production for_variable is exited.
func (s *BaseRedshiftParserListener) ExitFor_variable(ctx *For_variableContext) {}

// EnterStmt_foreach_a is called when production stmt_foreach_a is entered.
func (s *BaseRedshiftParserListener) EnterStmt_foreach_a(ctx *Stmt_foreach_aContext) {}

// ExitStmt_foreach_a is called when production stmt_foreach_a is exited.
func (s *BaseRedshiftParserListener) ExitStmt_foreach_a(ctx *Stmt_foreach_aContext) {}

// EnterForeach_slice is called when production foreach_slice is entered.
func (s *BaseRedshiftParserListener) EnterForeach_slice(ctx *Foreach_sliceContext) {}

// ExitForeach_slice is called when production foreach_slice is exited.
func (s *BaseRedshiftParserListener) ExitForeach_slice(ctx *Foreach_sliceContext) {}

// EnterStmt_exit is called when production stmt_exit is entered.
func (s *BaseRedshiftParserListener) EnterStmt_exit(ctx *Stmt_exitContext) {}

// ExitStmt_exit is called when production stmt_exit is exited.
func (s *BaseRedshiftParserListener) ExitStmt_exit(ctx *Stmt_exitContext) {}

// EnterExit_type is called when production exit_type is entered.
func (s *BaseRedshiftParserListener) EnterExit_type(ctx *Exit_typeContext) {}

// ExitExit_type is called when production exit_type is exited.
func (s *BaseRedshiftParserListener) ExitExit_type(ctx *Exit_typeContext) {}

// EnterStmt_return is called when production stmt_return is entered.
func (s *BaseRedshiftParserListener) EnterStmt_return(ctx *Stmt_returnContext) {}

// ExitStmt_return is called when production stmt_return is exited.
func (s *BaseRedshiftParserListener) ExitStmt_return(ctx *Stmt_returnContext) {}

// EnterOpt_return_result is called when production opt_return_result is entered.
func (s *BaseRedshiftParserListener) EnterOpt_return_result(ctx *Opt_return_resultContext) {}

// ExitOpt_return_result is called when production opt_return_result is exited.
func (s *BaseRedshiftParserListener) ExitOpt_return_result(ctx *Opt_return_resultContext) {}

// EnterStmt_raise is called when production stmt_raise is entered.
func (s *BaseRedshiftParserListener) EnterStmt_raise(ctx *Stmt_raiseContext) {}

// ExitStmt_raise is called when production stmt_raise is exited.
func (s *BaseRedshiftParserListener) ExitStmt_raise(ctx *Stmt_raiseContext) {}

// EnterOpt_stmt_raise_level is called when production opt_stmt_raise_level is entered.
func (s *BaseRedshiftParserListener) EnterOpt_stmt_raise_level(ctx *Opt_stmt_raise_levelContext) {}

// ExitOpt_stmt_raise_level is called when production opt_stmt_raise_level is exited.
func (s *BaseRedshiftParserListener) ExitOpt_stmt_raise_level(ctx *Opt_stmt_raise_levelContext) {}

// EnterOpt_raise_list is called when production opt_raise_list is entered.
func (s *BaseRedshiftParserListener) EnterOpt_raise_list(ctx *Opt_raise_listContext) {}

// ExitOpt_raise_list is called when production opt_raise_list is exited.
func (s *BaseRedshiftParserListener) ExitOpt_raise_list(ctx *Opt_raise_listContext) {}

// EnterOpt_raise_using is called when production opt_raise_using is entered.
func (s *BaseRedshiftParserListener) EnterOpt_raise_using(ctx *Opt_raise_usingContext) {}

// ExitOpt_raise_using is called when production opt_raise_using is exited.
func (s *BaseRedshiftParserListener) ExitOpt_raise_using(ctx *Opt_raise_usingContext) {}

// EnterOpt_raise_using_elem is called when production opt_raise_using_elem is entered.
func (s *BaseRedshiftParserListener) EnterOpt_raise_using_elem(ctx *Opt_raise_using_elemContext) {}

// ExitOpt_raise_using_elem is called when production opt_raise_using_elem is exited.
func (s *BaseRedshiftParserListener) ExitOpt_raise_using_elem(ctx *Opt_raise_using_elemContext) {}

// EnterOpt_raise_using_elem_list is called when production opt_raise_using_elem_list is entered.
func (s *BaseRedshiftParserListener) EnterOpt_raise_using_elem_list(ctx *Opt_raise_using_elem_listContext) {
}

// ExitOpt_raise_using_elem_list is called when production opt_raise_using_elem_list is exited.
func (s *BaseRedshiftParserListener) ExitOpt_raise_using_elem_list(ctx *Opt_raise_using_elem_listContext) {
}

// EnterStmt_assert is called when production stmt_assert is entered.
func (s *BaseRedshiftParserListener) EnterStmt_assert(ctx *Stmt_assertContext) {}

// ExitStmt_assert is called when production stmt_assert is exited.
func (s *BaseRedshiftParserListener) ExitStmt_assert(ctx *Stmt_assertContext) {}

// EnterOpt_stmt_assert_message is called when production opt_stmt_assert_message is entered.
func (s *BaseRedshiftParserListener) EnterOpt_stmt_assert_message(ctx *Opt_stmt_assert_messageContext) {
}

// ExitOpt_stmt_assert_message is called when production opt_stmt_assert_message is exited.
func (s *BaseRedshiftParserListener) ExitOpt_stmt_assert_message(ctx *Opt_stmt_assert_messageContext) {
}

// EnterLoop_body is called when production loop_body is entered.
func (s *BaseRedshiftParserListener) EnterLoop_body(ctx *Loop_bodyContext) {}

// ExitLoop_body is called when production loop_body is exited.
func (s *BaseRedshiftParserListener) ExitLoop_body(ctx *Loop_bodyContext) {}

// EnterStmt_execsql is called when production stmt_execsql is entered.
func (s *BaseRedshiftParserListener) EnterStmt_execsql(ctx *Stmt_execsqlContext) {}

// ExitStmt_execsql is called when production stmt_execsql is exited.
func (s *BaseRedshiftParserListener) ExitStmt_execsql(ctx *Stmt_execsqlContext) {}

// EnterStmt_dynexecute is called when production stmt_dynexecute is entered.
func (s *BaseRedshiftParserListener) EnterStmt_dynexecute(ctx *Stmt_dynexecuteContext) {}

// ExitStmt_dynexecute is called when production stmt_dynexecute is exited.
func (s *BaseRedshiftParserListener) ExitStmt_dynexecute(ctx *Stmt_dynexecuteContext) {}

// EnterOpt_execute_using is called when production opt_execute_using is entered.
func (s *BaseRedshiftParserListener) EnterOpt_execute_using(ctx *Opt_execute_usingContext) {}

// ExitOpt_execute_using is called when production opt_execute_using is exited.
func (s *BaseRedshiftParserListener) ExitOpt_execute_using(ctx *Opt_execute_usingContext) {}

// EnterOpt_execute_using_list is called when production opt_execute_using_list is entered.
func (s *BaseRedshiftParserListener) EnterOpt_execute_using_list(ctx *Opt_execute_using_listContext) {
}

// ExitOpt_execute_using_list is called when production opt_execute_using_list is exited.
func (s *BaseRedshiftParserListener) ExitOpt_execute_using_list(ctx *Opt_execute_using_listContext) {}

// EnterOpt_execute_into is called when production opt_execute_into is entered.
func (s *BaseRedshiftParserListener) EnterOpt_execute_into(ctx *Opt_execute_intoContext) {}

// ExitOpt_execute_into is called when production opt_execute_into is exited.
func (s *BaseRedshiftParserListener) ExitOpt_execute_into(ctx *Opt_execute_intoContext) {}

// EnterStmt_open is called when production stmt_open is entered.
func (s *BaseRedshiftParserListener) EnterStmt_open(ctx *Stmt_openContext) {}

// ExitStmt_open is called when production stmt_open is exited.
func (s *BaseRedshiftParserListener) ExitStmt_open(ctx *Stmt_openContext) {}

// EnterOpt_open_bound_list_item is called when production opt_open_bound_list_item is entered.
func (s *BaseRedshiftParserListener) EnterOpt_open_bound_list_item(ctx *Opt_open_bound_list_itemContext) {
}

// ExitOpt_open_bound_list_item is called when production opt_open_bound_list_item is exited.
func (s *BaseRedshiftParserListener) ExitOpt_open_bound_list_item(ctx *Opt_open_bound_list_itemContext) {
}

// EnterOpt_open_bound_list is called when production opt_open_bound_list is entered.
func (s *BaseRedshiftParserListener) EnterOpt_open_bound_list(ctx *Opt_open_bound_listContext) {}

// ExitOpt_open_bound_list is called when production opt_open_bound_list is exited.
func (s *BaseRedshiftParserListener) ExitOpt_open_bound_list(ctx *Opt_open_bound_listContext) {}

// EnterOpt_open_using is called when production opt_open_using is entered.
func (s *BaseRedshiftParserListener) EnterOpt_open_using(ctx *Opt_open_usingContext) {}

// ExitOpt_open_using is called when production opt_open_using is exited.
func (s *BaseRedshiftParserListener) ExitOpt_open_using(ctx *Opt_open_usingContext) {}

// EnterOpt_scroll_option is called when production opt_scroll_option is entered.
func (s *BaseRedshiftParserListener) EnterOpt_scroll_option(ctx *Opt_scroll_optionContext) {}

// ExitOpt_scroll_option is called when production opt_scroll_option is exited.
func (s *BaseRedshiftParserListener) ExitOpt_scroll_option(ctx *Opt_scroll_optionContext) {}

// EnterOpt_scroll_option_no is called when production opt_scroll_option_no is entered.
func (s *BaseRedshiftParserListener) EnterOpt_scroll_option_no(ctx *Opt_scroll_option_noContext) {}

// ExitOpt_scroll_option_no is called when production opt_scroll_option_no is exited.
func (s *BaseRedshiftParserListener) ExitOpt_scroll_option_no(ctx *Opt_scroll_option_noContext) {}

// EnterStmt_fetch is called when production stmt_fetch is entered.
func (s *BaseRedshiftParserListener) EnterStmt_fetch(ctx *Stmt_fetchContext) {}

// ExitStmt_fetch is called when production stmt_fetch is exited.
func (s *BaseRedshiftParserListener) ExitStmt_fetch(ctx *Stmt_fetchContext) {}

// EnterInto_target is called when production into_target is entered.
func (s *BaseRedshiftParserListener) EnterInto_target(ctx *Into_targetContext) {}

// ExitInto_target is called when production into_target is exited.
func (s *BaseRedshiftParserListener) ExitInto_target(ctx *Into_targetContext) {}

// EnterOpt_cursor_from is called when production opt_cursor_from is entered.
func (s *BaseRedshiftParserListener) EnterOpt_cursor_from(ctx *Opt_cursor_fromContext) {}

// ExitOpt_cursor_from is called when production opt_cursor_from is exited.
func (s *BaseRedshiftParserListener) ExitOpt_cursor_from(ctx *Opt_cursor_fromContext) {}

// EnterOpt_fetch_direction is called when production opt_fetch_direction is entered.
func (s *BaseRedshiftParserListener) EnterOpt_fetch_direction(ctx *Opt_fetch_directionContext) {}

// ExitOpt_fetch_direction is called when production opt_fetch_direction is exited.
func (s *BaseRedshiftParserListener) ExitOpt_fetch_direction(ctx *Opt_fetch_directionContext) {}

// EnterStmt_move is called when production stmt_move is entered.
func (s *BaseRedshiftParserListener) EnterStmt_move(ctx *Stmt_moveContext) {}

// ExitStmt_move is called when production stmt_move is exited.
func (s *BaseRedshiftParserListener) ExitStmt_move(ctx *Stmt_moveContext) {}

// EnterStmt_close is called when production stmt_close is entered.
func (s *BaseRedshiftParserListener) EnterStmt_close(ctx *Stmt_closeContext) {}

// ExitStmt_close is called when production stmt_close is exited.
func (s *BaseRedshiftParserListener) ExitStmt_close(ctx *Stmt_closeContext) {}

// EnterStmt_null is called when production stmt_null is entered.
func (s *BaseRedshiftParserListener) EnterStmt_null(ctx *Stmt_nullContext) {}

// ExitStmt_null is called when production stmt_null is exited.
func (s *BaseRedshiftParserListener) ExitStmt_null(ctx *Stmt_nullContext) {}

// EnterStmt_commit is called when production stmt_commit is entered.
func (s *BaseRedshiftParserListener) EnterStmt_commit(ctx *Stmt_commitContext) {}

// ExitStmt_commit is called when production stmt_commit is exited.
func (s *BaseRedshiftParserListener) ExitStmt_commit(ctx *Stmt_commitContext) {}

// EnterStmt_rollback is called when production stmt_rollback is entered.
func (s *BaseRedshiftParserListener) EnterStmt_rollback(ctx *Stmt_rollbackContext) {}

// ExitStmt_rollback is called when production stmt_rollback is exited.
func (s *BaseRedshiftParserListener) ExitStmt_rollback(ctx *Stmt_rollbackContext) {}

// EnterPlsql_opt_transaction_chain is called when production plsql_opt_transaction_chain is entered.
func (s *BaseRedshiftParserListener) EnterPlsql_opt_transaction_chain(ctx *Plsql_opt_transaction_chainContext) {
}

// ExitPlsql_opt_transaction_chain is called when production plsql_opt_transaction_chain is exited.
func (s *BaseRedshiftParserListener) ExitPlsql_opt_transaction_chain(ctx *Plsql_opt_transaction_chainContext) {
}

// EnterStmt_set is called when production stmt_set is entered.
func (s *BaseRedshiftParserListener) EnterStmt_set(ctx *Stmt_setContext) {}

// ExitStmt_set is called when production stmt_set is exited.
func (s *BaseRedshiftParserListener) ExitStmt_set(ctx *Stmt_setContext) {}

// EnterCursor_variable is called when production cursor_variable is entered.
func (s *BaseRedshiftParserListener) EnterCursor_variable(ctx *Cursor_variableContext) {}

// ExitCursor_variable is called when production cursor_variable is exited.
func (s *BaseRedshiftParserListener) ExitCursor_variable(ctx *Cursor_variableContext) {}

// EnterException_sect is called when production exception_sect is entered.
func (s *BaseRedshiftParserListener) EnterException_sect(ctx *Exception_sectContext) {}

// ExitException_sect is called when production exception_sect is exited.
func (s *BaseRedshiftParserListener) ExitException_sect(ctx *Exception_sectContext) {}

// EnterProc_exceptions is called when production proc_exceptions is entered.
func (s *BaseRedshiftParserListener) EnterProc_exceptions(ctx *Proc_exceptionsContext) {}

// ExitProc_exceptions is called when production proc_exceptions is exited.
func (s *BaseRedshiftParserListener) ExitProc_exceptions(ctx *Proc_exceptionsContext) {}

// EnterProc_exception is called when production proc_exception is entered.
func (s *BaseRedshiftParserListener) EnterProc_exception(ctx *Proc_exceptionContext) {}

// ExitProc_exception is called when production proc_exception is exited.
func (s *BaseRedshiftParserListener) ExitProc_exception(ctx *Proc_exceptionContext) {}

// EnterProc_conditions is called when production proc_conditions is entered.
func (s *BaseRedshiftParserListener) EnterProc_conditions(ctx *Proc_conditionsContext) {}

// ExitProc_conditions is called when production proc_conditions is exited.
func (s *BaseRedshiftParserListener) ExitProc_conditions(ctx *Proc_conditionsContext) {}

// EnterProc_condition is called when production proc_condition is entered.
func (s *BaseRedshiftParserListener) EnterProc_condition(ctx *Proc_conditionContext) {}

// ExitProc_condition is called when production proc_condition is exited.
func (s *BaseRedshiftParserListener) ExitProc_condition(ctx *Proc_conditionContext) {}

// EnterOpt_block_label is called when production opt_block_label is entered.
func (s *BaseRedshiftParserListener) EnterOpt_block_label(ctx *Opt_block_labelContext) {}

// ExitOpt_block_label is called when production opt_block_label is exited.
func (s *BaseRedshiftParserListener) ExitOpt_block_label(ctx *Opt_block_labelContext) {}

// EnterOpt_loop_label is called when production opt_loop_label is entered.
func (s *BaseRedshiftParserListener) EnterOpt_loop_label(ctx *Opt_loop_labelContext) {}

// ExitOpt_loop_label is called when production opt_loop_label is exited.
func (s *BaseRedshiftParserListener) ExitOpt_loop_label(ctx *Opt_loop_labelContext) {}

// EnterOpt_label is called when production opt_label is entered.
func (s *BaseRedshiftParserListener) EnterOpt_label(ctx *Opt_labelContext) {}

// ExitOpt_label is called when production opt_label is exited.
func (s *BaseRedshiftParserListener) ExitOpt_label(ctx *Opt_labelContext) {}

// EnterOpt_exitcond is called when production opt_exitcond is entered.
func (s *BaseRedshiftParserListener) EnterOpt_exitcond(ctx *Opt_exitcondContext) {}

// ExitOpt_exitcond is called when production opt_exitcond is exited.
func (s *BaseRedshiftParserListener) ExitOpt_exitcond(ctx *Opt_exitcondContext) {}

// EnterAny_identifier is called when production any_identifier is entered.
func (s *BaseRedshiftParserListener) EnterAny_identifier(ctx *Any_identifierContext) {}

// ExitAny_identifier is called when production any_identifier is exited.
func (s *BaseRedshiftParserListener) ExitAny_identifier(ctx *Any_identifierContext) {}

// EnterPlsql_unreserved_keyword is called when production plsql_unreserved_keyword is entered.
func (s *BaseRedshiftParserListener) EnterPlsql_unreserved_keyword(ctx *Plsql_unreserved_keywordContext) {
}

// ExitPlsql_unreserved_keyword is called when production plsql_unreserved_keyword is exited.
func (s *BaseRedshiftParserListener) ExitPlsql_unreserved_keyword(ctx *Plsql_unreserved_keywordContext) {
}

// EnterSql_expression is called when production sql_expression is entered.
func (s *BaseRedshiftParserListener) EnterSql_expression(ctx *Sql_expressionContext) {}

// ExitSql_expression is called when production sql_expression is exited.
func (s *BaseRedshiftParserListener) ExitSql_expression(ctx *Sql_expressionContext) {}

// EnterExpr_until_then is called when production expr_until_then is entered.
func (s *BaseRedshiftParserListener) EnterExpr_until_then(ctx *Expr_until_thenContext) {}

// ExitExpr_until_then is called when production expr_until_then is exited.
func (s *BaseRedshiftParserListener) ExitExpr_until_then(ctx *Expr_until_thenContext) {}

// EnterExpr_until_semi is called when production expr_until_semi is entered.
func (s *BaseRedshiftParserListener) EnterExpr_until_semi(ctx *Expr_until_semiContext) {}

// ExitExpr_until_semi is called when production expr_until_semi is exited.
func (s *BaseRedshiftParserListener) ExitExpr_until_semi(ctx *Expr_until_semiContext) {}

// EnterExpr_until_rightbracket is called when production expr_until_rightbracket is entered.
func (s *BaseRedshiftParserListener) EnterExpr_until_rightbracket(ctx *Expr_until_rightbracketContext) {
}

// ExitExpr_until_rightbracket is called when production expr_until_rightbracket is exited.
func (s *BaseRedshiftParserListener) ExitExpr_until_rightbracket(ctx *Expr_until_rightbracketContext) {
}

// EnterExpr_until_loop is called when production expr_until_loop is entered.
func (s *BaseRedshiftParserListener) EnterExpr_until_loop(ctx *Expr_until_loopContext) {}

// ExitExpr_until_loop is called when production expr_until_loop is exited.
func (s *BaseRedshiftParserListener) ExitExpr_until_loop(ctx *Expr_until_loopContext) {}

// EnterMake_execsql_stmt is called when production make_execsql_stmt is entered.
func (s *BaseRedshiftParserListener) EnterMake_execsql_stmt(ctx *Make_execsql_stmtContext) {}

// ExitMake_execsql_stmt is called when production make_execsql_stmt is exited.
func (s *BaseRedshiftParserListener) ExitMake_execsql_stmt(ctx *Make_execsql_stmtContext) {}

// EnterOpt_returning_clause_into is called when production opt_returning_clause_into is entered.
func (s *BaseRedshiftParserListener) EnterOpt_returning_clause_into(ctx *Opt_returning_clause_intoContext) {
}

// ExitOpt_returning_clause_into is called when production opt_returning_clause_into is exited.
func (s *BaseRedshiftParserListener) ExitOpt_returning_clause_into(ctx *Opt_returning_clause_intoContext) {
}
