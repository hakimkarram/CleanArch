package setting

import (
	"github.com/kataras/iris/v12"
)

type server struct {
	app *iris.Application
}

type Server interface {
	ServerConfiguration()
	Start() error
}

func NewServer() Server {
	application := iris.New()
	return &server{app: application}
}

// ServerConfiguration This method uses to pass another configuration to the server
func (s server) ServerConfiguration() {
	//TODO implement me
	panic("implement me")
}

func (s server) Start() error {
	err := s.app.Listen(":8081")
	if err != nil {
		return err
	}
	return nil
}
