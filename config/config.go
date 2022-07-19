package config

import (
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Application struct {
	Host          string `yaml:"host"`
	Port          string `yaml:"port"`
	ReadTimeout   int    `yaml:"read_timeout"`
	WriterTimeout int    `yaml:"writer_timeout"`
}

type Mysql struct {
	DSN             string `yaml:"Dsn"`
	MaxOpenConn     int    `yaml:"MaxOpenConn"` // open pool
	MaxIdleConn     int    `yaml:"MaxIdleConn"` // idle pool
	ConnMaxLifeTime int    `yaml:"ConnMaxLifeTime"`
	SlowThreshold   int    `yaml:"SlowThreshold"`
}

type Redis struct {
	Addr         string        `yaml:"Addr"`
	Password     string        `yaml:"Password"`
	DB           int           `yaml:"DB"`
	MinIdleConn  int           `yaml:"MinIdleConn"`
	DialTimeout  time.Duration `yaml:"DialTimeout"`
	ReadTimeout  time.Duration `yaml:"ReadTimeout"`
	WriteTimeout time.Duration `yaml:"WriteTimeout"`
	PoolSize     int           `yaml:"PoolSize"`
	PoolTimeout  time.Duration `yaml:"PoolTimeout"`
}

type Auth struct {
	JwtSecret string `yaml:"jwt_secret"`
}

type Configs struct {
	Application Application `yaml:"application"`
	Mysql       Mysql       `yaml:"mysql"`
	Redis       Redis       `yaml:"redis"`
	Auth        Auth        `yaml:"auth"`
}

func loadConfig() (settings Configs) {
	var config Configs
	viper.SetConfigType("yml")
	viper.SetConfigFile("config/dev.yml")
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	if err := viper.Unmarshal(&config); err != nil {
		panic(err)
	}
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		if err := viper.Unmarshal(&config); err != nil {
			panic(err)
		}
	})
	return config
}

var (
	ApplicationConfig = Application{}
	MysqlConfig       = Mysql{}
	RedisConfig       = Redis{}
	AuthConfig        = Auth{}
)

func init() {
	config := loadConfig()
	ApplicationConfig = config.Application
	MysqlConfig = config.Mysql
	RedisConfig = config.Redis
	AuthConfig = config.Auth
}
