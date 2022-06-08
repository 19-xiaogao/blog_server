package model

import "github.com/jinzhu/gorm"

type Article struct {
	ID           uint32 `gorm:"primary_key" json:"id"`
	Title        string `gorm:"title"`
	Describe     string `gorm:"describe"`
	CreateTime   string `gorm:"create_time"`
	Content      string `gorm:"content"`
	CommentId    int    `gorm:"comment_id"`
	MusicUrl     string `gorm:"music_url"`
	CoverUrl     string `gorm:"cover_url"`
	LickCount    uint32 `gorm:"lick_count"`
	LookCount    uint32 `gorm:"look_count"`
	CommentCount uint32 `gorm:"comment_count"`
}

func (a Article) TableName() string {
	return "blog_article"
}

func (a Article) Create(db *gorm.DB) error {
	return db.Create(&a).Error
}

func (a Article) Update(db *gorm.DB) error {
	return db.Model(&Article{}).Where("primary_key = ?", a.ID).Updates(map[string]interface{}{
		"title":    a.Title,
		"describe": a.Describe,
		"content":  a.Content,
		"musicUrl": a.MusicUrl,
		"coverUrl": a.CoverUrl,
	}).Error
}

func (a Article) Delete(db *gorm.DB) error {
	return db.Where("id = ?", a.ID).Delete(&a).Error
}

func (a Article) AllCount(db *gorm.DB) (int, error) {
	var count int
	err := db.Model(&a).Count(&count).Error

	if err != nil {
		return 0, err
	}

	return count, nil

}
