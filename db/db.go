package db

import (
	"encoding/json"
	"errors"
	"go_redis/redis"
	"log"

	"github.com/Unknwon/goconfig"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var Db *sqlx.DB

func init() {
	cfg, err := goconfig.LoadConfigFile("db/conf.ini")
	if err != nil {
		log.Fatalln(err.Error())
	}
	db_cfg, err := cfg.GetSection("mysql")
	if err != nil {
		log.Fatalln(err.Error())
	}
	db_user := db_cfg["db_user"] + ":" + db_cfg["db_pw"]
	db_addr := db_cfg["db_host"] + ":" + db_cfg["db_port"]
	db_name := db_cfg["db_name"]
	Db, err = sqlx.Open("mysql", db_user+"@tcp("+db_addr+")/"+db_name)
	if err != nil {
		log.Fatalln("open mysql failed,", err)
	}
}

type User struct {
	User  string `db:"user"` //primary key
	Sex   string `db:"sex"`
	Email string `db:"email"`
	Age   int    `db:"age"`
}

func (item *User) GetUser() ([]User, error) {
	var users []User
	err := Db.Select(&users, "select user, sex, email, age from user_info")
	if err != nil {
		return nil, errors.New("db exec failed, " + err.Error())
	}
	return users, nil
}

func (item *User) Add() (*int64, error) {
	if len(item.User) == 0 {
		return nil, errors.New("db insert error: pk user is empty")
	}
	r, err := Db.Exec(
		"insert into user_info(user, sex, email,age) values(?, ?, ?,?)",
		item.User, item.Sex, item.Email, item.Age)
	if err != nil {
		return nil, errors.New("db exec failed, " + err.Error())
	}
	id, err := r.LastInsertId()
	if err != nil {
		return nil, errors.New("db exec failed, " + err.Error())
	}
	return &id, nil
}

func (item *User) Update() (*int64, error) {
	if len(item.User) == 0 {
		return nil, errors.New("db update error: pk user is empty")
	}
	res, err := Db.Exec(
		"update user_info set sex=?, email=?, age=? where user=?",
		item.Sex, item.Email, item.Age, item.User)
	if err != nil {
		return nil, errors.New("db exec failed, " + err.Error())
	}
	row, err := res.RowsAffected()
	if err != nil {
		return nil, errors.New("db exec failed, " + err.Error())
	}
	return &row, nil
}

func (item *User) Delete() (*int64, error) {
	if len(item.User) == 0 {
		return nil, errors.New("db del error: pk user is empty")
	}
	res, err := Db.Exec("delete from user_info where user=?", item.User)
	if err != nil {
		return nil, errors.New("db exec failed, " + err.Error())
	}
	row, err := res.RowsAffected()
	if err != nil {
		return nil, errors.New("db exec failed, " + err.Error())
	}
	return &row, nil
}

func (item *User) UpdateRedis() (*string, error) {
	users, err := item.GetUser()
	if err != nil {
		return nil, err
	}
	user_str, err := json.Marshal(users)
	if err != nil {
		return nil, err
	}
	ex_time := "15"
	reply, err := redis.SetKV("user", string(user_str), &ex_time)
	if err != nil {
		return nil, err
	}
	return &reply, nil
}
