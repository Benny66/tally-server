package api

import (
	"github.com/gin-gonic/gin"
)

var MainApi *mainApi

func init() {
	MainApi = NewMainApi()
}

func NewMainApi() *mainApi {
	return &mainApi{}
}

type mainApi struct {
}

func (api *mainApi) SetTallyMainEdit(context *gin.Context) {

}
func (api *mainApi) GetTallyMainInfo(context *gin.Context) {

}
