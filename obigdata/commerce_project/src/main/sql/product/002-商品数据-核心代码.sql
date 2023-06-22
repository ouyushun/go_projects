-------------------------------------0.ods层数据--------------------------------------------------------------------
//SPU(Standard Product Unit)—标准化产品单元
//SKU(stock keeping unit)——SKU 即库存进出计量的单位，
DROP TABLE IF EXISTS ods_product_sku;
CREATE TABLE IF NOT EXISTS ods_product_sku (
  id BIGINT COMMENT 'skuId',
  spu_id BIGINT COMMENT 'spuId',
  name STRING COMMENT 'sku名称',
  catagory_id BIGINT COMMENT '所属分类id',
  brand_id BIGINT COMMENT '品牌id',
  default_image STRING COMMENT '默认图片',
  title STRING COMMENT '标题',
  subtitle STRING COMMENT '副标题',
  price DOUBLE COMMENT '价格',
  weight INT COMMENT '重量（克）'
) COMMENT 'sku信息' STORED AS orc;

DROP TABLE IF EXISTS ods_product_sku_attr_value;
CREATE TABLE IF NOT EXISTS ods_product_sku_attr_value (
  id BIGINT COMMENT 'id',
  sku_id BIGINT COMMENT 'sku_id',
  attr_id BIGINT COMMENT 'attr_id',
  attr_name STRING COMMENT '销售属性名',
  attr_value STRING COMMENT '销售属性值',
  sort INT COMMENT '顺序'
) COMMENT 'sku销售属性&值' STORED AS orc;

DROP TABLE IF EXISTS ods_product_spu;
CREATE TABLE IF NOT EXISTS ods_product_spu (
  id BIGINT COMMENT '商品id',
  name STRING COMMENT '商品名称',
  category_id BIGINT COMMENT '所属分类id',
  brand_id BIGINT COMMENT '品牌id',
  publish_status INT COMMENT '上架状态[0 - 下架，1 - 上架]',
  create_time STRING COMMENT '创建时间',
  update_time STRING COMMENT '更新时间'
) COMMENT 'spu信息' STORED AS orc;

DROP TABLE IF EXISTS ods_product_spu_attr_value;
CREATE TABLE IF NOT EXISTS ods_product_spu_attr_value (
  id BIGINT COMMENT 'id',
  spu_id BIGINT COMMENT '商品id',
  attr_id BIGINT COMMENT '属性id',
  attr_name STRING COMMENT '属性名',
  attr_value STRING COMMENT '属性值',
  sort INT COMMENT '顺序'
) COMMENT 'spu属性值' STORED AS orc;

DROP TABLE IF EXISTS ods_product_brand;
CREATE TABLE IF NOT EXISTS ods_product_brand (
  id BIGINT COMMENT '品牌id',
  name STRING COMMENT '品牌名',
  logo STRING COMMENT '品牌logo',
  status INT COMMENT '显示状态[0-不显示；1-显示]',
  first_letter STRING COMMENT '检索首字母',
  sort INT COMMENT '排序',
  remark string COMMENT '备注'
) COMMENT '品牌'  STORED AS orc;


DROP TABLE IF EXISTS ods_product_category;
CREATE TABLE IF NOT EXISTS ods_product_category (
  id BIGINT COMMENT '分类id',
  name STRING COMMENT '分类名称',
  parent_id BIGINT COMMENT '父分类id',
  status INT COMMENT '是否显示[0-不显示，1显示]',
  sort INT COMMENT '排序',
  icon STRING COMMENT '图标地址',
  unit STRING COMMENT '计量单位'
) COMMENT '商品三级分类' STORED AS orc;

DROP TABLE IF EXISTS ods_regioninfo;
CREATE TABLE IF NOT EXISTS ods_regioninfo (
  regionid STRING COMMENT '地区ID',
  parentid STRING COMMENT '父级区域ID',
  regionname STRING COMMENT '地区名称',
  regiontype INT COMMENT '区域类别（0国家/1省份/2城市/3区县）',
  agencyid INT COMMENT '无用字段',
  pt STRING COMMENT '系统更新时间'
) COMMENT '行政区划表'  STORED AS orc;

-------------------------------------1.dwd层数据--------------------------------------------------------------------
--part1:.sku 商品维度表
DROP TABLE IF EXISTS dim_commerce.dim_sku_detailed_info_full;
CREATE TABLE IF NOT EXISTS `dim_commerce.dim_sku_detailed_info_full`(
  `sku_id` bigint COMMENT '商品id', 
  `sku_name` string COMMENT '商品名称', 
  `catagory_id` bigint COMMENT '所属三级分类id', 
  `catagory_name` string COMMENT '所属三级分类名字', 
  `brand_id` bigint COMMENT '品牌id', 
  `brand_name` string COMMENT '品牌名称', 
  `sku_default_image` string COMMENT '默认图片', 
  `sku_title` string COMMENT '标题',  
  `sku_subtitle` string COMMENT '副标题', 
  `sku_price` double COMMENT '价格', 
  `spu_id` bigint COMMENT 'spu商品id', 
  `spu_name` string COMMENT 'spu商品名称', 
  `sku_attrs` array<struct<attr_id:bigint,attr_name:string,attr_value:string>> COMMENT 'sku 平台属性' 
  ) COMMENT 'sku 商品维度表'
STORED AS orc;


--核心代码
with product_sku as (
  select
    id,
    spu_id,
    name,
    catagory_id,
    brand_id,
    default_image,
    title,
    subtitle,
    price,
    weight
  from
    ods_commerce.ods_product_sku
),
product_spu as (
  select
    id as spu_id,
    name as spu_name
  from
    ods_commerce.ods_product_spu
),
product_category as (
  select
    id as category_id,
    name as catagory_name
  from
    ods_commerce.ods_product_category
),
product_brand as (
  select
    id as brand_id,
    name as brand_name
  from
    ods_commerce.ods_product_brand
),
sku_attr_value as (
  select
    sku_id,
    collect_set(named_struct('attr_id',attr_id, 'attr_name',attr_name, 'attr_value',attr_value)) as sku_attrs
  from
    ods_commerce.ods_product_sku_attr_value
  group by
    sku_id
)
insert overwrite table dim_commerce.dim_sku_detailed_info_full
select
  id as sku_id,
  name as sku_name,
  product_sku.catagory_id,
  product_category.catagory_name,
  product_sku.brand_id,
  product_brand.brand_name,
  default_image as sku_default_image,
  title as sku_title,
  subtitle as sku_subtitle,
  price as sku_price,
  product_sku.spu_id,
  product_spu.spu_name,
  sku_attr_value.sku_attrs
from
  product_sku
  left join product_spu on product_sku.spu_id = product_spu.spu_id
  left join product_category on product_sku.catagory_id = product_category.category_id
  left join product_brand on product_sku.brand_id = product_brand.brand_id
  left join sku_attr_value on product_sku.id = sku_attr_value.sku_id
;


select * from dim_commerce.dim_sku_detailed_info_full limit 10

-- 知识点的回顾与应用
-- 1.函数collect_set、named_struct的使用
-- 2.多表关联的写法




-- part2: spu 商品维度表
DROP TABLE IF EXISTS dim_commerce.dim_spu_detailed_info_full;
CREATE TABLE IF NOT EXISTS `dim_commerce.dim_spu_detailed_info_full`(
  `spu_id` bigint COMMENT 'spu商品id', 
  `spu_name` string COMMENT 'spu商品名称', 
  `spu_publish_status` int COMMENT '上架状态', 
  `spu_create_time` string COMMENT '创建时间', 
  `spu_update_time` string COMMENT '更新时间', 
  `category_id` bigint COMMENT '所属三级分类id', 
  `category_name` string COMMENT '所属三级分类名字', 
  `brand_id` bigint COMMENT '品牌id', 
  `brand_name` string COMMENT '品牌名称', 
  `spu_attrs` array<struct<col1:bigint,col2:string,col3:string>>COMMENT 'spu 平台属性' 
  ) COMMENT 'spu 商品维度表'
STORED AS orc;


-- 核心代码
with product_spu as (
  select
    id as spu_id,
    name as spu_name,
    category_id,
    brand_id,
    publish_status as spu_publish_status,
    create_time as spu_create_time,
    update_time as spu_update_time
  from
    ods_commerce.ods_product_spu
),
product_category as (
  select
    id as category_id,
    name as category_name
  from
    ods_commerce.ods_product_category
),
product_brand as (
  select
    id as brand_id,
    name as brand_name
  from
    ods_commerce.ods_product_brand
),
spu_attr_value as (
  select
    spu_id,
    collect_set(struct(attr_id, attr_name, attr_value)) as spu_attrs
  from
    ods_commerce.ods_product_spu_attr_value
  group by
    spu_id
)
insert overwrite table dim_commerce.dim_spu_detailed_info_full
select
  product_spu.spu_id,
  spu_name,
  spu_publish_status,
  spu_create_time,
  spu_update_time,
  product_spu.category_id,
  product_category.category_name,
  product_spu.brand_id,
  product_brand.brand_name,
  spu_attr_value.spu_attrs
from
  product_spu
  left join product_category on product_spu.category_id = product_category.category_id
  left join product_brand on product_spu.brand_id = product_brand.brand_id
  left join spu_attr_value on product_spu.spu_id = spu_attr_value.spu_id
;


select * from dim_commerce.dim_spu_detailed_info_full limit 10

-- 知识点的回顾与应用
-- 1.函数collect_set、named_struct的使用
-- 2.多表关联的写法


-- part3：品牌维度表
DROP TABLE IF EXISTS dim_commerce.dim_brand_detailed_info_full;
CREATE TABLE IF NOT EXISTS `dim_commerce.dim_brand_detailed_info_full`(
  `id` bigint COMMENT '品牌id', 
  `name` string COMMENT '品牌名称', 
  `logo` string COMMENT '品牌logo', 
  `update_time` string COMMENT '更新时间'
  ) COMMENT '品牌维度表'
STORED AS orc;


-- 核心代码
insert overwrite table dim_commerce.dim_brand_detailed_info_full
select
  id,
  name,
  logo,
  '20220529' AS update_time
from
  ods_commerce.ods_product_brand
union all
select
  cast(supplierid as bigint) as id,
  brandtype as name,
  '' as logo,
  pt as update_time
from
  ods_commerce.ods_product_goodsbrand;


select * from dim_commerce.dim_brand_detailed_info_full limit 10

-- 知识点的回顾与应用
-- 1.函数cast的使用
-- 2.表拼接一般有两种方式，上下拼接（union all）和左右拼接（join），上下拼接在字段数不一致的情况下该怎么做呢？


-- part4:商品分类维度表
DROP TABLE IF EXISTS dim_commerce.dim_category_detailed_info_full;
CREATE TABLE IF NOT EXISTS `dim_commerce.dim_category_detailed_info_full`(
  `category_3_id` bigint COMMENT '三级分类id', 
  `category_3_name` string COMMENT '三级分类名称', 
  `category_2_id` bigint COMMENT '二级分类id', 
  `category_2_name` string COMMENT '二级分类名称', 
  `category_1_id` bigint COMMENT '一级分类id', 
  `category_1_name` string COMMENT '一级分类名称'
  ) COMMENT '商品分类维度表'
STORED AS orc;


--核心代码
insert overwrite table dim_commerce.dim_category_detailed_info_full
select
  a.id as category_3_id,
  a.name as category_3_name,
  a.parent_id as category_2_id,
  b.name as category_2_name,
  b.parent_id as category_1_id,
  c.name as category_1_name
from
  ods_commerce.ods_product_category a
  inner join ods_commerce.ods_product_category b on a.parent_id = b.id
  inner join ods_commerce.ods_product_category c on b.parent_id = c.id;

select * from dim_commerce.dim_category_detailed_info_full limit 10

-- 知识点的回顾与应用
-- 1.维度退化的使用，怎么从一张表内，拆出多个字段？



-- part5:goods维度表
DROP TABLE IF EXISTS dim_commerce.dim_goods_detailed_info_full;
CREATE TABLE IF NOT EXISTS `dim_commerce.dim_goods_detailed_info_full`(
  `goodsid` string COMMENT '商品ID', 
  `brand_id` string COMMENT '品类ID', 
  `markid` string COMMENT '专场ID（商品售卖的位置）', 
  `goodstag` string COMMENT '进货渠道，档口的名字', 
  `brand_name` string COMMENT '品牌名称（不用，脱敏严重）', 
  `customtag` string COMMENT '商品的详情', 
  `goodsname` string COMMENT '竞价排名，BH为公司的补货', 
  `clickcount` int COMMENT '商品的点击次数', 
  `clickcr` int COMMENT '-', 
  `goodsnumber` int COMMENT '货号', 
  `goodsweight` int COMMENT '商品重量', 
  `marketprice` double COMMENT '进货价，成本', 
  `shopprice` double COMMENT '售价', 
  `addtime` string COMMENT '新款建档时间，在数据库里', 
  `isonsale` int COMMENT '是否在售（1在售，0否）', 
  `sales` int COMMENT '真实的销量+刷单的销量', 
  `realsales` int COMMENT '实际销量', 
  `extraprice` double COMMENT '特别价格（促销价）', 
  `goodsno` string COMMENT '货号ID，一个商品ID可能对应多个货号ID', 
  `update_time` string COMMENT '更新时间', 
  `category_3_id` int COMMENT '三级分类id', 
  `category_3_name` string COMMENT '三级分类名称', 
  `category_2_id` int COMMENT '二级分类id', 
  `category_2_name` string COMMENT '二级分类名称', 
  `category_1_id` int COMMENT '一级分类id', 
  `category_1_name` string COMMENT '一级分类名称'
  ) COMMENT 'goods维度表'
STORED AS orc;

-- 核心代码
spark.sql("""
insert overwrite table dim_commerce.dim_goods_detailed_info_full
select
  goodsid,
  typeid as brand_id,
  markid,
  goodstag,
  brandtag as brand_name,
  customtag,
  goodsname,
  clickcount,
  clickcr,
  goodsnumber,
  goodsweight,
  marketprice,
  shopprice,
  addtime,
  isonsale,
  sales,
  realsales,
  extraprice,
  goodsno,
  dt as update_time,
  nvl(category_base.category_3_id, 650) as category_3_id,
  nvl(category_base.category_3_name, 'T恤') as category_3_name,
  nvl(category_base.category_2_id, 76) as category_2_id,
  nvl(category_base.category_2_name, '女装') as category_2_name,
  nvl(category_base.category_1_id, 9) as category_1_id,
  nvl(category_base.category_1_name, '服饰内衣') as category_1_name
from
  ods_commerce.ods_product_goodsinfo2 goods
  left join dim_commerce.dim_category_detailed_info_full category_base on goods.goodsname = category_base.category_3_name
""")

select * from dim_commerce.dim_goods_detailed_info_full limit 10;


-- 知识点的回顾与应用
-- 1.hive 不支持不等值连接，需要用Spark SQL进行跑数，思考：什么是不等值连接？
-- 2.函数nvl、rlike的使用
-- 3.注意这里用的是  ods_commerce.ods_product_goodsinfo2


-- 6.地区维度表
DROP TABLE IF EXISTS dim_commerce.dim_region_info_full;
CREATE TABLE IF NOT EXISTS `dim_commerce.dim_region_info_full`(
  `county_id` string COMMENT '区县id', 
  `county_name` string COMMENT '区县名称' , 
  `city_id` string COMMENT '城市id' , 
  `city_name` string COMMENT '城市名称' , 
  `province_id` string COMMENT '省份id' , 
  `province_name`  string COMMENT '省份名称' , 
  `country_id`  string COMMENT '国家id', 
  `country_name` string  COMMENT '国家名称',
  `update_time` string COMMENT '更新时间' 
  )COMMENT '地区维度表'
STORED AS orc;

--核心代码
insert overwrite table dim_commerce.dim_region_info_full
select
  a.regionid as county_id,
  a.regionname as county_name,
  b.regionid as city_id,
  b.regionname as city_name,
  c.regionid as province_id,
  c.regionname as province_name,
  d.regionid as country_id,
  d.regionname as country_name,
  a.pt as update_time
from
  ods_commerce.ods_regioninfo a
  inner join ods_commerce.ods_regioninfo b on a.parentid = b.regionid
  inner join ods_commerce.ods_regioninfo c on b.parentid = c.regionid
  inner join ods_commerce.ods_regioninfo d on c.parentid = d.regionid
  ;


select * from dim_commerce.dim_region_info_full
 -- 知识点的回顾与应用
 -- 将维度的属性层次合并到单个维度中的操作称为反规范化。

select * from dim_commerce.dim_sku_detailed_info_full limit 10
select * from dim_commerce.dim_spu_detailed_info_full limit 10
select * from dim_commerce.dim_brand_detailed_info_full limit 10
select * from dim_commerce.dim_category_detailed_info_full limit 10
select * from dim_commerce.dim_goods_detailed_info_full limit 10
select * from dim_commerce.dim_region_info_full limit 10
  