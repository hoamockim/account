package db

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/tipee/account/db/models"
	"github.com/tipee/account/pkg/cache"
	"github.com/tipee/account/pkg/configs"
)

const (
	MySQLDialect = "mysql"
)

//dbFacade Build gorm run with raw query -> auto generate raw query
type dbFacade struct {
	orm   *gorm.DB
	cache cache.Adapter
}

var fcd *dbFacade

func init() {
	var (
		orm *gorm.DB
		err error
	)

	if orm, err = gorm.Open(MySQLDialect, configs.DBConnectionString()); err != nil {
		panic(err)
	}
	if err = orm.DB().Ping(); err != nil {
		panic(err)
	}

	orm.DB().SetMaxOpenConns(300)
	orm.DB().SetMaxIdleConns(10)
	fcd = new(dbFacade)
	fcd.orm = orm
}

//GetById get data by id
func GetById(id int, tableName string, data interface{ models.ModelCache }) error {
	/*	if data.IsCached() {
		if err := fcd.cache.Get(fmt.Sprintf("%v:db:%v:%v", fcd.app, tableName, id), data); err != nil {
			//TODO: log err
		}
	}*/
	return fcd.orm.New().Table(tableName).
		Where("id = ?", id).
		First(data).Error
}

//Save data into db
func Save(tableName string, data interface{}) error {
	return fcd.orm.New().Table(tableName).Save(data).Error
}

func Update(tableName string, data interface{}, conditions map[string]interface{}) error {
	query := fcd.orm.New().Table(tableName)
	for key, val := range conditions {
		query = query.Where(fmt.Sprintf("%v = ?", key), val)
	}
	return query.Update(data).Error
}

//Filter make a query with where condition
func Filter(tableName string, entities interface{}, conditions map[string]interface{}) error {
	query := fcd.orm.New().Table(tableName)
	for key, val := range conditions {
		query = query.Where(fmt.Sprintf("%v = ?", key), val)
	}
	return query.Find(entities).Error
}
