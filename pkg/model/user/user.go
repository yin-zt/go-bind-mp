package user

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(50);not null;unique;comment:'用户名'" json:"username"` // 用户名
	Password string `gorm:"size:255;not null;comment:'用户密码'" json:"password"`               // 用户密码
	Nickname string `gorm:"type:varchar(50);comment:'中文名'" json:"nickname"`                 // 昵称
}
