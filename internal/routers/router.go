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
	atricleApi := admin.NewArticle()
	upload := admin.NewUpload()
	{
		adminApi.GET("/article/:id", atricleApi.Query)
		adminApi.GET("/article_list", atricleApi.GetList)
		adminApi.DELETE("/article_delete/:id", atricleApi.Delete)
		adminApi.POST("/article_update", atricleApi.Update)
		adminApi.POST("/article_create", atricleApi.Create)
		adminApi.POST("/upload/file", upload.UploadFile)

	}

	return r
}
