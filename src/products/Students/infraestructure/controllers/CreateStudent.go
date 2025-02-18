package controllers

import (
	"arquitecturahex/src/products/Students/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (c *StudentController) CreateStudent(ctx *gin.Context) {
	var student entities.Student
	if err := ctx.ShouldBindJSON(&student); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	if err := c.createStudentUseCase.Execute(student); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Estudiante creado correctamente"})
}
