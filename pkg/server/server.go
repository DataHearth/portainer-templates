package server

import (
	"errors"
	"net/http"

	"github.com/datahearth/portainer-templates/pkg/db"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Server interface {
	Start(string) error
	RegisterHandlers() error
	getTemplateById(http.ResponseWriter, *http.Request)
	getAllTemplates(http.ResponseWriter, *http.Request)
	loadFromFile(http.ResponseWriter, *http.Request)
}

type server struct {
	logger logrus.FieldLogger
	router *mux.Router
	db     db.Database
}

func NewServer(logger logrus.FieldLogger, database db.Database) (Server, error) {
	if logger == nil {
		return nil, errors.New("logger is mandatory")
	}
	if database == nil {
		return nil, errors.New("database is mandatory")
	}

	return &server{
		router: mux.NewRouter(),
		logger: logger.WithField("pkg", "server"),
		db:     database,
	}, nil
}

func (s *server) Start(port string) error {
	s.logger.Infof("Server listening on port %s", port)
	if err := http.ListenAndServe(port, s.router); err != nil {
		return err
	}

	return nil
}

func (s *server) RegisterHandlers() error {
	s.router.HandleFunc("/templates", s.getAllTemplates).Methods("GET")
	s.router.HandleFunc("/templates/{type}/{id}", s.getTemplateById).Methods("GET")
	s.router.HandleFunc("/templates/load", s.loadFromFile).Methods("GET")

	return nil
}
