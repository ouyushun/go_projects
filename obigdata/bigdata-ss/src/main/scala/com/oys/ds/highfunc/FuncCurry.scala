package com.oys.ds.highfunc

/**
 * @Author ouyushun
 * @Date 2023/1/7
 * @Version 1.0
 */
object FuncCurry {
	def main(args: Array[String]): Unit = {
		def mul(x:Int, y:Int) = x * y
		def mulOne(x:Int) = (y:Int) => x * y
		// mulOne(8)的结果是函数(y: Int) => 6 * y，崦这个函数又被应用到 7，因此得到 42。


		//第一种写法
		def curry1(x:Int)(y:Int) = {
			x * y
		}
		val intToUnit : Int => Unit = curry1(10)
		println(intToUnit(20))
	}
}
