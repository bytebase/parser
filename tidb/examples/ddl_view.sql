;
#begin
-- Create view
create or replace view my_view1 as select 1 union select 2 limit 0,5;
create algorithm = merge view my_view2(col1, col2) as select * from t2 with cascaded check option;
create or replace definer = 'ivan'@'%' view my_view3 as select count(*) from t3;
create or replace definer = current_user sql security invoker view my_view4(c1, 1c, _, c1_2)
	as select * from  (t1 as tt1, t2 as tt2) inner join t1 on t1.col1 = tt1.col1;
create view v_some_table as (with a as (select * from some_table) select * from a);

#end

#begin
-- Drop view
DROP VIEW IF EXISTS my_view1;
