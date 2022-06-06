package model

import "github.com/jinzhu/gorm"

type Article struct {
	ID           uint32 `gorm:"primary_key" json:"id"`
	Title        string
	Describe     string
	CreateTime   string `json:"create_time"`
	Content      string
	CommentId    string `json:"comment_id"`
	MusicUrl     string `json:"music_url"`
	CoverUrl     string `json:"cover_url"`
	LickCount    uint32 `json:"lick_count"`
	LookCount    uint32 `json:"look_count"`
	CommentCount uint32 `json:"comment_count"`
}

func (a Article) TableName() string {
	return "blog_article"
}

func (a Article) Create(db *gorm.DB) error {
	return db.Create(&a).Error
}

func (a Article) Update(db *gorm.DB) error {
	return db.Model(&Article{}).Where("id = ?", a.ID).Delete(&a).Error
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
