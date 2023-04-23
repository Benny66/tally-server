package routers

import (
	"github.com/Benny66/tally-server/api"
	"github.com/Benny66/tally-server/middleware"
	"github.com/gin-gonic/gin"
)

// v1版本接口
func routerV1(group *gin.RouterGroup) {
	group.POST("auth-login", api.UserApi.AuthLogin)
	group.POST("benediction", api.UserApi.Benediction)

	// 	Route::post('upload-file','index/Api/uploadFile'); // 上传文件
	//    Route::post('get-weather','index/Api/getWeather'); // 获取天气预报

	group.POST("upload-file", api.UserApi.UploadFile)
	group.POST("get-weather", api.UserApi.GetWeather)

	group.Use(middleware.AuthMiddleware())
	{
		group.POST("get-user-info", api.UserApi.GetUserInfo)
		group.POST("set-user-info", api.UserApi.SetUserInfo)

		group.POST("get-user-book", api.BookApi.GetUserBooks)
		group.POST("get-book-info", api.BookApi.GetUserBookInfo)
		group.POST("set-book-edit", api.BookApi.SetBookEdit)

		group.POST("get-category-list", api.CategoryApi.GetCategoryList)
		group.POST("set-category-edit", api.CategoryApi.SetCategoryEdit)

		group.POST("get-tally-main-total", api.TallyApi.GetTallyMainTotal)
		group.POST("get-tally-main-list", api.TallyApi.GetTallyMainList)

		group.POST("set-tally-edit", api.TallyApi.SetTallyMainEdit)
		group.POST("get-tally-info", api.TallyApi.GetTallyMainInfo)

	}

	// group.POST("/user/login", api.UserApi.Login)
	// group.POST("/redis/test", api.RedisApi.Test)
	// group.Use(middleware.JWTMiddleware())
	// {
	// 	group.GET("/user/logout", api.UserApi.Logout)
	// 	group.GET("/user/refresh", api.UserApi.Refresh)
	// 	group.PUT("/user/update", api.UserApi.UpdatePassword)
	// }
}
