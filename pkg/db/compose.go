package db

import (
	"github.com/datahearth/portainer-templates/pkg/db/tables"
	"github.com/datahearth/portainer-templates/pkg/db/templates"
	"github.com/datahearth/portainer-templates/pkg/utils"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm/clause"
)

func (db *database) GetComposeTemplates() ([]tables.ComposeTable, error) {
	compose := []tables.ComposeTable{}
	res := db.Preload(clause.Associations).Preload("Envs.Selects").Find(&compose)
	if res.Error != nil {
		return nil, res.Error
	}

	if res.RowsAffected == 0 {
		db.logger.WithField("component", "getComposeTemplates").Warnln("no compose templates were retrieved")
	}

	return compose, nil
}

func (db *database) GetComposeById(id int) (*tables.ComposeTable, error) {
	var compose *tables.ComposeTable
	res := db.Preload(clause.Associations).Preload("Envs.Selects").Find(compose, "id = ?", id)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		db.logger.WithFields(logrus.Fields{
			"component": "getComposeById",
			"id":        id,
		}).Warnln("no compose template was found")
		return nil, nil
	}

	return compose, nil
}

func (db *database) AddComposeTemplates(composes []templates.Compose) error {
	for _, c := range composes {
		if err := db.AddComposeTemplate(c); err != nil {
			return err
		}
	}

	return nil
}

func (db *database) AddComposeTemplate(compose templates.Compose) error {
	sqlCompose := utils.JSONComposeToSQL(compose)

	if err := db.Where("title = ?", sqlCompose.Title).FirstOrCreate(&sqlCompose).Error; err != nil {
		db.logger.WithError(err).WithFields(logrus.Fields{
			"component":     "AddComposeTemplates",
			"compose-title": sqlCompose.Title,
		}).Errorln("failed to insert compose in database")
		
		return err
	}

	return nil
}
