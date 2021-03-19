package db

import "errors"

func (db *database) GetComposeTemplates() ([]Compose, error) {
	var compose []Compose
	res := db.Model(&Compose{}).Find(&compose)
	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, errors.New("no compose templates were retrieved")
	}

	return compose, nil
}

func (db *database) GetComposeById(id uint) (*Compose, error) {
	var compose *Compose
	res := db.Model(&Compose{}).Where("id = ?", id).Find(compose)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, errors.New("no compose template was found")
	}

	return compose, nil
}
