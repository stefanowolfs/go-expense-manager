package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/stefanowolf/go-expense-manager/internal/core/dto"
	"github.com/stefanowolf/go-expense-manager/internal/core/service"
	"github.com/stefanowolf/go-expense-manager/internal/infra/adapter"
	"github.com/stefanowolf/go-expense-manager/internal/infra/datasource"
	"github.com/stefanowolf/go-expense-manager/internal/infra/logger"
	"net/http"
)

func AddUserRoutes(rg *gin.RouterGroup) {
	userRouter := rg.Group("/users")

	userService := service.NewUserService(datasource.NewUserRepositoryDB(adapter.MYSQLDB))
	handler := UserHandler{service: userService}

	userRouter.GET("", handler.GetAllUsers)
	userRouter.POST("", handler.NewUser)
}

type UserHandler struct {
	service service.UserService
}

func (h UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.service.GetAll()
	if err != nil {
		c.JSON(err.Code, err.AsMessage())
	} else {
		c.JSON(http.StatusOK, users)
	}
}

func (h UserHandler) NewUser(c *gin.Context) {
	var body dto.NewUserRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	logger.Info("New user: " + body.Name)
	user, err := h.service.NewUser(body)
	if err != nil {
		c.JSON(err.Code, err.AsMessage())
	} else {
		c.JSON(http.StatusCreated, user)
	}
}
