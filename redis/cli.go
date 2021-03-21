package redis

import (
	"log"

	"github.com/Unknwon/goconfig"
	"github.com/gomodule/redigo/redis"
)

var addr string

func init() {
	cfg, err := goconfig.LoadConfigFile("redis/conf.ini")
	if err != nil {
		log.Fatalln(err.Error())
	}
	redis_cfg, err := cfg.GetSection("redis")
	if err != nil {
		log.Fatalln(err.Error())
	}
	addr = redis_cfg["host"] + ":" + redis_cfg["port"]
}

func exec(command string, args ...interface{}) (reply interface{}, err error) {
	con, err := redis.Dial("tcp", addr)
	if err != nil {
		return nil, err
	}
	defer con.Close()
	reply, err = con.Do(command, args...)
	if err != nil {
		log.Println("redis error:", err.Error())
		return nil, err
	}
	return reply, nil
}

func SetKV(key, val string, ex_second *string) (reply string, err error) {
	if ex_second != nil {
		return redis.String(exec("Set", key, val, "EX", *ex_second))
	}
	return redis.String(exec("Set", key, val))
}

func GetK(key string) (reply string, err error) {
	return redis.String(exec("Get", key))
}

func Exists(key string) (bool, error) {
	return redis.Bool(exec("EXISTS", key))
}
