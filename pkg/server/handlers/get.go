package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/datahearth/portainer-templates/pkg/utils"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func (h *handler) GetAllTemplates(rw http.ResponseWriter, r *http.Request) {
	logger := h.logger.WithField("component", "GetAllTemplates")

	containers, stacks, composes, err := h.db.GetAllTemplates()
	if err != nil {
		logger.WithError(err).Errorln("failed to get all templates")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	b, err := utils.FormatBody(containers, stacks, composes)
	if err != nil {
		logger.WithError(err).Errorln("error while formatting body")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	if _, err = rw.Write(b); err != nil {
		logger.WithError(err).Errorln("failed to write response body")
	}
}

func (h *handler) GetTemplateById(rw http.ResponseWriter, r *http.Request) {
	logger := h.logger.WithField("component", "GetAllTemplates")
	templateType := mux.Vars(r)["type"]
	id := mux.Vars(r)["id"]
	if templateType == "" || id == "" {
		logger.Errorln("id or template type is missing")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	logger = h.logger.WithFields(logrus.Fields{
		"id":            id,
		"template-type": templateType,
	})

	if ok := utils.CheckStringArray(templateType, []string{"compose", "stack", "container"}); !ok {
		logger.WithField("type", templateType).Errorln("invalide template type")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	template, err := h.db.GetTemplateById(templateType, id)
	if err != nil {
		logger.WithError(err).Errorln("failed to retrieve template by id")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	if template == nil {
		logger.WithError(err).Errorln("invalide template id")
		rw.WriteHeader(http.StatusBadRequest)
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
