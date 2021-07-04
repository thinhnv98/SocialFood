package services

import (
	"SocialFood/models"
	"context"
	"database/sql"
	"encoding/base64"
	"os"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/admin"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/volatiletech/null/v8"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type IUserService interface {
	Login(user models.Account) (bool, int, error)
	Register(user models.Account) (int, bool, error)
	IsExisted(user models.Account) (bool, error)
}

type UserService struct {
	Db *sql.DB
}

func (_self UserService) Login(user models.Account) (bool, int, error) {
	ctx := context.Background()
	count, err := models.Accounts(
		models.AccountWhere.Email.EQ(user.Email),
		models.AccountWhere.Password.EQ(user.Password),
	).Count(context.Background(), _self.Db)
	if err != nil {
		return false, 0, err
	}

	if count <= 0 {
		return false, 0, nil
	}

	acc, err := models.Accounts(
		models.AccountWhere.Email.EQ(user.Email),
		models.AccountWhere.Password.EQ(user.Password),
	).One(ctx, _self.Db)
	if err != nil {
		return false, 0, err
	}

	return true, acc.ID, nil
}

func (_self UserService) Register(user models.Account) (int, bool, error) {
	ctx := context.Background()
	//Upload image
	dec, _ := base64.StdEncoding.DecodeString(user.Profileimage.String)
	f, _ := os.Create("assets/imageName.jpeg")
	f.Write(dec)
	f.Sync()
	f.Close()
	cld, _ := cloudinary.NewFromParams("vietname", "254449548734168", "0xa-kDE5e55ecB2xXU9ypVLsDxk")
	_, err := cld.Upload.Upload(ctx, "assets/imageName.jpeg", uploader.UploadParams{PublicID: "newID"})

	resp, err := cld.Admin.Asset(ctx, admin.AssetParams{PublicID: "newID"})
	if err != nil {
		return 0, false, err
	}
	user.Profileimage = null.StringFrom(resp.SecureURL)

	err = user.Insert(ctx, _self.Db, boil.Infer())
	if err != nil {
		return 0, false, err
	}
	return user.ID, true, nil
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
