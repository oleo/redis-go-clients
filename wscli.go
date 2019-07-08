// Copyright 2015 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"flag"
	"log"
	"net/url"
	"os"
	"os/signal"
	"syscall"
	//"time"
	"fmt"

	"github.com/gorilla/websocket"
"github.com/go-redis/redis"
)

var addr = flag.String("addr", os.Args[1], "http service address")

func main() {
	flag.Parse()
	log.SetFlags(0)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: *addr, Path: "/nf"}
	log.Printf("connecting to %s", u.String())

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	//done := make(chan struct{})

	client := redis.NewClient(&redis.Options{ Addr: "192.168.0.201:6379", Password: "", DB:   0})

  defer client.Close()

  channel := os.Args[1]
  pubsub := client.Subscribe(channel)

  defer pubsub.Close()

	myc := make(chan os.Signal)
  signal.Notify(myc, os.Interrupt, syscall.SIGTERM)
  go func() {
		<-myc
		os.Exit(1)
  }()

  for {
		msg, err := pubsub.ReceiveMessage()
    	if err != nil {
    	  panic(err)
    	}
			if(len(msg.Payload) > 0) {
				fmt.Println("Got message, sending it forward: " + msg.Payload)
    		merr := c.WriteMessage(websocket.TextMessage, []byte(msg.Payload))
				fmt.Print("done....");
    		if merr != nil {
    		  panic(merr)
    		}
				fmt.Println("finished");
				msg.Payload=""
			}
  	}

}
