package main

import (
	"controleEstoque/internal/database"
	"controleEstoque/internal/routes"
)

func main() {
	database.GetDB()

	routes.InitRoutes()

	database.CloseDB(database.GetDB())
}
