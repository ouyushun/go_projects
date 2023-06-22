-------------------------------------1.dwd层数据--------------------------------------------------------------------
-- part1:用户信息表的处理
USE dwd_commerce;
DROP TABLE IF EXISTS dwd_user_userinfo;
CREATE TABLE IF NOT EXISTS dwd_user_userinfo (
   user_id STRING COMMENT '用户ID'
   ,user_name STRING COMMENT '用户名称'
   ,sex INT COMMENT '性别'
   ,user_money decimal(20,4) COMMENT '钱包'
   ,frozen_money decimal(20,4) COMMENT '近一个月的花费的总的金额'
   ,address_id STRING COMMENT '用户地址ID，0表示没有获取地址'
   ,reg_time STRING COMMENT '注册时间'
   ,last_login STRING COMMENT '最后登录时间'
) COMMENT '用户表' STORED AS ORC;

--ROW FORMAT DELIMITED FIELDS TERMINATED BY '#' STORED AS textfile;


insert overwrite table dwd_commerce.dwd_user_userinfo
select 
   userid   as user_id
  ,username as user_name
  ,sex
  ,cast(usermoney as decimal(20,4))   as user_money
  ,cast(frozenmoney as decimal(20,4)) as frozen_money
  ,addressid as address_id
  ,from_unixtime(cast(regtime as bigint),'yyyy-MM-dd HH:mm:ss')   as reg_time
  ,from_unixtime(cast(lastlogin as bigint),'yyyy-MM-dd HH:mm:ss') as last_login
from ods_commerce.ods_user_userinfo
;

select * from dwd_commerce.dwd_user_userinfo limit 100;

-- 知识点的回顾与应用
-- 1.日期函数：参考：https://blog.csdn.net/wangwangstone/article/details/110011860
-- 2.cast函数的用法：注意from_unixtime(cast(regtime as bigint),'yyyy-MM-dd HH:mm:ss')

-- part2:用户行为的处理
DROP TABLE IF EXISTS dwd_commerce.oys_dwd_user_behavior_detail_d;
CREATE TABLE IF NOT EXISTS dwd_commerce.oys_dwd_user_behavior_detail_d (
  user_id STRING COMMENT '用户标识'
  ,item_id STRING COMMENT '商品标识'
  ,category_id STRING COMMENT '商品分类标识'
  ,type STRING COMMENT '用户对商品的行为类型,浏览、收藏、加购物车、购买，对应取值分别是1、2、3、4'
  ,action_time STRING COMMENT '行为时间'
) 
COMMENT '用户行为明细表' 
PARTITIONED BY (dt string comment '日期')
STORED AS orc;



set hive.exec.dynamic.partition.mode=nonstrict;
insert overwrite table dwd_commerce.oys_dwd_user_behavior_detail_d partition(dt)
select 
   user_id
  ,item_id
  ,category_id
  ,type
  ,from_unixtime(cast(action_time as bigint),'yyyy-MM-dd HH:mm:ss')   as action_time
  ,to_date(from_unixtime(cast(action_time as bigint),'yyyy-MM-dd HH:mm:ss')) as dt
from ods_commerce.ods_userbehavior
;

-- 重点处理2017-11-25~2017-12-05的数据
select * from dwd_commerce.dwd_user_behavior_detail_d limit 100;

-- 知识点的回顾与应用
-- 1.to_date函数
-- 2.动态分区
-- 3.会话参数：set hive.exec.dynamic.partition.mode=nonstrict; -- 影响到分区写入
-- 4.查看、删除分区：
show partitions dwd_commerce.dwd_user_behavior_detail_d;
alter table dwd_commerce.dwd_user_behavior_detail_d drop if exists partition (dt<'2016-01-01');
alter table dwd_commerce.dwd_user_behavior_detail_d drop if exists partition (dt>='2022-05-20');

-------------------------------------2.dws层数据--------------------------------------------------------------------
--part1:dws层的数据建设 用户行为的处理
--1.用户粒度的汇总模型
DROP TABLE IF EXISTS dws_commerce.dws_user_behavior_sum_d;
CREATE TABLE IF NOT EXISTS dws_commerce.dws_user_behavior_sum_d (
  user_id STRING COMMENT '用户标识'
  ,view_cnt BIGINT COMMENT '浏览次数'
  ,fav_cnt BIGINT COMMENT '收藏次数'
  ,cart_cnt BIGINT COMMENT '加购物车次数'
  ,buy_cnt BIGINT COMMENT '购买次数'
) 
COMMENT '用户行为汇总表(用户粒度)' 
PARTITIONED BY (dt string comment '日期')
ROW FORMAT DELIMITED FIELDS TERMINATED BY ',' 
STORED AS orc;


insert overwrite table dws_commerce.dws_user_behavior_sum_d partition(dt)
select 
 user_id
,sum(if(type='pv',1,0))  as view_cnt
,sum(if(type='fav',1,0)) as fav_cnt
,sum(if(type='cart',1,0)) as cart_cnt
,sum(if(type='buy',1,0)) as buy_cnt
,dt
from dwd_commerce.dwd_user_behavior_detail_d
where dt between '2017-11-01' and '2017-12-03'
group by 
dt
,user_id
;
-- 数据验证：
select * from dws_commerce.dws_user_behavior_sum_d where dt='2017-11-25' limit 10;

--part2:dws层的数据建设 用户行为指标统计
--2.日期维度的汇总模型,基于用户粒度的向上汇总
DROP TABLE IF EXISTS dws_commerce.dws_user_behavior_analysis_d;
CREATE TABLE IF NOT EXISTS dws_commerce.dws_user_behavior_analysis_d (
  view_cnt BIGINT COMMENT '浏览次数'
  ,fav_cnt BIGINT COMMENT '收藏次数'
  ,cart_cnt BIGINT COMMENT '加购物车次数'
  ,buy_cnt BIGINT COMMENT '购买次数'
  ,view_uv BIGINT COMMENT '浏览人数'
  ,fav_uv BIGINT COMMENT '收藏人数'
  ,cart_uv BIGINT COMMENT '加购物车人数'
  ,buy_uv BIGINT COMMENT '购买人数'
) 
COMMENT '用户行为指标统计表' 
PARTITIONED BY (dt string comment '日期')
ROW FORMAT DELIMITED FIELDS TERMINATED BY ',' 
STORED AS orc;

insert overwrite table dws_commerce.dws_user_behavior_analysis_d partition(dt)
select 
 sum(if(type='pv',1,0))  as view_cnt
,sum(if(type='fav',1,0)) as fav_cnt
,sum(if(type='cart',1,0)) as cart_cnt
,sum(if(type='buy',1,0)) as buy_cnt
,count(distinct if(type='pv',user_id,null))  as view_uv
,count(distinct if(type='fav',user_id,null))  as fav_uv
,count(distinct if(type='cart',user_id,null))  as cart_uv
,count(distinct if(type='buy',user_id,null))  as buy_uv
,dt
from dwd_commerce.dwd_user_behavior_detail_d
where dt between '2017-11-01' and '2017-12-03'
group by 
dt
;

insert overwrite table dws_commerce.dws_user_behavior_analysis_d partition(dt)
select 
 sum(if(type='pv',1,0))  as view_cnt
,sum(if(type='fav',1,0)) as fav_cnt
,sum(if(type='cart',1,0)) as cart_cnt
,sum(if(type='buy',1,0)) as buy_cnt
,count(distinct if(type='pv',user_id,null))  as view_uv
,count(distinct if(type='fav',user_id,null))  as fav_uv
,count(distinct if(type='cart',user_id,null))  as cart_uv
,count(distinct if(type='buy',user_id,null))  as buy_uv
,dt
from dwd_commerce.dwd_user_behavior_detail_d
where dt between '2017-11-25' and '2017-12-03'
group by 
dt
;

-- 数据验证：
select * from dws_commerce.dws_user_behavior_analysis_d where dt='2017-12-03' limit 10;

--part3:dws层的数据建设 用户商品指标统计
--3. user_id*item_id粒度的用户行为统计数据
-- 收藏或者加购等行为会显著提升用户购买商品的转化率
DROP TABLE IF EXISTS dws_commerce.dws_user_goods_sum_d;
CREATE TABLE IF NOT EXISTS dws_commerce.dws_user_goods_sum_d (
  user_id STRING COMMENT '用户标识'
  ,goods_id STRING COMMENT '商品标识'
  ,goods_name STRING COMMENT '商品名称'
  ,category_id STRING COMMENT '商品分类标识'
  ,view_cnt BIGINT COMMENT '浏览次数'
  ,fav_cnt BIGINT COMMENT '收藏次数'
  ,cart_cnt BIGINT COMMENT '加购物车次数'
  ,buy_cnt BIGINT COMMENT '购买次数'
) 
COMMENT '用户商品指标统计表' 
PARTITIONED BY (dt string comment '日期')
ROW FORMAT DELIMITED FIELDS TERMINATED BY ',' 
STORED AS orc;

set hive.exec.dynamic.partition.mode=nonstrict;
insert overwrite table dws_commerce.dws_user_goods_sum_d partition(dt)
select 
 a.user_id
,a.item_id as goods_id
,b.goodsname  as goods_name
,a.category_id
,sum(case when a.type='pv' then 1 else 0 end) as view_cnt
,sum(case when a.type='cart' then 1 else 0 end) as cart_cnt
,sum(case when a.type='fav' then 1 else 0 end) as fav_cnt
,sum(case when a.type='buy' then 1 else 0 end) as buy_cnt
,a.dt
from dwd_commerce.dwd_user_behavior_detail_d as a
join ods_commerce.ods_product_goodsinfo as b
on a.item_id=b.goodsid
where a.dt between '2017-11-01' and '2017-12-03'
group by a.user_id,a.item_id,b.goodsname,a.category_id,a.dt;

-- 数据验证：
select * from dws_commerce.dws_user_goods_sum_d where dt='2017-11-25' limit 10;

-------------------------------------3.ads层数据-------------------------------------------------------------------- 
-- 场景一：业务指标分析
DROP TABLE IF EXISTS ads_commerce.ads_user_behavior_analysis_d;
CREATE TABLE IF NOT EXISTS ads_commerce.ads_user_behavior_analysis_d (
  avg_view_cnt BIGINT COMMENT '人均浏览次数'
  ,avg_fav_cnt BIGINT COMMENT '人均收藏次数'
  ,avg_cart_cnt BIGINT COMMENT '人均加购物车次数'
  ,avg_buy_cnt BIGINT COMMENT '人均购买次数'
) 
COMMENT '用户行为业务分析表' 
PARTITIONED BY (dt string comment '日期')
ROW FORMAT DELIMITED FIELDS TERMINATED BY ',' 
STORED AS orc;


-- OK 
set hive.exec.dynamic.partition.mode=nonstrict;
insert overwrite table ads_commerce.ads_user_behavior_analysis_d partition(dt)
select 
 view_cnt/view_uv as avg_view_cnt
,fav_cnt/fav_uv as avg_fav_cnt
,cart_cnt/cart_uv as av_cart_cnt
,buy_cnt/buy_uv as avg_buy_cnt
,dt
from dws_commerce.dws_user_behavior_analysis_d
where dt between '2017-11-01' and '2017-12-03'
;

-- 数据验证：
select * from ads_commerce.ads_user_behavior_analysis_d

--场景2：跳失率:在11月25日至12月3日期间只有浏览行为的用户/总用户数  OK 
DROP TABLE IF EXISTS ads_commerce.ads_user_lost_rate_d;
CREATE TABLE IF NOT EXISTS ads_commerce.ads_user_lost_rate_d (
  lost_rate decimal(20,4) COMMENT '跳失率'
) 
COMMENT '用户跳失率统计表' 
PARTITIONED BY (dt string comment '日期')
ROW FORMAT DELIMITED FIELDS TERMINATED BY ',' 
STORED AS orc;

set hive.exec.dynamic.partition.mode=nonstrict;
insert overwrite table ads_commerce.ads_user_lost_rate_d partition(dt)
select 
count(if(t.fav_cnt=0 and t.cart_cnt=0 and t.buy_cnt=0,user_id,null))/count(user_id)*100 as lost_rate
,dt
from dws_commerce.dws_user_behavior_sum_d as t
where dt between '2017-11-30' and '2017-12-03'
group by dt
;  

-- 如果不分层，直接统计
if(t.fav_cnt=0 and t.cart_cnt=0 and t.buy_cnt=0,user_id,null)

select count(if(t.fav_cnt=0 and t.cart_cnt=0 and t.buy_cnt=0,user_id,null))/count(user_id)*100 as lost_rate
from(
   select user_id,
             sum(if(type='fav',1,0)) fav_cnt,
             sum(if(type='cart',1,0)) cart_cnt,
             sum(if(type='buy',1,0)) buy_cnt
      from dwd_commerce.dwd_user_behavior_detail_d
      where dt='2017-11-30'
      group by user_id
) t;

--分层后统计更为直接
create table ads_commerce.ads_user_lost_rate_d_temp
stored as orc
as 
select 
count(if(t.fav_cnt=0 and t.cart_cnt=0 and t.buy_cnt=0,user_id,null))/count(user_id)*100 as lost_rate
,dt
from dws_commerce.dws_user_behavior_sum_d as t
where dt='2017-11-30'
group by dt
;

-- 数据验证：
select * from ads_commerce.ads_user_lost_rate_d
-- 结果：46.37%

--场景3：业务指标分析
--获取每个商品类目浏览量排名前10的商品、以及他们的浏览次数
DROP TABLE IF EXISTS ads_commerce.ads_goods_category_analysis_d;
CREATE TABLE IF NOT EXISTS ads_commerce.ads_goods_category_analysis_d (
  category_id STRING COMMENT '商品类别ID'
  ,goods_id STRING COMMENT '商品标识'
  ,goods_name STRING COMMENT '商品名称'
  ,view_cnt BIGINT COMMENT '浏览次数'
  ,view_rnk BIGINT COMMENT '浏览排名'
) 
COMMENT '商品类目运营排名表' 
PARTITIONED BY (dt string comment '日期')
ROW FORMAT DELIMITED FIELDS TERMINATED BY ',' 
STORED AS orc;

insert overwrite table ads_commerce.ads_goods_category_analysis_d partition(dt)
select category_id,goods_id,goods_name,view_cnt,rnk as view_rnk,dt
from (
select 
   dt
   ,category_id
   ,goods_id
   ,goods_name
   ,view_cnt
   ,RANK() OVER(PARTITION BY dt,category_id ORDER BY view_cnt desc) AS rnk
from (
  select 
    dt
    ,category_id
    ,goods_id
    ,goods_name
    ,sum(view_cnt) as view_cnt
  from dws_commerce.dws_user_goods_sum_d
  where dt between '2017-11-03' and '2017-11-03'
  group by dt,category_id,goods_id,goods_name
) as tmp
) as a where rnk<=10
;

-- 数据验证：
select * from ads_commerce.ads_goods_category_analysis_d

--场景4：用户分层  这里需要注意max_dd的字段类型
DROP TABLE IF EXISTS ads_commerce.ads_user_behavior_level_full;
CREATE TABLE IF NOT EXISTS ads_commerce.ads_user_behavior_level_full (
  user_id STRING COMMENT '用户ID'
  ,max_dd STRING COMMENT '最近访问日期'
  ,min_dd STRING COMMENT '最早访问日期'
  ,diff BIGINT COMMENT '浏览间隔天数'
  ,r_score BIGINT COMMENT '用户分层'
) 
COMMENT '用户分层表' 
STORED AS orc;

-- 分层数据加工逻辑
with ta as (
select user_id,to_date(from_unixtime(cast(action_time as bigint),'yyyy-MM-dd')) as dd
from ods_commerce.ods_userbehavior
where type='pv'
)
,tb as (
select 
user_id
,max(dd) as max_dd
,min(dd) as min_dd
,datediff(max(dd),min(dd)) as diff 
from ta
group by user_id
)
insert overwrite table ads_commerce.ads_user_behavior_level_full
select 
user_id
,max_dd
,min_dd
,diff
,(case when diff BETWEEN 0 and 2 then 4
      when diff BETWEEN 3 and 4 then 3
      when diff BETWEEN 5 and 6 then 2
      when diff BETWEEN 7 and 8 then 1
else 0 end ) as r_score
from tb
;

-- 数据验证
--1.查看执行结果
select * from ads_commerce.ads_user_behavior_level_full limit 100;
--2.查看用户分层分布
select r_score,count(*) from ads_commerce.ads_user_behavior_level_full group by r_score;
0       39638
1       801721
2       104106
3       32666
4       5983
--2.查看用户浏览天数分布
select diff,count(*) from ads_commerce.ads_user_behavior_level_full group by diff;

ods层用到3个表(用户信息表、用户行为表、商品表)，dwd层2个，dws层3个，ads层4个

