package main

import (
	"context"
	"fmt"
	"pdi-go-kafka-bd/internal/database"

	"pdi-go-kafka-bd/internal/entity"
	users "pdi-go-kafka-bd/internal/repository/users"
)

func main() {
	ctx := context.Background()
	db := database.GetConnection(ctx)
	user := entity.User{
		Name: "dia 7 de junho",
	}
	repository := users.CreateRepository(db)
	resp, err := repository.InsertUser(ctx, user)
	//resp, err := repository.GetUsers(ctx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", resp)
}
