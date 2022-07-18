package com.oys.bigdata.kafka;

import org.apache.kafka.clients.producer.KafkaProducer;
import org.apache.kafka.clients.producer.ProducerRecord;

import java.util.Properties;

public class ImoocKafkaProducer {


    public static void main(String[] args) {
        // 入口点

        Properties props = new Properties();
        props.put("bootstrap.servers", "localhost:9092");
        props.put("acks", "all");
        props.put("key.serializer", "org.apache.kafka.common.serialization.StringSerializer");
        props.put("value.serializer", "org.apache.kafka.common.serialization.StringSerializer");

        KafkaProducer<String, String> producer = new KafkaProducer<>(props);

        for(int i=0; i<100; i++) {
            System.out.println("-------");
            producer.send(new ProducerRecord<String, String>("my-replicated-topic",i+"",i+""));
        }
        System.out.println("消息发送完毕...");

        producer.close();
    }
}
