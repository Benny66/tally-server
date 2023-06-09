package middleware

import (
	"fmt"
	"github.com/Benny66/tally-server/utils/format"
	"github.com/Benny66/tally-server/utils/language"
	"github.com/Benny66/tally-server/utils/log"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

func Recover() gin.HandlerFunc {
	return func(context *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				log.SystemLog(fmt.Sprintf("%s", err))
				if gin.IsDebugging() {
					debug.PrintStack()
				}
				format.NewResponseJson(context).Error(language.SERVER_PANIC)
			}
		}()
		context.Next()
	}
}
