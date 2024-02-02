package settings

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type dataBase struct {
	HostName string
	UserName string
	Password string
	DbName   string
	Port     string
}

type DataBase interface {
	Configure() string
	Launch() (*gorm.DB, error)
}

func NewDataBase(hName, uName, pass, dbName, port string) DataBase {
	return &dataBase{
		HostName: hName,
		UserName: uName,
		Password: pass,
		DbName:   dbName,
		Port:     port,
	}
}
func (d *dataBase) Configure() string {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
		d.HostName, d.UserName, d.Password, d.DbName, d.Port)
	return dsn
}

func (d *dataBase) Launch() (*gorm.DB, error) {
	dsn := d.Configure()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
