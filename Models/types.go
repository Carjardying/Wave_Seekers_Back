package Models

type User struct {
	Email    string
	Password string
}

type Country struct {
	Name string
}

type Spot struct {
	UserID           int
	CountryID        int
	Destination      string
	Location         string
	Lat              float64
	Long             float64
	PeakSeasonStart  string
	PeakSeasonEnd    string
	DifficultyLevel  int
	SurfingCulture   string
	ImageURL         string
}
