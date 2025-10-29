package model

import (
	"errors"

	"gorm.io/gorm"
)

type Wallet struct {
	gorm.Model
	AccountID int64  `gorm:"not null"`
	Address   string `gorm:"not null"` // ウォレットアドレス
	ChainID   int64  `gorm:"not null"` // 1: Ethereum, 2: Polygon, 3: BSC, 4: Avalanche, 5: Fantom, ...
	IsMain    bool   `gorm:"not null"` // メインウォレットかどうか

	Account *Account `gorm:"constraint:OnDelete:CASCADE;foreignKey:AccountID"`
}

// BeforeCreate はGORMのBeforeCreateフック
func (w *Wallet) BeforeCreate(tx *gorm.DB) error {
	return w.validateMainWallet(tx)
}

// BeforeUpdate はGORMのBeforeUpdateフック
func (w *Wallet) BeforeUpdate(tx *gorm.DB) error {
	return w.validateMainWallet(tx)
}

// validateMainWallet はAccountIDにつきIsMain=trueが1つだけかどうかを検証
func (w *Wallet) validateMainWallet(tx *gorm.DB) error {
	if !w.IsMain {
		return nil // IsMain=falseの場合は検証不要
	}

	// 同じAccountIDでIsMain=trueのウォレットが既に存在するかチェック
	var count int64
	err := tx.Model(&Wallet{}).
		Where("account_id = ? AND is_main = ? AND id != ?", w.AccountID, true, w.ID).
		Count(&count).Error

	if err != nil {
		return err
	}

	if count > 0 {
		return errors.New("account can have only one main wallet")
	}

	return nil
}
