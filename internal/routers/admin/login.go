package admin

import (
	"github.com/gin-gonic/gin"
	"xiaolong_blog/global"
	"xiaolong_blog/internal/service"
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
	params := service.UserLoginRequest{}
	if err := c.ShouldBindJSON(&params); err != nil {
		global.Logger.Errorf("app.BindAndValid errs:%v", err)
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	svc := service.New(c)
	params.Password = auth.SHA256Secret(params.Password)
	err := svc.QueryUserExit(&params)
	if err == nil {
		response.ToErrorResponse(errcode.ErrorLoginNotExitUserFail)
		return
	}
	token, err := auth.CreateJwtToken(params.UserName, params.Password)
	if err != nil {
		global.Logger.Errorf("login err: %v", err)
		response.ToErrorResponse(errcode.ErrorLoginTokenFileFail.WithDetails(err.Error()))
		return
	}
	response.ToResponse(gin.H{
		"token": token,
	})
}

func (u Login) Register(c *gin.Context) {
	response := app.NewResponse(c)
	params := service.UserLoginRegisterRequest{}
	if err := c.ShouldBindJSON(&params); err != nil {
		global.Logger.Errorf("app.BindAndValid errs:%v", err)
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	svc := service.New(c)
	params.Password = auth.SHA256Secret(params.Password)
	err := svc.RegisterUser(&params)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorLoginNotExitUserFail)
		return
	}
	response.ToResponse(gin.H{
		"msg": "注册成功",
	})
}
