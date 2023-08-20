package http

import (
	"go-gacha-system/pkg/infra"
	"go-gacha-system/pkg/infra/mysql"
	"go-gacha-system/pkg/usecase"

	_ "go-gacha-system/docs"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.Use(cors.Default())

	mySQLConn := infra.NewMySQLConnector()
	userRepository := mysql.NewUserRepository(mySQLConn.DB)
	userUsecase := usecase.NewUserUsecase(userRepository)
	user := r.Group("/user")
	{
		handler := NewUserHandler(userUsecase)
		user.GET("/get", handler.GetUser())
		user.POST("/create", handler.CreateUser())
		user.PUT("/update", handler.UpdateName())
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
