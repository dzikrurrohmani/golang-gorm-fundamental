package model

import (
	"time"

	"gorm.io/gorm"
)

// type DeletedAt gorm.DeletedAt

type BaseModel struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	// DeletedAt sql.NullTime
	// DeletedAt DeletedAt `gorm:"index"` // kalau ini gabisa, oke
	DeletedAt gorm.DeletedAt `gorm:"index"` // Kalau ini bisa pak
}
