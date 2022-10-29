package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/tipee/account/db/models"
	"github.com/tipee/account/pkg/configs"
)

const (
	MySQLDialect = "mysql"
)

func New() *gorm.DB {
	var (
		client *gorm.DB
		err    error
	)

	if client, err = gorm.Open(MySQLDialect, configs.DBConnectionString()); err != nil {
		panic(err)
	}
	if err = client.DB().Ping(); err != nil {
		panic(err)
	}
	return client
}

//GetById get data by id
func GetById(db *gorm.DB, id int, data interface{ models.ModelMetadata }) error {
	return db.Table(data.GetTableName()).
		Where("id = ?", id).
		First(data).Error
}

//Save data into db
func Save(db *gorm.DB, data interface{ models.ModelMetadata }) error {
	return db.Table(data.GetTableName()).Save(data).Error
}

func Update(db *gorm.DB, data interface{ models.ModelMetadata }, conditions map[string]interface{}) error {
	query := db.Table(data.GetTableName())
	for key, val := range conditions {
		query = query.Where(fmt.Sprintf("%v = ?", key), val)
	}
	return query.Update(data).Error
}

//Filter make a query with where condition
func Filter(db *gorm.DB, data interface{ models.ModelMetadata }, conditions map[string]interface{}) (interface{}, error) {
	query := db.Table(data.GetTableName())
	for key, val := range conditions {
		query = query.Where(fmt.Sprintf("%v = ?", key), val)
	}

	err := query.Find(data).Error
	return data, err
}
