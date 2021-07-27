package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/jmnelson12/distributed-world/main-entry/internal/server"
)

func main() {
	port := "8080"
	s := server.NewHTTPServer(fmt.Sprintf(":%s", port))

	// start the server
	go func() {
		// l.Info("Starting server on port 9090")
		fmt.Printf("[INFO] Starting server on port %s\n", port)

		err := s.ListenAndServe()
		if err != nil {
			// l.Error("Error starting server", "error", err)
			fmt.Println("[ERROR] Unable to start server", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c

	// log.Println("Got signal:", sig)
	fmt.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	s.Shutdown(ctx)
}
