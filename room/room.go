package room

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/y0n0zawa/go-websocket-sample/client"
)

type Room struct {
	Forward chan []byte
	Join    chan *client.Client
	Leave   chan *client.Client
	Clients map[*client.Client]bool
}

func NewRoom() *Room {
	return &Room{
		Forward: make(chan []byte),
		Join:    make(chan *client.Client),
		Leave:   make(chan *client.Client),
		Clients: make(map[*client.Client]bool),
	}
}

func (r *Room) Run() {
	for {
		select {
		case client := <-r.Join:
			r.Clients[client] = true
		case client := <-r.Leave:
			delete(r.Clients, client)
			close(client.Send)
		case msg := <-r.Forward:
			for client := range r.Clients {
				select {
				case client.Send <- msg:
					// メッセージを送信する。
				default:
					// 送信に失敗したのでクライアントを切断する。
					delete(r.Clients, client)
					close(client.Send)
				}
			}
		}
	}
}

const (
	socketBufferSize  = 1024
	messageBufferSize = 256
)

var upgrader = &websocket.Upgrader{ReadBufferSize: socketBufferSize, WriteBufferSize: socketBufferSize}

func (r *Room) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	socket, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		log.Fatal("ServeHTTP", err)
		return
	}
	client := &client.Client{
		Socket:        socket,
		Send:          make(chan []byte, messageBufferSize),
		OnChatMessage: func(msg []byte) { r.Forward <- msg },
	}

	r.Join <- client
	defer func() { r.Leave <- client }()
	go client.Write()
	client.Read()
}
