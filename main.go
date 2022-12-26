package main

import (
	"context"
	"go-pgx/api"
	"go-pgx/database"
)

func main() {
	database.Init()
	defer database.DB.Close(context.Background())

	r := api.InitRoutes()
	r.Run(":6666")
}
