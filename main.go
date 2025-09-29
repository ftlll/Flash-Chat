package main

import (
	"flashchat/router"
)

func main() {
	r := router.Router()
	r.Run()
}
