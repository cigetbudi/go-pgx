package main

import (
	"go-pgx/api"
	"go-pgx/database"
)

func main() {
	database.Init()
	r := api.InitRoutes()
	r.Run(":6666")
}
