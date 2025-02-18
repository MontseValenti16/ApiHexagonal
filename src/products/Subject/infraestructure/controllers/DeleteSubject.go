package controllers

import (
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)



func (s *SubjectController) DeleteSubject(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	if err := s.deleteSubjectUseCase.Execute(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Asignatura eliminada correctamente"})
}