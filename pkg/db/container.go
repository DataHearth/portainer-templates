package db

import "errors"

func (db *database) GetContainerTemplates() ([]Container, error) {
	var containers []Container
	res := db.Model(&Container{}).Find(&containers)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, errors.New("no container templates were retrieved")
	}

	return containers, nil
}

func (db *database) GetContainerById(id uint) (*Container, error) {
	var container *Container
	res := db.Model(&Container{}).Where("id = ?", id).Find(container)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, errors.New("no container template was found")
	}

	return container, nil
}
