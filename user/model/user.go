package model

const (
	STATUS_USER = iota + 1
	STATUS_FORBIT
)

type FrontUser struct {
	Base
	Name     string `gorm:"column:name" db:"name" json:"name" form:"name"` //  用户名称
	Status   int64  `gorm:"column:status" db:"status" json:"status" form:"status"`
	Email    string `gorm:"column:email" db:"email" json:"email" form:"email"`             //  用户邮箱
	Password string `gorm:"column:password" db:"password" json:"password" form:"password"` //  密码
	Salt     string `gorm:"column:salt" db:"salt" json:"salt" form:"salt"`
}

func (FrontUser) TableName() string {
	return "front_user"
}
