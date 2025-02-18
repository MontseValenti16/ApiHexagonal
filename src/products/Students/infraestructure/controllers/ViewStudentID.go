package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func (c *StudentController) GetStudentById(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	student, err := c.viewStudentIDUseCase.Execute(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, student)
}
