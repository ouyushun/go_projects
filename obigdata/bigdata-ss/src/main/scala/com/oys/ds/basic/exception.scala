package com.oys.ds.basic

/**
 * @Author ouyushun
 * @Date 2022/12/17
 * @Version 1.0
 */
object exception {
  def main(args: Array[String]): Unit = {
	try {
		var res = 10 / 0
	}catch {
	  case exception: ArithmeticException => {println("divid by zero" + exception.printStackTrace())}
	} finally {
	  println("finally...")
	}
  }
}
