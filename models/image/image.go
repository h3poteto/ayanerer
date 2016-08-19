package image

import (
	"../ayaneru"
)

// Image is wrap struct for ayaneru
type Image struct {
	ID  int64
	URL string
}

// NewImage prepare struct for Image
func NewImage(id int64, url string) *Image {
	image := &Image{
		ID:  id,
		URL: url,
	}
	return image
}

// Save struct in database
func (u *Image) Save() error {
	ayane := ayaneru.NewAyaneru(u.ID, string(ayaneru.Image), u.URL)
	err := ayane.Save()
	return err
}
