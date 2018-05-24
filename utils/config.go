package utils

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/getsentry/raven-go"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/configor"
)

// Config is a struct for backend config
type Config struct {
	APPName string `default:"Gin App"`

	Server struct {
		Host      string `default:"127.0.0.1"`
		Port      string `default:"9012"`
		SecretKey string `default:"SecretKey"`
	}

	Mongo struct {
		Host     string `default:"127.0.0.1"`
		Port     string `default:"27017"`
		DB       string `default:"db"`
		Username string `default:"username"`
		Password string `default:"password"`
	}

	Dayu struct {
		AccessID  string `required:"true"`
		AccessKey string `required:"true"`
		Sign      string `default:""`
	}

	QcloudSMS struct {
		AppID  string `default:""`
		AppKey string `default:""`
		Sign   string `default:""`
	}

	Redis struct {
		Host     string `default:"127.0.0.1"`
		Port     string `default:"6379"`
		Password string `default:""`
		DB       int    `default:"db"`
	}

	Root struct {
		Username string `required:"true"`
		Password string `required:"true"`
	}

	Sentry struct {
		SDN string `default:""`
	}
}

// LoadConfiguration is a function to load cfg from file
func LoadConfiguration() Config {
	path, err := os.Getwd()

	if err != nil {
		raven.CaptureErrorAndWait(err, nil)
		log.Panicf("[loadAppConfig]: %s\n", err)
	}

	switch gin.Mode() {
	case "release":
		path = strings.Replace(path, "test", "", -1) + "/config.deploy.yml"
	case "debug":
		path = strings.Replace(strings.Replace(path, "test", "", -1), "/handler", "", -1) + "/config.yml"
	case "test":
		fmt.Println("Start test", gin.Mode())
		path = strings.Replace(strings.Replace(path, "test", "", -1), "/handler", "", -1) + "/config.yml"
	}

	var config Config
	configFile, err := os.Open(path)
	defer configFile.Close()
	if err != nil {
		log.Printf("[loadAppConfig]: %s\n", err)
	}

	configor.Load(&config, path)
	if count := len([]rune(config.Server.SecretKey)); count <= 32 {
		for i := 1; i <= 32-count; i++ {
			config.Server.SecretKey += "="
		}
	} else {
		config.Server.SecretKey = string([]byte(config.Server.SecretKey)[:32])
	}

	if gin.Mode() == "test" {
		config.Mongo.Host = os.Getenv("MONGO_HOST")
		config.Mongo.Port = os.Getenv("MONGO_PORT")
		config.Mongo.Username = os.Getenv("MONGO_USERNAME")
		config.Mongo.Password = os.Getenv("MONGO_PASSWORD")
		config.Mongo.DB = os.Getenv("MONGO_DB")
		config.Root.Username = os.Getenv("ROOT_USERNAME")
		config.Root.Password = os.Getenv("ROOT_PASSWORD")
	}

	return config
}

// AppConfig is a struct loaded from config file
var AppConfig = LoadConfiguration()
