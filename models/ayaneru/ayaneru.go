package ayaneru

import (
	"../db"

	"github.com/pkg/errors"
)

type Ayaneru struct {
	ID       int64
	Source   string
	URL      string
	database *db.Database
}

const (
	Image   = "image"
	Youtube = "youtube"
)

func NewAyaneru(id int64, source string, url string) *Ayaneru {
	ayaneru := &Ayaneru{
		ID:     id,
		Source: source,
		URL:    url,
	}
	ayaneru.Initialize()
	return ayaneru
}

func (u *Ayaneru) Initialize() {
	u.database = &db.Database{}
}

func (u *Ayaneru) Save() error {
	table := u.database.Init()
	defer table.Close()

	result, err := table.Exec("insert into ayanerus (source, url, created_at) values (?, ?, now());", u.Source, u.URL)
	if err != nil {
		return errors.Wrap(err, "sql execute error")
	}

	u.ID, _ = result.LastInsertId()
	return nil
}
