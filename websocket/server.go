package main

import (
	"github.com/gorilla/websocket"
	"net/http"
	"./impl"
	"time"
)

var (
	upgrader = websocket.Upgrader{
		// 允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func wsHandler (writer http.ResponseWriter, request *http.Request) {
	var (
		wsConn *websocket.Conn
		err error
		data []byte
		conn *impl.Connection
	)

	if wsConn, err = upgrader.Upgrade(writer, request, nil); err != nil {
		return
	}

	if conn, err = impl.InitConnection(wsConn); err != nil {
		goto ERR
	}

	go func() {
		var (
			err error
		)
		for {
			if err = conn.WriteMessage([]byte("heartbeat")); err != nil {
				return
			}
			time.Sleep(1 * time.Second)
		}
	}()

	for {
		if data, err = conn.ReadMessage(); err != nil {
			goto ERR
		}
		if err = conn.WriteMessage(data); err != nil {
			goto ERR
		}
	}

ERR:
	// TODO: 关闭连接
	conn.Close()
}

func main () {
	http.HandleFunc("/ws", wsHandler)
	http.ListenAndServe(":7777", nil)
}
