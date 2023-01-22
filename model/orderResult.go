package model

type OrderResult struct {
	BasketID string
	Books    []BookAndAmount
	UserId   string
}
