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
	Envs               []ContainerEnv
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
	Selects           []ContainerSelect
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
