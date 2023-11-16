package http

import (
	_ "go-gacha-system/pkg/domain/entity"
	"go-gacha-system/pkg/usecase"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type rankingHandler struct {
	usecase usecase.IRankingUsecase
}

func NewRankingHandler(usecase usecase.IRankingUsecase) *rankingHandler {
	return &rankingHandler{
		usecase: usecase,
	}
}

// GetRedisRanking godoc
//
//	@Summary		Redisランキング取得API
//	@Description	Redisからランキングを取得する処理
//	@Tags			ranking
//	@Accept			json
//	@Produce		json
//	@Param			limit	query		string	true	"取得個数"
//	@Success		200		{object}	[]entity.Ranking
//	@Failure		400		{object}	error
//	@Failure		500		{object}	error
//	@Router			/ranking/redis [get]
func (rh *rankingHandler) GetRedisRanking() gin.HandlerFunc {
	return func(c *gin.Context) {
		limitParam := c.DefaultQuery("limit", "0")
		limit, err := strconv.Atoi(limitParam)
		if err != nil {
			log.Print(err.Error())
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		ranking, err := rh.usecase.GetRedisRanking(c, int64(limit))
		if err != nil {
			log.Print(err)
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, ranking)
	}
}

// GetMysqlRanking godoc
//
//	@Summary		Mysqlランキング取得API
//	@Description	Mysqlからランキングを取得する処理
//	@Tags			ranking
//	@Accept			json
//	@Produce		json
//	@Param			limit	query		string	true	"取得個数"
//	@Success		200		{object}	[]entity.Ranking
//	@Failure		400		{object}	error
//	@Failure		500		{object}	error
//	@Router			/ranking/mysql [get]
func (rh *rankingHandler) GetMysqlRanking() gin.HandlerFunc {
	return func(c *gin.Context) {
		limitParam := c.DefaultQuery("limit", "0")
		limit, err := strconv.Atoi(limitParam)
		if err != nil {
			log.Print(err.Error())
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		ranking, err := rh.usecase.GetMysqlRanking(c, int64(limit))
		if err != nil {
			log.Print(err)
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, ranking)
	}
}
