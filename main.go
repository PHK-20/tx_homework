package main

import (
	"go_redis/db"
	"go_redis/svr"
	"log"
)

func prepare() {
	new_user := db.User{
		User:  "herry",
		Sex:   "boy",
		Email: "herry@foxmail.com",
		Age:   20,
	}
	_, err := db.AddUser(&new_user)
	if err != nil {
		log.Println(err.Error())
	}
	new_user = db.User{
		User:  "marry",
		Sex:   "girl",
		Email: "marry@foxmail.com",
		Age:   18,
	}
	_, err = db.AddUser(&new_user)
	if err != nil {
		log.Println(err.Error())
	}
}

func main() {
	defer db.Db.Close()
	prepare()
	svr.StartWebSvr()

}
