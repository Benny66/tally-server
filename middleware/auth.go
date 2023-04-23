package middleware

/*
 * @Descripttion:
 * @version: v1.0.0
 * @Author: shahao
 * @Date: 2021-04-07 09:20:20
 * @LastEditors: shahao
 * @LastEditTime: 2021-07-26 17:20:29
 */

import (
	"github.com/Benny66/tally-server/utils/format"
	"github.com/Benny66/tally-server/utils/language"

	"github.com/gin-gonic/gin"

	"github.com/Benny66/tally-server/models"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		token := context.GetHeader("token")
		if token == "" {
			format.NewResponseJson(context).Error(language.TOKEN_EMPTY)
			return
		}
		userInfo, err := models.NewUserDao().FindOneWhere("token", token)
		if err != nil {
			format.NewResponseJson(context).Error(language.TOKEN_EMPTY)
		}
		context.Set("user", userInfo)
		context.Set("token", token)
		context.Next()
	}
}
