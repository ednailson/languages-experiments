package main

import (
	"encoding/json"
	"log"

	"github.com/go-playground/validator/v10"
)

type User struct {
	// "validate" tag is the one used by the library to filter the rules
	// On ID field we set up as required field (the string cannot be empty) and it should also be an uuid
	ID string `validate:"required,uuid"`
	// Email is also required but the string should have email format
	Email string `validate:"required,email"`
	// Age should be greater or equal than 18
	Age uint `validate:"gte=18"`
	// Extra does not have any validation
	Extra string
}

func main() {
	// Initiation of validator
	validate := validator.New()

	var user1 User
	// In this example there isn't ID field, so the validation is going to return error
	_ = json.Unmarshal([]byte(`{"email": "junior@test.com", "age": 21}`), &user1)

	// On the unmarshal user1.ID is empty, so it will return error on the validation
	err := validate.Struct(user1)
	if err != nil {
		log.Printf("user1 validation error happened as predicted, error: %s", err.Error())
	}

	// User that respects all the rules, so it will pass on the validation.
	user2 := User{
		ID:    "b194a096-af97-43a7-9e29-27ad8c395322",
		Email: "junior@test.com",
		Age:   18,
		Extra: "",
	}

	err = validate.Struct(user2)
	log.Printf("user2 respects all the rules for all fields, so here error is equal to nil, error: %v", err)
}
