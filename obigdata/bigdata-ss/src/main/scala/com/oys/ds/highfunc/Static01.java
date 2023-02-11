package com.oys.ds.highfunc;

/**
 * @Author ouyushun
 * @Date 2023/1/5
 * @Version 1.0
 */
public  class Static01 {
    public static void main(String[] args) {
        test1();
        (new Static01()).test2();
    }

    public static void test1() {
        System.out.println("11111");
    }

    public void test2() {
        System.out.println("2222222");
    }
}
