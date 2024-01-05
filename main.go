package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type person struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
}

var people = []person{
	{ID: "1", Name: "Azuka", Gender: "M"},
	{ID: "2", Name: "Olisemelie", Gender: "M"},
	{ID: "3", Name: "David", Gender: "M"},
}

// Helper function
func findPersonById(id string) (*person, error) {
	for i, b := range people {
		if b.ID == id {
			return &people[i], nil
		}
	}
	return nil, errors.New("user not found")
}

func getPersonById(context *gin.Context) {
	id := context.Param("id")
	people, err := findPersonById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not Found"})
		return
	}
	context.IndentedJSON(http.StatusFound, people)
}

func getPeople(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, people)
}
func createPeople(context *gin.Context) {
	var newPerson person

	if err := context.BindJSON(&newPerson); err != nil {
		return
	}
	people = append(people, newPerson)
	context.IndentedJSON(http.StatusCreated, people)
}

func main() {
	router := gin.Default()
	router.GET("/person", getPeople)
	router.GET("/person/:id", getPersonById)
	router.POST("/person", createPeople)
	router.Run("localhost:8000")
}
