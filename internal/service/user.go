package service

type UserLoginRequest struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func (svc *Service) QueryUserExit(params *UserLoginRequest) error {
	return svc.dao.QueryUser(params.UserName, params.Password)

}
