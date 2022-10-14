package main

import (
	"log"
	"net/http"
	"time"

	"github.com/fvbock/endless"
)

type MyServer struct {
}

// server 实现 ServeHTTP interface
func (s *MyServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// 实现路由匹配 todo
	reqPath := r.URL.Path
	log.Println(reqPath)
	time.Sleep(4 * time.Second)
	w.Write([]byte("hello world!"))
}

func NewServer() *MyServer {
	s := &MyServer{}

	return s
}

func main() {
	s := NewServer()
	// listen and serve 1
	//if err := http.ListenAndServe(":8080", s); err != nil {
	//	log.Printf("err: %v \n", err)
	//}

	// listen and serve 2
	// support graceful restart
	server := endless.NewServer(":8080", s)
	if err := server.ListenAndServe(); err != nil {
		log.Printf("err: %v \n", err)
	}
}
