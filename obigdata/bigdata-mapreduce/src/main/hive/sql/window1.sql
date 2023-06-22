create table test.user_log_1 (user_id string, log_in_date string) row format delimited fields terminated by ' ';
load data local inpath '/var/lib/hadoop-hdfs/data/user_log.txt' into table test.user_log_1;

hive> select * from test.user_log_1 ;
OK
u0001 2019-10-10
u0001 2019-10-11
u0001 2019-10-12
u0001 2019-10-14
u0001 2019-10-15
u0001 2019-10-17
u0001 2019-10-18
u0001 2019-10-19
u0001 2019-10-20
u0002 2019-10-20


--1. 第一步：排序
--按照user_id分组，并且按照日期log_in_date排序：
select user_id,
       log_in_date,
       row_number() over(partition by user_id order by log_in_date) as rank
from test.user_log_1;

--2. 第二步：第二列与第三列做日期差值
--可以看出规律，日期小的，行号也小；如果将日期跟行号做差值，连续登录的差值应该是一样的。
select user_id,
       date_sub(log_in_date, rank) dts
from (
    select user_id, log_in_date, row_number() over(partition by user_id order by log_in_date) as rank
    from test.user_log_1
    ) m;

--3. 第三步:按第二列分组求和
select user_id,
       dts,
       count(1) num
from (
    select user_id,
           date_sub(log_in_date,rank) dts
    from (
        select user_id, log_in_date,
               row_number() over(partition by user_id order by log_in_date) as rank
        from test.user_log_1
        )m
    )m2 group by user_id, dts;



--4. 第四步：求最大次数
--已经算出了，每个用户连续登录天数序列，接下取每个用户最大登录天数最大值即可：

select user_id, max(num)
from (
    select user_id, dts, count(1) num
    from (
        select user_id, date_sub(log_in_date, rank) dts
        from (
            select user_id, log_in_date, row_number() over(partition by user_id order by log_in_date) as rank
            from test.user_log_1
            )m
        )m2 group by user_id, dts
    )m3 group by user_id;


-- 1. 一天多次的记录去重
with t1 as (
        select  customer_key, substr(create_date, 1, 10) as create_date from ods_sales_orders group by customer_key, substr(create_date, 1, 10)
    ),
--分别计算， 4天后的日期 and 当前记录后4条的日期
t2 as (
         select customer_key, create_date,
                date_add(create_date, 3) as 4_days_later,
                lead(create_date, 3) over(partition by customer_key order by create_date asc) as next_4_record_date
         from t1
     ),
--获得连续4天购买的用户ID
t3 as (
         select customer_key from t2 where 4_days_later = next_4_record_date group by customer_key
     )
select customer_key,
       count(1) as purchases_num,
       sum(unit_price) as total_const,
       avg(unit_price) as avg_const
from  ods_sales_orders where customer_key in (select customer_key from t3) group by customer_key;
