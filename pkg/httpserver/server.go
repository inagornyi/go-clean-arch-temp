package httpserver

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type HttpServer struct {
	server *http.Server
	notify chan error
}

func NewHttpServer() *HttpServer {
	router := gin.Default()
	return &HttpServer{
		server: &http.Server{
			Addr:    ":8080",
			Handler: router,
		},
		notify: make(chan error, 1),
	}
}

func (s *HttpServer) Run() error {
	go func() {
		s.notify <- s.server.ListenAndServe()
		close(s.notify)
	}()
	return nil
}

func (s *HttpServer) Notify() <-chan error {
	return s.notify
}

func (s *HttpServer) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	return s.server.Shutdown(ctx)
}
