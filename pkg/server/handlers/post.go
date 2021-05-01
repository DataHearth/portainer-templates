package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/datahearth/portainer-templates/pkg/db/templates"
	"github.com/datahearth/portainer-templates/pkg/utils"
)

func (h *handler) LoadFromFile(rw http.ResponseWriter, r *http.Request) {
	logger := h.logger.WithField("component", "LoadFromFile")

	templateFile := os.Getenv("TEMPLATE_FILE")
	if templateFile == "" {
		logger.Errorln("Environment variable TEMPLATE_FILE not provided")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	logger = logger.WithField("path", templateFile)

	b, err := ioutil.ReadFile(templateFile)
	if err != nil {
		logger.WithError(err).Errorln("failed to read templates file")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonTemplates := templates.Templates{}
	if err := json.Unmarshal(b, &jsonTemplates); err != nil {
		logger.WithError(err).Errorln("failed unmarshal templates data")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	containers, composes, stacks, err := utils.ReadJSONTemplates(jsonTemplates)
	if err != nil {
		logger.Errorln(err.Error())
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	// todo: use functions in goroutines for process time
	if err := h.db.AddComposeTemplates(composes); err != nil {
		logger.WithError(err).Errorln("failed to insert compose templates into the database")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := h.db.AddContainerTemplates(containers); err != nil {
		logger.WithError(err).Errorln("failed to insert container templates into the database")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := h.db.AddStackTemplates(stacks); err != nil {
		logger.WithError(err).Errorln("failed to insert Stacks templates into the database")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusCreated)
}

func (h *handler) InsertTemplates(rw http.ResponseWriter, r *http.Request) {
	logger := h.logger.WithField("component", "InsertTemplates")

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.WithError(err).Errorln("failed to read templates from body")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	var rBody templates.Templates
	if err := json.Unmarshal(b, &rBody); err != nil {
		logger.WithError(err).Errorln("failed to unmarshal JSON body")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	containers, composes, stacks, err := utils.ReadJSONTemplates(rBody)
	if err != nil {
		logger.Errorln(err.Error())
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	// todo: use functions in goroutines for process time
	if err := h.db.AddComposeTemplates(composes); err != nil {
		logger.WithError(err).Errorln("failed to insert compose templates into the database")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := h.db.AddContainerTemplates(containers); err != nil {
		logger.WithError(err).Errorln("failed to insert container templates into the database")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := h.db.AddStackTemplates(stacks); err != nil {
		logger.WithError(err).Errorln("failed to insert Stacks templates into the database")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	rw.WriteHeader(http.StatusCreated)
}
