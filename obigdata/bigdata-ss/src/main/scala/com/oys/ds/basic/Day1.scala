package com.oys.ds.basic

/**
 * @Author ouyushun
 * @Date 2022/12/13
 * @Version 1.0
 */
object Day1 {

  def main(args: Array[String]): Unit = {
    var a = -1
    println(a.abs)
  }

  def abs(num:Int): Int = {
    if (num > 0) {
      return num
    } else {
      return -num
    }
  }

  def fac1(num:Int) : Int = {
    var x = 1
    for (i <- 1 to num) x = x * i
    x
  }

  def fac2(num:Int) : Int = {
    if (num <= 2) num else num * fac2(num - 1)
  }

  def cal(x:Int,y:Int)= {
    (x+y)/(x-y)
  }



}
