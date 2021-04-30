package utils

import (
	"encoding/json"

	"github.com/datahearth/portainer-templates/pkg/db/tables"
)

func FormatBody(templates *tables.TemplatesArray) ([]byte, error) {
	t := new(tables.Templates)
	t.Version = "2"
	t.Templates = append(t.Templates, templates.Compose, templates.Container, templates.Stack)

	b, err := json.Marshal(t)
	if err != nil {
		return nil, err
	}

	return b, nil
}
