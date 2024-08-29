package entity

import "time"

type User struct {
	ID       	int
	Username 	string
	Password 	string
	Email    	string
}

type Product struct {
	ID          int
	Name        string
	Description string
	Price       float64
	Quantity    int
}

type Purchase struct {
	ID        	int
	UserID    	int
	ProductID 	int
	Quantity 	int
	Timestamp 	time.Time
}