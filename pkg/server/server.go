package server

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/datahearth/portainer-templates/pkg/db"
	"github.com/datahearth/portainer-templates/pkg/server/handlers"
	"github.com/datahearth/portainer-templates/pkg/server/middlewares"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Server interface {
	Start() error
	RegisterHandlers()
}

type server struct {
	logger   logrus.FieldLogger
	router   *mux.Router
	handlers handlers.Handler
	db       db.Database
	address  string
}

func NewServer(logger logrus.FieldLogger, database db.Database, address, port string) (Server, error) {
	if logger == nil {
		return nil, errors.New("logger is mandatory")
	}
	if database == nil {
		return nil, errors.New("database is mandatory")
	}
	if port == "" {
		return nil, errors.New("port is mandatory")
	}

	hs, err := handlers.NewHandlers(database, logger)
	if err != nil {
		return nil, err
	}

	return &server{
		handlers: hs,
		router:   mux.NewRouter(),
		logger:   logger.WithField("pkg", "server"),
		db:       database,
		address:  fmt.Sprintf("%s:%s", address, port),
	}, nil
}

func (srv *server) Start() error {
	srv.logger.Infof("Server listening on %s", srv.address)
	if err := http.ListenAndServe(srv.address, srv.router); err != nil {
		return err
	}

	return nil
}

func (srv *server) RegisterHandlers() {
	srv.router.Use(middlewares.HTTPLogger)
	srv.router.HandleFunc("/templates", srv.handlers.GetAllTemplates).Methods("GET")
	srv.router.HandleFunc("/templates/{type}/{id}", srv.handlers.GetTemplateById).Methods("GET")
	srv.router.HandleFunc("/templates/load", srv.handlers.LoadFromFile).Methods("POST")
	srv.router.HandleFunc("/templates/insert", srv.handlers.InsertTemplates).Methods("POST")
}
