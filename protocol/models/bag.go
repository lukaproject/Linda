package models

import (
	"time"

	"github.com/google/uuid"
)

type Bag struct {
	BagName        string `gorm:"primaryKey"`
	BagDisplayName string
	CreateTimeMs   int64
	UpdateTimeMs   int64
}

func NewBag(bagDisplayName string) *Bag {
	createTimeMs := time.Now().UnixMilli()
	return &Bag{
		BagName:        uuid.NewString(),
		BagDisplayName: bagDisplayName,
		CreateTimeMs:   createTimeMs,
		UpdateTimeMs:   createTimeMs,
	}
}
