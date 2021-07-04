package custom_models

type ReviewAndVote struct {
	UserID       int    `json:"userID"`
	RestaurantID int    `json:"restaurantID"`
	Rank         int    `json:"rank"`
	Text         string `json:"text"`
}

type ReviewUser struct {
	ID           int    `json:"id"`
	Username     string `json:"username"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	ProfileImage string `json:"profileImage"`
}

type Review struct {
	User   ReviewUser `json:"user"`
	Rating int        `json:"rating"`
	Text   string     `json:"text"`
}

type DishWithResID struct {
	RestaurantID int    `json:"restaurantID"`
	Name         string `json:"name"`
	Price        string `json:"price"`
	Photo        string `json:"photo"`
	NumPhotos    int    `json:"numPhotos"`
}

type Dish struct {
	Name      string `json:"name"`
	Price     string `json:"price"`
	Photo     string `json:"photo"`
	NumPhotos int    `json:"numPhotos"`
}

type PhotoDetail struct {
	Photo       string `json:"photo"`
	Description string `json:"description"`
	CreatedAt   string `json:"createdAt"`
}

type RestaurantFully struct {
	Rank          int           `json:"rank"`
	City          string        `json:"city"`
	Country       string        `json:"country"`
	Description   string        `json:"description"`
	Photos        []string      `json:"photos"`
	PhotoDetails  []PhotoDetail `json:"photoDetails"`
	PopularDishes []Dish        `json:"popularDishes"`
	Reviews       []Review      `json:"reviews"`
}

type NewRestaurantObj struct {
	Name        string `json:"name"`
	City        string `json:"city"`
	Country     string `json:"country"`
	ImageName   string `json:"imageName"`
	ImageData   string `json:"imageData"`
	Description string `json:"description"`
}
