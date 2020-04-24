package config

type unsplash struct {
	Search       []string
	SelectRandom bool
	Resolution   resolution
	Schedule     int
}

type resolution struct {
	Width  int
	Height int
}

var Unsplash = unsplash{
	Search: []string{
		"technology", "developer", "google", "github",
	},
	SelectRandom: false,
	Resolution: resolution{
		Width:  2560,
		Height: 1440,
	},
	Schedule: 1,
}
