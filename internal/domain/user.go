package domain

import (
	"context"
	"github.com/jsiebens/ionscale/internal/util"
)

type SystemRole string

const (
	SystemRoleNone  SystemRole = ""
	SystemRoleAdmin SystemRole = "admin"
)

func (s SystemRole) IsAdmin() bool {
	return s == SystemRoleAdmin
}

type UserType string

const (
	UserTypeService UserType = "service"
	UserTypePerson  UserType = "person"
)

type UserRole string

const (
	UserRoleNone   UserRole = ""
	UserRoleMember UserRole = "member"
	UserRoleAdmin  UserRole = "admin"
)

func (s UserRole) IsAdmin() bool {
	return s == UserRoleAdmin
}

type User struct {
	ID        uint64 `gorm:"primary_key;autoIncrement:false"`
	Name      string
	UserType  UserType
	TailnetID uint64
	Tailnet   Tailnet
	AccountID *uint64
	Account   *Account
}

type Users []User

func (r *repository) GetOrCreateServiceUser(ctx context.Context, tailnet *Tailnet) (*User, bool, error) {
	user := &User{}
	id := util.NextID()

	query := User{Name: tailnet.Name, TailnetID: tailnet.ID, UserType: UserTypeService}
	attrs := User{ID: id, Name: tailnet.Name, TailnetID: tailnet.ID, UserType: UserTypeService}

	tx := r.withContext(ctx).Where(query).Attrs(attrs).FirstOrCreate(user)

	if tx.Error != nil {
		return nil, false, tx.Error
	}

	return user, user.ID == id, nil
}

func (r *repository) ListUsers(ctx context.Context, tailnetID uint64) (Users, error) {
	var users = []User{}

	tx := r.withContext(ctx).Where("tailnet_id = ?", tailnetID).Find(&users)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return users, nil
}

func (r *repository) DeleteUsersByTailnet(ctx context.Context, tailnetID uint64) error {
	tx := r.withContext(ctx).Where("tailnet_id = ?", tailnetID).Delete(&User{})
	return tx.Error
}

func (r *repository) GetOrCreateUserWithAccount(ctx context.Context, tailnet *Tailnet, account *Account) (*User, bool, error) {
	user := &User{}
	id := util.NextID()

	query := User{AccountID: &account.ID, TailnetID: tailnet.ID}
	attrs := User{ID: id, Name: account.LoginName, TailnetID: tailnet.ID, AccountID: &account.ID, UserType: UserTypePerson}

	tx := r.withContext(ctx).Where(query).Attrs(attrs).FirstOrCreate(user)

	if tx.Error != nil {
		return nil, false, tx.Error
	}

	return user, user.ID == id, nil
}

func (r *repository) DeleteUsersByAuthMethod(ctx context.Context, authMethodID uint64) (int64, error) {
	subQuery := r.withContext(ctx).
		Select("users.id").
		Table("users").
		Joins("JOIN accounts a on a.id = users.account_id").
		Where("a.auth_method_id = ?", authMethodID)

	tx := r.withContext(ctx).
		Delete(&User{}, "id in (?)", subQuery)

	if tx.Error != nil {
		return 0, tx.Error
	}

	return tx.RowsAffected, nil
}
