package main

import (
	"github.com/dgozalo/aec-remote-executor/pkg/database/model"
	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm"
	"log"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "../query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	db, err := gorm.Open(postgres.Open("postgres://postgres:postgres@localhost:5432/aec_executor_dev?sslmode=disable"), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	g.UseDB(db)
	g.ModelPkgPath = "pkg/database/model"
	g.GenerateModel("alumni", gen.FieldRelateModel(field.Many2Many, "Subjects", model.Subject{}, &field.RelateConfig{
		GORMTag: map[string]string{
			"many2many": "alumni_subjects",
		},
	}))
	g.GenerateModel("professors", gen.FieldRelateModel(field.HasMany, "Subjects", model.Subject{}, &field.RelateConfig{
		GORMTag: map[string]string{
			"foreignKey": "professor_id",
		},
	}))
	g.GenerateModel("assignment_examples")
	g.GenerateModel("assignments", gen.FieldRelateModel(field.HasMany, "Examples", model.AssignmentExample{}, &field.RelateConfig{
		GORMTag: map[string]string{
			"foreignKey": "assignment_id",
		},
	}))
	g.GenerateModel("executions")
	g.GenerateModel("subjects", gen.FieldRelateModel(field.HasMany, "Assignments", model.Assignment{}, &field.RelateConfig{
		GORMTag: map[string]string{
			"foreignKey": "subject_id",
		},
	}))
	g.GenerateModel("alumni_assignments")
	g.GenerateModel("alumni_subjects")

	g.Execute()
}
