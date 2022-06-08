package routers

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "xiaolong_blog/docs"
	"xiaolong_blog/internal/routers/admin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	adminApi := r.Group("/api/admin")
	atricleApi := admin.NewArticle()
	{
		adminApi.GET("/article_list", atricleApi.GetList)
		adminApi.GET("/article_get", atricleApi.Get)
		adminApi.DELETE("/article_delete/:id", atricleApi.Delete)
		adminApi.POST("/article_update", atricleApi.Update)
		adminApi.POST("/article_create", atricleApi.Create)
	}

	return r
}
