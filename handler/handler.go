package handler

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"SE_MIM22_WEBSHOP_MONO/model"
	_ "github.com/go-sql-driver/mysql"
)

const post = "POST"

func InitDatabase(responseWriter http.ResponseWriter, request *http.Request) {
	db := openDB()
	defer closeDB(db)
	fmt.Println("init db was executed")

	/*fmt.Println("Creating table books...")
	err, _ := db.Query("USE books")
	if err != nil {
		fmt.Printf("Error Use Books: %s", err)
	}*/

	err, _ := db.Exec("CREATE TABLE IF NOT EXISTS `books` ( `Id` int(11) NOT NULL, `Titel` varchar(45) DEFAULT NULL, `EAN` varchar(45) DEFAULT NULL, `Content` varchar(45) DEFAULT NULL, `Price` float DEFAULT NULL, PRIMARY KEY (`Id`) ) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;")
	if err != nil {
		log.Printf("Error creating table: %s", err)
	}
	err, _ = db.Exec("CREATE TABLE IF NOT EXISTS `users` ( `Id` int(11) NOT NULL, `Username` varchar(45) DEFAULT NULL, `Password` varchar(45) DEFAULT NULL, `Firstname` varchar(45) DEFAULT NULL, `Lastname` varchar(45) DEFAULT NULL, `Housenumber` varchar(45) DEFAULT NULL, `Street` varchar(45) DEFAULT NULL, `Zipcode` varchar(45) DEFAULT NULL, `City` varchar(45) DEFAULT NULL, `Email` varchar(45) DEFAULT NULL, `Phone` varchar(45) DEFAULT NULL, PRIMARY KEY (`Id`) ) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;")
	if err != nil {
		log.Printf("Error creating table: %s", err)
	}

}

func Login(responseWriter http.ResponseWriter, request *http.Request) {
	switch request.Method {
	default:
		responseWriter.Write([]byte("THIS IS A POST REQUEST"))
	case post:
		if request.Body != nil {
			body, _ := ioutil.ReadAll(request.Body)
			user := model.User{}
			jsonErr := json.Unmarshal(body, &user)
			if jsonErr != nil {
				responseWriter.Write([]byte("{ERROR}"))
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
					responseWriter.Write([]byte("{true}"))
					return
				}
			}
			responseWriter.Write([]byte("{false}"))
			return
		}
		responseWriter.Write([]byte("{false}"))
	}
}

func Register(responseWriter http.ResponseWriter, request *http.Request) {
	switch request.Method {
	default:
		responseWriter.Write([]byte("THIS IS A POST REQUEST"))
	case post:
		fmt.Println("Register was executed")
		if request.Body != nil {
			fmt.Println("Body not nil")
			body, _ := ioutil.ReadAll(request.Body)
			user := model.User{}
			jsonErr := json.Unmarshal(body, &user)
			if jsonErr != nil {
				responseWriter.Write([]byte("{ERROR}"))
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
					responseWriter.Write([]byte("{already exists}"))
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
				responseWriter.Write([]byte("{true}"))
				return
			}
		}
	}
}

func GetAllBooks(responseWriter http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
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
		json, err := json.Marshal(books)
		errorHandler(err)
		responseWriter.Write(json)
	default:
		responseWriter.Write([]byte("THIS IS A GET REQUEST"))
	}
}

func GetBookByID(responseWriter http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
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
		json, err := json.Marshal(books)
		errorHandler(err)
		responseWriter.Write(json)
	default:
		responseWriter.Write([]byte("THIS IS A GET REQUEST"))
	}
}

func PlaceOrder(responseWriter http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case post:
		if request.Body != nil {
			body, _ := ioutil.ReadAll(request.Body)
			order := model.Order{}
			jsonErr := json.Unmarshal(body, &order)
			if jsonErr != nil {
				responseWriter.Write([]byte("{ERROR}"))
				return
			}
			db := openDB()
			defer closeDB(db)
			db.Query("INSERT INTO orders (produktId, userId, Amount) VALUES (?, ?, ?)",
				order.ProduktId, order.UserId, order.Amount)
			responseWriter.Write([]byte("{true}"))
			return
		}
	default:
		responseWriter.Write([]byte("THIS IS A POST REQUEST"))
	}
}
func GetOrdersByUserId(responseWriter http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		db := openDB()
		defer closeDB(db)
		result, err := db.Query("SELECT produktId,userId, amount FROM orders WHERE userId = ?", request.URL.Query().Get("id"))
		errorHandler(err)
		var orders []model.Order
		if result != nil {
			for result.Next() {
				var order model.Order
				err = result.Scan(&order.ProduktId, &order.UserId, &order.Amount)
				errorHandler(err)
				orders = append(orders, order)
			}
		}
		json, err := json.Marshal(orders)
		errorHandler(err)
		responseWriter.Write(json)
	default:
		responseWriter.Write([]byte("THIS IS A GET REQUEST"))
	}
}

func closeDB(db *sql.DB) {
	err := db.Close()
	errorHandler(err)
}

func openDB() *sql.DB {
	fmt.Println("Opening DB 2")
	db, err := sql.Open("mysql", "root:root@tcp(mysql:3306)/books")
	fmt.Println(db.Ping())
	fmt.Println(db.Stats())
	fmt.Println("Opening DB 3")
	db.SetMaxIdleConns(0)
	db, err = sql.Open("mysql", "root:root@tcp(docker.for.mac.localhost:3306)/books")
	db.SetMaxIdleConns(0)
	fmt.Println(db.Ping())
	fmt.Println(db.Stats())
	errorHandler(err)
	return db
}
func errorHandler(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
