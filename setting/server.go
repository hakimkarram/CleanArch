package settings

import (
	"github.com/kataras/iris/v12"
)

type server struct {
	app *iris.Application
}

type Server interface {
	Configure() error
	Start() error
	Launch() error
}

func EstablishServer() Server {
	application := iris.New()
	return &server{app: application}
}

// Configure This method uses to pass another configuration to the server
func (s *server) Configure() error {
	//TODO implement me
	panic("implement me")
}

// Start This method contains all server configurations, routers and registry
func (s *server) Start() error {
	_ = NewSettings(s.app)
	return nil
}

// Launch This method ise to start run all app server dependencies and listen to the port
func (s *server) Launch() error {
	err := s.app.Listen(":8081")
	if err != nil {
		return err
	}
	return nil
}
