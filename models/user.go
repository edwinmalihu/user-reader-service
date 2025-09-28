// user.go
package models

import (
	"gorm.io/gorm"
)

// User adalah model GORM untuk tabel 'users'
type User struct {
	gorm.Model
	Name  string `gorm:"type:varchar(100);not null" json:"name"`
	Email string `gorm:"type:varchar(100);uniqueIndex;not null" json:"email"`
}