package model

import "github.com/jinzhu/gorm"

type Register struct {
	ID         int `gorm:"id" json:"id"`
	Email      string
	VerifyCode string `gorm:"verify_code"`
	ExpireTime int64  `gorm:"expire_time"`
}

func (r Register) TableName() string {
	return "blog_register"
}

func (r Register) Create(db *gorm.DB) error {
	db.Where(&Register{Email: r.Email}).Delete(&Article{})
	return db.Create(&r).Error
}

func (r Register) Query(db *gorm.DB) (*Register, error) {
	var register = Register{}
	data := db.Where(&Register{Email: r.Email}).First(&Register{})
	if data.Error != nil {
		return nil, data.Error
	}

	return &register, nil
}
