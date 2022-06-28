package domain

import (
	"context"
	"errors"
	"github.com/jsiebens/ionscale/internal/util"
	"gorm.io/gorm"
)

type Account struct {
	ID uint64 `gorm:"primary_key;autoIncrement:false"`

	ExternalID   string
	LoginName    string
	AuthMethodID uint64
	AuthMethod   AuthMethod
}

func (r *repository) GetOrCreateAccount(ctx context.Context, authMethodID uint64, externalID, loginName string) (*Account, bool, error) {
	account := &Account{}
	id := util.NextID()

	tx := r.withContext(ctx).
		Where(Account{AuthMethodID: authMethodID, ExternalID: externalID}).
		Attrs(Account{ID: id, LoginName: loginName}).
		FirstOrCreate(account)

	if tx.Error != nil {
		return nil, false, tx.Error
	}

	return account, account.ID == id, nil
}

func (r *repository) GetAccount(ctx context.Context, id uint64) (*Account, error) {
	var account Account
	tx := r.withContext(ctx).Take(&account, "id = ?", id)

	if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if tx.Error != nil {
		return nil, tx.Error
	}

	return &account, nil
}

func (r *repository) DeleteAccountsByAuthMethod(ctx context.Context, authMethodID uint64) (int64, error) {
	tx := r.withContext(ctx).
		Delete(&Account{}, "auth_method_id = ?", authMethodID)

	if tx.Error != nil {
		return 0, tx.Error
	}

	return tx.RowsAffected, nil
}
