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
	templateType := r.URL.Query().Get("type")
	var body []byte

	if templateType != "" {
		switch templateType {
		case "container":
			containers, err := h.db.GetContainerTemplates()
			if err != nil {
				logger.WithError(err).Errorln("failed to get container templates")
				rw.WriteHeader(http.StatusInternalServerError)
				return
			}
			body, err = utils.FormatBody(containers, nil, nil)
			if err != nil {
				logger.WithError(err).Errorln("failed to format response body")
				rw.WriteHeader(http.StatusInternalServerError)
				return
			}
		case "compose":
			composes, err := h.db.GetComposeTemplates()
			if err != nil {
				logger.WithError(err).Errorln("failed to get compose templates")
				rw.WriteHeader(http.StatusInternalServerError)
				return
			}
			body, err = utils.FormatBody(nil, nil, composes)
			if err != nil {
				logger.WithError(err).Errorln("failed to format response body")
				rw.WriteHeader(http.StatusInternalServerError)
				return
			}
		case "stack":
			stacks, err := h.db.GetStackTemplates()
			if err != nil {
				logger.WithError(err).Errorln("failed to get stack templates")
				rw.WriteHeader(http.StatusInternalServerError)
				return
			}
			body, err = utils.FormatBody(nil, stacks, nil)
			if err != nil {
				logger.WithError(err).Errorln("failed to format response body")
				rw.WriteHeader(http.StatusInternalServerError)
				return
			}
		default:
			logger.Errorln("invalid templates type")
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
	} else {
		containers, stacks, composes, err := h.db.GetAllTemplates()
		if err != nil {
			logger.WithError(err).Errorln("failed to get all templates")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}

		body, err = utils.FormatBody(containers, stacks, composes)
		if err != nil {
			logger.WithError(err).Errorln("error while formatting body")
			rw.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	rw.Header().Add("content-type", "application/json")
	if _, err := rw.Write(body); err != nil {
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

	rw.Header().Add("content-type", "application/json")
	if _, err = rw.Write(b); err != nil {
		logger.WithError(err).Errorln("failed to write response body")
	}
}
