package main

import (
	"fmt"
	"log"
	"net/url"

	"./websocket"
)

var origin = "http://live.66boss.com/"
var request = "ws://live.66boss.com:8080/ws?"

func main() {
	v := url.Values{}
	v.Set("liveid", "z1.abc.co")
	v.Set("name", "踩踩")
	v.Set("userid", "123345")
	rl := request + v.Encode()
	ws, err := websocket.Dial(rl, "", origin)
	if err != nil {
		log.Fatal(err)
	}

	message := []byte("hello, world!")
	_, err = ws.Write(message)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Send: %s\n", message)

	var msg = make([]byte, 512)
	_, err = ws.Read(msg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Receive: %s\n", msg)
}
