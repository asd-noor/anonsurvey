package driverport

type HTTPServer struct{}

func NewHTTPServer() {}

type Router interface {
	Routes() []Route
}

func (s *HTTPServer) RegisterRoutes(rs ...Router) {}
