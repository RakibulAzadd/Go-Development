package cmd

import (
	"ecommerce/handlers"
	"ecommerce/util"
	"fmt"
	"net/http"
)

func Serve() {
	mux := http.NewServeMux()

	mux.Handle("GET /products", http.HandlerFunc(handlers.GetProducts))

	mux.Handle("POST /products", http.HandlerFunc(handlers.CreateProduct))

	mux.Handle("GET /products/{productId}",http.HandlerFunc(handlers.GetProductByID))

	fmt.Println("Server is running on port :3021")

	globalRouter := util.GlobalRouter(mux)

	err := http.ListenAndServe(":3021", globalRouter)

	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
