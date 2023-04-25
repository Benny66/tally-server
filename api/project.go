package api

import (
	"fmt"

	"github.com/Benny66/tally-server/models"
	"github.com/Benny66/tally-server/utils/format"
	"github.com/gin-gonic/gin"
)

var ProjectApi *projectApi

func init() {
	ProjectApi = NewProjectApi()
}

func NewProjectApi() *projectApi {
	return &projectApi{}
}

type projectApi struct {
}

func (api *projectApi) GetProjects(context *gin.Context) {
	projects, err := models.ProjectDao.OrderBy("sort", "asc").FindAllWhere("")
	if err != nil {
		format.NewResponseJson(context).Error(1, err.Error())
		return
	}
	fmt.Println(projects)
	format.NewResponseJson(context).Success(projects)
}
