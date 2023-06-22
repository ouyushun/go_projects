-- noinspection SqlCurrentSchemaInspectionForFile

CREATE TABLE row2colum (
                          name string comment '用户姓名',
                          orderdate string comment '订单日期',
                          const int comment '花费金额')
    comment '用户消费表'
row format delimited fields terminated by ','
stored as textfile;
load data local inpath '/Users/ouyushun/hive/row2colum' overwrite into table row2colum;


show databases ;

use default;
use test1;

show tables;

select * from row2colum;

select * from business;




select * from row2colum;

select orderdate,
        max(case when name = 'jack' then const else 0 end) as jack,
        max(case when name = 'marry' then const else 0 end) as marry,
        collect_list(case when name = 'rachul' then const else 0 end) as rachul
from row2colum
group by orderdate;


jack,2001-01-01,100
jack,2002-01-01,200
jack,2003-01-01,300
jack,2004-01-01,400
jack,2005-01-01,500
marry,2001-01-01,100
marry,2002-01-01,200
marry,2003-01-01,300
marry,2004-01-01,400
marry,2005-01-01,500
rachul,2018-01-01,100
rachul,2019-01-01,200
rachul,2020-01-01,300
rachul,2021-01-01,400
rachul,2022-01-01,500
rachul,2022-01-01,500
rachul,2022-01-01,600
rachul,2022-01-01,700