package tables

import (
	"time"

	"gorm.io/gorm"
)

type Stack struct {
	gorm.Model
	ID                int        `gorm:"primaryKey,autoIncrement"`
	CreatedAt         time.Time  `json:"-"`
	UpdatedAt         time.Time  `json:"-"`
	Type              int        `json:"type"`
	Title             string     `json:"title" gorm:"unique,index"`
	Description       string     `json:"description"`
	Note              string     `json:"note,omitempty"`
	Categories        []string   `json:"categories,omitempty"`
	Platform          string     `json:"platform,omitempty"`
	Logo              string     `json:"logo,omitempty"`
	Repository        Repository `json:"repository"`
	Env               []Env      `json:"env,omitempty"`
	AdministratorOnly bool       `json:"administrator_only,omitempty"`
	Name              string     `json:"name,omitempty"`
}
