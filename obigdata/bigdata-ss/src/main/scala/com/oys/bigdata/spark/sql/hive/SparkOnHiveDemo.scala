package com.oys.bigdata.spark.sql.hive


import org.apache.spark.sql.catalyst.TableIdentifier
import org.apache.spark.sql.functions.{column, lit}
import org.apache.spark.sql.{DataFrame, SaveMode, SparkSession}

object SparkOnHiveDemo extends App {
	private val spark: SparkSession = SparkSession.builder().master("local[1]").getOrCreate()
	// hive read
	spark.read.table("ds_spark.people").show()
	spark.sql("select * from ds_spark.people").show()
	spark.table("ds_spark.people").show()

	// hive write
	// saveAsTable insertInto
	spark.read.json("user_log.json")
	  .write.saveAsTable("ds_spark.ods_user_login")
	// alter spark
	spark.read.table("ds_spark.ods_user_login").count()
	spark.read.json("user_login.json")
	  .write.insertInto("ds_spark.ods_user_login")
	spark.read.table("ds_spark.ods_user_login").count()

	// 创建一个hive 外表
	spark.read.json("tmpdata/df-user.json").write.option("path", "ds_warehouse/ods_user_login_external").mode("overwrite").saveAsTable("ds_spark.ods_user_login_external")

	// 创建一个分区表怎么办？？？？
	spark.read.json("user_log.json").withColumn("dt_pt", lit("20220401"))
	  .write.partitionBy("dt_pt")
	  .saveAsTable("ds_spark.ods_user_login_pt")

	// insert 插入数据

	spark.read.json("user_log.json").withColumn("dt_pt", column("channel_id"))
	  .write.insertInto("ds_spark.ods_user_login_pt")

	spark.read.json("user_log.json").withColumn("dt_pt", lit("20220401"))
	  .withColumn("dt_pt2", lit("20220401"))
	  .write.mode("overwrite").partitionBy("dt_pt", "dt_pt2")
	  .saveAsTable("ds_spark.ods_user_login_pt2")


	spark.read.json("user_log.json").withColumn("dt_pt", lit("20220402"))
	  .withColumn("dt_pt2", lit("20220402"))
	  .write.insertInto("ds_spark.ods_user_login_pt2")

	spark.sql("set spark.sql.sources.partitionOverwriteMode = 'DYNAMIC' ")


	def saveAsHiveTable(
						 spark: SparkSession,
						 data: DataFrame,
						 dbName: String,
						 tableName: String,
						 tableSavePath: String = "",
						 ptCols: Array[String]
					   ): Unit = {

		// saveAsTable 建表
		// insert into 执行动态分区插入
		// if ... else ....
		// 指定内表 外表
		// 分区字段
		// create database ds_spark  local 'xxxx'

		val table = s"$dbName.$tableName"
		val tableIdentifier = TableIdentifier(tableName, Some(dbName))
		val catalog = spark.sessionState.catalog

		if (catalog.tableExists(tableIdentifier)) {
			// 如果表存在，则使用 insertInto 插入
			spark.conf.set("spark.sql.sources.partitionOverwriteMode", "DYNAMIC")
			// 分区字段放到末尾
			val sortCols = Array.concat(
				data.columns.filterNot(ptCols.contains(_)),
				ptCols
			)
			val resultDF = data.select(sortCols.head, sortCols.tail: _*)
			resultDF.write.mode(SaveMode.Overwrite).insertInto(table)

		} else {
			// 如果表不存在，则使用
			val path = s"$tableSavePath/$tableName"

			data.write
			  .mode(SaveMode.Overwrite)
			  .options(Map("path" -> path))
			  // array => string* 格式
			  .partitionBy(ptCols: _*)
			  .saveAsTable(table)
		}
	}
}
