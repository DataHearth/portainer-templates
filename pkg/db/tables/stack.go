package tables

import (
	"time"

	"gorm.io/gorm"
)

// ** Main table ** //
type StackTable struct {
	gorm.Model
	ID                int `gorm:"primaryKey,autoIncrement,not null"`
	CreatedAt         time.Time
	UpdatedAt         time.Time
	Type              int    `gorm:"not null"`
	Title             string `gorm:"unique,index,not null"`
	Description       string `gorm:"not null"`
	Note              string
	Categories        []StackCategory
	Platform          string
	Logo              string
	Repository        StackRepository
	Envs               []StackEnv
	AdministratorOnly bool
	Name              string
}

// ** Sub tables ** //
type StackRepository struct {
	gorm.Model
	ID           int `gorm:"primaryKey,autoIncrement,not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	URL          string `gorm:"not null"`
	Stackfile    string `gorm:"not null"`
	StackTableID int
}

type StackCategory struct {
	gorm.Model
	ID           int `gorm:"primaryKey,autoIncrement"`
	Name         string
	StackTableID int
}

type StackEnv struct {
	gorm.Model
	ID           int `gorm:"primaryKey,autoIncrement,not null"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	Name         string `gorm:"not null"`
	Label        string `gorm:"not null"`
	Description  string
	Default      string
	Preset       string
	Selects       []StackSelect
	StackTableID int
}

type StackSelect struct {
	gorm.Model
	ID         int `gorm:"primaryKey,autoIncrement,not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Text       string `gorm:"not null"`
	Value      string `gorm:"not null"`
	Default    bool
	StackEnvID int
}

// ** GORM table naming ** //
func (StackTable) TableName() string {
	return "stack"
}
