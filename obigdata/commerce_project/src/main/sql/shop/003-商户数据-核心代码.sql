-------------------------------------0.ods层数据--------------------------------------------------------------------
DROP TABLE IF EXISTS ods_shop_base_info;
CREATE TABLE IF NOT EXISTS ods_shop_base_info (
  shop_id STRING COMMENT '店铺ID',
  shop_name STRING COMMENT '店铺名称',
  shop_status int COMMENT '营业状态(0在线营业1暂时歇业2停业)'
) COMMENT '店铺表' STORED AS textfile;

DROP TABLE IF EXISTS ods_shop_spu_info;
CREATE TABLE IF NOT EXISTS ods_shop_spu_info (
  shop_id STRING COMMENT '店铺ID',
  shop_name STRING COMMENT '店铺名称',
  spu_id int COMMENT '商品ID'
) COMMENT '店铺商品表' STORED AS textfile;


-------------------------------------1.dim层数据--------------------------------------------------------------------
-- part1:店铺表的构建
DROP TABLE IF EXISTS dim_commerce.dim_shop_base_info_ss;
CREATE TABLE IF NOT EXISTS dim_commerce.dim_shop_base_info_ss (
  shop_id STRING COMMENT '店铺ID',
  shop_name STRING COMMENT '店铺名称',
  shop_status int COMMENT '营业状态(0在线营业1暂时歇业2停业)'
) 
COMMENT '店铺信息表'
PARTITIONED BY (dt string comment '数据日期') 
STORED AS orc;

--核心代码
set hive.exec.dynamic.partition.mode=nonstrict;
insert overwrite table dim_commerce.dim_shop_base_info_ss partition(dt)
select 
  shop_id,
  shop_name,
  shop_status,
  date_sub(current_date(),1) as dt
from ods_commerce.ods_shop_base_info
;

--验证数据:OK
select * from dim_commerce.dim_shop_base_info_ss where dt=date_sub(current_date(),1) limit 10;

-------------------------------------2.dwd层数据--------------------------------------------------------------------
-- part2:店铺-商品关系表的构建
DROP TABLE IF EXISTS dwd_commerce.dwd_shop_spu_info_ss;
CREATE TABLE IF NOT EXISTS dwd_commerce.dwd_shop_spu_info_ss (
  shop_id STRING COMMENT '店铺ID',
  shop_name STRING COMMENT '店铺名称',
  spu_id int COMMENT '商品ID',
  spu_name int COMMENT '商品名称'
) 
COMMENT '店铺商品表'
PARTITIONED BY (dt string comment '数据日期') 
STORED AS orc;

--核心代码
set hive.exec.dynamic.partition.mode=nonstrict;
insert overwrite table dwd_commerce.dwd_shop_spu_info_ss partition(dt)
select 
  a.shop_id,
  a.shop_name,
  a.spu_id,
  c.goodsname as spu_name,
  date_sub(current_date(),1) as dt
from ods_commerce.ods_shop_spu_info as a
join (
  select shop_id,shop_status 
  from dim_commerce.dim_shop_base_info_ss 
  where dt=date_sub(current_date(),1) and shop_status in (0,1)
  ) as b
on a.shop_id=b.shop_id
left join (
select goodsid,max(goodsname) as goodsname 
from ods_commerce.ods_product_goodsinfo
where goodsname is not null
group by goodsid
) as c  -- 此次对脏数据进行处理
on a.spu_id=c.goodsid
;

--验证数据:OK
select * from dwd_commerce.dwd_shop_spu_info_ss limit 10;

-------------------------------------3.dws层数据--------------------------------------------------------------------
--part3:店铺订单统计汇总表
DROP TABLE IF EXISTS dws_commerce.dws_shop_order_summary_full;
CREATE TABLE IF NOT EXISTS dws_commerce.dws_shop_order_summary_full (
  `shop_id` STRING COMMENT '店铺ID',
  `shop_name` STRING COMMENT '店铺名称',
  `order_cnt` BIGINT COMMENT '下单次数',
  `order_num` BIGINT COMMENT '下单件数',
  `order_coupon_cnt` BIGINT COMMENT '使用优惠券下单次数',
  `order_total_amount` DECIMAL(20,4) COMMENT '下单订单总额',
  `order_pay_amount` DECIMAL(20,4) COMMENT '下单应付总额',
  `order_freight_amount` DECIMAL(20,4) COMMENT '下单运费金额',
  `order_promotion_amount` DECIMAL(20,4) COMMENT '下单促销优化总金额（促销价、满减、阶梯价）',
  `order_integration_amount` DECIMAL(20,4) COMMENT '下单积分抵扣总金额',
  `order_coupon_amount` DECIMAL(20,4) COMMENT '下单优惠券抵扣总金额',
  `order_discount_amount` DECIMAL(20,4) COMMENT '下单后台调整订单使用的折扣总金额',
  `refund_payment_cnt` BIGINT COMMENT '被退款次数',
  `refund_payment_num` BIGINT COMMENT '被退款件数',
  `refund_payment_amount` DECIMAL(20,4) COMMENT '被退款金额',
  `browse_cnt` BIGINT COMMENT '商品浏览次数',
  `collection_cnt` BIGINT COMMENT '商品收藏次数',
  `shopping_cart_cnt` BIGINT COMMENT '商品加入购物车次数'
  )COMMENT "店铺订单统计汇总表"
stored as orc
;

-- 核心代码
set hive.exec.dynamic.partition.mode=nonstrict;
insert overwrite table  dws_commerce.dws_shop_order_summary_full
select
  shop_id,
  shop_name,
  sum(order_cnt) as order_cnt,
  sum(order_num) as order_num,
  sum(order_coupon_cnt) as order_coupon_cnt,
  sum(order_total_amount) as order_total_amount,
  sum(order_pay_amount) as order_pay_amount,
  sum(order_freight_amount) as order_freight_amount,
  sum(order_promotion_amount) as order_promotion_amount,
  sum(order_integration_amount) as order_integration_amount,
  sum(order_coupon_amount) as order_coupon_amount,
  sum(order_discount_amount) as order_discount_amount,
  sum(refund_payment_cnt) as refund_payment_cnt,
  sum(refund_payment_num) as refund_payment_num,
  sum(refund_payment_amount) as refund_payment_amount,
  sum(browse_cnt) as browse_cnt,
  sum(collection_cnt) as collection_cnt,
  sum(shopping_cart_cnt) as shopping_cart_cnt
from
  dwd_commerce.dwd_shop_spu_info_ss a
  left join dws_commerce.dws_sku_summary_full b on a.spu_id = b.spu_id
  where a.dt='2022-07-25' and a.shop_name is not null
group by
  shop_id,
  shop_name;

-- dws_commerce.dws_sku_summary_full 在交易主题域总进行建设
-- 数据验证 ok 
select * from dws_commerce.dws_shop_order_summary_full limit 10;


--part4:店铺商品的统计汇总表

DROP TABLE IF EXISTS dws_commerce.dws_shop_spu_order_summary_full;
CREATE TABLE IF NOT EXISTS dws_commerce.dws_shop_spu_order_summary_full (
  `shop_id` STRING COMMENT '店铺ID',
  `shop_name` STRING COMMENT '店铺名称',
  `spu_id` STRING COMMENT 'spu_id',
  `order_cnt` BIGINT COMMENT '下单次数',
  `order_num` BIGINT COMMENT '下单件数'
  )COMMENT "店铺商品的统计汇总表"
stored as orc
;


-- 核心代码
insert overwrite table  dws_commerce.dws_shop_spu_order_summary_full
select
  a.shop_id,
  a.shop_name,
  a.spu_id,
  sum(order_cnt) as order_cnt,
  sum(order_num) as order_num
from
  dwd_commerce.dwd_shop_spu_info_ss a
  left join dws_commerce.dws_sku_summary_full b on a.spu_id = b.spu_id
 where a.dt='2022-07-25' and a.shop_name is not null  
group by
  a.shop_id,
  a.shop_name,
  a.spu_id
;

-- 数据验证 ok 
select * from dws_commerce.dws_shop_spu_order_summary_full limit 10;

-------------------------------------4.ads层数据--------------------------------------------------------------------
--part5:店铺商品的下单排名表

DROP TABLE IF EXISTS ads_commerce.ads_shop_spu_order_rank_full;
CREATE TABLE IF NOT EXISTS ads_commerce.ads_shop_spu_order_rank_full (
  `shop_id` STRING COMMENT '店铺ID',
  `shop_name` STRING COMMENT '店铺名称',
  `spu_id` STRING COMMENT 'spu_id',
  `order_cnt` BIGINT COMMENT '下单次数',
  `order_num` BIGINT COMMENT '下单件数',
  `rnk` BIGINT COMMENT '排名'
  )COMMENT "店铺商品的下单排名表"
stored as orc
;

--核心代码
insert overwrite table  ads_commerce.ads_shop_spu_order_rank_full
select
  shop_id,
  shop_name,
  spu_id,
  order_cnt,
  order_num,
  rnk
from (
select
  shop_id,
  shop_name,
  spu_id,
  order_cnt,
  order_num,
  row_number() over (partition by shop_id order by order_cnt desc) as rnk
from dws_commerce.dws_shop_spu_order_summary_full
) as tmp where rnk<=3;

-- 数据验证 ok 
select * from ads_commerce.ads_shop_spu_order_rank_full limit 10;


ods层用到2个表，dim层1个，dwd层1个，dws层2个，ads层1个
