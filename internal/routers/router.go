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
	blogApi := admin.NewBlog()
	{
		adminApi.GET("/blog_list", blogApi.GetList)
		adminApi.GET("/blog_get", blogApi.Get)
		adminApi.POST("/blog_delete", blogApi.Delete)
		adminApi.POST("/blog_update", blogApi.Update)
	}

	return r
}
