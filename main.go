package main

import (
	"fmt"
	"net/http"
	"strconv"

	"go-crud-api/helper"

	"github.com/gin-gonic/gin"
)

type Person struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Gender string `json:"gender"`
	Amount int    `json:"amount"`
}

var People = []Person{
	{ID: "1", Name: "Azuka", Gender: "M", Amount: 500},
	{ID: "2", Name: "Olisemelie", Gender: "M", Amount: 2100},
	{ID: "3", Name: "Jenifer", Gender: "F", Amount: 1500},
	{ID: "4", Name: "Naomi", Gender: "F", Amount: 140},
}

func withdrawAmountFromPerson(context *gin.Context) {
	amountStr := context.Query("amount")
	id := context.Param("id")
	person, err := helper.FindPersonById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not Found"})
		return
	}
	if amountStr == "" {
		context.IndentedJSON(http.StatusNotAcceptable, gin.H{"message": "Query not found"})
		return
	}
	amount, err := strconv.Atoi(amountStr)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid amount"})
		return
	}
	if person.Amount < amount {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Insufficient Balance"})
		return
	}
	person.Amount -= amount
	fmt.Printf("%v, %T\n", person, person)
	context.IndentedJSON(http.StatusOK, gin.H{"data": fmt.Sprintf("Amount withdrawn: %d Available Balance: %d", amount, person.Amount), "message": "Withdrawal Successful"})
}
func getPersonById(context *gin.Context) {
	id := context.Param("id")
	person, err := helper.FindPersonById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not Found"})
		return
	}
	context.IndentedJSON(http.StatusFound, gin.H{"data": person, "message": "User found"})
}

func getPeople(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, gin.H{"data": People, "message": "All Users"})
}
func createPeople(context *gin.Context) {
	var newPerson Person

	if err := context.BindJSON(&newPerson); err != nil {
		return
	}
	People = append(People, newPerson)
	context.IndentedJSON(http.StatusCreated, People)
}

func main() {
	router := gin.Default()
	router.GET("/person", getPeople)
	router.GET("/person/:id", getPersonById)
	router.POST("/person", createPeople)
	router.POST("/withdraw/:id", withdrawAmountFromPerson)
	router.Run("localhost:8000")
}
