package database

import (
	"github.com/dgozalo/aec-remote-executor/pkg/graph/model"
	models "github.com/dgozalo/aec-remote-executor/pkg/model"
)

type Database interface {
	CreateExecution(newExec *models.InternalExecution) (*model.Execution, error)
	GetAll() ([]*model.Execution, error)
	GetOne(id string) (*models.InternalExecution, error)
}
