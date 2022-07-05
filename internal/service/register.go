package service

type UserEmailVerifyRequest struct {
	Email string `json:"email"`
}

func (svc *Service) QueryEmail(params *UserEmailVerifyRequest) (interface{}, error) {
	return svc.dao.QueryEmail(params.Email)
}
