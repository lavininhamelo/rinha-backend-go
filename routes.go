package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"rinha-basic/entities"
	"rinha-basic/usecases"
)

func SetupPersonRoutes(r *gin.Engine, usecase usecases.PersonUsecase) {
	r.POST("/pessoas", func(c *gin.Context) {
		var person entities.Person

		if err := c.ShouldBindJSON(&person); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if len(person.Nickname) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Apelido inválido."})
			return
		}

		id, err := usecase.CreatePerson(&person)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.Header("Location", "http://localhost:8080/pessoas/"+id)
		c.JSON(http.StatusCreated, gin.H{"id": id})
	})

	r.GET("/pessoas/:id", func(c *gin.Context) {
		id := c.Param("id")
		person, err := usecase.GetPerson(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, person)
	})

	r.GET("/pessoas", func(c *gin.Context) {
		term := c.DefaultQuery("t", "")
		if len(term) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"message": "O parâmetro 't' é obrigatório"})
			return
		}

		people, err := usecase.SearchPeople(term)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, people)
	})

	r.GET("/contagem-pessoas", func(c *gin.Context) {
		count, err := usecase.CountPeople()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"count": count})
	})
}
