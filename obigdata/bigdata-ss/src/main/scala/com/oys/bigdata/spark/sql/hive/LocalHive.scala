package com.oys.bigdata.spark.sql.hive

import com.oys.bigdata.spark.sql.hive.SparkOnHiveDemo.spark
import org.apache.spark.sql.{DataFrame, SparkSession}

object LocalHive extends App {
	private val spark = SparkSession.builder().master("local[1]").getOrCreate()
	spark.sql("show tables").show()

	val df: DataFrame = spark.read.json("tmpdata/df-user.json")

	//临时表
	df.createOrReplaceTempView("user0001")
	spark.sql("show tables").show()
	spark.sql("create table aa(id int)")
	//spark.read.json("tmpdata/df-user.json").write.saveAsTable("df_user003")
	spark.sql("sho w tables").show()
}
