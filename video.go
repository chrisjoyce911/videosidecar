package videosidecar

import (
	"time"
)

// A videos can have multiple images and sidecar files
type Video struct {
	VideoUUID         string
	VideoPath         string
	VideoName         string
	VideoTitle        string
	VideoTitleChanged bool
	VideoDescription  string
	VideoNotes        string
	VideoArtist       string
	VideoFavorite     bool
	VideoPrivate      bool
	VideoSensitive    bool
	VideoStory        bool
	VideoLat          float64
	VideoLong         float64
	VideoAltitude     int
	VideoFocalLength  int
	VideoIso          int
	VideoAperture     float64
	VideoExposure     string
	VideoViews        uint
	VideoAspectRatio  float64
	Camera            *Camera
	CameraID          uint
	CountryID         string
	CountryChanged    bool
	Location          *Location
	Country           *Country
	LocationID        uint
	LocationChanged   bool
	LocationEstimated bool
	TakenAt           time.Time
	TakenAtLocal      time.Time
	TakenAtChanged    bool
	TimeZone          string
	Labels            []*VideoLabel
	Files             []*File
	Albums            []*Album
}
