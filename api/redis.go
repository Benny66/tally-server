package api

import (
	"github.com/Benny66/tally-server/utils/format"
	"github.com/Benny66/tally-server/utils/log"
	myRedis "github.com/Benny66/tally-server/utils/redis"

	"github.com/gin-gonic/gin"
)

var RedisApi *redisApi

func init() {
	RedisApi = NewRedisApi()
}

func NewRedisApi() *redisApi {
	return &redisApi{}
}

type redisApi struct {
}

func (api *redisApi) Test(context *gin.Context) {
	key := "redis_test"
	result, err := myRedis.Set(key, 1000, 0)
	if err != nil {
		format.NewResponseJson(context).Error(51001, err.Error())
		return
	}
	if result {
		log.SystemLog("redis set success")
	}
	value, err := myRedis.Get(key)
	if err != nil {
		format.NewResponseJson(context).Error(51001, err.Error())
		return
	}
	format.NewResponseJson(context).Success(value)
}
