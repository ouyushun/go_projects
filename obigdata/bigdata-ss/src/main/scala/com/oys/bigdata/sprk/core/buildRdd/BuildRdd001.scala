package com.oys.bigdata.sprk.core.buildRdd

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
    val seq = Seq[Int](elems = 1,2,3,4,5,6)

    //val rdd = sc.parallelize(seq)
    //第二个参数表示分区的数量, 默认值为
    val rdd = sc.makeRDD(seq)
    rdd.saveAsTextFile("/tmp/bbb")

    sc.stop()
  }
}
