package http

import (
	"go-gacha-system/pkg/usecase"
	"go-gacha-system/pkg/usecase/schema"
	"log"
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

// GetUser godoc
//
//	@Summary		ユーザ情報取得API
//	@Description	ユーザ情報を取得します。\n ユーザの認証と特定の処理はリクエストヘッダのx-tokenを読み取ってデータベースに照会をします。
//	@Tags			user
//	@Param			x-token	header	string	true	"認証トークン"
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	schema.NameResponse
//	@Failure		400	{object}	error
//	@Failure		500	{object}	error
//	@Router			/user/get [get]
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

// CreateUser godoc
//
//	@Summary		ユーザ情報作成API
//	@Description	ユーザ情報を作成します。\n ユーザの名前情報をリクエストで受け取り、ユーザIDと認証用のトークンを生成しデータベースへ保存します。
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			body	body		schema.CreateUserPayload	true	"userの新規作成"
//	@Success		200		{object}	schema.TokenResponse
//	@Failure		400		{object}	error
//	@Failure		500		{object}	error
//	@Router			/user/create [post]
func (uh *userHandler) CreateUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		var createUser schema.CreateUserPayload
		err := c.ShouldBindJSON(&createUser)
		if err != nil {
			log.Print(err.Error())
			c.JSON(http.StatusBadRequest, err.Error())
		}
		token, err := uh.usecase.CreateUser(c, createUser.Name)
		if err != nil {
			log.Print(err.Error())
			c.JSON(http.StatusInternalServerError, err.Error())
		}
		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	}
}

// UpdateName godoc
//
//	@Summary		ユーザ情報更新API
//	@Description	ユーザ情報の更新をします。
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			x-token	header		string						true	"認証トークン"
//	@Param			body	body		schema.CreateUserPayload	true	"userの更新"
//	@Success		200		{object}	schema.NameResponse
//	@Failure		400		{object}	error
//	@Failure		500		{object}	error
//	@Router			/user/update [put]
func (uh *userHandler) UpdateName() gin.HandlerFunc {
	return func(c *gin.Context) {
		var createUser schema.CreateUserPayload
		err := c.ShouldBindJSON(&createUser)
		if err != nil {
			log.Print(err.Error())
			c.JSON(http.StatusBadRequest, err.Error())
		}

		token := c.Request.Header.Get("x-token")
		if token == "" {
			log.Print("x-token not found")
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "x-token not found",
			})
		}

		err = uh.usecase.UpdateName(c, createUser.Name, token)
		if err != nil {
			log.Print(err.Error())
			c.JSON(http.StatusInternalServerError, err.Error())
		}
		c.Status(http.StatusOK)
	}
}
