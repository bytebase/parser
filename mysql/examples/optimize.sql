;
#begin
OPTIMIZE TABLE t1;
OPTIMIZE TABLE t1, t2;
#OPTIMIZE TABLES t1;
#OPTIMIZE TABLES t1, t2; # MySQL docs say that TABLES is valid.
optimize local table t1;
#end

;