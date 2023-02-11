package com.oys.bigdata.sprk.core.CasePractice.top3

import org.apache.spark.rdd.RDD
import org.apache.spark.{SparkConf, SparkContext}

/**
 * @Author ouyushun
 * @Date 2022/7/23
 * @Version 1.0
 */
object Rdd_transform_top3 {
  def main(args: Array[String]): Unit = {
    //建立spark连接
    val sparkConf = new SparkConf().setAppName("WC").setMaster("local")

    val sc = new SparkContext(sparkConf)

    //时间戳 省份 城市 用户 广告
    val dataRdd = sc.textFile("tmpdata/top3.txt")

    //1. 时间戳 省份 城市 用户 广告 =》 ((省份， 广告), 1)
    val mapRdd = dataRdd.map(
      line => {
        val lineArr = line.split(" ")
        ((lineArr(1), lineArr(4)), 1)
      }
    )

    //2. ((省份， 广告), 1) => ((省份， 广告), sum)
    val reduceRdd = mapRdd.reduceByKey(_+_)

    //3.((省份， 广告), sum) => (省份， (广告, sum))
    val newMapRdd = reduceRdd.map {
      case ((peovince, ad), sum) => {
        (peovince, (ad, sum))
      }
    }

    //4.分组 (省份， (广告, sum))
    var groupRdd: RDD[(String, Iterable[(String, Int)])] = newMapRdd.groupByKey()

    //排序, 降序
    val resRdd = groupRdd.mapValues(
      iter => {
        iter.toList.sortWith(_._2 > _._2).take(3)
      }
    )

    resRdd.cache()
    resRdd.persist()
    sc.setCheckpointDir("./")
    resRdd.checkpoint()

    resRdd.collect().foreach(println)

    //关闭
    sc.stop()
  }

}
