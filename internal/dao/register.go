package dao

import "xiaolong_blog/internal/model"

func (d *Dao) QueryEmail(email string) (interface{}, error) {
	register := model.Register{Email: email}
	return register.Query(d.engine)
}
