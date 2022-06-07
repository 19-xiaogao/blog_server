package dao

import "xiaolong_blog/internal/model"

// 计算所有文章的总数
func (d *Dao) CountArticle() (int, error) {
	article := model.Article{}
	return article.AllCount(d.engine)
}

func (d *Dao) UpdateArticle(id uint32) error {
	article := model.Article{ID: id}
	return article.Update(d.engine)
}

// 删除指定文章...
func (d *Dao) DeleteArticle(id uint32) error {
	article := model.Article{ID: id}

	return article.Delete(d.engine)
}
