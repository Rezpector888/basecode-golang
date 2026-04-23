package models

import "time"

type File struct {
	ID string `gorm:"column:id;primaryKey;type:uuid;default:gen_random_uuid()" json:"id"`

	Name     string `gorm:"column:name" json:"name"`
	Path     string `gorm:"column:path" json:"path"`
	Url      string `gorm:"column:url" json:"url"`
	Mimetype string `gorm:"column:mimetype" json:"mimetype"`

	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated_at"`
}

func (File) TableName() string {
	return "file"
}
