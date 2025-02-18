package controllers

import (
	"arquitecturahex/src/products/Subject/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)


func (s *SubjectController) CreateSubject(c *gin.Context) {
	var subject entities.Subject
	if err := c.ShouldBindJSON(&subject); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	if err := s.createSubjectUseCase.Execute(subject); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Asignatura creada correctamente"})
}