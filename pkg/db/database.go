package db

import (
	"errors"
	"sync"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Database interface {
	GetContainerTemplates() ([]Container, error)
	GetContainerById(uint) (*Container, error)
	GetComposeTemplates() ([]Compose, error)
	GetComposeById(uint) (*Compose, error)
	GetStackTemplates() ([]Stack, error)
	GetStackById(uint) (*Stack, error)
	GetAllTemplates() (*TemplatesArray, error)
	Close()
}

type database struct {
	*gorm.DB
	logger logrus.FieldLogger
}

func NewDB(logger logrus.FieldLogger) (Database, error) {
	if logger == nil {
		return nil, errors.New("logger is mandatory")
	}
	logger = logger.WithField("pkg", "database")

	db, err := gorm.Open(sqlite.Open("dev.db"))
	if err != nil {
		return nil, err
	}

	if err := db.AutoMigrate(&Compose{}, &Container{}, &Stack{}); err != nil {
		return nil, err
	}

	return &database{
		DB:     db,
		logger: logger,
	}, nil
}

func (db *database) Close() {
	db.Close()
}

func (db *database) GetAllTemplates() (*TemplatesArray, error) {
	wg := new(sync.WaitGroup)
	templates := new(TemplatesArray)
	logger := db.logger.WithField("component", "GetAllTemplates")

	wg.Add(3)
	logger.Debugln("Start retrieving templates from database...")
	go func(wait *sync.WaitGroup, tmp *TemplatesArray) error {
		composes, err := db.GetComposeTemplates()
		if err != nil {
			return err
		}

		tmp.Compose = composes

		wg.Done()
		logger.Debugln("Compose templates retrieved")

		return nil
	}(wg, templates)
	go func(wait *sync.WaitGroup, tmp *TemplatesArray) error {
		containers, err := db.GetContainerTemplates()
		if err != nil {
			return err
		}

		tmp.Container = containers

		wg.Done()
		logger.Debugln("Container templates retrieved")

		return nil
	}(wg, templates)
	go func(wait *sync.WaitGroup, tmp *TemplatesArray) error {
		stacks, err := db.GetStackTemplates()
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
