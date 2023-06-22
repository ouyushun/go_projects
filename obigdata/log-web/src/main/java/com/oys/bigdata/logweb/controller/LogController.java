package com.oys.bigdata.logweb.controller;


import org.apache.log4j.Logger;
import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.ResponseBody;

/**
 * 日志服务:
 * client ==> /upload ==>  落地磁盘
 *                       ==> 1) 输出到控制台
 *                       ==> 2) 落地到磁盘
 */
@Controller
public class LogController {

    private static final Logger logger = Logger.getLogger(LogController.class);

    @PostMapping("/upload")
    @ResponseBody
    public void upload(@RequestBody String info) {
        logger.info(info);
    }
}

