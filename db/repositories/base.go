package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/tipee/account/db"
	"github.com/tipee/account/db/models"
)

type (
	repo struct {
		database *gorm.DB
	}

	//default repository
	Repository interface {
	}
)

//FilterField is used for where clause
type FilterField struct {
	Column    string
	Value     interface{}
	And       interface{}
	Or        interface{}
	condition string // =, !=, >=, <=, >, <
}

func New(db *gorm.DB) *repo {
	return &repo{
		database: db,
	}
}

//save
func (repo *repo) save(data interface {
	models.ModelCredential
	models.ModelMetadata
}) error {
	return db.Save(repo.database, data)
}

//getById
func (repo *repo) getById(data interface {
	models.ModelMetadata
}, id int) error {
	return db.GetById(repo.database, id, data)
}

//update
func (repo *repo) update(data interface {
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
	return db.Update(repo.database, data, conditions)
}

//filter
func (repo *repo) filter(data interface {
	models.ModelMetadata
}, filter ...FilterField) (interface{}, error) {
	var conditions = make(map[string]interface{})
	if len(filter) > 0 {
		for _, val := range filter {
			key := val.Column
			conditions[key] = val.Value
		}
	}
	return db.Filter(repo.database, data, conditions)
}
