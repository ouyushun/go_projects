package com.oys.bigdata.sprk.core.framwork.controller

import com.atguigu.bigdata.spark.core.framework.common.TController
import com.atguigu.bigdata.spark.core.framework.service.WordCountService
import com.oys.bigdata.sprk.core.framwork.common.TController
import com.oys.bigdata.sprk.core.framwork.service.WordCountService


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
