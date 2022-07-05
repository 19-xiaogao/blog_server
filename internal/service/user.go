package service

type UserLoginRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}
type UserLoginRegisterRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
	AuthCode string `json:"auth_code"`
	Email    string `json:"email"`
}

func (svc *Service) QueryUserExit(params *UserLoginRequest) error {
	return svc.dao.QueryUser(params.UserName, params.Password)
}

func (svc *Service) RegisterUser(params *UserLoginRegisterRequest) error {
	return svc.dao.RegisterUser(params.UserName, params.Password)
}
