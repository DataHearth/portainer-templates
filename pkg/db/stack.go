package db

import (
	"errors"

	"github.com/datahearth/portainer-templates/pkg/db/tables"
)

func (db *database) getStackTemplates() ([]tables.Stack, error) {
	var stack []tables.Stack
	res := db.Model(&tables.Stack{}).Find(&stack)
	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, errors.New("no stack templates were retrieved")
	}

	return stack, nil
}

func (db *database) getStackById(id int) (*tables.Stack, error) {
	var stack *tables.Stack
	res := db.Model(&tables.Stack{}).Where("id = ?", id).Find(stack)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, errors.New("no stack template was found")
	}

	return stack, nil
}
