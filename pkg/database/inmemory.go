package database

import (
	"github.com/dgozalo/aec-remote-executor/pkg/graph/model"
	models "github.com/dgozalo/aec-remote-executor/pkg/model"
	"github.com/google/uuid"
)

type InMemoryDB map[string]*models.InternalExecution

type InMemoryDBAccess struct {
	db InMemoryDB
}

func NewInMemoryDBAccess() *InMemoryDBAccess {
	return &InMemoryDBAccess{
		InMemoryDB{},
	}
}

func (d InMemoryDBAccess) CreateExecution(newExec *models.InternalExecution) (*model.Execution, error) {
	id := uuid.New()
	d.db[id.String()] = newExec
	return &model.Execution{
		ID:       id.String(),
		Language: newExec.Language,
		Code:     newExec.Code,
	}, nil
}

func (d InMemoryDBAccess) GetAll() ([]*model.Execution, error) {
	execs := make([]*model.Execution, 0, len(d.db))
	for k, v := range d.db {
		execs = append(execs, &model.Execution{
			ID:       k,
			Language: v.Language,
			Code:     v.Code,
		})
	}
	return execs, nil
}

func (d InMemoryDBAccess) GetOne(id string) (*models.InternalExecution, error) {
	exec, exists := d.db[id]
	if !exists {
		return nil, nil
	}
	return exec, nil
}
