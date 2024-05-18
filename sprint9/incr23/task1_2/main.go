package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	var srv = http.Server{Addr: ":8080"}
	idleConnsClosed := make(chan struct{})
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt)
	go func() {
		<-sig
		if err := srv.Shutdown(context.Background()); err != nil {
			log.Printf("shutdownd: %v", err)
		}
		close(idleConnsClosed)
	}()

	if err := srv.ListenAndServe(); err != http.ErrServerClosed {
		// ошибки запуска или остановки Listener
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	}
	<-idleConnsClosed
	fmt.Println("Server Shutdown gracefully")
}
