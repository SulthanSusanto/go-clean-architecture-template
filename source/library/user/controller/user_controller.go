package controller

import (
	"context"
	"fmt"
	"go-clean-architecture/source/entities"
	"math"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	validator "gopkg.in/go-playground/validator.v9"
)

type ResponseError struct {
	Message string `json:"message"`
}

type UserHandler struct {
	UsrUsecase entities.UserUsecase
}

func NewUserHandler(r *gin.Engine, uu entities.UserUsecase) {
	handler := &UserHandler{
		UsrUsecase: uu,
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

	if err := c.ShouldBindJSON(&usr); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	var ok bool
	if ok, err = isRequestValid(&usr); !ok {

		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		fmt.Println("test22")
		return
	}

	fmt.Println("test299")

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	result, err := user.UsrUsecase.InsertOne(ctx, &usr)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (user *UserHandler) FindOne(c *gin.Context) {
	id := c.Query("id")

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	result, err := user.UsrUsecase.FindOne(ctx, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (user *UserHandler) GetAll(c *gin.Context) {
	type Response struct {
		Total       int64           `json:"total"`
		PerPage     int64           `json:"per_page"`
		CurrentPage int64           `json:"current_page"`
		LastPage    int64           `json:"last_page"`
		From        int64           `json:"from"`
		To          int64           `json:"to"`
		User        []entities.User `json:"users"`
	}

	var (
		res   []entities.User
		count int64
	)

	rp, err := strconv.ParseInt(c.DefaultQuery("rp", "25"), 10, 64)
	if err != nil {
		rp = 25
	}

	page, err := strconv.ParseInt(c.DefaultQuery("p", "1"), 10, 64)
	if err != nil {
		page = 1
	}

	filters := bson.D{{"name", primitive.Regex{Pattern: ".*" + c.Query("name") + ".*", Options: "i"}}}

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	res, count, err = user.UsrUsecase.GetAllPagination(ctx, rp, page, filters, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	result := Response{
		Total:       count,
		PerPage:     rp,
		CurrentPage: page,
		LastPage:    int64(math.Ceil(float64(count) / float64(rp))),
		From:        page*rp - rp + 1,
		To:          page * rp,
		User:        res,
	}

	c.JSON(http.StatusOK, result)
}

func (user *UserHandler) UpdateOne(c *gin.Context) {
	id := c.Query("id")

	var (
		usr entities.User
		err error
	)

	if err := c.ShouldBindJSON(&usr); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
		return
	}

	ctx := c.Request.Context()
	if ctx == nil {
		ctx = context.Background()
	}

	result, err := user.UsrUsecase.UpdateOne(ctx, &usr, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
