package com.oys.bigdata.sprk.core.rdd

import org.apache.spark.rdd.RDD
import org.apache.spark.{HashPartitioner, SparkConf, SparkContext}

/**
 * @Author ouyushun
 * @Date 2022/7/23
 * @Version 1.0
 */
object Rdd_transform_partitionBy {
  def main(args: Array[String]): Unit = {
    //建立spark连接
    val sparkConf = new SparkConf().setAppName("WC").setMaster("local")

    val sc = new SparkContext(sparkConf)

    val rdd : RDD[Int] = sc.makeRDD(List(1, 2, 3, 4), 8)


    val mapRdd : RDD[(Int, Int)]= rdd.map((_, 1))

    //隐式转换(二次编译)
    //implicit def rddToPairRDDFunctions[K, V](rdd: RDD[(K, V)])
    //数据重分区
    mapRdd.partitionBy(new HashPartitioner(2))

    //可以自定义分区器

    //关闭
    sc.stop()
  }

}
