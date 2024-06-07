package db

import (
	"Linda/protocol/models"

	"github.com/lukaproject/xerr"
	"github.com/sirupsen/logrus"
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

	defer func() {
		if e := recover(); e != nil {
			logrus.Errorf("initial error failed, err is %v", e)
		}
	}()

	dbi = xerr.Must(gorm.Open(postgres.Open(dsn), &gorm.Config{}))
	xerr.Must0(dbi.AutoMigrate(&models.Bag{}))
	xerr.Must0(dbi.AutoMigrate(&models.Task{}))
}
