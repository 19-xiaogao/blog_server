package model

import (
	"github.com/jinzhu/gorm"
)

type Article struct {
	ID           int    `gorm:"id" json:"id"`
	Title        string `gorm:"title" json:"title"`
	Describe     string `gorm:"describe" json:"describe"`
	CreateTime   string `gorm:"create_time" json:"create_time"`
	Content      string `gorm:"content" json:"content"`
	CommentId    int    `gorm:"comment_id" json:"comment_id"`
	MusicUrl     string `gorm:"music_url" json:"music_url"`
	CoverUrl     string `gorm:"cover_url" json:"cover_url"`
	LickCount    uint32 `gorm:"lick_count" json:"lick_count"`
	LookCount    uint32 `gorm:"look_count" json:"look_count"`
	CommentCount uint32 `gorm:"comment_count" json:"comment_count"`
}

func (a Article) TableName() string {
	return "blog_article"
}

func (a Article) Create(db *gorm.DB) error {
	return db.Create(&a).Error
}

func (a Article) List(db *gorm.DB, pageOffset, pageSize int) ([]*Article, error) {
	var articles []*Article
	if pageOffset == 1 || pageOffset == 0 {
		pageOffset = 0
	}
	data := db.Offset(pageOffset).Limit(pageSize).Find(&articles)
	if data.Error != nil {
		return articles, data.Error
	}
	return articles, nil
}

func (a Article) Update(db *gorm.DB) error {
	return db.Model(&Article{}).Where(&Article{ID: a.ID}).Updates(map[string]interface{}{
		"title":     a.Title,
		"describe":  a.Describe,
		"content":   a.Content,
		"music_url": a.MusicUrl,
		"cover_url": a.CoverUrl,
	}).Error
}

func (a Article) Delete(db *gorm.DB) error {
	return db.Where(&Article{ID: a.ID}).Delete(&Article{}).Error
}
func (a Article) Query(db *gorm.DB) interface{} {
	data := db.Where(&Article{ID: a.ID}).First(&Article{})

	if data.Error != nil {
		return nil
	}
	return data.Value
}

func (a Article) AllCount(db *gorm.DB) (int, error) {
	var count int
	err := db.Model(&a).Count(&count).Error

	if err != nil {
		return 0, err
	}

	return count, nil

}
