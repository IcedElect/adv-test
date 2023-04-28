package models

import (
	"time"
)

type Photo struct {
	ID        uint       `json:"id" gorm:"primaryKey;autoIncrement:true;"`
	Url       *string    `json:"url" gorm:"not null;"`
	AdvID     uint       `json:"adv_id,omitempty" gorm:"not null;"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at" gorm:"index"`
}
