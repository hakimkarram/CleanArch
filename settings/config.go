package settings

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type DatabaseConfig struct {
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Database string `mapstructure:"database"`
}
type config struct {
	Database DatabaseConfig `mapstructure:"database"`
}

func (c *config) GetHostName() string {
	return c.Database.Host
}

type IConfig interface {
	initConfig(string) (err error)
	WatchConfig()
	GetHostName() string
}

func NewConfig(env string) IConfig {
	cfg := new(config)
	if err := cfg.initConfig(env); err != nil {
		return nil
	}
	return cfg
}
func (c *config) initConfig(env string) (err error) {

	viper.SetConfigFile(env)
	fmt.Printf("running with %s environment...\n", env)
	err = viper.ReadInConfig()
	if err != nil {
		return fmt.Errorf("failed to read config file: %w", err)
	}

	err = viper.Unmarshal(c)
	if err != nil {
		return fmt.Errorf("failed to unmarshal config: %w", err)
	}
	c.WatchConfig()
	return
}
func (c *config) WatchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
	})
}
