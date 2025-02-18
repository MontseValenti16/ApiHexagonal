package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
)


func (s *SubjectController) GetAllSubjects(c *gin.Context) {
	subjects, err := s.viewSubjectUseCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, subjects)
}