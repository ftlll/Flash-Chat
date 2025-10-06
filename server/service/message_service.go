package service

import (
	"flashchat/models"
	"flashchat/utils"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// avoid CSRF
var upGrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func SendMsg(c *gin.Context) {
	ws, err := upGrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(ws *websocket.Conn) {
		err = ws.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}(ws)
	fmt.Println("Client connected")
	MsgHandler(ws, c)
}

func MsgHandler(ws *websocket.Conn, c *gin.Context) {
	// keep long connectivity
	for {
		msg, err := utils.Subscribe(c, utils.PublishKey)
		if err != nil {
			fmt.Println("send message failed", err)
			return
		}
		fmt.Println("send message", msg)
		timestamp := time.Now().Format("2006-01-02 15:04:05")
		m := fmt.Sprintf("[ws][%s]:%s", timestamp, msg)
		err = ws.WriteMessage(1, []byte(m))
		if err != nil {
			fmt.Println(err)
		}
	}
}

func SendUserMsg(c *gin.Context) {
	// ws, err := upGrade.Upgrade(c.Writer, c.Request, nil)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// defer func(ws *websocket.Conn) {
	// 	err = ws.Close()
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		return
	// 	}
	// }(ws)
	// fmt.Println("Client connected")
	// MsgHandler(ws, c)
	models.Chat(c.Writer, c.Request)
}
