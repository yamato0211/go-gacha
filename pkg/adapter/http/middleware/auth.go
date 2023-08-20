package middleware

import (
	"go-gacha-system/pkg/domain/repository"
	"go-gacha-system/pkg/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type authMiddleware struct {
	repo repository.UserRepository
}

func NewAuthMiddleware(repo repository.UserRepository) *authMiddleware {
	return &authMiddleware{
		repo: repo,
	}
}

func (am *authMiddleware) Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("x-token")
		if token == "" {
			log.Print("x-token is empty")
			c.JSON(http.StatusBadRequest, "x-token is empty")
			return
		}

		userID, err := am.repo.SelectIDByToken(c, token)
		if err != nil {
			log.Print(err.Error())
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		utils.SetUserID(c, userID)
		c.Next()
	}
}
