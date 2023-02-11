use default ;

show tables ;
explain extended select name, count(*) from row2colum group by name, const limit 10;

