package main

import (
	"fmt"
	"gps/model"
	"log"
	"net"
	"os"
)

func Listener(port string, c chan model.Message, quit chan int) {
	ServerAddr, err := net.ResolveUDPAddr("udp4", fmt.Sprintf(":%s", port))
	model.Error_check(err)

	ServerConn, err := net.ListenUDP("udp4", ServerAddr)
	model.Error_check(err)
	defer ServerConn.Close()

	buf := make([]byte, 200)
	for {
		n, addr, err := ServerConn.ReadFromUDP(buf)
		model.Error_check(err)
		m := model.Message{
			Size:   n,
			Msg:    string(buf[0:n]),
			Source: addr.IP.String(),
		}
		log.Printf("Received message [%s] from [%s]", m.Msg, m.Source)
		c <- m
	}
	close(c)
	quit <- 0
}

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}

	log.Printf("[Start] start server port:%s \n", port)
	c := make(chan model.Message, 100)
	quit := make(chan int)
	go Listener(port, c, quit)
	<-quit
}
