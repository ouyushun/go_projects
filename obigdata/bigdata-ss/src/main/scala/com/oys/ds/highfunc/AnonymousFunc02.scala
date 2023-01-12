package com.oys.ds.highfunc

/**
 * @Author ouyushun
 * @Date 2023/1/6
 * @Version 1.0
 */
object Func01 {
	def main(args: Array[String]): Unit = {
		def fun(): Unit = {
			println("fun---")
		}
		def sum(x:Int, y:Int) : Int = x + y
		def diff(x:Int, y:Int) : Int = x - y

		def test(f:(Int, Int)=>Int): Int = {
			f(100, 10)
		}

		println(test(sum))
		println(test(diff))
		//名称不重要可以省略
		println(test((x:Int, y:Int) => {x+y}))
		println(test((x:Int, y:Int) => {x-y}))
		println(test((x:Int, y:Int) => {x*y}))
		println(test((x:Int, y:Int) => {x/y}))

		//至简原则
		println(test((x, y) => x+y))
		println(test(_ + _)) //只使用一次， 按顺序
		println(test((x, y) => x-y))
		println(test((x:Int, y:Int) => {x*y}))
		println(test((x:Int, y:Int) => {x/y}))
	}
}
