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

//@Summary 查询指定文章
//@Produce  json
//@Params id path int true "文章的ID"
//Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/admin/article/:id [get]
func (a Article) Query(c *gin.Context) {
	params := service.QueryArticleRequest{}
	response := app.NewResponse(c)

	if err := c.ShouldBindUri(&params); err != nil {
		global.Logger.Errorf("app.BindAndValid errs:%v", err)
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	svc := service.New(c.Request.Context())
	data := svc.QueryArticle(&params)

	response.ToResponse(data)
}

//@Summary 删除文章
//@Produce  json
//@Params id path int true "文章的ID"
//Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/admin/article_delete/:id [delete]
func (a Article) Delete(c *gin.Context) {
	params := service.DeleteArticleRequest{}
	response := app.NewResponse(c)
	if err := c.ShouldBindUri(&params); err != nil {
		global.Logger.Errorf("app.BindAndValid errs:%v", err)
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	svc := service.New(c.Request.Context())
	err := svc.DeleteArticle(&params)
	if err != nil {
		global.Logger.Errorf("svc.Create err :%v", err)
		response.ToErrorResponse(errcode.ErrorCreateArticleFail)
		return
	}
	response.ToResponse(gin.H{
		"code": 200,
		"msg":  "删除成功",
	})
}

//@Summary 更新文章
//@Produce  json
//@Params id body int true "文章的ID"
//@Params title body string true "文章的标题"
//@Params describe body string true "文章的描述"
//@Params content body string true "文章的内容"
//@Params musicUrl body int string "文章的背景音乐url"
//@Params coverUrl body int string "文章的封面的url"
//Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/admin/article_update [post]
func (a Article) Update(c *gin.Context) {
	params := service.UpdateArticleRequest{}
	response := app.NewResponse(c)
	if err := c.ShouldBindJSON(&params); err != nil {
		global.Logger.Errorf("app.BindAndValid errs:%v", err)
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	svc := service.New(c.Request.Context())

	err := svc.UpdateArticle(&params)
	if err != nil {
		global.Logger.Errorf("svc.Create err :%v", err)
		response.ToErrorResponse(errcode.ErrorCreateArticleFail)
		return
	}
	response.ToResponse(gin.H{
		"code": 200,
		"msg":  "更新成功",
	})
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
	response.ToResponse(gin.H{
		"code": 200,
		"msg":  "创建成功",
	})
}
