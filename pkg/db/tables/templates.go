package tables

import (
	"time"

	"gorm.io/gorm"
)

// ** Global **
type Templates struct {
	Version   string        `json:"version"`
	Templates []interface{} `json:"templates"`
}

type TemplatesArray struct {
	Container []Container
	Compose   []Compose
	Stack     []Stack
}

type Volumes struct {
	gorm.Model
	ID        int       `gorm:"primaryKey,autoIncrement"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	Container string    `json:"container"`
	Bind      string    `json:"bind,omitempty"`
	ReadOnly  bool      `json:"readonly,omitempty"`
}

type Env struct {
	gorm.Model
	ID          int       `gorm:"primaryKey,autoIncrement"`
	CreatedAt   time.Time `json:"-"`
	UpdatedAt   time.Time `json:"-"`
	Name        string    `json:"name"`
	Label       string    `json:"label"`
	Description string    `json:"description,omitempty"`
	Default     string    `json:"default,omitempty"`
	Preset      string    `json:"preset,omitempty"`
	Select      []Select  `json:"select,omitempty"`
}

type Select struct {
	gorm.Model
	ID        int       `gorm:"primaryKey,autoIncrement"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	Text      string    `json:"text"`
	Value     string    `json:"value"`
	Default   bool      `json:"default,omitempty"`
	EnvID     int
}

type Repository struct {
	gorm.Model
	ID        int       `gorm:"primaryKey,autoIncrement"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	URL       string    `json:"url"`
	Stackfile string    `json:"stackfile"`
}

type Label struct {
	gorm.Model
	ID        int       `gorm:"primaryKey,autoIncrement"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	Name      string    `json:"name"`
	Value     string    `json:"value"`
}
