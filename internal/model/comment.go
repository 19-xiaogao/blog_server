package model

type Comment struct {
	ID         int    `gorm:"id" json:"id"`
	CreateTime string `gorm:"create_time" json:"create_time"`
	Content    string `gorm:"content" json:"content"`
	Observer   string `gorm:"observer" json:"observer"`
	ArticleId  int    `gorm:"article_id" json:"article_id"`
	GiveLike   uint32 `gorm:"give_like" json:"give_like"`
	LatestId   int    `gorm:"latest_id" json:"latest_id"`
}

func (a Comment) TableName() string {
	return "blog_comment"
}
