package http

import (
	"errors"
	_ "go-gacha-system/pkg/domain/entity"
	"go-gacha-system/pkg/usecase"
	"go-gacha-system/pkg/usecase/schema"
	"go-gacha-system/pkg/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type gachaHandler struct {
	usecase usecase.IGachaUsecase
}

func NewGachaHandler(usecase usecase.IGachaUsecase) *gachaHandler {
	return &gachaHandler{
		usecase: usecase,
	}
}

// DrawGacha godoc
//
//	@Summary		ガチャ実行API
//	@Description	ガチャを引いてキャラクターを取得する処理
//	@Tags			gacha
//	@Accept			json
//	@Produce		json
//	@Param			x-token	header		string					true	"認証トークン"
//	@Param			body	body		schema.DrawGachaPayload	true	"ガチャを引く"
//	@Success		200		{object}	[]entity.Character
//	@Failure		400		{object}	error
//	@Failure		500		{object}	error
//	@Router			/gacha/draw [post]
func (gh *gachaHandler) DrawGacha() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := utils.GetUserIDFromContext(c)
		var drawGachaPayload schema.DrawGachaPayload
		if err := c.ShouldBindJSON(&drawGachaPayload); err != nil {
			log.Print(err.Error())
			log.Print(errors.New("should bind json failed").Error())
			c.JSON(http.StatusBadRequest, err.Error())
			return
		}
		characters, err := gh.usecase.DrawGacha(c, userID, drawGachaPayload.Count)
		if err != nil {
			log.Print(err.Error())
			log.Print(errors.New("usecase DrawGacha failed").Error())
			c.JSON(http.StatusInternalServerError, err.Error())
			return
		}
		c.JSON(http.StatusOK, characters)
	}
}
