# MONOLITH BOOKSTORE

    About to listen on Port: 8080.
    
    SUPPORTED REQUESTS:
    GET:
    Get All Books: http://127.0.0.1:8080/getAllBooks
    Get Book By ID: http://127.0.0.1:8080/getBookById?id=1 requiers a url parameter id
    Go to http://127.0.0.1:8080/init to initialise the Database.
    Get Order By ID: http://127.0.0.1:8080/getOrdersByUserId?id=1 requiers a url parameter id
    Create Error on: http://127.0.0.1:8080/error
    
    
    POST:
    Login on: http://127.0.0.1:8080/login requires a JSON Body with the following format:
    {
    "Username": "mmuster",
    "Password": "password"
    }
    Register on: http://127.0.0.1:8080/register requires a JSON Body with the following format:
    {
    "Username": "mmuster",
    "Password": "password",
    "Firstname": "Max",
    "Lastname": "Muster",
    "Housenumber": "1",
    "Street": "Musterstr.",
    "Zipcode": "01234",
    "City": "Musterstadt",
    "Email": "max.muster@mail.com",
    "Phone": "012345678910"
    }
    Place Order: http://127.0.0.1:8080/placeOrder requiers a Body with following json:
    {
    "produktId": "1",
    "userId": "1",
    "amount": "1"
    }

