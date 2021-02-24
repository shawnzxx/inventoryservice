package product

import (
	"fmt"
	"log"
	"time"

	"golang.org/x/net/websocket"
)

type message struct {
	Data string `json:"data"`
	Type string `json:"type"`
}

func productSocket(ws *websocket.Conn) {
	// we can verify that the origin is an allowed origin
	fmt.Printf("origin: %s\n", ws.Config().Origin)

	done := make(chan struct{})
	fmt.Println("new websocket connection established")
	go func(c *websocket.Conn) {
		for {
			var msg message
			if err := websocket.JSON.Receive(ws, &msg); err != nil {
				log.Println(err)
				break
			}
			fmt.Printf("received message: data: %s, type: %s\n", msg.Data, msg.Type)
		}
		close(done)
	}(ws)
loop:
	for {
		select {
		case <-done:
			fmt.Println("connection was closed, lets break out of here")
			break loop
		default:
			{
				fmt.Println("sending top 10 product list to the client")
				products, err := GetTopTenProducts()
				if err != nil {
					log.Println(err)
					break
				}
				if err := websocket.JSON.Send(ws, products); err != nil {
					log.Println(err)
					break
				}
				// pause for 10 seconds before sending again
				time.Sleep(10 * time.Second)
			}
		}
	}
	fmt.Println("closing the connection")
	defer ws.Close()
}
