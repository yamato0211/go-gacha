package http

import (
	"go-gacha-system/pkg/usecase"
	"go-gacha-system/pkg/usecase/schema"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	usecase usecase.IUserUsecase
}

func NewUserHandler(usecase usecase.IUserUsecase) *userHandler {
	return &userHandler{
		usecase: usecase,
	}
}

func (uh *userHandler) GetUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("x-token")
		if token == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "x-token not found",
			})
		}
		name, err := uh.usecase.GetUser(c, token)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
		}
		c.JSON(http.StatusOK, gin.H{
			"name": name,
		})
	}
}

func (uh *userHandler) CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var createUser schema.CreateUserPayload
		err := c.ShouldBindJSON(&createUser)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
		}
		token, err := uh.usecase.CreateUser(c, createUser.Name)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
		}
		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	}
}

func (uh *userHandler) UpdateName() gin.HandlerFunc {
	return func(c *gin.Context) {
		var createUser schema.CreateUserPayload
		err := c.ShouldBindJSON(&createUser)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
		}

		token := c.Request.Header.Get("x-token")
		if token == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "x-token not found",
			})
		}

		err = uh.usecase.UpdateName(c, createUser.Name, token)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err.Error())
		}
		c.Status(http.StatusOK)
	}
}
