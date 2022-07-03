package admin

import (
	"github.com/gin-gonic/gin"
	"xiaolong_blog/global"
	"xiaolong_blog/pkg/app"
	"xiaolong_blog/pkg/auth"
	"xiaolong_blog/pkg/errcode"
)

type UserLogin struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type Login struct {
}

func NewLogin() Login {
	return Login{}
}

func (u Login) Login(c *gin.Context) {
	response := app.NewResponse(c)
	params := UserLogin{}
	if err := c.ShouldBindJSON(&params); err != nil {
		global.Logger.Errorf("app.BindAndValid errs:%v", err)
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}

	// 假如它用户密码正确...暂时没有写用户表...
	token, err := auth.CreateJwtToken(params.UserName, params.Password)
	if err != nil {
		global.Logger.Errorf("login err: %v", err)
		response.ToErrorResponse(errcode.ErrorLoginTokenFileFail.WithDetails(err.Error()))
	}
	response.ToResponse(gin.H{
		"token": token,
	})
}
