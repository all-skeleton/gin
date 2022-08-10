package models

import (
	"gorm.io/gorm"
)

type User struct {
	ID        int      `gorm:"id" json:"id,omitempty"`
	Username  string   `json:"username,omitempty"`
	Password  string   `json:"password,omitempty"`
	Name      string   `json:"name"`
	Avatar    string   `json:"avatar"`
	CreatedAt JsonTime `json:"created_at,omitempty"`
	UpdatedAt JsonTime `json:"updated_at,omitempty"`
	DeletedAt gorm.DeletedAt `json:"deleted_at,omitempty"`
}

func (User) TableName() string {
	return "user"
}

func (t *User) AfterFind(tx *gorm.DB) (err error) {
	// todo
	return
}
