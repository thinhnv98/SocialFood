package custom_models

type DestinationFully struct {
	ID          int      `json:"id"`
	Name        string   `json:"name"`
	Country     string   `json:"country"`
	ImageName   string   `json:"imageName"`
	ImageData   string   `json:"imageData"`
	Description string   `json:"description"`
	Latitude    float64  `json:"latitude"`
	Longitude   float64  `json:"longitude"`
	Vote        float64  `json:"vote"`
	Photos      []string `json:"photos"`
}
