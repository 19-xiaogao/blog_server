package routers

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
	_ "xiaolong_blog/docs"
	"xiaolong_blog/global"
	"xiaolong_blog/internal/routers/admin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))

	adminApi := r.Group("/api/admin")
	articleApi := admin.NewArticle()
	upload := admin.NewUpload()
	login := admin.NewLogin()
	{
		adminApi.POST("/login", login.Login)
		adminApi.POST("/register", login.Register)
		adminApi.GET("/article/:id", articleApi.Query)
		adminApi.GET("/article_list", articleApi.GetList)
		adminApi.DELETE("/article_delete/:id", articleApi.Delete)
		adminApi.POST("/article_update", articleApi.Update)
		adminApi.POST("/article_create", articleApi.Create)
		adminApi.POST("/upload/file", upload.UploadFile)

	}

	return r
}
