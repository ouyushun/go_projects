package com.oys.bigdata.sprk.core

import org.apache.spark.{SparkConf, SparkContext}

/**
 * @Author ouyushun
 * @Date 2023/2/8
 * @Version 1.0
 */
object Test01 {
	def main(args: Array[String]): Unit = {
		//建立spark连接
		val sparkConf = new SparkConf().setAppName("WC").setMaster("local")
		val sc = new SparkContext(sparkConf)

		val rdd = sc.makeRDD(List((1,1),(2,2),(3,3)))
		val rdd2 = rdd.flatMap(
			List(_)
		)
		rdd2.collect().foreach(println)
	}
}
