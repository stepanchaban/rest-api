package userService

import (
	"time"

	"github.com/oapi-codegen/runtime/types"
)

type User struct {
	ID        types.UUID `gorm:"primaryKey" json:"id"`
	Email     string     `gorm:"uniqueIndex;not null"`
	Password  string     `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `gorm:"index"`
}