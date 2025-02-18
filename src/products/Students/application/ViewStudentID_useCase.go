package application

import (
	"arquitecturahex/src/products/Students/domain/entities"
	"arquitecturahex/src/products/Students/domain/repositories"
)

type ViewStudentIdUseCase struct {
	StudentRepository repositories.StudentRepository
}

func NewViewStudentIdUseCase(repo repositories.StudentRepository) *ViewStudentIdUseCase {
	return &ViewStudentIdUseCase{StudentRepository: repo}
}

func (u *ViewStudentIdUseCase) Execute(id int) (*entities.Student, error) {
	return u.StudentRepository.GetById(id)
}
