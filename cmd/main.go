package main

import (
	"ewangsong/goradius"
	"fmt"
	"log"
	"net"
	"os"

	_ "github.com/Go-SQL-Driver/MySQL"
)

func main() {
	var ch chan int
	ch = make(chan int)
	logfile, err := os.OpenFile("./goradius.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		os.Exit(-1)
	}
	defer logfile.Close()
	log.SetOutput(logfile)
	log.SetPrefix("[INFO]")
	log.SetFlags(log.Llongfile | log.Ldate | log.Ltime)

	goradius.DbLive() //判断数据库是否存在

	udpaddr, err := net.ResolveUDPAddr("udp", ":1812")
	if err != nil {
		fmt.Println(err)
	}

	udpconn, err2 := net.ListenUDP("udp", udpaddr)
	if err2 != nil {
		log.Println(err2)
	}
	go goradius.Server(udpconn)

	<-ch
}
