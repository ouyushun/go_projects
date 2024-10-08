package com.oys.bd.commonfriends;

import org.apache.hadoop.io.Text;
import org.apache.hadoop.mapreduce.Reducer;

import java.io.IOException;
import java.util.ArrayList;
import java.util.List;

public class CFStep2Reducer extends Reducer<Text, Text, Text, Text> {
    @Override
    protected void reduce(Text key, Iterable<Text> values, Reducer<Text, Text, Text, Text>.Context context) throws IOException, InterruptedException {
        List<String> valArr = new ArrayList();
        for (Text person : values) {
            valArr.add(person.toString());
        }
        context.write(key,new Text(String.join(",", valArr)));
    }
}
