package handlers

import (
	"errors"
	"net/http"

	"github.com/datahearth/portainer-templates/pkg/db"
	"github.com/sirupsen/logrus"
)

type Handler interface {
	GetAllTemplates(rw http.ResponseWriter, r *http.Request)
	GetTemplateById(rw http.ResponseWriter, r *http.Request)
	LoadFromFile(rw http.ResponseWriter, r *http.Request)
	InsertTemplates(rw http.ResponseWriter, r *http.Request)
}

type handler struct {
	logger logrus.FieldLogger
	db     db.Database
}

func NewHandlers(db db.Database, logger logrus.FieldLogger) (Handler, error) {
	if logger == nil {
		return nil, errors.New("logger is mandatory")
	}
	if db == nil {
		return nil, errors.New("database is mandatory")
	}

	return &handler{
		logger: logger.WithField("pkg", "handler"),
		db:     db,
	}, nil
}
