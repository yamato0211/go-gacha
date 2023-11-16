package http

import (
	"context"
	"go-gacha-system/pkg/adapter/http/middleware"
	"go-gacha-system/pkg/infra"
	"go-gacha-system/pkg/infra/mysql"
	"go-gacha-system/pkg/infra/redis"
	"go-gacha-system/pkg/usecase"
	"log"

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
	redisConn := infra.NewRedisConnector()
	userRepository := mysql.NewUserRepository(mySQLConn.DB)
	characterRepository := mysql.NewCharacterRepository(mySQLConn.DB)
	gachaRepository := mysql.NewGachaRepository(mySQLConn.DB)
	rankingRepository := redis.NewRankingRepository(redisConn.Client)
	userUsecase := usecase.NewUserUsecase(userRepository)
	characterUsecase := usecase.NewCharacterUsecase(characterRepository, userRepository)
	gachaUsecase := usecase.NewGachaUsecase(gachaRepository, characterRepository, userRepository)
	rankingUsecase := usecase.NewRankingUsecase(characterRepository, rankingRepository)
	middleware := middleware.NewAuthMiddleware(userRepository)

	//Redis Insert
	if err := rankingUsecase.InsertRedisRanking(context.TODO()); err != nil {
		log.Fatal(err)
	}

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
	ranking := r.Group("/ranking")
	{
		handler := NewRankingHandler(rankingUsecase)
		ranking.GET("/redis", handler.GetRedisRanking())
		ranking.GET("/mysql", handler.GetMysqlRanking())
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}
