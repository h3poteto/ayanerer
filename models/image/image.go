package image

import (
	"../ayaneru"
)

type Image struct {
	ID  int64
	URL string
}

func NewImage(id int64, url string) *Image {
	image := &Image{
		ID:  id,
		URL: url,
	}
	return image
}

func (u *Image) Save() error {
	ayane := ayaneru.NewAyaneru(u.ID, string(ayaneru.Image), u.URL)
	err := ayane.Save()
	return err
}
