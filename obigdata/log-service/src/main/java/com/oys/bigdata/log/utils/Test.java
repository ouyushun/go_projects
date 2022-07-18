package com.oys.bigdata.log.utils;

import com.imooc.bigdata.gen.LogGenerator;

public class Test {
    public static void main(String[] args)throws Exception {
        String url = "http://localhost:9527/log-web/upload";
        String code = "FCD8BBC4A6CAF839";
        LogGenerator.generator(url, code);
    }
}
