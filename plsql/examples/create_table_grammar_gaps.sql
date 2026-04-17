-- Regression fixtures for CREATE TABLE grammar gaps closed against Oracle 19c docs.
-- Each block below previously failed to parse; see PR description for details.

--------------------------------------------------------------------------------
-- Tier 1: partition value clauses accept full expressions (fixes BYT-9302)
--------------------------------------------------------------------------------

-- DATE typed literal in RANGE bound with INTERVAL partitioning
CREATE TABLE gcp.lead_drop_mc_native_data (
    txn_date  DATE,
    userid    VARCHAR2(100),
    custid    VARCHAR2(100),
    screenid  VARCHAR2(500),
    eventtime DATE,
    status    NUMBER
)
PARTITION BY RANGE (txn_date)
INTERVAL (NUMTODSINTERVAL(1,'DAY'))
(PARTITION p0 VALUES LESS THAN (DATE '2026-01-01'));

-- TIMESTAMP typed literal in RANGE bound
CREATE TABLE range_ts (d TIMESTAMP)
PARTITION BY RANGE (d)
(PARTITION p0 VALUES LESS THAN (TIMESTAMP '2026-01-01 00:00:00'));

-- Arithmetic expression in RANGE bound
CREATE TABLE range_expr (c NUMBER)
PARTITION BY RANGE (c)
(PARTITION p0 VALUES LESS THAN (10*10 + 5));

-- NULL in LIST values
CREATE TABLE list_null (c VARCHAR2(10))
PARTITION BY LIST (c)
(PARTITION p0 VALUES (NULL, 'a', 'b'));

-- Expression in LIST values
CREATE TABLE list_expr (c NUMBER)
PARTITION BY LIST (c)
(PARTITION p0 VALUES (1+2, 3*3));

-- Multi-column LIST partitioning with tuple form
CREATE TABLE list_multi_col (a NUMBER, b NUMBER)
PARTITION BY LIST (a, b)
(
    PARTITION p0 VALUES ((1,2),(3,4)),
    PARTITION p1 VALUES ((5,6))
);

--------------------------------------------------------------------------------
-- Tier 2: IF NOT EXISTS, LIST AUTOMATIC, blockchain flexibility
--------------------------------------------------------------------------------

-- IF NOT EXISTS on CREATE TABLE
CREATE TABLE IF NOT EXISTS t_ine (a NUMBER, b VARCHAR2(10));

-- AUTOMATIC list partitioning
CREATE TABLE list_auto (c NUMBER)
PARTITION BY LIST (c) AUTOMATIC
(PARTITION p0 VALUES (1));

-- Blockchain: new DAYS IDLE form on retention clause, with single-quoted algorithm/version
CREATE BLOCKCHAIN TABLE bc_idle (a NUMBER)
    NO DROP UNTIL 0 DAYS IDLE
    NO DELETE UNTIL 16 DAYS IDLE
    HASHING USING 'SHA2_512' VERSION 'v1';

-- Blockchain: drop/retention clauses both omitted (only the hash clause is required)
CREATE BLOCKCHAIN TABLE bc_hash_only (a NUMBER)
    HASHING USING "SHA2_512" VERSION "v1";

--------------------------------------------------------------------------------
-- Tier 3: CONSISTENT HASH and PARTITIONSET
--------------------------------------------------------------------------------

-- Consistent hash partitioning
CREATE TABLE ch_simple (a NUMBER)
PARTITION BY CONSISTENT HASH (a)
PARTITIONS AUTO;

-- Consistent hash with range sub-partitioning
CREATE TABLE ch_sub (a NUMBER, d DATE)
PARTITION BY CONSISTENT HASH (a)
PARTITIONS AUTO
SUBPARTITION BY RANGE (d)
SUBPARTITIONS AUTO;

-- Range partitionset
CREATE TABLE range_ps (a NUMBER, b NUMBER)
PARTITION BY PARTITIONSET RANGE (a)
(
    PARTITIONSET ps1 VALUES LESS THAN (100)
    (PARTITION p1 VALUES LESS THAN (10))
);

-- List partitionset
CREATE TABLE list_ps (a NUMBER, b NUMBER)
PARTITION BY PARTITIONSET LIST (a)
(
    PARTITIONSET ps1 VALUES (1,2,3)
    (PARTITION p1 VALUES (1))
);

--------------------------------------------------------------------------------
-- Tier 4: subpartition read_only / indexing clauses + column annotations
--------------------------------------------------------------------------------

-- read_only_clause on range_subpartition_desc (inside composite range partitioning)
CREATE TABLE comp_range_ro (a NUMBER, d DATE)
PARTITION BY RANGE (d)
SUBPARTITION BY RANGE (a)
(PARTITION p0 VALUES LESS THAN (DATE '2026-01-01')
  (SUBPARTITION sp0 VALUES LESS THAN (100) READ ONLY,
   SUBPARTITION sp1 VALUES LESS THAN (MAXVALUE) READ WRITE INDEXING OFF));

-- indexing_clause on individual_hash_subparts
CREATE TABLE comp_hash_idx (a NUMBER, b NUMBER)
PARTITION BY RANGE (a)
SUBPARTITION BY HASH (b)
(PARTITION p0 VALUES LESS THAN (100)
  (SUBPARTITION sp0 INDEXING ON,
   SUBPARTITION sp1 INDEXING OFF));

-- annotations_clause on column definition (Oracle 23c)
CREATE TABLE annotated_cols (
    id   NUMBER ANNOTATIONS (Display 'Identifier'),
    name VARCHAR2(100) ANNOTATIONS (ADD Display 'Name', MaxLength '100')
);
