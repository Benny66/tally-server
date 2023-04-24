// config/config.go

package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var Config *config

type config struct {
	Mode      string
	IPAddress string
	Port      string

	DBHost     string
	DBPort     string
	DBUsername string
	DBPassword string
	DBName     string

	//redis
	RedisUsed     string
	RedisAddr     string
	RedisPassword string

	LogExpire   int //日志保存时间(单位：天)
	AppName     string
	AppVersion  string
	TokenSecret string
	Language    string

	WxAPPID  string
	WxSecret string

	AppUrl        string
	AppUploadType string
	AppLocalPath  string

	QiNiuHost         string
	QiNiuRegionHost   string
	QiNiuAccessId     string
	QiNiuAccessSecret string
	QiNiuBucket       string
	QiNiuStyleDetail  string
}

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err.Error())
	}

	Config = &config{
		Mode:      getEnv("MODE", "debug"),
		IPAddress: getEnv("IP_ADDRESS", "0.0.0.0"),
		Port:      getEnv("PORT", "8080"),

		DBHost:     getEnv("DB_HOST", ""),
		DBPort:     getEnv("DB_PORT", ""),
		DBUsername: getEnv("DB_USERNAME", ""),
		DBPassword: getEnv("DB_PASSWORD", ""),
		DBName:     getEnv("DB_NAME", ""),

		RedisUsed:     getEnv("DB_REDIS_USED", "false"),
		RedisAddr:     getEnv("DB_REDIS_ADDRESS", ""),
		RedisPassword: getEnv("DB_REDIS_PASSWORD", ""),
		AppName:       getEnv("APP_NAME", "gin-server"),
		AppVersion:    getEnv("APP_VERSION", "v1.0.0"),
		TokenSecret:   getEnv("TOKEN_SECRET", "akhdfijwfwsefsdf"),
		Language:      getEnv("LANGUAGE", "zh-cn"),

		WxAPPID:       getEnv("WX_APPID", "xxxxx"),
		WxSecret:      getEnv("WX_SECRET", "xxxx"),
		AppUrl:        getEnv("APP_URL", "http://127.0.0.1:8080"),
		AppUploadType: getEnv("APP_UPLOAD_TYPE", "qiniu"),

		AppLocalPath: getEnv("APP_LOCAL_PATH", "./public"),

		QiNiuHost:         getEnv("QINIU_HOST", ""),
		QiNiuRegionHost:   getEnv("QINIU_REGION_HOST", ""),
		QiNiuAccessId:     getEnv("QINIU_ACCESS_ID", ""),
		QiNiuAccessSecret: getEnv("QINIU_ACCESS_SECRET", ""),
		QiNiuBucket:       getEnv("QINIU_BUCKET", ""),
		QiNiuStyleDetail:  getEnv("QINIU_STYLE_DETAIL", "imageView2/2/w/750"),
	}
	Config.LogExpire, _ = strconv.Atoi(getEnv("LOG_EXPIRE", "30")) //日志保存时间(单位：天")

}

func getEnv(key string, def string) string {
	value := os.Getenv(key)
	if value == "" {
		return def
	}
	return value
}
