package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Bag struct {
	BagName        string `gorm:"primaryKey"`
	BagDisplayName string
	CreateTimeMs   int64
	UpdateTimeMs   int64
}

func (b *Bag) BeforeCreate(tx *gorm.DB) (err error) {
	b.BagName = uuid.NewString()
	createTimeMs := time.Now().UnixMilli()
	b.CreateTimeMs = createTimeMs
	b.UpdateTimeMs = createTimeMs
	return
}

func (b *Bag) BeforeSave(tx *gorm.DB) (err error) {
	b.UpdateTimeMs = time.Now().UnixMilli()
	return
}
