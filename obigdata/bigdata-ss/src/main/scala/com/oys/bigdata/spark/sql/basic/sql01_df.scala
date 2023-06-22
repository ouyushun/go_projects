package com.oys.bigdata.spark.sql.basic

import org.apache.spark.SparkConf
import org.apache.spark.sql.{DataFrame, Row, SparkSession}

/**
 * @Author ouyushun
 * @Date 2023/2/13
 * @Version 1.0
 */
object sql01_df {
	def main(args: Array[String]): Unit = {
		val sc = new SparkConf().setMaster("local[*]").setAppName("SparkSql-df")
		val spark = SparkSession.builder().config(sc).getOrCreate()

		//隐式转换
		import spark.implicits._

		// TODO 执行逻辑操作

		// TODO DataFrame
		val df: DataFrame = spark.read.json("tmpdata/df-user.json")

		df.show()

		val frame = df.select().sample(false, 0.1)



		// DataFrame => SQL
		df.createOrReplaceTempView("user")

		spark.sql("select * from user").show
		spark.sql("select age, name from user").show
		spark.sql("select avg(age) from user").show

		// DataFrame => DSL
		// 在使用DataFrame时，如果涉及到转换操作，需要引入转换规则

		df.select("age", "name").show
		df.select($"age" + 1).show
		df.select('age + 1).show


		spark.close()
	}
}
