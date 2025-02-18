package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)
func (c *StudentController) GetStudents(ctx *gin.Context) {
	students, err := c.viewStudentUseCase.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, students)
}
