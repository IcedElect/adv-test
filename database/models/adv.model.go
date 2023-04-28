package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Adv struct {
	ID          uint            `json:"id" gorm:"primaryKey;autoIncrement:true;"`
	Name        string          `json:"name" gorm:"not null;"`
	Description *string         `json:"description" gorm:"type:text;"`
	Price       decimal.Decimal `json:"price" gorm:"type:decimal(10,2);"`
	Photos      []Photo         `json:"photos" gorm:"foreignKey:AdvID;"`
	CreatedAt   time.Time       `json:"created_at"`
	UpdatedAt   time.Time       `json:"updated_at"`
	DeletedAt   *time.Time      `json:"deleted_at" gorm:"index"`
}

type ShortAdv struct {
	ID    uint            `json:"id"`
	Name  string          `json:"name"`
	Price decimal.Decimal `json:"price" gorm:"type:decimal(10,2);"`
}

type FullAdv struct {
	ShortAdv
	Description *string `json:"description"`
	Photos      []Photo `json:"photos"`
}
