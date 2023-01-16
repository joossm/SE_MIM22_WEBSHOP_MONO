package model

type User struct {
	Id          int
	Username    string
	Password    string
	Firstname   string
	Lastname    string
	HouseNumber string
	Street      string
	ZipCode     string
	City        string
	Email       string
	Phone       string
}

// Json of User
/*{
	"Username":"mjooss",
	"Password":"root",
	"Firstname":"Matthias",
	"Lastname":"Jooss",
	"HouseNumber":"11",
	"Street":"Haertelstr.",
	"ZipCode":"04420",
	"City":"Markranstädt",
	"Email":"jooss.matthias@gmail.com",
	"Phone":"015225444017"
}*/
