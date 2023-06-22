package com.oys.ds.控制抽象

/**
 * 什么是抽象控制？
 * 1) 抽象控制是一个方法，但是方法的参数a是一个函数
 * 2）函数a的没有参数，也没有返回值
 * 主要目的：在一个方法中调用某个代码块，也可一理解成() => Unit的匿名函数,可以简写成 =>Unit
 */
object TestAbstractControl {
	def main(args: Array[String]): Unit = {


		//1
		def controlAbstract(f: () => Unit): Unit = {
			f()
		}

		//2，当函数的参数为()时，()可以省略，调用的时候也可以省略
		def controlAbstract2(f: => Unit): Unit = {
			f
		}

		//1
		controlAbstract {
			() =>
				new Thread {
					println("开始干活，干活中ing！～～")
					Thread.sleep(1000)
					println("活干完了")
				}.start()
		}


		//2
		controlAbstract2(
			new Thread() {
				override def run(): Unit = {
					println("开始干活")
					Thread.sleep(1000)
					println("ganwanle ")
				}
			}.start()
		)
	}
}

