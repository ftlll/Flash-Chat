package main

import (
	"flashchat/router"
	"flashchat/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMySQL()
	utils.InitRedis()

	r := router.Router()
	r.Run()
}
