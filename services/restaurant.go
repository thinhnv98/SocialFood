package services

import (
	"SocialFood/custom_models"
	"SocialFood/models"
	"context"
	"database/sql"
	"encoding/base64"
	"math"
	"os"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/admin"
	"github.com/cloudinary/cloudinary-go/api/uploader"

	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type IRestaurantService interface {
	GetRestaurants() (models.RestaurantSlice, error)
	GetRestaurantDetailByID(id int) (custom_models.RestaurantFully, error)
	CreateRestaurant(restaurantFully custom_models.NewRestaurantObj) error
	CreateDish(dish custom_models.DishWithResID) error
	GetDishesByResID(restaurantID int) (models.RestaurantDishSlice, error)
	Review(input custom_models.ReviewAndVote) (float64, []custom_models.Review, error)
}

type RestaurantService struct {
	Db *sql.DB
}

func (_self RestaurantService) GetRestaurants() (models.RestaurantSlice, error) {
	ctx := context.Background()
	return models.Restaurants(
		qm.InnerJoin(models.TableNames.RestaurantRank+" dr ON "+models.TableNames.Restaurant+"."+models.RestaurantColumns.ID+" = dr."+models.RestaurantRankColumns.RestaurantID),
		qm.OrderBy("dr."+models.RestaurantRankColumns.TotalOfVote+" DESC"),
		qm.OrderBy("dr."+models.RestaurantRankColumns.Vote+" DESC"),
	).All(ctx, _self.Db)
}

func (_self RestaurantService) GetRestaurantDetailByID(id int) (custom_models.RestaurantFully, error) {
	ctx := context.Background()
	restaurant, err := models.Restaurants(
		models.RestaurantWhere.ID.EQ(id),
		qm.Load(models.RestaurantRels.RestaurantDishes),
		qm.Load(models.RestaurantRels.RestaurantPhotos),
		qm.Load(models.RestaurantRels.RestaurantRanks),
	).One(ctx, _self.Db)
	if err != nil {
		return custom_models.RestaurantFully{}, err
	}

	restaurantReviews, err := models.RestaurantReviews(
		models.RestaurantReviewWhere.RestaurantID.EQ(id),
		qm.Load(models.RestaurantReviewRels.Account),
	).All(ctx, _self.Db)

	photos := []string{}
	photoDetails := []custom_models.PhotoDetail{}
	dishes := []custom_models.Dish{}
	reviews := []custom_models.Review{}
	for _, photo := range restaurant.R.RestaurantPhotos {
		photos = append(photos, photo.Photourl.String)
		photoDetails = append(photoDetails, custom_models.PhotoDetail{
			Photo:       photo.Photourl.String,
			Description: photo.Description.String,
			CreatedAt:   photo.CreatedOn.String,
		})
	}

	for _, dish := range restaurant.R.RestaurantDishes {
		dishes = append(dishes, custom_models.Dish{
			Name:      dish.Name.String,
			Price:     dish.Price.String,
			NumPhotos: dish.Numphotos.Int,
			Photo:     dish.Photo.String,
		})
	}

	for _, review := range restaurantReviews {
		reviews = append(reviews, custom_models.Review{
			User: custom_models.ReviewUser{
				ID:           review.AccountID,
				Username:     review.R.Account.Username.String,
				FirstName:    review.R.Account.Firstname.String,
				LastName:     review.R.Account.Lastname.String,
				ProfileImage: review.R.Account.Profileimage.String,
			},
			Rating: review.Rating.Int,
			Text:   review.Text.String,
		})
	}

	result := custom_models.RestaurantFully{
		Rank:          int(math.RoundToEven(restaurant.R.RestaurantRanks[0].Vote)),
		City:          restaurant.City.String,
		Country:       restaurant.Country.String,
		Description:   restaurant.Description.String,
		Photos:        photos,
		PhotoDetails:  photoDetails,
		PopularDishes: dishes,
		Reviews:       reviews,
	}

	return result, nil
}

func (_self RestaurantService) CreateRestaurant(restaurantFully custom_models.NewRestaurantObj) error {
	//Upload image
	dec, _ := base64.StdEncoding.DecodeString(restaurantFully.ImageData)
	f, _ := os.Create("assets/imageName.jpeg")
	f.Write(dec)
	f.Sync()
	f.Close()

	ctx := context.Background()
	cld, _ := cloudinary.NewFromParams("vietname", "254449548734168", "0xa-kDE5e55ecB2xXU9ypVLsDxk")
	_, err := cld.Upload.Upload(ctx, "assets/imageName.jpeg", uploader.UploadParams{PublicID: "newRestaurant"})

	resp, err := cld.Admin.Asset(ctx, admin.AssetParams{PublicID: "newRestaurant"})

	restaurantFully.ImageName = resp.SecureURL
	restaurantFully.ImageData = ""

	newRestaurant := models.Restaurant{
		Name:        null.StringFrom(restaurantFully.Name),
		Imagename:   null.StringFrom(restaurantFully.ImageName),
		City:        null.StringFrom(restaurantFully.City),
		Country:     null.StringFrom(restaurantFully.Country),
		Description: null.StringFrom(restaurantFully.Description),
		//CreatedBy:   null.IntFrom(restaurantFully.),
	}

	err = newRestaurant.Insert(context.Background(), _self.Db, boil.Infer())
	if err != nil {
		return err
	}

	rank := models.RestaurantRank{
		RestaurantID: newRestaurant.ID,
		Vote:         5,
		TotalOfVote:  1,
	}
	err = rank.Insert(context.Background(), _self.Db, boil.Infer())
	if err != nil {
		return err
	}

	return nil
}

func (_self RestaurantService) CreateDish(dish custom_models.DishWithResID) error {
	ctx := context.Background()

	//Upload image
	dec, _ := base64.StdEncoding.DecodeString(dish.Photo)
	f, _ := os.Create("assets/imageName.jpeg")
	f.Write(dec)
	f.Sync()
	f.Close()
	cld, _ := cloudinary.NewFromParams("vietname", "254449548734168", "0xa-kDE5e55ecB2xXU9ypVLsDxk")
	_, err := cld.Upload.Upload(ctx, "assets/imageName.jpeg", uploader.UploadParams{PublicID: "newID"})

	resp, err := cld.Admin.Asset(ctx, admin.AssetParams{PublicID: "newID"})
	if err != nil {
		return err
	}
	dish.Photo = resp.SecureURL

	dishNew := models.RestaurantDish{
		RestaurantID: dish.RestaurantID,
		Name:         null.StringFrom(dish.Name),
		Price:        null.StringFrom(dish.Price),
		Numphotos:    null.IntFrom(dish.NumPhotos),
		Photo:        null.StringFrom(dish.Photo),
	}

	return dishNew.Insert(ctx, _self.Db, boil.Infer())
}

func (_self RestaurantService) GetDishesByResID(restaurantID int) (models.RestaurantDishSlice, error) {
	ctx := context.Background()
	dishes, err := models.RestaurantDishes(
		models.RestaurantDishWhere.RestaurantID.EQ(restaurantID),
	).All(ctx, _self.Db)
	if err != nil {
		return nil, err
	}

	return dishes, nil
}

func (_self RestaurantService) Review(input custom_models.ReviewAndVote) (float64, []custom_models.Review, error) {
	ctx := context.Background()
	review := models.RestaurantReview{
		RestaurantID: input.RestaurantID,
		AccountID:    input.UserID,
		Rating:       null.IntFrom(input.Rank),
		Text:         null.StringFrom(input.Text),
	}
	//insert review
	err := review.Insert(ctx, _self.Db, boil.Infer())
	if err != nil {
		return 0, nil, err
	}

	currentRank, err := models.RestaurantRanks(
		models.RestaurantRankWhere.RestaurantID.EQ(input.RestaurantID),
	).One(ctx, _self.Db)
	if err != nil {
		return 0, nil, err
	}

	rank := models.RestaurantRank{
		ID:           currentRank.ID,
		RestaurantID: input.RestaurantID,
		TotalOfVote:  currentRank.TotalOfVote + 1,
	}
	rank.Vote = (float64(input.Rank) + currentRank.Vote*float64(currentRank.TotalOfVote)) / (float64(currentRank.TotalOfVote) + 1)

	_, err = rank.Update(ctx, _self.Db, boil.Infer())
	if err != nil {
		return 0, nil, err
	}

	//get review
	reviews := []custom_models.Review{}
	rvs, err := models.RestaurantReviews(
		models.RestaurantReviewWhere.RestaurantID.EQ(input.RestaurantID),
		qm.Load(models.RestaurantReviewRels.Account),
	).All(ctx, _self.Db)
	for _, rv := range rvs {
		reviews = append(reviews, custom_models.Review{
			User: custom_models.ReviewUser{
				ID:           rv.R.Account.ID,
				Username:     rv.R.Account.Username.String,
				FirstName:    rv.R.Account.Firstname.String,
				LastName:     rv.R.Account.Lastname.String,
				ProfileImage: rv.R.Account.Profileimage.String,
			},
			Rating: rv.Rating.Int,
			Text:   rv.Text.String,
		})
	}

	return math.RoundToEven(rank.Vote), reviews, nil
}
