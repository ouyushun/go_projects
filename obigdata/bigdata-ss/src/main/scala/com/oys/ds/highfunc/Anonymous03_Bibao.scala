package com.oys.ds.highfunc

/**
 * @Author ouyushun
 * @Date 2023/1/7
 * @Version 1.0
 */
object Anonymous03_Bibao {
	def main(args: Array[String]): Unit = {

		
		def outer(x:Int) = {
			def inner(y:Int) = {
				x + y
			}
			inner _
		}

		println(outer(10)(20))

		def funA() = {
			var a = 0
			def funB() : Int = {
				a = a + 1
				a
			}
			funB _
		}

		val f = funA()

		println(f())
		println(f())
		println(f())
		println(f())
	}
}
