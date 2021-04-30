package server

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"path"

	"github.com/datahearth/portainer-templates/pkg/db/templates"
	"github.com/datahearth/portainer-templates/pkg/utils"
	"github.com/gorilla/mux"
	"github.com/mitchellh/mapstructure"
	"github.com/sirupsen/logrus"
)

func (srv *server) getAllTemplates(rw http.ResponseWriter, r *http.Request) {
	logger := srv.logger.WithField("component", "getAllTemplates")

	containers, stacks, composes, err := srv.db.GetAllTemplates()
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
		logger.WithField("type", templateType).Errorln("invalide template type")
		rw.WriteHeader(http.StatusBadRequest)
		return
	}

	template, err := srv.db.GetTemplateById(templateType, id)
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

func (srv *server) loadFromFile(rw http.ResponseWriter, r *http.Request) {
	logger := srv.logger.WithField("component", "loadFromFile")

	pwd, err := os.Getwd()
	if err != nil {
		logger.WithError(err).Errorln("failed to get current working directory")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	templatesPath := path.Join(pwd, "templates.json")
	b, err := ioutil.ReadFile(templatesPath)
	if err != nil {
		logger.WithError(err).WithField("path", templatesPath).Errorln("failed to read templates file")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	jsonTemplates := templates.Templates{}
	if err := json.Unmarshal(b, &jsonTemplates); err != nil {
		logger.WithError(err).WithField("path", templatesPath).Errorln("failed unmarshal templates data")
		rw.WriteHeader(http.StatusInternalServerError)
		return
	}

	for i, t := range jsonTemplates.Templates {
		switch int(t.(map[string]interface{})["type"].(float64)) {
		case 1:
			var container templates.Container
			if err := mapstructure.Decode(t, &container); err != nil {
				logger.WithError(err).WithField("path", templatesPath).Errorf("failed to insert container data at index %d", i)
				rw.WriteHeader(http.StatusInternalServerError)
				return
			}
			if err := srv.db.AddContainerTemplate(container); err != nil {
				logger.WithError(err).WithField("path", templatesPath).Errorf("failed to insert container data at index %d", i)
				rw.WriteHeader(http.StatusInternalServerError)
				return
			}

			continue

		case 2:
			var stack templates.Stack
			if err := mapstructure.Decode(t, &stack); err != nil {
				logger.WithError(err).WithField("path", templatesPath).Errorf("failed to decode stack data at index %d", i)
				rw.WriteHeader(http.StatusInternalServerError)
				return
			}

			if err := srv.db.AddStackTemplate(stack); err != nil {
				logger.WithError(err).WithField("path", templatesPath).Errorf("failed to insert stack data at index %d", i)
				rw.WriteHeader(http.StatusInternalServerError)
				return
			}

			continue

		case 3:
			var compose templates.Compose
			if err := mapstructure.Decode(t, &compose); err != nil {
				logger.WithError(err).WithField("path", templatesPath).Errorf("failed to decode compose data at index %d", i)
				rw.WriteHeader(http.StatusInternalServerError)
				return
			}
			if err := srv.db.AddComposeTemplate(compose); err != nil {
				logger.WithError(err).WithField("path", templatesPath).Errorf("failed to insert compose data at index %d", i)
				rw.WriteHeader(http.StatusInternalServerError)
				return
			}

			continue

		default:
			logger.WithField("path", templatesPath).Errorf("invalid data type. Please check your data at index %d\n", i)
			rw.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	rw.WriteHeader(http.StatusCreated)
}
