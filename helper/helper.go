package helper

import (
	"errors"
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

// Helper function
func FindPersonById(id string) (*Person, error) {
	for i, b := range People {
		if b.ID == id {
			return &People[i], nil
		}
	}
	return nil, errors.New("user not found")
}
