// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameProfessor = "professors"

// Professor mapped from table <professors>
type Professor struct {
	ProfessorID int32  `gorm:"column:professor_id;primaryKey;autoIncrement:true" json:"professor_id"`
	FirstName   string `gorm:"column:first_name" json:"first_name"`
	LastName    string `gorm:"column:last_name" json:"last_name"`
	Email       string `gorm:"column:email" json:"email"`
}

// TableName Professor's table name
func (*Professor) TableName() string {
	return TableNameProfessor
}
