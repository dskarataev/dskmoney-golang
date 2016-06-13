package dskmoney

import (
	"dskmoney-golang/dskmoney/accounts"
	"fmt"
)

func (this *DSKMoney) runMigrations() error {
	if err := logMigrationResult(accounts.RunMigrations(this.DB)); err != nil {
		return err
	}

	return nil
}

func logMigrationResult(appName string, oldVersion, newVersion int64, err error) error {
	if err != nil {
		fmt.Printf("App %s was not migrated because of error\n", appName)
		return err
	}

	fmt.Printf("App %s was migrated from version %d to %d", appName, oldVersion, newVersion)
	return nil
}
