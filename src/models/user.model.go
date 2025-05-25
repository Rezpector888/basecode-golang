package models

import "time"

type User struct {
	ID string `gorm:"column:id;primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`

	Fullname string `gorm:"column:full_name" json:"full_name"`
	Email    string `gorm:"column:email;unique:true;" json:"email"`
	Password string `gorm:"column:password;" json:"-"`

	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (User) TableName() string {
	return "user"
}
