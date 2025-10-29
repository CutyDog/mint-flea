package db

import (
	"github.com/CutyDog/mint-flea/services/account/internal/model"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

// getMigrations マイグレーションの定義を取得する
func getMigrations() []*gormigrate.Migration {
	return []*gormigrate.Migration{
		{
			ID: "202510191400_create_accounts",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&model.Account{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable(&model.Account{})
			},
		},
		{
			ID: "202510261400_create_account_wallets",
			Migrate: func(tx *gorm.DB) error {
				return tx.AutoMigrate(&model.Wallet{})
			},
			Rollback: func(tx *gorm.DB) error {
				return tx.Migrator().DropTable(&model.Wallet{})
			},
		},
	}
}
