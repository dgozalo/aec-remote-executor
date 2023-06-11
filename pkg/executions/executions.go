package executions

import (
	"github.com/dgozalo/aec-remote-executor/pkg/compiler"
	"github.com/dgozalo/aec-remote-executor/pkg/database"
	dbmodels "github.com/dgozalo/aec-remote-executor/pkg/database/model"
	"github.com/dgozalo/aec-remote-executor/pkg/graph/model"
	"github.com/pkg/errors"
	"gorm.io/gorm/clause"
	"strconv"
)

// ExecutionService is the service that handles the code executions of the application
type ExecutionService struct {
	*database.PostgresDBAccess
}

// NewExecutionService creates a new execution service
func NewExecutionService(db *database.PostgresDBAccess) *ExecutionService {
	return &ExecutionService{
		db,
	}
}

// CreateExecution creates a new execution in the database
func (e ExecutionService) CreateExecution(input model.NewExecution, temporalExec compiler.TemporalExecution) (*dbmodels.Execution, error) {
	i64Id, err := strconv.ParseInt(input.AssignmentID, 10, 32)
	if err != nil {
		return nil, errors.Wrap(err, "there was an error parsing the execution ID")
	}
	execution := dbmodels.Execution{
		Language:     input.Language,
		WorkflowID:   temporalExec.WorkflowID,
		RunID:        temporalExec.RunID,
		Code:         input.Code,
		AssignmentID: int32(i64Id),
	}

	result := e.DB.Clauses(clause.Returning{}).Create(&execution)
	if result.Error != nil {
		return nil, errors.Wrap(result.Error, "could not store the execution in the database")
	}
	return &execution, nil
}

// GetExecutions retrieves all the executions from the database
func (e ExecutionService) GetExecutions() ([]dbmodels.Execution, error) {
	var dbExecutions []dbmodels.Execution
	result := e.DB.Find(&dbExecutions)
	if result.Error != nil {
		return nil, errors.Wrap(result.Error, "could not retrieve all the executions from the database")
	}
	return dbExecutions, nil
}

// GetExecution retrieves a single execution from the database
func (e ExecutionService) GetExecution(id string) (*dbmodels.Execution, error) {
	i64Id, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return nil, errors.Wrap(err, "there was an error parsing the execution ID")
	}
	execution := dbmodels.Execution{
		ID: int32(i64Id),
	}
	result := e.DB.First(&execution)
	if result.Error != nil {
		return nil, errors.Wrap(result.Error, "could not retrieve an execution from the database")
	}
	return &execution, nil
}
