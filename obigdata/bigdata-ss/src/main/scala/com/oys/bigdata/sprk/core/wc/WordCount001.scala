package com.oys.bigdata.sprk.core.wc

import org.apache.spark.rdd.RDD
import org.apache.spark.{SparkConf, SparkContext}

/**
 * @Author ouyushun
 * @Date 2022/7/23
 * @Version 1.0
 */
object WordCount001 {
  def main(args: Array[String]): Unit = {
    //建立spark连接
    val sparkConf = new SparkConf().setAppName("WC").setMaster("local")
    var context: SparkContext = new SparkContext(sparkConf)
    val sc = context

    //操作
    val lines: RDD[String] = sc.textFile("wc.txt")
    val words: RDD[String] = lines.flatMap(_.split(""))
    val wordToOne: RDD[(String, Int)] = words.map(word => (word, 1))
    val wordTosum: RDD[(String, Int)] = wordToOne.reduceByKey(_ + _)

    var tuples: Array[(String, Int)] = wordTosum.collect()
    //关闭

  }

}
