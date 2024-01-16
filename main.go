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

func transferAmountToOtherPerson(context *gin.Context) {
	type requestBody struct {
		Sender_ID   string `json:"sender_id"`
		Receiver_ID string `json:"receiver_id"`
		Amount      int    `json:"amount"`
	}
	var newRequest requestBody

	if err := context.BindJSON(&newRequest); err != nil {
		return
	}

	if newRequest.Amount == 0 || newRequest.Sender_ID == "" || newRequest.Receiver_ID == "" {
		context.IndentedJSON(http.StatusNotAcceptable, gin.H{"message": "Missing Data. Please check input again"})
		return
	}
	reciever, err := helper.FindPersonById(newRequest.Receiver_ID)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Receiver not Found"})
		return
	}
	sender, err := helper.FindPersonById(newRequest.Sender_ID)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "Sender not Found"})
		return
	}
	if newRequest.Amount > sender.Amount {
		context.IndentedJSON(http.StatusNotAcceptable, gin.H{"message": "Insufficient Balance"})
		return
	}
	reciever.Amount += newRequest.Amount
	sender.Amount -= newRequest.Amount
	context.IndentedJSON(http.StatusOK, gin.H{"data": fmt.Sprintf("Amount Transfered: %d to %s from %s", newRequest.Amount, reciever.Name, sender.Name), "message": "Deposit Successful"})
}
func withdrawOrDepositAmountFromUser(context *gin.Context) {
	amountStr := context.Query("amount")
	typeOfTransaction := context.Query("operation")
	id := context.Param("id")
	person, err := helper.FindPersonById(id)

	if err != nil {
		context.IndentedJSON(http.StatusNotFound, gin.H{"message": "User not Found"})
		return
	}
	if amountStr == "" || typeOfTransaction == "" {
		context.IndentedJSON(http.StatusNotAcceptable, gin.H{"message": "Query not found"})
		return
	}
	amount, err := strconv.Atoi(amountStr)
	if err != nil {
		context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid amount"})
		return
	}
	// If transaction is deposit, it adds to users amount
	// If transaction is withdrawal, it deducts from users amount
	switch {
	case typeOfTransaction == "withdraw":
		if person.Amount < amount {
			context.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Insufficient Balance"})
			return
		}
		person.Amount -= amount
		context.IndentedJSON(http.StatusOK, gin.H{"data": fmt.Sprintf("Amount withdrawn: %d Available Balance: %d", amount, person.Amount), "message": "Withdrawal Successful"})
	case typeOfTransaction == "deposit":
		person.Amount += amount
		context.IndentedJSON(http.StatusOK, gin.H{"data": fmt.Sprintf("Amount deposited: %d Available Balance: %d", amount, person.Amount), "message": "Deposit Successful"})
	}
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
	router.POST("/pay/:id", withdrawOrDepositAmountFromUser)
	router.POST("/transfer", transferAmountToOtherPerson)
	router.Run("localhost:8000")
}
