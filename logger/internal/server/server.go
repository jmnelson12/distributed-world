package server

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	Address string
}

func NewHTTPServer(sc HttpServer) (*http.Server, error) {
	r := SetupRouter()

	return &http.Server{
		Addr:         sc.Address,
		Handler:      r,
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}, nil
}

func SetupRouter() *gin.Engine {
	httpsrv := newHTTPServer()
	r := gin.Default()

	// may move to a handler once we introduce more endpoints
	r.GET("/status", httpsrv.handleStatus)

	return r
}

func newHTTPServer() *HttpServer {
	return &HttpServer{}
}

func (s *HttpServer) handleStatus(c *gin.Context) {
	c.JSON(200, gin.H{
		"status": "ok",
	})
}
