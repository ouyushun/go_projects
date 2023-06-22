-- 重新计算dw层拉链表中的失效时间
select
    t1.goods_id,                -- 商品编号
    t1.goods_status,            -- 商品状态
    t1.createtime,              -- 商品创建时间
    t1.modifytime,              -- 商品修改时间
    t1.dw_start_date,           -- 生效日期（生效日期无需重新计算）
    case when (t2.goods_id is not null and t1.dw_end_date > '2019-12-21')
             then '2019-12-21'
         else t1.dw_end_date     -- 小的是以前修改的，不用修改，只修改9999-12-31的数据
        end as dw_end_date       -- 更新生效日期（需要重新计算）
from
    `demo`.`dw_product_2` t1
        left join
    (select * from `demo`.`ods_product_2` where dt='2019-12-21') t2
    on t1.goods_id = t2.goods_id



    insert overwrite table `demo`.`dw_product_2`
select
    t1.goods_id,                -- 商品编号
    t1.goods_status,            -- 商品状态
    t1.createtime,              -- 商品创建时间
    t1.modifytime,              -- 商品修改时间
    t1.dw_start_date,           -- 生效日期（生效日期无需重新计算）
    case when (t2.goods_id is not null and t1.dw_end_date > '2019-12-21')
             then '2019-12-21'
         else t1.dw_end_date
        end as dw_end_date       -- 更新生效日期（需要重新计算）
from
    `demo`.`dw_product_2` t1
        left join
    (select * from `demo`.`ods_product_2` where dt='2019-12-21') t2
    on t1.goods_id = t2.goods_id
union all
select
    goods_id,                -- 商品编号
    goods_status,            -- 商品状态
    createtime,              -- 商品创建时间
    modifytime,              -- 商品修改时间
    modifytime as dw_start_date,  -- 生效日期
    '9999-12-31' as dw_end_date   -- 失效日期
from
    `demo`.`ods_product_2` where dt='2019-12-21'    -- 只有新增和修改的数据
order by dw_start_date, goods_id;
