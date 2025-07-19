package main

import (
	"ComputerInfoAPI/routes"
	"log"
	"net/http"
)

func main() {
	router := routes.NewRouter()

	// 设置HTTP服务器
	server := &http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Println("Starting server on port 8080")
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Printf("Server failed to start: %v", err)
		log.Println("Shutting down server...")
	}
}
