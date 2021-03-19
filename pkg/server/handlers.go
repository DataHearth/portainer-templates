package server

import (
	"encoding/json"
	"net/http"

	"github.com/datahearth/portainer-templates/pkg/db"
	"github.com/gorilla/mux"
)

func (srv *Server) getAllTemplates(rw http.ResponseWriter, r *http.Request) {
	logger := srv.logger.WithField("component", "getAllTemplates")

	templates, err := srv.db.GetAllTemplates()
	if err != nil {
		logger.WithError(err).Errorln("failed to get all templates")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	b, err := formatBody(templates)
	if err != nil {
		logger.WithError(err).Errorln("error while formatting body")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	_, err = rw.Write(b)
	if err != nil {
		logger.WithError(err).Errorln("failed to write response body")
		return
	}
}

func (srv *Server) getTemplateById(rw http.ResponseWriter, r *http.Request) {
	_ = mux.Vars(r)["id"]

	return
}

func formatBody(templates *db.TemplatesArray) ([]byte, error) {
	t := new(db.Templates)
	t.Version = "2"
	t.Templates = append(t.Templates, templates.Compose, templates.Container, templates.Stack)

	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}

	return b, nil
}
