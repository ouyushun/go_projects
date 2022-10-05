package com.oys.bigdata.core.rdd;

import org.apache.spark.SparkConf;
import org.apache.spark.SparkContext;
import org.apache.spark.api.java.JavaRDD;
import org.apache.spark.api.java.JavaSparkContext;
import org.apache.spark.api.java.function.FlatMapFunction;
import org.apache.spark.api.java.function.Function;

import java.util.ArrayList;
import java.util.Arrays;
import java.util.Iterator;

/**
 * @Author ouyushun
 * @Date 2022/7/24
 * @Version 1.0
 */
public class RddTransferMap {
    public static void main(String[] args) {

        SparkConf conf = new SparkConf();
        conf.setMaster("local[1]");
        conf.setAppName("SPARK ES");
        JavaSparkContext sc = new JavaSparkContext(conf);
        JavaRDD<String> javaRdd = sc.parallelize(Arrays.asList("a", "b", "c", "d", "e"));

        System.out.println("map(func)：通过函数func传递的每个元素，返回一个新的RDD。");
        JavaRDD<Object> map = javaRdd.map(
                 (Function<String, Object>) item -> "new" + item
        );

        map.foreach(x -> System.out.println(x));


        System.out.println("filter(func)：筛选通过func处理后返回 true 的元素，返回一个新的RDD。");
        JavaRDD<String> filter = javaRdd.filter(item -> item.equals("a") || item.equals("b"));
        filter.foreach(x -> System.out.println(x));

        System.out.println("flatMap(func)：类似于 map，但每个输入项可以映射到 0 个或更多输出项。");
        JavaRDD<String> rdd2 = sc.parallelize(Arrays.asList("a,b", "c,d,e", "f,g"));
        JavaRDD<String> flatMap = rdd2.flatMap((FlatMapFunction<String, String>)
                s -> Arrays.asList(s.split(",")).iterator());
        flatMap.foreach(x -> System.out.println(x));


        System.out.println("mapPartitions  mapPartitions(func)：类似于map，但该函数是在RDD每个partition上单独运行，因此入参会是Iterator<Object>");
        JavaRDD<String> mapPartitions = javaRdd.mapPartitions((FlatMapFunction<Iterator<String>, String>) stringIterator -> {
            ArrayList<String> list = new ArrayList<>();
            while (stringIterator.hasNext()) {
                list.add(stringIterator.next());
            }
            return list.iterator();
        });
        mapPartitions.foreach(x -> System.out.println(x));








        sc.stop();
    }
}
