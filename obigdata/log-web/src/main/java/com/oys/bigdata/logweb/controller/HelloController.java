package com.oys.bigdata.logweb.controller;

import org.springframework.stereotype.Controller;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.ResponseBody;

/**
 * @Author ouyushun
 * @Date 2022/7/15
 * @Version 1.0
 */
@Controller
public class HelloController {

    @ResponseBody
    @RequestMapping("/hello")
    public String sayHello() {
        return "9HelHello OYS!";
    }


    @ResponseBody
    @RequestMapping("/hello3")
    public String sayHelloo() {
        return "oooooo Hello World! Hello OYS!";
    }

}
