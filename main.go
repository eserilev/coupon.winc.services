package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/eserilev/migration.winc.services/server"
)

func main() {
	var x os.Signal
	s := make(chan os.Signal, 1)
	d := make(chan bool, 1)
	fmt.Println("migration.winc.services")
	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	go func() {
		x = <-s
		d <- true
	}()
	fmt.Println(x.String())
	h := &http.Server{
		Addr:    ":4000",
		Handler: &server.HTTPServer{},
	}
	server.Start()
	go func() {
		if err := h.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatalf("ListenAndServe(): %v", err)
		}
	}()
	<-d
}
