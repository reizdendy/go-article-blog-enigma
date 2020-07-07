package main

import (
	"blog-article-go/configs"
	"blog-article-go/master"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	db, _ := configs.ConnectDB()
	router := configs.CreateRouter()
	master.Init(db, router)
	configs.RunServer(router)
}
