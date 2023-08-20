package utils

import (
	"github.com/gin-gonic/gin"
)

type key string

const (
	userIDKey key = "userID"
)

func SetUserID(c *gin.Context, userID int64) {
	c.Set(string(userIDKey), userID)
}

func GetUserIDFromContext(c *gin.Context) int64 {
	userID := c.MustGet(string(userIDKey)).(int64)
	return userID
}
