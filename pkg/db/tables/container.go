package tables

import (
	"time"

	"gorm.io/gorm"
)

// ** Main table ** //
type ContainerTable struct {
	gorm.Model
	ID                int `gorm:"primaryKey,autoIncrement,not null"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	Type              int
	Title             string `gorm:"unique,index,not null"`
	Description       string
	Categories        []ContainerCategory
	Platform          string
	Logo              string
	Image             string `gorm:"not null"`
	Ports             []ContainerPort
	Volumes           []ContainerVolume
	AdministratorOnly bool
	Name              string
	Registry          string
	Command           string
	Env               []ContainerEnv
	Network           string
	Labels            []ContainerLabel
	Privileged        bool
	Interactive       bool
	RestartPolicy     string
	Hostname          string
	Note              string
}

// ** GORM table naming ** //
func (ContainerTable) TableName() string {
	return "container"
}

// ** Sub tables ** //
type ContainerCategory struct {
	gorm.Model
	ID               int `gorm:"primaryKey,autoIncrement,not null"`
	Name             string
	ContainerTableID int
}

type ContainerPort struct {
	gorm.Model
	ID               int `gorm:"primaryKey,autoIncrement,not null"`
	Port             string
	ContainerTableID int
}

type ContainerVolume struct {
	gorm.Model
	ID               int `gorm:"primaryKey,autoIncrement,not null"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	Container        string `gorm:"not null"`
	Bind             string
	ReadOnly         bool
	ContainerTableID int
}

type ContainerEnv struct {
	gorm.Model
	ID               int `gorm:"primaryKey,autoIncrement,not null"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	Name             string `gorm:"not null"`
	Label            string `gorm:"not null"`
	Description      string
	Default          string
	Preset           string
	Select           []ContainerSelect
	ContainerTableID int
}

type ContainerSelect struct {
	gorm.Model
	ID             int `gorm:"primaryKey,autoIncrement,not null"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Text           string `gorm:"not null"`
	Value          string `gorm:"not null"`
	Default        bool
	ContainerEnvID int
}

type ContainerLabel struct {
	gorm.Model
	ID               int `gorm:"primaryKey,autoIncrement,not null"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	Name             string `gorm:"not null"`
	Value            string `gorm:"not null"`
	ContainerTableID int
}

// ** JSON ** //
type Container struct {
	ID                int       `json:"id"`
	Type              int       `json:"type"`
	Title             string    `json:"title"`
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
