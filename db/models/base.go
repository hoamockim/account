package models

import "time"

type BaseModel struct {
	Id        int
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

type ModelCredential interface {
	Validate() bool
}

type ModelMetadata interface {
	GetTableName() string
}

const (
	DataInvalid string = "data is invalid"
)
