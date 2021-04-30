package db

import (
	"github.com/datahearth/portainer-templates/pkg/db/tables"
	"github.com/datahearth/portainer-templates/pkg/db/templates"
	"github.com/datahearth/portainer-templates/pkg/utils"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm/clause"
)

func (db *database) getContainerTemplates() ([]tables.ContainerTable, error) {
	containers := []tables.ContainerTable{}
	res := db.Preload(clause.Associations).Preload("Envs.Selects").Find(&containers)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		db.logger.WithField("component", "getContainerTemplates").Warnln("no container templates were retrieved")
	}

	return containers, nil
}

func (db *database) getContainerById(id int) (*tables.ContainerTable, error) {
	var container *tables.ContainerTable
	res := db.Preload(clause.Associations).Preload("Envs.Selects").Find(container, "id = ?", id)
	if res.Error != nil {
		return nil, res.Error
	}
	if res.RowsAffected == 0 {
		db.logger.WithFields(logrus.Fields{
			"component": "getContainerById",
			"id":        id,
		}).Warnln("no container template was found")
		return nil, nil
	}

	return container, nil
}

func (db *database) AddContainerTemplates(containers []templates.Container) error {
	for _, c := range containers {
		if err := db.AddContainerTemplate(c); err != nil {
			return err
		}
	}

	return nil
}

func (db *database) AddContainerTemplate(container templates.Container) error {
	sqlContainer := utils.JSONContainerToSQL(container)

	if err := db.Where("title = ?", sqlContainer.Title).FirstOrCreate(&sqlContainer).Error; err != nil {
		db.logger.WithError(err).WithFields(logrus.Fields{
			"component":       "AddContainerTemplates",
			"container-title": sqlContainer.Title,
		}).Errorln("failed to insert container in database")

		return err
	}

	return nil
}
