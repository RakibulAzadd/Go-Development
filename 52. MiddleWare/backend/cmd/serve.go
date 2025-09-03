package cmd

import (
	"ecommerce/middleware"
	"ecommerce/util"
	"fmt"
	"net/http"
)

func Serve() {
	manager := middleware.NewManager()

	manager.Use(middleware.Logger, middleware.Hudai)

	mux := http.NewServeMux()

	InitRoute(mux, manager)

	fmt.Println("Server is running on port :3021")

	globalRouter := util.GlobalRouter(mux)

	err := http.ListenAndServe(":3021", globalRouter)

	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}
