GROUP BY a，b，c WITH CUBE
等价于:
GROUP BY a，b，c GROUPING SETS（（a，b，c），（a，b），（b，c），（a，c），（a），（b），（c），（））。

GROUP BY a，b，c WITH ROLLUP
等价于:
GROUP BY a，b，c GROUPING SETS（（a，b，c），（a，b），（a），（）



select
       case when id&1=1 then a.province_code else 'all' end as province_code,
       case when id&2=2 then a.city_code else 'all' end as city_code,
       case when id&4=4 then a.district else 'all' end as district,
       concat(if(id&1=1,'province-',''), if(id&2=2,'city-',''), if(id&4=4,'district-','')) as flag,
       all_num,
       male_num,
       femal_num
from (
         select
             province_code, city_code, district,
             sum(pop_cnt) as all_num, sum(man_cnt) as male_num, sum(female_cnt) as femal_num,
             GROUPING__ID as id
         from dww_china_national_population_census_7th
         where province_code = 110000
         group by province_code, city_code, district WITH CUBE
         ORDER BY id
    ) as a ;



select
    case when id&1=1 then a.province_code else 'all' end as province_code,
    case when id&2=2 then a.city_code else 'all' end as city_code,
    case when id&4=4 then a.district else 'all' end as district,
    concat(if(id&1=1,'province-',''), if(id&2=2,'city-',''), if(id&4=4,'district-','')) as flag,
    all_num,
    male_num,
    femal_num
from (
         select
             province_code, city_code, district,
             sum(pop_cnt) as all_num, sum(man_cnt) as male_num, sum(female_cnt) as femal_num,
             GROUPING__ID as id
         from dww_china_national_population_census_7th
         group by province_code, city_code, district WITH CUBE
         ORDER BY id
     ) as a ;



select
    province_code, city_code, district,
    sum(pop_cnt) as all_num, sum(man_cnt) as male_num, sum(female_cnt) as femal_num,
    GROUPING__ID
from dww_china_national_population_census_7th
group by province_code, city_code, district
GROUPING SETS ( province_code, (province_code, city_code), (city_code, district))
ORDER BY GROUPING__ID;





select * from (
    select province_code, sum(female_cnt) as femal_num, row_number() over (
        partition by province_code order by femal_num desc) as rnk
    from dww_china_national_population_census_7th
    ) t
where rnk < CEILING(max(rnk) / 10);





select
    `province_code`, `city_code`, `district`, `femal_num`, `id`,
    row_number() over(partition by id order by femal_num desc) as rank
from (select
                   province_code, city_code, district,
                   sum(female_cnt) as femal_num,
                   GROUPING__ID as id
               from dww_china_national_population_census_7th
               group by province_code, city_code, district
                   GROUPING SETS ( province_code, city_code, district)
               ORDER BY id desc
    ) as t
where rank < CEILING(max(rank) / 10);
;


select `province_code`, `city_code`, `district`, `femal_num`, `id` ,
       case when id&1=1 then 'top-省'
           when id&2=2 then 'top-市'
           when id&4=4 then 'top-县' else "" end as flag
       from (
                  select
                      `province_code`, `city_code`, `district`, `femal_num`, `id`,
                      row_number() over(partition by id order by femal_num desc) as rank,
                      count(1) over(partition by id) total_rank_count
                  from (
                           select
                               province_code, city_code, district,
                               sum(female_cnt) as femal_num,
                               GROUPING__ID as id
                           from dww_china_national_population_census_7th
                           group by province_code, city_code, district
                               GROUPING SETS (province_code, city_code, district)
                           ORDER BY id, femal_num asc
                ) as t1
        ) as t2
where  rank / total_rank_count <= 0.1;



select id, count(1) from (
                             select
                                 province_code, city_code, district,
                                 sum(female_cnt) as femal_num,
                                 GROUPING__ID as id
                             from dww_china_national_population_census_7th
                             group by province_code, city_code, district
                                 GROUPING SETS (province_code, city_code, district, (province_code,city_code, district))
                             ORDER BY id asc
                             ) t group by id ;
1	32
2	366
4	2810
7	2848