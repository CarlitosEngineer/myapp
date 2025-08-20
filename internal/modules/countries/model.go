package countries

import (
	"time"

	"gorm.io/gorm"
)

type Country struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Name      string         `json:"name" gorm:"size:100;not null;uniqueIndex"`
	Code      string         `json:"code" gorm:"size:2;not null;uniqueIndex"` // ISO2 (ej: CO, US)
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}

func (Country) TableName() string { return "countries" }
