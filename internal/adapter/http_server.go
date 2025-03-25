package adapter

import (
	"anonsurvey/logger"
	"fmt"

	"github.com/labstack/echo/v4"
)

type HTTPRouter interface {
	RegisterRoutes(e *echo.Echo)
}

type HTTPServer struct {
	engine *echo.Echo
	port   string
}

type HTTPServerOpt func(*HTTPServer)

func HTTPServerPort(port int) func(*HTTPServer) {
	return func(s *HTTPServer) {
		s.port = fmt.Sprintf(":%d", port)
	}
}

func NewHTTPServer(opts ...HTTPServerOpt) HTTPServer {
	s := HTTPServer{
		engine: echo.New(),
		port:   ":8080",
	}

	for _, opt := range opts {
		opt(&s)
	}

	return s
}

func (s *HTTPServer) Start(routers []HTTPRouter) {
	for _, router := range routers {
		router.RegisterRoutes(s.engine)
	}

	if err := s.engine.Start(s.port); err != nil {
		logger.Fatal("cannot start echo server", "error", err)
	}
}
