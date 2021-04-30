package db

import (
	"errors"

	"github.com/datahearth/portainer-templates/pkg/db/tables"
)

func (db *database) getComposeTemplates() ([]tables.Compose, error) {
	var compose []tables.Compose
	res := db.Model(&tables.Compose{}).Find(&compose)
	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, errors.New("no compose templates were retrieved")
	}

	return compose, nil
}

func (db *database) getComposeById(id int) (*tables.Compose, error) {
	var compose *tables.Compose
	res := db.Model(&tables.Compose{}).Where("id = ?", id).Find(compose)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, errors.New("no compose template was found")
	}

	return compose, nil
}
