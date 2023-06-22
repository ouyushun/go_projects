package com.oys.bigdata.spark.sparkOpt

import org.apache.spark.SparkConf
import org.apache.spark.sql.functions._
import org.apache.spark.sql.{DataFrame, Row, SparkSession}
import org.threeten.extra.scale.UtcRules.system

/**
 * @Author ouyushun
 * @Date 2023/3/5
 * @Version 1.0
 */
object Opt03 extends App {
	val sc = new SparkConf().setMaster("local[*]").setAppName("SparkSql-opt01")
	val spark = SparkSession.builder().config(sc).getOrCreate()

	//小王同学看到这个需求，说简单啊，然后三下五除二就完成了下面的代码：
	val item_df = spark.read.parquet("./data/game/parquet/item_purchase")
	item_df.show()
	item_df.printSchema()
/*
* +----------+--------------+----------+--------------------+-----------+-------+-------+-------------------+---------+---------+-------+
|channel_id|     client_ip|event_time|           item_info|online_time|op_type|plat_id|            role_id|role_name|server_id|user_id|
+----------+--------------+----------+--------------------+-----------+-------+-------+-------------------+---------+---------+-------+
|         2|532.44.442.021|1570620194|[{1, 1001}, {3, 1002}|          0|      1|      1|3102579400731348833|       头|        5|0107368|
+----------+--------------+----------+--------------------+-----------+-------+-------+-------------------+---------+---------+-------+
* */
	val resultDF = item_df
	  .withColumn("s", explode(col("item_info")))
	  .filter("s.id = '1002' ")
	  .filter("channel_id = '2' ")
	  .filter("plat_id = '1' ")
	  .groupBy("s.id", "channel_id", "plat_id")
	  .agg(
		countDistinct("user_id").alias("user_cnt"),
		sum("s.cnt").alias("sum_item_cnt"),
	).withColumn("avg_item_cnt", col("sum_item_cnt") / col("user_cnt"))

	resultDF.show(10,false)


	//优化一
	val resultOptDF = item_df
	  .filter("channel_id = '2' ")
	  .filter("plat_id = '1' ")
	  .select("item_info", "user_id")
	  .withColumn("s", explode(col("item_info"))).filter("s.id = '1002' ")
	  .groupBy("s.id")
	  .agg(
		countDistinct("user_id").alias("user_cnt"),
		sum("s.cnt").alias("sum_item_cnt"),
	).withColumn("avg_item_cnt", col("sum_item_cnt") / col("user_cnt"))
	  .withColumn("channel_id", lit("2"))
	  .withColumn("plat_id", lit("1"))

	resultOptDF.show(10, false)

	resultDF.show()

	//优化二
	/*
	* 先explode 会使一行变多行，使内存数据量增大
没进行列裁剪，选出所需的字段，这样夹带这许多我们并不需要的字段
没事先进行过滤，导致扫描的数据量变多
因此，我们得把过滤和列剪枝这些可以节省数据访问量的操作尽可能地往前推，
* 把计算开销较大的操作如 Shuffle 尽量往后拖，从而在整体上降低数据处理的负载和开销。
	* */
	val resultSuperOptDF2 = item_df
	  .filter("channel_id = '2' ")
	  .filter("plat_id = '1' ")
	  .filter("cast(item_info as string) rlike '1002' ") // 方式1 这种有可能会有 100200这样的item_id,会导致结果错误
	  .filter(" array_contains(item_info.id,1002) ") // 方式2
	  .filter(" exists(item_info.id, x-> x == 1002) ") // 方式3
	  .select("item_info", "user_id")
	  .withColumn("cnt", expr("element_at(map_from_arrays(item_info.id,item_info.cnt),1002)"))
	  .agg(
		  countDistinct("user_id").alias("user_cnt"),
		  sum("cnt").alias("sum_item_cnt"),
	  ).withColumn("avg_item_cnt", col("sum_item_cnt") / col("user_cnt"))
	  .withColumn("channel_id", lit("2"))
	  .withColumn("plat_id", lit("1"))
	  .withColumn("id", lit("1002"))

	resultSuperOptDF2.show(10, false)


	//优化三
	//能不能在explode之前就把id =1002的数据过滤出来呢？
	//能不能不用explode呢？直接把id =1002的cnt 取出来？
	val resultSuperOptDF3 = item_df
	  .filter("channel_id = '2' ")
	  .filter("plat_id = '1' ")
	  .filter("cast(item_info as string) rlike '1002' ") // 方式1 这种有可能会有 100200这样的item_id,会导致结果错误
	  .filter(" array_contains(item_info.id,1002) ") // 方式2
	  .filter(" exists(item_info.id, x-> x == 1002) ") // 方式3
	  .select("item_info", "user_id")
	  .withColumn("cnt", expr("element_at(map_from_arrays(item_info.id,item_info.cnt),1002)"))
	  .agg(
		  countDistinct("user_id").alias("user_cnt"),
		  sum("cnt").alias("sum_item_cnt"),
	  ).withColumn("avg_item_cnt", col("sum_item_cnt") / col("user_cnt"))
	  .withColumn("channel_id", lit("2"))
	  .withColumn("plat_id", lit("1"))
	  .withColumn("id", lit("1002"))
	resultSuperOptDF3.show(10, false)

	// 追加需求，想求最近一段时间，所有物品销量的总和
	val result2DF = item_df
	  .select(expr("aggregate(item_info.cnt,0L,(acc,x) -> acc + x)").alias("cnt"))

	result2DF.show()

	// 如果想求 物品id = 1002 的物品销量总和呢？
	val result3DF = item_df
	  .withColumn(
		"filter_item_info", expr("filter(item_info, x -> x.id == 1002)")
	)
	  .select(expr("aggregate(filter_item_info.cnt,0L,(acc,x) -> acc + x)").alias("cnt"))
	result3DF.show()

}
