package main

import (
	"controleEstoque/internal/database"
	"controleEstoque/internal/products"
	"controleEstoque/internal/routes"
)

func main() {
	db := database.Connect()

	//Auto migrate
	db.AutoMigrate(&products.Product{})

	// DI (injeção de dependência)
	repo := products.NewProductRepository(db)
	service := products.NewProductService(repo)
	handler := products.NewProductHandler(service)

	//router
	r := routes.SetupRouter(handler)

	r.Run(":8080")
}
