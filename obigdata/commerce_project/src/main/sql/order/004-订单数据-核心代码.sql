-------------------------------------0.ods层数据--------------------------------------------------------------------
DROP TABLE IF EXISTS ods_trade_order_item;
CREATE TABLE IF NOT EXISTS ods_trade_order_item (
  id BIGINT COMMENT 'id',
  order_id BIGINT COMMENT 'order_id',
  order_sn STRING COMMENT 'order_sn',
  spu_id BIGINT COMMENT 'spu_id',
  spu_name STRING COMMENT 'spu_name',
  spu_pic STRING COMMENT 'spu_pic',
  spu_brand STRING COMMENT '品牌',
  category_id BIGINT COMMENT '商品分类id',
  sku_id BIGINT COMMENT '商品sku编号',
  sku_name STRING COMMENT '商品sku名字',
  sku_pic STRING COMMENT '商品sku图片',
  sku_price DOUBLE COMMENT '商品sku价格',
  sku_quantity INT COMMENT '商品购买的数量',
  sku_attrs_vals STRING COMMENT '商品销售属性组合（JSON）',
  promotion_amount DOUBLE COMMENT '商品促销分解金额',
  coupon_amount DOUBLE COMMENT '优惠券优惠分解金额',
  integration_amount DOUBLE COMMENT '积分优惠分解金额',
  real_amount DOUBLE COMMENT '该商品经过优惠后的分解金额',
  gift_integration INT COMMENT '赠送积分',
  gift_growth INT COMMENT '赠送成长值'
) COMMENT '订单项信息' STORED AS textfile;


DROP TABLE IF EXISTS ods_trade_order2;
CREATE TABLE IF NOT EXISTS ods_trade_order2 (
  id BIGINT COMMENT 'id',
  user_id BIGINT COMMENT 'member_id',
  order_sn STRING COMMENT '订单号',
  coupon_id BIGINT COMMENT '使用的优惠券',
  create_time STRING COMMENT '创建时间',
  username STRING COMMENT '用户名',
  total_amount DOUBLE COMMENT '订单总额',
  pay_amount DOUBLE COMMENT '应付总额',
  freight_amount DOUBLE COMMENT '运费金额',
  promotion_amount DOUBLE COMMENT '促销优化金额（促销价、满减、阶梯价）',
  integration_amount DOUBLE COMMENT '积分抵扣金额',
  coupon_amount DOUBLE COMMENT '优惠券抵扣金额',
  discount_amount DOUBLE COMMENT '后台调整订单使用的折扣金额',
  pay_type INT COMMENT '支付方式【1->支付宝；2->微信；3->银联； 4->货到付款；】',
  source_type INT COMMENT '订单来源[0->PC订单；1->app订单]',
  status INT COMMENT '订单状态【0->待付款；1->待发货；2->已发货；3->已完成；4->已关闭；5->无效订单】',
  delivery_company STRING COMMENT '物流公司(配送方式)',
  delivery_sn STRING COMMENT '物流单号',
  auto_confirm_day INT COMMENT '自动确认时间（天）',
  integration INT COMMENT '可以获得的积分',
  growth INT COMMENT '可以获得的成长值',
  bill_type INT COMMENT '发票类型[0->不开发票；1->电子发票；2->纸质发票]',
  bill_header STRING COMMENT '发票抬头',
  bill_content STRING COMMENT '发票内容',
  bill_receiver_phone STRING COMMENT '收票人电话',
  bill_receiver_email STRING COMMENT '收票人邮箱',
  receiver_name STRING COMMENT '收货人姓名',
  receiver_phone STRING COMMENT '收货人电话',
  receiver_post_code STRING COMMENT '收货人邮编',
  receiver_province STRING COMMENT '省份/直辖市',
  receiver_city STRING COMMENT '城市',
  receiver_region STRING COMMENT '区',
  receiver_address STRING COMMENT '详细地址',
  confirm_status INT COMMENT '确认收货状态[0->未确认；1->已确认]',
  delete_status INT COMMENT '删除状态【0->未删除；1->已删除】',
  use_integration INT COMMENT '下单时使用的积分',
  payment_time STRING COMMENT '支付时间',
  delivery_time STRING COMMENT '发货时间',
  receive_time STRING COMMENT '确认收货时间',
  comment_time STRING COMMENT '评价时间',
  modify_time STRING COMMENT '修改时间',
  remark STRING COMMENT '订单备注'
) COMMENT '订单' STORED AS textfile;


DROP TABLE IF EXISTS ods_trade_refund_info;
CREATE TABLE IF NOT EXISTS ods_trade_refund_info (
  id BIGINT COMMENT 'id',
  order_return_id BIGINT COMMENT '退款的订单',
  refund DOUBLE COMMENT '退款金额',
  refund_sn STRING COMMENT '退款交易流水号',
  refund_status INT COMMENT '退款状态',
  refund_channel INT COMMENT '退款渠道[1-支付宝，2-微信，3-银联，4-汇款]',
  refund_content STRING COMMENT ''
) COMMENT '退款信息' STORED AS textfile;


-------------------------------------1.dwd层数据--------------------------------------------------------------------
--part1.订单项明细事实表
DROP TABLE IF EXISTS dwd_commerce.dwd_trade_order_item_detail_d;
CREATE TABLE IF NOT EXISTS dwd_commerce.dwd_trade_order_item_detail_d(
  `order_id` bigint COMMENT '订单id', 
  `order_sn` string COMMENT '订单编号', 
  `user_id` string COMMENT 'user_id', 
  `spu_id` bigint COMMENT 'spu_id', 
  `spu_name` string COMMENT 'spu_name', 
  `spu_pic` string COMMENT 'spu_pic', 
  `spu_brand` string COMMENT '品牌', 
  `category_id` bigint COMMENT '商品分类id', 
  `sku_id` bigint COMMENT '商品sku编号', 
  `sku_name` string COMMENT '商品sku名字', 
  `sku_pic` string COMMENT '商品sku图片', 
  `sku_price` double COMMENT '商品sku价格', 
  `sku_quantity` int COMMENT '商品购买的数量', 
  `color_id` string COMMENT '颜色ID', 
  `size_id` string COMMENT '尺码ID',
  `coupon_id` bigint COMMENT '使用的优惠券', 
  `create_time` string COMMENT '创建时间', 
  `total_amount` double COMMENT '订单总额', 
  `pay_amount` double COMMENT '应付总额', 
  `freight_amount` double COMMENT '运费金额', 
  `promotion_amount` double COMMENT '促销优化金额（促销价、满减、阶梯价）', 
  `integration_amount` double COMMENT '积分抵扣金额', 
  `coupon_amount` double COMMENT '优惠券抵扣金额', 
  `discount_amount` double COMMENT '后台调整订单使用的折扣金额', 
  `receiver_province` string COMMENT '省份/直辖市', 
  `receiver_city` string COMMENT '城市', 
  `receiver_region` string COMMENT '区', 
  `receiver_address` string COMMENT '详细地址'
  ) COMMENT '订单项明细事实表'
PARTITIONED BY (dt string comment '日期')
STORED AS orc;


set hive.exec.dynamic.partition.mode=nonstrict;
insert overwrite table dwd_commerce.dwd_trade_order_item_detail_d partition(dt)
select
  a.order_id,
  a.order_sn,
  a.user_id,
  a.spu_id,
  a.spu_name,
  a.spu_pic,
  a.spu_brand,
  a.category_id,
  a.sku_id,
  a.sku_name,
  a.sku_pic,
  a.sku_price,
  a.sku_quantity,
  a.color_id,
  a.size_id,
  b.coupon_id,
  b.create_time,
  b.total_amount,
  b.pay_amount,
  b.freight_amount,
  b.promotion_amount,
  b.integration_amount,
  b.coupon_amount,
  b.discount_amount,
  b.receiver_province,
  b.receiver_city,
  b.receiver_region,
  b.receiver_address,
  b.dt
from
  ods_commerce.ods_trade_order_item a
  left join ods_commerce.ods_trade_order2 b on a.order_id = b.id;

-- 数据验证:OK
select * from dwd_commerce.dwd_trade_order_item_detail_d limit 10;



-- part2.订单信息明细事实表
DROP TABLE IF EXISTS dwd_commerce.dwd_trade_order_info_detail_d;
CREATE TABLE IF NOT EXISTS dwd_commerce.dwd_trade_order_info_detail_d(
  `order_id` bigint COMMENT '订单id', 
  `user_id` bigint COMMENT 'member_id', 
  `order_sn` string COMMENT '订单号', 
  `coupon_id` bigint COMMENT '使用的优惠券', 
  `create_time` string COMMENT '创建时间', 
  `username` string COMMENT '用户名', 
  `total_amount` double COMMENT '订单总额', 
  `pay_amount` double COMMENT '应付总额', 
  `freight_amount` double COMMENT '运费金额', 
  `promotion_amount` double COMMENT '促销优化金额（促销价、满减、阶梯价）', 
  `integration_amount` double COMMENT '积分抵扣金额', 
  `coupon_amount` double COMMENT '优惠券抵扣金额', 
  `discount_amount` double COMMENT '后台调整订单使用的折扣金额', 
  `pay_type` int COMMENT '支付方式【1->支付宝；2->微信；3->银联； 4->货到付款；】', 
  `source_type` int COMMENT '订单来源[0->PC订单；1->app订单]', 
  `status` int COMMENT '订单状态【0->待付款；1->待发货；2->已发货；3->已完成；4->已关闭；5->无效订单】', 
  `delivery_company` string COMMENT '物流公司(配送方式)', 
  `delivery_sn` string COMMENT '物流单号', 
  `auto_confirm_day` int COMMENT '自动确认时间（天）', 
  `integration` int COMMENT '可以获得的积分', 
  `growth` int COMMENT '可以获得的成长值', 
  `receiver_name` string COMMENT '收货人姓名', 
  `receiver_phone` string COMMENT '收货人电话', 
  `receiver_post_code` string COMMENT '收货人邮编', 
  `receiver_province` string COMMENT '省份/直辖市', 
  `receiver_city` string COMMENT '城市', 
  `receiver_region` string COMMENT '区', 
  `receiver_address` string COMMENT '详细地址', 
  `confirm_status` int COMMENT '确认收货状态[0->未确认；1->已确认]', 
  `delete_status` int COMMENT '删除状态【0->未删除；1->已删除】', 
  `use_integration` int COMMENT '下单时使用的积分', 
  `payment_time` string COMMENT '支付时间', 
  `delivery_time` string COMMENT '发货时间', 
  `receive_time` string COMMENT '确认收货时间', 
  `comment_time` string COMMENT '评价时间', 
  `modify_time` string COMMENT '修改时间',
  `is_refund` int COMMENT '是否为退款订单', 
  `refund` double COMMENT '退款金额', 
  `refund_sn` string COMMENT '退款交易流水号', 
  `refund_status` int COMMENT '退款状态', 
  `refund_channel` int COMMENT '退款渠道[1-支付宝，2-微信，3-银联，4-汇款]'
  ) COMMENT '订单信息明细事实表'
PARTITIONED BY (dt string comment '日期')
STORED AS orc;

-- 核心代码
set hive.exec.dynamic.partition.mode=nonstrict;

insert overwrite table dwd_commerce.dwd_trade_order_info_detail_d partition(dt)
select
  a.id as order_id,
  user_id,
  order_sn,
  coupon_id,
  create_time,
  username,
  total_amount,
  pay_amount,
  freight_amount,
  promotion_amount,
  integration_amount,
  coupon_amount,
  discount_amount,
  pay_type,
  source_type,
  status,
  delivery_company,
  delivery_sn,
  auto_confirm_day,
  integration,
  growth,
  receiver_name,
  receiver_phone,
  receiver_post_code,
  receiver_province,
  receiver_city,
  receiver_region,
  receiver_address,
  confirm_status,
  delete_status,
  use_integration,
  payment_time,
  delivery_time,
  receive_time,
  comment_time,
  modify_time,
  if(order_return_id is not null, 1, 0) as is_refund,
  refund,
  refund_sn,
  refund_status,
  refund_channel,
  a.dt
from ods_commerce.ods_trade_order2 a 
left join ods_commerce.ods_trade_refund_info b on a.id = b.order_return_id;

-- 数据验证:OK
select * from dwd_commerce.dwd_trade_order_info_detail_d limit 10;


-- part3.订单退款明细事实表
DROP TABLE IF EXISTS dwd_commerce.dwd_trade_order_refund_info_d;
CREATE TABLE IF NOT EXISTS dwd_commerce.dwd_trade_order_refund_info_d(
  `id` bigint COMMENT 'id', 
  `order_return_id` bigint COMMENT '退款的订单', 
  `refund` double COMMENT '退款金额', 
  `refund_sn` string COMMENT '退款交易流水号', 
  `refund_status` int COMMENT '退款状态', 
  `refund_channel` int COMMENT '退款渠道[1-支付宝，2-微信，3-银联，4-汇款]', 
  `refund_content` string COMMENT '',
  `receiver_province` string COMMENT '省份/直辖市', 
  `receiver_city` string COMMENT '城市', 
  `receiver_region` string COMMENT '区',  
  `user_id` bigint COMMENT 'member_id', 
  `username` string COMMENT '用户名'
  ) COMMENT '订单退款明细事实表'
PARTITIONED BY (dt string comment '日期')
STORED AS orc;

-- 核心代码
set hive.exec.dynamic.partition.mode=nonstrict;

insert overwrite table dwd_commerce.dwd_trade_order_refund_info_d partition(dt)
select
  a.id,
  a.order_return_id,
  a.refund,
  a.refund_sn,
  a.refund_status,
  a.refund_channel,
  a.refund_content,
  b.receiver_province,
  b.receiver_city,
  b.receiver_region,
  b.user_id,
  b.username,
  b.dt
from ods_commerce.ods_trade_refund_info a 
left join ods_commerce.ods_trade_order2 b on  b.id = a.order_return_id;


-- 数据验证:OK
select * from dwd_commerce.dwd_trade_order_refund_info_d limit 10;  


-------------------------------------3.dws层数据--------------------------------------------------------------------
-- 1.dws 商品统计汇总表
DROP TABLE IF EXISTS dws_commerce.dws_spu_summary_d;
CREATE TABLE IF NOT EXISTS dws_commerce.dws_spu_summary_d(
  `spu_id` bigint COMMENT 'spu_id', 
  `order_cnt` bigint  COMMENT '下单次数', 
  `order_num` bigint COMMENT '下单件数', 
  `order_coupon_cnt` bigint COMMENT '使用优惠券下单次数', 
  `order_total_amount` decimal(20,4) COMMENT '下单订单总额', 
  `order_pay_amount` decimal(20,4) COMMENT '下单应付总额', 
  `order_freight_amount` decimal(20,4) COMMENT '下单运费金额', 
  `order_promotion_amount` decimal(20,4) COMMENT '下单促销优化总金额（促销价、满减、阶梯价）', 
  `order_integration_amount` decimal(20,4) COMMENT '下单积分抵扣总金额', 
  `order_coupon_amount` decimal(20,4) COMMENT '下单优惠券抵扣总金额', 
  `order_discount_amount` decimal(20,4) COMMENT '下单后台调整订单使用的折扣总金额', 
  `refund_payment_cnt` bigint COMMENT '被退款次数', 
  `refund_payment_num` bigint  COMMENT '被退款件数', 
  `refund_payment_amount` decimal(20,4) COMMENT '被退款金额', 
  `browse_cnt` bigint COMMENT '商品浏览次数', 
  `collection_cnt` bigint COMMENT '商品收藏次数', 
  `shopping_cart_cnt` bigint COMMENT '商品加入购物车次数'
  )COMMENT "商品统计汇总表"
PARTITIONED BY (dt string comment '日期')
stored as orc;


set hive.exec.dynamic.partition.mode=nonstrict;

with order_refund as (
  select
    b.spu_id,
    count(1) as refund_payment_cnt,
    sum(b.sku_quantity) as refund_payment_num,
    sum(refund) as refund_payment_amount,
    a.dt
  from
    dwd_commerce.dwd_trade_order_refund_info_d a
    left join dwd_commerce.dwd_trade_order_item_detail_d b on b.order_id = a.order_return_id
  group by
    b.spu_id,a.dt
),
order_info as (
  select
    spu_id,
    count(order_id) as order_cnt,
    sum(sku_quantity) as order_num,
    sum(if(coupon_id is not null, 1, 0)) as order_coupon_cnt,
    sum(total_amount) as order_total_amount,
    sum(pay_amount) as order_pay_amount,
    sum(freight_amount) as order_freight_amount,
    sum(promotion_amount) as order_promotion_amount,
    sum(integration_amount) as order_integration_amount,
    sum(coupon_amount) as order_coupon_amount,
    sum(discount_amount) as order_discount_amount,
    dt
  from
    dwd_commerce.dwd_trade_order_item_detail_d a
  group by
    spu_id,dt
),
user_behavior as (
  select
    cast(item_id as bigint) as spu_id,
    sum(if(type = 1, 1, 0)) as browse_cnt,
    sum(if(type = 2, 1, 0)) as collection_cnt,
    sum(if(type = 3, 1, 0)) as shopping_cart_cnt,
    dt
  from
    dwd_commerce.dwd_user_behavior_detail_d
  where
    type in (1, 2, 3)
  group by
    item_id,dt
),
union_all_table as (
select
  spu_id,
  order_cnt,
  order_num,
  order_coupon_cnt,
  order_total_amount,
  order_pay_amount,
  order_freight_amount,
  order_promotion_amount,
  order_integration_amount,
  order_coupon_amount,
  order_discount_amount,
  0 as refund_payment_cnt,
  0 as refund_payment_num,
  0 as refund_payment_amount,
  0 as browse_cnt,
  0 as collection_cnt,
  0 as shopping_cart_cnt,
  dt
from
  order_info
union all
select
  spu_id,
  0 as order_cnt,
  0 as order_num,
  0 as order_coupon_cnt,
  0 as order_total_amount,
  0 as order_pay_amount,
  0 as order_freight_amount,
  0 as order_promotion_amount,
  0 as order_integration_amount,
  0 as order_coupon_amount,
  0 as order_discount_amount,
  refund_payment_cnt,
  refund_payment_num,
  refund_payment_amount,
  0 as browse_cnt,
  0 as collection_cnt,
  0 as shopping_cart_cnt,
  dt
from
  order_refund
union all
select
  spu_id,
  0 as order_cnt,
  0 as order_num,
  0 as order_coupon_cnt,
  0 as order_total_amount,
  0 as order_pay_amount,
  0 as order_freight_amount,
  0 as order_promotion_amount,
  0 as order_integration_amount,
  0 as order_coupon_amount,
  0 as order_discount_amount,
  0 as refund_payment_cnt,
  0 as refund_payment_num,
  0 as refund_payment_amount,
  browse_cnt,
  collection_cnt,
  shopping_cart_cnt,
  dt
from
  user_behavior
),
sum_table as (
select
  spu_id,
  sum(order_cnt)                as order_cnt,
  sum(order_num)                as order_num,
  sum(order_coupon_cnt)         as order_coupon_cnt,
  sum(order_total_amount)       as order_total_amount,
  sum(order_pay_amount)         as order_pay_amount,
  sum(order_freight_amount)     as order_freight_amount,
  sum(order_promotion_amount)   as order_promotion_amount,
  sum(order_integration_amount) as order_integration_amount,
  sum(order_coupon_amount)      as order_coupon_amount,
  sum(order_discount_amount)    as order_discount_amount,
  sum(refund_payment_cnt)       as refund_payment_cnt,
  sum(refund_payment_num)       as refund_payment_num,
  sum(refund_payment_amount)    as refund_payment_amount,
  sum(browse_cnt)               as browse_cnt,
  sum(collection_cnt)           as collection_cnt,
  sum(shopping_cart_cnt)        as shopping_cart_cnt,
  dt
from union_all_table
group by spu_id,dt
)
insert overwrite table  dws_commerce.dws_spu_summary_d partition(dt)
select
spu_id,
order_cnt,
order_num,
order_coupon_cnt,
order_total_amount,
order_pay_amount,
order_freight_amount,
order_promotion_amount,
order_integration_amount,
order_coupon_amount,
order_discount_amount,
refund_payment_cnt,
refund_payment_num,
refund_payment_amount,
browse_cnt,
collection_cnt,
shopping_cart_cnt,
dt
from sum_table
;

-- 数据验证 ok 
select * from dws_commerce.dws_spu_summary_d limit 10;


-- 2.dws 用户交易行为汇总表
DROP TABLE IF EXISTS dws_commerce.dws_user_summary_d;
CREATE TABLE IF NOT EXISTS dws_commerce.dws_user_summary_d(
  `user_id` bigint COMMENT 'member_id', 
  `browse_cnt` bigint COMMENT '商品浏览次数', 
  `collection_cnt` bigint COMMENT '商品收藏次数', 
  `shopping_cart_cnt` bigint COMMENT '商品加入购物车次数',
  `purchase_cnt` bigint COMMENT '商品购买次数',
  `order_cnt` bigint  COMMENT '下单次数', 
  `order_coupon_cnt` bigint COMMENT '使用优惠券下单次数', 
  `order_total_amount` decimal(20,4) COMMENT '下单订单总额', 
  `order_pay_amount` decimal(20,4) COMMENT '下单应付总额', 
  `order_freight_amount` decimal(20,4) COMMENT '下单运费金额', 
  `order_promotion_amount` decimal(20,4) COMMENT '下单促销优化总金额（促销价、满减、阶梯价）', 
  `order_integration_amount` decimal(20,4) COMMENT '下单积分抵扣总金额', 
  `order_coupon_amount` decimal(20,4) COMMENT '下单优惠券抵扣总金额', 
  `order_discount_amount` decimal(20,4) COMMENT '下单后台调整订单使用的折扣总金额', 
  `refund_payment_cnt` bigint COMMENT '被退款次数', 
  `refund_payment_num` bigint  COMMENT '被退款件数', 
  `refund_payment_amount` decimal(20,4) COMMENT '被退款金额'
  ) COMMENT '用户交易行为汇总表'
PARTITIONED BY (dt string comment '日期')
stored as orc;

-- 核心代码
set hive.exec.dynamic.partition.mode=nonstrict;
with user_behavior as (
  select
    cast(user_id as bigint) as user_id,
    count(1) as login_cnt,
    sum(if(type = 1, 1, 0)) as browse_cnt,
    sum(if(type = 2, 1, 0)) as collection_cnt,
    sum(if(type = 3, 1, 0)) as shopping_cart_cnt,
    sum(if(type = 4, 1, 0)) as purchase_cnt,
    dt
  from
    dwd_commerce.dwd_user_behavior_detail_d
  group by
    user_id,dt
),
user_order as (
  select
    cast(user_id as bigint) as user_id,
    count(order_id) as order_cnt,
    sum(if(coupon_id is not null, 1, 0)) as order_coupon_cnt,
    sum(total_amount) as order_total_amount,
    sum(pay_amount) as order_pay_amount,
    sum(freight_amount) as order_freight_amount,
    sum(promotion_amount) as order_promotion_amount,
    sum(integration_amount) as order_integration_amount,
    sum(coupon_amount) as order_coupon_amount,
    sum(discount_amount) as order_discount_amount,
    dt
  from
    dwd_commerce.dwd_trade_order_item_detail_d a
  group by
    user_id,dt
),
user_order_refund as (
  select
    a.user_id,
    count(1) as refund_payment_cnt,
    sum(b.sku_quantity) as refund_payment_num,
    sum(refund) as refund_payment_amount,
    a.dt
  from
    dwd_commerce.dwd_trade_order_refund_info_d a
  left join dwd_commerce.dwd_trade_order_item_detail_d b on b.order_id = a.order_return_id
  group by
  a.user_id,a.dt
),
union_all_table as (
select
  user_id,
  0 as browse_cnt,
  0 as collection_cnt,
  0 as shopping_cart_cnt,
  0 as purchase_cnt,
  order_cnt,
  order_coupon_cnt,
  order_total_amount,
  order_pay_amount,
  order_freight_amount,
  order_promotion_amount,
  order_integration_amount,
  order_coupon_amount,
  order_discount_amount,
  0 as refund_payment_cnt,
  0 as refund_payment_num,
  0 as refund_payment_amount,
  dt
from
  user_order
union all
select
  user_id,
  browse_cnt,
  collection_cnt,
  shopping_cart_cnt,
  purchase_cnt,
  0 as order_cnt,
  0 as order_coupon_cnt,
  0 as order_total_amount,
  0 as order_pay_amount,
  0 as order_freight_amount,
  0 as order_promotion_amount,
  0 as order_integration_amount,
  0 as order_coupon_amount,
  0 as order_discount_amount,
  0 as refund_payment_cnt,
  0 as refund_payment_num,
  0 as refund_payment_amount,
  dt
from
  user_behavior
union all
select
  user_id,
  0 as browse_cnt,
  0 as collection_cnt,
  0 as shopping_cart_cnt,
  0 as purchase_cnt,
  0 as order_cnt,
  0 as order_coupon_cnt,
  0 as order_total_amount,
  0 as order_pay_amount,
  0 as order_freight_amount,
  0 as order_promotion_amount,
  0 as order_integration_amount,
  0 as order_coupon_amount,
  0 as order_discount_amount,
  refund_payment_cnt,
  refund_payment_num,
  refund_payment_amount,
  dt
from
  user_order_refund
),
sum_table as (
select
  user_id,
  sum(browse_cnt)               as browse_cnt,
  sum(collection_cnt)           as collection_cnt,
  sum(shopping_cart_cnt)        as shopping_cart_cnt,
  sum(purchase_cnt)             as purchase_cnt,
  sum(order_cnt)                as order_cnt,
  sum(order_coupon_cnt)         as order_coupon_cnt,
  sum(order_total_amount)       as order_total_amount,
  sum(order_pay_amount)         as order_pay_amount,
  sum(order_freight_amount)     as order_freight_amount,
  sum(order_promotion_amount)   as order_promotion_amount,
  sum(order_integration_amount) as order_integration_amount,
  sum(order_coupon_amount)      as order_coupon_amount,
  sum(order_discount_amount)    as order_discount_amount,
  sum(refund_payment_cnt)       as refund_payment_cnt,
  sum(refund_payment_num)       as refund_payment_num,
  sum(refund_payment_amount)    as refund_payment_amount,
  dt
from union_all_table
group by user_id,dt
)
insert overwrite table  dws_commerce.dws_user_summary_d partition(dt)
select 
    user_id,
    browse_cnt,
    collection_cnt,
    shopping_cart_cnt,
    purchase_cnt,
    order_cnt,
    order_coupon_cnt,
    order_total_amount,
    order_pay_amount,
    order_freight_amount,
    order_promotion_amount,
    order_integration_amount,
    order_coupon_amount,
    order_discount_amount,
    refund_payment_cnt,
    refund_payment_num,
    refund_payment_amount,
    dt
from sum_table
;


-- 数据验证 ok 有数据、待优化
select * from dws_commerce.dws_user_summary_d limit 10;


-- 3.用户交易地区汇总表
DROP TABLE IF EXISTS dws_commerce.dws_region_summary_d;
CREATE TABLE IF NOT EXISTS dws_commerce.dws_region_summary_d(
  `province` string  COMMENT '省份/直辖市', 
  `city` string  COMMENT '城市', 
  `region` string  COMMENT '区', 
  `order_cnt` bigint  COMMENT '下单次数', 
  `order_num` bigint COMMENT '下单件数', 
  `order_coupon_cnt` bigint COMMENT '使用优惠券下单次数', 
  `order_total_amount` decimal(20,4) COMMENT '下单订单总额', 
  `order_pay_amount` decimal(20,4) COMMENT '下单应付总额', 
  `order_freight_amount` decimal(20,4) COMMENT '下单运费金额', 
  `order_promotion_amount` decimal(20,4) COMMENT '下单促销优化总金额（促销价、满减、阶梯价）', 
  `order_integration_amount` decimal(20,4) COMMENT '下单积分抵扣总金额', 
  `order_coupon_amount` decimal(20,4) COMMENT '下单优惠券抵扣总金额', 
  `order_discount_amount` decimal(20,4) COMMENT '下单后台调整订单使用的折扣总金额', 
  `refund_payment_cnt` bigint COMMENT '被退款次数', 
  `refund_payment_num` bigint  COMMENT '被退款件数', 
  `refund_payment_amount` decimal(20,4) COMMENT '被退款金额',
  `browse_cnt` bigint COMMENT '商品浏览次数', 
  `collection_cnt` bigint COMMENT '商品收藏次数', 
  `shopping_cart_cnt` bigint COMMENT '商品加入购物车次数',
  `purchase_cnt` bigint COMMENT '商品购买次数'
  )COMMENT '用户交易地区汇总表'
PARTITIONED BY (dt string comment '日期')
stored as orc;


-- 核心代码:该代码需要运行很久
set hive.exec.dynamic.partition.mode=nonstrict;
with order_refund as (
  select
    a.receiver_province,
    a.receiver_city,    
    a.receiver_region,  
    count(1) as refund_payment_cnt,
    sum(b.sku_quantity) as refund_payment_num,
    sum(refund) as refund_payment_amount,
    b.dt
  from
    dwd_commerce.dwd_trade_order_refund_info_d a
    left join dwd_commerce.dwd_trade_order_item_detail_d b on b.order_id = a.order_return_id
  group by
    a.receiver_province,
    a.receiver_city,    
    a.receiver_region,
    b.dt 
),
order_info as (
  select
    receiver_province,
    receiver_city,    
    receiver_region,
    count(order_id) as order_cnt,
    sum(sku_quantity) as order_num,
    sum(if(coupon_id is not null, 1, 0)) as order_coupon_cnt,
    sum(total_amount) as order_total_amount,
    sum(pay_amount) as order_pay_amount,
    sum(freight_amount) as order_freight_amount,
    sum(promotion_amount) as order_promotion_amount,
    sum(integration_amount) as order_integration_amount,
    sum(coupon_amount) as order_coupon_amount,
    sum(discount_amount) as order_discount_amount,
    dt
  from
    dwd_commerce.dwd_trade_order_item_detail_d a
  group by
    receiver_province,
    receiver_city,
    receiver_region,
    dt
),
--应该需要做个窗口函数去重
user_address as (
select
user_id,
receiver_province,
receiver_city,
receiver_region,
dt
from
dwd_commerce.dwd_trade_order_info_detail_d
group by
user_id,
receiver_province,
receiver_city,
receiver_region,
dt
),
user_behavior as (
  select
    receiver_province,
    receiver_city,
    receiver_region,
    sum(if(type = 1, 1, 0)) as browse_cnt,
    sum(if(type = 2, 1, 0)) as collection_cnt,
    sum(if(type = 3, 1, 0)) as shopping_cart_cnt,
    sum(if(type = 4, 1, 0)) as purchase_cnt,
    a.dt
  from
    dwd_commerce.dwd_user_behavior_detail_d a
   left join user_address b
   on cast(a.user_id as bigint) = b.user_id
  group by
    receiver_province,
    receiver_city,
    receiver_region,
    a.dt
),
union_all_table as (
select
  receiver_province,
  receiver_city,    
  receiver_region,
  order_cnt,
  order_num,
  order_coupon_cnt,
  order_total_amount,
  order_pay_amount,
  order_freight_amount,
  order_promotion_amount,
  order_integration_amount,
  order_coupon_amount,
  order_discount_amount,
  0 as refund_payment_cnt,
  0 as refund_payment_num,
  0 as refund_payment_amount,
  0 as browse_cnt,
  0 as collection_cnt,
  0 as shopping_cart_cnt,
  0 as purchase_cnt,
  dt
from
  order_info
union all
select
  receiver_province,
  receiver_city,    
  receiver_region,
  0 as order_cnt,
  0 as order_num,
  0 as order_coupon_cnt,
  0 as order_total_amount,
  0 as order_pay_amount,
  0 as order_freight_amount,
  0 as order_promotion_amount,
  0 as order_integration_amount,
  0 as order_coupon_amount,
  0 as order_discount_amount,
  refund_payment_cnt,
  refund_payment_num,
  refund_payment_amount,
  0 as browse_cnt,
  0 as collection_cnt,
  0 as shopping_cart_cnt,
  0 as purchase_cnt,
  dt
from
  order_refund
union all
select
  receiver_province,
  receiver_city,    
  receiver_region,
  0 as order_cnt,
  0 as order_num,
  0 as order_coupon_cnt,
  0 as order_total_amount,
  0 as order_pay_amount,
  0 as order_freight_amount,
  0 as order_promotion_amount,
  0 as order_integration_amount,
  0 as order_coupon_amount,
  0 as order_discount_amount,
  0 as refund_payment_cnt,
  0 as refund_payment_num,
  0 as refund_payment_amount,
  browse_cnt,
  collection_cnt,
  shopping_cart_cnt,
  purchase_cnt,
  dt
from
  user_behavior
),
sum_table as (
select
  receiver_province,
  receiver_city,    
  receiver_region,
  sum(order_cnt)                as order_cnt,
  sum(order_num)                as order_num,
  sum(order_coupon_cnt)         as order_coupon_cnt,
  sum(order_total_amount)       as order_total_amount,
  sum(order_pay_amount)         as order_pay_amount,
  sum(order_freight_amount)     as order_freight_amount,
  sum(order_promotion_amount)   as order_promotion_amount,
  sum(order_integration_amount) as order_integration_amount,
  sum(order_coupon_amount)      as order_coupon_amount,
  sum(order_discount_amount)    as order_discount_amount,
  sum(refund_payment_cnt)       as refund_payment_cnt,
  sum(refund_payment_num)       as refund_payment_num,
  sum(refund_payment_amount)    as refund_payment_amount,
  sum(browse_cnt)               as browse_cnt,
  sum(collection_cnt)           as collection_cnt,
  sum(shopping_cart_cnt)        as shopping_cart_cnt,
  sum(purchase_cnt)             as purchase_cnt,
  dt
from union_all_table
group by   
  receiver_province,
  receiver_city,    
  receiver_region,
  dt
)
insert overwrite table  dws_commerce.dws_region_summary_d partition(dt)
select
    receiver_province as province,
    receiver_city as city,    
    receiver_region as region,
    order_cnt,
    order_num,
    order_coupon_cnt,
    order_total_amount,
    order_pay_amount,
    order_freight_amount,
    order_promotion_amount,
    order_integration_amount,
    order_coupon_amount,
    order_discount_amount,
    refund_payment_cnt,
    refund_payment_num,
    refund_payment_amount,
    browse_cnt,
    collection_cnt,
    shopping_cart_cnt,
    purchase_cnt,
    dt
from sum_table
;

-- 数据校验：OK 待优化
select * from dws_commerce.dws_region_summary_d limit 10;



-------------------------------------4.ads层数据--------------------------------------------------------------------
--part1:不同时间段内的订单汇总统计
drop table if exists ads_commerce.ads_order_sum_full;
create table ads_commerce.ads_order_sum_full 
stored as orc
  as
select
  recent_days,
  sum(order_cnt) as order_cnt,
  sum(order_total_amount) as order_total_amount,
  sum(order_user_cnt) as order_user_cnt
from
  (
    select
      1 as recent_days,
      sum(order_cnt) as order_cnt,
      sum(order_total_amount) as order_total_amount,
      sum(if(order_total_amount>0,1,0)) as order_user_cnt
    from
      dws_commerce.dws_spu_summary_d
    where dt between '2019-10-31' and '2019-10-31' group by 1
    union all
    select
      3 as recent_days,
      sum(order_cnt) as order_cnt,
      sum(order_total_amount) as order_total_amount,
      sum(if(order_total_amount>0,1,0)) as order_user_cnt
    from
      dws_commerce.dws_spu_summary_d
     where dt between '2019-10-29' and '2019-10-31'  group by 1
    union all
    select
      7 as recent_days,
      sum(order_cnt) as order_cnt,
      sum(order_total_amount) as order_total_amount,
      sum(if(order_total_amount>0,1,0)) as order_user_cnt
    from
      dws_commerce.dws_spu_summary_d
    where dt between '2019-10-25' and '2019-10-31'  group by 1
  ) as tmp
group by
  recent_days;

-- 数据验证：OK
select * from ads_commerce.ads_order_sum_full limit 10;  



--part2:不同时间段的各地区订单统计
drop table if exists ads_commerce.ads_province_order_sum_full;
create table ads_commerce.ads_province_order_sum_full 
stored as orc
  as
select
  recent_days,
  province,
  sum(order_cnt) as order_cnt,
  sum(order_total_amount) as order_total_amount,
  sum(order_user_cnt) as order_user_cnt
from
  (
    select
      1 as recent_days,
      province,
      sum(order_cnt) as order_cnt,
      sum(order_total_amount) as order_total_amount,
      sum(if(order_total_amount>0,1,0)) as order_user_cnt
    from
      dws_commerce.dws_region_summary_d
      where dt between '2019-10-31' and '2019-10-31'
      group by province
    union all
    select
      3 as recent_days,
      province,
      sum(order_cnt) as order_cnt,
      sum(order_total_amount) as order_total_amount,
      sum(if(order_total_amount>0,1,0)) as order_user_cnt
    from
      dws_commerce.dws_region_summary_d
      where dt between '2019-10-29' and '2019-10-31'
      group by province
    union all
    select
      7 as recent_days,
      province,
      sum(order_cnt) as order_cnt,
      sum(order_total_amount) as order_total_amount,
      sum(if(order_total_amount>0,1,0)) as order_user_cnt
    from
      dws_commerce.dws_region_summary_d
    where dt between '2019-10-25' and '2019-10-31'
      group by province
  ) as tmp
group by
  recent_days,
  province;


-- 数据验证：OK
select * from ads_commerce.ads_province_order_sum_full limit 10;


ods层2个表，dwd层3个表，dws层3个表,ads层2个表