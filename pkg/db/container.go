package db

import (
	"errors"

	"github.com/datahearth/portainer-templates/pkg/db/tables"
)

func (db *database) getContainerTemplates() ([]tables.Container, error) {
	var containers []tables.Container
	res := db.Model(&tables.Container{}).Find(&containers)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, errors.New("no container templates were retrieved")
	}

	return containers, nil
}

func (db *database) getContainerById(id int) (*tables.Container, error) {
	var container *tables.Container
	res := db.Model(&tables.Container{}).Where("id = ?", id).Find(container)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, errors.New("no container template was found")
	}

	return container, nil
}
