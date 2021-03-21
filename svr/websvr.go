package svr

import (
	"encoding/json"
	"fmt"
	"go_redis/db"
	"go_redis/redis"
	"log"
	"net/http"
)

func handle(w http.ResponseWriter, r *http.Request) {
	var users []db.User
	hasCache, err := redis.Exists("user")
	if hasCache {
		user_str, err := redis.GetK("user")
		if err != nil {
			log.Println(err.Error())
		}
		err = json.Unmarshal([]byte(user_str), &users)
		if err != nil {
			log.Println(err.Error())
		}
		log.Println("data from redis")
		log.Println(users)
	} else {
		users, err = db.GetUser()
		if err != nil {
			log.Println(err.Error())
			return
		}
		log.Println("data from db")
		log.Println(users)
		user_str, err := json.Marshal(users)
		if err != nil {
			log.Println(err.Error())
			return
		}
		ex_time := "15"
		reply, err := redis.SetKV("user", string(user_str), &ex_time)
		if err != nil {
			log.Println(err.Error())
		}
		log.Println(fmt.Sprintf("Redis Set Key: user reply: %v", reply))
	}
}
func StartWebSvr() {
	log.Println("The Web Server is running in localhost:9000/GetUser")
	http.HandleFunc("/GetUser", handle)
	http.ListenAndServe("localhost:9000", nil)
}
