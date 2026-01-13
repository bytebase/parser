 /**
 * Oracle(c) PL/SQL 11g Parser
 *
 * Copyright (c) 2009-2011 Alexandre Porcelli <alexandre.porcelli@gmail.com>
 * Copyright (c) 2015-2019 Ivan Kochurkin (KvanTTT, kvanttt@gmail.com, Positive Technologies).
 * Copyright (c) 2017-2018 Mark Adams <madams51703@gmail.com>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

parser grammar PlSqlParser;

options {
    tokenVocab=PlSqlLexer;
    superClass=PlSqlParserBase;
}

@parser::postinclude {
#include <PlSqlParserBase.h>
}

sql_script
    : (
        (sql_plus_command | unit_statement) SEMICOLON?
        | SEMICOLON
      )* EOF
    ;

unit_statement
    : alter_analytic_view
    | alter_attribute_dimension
    | alter_audit_policy
    | alter_cluster
    | alter_database
    | alter_database_link
    | alter_dimension
    | alter_diskgroup
    | alter_flashback_archive
    | alter_function
    | alter_hierarchy
    | alter_index
    | alter_inmemory_join_group
    | alter_java
    | alter_library
    | alter_lockdown_profile
    | alter_materialized_view
    | alter_materialized_view_log
    | alter_materialized_zonemap
    | alter_operator
    | alter_outline
    | alter_package
    | alter_pmem_filestore
    | alter_procedure
    | alter_resource_cost
    | alter_role
    | alter_rollback_segment
    | alter_sequence
    | alter_session
    | alter_synonym
    | alter_table
    | alter_tablespace
    | alter_tablespace_set
    | alter_trigger
    | alter_type
    | alter_user
    | alter_view

    | create_analytic_view
    | create_attribute_dimension
    | create_audit_policy
    | create_cluster
    | create_context
    | create_controlfile
    | create_database
    | create_database_link
    | create_dimension
    | create_directory
    | create_diskgroup
    | create_edition
    | create_flashback_archive
    | create_function_body
    | create_hierarchy
    | create_index
    | create_inmemory_join_group
    | create_java
    | create_library
    | create_lockdown_profile
    | create_materialized_view
    | create_materialized_view_log
    | create_materialized_zonemap
    | create_operator
    | create_outline
    | create_package
    | create_package_body
    | create_pmem_filestore
    | create_procedure_body
    | create_profile
    | create_restore_point
    | create_role
    | create_rollback_segment
    | create_schema
    | create_sequence
    | create_spfile
    | create_synonym
    | create_table
    | create_tablespace
    | create_tablespace_set
    | create_trigger
    | create_type
    | create_user
    | create_view

    | drop_analytic_view
    | drop_attribute_dimension
    | drop_audit_policy
    | drop_cluster
    | drop_context
    | drop_database
    | drop_database_link
    | drop_directory
    | drop_diskgroup
    | drop_edition
    | drop_flashback_archive
    | drop_function
    | drop_hierarchy
    | drop_index
    | drop_indextype
    | drop_inmemory_join_group
    | drop_java
    | drop_library
    | drop_lockdown_profile
    | drop_materialized_view
    | drop_materialized_view_log
    | drop_materialized_zonemap
    | drop_operator
    | drop_outline
    | drop_package
    | drop_pmem_filestore
    | drop_procedure
    | drop_restore_point
    | drop_role
    | drop_rollback_segment
    | drop_sequence
    | drop_synonym
    | drop_table
    | drop_tablespace
    | drop_tablespace_set
    | drop_trigger
    | drop_type
    | drop_user
    | drop_view

    | administer_key_management
    | analyze
    | anonymous_block
    | associate_statistics
    | audit_traditional
    | comment_on_column
    | comment_on_materialized
    | comment_on_table
    | data_manipulation_language_statements
    | disassociate_statistics
    | flashback_table
    | grant_statement
    | noaudit_statement
    | purge_statement
    | rename_object
    | revoke_statement
    | transaction_control_statements
    | truncate_cluster
    | truncate_table
    | unified_auditing
    | sql_call_statement // put call statement after all statements. BYT-8268: SQL level requires CALL keyword
    ;

// https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/ALTER-DISKGROUP.html
alter_diskgroup
    : ALTER DISKGROUP ( id_expression (((add_disk_clause | drop_disk_clause)+ | resize_disk_clause) rebalance_diskgroup_clause?
                                      | ( replace_disk_clause
                                        | rename_disk_clause
                                        | disk_online_clause
                                        | disk_offline_clause
                                        | rebalance_diskgroup_clause
                                        | check_diskgroup_clause
                                        | diskgroup_template_clauses
                                        | diskgroup_directory_clauses
                                        | diskgroup_alias_clauses
                                        | diskgroup_volume_clauses
                                        | diskgroup_attributes
                                        | drop_diskgroup_file_clause
                                        | convert_redundancy_clause
                                        | usergroup_clauses
                                        | user_clauses
                                        | file_permissions_clause
                                        | file_owner_clause
                                        | scrub_clause
                                        | quotagroup_clauses
                                        | filegroup_clauses
                                        )
                                      )
                      | (id_expression (COMMA id_expression)* | ALL) (undrop_disk_clause | diskgroup_availability | enable_disable_volume)
                      )
    ;

add_disk_clause
    : ADD ( (SITE sn=id_expression)? quorum_regular? (FAILGROUP fgn=id_expression)? DISK qualified_disk_clause (COMMA qualified_disk_clause)*)+
    ;

drop_disk_clause
    : DROP ( quorum_regular? DISK id_expression force_noforce? (COMMA id_expression force_noforce?)*
           | DISKS IN quorum_regular? FAILGROUP id_expression force_noforce? (COMMA id_expression force_noforce?)*
           )
    ;

resize_disk_clause
    : RESIZE ALL (SIZE size_clause)?
    ;

replace_disk_clause
    : REPLACE DISK id_expression WITH CHAR_STRING force_noforce? (COMMA id_expression WITH CHAR_STRING force_noforce?)*
        (POWER numeric)?
        wait_nowait?
    ;

wait_nowait
    : WAIT
    | NOWAIT
    ;

rename_disk_clause
    : RENAME ( DISK id_expression TO id_expression (COMMA id_expression TO id_expression)*
             | DISKS ALL
             )
    ;

disk_online_clause
    : ONLINE ( (quorum_regular? DISK id_expression (COMMA id_expression)* | DISKS IN quorum_regular? FAILGROUP id_expression (COMMA id_expression)*)+
             | ALL
             )
        (POWER numeric)?
        wait_nowait?
    ;

disk_offline_clause
    : OFFLINE ( quorum_regular? DISK id_expression (COMMA id_expression)*
              | DISKS IN quorum_regular? FAILGROUP id_expression (COMMA id_expression)*
              ) timeout_clause?
    ;

timeout_clause
    : DROP AFTER numeric (M_LETTER | H_LETTER)
    ;

rebalance_diskgroup_clause
    : REBALANCE ( ((WITH | WITHOUT) phase+)? (POWER numeric) (WAIT | NOWAIT)?
                | MODIFY POWER numeric?
                )
    ;

phase
    : id_expression //TODO
    ;

check_diskgroup_clause
    : CHECK ALL? (REPAIR | NOREPAIR)? //inconsistent documentation
    ;

diskgroup_template_clauses
    : (ADD | MODIFY) TEMPLATE id_expression qualified_template_clause (COMMA id_expression qualified_template_clause)*
    | DROP TEMPLATE id_expression (COMMA id_expression)*
    ;

qualified_template_clause
    : ATTRIBUTES LEFT_PAREN redundancy_clause? striping_clause? RIGHT_PAREN //inconsistent documentation
    ;

redundancy_clause
    : MIRROR
    | HIGH
    | UNPROTECTED
    | PARITY
    | DOUBLE
    ;

striping_clause
    : FINE
    | COARSE
    ;

force_noforce
    : FORCE
    | NOFORCE
    ;

diskgroup_directory_clauses
    : ADD DIRECTORY filename (COMMA filename)*
    | DROP DIRECTORY filename force_noforce? (COMMAfilename force_noforce?)*
    | RENAME DIRECTORY dir_name TO dir_name (COMMA dir_name TO dir_name)*
    ;

dir_name
    : CHAR_STRING
    ;

diskgroup_alias_clauses
    : ADD ALIAS CHAR_STRING FOR CHAR_STRING (COMMA CHAR_STRING FOR CHAR_STRING)*
    | DROP ALIAS CHAR_STRING (COMMA CHAR_STRING)*
    | RENAME ALIAS CHAR_STRING TO CHAR_STRING (COMMA CHAR_STRING TO CHAR_STRING)*
    ;

diskgroup_volume_clauses
    : add_volume_clause
    | modify_volume_clause
    | RESIZE VOLUME id_expression SIZE size_clause
    | DROP VOLUME id_expression
    ;

add_volume_clause
    : ADD VOLUME id_expression SIZE size_clause redundancy_clause?
        (STRIPE_WIDTH numeric (K_LETTER | M_LETTER))?
        (STRIPE_COLUMNS numeric)?
    ;

modify_volume_clause
    : MODIFY VOLUME id_expression (MOUNTPATH CHAR_STRING)?
        (USAGE CHAR_STRING)?
    ;

diskgroup_attributes
    : SET ATTRIBUTE CHAR_STRING EQUALS_OP CHAR_STRING
    ;

drop_diskgroup_file_clause
    : DROP FILE filename (COMMA filename)*
    ;

convert_redundancy_clause
    : CONVERT REDUNDANCY TO FLEX
    ;

usergroup_clauses
    : ADD USERGROUP CHAR_STRING WITH MEMBER CHAR_STRING (COMMA CHAR_STRING)*
    | MODIFY USERGROUP CHAR_STRING (ADD | DROP) MEMBER CHAR_STRING (COMMA CHAR_STRING)*
    | DROP USERGROUP CHAR_STRING
    ;

user_clauses
    : ADD USER CHAR_STRING (COMMA CHAR_STRING)*
    | DROP USER CHAR_STRING (COMMA CHAR_STRING)* CASCADE?
    | REPLACE USER CHAR_STRING WITH CHAR_STRING (COMMA CHAR_STRING WITH CHAR_STRING)*
    ;

file_permissions_clause
    : SET PERMISSION (OWNER | GROUP | OTHER) EQUALS_OP (NONE | READ (ONLY | WRITE)) (COMMA (OWNER | GROUP | OTHER) EQUALS_OP (NONE | READ (ONLY | WRITE)))*
        FOR FILE CHAR_STRING (COMMA CHAR_STRING)*
    ;

file_owner_clause
    : SET OWNERSHIP (OWNER | GROUP) EQUALS_OP CHAR_STRING (COMMA (OWNER | GROUP) EQUALS_OP CHAR_STRING)* FOR FILE CHAR_STRING (COMMA CHAR_STRING)*
    ;

scrub_clause
    : SCRUB (FILE CHAR_STRING | DISK id_expression)?
        (REPAIR | NOREPAIR)?
        (POWER (AUTO | LOW | HIGH | MAX))?
        wait_nowait?
        force_noforce?
        STOP?
    ;

quotagroup_clauses
    : ADD QUOTAGROUP id_expression (SET property_name EQUALS_OP property_value)?
    | MODIFY QUOTAGROUP id_expression SET property_name EQUALS_OP property_value
    | MOVE QUOTAGROUP id_expression TO id_expression
    | DROP QUOTAGROUP id_expression
    ;

property_name
    : id_expression
    ;

property_value
    : id_expression
    ;

filegroup_clauses
    : add_filegroup_clause
    | modify_filegroup_clause
    | move_to_filegroup_clause
    | drop_filegroup_clause
    ;

add_filegroup_clause
    : ADD FILEGROUP id_expression ((DATABASE | CLUSTER | VOLUME) id_expression | TEMPLATE) (FROM TEMPLATE id_expression)?
        (SET CHAR_STRING EQUALS_OP CHAR_STRING)?
    ;

modify_filegroup_clause
    : MODIFY FILEGROUP id_expression
        SET CHAR_STRING EQUALS_OP CHAR_STRING
    ;

move_to_filegroup_clause
    : MOVE FILE CHAR_STRING TO FILEGROUP id_expression
    ;

drop_filegroup_clause
    : DROP FILEGROUP id_expression CASCADE?
    ;


quorum_regular
    : QUORUM
    | REGULAR
    ;

undrop_disk_clause
    : UNDROP DISKS
    ;

diskgroup_availability
    : MOUNT (RESTRICTED | NORMAL)? (FORCE | NOFORCE)?
    | DISMOUNT (FORCE | NOFORCE)?
    ;

enable_disable_volume
    : (ENABLE | DISABLE) VOLUME ( id_expression (COMMA id_expression)* | ALL)
    ;

// DDL -> SQL Statements for Stored PL/SQL Units

// Function DDLs

drop_function
    : DROP FUNCTION function_name SEMICOLON
    ;

// https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/ALTER-FLASHBACK-ARCHIVE.html
alter_flashback_archive
    : ALTER FLASHBACK ARCHIVE fa=id_expression
        ( SET DEFAULT
        | (ADD | MODIFY) TABLESPACE ts=id_expression flashback_archive_quota?
        | REMOVE TABLESPACE rts=id_expression
        | MODIFY /*RETENTION*/ flashback_archive_retention // inconsistent documentation
        | PURGE (ALL | BEFORE (SCN expression | TIMESTAMP expression))
        | NO? OPTIMIZE DATA
        )
    ;

// https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/ALTER-HIERARCHY.html
alter_hierarchy
    : ALTER HIERARCHY (schema_name PERIOD)? hn=id_expression (RENAME TO nhn=id_expression | COMPILE)
    ;

alter_function
    : ALTER FUNCTION function_name COMPILE DEBUG? compiler_parameters_clause* (REUSE SETTINGS)? SEMICOLON
    ;

// https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/ALTER-JAVA.html
alter_java
    : ALTER JAVA (SOURCE | CLASS) (schema_name PERIOD)? o=id_expression
        (RESOLVER LEFT_PAREN (LEFT_PAREN match_string COMMA? (schema_name | '-') RIGHT_PAREN)+ RIGHT_PAREN)?
        (COMPILE | RESOLVE | invoker_rights_clause)
    ;

match_string
    : DELIMITED_ID
    | '*'
    ;

create_function_body
    : CREATE (OR REPLACE)? FUNCTION function_name (LEFT_PAREN parameter (COMMA parameter)* RIGHT_PAREN)?
      RETURN type_spec (invoker_rights_clause | accessible_by_clause | parallel_enable_clause | result_cache_clause | DETERMINISTIC | default_collation_clause | annotations_clause)*
      ((PIPELINED? (IS | AS) (DECLARE? seq_of_declare_specs? body | call_spec))
        | pipelined_using_clause
        | AGGREGATE USING implementation_type_name
        | sql_macro_body
      ) SEMICOLON
    ;

sql_macro_body
    : SQL_MACRO IS BEGIN RETURN quoted_string SEMICOLON END
    ;

// Creation Function - Specific Clauses

parallel_enable_clause
    : PARALLEL_ENABLE partition_by_clause?
    ;

partition_by_clause
    : LEFT_PAREN PARTITION expression BY (ANY | (HASH | RANGE | LIST) paren_column_list) streaming_clause? RIGHT_PAREN
    ;

result_cache_clause
    : RESULT_CACHE relies_on_part?
    ;

relies_on_part
    : RELIES_ON LEFT_PAREN tableview_name (COMMA tableview_name)* RIGHT_PAREN
    ;

streaming_clause
    : (ORDER | CLUSTER) expression BY paren_column_list
    ;

// https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/ALTER-OUTLINE.html
alter_outline
    : ALTER OUTLINE (PUBLIC | PRIVATE)? o=id_expression
        outline_options+
    ;

outline_options
    : REBUILD
    | RENAME TO non=id_expression
    | CHANGE CATEGORY TO ncn=id_expression
    | ENABLE
    | DISABLE
    ;

// https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/ALTER-LOCKDOWN-PROFILE.html
alter_lockdown_profile
    : ALTER LOCKDOWN PROFILE id_expression (lockdown_feature | lockdown_options | lockdown_statements)
        (USERS EQUALS_OP (ALL | COMMON | LOCAL))?
    ;

lockdown_feature
    : disable_enable FEATURE ( EQUALS_OP LEFT_PAREN string_list RIGHT_PAREN
                             | ALL (EXCEPT EQUALS_OP LEFT_PAREN string_list RIGHT_PAREN)?
                             )
    ;

lockdown_options
    : disable_enable OPTION ( EQUALS_OP LEFT_PAREN string_list RIGHT_PAREN
                            | ALL (EXCEPT EQUALS_OP LEFT_PAREN string_list RIGHT_PAREN)?
                            )
    ;

lockdown_statements
    : disable_enable STATEMENT ( EQUALS_OP LEFT_PAREN string_list RIGHT_PAREN
                               | EQUALS_OP LEFT_PAREN CHAR_STRING RIGHT_PAREN statement_clauses
                               | ALL (EXCEPT EQUALS_OP LEFT_PAREN string_list RIGHT_PAREN)?
                               )
    ;

statement_clauses
    : CLAUSE ( EQUALS_OP LEFT_PAREN string_list RIGHT_PAREN
             | EQUALS_OP LEFT_PAREN CHAR_STRING RIGHT_PAREN clause_options
             | ALL (EXCEPT EQUALS_OP LEFT_PAREN string_list RIGHT_PAREN)?
             )
    ;

clause_options
    : OPTION ( EQUALS_OP LEFT_PAREN string_list RIGHT_PAREN
             | EQUALS_OP LEFT_PAREN CHAR_STRING RIGHT_PAREN option_values+
             | ALL (EXCEPT EQUALS_OP LEFT_PAREN string_list RIGHT_PAREN)?
             )
    ;

option_values
    : VALUE EQUALS_OP LEFT_PAREN string_list RIGHT_PAREN
    | (MINVALUE | MAXVALUE) EQUALS_OP CHAR_STRING
    ;

string_list
    : CHAR_STRING (COMMA CHAR_STRING)*
    ;

disable_enable
    : DISABLE
    | ENABLE
    ;

// https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/DROP-LOCKDOWN-PROFILE.html
drop_lockdown_profile
    : DROP LOCKDOWN PROFILE p=id_expression
    ;

// Package DDLs

drop_package
    : DROP PACKAGE BODY? (schema_object_name PERIOD)? package_name SEMICOLON
    ;

alter_package
    : ALTER PACKAGE package_name COMPILE DEBUG? (PACKAGE | BODY | SPECIFICATION)? compiler_parameters_clause* (REUSE SETTINGS)? SEMICOLON
    ;

create_package
    : CREATE (OR REPLACE)? PACKAGE (schema_object_name PERIOD)? package_name
      (invoker_rights_clause | accessible_by_clause | default_collation_clause | annotations_clause)*
      (IS | AS) package_obj_spec* END package_name? SEMICOLON
    ;

create_package_body
    : CREATE (OR REPLACE)? PACKAGE BODY (schema_object_name PERIOD)? package_name (IS | AS) package_obj_body* (BEGIN seq_of_statements)? END package_name? SEMICOLON
    ;

// Create Package Specific Clauses

package_obj_spec
    : pragma_declaration
    | exception_declaration
    | variable_declaration
    | subtype_declaration
    | cursor_declaration
    | type_declaration
    | procedure_spec
    | function_spec
    ;

procedure_spec
    : PROCEDURE identifier (LEFT_PAREN parameter ( COMMA parameter )* RIGHT_PAREN)? SEMICOLON
    ;

function_spec
    : FUNCTION identifier (LEFT_PAREN parameter ( COMMA parameter)* RIGHT_PAREN)?
      RETURN type_spec PIPELINED? DETERMINISTIC? (RESULT_CACHE)? SEMICOLON
    ;

package_obj_body
    : exception_declaration
    | subtype_declaration
    | cursor_declaration
    | variable_declaration
    | type_declaration
    | procedure_body
    | function_body
    | procedure_spec
    | function_spec
    ;

// https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/alter-pmem-filestore.html
alter_pmem_filestore
    : ALTER PMEM FILESTORE fsn=id_expression
        ( RESIZE size_clause
        | autoextend_clause
        | MOUNT (MOUNTPOINT file_path)? (BACKINGFILE filename)? FORCE? //inconsistent documentation
        | DISMOUNT
        )
    ;

// https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/drop-pmem-filestore.html
drop_pmem_filestore
    : DROP PMEM FILESTORE fsn=id_expression
        ((FORCE? INCLUDING | EXCLUDING) CONTENTS)?
    ;

// Procedure DDLs

drop_procedure
    : DROP PROCEDURE procedure_name SEMICOLON
    ;

alter_procedure
    : ALTER PROCEDURE procedure_name COMPILE DEBUG? compiler_parameters_clause* (REUSE SETTINGS)? SEMICOLON
    ;

function_body
    : FUNCTION identifier (LEFT_PAREN parameter (COMMA parameter)* RIGHT_PAREN)?
      RETURN type_spec (invoker_rights_clause | parallel_enable_clause | result_cache_clause | DETERMINISTIC)*
      ((PIPELINED? DETERMINISTIC? (IS | AS) (DECLARE? seq_of_declare_specs? body | call_spec)) | (PIPELINED | AGGREGATE) USING implementation_type_name) SEMICOLON
    ;

procedure_body
    : PROCEDURE identifier (LEFT_PAREN parameter (COMMA parameter)* RIGHT_PAREN)? (IS | AS)
      (DECLARE? seq_of_declare_specs? body | call_spec | EXTERNAL) SEMICOLON
    ;

create_procedure_body
    : CREATE (OR REPLACE)? PROCEDURE procedure_name (LEFT_PAREN parameter (COMMA parameter)* RIGHT_PAREN)?
      (invoker_rights_clause | accessible_by_clause | default_collation_clause | annotations_clause)* (IS | AS)
      (DECLARE? seq_of_declare_specs? body | call_spec | EXTERNAL) SEMICOLON
    ;

// https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/ALTER-RESOURCE-COST.html
alter_resource_cost
    : ALTER RESOURCE COST ((CPU_PER_SESSION | CONNECT_TIME | LOGICAL_READS_PER_SESSION | PRIVATE_SGA) UNSIGNED_INTEGER)+
    ;

// https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/DROP-OUTLINE.html
drop_outline
    : DROP OUTLINE o=id_expression
    ;

// Rollback Segment DDLs

//https://docs.oracle.com/cd/E11882_01/server.112/e41084/statements_2011.htm#SQLRF00816
alter_rollback_segment
    : ALTER ROLLBACK SEGMENT rollback_segment_name (ONLINE | OFFLINE | storage_clause | SHRINK (TO size_clause)?)
    ;

// https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/CREATE-RESTORE-POINT.html
drop_restore_point
    : DROP RESTORE POINT rp=id_expression (FOR PLUGGABLE DATABASE pdb=id_expression)?
    ;

drop_rollback_segment
    : DROP ROLLBACK SEGMENT rollback_segment_name
    ;

drop_role
    : DROP ROLE role_name SEMICOLON
    ;

// https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/create-pmem-filestore.html
create_pmem_filestore
    : CREATE PMEM FILESTORE fsn=id_expression
        pmem_filestore_options+
    ;

pmem_filestore_options
    : MOUNTPOINT file_path
    | BACKINGFILE filename REUSE?
    | (SIZE | BLOCKSIZE) size_clause
    | autoextend_clause
    ;

file_path
    : CHAR_STRING
    ;

create_rollback_segment
    : CREATE PUBLIC? ROLLBACK SEGMENT rollback_segment_name (TABLESPACE tablespace | storage_clause)*
    ;

// Trigger DDLs

drop_trigger
    : DROP TRIGGER trigger_name SEMICOLON
    ;

alter_trigger
    : ALTER TRIGGER alter_trigger_name=trigger_name
      ((ENABLE | DISABLE) | RENAME TO rename_trigger_name=trigger_name | COMPILE DEBUG? compiler_parameters_clause* (REUSE SETTINGS)?) SEMICOLON
    ;

create_trigger
    : CREATE ( OR REPLACE )? TRIGGER trigger_name
      (simple_dml_trigger | compound_dml_trigger | non_dml_trigger)
      trigger_follows_clause? (ENABLE | DISABLE)? trigger_when_clause? trigger_body SEMICOLON
    ;

trigger_follows_clause
    : FOLLOWS trigger_name (COMMA trigger_name)*
    ;

trigger_when_clause
    : WHEN LEFT_PAREN condition RIGHT_PAREN
    ;

// Create Trigger Specific Clauses

simple_dml_trigger
    : (BEFORE | AFTER | INSTEAD OF) dml_event_clause referencing_clause? for_each_row?
    ;

for_each_row
    : FOR EACH ROW
    ;

compound_dml_trigger
    : FOR dml_event_clause referencing_clause?
    ;

non_dml_trigger
    : (BEFORE | AFTER) non_dml_event (OR non_dml_event)* ON (DATABASE | (schema_name PERIOD)? SCHEMA)
    ;

trigger_body
    : compound_trigger_block
    | CALL identifier
    | trigger_block
    ;

non_dml_event
    : ALTER
    | ANALYZE
    | ASSOCIATE STATISTICS
    | AUDIT
    | COMMENT
    | CREATE
    | DISASSOCIATE STATISTICS
    | DROP
    | GRANT
    | NOAUDIT
    | RENAME
    | REVOKE
    | TRUNCATE
    | DDL
    | STARTUP
    | SHUTDOWN
    | DB_ROLE_CHANGE
    | LOGON
    | LOGOFF
    | SERVERERROR
    | SUSPEND
    | DATABASE
    | SCHEMA
    | FOLLOWS
    ;

dml_event_clause
    : dml_event_element (OR dml_event_element)* ON dml_event_nested_clause? tableview_name
    ;

dml_event_element
    : (DELETE | INSERT | UPDATE) (OF column_list)?
    ;

dml_event_nested_clause
    : NESTED TABLE tableview_name OF
    ;

referencing_clause
    : REFERENCING referencing_element+
    ;

referencing_element
    : (NEW | OLD | PARENT) column_alias
    ;

// DDLs

drop_type
    : DROP TYPE BODY? type_name (FORCE | VALIDATE)? SEMICOLON
    ;

alter_type
    : ALTER TYPE type_name
    (compile_type_clause
    | replace_type_clause
    //TODO | {input.LT(2).getText().equalsIgnoreCase("attribute")}? alter_attribute_definition
    | alter_method_spec
    | alter_collection_clauses
    | modifier_clause
    | overriding_subprogram_spec
    ) dependent_handling_clause? SEMICOLON
    ;

// Alter Type Specific Clauses

compile_type_clause
    : COMPILE DEBUG? (SPECIFICATION | BODY)? compiler_parameters_clause* (REUSE SETTINGS)?
    ;

replace_type_clause
    : REPLACE invoker_rights_clause? AS OBJECT LEFT_PAREN object_member_spec (COMMA object_member_spec)* RIGHT_PAREN
    ;

alter_method_spec
    : alter_method_element (COMMA alter_method_element)*
    ;

alter_method_element
    : (ADD | DROP) (map_order_function_spec | subprogram_spec)
    ;

alter_collection_clauses
    : MODIFY (LIMIT expression | ELEMENT TYPE type_spec)
    ;

dependent_handling_clause
    : INVALIDATE
    | CASCADE (CONVERT TO SUBSTITUTABLE | NOT? INCLUDING TABLE DATA)? dependent_exceptions_part?
    ;

dependent_exceptions_part
    : FORCE? EXCEPTIONS INTO tableview_name
    ;

create_type
    : CREATE (OR REPLACE)? TYPE (type_definition | type_body) SEMICOLON
    ;

// Create Type Specific Clauses

type_definition
    : type_name (OID CHAR_STRING)? FORCE? object_type_def?
    ;

object_type_def
    : invoker_rights_clause? (object_as_part | object_under_part) sqlj_object_type?
      (LEFT_PAREN object_member_spec (COMMA object_member_spec)* RIGHT_PAREN)? modifier_clause*
    ;

object_as_part
    : (IS | AS) (OBJECT | varray_type_def | nested_table_type_def)
    ;

object_under_part
    : UNDER type_spec
    ;

nested_table_type_def
    : TABLE OF type_spec (NOT NULL_)?
    ;

sqlj_object_type
    : EXTERNAL NAME expression LANGUAGE JAVA USING (SQLDATA | CUSTOMDATUM | ORADATA)
    ;

type_body
    : BODY type_name (IS | AS) (type_body_elements)+ END
    ;

type_body_elements
    : map_order_func_declaration
    | subprog_decl_in_type
    | overriding_subprogram_spec
    ;

map_order_func_declaration
    : (MAP | ORDER) MEMBER func_decl_in_type
    ;

subprog_decl_in_type
    : (MEMBER | STATIC) (proc_decl_in_type | func_decl_in_type | constructor_declaration)
    ;

proc_decl_in_type
    : PROCEDURE procedure_name LEFT_PAREN type_elements_parameter (COMMA type_elements_parameter)* RIGHT_PAREN
      (IS | AS) (call_spec | DECLARE? seq_of_declare_specs? body SEMICOLON)
    ;

func_decl_in_type
    : FUNCTION function_name (LEFT_PAREN type_elements_parameter (COMMA type_elements_parameter)* RIGHT_PAREN)?
      RETURN type_spec (IS | AS) (call_spec | DECLARE? seq_of_declare_specs? body SEMICOLON)
    ;

constructor_declaration
    : FINAL? INSTANTIABLE? CONSTRUCTOR FUNCTION type_spec
      (LEFT_PAREN (SELF IN OUT type_spec COMMA) type_elements_parameter (COMMA type_elements_parameter)*  RIGHT_PAREN)?
      RETURN SELF AS RESULT (IS | AS) (call_spec | DECLARE? seq_of_declare_specs? body SEMICOLON)
    ;

// Common Type Clauses

modifier_clause
    : NOT? (INSTANTIABLE | FINAL | OVERRIDING)
    ;

object_member_spec
    : identifier type_spec sqlj_object_type_attr?
    | element_spec
    ;

sqlj_object_type_attr
    : EXTERNAL NAME expression
    ;

element_spec
    : modifier_clause? element_spec_options+ (COMMA pragma_clause)?
    ;

element_spec_options
    : subprogram_spec
    | constructor_spec
    | map_order_function_spec
    ;

subprogram_spec
    : (MEMBER | STATIC) (type_procedure_spec | type_function_spec)
    ;

// TODO: should be refactored such as Procedure body and Function body, maybe Type_Function_Body and overriding_function_body
overriding_subprogram_spec
    : OVERRIDING MEMBER overriding_function_spec
    ;

overriding_function_spec
    : FUNCTION function_name (LEFT_PAREN type_elements_parameter (COMMA type_elements_parameter)* RIGHT_PAREN)?
      RETURN (type_spec | SELF AS RESULT)
     (PIPELINED? (IS | AS) (DECLARE? seq_of_declare_specs? body))? SEMICOLON?
    ;

type_procedure_spec
    : PROCEDURE procedure_name LEFT_PAREN type_elements_parameter (COMMA type_elements_parameter)* RIGHT_PAREN ((IS | AS) call_spec)?
    ;

type_function_spec
    : FUNCTION function_name (LEFT_PAREN type_elements_parameter (COMMA type_elements_parameter)* RIGHT_PAREN)?
      RETURN (type_spec | SELF AS RESULT) ((IS | AS) call_spec | EXTERNAL VARIABLE? NAME expression)?
    ;

constructor_spec
    : FINAL? INSTANTIABLE? CONSTRUCTOR FUNCTION
      type_spec (LEFT_PAREN (SELF IN OUT type_spec COMMA) type_elements_parameter (COMMA type_elements_parameter)*  RIGHT_PAREN)?
      RETURN SELF AS RESULT ((IS | AS) call_spec)?
    ;

map_order_function_spec
    : (MAP | ORDER) MEMBER type_function_spec
    ;

pragma_clause
    : PRAGMA RESTRICT_REFERENCES LEFT_PAREN pragma_elements (COMMA pragma_elements)* RIGHT_PAREN
    ;

pragma_elements
    : identifier
    | DEFAULT
    ;

type_elements_parameter
    : parameter_name type_spec
    ;

// Sequence DDLs

drop_sequence
    : DROP SEQUENCE sequence_name SEMICOLON
    ;

alter_sequence
    : ALTER SEQUENCE sequence_name sequence_spec+ SEMICOLON
    ;

// https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/ALTER-SESSION.html
alter_session
    : ALTER SESSION (
        ADVISE ( COMMIT | ROLLBACK | NOTHING )
        | CLOSE DATABASE LINK parameter_name
        | enable_or_disable COMMIT IN PROCEDURE
        | enable_or_disable GUARD
        | (enable_or_disable | FORCE) PARALLEL (DML | DDL | QUERY) (PARALLEL (literal | parameter_name))?
        | SET alter_session_set_clause
    )
    ;

alter_session_set_clause
    : (parameter_name EQUALS_OP parameter_value)+
    | EDITION EQUALS_OP en=id_expression
    | CONTAINER EQUALS_OP cn=id_expression (SERVICE EQUALS_OP sn=id_expression)?
    | ROW ARCHIVAL VISIBILITY EQUALS_OP (ACTIVE | ALL)
    | DEFAULT_COLLATION EQUALS_OP (c=id_expression | NONE)
    ;

create_sequence
    : CREATE SEQUENCE (IF NOT EXISTS)? sequence_name sequence_spec* (SHARING EQUALS_OP (METADATA | DATA | NONE))? SEMICOLON
    ;

// Common Sequence

sequence_spec
    : INCREMENT BY UNSIGNED_INTEGER
    | sequence_start_clause
    | MAXVALUE UNSIGNED_INTEGER
    | NOMAXVALUE
    | MINVALUE UNSIGNED_INTEGER
    | NOMINVALUE
    | CYCLE
    | NOCYCLE
    | CACHE UNSIGNED_INTEGER
    | NOCACHE
    | ORDER
    | NOORDER
    | KEEP
    | NOKEEP
    | SCALE (EXTEND | NOEXTEND)?
    | NOSCALE
    | SHARD (EXTEND | NOEXTEND)?
    | NOSHARD
    | SESSION
    | GLOBAL
    ;

sequence_start_clause
    : START WITH UNSIGNED_INTEGER
    ;

// https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/CREATE-ANALYTIC-VIEW.html
create_analytic_view
    : CREATE (OR REPLACE)? (NOFORCE | FORCE)? ANALYTIC VIEW av=id_expression
        (SHARING EQUALS_OP (METADATA | NONE))?
        classification_clause*
        cav_using_clause?
        dim_by_clause?
        measures_clause?
        default_measure_clause?
        default_aggregate_clause?
        cache_clause?
        fact_columns_clause?
        qry_transform_clause?
    ;

classification_clause
// : (CAPTION c=quoted_string)? (DESCRIPTION d=quoted_string)? classification_item*
// to handle - 'rule contains a closure with at least one alternative that can match an empty string'
    : (caption_clause description_clause? | caption_clause? description_clause) classification_item*
    | caption_clause? description_clause? classification_item+
    ;

caption_clause
    : CAPTION c=quoted_string
    ;

description_clause
    : DESCRIPTION d=quoted_string
    ;

classification_item
    : CLASSIFICATION cn=id_expression (VALUE cv=quoted_string)? (LANGUAGE language)?
    ;

language
    : NULL_
    | nls=id_expression
    ;

cav_using_clause
    : USING (schema_name PERIOD)? t=id_expression REMOTE? (AS? ta=id_expression)?
    ;

dim_by_clause
    : DIMENSION BY LEFT_PAREN dim_key (COMMA dim_key)* RIGHT_PAREN
    ;

dim_key
    : dim_ref classification_clause*
        KEY ( LEFT_PAREN (a=id_expression PERIOD)? f=column_name (COMMA (a=id_expression PERIOD)? f=column_name)* RIGHT_PAREN
            |  (a=id_expression PERIOD)? f=column_name
            )
        REFERENCES DISTINCT? (LEFT_PAREN attribute_name (COMMA attribute_name) RIGHT_PAREN | attribute_name)
        HIERARCHIES LEFT_PAREN hier_ref (COMMA hier_ref)* RIGHT_PAREN
    ;

dim_ref
    : (schema_name PERIOD)? ad=id_expression (AS? da=id_expression)?
    ;

hier_ref
    : (schema_name PERIOD)? h=id_expression (AS? ha=id_expression)? DEFAULT?
    ;

measures_clause
    : MEASURES LEFT_PAREN av_measure (COMMA av_measure)* RIGHT_PAREN
    ;

av_measure
    : mn=id_expression (base_meas_clause | calc_meas_clause)? //classification_clause*
    ;

base_meas_clause
    : FACT /*FOR MEASURE*/ bm=id_expression meas_aggregate_clause? //FIXME inconsistent documentation
    ;

meas_aggregate_clause
    : AGGREGATE BY aggregate_function_name
    ;

calc_meas_clause
    : /*m=id_expression*/ AS LEFT_PAREN expression RIGHT_PAREN //FIXME inconsistent documentation
    ;

default_measure_clause
    : DEFAULT MEASURE m=id_expression
    ;

default_aggregate_clause
    : DEFAULT AGGREGATE BY aggregate_function_name
    ;

cache_clause
    : CACHE cache_specification (COMMA cache_specification)*
    ;

cache_specification
    : MEASURE GROUP (ALL | LEFT_PAREN id_expression (COMMA id_expression)* RIGHT_PAREN levels_clause (COMMA levels_clause)* )
    ;

levels_clause
    : LEVELS LEFT_PAREN level_specification (COMMA level_specification)* RIGHT_PAREN level_group_type
    ;

level_specification
    : LEFT_PAREN ((d=id_expression PERIOD)? h=id_expression PERIOD)? l=id_expression RIGHT_PAREN
    ;

level_group_type
    : DYNAMIC
    | MATERIALIZED (USING (schema_name PERIOD)? t=id_expression)?
    ;

fact_columns_clause
    : FACT COLUMN f=column_name (AS? fa=id_expression (COMMA AS? fa=id_expression)* )?
    ;

qry_transform_clause
    : ENABLE QUERY TRANSFORM (RELY | NORELY)?
    ;

// https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/CREATE-ATTRIBUTE-DIMENSION.html
create_attribute_dimension
    : CREATE (OR REPLACE)? (NOFORCE | FORCE)? ATTRIBUTE DIMENSION (schema_name PERIOD)? ad=id_expression
        (SHARING EQUALS_OP (METADATA | NONE))?
        classification_clause*
        (DIMENSION TYPE (STANDARD | TIME))?
        ad_using_clause
        attributes_clause
        ad_level_clause+
        all_clause?
    ;

ad_using_clause
    : USING source_clause (COMMA source_clause)* join_path_clause*
    ;

source_clause
    : (schema_name PERIOD)? ftov=id_expression REMOTE? (AS? a=id_expression)?
    ;

join_path_clause
    : JOIN PATH jpn=id_expression ON join_condition
    ;

join_condition
    : join_condition_item (AND join_condition_item)*
    ;

join_condition_item
    : (a=id_expression PERIOD)? column_name EQUALS_OP (b=id_expression PERIOD)? column_name
    ;

attributes_clause
    : ATTRIBUTES LEFT_PAREN ad_attributes_clause (COMMA ad_attributes_clause)* RIGHT_PAREN
    ;

ad_attributes_clause
    : (a=id_expression PERIOD)? column_name (AS? an=id_expression)?
        classification_clause*
    ;

ad_level_clause
    : LEVEL l=id_expression (NOT NULL_ | SKIP_ WHEN NULL_)?
        (LEVEL TYPE (STANDARD | YEARS | HALF_YEARS | QUARTERS | MONTHS | WEEKS | DAYS | HOURS | MINUTES | SECONDS))?
        classification_clause* //inconsistent documentation - LEVEL TYPE goes after the classification_clause rule
        key_clause
        alternate_key_clause?
        (MEMBER NAME expression)?
        (MEMBER CAPTION expression)?
        (MEMBER DESCRIPTION expression)?
        (ORDER BY (MIN | MAX)? dim_order_clause (COMMA (MIN | MAX)? dim_order_clause)*)?
        (DETERMINES LEFT_PAREN id_expression (COMMA id_expression)* RIGHT_PAREN)?
    ;

key_clause
    : KEY (a=id_expression | LEFT_PAREN id_expression (COMMA id_expression)* RIGHT_PAREN)
    ;

alternate_key_clause
    : ALTERNATE key_clause
    ;

dim_order_clause
    : a=id_expression (ASC | DESC)? (NULLS (FIRST | LAST))?
    ;

all_clause
    : ALL MEMBER ( NAME expression (MEMBER CAPTION expression)?
                 | CAPTION expression (MEMBER DESCRIPTION expression)?
                 | DESCRIPTION expression
                 )
    ;

// https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/CREATE-AUDIT-POLICY-Unified-Auditing.html
create_audit_policy
    : CREATE AUDIT POLICY p=id_expression
        privilege_audit_clause?
        action_audit_clause?
        role_audit_clause?
        (WHEN quoted_string EVALUATE PER (STATEMENT | SESSION | INSTANCE))?
        (ONLY TOPLEVEL)?
        container_clause?
    ;

privilege_audit_clause
    : PRIVILEGES system_privilege (COMMA system_privilege)*
    ;

action_audit_clause
    : (standard_actions | component_actions)+
    ;

standard_actions
    : ACTIONS actions_clause (COMMA actions_clause)*
    ;

actions_clause
    : (object_action | ALL) ON (DIRECTORY directory_name | (MINING MODEL)? (schema_name PERIOD)? id_expression)
    | (system_action | ALL)
    ;

object_action
    : ALTER
    | GRANT
    | READ
    | EXECUTE
    | AUDIT
    | COMMENT
    | DELETE
    | INDEX
    | INSERT
    | LOCK
    | SELECT
    | UPDATE
    | FLASHBACK
    | RENAME
    ;

// BYT-8268: Added support for multi-word system actions
// Note: Complete list varies by Oracle version, query AUDITABLE_SYSTEM_ACTIONS view for full list
system_action
    : ADMINISTER KEY MANAGEMENT  // TDE operations
    | (CREATE | ALTER | DROP) JAVA
    | LOCK TABLE
    | (READ | WRITE | EXECUTE) DIRECTORY
    | id_expression id_expression id_expression  // 3-word actions (ALTER AUDIT POLICY, etc.)
    | id_expression id_expression  // 2-word actions (ALTER ASSEMBLY, CREATE TABLE, etc.)
    | id_expression  // Single-word actions
    ;

component_actions
    : ACTIONS COMPONENT EQUALS_OP ( (DATAPUMP | DIRECT_LOAD | OLS | XS) component_action (COMMA component_action)*
                            | DV component_action ON id_expression (COMMA component_action ON id_expression)*
                            | PROTOCOL (FTP | HTTP | AUTHENTICATION)
                            )
    ;

component_action
    : id_expression // SELECT name FROM auditable_system_actions WHERE component = 'Datapump';
    ;

role_audit_clause
    : ROLES role_name (COMMA role_name)*
    ;

// https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/CREATE-CONTROLFILE.html
create_controlfile
    : CREATE CONTROLFILE REUSE? SET? DATABASE d=id_expression
        logfile_clause? (RESETLOGS | NORESETLOGS)
        (DATAFILE file_specification (COMMA file_specification)*)?
        controlfile_options*
        character_set_clause?
    ;

controlfile_options
    : MAXLOGFILES numeric
    | MAXLOGMEMBERS numeric
    | MAXLOGHISTORY numeric
    | MAXDATAFILES numeric
    | MAXINSTANCES numeric
    | ARCHIVELOG
    | NOARCHIVELOG
    | FORCE LOGGING
    | SET STANDBY NOLOGGING FOR (DATA AVAILABILITY | LOAD PERFORMANCE)
    ;

logfile_clause
    : LOGFILE (GROUP? numeric)? file_specification (COMMA (GROUP? numeric)? file_specification)*
    ;

character_set_clause
    : CHARACTER SET cs=id_expression
    ;

file_specification
    : datafile_tempfile_spec
    | redo_log_file_spec
    ;

// https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/CREATE-DISKGROUP.html
create_diskgroup
    : CREATE DISKGROUP id_expression ((HIGH | NORMAL | FLEX | EXTENDED (SITE sn=id_expression)? | EXTERNAL) REDUNDANCY)?
        (quorum_regular? (FAILGROUP fg=id_expression)? DISK qualified_disk_clause (COMMA qualified_disk_clause)*)+
        (ATTRIBUTE an=CHAR_STRING EQUALS_OP av=CHAR_STRING (COMMA CHAR_STRING EQUALS_OP CHAR_STRING)*)?
    ;

qualified_disk_clause
    : ss=CHAR_STRING (NAME dn=id_expression)? (SIZE size_clause)? force_noforce?
    ;

// https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/CREATE-EDITION.html
create_edition
    : CREATE EDITION e=id_expression (AS CHILD OF pe=id_expression)?
    ;

// https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/CREATE-FLASHBACK-ARCHIVE.html
create_flashback_archive
    : CREATE FLASHBACK ARCHIVE DEFAULT? fa=id_expression TABLESPACE ts=id_expression
        flashback_archive_quota?
        (NO? OPTIMIZE DATA)?
        flashback_archive_retention
    ;

flashback_archive_quota
    : QUOTA UNSIGNED_INTEGER (M_LETTER | G_LETTER | T_LETTER | P_LETTER | E_LETTER)
    ;

flashback_archive_retention
    : RETENTION UNSIGNED_INTEGER (YEAR | MONTH | DAY)
    ;

// https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/CREATE-HIERARCHY.html
create_hierarchy
    : CREATE (OR REPLACE)? (NO? FORCE)? HIERARCHY (schema_name PERIOD)? h=id_expression
        (SHARING EQUALS_OP (METADATA | NONE))?
        classification_clause*
        hier_using_clause
        level_hier_clause
        hier_attrs_clause?
    ;

hier_using_clause
    : USING (schema_name PERIOD)? ad=id_expression
    ;

level_hier_clause
    : LEFT_PAREN (l=id_expression (CHILD OF)?)+ RIGHT_PAREN
    ;

hier_attrs_clause
    : HIERARCHICAL ATTRIBUTES LEFT_PAREN hier_attr_clause RIGHT_PAREN
    ;

hier_attr_clause
    : hier_attr_name classification_clause*
    ;

hier_attr_name
    : MEMBER_NAME
    | MEMBER_UNIQUE_NAME
    | MEMBER_CAPTION
    | MEMBER_DESCRIPTION
    | LEVEL_NAME
    | HIER_ORDER
    | DEPTH
    | IS_LEAF
    | PARENT_LEVEL_NAME
    | PARENT_UNIQUE_NAME
    ;

create_index
    : CREATE (UNIQUE | BITMAP)? INDEX index_name
       ON (cluster_index_clause | table_index_clause | bitmap_join_index_clause)
       (USABLE | UNUSABLE)?
       SEMICOLON
    ;

cluster_index_clause
    : CLUSTER cluster_name index_attributes?
    ;

cluster_name
    : (id_expression PERIOD)? id_expression
    ;

table_index_clause
    : tableview_name table_alias? LEFT_PAREN index_expr_option  (COMMA index_expr_option)* RIGHT_PAREN
          index_properties?
    ;
bitmap_join_index_clause
    : tableview_name LEFT_PAREN (tableview_name | table_alias)? column_name (ASC | DESC)?  (COMMA (tableview_name | table_alias)? column_name (ASC | DESC)? )* RIGHT_PAREN
        FROM tableview_name table_alias (COMMA tableview_name table_alias)*
        where_clause local_partitioned_index? index_attributes?
    ;
    
index_expr_option
    : index_expr (ASC | DESC)?
    ;

index_expr
    : column_name
    | expression
    ;

index_properties
    : (global_partitioned_index | local_partitioned_index | index_attributes)+
    | INDEXTYPE IS (domain_index_clause | xmlindex_clause)
    ;

domain_index_clause
    : indextype local_domain_index_clause? parallel_clause? (PARAMETERS LEFT_PAREN odci_parameters RIGHT_PAREN )?
    ;

local_domain_index_clause
    : LOCAL (LEFT_PAREN PARTITION partition_name (PARAMETERS LEFT_PAREN odci_parameters RIGHT_PAREN )?  (COMMA PARTITION partition_name (PARAMETERS LEFT_PAREN odci_parameters RIGHT_PAREN )? )* RIGHT_PAREN )?
    ;

xmlindex_clause
    : (XDB PERIOD)? XMLINDEX local_xmlindex_clause?
        parallel_clause? //TODO xmlindex_parameters_clause?
    ;

local_xmlindex_clause
    : LOCAL (LEFT_PAREN PARTITION partition_name (COMMA PARTITION partition_name //TODO xmlindex_parameters_clause?
                                                       )* RIGHT_PAREN)?
    ;

global_partitioned_index
    : GLOBAL (
                PARTITION BY (RANGE LEFT_PAREN column_name (COMMA column_name)* RIGHT_PAREN LEFT_PAREN index_partitioning_clause (COMMA index_partitioning_clause)* RIGHT_PAREN
                            | HASH LEFT_PAREN column_name (COMMA column_name)* RIGHT_PAREN
                                            (individual_hash_partitions
                                            | hash_partitions_by_quantity
                                            )
                             )
             )?
    ;

index_partitioning_clause
    : PARTITION partition_name? VALUES LESS THAN LEFT_PAREN literal (COMMA literal)* RIGHT_PAREN
        segment_attributes_clause?
    ;

local_partitioned_index
    : LOCAL (on_range_partitioned_table
            | on_list_partitioned_table
            | on_hash_partitioned_table
            | on_comp_partitioned_table
            )?
    ;

on_range_partitioned_table
    : LEFT_PAREN partitioned_table (COMMA partitioned_table)* RIGHT_PAREN
    ;

on_list_partitioned_table
    : LEFT_PAREN partitioned_table (COMMA partitioned_table)* RIGHT_PAREN
    ;

partitioned_table
    :  PARTITION partition_name?
        (segment_attributes_clause | key_compression)*
        UNUSABLE?
    ;

on_hash_partitioned_table
    : STORE IN LEFT_PAREN tablespace (COMMA tablespace)* RIGHT_PAREN
    | LEFT_PAREN on_hash_partitioned_clause (COMMA on_hash_partitioned_clause)* RIGHT_PAREN
    ;

on_hash_partitioned_clause
    : PARTITION partition_name? (TABLESPACE tablespace)?
        key_compression? UNUSABLE?
    ;
on_comp_partitioned_table
    : (STORE IN LEFT_PAREN tablespace (COMMA tablespace)* RIGHT_PAREN )?
        LEFT_PAREN on_comp_partitioned_clause (COMMA on_comp_partitioned_clause)* RIGHT_PAREN
    ;

on_comp_partitioned_clause
    : PARTITION partition_name?
        (segment_attributes_clause | key_compression)*
        UNUSABLE index_subpartition_clause?
    ;

index_subpartition_clause
    : STORE IN LEFT_PAREN tablespace (COMMA tablespace)* RIGHT_PAREN
    | LEFT_PAREN index_subpartition_subclause (COMMA index_subpartition_subclause)* RIGHT_PAREN
    ;

index_subpartition_subclause
    : SUBPARTITION subpartition_name? (TABLESPACE tablespace)?
        key_compression? UNUSABLE?
    ;

odci_parameters
    : CHAR_STRING
    ;

indextype
    : (id_expression PERIOD)? id_expression
    ;

//https://docs.oracle.com/cd/E11882_01/server.112/e41084/statements_1010.htm#SQLRF00805
alter_index
    : ALTER INDEX index_name (alter_index_ops_set1 | alter_index_ops_set2) SEMICOLON
    ;

alter_index_ops_set1
    : ( deallocate_unused_clause
      | allocate_extent_clause
      | shrink_clause
      | parallel_clause
      | physical_attributes_clause
      | logging_clause
      )+
    ;

alter_index_ops_set2
    : rebuild_clause
    | PARAMETERS LEFT_PAREN odci_parameters RIGHT_PAREN
    | COMPILE
    | enable_or_disable
    | UNUSABLE
    | visible_or_invisible
    | RENAME TO new_index_name
    | COALESCE
    | monitoring_nomonitoring USAGE
    | UPDATE BLOCK REFERENCES
    | alter_index_partitioning
    ;

visible_or_invisible
    : VISIBLE
    | INVISIBLE
    ;

monitoring_nomonitoring
    : MONITORING
    | NOMONITORING
    ;

rebuild_clause
    : REBUILD ( PARTITION partition_name
              | SUBPARTITION subpartition_name
              | REVERSE
              | NOREVERSE
              )?
              ( parallel_clause
              | TABLESPACE tablespace
              | PARAMETERS LEFT_PAREN odci_parameters RIGHT_PAREN
//TODO        | xmlindex_parameters_clause
              | ONLINE
              | physical_attributes_clause
              | key_compression
              | logging_clause
              )*
    ;

alter_index_partitioning
    : modify_index_default_attrs
    | add_hash_index_partition
    | modify_index_partition
    | rename_index_partition
    | drop_index_partition
    | split_index_partition
    | coalesce_index_partition
    | modify_index_subpartition
    ;

modify_index_default_attrs
    : MODIFY DEFAULT ATTRIBUTES (FOR PARTITION partition_name)?
         ( physical_attributes_clause
         | TABLESPACE (tablespace | DEFAULT)
         | logging_clause
         )
    ;

add_hash_index_partition
    : ADD PARTITION partition_name? (TABLESPACE tablespace)?
        key_compression? parallel_clause?
    ;

coalesce_index_partition
    : COALESCE PARTITION parallel_clause?
    ;

modify_index_partition
    : MODIFY PARTITION partition_name
        ( modify_index_partitions_ops+
        | PARAMETERS LEFT_PAREN odci_parameters RIGHT_PAREN
        | COALESCE
        | UPDATE BLOCK REFERENCES
        | UNUSABLE
        )
    ;

modify_index_partitions_ops
    : deallocate_unused_clause
    | allocate_extent_clause
    | physical_attributes_clause
    | logging_clause
    | key_compression
    ;

rename_index_partition
    : RENAME (PARTITION partition_name | SUBPARTITION subpartition_name)
         TO new_partition_name
    ;

drop_index_partition
    : DROP PARTITION partition_name
    ;

split_index_partition
    : SPLIT PARTITION partition_name_old AT LEFT_PAREN literal (COMMA literal)* RIGHT_PAREN
        (INTO LEFT_PAREN index_partition_description COMMA index_partition_description RIGHT_PAREN ) ? parallel_clause?
    ;

index_partition_description
    : PARTITION (partition_name ( (segment_attributes_clause | key_compression)+
                                | PARAMETERS LEFT_PAREN odci_parameters RIGHT_PAREN
                                )
                                UNUSABLE?
                )?
    ;

modify_index_subpartition
    : MODIFY SUBPARTITION subpartition_name (UNUSABLE
                                            | allocate_extent_clause
                                            | deallocate_unused_clause
                                            )
    ;

partition_name_old
    : partition_name
    ;

new_partition_name
    : partition_name
    ;

new_index_name
    : index_name
    ;

// https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/ALTER-INMEMORY-JOIN-GROUP.html
alter_inmemory_join_group
    : ALTER INMEMORY JOIN GROUP (schema_name PERIOD)? jg=id_expression
        (ADD | REMOVE) LEFT_PAREN (schema_name PERIOD)? t=id_expression LEFT_PAREN c=id_expression RIGHT_PAREN RIGHT_PAREN
    ;

create_user
    : CREATE USER
      user_object_name
        ( identified_by
          | identified_other_clause
          | user_tablespace_clause
          | quota_clause
          | profile_clause
          | password_expire_clause
          | user_lock_clause
          | user_editions_clause
          | container_clause
        )+ SEMICOLON
    ;

// The standard clauses only permit one user per statement.
// The proxy clause allows multiple users for a proxy designation.
alter_user
    : ALTER USER
      user_object_name
        ( alter_identified_by
        | identified_other_clause
        | user_tablespace_clause
        | quota_clause
        | profile_clause
        | user_default_role_clause
        | password_expire_clause
        | user_lock_clause
        | alter_user_editions_clause
        | container_clause
        | container_data_clause
        )+
      SEMICOLON
      | user_object_name (COMMA user_object_name)* proxy_clause SEMICOLON
    ;

drop_user
    : DROP USER user_object_name CASCADE?
    ;

alter_identified_by
    : identified_by (REPLACE id_expression)?
    ;

identified_by
    : IDENTIFIED BY id_expression
    ;

identified_other_clause
    : IDENTIFIED (EXTERNALLY | GLOBALLY) (AS quoted_string)?
    ;

user_tablespace_clause
    : (DEFAULT | TEMPORARY) TABLESPACE id_expression
    ;

quota_clause
    : QUOTA (size_clause | UNLIMITED) ON id_expression
    ;

profile_clause
    : PROFILE id_expression
    ;

role_clause
    : role_name (COMMA role_name)*
    | ALL (EXCEPT role_name (COMMA role_name)*)*
    ;

user_default_role_clause
    : DEFAULT ROLE (NONE | role_clause)
    ;

password_expire_clause
    : PASSWORD EXPIRE
    ;

user_lock_clause
    : ACCOUNT (LOCK | UNLOCK)
    ;

user_editions_clause
    : ENABLE EDITIONS
    ;

alter_user_editions_clause
    : user_editions_clause (FOR regular_id (COMMA regular_id)*)? FORCE?
    ;

proxy_clause
    : REVOKE CONNECT THROUGH (ENTERPRISE USERS | user_object_name)
    | GRANT CONNECT THROUGH
        ( ENTERPRISE USERS
        | user_object_name
            (WITH (NO ROLES | ROLE role_clause))?
            (AUTHENTICATION REQUIRED)?
            (AUTHENTICATED USING (PASSWORD | CERTIFICATE | DISTINGUISHED NAME))?
        )
    ;

container_names
    : LEFT_PAREN id_expression (COMMA id_expression)* RIGHT_PAREN
    ;

set_container_data
    : SET CONTAINER_DATA EQUALS_OP (ALL | DEFAULT | container_names)
    ;

add_rem_container_data
    : (ADD | REMOVE) CONTAINER_DATA EQUALS_OP container_names
    ;

container_data_clause
    : set_container_data
    | add_rem_container_data (FOR container_tableview_name)?
    ;

// https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/ADMINISTER-KEY-MANAGEMENT.html
administer_key_management
    : ADMINISTER KEY MANAGEMENT ( keystore_management_clauses
                                | key_management_clauses
                                | secret_management_clauses
                                | zero_downtime_software_patching_clauses
                                )
        SEMICOLON
    ;

keystore_management_clauses
    : create_keystore
    | open_keystore
    | close_keystore
    | backup_keystore
    | alter_keystore_password
    | merge_into_new_keystore
    | merge_into_existing_keystore
    | isolate_keystore
    | unite_keystore
    ;

create_keystore
    : CREATE ( KEYSTORE ksl=CHAR_STRING
             | LOCAL? AUTO_LOGIN KEYSTORE FROM KEYSTORE ksl=CHAR_STRING
             ) IDENTIFIED BY keystore_password
    ;

open_keystore
    : SET KEYSTORE OPEN force_keystore?
        identified_by_store
        container_clause?
    ;

force_keystore
    : FORCE KEYSTORE
    ;

close_keystore
    : SET KEYSTORE CLOSE identified_by_store?
        container_clause?
    ;

backup_keystore
    : BACKUP KEYSTORE (USING bi=CHAR_STRING)? force_keystore?
        identified_by_store
        (TO ksl=CHAR_STRING)?
    ;

alter_keystore_password
    : ALTER KEYSTORE PASSWORD force_keystore? IDENTIFIED BY o=keystore_password
        SET n=keystore_password with_backup_clause?
    ;

merge_into_new_keystore
    : MERGE KEYSTORE ksl1=CHAR_STRING identified_by_password_clause?
        AND KEYSTORE ksl2=CHAR_STRING identified_by_password_clause?
        INTO NEW KEYSTORE ksl2=CHAR_STRING identified_by_password_clause
    ;

merge_into_existing_keystore
    : MERGE KEYSTORE ksl1=CHAR_STRING identified_by_password_clause?
        INTO EXISTING KEYSTORE ksl2=CHAR_STRING identified_by_password_clause
        with_backup_clause?
    ;

isolate_keystore
    : FORCE? ISOLATE KEYSTORE IDENTIFIED BY i=keystore_password FROM ROOT KEYSTORE
        force_keystore?
        identified_by_store
        with_backup_clause?
    ;

unite_keystore
    : UNITE KEYSTORE IDENTIFIED BY i=keystore_password WITH ROOT KEYSTORE
        force_keystore?
        identified_by_store
        with_backup_clause?
    ;

key_management_clauses
    : set_key
    | create_key
    | use_key
    | set_key_tag
    | export_keys
    | import_keys
    | migrate_keys
    | reverse_migrate_keys
    | move_keys
    ;

set_key
    : SET ENCRYPTION? KEY ((mkid ':')? mk)? using_tag_clause?
        using_algorithm_clause? force_keystore?
        identified_by_store
        with_backup_clause?
        container_clause?
    ;

create_key
    : CREATE ENCRYPTION? KEY ((mkid ':')? mk)? using_tag_clause?
        using_algorithm_clause? force_keystore?
        identified_by_store
        with_backup_clause?
        container_clause?
    ;

mkid
    : CHAR_STRING
    ;

mk
    : CHAR_STRING
    ;

use_key
    : USE ENCRYPTION? KEY k=CHAR_STRING using_tag_clause? force_keystore?
        identified_by_store
        with_backup_clause?
    ;

set_key_tag
    : SET TAG t=CHAR_STRING FOR k=CHAR_STRING force_keystore?
        identified_by_store
        with_backup_clause?
    ;

export_keys
    : EXPORT ENCRYPTION? KEYS WITH SECRET secret TO filename
        force_keystore? identified_by_store
        (WITH IDENTIFIER IN (CHAR_STRING (COMMA CHAR_STRING)* | LEFT_PAREN subquery RIGHT_PAREN))?
    ;

import_keys
    : IMPORT ENCRYPTION? KEYS WITH SECRET secret FROM filename
        force_keystore?
        identified_by_store
        with_backup_clause?
    ;

migrate_keys
    : SET ENCRYPTION? KEY IDENTIFIED BY hsm=secret
        force_keystore?
        MIGRATE USING keystore_password
        with_backup_clause?
    ;

reverse_migrate_keys
    : SET ENCRYPTION? KEY IDENTIFIED BY s=secret
        force_keystore?
        REVERSE MIGRATE USING hsm=secret
    ;

move_keys
    : MOVE ENCRYPTION? KEYS TO NEW KEYSTORE ksl1=CHAR_STRING
        IDENTIFIED BY ksp1=keystore_password FROM FORCE? KEYSTORE IDENTIFIED BY ksp=keystore_password
        (WITH IDENTIFIER IN (CHAR_STRING (COMMA CHAR_STRING)* | subquery))?
        with_backup_clause?
    ;

identified_by_store
    : IDENTIFIED BY (EXTERNAL STORE | keystore_password)
    ;

using_algorithm_clause
    : USING ALGORITHM ea=CHAR_STRING
    ;

using_tag_clause
    : USING TAG t=CHAR_STRING
    ;

secret_management_clauses
    : add_update_secret
    | delete_secret
    | add_update_secret_seps
    | delete_secret_seps
    ;

add_update_secret
    : (ADD | UPDATE) SECRET s=CHAR_STRING FOR CLIENT ci=CHAR_STRING
        using_tag_clause?
        force_keystore?
        identified_by_store?
        with_backup_clause?
    ;

delete_secret
    : DELETE SECRET FOR CLIENT ci=CHAR_STRING
        force_keystore?
        identified_by_store
        with_backup_clause?
    ;

add_update_secret_seps
    : (ADD | UPDATE) SECRET s=CHAR_STRING FOR CLIENT ci=CHAR_STRING
        using_tag_clause?
        TO LOCAL? AUTO_LOGIN KEYSTORE directory_path
    ;

delete_secret_seps
    : DELETE SECRET s=CHAR_STRING SQ FOR CLIENT ci=CHAR_STRING
        FROM LOCAL? AUTO_LOGIN KEYSTORE directory_path
    ;

zero_downtime_software_patching_clauses
    : SWITCHOVER TO? LIBRARY path FOR ALL CONTAINERS //inconsistent documentation
    ;

with_backup_clause
    : WITH BACKUP (USING bi=CHAR_STRING)?
    ;

identified_by_password_clause
    : IDENTIFIED BY keystore_password
    ;

keystore_password
    : DELIMITED_ID
    ;

path
    : CHAR_STRING
    ;

secret
    : DELIMITED_ID
    ;

// https://docs.oracle.com/cd/E11882_01/server.112/e41084/statements_4005.htm#SQLRF01105
analyze
    : ( ANALYZE (TABLE tableview_name | INDEX index_name) partition_extention_clause?
      | ANALYZE CLUSTER cluster_name
      )

      ( validation_clauses
      | LIST CHAINED ROWS into_clause1?
      | DELETE SYSTEM? STATISTICS
      )
      SEMICOLON
    ;

partition_extention_clause
    : PARTITION ( LEFT_PAREN partition_name RIGHT_PAREN
                | FOR LEFT_PAREN partition_key_value (COMMA partition_key_value)* RIGHT_PAREN
                )
    | SUBPARTITION ( LEFT_PAREN subpartition_name RIGHT_PAREN
                   | FOR LEFT_PAREN subpartition_key_value (COMMA subpartition_key_value)* RIGHT_PAREN
                   )
    ;

validation_clauses
    : VALIDATE REF UPDATE (SET DANGLING TO NULL_)?
    | VALIDATE STRUCTURE
        ( CASCADE FAST
        | CASCADE online_or_offline? into_clause?
        | CASCADE
        )?
        online_or_offline? into_clause?
    ;

compute_clauses
    : COMPUTE SYSTEM? STATISTICS for_clause?
    ;
for_clause
     : FOR ( TABLE for_clause*
           | ALL (INDEXED? COLUMNS (SIZE UNSIGNED_INTEGER)? for_clause* | LOCAL? INDEXES)
           | COLUMNS (SIZE UNSIGNED_INTEGER)? (column_name SIZE UNSIGNED_INTEGER)+ for_clause*
           )
    ;

online_or_offline
    : OFFLINE
    | ONLINE
    ;

into_clause1
    : INTO tableview_name?
    ;

//Making assumption on partition ad subpartition key value clauses
partition_key_value
    : literal
    | TIMESTAMP quoted_string
    ;

subpartition_key_value
    : literal
    | TIMESTAMP quoted_string
    ;

//https://docs.oracle.com/cd/E11882_01/server.112/e41084/statements_4006.htm#SQLRF01106
associate_statistics
    : ASSOCIATE STATISTICS
        WITH (column_association | function_association)
        storage_table_clause?
      SEMICOLON
    ;

column_association
    : COLUMNS tableview_name PERIOD column_name (COMMA tableview_name PERIOD column_name)* using_statistics_type
    ;

function_association
    : ( FUNCTIONS function_name (COMMA function_name)*
      | PACKAGES package_name (COMMA package_name)*
      | TYPES type_name (COMMA type_name)*
      | INDEXES index_name (COMMA index_name)*
      | INDEXTYPES indextype_name (COMMA indextype_name)*
      )

      ( using_statistics_type
      | default_cost_clause (COMMA default_selectivity_clause)?
      | default_selectivity_clause (COMMA default_cost_clause)?
      )
    ;

indextype_name
    : id_expression
    ;


using_statistics_type
    : USING (statistics_type_name | NULL_)
    ;

statistics_type_name
    : regular_id
    ;

default_cost_clause
    : DEFAULT COST LEFT_PAREN cpu_cost COMMA io_cost COMMA network_cost RIGHT_PAREN
    ;

cpu_cost
    : UNSIGNED_INTEGER
    ;

io_cost
    : UNSIGNED_INTEGER
    ;

network_cost
    : UNSIGNED_INTEGER
    ;

default_selectivity_clause
    : DEFAULT SELECTIVITY default_selectivity
    ;

default_selectivity
    : UNSIGNED_INTEGER
    ;

storage_table_clause
    : WITH (SYSTEM | USER) MANAGED STORAGE TABLES
    ;

// https://docs.oracle.com/database/121/SQLRF/statements_4008.htm#SQLRF56110
unified_auditing
    : {p.isVersion12()}?
      AUDIT (POLICY policy_name ((BY | EXCEPT) audit_user (COMMA audit_user)* )?
                                (WHENEVER NOT? SUCCESSFUL)?
            | CONTEXT NAMESPACE oracle_namespace
                      ATTRIBUTES attribute_name (COMMA attribute_name)* (BY audit_user (COMMA audit_user)*)?
            )
      SEMICOLON
    ;

policy_name
    : identifier
    ;

// https://docs.oracle.com/cd/E11882_01/server.112/e41084/statements_4007.htm#SQLRF01107
// https://docs.oracle.com/database/121/SQLRF/statements_4007.htm#SQLRF01107

audit_traditional
    : AUDIT ( audit_operation_clause (auditing_by_clause | IN SESSION CURRENT)?
            | audit_schema_object_clause
            | NETWORK
            | audit_direct_path
            )
        (BY (SESSION | ACCESS) )? (WHENEVER NOT? SUCCESSFUL)?
        audit_container_clause?
      SEMICOLON
    ;

audit_direct_path
    : {p.isVersion12()}? DIRECT_PATH auditing_by_clause
    ;

audit_container_clause
    : {p.isVersion12()}? (CONTAINER EQUALS_OP (CURRENT | ALL))
    ;

audit_operation_clause
    : ( (sql_statement_shortcut | ALL STATEMENTS?)  (COMMA (sql_statement_shortcut | ALL STATEMENTS?) )*
      | (system_privilege | ALL PRIVILEGES)  (COMMA (system_privilege | ALL PRIVILEGES) )*
      )
    ;

auditing_by_clause
    : BY audit_user (COMMA audit_user)*
    ;

audit_user
    : regular_id
    ;

audit_schema_object_clause
    : ( sql_operation (COMMA sql_operation)* | ALL) auditing_on_clause
    ;

sql_operation
    : ALTER
    | AUDIT
    | COMMENT
    | DELETE
    | EXECUTE
    | FLASHBACK
    | GRANT
    | INDEX
    | INSERT
    | LOCK
    | READ
    | RENAME
    | SELECT
    | UPDATE
    ;

auditing_on_clause
    : ON ( object_name
         | DIRECTORY regular_id
         | MINING MODEL model_name
         | {p.isVersion12()}? SQL TRANSLATION PROFILE profile_name
         | DEFAULT
         )
    ;

model_name
    : (id_expression PERIOD)? id_expression
    ;

object_name
    : (id_expression PERIOD)? id_expression
    ;

profile_name
    : (id_expression PERIOD)? id_expression
    ;

sql_statement_shortcut
    : ALTER SYSTEM
    | CLUSTER
    | CONTEXT
    | DATABASE LINK
    | DIMENSION
    | DIRECTORY
    | INDEX
    | MATERIALIZED VIEW
    | NOT EXISTS
    | OUTLINE
    | {p.isVersion12()}? PLUGGABLE DATABASE
    | PROCEDURE
    | PROFILE
    | PUBLIC DATABASE LINK
    | PUBLIC SYNONYM
    | ROLE
    | ROLLBACK SEGMENT
    | SEQUENCE
    | SESSION
    | SYNONYM
    | SYSTEM AUDIT
    | SYSTEM GRANT
    | TABLE
    | TABLESPACE
    | TRIGGER
    | TYPE
    | USER
    | VIEW
    | ALTER SEQUENCE
    | ALTER TABLE
    | COMMENT TABLE
    | DELETE TABLE
    | EXECUTE PROCEDURE
    | GRANT DIRECTORY
    | GRANT PROCEDURE
    | GRANT SEQUENCE
    | GRANT TABLE
    | GRANT TYPE
    | INSERT TABLE
    | LOCK TABLE
    | SELECT SEQUENCE
    | SELECT TABLE
    | UPDATE TABLE
    ;

drop_index
    : DROP INDEX index_name SEMICOLON
    ;

// https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/DISASSOCIATE-STATISTICS.html
disassociate_statistics
    : DISASSOCIATE STATISTICS FROM
        ( COLUMNS (schema_name PERIOD)? tb=id_expression PERIOD c=id_expression (COMMA (schema_name PERIOD)? tb=id_expression PERIOD c=id_expression)*
        | FUNCTIONS (schema_name PERIOD)? fn=id_expression (COMMA (schema_name PERIOD)? fn=id_expression)*
        | PACKAGES (schema_name PERIOD)? pkg=id_expression (COMMA (schema_name PERIOD)? pkg=id_expression)*
        | TYPES (schema_name PERIOD)? t=id_expression (COMMA (schema_name PERIOD)? t=id_expression)*
        | INDEXES (schema_name PERIOD)? ix=id_expression (COMMA (schema_name PERIOD)? ix=id_expression)*
        | INDEXTYPES (schema_name PERIOD)? it=id_expression (COMMA (schema_name PERIOD)? it=id_expression)*
        )
        FORCE?
    ;

// https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/DROP-INDEXTYPE.html
drop_indextype
    : DROP INDEXTYPE (schema_name PERIOD)? it=id_expression FORCE?
    ;

// https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/DROP-INMEMORY-JOIN-GROUP.html
drop_inmemory_join_group
    : DROP INMEMORY JOIN GROUP (schema_name PERIOD)? jg=id_expression
    ;

flashback_table
    : FLASHBACK TABLE tableview_name (COMMA tableview_name)* TO
      ( ((SCN | TIMESTAMP) expression | RESTORE POINT restore_point) ((ENABLE | DISABLE) TRIGGERS)?
        | BEFORE DROP (RENAME TO tableview_name)?
      )
    ;

restore_point
    : identifier (PERIOD id_expression)*
    ;

// https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/PURGE.html
purge_statement
    : PURGE ( (TABLE | INDEX) id_expression
            | TABLESPACE SET? ts=id_expression (USER u=id_expression)?
            | RECYCLEBIN
            | DBA_RECYCLEBIN
            )
    ;

// https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/NOAUDIT-Traditional-Auditing.html
noaudit_statement
    : NOAUDIT
        ( audit_operation_clause auditing_by_clause?
        | audit_schema_object_clause
        | NETWORK
        | DIRECT_PATH LOAD auditing_by_clause?
        )
        (WHENEVER NOT? SUCCESSFUL)?
        container_clause?
    ;

rename_object
    : RENAME object_name TO object_name SEMICOLON
    ;

grant_statement
    : GRANT
        ( COMMA?
          (object_privilege paren_column_list?
          | system_privilege
          | role_name
          )
        )+
      (ON grant_object_name)?
      TO (grantee_name | PUBLIC) (COMMA (grantee_name | PUBLIC) )*
      (WITH (ADMIN | DELEGATE) OPTION)?
      (WITH HIERARCHY OPTION)?
      (WITH GRANT OPTION)?
      container_clause? SEMICOLON
    ;

container_clause
    : CONTAINER EQUALS_OP (CURRENT | ALL)
    ;

// https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/REVOKE.html
revoke_statement
    : REVOKE ((revoke_object_privileges | revoke_system_privilege) container_clause? | revoke_roles_from_programs)
    ;

revoke_system_privilege
    : (system_privilege | role_name | ALL PRIVILEGES) FROM revokee_clause
    ;

revokee_clause
    : (id_expression | PUBLIC) (COMMA (id_expression | PUBLIC))*
    ;

revoke_object_privileges
    : (object_privilege | ALL PRIVILEGES?) (COMMA (object_privilege | ALL PRIVILEGES?))* on_object_clause
        FROM revokee_clause (CASCADE CONSTRAINTS | FORCE)?
    ;

on_object_clause
    : ON ( (schema_name PERIOD)? o=id_expression
         | USER id_expression (COMMA id_expression)*
         | DIRECTORY directory_name
         | EDITION edition_name
         | MINING MODEL (schema_name PERIOD)? mmn=id_expression
         | JAVA (SOURCE | RESOURCE) (schema_name PERIOD)? o2=id_expression
         | SQL TRANSLATION PROFILE (schema_name PERIOD)? p=id_expression
         )
    ;

revoke_roles_from_programs
    : (role_name (COMMA role_name)* | ALL) FROM program_unit (COMMA program_unit)*
    ;

program_unit
    : (FUNCTION | PROCEDURE | PACKAGE) (schema_name PERIOD)? id_expression
    ;

create_dimension
    : CREATE DIMENSION identifier  level_clause+ (hierarchy_clause | attribute_clause | extended_attribute_clause)+
    ;

// https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/CREATE-DIRECTORY.html
create_directory
    : CREATE (OR REPLACE)? DIRECTORY directory_name
        (SHARING EQUALS_OP (METADATA | NONE))?
        AS directory_path
      SEMICOLON
    ;

directory_name
    : id_expression
    ;

directory_path
    : CHAR_STRING
    ;

// https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/CREATE-INMEMORY-JOIN-GROUP.html
create_inmemory_join_group
    : CREATE INMEMORY JOIN GROUP (schema_name PERIOD)? jg=id_expression
        LEFT_PAREN (schema_name PERIOD)? t=id_expression LEFT_PAREN c=id_expression RIGHT_PAREN (COMMA (schema_name PERIOD)? t=id_expression LEFT_PAREN c=id_expression RIGHT_PAREN)+ RIGHT_PAREN
    ;

// https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/DROP-HIERARCHY.html
drop_hierarchy
    : DROP HIERARCHY (schema_name PERIOD)? hn=id_expression
    ;

// https://docs.oracle.com/cd/E11882_01/appdev.112/e25519/alter_library.htm#LNPLS99946
// https://docs.oracle.com/database/121/LNPLS/alter_library.htm#LNPLS99946
alter_library
    : ALTER LIBRARY library_name
       ( COMPILE library_debug? compiler_parameters_clause* (REUSE SETTINGS)?
       | library_editionable
       )
     SEMICOLON
    ;

// https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/DROP-JAVA.html
drop_java
    : DROP JAVA (SOURCE | CLASS | RESOURCE) (schema_name PERIOD)? id_expression
    ;

drop_library
    : DROP LIBRARY library_name
    ;

// https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/CREATE-JAVA.html
create_java
    : CREATE (OR REPLACE)? (AND (RESOLVE | COMPILE))? NOFORCE?
        JAVA ( (SOURCE | RESOURCE) NAMED (schema_name PERIOD)? pn=id_expression
             | CLASS (SCHEMA id_expression)?
             )
        (SHARING EQUALS_OP (METADATA | NONE))?
        invoker_rights_clause?
        (RESOLVER LEFT_PAREN (LEFT_PAREN CHAR_STRING COMMA? (sn=id_expression | '-') RIGHT_PAREN)+ RIGHT_PAREN)?
        ( USING (BFILE LEFT_PAREN d=id_expression COMMA filename RIGHT_PAREN | (CLOB | BLOB | BFILE) subquery | CHAR_STRING)
        | AS CHAR_STRING
        )
    ;

create_library
    : CREATE (OR REPLACE)? (EDITIONABLE | NONEDITIONABLE)? LIBRARY plsql_library_source
    ;

plsql_library_source
    : library_name (IS | AS) quoted_string (IN directory_name)?
        (AGENT quoted_string)? (CREDENTIAL credential_name)?
    ;

credential_name
    : (id_expression PERIOD)? id_expression
    ;

library_editionable
    : {p.isVersion12()}? (EDITIONABLE | NONEDITIONABLE)
    ;

library_debug
    : {p.isVersion12()}? DEBUG
    ;


compiler_parameters_clause
    : parameter_name EQUALS_OP parameter_value
    ;

parameter_value
    : regular_id
    | CHAR_STRING
    ;

library_name
    : (regular_id PERIOD)? regular_id
    ;

alter_dimension
    : ALTER DIMENSION identifier
    ( (ADD (level_clause | hierarchy_clause | attribute_clause |  extended_attribute_clause))+
     | (DROP (LEVEL identifier (RESTRICT | CASCADE)? | HIERARCHY identifier | ATTRIBUTE identifier (LEVEL identifier (COLUMN column_name (COMMA COLUMN column_name)*)?)? ))+
     | COMPILE
    )
    ;

level_clause
    : LEVEL identifier IS (table_name PERIOD column_name | LEFT_PAREN table_name PERIOD column_name (COMMA table_name PERIOD column_name)* RIGHT_PAREN)
        (SKIP_ WHEN NULL_)?
    ;

hierarchy_clause
    : HIERARCHY identifier LEFT_PAREN identifier (CHILD OF identifier)+ dimension_join_clause? RIGHT_PAREN
    ;

dimension_join_clause
    : (JOIN KEY column_one_or_more_sub_clause REFERENCES identifier)+
    ;

attribute_clause
    : (ATTRIBUTE identifier DETERMINES column_one_or_more_sub_clause)+
    ;

extended_attribute_clause
    : ATTRIBUTE identifier (LEVEL identifier DETERMINES column_one_or_more_sub_clause )+
    ;

column_one_or_more_sub_clause
    : column_name
    | LEFT_PAREN column_name (COMMA column_name)* RIGHT_PAREN
    ;

// https://docs.oracle.com/cd/E11882_01/server.112/e41084/statements_4004.htm#SQLRF01104
// https://docs.oracle.com/database/121/SQLRF/statements_4004.htm#SQLRF01104
alter_view
    : ALTER VIEW tableview_name
       ( ADD out_of_line_constraint
       | MODIFY CONSTRAINT constraint_name (RELY | NORELY)
       | DROP ( CONSTRAINT constraint_name
              | PRIMARY KEY
              | UNIQUE LEFT_PAREN column_name (COMMA column_name)* RIGHT_PAREN
              )
       | COMPILE
       | READ (ONLY | WRITE)
       | alter_view_editionable?
       )
      SEMICOLON
    ;

alter_view_editionable
    : {p.isVersion12()}? (EDITIONABLE | NONEDITIONABLE)
    ;

// https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/CREATE-VIEW.html
create_view
    : CREATE (OR REPLACE)? (NO? FORCE)? editioning_clause? VIEW (schema_name PERIOD)? v=id_expression
      (SHARING EQUALS_OP (METADATA | EXTENDED? DATA | NONE))?
      view_options?
      (DEFAULT COLLATION cn=id_expression)?
      (BEQUEATH (CURRENT_USER | DEFINER))?
      AS select_only_statement subquery_restriction_clause?
      (CONTAINER_MAP | CONTAINERS_DEFAULT)?
    ;

editioning_clause
    : EDITIONING
    | EDITIONABLE EDITIONING?
    | NONEDITIONABLE
    ;

view_options
    : view_alias_constraint
    | object_view_clause
    | xmltype_view_clause
    ;

view_alias_constraint
    : LEFT_PAREN ( COMMA? (table_alias inline_constraint* | out_of_line_constraint) )+ RIGHT_PAREN
    ;

object_view_clause
    : OF (schema_name PERIOD)? tn=id_expression
       ( WITH OBJECT (IDENTIFIER | ID) (DEFAULT | LEFT_PAREN REGULAR_ID (COMMA REGULAR_ID)* RIGHT_PAREN)
       | UNDER (schema_name PERIOD)? sv=id_expression
       )
       (LEFT_PAREN (COMMA? (out_of_line_constraint | REGULAR_ID inline_constraint))+ RIGHT_PAREN)*
    ;

inline_constraint
    : (CONSTRAINT constraint_name)?
        ( NOT? NULL_
        | UNIQUE
        | PRIMARY KEY
        | references_clause
        | check_constraint
        )
      constraint_state?
    ;

inline_ref_constraint
    : SCOPE IS tableview_name
    | WITH ROWID
    | (CONSTRAINT constraint_name)? references_clause constraint_state?
    ;

out_of_line_ref_constraint
    : SCOPE FOR LEFT_PAREN ref_col_or_attr=regular_id RIGHT_PAREN IS tableview_name
    | REF LEFT_PAREN ref_col_or_attr=regular_id RIGHT_PAREN WITH ROWID
    | (CONSTRAINT constraint_name)? FOREIGN KEY LEFT_PAREN ( COMMA? ref_col_or_attr=regular_id)+ RIGHT_PAREN references_clause constraint_state?
    ;

out_of_line_constraint
    : ( (CONSTRAINT constraint_name)?
          ( UNIQUE LEFT_PAREN column_name (COMMA column_name)* RIGHT_PAREN
          | PRIMARY KEY LEFT_PAREN column_name (COMMA column_name)* RIGHT_PAREN
          | foreign_key_clause
          | CHECK LEFT_PAREN condition RIGHT_PAREN
          )
       )
      constraint_state?
    ;

constraint_state
    : ( NOT? DEFERRABLE
      | INITIALLY (IMMEDIATE | DEFERRED)
      | (RELY | NORELY)
      | (ENABLE | DISABLE)
      | (VALIDATE | NOVALIDATE)
      | using_index_clause
      )+
    ;

xmltype_view_clause
    : OF XMLTYPE xml_schema_spec? WITH OBJECT (IDENTIFIER | ID) (DEFAULT | LEFT_PAREN expression (COMMA expression)* RIGHT_PAREN)
    ;

xml_schema_spec
    : (XMLSCHEMA xml_schema_url)? ELEMENT (element | xml_schema_url '#' element)
        (STORE ALL VARRAYS AS (LOBS | TABLES))?
        (allow_or_disallow NONSCHEMA)?
        (allow_or_disallow ANYSCHEMA)?
    ;

xml_schema_url
    : DELIMITED_ID
    ;

element
    : DELIMITED_ID
    ;

alter_tablespace
    : ALTER TABLESPACE tablespace
       ( DEFAULT table_compression? storage_clause?
       | MINIMUM EXTENT size_clause
       | RESIZE size_clause
       | COALESCE
       | SHRINK SPACE_KEYWORD (KEEP size_clause)?
       | RENAME TO new_tablespace_name
       | begin_or_end BACKUP
       | datafile_tempfile_clauses
       | tablespace_logging_clauses
       | tablespace_group_clause
       | tablespace_state_clauses
       | autoextend_clause
       | flashback_mode_clause
       | tablespace_retention_clause
       )
     SEMICOLON
    ;

datafile_tempfile_clauses
    : ADD (datafile_specification | tempfile_specification)
    | DROP (DATAFILE | TEMPFILE) (filename | UNSIGNED_INTEGER) (KEEP size_clause)?
    | SHRINK TEMPFILE (filename | UNSIGNED_INTEGER) (KEEP size_clause)?
    | RENAME DATAFILE filename (COMMA filename)* TO filename (COMMA filename)*
    | (DATAFILE | TEMPFILE) (online_or_offline)
    ;

tablespace_logging_clauses
    : logging_clause
    | NO? FORCE LOGGING
    ;

tablespace_group_clause
    : TABLESPACE GROUP (tablespace_group_name | CHAR_STRING)
    ;

tablespace_group_name
    : regular_id
    ;

tablespace_state_clauses
    : ONLINE
    | OFFLINE (NORMAL | TEMPORARY | IMMEDIATE)?
    | READ (ONLY | WRITE)
    | PERMANENT
    | TEMPORARY
    ;

flashback_mode_clause
    : FLASHBACK (ON | OFF)
    ;

new_tablespace_name
    : tablespace
    ;

create_tablespace
    : CREATE (BIGFILE | SMALLFILE)?
        ( permanent_tablespace_clause
        | temporary_tablespace_clause
        | undo_tablespace_clause
        )
      SEMICOLON
    ;

permanent_tablespace_clause
    : TABLESPACE id_expression datafile_specification?
        ( MINIMUM EXTENT size_clause
        | BLOCKSIZE size_clause
        | logging_clause
        | FORCE LOGGING
        | (ONLINE | OFFLINE)
        | ENCRYPTION tablespace_encryption_spec
        | DEFAULT //TODO table_compression? storage_clause?
        | extent_management_clause
        | segment_management_clause
        | flashback_mode_clause
        )*
    ;

tablespace_encryption_spec
    : USING encrypt_algorithm=CHAR_STRING
    ;

logging_clause
    : LOGGING
     | NOLOGGING
     | FILESYSTEM_LIKE_LOGGING
    ;

extent_management_clause
    : EXTENT MANAGEMENT LOCAL
        ( AUTOALLOCATE
        | UNIFORM (SIZE size_clause)?
        )?
    ;

segment_management_clause
    : SEGMENT SPACE_KEYWORD MANAGEMENT (AUTO | MANUAL)
    ;

temporary_tablespace_clause
    : TEMPORARY TABLESPACE tablespace_name=id_expression
        tempfile_specification?
        tablespace_group_clause? extent_management_clause?
    ;

undo_tablespace_clause
    : UNDO TABLESPACE tablespace_name=id_expression
        datafile_specification?
        extent_management_clause? tablespace_retention_clause?
    ;

tablespace_retention_clause
    : RETENTION (GUARANTEE | NOGUARANTEE)
    ;

// https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/CREATE-TABLESPACE-SET.html
create_tablespace_set
    : CREATE TABLESPACE SET tss=id_expression (IN SHARDSPACE ss=id_expression)?
        (USING TEMPLATE LEFT_PAREN (DATAFILE file_specification (COMMA file_specification)*)? permanent_tablespace_attrs+ RIGHT_PAREN)?
    ;

permanent_tablespace_attrs
    : MINIMUM EXTENT size_clause
    | BLOCKSIZE numeric K_LETTER?
    | logging_clause
    | FORCE LOGGING
    | tablespace_encryption_clause
    | default_tablespace_params
    | ONLINE
    | OFFLINE
    | extent_management_clause
    | segment_management_clause
    | flashback_mode_clause
    | lost_write_protection
    ;

tablespace_encryption_clause
    : ENCRYPTION (tablespace_encryption_spec? ENCRYPT | DECRYPT)
    ;

default_tablespace_params
    : DEFAULT default_table_compression? default_index_compression? inmmemory_clause? ilm_clause? storage_clause?
    ;

default_table_compression
    : TABLE ( COMPRESS FOR (OLTP | QUERY low_high | ARCHIVE low_high)
            | NOCOMPRESS
            )
    ;

low_high
    : LOW
    | HIGH
    ;

default_index_compression
    : INDEX (COMPRESS ADVANCED low_high | NOCOMPRESS)
    ;

inmmemory_clause
    : INMEMORY inmemory_attributes? (TEXT (column_name (COMMA column_name)* | column_name USING policy_name (COMMA column_name USING policy_name)*))?
    | NO INMEMORY
    ;

// asm_filename is just a charater string.  Would need to parse the string
// to find diskgroup...
datafile_specification
    : DATAFILE
	  (COMMA? datafile_tempfile_spec)
    ;

tempfile_specification
    : TEMPFILE
	  (COMMA? datafile_tempfile_spec)
    ;

datafile_tempfile_spec
    : CHAR_STRING? (SIZE size_clause)? REUSE? autoextend_clause?
    ;

redo_log_file_spec
    : ( filename
      | LEFT_PAREN filename (COMMA filename)* RIGHT_PAREN
      )
        (SIZE size_clause)?
        (BLOCKSIZE size_clause)?
        REUSE?
    ;

autoextend_clause
    : AUTOEXTEND (OFF | ON (NEXT size_clause)? maxsize_clause? )
    ;

maxsize_clause
    : MAXSIZE (UNLIMITED | size_clause)
    ;

build_clause
    : BUILD (IMMEDIATE | DEFERRED)
    ;

partial_index_clause
    : INDEXING (PARTIAL | FULL)
    ;

parallel_clause
    : NOPARALLEL
    | PARALLEL parallel_count=UNSIGNED_INTEGER?
    ;

alter_materialized_view
    : ALTER MATERIALIZED VIEW tableview_name
       ( physical_attributes_clause
       | modify_mv_column_clause
       | table_compression
       | lob_storage_clause (COMMA lob_storage_clause)*
       | modify_lob_storage_clause (COMMA modify_lob_storage_clause)*
//TODO | alter_table_partitioning
       | parallel_clause
       | logging_clause
       | allocate_extent_clause
       | deallocate_unused_clause
       | shrink_clause
       | (cache_or_nocache)
       )?
       alter_iot_clauses?
       (USING INDEX physical_attributes_clause)?
       alter_mv_option1?
       ( enable_or_disable QUERY REWRITE
       | COMPILE
       | CONSIDER FRESH
       )?
     SEMICOLON
    ;

alter_mv_option1
    : alter_mv_refresh
//TODO  | MODIFY scoped_table_ref_constraint
    ;

alter_mv_refresh
    : REFRESH ( FAST
              | COMPLETE
              | FORCE
              | ON (DEMAND | COMMIT)
              | START WITH expression
              | NEXT expression
              | WITH PRIMARY KEY
              | USING DEFAULT? MASTER ROLLBACK SEGMENT rollback_segment?
              | USING (ENFORCED | TRUSTED) CONSTRAINTS
              )+
    ;

rollback_segment
    : regular_id
    ;

modify_mv_column_clause
    : MODIFY LEFT_PAREN column_name (ENCRYPT encryption_spec | DECRYPT)? RIGHT_PAREN
    ;

alter_materialized_view_log
    : ALTER MATERIALIZED VIEW LOG FORCE? ON tableview_name
       ( physical_attributes_clause
       | add_mv_log_column_clause
//TODO | alter_table_partitioning
       | parallel_clause
       | logging_clause
       | allocate_extent_clause
       | shrink_clause
       | move_mv_log_clause
       | cache_or_nocache
       )?
       mv_log_augmentation? mv_log_purge_clause?
      SEMICOLON
    ;
add_mv_log_column_clause
    : ADD LEFT_PAREN column_name RIGHT_PAREN
    ;

move_mv_log_clause
    : MOVE segment_attributes_clause parallel_clause?
    ;

mv_log_augmentation
    : ADD ( ( OBJECT ID
            | PRIMARY KEY
            | ROWID
            | SEQUENCE
            )
            (LEFT_PAREN column_name (COMMA column_name)* RIGHT_PAREN)?

          | LEFT_PAREN column_name (COMMA column_name)* RIGHT_PAREN
          )
          new_values_clause?
    ;

// Should bound this to just date/time expr
datetime_expr
    : expression
    ;

create_materialized_view_log
    : CREATE MATERIALIZED VIEW LOG ON tableview_name
        ( ( physical_attributes_clause
          | TABLESPACE tablespace_name=id_expression
          | logging_clause
          | (CACHE | NOCACHE)
          )+
         )?
        parallel_clause?
        // table_partitioning_clauses TODO
        ( WITH
           ( COMMA?
             ( OBJECT ID
             | PRIMARY KEY
             | ROWID
             | SEQUENCE
             | COMMIT SCN
             )
           )*
           (LEFT_PAREN ( COMMA? regular_id )+ RIGHT_PAREN new_values_clause? )?
           mv_log_purge_clause?
        )*
    ;

new_values_clause
    : (INCLUDING | EXCLUDING ) NEW VALUES
    ;

mv_log_purge_clause
    : PURGE
         ( IMMEDIATE (SYNCHRONOUS | ASYNCHRONOUS)?
      // |START WITH CLAUSES TODO
         )
    ;

create_materialized_zonemap
    : CREATE MATERIALIZED ZONEMAP zonemap_name (LEFT_PAREN column_list RIGHT_PAREN)? zonemap_attributes? zonemap_refresh_clause?
        ((ENABLE | DISABLE) PRUNING)? (create_zonemap_on_table | create_zonemap_as_subquery)
    ;

alter_materialized_zonemap
    : ALTER MATERIALIZED ZONEMAP zonemap_name
        ( zonemap_attributes | zonemap_refresh_clause| (ENABLE | DISABLE) PRUNING | COMPILE | REBUILD | UNUSABLE
        )
    ;

drop_materialized_zonemap
    : DROP MATERIALIZED ZONEMAP zonemap_name
    ;

zonemap_refresh_clause
    : REFRESH (FAST | COMPILE | FORCE)? (ON (DEMAND | COMMIT | LOAD | DATA MOVEMENT | LOAD DATA MOVEMENT))?
    ;

zonemap_attributes
    : (PCTFREE numeric | PCTUSED numeric | SCALE numeric | TABLESPACE tablespace | (CACHE | NOCACHE))+
    ;

zonemap_name
    : identifier (PERIOD id_expression)?
    ;

operator_name
    : identifier (PERIOD id_expression)?
    ;

operator_function_name
    : identifier (PERIOD id_expression)*
    ;

create_zonemap_on_table
    : ON tableview_name LEFT_PAREN column_list RIGHT_PAREN
    ;

create_zonemap_as_subquery
    : AS subquery
    ;

alter_operator
    : ALTER OPERATOR operator_name (add_binding_clause | drop_binding_clause | COMPILE)
    ;

drop_operator
    : DROP OPERATOR operator_name FORCE?
    ;

create_operator
    : CREATE (OR REPLACE)? OPERATOR operator_name BINDING binding_clause (COMMA binding_clause)*
        (SHARING EQUALS_OP (METADATA | NONE))?
    ;

binding_clause
    : LEFT_PAREN datatype (COMMA datatype)* RIGHT_PAREN
          RETURN LEFT_PAREN? datatype RIGHT_PAREN? implementation_clause? using_function_clause
    ;

add_binding_clause
    : ADD BINDING binding_clause
    ;

implementation_clause
    : ANCILLARY TO primary_operator_list
    | operator_context_clause
    ;

primary_operator_list
    : primary_operator_item (COMMA primary_operator_item)*
    ;

primary_operator_item
    : schema_object_name LEFT_PAREN datatype (COMMA datatype)* RIGHT_PAREN
    ;

operator_context_clause
    : WITH INDEX CONTEXT COMMA SCAN CONTEXT implementation_type_name (COMPUTE ANCILLARY DATA)?
        (WITH COLUMN CONTEXT)?
    ;

using_function_clause
    : USING operator_function_name
    ;

drop_binding_clause
    : DROP BINDING LEFT_PAREN datatype (COMMA datatype)* RIGHT_PAREN FORCE?
    ;

create_materialized_view
    : CREATE MATERIALIZED VIEW tableview_name
      (OF type_name )?
        ( LEFT_PAREN (scoped_table_ref_constraint | mv_column_alias) (COMMA (scoped_table_ref_constraint | mv_column_alias))* RIGHT_PAREN )?
        ( ON PREBUILT TABLE ( (WITH | WITHOUT) REDUCED PRECISION)?
        | physical_properties?  (CACHE | NOCACHE)? parallel_clause? build_clause?
        )
        ( USING INDEX ( (physical_attributes_clause | TABLESPACE mv_tablespace=id_expression)+ )*
        | USING NO INDEX
        )?
        create_mv_refresh?
        (FOR UPDATE)?
        ( (DISABLE | ENABLE) QUERY REWRITE )?
        AS select_only_statement
        SEMICOLON
    ;

scoped_table_ref_constraint
    : SCOPE FOR LEFT_PAREN ref_column_or_attribute=identifier RIGHT_PAREN IS (schema_name PERIOD)? scope_table_name_or_c_alias=identifier
    ;

mv_column_alias
    : (identifier | quoted_string) (ENCRYPT encryption_spec)?
    ;

create_mv_refresh
    : ( NEVER REFRESH
      | REFRESH
         ( (FAST | COMPLETE | FORCE)
         | ON (DEMAND | COMMIT)
         | (START WITH | NEXT) //date goes here TODO
         | WITH (PRIMARY KEY | ROWID)
         | USING
             ( DEFAULT (MASTER | LOCAL)? ROLLBACK SEGMENT
             | (MASTER | LOCAL)? ROLLBACK SEGMENT rb_segment=REGULAR_ID
             )
         | USING (ENFORCED | TRUSTED) CONSTRAINTS
         )+
      )
    ;

drop_materialized_view
    : DROP MATERIALIZED VIEW tableview_name (PRESERVE TABLE)?
        SEMICOLON
    ;

create_context
    : CREATE (OR REPLACE)? CONTEXT oracle_namespace USING (schema_object_name PERIOD)? package_name
           (INITIALIZED (EXTERNALLY | GLOBALLY)
           | ACCESSED GLOBALLY
           )?
      SEMICOLON
    ;

oracle_namespace
    : id_expression
    ;

//https://docs.oracle.com/cd/E11882_01/server.112/e41084/statements_5001.htm#SQLRF01201
create_cluster
    : CREATE CLUSTER  cluster_name LEFT_PAREN column_name datatype SORT? (COMMA column_name datatype SORT?)* RIGHT_PAREN
          ( physical_attributes_clause
          | SIZE size_clause
          | TABLESPACE tablespace
          | INDEX
          | (SINGLE TABLE)? HASHKEYS UNSIGNED_INTEGER (HASH IS expression)?
          )*
          parallel_clause? (ROWDEPENDENCIES | NOROWDEPENDENCIES)?
          (CACHE | NOCACHE)?
          SEMICOLON
    ;

// https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/CREATE-PROFILE.html
create_profile
    : CREATE MANDATORY? PROFILE p=id_expression
        LIMIT (resource_parameters | password_parameters)+
        container_clause?
    ;

resource_parameters
    : ( SESSIONS_PER_USER
      | CPU_PER_SESSION
      | CPU_PER_CALL
      | CONNECT_TIME
      | IDLE_TIME
      | LOGICAL_READS_PER_SESSION
      | LOGICAL_READS_PER_CALL
      | COMPOSITE_LIMIT
      ) (UNSIGNED_INTEGER | UNLIMITED | DEFAULT)
      | PRIVATE_SGA (size_clause | UNLIMITED | DEFAULT)
    ;

password_parameters
    : ( FAILED_LOGIN_ATTEMPTS
      | PASSWORD_LIFE_TIME
      | PASSWORD_REUSE_TIME
      | PASSWORD_REUSE_MAX
      | PASSWORD_LOCK_TIME
      | PASSWORD_GRACE_TIME
      | INACTIVE_ACCOUNT_TIME
      ) (expression | UNLIMITED | DEFAULT)
      | PASSWORD_VERIFY_FUNCTION (function_name | NULL_ | DEFAULT)
      | PASSWORD_ROLLOVER_TIME (expression | DEFAULT)
    ;

// https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/CREATE-LOCKDOWN-PROFILE.html
create_lockdown_profile
    : CREATE LOCKDOWN PROFILE id_expression (static_base_profile | dynamic_base_profile)?
    ;

static_base_profile
    : FROM bp=id_expression
    ;

dynamic_base_profile
    : INCLUDING bp=id_expression
    ;

// https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/CREATE-OUTLINE.html
create_outline
    : CREATE (OR REPLACE)? (PUBLIC | PRIVATE)? OUTLINE (o=id_expression)?
        (FROM (PUBLIC | PRIVATE)? so=id_expression)?
        (FOR CATEGORY c=id_expression)?
        (ON statement)?
    ;

// https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/CREATE-RESTORE-POINT.html
create_restore_point
    : CREATE CLEAN? RESTORE POINT rp=id_expression (FOR PLUGGABLE DATABASE pdb=id_expression)?
        (AS OF (TIMESTAMP | SCN) expression)?
        (PRESERVE | GUARANTEE FLASHBACK DATABASE)?
    ;

create_role
    : CREATE ROLE role_name role_identified_clause? container_clause?
    ;

create_table
    : CREATE
            ( (GLOBAL | PRIVATE) TEMPORARY
            | SHARDED
            | DUPLICATED
            | IMMUTABLE? BLOCKCHAIN
            | IMMUTABLE
            )?
        TABLE (schema_name PERIOD)? table_name
        (SHARING EQUALS_OP (METADATA | EXTENDED? DATA | NONE))?
        (relational_table | object_table | xmltype_table)
        (MEMOPTIMIZE FOR READ)?
        (MEMOPTIMIZE FOR WRITE)?
        (PARENT tableview_name)?
      SEMICOLON
    ;

xmltype_table
    : OF XMLTYPE (LEFT_PAREN object_properties RIGHT_PAREN)?
         (XMLTYPE xmltype_storage)?
         xmlschema_spec?
         xmltype_virtual_columns?
         (ON COMMIT (DELETE | PRESERVE) ROWS)?
         oid_clause?
         oid_index_clause?
         physical_properties?
         table_properties
    ;

xmltype_virtual_columns
    : VIRTUAL COLUMNS LEFT_PAREN column_name AS LEFT_PAREN expression RIGHT_PAREN (COMMA column_name AS LEFT_PAREN expression RIGHT_PAREN)* RIGHT_PAREN
    ;

xmltype_column_properties
    : XMLTYPE COLUMN? column_name xmltype_storage? xmlschema_spec?
    ;

xmltype_storage
    : STORE AS ( OBJECT RELATIONAL
               | (SECUREFILE | BASICFILE)? (CLOB | BINARY XML) (lob_segname (LEFT_PAREN lob_parameters RIGHT_PAREN)? | LEFT_PAREN lob_parameters RIGHT_PAREN)?
               )
    | STORE VARRAYS AS (LOBS | TABLES)
    ;

xmlschema_spec
    : (XMLSCHEMA DELIMITED_ID)? ELEMENT DELIMITED_ID
         (allow_or_disallow NONSCHEMA)?
         (allow_or_disallow ANYSCHEMA)?
    ;

object_table
    : OF (schema_name PERIOD)? object_type
      object_table_substitution?
      (LEFT_PAREN object_properties (COMMA object_properties)* RIGHT_PAREN)?
      (ON COMMIT (DELETE | PRESERVE) ROWS)?
      oid_clause?
      oid_index_clause?
      physical_properties?
      table_properties
    ;

object_type
    : regular_id
    ;

oid_index_clause
    : OIDINDEX index_name? LEFT_PAREN (physical_attributes_clause | TABLESPACE tablespace)+ RIGHT_PAREN
    ;

oid_clause
    : OBJECT IDENTIFIER IS (SYSTEM GENERATED | PRIMARY KEY)
    ;

object_properties
    : (column_name | attribute_name) (DEFAULT expression)? (inline_constraint (COMMA inline_constraint)* | inline_ref_constraint)?
    | out_of_line_constraint
    | out_of_line_ref_constraint
    | supplemental_logging_props
    ;

object_table_substitution
    : NOT? SUBSTITUTABLE AT ALL LEVELS
    ;

relational_table
    : (LEFT_PAREN relational_property (COMMA relational_property)* RIGHT_PAREN)?
      immutable_table_clauses
      blockchain_table_clauses?
      (DEFAULT COLLATION collation_name)?
      (ON COMMIT (DROP | PRESERVE) DEFINITION)?
      (ON COMMIT (DELETE | PRESERVE) ROWS)?
      physical_properties?
      table_properties
    ;

immutable_table_clauses
    : immutable_table_no_drop_clause? immutable_table_no_delete_clause?
    ;

immutable_table_no_drop_clause
    : NO DROP (UNTIL numeric DAYS IDLE)?
    ;

immutable_table_no_delete_clause
    : NO DELETE (LOCKED? | UNTIL numeric DAYS AFTER INSERT LOCKED?)
    ;

blockchain_table_clauses
    : blockchain_drop_table_clause blockchain_row_retention_clause blockchain_hash_and_data_format_clause
    ;

blockchain_drop_table_clause
    : NO DROP (UNTIL numeric DAYS IDLE)?
    ;

blockchain_row_retention_clause
    : NO DELETE (LOCKED? | UNTIL numeric DAYS AFTER INSERT LOCKED?)
    ;

blockchain_hash_and_data_format_clause
    : HASHING USING SHA2_512_Q VERSION V1_Q
    ;

collation_name
    : identifier
    ;

table_properties
    : column_properties?
        read_only_clause?
        indexing_clause?
        table_partitioning_clauses?
        attribute_clustering_clause?
        (CACHE | NOCACHE)?
        result_cache_clause?
        parallel_clause?
        (ROWDEPENDENCIES | NOROWDEPENDENCIES)?
        enable_disable_clause*
        row_movement_clause?
        logical_replication_clause?
        flashback_archive_clause?
        physical_properties?
        (ROW ARCHIVAL)?
        (AS select_only_statement | FOR EXCHANGE WITH TABLE (schema_name PERIOD)? table_name)?
    ;

read_only_clause
    : READ (ONLY | WRITE)
    ;

indexing_clause
    : INDEXING (ON | OFF)
    ;

attribute_clustering_clause
    : CLUSTERING
        clustering_join?
        cluster_clause
        (yes_no? ON LOAD)?
        (yes_no? ON DATA MOVEMENT)?
        zonemap_clause?
    ;

clustering_join
    : (schema_name PERIOD)? table_name clustering_join_item (COMMA clustering_join_item)*
    ;

clustering_join_item
    : JOIN (schema_name PERIOD)? table_name ON LEFT_PAREN equijoin_condition RIGHT_PAREN
    ;

equijoin_condition
    : expression
    ;

cluster_clause
    : BY (LINEAR | INTERLEAVED)? ORDER clustering_columns
    ;

clustering_columns
    : clustering_column_group
    | LEFT_PAREN clustering_column_group (COMMA clustering_column_group)* RIGHT_PAREN
    ;

clustering_column_group
    : LEFT_PAREN column_name (COMMA column_name)* RIGHT_PAREN
    ;

yes_no
    : YES
    | NO
    ;

zonemap_clause
    : WITH MATERIALIZED ZONEMAP (LEFT_PAREN zonemap_name RIGHT_PAREN)?
    | WITHOUT MATERIALIZED ZONEMAP
    ;

logical_replication_clause
    : DISABLE LOGICAL REPLICATION
    | ENABLE LOGICAL REPLICATION ((ALL | ALLOW NOVALIDATE) KEYS)?
    ;

table_name
    : identifier
    ;

relational_property
    : column_definition
    | virtual_column_definition
    | period_definition
    | out_of_line_constraint
    | out_of_line_ref_constraint
    | supplemental_logging_props
    ;

table_partitioning_clauses
    : range_partitions
    | list_partitions
    | hash_partitions
    | composite_range_partitions
    | composite_list_partitions
    | composite_hash_partitions
    | reference_partitioning
    | system_partitioning
    ;

range_partitions
    : PARTITION BY RANGE LEFT_PAREN column_name (COMMA column_name)* RIGHT_PAREN
        (INTERVAL LEFT_PAREN expression RIGHT_PAREN (STORE IN LEFT_PAREN tablespace (COMMA tablespace)* RIGHT_PAREN )? )?
          LEFT_PAREN PARTITION partition_name? range_values_clause table_partition_description (COMMA PARTITION partition_name? range_values_clause table_partition_description)* RIGHT_PAREN
    ;

list_partitions
    : PARTITION BY LIST LEFT_PAREN column_name RIGHT_PAREN
        LEFT_PAREN PARTITION partition_name? list_values_clause table_partition_description  (COMMA PARTITION partition_name? list_values_clause table_partition_description )* RIGHT_PAREN
    ;

hash_partitions
    : PARTITION BY HASH LEFT_PAREN column_name (COMMA column_name)* RIGHT_PAREN
        (individual_hash_partitions | hash_partitions_by_quantity)
    ;

individual_hash_partitions
    : LEFT_PAREN PARTITION partition_name? partitioning_storage_clause? (COMMA PARTITION partition_name? partitioning_storage_clause?)* RIGHT_PAREN
    ;

hash_partitions_by_quantity
    : PARTITIONS hash_partition_quantity
       (STORE IN LEFT_PAREN tablespace (COMMA tablespace)* RIGHT_PAREN)?
         (table_compression | key_compression)?
         (OVERFLOW STORE IN LEFT_PAREN tablespace (COMMA tablespace)* RIGHT_PAREN )?
    ;

hash_partition_quantity
    : UNSIGNED_INTEGER
    ;

composite_range_partitions
    : PARTITION BY RANGE LEFT_PAREN column_name (COMMA column_name)* RIGHT_PAREN
       (INTERVAL LEFT_PAREN expression RIGHT_PAREN (STORE IN LEFT_PAREN tablespace (COMMA tablespace)* RIGHT_PAREN )? )?
       (subpartition_by_range | subpartition_by_list | subpartition_by_hash)
         LEFT_PAREN range_partition_desc (COMMA range_partition_desc)* RIGHT_PAREN
    ;

composite_list_partitions
    : PARTITION BY LIST LEFT_PAREN column_name RIGHT_PAREN
       (subpartition_by_range | subpartition_by_list | subpartition_by_hash)
        LEFT_PAREN list_partition_desc (COMMA list_partition_desc)* RIGHT_PAREN
    ;

composite_hash_partitions
    : PARTITION BY HASH LEFT_PAREN (COMMA column_name)+ RIGHT_PAREN
       (subpartition_by_range | subpartition_by_list | subpartition_by_hash)
         (individual_hash_partitions | hash_partitions_by_quantity)
    ;

reference_partitioning
    : PARTITION BY REFERENCE LEFT_PAREN regular_id RIGHT_PAREN
             (LEFT_PAREN reference_partition_desc (COMMA reference_partition_desc)* RIGHT_PAREN)?
    ;

reference_partition_desc
    : PARTITION partition_name? table_partition_description
    ;

system_partitioning
    : PARTITION BY SYSTEM
       (PARTITIONS UNSIGNED_INTEGER | reference_partition_desc (COMMA reference_partition_desc)*)?
    ;

range_partition_desc
    : PARTITION partition_name? range_values_clause table_partition_description
        ( ( LEFT_PAREN ( range_subpartition_desc (COMMA range_subpartition_desc)*
                | list_subpartition_desc (COMMA list_subpartition_desc)*
                | individual_hash_subparts (COMMA individual_hash_subparts)*
                )
            RIGHT_PAREN
          | hash_subparts_by_quantity
          )
        )?
    ;

list_partition_desc
    : PARTITION partition_name? list_values_clause table_partition_description
        ( ( LEFT_PAREN ( range_subpartition_desc (COMMA range_subpartition_desc)*
                | list_subpartition_desc (COMMA list_subpartition_desc)*
                | individual_hash_subparts (COMMA individual_hash_subparts)*
                )
            RIGHT_PAREN
          | hash_subparts_by_quantity
          )
        )?
    ;

subpartition_template
    : SUBPARTITION TEMPLATE
        ( ( LEFT_PAREN ( range_subpartition_desc (COMMA range_subpartition_desc)*
                | list_subpartition_desc (COMMA list_subpartition_desc)*
                | individual_hash_subparts (COMMA individual_hash_subparts)*
                )
            RIGHT_PAREN
          | hash_subpartition_quantity
          )
        )
    ;

hash_subpartition_quantity
    : UNSIGNED_INTEGER
    ;

subpartition_by_range
    : SUBPARTITION BY RANGE LEFT_PAREN column_name (COMMA column_name)* RIGHT_PAREN subpartition_template?
    ;

subpartition_by_list
    : SUBPARTITION BY LIST LEFT_PAREN column_name RIGHT_PAREN subpartition_template?
    ;

subpartition_by_hash
    : SUBPARTITION BY HASH LEFT_PAREN column_name (COMMA column_name)* RIGHT_PAREN
       (SUBPARTITIONS UNSIGNED_INTEGER (STORE IN LEFT_PAREN tablespace (COMMA tablespace)* RIGHT_PAREN )?
       | subpartition_template
       )?
    ;

subpartition_name
    : partition_name
    ;

range_subpartition_desc
    : SUBPARTITION subpartition_name? range_values_clause partitioning_storage_clause?
    ;

list_subpartition_desc
    : SUBPARTITION subpartition_name? list_values_clause partitioning_storage_clause?
    ;

individual_hash_subparts
    : SUBPARTITION subpartition_name? partitioning_storage_clause?
    ;

hash_subparts_by_quantity
    : SUBPARTITIONS UNSIGNED_INTEGER (STORE IN LEFT_PAREN tablespace (COMMA tablespace)* RIGHT_PAREN )?
    ;

range_values_clause
    : VALUES LESS THAN LEFT_PAREN literal (COMMA literal)* RIGHT_PAREN
    ;

list_values_clause
    : VALUES LEFT_PAREN (literal (COMMA literal)* | TIMESTAMP literal (COMMA TIMESTAMP literal)* | DEFAULT) RIGHT_PAREN
    ;

table_partition_description
    : deferred_segment_creation? segment_attributes_clause?
        (table_compression | key_compression)?
        (OVERFLOW segment_attributes_clause? )?
        (lob_storage_clause | varray_col_properties | nested_table_col_properties)*
    ;

partitioning_storage_clause
    : ( TABLESPACE tablespace
      | OVERFLOW (TABLESPACE tablespace)?
      | table_compression
      | key_compression
      | lob_partitioning_storage
      | VARRAY varray_item STORE AS (BASICFILE | SECUREFILE)? LOB lob_segname
      )+
    ;

lob_partitioning_storage
    : LOB LEFT_PAREN lob_item RIGHT_PAREN
       STORE AS (BASICFILE | SECUREFILE)?
               (lob_segname (LEFT_PAREN TABLESPACE tablespace RIGHT_PAREN )?
               | LEFT_PAREN TABLESPACE tablespace RIGHT_PAREN
               )
    ;

// https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/size_clause.html
// Technically, this should only allow 'K' | 'M' | 'G' | 'T' | 'P' | 'E'
// but having issues with examples/numbers01.sql line 11 "sysdate -1m"
size_clause
    : UNSIGNED_INTEGER (K_LETTER | M_LETTER | G_LETTER | T_LETTER | P_LETTER | E_LETTER)?
    ;


table_compression
    : COMPRESS
        ( BASIC
        | FOR ( OLTP
              | (QUERY | ARCHIVE) (LOW | HIGH)?
              )
        )?
    | NOCOMPRESS
    ;

// avoid to match an empty string in
inmemory_table_clause
    : inmemory_column_clause+
    | (INMEMORY inmemory_attributes? | NO INMEMORY) inmemory_column_clause*
    ;

// avoid to match an empty string in
inmemory_attributes
    : inmemory_memcompress inmemory_priority? inmemory_distribute? inmemory_duplicate?
    | inmemory_priority inmemory_distribute? inmemory_duplicate?
    | inmemory_distribute inmemory_duplicate?
    | inmemory_duplicate
    ;

inmemory_memcompress
    : MEMCOMPRESS FOR (DML | (QUERY | CAPACITY) (LOW | HIGH)?)
    | NO MEMCOMPRESS
    ;

inmemory_priority
    : PRIORITY (NONE | LOW | MEDIUM | HIGH | CRITICAL)
    ;

inmemory_distribute
    : DISTRIBUTE
        (AUTO | BY (ROWID RANGE | PARTITION | SUBPARTITION))?
        (FOR SERVICE (DEFAULT | ALL | identifier | NONE))?
    ;

inmemory_duplicate
    : DUPLICATE ALL?
    | NO DUPLICATE
    ;

inmemory_column_clause
    : (INMEMORY inmemory_memcompress? | NO INMEMORY) LEFT_PAREN column_list RIGHT_PAREN
    ;

physical_attributes_clause
    : (PCTFREE pctfree=UNSIGNED_INTEGER
      | PCTUSED pctused=UNSIGNED_INTEGER
      | INITRANS inittrans=UNSIGNED_INTEGER
      | MAXTRANS maxtrans=UNSIGNED_INTEGER
      | COMPUTE STATISTICS
      | storage_clause
      | compute_clauses
      )+
    ;

storage_clause
    : STORAGE LEFT_PAREN
         (INITIAL initial_size=size_clause
         | NEXT next_size=size_clause
         | MINEXTENTS minextents=(UNSIGNED_INTEGER | UNLIMITED)
         | MAXEXTENTS minextents=(UNSIGNED_INTEGER | UNLIMITED)
         | PCTINCREASE pctincrease=UNSIGNED_INTEGER
         | FREELISTS freelists=UNSIGNED_INTEGER
         | FREELIST GROUPS freelist_groups=UNSIGNED_INTEGER
         | OPTIMAL (size_clause | NULL_ )
         | BUFFER_POOL (KEEP | RECYCLE | DEFAULT)
         | FLASH_CACHE (KEEP | NONE | DEFAULT)
         | CELL_FLASH_CACHE (KEEP | NONE | DEFAULT)
         | ENCRYPT
         )+
       RIGHT_PAREN
    ;

deferred_segment_creation
    : SEGMENT CREATION (IMMEDIATE | DEFERRED)
    ;

segment_attributes_clause
    : ( physical_attributes_clause
      | TABLESPACE (tablespace_name=id_expression | SET? identifier)
      | table_compression
      | logging_clause
      )+
    ;

physical_properties
    : deferred_segment_creation? segment_attributes_clause table_compression? inmemory_table_clause? ilm_clause?
    | deferred_segment_creation? ( ORGANIZATION ( HEAP segment_attributes_clause? heap_org_table_clause
                                                | INDEX segment_attributes_clause? index_org_table_clause
                                                | EXTERNAL external_table_clause
                                                )
                                 | EXTERNAL PARTITION ATTRIBUTES external_table_clause (REJECT LIMIT)?
                                 )
    | CLUSTER cluster_name LEFT_PAREN column_name (COMMA column_name)* RIGHT_PAREN
    ;

ilm_clause
    : ILM ( ADD POLICY ilm_policy_clause
          | (DELETE | ENABLE | DISABLE) POLICY ilm_policy_clause
          | DELETE_ALL
          | ENABLE_ALL
          | DISABLE_ALL
          )
    ;

ilm_policy_clause
    : ilm_compression_policy
    | ilm_tiering_policy
    | ilm_inmemory_policy
    ;

ilm_compression_policy
    : table_compression segment_group ilm_after_on
    | ((ROW | COLUMN) STORE COMPRESS (ADVANCED | FOR QUERY)) ROW AFTER ilm_time_period OF NO MODIFICATION
    ;

ilm_tiering_policy
    : TIER TO tablespace ( segment_group? (ON function_name)?
                         | READ ONLY segment_group? ilm_after_on
                         )
    ;

ilm_after_on
    : AFTER ilm_time_period OF (NO (ACCESS | MODIFICATION) | CREATION)
    | ON function_name
    ;

segment_group
    : SEGMENT
    | GROUP
    ;

ilm_inmemory_policy
    : ( SET INMEMORY inmemory_attributes?
      | MODIFY INMEMORY inmemory_memcompress
      | NO INMEMORY
      ) SEGMENT? ilm_after_on
    ;

ilm_time_period
    : numeric ( DAY
              | DAYS
              | MONTH
              | MONTHS
              | YEAR
              | YEARS
              )
    ;

heap_org_table_clause
    : table_compression? inmemory_table_clause? ilm_clause?
    ;

external_table_clause
    : LEFT_PAREN (TYPE access_driver_type)? external_table_data_props RIGHT_PAREN (REJECT LIMIT (numeric | UNLIMITED))? inmemory_table_clause?
    ;

access_driver_type
    : ORACLE_LOADER
    | ORACLE_DATAPUMP
    | ORACLE_HDFS
    | ORACLE_HIVE
    ;

external_table_data_props
    : (DEFAULT DIRECTORY directory_name)?
        (ACCESS PARAMETERS ( LEFT_PAREN CHAR_STRING RIGHT_PAREN
                           | LEFT_PAREN opaque_format_spec RIGHT_PAREN
                           | USING CLOB select_only_statement
                           ))?
        (LOCATION LEFT_PAREN directory_name COLON CHAR_STRING (COMMA directory_name COLON CHAR_STRING)* RIGHT_PAREN )?
    ;

// https://docs.oracle.com/en/database/oracle/oracle-database/21/sutil/oracle-external-tables.html
opaque_format_spec
    : record_format_info? field_definitions? column_transforms?
    ;
    
record_format_info
    : RECORDS   (
                    (FIXED | VARIABLE) UNSIGNED_INTEGER
                    | DELIMITED BY (DETECTED? NEWLINE_ | CHAR_STRING)
                    | XMLTAG CHAR_STRING
                ) et_record_spec_options?
    ;

// The string for external table data properties.
// Reference: https://docs.oracle.com/en/database/oracle/oracle-database/21/sutil/oracle_loader-access-driver.html#GUID-03FD7309-EAE9-4880-9BD1-945B32311478
et_string
    : CHAR_STRING
    | DELIMITED_ID
    | HEX_STRING_LIT
    ;

et_record_spec_options
    : et_record_spec_option+
    ;
    
et_record_spec_option
    : CHARACTERSET char_set_name
    | PREPROCESSOR (directory_spec COLON)? file_spec
    | (LANGUAGE | TERRITORY) CHAR_STRING
    | DATA IS (LITTLE | BIG) ENDIAN
    | BYTEORDERMARK (CHECK | NOCHECK)
    | STRING SIZES ARE IN (BYTES | CHARACTERS)
    | LOAD WHEN condition
    | et_output_files
    | READSIZE UNSIGNED_INTEGER
    | DISABLE_DIRECTORY_LINK_CHECK
    | (DATE_CACHE | SKIP_) UNSIGNED_INTEGER
    | FIELD_NAMES (
                    FIRST FILE IGNORE?
                    | ALL FILES IGNORE?
                    | NONE
                  )
    | IO_OPTIONS LEFT_PAREN (DIRECTIO | NODIRECTIO) RIGHT_PAREN
    | (DNFS_ENABLE | DNFS_DISABLE)
    | DNFS_READBUFFERS UNSIGNED_INTEGER
    ;
    
et_output_files
    : et_output_file+
    ;

et_output_file
    : NOBADFILE
    | BADFILE (directory_spec COLON)? file_spec?
    | NODISCARDFILE
    | DISCARDFILE (directory_spec COLON)? file_spec?
    | NOLOGFILE
    | LOGFILE (directory_spec COLON)? file_spec?
    ;
    
directory_spec
    : directory_name
    ;
    
file_spec
    : CHAR_STRING
    ;
    
field_definitions
    : FIELDS field_options? field_list?
    ;
    
field_options
    : field_option+
    ;

field_option
    : IGNORE_CHARS_AFTER_EOR
      | CSV ((WITH | WITHOUT) EMBEDDED)?
      | delim_spec
      | trim_spec
      | ALL FIELDS OVERRIDE THESE FIELDS
      | MISSING FIELD VALUES ARE NULL_
      | REJECT ROWS WITH ALL NULL_ FIELDS
      | field_date_format+
      | (NULLIF | NONULLIF)
    ;
    
field_list
    : LEFT_PAREN field_item (COMMA field_item)* RIGHT_PAREN
    ;
    
field_item
    : field_name pos_spec? datatype_spec? init_spec? lls_clause?
    ;
    
field_name
    : column_name
    ;

pos_spec
    : POSITION? LEFT_PAREN (pos_start | '*' (PLUS_SIGN | '-') pos_increment) ((':' | '-') (pos_end | pos_length))? RIGHT_PAREN
    ;

pos_start
    : UNSIGNED_INTEGER
    ;

pos_increment
    : UNSIGNED_INTEGER
    ;

pos_end
    : UNSIGNED_INTEGER
    ;

pos_length
    : UNSIGNED_INTEGER
    ;
    
datatype_spec
    : UNSIGNED? INTEGER EXTERNAL? (LEFT_PAREN UNSIGNED_INTEGER RIGHT_PAREN)? delim_spec?
    | (DECIMAL | ZONED) (
                            EXTERNAL (LEFT_PAREN UNSIGNED_INTEGER RIGHT_PAREN)? delim_spec?
                            | precision_part
                        )?
    | ORACLE_DATE
    | ORACLE_NUMBER COUNTED?
    | FLOAT EXTERNAL? (LEFT_PAREN UNSIGNED_INTEGER RIGHT_PAREN)? delim_spec?
    | DOUBLE
    | BINARY_FLOAT EXTERNAL? (LEFT_PAREN UNSIGNED_INTEGER RIGHT_PAREN)? delim_spec?
    | BINARY_DOUBLE
    | RAW (LEFT_PAREN UNSIGNED_INTEGER RIGHT_PAREN)?
    | CHAR (LEFT_PAREN UNSIGNED_INTEGER RIGHT_PAREN)? delim_spec? trim_spec? field_date_format?
    | (
        VARCHAR
        | VARRAW
        | VARCHARC
        | VARRAWC
      ) LEFT_PAREN (numeric COMMA)? numeric RIGHT_PAREN
    ;
    
init_spec
    : (DEFAULTIF | NULLIF) condition
    ;
    
lls_clause
    : LLS (directory_spec COLON)? file_spec?
    ;
    
delim_spec
    : ENCLOSED BY et_string (AND et_string)?
    | TERMINATED BY (et_string | WHITESPACE) (OPTIONALLY? ENCLOSED BY et_string (AND et_string)? )?
    ;
    
trim_spec
    : LRTRIM
    | NOTRIM
    | LTRIM
    | RTRIM
    | LDRTRIM
    ;

field_date_format
    : DATE_FORMAT (DATE | TIMESTAMP)? MASK datetime_expr
    ;
    
column_transforms
    : COLUMN TRANSFORMS LEFT_PAREN transform (COMMA transform)* RIGHT_PAREN
    ;
    
transform
    : column_name FROM  (
                            NULL_
                            | CONSTANT CHAR_STRING
                            | CONCAT LEFT_PAREN concat_item (COMMA concat_item)* RIGHT_PAREN
                            | LOBFILE LEFT_PAREN lobfile_item (COMMA lobfile_item)* RIGHT_PAREN lobfile_attr_list?
                            | STARTOF source_field LEFT_PAREN UNSIGNED_INTEGER RIGHT_PAREN
                        )
    ;
    
source_field
    : column_name
    ;
    
lobfile_item
    : column_name
    | CONSTANT CHAR_STRING COLON
    ;
    
lobfile_attr_list
    : FROM LEFT_PAREN (directory_spec COLON)? file_spec? (COMMA (directory_spec COLON)? file_spec?)* RIGHT_PAREN
    | CLOB
    | BLOB
    | CHARACTERSET EQUALS_OP char_set_name
    ;
    
concat_item
    : column_name
    | CONSTANT CHAR_STRING
    ;

row_movement_clause
    : (ENABLE | DISABLE)? ROW MOVEMENT
    ;

flashback_archive_clause
    : FLASHBACK ARCHIVE fa=id_expression?
    | NO FLASHBACK ARCHIVE
    ;

log_grp
    : UNSIGNED_INTEGER
    ;

supplemental_table_logging
    : ADD SUPPLEMENTAL LOG  (supplemental_log_grp_clause | supplemental_id_key_clause)
       (COMMA SUPPLEMENTAL LOG  (supplemental_log_grp_clause | supplemental_id_key_clause) )*
    | DROP SUPPLEMENTAL LOG (supplemental_id_key_clause | GROUP log_grp)
        (COMMA SUPPLEMENTAL LOG (supplemental_id_key_clause | GROUP log_grp) )*
    ;

supplemental_log_grp_clause
    : GROUP log_grp LEFT_PAREN regular_id (NO LOG)? (COMMA regular_id (NO LOG)?)* RIGHT_PAREN ALWAYS?
    ;

supplemental_id_key_clause
    : DATA LEFT_PAREN( COMMA? ( ALL
                     | PRIMARY KEY
                     | UNIQUE
                     | FOREIGN KEY
                     )
              )+
           RIGHT_PAREN
      COLUMNS
    ;

allocate_extent_clause
    : ALLOCATE EXTENT
       ( LEFT_PAREN ( SIZE size_clause
             | DATAFILE datafile=CHAR_STRING
             | INSTANCE inst_num=UNSIGNED_INTEGER
             )+
         RIGHT_PAREN
       )?
    ;

deallocate_unused_clause
    : DEALLOCATE UNUSED (KEEP size_clause)?
    ;

shrink_clause
    : SHRINK SPACE_KEYWORD COMPACT? CASCADE?
    ;

records_per_block_clause
    : (MINIMIZE | NOMINIMIZE)? RECORDS_PER_BLOCK
    ;

upgrade_table_clause
    : UPGRADE (NOT? INCLUDING DATA) column_properties
    ;

truncate_table
    : TRUNCATE TABLE tableview_name PURGE? SEMICOLON
    ;

// https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/DROP-TABLE.html
drop_table
    : DROP TABLE tableview_name (CASCADE CONSTRAINTS)? PURGE? SEMICOLON
    ;

// https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/DROP-TABLESPACE.html
drop_tablespace
    : DROP TABLESPACE ts=id_expression ((DROP | KEEP) QUOTA?)?
        including_contents_clause?
    ;

// https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/DROP-TABLESPACE-SET.html
drop_tablespace_set
    : DROP TABLESPACE SET tss=id_expression
        including_contents_clause?
    ;

including_contents_clause
    : INCLUDING CONTENTS ((AND | KEEP) DATAFILES)? (CASCADE CONSTRAINTS)?
    ;

drop_view
    : DROP VIEW tableview_name (CASCADE CONSTRAINT)? SEMICOLON
    ;

comment_on_column
    : COMMENT ON COLUMN column_name IS quoted_string
    ;

enable_or_disable
    : ENABLE
    | DISABLE
    ;
allow_or_disallow
    : ALLOW
    | DISALLOW
    ;

// Synonym DDL Clauses

alter_synonym
    : ALTER PUBLIC? SYNONYM (schema_name PERIOD)? synonym_name (EDITIONABLE | NONEDITIONABLE | COMPILE)
    ;

create_synonym
    // Synonym's schema cannot be specified for public synonyms
    : CREATE (OR REPLACE)? PUBLIC SYNONYM synonym_name FOR (schema_name PERIOD)? schema_object_name (AT_SIGN link_name)?
    | CREATE (OR REPLACE)? SYNONYM (schema_name PERIOD)? synonym_name FOR (schema_name PERIOD)? schema_object_name (AT_SIGN link_name)?
    ;

drop_synonym
    : DROP PUBLIC? SYNONYM (schema_name PERIOD)? synonym_name FORCE?
    ;

// https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/CREATE-SPFILE.html
create_spfile
    : CREATE SPFILE (EQUALS_OP spfile_name)?
        FROM (PFILE (EQUALS_OP pfile_name)? (AS COPY)? | MEMORY)
    ;

spfile_name
    : CHAR_STRING
    ;

pfile_name
    : CHAR_STRING
    ;

comment_on_table
    : COMMENT ON TABLE tableview_name IS quoted_string
    ;

comment_on_materialized
    : COMMENT ON MATERIALIZED VIEW tableview_name IS quoted_string
    ;

alter_analytic_view
    : ALTER ANALYTIC VIEW (schema_name PERIOD)? av=id_expression
        ( RENAME TO id_expression
        | COMPILE
        | alter_add_cache_clause
        | alter_drop_cache_clause
        )
    ;

alter_add_cache_clause
    : ADD CACHE MEASURE GROUP LEFT_PAREN (ALL | measure_list)? RIGHT_PAREN
        LEVELS LEFT_PAREN levels_item (COMMA levels_item)* RIGHT_PAREN
    ;

levels_item
    : ((d=id_expression PERIOD)? h=id_expression PERIOD)? l=id_expression
    ;

measure_list
    : id_expression (COMMA id_expression)*
    ;

alter_drop_cache_clause
    : DROP CACHE MEASURE GROUP LEFT_PAREN (ALL | measure_list)? RIGHT_PAREN
        LEVELS LEFT_PAREN levels_item (COMMA levels_item)* RIGHT_PAREN
    ;

// https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/ALTER-ATTRIBUTE-DIMENSION.html
alter_attribute_dimension
    : ALTER ATTRIBUTE DIMENSION (schema_name PERIOD)? ad=id_expression
        (RENAME TO nad=id_expression | COMPILE)
    ;

// https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/ALTER-AUDIT-POLICY-Unified-Auditing.html
alter_audit_policy
    : ALTER AUDIT POLICY p=id_expression
        ADD? (privilege_audit_clause? action_audit_clause? role_audit_clause? | (ONLY TOPLEVEL)?)
        DROP? (privilege_audit_clause? action_audit_clause? role_audit_clause? | (ONLY TOPLEVEL)?)
        (CONDITION ( DROP
                   | CHAR_STRING EVALUATE PER (STATEMENT | SESSION | INSTANCE))
                   )?
    ;

alter_cluster
    : ALTER CLUSTER  cluster_name
        ( physical_attributes_clause
        | SIZE size_clause
        | allocate_extent_clause
        | deallocate_unused_clause
        | cache_or_nocache
        )+
        parallel_clause?
        SEMICOLON
    ;

drop_analytic_view
    : DROP ANALYTIC VIEW (schema_name PERIOD)? av=id_expression
    ;

// https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/DROP-ATTRIBUTE-DIMENSION.html
drop_attribute_dimension
    : DROP ATTRIBUTE DIMENSION (schema_name PERIOD)? ad=id_expression
    ;

// https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/DROP-AUDIT-POLICY-Unified-Auditing.html
drop_audit_policy
    : DROP AUDIT POLICY p=id_expression
    ;

// https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/DROP-FLASHBACK-ARCHIVE.html
drop_flashback_archive
    : DROP FLASHBACK ARCHIVE fa=id_expression
    ;

drop_cluster
    : DROP CLUSTER cluster_name (INCLUDING TABLES (CASCADE CONSTRAINTS)?)?
    ;

// https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/DROP-CONTEXT.html
drop_context
    : DROP CONTEXT ns=id_expression
    ;

// https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/DROP-DIRECTORY.html
drop_directory
    : DROP DIRECTORY dn=id_expression
    ;

// https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/DROP-DISKGROUP.html
drop_diskgroup
    : DROP DISKGROUP dgn=id_expression ((FORCE? INCLUDING | EXCLUDING) CONTENTS)?
    ;

// https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/DROP-EDITION.html
drop_edition
    : DROP EDITION e=id_expression CASCADE?
    ;

truncate_cluster
    : TRUNCATE CLUSTER cluster_name ((DROP | REUSE) STORAGE)?
    ;

cache_or_nocache
    : CACHE
    | NOCACHE
    ;

database_name
    : regular_id
    ;

alter_database
    : ALTER database_clause
       ( startup_clauses
       | recovery_clauses
       | database_file_clauses
       | logfile_clauses
       | controlfile_clauses
       | standby_database_clauses
       | default_settings_clause
       | instance_clauses
       | security_clause
       | prepare_clause
       | drop_mirror_clause
       | lost_write_protection
       | cdb_fleet_clauses
       | property_clauses
       | replay_upgrade_clauses
       )
      SEMICOLON
    ;

database_clause
    : PLUGGABLE? DATABASE database_name?
    ;

startup_clauses
    : MOUNT ((STANDBY | CLONE) DATABASE)?
    | OPEN (READ WRITE)? resetlogs_or_noresetlogs? upgrade_or_downgrade?
    | OPEN READ ONLY
    ;

resetlogs_or_noresetlogs
    : RESETLOGS
    | NORESETLOGS
    ;

upgrade_or_downgrade
    : UPGRADE
    | DOWNGRADE
    ;

recovery_clauses
    : general_recovery
    | managed_standby_recovery
    | begin_or_end BACKUP
    ;

begin_or_end
    : BEGIN
    | END
    ;

general_recovery
    : RECOVER AUTOMATIC? (FROM CHAR_STRING)?
       ( (full_database_recovery | partial_database_recovery | LOGFILE CHAR_STRING )?
         ((TEST | ALLOW UNSIGNED_INTEGER CORRUPTION | parallel_clause)+ )?
       | CONTINUE DEFAULT?
       | CANCEL
       )
    ;

//Need to come back to
full_database_recovery
    : STANDBY? DATABASE
          ((UNTIL (CANCEL |TIME CHAR_STRING | CHANGE UNSIGNED_INTEGER | CONSISTENT)
           | USING BACKUP CONTROLFILE
           | SNAPSHOT TIME CHAR_STRING
           )+
          )?
    ;

partial_database_recovery
    : TABLESPACE tablespace (COMMA tablespace)*
    | DATAFILE CHAR_STRING | filenumber (COMMA CHAR_STRING | filenumber)*
    | partial_database_recovery_10g
    ;

partial_database_recovery_10g
    : {p.isVersion10()}? STANDBY
      ( TABLESPACE tablespace (COMMA tablespace)*
      | DATAFILE CHAR_STRING | filenumber (COMMA CHAR_STRING | filenumber)*
      )
      UNTIL (CONSISTENT WITH)? CONTROLFILE
    ;


managed_standby_recovery
    : RECOVER (MANAGED STANDBY DATABASE
               ((USING CURRENT LOGFILE
                | DISCONNECT (FROM SESSION)?
                | NODELAY
                | UNTIL CHANGE UNSIGNED_INTEGER
                | UNTIL CONSISTENT
                | parallel_clause
                )+
               | FINISH
               | CANCEL
               )?
              | TO LOGICAL STANDBY (db_name | KEEP IDENTITY)
              )
    ;

db_name
    : regular_id
    ;
database_file_clauses
    : RENAME FILE filename (COMMA filename)* TO filename
    | create_datafile_clause
    | alter_datafile_clause
    | alter_tempfile_clause
    | move_datafile_clause
    ;

create_datafile_clause
    : CREATE DATAFILE (filename | filenumber) (COMMA (filename | filenumber) )*
        (AS (//TODO (COMMA? file_specification)+ |
              NEW) )?
    ;

alter_datafile_clause
    : DATAFILE (filename | filenumber) (COMMA (filename | filenumber) )*
        ( ONLINE
        | OFFLINE (FOR DROP)?
        | RESIZE size_clause
        | autoextend_clause
        | END BACKUP
        )
    ;

alter_tempfile_clause
    : TEMPFILE (filename | filenumber) (COMMA (filename | filenumber) )*
        ( RESIZE size_clause
        | autoextend_clause
        | DROP (INCLUDING DATAFILES)
        | ONLINE
        | OFFLINE
        )
    ;

move_datafile_clause
    : MOVE DATAFILE (filename | filenumber) (COMMA (filename | filenumber) )*
        (TO filename)? REUSE? KEEP?
    ;

logfile_clauses
    : (ARCHIVELOG MANUAL? | NOARCHIVELOG)
    | NO? FORCE LOGGING
    | SET STANDBY NOLOGGING FOR (DATA AVAILABILITY | LOAD PERFORMANCE)
    | RENAME FILE filename (COMMA filename)* TO filename
    | CLEAR UNARCHIVED? LOGFILE logfile_descriptor (COMMA logfile_descriptor)* (UNRECOVERABLE DATAFILE)?
    | add_logfile_clauses
    | drop_logfile_clauses
    | switch_logfile_clause
    | supplemental_db_logging
    ;

add_logfile_clauses
    : ADD STANDBY? LOGFILE
        ( (INSTANCE CHAR_STRING | THREAD UNSIGNED_INTEGER)? group_redo_logfile+
        | MEMBER filename REUSE? (COMMA filename REUSE?)* TO logfile_descriptor (COMMA logfile_descriptor)*
        )
    ;

group_redo_logfile
    : (GROUP UNSIGNED_INTEGER)? redo_log_file_spec
    ;

drop_logfile_clauses
    : DROP STANDBY?
          LOGFILE (logfile_descriptor (COMMA logfile_descriptor)*
                  | MEMBER filename (COMMA filename)*
                  )
    ;

switch_logfile_clause
    : SWITCH ALL LOGFILES TO BLOCKSIZE UNSIGNED_INTEGER
    ;

supplemental_db_logging
    :  add_or_drop
          SUPPLEMENTAL LOG (DATA
                           | supplemental_id_key_clause
                           | supplemental_plsql_clause
                           )
    ;

add_or_drop
    : ADD
    | DROP
    ;

supplemental_plsql_clause
    : DATA FOR PROCEDURAL REPLICATION
    ;

logfile_descriptor
    : GROUP UNSIGNED_INTEGER
    | LEFT_PAREN filename (COMMA filename)* RIGHT_PAREN
    | filename
    ;

controlfile_clauses
    : CREATE (LOGICAL | PHYSICAL)? STANDBY CONTROLFILE AS filename REUSE?
    | BACKUP CONTROLFILE TO (filename REUSE? | trace_file_clause)
    ;

trace_file_clause
    : TRACE (AS filename REUSE?)? (RESETLOGS|NORESETLOGS)?
    ;

standby_database_clauses
    : ( activate_standby_db_clause
      | maximize_standby_db_clause
      | register_logfile_clause
      | commit_switchover_clause
      | start_standby_clause
      | stop_standby_clause
      | convert_database_clause
      )
      parallel_clause?
    ;

activate_standby_db_clause
    : ACTIVATE (PHYSICAL | LOGICAL)? STANDBY DATABASE (FINISH APPLY)?
    ;

maximize_standby_db_clause
    : SET STANDBY DATABASE TO MAXIMIZE (PROTECTION | AVAILABILITY | PERFORMANCE)
    ;

register_logfile_clause
    : REGISTER (OR REPLACE)? (PHYSICAL | LOGICAL) LOGFILE //TODO (COMMA? file_specification)+
    //TODO   (FOR logminer_session_name)?
    ;

commit_switchover_clause
    : (PREPARE | COMMIT) TO SWITCHOVER
        ((TO (((PHYSICAL | LOGICAL)? PRIMARY |  PHYSICAL? STANDBY)
           ((WITH | WITHOUT)? SESSION SHUTDOWN (WAIT | NOWAIT) )?
          | LOGICAL STANDBY
          )
         | LOGICAL STANDBY
         )
        | CANCEL
        )?
    ;

start_standby_clause
    : START LOGICAL STANDBY APPLY IMMEDIATE? NODELAY?
        ( NEW PRIMARY regular_id
        | INITIAL scn_value=UNSIGNED_INTEGER?
        | SKIP_ FAILED TRANSACTION
        | FINISH
        )?
    ;

stop_standby_clause
    : (STOP | ABORT) LOGICAL STANDBY APPLY
    ;

convert_database_clause
    : CONVERT TO (PHYSICAL | SNAPSHOT) STANDBY
    ;

default_settings_clause
    : DEFAULT EDITION EQUALS_OP edition_name
    | SET DEFAULT (BIGFILE | SMALLFILE) TABLESPACE
    | DEFAULT TABLESPACE tablespace
    | DEFAULT TEMPORARY TABLESPACE (tablespace | tablespace_group_name)
    | RENAME GLOBAL_NAME TO database (PERIOD domain)+
    | ENABLE BLOCK CHANGE TRACKING (USING FILE filename REUSE?)?
    | DISABLE BLOCK CHANGE TRACKING
    | flashback_mode_clause
    | set_time_zone_clause
    ;

set_time_zone_clause
    : SET TIMEZONE EQUALS_OP CHAR_STRING
    ;

instance_clauses
    : enable_or_disable INSTANCE CHAR_STRING
    ;

security_clause
    : GUARD (ALL | STANDBY | NONE)
    ;

domain
    : regular_id
    ;

database
    : regular_id
    ;

edition_name
    : regular_id
    ;

filenumber
    : UNSIGNED_INTEGER
    ;

filename
    : CHAR_STRING
    ;

prepare_clause
    : PREPARE MIRROR COPY c=id_expression (WITH (UNPROTECTED | MIRROR | HIGH) REDUNDANCY)?
        (FOR DATABASE id_expression)?
    ;

drop_mirror_clause
    : DROP MIRROR COPY mn=id_expression
    ;

lost_write_protection
    : (ENABLE | DISABLE | REMOVE | SUSPEND) LOST WRITE PROTECTION
    ;

cdb_fleet_clauses
    : lead_cdb_clause
    | lead_cdb_uri_clause
    ;

lead_cdb_clause
    : SET LEAD_CDB EQUALS_OP (TRUE | FALSE)
    ;

lead_cdb_uri_clause
    : SET LEAD_CDB_URI EQUALS_OP CHAR_STRING
    ;

property_clauses
    : PROPERTY (SET | REMOVE) DEFAULT_CREDENTIAL EQUALS_OP qcn=id_expression
    ;

replay_upgrade_clauses
    : UPGRADE SYNC (ON | OFF)
    ;

alter_database_link
    : ALTER SHARED? PUBLIC? DATABASE LINK link_name (CONNECT TO user_object_name IDENTIFIED BY password_value link_authentication? | link_authentication)
    ;

password_value
    : id_expression
    | numeric
    ;

link_authentication
    : AUTHENTICATED BY user_object_name IDENTIFIED BY password_value
    ;

// added by zrh
create_database
    : CREATE DATABASE database_name
    ( USER (SYS | SYSTEM) IDENTIFIED BY password_value
     | CONTROLFILE REUSE
     | (MAXDATAFILES | MAXINSTANCES) UNSIGNED_INTEGER
     | NATIONAL? CHARACTER SET char_set_name
     | SET DEFAULT (BIGFILE | SMALLFILE) TABLESPACE
     | database_logging_clauses
     | tablespace_clauses
     | set_time_zone_clause
     | (BIGFILE | SMALLFILE)? USER_DATA TABLESPACE tablespace_group_name DATAFILE datafile_tempfile_spec (COMMA datafile_tempfile_spec)*
     | enable_pluggable_database
    )+
    ;

database_logging_clauses
    : LOGFILE database_logging_sub_clause (COMMA database_logging_sub_clause)*
    | (MAXLOGFILES | MAXLOGMEMBERS | MAXLOGHISTORY) UNSIGNED_INTEGER
    | ARCHIVELOG
    | NOARCHIVELOG
    | FORCE LOGGING
    ;

database_logging_sub_clause
    : (GROUP UNSIGNED_INTEGER)? file_specification
    ;

tablespace_clauses
    : EXTENT MANAGEMENT LOCAL
    | SYSAUX? DATAFILE file_specification (COMMA file_specification)*
    | default_tablespace
    | default_temp_tablespace
    | undo_tablespace
    ;

enable_pluggable_database
    : ENABLE PLUGGABLE DATABASE (SEED file_name_convert? (SYSTEM tablespace_datafile_clauses)? (SYSAUX tablespace_datafile_clauses)? )? undo_mode_clause?
    ;

file_name_convert
    : FILE_NAME_CONVERT EQUALS_OP ( LEFT_PAREN filename_convert_sub_clause (COMMA filename_convert_sub_clause)* RIGHT_PAREN | NONE)
    ;

filename_convert_sub_clause
    : CHAR_STRING (COMMA CHAR_STRING)?
    ;

tablespace_datafile_clauses
    : DATAFILES (SIZE size_clause | autoextend_clause)+
    ;

undo_mode_clause
    : LOCAL UNDO (ON | OFF)
    ;

default_tablespace
    : DEFAULT TABLESPACE tablespace (DATAFILE datafile_tempfile_spec)? extent_management_clause?
    ;

default_temp_tablespace
    : (BIGFILE | SMALLFILE)? DEFAULT (TEMPORARY TABLESPACE | LOCAL TEMPORARY TABLESPACE FOR (ALL | LEAF)) tablespace
        (TEMPFILE file_specification (COMMA file_specification)*)? extent_management_clause?
    ;

undo_tablespace
    : (BIGFILE | SMALLFILE)? UNDO TABLESPACE tablespace (DATAFILE file_specification (COMMA file_specification)*)?
    ;

// https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/DROP-DATABASE.html
drop_database
    : DROP DATABASE (INCLUDING BACKUPS)? NOPROMPT?
    ;

// https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/CREATE-DATABASE-LINK.html
create_database_link
    : CREATE SHARED? PUBLIC? DATABASE LINK dblink ( CONNECT TO ( CURRENT_USER
                                                               | user_object_name IDENTIFIED BY password_value link_authentication?
                                                               )
                                                  | link_authentication
                                                  )* (USING CHAR_STRING)?
    ;

dblink
    : database_name (PERIOD d=id_expression)* (AT_SIGN cq=id_expression)?
    ;

drop_database_link
    : DROP PUBLIC? DATABASE LINK dblink
    ;

// https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/ALTER-TABLESPACE-SET.html
alter_tablespace_set
    : ALTER TABLESPACE SET tss=id_expression alter_tablespace_attrs
    ;

alter_tablespace_attrs
    : default_tablespace_params
    | MINIMUM EXTENT size_clause
    | RESIZE size_clause
    | COALESCE
    | SHRINK SPACE_KEYWORD (KEEP size_clause)?
    | RENAME TO nts=id_expression
    | (BEGIN | END) BACKUP
    | datafile_tempfile_clauses
    | tablespace_logging_clauses
    | tablespace_group_clause
    | tablespace_state_clauses
    | autoextend_clause
    | flashback_mode_clause
    | tablespace_retention_clause
    | alter_tablespace_encryption
    | lost_write_protection
    ;

alter_tablespace_encryption
    : ENCRYPTION ( OFFLINE (tablespace_encryption_spec? ENCRYPT | DECRYPT)
                 | ONLINE (tablespace_encryption_spec? (ENCRYPT | REKEY) | DECRYPT) ts_file_name_convert?
                 | FINISH (ENCRYPT | REKEY | DECRYPT) ts_file_name_convert?
                 )
    ;

ts_file_name_convert
    : FILE_NAME_CONVERT EQUALS_OP LEFT_PAREN CHAR_STRING COMMA CHAR_STRING (COMMA CHAR_STRING COMMA CHAR_STRING)* RIGHT_PAREN KEEP?
    ;

alter_role
    : ALTER ROLE role_name role_identified_clause container_clause?
    ;

role_identified_clause
    : NOT IDENTIFIED
    | IDENTIFIED ( BY identifier
                 | USING identifier (PERIOD id_expression)?
                 | EXTERNALLY
                 | GLOBALLY (AS CHAR_STRING)?
                 )
    ;

alter_table
    : ALTER TABLE tableview_name memoptimize_read_write_clause*
      (
      | alter_table_properties
      | constraint_clauses
      | column_clauses
      | alter_table_partitioning
//TODO      | alter_external_table
      | move_table_clause
      )
      ((enable_disable_clause | enable_or_disable (TABLE LOCK | ALL TRIGGERS) )+)?
      SEMICOLON
    ;

memoptimize_read_write_clause
    : NO? MEMOPTIMIZE FOR (READ | WRITE)
    ;

alter_table_properties
    : alter_table_properties_1
    | RENAME TO tableview_name
    | shrink_clause
    | READ ONLY
    | READ WRITE
    | REKEY CHAR_STRING
    ;

alter_table_partitioning
    : add_table_partition
    | drop_table_partition
    | merge_table_partition
    | modify_table_partition
    | split_table_partition
    | truncate_table_partition
    | exchange_table_partition
    | coalesce_table_partition
    | alter_interval_partition
    ;


add_table_partition
    : ADD ( range_partition_desc
          | list_partition_desc
          | PARTITION partition_name? (TABLESPACE tablespace)? key_compression? UNUSABLE?
      )
    ;

drop_table_partition
    : DROP (partition_extended_names | subpartition_extended_names) (update_index_clauses parallel_clause?)?
    ;

merge_table_partition
    : MERGE PARTITION partition_name AND partition_name INTO PARTITION partition_name
    ;

modify_table_partition
    : MODIFY 
      (
        | table_partitioning_clauses
        | (PARTITION partition_name ((ADD | DROP) list_values_clause)? (ADD range_subpartition_desc)? (REBUILD? UNUSABLE LOCAL INDEXES)?)
      )
      (ONLINE)?
      (update_index_clauses parallel_clause?)?
    ;

split_table_partition
    : SPLIT PARTITION partition_name INTO LEFT_PAREN
         (range_partition_desc (COMMA range_partition_desc)*
         | list_partition_desc (COMMA list_partition_desc)* )
     RIGHT_PAREN
    ;

truncate_table_partition
    : TRUNCATE (partition_extended_names | subpartition_extended_names)
            ((DROP ALL? | REUSE)? STORAGE)? CASCADE? (update_index_clauses parallel_clause?)?
    ;

exchange_table_partition
    : EXCHANGE PARTITION partition_name WITH TABLE tableview_name
            ((INCLUDING | EXCLUDING) INDEXES)?
            ((WITH | WITHOUT) VALIDATION)?
    ;

coalesce_table_partition
    : COALESCE PARTITION parallel_clause? (allow_or_disallow CLUSTERING)?
    ;

alter_interval_partition
    : SET INTERVAL LEFT_PAREN (constant | expression)? RIGHT_PAREN
    ;


partition_extended_names
    : (PARTITION | PARTITIONS) ( partition_name
                               | LEFT_PAREN partition_name (COMMA partition_name)* RIGHT_PAREN
                               | FOR LEFT_PAREN? partition_key_value (COMMA partition_key_value)* RIGHT_PAREN? )
    ;

subpartition_extended_names
    : (SUBPARTITION | SUBPARTITIONS) ( partition_name (UPDATE INDEXES)?
                                     | LEFT_PAREN partition_name (COMMA partition_name)* RIGHT_PAREN
                                     | FOR LEFT_PAREN? subpartition_key_value (COMMA subpartition_key_value)* RIGHT_PAREN? )
    ;

alter_table_properties_1
    : ( physical_attributes_clause
      | logging_clause
      | table_compression
      | inmemory_table_clause
      | supplemental_table_logging
      | allocate_extent_clause
      | deallocate_unused_clause
      | (CACHE | NOCACHE)
      | RESULT_CACHE LEFT_PAREN MODE (DEFAULT | FORCE) RIGHT_PAREN
      | upgrade_table_clause
      | records_per_block_clause
      | parallel_clause
      | row_movement_clause
      | flashback_archive_clause
      )+
      alter_iot_clauses?
    ;

alter_iot_clauses
    : index_org_table_clause
    | alter_overflow_clause
    | alter_mapping_table_clause
    | COALESCE
    ;

alter_mapping_table_clause
    : MAPPING TABLE (allocate_extent_clause | deallocate_unused_clause)
    ;

alter_overflow_clause
    : add_overflow_clause
    | OVERFLOW (segment_attributes_clause | allocate_extent_clause | shrink_clause | deallocate_unused_clause)+
    ;

add_overflow_clause
    : ADD OVERFLOW segment_attributes_clause? (LEFT_PAREN PARTITION segment_attributes_clause? (COMMA PARTITION segment_attributes_clause?)*  RIGHT_PAREN )?
    ;


update_index_clauses
    : update_global_index_clause
    | update_all_indexes_clause
    ;

update_global_index_clause
    : (UPDATE | INVALIDATE) GLOBAL INDEXES
    ;

update_all_indexes_clause
    : UPDATE INDEXES (LEFT_PAREN update_all_indexes_index_clause RIGHT_PAREN)?
    ;

update_all_indexes_index_clause
    : index_name LEFT_PAREN (update_index_partition | update_index_subpartition) RIGHT_PAREN (COMMA update_all_indexes_clause)*
    ;

update_index_partition
    : index_partition_description index_subpartition_clause? (COMMA update_index_partition)*
    ;

update_index_subpartition
    : SUBPARTITION subpartition_name? (TABLESPACE tablespace)? (COMMA update_index_subpartition)*
    ;

enable_disable_clause
    : (ENABLE | DISABLE) (VALIDATE | NOVALIDATE)?
         (UNIQUE LEFT_PAREN column_name (COMMA column_name)* RIGHT_PAREN
         | PRIMARY KEY
         | CONSTRAINT constraint_name
         ) using_index_clause? exceptions_clause?
         CASCADE? ((KEEP | DROP) INDEX)?
    ;

using_index_clause
    : USING INDEX (index_name | LEFT_PAREN create_index RIGHT_PAREN | index_attributes )?
    ;

index_attributes
    : ( physical_attributes_clause
      | logging_clause
      | ONLINE
      | TABLESPACE (tablespace | DEFAULT)
      | key_compression
      | sort_or_nosort
      | REVERSE
      | visible_or_invisible
      | partial_index_clause
      | parallel_clause
      )+
    ;

sort_or_nosort
    : SORT
    | NOSORT
    ;

exceptions_clause
    : EXCEPTIONS INTO tableview_name
    ;

move_table_clause
    : MOVE ONLINE?
        segment_attributes_clause?
        table_compression?
        index_org_table_clause?
        (lob_storage_clause | varray_col_properties)*
        parallel_clause?
    ;

index_org_table_clause
    : (mapping_table_clause | PCTTHRESHOLD UNSIGNED_INTEGER | key_compression)+ index_org_overflow_clause?
    | index_org_overflow_clause // rule move_table_clause contains an optional block with at least one alternative that can match an empty string
    ;

mapping_table_clause
    : MAPPING TABLE
    | NOMAPPING
    ;

key_compression
    : NOCOMPRESS
    | COMPRESS UNSIGNED_INTEGER
    ;

index_org_overflow_clause
    : (INCLUDING column_name)? OVERFLOW segment_attributes_clause?
    ;

column_clauses
    : add_modify_drop_column_clauses
    | rename_column_clause
    | modify_collection_retrieval
    | modify_lob_storage_clause
    ;

modify_collection_retrieval
    : MODIFY NESTED TABLE collection_item RETURN AS (LOCATOR | VALUE)
    ;

collection_item
    : tableview_name
    ;

rename_column_clause
    : RENAME COLUMN old_column_name TO new_column_name
    ;

old_column_name
    : column_name
    ;

new_column_name
    : column_name
    ;

add_modify_drop_column_clauses
    : (constraint_clauses
      |add_column_clause
      |modify_column_clauses
      |drop_column_clause
      )+
    ;

drop_column_clause
    : SET UNUSED (COLUMN column_name| (LEFT_PAREN column_name (COMMA column_name)* RIGHT_PAREN )) (CASCADE CONSTRAINTS | INVALIDATE)*
    | DROP (COLUMN column_name | LEFT_PAREN column_name (COMMA column_name)* RIGHT_PAREN ) (CASCADE CONSTRAINTS | INVALIDATE)* (CHECKPOINT UNSIGNED_INTEGER)?
    | DROP (UNUSED COLUMNS | COLUMNS CONTINUE) (CHECKPOINT UNSIGNED_INTEGER)
    ;

modify_column_clauses
    : MODIFY (LEFT_PAREN modify_col_properties (COMMA modify_col_properties)* RIGHT_PAREN
             | LEFT_PAREN modify_col_visibility (COMMA modify_col_visibility) RIGHT_PAREN
             | modify_col_properties
             | modify_col_visibility
             | modify_col_substitutable
             )
    ;
    
modify_col_visibility
    : column_name (VISIBLE | INVISIBLE)
    ;

modify_col_properties
    : column_name datatype? (COLLATE column_collation_name)? (DEFAULT (ON NULL_)? expression)? (ENCRYPT encryption_spec | DECRYPT)? inline_constraint* lob_storage_clause? //TODO alter_xmlschema_clause
    ;

modify_col_substitutable
    : COLUMN column_name NOT? SUBSTITUTABLE AT ALL LEVELS FORCE?
    ;

add_column_clause
    : ADD (LEFT_PAREN (column_definition | virtual_column_definition) (COMMA (column_definition
              | virtual_column_definition)
              )*
          RIGHT_PAREN
          | ( column_definition | virtual_column_definition ))
       column_properties?
//TODO       (COMMA? out_of_line_part_storage )
    ;

varray_col_properties
    : VARRAY varray_item ( substitutable_column_clause? varray_storage_clause
                         | substitutable_column_clause
                         )
    ;

varray_storage_clause
    : STORE AS (SECUREFILE|BASICFILE)? LOB ( lob_segname? LEFT_PAREN lob_storage_parameters RIGHT_PAREN
                                           | lob_segname
                                           )
    ;

lob_segname
    : regular_id
    ;

lob_item
    : regular_id
    | quoted_string
    ;

lob_storage_parameters
    :  TABLESPACE tablespace_name=id_expression | (lob_parameters storage_clause? )
    |  storage_clause
    ;

lob_storage_clause
    : LOB ( LEFT_PAREN lob_item (COMMA lob_item)* RIGHT_PAREN STORE AS ( (SECUREFILE|BASICFILE) | LEFT_PAREN lob_storage_parameters* RIGHT_PAREN )+
          | LEFT_PAREN lob_item RIGHT_PAREN STORE AS ( (SECUREFILE | BASICFILE) | lob_segname | LEFT_PAREN lob_storage_parameters* RIGHT_PAREN )+
          )
    ;

modify_lob_storage_clause
    : MODIFY LOB LEFT_PAREN lob_item RIGHT_PAREN LEFT_PAREN modify_lob_parameters RIGHT_PAREN
    ;

modify_lob_parameters
    : ( storage_clause
      | (PCTVERSION | FREEPOOLS) UNSIGNED_INTEGER
      | REBUILD FREEPOOLS
      | lob_retention_clause
      | lob_deduplicate_clause
      | lob_compression_clause
      | ENCRYPT encryption_spec
      | DECRYPT
      | CACHE
      | (CACHE | NOCACHE | CACHE READS) logging_clause?
      | allocate_extent_clause
      | shrink_clause
      | deallocate_unused_clause
     )+
    ;

lob_parameters
    : ( (ENABLE | DISABLE) STORAGE IN ROW
      | CHUNK UNSIGNED_INTEGER
      | PCTVERSION UNSIGNED_INTEGER
      | FREEPOOLS UNSIGNED_INTEGER
      | lob_retention_clause
      | lob_deduplicate_clause
      | lob_compression_clause
      | ENCRYPT encryption_spec
      | DECRYPT
      | (CACHE | NOCACHE | CACHE READS) logging_clause?
      )+
    ;

lob_deduplicate_clause
    : DEDUPLICATE
    | KEEP_DUPLICATES
    ;

lob_compression_clause
    : NOCOMPRESS
    | COMPRESS (HIGH | MEDIUM | LOW)?
    ;

lob_retention_clause
    : RETENTION (MAX | MIN UNSIGNED_INTEGER | AUTO | NONE)?
    ;

encryption_spec
    : (USING CHAR_STRING)?
        (IDENTIFIED BY REGULAR_ID)?
        CHAR_STRING? (NO? SALT)?
    ;
tablespace
    : id_expression
    ;

varray_item
    : (id_expression PERIOD)? (id_expression PERIOD)? id_expression
    ;

column_properties
    : (object_type_col_properties
    | nested_table_col_properties
    | (varray_col_properties | lob_storage_clause) (LEFT_PAREN lob_partition_storage (COMMA lob_partition_storage)* RIGHT_PAREN)? //TODO LEFT_PAREN ( COMMA? lob_partition_storage)+ RIGHT_PAREN
    | xmltype_column_properties)+
    ;

lob_partition_storage
    : LOB (LEFT_PAREN lob_item (COMMA lob_item) RIGHT_PAREN STORE AS ((SECUREFILE | BASICFILE) | LEFT_PAREN lob_storage_parameters RIGHT_PAREN)+
          | LEFT_PAREN lob_item RIGHT_PAREN STORE AS ((SECUREFILE | BASICFILE) | lob_segname | LEFT_PAREN lob_storage_parameters RIGHT_PAREN)+
          )
    ;

period_definition
    : {p.isVersion12()}? PERIOD FOR column_name
        ( LEFT_PAREN start_time_column COMMA end_time_column RIGHT_PAREN )?
    ;

start_time_column
    : column_name
    ;

end_time_column
    : column_name
    ;

column_definition
    : column_name
         ( (datatype | regular_id) (COLLATE column_collation_name)?)?
         SORT?
         (VISIBLE | INVISIBLE)?
         (DEFAULT (ON NULL_)? expression | identity_clause)?
         (ENCRYPT encryption_spec)?
         (inline_constraint+ | inline_ref_constraint)?
    ;

column_collation_name
    : id_expression
    ;

identity_clause
    : GENERATED (ALWAYS | BY DEFAULT (ON NULL_)?)? AS IDENTITY identity_options_parentheses?
    ;

//https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/CREATE-TABLE.html#GUID-F9CE0CC3-13AE-4744-A43C-EAC7A71AAAB6
// NOTE to identity options
// according to the SQL Reference, identity_options be nested in parentheses.
// But statements without parentheses can also be executed successfully on a oracle database.
// See this issue for more details: https://github.com/antlr/grammars-v4/issues/3183
identity_options_parentheses
    : identity_options+
    | LEFT_PAREN identity_options+ RIGHT_PAREN
    ;

identity_options
    : START WITH (numeric | LIMIT VALUE)
    | INCREMENT BY numeric
    | MAXVALUE numeric
    | NOMAXVALUE
    | MINVALUE numeric
    | NOMINVALUE
    | CYCLE
    | NOCYCLE
    | CACHE numeric
    | NOCACHE
    | ORDER
    | NOORDER
    ;

virtual_column_definition
    : column_name (datatype COLLATE column_collation_name)?
        (VISIBLE | INVISIBLE)?
        autogenerated_sequence_definition?
        VIRTUAL?
        evaluation_edition_clause?
        (UNUSABLE BEFORE (CURRENT EDITION | EDITION edition_name))?
                (UNUSABLE BEGINNING WITH ((CURRENT | NULL_) EDITION | EDITION edition_name))?
        inline_constraint*
    ;

autogenerated_sequence_definition
    : GENERATED (ALWAYS | BY DEFAULT (ON NULL_)?)? AS IDENTITY ( LEFT_PAREN (sequence_start_clause | sequence_spec)* RIGHT_PAREN )?
    ;

evaluation_edition_clause
    : EVALUATE USING ((CURRENT | NULL_) EDITION | EDITION edition_name)
    ;

nested_table_col_properties
    : NESTED TABLE  (nested_item | COLUMN_VALUE) substitutable_column_clause? (LOCAL | GLOBAL)?
       STORE AS tableview_name ( LEFT_PAREN ( LEFT_PAREN object_properties RIGHT_PAREN
                                     | physical_properties
                                     | column_properties
                                     )+
                                  RIGHT_PAREN
                               )?
        (RETURN AS? (LOCATOR | VALUE) )?
     ;

nested_item
    : regular_id
    ;

substitutable_column_clause
    : ELEMENT? IS OF TYPE? LEFT_PAREN type_name RIGHT_PAREN
    | NOT? SUBSTITUTABLE AT ALL LEVELS
    ;

partition_name
    : regular_id
    | DELIMITED_ID
    ;

supplemental_logging_props
    : SUPPLEMENTAL LOG (supplemental_log_grp_clause | supplemental_id_key_clause)
    ;

object_type_col_properties
    : COLUMN column=regular_id substitutable_column_clause
    ;

constraint_clauses
    : ADD LEFT_PAREN (out_of_line_constraint* | out_of_line_ref_constraint) RIGHT_PAREN
    | ADD  (out_of_line_constraint* | out_of_line_ref_constraint)
    | MODIFY (CONSTRAINT constraint_name | PRIMARY KEY | UNIQUE LEFT_PAREN column_name (COMMA column_name)* RIGHT_PAREN)  constraint_state CASCADE?
    | RENAME CONSTRAINT old_constraint_name TO new_constraint_name
    | drop_constraint_clause+
    ;

old_constraint_name
    : constraint_name
    ;

new_constraint_name
    : constraint_name
    ;

drop_constraint_clause
    : DROP  drop_primary_key_or_unique_or_generic_clause
    ;

drop_primary_key_or_unique_or_generic_clause
    : (PRIMARY KEY | UNIQUE LEFT_PAREN column_name (COMMA column_name)* RIGHT_PAREN) CASCADE? (KEEP | DROP)?
    | CONSTRAINT constraint_name CASCADE?
    ;

check_constraint
    : CHECK LEFT_PAREN condition RIGHT_PAREN DISABLE?
    ;

foreign_key_clause
    : FOREIGN KEY paren_column_list references_clause on_delete_clause?
    ;

references_clause
    : REFERENCES tableview_name paren_column_list? (ON DELETE (CASCADE | SET NULL_))?
    ;

on_delete_clause
    : ON DELETE (CASCADE | SET NULL_)
    ;

// Anonymous PL/SQL code block

anonymous_block
    : (DECLARE seq_of_declare_specs)? BEGIN seq_of_statements (EXCEPTION exception_handler+)? END SEMICOLON
    ;

// Common DDL Clauses

invoker_rights_clause
    : AUTHID (CURRENT_USER | DEFINER)
    ;

call_spec
    : LANGUAGE (java_spec | c_spec)
    ;

// Call Spec Specific Clauses

java_spec
    : JAVA NAME CHAR_STRING
    ;

c_spec
    : C_LETTER (NAME CHAR_STRING)? LIBRARY identifier c_agent_in_clause? (WITH CONTEXT)? c_parameters_clause?
    ;

c_agent_in_clause
    : AGENT IN LEFT_PAREN expressions RIGHT_PAREN
    ;

c_parameters_clause
    : PARAMETERS LEFT_PAREN (expressions | PERIOD PERIOD PERIOD) RIGHT_PAREN
    ;

parameter
    : parameter_name (IN | OUT | INOUT | NOCOPY)* type_spec? default_value_part?
    ;

default_value_part
    : (ASSIGN_OP | DEFAULT) expression
    ;

// Elements Declarations

seq_of_declare_specs
    : declare_spec+
    ;

declare_spec
    : pragma_declaration
    | exception_declaration
    | variable_declaration
    | subtype_declaration
    | cursor_declaration
    | type_declaration
    | procedure_spec
    | function_spec
    | procedure_body
    | function_body
    ;

// incorporates constant_declaration
variable_declaration
    : identifier CONSTANT? type_spec (NOT NULL_)? default_value_part? SEMICOLON
    ;

subtype_declaration
    : SUBTYPE identifier IS type_spec (RANGE expression '..' expression)? (NOT NULL_)? SEMICOLON
    ;

// cursor_declaration incorportates curscursor_body and cursor_spec

cursor_declaration
    : CURSOR identifier (LEFT_PAREN parameter_spec (COMMA parameter_spec)* RIGHT_PAREN )? (RETURN type_spec)? (IS select_statement)? SEMICOLON
    ;

parameter_spec
    : parameter_name (IN? type_spec)? default_value_part?
    ;

exception_declaration
    : identifier EXCEPTION SEMICOLON
    ;

pragma_declaration
    : PRAGMA (SERIALLY_REUSABLE
    | AUTONOMOUS_TRANSACTION
    | EXCEPTION_INIT LEFT_PAREN exception_name COMMA numeric_negative RIGHT_PAREN
    | INLINE LEFT_PAREN id1=identifier COMMA expression RIGHT_PAREN
    | RESTRICT_REFERENCES LEFT_PAREN (identifier | DEFAULT) (COMMA identifier)+ RIGHT_PAREN) SEMICOLON
    ;

// Record Declaration Specific Clauses

// incorporates ref_cursor_type_definition

record_type_def
    : RECORD LEFT_PAREN field_spec (COMMA field_spec)* RIGHT_PAREN
    ;

field_spec
    : column_name type_spec? (NOT NULL_)? default_value_part?
    ;

ref_cursor_type_def
    : REF CURSOR (RETURN type_spec)?
    ;

type_declaration
    : TYPE identifier IS (table_type_def | varray_type_def | record_type_def | ref_cursor_type_def) SEMICOLON
    ;

table_type_def
    : TABLE OF type_spec table_indexed_by_part? (NOT NULL_)?
    ;

table_indexed_by_part
    : (idx1=INDEXED | idx2=INDEX) BY type_spec
    ;

varray_type_def
    : (VARRAY | VARYING ARRAY) LEFT_PAREN expression RIGHT_PAREN OF type_spec (NOT NULL_)?
    ;

// Statements

seq_of_statements
    : (statement (SEMICOLON | EOF) | label_declaration)+
    ;

label_declaration
    : ltp1= '<' '<' label_name '>' '>'
    ;

statement
    : body
    | block
    | assignment_statement
    | continue_statement
    | exit_statement
    | goto_statement
    | if_statement
    | loop_statement
    | forall_statement
    | null_statement
    | raise_statement
    | return_statement
    | case_statement
    | sql_statement
    | pipe_row_statement
    | plsql_call_statement // put call statement after others. BYT-8268: PL/SQL level allows optional CALL
    ;

assignment_statement
    : (general_element | bind_variable) ASSIGN_OP expression
    ;

continue_statement
    : CONTINUE label_name? (WHEN condition)?
    ;

exit_statement
    : EXIT label_name? (WHEN condition)?
    ;

goto_statement
    : GOTO label_name
    ;

if_statement
    : IF condition THEN seq_of_statements elsif_part* else_part? END IF
    ;

elsif_part
    : ELSIF condition THEN seq_of_statements
    ;

else_part
    : ELSE seq_of_statements
    ;

loop_statement
    : label_declaration? (WHILE condition | FOR cursor_loop_param)? LOOP seq_of_statements END LOOP label_name?
    ;

// Loop Specific Clause

cursor_loop_param
    : index_name IN REVERSE? lower_bound range_separator='..' upper_bound
    | record_name IN (cursor_name (LEFT_PAREN expressions? RIGHT_PAREN)? | LEFT_PAREN select_statement RIGHT_PAREN)
    ;

forall_statement
    : FORALL index_name IN bounds_clause sql_statement (SAVE EXCEPTIONS)?
    ;

bounds_clause
    : lower_bound '..' upper_bound
    | INDICES OF collection_name between_bound?
    | VALUES OF index_name
    ;

between_bound
    : BETWEEN lower_bound AND upper_bound
    ;

lower_bound
    : concatenation
    ;

upper_bound
    : concatenation
    ;

null_statement
    : NULL_
    ;

raise_statement
    : RAISE exception_name?
    ;

return_statement
    : RETURN expression?
    ;

// BYT-8268: Split into two rules for different contexts
// SQL level: CALL keyword is MANDATORY to prevent misidentifying keywords like CASCADE as procedure calls
sql_call_statement
    : CALL routine_name function_argument? (PERIOD id_expression function_argument?)* (INTO bind_variable)?
    ;

// PL/SQL block level: CALL keyword is OPTIONAL (standard PL/SQL allows direct procedure calls)
// Support method chaining: obj_constructor(...).method(...) or CALL obj_constructor(...).method(...)
plsql_call_statement
    : CALL? routine_name function_argument? (PERIOD id_expression function_argument?)* (INTO bind_variable)?
    ;

pipe_row_statement
    : PIPE ROW LEFT_PAREN expression RIGHT_PAREN;

body
    : BEGIN seq_of_statements (EXCEPTION exception_handler+)? END label_name?
    ;

// Body Specific Clause

exception_handler
    : WHEN exception_name (OR exception_name)* THEN seq_of_statements
    ;

trigger_block
    : (DECLARE declare_spec*)? body
    ;

block
    : DECLARE? declare_spec+ body
    ;

// SQL Statements

sql_statement
    : execute_immediate
    | data_manipulation_language_statements
    | cursor_manipulation_statements
    | transaction_control_statements
    ;

execute_immediate
    : EXECUTE IMMEDIATE expression (into_clause using_clause? | using_clause dynamic_returning_clause? | dynamic_returning_clause)?
    ;

// Execute Immediate Specific Clause

dynamic_returning_clause
    : (RETURNING | RETURN) into_clause
    ;

// DML Statements

data_manipulation_language_statements
    : merge_statement
    | lock_table_statement
    | select_statement
    | update_statement
    | delete_statement
    | insert_statement
    | explain_statement
    ;

// Cursor Manipulation Statements

cursor_manipulation_statements
    : close_statement
    | open_statement
    | fetch_statement
    | open_for_statement
    ;

close_statement
    : CLOSE cursor_name
    ;

open_statement
    : OPEN cursor_name (LEFT_PAREN expressions? RIGHT_PAREN)?
    ;

fetch_statement
    : FETCH cursor_name (it1=INTO variable_name (COMMA variable_name)* | BULK COLLECT INTO variable_name (COMMA variable_name)* (LIMIT (numeric | variable_name))?)
    ;

open_for_statement
    : OPEN variable_name FOR (select_statement | expression) using_clause?
    ;

// Transaction Control SQL Statements

transaction_control_statements
    : set_transaction_command
    | set_constraint_command
    | commit_statement
    | rollback_statement
    | savepoint_statement
    ;

set_transaction_command
    : SET TRANSACTION
      (READ (ONLY | WRITE) | ISOLATION LEVEL (SERIALIZABLE | READ COMMITTED) | USE ROLLBACK SEGMENT rollback_segment_name)?
      (NAME quoted_string)?
    ;

set_constraint_command
    : SET (CONSTRAINT | CONSTRAINTS) (ALL | constraint_name (COMMA constraint_name)*) (IMMEDIATE | DEFERRED)
    ;

// https://docs.oracle.com/cd/E18283_01/server.112/e17118/statements_4010.htm#SQLRF01110
// https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/COMMIT.html
// BYT-8268: Fixed to allow WRITE clause independently, not just after COMMENT
// Correct Oracle syntax: COMMIT [WORK] [COMMENT 'text'] [WRITE [WAIT|NOWAIT] [IMMEDIATE|BATCH]] [FORCE...]
commit_statement
    : COMMIT WORK? (COMMENT CHAR_STRING)? write_clause?
      ( FORCE ( CHAR_STRING (COMMA numeric)?
              | CORRUPT_XID CHAR_STRING
              | CORRUPT_XID_ALL
              )
      )?
    ;

write_clause
    : WRITE (WAIT | NOWAIT)? (IMMEDIATE | BATCH)?
    ;

rollback_statement
    : ROLLBACK WORK? (TO SAVEPOINT? savepoint_name | FORCE quoted_string)?
    ;

savepoint_statement
    : SAVEPOINT savepoint_name
    ;

// Dml

/* TODO
//SHOULD BE OVERRIDEN!
compilation_unit
    : seq_of_statements* EOF
    ;

//SHOULD BE OVERRIDEN!
seq_of_statements
    : select_statement
    | update_statement
    | delete_statement
    | insert_statement
    | lock_table_statement
    | merge_statement
    | explain_statement
//    | case_statement[true]
    ;
*/

explain_statement
    : EXPLAIN PLAN (SET STATEMENT_ID EQUALS_OP quoted_string)? (INTO tableview_name)?
      FOR (select_statement | update_statement | delete_statement | insert_statement | merge_statement)
    ;

select_only_statement
    : subquery
    ;

select_statement
    : select_only_statement (for_update_clause | order_by_clause | offset_clause | fetch_clause)*
    ;

// Select Specific Clauses

subquery_factoring_clause
    : WITH factoring_element (COMMA factoring_element)*
    ;

factoring_element
    : query_name paren_column_list? AS LEFT_PAREN subquery order_by_clause? RIGHT_PAREN
      search_clause? cycle_clause?
    ;

search_clause
    : SEARCH (DEPTH | BREADTH) FIRST BY column_name ASC? DESC? (NULLS FIRST)? (NULLS LAST)?
      (COMMA column_name ASC? DESC? (NULLS FIRST)? (NULLS LAST)?)* SET column_name
    ;

cycle_clause
    : CYCLE column_list SET column_name TO expression DEFAULT expression
    ;

subquery
    : subquery_basic_elements subquery_operation_part*
    ;

subquery_basic_elements
    : query_block
    | LEFT_PAREN subquery RIGHT_PAREN
    ;

subquery_operation_part
    : (UNION ALL? | INTERSECT | MINUS) subquery_basic_elements
    ;

query_block
    : subquery_factoring_clause? SELECT (DISTINCT | UNIQUE | ALL)? selected_list
      into_clause? from_clause where_clause? hierarchical_query_clause? group_by_clause? model_clause? order_by_clause? fetch_clause?
    ;

selected_list
    : '*'
    | select_list_elements (COMMA select_list_elements)*
    ;

from_clause
    : FROM table_ref_list
    ;

select_list_elements
    : table_wild
    | expression column_alias?
    ;
    
table_wild
    : tableview_name PERIOD ASTERISK
    ;

table_ref_list
    : table_ref (COMMA table_ref)*
    ;

// NOTE to PIVOT clause
// according the SQL reference this should not be possible
// according to he reality it is. Here we probably apply pivot/unpivot onto whole join clause
// eventhough it is not enclosed in parenthesis. See pivot examples 09,10,11

table_ref
    : table_ref_aux join_clause* (pivot_clause | unpivot_clause)?
    ;

table_ref_aux
    : table_ref_aux_internal flashback_query_clause* ({p.IsTableAlias()}? table_alias)?
    ;

table_ref_aux_internal
    : dml_table_expression_clause (pivot_clause | unpivot_clause)?                 # table_ref_aux_internal_one
    | LEFT_PAREN table_ref subquery_operation_part* RIGHT_PAREN (pivot_clause | unpivot_clause)?  # table_ref_aux_internal_two
    | ONLY LEFT_PAREN dml_table_expression_clause RIGHT_PAREN                                     # table_ref_aux_internal_three
    ;

join_clause
    : query_partition_clause? (CROSS | NATURAL)? (INNER | outer_join_type)?
      JOIN table_ref_aux query_partition_clause? (join_on_part | join_using_part)*
    ;

join_on_part
    : ON condition
    ;

join_using_part
    : USING paren_column_list
    ;

outer_join_type
    : (FULL | LEFT | RIGHT) OUTER?
    ;

query_partition_clause
    : PARTITION BY ((LEFT_PAREN (subquery | expressions)? RIGHT_PAREN) | expressions)
    ;

flashback_query_clause
    : VERSIONS BETWEEN (SCN | TIMESTAMP) expression
    | AS OF (SCN | TIMESTAMP | SNAPSHOT) expression
    ;

pivot_clause
    : PIVOT XML? LEFT_PAREN pivot_element (COMMA pivot_element)* pivot_for_clause pivot_in_clause RIGHT_PAREN
    ;

pivot_element
    : aggregate_function_name LEFT_PAREN expression RIGHT_PAREN column_alias?
    ;

pivot_for_clause
    : FOR (column_name | paren_column_list)
    ;

pivot_in_clause
    : IN LEFT_PAREN (subquery | ANY (COMMA ANY)* | pivot_in_clause_element (COMMA pivot_in_clause_element)*) RIGHT_PAREN
    ;

pivot_in_clause_element
    : pivot_in_clause_elements column_alias?
    ;

pivot_in_clause_elements
    : expression
    | LEFT_PAREN expressions? RIGHT_PAREN
    ;

unpivot_clause
    : UNPIVOT ((INCLUDE | EXCLUDE) NULLS)?
    LEFT_PAREN (column_name | paren_column_list) pivot_for_clause unpivot_in_clause RIGHT_PAREN
    ;

unpivot_in_clause
    : IN LEFT_PAREN unpivot_in_elements (COMMA unpivot_in_elements)* RIGHT_PAREN
    ;

unpivot_in_elements
    : (column_name | paren_column_list)
      (AS (constant | LEFT_PAREN constant (COMMA constant)* RIGHT_PAREN))?
    ;

hierarchical_query_clause
    : CONNECT BY NOCYCLE? condition start_part?
    | start_part CONNECT BY NOCYCLE? condition
    ;

start_part
    : START WITH condition
    ;

group_by_clause
    : GROUP BY group_by_elements (COMMA group_by_elements)* having_clause?
    | having_clause (GROUP BY group_by_elements (COMMA group_by_elements)*)?
    ;

group_by_elements
    : grouping_sets_clause
    | rollup_cube_clause
    | expression
    ;

rollup_cube_clause
    : (ROLLUP | CUBE) LEFT_PAREN grouping_sets_elements (COMMA grouping_sets_elements)* RIGHT_PAREN
    ;

grouping_sets_clause
    : GROUPING SETS LEFT_PAREN grouping_sets_elements (COMMA grouping_sets_elements)* RIGHT_PAREN
    ;

grouping_sets_elements
    : rollup_cube_clause
    | LEFT_PAREN expressions? RIGHT_PAREN
    | expression
    ;

having_clause
    : HAVING condition
    ;

model_clause
    : MODEL cell_reference_options* return_rows_clause? reference_model* main_model
    ;

cell_reference_options
    : (IGNORE | KEEP) NAV
    | UNIQUE (DIMENSION | SINGLE REFERENCE)
    ;

return_rows_clause
    : RETURN (UPDATED | ALL) ROWS
    ;

reference_model
    : REFERENCE reference_model_name ON LEFT_PAREN subquery RIGHT_PAREN model_column_clauses cell_reference_options*
    ;

main_model
    : (MAIN main_model_name)? model_column_clauses cell_reference_options* model_rules_clause
    ;

model_column_clauses
    : model_column_partition_part? DIMENSION BY model_column_list MEASURES model_column_list
    ;

model_column_partition_part
    : PARTITION BY model_column_list
    ;

model_column_list
    : LEFT_PAREN model_column (COMMA model_column)*  RIGHT_PAREN
    ;

model_column
    : (expression | query_block) column_alias?
    ;

model_rules_clause
    : model_rules_part? LEFT_PAREN (model_rules_element (COMMA model_rules_element)*)? RIGHT_PAREN
    ;

model_rules_part
    : RULES (UPDATE | UPSERT ALL?)? ((AUTOMATIC | SEQUENTIAL) ORDER)? model_iterate_clause?
    ;

model_rules_element
    : (UPDATE | UPSERT ALL?)? cell_assignment order_by_clause? EQUALS_OP expression
    ;

cell_assignment
    : model_expression
    ;

model_iterate_clause
    : ITERATE LEFT_PAREN expression RIGHT_PAREN until_part?
    ;

until_part
    : UNTIL LEFT_PAREN condition RIGHT_PAREN
    ;

order_by_clause
    : ORDER SIBLINGS? BY order_by_elements (COMMA order_by_elements)*
    ;

order_by_elements
    : expression (ASC | DESC)? (NULLS (FIRST | LAST))?
    ;

offset_clause
    : OFFSET expression (ROW | ROWS)
    ;

fetch_clause
    : FETCH (FIRST | NEXT) (expression PERCENT_KEYWORD?)? (ROW | ROWS) (ONLY | WITH TIES)
    ;

for_update_clause
    : FOR UPDATE for_update_of_part? for_update_options?
    ;

for_update_of_part
    : OF column_list
    ;

for_update_options
    : SKIP_ LOCKED
    | NOWAIT
    | WAIT expression
    ;

update_statement
    : UPDATE general_table_ref update_set_clause where_clause? static_returning_clause? error_logging_clause?
    ;

// Update Specific Clauses

update_set_clause
    : SET
      (column_based_update_set_clause (COMMA column_based_update_set_clause)* | VALUE LEFT_PAREN identifier RIGHT_PAREN EQUALS_OP expression)
    ;

column_based_update_set_clause
    : column_name EQUALS_OP expression
    | paren_column_list EQUALS_OP subquery
    | LEFT_PAREN column_name RIGHT_PAREN EQUALS_OP expression
    ;

delete_statement
    : DELETE FROM? general_table_ref where_clause? static_returning_clause? error_logging_clause?
    ;

insert_statement
    : INSERT (single_table_insert | multi_table_insert)
    ;

// Insert Specific Clauses

single_table_insert
    : insert_into_clause (values_clause static_returning_clause? | select_statement) error_logging_clause?
    ;

multi_table_insert
    : (ALL multi_table_element+ | conditional_insert_clause) select_statement
    ;

multi_table_element
    : insert_into_clause values_clause? error_logging_clause?
    ;

conditional_insert_clause
    : (ALL | FIRST)? conditional_insert_when_part+ conditional_insert_else_part?
    ;

conditional_insert_when_part
    : WHEN condition THEN multi_table_element+
    ;

conditional_insert_else_part
    : ELSE multi_table_element+
    ;

insert_into_clause
    : INTO general_table_ref paren_column_list?
    ;

values_clause
    : VALUES (expression | LEFT_PAREN expressions RIGHT_PAREN)
    ;

merge_statement
    : MERGE INTO selected_tableview USING selected_tableview ON LEFT_PAREN condition RIGHT_PAREN
      (merge_update_clause merge_insert_clause? | merge_insert_clause merge_update_clause?)?
      error_logging_clause?
    ;

// Merge Specific Clauses

merge_update_clause
    : WHEN MATCHED THEN UPDATE SET merge_element (COMMA merge_element)* where_clause? merge_update_delete_part?
    ;

merge_element
    : column_name EQUALS_OP expression
    ;

merge_update_delete_part
    : DELETE where_clause
    ;

merge_insert_clause
    : WHEN NOT MATCHED THEN INSERT paren_column_list?
      values_clause where_clause?
    ;

selected_tableview
    : (tableview_name | LEFT_PAREN select_statement RIGHT_PAREN) table_alias?
    ;

lock_table_statement
    : LOCK TABLE lock_table_element (COMMA lock_table_element)* IN lock_mode MODE wait_nowait_part?
    ;

wait_nowait_part
    : WAIT expression
    | NOWAIT
    ;

// Lock Specific Clauses

lock_table_element
    : tableview_name partition_extension_clause?
    ;

lock_mode
    : ROW SHARE
    | ROW EXCLUSIVE
    | SHARE UPDATE?
    | SHARE ROW EXCLUSIVE
    | EXCLUSIVE
    ;

// Common DDL Clauses

general_table_ref
    : (dml_table_expression_clause | ONLY LEFT_PAREN dml_table_expression_clause RIGHT_PAREN) table_alias?
    ;

static_returning_clause
    : (RETURNING | RETURN) expressions into_clause
    ;

error_logging_clause
    : LOG ERRORS error_logging_into_part? expression? error_logging_reject_part?
    ;

error_logging_into_part
    : INTO tableview_name
    ;

error_logging_reject_part
    : REJECT LIMIT (UNLIMITED | expression)
    ;

dml_table_expression_clause
    : table_collection_expression
    | LEFT_PAREN select_statement subquery_restriction_clause? RIGHT_PAREN
    | tableview_name sample_clause?
    | json_table_clause (AS identifier)?
    ;

table_collection_expression
    : (TABLE | THE) (LEFT_PAREN subquery RIGHT_PAREN | LEFT_PAREN expression RIGHT_PAREN outer_join_sign?)
    ;

subquery_restriction_clause
    : WITH (READ ONLY | CHECK OPTION (CONSTRAINT constraint_name)?)
    ;

sample_clause
    : SAMPLE BLOCK? LEFT_PAREN expression (COMMA expression)? RIGHT_PAREN seed_part?
    ;

seed_part
    : SEED LEFT_PAREN expression RIGHT_PAREN
    ;

// Expression & Condition

condition
    : expression
    | json_condition
    ;

json_condition
    : column_name IS NOT? JSON (FORMAT JSON)? (STRICT|LAX)? ((WITH|WITHOUT) UNIQUE KEYS)?
    | JSON_EQUAL LEFT_PAREN expressions RIGHT_PAREN
    ;

expressions
    : expression (COMMA expression)*
    ;

expression
    : cursor_expression
    | logical_expression
    ;

cursor_expression
    : CURSOR LEFT_PAREN subquery RIGHT_PAREN
    ;

logical_expression
    : unary_logical_expression
    | logical_expression AND logical_expression
    | logical_expression OR logical_expression
    ;

unary_logical_expression
    : NOT? multiset_expression (IS NOT? logical_operation)*
    ;

logical_operation:
    ( NULL_
    | NAN | PRESENT
    | INFINITE | A_LETTER SET | EMPTY
    | OF TYPE? LEFT_PAREN ONLY? type_spec (COMMA type_spec)* RIGHT_PAREN
    )
    ;

multiset_expression
    : relational_expression (multiset_type=(MEMBER | SUBMULTISET) OF? concatenation)?
    ;

relational_expression
    : relational_expression relational_operator relational_expression
    | compound_expression
    ;

compound_expression
    : concatenation
      (NOT? ( IN in_elements
            | BETWEEN between_elements
            | like_type=(LIKE | LIKEC | LIKE2 | LIKE4) concatenation (ESCAPE concatenation)?))?
    ;

relational_operator
    : EQUALS_OP
    | (NOT_EQUAL_OP | '<' '>' | '!' EQUALS_OP | '^' EQUALS_OP)
    | ('<' | '>') EQUALS_OP?
    ;

in_elements
    : LEFT_PAREN subquery RIGHT_PAREN
    | LEFT_PAREN concatenation (COMMA concatenation)* RIGHT_PAREN
    | constant
    | bind_variable
    | general_element
    ;

between_elements
    : concatenation AND concatenation
    ;

concatenation
    : model_expression
        (AT (LOCAL | TIME ZONE concatenation) | interval_expression)?
        (ON OVERFLOW (TRUNCATE | ERROR))?
    | concatenation op=(ASTERISK | SOLIDUS) concatenation
    | concatenation op=(PLUS_SIGN | MINUS_SIGN) concatenation
    | concatenation BAR BAR concatenation
    ;

interval_expression
    : DAY (LEFT_PAREN concatenation RIGHT_PAREN)? TO SECOND (LEFT_PAREN concatenation RIGHT_PAREN)?
    | YEAR (LEFT_PAREN concatenation RIGHT_PAREN)? TO MONTH
    | concatenation (SECOND | DAY | MONTH | YEAR)
    ;

model_expression
    : unary_expression ('[' model_expression_element ']')?
    ;

model_expression_element
    : (ANY | expression) (COMMA (ANY | expression))*
    | single_column_for_loop (COMMA single_column_for_loop)*
    | multi_column_for_loop
    ;

single_column_for_loop
    : FOR column_name
       ( IN LEFT_PAREN expressions? RIGHT_PAREN
       | (LIKE expression)? FROM fromExpr=expression TO toExpr=expression
         action_type=(INCREMENT | DECREMENT) action_expr=expression)
    ;

multi_column_for_loop
    : FOR paren_column_list
      IN LEFT_PAREN (subquery | LEFT_PAREN expressions? RIGHT_PAREN) RIGHT_PAREN
    ;

unary_expression
    : ('-' | PLUS_SIGN) unary_expression
    | PRIOR unary_expression
    | CONNECT_BY_ROOT unary_expression
    | /*TODO {input.LT(1).getText().equalsIgnoreCase("new") && !input.LT(2).getText().equals(".")}?*/ NEW unary_expression
    |  DISTINCT unary_expression
    |  ALL unary_expression
    |  /*TODO{(input.LA(1) == CASE || input.LA(2) == CASE)}?*/ case_statement/*[false]*/
    |  quantified_expression
    |  standard_function
    |  atom
    ;

case_statement /*TODO [boolean isStatementParameter]
TODO scope    {
    boolean isStatement;
}
@init    {$case_statement::isStatement = $isStatementParameter;}*/
    : searched_case_statement
    | simple_case_statement
    ;

// CASE

simple_case_statement
    : label_name? ck1=CASE expression simple_case_when_part+  case_else_part? END CASE? label_name?
    ;

simple_case_when_part
    : WHEN expression THEN (/*TODO{$case_statement::isStatement}?*/ seq_of_statements | expression)
    ;

searched_case_statement
    : label_name? ck1=CASE searched_case_when_part+ case_else_part? END CASE? label_name?
    ;

searched_case_when_part
    : WHEN expression THEN (/*TODO{$case_statement::isStatement}?*/ seq_of_statements | expression)
    ;

case_else_part
    : ELSE (/*{$case_statement::isStatement}?*/ seq_of_statements | expression)
    ;

atom
    : table_element outer_join_sign
    | bind_variable
    | constant_without_variable
    | general_element
    | LEFT_PAREN subquery RIGHT_PAREN subquery_operation_part*
    | LEFT_PAREN expressions RIGHT_PAREN
    ;

quantified_expression
    : (SOME | EXISTS | ALL | ANY) (LEFT_PAREN select_only_statement RIGHT_PAREN | LEFT_PAREN expression RIGHT_PAREN)
    ;

string_function
    : SUBSTR LEFT_PAREN expression COMMA expression (COMMA expression)? RIGHT_PAREN
    | TO_CHAR LEFT_PAREN (table_element | standard_function | expression)
                  (COMMA quoted_string)? (COMMA quoted_string)? RIGHT_PAREN
    | DECODE LEFT_PAREN expressions  RIGHT_PAREN
    | CHR LEFT_PAREN concatenation USING NCHAR_CS RIGHT_PAREN
    | NVL LEFT_PAREN expression COMMA expression RIGHT_PAREN
    | TRIM LEFT_PAREN ((LEADING | TRAILING | BOTH)? quoted_string? FROM)? concatenation RIGHT_PAREN
    | TO_DATE LEFT_PAREN (table_element | standard_function | expression)
       (DEFAULT concatenation ON CONVERSION ERROR)? (COMMA quoted_string  (COMMA quoted_string)? )? RIGHT_PAREN
    ;

standard_function
    : string_function
    | numeric_function_wrapper
    | json_function
    | other_function
    ;

//see as https://docs.oracle.com/en/database/oracle/oracle-database/21/sqlrf/JSON_ARRAY.html#GUID-46CDB3AF-5795-455B-85A8-764528CEC43B
json_function
    : JSON_ARRAY LEFT_PAREN  json_array_element ( COMMA json_array_element)* json_on_null_clause? json_return_clause? STRICT? RIGHT_PAREN
    | JSON_ARRAYAGG LEFT_PAREN expression (FORMAT JSON)? order_by_clause? json_on_null_clause? json_return_clause? STRICT? RIGHT_PAREN
    | JSON_OBJECT LEFT_PAREN json_object_content RIGHT_PAREN
    | JSON_OBJECTAGG LEFT_PAREN KEY? expression VALUE expression ((NULL_ | ABSENT) ON NULL_)? (RETURNING ( VARCHAR2 (LEFT_PAREN UNSIGNED_INTEGER ( BYTE | CHAR )? RIGHT_PAREN)? | CLOB | BLOB ))?  STRICT? (WITH UNIQUE KEYS)?RIGHT_PAREN
    | JSON_QUERY LEFT_PAREN expression (FORMAT JSON)? COMMA CHAR_STRING json_query_returning_clause? json_query_wrapper_clause? json_query_on_error_clause? json_query_on_empty_clause? RIGHT_PAREN
    | JSON_SERIALIZE LEFT_PAREN CHAR_STRING (RETURNING json_query_return_type)? PRETTY? ASCII? TRUNCATE? ((NULL_ | ERROR | EMPTY (ARRAY | OBJECT)) ON ERROR)? RIGHT_PAREN
    | JSON_TRANSFORM LEFT_PAREN expression COMMA json_transform_op (COMMA json_transform_op)* RIGHT_PAREN
    | JSON_VALUE LEFT_PAREN  expression (FORMAT JSON)? (COMMA CHAR_STRING? json_value_return_clause? ((ERROR | NULL_ | DEFAULT literal)? ON ERROR)? ((ERROR | NULL_ | DEFAULT literal)? ON EMPTY)? json_value_on_mismatch_clause?RIGHT_PAREN)?
    ;

json_object_content
    : (json_object_entry (COMMA json_object_entry)* | '*') json_on_null_clause? json_return_clause? STRICT? (WITH UNIQUE KEYS)?
    ;

json_object_entry
    : KEY? expression (VALUE | IS)? expression
    | expression ':' expression (FORMAT JSON)?
    | identifier
    ;

json_table_clause
    : JSON_TABLE LEFT_PAREN expression (FORMAT JSON)? (COMMA CHAR_STRING)? ((ERROR | NULL_) ON ERROR)? ((EMPTY | NULL_) ON EMPTY)? json_column_clause? RIGHT_PAREN
    ;

json_array_element
    : (expression | CHAR_STRING | NULL_ |  UNSIGNED_INTEGER | json_function) (FORMAT JSON)?
    ;

json_on_null_clause
    : (NULL_ | ABSENT) ON NULL_
    ;

json_return_clause
    : RETURNING ( VARCHAR2 (LEFT_PAREN UNSIGNED_INTEGER ( BYTE | CHAR )? RIGHT_PAREN)? | CLOB | BLOB )
    ;

json_transform_op
    : REMOVE CHAR_STRING ((IGNORE | ERROR)? ON MISSING)?
    | INSERT CHAR_STRING EQUALS_OP CHAR_STRING ((REPLACE | IGNORE | ERROR) ON EXISTING)? ((NULL_ | IGNORE | ERROR | REMOVE)? ON NULL_)?
    | REPLACE CHAR_STRING EQUALS_OP CHAR_STRING ((CREATE | IGNORE | ERROR) ON MISSING)? ((NULL_ | IGNORE | ERROR)? ON NULL_)?
    | expression (FORMAT JSON)?
    | APPEND CHAR_STRING EQUALS_OP CHAR_STRING ((CREATE | IGNORE | ERROR) ON MISSING)? ((NULL_ | IGNORE | ERROR)? ON NULL_)?
    | SET CHAR_STRING EQUALS_OP expression (FORMAT JSON)? ((REPLACE | IGNORE | ERROR) ON EXISTING)? ((CREATE | IGNORE | ERROR) ON MISSING)? ((NULL_ | IGNORE | ERROR)? ON NULL_)?
    ;

json_column_clause
    : COLUMNS LEFT_PAREN json_column_definition (COMMA json_column_definition)* RIGHT_PAREN
    ;

json_column_definition
    : expression json_value_return_type? (EXISTS? PATH CHAR_STRING | TRUNCATE (PATH CHAR_STRING)?)? json_query_on_error_clause? json_query_on_empty_clause?
    | expression json_query_return_type? TRUNCATE? FORMAT JSON  json_query_wrapper_clause? PATH CHAR_STRING
    | NESTED PATH? expression ('[' ASTERISK ']')? json_column_clause
    | expression FOR ORDINALITY
    ;

json_query_returning_clause
    : (RETURNING json_query_return_type)? PRETTY? ASCII?
    ;

json_query_return_type
    : VARCHAR2 (LEFT_PAREN UNSIGNED_INTEGER ( BYTE | CHAR )? RIGHT_PAREN)? | CLOB | BLOB
    ;

json_query_wrapper_clause
    : (WITHOUT ARRAY? WRAPPER) | (WITH (UNCONDITIONAL | CONDITIONAL)? ARRAY? WRAPPER)
    ;

json_query_on_error_clause
    : (ERROR | NULL_ | EMPTY | EMPTY ARRAY | EMPTY OBJECT)? ON ERROR
    ;

json_query_on_empty_clause
    : (ERROR | NULL_ | EMPTY | EMPTY ARRAY | EMPTY OBJECT)? ON EMPTY
    ;

json_value_return_clause
    : RETURNING json_value_return_type? ASCII?
    ;

json_value_return_type
    : VARCHAR2 (LEFT_PAREN UNSIGNED_INTEGER  ( BYTE | CHAR )? RIGHT_PAREN)? TRUNCATE?
    | CLOB
    | DATE
    | NUMBER LEFT_PAREN INTEGER (COMMA INTEGER)? RIGHT_PAREN
    | TIMESTAMP (WITH TIMEZONE)?
    | SDO_GEOMETRY
    | expression (USING CASESENSITIVE MAPPING)?
    ;

json_value_on_mismatch_clause
    : (IGNORE | ERROR | NULL_) ON MISMATCH (LEFT_PAREN MISSING DATA | EXTRA DATA | TYPE ERROR RIGHT_PAREN)?
    ;

literal
    : CHAR_STRING
    | string_function
    | numeric
    | MAXVALUE
    ;

numeric_function_wrapper
    : numeric_function (single_column_for_loop | multi_column_for_loop)?
    ;

numeric_function
   : SUM LEFT_PAREN (DISTINCT | ALL)? expression RIGHT_PAREN
   | COUNT LEFT_PAREN ( ASTERISK | ((DISTINCT | UNIQUE | ALL)? concatenation)? ) RIGHT_PAREN over_clause?
   | ROUND LEFT_PAREN expression (COMMA UNSIGNED_INTEGER)?  RIGHT_PAREN
   | AVG LEFT_PAREN (DISTINCT | ALL)? expression RIGHT_PAREN
   | MAX LEFT_PAREN (DISTINCT | ALL)? expression RIGHT_PAREN
   | LEAST LEFT_PAREN expressions RIGHT_PAREN
   | GREATEST LEFT_PAREN expressions RIGHT_PAREN
   ;

listagg_overflow_clause
    : ON OVERFLOW (ERROR | TRUNCATE) CHAR_STRING? ((WITH | WITHOUT) COUNT)?
    ;

other_function
    : over_clause_keyword function_argument_analytic over_clause?
    | /*TODO stantard_function_enabling_using*/ regular_id function_argument_modeling using_clause?
    | COUNT LEFT_PAREN ( ASTERISK | (DISTINCT | UNIQUE | ALL)? concatenation) RIGHT_PAREN over_clause?
    | (CAST | XMLCAST) LEFT_PAREN (MULTISET LEFT_PAREN subquery RIGHT_PAREN | concatenation) AS type_spec
      (DEFAULT concatenation ON CONVERSION ERROR)? (COMMA quoted_string (COMMA quoted_string)? )? RIGHT_PAREN
    | COALESCE LEFT_PAREN table_element (COMMA (numeric | quoted_string))? RIGHT_PAREN
    | COLLECT LEFT_PAREN (DISTINCT | UNIQUE)? concatenation collect_order_by_part? RIGHT_PAREN
    | within_or_over_clause_keyword function_argument within_or_over_part+
    | LISTAGG LEFT_PAREN (ALL | DISTINCT | UNIQUE)? argument (COMMA CHAR_STRING)? listagg_overflow_clause? RIGHT_PAREN
      (WITHIN GROUP LEFT_PAREN order_by_clause RIGHT_PAREN)? over_clause?
    | cursor_name ( PERCENT_ISOPEN | PERCENT_FOUND | PERCENT_NOTFOUND | PERCENT_ROWCOUNT )
    | DECOMPOSE LEFT_PAREN concatenation (CANONICAL | COMPATIBILITY)? RIGHT_PAREN
    | EXTRACT LEFT_PAREN regular_id FROM concatenation RIGHT_PAREN
    | (FIRST_VALUE | LAST_VALUE) function_argument_analytic respect_or_ignore_nulls? over_clause
    | standard_prediction_function_keyword
      LEFT_PAREN expressions cost_matrix_clause? using_clause? RIGHT_PAREN
    | (TO_BINARY_DOUBLE | TO_BINARY_FLOAT | TO_NUMBER | TO_TIMESTAMP | TO_TIMESTAMP_TZ)
      LEFT_PAREN concatenation (DEFAULT concatenation ON CONVERSION ERROR)? (COMMA quoted_string (COMMA quoted_string)? )? RIGHT_PAREN
    | (TO_DSINTERVAL | TO_YMINTERVAL) LEFT_PAREN concatenation (DEFAULT concatenation ON CONVERSION ERROR)? RIGHT_PAREN
    | TRANSLATE LEFT_PAREN expression (USING (CHAR_CS | NCHAR_CS))? (COMMA expression)* RIGHT_PAREN
    | TREAT LEFT_PAREN expression AS REF? type_spec RIGHT_PAREN
    | TRIM LEFT_PAREN ((LEADING | TRAILING | BOTH)? quoted_string? FROM)? concatenation RIGHT_PAREN
    | VALIDATE_CONVERSION LEFT_PAREN concatenation AS type_spec (COMMA quoted_string (COMMA quoted_string)? )? RIGHT_PAREN
    | XMLAGG LEFT_PAREN expression order_by_clause? RIGHT_PAREN (PERIOD general_element)?
    | (XMLCOLATTVAL | XMLFOREST)
      LEFT_PAREN xml_multiuse_expression_element (COMMA xml_multiuse_expression_element)* RIGHT_PAREN (PERIOD general_element)?
    | XMLELEMENT
      LEFT_PAREN (ENTITYESCAPING | NOENTITYESCAPING)? (NAME | EVALNAME)? expression
       (/*TODO{input.LT(2).getText().equalsIgnoreCase("xmlattributes")}?*/ COMMA xml_attributes_clause)?
       (COMMA expression column_alias?)* RIGHT_PAREN (PERIOD general_element)?
    | XMLEXISTS LEFT_PAREN expression xml_passing_clause? RIGHT_PAREN
    | XMLPARSE LEFT_PAREN (DOCUMENT | CONTENT) concatenation WELLFORMED? RIGHT_PAREN (PERIOD general_element)?
    | XMLPI
      LEFT_PAREN (NAME identifier | EVALNAME concatenation) (COMMA concatenation)? RIGHT_PAREN (PERIOD general_element)?
    | XMLQUERY
      LEFT_PAREN concatenation xml_passing_clause? RETURNING CONTENT (NULL_ ON EMPTY)? RIGHT_PAREN (PERIOD general_element)?
    | XMLROOT
      LEFT_PAREN concatenation (COMMA xmlroot_param_version_part)? (COMMA xmlroot_param_standalone_part)? RIGHT_PAREN (PERIOD general_element)?
    | XMLSERIALIZE
      LEFT_PAREN (DOCUMENT | CONTENT) concatenation (AS type_spec)?
      xmlserialize_param_enconding_part? xmlserialize_param_version_part? xmlserialize_param_ident_part? ((HIDE | SHOW) DEFAULTS)? RIGHT_PAREN
      (PERIOD general_element)?
    | TIME CHAR_STRING
    | xmltable
    ;

over_clause_keyword
    : AVG
    | CORR
    | LAG
    | LAG_DIFF
    | LAG_DIFF_PERCENT
    | LEAD
    | MAX
    | MEDIAN
    | MIN
    | NTILE
    | RATIO_TO_REPORT
    | ROW_NUMBER
    | SUM
    | VARIANCE
    | REGR_
    | STDDEV
    | VAR_
    | COVAR_
    ;

within_or_over_clause_keyword
    : CUME_DIST
    | DENSE_RANK
    | PERCENT_RANK
    | PERCENTILE_CONT
    | PERCENTILE_DISC
    | RANK
    ;

standard_prediction_function_keyword
    : PREDICTION
    | PREDICTION_BOUNDS
    | PREDICTION_COST
    | PREDICTION_DETAILS
    | PREDICTION_PROBABILITY
    | PREDICTION_SET
    ;

over_clause
    : OVER LEFT_PAREN ( query_partition_clause? (order_by_clause windowing_clause?)?
               | HIERARCHY th=id_expression OFFSET numeric (ACROSS ANCESTOR AT LEVEL id_expression)?
               )
           RIGHT_PAREN
    ;

windowing_clause
    : windowing_type
      (BETWEEN windowing_elements AND windowing_elements | windowing_elements)
    ;

windowing_type
    : ROWS
    | RANGE
    ;

windowing_elements
    : UNBOUNDED PRECEDING
    | CURRENT ROW
    | concatenation (PRECEDING | FOLLOWING)
    ;

using_clause
    : USING (ASTERISK | using_element (COMMA using_element)*)
    ;

using_element
    : (IN OUT? | OUT)? select_list_elements
    ;

collect_order_by_part
    : ORDER BY concatenation
    ;

within_or_over_part
    : WITHIN GROUP LEFT_PAREN order_by_clause RIGHT_PAREN
    | over_clause
    ;

cost_matrix_clause
    : COST (MODEL AUTO? | LEFT_PAREN cost_class_name (COMMA cost_class_name)* RIGHT_PAREN VALUES LEFT_PAREN expressions? RIGHT_PAREN)
    ;

xml_passing_clause
    : PASSING (BY VALUE)? expression column_alias? (COMMA expression column_alias?)*
    ;

xml_attributes_clause
    : XMLATTRIBUTES
     LEFT_PAREN (ENTITYESCAPING | NOENTITYESCAPING)? (SCHEMACHECK | NOSCHEMACHECK)?
     xml_multiuse_expression_element (COMMA xml_multiuse_expression_element)* RIGHT_PAREN
    ;

xml_namespaces_clause
    : XMLNAMESPACES
      LEFT_PAREN (concatenation column_alias)? (COMMA concatenation column_alias)* xml_general_default_part? RIGHT_PAREN
    ;

xml_table_column
    : xml_column_name
      (FOR ORDINALITY | type_spec (PATH concatenation)? xml_general_default_part?)
    ;

xml_general_default_part
    : DEFAULT concatenation
    ;

xml_multiuse_expression_element
    : expression (AS (id_expression | EVALNAME concatenation))?
    ;

xmlroot_param_version_part
    : VERSION (NO VALUE | expression)
    ;

xmlroot_param_standalone_part
    : STANDALONE (YES | NO VALUE?)
    ;

xmlserialize_param_enconding_part
    : ENCODING concatenation
    ;

xmlserialize_param_version_part
    : VERSION concatenation
    ;

xmlserialize_param_ident_part
    : NO INDENT
    | INDENT (SIZE EQUALS_OP concatenation)?
    ;

// SqlPlus

sql_plus_command
    : '/'
    | EXIT
    | PROMPT_MESSAGE
    | SHOW (ERR | ERRORS)
    | START_CMD
    | whenever_command
    | set_command
    | timing_command
    ;

whenever_command
    : WHENEVER (SQLERROR | OSERROR)
         ( EXIT (SUCCESS | FAILURE | WARNING | variable_name | numeric) (COMMIT | ROLLBACK)?
         | CONTINUE (COMMIT | ROLLBACK | NONE)?)
    ;

set_command
    : SET regular_id (CHAR_STRING | ON | OFF | /*EXACT_NUM_LIT*/numeric | regular_id)
    ;

timing_command
    : TIMING (START timing_text=id_expression* | SHOW | STOP)?
    ;

// Common

partition_extension_clause
    : (SUBPARTITION | PARTITION) FOR? LEFT_PAREN expressions? RIGHT_PAREN
    ;

column_alias
    : AS? (identifier | quoted_string)
    | AS
    ;

table_alias
    : identifier
    | quoted_string
    ;

where_clause
    : WHERE (CURRENT OF cursor_name | expression | quantitative_where_stmt)
    ;

quantitative_where_stmt
    : expression relational_operator (SOME | ALL | ANY) LEFT_PAREN expression (COMMA expression)* RIGHT_PAREN
    ;

into_clause
    : (BULK COLLECT)? INTO (general_element | bind_variable) (COMMA (general_element | bind_variable))*
    ;

// Common Named Elements

xml_column_name
    : identifier
    | quoted_string
    ;

cost_class_name
    : identifier
    ;

attribute_name
    : identifier
    ;

savepoint_name
    : identifier
    ;

rollback_segment_name
    : identifier
    ;

schema_name
    : identifier
    ;

routine_name
    : identifier (PERIOD id_expression)* (AT_SIGN link_name)?
    ;

package_name
    : identifier
    ;

implementation_type_name
    : identifier (PERIOD id_expression)?
    ;

parameter_name
    : identifier
    ;

reference_model_name
    : identifier
    ;

main_model_name
    : identifier
    ;

container_tableview_name
    : identifier (PERIOD id_expression)?
    ;

aggregate_function_name
    : identifier (PERIOD id_expression)*
    ;

query_name
    : identifier
    ;

grantee_name
    : id_expression identified_by?
    ;

role_name
    : id_expression
    | CONNECT
    | RESOURCE
    ;

constraint_name
    : identifier (PERIOD id_expression)* (AT_SIGN link_name)?
    ;

label_name
    : id_expression
    ;

type_name
    : id_expression (PERIOD id_expression)*
    ;

sequence_name
    : id_expression (PERIOD id_expression)*
    ;

exception_name
    : identifier (PERIOD id_expression)*
    ;

function_name
    : identifier (PERIOD id_expression)?
    ;

procedure_name
    : identifier (PERIOD id_expression)?
    ;

trigger_name
    : identifier (PERIOD id_expression)?
    ;

variable_name
    : (INTRODUCER char_set_name)? id_expression (PERIOD id_expression)?
    | bind_variable
    ;

index_name
    : identifier (PERIOD id_expression)?
    ;

cursor_name
    : general_element
    | bind_variable
    ;

record_name
    : identifier
    | bind_variable
    ;

collection_name
    : identifier (PERIOD id_expression)?
    ;

// BYT-8268: Database link names can include domain qualifiers (e.g., linkname.domain)
link_name
    : identifier (PERIOD id_expression)*
    ;

column_name
    : identifier (PERIOD id_expression)*
    ;

tableview_name
    : identifier (PERIOD id_expression)?
          (AT_SIGN link_name (PERIOD link_name)* | /*TODO{!(input.LA(2) == BY)}?*/ partition_extension_clause)?
    | xmltable outer_join_sign?
    ;

xmltable
    : XMLTABLE LEFT_PAREN (xml_namespaces_clause COMMA)? concatenation xml_passing_clause? (COLUMNS xml_table_column (COMMA xml_table_column)*)? RIGHT_PAREN (PERIOD general_element)?
    ;

char_set_name
    : id_expression (PERIOD id_expression)*
    ;

synonym_name
    : identifier
    ;

// Represents a valid DB object name in DDL commands which are valid for several DB (or schema) objects.
// For instance, create synonym ... for <DB object name>, or rename <old DB object name> to <new DB object name>.
// Both are valid for sequences, tables, views, etc.
schema_object_name
    : id_expression
    ;

dir_object_name
    : id_expression
    ;

user_object_name
    : id_expression
    ;

grant_object_name
    : tableview_name
    | USER user_object_name (COMMA user_object_name)*
    | DIRECTORY dir_object_name
    | EDITION schema_object_name
    | MINING MODEL schema_object_name
    | JAVA (SOURCE | RESOURCE) schema_object_name
    | SQL TRANSLATION PROFILE schema_object_name
    ;

column_list
    : column_name (COMMA column_name)*
    ;

paren_column_list
    : LEFT_PAREN column_list RIGHT_PAREN
    ;

// PL/SQL Specs

// NOTE: In reality this applies to aggregate functions only
keep_clause
    : KEEP LEFT_PAREN DENSE_RANK (FIRST | LAST) order_by_clause RIGHT_PAREN over_clause?
    ;

function_argument
    : LEFT_PAREN (argument (COMMA argument)*)? RIGHT_PAREN keep_clause?
    ;

function_argument_analytic
    : LEFT_PAREN (argument respect_or_ignore_nulls? (COMMA argument respect_or_ignore_nulls?)*)? RIGHT_PAREN keep_clause?
    ;

function_argument_modeling
    : LEFT_PAREN column_name (COMMA (numeric | NULL_) (COMMA (numeric | NULL_))?)?
      USING (tableview_name PERIOD ASTERISK | ASTERISK | expression column_alias? (COMMA expression column_alias?)*)
      RIGHT_PAREN keep_clause?
    ;

respect_or_ignore_nulls
    : (RESPECT | IGNORE) NULLS
    ;

argument
    : (identifier EQUALS_OP '>')? expression
    ;

type_spec
    : datatype
    | REF? type_name (PERCENT_ROWTYPE | PERCENT_TYPE)?
    ;

datatype
    : native_datatype_element precision_part? (WITH LOCAL? TIME ZONE | CHARACTER SET char_set_name)?
    | INTERVAL (YEAR | DAY) (LEFT_PAREN expression RIGHT_PAREN)? TO (MONTH | SECOND) (LEFT_PAREN expression RIGHT_PAREN)?
    ;

precision_part
    : LEFT_PAREN (numeric | ASTERISK) (COMMA (numeric | numeric_negative))? (CHAR | BYTE)? RIGHT_PAREN
    ;

native_datatype_element
    : BINARY_INTEGER
    | PLS_INTEGER
    | NATURAL
    | BINARY_FLOAT
    | BINARY_DOUBLE
    | NATURALN
    | POSITIVE
    | POSITIVEN
    | SIGNTYPE
    | SIMPLE_INTEGER
    | NVARCHAR2
    | DEC
    | INTEGER
    | INT
    | NUMERIC
    | SMALLINT
    | NUMBER
    | DECIMAL
    | DOUBLE PRECISION?
    | FLOAT
    | REAL
    | NCHAR
    | LONG RAW?
    | CHAR
    | CHARACTER
    | VARCHAR2
    | VARCHAR
    | STRING
    | RAW
    | BOOLEAN
    | DATE
    | ROWID
    | UROWID
    | YEAR
    | MONTH
    | DAY
    | HOUR
    | MINUTE
    | SECOND
    | TIMEZONE_HOUR
    | TIMEZONE_MINUTE
    | TIMEZONE_REGION
    | TIMEZONE_ABBR
    | TIMESTAMP
    | TIMESTAMP_UNCONSTRAINED
    | TIMESTAMP_TZ_UNCONSTRAINED
    | TIMESTAMP_LTZ_UNCONSTRAINED
    | YMINTERVAL_UNCONSTRAINED
    | DSINTERVAL_UNCONSTRAINED
    | BFILE
    | BLOB
    | CLOB
    | NCLOB
    | MLSLABEL
    ;

bind_variable
    : (BINDVAR | ':' UNSIGNED_INTEGER)
      // Pro*C/C++ indicator variables
      (INDICATOR? (BINDVAR | ':' UNSIGNED_INTEGER))?
      (PERIOD general_element)?
    ;

general_element
    : general_element_part (PERIOD general_element_part)*
    ;

// general_element_part contains the variable_name rule.
general_element_part
    : (INTRODUCER char_set_name)? id_expression (AT_SIGN link_name)? function_argument?
    ;

table_element
    : (INTRODUCER char_set_name)? id_expression (PERIOD id_expression)*
    ;

object_privilege
    : ALL PRIVILEGES?
    | ALTER
    | DEBUG
    | DELETE
    | EXECUTE
    | FLASHBACK ARCHIVE
    | INDEX
    | INHERIT PRIVILEGES
    | INSERT
    | KEEP SEQUENCE
    | MERGE VIEW
    | ON COMMIT REFRESH
    | QUERY REWRITE
    | READ
    | REFERENCES
    | SELECT
    | TRANSLATE SQL
    | UNDER
    | UPDATE
    | USE
    | WRITE
    ;

//Ordered by type rather than alphabetically
system_privilege
    : ALL PRIVILEGES
    | ADVISOR
    | ADMINISTER ANY? SQL TUNING SET
    | (ALTER | CREATE | DROP) ANY SQL PROFILE
    | ADMINISTER SQL MANAGEMENT OBJECT
    | CREATE ANY? CLUSTER
    | (ALTER | DROP) ANY CLUSTER
    | (CREATE | DROP) ANY CONTEXT
    | EXEMPT REDACTION POLICY
    | ALTER DATABASE
    | (ALTER | CREATE) PUBLIC? DATABASE LINK
    | DROP PUBLIC DATABASE LINK
    | DEBUG CONNECT SESSION
    | DEBUG ANY PROCEDURE
    | ANALYZE ANY DICTIONARY
    | CREATE ANY? DIMENSION
    | (ALTER | DROP) ANY DIMENSION
    | (CREATE | DROP) ANY DIRECTORY
    | (CREATE | DROP) ANY EDITION
    | FLASHBACK (ARCHIVE ADMINISTER | ANY TABLE)
    | (ALTER | CREATE | DROP) ANY INDEX
    | CREATE ANY? INDEXTYPE
    | (ALTER | DROP | EXECUTE) ANY INDEXTYPE
    | CREATE (ANY | EXTERNAL)? JOB
    | EXECUTE ANY (CLASS | PROGRAM)
    | MANAGE SCHEDULER
    | ADMINISTER KEY MANAGEMENT
    | CREATE ANY? LIBRARY
    | (ALTER | DROP | EXECUTE) ANY LIBRARY
    | LOGMINING
    | CREATE ANY? MATERIALIZED VIEW
    | (ALTER | DROP) ANY MATERIALIZED VIEW
    | GLOBAL? QUERY REWRITE
    | ON COMMIT REFRESH
    | CREATE ANY? MINING MODEL
    | (ALTER | DROP | SELECT | COMMENT) ANY MINING MODEL
    | CREATE ANY? CUBE
    | (ALTER | DROP | SELECT | UPDATE) ANY CUBE
    | CREATE ANY? MEASURE FOLDER
    | (DELETE | DROP | INSERT) ANY MEASURE FOLDER
    | CREATE ANY? CUBE DIMENSION
    | (ALTER | DELETE | DROP | INSERT | SELECT | UPDATE) ANY CUBE DIMENSION
    | CREATE ANY? CUBE BUILD PROCESS
    | (DROP | UPDATE) ANY CUBE BUILD PROCESS
    | CREATE ANY? OPERATOR
    | (ALTER | DROP | EXECUTE) ANY OPERATOR
    | (CREATE | ALTER | DROP) ANY OUTLINE
    | CREATE PLUGGABLE DATABASE
    | SET CONTAINER
    | CREATE ANY? PROCEDURE
    | (ALTER | DROP | EXECUTE) ANY PROCEDURE
    | (CREATE | ALTER | DROP ) PROFILE
    | CREATE ROLE
    | (ALTER | DROP | GRANT) ANY ROLE
    | (CREATE | ALTER | DROP) ROLLBACK SEGMENT
    | CREATE ANY? SEQUENCE
    | (ALTER | DROP | SELECT) ANY SEQUENCE
    | (ALTER | CREATE | RESTRICTED) SESSION
    | ALTER RESOURCE COST
    | CREATE ANY? SQL TRANSLATION PROFILE
    | (ALTER | DROP | USE) ANY SQL TRANSLATION PROFILE
    | TRANSLATE ANY SQL
    | CREATE ANY? SYNONYM
    | DROP ANY SYNONYM
    | (CREATE | DROP) PUBLIC SYNONYM
    | CREATE ANY? TABLE
    | (ALTER | BACKUP | COMMENT | DELETE | DROP | INSERT | LOCK | READ | SELECT | UPDATE) ANY TABLE
    | (CREATE | ALTER | DROP | MANAGE | UNLIMITED) TABLESPACE
    | CREATE ANY? TRIGGER
    | (ALTER | DROP) ANY TRIGGER
    | ADMINISTER DATABASE TRIGGER
    | CREATE ANY? TYPE
    | (ALTER | DROP | EXECUTE | UNDER) ANY TYPE
    | (CREATE | ALTER | DROP) USER
    | CREATE ANY? VIEW
    | (DROP | UNDER | MERGE) ANY VIEW
    | (ANALYZE | AUDIT) ANY
    | BECOME USER
    | CHANGE NOTIFICATION
    | EXEMPT ACCESS POLICY
    | FORCE ANY? TRANSACTION
    | GRANT ANY OBJECT? PRIVILEGE
    | INHERIT ANY PRIVILEGES
    | KEEP DATE TIME
    | KEEP SYSGUID
    | PURGE DBA_RECYCLEBIN
    | RESUMABLE
    | SELECT ANY (DICTIONARY | TRANSACTION)
    | SYSBACKUP
    | SYSDBA
    | SYSDG
    | SYSKM
    | SYSOPER
    ;

// $>

// $<Lexer Mappings

constant
    : TIMESTAMP (quoted_string | bind_variable) (AT TIME ZONE quoted_string)?
    | INTERVAL (quoted_string | bind_variable | general_element)
      (YEAR | MONTH | DAY | HOUR | MINUTE | SECOND)
      (LEFT_PAREN (UNSIGNED_INTEGER | bind_variable) (COMMA (UNSIGNED_INTEGER | bind_variable) )? RIGHT_PAREN)?
      (TO ( DAY | HOUR | MINUTE | SECOND (LEFT_PAREN (UNSIGNED_INTEGER | bind_variable) RIGHT_PAREN)?))?
    | numeric
    | DATE quoted_string
    | quoted_string
    | NULL_
    | TRUE
    | FALSE
    | DBTIMEZONE
    | SESSIONTIMEZONE
    | MINVALUE
    | MAXVALUE
    | DEFAULT
    ;
    
constant_without_variable
    : TIMESTAMP (quoted_string | bind_variable) (AT TIME ZONE quoted_string)?
    | INTERVAL (quoted_string | bind_variable | general_element)
      (YEAR | MONTH | DAY | HOUR | MINUTE | SECOND)
      (LEFT_PAREN (UNSIGNED_INTEGER | bind_variable) (COMMA (UNSIGNED_INTEGER | bind_variable) )? RIGHT_PAREN)?
      (TO ( DAY | HOUR | MINUTE | SECOND (LEFT_PAREN (UNSIGNED_INTEGER | bind_variable) RIGHT_PAREN)?))?
    | numeric
    | DATE quoted_string
    | char_str
    | NULL_
    | TRUE
    | FALSE
    | DBTIMEZONE
    | SESSIONTIMEZONE
    | MINVALUE
    | MAXVALUE
    | DEFAULT
    ;

numeric
    : UNSIGNED_INTEGER
    | APPROXIMATE_NUM_LIT
    ;

numeric_negative
    : MINUS_SIGN numeric
    ;

quoted_string
    : variable_name
    | CHAR_STRING
    //| CHAR_STRING_PERL
    | NATIONAL_CHAR_STRING_LIT
    ;
    
char_str
    : CHAR_STRING
    | NATIONAL_CHAR_STRING_LIT
    ;

identifier
    : (INTRODUCER char_set_name)? id_expression
    ;

id_expression
    : regular_id
    | DELIMITED_ID
    ;

outer_join_sign
    : LEFT_PAREN PLUS_SIGN RIGHT_PAREN
    ;

// === Rules added from upstream for enhanced Oracle support ===

// Oracle 21c+ annotations support
annotations_clause
    : ANNOTATIONS '(' annotations_list ')'
    ;

annotations_list
    : (ADD (IF NOT EXISTS | OR REPLACE)? | DROP (IF EXISTS)? | REPLACE)? annotation (',' annotations_list)*
    ;

annotation
    : identifier CHAR_STRING?
    ;

// CREATE SCHEMA statement
create_schema
    : CREATE SCHEMA AUTHORIZATION schema_name (create_table | create_view | grant_statement)*
    ;

// DROP MATERIALIZED VIEW LOG statement
drop_materialized_view_log
    : DROP MATERIALIZED VIEW LOG (IF EXISTS)? ON tableview_name
    ;

// Compound trigger support
compound_trigger_block
    : COMPOUND TRIGGER seq_of_declare_specs? timing_point_section+ END trigger_name?
    ;

timing_point_section
    : bk = BEFORE STATEMENT IS tps_block BEFORE STATEMENT ';'
    | bk = BEFORE EACH ROW IS tps_block BEFORE EACH ROW ';'
    | ak = AFTER STATEMENT IS tps_block AFTER STATEMENT ';'
    | ak = AFTER EACH ROW IS tps_block AFTER EACH ROW ';'
    ;

tps_block
    : declare_spec* body
    ;

// Partition operations
move_table_partition
    : MOVE (partition_extended_names (MAPPING TABLE)? table_partition_description
          | subpartition_extended_names indexing_clause? partitioning_storage_clause?)
      (filter_condition | update_index_clauses | parallel_clause | allow_or_disallow CLUSTERING | ONLINE)*
    ;

rename_table_partition
    : RENAME (partition_extended_names | subpartition_extended_names) TO partition_name
    ;

// Implicit cursor expressions (SQL%BULK_ROWCOUNT, SQL%BULK_EXCEPTIONS)
implicit_cursor_expression
    : SQL (PERCENT_BULK_ROWCOUNT '(' expression ')'
         | PERCENT_BULK_EXCEPTIONS ('.' COUNT | '(' expression ')' '.' (ERROR_INDEX | ERROR_CODE)))
    ;

// Preprocessor directives
inquiry_directive
    : INQUIRY_DIRECTIVE
    ;

error_directive
    : DOLLAR_ERROR concatenation DOLLAR_END
    ;

selection_directive
    : DOLLAR_IF condition DOLLAR_THEN selection_directive_body
      (DOLLAR_ELSIF selection_directive_body)* (DOLLAR_ELSE selection_directive_body)? DOLLAR_END
    ;

selection_directive_body
    : (pragma_declaration? statement ';' | variable_declaration | error_directive | function_body | procedure_body)+
    ;

// Pipelined functions with USING clause
pipelined_using_clause
    : PIPELINED ((ROW | TABLE) POLYMORPHIC)? USING implementation_type_name
    ;

// Accessible by clause for package visibility
accessible_by_clause
    : ACCESSIBLE BY '(' accessor (',' accessor)* ')'
    ;

accessor
    : (FUNCTION | PROCEDURE | PACKAGE | TRIGGER | TYPE) function_name
    ;

// Default collation clause
default_collation_clause
    : DEFAULT COLLATION USING_NLS_COMP
    ;

// Helper rule for filter condition in partition operations
filter_condition
    : INCLUDING ROWS where_clause
    ;

// === Extended upstream rules for more complete Oracle support ===

// C external call parameter support (extends c_parameters_clause)
c_external_parameter
    : CONTEXT
    | SELF (TDO | c_property)?
    | (parameter_name | RETURN) c_property? (BY REFERENCE)? external_datatype=regular_id?
    ;

c_property
    : INDICATOR (STRUCT | TDO)?
    | LENGTH
    | DURATION
    | MAXLEN
    | CHARSETID
    | CHARSETFORM
    ;

// Analytic view hierarchies
hierarchies_clause
    : HIERARCHIES '(' hier_alias+=object_name (',' hier_alias+=object_name)* ')'
    ;

// Filter clauses for analytic views
filter_clause
    : (MEASURES | hier_alias=object_name) TO condition
    ;

filter_clauses
    : FILTER FACT '(' filter_clause (',' filter_clause)* ')'
    ;

// Subanalytic view support
subav_clause
    : USING subav_name=object_name hierarchies_clause? filter_clauses? add_calcs_clause?
    ;

subav_factoring_clause
    : subav_name=id_expression ANALYTIC VIEW AS '(' subav_clause ')'
    ;

// Analytic view calculated measures
add_calc_meas_clause
    : meas_name=id_expression AS '(' expression ')'
    ;

add_calcs_clause
    : ADD MEASURES '(' add_calc_meas_clause (',' add_calc_meas_clause)* ')'
    ;

// Aggregate clause for user-defined aggregate functions
aggregate_clause
    : AGGREGATE USING implementation_type_name
    ;

// Parallel instances clause
parallel_instances_clause
    : INSTANCES (UNSIGNED_INTEGER | DEFAULT)
    ;

// Overriding procedure spec for object types
overriding_procedure_spec
    : PROCEDURE procedure_name ('(' type_elements_parameter (',' type_elements_parameter)* ')')?
      (IS | AS) (call_spec | DECLARE? seq_of_declare_specs? body ';')
    ;

// Assignable element (assignment target)
assignable_element
    : general_element
    | bind_variable
    ;

// Database link connection qualifier
connection_qualifier
    : identifier
    ;

local_link_name
    : identifier
    ;

// Statistics clause
by_user_for_statistics_clause
    : BY USER FOR STATISTICS
    ;

// Unary logical operation
unary_logical_operation
    : IS NOT? logical_operation
    ;

// Variable or collection reference
variable_or_collection
    : variable_name
    | collection_expression
    ;

// Collection expression
collection_expression
    : collation_name '(' expression ')' ('.' general_element_part)*
    ;

// Virtual column expression definition
virtual_column_expression
    : autogenerated_sequence_definition
    | (GENERATED ALWAYS?)? AS '(' expression ')'
    ;

// Partition values lists
index_partitioning_values_list
    : literal (',' literal)*
    | TIMESTAMP literal (',' TIMESTAMP literal)*
    ;

range_values_list
    : literal (',' literal)*
    | TIMESTAMP literal (',' TIMESTAMP literal)*
    ;

// Expression list (with underscore suffix per upstream naming)
expressions_
    : expression (',' expression)*
    ;

// String delimiter helper
string_delimiter
    : CHAR_STRING
    | string_function
    | string_delimiter BAR BAR string_delimiter
    | '(' string_delimiter ')'
    | id_expression
    ;

// SQL*Plus clear command
clear_command
    : CLEAR (COLUMN? regular_id | ALL)
    ;

// SQL*Plus start command
start_command
    : START_CMD id_expression PERIOD (SQL | FILE_EXT)
    ;

// SQL*Plus command without semicolon
sql_plus_command_no_semicolon
    : set_command
    ;

// System actions for audit
system_actions
    : ACTIONS system_privilege (',' system_privilege)*
    ;

regular_id
    : non_reserved_keywords_pre12c
    | non_reserved_keywords_in_12c
    | REGULAR_ID
    | ABSENT
    | A_LETTER
    | BYTES
    | CHARSETFORM
    | CHARSETID
    | DURATION
    | ENABLED
    | ERROR_CODE
    | ERROR_INDEX
    | EXTEND
    | FIELD
    | ITEMS
    | LINES
    | MASK
    | MAXLEN
    | NEWLINE_
    | NOEXTEND
    | NOSCALE
    | NOSHARD
    | POLYMORPHIC
    | RECORDS
    | SHARD
    | AGENT
    | AGGREGATE
    | ANALYZE
    | AUTONOMOUS_TRANSACTION
    | BACKINGFILE
    | BATCH
    | BINARY_INTEGER
    | BOOLEAN
    | C_LETTER
    | CHAR
    | CLUSTER
    | CONSTRUCTOR
    | CUSTOMDATUM
    | CASESENSITIVE
    | DECIMAL
    | DELETE
    | DETERMINISTIC
    | DSINTERVAL_UNCONSTRAINED
    | E_LETTER
    | ERR
    | EXCEPTION
    | EXCEPTION_INIT
    | EXCEPTIONS
    | EXISTS
    | EXIT
    | FILESTORE
    | FLOAT
    | FORALL
    | G_LETTER
    | INDICES
    | INOUT
    | INTEGER
    | JSON_TRANSFORM
    | K_LETTER
    | LANGUAGE
    | LONG
    | LOOP
    | MOUNTPOINT
    | M_LETTER
    | MISSING
    | MISMATCH
    | NUMBER
    | ORADATA
    | OSERROR
    | OUT
    | OVERRIDING
    | P_LETTER
    | PARALLEL_ENABLE
    | PIPELINED
    | PLS_INTEGER
    | PMEM
    | POSITIVE
    | POSITIVEN
    | PRAGMA
    | PUBLIC
    | RAISE
    | RAW
    | RECORD
    | REF
    | RENAME
    | RESTRICT_REFERENCES
    | RESULT
    | SDO_GEOMETRY
    | SELF
    | SERIALLY_REUSABLE
    | SET
    | SHARDSPACE
    | SIGNTYPE
    | SIMPLE_INTEGER
    | SMALLINT
    | SQLDATA
    | SQLERROR
    | SUBTYPE
    | T_LETTER
    | TIMESTAMP_LTZ_UNCONSTRAINED
    | TIMESTAMP_TZ_UNCONSTRAINED
    | TIMESTAMP_UNCONSTRAINED
    | TRIGGER
    | VARCHAR
    | VARCHAR2
    | VARIABLE
    | WARNING
    | WHILE
    | XMLAGG
    | YMINTERVAL_UNCONSTRAINED
    | REGR_
    | VAR_
    | VALUE
    | COVAR_
    ;

non_reserved_keywords_in_12c
    : ACL
    | ACROSS
    | ACTION
    | ACTIONS
    | ACTIVE
    | ACTIVE_DATA
    | ACTIVITY
    | ADAPTIVE_PLAN
    | ADVANCED
    | AFD_DISKSTRING
    | ALTERNATE
    | ALGORITHM
    | ANALYTIC
    | ANCESTOR
    | ANOMALY
    | ANSI_REARCH
    | APPLICATION
    | APPROX_COUNT_DISTINCT
    | ARCHIVAL
    | ARCHIVED
    | ASIS
    | ASSIGN
    | AUTO_LOGIN
    | AUTO_REOPTIMIZE
    | AVRO
    | BACKGROUND
    | BACKUPS
    | BATCHSIZE
    | BATCH_TABLE_ACCESS_BY_ROWID
    | BEGINNING
    | BEQUEATH
    | BITMAP_AND
    | BLOCKCHAIN
    | BSON
    | CACHING
    | CALCULATED
    | CALLBACK
    | CAPACITY
    | CAPTION
    | CDBDEFAULT
    | CLASSIFICATION
    | CLASSIFIER
    | CLAUSE
    | CLEAN
    | CLEANUP
    | CLIENT
    | CLUSTERING
    | CLUSTER_DETAILS
    | CLUSTER_DISTANCE
    | COLLATE
    | COLLATION
    | COMMON
    | COMMON_DATA
    | COMPONENT
    | COMPONENTS
    | CONDITION
    | CONDITIONAL
    | CONTAINERS
    | CONTAINERS_DEFAULT
    | CONTAINER_DATA
    | CONTAINER_MAP
    | CONVERSION
    | CON_DBID_TO_ID
    | CON_GUID_TO_ID
    | CON_ID
    | CON_NAME_TO_ID
    | CON_UID_TO_ID
    | COOKIE
    | COPY
    | CREATE_FILE_DEST
    | CREDENTIAL
    | CRITICAL
    | CUBE_AJ
    | CUBE_SJ
    | DATAMOVEMENT
    | DATAOBJ_TO_MAT_PARTITION
    | DATAPUMP
    | DATA_SECURITY_REWRITE_LIMIT
    | DAYS
    | DB_UNIQUE_NAME
    | DECORRELATE
    | DEFAULT_CREDENTIAL
    | DEFAULT_COLLATION
    | DEFINE
    | DEFINITION
    | DELEGATE
    | DELETE_ALL
    | DESCRIPTION
    | DESTROY
    | DIMENSIONS
    | DISABLE_ALL
    | DISABLE_PARALLEL_DML
    | DISCARD
    | DISTRIBUTE
    | DUPLICATE
    | DUPLICATED
    | DV
    | EDITIONABLE
    | ELIM_GROUPBY
    | EM
    | ENABLE_ALL
    | ENABLE_PARALLEL_DML
    | EQUIPART
    | EVAL
    | EVALUATE
    | EXISTING
    | EXPRESS
    | EXTENDED
    | EXTRACTCLOBXML
    | FACTOR
    | FAILOVER
    | FAILURE
    | FAMILY
    | FAR
    | FASTSTART
    | FEATURE
    | FEATURE_DETAILS
    | FETCH
    | FILE_NAME_CONVERT
    | FILEGROUP
    | FIXED_VIEW_DATA
    | FLEX
    | FORMAT
    | FTP
    | GATHER_OPTIMIZER_STATISTICS
    | GET
    | HALF_YEARS
    | HASHING
    | HIER_ORDER
    | HIERARCHICAL
    | HOURS
    | HTTP
    | H_LETTER
    | IDLE
    | ILM
    | IMMUTABLE
    | INACTIVE
    | INACTIVE_ACCOUNT_TIME
    | INDEXING
    | INHERIT
    | INMEMORY
    | INMEMORY_PRUNING
    | INPLACE
    | INTERLEAVED
    | ISOLATE
    | IS_LEAF
    | JSON
    | JSONGET
    | JSONPARSE
    | JSON_ARRAY
    | JSON_ARRAYAGG
    | JSON_EQUAL
    | JSON_EXISTS
    | JSON_EXISTS2
    | JSON_OBJECT
    | JSON_OBJECTAGG
    | JSON_QUERY
    | JSON_SERIALIZE
    | JSON_TABLE
    | JSON_TEXTCONTAINS
    | JSON_TEXTCONTAINS2
    | JSON_VALUE
    | KEYSTORE
    | LABEL
    | LAX
    | LEAD_CDB
    | LEAD_CDB_URI
    | LEVEL_NAME
    | LIFECYCLE
    | LINEAR
    | LOCKDOWN
    | LOCKING
    | LOGMINING
    | LOST
    | MANDATORY
    | MAP
    | MATCH
    | MATCHES
    | MATCH_NUMBER
    | MATCH_RECOGNIZE
    | MAX_SHARED_TEMP_SIZE
    | MEMCOMPRESS
    | METADATA
    | MEMBER_CAPTION
    | MEMBER_DESCRIPTION
    | MEMBER_NAME
    | MEMBER_UNIQUE_NAME
    | MEMOPTIMIZE
    | MINUTES
    | MODEL_NB
    | MODEL_SV
    | MODIFICATION
    | MODULE
    | MONTHS
    | MULTIDIMENSIONAL
    | NEG
    | NOCOPY
    | NOKEEP
    | NONEDITIONABLE
    | NOPARTITION
    | NORELOCATE
    | NOREPLAY
    | NO_ADAPTIVE_PLAN
    | NO_ANSI_REARCH
    | NO_AUTO_REOPTIMIZE
    | NO_BATCH_TABLE_ACCESS_BY_ROWID
    | NO_CLUSTERING
    | NO_COMMON_DATA
    | NO_DATA_SECURITY_REWRITE
    | NO_DECORRELATE
    | NO_ELIM_GROUPBY
    | NO_GATHER_OPTIMIZER_STATISTICS
    | NO_INMEMORY
    | NO_INMEMORY_PRUNING
    | NO_OBJECT_LINK
    | NO_PARTIAL_JOIN
    | NO_PARTIAL_ROLLUP_PUSHDOWN
    | NO_PQ_CONCURRENT_UNION
    | NO_PQ_REPLICATE
    | NO_PQ_SKEW
    | NOPROMPT
    | NO_PX_FAULT_TOLERANCE
    | NO_ROOT_SW_FOR_LOCAL
    | NO_SQL_TRANSLATION
    | NO_USE_CUBE
    | NO_USE_VECTOR_AGGREGATION
    | NO_VECTOR_TRANSFORM
    | NO_VECTOR_TRANSFORM_DIMS
    | NO_VECTOR_TRANSFORM_FACT
    | NO_ZONEMAP
    | OBJ_ID
    | OFFSET
    | OLS
    | OMIT
    | ONE
    | ORACLE_DATAPUMP
    | ORACLE_HDFS
    | ORACLE_HIVE
    | ORACLE_LOADER
    | ORA_CHECK_ACL
    | ORA_CHECK_PRIVILEGE
    | ORA_CLUSTERING
    | ORA_INVOKING_USER
    | ORA_INVOKING_USERID
    | ORA_INVOKING_XS_USER
    | ORA_INVOKING_XS_USER_GUID
    | ORA_RAWCOMPARE
    | ORA_RAWCONCAT
    | ORA_WRITE_TIME
    | PARENT_LEVEL_NAME
    | PARENT_UNIQUE_NAME
    | PASSWORD_ROLLOVER_TIME
    | PARTIAL
    | PARTIAL_JOIN
    | PARTIAL_ROLLUP_PUSHDOWN
    | PAST
    | PATCH
    | PATH_PREFIX
    | PATTERN
    | PER
    | PERIOD
    | PERIOD_KEYWORD
    | PERMUTE
    | PLUGGABLE
    | POOL_16K
    | POOL_2K
    | POOL_32K
    | POOL_4K
    | POOL_8K
    | PQ_CONCURRENT_UNION
    | PQ_DISTRIBUTE_WINDOW
    | PQ_FILTER
    | PQ_REPLICATE
    | PQ_SKEW
    | PRELOAD
    | PRETTY
    | PREV
    | PRINTBLOBTOCLOB
    | PRIORITY
    | PRIVILEGED
    | PROPERTY
    | PROTOCOL
    | PROXY
    | PRUNING
    | PX_FAULT_TOLERANCE
    | QUARTERS
    | QUOTAGROUP
    | REALM
    | REDEFINE
    | RELOCATE
    | REMOTE
    | RESTART
    | ROLESET
    | ROWID_MAPPING_TABLE
    | RUNNING
    | SAVE
    | SCRUB
    | SDO_GEOM_MBR
    | SECONDS
    | SECRET
    | SERIAL
    | SERVICES
    | SERVICE_NAME_CONVERT
    | SHARDED
    | SHARING
    | SHELFLIFE
    | SITE
    | SOURCE_FILE_DIRECTORY
    | SOURCE_FILE_NAME_CONVERT
    | SQL_TRANSLATION_PROFILE
    | STANDARD
    | STANDARD_HASH
    | STANDBYS
    | STATE
    | STATEMENT
    | STREAM
    | SUBSCRIBE
    | SUBSET
    | SUCCESS
    | SYS
    | SYSBACKUP
    | SYSDG
    | SYSGUID
    | SYSKM
    | SYSOBJ
    | SYS_CHECK_PRIVILEGE
    | SYS_GET_COL_ACLIDS
    | SYS_MKXTI
    | SYS_OP_CYCLED_SEQ
    | SYS_OP_HASH
    | SYS_OP_KEY_VECTOR_CREATE
    | SYS_OP_KEY_VECTOR_FILTER
    | SYS_OP_KEY_VECTOR_FILTER_LIST
    | SYS_OP_KEY_VECTOR_SUCCEEDED
    | SYS_OP_KEY_VECTOR_USE
    | SYS_OP_PART_ID
    | SYS_OP_ZONE_ID
    | SYS_RAW_TO_XSID
    | SYS_XSID_TO_RAW
    | SYS_ZMAP_FILTER
    | SYS_ZMAP_REFRESH
    | TAG
    | TEXT
    | TIER
    | TIES
    | TO_ACLID
    | TRANSFORM
    | TRANSLATION
    | TRUST
    | UCS2
    | UNCONDITIONAL
    | UNITE
    | UNMATCHED
    | UNPLUG
    | UNSUBSCRIBE
    | USABLE
    | USER_DATA
    | USER_TABLESPACES
    | USE_CUBE
    | USE_HIDDEN_PARTITIONS
    | USE_VECTOR_AGGREGATION
    | USING_NO_EXPAND
    | UTF16BE
    | UTF16LE
    | UTF32
    | UTF8
    | V1
    | V2
    | VALIDATE_CONVERSION
    | VALID_TIME_END
    | VECTOR_TRANSFORM
    | VECTOR_TRANSFORM_DIMS
    | VECTOR_TRANSFORM_FACT
    | VERIFIER
    | VIOLATION
    | VISIBILITY
    | WEEK
    | WEEKS
    | WITH_PLSQL
    | WRAPPER
    | XS
    | YEARS
    | ZONEMAP
    | DEFAULTIF            
    | LLS
    | ENCLOSED
    | TERMINATED
    | OPTIONALLY
    | LRTRIM
    | NOTRIM                      
    | LDRTRIM                     
    | DATE_FORMAT        
    | MASK                        
    | TRANSFORMS                  
    | LOBFILE                     
    | STARTOF                     
    | CHARACTERSET                
    | RECORDS                     
    | FIXED                       
    | DELIMITED                   
    | XMLTAG                      
    | PREPROCESSOR                
    | TERRITORY                   
    | LITTLE                      
    | BIG                         
    | ENDIAN                      
    | BYTEORDERMARK               
    | NOCHECK                     
    | SIZES                       
    | ARE                         
    | BYTES                       
    | CHARACTERS                  
    | READSIZE                    
    | DISABLE_DIRECTORY_LINK_CHECK
    | DATE_CACHE                  
    | FIELD_NAMES                 
    | FILES                       
    | IO_OPTIONS                  
    | DIRECTIO                    
    | NODIRECTIO                  
    | DNFS_ENABLE                 
    | DNFS_DISABLE                
    | DNFS_READBUFFERS            
    | NOBADFILE                   
    | BADFILE                     
    | NODISCARDFILE               
    | DISCARDFILE                 
    | NOLOGFILE                   
    | FIELDS                      
    | IGNORE_CHARS_AFTER_EOR      
    | CSV                         
    | EMBEDDED                    
    | OVERRIDE                    
    | THESE                       
    | FIELD                       
    | NONULLIF                    
    | POSITION                    
    | NEWLINE_
    | DETECTED
    | UNSIGNED
    | ZONED
    | ORACLE_DATE
    | ORACLE_NUMBER
    | COUNTED
    | VARRAW
    | VARCHARC
    | VARRAWC
    ;

non_reserved_keywords_pre12c
    : ABORT
    | ABS
    | ACCESSED
    | ACCESS
    | ACCOUNT
    | ACOS
    | ACTIVATE
    | ACTIVE_COMPONENT
    | ACTIVE_FUNCTION
    | ACTIVE_TAG
    | ADD_COLUMN
    | ADD_GROUP
    | ADD_MONTHS
    | ADD
    | ADJ_DATE
    | ADMINISTER
    | ADMINISTRATOR
    | ADMIN
    | ADVISE
    | ADVISOR
    | AFTER
    | ALIAS
    | ALLOCATE
    | ALLOW
    | ALL_ROWS
    | ALWAYS
    | ANALYZE
    | ANCILLARY
    | AND_EQUAL
    | ANTIJOIN
    | ANYSCHEMA
    | APPENDCHILDXML
    | APPEND
    | APPEND_VALUES
    | APPLY
    | ARCHIVELOG
    | ARCHIVE
    | ARRAY
    | ASCII
    | ASCIISTR
    | ASIN
    | ASSEMBLY
    | ASSOCIATE
    | ASYNCHRONOUS
    | ASYNC
    | ATAN2
    | ATAN
    | AT
    | ATTRIBUTE
    | ATTRIBUTES
    | AUTHENTICATED
    | AUTHENTICATION
    | AUTHID
    | AUTHORIZATION
    | AUTOALLOCATE
    | AUTOEXTEND
    | AUTOMATIC
    | AUTO
    | AVAILABILITY
    | AVG
    | BACKUP
    | BASICFILE
    | BASIC
    | BATCH
    | BECOME
    | BEFORE
    | BEGIN
    | BEGIN_OUTLINE_DATA
    | BEHALF
    | BFILE
    | BFILENAME
    | BIGFILE
    | BINARY_DOUBLE_INFINITY
    | BINARY_DOUBLE
    | BINARY_DOUBLE_NAN
    | BINARY_FLOAT_INFINITY
    | BINARY_FLOAT
    | BINARY_FLOAT_NAN
    | BINARY
    | BIND_AWARE
    | BINDING
    | BIN_TO_NUM
    | BITAND
    | BITMAP
    | BITMAPS
    | BITMAP_TREE
    | BITS
    | BLOB
    | BLOCK
    | BLOCK_RANGE
    | BLOCKSIZE
    | BLOCKS
    | BODY
    | BOTH
    | BOUND
    | BRANCH
    | BREADTH
    | BROADCAST
    | BUFFER_CACHE
    | BUFFER
    | BUFFER_POOL
    | BUILD
    | BULK
    | BYPASS_RECURSIVE_CHECK
    | BYPASS_UJVC
    | BYTE
    | CACHE_CB
    | CACHE_INSTANCES
    | CACHE
    | CACHE_TEMP_TABLE
    | CALL
    | CANCEL
    | CARDINALITY
    | CASCADE
    | CASE
    | CAST
    | CATEGORY
    | CEIL
    | CELL_FLASH_CACHE
    | CERTIFICATE
    | CFILE
    | CHAINED
    | CHANGE_DUPKEY_ERROR_INDEX
    | CHANGE
    | CHARACTER
    | CHAR_CS
    | CHARTOROWID
    | CHECK_ACL_REWRITE
    | CHECKPOINT
    | CHILD
    | CHOOSE
    | CHR
    | CHUNK
    | CLASS
    | CLEAR
    | CLOB
    | CLONE
    | CLOSE_CACHED_OPEN_CURSORS
    | CLOSE
    | CLUSTER_BY_ROWID
    | CLUSTER_ID
    | CLUSTERING_FACTOR
    | CLUSTER_PROBABILITY
    | CLUSTER_SET
    | COALESCE
    | COALESCE_SQ
    | COARSE
    | CO_AUTH_IND
    | COLD
    | COLLECT
    | COLUMNAR
    | COLUMN_AUTH_INDICATOR
    | COLUMN
    | COLUMNS
    | COLUMN_STATS
    | COLUMN_VALUE
    | COMMENT
    | COMMIT
    | COMMITTED
    | COMPACT
    | COMPATIBILITY
    | COMPILE
    | COMPLETE
    | COMPLIANCE
    | COMPOSE
    | COMPOSITE_LIMIT
    | COMPOSITE
    | COMPOUND
    | COMPUTE
    | CONCAT
    | CONFIRM
    | CONFORMING
    | CONNECT_BY_CB_WHR_ONLY
    | CONNECT_BY_COMBINE_SW
    | CONNECT_BY_COST_BASED
    | CONNECT_BY_ELIM_DUPS
    | CONNECT_BY_FILTERING
    | CONNECT_BY_ISCYCLE
    | CONNECT_BY_ISLEAF
    | CONNECT_BY_ROOT
    | CONNECT_TIME
    | CONSIDER
    | CONSISTENT
    | CONSTANT
    | CONST
    | CONSTRAINT
    | CONSTRAINTS
    | CONTAINER
    | CONTENT
    | CONTENTS
    | CONTEXT
    | CONTINUE
    | CONTROLFILE
    | CONVERT
    | CORR_K
    | CORR
    | CORR_S
    | CORRUPTION
    | CORRUPT_XID_ALL
    | CORRUPT_XID
    | COSH
    | COS
    | COST
    | COST_XML_QUERY_REWRITE
    | COUNT
    | COVAR_POP
    | COVAR_SAMP
    | CPU_COSTING
    | CPU_PER_CALL
    | CPU_PER_SESSION
    | CRASH
    | CREATE_STORED_OUTLINES
    | CREATION
    | CROSSEDITION
    | CROSS
    | CSCONVERT
    | CUBE_GB
    | CUBE
    | CUME_DISTM
    | CUME_DIST
    | CURRENT_DATE
    | CURRENT
    | CURRENT_SCHEMA
    | CURRENT_TIME
    | CURRENT_TIMESTAMP
    | CURRENT_USER
    | CURRENTV
    | CURSOR
    | CURSOR_SHARING_EXACT
    | CURSOR_SPECIFIC_SEGMENT
    | CV
    | CYCLE
    | DANGLING
    | DATABASE
    | DATAFILE
    | DATAFILES
    | DATA
    | DATAOBJNO
    | DATAOBJ_TO_PARTITION
    | DATE_MODE
    | DAY
    | DBA
    | DBA_RECYCLEBIN
    | DBMS_STATS
    | DB_ROLE_CHANGE
    | DBTIMEZONE
    | DB_VERSION
    | DDL
    | DEALLOCATE
    | DEBUGGER
    | DEBUG
    // | DECLARE
    | DEC
    | DECOMPOSE
    | DECREMENT
    | DECR
    | DECRYPT
    | DEDUPLICATE
    | DEFAULTS
    | DEFERRABLE
    | DEFERRED
    | DEFINED
    | DEFINER
    | DEGREE
    | DELAY
    | DELETEXML
    | DEMAND
    | DENSE_RANKM
    | DENSE_RANK
    | DEPENDENT
    | DEPTH
    | DEQUEUE
    | DEREF
    | DEREF_NO_REWRITE
    | DETACHED
    | DETERMINES
    | DICTIONARY
    | DIMENSION
    | DIRECT_LOAD
    | DIRECTORY
    | DIRECT_PATH
    | DISABLE
    | DISABLE_PRESET
    | DISABLE_RPKE
    | DISALLOW
    | DISASSOCIATE
    | DISCONNECT
    | DISKGROUP
    | DISK
    | DISKS
    | DISMOUNT
    | DISTINGUISHED
    | DISTRIBUTED
    | DML
    | DML_UPDATE
    | DOCFIDELITY
    | DOCUMENT
    | DOMAIN_INDEX_FILTER
    | DOMAIN_INDEX_NO_SORT
    | DOMAIN_INDEX_SORT
    | DOUBLE
    | DOWNGRADE
    | DRIVING_SITE
    | DROP_COLUMN
    | DROP_GROUP
    | DST_UPGRADE_INSERT_CONV
    | DUMP
    | DYNAMIC
    | DYNAMIC_SAMPLING_EST_CDN
    | DYNAMIC_SAMPLING
    | EACH
    | EDITIONING
    | EDITION
    | EDITIONS
    | ELEMENT
    | ELIMINATE_JOIN
    | ELIMINATE_OBY
    | ELIMINATE_OUTER_JOIN
    | EMPTY_BLOB
    | EMPTY_CLOB
    | EMPTY
    | ENABLE
    | ENABLE_PRESET
    | ENCODING
    | ENCRYPTION
    | ENCRYPT
    | END_OUTLINE_DATA
    | ENFORCED
    | ENFORCE
    | ENQUEUE
    | ENTERPRISE
    | ENTITYESCAPING
    | ENTRY
    | ERROR_ARGUMENT
    | ERROR
    | ERROR_ON_OVERLAP_TIME
    | ERRORS
    | ESCAPE
    | ESTIMATE
    | EVALNAME
    | EVALUATION
    | EVENTS
    | EVERY
    | EXCEPTIONS
    | EXCEPT
    | EXCHANGE
    | EXCLUDE
    | EXCLUDING
    | EXECUTE
    | EXEMPT
    | EXISTSNODE
    | EXPAND_GSET_TO_UNION
    | EXPAND_TABLE
    | EXPIRE
    | EXPLAIN
    | EXPLOSION
    | EXP
    | EXPORT
    | EXPR_CORR_CHECK
    | EXTENDS
    | EXTENT
    | EXTENTS
    | EXTERNALLY
    | EXTERNAL
    | EXTRACT
    | EXTRACTVALUE
    | EXTRA
    | FACILITY
    | FACT
    | FACTORIZE_JOIN
    | FAILED_LOGIN_ATTEMPTS
    | FAILED
    | FAILGROUP
    | FALSE
    | FAST
    | FBTSCAN
    | FEATURE_ID
    | FEATURE_SET
    | FEATURE_VALUE
    | FILE
    | FILESYSTEM_LIKE_LOGGING
    | FILTER
    | FINAL
    | FINE
    | FINISH
    | FIRSTM
    | FIRST
    | FIRST_ROWS
    | FIRST_VALUE
    | FLAGGER
    | FLASHBACK
    | FLASH_CACHE
    | FLOB
    | FLOOR
    | FLUSH
    | FOLDER
    | FOLLOWING
    | FOLLOWS
    | FORCE
    | FORCE_XML_QUERY_REWRITE
    | FOREIGN
    | FOREVER
    | FORWARD
    | FRAGMENT_NUMBER
    | FREELIST
    | FREELISTS
    | FREEPOOLS
    | FRESH
    | FROM_TZ
    | FULL
    | FULL_OUTER_JOIN_TO_OUTER
    | FUNCTION
    | FUNCTIONS
    | GATHER_PLAN_STATISTICS
    | GBY_CONC_ROLLUP
    | GBY_PUSHDOWN
    | GENERATED
    | GLOBALLY
    | GLOBAL
    | GLOBAL_NAME
    | GLOBAL_TOPIC_ENABLED
    | GREATEST
    | GROUP_BY
    | GROUP_ID
    | GROUPING_ID
    | GROUPING
    | GROUPS
    | GUARANTEED
    | GUARANTEE
    | GUARD
    | HASH_AJ
    | HASHKEYS
    | HASH
    | HASH_SJ
    | HEADER
    | HEAP
    | HELP
    | HEXTORAW
    | HEXTOREF
    | HIDDEN_KEYWORD
    | HIDE
    | HIERARCHY
    | HIGH
    | HINTSET_BEGIN
    | HINTSET_END
    | HOT
    | HOUR
    | HWM_BROKERED
    | HYBRID
    | IDENTIFIER
    | IDENTITY
    | IDGENERATORS
    | IDLE_TIME
    | ID
    | IF
    | IGNORE
    | IGNORE_OPTIM_EMBEDDED_HINTS
    | IGNORE_ROW_ON_DUPKEY_INDEX
    | IGNORE_WHERE_CLAUSE
    | IMMEDIATE
    | IMPACT
    | IMPORT
    | INCLUDE
    | INCLUDE_VERSION
    | INCLUDING
    | INCREMENTAL
    | INCREMENT
    | INCR
    | INDENT
    | INDEX_ASC
    | INDEX_COMBINE
    | INDEX_DESC
    | INDEXED
    | INDEXES
    | INDEX_FFS
    | INDEX_FILTER
    | INDEX_JOIN
    | INDEX_ROWS
    | INDEX_RRS
    | INDEX_RS_ASC
    | INDEX_RS_DESC
    | INDEX_RS
    | INDEX_SCAN
    | INDEX_SKIP_SCAN
    | INDEX_SS_ASC
    | INDEX_SS_DESC
    | INDEX_SS
    | INDEX_STATS
    | INDEXTYPE
    | INDEXTYPES
    | INDICATOR
    | INFINITE
    | INFORMATIONAL
    | INITCAP
    | INITIALIZED
    | INITIALLY
    | INITIAL
    | INITRANS
    | INLINE
    | INLINE_XMLTYPE_NT
    | IN_MEMORY_METADATA
    | INNER
    | INSERTCHILDXMLAFTER
    | INSERTCHILDXMLBEFORE
    | INSERTCHILDXML
    | INSERTXMLAFTER
    | INSERTXMLBEFORE
    | INSTANCE
    | INSTANCES
    | INSTANTIABLE
    | INSTANTLY
    | INSTEAD
    | INSTR2
    | INSTR4
    | INSTRB
    | INSTRC
    | INSTR
    | INTERMEDIATE
    | INTERNAL_CONVERT
    | INTERNAL_USE
    | INTERPRETED
    | INTERVAL
    | INT
    | INVALIDATE
    | INVISIBLE
    | IN_XQUERY
    | ISOLATION_LEVEL
    | ISOLATION
    | ITERATE
    | ITERATION_NUMBER
    | JAVA
    | JOB
    | JOIN
    | KEEP_DUPLICATES
    | KEEP
    | KERBEROS
    | KEY_LENGTH
    | KEY
    | KEYSIZE
    | KEYS
    | KILL
    | LAG
    | LAST_DAY
    | LAST
    | LAST_VALUE
    | LATERAL
    | LAYER
    | LDAP_REGISTRATION_ENABLED
    | LDAP_REGISTRATION
    | LDAP_REG_SYNC_INTERVAL
    | LEADING
    | LEAD
    | LEAF
    | LEAST
    | LEFT
    | LENGTH2
    | LENGTH4
    | LENGTHB
    | LENGTHC
    | LENGTH
    | LESS
    | LEVEL
    | LEVELS
    | LIBRARY
    | LIFE
    | LIFETIME
    | LIKE2
    | LIKE4
    | LIKEC
    | LIKE_EXPAND
    | LIMIT
    | LINK
    | LISTAGG
    | LIST
    | LN
    | LNNVL
    | LOAD
    | LOB
    | LOBNVL
    | LOBS
    | LOCAL_INDEXES
    | LOCAL
    | LOCALTIME
    | LOCALTIMESTAMP
    | LOCATION
    | LOCATOR
    | LOCKED
    | LOGFILE
    | LOGFILES
    | LOGGING
    | LOGICAL
    | LOGICAL_READS_PER_CALL
    | LOGICAL_READS_PER_SESSION
    | LOG
    | LOGOFF
    | LOGON
    | LOG_READ_ONLY_VIOLATIONS
    | LOWER
    | LOW
    | LPAD
    | LTRIM
    | MAIN
    | MAKE_REF
    | MANAGED
    | MANAGEMENT
    | MANAGE
    | MANAGER
    | MANUAL
    | MAPPING
    | MASTER
    | MATCHED
    | MATERIALIZED
    | MATERIALIZE
    | MAXARCHLOGS
    | MAXDATAFILES
    | MAXEXTENTS
    | MAXIMIZE
    | MAXINSTANCES
    | MAXLOGFILES
    | MAXLOGHISTORY
    | MAXLOGMEMBERS
    | MAX
    | MAXSIZE
    | MAXTRANS
    | MAXVALUE
    | MEASURE
    | MEASURES
    | MEDIAN
    | MEDIUM
    | MEMBER
    | MEMOPTIMIZE
    | MEMORY
    | MERGEACTIONS
    | MERGE_AJ
    | MERGE_CONST_ON
    | MERGE
    | MERGE_SJ
    | METHOD
    | MIGRATE
    | MIGRATION
    | MINEXTENTS
    | MINIMIZE
    | MINIMUM
    | MINING
    | MIN
    | MINUS_NULL
    | MINUTE
    | MINVALUE
    | MIRRORCOLD
    | MIRRORHOT
    | MIRROR
    | MLSLABEL
    | MODEL_COMPILE_SUBQUERY
    | MODEL_DONTVERIFY_UNIQUENESS
    | MODEL_DYNAMIC_SUBQUERY
    | MODEL_MIN_ANALYSIS
    | MODEL
    | MODEL_NO_ANALYSIS
    | MODEL_PBY
    | MODEL_PUSH_REF
    | MODIFY_COLUMN_TYPE
    | MODIFY
    | MOD
    | MONITORING
    | MONITOR
    | MONTH
    | MONTHS_BETWEEN
    | MOUNT
    | MOUNTPATH
    | MOVEMENT
    | MOVE
    | MULTISET
    | MV_MERGE
    | NAMED
    | NAME
    | NAMESPACE
    | NAN
    | NANVL
    | NATIONAL
    | NATIVE_FULL_OUTER_JOIN
    | NATIVE
    | NATURAL
    | NAV
    | NCHAR_CS
    | NCHAR
    | NCHR
    | NCLOB
    | NEEDED
    | NESTED
    | NESTED_TABLE_FAST_INSERT
    | NESTED_TABLE_GET_REFS
    | NESTED_TABLE_ID
    | NESTED_TABLE_SET_REFS
    | NESTED_TABLE_SET_SETID
    | NETWORK
    | NEVER
    | NEW
    | NEW_TIME
    | NEXT_DAY
    | NEXT
    | NL_AJ
    | NLJ_BATCHING
    | NLJ_INDEX_FILTER
    | NLJ_INDEX_SCAN
    | NLJ_PREFETCH
    | NLS_CALENDAR
    | NLS_CHARACTERSET
    | NLS_CHARSET_DECL_LEN
    | NLS_CHARSET_ID
    | NLS_CHARSET_NAME
    | NLS_COMP
    | NLS_CURRENCY
    | NLS_DATE_FORMAT
    | NLS_DATE_LANGUAGE
    | NLS_INITCAP
    | NLS_ISO_CURRENCY
    | NL_SJ
    | NLS_LANG
    | NLS_LANGUAGE
    | NLS_LENGTH_SEMANTICS
    | NLS_LOWER
    | NLS_NCHAR_CONV_EXCP
    | NLS_NUMERIC_CHARACTERS
    | NLS_SORT
    | NLSSORT
    | NLS_SPECIAL_CHARS
    | NLS_TERRITORY
    | NLS_UPPER
    | NO_ACCESS
    | NOAPPEND
    | NOARCHIVELOG
    | NOAUDIT
    | NO_BASETABLE_MULTIMV_REWRITE
    | NO_BIND_AWARE
    | NO_BUFFER
    | NOCACHE
    | NO_CARTESIAN
    | NO_CHECK_ACL_REWRITE
    | NO_CLUSTER_BY_ROWID
    | NO_COALESCE_SQ
    | NO_CONNECT_BY_CB_WHR_ONLY
    | NO_CONNECT_BY_COMBINE_SW
    | NO_CONNECT_BY_COST_BASED
    | NO_CONNECT_BY_ELIM_DUPS
    | NO_CONNECT_BY_FILTERING
    | NO_COST_XML_QUERY_REWRITE
    | NO_CPU_COSTING
    | NOCPU_COSTING
    | NOCYCLE
    | NODELAY
    | NO_DOMAIN_INDEX_FILTER
    | NO_DST_UPGRADE_INSERT_CONV
    | NO_ELIMINATE_JOIN
    | NO_ELIMINATE_OBY
    | NO_ELIMINATE_OUTER_JOIN
    | NOENTITYESCAPING
    | NO_EXPAND_GSET_TO_UNION
    | NO_EXPAND
    | NO_EXPAND_TABLE
    | NO_FACT
    | NO_FACTORIZE_JOIN
    | NO_FILTERING
    | NOFORCE
    | NO_FULL_OUTER_JOIN_TO_OUTER
    | NO_GBY_PUSHDOWN
    | NOGUARANTEE
    | NO_INDEX_FFS
    | NO_INDEX
    | NO_INDEX_SS
    | NO_LOAD
    | NOLOCAL
    | NOLOGGING
    | NOMAPPING
    | NOMAXVALUE
    | NO_MERGE
    | NOMINIMIZE
    | NOMINVALUE
    | NO_MODEL_PUSH_REF
    | NO_MONITORING
    | NOMONITORING
    | NO_MONITOR
    | NO_MULTIMV_REWRITE
    | NO
    | NO_NATIVE_FULL_OUTER_JOIN
    | NONBLOCKING
    | NONE
    | NO_NLJ_BATCHING
    | NO_NLJ_PREFETCH
    | NONSCHEMA
    | NOORDER
    | NO_ORDER_ROLLUPS
    | NO_OUTER_JOIN_TO_ANTI
    | NO_OUTER_JOIN_TO_INNER
    | NOOVERRIDE
    | NO_PARALLEL_INDEX
    | NOPARALLEL_INDEX
    | NO_PARALLEL
    | NOPARALLEL
    | NO_PARTIAL_COMMIT
    | NO_PLACE_DISTINCT
    | NO_PLACE_GROUP_BY
    | NO_PQ_MAP
    | NO_PRUNE_GSETS
    | NO_PULL_PRED
    | NO_PUSH_PRED
    | NO_PUSH_SUBQ
    | NO_PX_JOIN_FILTER
    | NO_QKN_BUFF
    | NO_QUERY_TRANSFORMATION
    | NO_REF_CASCADE
    | NORELY
    | NOREPAIR
    | NORESETLOGS
    | NO_RESULT_CACHE
    | NOREVERSE
    | NO_REWRITE
    | NOREWRITE
    | NORMAL
    | NOROWDEPENDENCIES
    | NOSCHEMACHECK
    | NOSEGMENT
    | NO_SEMIJOIN
    | NO_SEMI_TO_INNER
    | NO_SET_TO_JOIN
    | NOSORT
    | NO_SQL_TUNE
    | NO_STAR_TRANSFORMATION
    | NO_STATEMENT_QUEUING
    | NO_STATS_GSETS
    | NOSTRICT
    | NO_SUBQUERY_PRUNING
    | NO_SUBSTRB_PAD
    | NO_SWAP_JOIN_INPUTS
    | NOSWITCH
    | NO_TABLE_LOOKUP_BY_NL
    | NO_TEMP_TABLE
    | NOTHING
    | NOTIFICATION
    | NO_TRANSFORM_DISTINCT_AGG
    | NO_UNNEST
    | NO_USE_HASH_AGGREGATION
    | NO_USE_HASH_GBY_FOR_PUSHDOWN
    | NO_USE_HASH
    | NO_USE_INVISIBLE_INDEXES
    | NO_USE_MERGE
    | NO_USE_NL
    | NOVALIDATE
    | NO_XDB_FASTPATH_INSERT
    | NO_XML_DML_REWRITE
    | NO_XMLINDEX_REWRITE_IN_SELECT
    | NO_XMLINDEX_REWRITE
    | NO_XML_QUERY_REWRITE
    | NTH_VALUE
    | NTILE
    | NULLIF
    | NULLS
    | NUMERIC
    | NUM_INDEX_KEYS
    | NUMTODSINTERVAL
    | NUMTOYMINTERVAL
    | NVARCHAR2
    | NVL2
    | NVL
    | OBJECT2XML
    | OBJECT
    | OBJNO
    | OBJNO_REUSE
    | OCCURENCES
    | OFFLINE
    | OFF
    | OIDINDEX
    | OID
    | OLAP
    | OLD
    | OLD_PUSH_PRED
    | OLTP
    | ONLINE
    | ONLY
    | OPAQUE
    | OPAQUE_TRANSFORM
    | OPAQUE_XCANONICAL
    | OPCODE
    | OPEN
    | OPERATIONS
    | OPERATOR
    | OPT_ESTIMATE
    | OPTIMAL
    | OPTIMIZE
    | OPTIMIZER_FEATURES_ENABLE
    | OPTIMIZER_GOAL
    | OPT_PARAM
    | ORA_BRANCH
    | ORADEBUG
    | ORA_DST_AFFECTED
    | ORA_DST_CONVERT
    | ORA_DST_ERROR
    | ORA_GET_ACLIDS
    | ORA_GET_PRIVILEGES
    | ORA_HASH
    | ORA_ROWSCN
    | ORA_ROWSCN_RAW
    | ORA_ROWVERSION
    | ORA_TABVERSION
    | ORDERED
    | ORDERED_PREDICATES
    | ORDINALITY
    | OR_EXPAND
    | ORGANIZATION
    | OR_PREDICATES
    | OTHER
    | OUTER_JOIN_TO_ANTI
    | OUTER_JOIN_TO_INNER
    | OUTER
    | OUTLINE_LEAF
    | OUTLINE
    | OUT_OF_LINE
    | OVERFLOW
    | OVERFLOW_NOMOVE
    | OVERLAPS
    | OVER
    | OWNER
    | OWNERSHIP
    | OWN
    | PACKAGE
    | PACKAGES
    | PARALLEL_INDEX
    | PARALLEL
    | PARAMETERS
    | PARAM
    | PARENT
    | PARITY
    | PARTIALLY
    | PARTITION_HASH
    | PARTITION_LIST
    | PARTITION
    | PARTITION_RANGE
    | PARTITIONS
    | PARTNUMINST
    | PASSING
    | PASSWORD_GRACE_TIME
    | PASSWORD_LIFE_TIME
    | PASSWORD_LOCK_TIME
    | PASSWORD
    | PASSWORD_REUSE_MAX
    | PASSWORD_REUSE_TIME
    | PASSWORD_VERIFY_FUNCTION
    | PATH
    | PATHS
    | PBL_HS_BEGIN
    | PBL_HS_END
    | PCTINCREASE
    | PCTTHRESHOLD
    | PCTUSED
    | PCTVERSION
    | PENDING
    | PERCENTILE_CONT
    | PERCENTILE_DISC
    | PERCENT_KEYWORD
    | PERCENT_RANKM
    | PERCENT_RANK
    | PERFORMANCE
    | PERMANENT
    | PERMISSION
    | PFILE
    | PHYSICAL
    | PIKEY
    | PIV_GB
    | PIVOT
    | PIV_SSF
    | PLACE_DISTINCT
    | PLACE_GROUP_BY
    | PLAN
    | PLSCOPE_SETTINGS
    | PLSQL_CCFLAGS
    | PLSQL_CODE_TYPE
    | PLSQL_DEBUG
    | PLSQL_OPTIMIZE_LEVEL
    | PLSQL_WARNINGS
    | POINT
    | POLICY
    | POST_TRANSACTION
    | POWERMULTISET_BY_CARDINALITY
    | POWERMULTISET
    | POWER
    | PQ_DISTRIBUTE
    | PQ_MAP
    | PQ_NOMAP
    | PREBUILT
    | PRECEDES
    | PRECEDING
    | PRECISION
    | PRECOMPUTE_SUBQUERY
    | PREDICATE_REORDERS
    | PREDICTION_BOUNDS
    | PREDICTION_COST
    | PREDICTION_DETAILS
    | PREDICTION
    | PREDICTION_PROBABILITY
    | PREDICTION_SET
    | PREPARE
    | PRESENT
    | PRESENTNNV
    | PRESENTV
    | PRESERVE
    | PRESERVE_OID
    | PREVIOUS
    | PRIMARY
    | PRIVATE
    | PRIVATE_SGA
    | PRIVILEGE
    | PRIVILEGES
    | PROCEDURAL
    | PROCEDURE
    | PROCESS
    | PROFILE
    | PROGRAM
    | PROJECT
    | PROPAGATE
    | PROTECTED
    | PROTECTION
    | PULL_PRED
    | PURGE
    | PUSH_PRED
    | PUSH_SUBQ
    | PX_GRANULE
    | PX_JOIN_FILTER
    | QB_NAME
    | QUERY_BLOCK
    | QUERY
    | QUEUE_CURR
    | QUEUE
    | QUEUE_ROWP
    | QUIESCE
    | QUORUM
    | QUOTA
    | RANDOM_LOCAL
    | RANDOM
    | RANGE
    | RANKM
    | RANK
    | RAPIDLY
    | RATIO_TO_REPORT
    | RAWTOHEX
    | RAWTONHEX
    | RBA
    | RBO_OUTLINE
    | RDBA
    | READ
    | READS
    | REAL
    | REBALANCE
    | REBUILD
    | RECORDS_PER_BLOCK
    | RECOVERABLE
    | RECOVER
    | RECOVERY
    | RECYCLEBIN
    | RECYCLE
    | REDACTION
    | REDO
    | REDUCED
    | REDUNDANCY
    | REF_CASCADE_CURSOR
    | REFERENCED
    | REFERENCE
    | REFERENCES
    | REFERENCING
    | REF
    | REFRESH
    | REFTOHEX
    | REGEXP_COUNT
    | REGEXP_INSTR
    | REGEXP_LIKE
    | REGEXP_REPLACE
    | REGEXP_SUBSTR
    | REGISTER
    | REGR_AVGX
    | REGR_AVGY
    | REGR_COUNT
    | REGR_INTERCEPT
    | REGR_R2
    | REGR_SLOPE
    | REGR_SXX
    | REGR_SXY
    | REGR_SYY
    | REGULAR
    | REJECT
    | REKEY
    | RELATIONAL
    | RELY
    | REMAINDER
    | REMOTE_MAPPED
    | REMOVE
    | REPAIR
    | REPEAT
    | REPLACE
    | REPLICATION
    | REQUIRED
    | RESETLOGS
    | RESET
    | RESIZE
    | RESOLVE
    | RESOLVER
    | RESPECT
    | RESTORE_AS_INTERVALS
    | RESTORE
    | RESTRICT_ALL_REF_CONS
    | RESTRICTED
    | RESTRICT
    | RESULT_CACHE
    | RESUMABLE
    | RESUME
    | RETENTION
    | RETRY_ON_ROW_CHANGE
    | RETURNING
    | RETURN
    | REUSE
    | REVERSE
    | REWRITE
    | REWRITE_OR_ERROR
    | RIGHT
    | ROLE
    | ROLES
    | ROLLBACK
    | ROLLING
    | ROLLUP
    | ROOT
    | ROUND
    | ROWDEPENDENCIES
    | ROWID
    | ROWIDTOCHAR
    | ROWIDTONCHAR
    | ROW_LENGTH
    | ROW
    | ROW_NUMBER
    | ROWNUM
    | ROWS
    | RPAD
    | RTRIM
    | RULE
    | RULES
    | SALT
    | SAMPLE
    | SAVE_AS_INTERVALS
    | SAVEPOINT
    | SB4
    | SCALE
    | SCALE_ROWS
    | SCAN_INSTANCES
    | SCAN
    | SCHEDULER
    | SCHEMACHECK
    | SCHEMA
    | SCN_ASCENDING
    | SCN
    | SCOPE
    | SD_ALL
    | SD_INHIBIT
    | SD_SHOW
    | SEARCH
    | SECOND
    | SECUREFILE_DBA
    | SECUREFILE
    | SECURITY
    | SEED
    | SEG_BLOCK
    | SEG_FILE
    | SEGMENT
    | SELECTIVITY
    | SEMIJOIN_DRIVER
    | SEMIJOIN
    | SEMI_TO_INNER
    | SEQUENCED
    | SEQUENCE
    | SEQUENTIAL
    | SERIALIZABLE
    | SERVERERROR
    | SERVICE
    | SESSION_CACHED_CURSORS
    | SESSION
    | SESSIONS_PER_USER
    | SESSIONTIMEZONE
    | SESSIONTZNAME
    | SETS
    | SETTINGS
    | SET_TO_JOIN
    | SEVERE
    | SHARED
    | SHARED_POOL
    | SHOW
    | SHRINK
    | SHUTDOWN
    | SIBLINGS
    | SID
    | SIGNAL_COMPONENT
    | SIGNAL_FUNCTION
    | SIGN
    | SIMPLE
    | SINGLE
    | SINGLETASK
    | SINH
    | SIN
    | SKIP_EXT_OPTIMIZER
    | SKIP_
    | SKIP_UNQ_UNUSABLE_IDX
    | SKIP_UNUSABLE_INDEXES
    | SMALLFILE
    | SNAPSHOT
    | SOME
    | SORT
    | SOUNDEX
    | SOURCE
    | SPACE_KEYWORD
    | SPECIFICATION
    | SPFILE
    | SPLIT
    | SPREADSHEET
    | SQLLDR
    | SQL
    | SQL_TRACE
    | SQL_MACRO
    | SQRT
    | STALE
    | STANDALONE
    | STANDBY_MAX_DATA_DELAY
    | STANDBY
    | STAR
    | STAR_TRANSFORMATION
    | STARTUP
    | STATEMENT_ID
    | STATEMENT_QUEUING
    | STATEMENTS
    | STATIC
    | STATISTICS
    | STATS_BINOMIAL_TEST
    | STATS_CROSSTAB
    | STATS_F_TEST
    | STATS_KS_TEST
    | STATS_MODE
    | STATS_MW_TEST
    | STATS_ONE_WAY_ANOVA
    | STATS_T_TEST_INDEP
    | STATS_T_TEST_INDEPU
    | STATS_T_TEST_ONE
    | STATS_T_TEST_PAIRED
    | STATS_WSR_TEST
    | STDDEV
    | STDDEV_POP
    | STDDEV_SAMP
    | STOP
    | STORAGE
    | STORE
    | STREAMS
    | STRICT
    | STRING
    | STRIPE_COLUMNS
    | STRIPE_WIDTH
    | STRIP
    | STRUCTURE
    | SUBMULTISET
    | SUBPARTITION
    | SUBPARTITION_REL
    | SUBPARTITIONS
    | SUBQUERIES
    | SUBQUERY_PRUNING
    | SUBSTITUTABLE
    | SUBSTR2
    | SUBSTR4
    | SUBSTRB
    | SUBSTRC
    | SUBSTR
    | SUCCESSFUL
    | SUMMARY
    | SUM
    | SUPPLEMENTAL
    | SUSPEND
    | SWAP_JOIN_INPUTS
    | SWITCH
    | SWITCHOVER
    | SYNCHRONOUS
    | SYNC
    | SYS
    | SYSASM
    | SYS_AUDIT
    | SYSAUX
    | SYS_CHECKACL
    | SYS_CONNECT_BY_PATH
    | SYS_CONTEXT
    | SYSDATE
    | SYSDBA
    | SYS_DBURIGEN
    | SYS_DL_CURSOR
    | SYS_DM_RXFORM_CHR
    | SYS_DM_RXFORM_NUM
    | SYS_DOM_COMPARE
    | SYS_DST_PRIM2SEC
    | SYS_DST_SEC2PRIM
    | SYS_ET_BFILE_TO_RAW
    | SYS_ET_BLOB_TO_IMAGE
    | SYS_ET_IMAGE_TO_BLOB
    | SYS_ET_RAW_TO_BFILE
    | SYS_EXTPDTXT
    | SYS_EXTRACT_UTC
    | SYS_FBT_INSDEL
    | SYS_FILTER_ACLS
    | SYS_FNMATCHES
    | SYS_FNREPLACE
    | SYS_GET_ACLIDS
    | SYS_GET_PRIVILEGES
    | SYS_GETTOKENID
    | SYS_GETXTIVAL
    | SYS_GUID
    | SYS_MAKEXML
    | SYS_MAKE_XMLNODEID
    | SYS_MKXMLATTR
    | SYS_OP_ADT2BIN
    | SYS_OP_ADTCONS
    | SYS_OP_ALSCRVAL
    | SYS_OP_ATG
    | SYS_OP_BIN2ADT
    | SYS_OP_BITVEC
    | SYS_OP_BL2R
    | SYS_OP_BLOOM_FILTER_LIST
    | SYS_OP_BLOOM_FILTER
    | SYS_OP_C2C
    | SYS_OP_CAST
    | SYS_OP_CEG
    | SYS_OP_CL2C
    | SYS_OP_COMBINED_HASH
    | SYS_OP_COMP
    | SYS_OP_CONVERT
    | SYS_OP_COUNTCHG
    | SYS_OP_CSCONV
    | SYS_OP_CSCONVTEST
    | SYS_OP_CSR
    | SYS_OP_CSX_PATCH
    | SYS_OP_DECOMP
    | SYS_OP_DESCEND
    | SYS_OP_DISTINCT
    | SYS_OP_DRA
    | SYS_OP_DUMP
    | SYS_OP_DV_CHECK
    | SYS_OP_ENFORCE_NOT_NULL
    | SYSOPER
    | SYS_OP_EXTRACT
    | SYS_OP_GROUPING
    | SYS_OP_GUID
    | SYS_OP_IIX
    | SYS_OP_ITR
    | SYS_OP_LBID
    | SYS_OP_LOBLOC2BLOB
    | SYS_OP_LOBLOC2CLOB
    | SYS_OP_LOBLOC2ID
    | SYS_OP_LOBLOC2NCLOB
    | SYS_OP_LOBLOC2TYP
    | SYS_OP_LSVI
    | SYS_OP_LVL
    | SYS_OP_MAKEOID
    | SYS_OP_MAP_NONNULL
    | SYS_OP_MSR
    | SYS_OP_NICOMBINE
    | SYS_OP_NIEXTRACT
    | SYS_OP_NII
    | SYS_OP_NIX
    | SYS_OP_NOEXPAND
    | SYS_OP_NTCIMG
    | SYS_OP_NUMTORAW
    | SYS_OP_OIDVALUE
    | SYS_OP_OPNSIZE
    | SYS_OP_PAR_1
    | SYS_OP_PARGID_1
    | SYS_OP_PARGID
    | SYS_OP_PAR
    | SYS_OP_PIVOT
    | SYS_OP_R2O
    | SYS_OP_RAWTONUM
    | SYS_OP_RDTM
    | SYS_OP_REF
    | SYS_OP_RMTD
    | SYS_OP_ROWIDTOOBJ
    | SYS_OP_RPB
    | SYS_OPTLOBPRBSC
    | SYS_OP_TOSETID
    | SYS_OP_TPR
    | SYS_OP_TRTB
    | SYS_OPTXICMP
    | SYS_OPTXQCASTASNQ
    | SYS_OP_UNDESCEND
    | SYS_OP_VECAND
    | SYS_OP_VECBIT
    | SYS_OP_VECOR
    | SYS_OP_VECXOR
    | SYS_OP_VERSION
    | SYS_OP_VREF
    | SYS_OP_VVD
    | SYS_OP_XMLCONS_FOR_CSX
    | SYS_OP_XPTHATG
    | SYS_OP_XPTHIDX
    | SYS_OP_XPTHOP
    | SYS_OP_XTXT2SQLT
    | SYS_ORDERKEY_DEPTH
    | SYS_ORDERKEY_MAXCHILD
    | SYS_ORDERKEY_PARENT
    | SYS_PARALLEL_TXN
    | SYS_PATHID_IS_ATTR
    | SYS_PATHID_IS_NMSPC
    | SYS_PATHID_LASTNAME
    | SYS_PATHID_LASTNMSPC
    | SYS_PATH_REVERSE
    | SYS_PXQEXTRACT
    | SYS_RID_ORDER
    | SYS_ROW_DELTA
    | SYS_SC_2_XMLT
    | SYS_SYNRCIREDO
    | SYSTEM_DEFINED
    | SYSTEM
    | SYSTIMESTAMP
    | SYS_TYPEID
    | SYS_UMAKEXML
    | SYS_XMLANALYZE
    | SYS_XMLCONTAINS
    | SYS_XMLCONV
    | SYS_XMLEXNSURI
    | SYS_XMLGEN
    | SYS_XMLI_LOC_ISNODE
    | SYS_XMLI_LOC_ISTEXT
    | SYS_XMLINSTR
    | SYS_XMLLOCATOR_GETSVAL
    | SYS_XMLNODEID_GETCID
    | SYS_XMLNODEID_GETLOCATOR
    | SYS_XMLNODEID_GETOKEY
    | SYS_XMLNODEID_GETPATHID
    | SYS_XMLNODEID_GETPTRID
    | SYS_XMLNODEID_GETRID
    | SYS_XMLNODEID_GETSVAL
    | SYS_XMLNODEID_GETTID
    | SYS_XMLNODEID
    | SYS_XMLT_2_SC
    | SYS_XMLTRANSLATE
    | SYS_XMLTYPE2SQL
    | SYS_XQ_ASQLCNV
    | SYS_XQ_ATOMCNVCHK
    | SYS_XQBASEURI
    | SYS_XQCASTABLEERRH
    | SYS_XQCODEP2STR
    | SYS_XQCODEPEQ
    | SYS_XQCON2SEQ
    | SYS_XQCONCAT
    | SYS_XQDELETE
    | SYS_XQDFLTCOLATION
    | SYS_XQDOC
    | SYS_XQDOCURI
    | SYS_XQDURDIV
    | SYS_XQED4URI
    | SYS_XQENDSWITH
    | SYS_XQERRH
    | SYS_XQERR
    | SYS_XQESHTMLURI
    | SYS_XQEXLOBVAL
    | SYS_XQEXSTWRP
    | SYS_XQEXTRACT
    | SYS_XQEXTRREF
    | SYS_XQEXVAL
    | SYS_XQFB2STR
    | SYS_XQFNBOOL
    | SYS_XQFNCMP
    | SYS_XQFNDATIM
    | SYS_XQFNLNAME
    | SYS_XQFNNM
    | SYS_XQFNNSURI
    | SYS_XQFNPREDTRUTH
    | SYS_XQFNQNM
    | SYS_XQFNROOT
    | SYS_XQFORMATNUM
    | SYS_XQFTCONTAIN
    | SYS_XQFUNCR
    | SYS_XQGETCONTENT
    | SYS_XQINDXOF
    | SYS_XQINSERT
    | SYS_XQINSPFX
    | SYS_XQIRI2URI
    | SYS_XQLANG
    | SYS_XQLLNMFRMQNM
    | SYS_XQMKNODEREF
    | SYS_XQNILLED
    | SYS_XQNODENAME
    | SYS_XQNORMSPACE
    | SYS_XQNORMUCODE
    | SYS_XQ_NRNG
    | SYS_XQNSP4PFX
    | SYS_XQNSPFRMQNM
    | SYS_XQPFXFRMQNM
    | SYS_XQ_PKSQL2XML
    | SYS_XQPOLYABS
    | SYS_XQPOLYADD
    | SYS_XQPOLYCEL
    | SYS_XQPOLYCSTBL
    | SYS_XQPOLYCST
    | SYS_XQPOLYDIV
    | SYS_XQPOLYFLR
    | SYS_XQPOLYMOD
    | SYS_XQPOLYMUL
    | SYS_XQPOLYRND
    | SYS_XQPOLYSQRT
    | SYS_XQPOLYSUB
    | SYS_XQPOLYUMUS
    | SYS_XQPOLYUPLS
    | SYS_XQPOLYVEQ
    | SYS_XQPOLYVGE
    | SYS_XQPOLYVGT
    | SYS_XQPOLYVLE
    | SYS_XQPOLYVLT
    | SYS_XQPOLYVNE
    | SYS_XQREF2VAL
    | SYS_XQRENAME
    | SYS_XQREPLACE
    | SYS_XQRESVURI
    | SYS_XQRNDHALF2EVN
    | SYS_XQRSLVQNM
    | SYS_XQRYENVPGET
    | SYS_XQRYVARGET
    | SYS_XQRYWRP
    | SYS_XQSEQ2CON4XC
    | SYS_XQSEQ2CON
    | SYS_XQSEQDEEPEQ
    | SYS_XQSEQINSB
    | SYS_XQSEQRM
    | SYS_XQSEQRVS
    | SYS_XQSEQSUB
    | SYS_XQSEQTYPMATCH
    | SYS_XQSTARTSWITH
    | SYS_XQSTATBURI
    | SYS_XQSTR2CODEP
    | SYS_XQSTRJOIN
    | SYS_XQSUBSTRAFT
    | SYS_XQSUBSTRBEF
    | SYS_XQTOKENIZE
    | SYS_XQTREATAS
    | SYS_XQ_UPKXML2SQL
    | SYS_XQXFORM
    | TABLE
    | TABLE_LOOKUP_BY_NL
    | TABLES
    | TABLESPACE
    | TABLESPACE_NO
    | TABLE_STATS
    | TABNO
    | TANH
    | TAN
    | TBLORIDXPARTNUM
    | TEMPFILE
    | TEMPLATE
    | TEMPORARY
    | TEMP_TABLE
    | TEST
    | THAN
    | THE
    | THEN
    | THREAD
    | THROUGH
    | TIME
    | TIMING
    | TIMEOUT
    | TIMES
    | TIMESTAMP
    | TIMEZONE_ABBR
    | TIMEZONE_HOUR
    | TIMEZONE_MINUTE
    | TIME_ZONE
    | TIMEZONE_OFFSET
    | TIMEZONE_REGION
    | TIV_GB
    | TIV_SSF
    | TO_BINARY_DOUBLE
    | TO_BINARY_FLOAT
    | TO_BLOB
    | TO_CHAR
    | TO_CLOB
    | TO_DATE
    | TO_DSINTERVAL
    | TO_LOB
    | TO_MULTI_BYTE
    | TO_NCHAR
    | TO_NCLOB
    | TO_NUMBER
    | TOPLEVEL
    | TO_SINGLE_BYTE
    | TO_TIME
    | TO_TIMESTAMP
    | TO_TIMESTAMP_TZ
    | TO_TIME_TZ
    | TO_YMINTERVAL
    | TRACE
    | TRACING
    | TRACKING
    | TRAILING
    | TRANSACTION
    | TRANSFORM_DISTINCT_AGG
    | TRANSITIONAL
    | TRANSITION
    | TRANSLATE
    | TREAT
    | TRIGGERS
    | TRIM
    | TRUE
    | TRUNCATE
    | TRUNC
    | TRUSTED
    | TUNING
    | TX
    | TYPE
    | TYPES
    | TZ_OFFSET
    | UB2
    | UBA
    | UID
    | UNARCHIVED
    | UNBOUNDED
    | UNBOUND
    | UNDER
    | UNDO
    | UNDROP
    | UNIFORM
    | UNISTR
    | UNLIMITED
    | UNLOAD
    | UNLOCK
    | UNNEST_INNERJ_DISTINCT_VIEW
    | UNNEST
    | UNNEST_NOSEMIJ_NODISTINCTVIEW
    | UNNEST_SEMIJ_VIEW
    | UNPACKED
    | UNPIVOT
    | UNPROTECTED
    | UNQUIESCE
    | UNRECOVERABLE
    | UNRESTRICTED
    | UNTIL
    | UNUSABLE
    | UNUSED
    | UPDATABLE
    | UPDATED
    | UPDATEXML
    | UPD_INDEXES
    | UPD_JOININDEX
    | UPGRADE
    | UPPER
    | UPSERT
    | UROWID
    | USAGE
    | USE_ANTI
    | USE_CONCAT
    | USE_HASH_AGGREGATION
    | USE_HASH_GBY_FOR_PUSHDOWN
    | USE_HASH
    | USE_INVISIBLE_INDEXES
    | USE_MERGE_CARTESIAN
    | USE_MERGE
    | USE
    | USE_NL
    | USE_NL_WITH_INDEX
    | USE_PRIVATE_OUTLINES
    | USER_DEFINED
    | USERENV
    | USERGROUP
    | USER
    | USER_RECYCLEBIN
    | USERS
    | USE_SEMI
    | USE_STORED_OUTLINES
    | USE_TTT_FOR_GSETS
    | USE_WEAK_NAME_RESL
    | USING
    | VALIDATE
    | VALIDATION
    | VALUE
    | VARIANCE
    | VAR_POP
    | VARRAY
    | VARRAYS
    | VAR_SAMP
    | VARYING
    | VECTOR_READ
    | VECTOR_READ_TRACE
    | VERIFY
    | VERSIONING
    | VERSION
    | VERSIONS_ENDSCN
    | VERSIONS_ENDTIME
    | VERSIONS
    | VERSIONS_OPERATION
    | VERSIONS_STARTSCN
    | VERSIONS_STARTTIME
    | VERSIONS_XID
    | VIRTUAL
    | VISIBLE
    | VOLUME
    | VSIZE
    | WAIT
    | WALLET
    | WELLFORMED
    | WHENEVER
    | WHEN
    | WHITESPACE
    | WIDTH_BUCKET
    | WITHIN
    | WITHOUT
    | WORK
    | WRAPPED
    | WRITE
    | XDB_FASTPATH_INSERT
    | X_DYN_PRUNE
    | XID
    | XML2OBJECT
    | XMLATTRIBUTES
    | XMLCAST
    | XMLCDATA
    | XMLCOLATTVAL
    | XMLCOMMENT
    | XMLCONCAT
    | XMLDIFF
    | XML_DML_RWT_STMT
    | XMLELEMENT
    | XMLEXISTS2
    | XMLEXISTS
    | XMLFOREST
    | XMLINDEX_REWRITE_IN_SELECT
    | XMLINDEX_REWRITE
    | XMLINDEX_SEL_IDX_TBL
    | XMLISNODE
    | XMLISVALID
    | XML
    | XMLNAMESPACES
    | XMLPARSE
    | XMLPATCH
    | XMLPI
    | XMLQUERY
    | XMLQUERYVAL
    | XMLROOT
    | XMLSCHEMA
    | XMLSERIALIZE
    | XMLTABLE
    | XMLTRANSFORMBLOB
    | XMLTRANSFORM
    | XMLTYPE
    | XPATHTABLE
    | XS_SYS_CONTEXT
    | YEAR
    | YES
    | ZONE
    ;