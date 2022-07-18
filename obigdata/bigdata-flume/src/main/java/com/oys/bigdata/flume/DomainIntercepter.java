package com.oys.bigdata.flume;

import org.apache.flume.Context;
import org.apache.flume.Event;
import org.apache.flume.interceptor.Interceptor;

import java.util.ArrayList;
import java.util.List;
import java.util.Map;

/**
 * @Author ouyushun
 * @Date 2022/7/17
 * @Version 1.0
 */
public class DomainIntercepter implements Interceptor {

    List<Event> events;
    //初始化
    @Override
    public void initialize() {
        events = new ArrayList<>();
    }

    @Override
    public Event intercept(Event event) {
        Map<String, String> header = event.getHeaders();
        String body = new String(event.getBody());
        if (body.contains("imooc")) {
            header.put("type", "imooc");
        } else {
            header.put("type", "other");
        }
        return event;
    }

    @Override
    public List<Event> intercept(List<Event> list) {
        events.clear();
        for (Event event : list) {
            events.add(intercept(event));
        }
        return events;
    }

    @Override
    public void close() {
        events = null;
    }

    public static class Builder implements Interceptor.Builder {

        @Override
        public Interceptor build() {
            return new DomainIntercepter();
        }

        @Override
        public void configure(Context context) {

        }
    }
}
