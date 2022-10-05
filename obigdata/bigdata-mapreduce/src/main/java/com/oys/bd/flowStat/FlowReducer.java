package com.oys.bd.flowStat;

import com.oys.bd.flowStat.writable.FlowBean;
import org.apache.hadoop.io.Text;
import org.apache.hadoop.mapreduce.Reducer;
import org.apache.hadoop.yarn.webapp.hamlet2.HamletSpec;

import java.io.IOException;

/**
 * @Author ouyushun
 * @Date 2022/10/5
 * @Version 1.0
 */
public class FlowReducer extends Reducer<Text, FlowBean, Text, FlowBean> {

    private FlowBean outV = new FlowBean();

    @Override
    protected void reduce(Text key, Iterable<FlowBean> values, Reducer<Text, FlowBean, Text, FlowBean>.Context context) throws IOException, InterruptedException {
        long upTotal = 0;
        long downTotal = 0;
        for (FlowBean flowBean : values) {
            upTotal += flowBean.getUpFlow();
            downTotal += flowBean.getDownFlow();
        }

        outV.setUpFlow(upTotal);
        outV.setDownFlow(downTotal);
        outV.setSumFlow();

        //写出
        context.write(key, outV);
    }
}
