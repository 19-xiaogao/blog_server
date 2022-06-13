package admin

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"xiaolong_blog/global"
	"xiaolong_blog/internal/service"
	"xiaolong_blog/pkg/app"
	"xiaolong_blog/pkg/errcode"
	"xiaolong_blog/pkg/upload"
)

type Upload struct {
}

func NewUpload() Upload {
	return Upload{}
}

// UploadFile @Summary 上传图片
//@Produce  json
//@Params file file true "图片"
//@Params type int true "图片类型"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errcode.Error "请求错误"
// @Failure 500 {object} errcode.Error "内部错误"
// @Router /api/admin/upload/file [post]
func (u Upload) UploadFile(c *gin.Context) {
	response := app.NewResponse(c)
	file, fileHeader, err := c.Request.FormFile("file")
	if err != nil {
		response.ToErrorResponse(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}
	fileType, _ := strconv.Atoi(c.PostForm("type"))
	if fileHeader == nil || fileType <= 0 {
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	svc := service.New(c.Request.Context())
	fileInfo, err := svc.UploadFile(upload.FileType(fileType), file, fileHeader)

	if fileInfo != nil {
		global.Logger.Errorf("svc.UploadFile err: %v", err)
		response.ToErrorResponse(errcode.ErrorUploadFileFail.WithDetails(err.Error()))
		return
	}
	response.ToResponse(gin.H{
		"file_access_url": fileInfo.AccessUrl,
	})
}
