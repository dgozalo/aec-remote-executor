package service

import (
	"github.com/dgozalo/aec-remote-executor/pkg/compiler"
	"github.com/dgozalo/aec-remote-executor/pkg/database"
	dbmodels "github.com/dgozalo/aec-remote-executor/pkg/database/model"
	"github.com/dgozalo/aec-remote-executor/pkg/graph/model"
	"github.com/pkg/errors"
	"gorm.io/gorm/clause"
	"strconv"
)

type ExecutionService struct {
	*database.PostgresDBAccess
}

func NewExecutionService(db *database.PostgresDBAccess) *ExecutionService {
	return &ExecutionService{
		db,
	}
}

func (e ExecutionService) CreateExecution(input model.NewExecution, temporalExec compiler.TemporalExecution) (*dbmodels.Execution, error) {
	execution := dbmodels.Execution{
		Language:     input.Language,
		WorkflowID:   temporalExec.WorkflowID,
		RunID:        temporalExec.RunID,
		Code:         input.Code,
		AssignmentID: 1,
	}

	result := e.DB.Clauses(clause.Returning{}).Create(&execution)
	if result.Error != nil {
		return nil, errors.Wrap(result.Error, "could not store the execution in the database")
	}
	return &execution, nil
}

func (e ExecutionService) GetExecutions() ([]dbmodels.Execution, error) {
	var dbExecutions []dbmodels.Execution
	result := e.DB.Find(&dbExecutions)
	if result.Error != nil {
		return nil, errors.Wrap(result.Error, "could not retrieve all the executions from the database")
	}
	return dbExecutions, nil
}

func (e ExecutionService) GetExecution(id string) (*dbmodels.Execution, error) {
	i64Id, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return nil, errors.Wrap(err, "there was an error parsing the execution ID")
	}
	execution := dbmodels.Execution{
		ExecutionID: int32(i64Id),
	}
	result := e.DB.First(&execution)
	if result.Error != nil {
		return nil, errors.Wrap(result.Error, "could not retrieve an execution from the database")
	}
	return &execution, nil
}
