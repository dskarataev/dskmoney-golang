package accounts

import (
	"github.com/dskarataev/migrations"
)

func RunMigrations(db migrations.DB) (appName string, oldVersion int64, newVersion int64, err error) {
	migrations.Register(
		1,
		"create initial tables",
		// up
		func(db migrations.DB) error {
			return nil
		},
		// down
		func(db migrations.DB) error {
			return nil
		},
	)

	return migrations.MigrateApp(db, "accounts")
}
