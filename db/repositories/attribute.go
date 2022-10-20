package repositories

import (
	"github.com/tipee/account/db/models"
)

type UserAttributeRepository interface {
	GetUserAttribute(userId int) (*models.UserAttribute, error)
}

func (repo repository) GetUserAttribute(userId int) (*models.UserAttribute, error) {
	var entity models.UserAttribute
	err := repo.getById(&entity, userId)
	if err != nil {
		return nil, err
	}
	return &entity, nil
}
