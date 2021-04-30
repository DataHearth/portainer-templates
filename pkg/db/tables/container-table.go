package tables

import (
	"time"

	"gorm.io/gorm"
)

type Container struct {
	gorm.Model
	ID                int       `gorm:"primaryKey,autoIncrement"`
	CreatedAt         time.Time `json:"-"`
	UpdatedAt         time.Time `json:"-"`
	Type              int       `json:"type"`
	Title             string    `json:"title" gorm:"unique,index"`
	Description       string    `json:"description"`
	Categories        []string  `json:"categories,omitempty"`
	Platform          string    `json:"platform,omitempty"`
	Logo              string    `json:"logo,omitempty"`
	Image             string    `json:"image"`
	Ports             []string  `json:"ports,omitempty"`
	Volumes           []Volumes `json:"volumes"`
	AdministratorOnly bool      `json:"administrator_only,omitempty"`
	Name              string    `json:"name,omitempty"`
	Registry          string    `json:"registry,omitempty"`
	Command           string    `json:"command,omitempty"`
	Env               []Env     `json:"env,omitempty"`
	Network           string    `json:"network,omitempty"`
	Labels            []Label   `json:"labels,omitempty"`
	Privileged        bool      `json:"privileged,omitempty"`
	Interactive       bool      `json:"interactive,omitempty"`
	RestartPolicy     string    `json:"restart_policy,omitempty"`
	Hostname          string    `json:"hostname,omitempty"`
	Note              string    `json:"note,omitempty"`
}
