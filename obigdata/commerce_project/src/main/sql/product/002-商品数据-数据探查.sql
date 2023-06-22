--1.查看用户量
select * from ods_commerce.ods_product_sku limit 10
select * from ods_commerce.ods_product_sku_attr_value limit 10
select * from ods_commerce.ods_product_spu limit 10
select * from ods_commerce.ods_product_spu_attr_value limit 10
select * from ods_commerce.ods_product_brand limit 10
select * from ods_commerce.ods_product_category limit 10
select * from ods_commerce.ods_regioninfo limit 10

desc ods_commerce.ods_product_sku;
desc ods_commerce.ods_product_sku_attr_value;

--2.查看行政区划表的数据
select regiontype,count(*) from ods_commerce.ods_regioninfo group by regiontype



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

select
  count(*)
from
  product_spu
  left join product_category on product_spu.category_id = product_category.category_id
  left join product_brand on product_spu.brand_id = product_brand.brand_id
;

select
  count(*)
from
    ods_commerce.ods_product_spu