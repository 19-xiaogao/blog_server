package admin

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
	"xiaolong_blog/global"
	"xiaolong_blog/internal/service"
	"xiaolong_blog/pkg/app"
	"xiaolong_blog/pkg/auth"
	"xiaolong_blog/pkg/errcode"
	"xiaolong_blog/pkg/util"
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

	// 验证邮箱
	svc := service.New(c)
	data, err := svc.VerifyEmail(params.Email, params.AuthCode)
	fmt.Println(data.IsEmpty())
	fmt.Println(data.VerifyCode)
	fmt.Println(params.AuthCode)
	fmt.Println(time.Now().Unix() * 1000)
	fmt.Println(data.ExpireTime)
	if data.IsEmpty() || data.VerifyCode != params.AuthCode || time.Now().Unix()*1000 < data.ExpireTime {
		response.ToErrorResponse(errcode.ErrorAuthEmailFail)
		return
	}

	params.Password = auth.SHA256Secret(params.Password)
	err = svc.RegisterUser(&params)
	if err != nil {
		response.ToErrorResponse(errcode.ErrorLoginNotExitUserFail)
		return
	}

	response.ToResponse(gin.H{
		"msg": "注册成功",
	})
}

func (u Login) SendVerifyEmail(c *gin.Context) {
	response := app.NewResponse(c)
	params := service.UserEmailVerifyRequest{}
	if err := c.ShouldBindJSON(&params); err != nil {
		global.Logger.Errorf("app.BindAndValid errs:%v", err)
		response.ToErrorResponse(errcode.InvalidParams)
		return
	}
	svc := service.New(c)
	data, _ := svc.QueryEmail(&params)

	if data != nil {
		if !data.IsEmpty() && time.Now().Unix()*1000 < data.ExpireTime {
			response.ToResponse(gin.H{
				"msg": "邮箱已经发送,请查收，过期时间10分钟。",
			})
			return
		}
	}
	roundNumber := util.CreateSixNumber()

	err := auth.SendEmail(params.Email, roundNumber)
	if err != nil {
		global.Logger.Errorf("app.sendVerifyEmail errs%v", err)
		response.ToErrorResponse(errcode.ErrorRegisterSendEmailFail.WithDetails(err.Error()))
		return
	}
	expireTime := time.Now().Add(time.Minute*5).Unix() * 1000
	err = svc.AddVerifyEmail(params.Email, roundNumber, expireTime)

	if err != nil {
		global.Logger.Errorf("app.sendVerifyEmail errs%v", err)
		response.ToErrorResponse(errcode.ErrorRegisterSendEmailFail.WithDetails(err.Error()))
		return
	}

	response.ToResponse(gin.H{
		"msg": "邮件发送成功",
	})

}
