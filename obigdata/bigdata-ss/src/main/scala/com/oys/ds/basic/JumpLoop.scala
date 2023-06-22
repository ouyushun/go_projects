package com.oys.ds.basic

import scala.util.control.Breaks.{break, breakable}

/**
 * @Author ouyushun
 * @Date 2022/12/17
 * @Version 1.0
 */
object JumpLoop {
  def main(args: Array[String]): Unit = {
    printMultiTable()
  }

  def jump(): Unit = {
    val max=100
    var sum=0
    breakable{
      for(i<-1 to max){
        sum+=i//sumnum
        if(sum>20){
          println("i="+i)
          break()
          }
        }
    }
    for(i<-1 to 10){
      if(i!=4 && i!=5){
        println("i="+i)
      }
    }
  }


  def printMultiTable()={
    var i=1
    //只有i在作用域内
    while(i<=10){
      var j=1
      //i和j在作用域内
      while(j<=10){
        val prod=(i*j).toString
        //i、j和prod在作用域内
        var k=prod.length
        //i、j、prod和k在作用域内
        while(k<4){
          print("-")
          k+=1
        }
        print(prod)
        j+=1
      }
      //i和j仍在作用域内，prod和k超出了作用域
      println()
      i+=1
      println("before: j=" , j)
      j+=3
      println("after: j=" , j)
    }
    //i仍在作用域内，j、prod和k超出了作用域
  }

  val temp7 = for (
    x <- List(1, 2);
    y <- List("one", "two")
  ) yield (x, y)
  println(temp7)
}
