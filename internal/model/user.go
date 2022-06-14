package model

type User struct {
	ID                 int    `gorm:"id" json:"id"`
	UserName           string `gorm:"user_name" json:"user_name"`
	Describe           string `gorm:"describe" json:"describe"`
	Content            string `gorm:"content" json:"content"`
	Email              string `gorm:"email" json:"email"`
	WeChat             string `gorm:"we_chat" json:"we_chat"`
	Github             string `gorm:"github" json:"github"`
	CreateTime         string `gorm:"create_time" json:"create_time"`
	Avatar             string `gorm:"avatar" json:"avatar"`
	BackgroundMusicUrl string `gorm:"background_music_url" json:"background_music_url"`
}

func (u *User) TableName() string {
	return "blog_user"
}
