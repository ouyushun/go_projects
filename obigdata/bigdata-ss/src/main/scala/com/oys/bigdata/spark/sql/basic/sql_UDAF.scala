package com.oys.bigdata.spark.sql.basic

import org.apache.hadoop.shaded.org.eclipse.jetty.websocket.common.frames.DataFrame
import org.apache.spark.{SparkConf, sql}
import org.apache.spark.sql.SparkSession

/**
 * @Author ouyushun
 * @Date 2023/2/13
 * @Version 1.0
 */
object sql_UDAF {
	def main(args: Array[String]): Unit = {
		val sc = new SparkConf().setMaster("local[*]").setAppName("SparkSql-ds")
		val spark = SparkSession.builder().config(sc).getOrCreate()
		// TODO DataFrame
		val df: sql.DataFrame = spark.read.json("tmpdata/df-user.json")

		df.show()

		// DataFrame => SQL
		df.createOrReplaceTempView("user")

		spark.udf.register("prefixName", (name: String) => {
			"Name: " + name
		})
		spark.sql("select age, name from user").show
		spark.sql("select age, prefixName(name) from user").show


		spark.close()
	}

}

