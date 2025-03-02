package models

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	UserID         int64     `gorm:"userid;autoIncrement"`
	UserName       string    `gorm:"primaryKey;column:username;unique;binding:required"`
	Password       string    `gorm:"password" ;binding:"required"`
	Email          string    `gorm:"email"`
	NickName       string    `gorm:"nickname"`
	ResigterTime   time.Time `gorm:"resigtertime"`
	ChangeTime     time.Time `gorm:"changtime"`
	gorm.DeletedAt `gorm:"index"`
}

// BeforeCreate 钩子函数，在创建记录前执行
func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	if u.NickName == "" { // 如果 NickName 为空，设置默认值
		u.NickName = "user"
	}
	if u.Email == "" { // 如果 Email 为空，设置默认值
		u.Email = ""
	}

	return nil
}
