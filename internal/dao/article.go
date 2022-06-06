package dao

import "xiaolong_blog/internal/model"

func (d *Dao) CountArticle() (int, error) {
	article := model.Article{}
	return article.AllCount(d.engine)
}
