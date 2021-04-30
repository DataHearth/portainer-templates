package db

import (
	"github.com/datahearth/portainer-templates/pkg/db/tables"
	"github.com/datahearth/portainer-templates/pkg/db/templates"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm/clause"
)

func (db *database) getComposeTemplates() ([]tables.ComposeTable, error) {
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

func (db *database) getComposeById(id int) (*tables.ComposeTable, error) {
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

func (db *database) AddComposeTemplates([]templates.Compose) {

}
