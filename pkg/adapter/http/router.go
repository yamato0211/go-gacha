package http

import (
	"go-gacha-system/pkg/adapter/http/middleware"
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
	//DI
	mySQLConn := infra.NewMySQLConnector()
	userRepository := mysql.NewUserRepository(mySQLConn.DB)
	characterRepository := mysql.NewCharacterRepository(mySQLConn.DB)
	gachaRepository := mysql.NewGachaRepository(mySQLConn.DB)
	userUsecase := usecase.NewUserUsecase(userRepository)
	characterUsecase := usecase.NewCharacterUsecase(characterRepository, userRepository)
	gachaUsecase := usecase.NewGachaUsecase(gachaRepository, characterRepository, userRepository)
	middleware := middleware.NewAuthMiddleware(userRepository)
	user := r.Group("/user")
	{
		handler := NewUserHandler(userUsecase)
		user.GET("/get", middleware.Authenticate(), handler.GetUser())
		user.POST("/create", handler.CreateUser())
		user.PUT("/update", middleware.Authenticate(), handler.UpdateName())
	}
	character := r.Group("/character")
	{
		handler := NewCharacterHandler(characterUsecase)
		character.GET("/list", middleware.Authenticate(), handler.GetCharacters())
	}
	gacha := r.Group("/gacha")
	{
		handler := NewGachaHandler(gachaUsecase)
		gacha.POST("/draw", middleware.Authenticate(), handler.DrawGacha())
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
