package service

import "xiaolong_blog/internal/model"

type UserEmailVerifyRequest struct {
	Email string `json:"email"`
}

func (svc *Service) QueryEmail(params *UserEmailVerifyRequest) (*model.Register, error) {
	return svc.dao.QueryEmail(params.Email)
}

func (svc *Service) AddVerifyEmail(email string, verifyCode string, expireTime int64) error {
	return svc.dao.AddVerifyEmail(email, verifyCode, expireTime)
}
