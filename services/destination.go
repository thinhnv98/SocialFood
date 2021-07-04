package services

import (
	"SocialFood/custom_models"
	"SocialFood/models"
	"context"
	"database/sql"
	"encoding/base64"
	"math"
	"os"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/admin"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/volatiletech/null/v8"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type IDestinationService interface {
	CreateDestination(destination custom_models.DestinationFully) error
	GetDestination() (models.DestinationSlice, error)
	GetDestinationDetailByID(id int) (custom_models.DestinationFully, error)
	Vote(rank models.DestinationRank) (int, error)
}

type DestinationService struct {
	Db *sql.DB
}

func (_self DestinationService) GetDestination() (models.DestinationSlice, error) {
	destinations, err := models.Destinations(
		qm.InnerJoin(models.TableNames.DestinationRank+" dr ON "+models.TableNames.Destination+"."+models.DestinationColumns.ID+" = dr."+models.DestinationRankColumns.DestinationID),
		qm.OrderBy("dr."+models.DestinationRankColumns.TotalOfVote+" DESC"),
		qm.OrderBy("dr."+models.DestinationRankColumns.Vote+" DESC"),
	).All(context.Background(), _self.Db)
	return destinations, err
}

func (_self DestinationService) CreateDestination(destinationFully custom_models.DestinationFully) error {
	//Upload image
	dec, _ := base64.StdEncoding.DecodeString(destinationFully.ImageData)
	f, _ := os.Create("assets/imageName.jpeg")
	f.Write(dec)
	f.Sync()
	f.Close()

	ctx := context.Background()
	cld, _ := cloudinary.NewFromParams("vietname", "254449548734168", "0xa-kDE5e55ecB2xXU9ypVLsDxk")
	_, err := cld.Upload.Upload(ctx, "assets/imageName.jpeg", uploader.UploadParams{PublicID: "newID"})

	resp, err := cld.Admin.Asset(ctx, admin.AssetParams{PublicID: "newID"})
	if err != nil {
		return err
	}

	destinationFully.ImageName = resp.SecureURL
	destinationFully.ImageData = ""

	destination := models.Destination{
		Name:        null.StringFrom(destinationFully.Name),
		Country:     null.StringFrom(destinationFully.Country),
		Description: destinationFully.Description,
		Imagename:   null.StringFrom(destinationFully.ImageName),
		Imagedata:   null.StringFrom(destinationFully.ImageData),
		Latitude:    null.Float64From(destinationFully.Latitude),
		Longitude:   null.Float64From(destinationFully.Longitude),
	}
	err = destination.Insert(context.Background(), _self.Db, boil.Infer())
	if err != nil {
		return err
	}

	detail := models.DestinationDetail{
		DestinationID: null.IntFrom(destination.ID),
		Photourl:      null.StringFrom(destinationFully.ImageName),
	}

	err = detail.Insert(context.Background(), _self.Db, boil.Infer())
	if err != nil {
		return err
	}

	rank := models.DestinationRank{
		DestinationID: destination.ID,
		Vote:          5,
		TotalOfVote:   1,
	}
	err = rank.Insert(context.Background(), _self.Db, boil.Infer())
	if err != nil {
		return err
	}

	return nil
}

func (_self DestinationService) GetDestinationDetailByID(id int) (custom_models.DestinationFully, error) {
	ctx := context.Background()

	destination, err := models.Destinations(
		models.DestinationWhere.ID.EQ(id),
	).One(ctx, _self.Db)
	if err != nil {
		return custom_models.DestinationFully{}, err
	}

	destinationDetail, err := models.DestinationDetails(
		models.DestinationDetailWhere.DestinationID.EQ(null.IntFrom(destination.ID)),
	).All(ctx, _self.Db)
	if err != nil {
		return custom_models.DestinationFully{}, err
	}

	rank, err := models.DestinationRanks(
		models.DestinationRankWhere.DestinationID.EQ(destination.ID),
	).One(ctx, _self.Db)
	if err != nil {
		return custom_models.DestinationFully{}, err
	}

	result := custom_models.DestinationFully{
		ID:          destination.ID,
		Name:        destination.Name.String,
		Country:     destination.Country.String,
		Description: destination.Description,
		Vote:        math.RoundToEven(rank.Vote),
		Photos:      nil,
	}

	for _, detail := range destinationDetail {
		result.Photos = append(result.Photos, detail.Photourl.String)
	}
	if result.Photos == nil {
		result.Photos = append(result.Photos, "")
	}

	return result, nil
}

func (_self DestinationService) Vote(rank models.DestinationRank) (int, error) {
	ctx := context.Background()
	currentRank, err := models.DestinationRanks(
		models.DestinationRankWhere.DestinationID.EQ(rank.DestinationID),
	).One(ctx, _self.Db)
	if err != nil {
		return 0, err
	}

	rank.ID = currentRank.ID
	rank.Vote = (rank.Vote + currentRank.Vote*float64(currentRank.TotalOfVote)) / (float64(currentRank.TotalOfVote) + 1)
	rank.TotalOfVote = currentRank.TotalOfVote + 1

	_, err = rank.Update(ctx, _self.Db, boil.Infer())
	if err != nil {
		return 0, err
	}

	return int(math.RoundToEven(rank.Vote)), nil
}
