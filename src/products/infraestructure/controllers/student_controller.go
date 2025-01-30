package controllers

import (
    "arquitecturahex/src/products/application"
    "arquitecturahex/src/products/domain/entities"
    "github.com/gin-gonic/gin"
    "net/http"
    "strconv"
)
	
	type StudentController struct {
		createStudentUseCase *application.CreateStudentUseCase
		viewStudentUseCase   *application.ViewStudentUseCase
		updateStudentUseCase *application.UpdateStudentUseCase
		deleteStudentUseCase *application.DeleteStudentUseCase
	}
	
	func NewStudentController(
		createStudentUseCase *application.CreateStudentUseCase,
		viewStudentUseCase *application.ViewStudentUseCase,
		updateStudentUseCase *application.UpdateStudentUseCase,
		deleteStudentUseCase *application.DeleteStudentUseCase,
	) *StudentController {
		return &StudentController{
			createStudentUseCase: createStudentUseCase,
			viewStudentUseCase:   viewStudentUseCase,
			updateStudentUseCase: updateStudentUseCase,
			deleteStudentUseCase: deleteStudentUseCase,
		}
	}

func (s *StudentController) CreateStudent(c *gin.Context) {
    var student entities.Student
    if err := c.ShouldBindJSON(&student); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
        return
    }

    if err := s.createStudentUseCase.Execute(student); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo guardar el estudiante"})
        return
    }

    c.JSON(http.StatusCreated, gin.H{"message": "Estudiante creado correctamente"})
}

func (s *StudentController) GetAllStudents(c *gin.Context) {
    students, err := s.viewStudentUseCase.Execute()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudieron obtener los estudiantes"})
        return
    }
    c.JSON(http.StatusOK, students)
}

func (s *StudentController) UpdateStudent(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
        return
    }

    var student entities.Student
    if err := c.ShouldBindJSON(&student); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
        return
    }

    student.ID = id
    if err := s.updateStudentUseCase.Execute(id, student); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo actualizar el estudiante"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Estudiante actualizado correctamente"})
}

func (s *StudentController) DeleteStudent(c *gin.Context) {
    id, err := strconv.Atoi(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
        return
    }

    if err := s.deleteStudentUseCase.Execute(id); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo eliminar el estudiante"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Estudiante eliminado correctamente"})
}