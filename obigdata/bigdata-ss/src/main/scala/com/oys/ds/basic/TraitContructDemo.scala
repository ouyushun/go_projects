package com.oys.ds.basic

/**
 * @Author ouyushun
 * @Date 2022/12/31
 * @Version 1.0
 */
object TraitContructDemo {
	def main(args: Array[String]): Unit = {

	}


	class E {
		println("E------")
	}
	class F extends E with C with D{
		println("F------")
	}
	class G
	trait A {
		println("A-----")
	}
	trait B extends A {
		println("B-----")
	}
	trait C extends B {
		println("B-----")
	}
	trait D extends B {
		println("B-----")
	}
}
