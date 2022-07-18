package com.oys.bigdata.kafka;

import org.apache.kafka.clients.consumer.ConsumerRecord;
import org.apache.kafka.clients.consumer.ConsumerRecords;
import org.apache.kafka.clients.consumer.KafkaConsumer;

import java.time.Duration;
import java.util.Arrays;
import java.util.Properties;

public class ImoocKafkaConsumer {


    public static void main(String[] args) {

        Properties properties = new Properties();
        properties.put("bootstrap.servers","localhost:9092");
        properties.put("group.id","pk-group");
        properties.put("enable.auto.commit","true");
        properties.put("key.deserializer","org.apache.kafka.common.serialization.StringDeserializer");
        properties.put("value.deserializer","org.apache.kafka.common.serialization.StringDeserializer");

        // 入口点
        KafkaConsumer<String, String> consumer = new KafkaConsumer<>(properties);

        consumer.subscribe(Arrays.asList("my-replicated-topic"));

        while(true) {
            System.out.println("wait....");
            ConsumerRecords<String, String> records = consumer.poll(Duration.ofMillis(5000));
            for(ConsumerRecord record : records) {
                System.out.println(record.key() + "\t"
                + record.value() + "\t"
                        + record.topic() + "\t"
                        + record.offset() + "\t"
                        + record.partition()
                );
            }
        }

    }
}
