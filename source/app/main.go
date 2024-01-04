package main

import (
	"fmt"
	"go-clean-architecture/source/config"
	"log"
	"time"

	"github.com/gin-gonic/gin"

	_userRepo "go-clean-architecture/source/library/user/repository/mongo"
	_userUcase "go-clean-architecture/source/library/user/usecase"

	_userHttp "go-clean-architecture/source/library/user/controller"
)

func main() {
	r := gin.Default()

	timeoutContext := time.Duration(config.App.Config.GetInt("context.timeout")) * time.Second

	database := config.App.Mongo.Database(config.App.Config.GetString("mongo.name"))

	userRepo := _userRepo.NewMongoRepository(database)
	usrUsecase := _userUcase.NewUserUsecase(userRepo, timeoutContext)
	_userHttp.NewUserHandler(r, usrUsecase)

	appPort := fmt.Sprintf(":%s", config.App.Config.GetString("server.address"))
	log.Fatal(r.Run(appPort))
}
