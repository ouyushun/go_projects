--1.查看用户量
select 
	sex
	,count(userid) as user_cnt
	,count(distinct userid) as user_cnt1
from ods_commerce.ods_user_userinfo
group by sex
;

--2.查看用户行为表的分区情况
show partitions ods_commerce.ods_userbehavior

--3.查看用户行为表的枚举值分布
select type,count(*) as pv,count(distinct user_id) as uv
from ods_commerce.ods_userbehavior
group by type
;