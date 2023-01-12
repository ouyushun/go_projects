package com.oys.bd.commonfriends;

import org.apache.hadoop.io.Text;
import org.apache.hadoop.mapreduce.Reducer;

import java.io.IOException;

/**
 * @Author ouyushun
 * @Date 2022/10/30
 * @Version 1.0
 */
public class CFStep1Reducer extends Reducer<Text, Text, Text, Text> {

    @Override
    protected void reduce(Text key, Iterable<Text> values, Reducer<Text, Text, Text, Text>.Context context) throws IOException, InterruptedException {
        StringBuilder sb = new StringBuilder();
        for (Text person : values) {
            sb.append(person + ",");
        }
        context.write(key,new Text(sb.toString()));
    }
}
