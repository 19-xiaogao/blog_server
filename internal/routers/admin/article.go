package admin

import (
	"github.com/gin-gonic/gin"
	"xiaolong_blog/global"
	"xiaolong_blog/internal/service"
	"xiaolong_blog/pkg/app"
	"xiaolong_blog/pkg/errcode"
)

type Article struct {
}

func NewArticle() Article {
	return Article{}
}

func (a Article) Get(c *gin.Context) {

}

func (a Article) GetList(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": "200",
	})

}
func (a Article) Delete(c *gin.Context) {

}
func (a Article) Update(c *gin.Context) {

}

// @Summary 创建文章
// @Produce  json
// @Param title body string true "文章名称" minlength(1) maxlength(50)
// @Param describe body string true "文章描述" minlength(10) maxlength(200)
// @Param content body string true "文章内容" minlength(1)
// @Param musicUrl body string true "文章背景音乐url"
// @Param coverUrl body string true "文章封面"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/admin/article_create [post]
func (a Article) Create(c *gin.Context) {

	params := service.CreateArticleRequest{}
	response := app.NewResponse(c)
	if err := c.ShouldBindJSON(&params); err != nil {
		global.Logger.Errorf("app.BindAndValid errs:%v", err)
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	svc := service.New(c.Request.Context())
	err := svc.CreateArticle(&params)

	if err != nil {
		global.Logger.Errorf("svc.Create err :%v", err)
		response.ToErrorResponse(errcode.ErrorCreateArticleFail)
		return
	}
	response.ToResponse(gin.H{})
	//valid, errs := app.BindAndValid(c, &params)
	//
	//if !valid {
	//	global.Logger.Errorf("app.BindAndValid errs:%v", errs)
	//	response.ToErrorResponse(errcode.InvalidParams.WithDetails(errs.Error()))
	//	return
	//}

}
