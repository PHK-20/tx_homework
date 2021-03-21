package db

import (
	"errors"
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
	db_name := db_cfg["db_name"]

	Db, err = sqlx.Open("mysql", db_user+"@tcp(127.0.0.1:3306)/"+db_name)
	if err != nil {
		log.Fatalln("open mysql failed,", err)
	}
}

type User struct {
	User  string `db:"user"`
	Sex   string `db:"sex"`
	Email string `db:"email"`
	Age   int    `db:"age"`
}

func GetUser() ([]User, error) {
	var users []User
	err := Db.Select(&users, "select user, sex, email, age from user_info")
	if err != nil {
		return nil, errors.New("exec failed, " + err.Error())
	}
	return users, nil
}
func AddUser(item *User) (*int64, error) {
	r, err := Db.Exec(
		"insert into user_info(user, sex, email,age) values(?, ?, ?,?)",
		item.User, item.Sex, item.Email, item.Age)
	if err != nil {
		return nil, errors.New("exec failed, " + err.Error())
	}
	id, err := r.LastInsertId()
	if err != nil {
		return nil, errors.New("exec failed, " + err.Error())
	}
	return &id, nil
}
