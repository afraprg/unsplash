package config

type unsplash struct {
	Search       []string
	SelectRandom bool
	Resolution   resolution
	Interval     int64
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
	Interval: 1,
}
