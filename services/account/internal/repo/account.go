package repo

import (
	"context"
	"errors"

	"github.com/CutyDog/mint-flea/services/account/internal/model"
	"gorm.io/gorm"
)

type AccountRepository interface {
	FindByID(ctx context.Context, id int64) (*model.Account, error)
	FindByUID(ctx context.Context, uid string) (*model.Account, error)
	Create(ctx context.Context, uid string) (*model.Account, error)
}

type accountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) AccountRepository {
	return &accountRepository{db: db}
}

func (r *accountRepository) FindByID(ctx context.Context, id int64) (*model.Account, error) {
	var account model.Account
	if err := r.db.WithContext(ctx).First(&account, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &account, nil
}

func (r *accountRepository) FindByUID(ctx context.Context, uid string) (*model.Account, error) {
	var account model.Account
	if err := r.db.WithContext(ctx).Where("uid = ?", uid).First(&account).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &account, nil
}

func (r *accountRepository) Create(ctx context.Context, uid string) (*model.Account, error) {
	account := &model.Account{
		UID: uid,
	}
	if err := r.db.WithContext(ctx).Create(&account).Error; err != nil {
		return nil, err
	}
	return account, nil
}
