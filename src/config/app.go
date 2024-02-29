package config

import (
	"os"

	"github.com/joho/godotenv"
)

type App struct {
	Database struct {
		Hostname string
		Password string
		Username string
		DBName   string
    Port string
	}

	Server struct {
		Port string
		Host string
		Key  string
	}

	Admin struct {
		Email    string
		Password string
	}
}

var app *App

func GetConfig() *App {
  if app == nil {
    app = initConfig()
  }

  return app
}

func initConfig() *App {
	conf := App{}
	err := godotenv.Load()
  if err != nil {
    conf.Server.Host = ""
    conf.Server.Port = ""
    conf.Server.Key = ""
    
    conf.Database.Hostname = ""
    conf.Database.Username = ""
    conf.Database.Password = ""
    conf.Database.DBName = ""
    conf.Database.Port = ""

    conf.Admin.Email = ""
    conf.Admin.Password = ""

    return &conf
  }

  conf.Server.Host = os.Getenv("APP_HOST")
  conf.Server.Port = os.Getenv("APP_PORT")
  conf.Server.Key = os.Getenv("APP_KEY")

  conf.Database.Hostname = os.Getenv("DB_HOST")
  conf.Database.Username = os.Getenv("DB_USER")
  conf.Database.Password = os.Getenv("DB_PASS")
  conf.Database.DBName = os.Getenv("DB_NAME")
  conf.Database.Port = os.Getenv("DB_PORT")

  conf.Admin.Email = os.Getenv("ADMIN_EMAIL")
  conf.Admin.Password = os.Getenv("ADMIN_PASSWORD")
  return &conf
}
