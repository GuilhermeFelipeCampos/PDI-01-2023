package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"pdi-go-kafka-bd/internal/database"
	"pdi-go-kafka-bd/internal/entity"
	"syscall"

	users "pdi-go-kafka-bd/internal/repository/users"
)

// func finishApp(signal os.Signal) {
// 	if signal == syscall.SIGINT {
// 		ctx := context.Background()
// 		db := database.GetConnection(ctx)
// 		database.DbConnection.Close(db)
// 		os.Exit(0)
// 	}
// }

func sinal(signal chan os.Signal, exitdb chan int) {
	fmt.Println("Aguardando Sinal...")
	for {
		select {
		case <-signal:
			ctx := context.Background()
			db := database.GetConnection(ctx)
			database.DbConnection.Close(db)
			fmt.Println("Fechando conexão")
			exitdb <- 0
			return
		}

	}

}

func main() {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)
	exitdb := make(chan int, 1)

	ctx := context.Background()
	db := database.GetConnection(ctx)
	user := entity.User{
		Name: "Testando signals",
	}
	repository := users.CreateRepository(db)
	resp, err := repository.InsertUser(ctx, user)
	//resp, err := repository.GetUsers(ctx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%+v\n", resp)
	go sinal(sig, exitdb)

	exitcode := <-exitdb
	fmt.Println("Fechando APP...")
	os.Exit(exitcode)
}
