package model

import "time"

type Base struct {
	CreatedAt int64 `gorm:"column:created_at" db:"created_at" json:"created_at" form:"created_at"`
	UpdatedAt int64 `gorm:"column:updated_at" db:"updated_at" json:"updated_at" form:"updated_at"`
	DeletedAt int64 `gorm:"column:deleted_at" db:"deleted_at" json:"deleted_at" form:"deleted_at"`
}

func (base *Base) BeforeCreate() {
	base.CreatedAt = time.Now().Unix()
	base.UpdatedAt = time.Now().Unix()
}

func (base *Base) BeforeUpdate() {
	base.UpdatedAt = time.Now().Unix()
}

func (base *Base) BeforeDelete() {
	base.DeletedAt = time.Now().Unix()
}
