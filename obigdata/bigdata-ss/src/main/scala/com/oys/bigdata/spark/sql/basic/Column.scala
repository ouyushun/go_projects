package com.oys.bigdata.spark.sql.basic
import org.apache.spark.sql
import org.apache.spark.sql.{ColumnName, DataFrame, Dataset, SparkSession}
/**
 * @Author ouyushun
 * @Date 2023/3/26
 * @Version 1.0
 */

object Column extends App {
	val spark: SparkSession = SparkSession.builder()
	  .master("local[6]")
	  .appName(this.getClass.getSimpleName)
	  .getOrCreate()


	import spark.implicits._
	import org.apache.spark.sql.functions._

	def creation(): Unit ={

		val ds: Dataset[Person] = Seq(Person("zhangsan", 20),Person("zhangsan", 18), Person("lisi", 15)).toDS
		val df: DataFrame       = Seq(Person("zhangsan", 20),Person("zhangsan", 18), Person("lisi", 15)).toDF
		//1. '
		val column: Symbol = 'name


		//2. $
		val column1: ColumnName = $"name"


		//3. col (必须导入functions)import org.apache.spark.sql.functions._
		val column2: sql.Column = col("name")

		//4. column (必须导入functions)
		val column3: sql.Column = column("name")

		/**
		 * 这四种创建方式有关联的DataSet么
		 */
		ds.select(column).show()

		//DataSet可以 DataFrame 可以使用 column 对象选中行吗
		df.select(column).show()

		//select 方法可以使用column 对象来选中某个列，那么其他的算子行么
		df.where(column === "zhangsan").show()


		/**
		 * column 有几个创建方式？ -> 四种
		 * column 对象可以用作与DataSet 和 DataFrame 中
		 * column 可以和命令式的弱类型的API 配合使用select where
		 *
		 */
		//5. dataset.col
		val column4 = ds.col("name")
		val column5 = df.col("name")

		//报错
		//ds.select(column5).show()

		//为什么要和dataset 来绑定呢
		// ds.join(df,ds.col("name") === df.col("name"))

		//6.dataset.apple
		val column6 = ds.apply("name")
		val column7 = ds("name")


	}

	def as(): Unit ={
		val ds: Dataset[Person] = Seq(Person("zhangsan", 20),Person("zhangsan", 18), Person("lisi", 15)).toDS

		ds.select ('name as "new_name").show()


		ds.select('age.as[Long]).show()
	}


	def api(): Unit ={
		val ds: Dataset[Person] = Seq(Person("zhangsan", 20),Person("zhangsan", 18), Person("lisi", 15)).toDS
		//需求一：ds增加列，双倍年龄
		// age*2 其实本质上就是是将一个表达式(逻辑计划表达式) 附着到column对象上
		//表达式在执行的时候对应每一条数据进行操作
		ds.withColumn("doubled",'age * 2).show()
		//需求二：模糊查询
		//select * from table where name like zhang%
		ds.where('name like "zhang%")

		//需求三：排序，正反序
		ds.sort('age asc).show()

		//需求四：枚举判断
		ds.where('name isin("zhangsan","wangwu","zhaoliu")).show()
	}

	as()
}
case class Person(name:String,age:Int)

