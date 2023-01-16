package main

import (
	"log"
	"net/http"
	"time"

	"SE_MIM22_WEBSHOP_MONO/handler"
)

func main() {
	// Server
	var serveMux = http.NewServeMux()
	serveMux.HandleFunc("/login", handler.Login)
	serveMux.HandleFunc("/register", handler.Register)
	serveMux.HandleFunc("/getAllBooks", handler.GetAllBooks)
	serveMux.HandleFunc("/getBookById", handler.GetBookByID)
	serveMux.HandleFunc("/placeOrder", handler.PlaceOrder)
	serveMux.HandleFunc("/getOrdersByUserId", handler.GetOrdersByUserId)
	serveMux.HandleFunc("/initDatabase", handler.InitDatabase)
	log.Printf("About to listen on 8080. Go to http://127.0.0.1:8080/register\n Go to http://127.0.0.1:8080/login")
	server := &http.Server{
		Addr:              ":8080",
		ReadHeaderTimeout: 3 * time.Second,
		WriteTimeout:      3 * time.Second,
		IdleTimeout:       3 * time.Second,
		Handler:           serveMux,
	}
	log.Fatal(server.ListenAndServe())
}
