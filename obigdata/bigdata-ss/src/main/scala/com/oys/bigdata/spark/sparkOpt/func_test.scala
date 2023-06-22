package com.oys.bigdata.spark.sparkOpt

import org.apache.spark.SparkConf
import org.apache.spark.sql.{DataFrame, SparkSession}
import org.apache.spark.sql.functions._

/**
 * @Author ouyushun
 * @Date 2023/3/5
 * @Version 1.0
 */
object func_test  {
	def main(args: Array[String]): Unit = {
		val sc = new SparkConf().setMaster("local[*]").setAppName("SparkSql-opt01")
		val spark = SparkSession.builder().config(sc).getOrCreate()

		val data = Seq(("Prashant","Banglore",25, 58, "2022-08-01", 1),
			("Ankit","Banglore",26,54,"2021-05-02", 2),
			("Ramakant","Gurugram",24, 60, "2022-06-02", 3),
			("Brijesh","Gazipur", 26,75,"2022-07-04", 4),
			("Devendra","Gurugram", 27, 62, "2022-04-03", 5),
			("Ajay","Chandigarh", 25,72,"2022-02-01", 6))

		import spark.implicits._
		val df_friends =data.toDF("friends_name","location", "age", "weight", "meetup_date", "offset")


		exprTest(spark, df_friends)
	}

	def exprTest(spark: SparkSession, df: DataFrame) = {
		val df_concat = df.withColumn(
			"name-age-location", expr("friends_name|| '-'|| age || '-' || location")
		)

		val df_concat2 = df_concat.withColumn(
			"concat_2", concat_ws("@", col("friends_name"), col("age"))
		)

		val df_concat3 = df_concat2.withColumn(
			"Exercise_Need", expr("CASE WHEN weight >= 60  THEN 'Yes' " + "WHEN  weight < 55  THEN 'No' ELSE 'Enjoy' END")
		)

		df_concat3.show()
	}

}
