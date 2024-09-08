// 将本地文件上传到hdfs
hdfs dfs -moveFromLocal ./olist_customers_dataset.csv  /tmp/olist_customers_dataset

//2.载入数据
load data local inpath '/home/ds_teacher/userdd.txt' overwrite into table user_info;
load data local inpath '/Users/ouyushun/hive/user_info' overwrite into table user_info;

// hive从hdfs加载
load data inpath '/home/ds_teacher/userdd.txt' overwrite into table user_info;
load data local inpath '/home/ouyushun0919/hive/user_info' overwrite into table user_info;
show databases;

add jar 本地目录
jar lists;
delete jar ... ;

create temporary function "方法名" as "全路径";
create temporary function oys_encrypt_phone as "com/hive/udf/test/EncryptNumber.java";



--1.新建表
CREATE TABLE user_info (
id int comment '用户id',
name string comment '用户姓名',
age int comment '用户年龄',
city string comment '城市名称',
birthday date comment '用户出生年月')
comment '用户信息表'
    row format delimited fields terminated by ','
    stored as textfile;
