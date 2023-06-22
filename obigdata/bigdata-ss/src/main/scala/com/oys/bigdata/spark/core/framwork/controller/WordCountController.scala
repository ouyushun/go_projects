package com.oys.bigdata.spark.core.framwork.controller

import com.oys.bigdata.spark.core.framwork.common.TController
import com.oys.bigdata.spark.core.framwork.service.WordCountService


/**
  * 控制层
  */
class WordCountController extends TController {

    private val wordCountService = new WordCountService()

    // 调度
    def dispatch(): Unit = {
        // TODO 执行业务操作
        val array = wordCountService.dataAnalysis()
        array.foreach(println)
    }
}
