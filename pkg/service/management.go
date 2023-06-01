package service

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
	result := m.DB.Preload("Subjects").Preload("Subjects.Assignments").Find(&alum)
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

func (m ManagementService) GetSubjects() ([]dbmodels.Subject, error) {
	var subjects []dbmodels.Subject
	result := m.DB.Preload("Assignments").Find(&subjects)
	if result.Error != nil {
		return nil, errors.Wrap(result.Error, "could not retrieve all the executions from the database")
	}
	return subjects, nil
}
