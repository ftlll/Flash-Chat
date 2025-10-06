package models

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"sync"

	"github.com/fatih/set"
	"github.com/gorilla/websocket"
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	SrcId       uint64
	DstId       uint64
	Type        int
	Media       string //
	Content     string
	Image       string
	Url         string
	Description string
	Amount      int
}

func (table *Message) TableName() string {
	return "message"
}

type Node struct {
	Conn      *websocket.Conn
	DataQueue chan []byte
	GroupSets set.Interface
}

var clientMap map[int64]*Node = make(map[int64]*Node, 0)
var rwLock sync.RWMutex

func Chat(writer http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	id := query.Get("userId")
	userId, _ := strconv.ParseInt(id, 10, 64)
	// msgType := query.Get("Type")
	// targetId := query.Get("targetId")
	// context := query.Get("context")
	conn, err := (&websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}).Upgrade(writer, request, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	node := &Node{
		Conn:      conn,
		DataQueue: make(chan []byte),
		GroupSets: set.New(set.ThreadSafe),
	}

	rwLock.Lock()
	clientMap[userId] = node
	rwLock.Unlock()

	go sendProc(node)
	go receiveProc(node)
}

func sendProc(node *Node) {
	for {
		select {
		case data := <-node.DataQueue:
			err := node.Conn.WriteMessage(websocket.TextMessage, data)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}

func receiveProc(node *Node) {
	for {
		_, data, err := node.Conn.ReadMessage()
		if err != nil {
			fmt.Println(err)
			return
		}
		broadMsg(data)
		fmt.Println("[ws] <<<< ", data)
	}
}

var udpChan chan []byte = make(chan []byte)

func broadMsg(data []byte) {
	udpChan <- data
}

func init() {
	go udpSendProc()
	go udpReceiveProc()
}

// use UDP to send process
func udpSendProc() {
	con, err := net.DialUDP("udp", nil, &net.UDPAddr{
		IP:   net.IPv4(127, 0, 0, 1),
		Port: 3000,
	})
	defer con.Close()
	if err != nil {
		fmt.Println(err)
	}

	for {
		select {
		case data := <-udpChan:
			_, err := con.Write(data)
			if err != nil {
				fmt.Println(err)
				return
			}
		}
	}
}

func udpReceiveProc() {
	con, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4zero,
		Port: 3000,
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer con.Close()
	for {
		var buffer [512]byte
		n, err := con.Read(buffer[0:])
		if err != nil {
			fmt.Println(err)
			return
		}
		dispatch(buffer[0:n])
	}
}

func dispatch(data []byte) {
	msg := Message{}
	err := json.Unmarshal(data, &msg)
	fmt.Println(msg)
	if err != nil {
		fmt.Println(err)
		return
	}
	switch msg.Type {
	case 1:
		sendMsg(int64(msg.DstId), data)
		// case 2:
		// 	sendGroupMsg()
		// case 3:
		// 	sendAllMsg()
		// case 4:
	}
}

func sendMsg(userId int64, msg []byte) {
	rwLock.RLock()
	node, ok := clientMap[userId]
	rwLock.RUnlock()
	if ok {
		node.DataQueue <- msg
	}
}
