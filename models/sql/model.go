package sql

import (
	"github.com/jinzhu/gorm"
)

const (
	ValueTypeJson = "json"
	ValueTypeInt = "int"
	ValueTypeString = "string"
	ValueTypeFloat = "float"
)

type User struct {
	gorm.Model
	Name             string    `json:"name" gorm:"not null"`
	Password         string    `json:"password" gorm:"not null"`
	Projects         []Project `json:"projects"`
	OperatedProjects []Project `json:"operated_projects" gorm:"many2many:user_projects;"`
}

type Project struct {
	gorm.Model
	Name      string  `json:"name" gorm:"not null"`
	AccessKey string  `json:"access_key" gorm:"not null"`
	UserID    uint    `json:"user_id" gorm:"not null"`
	Operators []User  `json:"operators" gorm:"many2many:user_projects;"`
	Entries   []Entry `json:"entries"`
}

type Entry struct {
	gorm.Model
	Key   string `json:"key" gorm:"not null;index"`
	Value string `json:"value" gorm:"not null;type:text;"`
	Type  string `json:"type" gorm:"not null;default:'string'"`
	ProjectID uint `json:"project_id"`
}
