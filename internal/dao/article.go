package dao

import (
	"strconv"
	"time"
	"xiaolong_blog/internal/model"
)

// 计算所有文章的总数
func (d *Dao) CountArticle() (int, error) {
	article := model.Article{}
	return article.AllCount(d.engine)
}

// 更新指定文章
func (d *Dao) UpdateArticle(id int, title string, describe string, content string, musicUrl string, coverUrl string) error {
	article := model.Article{ID: id, Title: title, Describe: describe, CoverUrl: coverUrl, Content: content, MusicUrl: musicUrl}
	return article.Update(d.engine)
}
func (d *Dao) QueryArticle(id int) interface{} {
	article := model.Article{ID: id}
	return article.Query(d.engine)
}

// 删除指定文章...
func (d *Dao) DeleteArticle(id int) error {
	article := model.Article{ID: id}

	return article.Delete(d.engine)
}

// 创建文章
func (d *Dao) CreateArticle(title string, describe string, content string, musicUrl string, coverUrl string) error {
	article := model.Article{
		Title:        title,
		Describe:     describe,
		Content:      content,
		MusicUrl:     musicUrl,
		CoverUrl:     coverUrl,
		CreateTime:   strconv.FormatInt(time.Now().Unix(), 10),
		CommentId:    0,
		LickCount:    0,
		LookCount:    0,
		CommentCount: 0,
	}
	return article.Create(d.engine)
}
