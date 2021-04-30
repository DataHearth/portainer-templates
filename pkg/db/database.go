package db

import (
	"errors"
	"strconv"
	"sync"

	"github.com/datahearth/portainer-templates/pkg/db/tables"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Database interface {
	getContainerTemplates() ([]tables.Container, error)
	getContainerById(int) (*tables.Container, error)
	getComposeTemplates() ([]tables.Compose, error)
	getComposeById(int) (*tables.Compose, error)
	getStackTemplates() ([]tables.Stack, error)
	getStackById(int) (*tables.Stack, error)
	GetAllTemplates() (*tables.TemplatesArray, error)
	GetTemplateById(string, string) (interface{}, error)
}

type database struct {
	*gorm.DB
	logger logrus.FieldLogger
}

func NewDB(logger logrus.FieldLogger) (Database, error) {
	if logger == nil {
		return nil, errors.New("logger is mandatory")
	}

	db, err := gorm.Open(sqlite.Open("dev.db"), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, err
	}

	return &database{
		DB:     db,
		logger: logger.WithField("pkg", "database"),
	}, nil
}

func (db *database) GetAllTemplates() (*tables.TemplatesArray, error) {
	wg := new(sync.WaitGroup)
	templates := new(tables.TemplatesArray)
	logger := db.logger.WithField("component", "GetAllTemplates")

	logger.Debugln("Start retrieving templates from database...")
	wg.Add(3)
	go func(wait *sync.WaitGroup, tmp *tables.TemplatesArray) error {
		composes, err := db.getComposeTemplates()
		if err != nil {
			return err
		}

		tmp.Compose = composes

		wg.Done()
		logger.Debugln("Compose templates retrieved")

		return nil
	}(wg, templates)
	go func(wait *sync.WaitGroup, tmp *tables.TemplatesArray) error {
		containers, err := db.getContainerTemplates()
		if err != nil {
			return err
		}

		tmp.Container = containers

		wg.Done()
		logger.Debugln("Container templates retrieved")

		return nil
	}(wg, templates)
	go func(wait *sync.WaitGroup, tmp *tables.TemplatesArray) error {
		stacks, err := db.getStackTemplates()
		if err != nil {
			return err
		}

		tmp.Stack = stacks

		wg.Done()
		logger.Debugln("Stack templates retrieved")

		return nil
	}(wg, templates)

	wg.Wait()

	return templates, nil
}

func (db *database) GetTemplateById(templateType, id string) (interface{}, error) {
	var template interface{}
	ID, err := strconv.Atoi(id)
	if err != nil {
		return nil, errors.New("failed to convert string to int")
	}

	switch templateType {
	case "container":
		container, err := db.getContainerById(ID)
		if err != nil {
			return nil, err
		}
		template = container
	case "stack":
		stack, err := db.getStackById(ID)
		if err != nil {
			return nil, err
		}
		template = stack
	case "compose":
		compose, err := db.getComposeById(ID)
		if err != nil {
			return nil, err
		}
		template = compose
	}

	return template, nil
}
