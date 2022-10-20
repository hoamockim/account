package repositories

import (
	"github.com/tipee/account/db"
	"github.com/tipee/account/db/models"
)

type repository struct {
}

//FilterField is used for where clause
type FilterField struct {
	Column    string
	Value     interface{}
	And       interface{}
	Or        interface{}
	condition string // =, !=, >=, <=, >, <
}

var repo repository

func New() *repository {
	return &repo
}

//save
func (repo *repository) save(data interface {
	models.ModelCredential
	models.ModelMetadata
}) error {
	return db.Save(data.GetTableName(), data)
}

//getById
func (repo *repository) getById(data interface {
	models.ModelMetadata
	models.ModelCache
}, id int) error {
	return db.GetById(id, data.GetTableName(), data)
}

//update
func (repo *repository) update(data interface {
	models.ModelCredential
	models.ModelMetadata
}, filter ...FilterField) error {
	var conditions map[string]interface{}
	if len(filter) > 0 {
		for _, val := range filter {
			key := val.Column
			conditions[key] = val.Value
		}
	}
	return db.Update(data.GetTableName(), data, conditions)
}

//filter
func (repo *repository) filter(data interface {
	models.ModelMetadata
	models.ModelCache
}, entities interface{}, filter ...FilterField) error {
	var conditions = make(map[string]interface{})
	if len(filter) > 0 {
		for _, val := range filter {
			key := val.Column
			conditions[key] = val.Value
		}
	}
	return db.Filter(data.GetTableName(), entities, conditions)
}
