package models

import (
	"time"

	"github.com/google/uuid"
)

type Task struct {
	ID        uuid.UUID `gorm:"primaryKey"`
	Desc      string
	Complete  bool
	CreatedAt time.Time `gorm:"autoUpdateTime:nano"`
	UpdatedAt time.Time `gorm:"autoUpdateTime:nano"`
}
