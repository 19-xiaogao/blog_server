package dao

import "xiaolong_blog/internal/model"

func (d *Dao) QueryUser(username string, password string) error {
	user := model.User{UserName: username, Password: password}
	return user.QueryUserExit(d.engine)
}
