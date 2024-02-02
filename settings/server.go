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
	Db := NewDataBase("localhost", "postgres", "shahed1986", "CleanArchDB", "5432")
	DbLaunched, err := Db.Launch()
	if err != nil {
		return err
	}
	Settings(s.app, DbLaunched)
	return nil
}

//func (s *server) LaunchDb() (*gorm.DB, error) {
//	dsn := "host=localhost user=postgres password=shahed1986 dbname=CleanArchDB port=5432 sslmode=disable TimeZone=Asia/Shanghai"
//	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
//	if err != nil {
//		return nil, err
//	}
//	return db, nil
//}

// Launch This method is to start run all app server dependencies and listen to the port
func (s *server) Launch() error {
	err := s.app.Listen(":8081")
	if err != nil {
		return err
	}
	return nil
}
