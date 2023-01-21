package handler

import (
	"SE_MIM22_WEBSHOP_MONO/model"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"io"
	"log"
	"net/http"
	"strconv"
)

const post, get = "POST", "GET"

func InitDatabase(responseWriter http.ResponseWriter, request *http.Request) {

	switch request.Method {
	case get:
		db := openDB()
		defer closeDB(db)
		fmt.Println("init db was executed")

		_, err := db.Exec("CREATE TABLE IF NOT EXISTS `books` ( `Id` int(11) NOT NULL, `Titel` varchar(45) DEFAULT NULL, `EAN` varchar(45) DEFAULT NULL, `Content` varchar(45) DEFAULT NULL, `Price` float DEFAULT NULL, PRIMARY KEY (`Id`) ) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;")
		if err != nil {
			log.Printf("Error creating table: %s", err)
		}
		_, err = db.Exec("CREATE TABLE IF NOT EXISTS `users` ( `Id` int(11) NOT NULL, `Username` varchar(45) DEFAULT NULL, `Password` varchar(45) DEFAULT NULL, `Firstname` varchar(45) DEFAULT NULL, `Lastname` varchar(45) DEFAULT NULL, `Housenumber` varchar(45) DEFAULT NULL, `Street` varchar(45) DEFAULT NULL, `Zipcode` varchar(45) DEFAULT NULL, `City` varchar(45) DEFAULT NULL, `Email` varchar(45) DEFAULT NULL, `Phone` varchar(45) DEFAULT NULL, PRIMARY KEY (`Id`) ) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;")
		if err != nil {
			log.Printf("Error creating table: %s", err)
		}
		_, err = db.Exec("CREATE TABLE IF NOT EXISTS `orders` ( `id` int(11) NOT NULL AUTO_INCREMENT, `produktId` varchar(45) DEFAULT NULL, `userId` varchar(45) DEFAULT NULL, `amount` varchar(45) DEFAULT NULL, PRIMARY KEY (`id`) ) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;")
		if err != nil {
			log.Printf("Error creating table: %s", err)
		}
		js, err := json.Marshal("Success")
		errorHandler(err)
		_, responseErr := responseWriter.Write(js)
		errorHandler(responseErr)
		return
	default:
		js, err := json.Marshal("THIS IS A GET REQUEST")
		errorHandler(err)
		_, responseErr := responseWriter.Write(js)
		errorHandler(responseErr)
		return
	}

}

func Login(responseWriter http.ResponseWriter, request *http.Request) {

	switch request.Method {
	case post:
		if request.Body != nil {
			body, _ := io.ReadAll(request.Body)
			user := model.User{}
			jsonErr := json.Unmarshal(body, &user)
			if jsonErr != nil {
				js, err := json.Marshal("Error")
				errorHandler(err)
				_, responseErr := responseWriter.Write(js)
				errorHandler(responseErr)
				return
			}
			db := openDB()
			defer closeDB(db)
			result, err := db.Query("SELECT Id, Username, Password FROM users WHERE Username = ? AND Password = ?", user.Username, user.Password)
			errorHandler(err)
			var users []model.User
			if result != nil {
				for result.Next() {
					var user model.User
					err = result.Scan(&user.Id, &user.Username, &user.Password)
					errorHandler(err)
					users = append(users, user)
				}
			}
			for _, iUser := range users {
				fmt.Println(user.Username + " " + user.Password)
				fmt.Println(iUser.Username + " " + iUser.Password)
				if iUser.Username == user.Username && iUser.Password == user.Password {
					js, err := json.Marshal("true")
					errorHandler(err)
					_, responseErr := responseWriter.Write(js)
					errorHandler(responseErr)
					return
				}
			}
			js, err := json.Marshal("false")
			errorHandler(err)
			_, responseErr := responseWriter.Write(js)
			errorHandler(responseErr)
			return
		}
		js, err := json.Marshal("false")
		errorHandler(err)
		_, responseErr := responseWriter.Write(js)
		errorHandler(responseErr)
		return
	default:
		js, err := json.Marshal("THIS IS A POST REQUEST")
		errorHandler(err)
		_, responseErr := responseWriter.Write(js)
		errorHandler(responseErr)
		return
	}
}

/*responseWriter.Header().Set("Content-Type", "application/json")
responseWriter.Write(js)*/

func Register(responseWriter http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case post:
		fmt.Println("Register was executed")
		if request.Body != nil {
			fmt.Println("Body not nil")
			body, _ := io.ReadAll(request.Body)
			user := model.User{}
			jsonErr := json.Unmarshal(body, &user)
			if jsonErr != nil {
				js, err := json.Marshal("Error")
				errorHandler(err)
				_, responseErr := responseWriter.Write(js)
				errorHandler(responseErr)
				return
			}
			fmt.Println("No json error")
			db := openDB()
			defer closeDB(db)
			result, err := db.Query("SELECT Username FROM users WHERE Username = ?", user.Username)
			fmt.Println("result: ", result)
			errorHandler(err)
			fmt.Println("Query executed")
			var users []model.User
			if result.Next() == true {
				for result.Next() {
					var user model.User
					err = result.Scan(&user.Id, &user.Username, &user.Password)
					fmt.Println("user: ", user.Username, user.Password)
					users = append(users, user)
				}
				if users != nil {
					js, err := json.Marshal("already exists")
					errorHandler(err)
					_, responseErr := responseWriter.Write(js)
					errorHandler(responseErr)
					return
				}
			} else {
				// GET MAX ID
				result, err := db.Query("SELECT MAX(Id) FROM users")
				errorHandler(err)
				var maxId int
				if result != nil {
					for result.Next() {
						err = result.Scan(&maxId)
						errorHandler(err)
					}
				}
				maxId++
				fmt.Println("result is nil | execute insert")
				res, err := db.Query("INSERT INTO users (Id, Username, Password, Firstname, Lastname, HouseNumber, Street, ZipCode, City, Email, Phone) VALUES (?,?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
					maxId, user.Username, user.Password, user.Firstname, user.Lastname, user.HouseNumber, user.Street, user.ZipCode, user.City, user.Email, user.Phone)
				fmt.Println(res)
				errorHandler(err)
				js, err := json.Marshal("true")
				_, responseErr := responseWriter.Write(js)
				errorHandler(responseErr)
				return
			}
		}
	default:
		js, err := json.Marshal("THIS IS A POST REQUEST")
		errorHandler(err)
		_, responseErr := responseWriter.Write(js)
		errorHandler(responseErr)
		return
	}
}

func GetAllBooks(responseWriter http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case get:
		db := openDB()
		defer closeDB(db)
		result, err := db.Query("SELECT * FROM books")
		errorHandler(err)
		var books []model.Book
		if result != nil {
			for result.Next() {
				var book model.Book
				err = result.Scan(&book.Id, &book.Titel, &book.EAN, &book.Content, &book.Price)
				errorHandler(err)
				books = append(books, book)
			}
		}
		jsonBook, err := json.Marshal(books)
		errorHandler(err)
		_, responseErr := responseWriter.Write(jsonBook)
		errorHandler(responseErr)
		return
	default:
		js, err := json.Marshal("THIS IS A GET REQUEST")
		errorHandler(err)
		_, responseErr := responseWriter.Write(js)
		errorHandler(responseErr)
		return
	}
}

func GetBookByID(responseWriter http.ResponseWriter, request *http.Request) {

	switch request.Method {
	case get:
		db := openDB()
		defer closeDB(db)
		result, err := db.Query("SELECT * FROM books WHERE Id = ?", request.URL.Query().Get("id"))
		errorHandler(err)
		var books []model.Book
		if result != nil {
			for result.Next() {
				var book model.Book
				err = result.Scan(&book.Id, &book.Titel, &book.EAN, &book.Content, &book.Price)
				errorHandler(err)
				books = append(books, book)
			}
		}
		jsonBook, err := json.Marshal(books)
		errorHandler(err)
		_, responseErr := responseWriter.Write(jsonBook)
		errorHandler(responseErr)
		return
	default:
		js, err := json.Marshal("THIS IS A GET REQUEST")
		errorHandler(err)
		_, responseErr := responseWriter.Write(js)
		errorHandler(responseErr)
		return
	}
}

func PlaceOrder(responseWriter http.ResponseWriter, request *http.Request) {

	switch request.Method {
	case post:
		if request.Body != nil {
			body, _ := io.ReadAll(request.Body)
			order := model.Order{}
			jsonErr := json.Unmarshal(body, &order)
			if jsonErr != nil {
				js, err := json.Marshal("Error")
				errorHandler(err)
				_, responseErr := responseWriter.Write(js)
				errorHandler(responseErr)
				return
			}
			db := openDB()
			defer closeDB(db)
			_, insertErr := db.Query("INSERT INTO orders (produktId, userId, Amount) VALUES (?, ?, ?)",
				order.ProduktId, order.UserId, order.Amount)
			errorHandler(insertErr)
			js, err := json.Marshal("true")
			errorHandler(err)
			_, responseErr := responseWriter.Write(js)
			errorHandler(responseErr)
			return
		}
	default:
		js, err := json.Marshal("THIS IS A POST REQUEST")
		errorHandler(err)
		_, responseErr := responseWriter.Write(js)
		errorHandler(responseErr)
		return
	}
}

type BookAndAmount struct {
	Book   model.Book
	Amount string
}

type orderResult struct {
	BasketID string
	Books    []BookAndAmount
	UserId   string
}

func GetOrdersByUserId(responseWriter http.ResponseWriter, request *http.Request) {

	switch request.Method {
	case get:
		db := openDB()
		defer closeDB(db)
		result, err := db.Query("SELECT id, produktId, userId, amount FROM orders WHERE userId = ?", request.URL.Query().Get("id"))
		errorHandler(err)
		var orders []model.Order
		if result != nil {
			for result.Next() {
				var order model.Order
				err = result.Scan(&order.Id, &order.ProduktId, &order.UserId, &order.Amount)
				errorHandler(err)
				orders = append(orders, order)
			}
		}
		// find orders from same customer
		var orderFromSameCustomer []model.Order
		for _, order := range orders {
			if order.UserId == request.URL.Query().Get("id") {
				orderFromSameCustomer = append(orderFromSameCustomer, order)
			}
		}
		// sum amount of same books from same customer
		var orderFromSameCustomerWithSummedAmount []model.Order
		for _, order := range orderFromSameCustomer {
			for i, order2 := range orderFromSameCustomer {
				if order.ProduktId == order2.ProduktId {
					orderA1, _ := strconv.Atoi(order.Amount)
					orderA2, _ := strconv.Atoi(order2.Amount)
					orderAGesamt := orderA1 + orderA2
					order.Amount = strconv.Itoa(orderAGesamt)
					// remove doubled order from order
					orderFromSameCustomer = append(orderFromSameCustomer[:i], orderFromSameCustomer[i+1:]...)
				}
			}
		}
		orderFromSameCustomerWithSummedAmount = append(orderFromSameCustomerWithSummedAmount, orderFromSameCustomer...)
		var results []orderResult
		for _, order := range orderFromSameCustomerWithSummedAmount {
			// get book from order
			var books []model.Book
			resultBook, err := db.Query("SELECT * FROM books WHERE Id = ?", order.ProduktId)
			errorHandler(err)
			if resultBook != nil {
				for resultBook.Next() {
					var book model.Book
					err = resultBook.Scan(&book.Id, &book.Titel, &book.EAN, &book.Content, &book.Price)
					errorHandler(err)
					books = append(books, book)
				}
			}
			var result orderResult
			result.BasketID = order.Id
			result.UserId = order.UserId
			result.Books = append(result.Books, BookAndAmount{Book: books[0], Amount: order.Amount})
			results = append(results, result)
		}

		jsonResults, err := json.Marshal(results)
		errorHandler(err)
		_, responseErr := responseWriter.Write(jsonResults)
		errorHandler(responseErr)
		return
	default:
		js, err := json.Marshal("THIS IS A GET REQUEST")
		errorHandler(err)
		_, responseErr := responseWriter.Write(js)
		errorHandler(responseErr)
		return
	}
}

func Error(responseWriter http.ResponseWriter, request *http.Request) {
	// This is just a test function to create an error
	Error(responseWriter, request)
	panic("ERROR")
}

func closeDB(db *sql.DB) {
	err := db.Close()
	errorHandler(err)
}

func openDB() *sql.DB {
	fmt.Println("Opening DB")
	db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/books")
	fmt.Println(db.Ping())
	fmt.Println(db.Stats())
	db.SetMaxIdleConns(0)
	errorHandler(err)
	return db
}

func errorHandler(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
