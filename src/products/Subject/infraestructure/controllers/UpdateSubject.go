package controllers

import (
	"arquitecturahex/src/products/Subject/domain/entities"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
)


func (s *SubjectController) UpdateSubject(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	var subject entities.Subject
	if err := c.ShouldBindJSON(&subject); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Datos inválidos"})
		return
	}

	subject.ID = id
	if err := s.updateSubjectUseCase.Execute(id, subject); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Asignatura actualizada correctamente"})
}