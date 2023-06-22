1.创建SBT项目

2.Maven导包 -> sbt 导包???
https://mvnrepository.com/


3.在企业中,怎么用Scala代码开发数仓项目
数据接入
数据仓库
数据出仓


4.SQL

dwd_user_userinfo

CREATE TABLE IF NOT EXISTS dwd_commerce.dwd_user_userinfo (
   user_id STRING COMMENT '用户ID'
   ,user_name STRING COMMENT '用户名称'
   ,sex INT COMMENT '性别'
   ,user_money decimal(20,4) COMMENT '钱包'
   ,frozen_money decimal(20,4) COMMENT '近一个月的花费的总的金额'
   ,address_id STRING COMMENT '用户地址ID，0表示没有获取地址'
   ,reg_time STRING COMMENT '注册时间'
   ,last_login STRING COMMENT '最后登录时间'
) COMMENT '用户表' STORED AS ORC;


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

5.spark 打包 -> 运行
scp /Users/cpa/workspace/ds_dw/target/scala-2.11/ds_dw-assembly-0.1.jar ds_teacher@47.108.235.18:.

