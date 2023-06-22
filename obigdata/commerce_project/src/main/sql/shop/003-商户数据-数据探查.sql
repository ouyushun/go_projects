1.查看商品表的数据
select * from ods_commerce.ods_product_goodsinfo

2.查看商品是否存在重复
select goodsid,count(*) from ods_commerce.ods_product_goodsinfo group by goodsid having count(*)>1

3.对重复数据进行处理
select goodsid,max(goodsname) as goodsname 
from ods_commerce.ods_product_goodsinfo
where goodsname is not null
group by goodsid