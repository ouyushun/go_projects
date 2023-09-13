package com.oys.bigdata.spark.sql.basic

import org.apache.spark.SparkConf
import org.apache.spark.rdd.RDD
import org.apache.spark.sql.{DataFrame, Dataset, Row, SparkSession}

/**
 * @Author ouyushun
 * @Date 2023/2/13
 * @Version 1.0
 */
object sql01_ds {
	def main(args: Array[String]): Unit = {
		val sc = new SparkConf().setMaster("local[*]").setAppName("SparkSql-ds")
		val spark = SparkSession.builder().config(sc)
		  //.enableHiveSupport() //开启支持hive相关操作
		  .getOrCreate()


		import spark.implicits._

		// TODO DataSet
		// DataFrame其实是特定泛型的DataSet

		// RDD <=> DataFrame
		val rdd = spark.sparkContext.makeRDD(List((1, "zhangsan", 30), (2, "lisi", 40)))
		val df: DataFrame = rdd.toDF("id", "name", "age")
		val rdd1 = df.rdd

		// DataFrame <=> DataSet
		val ds: Dataset[User] = df.as[User]
		val df1: DataFrame = ds.toDF()

		// RDD <=> DataSet
		val ds1: Dataset[User] = rdd.map {
			case (id, name, age) => {
				User(id, name, age)
			}
		}.toDS()
		val userRDD: RDD[User] = ds1.rdd

		spark.close()
	}

	case class User(id: Int, name: String, age: Int)
}
