package main

import (
	"app"
	"config"
	"fmt"
	"net/http"
)

func main() {
	config.Load()

	mux := http.NewServeMux()
	app.Mount(mux)
	port := fmt.Sprintf(":%d", config.PORT)
	fmt.Printf("Server running on port %d\n", config.PORT)
	http.ListenAndServe(port, mux)
}
