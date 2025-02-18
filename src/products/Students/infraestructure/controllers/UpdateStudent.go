package controllers

import (
	"arquitecturahex/src/products/Students/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (c *StudentController) UpdateStudent(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var student entities.Student
	if err := ctx.ShouldBindJSON(&student); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	if err := c.updateStudentUseCase.Execute(id, student); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Estudiante actualizado correctamente"})
}
