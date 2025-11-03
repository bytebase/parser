;
#begin
(select 1 as a, 10 as a2) except distinct (select 1 as a, 10 as a2);
#end

#begin
(select 1 as a, 10 as a2) intersect (select 1 as a, 10 as a2);
#end

#begin
select * from t1 intersect select * from t2;
#end

#begin
select 1 union select 2;
#end

#begin
select 1 as a1, 10 as a2 union all select 2, 20 union distinct select 3, 30 union distinct select 2, 20 union all select 3, 30;

#end

#begin
(select 1 as a1, 10 as a2) union all (select 2, 20);
#end

#begin
(select 1 as a1, 10 as a2) union all (select 2, 20) union distinct (select 3, 30);
#end

#begin
select 1 as a1, 10 as a2 union all select 2, 20 union distinct (select 3, 30);
#end

#begin
select 1 as a1, 10 as a2 union all (select 2, 20) union distinct (select 3, 30);
#end

#begin
select 1 as a1, 10 as a2 union all (select 2, 20) union distinct select 3, 30;
#end

#begin
select 1 as a1, 10 as a2 union all (select 2, 20) union distinct (select 3, 30) union distinct select 2, 20 union all select 3, 30;
#end

#begin
select 1 as a1, 10 as a2 union all (select 2, 20) union distinct select 3, 30 union distinct select 2, 20 union all select 3, 30;
select 1 as a1, 10 as a2 union all (select 2, 20) union distinct select 3, 30 union distinct (select 2, 20) union all select 3, 30;
#end

#begin
((select 1 as a1, 10 as a2)) union all (((select 2, 20))) union distinct (select 3, 30);
#end

#begin
((select 1 as a1, 10 as a2)) union all (((select 2, 20))) union distinct (select 3, 30 into outfile 'test.dump');
#end

#begin
select 1 as a1, 10 as a2 union all (select 2, 20) union distinct (select 3, 30) union distinct select 2, 20 union all select 3, 30 into outfile 'test.dump';
#end

#begin
select 1 as a1, 10 as a2 union all (select 2, 20) union distinct select 3, 30 order by 1;
select 1 as a1, 10 as a2 union all (select 2, 20 order by 2) union distinct select 3, 30 order by 1;
select 1 as a1, 10 as a2 union all (select 2, 20 order by 2) union distinct (select 3, 30 order by 1);
select 1 as a1, 10 as a2 union all (select 2, 20 order by 2) union distinct (select 3, 30 order by 1) order by 2;
#end
;