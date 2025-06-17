package maps

type MapMetadata struct {
	ID       int     `json:"id"`
	Title    string  `json:"title"`
	ImageURL string  `json:"imageUrl"`
	BBox     string  `json:"bbox"`
	Opacity  float64 `json:"opacity"`
}
