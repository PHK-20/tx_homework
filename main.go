package main

import (
	"go_redis/db"
	"go_redis/svr"
)


func main() {
	defer db.Db.Close()
	svr.StartWebSvr()
}
