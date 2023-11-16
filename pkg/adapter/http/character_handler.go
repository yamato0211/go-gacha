package http

import (
	"go-gacha-system/pkg/usecase"
	"go-gacha-system/pkg/utils"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type characterHandler struct {
	usecase usecase.ICharacterUsecase
}

func NewCharacterHandler(usecase usecase.ICharacterUsecase) *characterHandler {
	return &characterHandler{
		usecase: usecase,
	}
}

//	GetCharacters godoc
//
//	@Summary		キャラクター関連API
//	@Description	ユーザが所持しているキャラクター一覧情報を取得
//	@Tags			character
//	@Accept			json
//	@Produce		json
//	@Param			x-token	header		string	true	"認証トークン"
//	@Success		200		{object}	[]entity.Character
//	@Failure		400		{object}	error
//	@Failure		500		{object}	error
//	@Router			/character/list [get]
func (ch *characterHandler) GetCharacters() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := utils.GetUserIDFromContext(c)
		characters, err := ch.usecase.GetCharacters(c, userID)
		if err != nil {
			log.Print(err.Error())
			c.JSON(http.StatusInternalServerError, err.Error())
		}

		c.JSON(http.StatusOK, characters)
	}
}
