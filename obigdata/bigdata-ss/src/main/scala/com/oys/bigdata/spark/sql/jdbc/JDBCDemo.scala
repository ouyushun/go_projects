package com.oys.bigdata.spark.sql.jdbc

import org.apache.spark.sql.SparkSession

import java.util.Properties

object JDBCDemo extends App {
	private val spark: SparkSession = SparkSession.builder().master("local[1]").getOrCreate()
	// maven 导包

	/*
			<dependency>
			  <groupId>mysql</groupId>
			  <artifactId>mysql-connector-java</artifactId>
			  <version>8.0.28</version>
		  </dependency>

	 */

	// spark 添加外部依赖
	//  spark.sparkContext.addJar("./jar/mysql-connector-java-8.0.27.jar")

	// jar spark-submit --jar  /user/xxxx/mysql-connector-java-8.0.27.jar



	//  spark.read.jdbc()
	val mysqlUrl = "jdbc:mysql://localhost:3306/ds_mysql?useSSL=false"
	val mysqlDriver = "com.mysql.cj.jdbc.Driver"
	val mysqlTableName = "ds_test"
	val mysqlUser = "root"
	val mysqlPassword = "123456"

	val jdbcMap = Map(
		"url" -> mysqlUrl,
		"driver" -> mysqlDriver,
		"dbtable" -> mysqlTableName,
		"user" -> mysqlUser,
		"password" -> mysqlPassword
	)
	//jdbc 的读取操作


	//  val df = spark.read.format("jdbc")
	//    .options(jdbcMap).load()
	//  df.show(10,false)
	// spark write mysql 自动推断类型，自动建表，
	//
	//  spark.read.json("./data/people.json")
	//    .write.format("jdbc").options(jdbcMap).save()
	val connProp = new Properties()
	connProp.put("user",mysqlUser)
	connProp.put("password",mysqlPassword)
	connProp.put("driver",mysqlDriver)

	//  spark.read.jdbc(mysqlUrl,mysqlTableName,connProp).show()
	// mysql 5.0x com.mysql.jdbc.Driver  mysql 8.0x com.mysql.cj.jdbc.Driver

	// 执行一条sql语句
	val jdbcMap2 = Map(
		"url" -> mysqlUrl,
		"driver" -> mysqlDriver,
		"user" -> mysqlUser,
		"password" -> mysqlPassword
	)
	val query = "select name from ds_test_2"

	val df = spark.read.format("jdbc")
	  .options(jdbcMap2).option("query",query).load()
	//  df.show()

	// spark 自动建表不被期望，mysql  index, pri key,unique key

	val jdbcMap3 = Map(
		"url" -> mysqlUrl,
		"driver" -> mysqlDriver,
		"user" -> mysqlUser,
		"password" -> mysqlPassword,
		"dbtable" -> "ds_test_3",
		"truncate"->"true"

	)
	//    spark.read.json("./data/people.json")
	//      .write.mode("overwrite").format("jdbc").options(jdbcMap3).save()

	//读取数据时自定义schema
	val schema = "event_time string,online_time string"


	val df1 = spark.read.format("jdbc")
	  .options(jdbcMap)
	  .option("customSchema",schema).load()
	//  df1.show()
	//  df1.printSchema()

	// mysql id text => int   hbase id =int

	val jdbcMap4 = Map(
		"url" -> mysqlUrl,
		"driver" -> mysqlDriver,
		"user" -> mysqlUser,
		"password" -> mysqlPassword,
		"dbtable" -> "ds_test_3",
	)
	val writeSchema = "name varchar(255),age varchar(100)"

	spark.read.json("./data/people.json")
	  .write.mode("overwrite")
	  .option("createTableColumnTypes",writeSchema)
	  .format("jdbc")
	  .options(jdbcMap4).save()
}


