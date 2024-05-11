// mysqldbmodels/teacher.go

package mysqldbmodels

import "fmt"

type TeacherRepository interface {
	CreateTeacher(teacher *Teacher) error
	GetTeacherByID(id uint) (*Teacher, error)
	// Add other methods as needed
}

type Teacher struct {
	ID   uint   `gorm:"primaryKey"`
	Name string
}

func MainTeacher(repo TeacherRepository) {
	// Create a new teacher
	teacher := &Teacher{Name: "Shubham"}
	if err := repo.CreateTeacher(teacher); err != nil {
		panic("failed to create teacher")
	}
	fmt.Println("Teacher created successfully")

	// Read teacher
	retrievedTeacher, err := repo.GetTeacherByID(teacher.ID)
	if err != nil {
		panic("failed to retrieve teacher")
	}
	fmt.Printf("Retrieved Teacher: %+v\n", *retrievedTeacher)

	// Add other CRUD operations
}
