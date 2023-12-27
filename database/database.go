package database

import (
	"fmt"
	"time"

	"go-oauth2-server/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	// Drivers
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
)

var gormconfig gorm.Config

func init() {
	gormconfig = gorm.Config{
		NowFunc: func() time.Time {
			return time.Now().Local()
		}}
}

// NewDatabase returns a gorm.DB struct, gorm.DB.DB() returns a database handle
// see http://golang.org/pkg/database/sql/#DB
func NewDatabase(cnf *config.Config) (db *gorm.DB, err error) {
	gormconfig.Logger = logger.Default
	if cnf.IsDevelopment {
		gormconfig.Logger.LogMode(logger.Info)
	} else {
		gormconfig.Logger.LogMode(logger.Silent)
	}
	switch cnf.Database.Type {
	// Postgres
	case "postgres":
		// Connection args
		// see https://godoc.org/github.com/lib/pq#hdr-Connection_String_Parameters
		dsn := fmt.Sprintf(
			"sslmode=disable host=%s port=%d user=%s password='%s' dbname=%s",
			cnf.Database.Host,
			cnf.Database.Port,
			cnf.Database.User,
			cnf.Database.Password,
			cnf.Database.DatabaseName,
		)
		db, err = gorm.Open(postgres.New(postgres.Config{
			DSN:                  dsn,
			PreferSimpleProtocol: true,
		}), &gormconfig)
		if err != nil {
			return
		}

	case "sqlite":
		db, err = gorm.Open(sqlite.Open(cnf.Database.DatabaseName), &gormconfig)
		if err != nil {
			return
		}
	case "mysql":
		dsn := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
			cnf.Database.User,
			cnf.Database.Password,
			cnf.Database.Host,
			cnf.Database.Port,
			cnf.Database.DatabaseName,
		)
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			return
		}
	}
	if db == nil {
		// Database type not supported
		return nil, fmt.Errorf("database type %s not suppported", cnf.Database.Type)
	}

	sqlDB, err := db.DB()
	if err == nil {
		// Max idle connections
		sqlDB.SetMaxIdleConns(cnf.Database.MaxIdleConns)

		// Max open connections
		sqlDB.SetMaxOpenConns(cnf.Database.MaxOpenConns)
	}
	return db, nil
}
