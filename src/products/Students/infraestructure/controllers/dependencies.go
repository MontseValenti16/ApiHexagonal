package controllers

import(
	"arquitecturahex/src/products/Students/application"
)

type StudentController struct {
	createStudentUseCase *application.CreateStudentUseCase
	viewStudentUseCase   *application.ViewStudentUseCase
	updateStudentUseCase *application.UpdateStudentUseCase
	deleteStudentUseCase *application.DeleteStudentUseCase
	viewStudentIDUseCase *application.ViewStudentIdUseCase 
}

func NewStudentController(
	createStudentUseCase *application.CreateStudentUseCase,
	viewStudentUseCase *application.ViewStudentUseCase,
	updateStudentUseCase *application.UpdateStudentUseCase,
	deleteStudentUseCase *application.DeleteStudentUseCase,
	viewStudentIDUseCase *application.ViewStudentIdUseCase,
) *StudentController {
	return &StudentController{
		createStudentUseCase: createStudentUseCase,
		viewStudentUseCase:   viewStudentUseCase,
		updateStudentUseCase: updateStudentUseCase,
		deleteStudentUseCase: deleteStudentUseCase,
		viewStudentIDUseCase: viewStudentIDUseCase,
	}
}
