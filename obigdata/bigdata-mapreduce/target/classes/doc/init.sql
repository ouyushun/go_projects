CREATE TABLE olist_customers_dataset (
        customer_id string comment '订单数据集的主键。每个订单都有一个唯一的customer_id',
        customer_unique_id string comment '客户的唯一标识符',
        customer_zip_code_prefix string comment '客户邮政编码的前五位数字',
        customer_city string comment '客户城市名称',
        customer_state string comment '客户状态'
) comment '客户及其位置的信息'
row format delimited
fields terminated by ','
stored as textfile;


CREATE TABLE eeee (
                                         customer_id string comment '订单数据集的主键。每个订单都有一个唯一的customer_id',
                                         customer_unique_id string comment '客户的唯一标识符',
                                         customer_zip_code_prefix string comment '客户邮政编码的前五位数字',
                                         customer_city string comment '客户城市名称',
                                         customer_state string comment '客户状态'
) comment '客户及其位置的信息'
row format delimited
fields terminated by ','
stored as textfile;

explain select * from row2colum;


load data local inpath '/home/ds_teacher/olist_customers_dataset.csv' overwrite into table olist_customers_dataset;
load data local inpath '/root/olist_customers_dataset.csv' overwrite into table olist_customers_dataset;

CREATE TABLE olist_geolocation_dataset (
                                           geolocation_zip_code_prefix string comment '邮政编码的前5位数字',
                                           geolocation_lat string comment '纬度',
                                           geolocation_lng string comment '经度',
                                           geolocation_city string comment '城市名称',
                                           geolocation_state string comment '州'
) comment '巴西邮政编码及其纬度/经度坐标信息'
row format delimited
fields terminated by ','
stored as textfile;

load data local inpath '/home/ds_teacher/olist_geolocation_dataset.csv' overwrite into table olist_geolocation_dataset;

CREATE TABLE olist_order_items_dataset (
                                           order_id string comment '订单ID',
                                           order_item_id int comment '序号，用于标识同一订单中包含的商品数量',
                                           product_id string comment '商品ID',
                                           seller_id string comment '商家ID',
                                           shipping_limit_date string comment '将订单处理到物流合作伙伴的卖家发货限制日期',
                                           price decimal(20,2) comment '商品价格',
                                           freight_value decimal(20,2) comment '运费，物品运费价值物品（如果订单包含多个物品，则运费价值将在物品之间分配）'
) comment '订单项明细(每个订单中购买的商品的数据)'
row format delimited
fields terminated by ','
stored as textfile;

load data local inpath '/home/ds_teacher/olist_order_items_dataset.csv' overwrite into table olist_order_items_dataset;

CREATE TABLE olist_order_payments_dataset (
                                              order_id string comment '订单ID',
                                              payment_sequential int comment '客户可以使用多种付款方式支付订单',
                                              payment_type string comment '客户选择的付款方式',
                                              payment_installments int comment '客户选择的分期付款数量',
                                              payment_value decimal(20,2) comment '交易价值'
) comment '订单付款数据'
row format delimited
fields terminated by ','
stored as textfile;

load data local inpath '/home/ds_teacher/olist_order_payments_dataset.csv' overwrite into table olist_order_payments_dataset;

CREATE TABLE olist_order_reviews_dataset (
                                             review_id string comment '评论ID',
                                             order_id string comment '订单ID',
                                             review_score int comment '客户满意度评分：1到5',
                                             review_comment_title string comment '用户评论标题，葡萄牙语',
                                             review_comment_message string comment '用户评论信息，葡萄牙语',
                                             review_creation_date string comment '显示向客户发送满意度调查的日期',
                                             review_answer_timestamp string comment '显示满意度调查答案时间戳'
) comment '客户评论数据'
row format delimited
fields terminated by ','
stored as textfile;

load data local inpath '/home/ds_teacher/olist_order_reviews_dataset.csv' overwrite into table olist_order_reviews_dataset;


CREATE TABLE olist_orders_dataset (
                                      order_id string comment '订单ID',
                                      customer_id string comment '订单对应的用户ID',
                                      order_status string comment '订单状态',
                                      order_purchase_timestamp string comment '显示购买时间戳',
                                      order_approved_at string comment '显示付款审批时间戳',
                                      order_delivered_carrier_date string comment '显示订单过帐时间戳。当它被处理给后勤合作伙伴时',
                                      order_delivered_customer_date string comment '显示客户的实际订单交货日期',
                                      order_estimated_delivery_date string comment '显示在购买时通知客户的预计交货日期'
) comment '订单信息'
row format delimited
fields terminated by ','
stored as textfile;

load data local inpath '/home/ds_teacher/olist_orders_dataset.csv' overwrite into table olist_orders_dataset;



CREATE TABLE olist_products_dataset (
                                        product_id string comment '商品ID',
                                        product_category_name string comment '产品的根类别',
                                        product_name_lenght int comment '从产品名称中提取的字符数',
                                        product_description_lenght int comment '从产品说明中提取的字符数',
                                        product_photos_qty int comment '产品发布数量',
                                        product_weight_g int comment '产品重量以克计',
                                        product_length_cm int comment '产品长度以厘米为单位',
                                        product_height_cm int comment '产品高度以厘米为单位',
                                        product_width_cm int comment '产品宽度以厘米为单位'
) comment '产品信息'
row format delimited
fields terminated by ','
stored as textfile;

load data local inpath '/home/ds_teacher/olist_products_dataset.csv' overwrite into table olist_products_dataset;


CREATE TABLE olist_sellers_dataset (
                                       seller_id string comment '卖家ID',
                                       seller_zip_code_prefix string comment '卖家邮政编码的前5位数字',
                                       seller_city string comment '卖方城市名称',
                                       seller_state string comment '卖方国家'
) comment '订单的卖家的数据'
row format delimited
fields terminated by ','
stored as textfile;

load data local inpath '/home/ds_teacher/olist_sellers_dataset.csv' overwrite into table olist_sellers_dataset;

CREATE TABLE product_category_name_translation (
                                                   product_category_name string comment '葡萄牙语的类别名称',
                                                   product_category_name_english string comment '英语的类别名称'
) comment '将商品名从葡萄牙语翻译为英语'
row format delimited
fields terminated by ','
stored as textfile;

load data local inpath '/home/ds_teacher/product_category_name_translation.csv' overwrite into table product_category_name_translation;
