package model

import (
	"time"

	"github.com/google/uuid"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type User struct {
	Id        uuid.UUID  `json:"id" gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name      string     `json:"name"`
	Email     string     `gorm:"unique" json:"email"`
	Password  string     `json:"-"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time  `gorm:"autoUpdateTime" json:"updated_at"`
	DeletedAt *time.Time `gorm:"index" json:"deleted_at"`
}
