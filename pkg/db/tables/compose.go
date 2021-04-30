package tables

import (
	"time"

	"gorm.io/gorm"
)

// ** Main table ** //
type ComposeTable struct {
	gorm.Model
	ID                int `gorm:"primaryKey,autoIncrement,not null"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	Type              int    `gorm:"not null"`
	Title             string `gorm:"unique,index,not null"`
	Description       string `gorm:"not null"`
	Note              string
	Categories        []ComposeCategory `gorm:"foreignKey:ID"`
	Platform          string
	Logo              string
	Repository        ComposeRepository
	Envs               []ComposeEnv `gorm:"foreignKey:ID"`
	AdministratorOnly bool
	Name              string
}

// ** Sub tables ** //
type ComposeRepository struct {
	gorm.Model
	ID             int `gorm:"primaryKey,autoIncrement,not null"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
	URL            string `gorm:"not null"`
	Stackfile      string `gorm:"not null"`
	ComposeTableID int
}

type ComposeCategory struct {
	gorm.Model
	ID             int `gorm:"primaryKey,autoIncrement"`
	Name           string
	ComposeTableID int
}

type ComposeEnv struct {
	gorm.Model
	ID                int `gorm:"primaryKey,autoIncrement,not null"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	Name              string `gorm:"not null"`
	Label             string `gorm:"not null"`
	Description       string
	Default           string
	Preset            string
	Selects            []ComposeSelect
	ComposeCategoryID int
}

type ComposeSelect struct {
	gorm.Model
	ID           int `gorm:"primaryKey,autoIncrement,not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Text         string `gorm:"not null"`
	Value        string `gorm:"not null"`
	Default      bool
	ComposeEnvID int
}

// ** GORM table naming ** //
func (ComposeTable) TableName() string {
	return "compose"
}
