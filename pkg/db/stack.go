package db

import "errors"

func (db *database) GetStackTemplates() ([]Stack, error) {
	var stack []Stack
	res := db.Model(&Stack{}).Find(&stack)
	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		return nil, errors.New("no stack templates were retrieved")
	}

	return stack, nil
}

func (db *database) GetStackById(id uint) (*Stack, error) {
	var stack *Stack
	res := db.Model(&Stack{}).Where("id = ?", id).Find(stack)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		return nil, errors.New("no stack template was found")
	}

	return stack, nil
}
