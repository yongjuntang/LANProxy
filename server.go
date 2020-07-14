package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strconv"
	"time"
)

type Config struct {
	Network string
	Ip	string
	RemotePort uint32
	Heart	int //seconds
}

type Server struct {
	l net.Listener
	conn net.Conn
	config *Config
}

var S *Server

func LoadConfig() *Config{
	f,err := os.Open("./server.json")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	bdata,err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	c := Config{}
	err = json.Unmarshal(bdata,&c)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	return &c
}

func NewServer(){
	c := LoadConfig()
	S = new(Server)
	S.config = c
}

func (this *Server) Listen(){
	l,err := net.Listen(this.config.Network,this.config.Ip+":"+strconv.FormatUint(uint64(this.config.RemotePort),10))
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	this.l = l
}


func (this *Server)Dispatch(){
	for {
		conn := this.l.Accept()

	}
}


func (this *Server) Handle() {
	for {
		conn,err := this.l.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

	}
}


/**
提供
 */
func main(){
	go ListenA()
}

func ListenA(){
	l,err := net.Listen("tcp",":8000")
	if err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}

}


func handle
