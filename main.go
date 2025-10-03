package main

import (
	"flashchat/router"
	"flashchat/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMySQL()
	r := router.Router()
	r.Run()
}
