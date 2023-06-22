-------------------------------------0.ods层数据--------------------------------------------------------------------
DROP TABLE IF EXISTS ods_market_coupon;
CREATE TABLE IF NOT EXISTS ods_market_coupon (
  id BIGINT COMMENT 'id',
  coupon_type INT COMMENT '优惠卷类型[0->全场赠券；1->会员赠券；2->购物赠券；3->注册赠券]',
  coupon_img STRING COMMENT '优惠券图片',
  coupon_name STRING COMMENT '优惠卷名字',
  num INT COMMENT '数量',
  amount DOUBLE COMMENT '金额',
  per_limit INT COMMENT '每人限领张数',
  min_point DOUBLE COMMENT '使用门槛',
  start_time STRING COMMENT '开始时间',
  end_time STRING COMMENT '结束时间',
  use_type INT COMMENT '使用类型[0->全场通用；1->指定分类；2->指定商品]',
  note STRING COMMENT '备注',
  publish_count INT COMMENT '发行数量',
  use_count INT COMMENT '已使用数量',
  receive_count INT COMMENT '领取数量',
  enable_start_time STRING COMMENT '可以领取的开始日期',
  enable_end_time STRING COMMENT '可以领取的结束日期',
  code STRING COMMENT '优惠码',
  member_level INT COMMENT '可以领取的会员等级[0->不限等级，其他-对应等级]',
  publish INT COMMENT '发布状态[0-未发布，1-已发布]'
) COMMENT '优惠券信息' STORED AS textfile;
USE dwd_commerce;
DROP TABLE IF EXISTS dwd_user_userinfo;
CREATE TABLE IF NOT EXISTS dwd_user_userinfo (
                                                 user_id STRING COMMENT 'mistake'
    ,user_name STRING COMMENT ' garbage stem'
    ,sex INT COMMENT ' leave ice aisle thumb '
    ,user_money decimal(20,4) COMMENT 'sport input alien wood settle'
    ,frozen_money decimal(20,4) COMMENT '近一个月的花费的总的金额'
    ,address_id STRING COMMENT '用户地址ID，0表示没有获取地址'
    ,reg_time STRING COMMENT '注册时间'
    ,last_login STRING COMMENT '最后登录时间'
    ) COMMENT '用户表' STORED AS ORC;

--ROW FORMAT DELIMITED FIELDS TERMINATED BY '#' STORED AS textfile;
DROP TABLE IF EXISTS ods_market_coupon_spu_category;
CREATE TABLE IF NOT EXISTS ods_market_coupon_spu_category (
  id BIGINT COMMENT 'id',
  coupon_id BIGINT COMMENT '优惠券id',
  category_id BIGINT COMMENT '产品分类id',
  category_name STRING COMMENT '产品分类名称'
) COMMENT '优惠券分类关联' STORED AS textfile;

-------------------------------------1.dim层数据--------------------------------------------------------------------

-- part1:优惠券基础数据表的构建
DROP TABLE IF EXISTS dim_commerce.dim_market_coupon_info_full;
CREATE TABLE IF NOT EXISTS dim_commerce.dim_market_coupon_info_full (
  coupon_id BIGINT COMMENT 'id',
  coupon_type INT COMMENT '优惠券类型[0->全场赠券；1->会员赠券；2->购物赠券；3->注册赠券]',
  coupon_img STRING COMMENT '优惠券图片',
  coupon_name STRING COMMENT '优惠券名字',
  num INT COMMENT '数量',
  amount DOUBLE COMMENT '金额',
  per_limit INT COMMENT '每人限领张数',
  min_point DOUBLE COMMENT '使用门槛',
  start_time STRING COMMENT '开始时间',
  end_time STRING COMMENT '结束时间',
  use_type INT COMMENT '使用类型[0->全场通用；1->指定分类；2->指定商品]',
  note STRING COMMENT '备注',
  publish_count INT COMMENT '发行数量',
  use_count INT COMMENT '已使用数量',
  receive_count INT COMMENT '领取数量',
  enable_start_time STRING COMMENT '可以领取的开始日期',
  enable_end_time STRING COMMENT '可以领取的结束日期',
  code STRING COMMENT '优惠码',
  member_level INT COMMENT '可以领取的会员等级[0->不限等级，其他-对应等级]',
  publish INT COMMENT '发布状态[0-未发布，1-已发布]',
  category_id BIGINT COMMENT '产品分类id',
  category_name STRING COMMENT '产品分类名称'
) 
COMMENT '优惠券基础信息表'
STORED AS orc;

--核心代码
insert overwrite table dim_commerce.dim_market_coupon_info_full
select 
  a.id as coupon_id,
  coupon_type,
  coupon_img,
  coupon_name,
  num,
  amount,
  per_limit,
  min_point,
  start_time,
  end_time,
  use_type,
  note,
  publish_count,
  use_count,
  receive_count,
  enable_start_time,
  enable_end_time,
  code,
  member_level,
  publish,
  b.category_id,
  b.category_name
from ods_commerce.ods_market_coupon as a
left join ods_commerce.ods_market_coupon_spu_category as b
on a.id=b.coupon_id
;


--数据验证 OK
select * from dim_commerce.dim_market_coupon_info_full limit 10;

-------------------------------------2.dwd层数据--------------------------------------------------------------------
-- part2:优惠券-商品关系表的构建
DROP TABLE IF EXISTS dwd_commerce.dwd_market_coupon_spu_d;
CREATE TABLE IF NOT EXISTS dwd_commerce.dwd_market_coupon_spu_d (
  id BIGINT COMMENT 'id',
  coupon_id BIGINT COMMENT '优惠券id',
  coupon_name BIGINT COMMENT '优惠券名称',
  spu_id BIGINT COMMENT '商品id',
  spu_name STRING COMMENT '商品名称'
) 
COMMENT '优惠券与产品关联关系表'
PARTITIONED BY (dt string comment '日期') 
STORED AS orc;

set hive.exec.dynamic.partition.mode=nonstrict;
insert overwrite table dwd_commerce.dwd_market_coupon_spu_d partition(dt)
select 
  a.id,
  a.coupon_id,
  b.coupon_name,
  a.spu_id,
  a.spu_name,
  date_sub(current_date(),1) as dt
from ods_commerce.ods_market_coupon_spu as a 
left join dim_commerce.dim_market_coupon_info_full as b 
on a.coupon_id=b.coupon_id
;

-- 数据验证：OK
select * from dwd_commerce.dwd_market_coupon_spu_d limit 10;
--select * from dwd_commerce.dwd_market_coupon_spu_d where coupon_name is not null limit 10;  -- coupon_name为空，数据待构建


-------------------------------------3.dws层数据--------------------------------------------------------------------
-- part3:优惠券使用情况统计表
DROP TABLE IF EXISTS dws_commerce.dws_market_coupon_order_d;
CREATE TABLE IF NOT EXISTS dws_commerce.dws_market_coupon_order_d (
  coupon_type INT COMMENT '优惠卷类型[0->全场赠券；1->会员赠券；2->购物赠券；3->注册赠券]',
  coupon_id BIGINT COMMENT '优惠券id',
  coupon_name string COMMENT '优惠券名称',
  receiver_province string COMMENT '省份/直辖市', 
  receiver_city string COMMENT '城市',
  coupon_amount double COMMENT '优惠券抵扣金额',
  coupon_order_cnt BIGINT COMMENT '用券订单数',
  coupon_refund_cnt BIGINT COMMENT '用券退单数'
) 
COMMENT '优惠券使用情况统计表'
PARTITIONED BY (dt string comment '日期') 
STORED AS orc;


set hive.exec.dynamic.partition.mode=nonstrict;
insert overwrite table dws_commerce.dws_market_coupon_order_d partition(dt)
select 
a.coupon_type
,a.coupon_id
,coalesce(a.coupon_name,concat('优惠券',a.coupon_id)) as coupon_name
,coalesce(b.receiver_province,'0') as receiver_province
,coalesce(b.receiver_city,'0') as receiver_city
,sum(b.coupon_amount) as coupon_amount
,count(distinct b.order_id) as coupon_order_cnt
,count(distinct if(b.is_refund=1,b.order_id,null)) as  coupon_refund_cnt
,date_sub(current_date(),1) as dt
from dim_commerce.dim_market_coupon_info_full as a
join dwd_commerce.dwd_trade_order_info_detail_d as b
on a.coupon_id=b.coupon_id
group by 
a.coupon_type
,a.coupon_id
,coalesce(a.coupon_name,concat('优惠券',a.coupon_id))
,coalesce(b.receiver_province,'0')
,coalesce(b.receiver_city,'0')
;

--验证数据：OK
select * from dws_commerce.dws_market_coupon_order_d where dt=date_sub(current_date(),1) limit 10;  -- 备注：coupon_name为空、coupon_amount为0，属于缺少底层数据的问题


-------------------------------------4.ads层数据--------------------------------------------------------------------
-- part4:应用层的统计表
--场景1：获取每种类别的消费券在每个省份、每个城市的下单数排名前3的数据，要求输出券类型、券ID和名称、省份、城市、下单订单数、排名
DROP TABLE IF EXISTS ads_commerce.ads_market_coupon_analysis_d;
CREATE TABLE IF NOT EXISTS ads_commerce.ads_market_coupon_analysis_d (
  coupon_type INT COMMENT '优惠卷类型[0->全场赠券；1->会员赠券；2->购物赠券；3->注册赠券]',
  coupon_id BIGINT COMMENT '优惠券id',
  coupon_name string COMMENT '优惠券名称',
  receiver_province string COMMENT '省份/直辖市', 
  receiver_city string COMMENT '城市',
  coupon_order_cnt BIGINT COMMENT '用券订单数',
  rank BIGINT COMMENT '用券订单量排名'
) 
COMMENT '优惠券使用情况分析表' 
STORED AS orc;

--核心代码
with tmp as(
select
coupon_type
,coupon_id
,coupon_name
,receiver_province
,receiver_city
,coupon_order_cnt
,row_number() over (partition by coupon_type,receiver_province,receiver_city order by coupon_order_cnt desc) as rnk
from dws_commerce.dws_market_coupon_order_d
)
insert overwrite table ads_commerce.ads_market_coupon_analysis_d
select 
coupon_type
,coupon_id
,coupon_name
,receiver_province
,receiver_city
,coupon_order_cnt
,rnk
from tmp where rnk<=3
; 

--验证数据：OK
select * from ads_commerce.ads_market_coupon_analysis_d limit 10;

--场景2：获取消费券在每个省份、每个城市的下单数排名前3的数据，要求输出省份、城市、下单订单数、排名，目的是为了看优惠券在各个城市的投放效果
DROP TABLE IF EXISTS ads_commerce.ads_market_coupon_city_d;
CREATE TABLE IF NOT EXISTS ads_commerce.ads_market_coupon_city_d (
  receiver_province string COMMENT '省份/直辖市', 
  receiver_city string COMMENT '城市',
  coupon_order_cnt BIGINT COMMENT '用券订单数',
  rank BIGINT COMMENT '用券订单量排名'
) 
COMMENT '优惠券在各个城市的使用情况统计表' 
STORED AS orc;

-- 核心代码
with ta as(
select
receiver_province
,receiver_city
,sum(coupon_order_cnt) as coupon_order_cnt
from dws_commerce.dws_market_coupon_order_d
group by 
receiver_province
,receiver_city
)
,tb as(
select
receiver_province
,receiver_city
,coupon_order_cnt
,row_number() over (partition by receiver_province,receiver_city order by coupon_order_cnt desc) as rnk
from ta
)
insert overwrite table ads_commerce.ads_market_coupon_city_d
select 
receiver_province
,receiver_city
,coupon_order_cnt
,rnk
from tb where rnk<=3
;

--验证数据：OK
select * from ads_commerce.ads_market_coupon_city_d limit 10;

ods层用到3个表(优惠券表、券分类表、券-商品关系表)，dim层1个，dwd层1个，dws层1个，ads层2个
ods_market_coupon_spu
ods_market_coupon
ods_market_coupon_spu_category