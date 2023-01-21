package main

import (
	"log"
	"net/http"
	"time"

	"SE_MIM22_WEBSHOP_MONO/handler"
)

func main() {
	var serveMux = http.NewServeMux()
	serveMux.HandleFunc("/login", handler.Login)
	serveMux.HandleFunc("/register", handler.Register)
	serveMux.HandleFunc("/getAllBooks", handler.GetAllBooks)
	serveMux.HandleFunc("/getBookById", handler.GetBookByID)
	serveMux.HandleFunc("/placeOrder", handler.PlaceOrder)
	serveMux.HandleFunc("/getOrdersByUserId", handler.GetOrdersByUserId)
	serveMux.HandleFunc("/initDatabase", handler.InitDatabase)
	serveMux.HandleFunc("/error", handler.Error)
	log.Printf("\n\n\tMONOLITH BOOKSTORE\n\n" +
		"About to listen on Port: 8080." +
		"\n\nSUPPORTED REQUESTS:" +
		"\nGET:" +
		"\nGet All Books: http://127.0.0.1:8080/getAllBooks" +
		"\nGet Book By ID: http://127.0.0.1:8080/getBookById?id=1 requiers a url parameter id" +
		"\nGo to http://127.0.0.1:8080/init to initialise the Database." +
		"\nGet Order By ID: http://127.0.0.1:8080/getOrdersByUserId?id=1 requiers a url parameter id" +
		"\nCreate Error on: http://127.0.0.1:8080/error\n" +
		"\n\nPOST:" +
		"\nLogin on: http://127.0.0.1:8080/login requires a JSON Body with the following format:\n" +
		"{\n    \"Username\": \"mmuster\",\n    \"Password\": \"password\"\n}" +
		"\nRegister on: http://127.0.0.1:8080/register requires a JSON Body with the following format:\n" +
		"{\n    \"Username\": \"mmuster\",\n    \"Password\": \"password\",\n    \"Firstname\": \"Max\",\n   " +
		" \"Lastname\": \"Muster\",\n    \"Housenumber\": \"1\",\n    \"Street\": \"Musterstr.\",\n  " +
		"  \"Zipcode\": \"01234\",\n    \"City\": \"Musterstadt\",\n    \"Email\": \"max.muster@mail.com\",\n  " +
		"  \"Phone\": \"012345678910\"\n  }" +
		"\nPlace Order: http://127.0.0.1:8080/placeOrder requiers a Body with following json:\n{\n    \"produktId\": \"1\",\n    \"userId\": \"1\",\n    \"amount\": \"1\"\n}")
	server := &http.Server{
		Addr:              ":8080",
		ReadHeaderTimeout: 3 * time.Second,
		WriteTimeout:      3 * time.Second,
		IdleTimeout:       3 * time.Second,
		Handler:           serveMux,
	}
	log.Fatal(server.ListenAndServe())
}
