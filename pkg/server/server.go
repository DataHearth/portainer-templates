package server

import (
	"errors"
	"net/http"

	"github.com/datahearth/portainer-templates/pkg/db"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Server struct {
	logger logrus.FieldLogger
	router *mux.Router
	db     db.Database
}

func NewServer(logger logrus.FieldLogger, database db.Database) (*Server, error) {
	if logger == nil {
		return nil, errors.New("logger is mandatory")
	}
	if database == nil {
		return nil, errors.New("database is mandatory")
	}

	return &Server{
		router: mux.NewRouter(),
		logger: logger.WithField("pkg", "server"),
		db:     database,
	}, nil
}

func (s *Server) Start(port string) error {
	if err := http.ListenAndServe(port, s.router); err != nil {
		return err
	}

	return nil
}

func (s *Server) RegisterHandlers() error {
	s.router.HandleFunc("/templates", s.getAllTemplates).Methods("GET")
	s.router.HandleFunc("/templates/{id}", s.getTemplateById).Methods("GET")

	return nil
}
