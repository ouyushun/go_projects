package com.oys.bigdata.sprk.core.CasePractice.user_analysis

import org.apache.spark.rdd.RDD
import org.apache.spark.{SparkConf, SparkContext}

/**
 * @Author ouyushun
 * @Date 2022/7/23
 * @Version 1.0
 */
object HotCategoryTop10Analysis {
  def main(args: Array[String]): Unit = {
    fun1() //V1
    fun2() //V2
  }

  //优化逻辑
  def fun2() = {
    //建立spark连接
    val sparkConf = new SparkConf().setAppName("Hot Category Top10").setMaster("local")
    val sc = new SparkContext(sparkConf)

    //1. 读取全部数据
    val lines: RDD[String] = sc.textFile("tmpdata/user_action.csv")

    lines.cache()

    // 直接组装成 （点击数， 下单数，支付数）
    val flatMapRdd = lines.flatMap(
      action => {
        val datas = action.split("_")
        if (datas(6) != "-1") {
          List((datas(6), (1, 0, 0)))
        } else if (datas(8) != "null") {
          val idList = datas(8).split(",")
          idList.map(
            id => (id, (0, 1, 0))
          )
        } else if (datas(10) != "null") {
          val cids = datas(10).split(",")
          cids.map(
            id=>(id, (0, 0, 1))
          )
        } else {
          Nil
        }
      }
    )

    val analysisRdd = flatMapRdd.reduceByKey(
      (tuple1, tuple2) => {
        (tuple1._1 + tuple2._1, tuple1._2 + tuple2._2, tuple1._3 + tuple2._3)
      }
    )

    val resultRDD = analysisRdd.sortBy(_._2, false).take(10)

    // 6. 将结果采集到控制台打印出来
    resultRDD.foreach(println)

    //关闭
    sc.stop()
  }


  def fun1(): Unit = {
    //建立spark连接
    val sparkConf = new SparkConf().setAppName("Hot Category Top10").setMaster("local")
    val sc = new SparkContext(sparkConf)

    //1. 读取全部数据
    val lines: RDD[String] = sc.textFile("tmpdata/user_action.csv")

    lines.cache()

    // 2. 统计品类的点击数量：（品类ID，点击数量）
    val clickActionRdd = lines.filter(
      action => {
        val datas = action.split("_")
        datas(6) != "-1"
      }
    )

    val clickCountRdd = clickActionRdd.map(
      data => {
        val datas = data.split("_")
        (datas(6), 1)
      }
    ).reduceByKey(_+_)


    // 3. 统计品类的下单数量：（品类ID，下单数量）
    val orderActionRdd = lines.filter(
      line => {
        val datas = line.split("_")
        datas(8) != "null"
      }
    )

    /*ids = 1,2,3*/
    val orderCountRdd = orderActionRdd.flatMap(
      line => {
        val datas = line.split("_")
        val idsStr = datas(8)
        val idList = idsStr.split(",")
        idList.map(
          id => (id, 1)
        )
      }
    ).reduceByKey(_ + _)

    // 4. 统计品类的支付数量：（品类ID，支付数量）
    val payActionRDD = lines.filter(
      action => {
        val datas = action.split("_")
        datas(10) != "null"
      }
    )

    // orderid => 1,2,3
    // 【(1,1)，(2,1)，(3,1)】
    val payCountRDD = payActionRDD.flatMap(
      action => {
        val datas = action.split("_")
        val cid = datas(10)
        val cids = cid.split(",")
        cids.map(id=>(id, 1))
      }
    ).reduceByKey(_+_)

    // 5. 将品类进行排序，并且取前10名
    //    点击数量排序，下单数量排序，支付数量排序
    //    元组排序：先比较第一个，再比较第二个，再比较第三个，依此类推
    //    ( 品类ID, ( 点击数量, 下单数量, 支付数量 ) )
    //
    val cogroupRDD: RDD[(String, (Iterable[Int], Iterable[Int], Iterable[Int]))] =
    clickCountRdd.cogroup(orderCountRdd, payCountRDD)

    val analysisRDD = cogroupRDD.mapValues{
      case ( clickIter, orderIter, payIter ) => {

        var clickCnt = 0
        val iter1 = clickIter.iterator
        if ( iter1.hasNext ) {
          clickCnt = iter1.next()
        }
        var orderCnt = 0
        val iter2 = orderIter.iterator
        if ( iter2.hasNext ) {
          orderCnt = iter2.next()
        }
        var payCnt = 0
        val iter3 = payIter.iterator
        if ( iter3.hasNext ) {
          payCnt = iter3.next()
        }

        ( clickCnt, orderCnt, payCnt )
      }
    }

    val resultRDD = analysisRDD.sortBy(_._2, false).take(10)

    // 6. 将结果采集到控制台打印出来
    resultRDD.foreach(println)

    //关闭
    sc.stop()
  }


}
