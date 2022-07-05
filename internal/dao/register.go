package dao

import "xiaolong_blog/internal/model"

func (d *Dao) QueryEmail(email string) (*model.Register, error) {
	register := model.Register{Email: email}
	return register.Query(d.engine)
}

func (d *Dao) AddVerifyEmail(email string, verifyCode string, expireTime int64) error {
	register := model.Register{Email: email, VerifyCode: verifyCode, ExpireTime: expireTime}
	return register.Create(d.engine)
}
