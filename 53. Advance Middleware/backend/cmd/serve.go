package cmd

import (
	"ecommerce/middleware"
	"fmt"
	"net/http"
)

func Serve() {
	manager := middleware.NewManager()

	manager.Use(
		middleware.Preflight, 
		middleware.Cors,
		middleware.Logger, 
	)
		 

	mux := http.NewServeMux()

    wrappedMux :=  manager.WrapMux(mux)

	InitRoute(mux, manager)

	fmt.Println("Server is running on port :3021")

	err := http.ListenAndServe(":3021", wrappedMux)

	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
