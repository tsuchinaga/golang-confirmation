package main

import (
	"context"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Llongfile)

	ln, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		log.Fatalln(err)
	}
	defer ln.Close()

	srv := &http.Server{Handler: http.FileServer(new(fs))}
	go func() {
		if err := srv.Serve(ln); err != nil && err != http.ErrServerClosed {
			log.Println(err)
		}
	}()

	<-time.After(3 * time.Second)
	if err := srv.Shutdown(context.Background()); err != nil {
		log.Println(err)
	}
	<-time.After(3 * time.Second)
}

type fs struct{}

func (f fs) Open(name string) (http.File, error) {
	panic("implement me")
}
