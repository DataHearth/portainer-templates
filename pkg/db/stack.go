package db

import (
	"github.com/datahearth/portainer-templates/pkg/db/tables"
	"github.com/datahearth/portainer-templates/pkg/db/templates"
	"github.com/datahearth/portainer-templates/pkg/utils"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm/clause"
)

func (db *database) getStackTemplates() ([]tables.StackTable, error) {
	stack := []tables.StackTable{}
	res := db.Preload(clause.Associations).Preload("Envs.Selects").Find(&stack)
	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		db.logger.WithField("component", "getStackTemplates").Warnln("no stack templates were retrieved")
	}

	return stack, nil
}

func (db *database) getStackById(id int) (*tables.StackTable, error) {
	var stack *tables.StackTable
	res := db.Preload(clause.Associations).Preload("Envs.Selects").Find(stack, "id = ?", id)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		db.logger.WithFields(logrus.Fields{
			"component": "getStackById",
			"id":        id,
		}).Warnln("no stack template was found")
		return nil, nil
	}

	return stack, nil
}

func (db *database) AddStackTemplates(stacks []templates.Stack) error {
	for _, s := range stacks {
		if err := db.AddStackTemplate(s); err != nil {
			return err
		}
	}

	return nil
}

func (db *database) AddStackTemplate(stack templates.Stack) error {
	sqlStacks := utils.JSONStackToSQL(stack)

	if err := db.Where("title = ?", sqlStacks.Title).FirstOrCreate(&sqlStacks).Error; err != nil {
		db.logger.WithError(err).WithFields(logrus.Fields{
			"component":   "AddStackTemplates",
			"stack-title": sqlStacks.Title,
		}).Errorln("failed to insert stack in database")

		return err
	}

	return nil
}
