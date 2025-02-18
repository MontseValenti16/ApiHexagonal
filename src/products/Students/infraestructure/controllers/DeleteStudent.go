package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (c *StudentController) DeleteStudent(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := c.deleteStudentUseCase.Execute(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Estudiante eliminado correctamente"})
}
