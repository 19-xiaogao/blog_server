package errcode

var (
	ErrorGetArticleListFail    = NewError(20010001, "获取文章列表失败")
	ErrorCreateArticleFail     = NewError(20010002, "创建文章失败")
	ErrorUpdateArticleFail     = NewError(20010003, "更新文章失败")
	ErrorDeleteArticleFail     = NewError(20010004, "删除文章失败")
	ErrorCountArticleFail      = NewError(20010005, "统计文章失败")
	ErrorUploadFileFail        = NewError(20030001, "上传文件失败")
	ErrorLoginTokenFileFail    = NewError(20030002, "token生成失败")
	ErrorLoginNotExitUserFail  = NewError(20030003, "没有查询到用户")
	ErrorLoginExitUserFail     = NewError(20030004, "用户已经存在")
	ErrorRegisterSendEmailFail = NewError(20030005, "邮件发送失败")
	ErrorAuthEmailFail         = NewError(20030006, "邮箱验证失败")
)
