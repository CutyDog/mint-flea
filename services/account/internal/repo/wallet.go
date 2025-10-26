package repo

import (
	"context"

	"github.com/CutyDog/mint-flea/services/account/internal/model"
	"gorm.io/gorm"
)

type WalletRepository interface {
	ListByAccountID(ctx context.Context, accountID int64) ([]*model.Wallet, error)
	Create(ctx context.Context, params CreateWalletParams) (*model.Wallet, error)
	Delete(ctx context.Context, id int64) (bool, error)
	SetMain(ctx context.Context, id int64) (*model.Wallet, error)
	GetMain(ctx context.Context, accountID int64) (*model.Wallet, error)
}

type walletRepository struct {
	db *gorm.DB
}

type CreateWalletParams struct {
	AccountID int64
	Address   string
	ChainID   int64
	IsMain    bool
}

func NewWalletRepository(db *gorm.DB) WalletRepository {
	return &walletRepository{db: db}
}

func (r *walletRepository) ListByAccountID(ctx context.Context, accountID int64) ([]*model.Wallet, error) {
	var wallets []*model.Wallet
	if err := r.db.WithContext(ctx).Where("account_id = ?", accountID).Find(&wallets).Error; err != nil {
		return nil, err
	}
	return wallets, nil
}

func (r *walletRepository) Create(ctx context.Context, params CreateWalletParams) (*model.Wallet, error) {
	wallet := &model.Wallet{
		AccountID: params.AccountID,
		Address:   params.Address,
		ChainID:   params.ChainID,
		IsMain:    params.IsMain,
	}
	if err := r.db.WithContext(ctx).Create(wallet).Error; err != nil {
		return nil, err
	}
	return wallet, nil
}

func (r *walletRepository) Delete(ctx context.Context, id int64) (bool, error) {
	if err := r.db.WithContext(ctx).Delete(&model.Wallet{}, id).Error; err != nil {
		return false, err
	}
	return true, nil
}

func (r *walletRepository) GetMain(ctx context.Context, accountID int64) (*model.Wallet, error) {
	var wallet model.Wallet
	if err := r.db.WithContext(ctx).Where("account_id = ? AND is_main = true", accountID).First(&wallet).Error; err != nil {
		return nil, err
	}
	return &wallet, nil
}

func (r *walletRepository) SetMain(ctx context.Context, id int64) (*model.Wallet, error) {
	var wallet model.Wallet
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&wallet).Error; err != nil {
		return nil, err
	}

	accountID := wallet.AccountID
	// 同じアカウントIDでメインウォレットに設定していたウォレットのIsMainをfalseにする
	if err := r.db.WithContext(ctx).Model(&model.Wallet{}).Where("account_id = ? AND is_main = true", accountID).Update("is_main", false).Error; err != nil {
		return nil, err
	}

	wallet.IsMain = true
	if err := r.db.WithContext(ctx).Save(&wallet).Error; err != nil {
		return nil, err
	}
	return &wallet, nil
}
