package socket

import (
	"errors"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

const HOST = ":8082"

type SocketService struct {
	Upgrader websocket.Upgrader
	Client   *websocket.Conn
}

func NewSocketService() *SocketService {
	return &SocketService{
		Upgrader: websocket.Upgrader{
			ReadBufferSize:  1024,
			WriteBufferSize: 1024,
		},
	}
}

func (s *SocketService) Active() {
	log.Println("socket service active")
	log.Println("listen: " + HOST)

	http.HandleFunc("/ws", s.Ws)
	go func() {
		log.Fatal(http.ListenAndServe(HOST, nil))
	}()
}

func (s *SocketService) Ws(w http.ResponseWriter, r *http.Request) {
	s.Upgrader.CheckOrigin = func(r *http.Request) bool {
		return true
	}
	ws, err := s.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}
	log.Println("Client Connected")
	s.Client = ws
	go func() {
		for {
			_ = s.Say("not sleep")
			time.Sleep(time.Second * 1)
		}
	}()
}

func (c *SocketService) Say(message string) error {
	if c.Client != nil {
		err := c.Client.WriteMessage(1, []byte(message))
		if err != nil {
			return err
		}
		return nil
	}
	return errors.New("client not connected")
}

//func (s *SocketService) reader(conn *websocket.Conn) {
//	for {
//		// read in a message
//		messageType, p, err := conn.ReadMessage()
//		if err != nil {
//			log.Println(err)
//			return
//		}
//		// print out that message for clarity
//		fmt.Println(string(p))
//
//		if err := conn.WriteMessage(messageType, p); err != nil {
//			log.Println(err)
//			return
//		}
//
//	}
//}
