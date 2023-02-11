hive (ds_hiveclass_test)> select * from dataexchange_test ;
a	b	1
a	b	2
a	b	3
c	d	4
c	d	5
c	d	6
use ds_hiveclass_test;
select str1, str2, concat_ws("-", collect_list(cast(id as string))) from dataexchange_test group by str1, str2;



select str1, str2, group_concat(",", id)  from dataexchange_test group by str1, str2;

select str1, str2, id2 from dataexchange_test  lateral view explode(split(id, ',')) newTable as id2;


--累加级联求和
select userid, visitdate, visitcount,
       sum(visitcount) over(partition by userid order by visitdate) as total
from dws_user_action;


select *