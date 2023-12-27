package cmd

import (
	"go-oauth2-server/config"
	"go-oauth2-server/models"
	"go-oauth2-server/util/migrations"
)

// Migrate runs database migrations
func Migrate(configBackend config.ConfigBackendType) error {
	_, db, err := initConfigDB(true, false, configBackend)
	if err != nil {
		return err
	}
	defer func() {
		d, err := db.DB()
		if err == nil {
			d.Close()
		}
	}()

	// Bootstrap migrations
	if err := migrations.Bootstrap(db); err != nil {
		return err
	}

	// Run migrations for the oauth service
	if err := models.MigrateAll(db); err != nil {
		return err
	}

	return nil
}
