package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	var address string

	flag.StringVar(&address, "address", ":8080", "HTTP Server Address")
	flag.Parse()

	errC, err := run(address)
	if err != nil {
		log.Fatalf("Couldn't run: %s", err)
	}

	if err := <-errC; err != nil {
		log.Fatalf("Error while running: %s", err)
	}
}

func run(addr string) (<-chan error, error) {
	srv, err := newServer(serverConfig{
		Address: addr,
	})

	if err != nil {
		return nil, fmt.Errorf("newServer %w", err)
	}

	errC := make(chan error, 1)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGQUIT)
	go func() {
		<-ctx.Done()

		fmt.Println("[INFO] Shutdown signal received")

		// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
		ctxTimeout, cancel := context.WithTimeout(context.Background(), 15*time.Second)

		defer func() {
			stop()
			cancel()
			close(errC)
		}()

		srv.SetKeepAlivesEnabled(false)
		if err := srv.Shutdown(ctxTimeout); err != nil {
			errC <- err
		}

		fmt.Println("[INFO] Shutdown completed")
	}()

	// start the server
	go func() {
		fmt.Printf("[INFO] Listening and serving at %s\n", addr)

		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			errC <- err
		}
	}()

	return errC, nil
}

type serverConfig struct {
	Address string
}

func newServer(sc serverConfig) (*http.Server, error) {
	r := gin.Default()

	r.GET("/status", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	return &http.Server{
		Addr:         sc.Address,
		Handler:      r,
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}, nil
}
