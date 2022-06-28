package main

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type DeletedAt gorm.DeletedAt

type BaseModel struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt sql.NullTime
	// DeletedAt DeletedAt `gorm:"index"`
}
