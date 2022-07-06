package middleware

import (
	"xiaolong_blog/pkg/app"
	"xiaolong_blog/pkg/auth"
	"xiaolong_blog/pkg/errcode"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			token string
			ecode = errcode.Success
		)
		token = c.GetHeader("token")
		if len(token) == 0 {
			ecode = errcode.UnauthorizedTokenError
		} else {
			_, err := auth.ValidateToken(token)
			if err != nil {
				switch err.(*jwt.ValidationError).Errors {
				case jwt.ValidationErrorExpired:
					ecode = errcode.UnauthorizedTokenTimeout
				default:
					ecode = errcode.UnauthorizedTokenError
				}
			}
		}
		if ecode != errcode.Success {
			response := app.NewResponse(c)
			response.ToErrorResponse(ecode)
			c.Abort()
			return
		}
		c.Next()
	}
}
