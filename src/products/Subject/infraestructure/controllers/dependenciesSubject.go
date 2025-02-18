package controllers

import (
	"arquitecturahex/src/products/Subject/application"
)

type SubjectController struct {
    createSubjectUseCase *application.CreateSubjectUseCase
    viewSubjectUseCase   *application.ViewSubjectUseCase
    updateSubjectUseCase *application.UpdateSubjectUseCase
    deleteSubjectUseCase *application.DeleteSubjectUseCase
}

func NewSubjectController(
    createSubjectUseCase *application.CreateSubjectUseCase,
    viewSubjectUseCase *application.ViewSubjectUseCase,
    updateSubjectUseCase *application.UpdateSubjectUseCase,
    deleteSubjectUseCase *application.DeleteSubjectUseCase,
) *SubjectController {
    return &SubjectController{
        createSubjectUseCase: createSubjectUseCase,
        viewSubjectUseCase:   viewSubjectUseCase,
        updateSubjectUseCase: updateSubjectUseCase,
        deleteSubjectUseCase: deleteSubjectUseCase,
    }
}