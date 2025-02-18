package application

import (
	"arquitecturahex/src/products/Subject/domain/entities"
	"arquitecturahex/src/products/Subject/domain/repositories"
)

type ViewSubjectUseCase struct {
	SubjectRepo repositories.SubjectRepository
}

func NewViewSubjectUseCase(repo repositories.SubjectRepository) *ViewSubjectUseCase {
	return &ViewSubjectUseCase{SubjectRepo: repo}
}

func (uc *ViewSubjectUseCase) Execute() ([]entities.Subject, error) {
	return uc.SubjectRepo.GetAllSubjects()
}
