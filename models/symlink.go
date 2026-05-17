package models

import "gorm.io/gorm"

type Symlink struct {
	gorm.Model
	Source string
	Target string `gorm:"unique"`

	UserID uint
}
