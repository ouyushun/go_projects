package com.oys.bigdata.spark.core.framwork.application

import com.oys.bigdata.spark.core.framwork.common.TApplication
import com.oys.bigdata.spark.core.framwork.controller.WordCountController

object WordCountApplication extends App with TApplication{

    // 启动应用程序
    start(){
        val controller = new WordCountController()
        controller.dispatch()
    }
}
