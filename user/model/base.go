package model

import (
	"gorm.io/gorm"
	"time"
)

type Base struct {
	CreatedAt int64 `gorm:"column:created_at" db:"created_at" json:"created_at" form:"created_at"`
	UpdatedAt int64 `gorm:"column:updated_at" db:"updated_at" json:"updated_at" form:"updated_at"`
	DeletedAt int64 `gorm:"column:deleted_at" db:"deleted_at" json:"deleted_at" form:"deleted_at"`
}

func (base *Base) BeforeCreate(tx *gorm.DB) (err error) {
	base.CreatedAt = time.Now().Unix()
	base.UpdatedAt = time.Now().Unix()
	return
}

func (base *Base) BeforeUpdate(tx *gorm.DB) (err error) {
	base.UpdatedAt = time.Now().Unix()
	return
}

func (base *Base) BeforeDelete(tx *gorm.DB) (err error) {
	base.DeletedAt = time.Now().Unix()
	return
}
