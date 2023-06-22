package com.oys.ds.basic

/**
 * @Author ouyushun
 * @Date 2022/12/31
 * @Version 1.0
 */
object AssociatedDemo{
	// object中的apply()方法
	def apply(param:String){
		println("object apply method called:" + param)
	}
	def main(args: Array[String]): Unit = {
		// class 的apply()
		val ad2 = new AssociatedDemo();
		ad2("AAA")
		ad2("BBB")

		// object 的apply()
		AssociatedDemo("CCC")
		AssociatedDemo("DDD")
	}

	class AssociatedDemo {
		// class中的apply()方法
		def apply(param:String){
			println("class apply method called:" + param)
		}
	}
}