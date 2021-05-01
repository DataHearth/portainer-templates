package server

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/datahearth/portainer-templates/pkg/db"
	"github.com/datahearth/portainer-templates/pkg/server/handlers"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type Server interface {
	Start() error
	RegisterHandlers() error
	getTemplateById(http.ResponseWriter, *http.Request)
	getAllTemplates(http.ResponseWriter, *http.Request)
	loadFromFile(http.ResponseWriter, *http.Request)
}

type server struct {
	logger  logrus.FieldLogger
	router  *mux.Router
	db      db.Database
	address string
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

	return &server{
		router:  mux.NewRouter(),
		logger:  logger.WithField("pkg", "server"),
		db:      database,
		address: fmt.Sprintf("%s:%s", address, port),
	}, nil
}

func (srv *server) Start() error {
	srv.logger.Infof("Server listening on %s", srv.address)
	if err := http.ListenAndServe(srv.address, srv.router); err != nil {
		return err
	}

	return nil
}

func (srv *server) RegisterHandlers() error {
	srv.router.Use(handlers.HTTPLogger)
	srv.router.HandleFunc("/templates", srv.getAllTemplates).Methods("GET")
	srv.router.HandleFunc("/templates/{type}/{id}", srv.getTemplateById).Methods("GET")
	srv.router.HandleFunc("/templates/load", srv.loadFromFile).Methods("GET")

	return nil
}
