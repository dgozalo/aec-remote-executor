package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gen"
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
	g.GenerateModel("alumni")
	g.GenerateModel("professors")
	g.GenerateModel("assignments")
	g.GenerateModel("executions")
	g.GenerateModel("subjects")
	g.GenerateModel("alumni_assignments")
	g.GenerateModel("alumni_subjects")

	g.Execute()
}
