package server

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func NewHTTPServer(addr string) *http.Server {
	router := gin.Default()

	router.GET("/status", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	return &http.Server{
		Addr:         addr,
		Handler:      router,
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}
}
