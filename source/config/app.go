package config

import (
	"database/sql"
	"go-clean-architecture/mongo"

	"github.com/spf13/viper"
)

var (
	App *Application
)

type Application struct {
	Config *viper.Viper
	MySql  *sql.DB
	Mongo  mongo.Client
}

func init() {
	AppInit()
}

func AppInit() {
	App = &Application{}
	App.Config = InitConfig()
	App.Mongo = InitMongoDatabase()
}
