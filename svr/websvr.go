package svr

import (
	"encoding/json"
	"errors"
	"fmt"
	"go_redis/db"
	"go_redis/redis"
	"log"
	"net/http"
	"strconv"
)

func GetUserSvr(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			e := err.(error)
			log.Println(e.Error())
		}
	}()
	var users []db.User
	hasCache, err := redis.Exists("user")
	if hasCache {
		user_str, err := redis.GetK("user")
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal([]byte(user_str), &users)
		if err != nil {
			panic(err)
		}
		log.Println(fmt.Sprintf("data from redis\n%v", users))
	} else {
		item := db.User{}
		users, err = item.GetUser()
		if err != nil {
			panic(err)
		}
		log.Println(fmt.Sprintf("data from db\n%v", users))
		_, err = item.UpdateRedis()
		if err != nil {
			panic(err)
		}
	}
}

func AddUserSvr(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			e := err.(error)
			log.Println(e.Error())
			w.Write([]byte(e.Error()))
		}
	}()
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}
	user_age, _ := strconv.Atoi(r.Form.Get("age"))
	user := db.User{
		User:  r.FormValue("user"),
		Age:   user_age,
		Sex:   r.FormValue("sex"),
		Email: r.FormValue("email"),
	}
	_, err = user.Add()
	if err != nil {
		panic(err)
	}
	log.Println(fmt.Sprintf("Add User: %v", user))
	_, err = user.UpdateRedis()
	if err != nil {
		panic(err)
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			e := err.(error)
			log.Println(e.Error())
			w.Write([]byte(e.Error()))
		}
	}()
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}
	user_age, _ := strconv.Atoi(r.Form.Get("age"))
	user := db.User{
		User:  r.FormValue("user"),
		Age:   user_age,
		Sex:   r.FormValue("sex"),
		Email: r.FormValue("email"),
	}
	fmt.Println(user)
	rows, err := user.Update()
	if *rows == 0 {
		panic(errors.New(fmt.Sprintf("do not exists user: %v", user)))
	}
	if err != nil {
		panic(err)
	}
	log.Println(fmt.Sprintf("Update User: %v", user))
	_, err = user.UpdateRedis()
	if err != nil {
		panic(err)
	}

}

func DelUser(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			e := err.(error)
			log.Println(e.Error())
			w.Write([]byte(e.Error()))
		}
	}()
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}
	user := db.User{
		User: r.FormValue("user"),
	}
	rows, err := user.Delete()
	if *rows == 0 {
		panic(errors.New(fmt.Sprintf("do not exists user: %v", user)))
	}
	if err != nil {
		panic(err)
	}
	log.Println(fmt.Sprintf("Del User: %v", user))
	_, err = user.UpdateRedis()
	if err != nil {
		panic(err)
	}
}
func StartWebSvr() {
	log.Println("The Web Server is running in localhost:9000/GetUser")
	http.HandleFunc("/GetUser", GetUserSvr)
	http.HandleFunc("/AddUser", AddUserSvr)
	http.HandleFunc("/UpdateUser", UpdateUser)
	http.HandleFunc("/DelUser", DelUser)

	http.ListenAndServe("localhost:9000", nil)
}
