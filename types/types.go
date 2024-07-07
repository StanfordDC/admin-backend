package types

import (
	"cloud.google.com/go/firestore"
)

type User struct{
	Id string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email string `json:"email"`
}

type UserStore interface{
	GetAllUsers() *firestore.DocumentIterator
	CreateUser(user User) error
	UpdateUser(user User) error
	CheckIfUserExists(email string) bool
}

type WastetypeResponse struct{
	ImageUrl string `json:"imageUrl"`
	Objects map[string]int `json:"objects"`
}

type WastetypeResponseStore interface{
	GetAll() *firestore.DocumentIterator
}

type WasteTypeStore interface{
	Create(w WasteType) error
	GetAll() *firestore.DocumentIterator
	GetAllByItem(item string) *firestore.DocumentSnapshot
	DeleteItemByName(item string) (bool, error)
	Update(w WasteType) error
}

type WasteType struct{
	Id string `json:"id"`
	Instructions string `json:"instructions"`
	Item string `json:"item"`
	Links []string `json:"links"`
	Material string `json:"material"`
	Recyclable bool `json:"recyclable"`
}