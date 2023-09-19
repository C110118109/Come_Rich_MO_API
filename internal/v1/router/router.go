package router

import (
	"github.com/gin-gonic/gin"
	"net/http"

	"time"
	"github.com/gin-contrib/cors"
)

func Default() *gin.Engine {

	router := gin.New()

	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	router.StaticFS("/public",http.Dir("./storage")) //啟動靜態資料夾存取

	corsConfig := cors.DefaultConfig() //CORS非簡單請求設定
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowMethods = []string{"GET", "POST", "OPTIONS", "PATCH", "DELETE"}
	corsConfig.AllowHeaders = []string{"Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "accept", "origin", "Cache-Control", "X-Requested-With"}
	corsConfig.AllowCredentials = true
	corsConfig.MaxAge = 24 * time.Hour
	router.Use(cors.New(corsConfig))

	return router
}
