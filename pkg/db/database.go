package db

import (
	"errors"
	"os"
	"strconv"
	"sync"

	"github.com/datahearth/portainer-templates/pkg/db/tables"
	"github.com/datahearth/portainer-templates/pkg/db/templates"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type Database interface {
	GetContainerTemplates() ([]tables.ContainerTable, error)
	GetContainerById(int) (*tables.ContainerTable, error)
	GetComposeTemplates() ([]tables.ComposeTable, error)
	GetComposeById(int) (*tables.ComposeTable, error)
	GetStackTemplates() ([]tables.StackTable, error)
	GetStackById(int) (*tables.StackTable, error)
	AddStackTemplates([]templates.Stack) error
	AddComposeTemplates([]templates.Compose) error
	AddContainerTemplates([]templates.Container) error
	AddComposeTemplate(templates.Compose) error
	AddStackTemplate(stack templates.Stack) error
	AddContainerTemplate(container templates.Container) error
	GetAllTemplates() ([]tables.ContainerTable, []tables.StackTable, []tables.ComposeTable, error)
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

	dbLocation := "portainer-templates.db"
	if os.Getenv("DB_FILE") != "" {
		dbLocation = os.Getenv("DB_FILE")
	}

	db, err := gorm.Open(sqlite.Open(dbLocation), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(
		&tables.StackTable{}, &tables.StackSelect{}, &tables.StackRepository{}, &tables.StackEnv{}, &tables.StackCategory{},
		&tables.ComposeTable{}, &tables.ComposeSelect{}, &tables.ComposeRepository{}, &tables.ComposeEnv{}, &tables.ComposeCategory{},
		&tables.ContainerTable{}, &tables.ContainerSelect{}, &tables.ContainerPort{}, &tables.ContainerLabel{}, &tables.ContainerEnv{}, &tables.ContainerCategory{}, &tables.ContainerVolume{},
	); err != nil {
		return nil, err
	}

	return &database{
		DB:     db,
		logger: logger.WithField("pkg", "database"),
	}, nil
}

func (db *database) GetAllTemplates() ([]tables.ContainerTable, []tables.StackTable, []tables.ComposeTable, error) {
	wg := new(sync.WaitGroup)
	err := make(chan error)
	wgDone := make(chan bool)
	composes := []tables.ComposeTable{}
	stacks := []tables.StackTable{}
	containers := []tables.ContainerTable{}
	logger := db.logger.WithField("component", "GetAllTemplates")

	logger.Debugln("Start retrieving templates from database...")
	wg.Add(3)
	go func() {
		var e error
		composes, e = db.GetComposeTemplates()
		if e != nil {
			err <- e
		} else {
			logger.Debugln("Compose templates retrieved")
		}
		wg.Done()
	}()
	go func() {
		var e error
		stacks, e = db.GetStackTemplates()
		if e != nil {
			err <- e
		} else {
			logger.Debugln("Stacks templates retrieved")
		}
		wg.Done()
	}()
	go func() {
		var e error
		containers, e = db.GetContainerTemplates()
		if e != nil {
			err <- e
		} else {
			logger.Debugln("Containers templates retrieved")
		}
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(wgDone)
	}()

	select {
	case <-wgDone:
		break
	case e := <-err:
		close(err)
		return nil, nil, nil, e
	}

	return containers, stacks, composes, nil
}

func (db *database) GetTemplateById(templateType, id string) (interface{}, error) {
	var template interface{}
	ID, err := strconv.Atoi(id)
	if err != nil {
		return nil, errors.New("failed to convert string to int")
	}

	switch templateType {
	case "container":
		container, err := db.GetContainerById(ID)
		if err != nil {
			return nil, err
		}
		template = container
	case "stack":
		stack, err := db.GetStackById(ID)
		if err != nil {
			return nil, err
		}
		template = stack
	case "compose":
		compose, err := db.GetComposeById(ID)
		if err != nil {
			return nil, err
		}
		template = compose
	}

	return template, nil
}
