package server

import (
	"encoding/json"
	"net/http"

	"github.com/datahearth/portainer-templates/pkg/utils"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func (srv *server) getAllTemplates(rw http.ResponseWriter, r *http.Request) {
	logger := srv.logger.WithField("component", "getAllTemplates")

	templates, err := srv.db.GetAllTemplates()
	if err != nil {
		logger.WithError(err).Errorln("failed to get all templates")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	b, err := utils.FormatBody(templates)
	if err != nil {
		logger.WithError(err).Errorln("error while formatting body")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	if _, err = rw.Write(b); err != nil {
		logger.WithError(err).Errorln("failed to write response body")
	}
}

func (srv *server) getTemplateById(rw http.ResponseWriter, r *http.Request) {
	logger := srv.logger.WithField("component", "getAllTemplates")
	templateType := mux.Vars(r)["type"]
	id := mux.Vars(r)["id"]
	if templateType == "" || id == "" {
		logger.Errorln("id or template type is missing")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	logger = srv.logger.WithFields(logrus.Fields{
		"id":            id,
		"template-type": templateType,
	})

	if ok := utils.CheckStringArray(templateType, []string{"compose", "stack", "container"}); !ok {
		logger.WithField("type", templateType).Errorln("invalide template type in url")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	template, err := srv.db.GetTemplateById(templateType, id)
	if err != nil {
		logger.WithError(err).Errorln("invalide template type in url")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	b, err := json.Marshal(template)
	if err != nil {
		logger.WithError(err).WithFields(logrus.Fields{
			"id":            id,
			"template-type": templateType,
		}).Errorln("failed to marshal template in JSON")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	if _, err = rw.Write(b); err != nil {
		logger.WithError(err).Errorln("failed to write response body")
	}
}
