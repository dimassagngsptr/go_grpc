package main

import (
	"fmt"
	"go_grpc/api"
	"go_grpc/app"
	"go_grpc/model"
)

func main() {
	HOST := "localhost"
	PORT := 8000

	//init database
	db := model.InitDatabase("pgx", "0.0.0.0", "user", "password", "database", 5432)
	//service definition
	addService := app.AppService(db)
	grpcAPI := api.InitApi(HOST, PORT, addService, true)
	grpcAPI.Start()
	fmt.Printf("server started in %s:%d", HOST, PORT)
}
