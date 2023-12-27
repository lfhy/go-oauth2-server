package cmd

import (
	"go-oauth2-server/config"

	"github.com/RichardKnop/go-fixtures"
)

// LoadData loads fixtures
func LoadData(paths []string, configBackend config.ConfigBackendType) error {
	cnf, db, err := initConfigDB(true, false, configBackend)
	if err != nil {
		return err
	}
	d, err := db.DB()
	if err != nil {
		return err
	}
	defer d.Close()
	return fixtures.LoadFiles(paths, d, cnf.Database.Type)
}
