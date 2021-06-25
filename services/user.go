package services

import (
	"SocialFood/models"
	"context"
	"database/sql"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type IUserService interface {
	Login(user models.Account) (bool, error)
	Register(user models.Account) (bool, error)
	IsExisted(user models.Account) (bool, error)
}

type UserService struct {
	Db *sql.DB
}

func (_self UserService) Login(user models.Account) (bool, error) {
	count, err := models.Accounts(
		models.AccountWhere.Email.EQ(user.Email),
		models.AccountWhere.Password.EQ(user.Password),
	).Count(context.Background(), _self.Db)
	if err != nil {
		return false, err
	}

	if count <= 0 {
		return false, nil
	}

	return true, nil
}

func (_self UserService) Register(user models.Account) (bool, error) {
	err := user.Insert(context.Background(), _self.Db, boil.Infer())
	if err != nil {
		return false, err
	}
	return true, nil
}

func (_self UserService) IsExisted(user models.Account) (bool, error) {
	count, err := models.Accounts(
		models.AccountWhere.Email.EQ(user.Email),
	).Count(context.Background(), _self.Db)
	if err != nil {
		return false, err
	}

	if count > 0 {
		return true, nil
	}

	return false, nil
}
