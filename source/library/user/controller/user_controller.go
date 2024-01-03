package controller

import (
	"go-clean-architecture/source/entities"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

type ResponseError struct {
	Message string `json:"message"`
}

type UserHandler struct {
	UserUsecase entities.UserUsecase
}

func NewUserHandler(r *gin.Context, uu entities.UserUsecase) {
	handler := &UserHandler{
		UserUsecase: uu,
	}

	r.POST("/user", handler.InsertOne)
	r.GET("/user", handler.FindOne)
	r.GET("/users", handler.GetAll)
	r.PUT("/user", handler.UpdateOne)
}

func isRequestValid(m *entities.User) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (user *UserHandler) InsertOne(c *gin.Context) {
	var (
		usr entities.User
		err error
	)

	err = c.Bind(&usr)

}
