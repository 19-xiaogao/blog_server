package routers

import (
	"net/http"
	_ "xiaolong_blog/docs"
	"xiaolong_blog/global"
	"xiaolong_blog/internal/middleware"
	"xiaolong_blog/internal/routers/admin"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))

	login := admin.NewLogin()
	r.POST("/api/login", login.Login)
	r.POST("/api/register", login.Register)
	r.POST("/api/send_verify_email", login.SendVerifyEmail)

	adminApi := r.Group("/api/admin")
	adminApi.Use(middleware.JWT())
	articleApi := admin.NewArticle()
	upload := admin.NewUpload()
	{

		adminApi.GET("/article/:id", articleApi.Query)
		adminApi.GET("/article_list", articleApi.GetList)
		adminApi.DELETE("/article_delete/:id", articleApi.Delete)
		adminApi.POST("/article_update", articleApi.Update)
		adminApi.POST("/article_create", articleApi.Create)
		adminApi.POST("/upload/file", upload.UploadFile)

	}

	return r
}
