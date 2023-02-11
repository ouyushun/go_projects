package com.oys.ds.basic
import scala.collection.mutable
/**
 * @Author ouyushun
 * @Date 2022/12/18
 * @Version 1.0
 */
object PartialFunction {def main(args: Array[String]): Unit = {


	val p:PartialFunction[Int, String] = {
		case 1 => "One"
		case 2 => "two"
		case _	=> "ss"
	}
	println(p(2))
	println(p(5))
}


}

