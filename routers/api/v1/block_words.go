package v1

import (
	"pp-bakcend/pkg/app"
	"pp-bakcend/pkg/enums"
	"pp-bakcend/pkg/logging"
	blockwordsservice "pp-bakcend/service/block_words_service"

	"github.com/gin-gonic/gin"
)

type HandleBlockWordsForm struct {
	Shield    string `json:"shield" binding:"required"`
	Handle    string `json:"handle" binding:"required"`
	Operation bool   `json:"operation" binding:"required"` // true: add or update
}

func HandleBlockWords(c *gin.Context) {
	App := app.Gin{C: c}
	var form HandleBlockWordsForm
	var err error
	if err = c.ShouldBindJSON(&form); err != nil {
		App.Response(200, 400, "")
		logging.Info(err)
		return
	}

	mid, err := c.Cookie("mid")
	if err != nil {
		App.Response(200, enums.INVALID_PARAMS, "")
		return
	}
	if form.Operation {
		err = blockwordsservice.AddOrUpdate(mid, form.Shield, form.Handle)
		if err != nil {
			logging.Error(err)
			App.Response(200, enums.SERVER_ERROR, "")
			return
		}

	} else {
		err = blockwordsservice.Delete(mid, form.Shield)
		if err != nil {
			logging.Error(err)
			App.Response(200, enums.SERVER_ERROR, "")
			return
		}
	}

	App.Response(200, enums.SUCCESS, "")

}

func GetAllBlockWords(c *gin.Context) {
	App := app.Gin{C: c}
	infos, err := blockwordsservice.GetAllBlockWords()
	if err != nil {
		logging.Error(err)
		App.Response(200, enums.SERVER_ERROR, "")
		return
	}

	App.Response(200, enums.SUCCESS, infos)
}

type SetWordVisibilityForm struct {
	Mid     string `json:"mid" binding:"required"`
	Shield  string `json:"shield" binding:"required"`
	Visible bool   `json:"visible" binding:"required"`
}

func SetWordVisibility(c *gin.Context) {
	App := app.Gin{C: c}
	var form SetWordVisibilityForm
	var err error
	if err = c.ShouldBindJSON(&form); err != nil {
		App.Response(200, enums.INVALID_PARAMS, "")
		logging.Info(err)
		return
	}
	App.Response(200, enums.SUCCESS, "")

}

type UserGetBlockWordForm struct {
	use bool
}

func UserGetBlockWord(c *gin.Context) {
	App := app.Gin{C: c}
	mid, _ := c.Cookie("mid")
	var form UserGetBlockWordForm
	var err error
	if err = c.ShouldBindJSON(&form); err != nil {
		App.Response(200, enums.INVALID_PARAMS, "")
		logging.Info(err)
		return
	}
	var infos []blockwordsservice.BlockWordsInfo
	if form.use {
		infos, err = blockwordsservice.GetBlockWordsWithoutRepetition(mid)
		if err != nil {
			logging.Error(err)
			App.Response(200, enums.SERVER_ERROR, "")
			return
		}
	} else {
		infos, err = blockwordsservice.GetUserBlockWords(mid)
		if err != nil {
			logging.Error(err)
			App.Response(200, enums.SERVER_ERROR, "")
			return
		}
	}

	App.Response(200, enums.SUCCESS, infos)

}
