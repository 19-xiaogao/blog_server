package errcode

var (
	ErrorGetArticleListFail = NewError(20010001, "获取文章列表失败")
	ErrorCreateArticleFail  = NewError(20010002, "创建文章失败")
	ErrorUpdateArticleFail  = NewError(20010003, "更新文章失败")
	ErrorDeleteArticleFail  = NewError(20010004, "删除文章失败")
	ErrorCountArticleFail   = NewError(20010005, "统计文章失败")
	ErrorUploadFileFail     = NewError(20030001, "上传文件失败")
)
