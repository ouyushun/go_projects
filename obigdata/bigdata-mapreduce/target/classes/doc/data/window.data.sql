CREATE TABLE business (
id int comment '用户id',
name string comment '用户姓名',
orderdate string comment '订单日期',
const int comment '花费金额')
comment '用户消费表'
row format delimited fields terminated by ','
stored as textfile;
load data local inpath '/Users/ouyushun/hive/business' overwrite into table business;

// 向前三行

select *, sum(const) over(partition by id order by const) from business ;
select *,  from business ;


show tables;

show databases;

select * from user_info;


select * from business;

select count(1) over() from  business;











1,jack,2001-01-01,100
1,jack,2002-01-01,200
1,jack,2003-01-01,300
1,jack,2004-01-01,400
1,jack,2005-01-01,500
2,marry,2001-01-01,100
2,marry,2002-01-01,200
2,marry,2003-01-01,300
2,marry,2004-01-01,400
2,marry,2005-01-01,500
3,rachul,2001-01-01,100
3,rachul,2002-01-01,200
3,rachul,2003-01-01,300
3,rachul,2004-01-01,400
3,rachul,2005-01-01,500