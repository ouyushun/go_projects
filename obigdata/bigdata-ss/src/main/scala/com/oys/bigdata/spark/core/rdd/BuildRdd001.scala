package com.oys.bigdata.spark.core.rdd

import org.apache.spark.{SparkConf, SparkContext}

/**
 * @Author ouyushun
 * @Date 2022/7/23
 * @Version 1.0
 */
object BuildRdd001 {
	def main(args: Array[String]): Unit = {
		var conf: SparkConf = new SparkConf().setMaster("local[*]").setAppName("BuILD RDD")
		val sc = new SparkContext(conf)
		val seq = Seq[Int](1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16)

		//val rdd = sc.parallelize(seq)
		//第二个参数表示分区的数量, 默认值为
		//从内存中创建RDD， 内存中的集合作为数据源
		val rdd = sc.makeRDD(seq)
		rdd.saveAsTextFile("/tmp/bbb")


		sc.stop()
	}
}
