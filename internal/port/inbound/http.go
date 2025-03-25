package driverport

type Route struct {
	Method      string
	Endpoint    string
	Handler     string
	Middlewares []string
}
