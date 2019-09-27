package main

import (
	_ "web_haoge/database"
	orm "web_haoge/database"
	"web_haoge/router"
)

func main() {
	defer orm.Eloquent.Close()
	router := router.InitRouter()
	router.Run(":80")
}
