package db

import (
	"Linda/protocol/models"

	"github.com/lukaproject/xerr"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbi *gorm.DB = nil

func Instance() *gorm.DB {
	return dbi
}

func InitialWithDSN(dsn string) {
	if dbi != nil {
		return
	}
	ReInitialWithDSN(dsn)
}

func ReInitialWithDSN(dsn string) {
	dbi = xerr.Must(gorm.Open(postgres.Open(dsn), &gorm.Config{}))
	xerr.Must0(dbi.AutoMigrate(&models.Bag{}))
	xerr.Must0(dbi.AutoMigrate(&models.Task{}))
	xerr.Must0(dbi.AutoMigrate(&models.NodeInfo{}))
}
