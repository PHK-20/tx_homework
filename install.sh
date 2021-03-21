#!/bin/bash
# mysql conf
host="127.0.0.1"
port="3306"
db_user="root"
db_pw="123456"
db_name="go_demo"

raw_sql="
    CREATE DATABASE IF NOT EXISTS go_demo 
"
mysql -h${host} -P${port} -u${db_user} -p${db_pw} -e "${raw_sql}"

raw_sql="
    CREATE TABLE IF NOT EXISTS user_info (
        user VARCHAR(20) NOT NULL,
        sex VARCHAR(10)  NOT NULL,
        email VARCHAR(32),
        age INT,
        PRIMARY KEY (user)
    );
"

mysql -h${host} -P${port} -u${db_user} -p${db_pw} -D ${db_name} -e "${raw_sql}"

