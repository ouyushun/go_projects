package com.oys.bigdata.spark.sql.join

import org.apache.spark.sql.functions.broadcast
import org.apache.spark.sql.{SaveMode, SparkSession}

/**
 * @Author ouyushun
 * @Date 2023/2/23
 * @Version 1.0
 */
object SortMergeBucketJoinDemo01 extends App {
	private val spark = SparkSession.builder().master("local[1]").getOrCreate()

	//https://blog.51cto.com/u_15127538/2659053

	def useBucket(sparkSession: SparkSession) = {
		sparkSession.read.json("/user/atguigu/ods/coursepay.log")
		  .write.partitionBy("dt", "dn")
		  .format("parquet")
		  .bucketBy(10, "orderid")
		  .sortBy("orderid")
		  .mode(SaveMode.Overwrite)
		  .saveAsTable("dwd.dwd_course_pay_cluster")

		sparkSession.read.json("/user/atguigu/ods/courseshoppingcart.log")
		  .write.partitionBy("dt", "dn")
		  .bucketBy(10, "orderid")
		  .format("parquet")
		  .sortBy("orderid")
		  .mode(SaveMode.Overwrite)
		  .saveAsTable("dwd.dwd_course_shopping_cart_cluster")  }

	def useSMBJoin(sparkSession: SparkSession) = {
		//查询出三张表 并进行join 插入到最终表中
		val saleCourse = sparkSession.sql("select *from dwd.dwd_sale_course")

		val coursePay = sparkSession.sql("select * from dwd.dwd_course_pay_cluster")
		  .withColumnRenamed("discount", "pay_discount")
		  .withColumnRenamed("createtime", "pay_createtime")

		val courseShoppingCart = sparkSession.sql("select *from dwd.dwd_course_shopping_cart_cluster")
		  .drop("coursename")
		  .withColumnRenamed("discount", "cart_discount")
		  .withColumnRenamed("createtime", "cart_createtime")

		//两张分桶表进行 join
		val tmpdata = courseShoppingCart.join(coursePay, Seq("orderid"), "left")

		//再与小表进行 join
		val result = broadcast(saleCourse).join(tmpdata, Seq("courseid"), "right")

		result.select("courseid", "coursename", "status", "pointlistid", "majorid", "chapterid", "chaptername", "edusubjectid"
			, "edusubjectname", "teacherid", "teachername", "coursemanager", "money", "orderid", "cart_discount", "sellmoney",
			"cart_createtime", "pay_discount", "paymoney", "pay_createtime", "dwd.dwd_sale_course.dt", "dwd.dwd_sale_course.dn")
		  .write.mode(SaveMode.Overwrite).saveAsTable("dws.dws_salecourse_detail_2")
	}
}