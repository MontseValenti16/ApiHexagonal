package application

import (
	"arquitecturahex/src/products/Students/domain/entities"
	"arquitecturahex/src/products/Students/domain/repositories"
)

type ViewStudentUseCase struct {
	StudentRepo repositories.StudentRepository
}

func NewViewStudentUseCase(repo repositories.StudentRepository) *ViewStudentUseCase {
	return &ViewStudentUseCase{StudentRepo: repo}
}

func (uc *ViewStudentUseCase) Execute() ([]entities.Student, error) {
	return uc.StudentRepo.GetAllStudents()
}
