package management

import (
	"github.com/dgozalo/aec-remote-executor/pkg/database"
	dbmodels "github.com/dgozalo/aec-remote-executor/pkg/database/model"
	"github.com/pkg/errors"
	"strconv"
)

type ManagementService struct {
	*database.PostgresDBAccess
}

func NewManagementService(db *database.PostgresDBAccess) *ManagementService {
	return &ManagementService{
		db,
	}
}

func (m ManagementService) GetAlumnus() ([]dbmodels.Alumni, error) {
	var alumnus []dbmodels.Alumni
	result := m.DB.Preload("Subjects").Find(&alumnus)
	if result.Error != nil {
		return nil, errors.Wrap(result.Error, "could not retrieve all the executions from the database")
	}
	return alumnus, nil
}

func (m ManagementService) GetAlum(id string) (*dbmodels.Alumni, error) {
	i64Id, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return nil, errors.Wrap(err, "there was an error parsing the execution ID")
	}
	alum := &dbmodels.Alumni{
		ID: int32(i64Id),
	}
	result := m.DB.Preload("Subjects").
		Preload("Subjects.Assignments").
		Preload("Subjects.Assignments.Examples").
		Preload("Subjects.Assignments.CodeTemplates").Find(&alum)
	if result.Error != nil {
		return nil, errors.Wrap(result.Error, "could not retrieve all the executions from the database")
	}
	return alum, nil
}

func (m ManagementService) GetProfessors() ([]dbmodels.Professor, error) {
	var professors []dbmodels.Professor
	result := m.DB.Preload("Subjects").Find(&professors)
	if result.Error != nil {
		return nil, errors.Wrap(result.Error, "could not retrieve all the executions from the database")
	}
	return professors, nil
}

func (m ManagementService) GetProfessor(id string) (*dbmodels.Professor, error) {
	i64Id, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return nil, errors.Wrap(err, "there was an error parsing the execution ID")
	}
	professor := &dbmodels.Professor{
		ID: int32(i64Id),
	}
	result := m.DB.Preload("Subjects").Find(&professor)
	if result.Error != nil {
		return nil, errors.Wrap(result.Error, "could not retrieve all the executions from the database")
	}
	return professor, nil
}

func (m ManagementService) GetAssignment(id string) (*dbmodels.Assignment, error) {
	i64Id, err := strconv.ParseInt(id, 10, 32)
	if err != nil {
		return nil, errors.Wrap(err, "there was an error parsing the execution ID")
	}
	assignment := &dbmodels.Assignment{
		ID: int32(i64Id),
	}
	result := m.DB.Preload("Examples").Preload("CodeTemplates").Find(&assignment)
	if result.Error != nil {
		return nil, errors.Wrap(result.Error, "could not retrieve all the executions from the database")
	}
	return assignment, nil
}

func (m ManagementService) GetAssignmentCodeTemplateForCode(assignmentID string, language string) (*dbmodels.AssignmentCodeTemplate, error) {
	assignment, err := m.GetAssignment(assignmentID)
	if err != nil {
		return nil, errors.Wrap(err, "there was an error retrieving the assignment")
	}
	for _, codeTemplate := range assignment.CodeTemplates {
		if codeTemplate.Language == language {
			return &codeTemplate, nil
		}
	}
	return nil, errors.New("could not find the code template for the assignment and language")
}

func (m ManagementService) GetSubjects() ([]dbmodels.Subject, error) {
	var subjects []dbmodels.Subject
	result := m.DB.Preload("Assignments").Preload("Assignments.Examples").Find(&subjects)
	if result.Error != nil {
		return nil, errors.Wrap(result.Error, "could not retrieve all the executions from the database")
	}
	return subjects, nil
}
