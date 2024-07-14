package types

import (
	"time"

	"cloud.google.com/go/firestore"
)

type User struct{
	Id string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserStore interface{
	GetAllUsers() *firestore.DocumentIterator
	CreateUser(user User) error
	UpdateUser(user User) error
	CheckIfUserExists(username string) bool
	DeleteUserByUsername(username string) (bool, error)
}

type WasteTypeFeedback struct{
	Item string `json:"item"`
	Feedback int `json:"feedback"`
	Source string `json:"source"`
}

type WastetypeResponse struct{
	ImageUrl string `json:"imageUrl"`
	Items []WasteTypeFeedback `json:"items"`
	CreateTime time.Time `json:"createTime"`
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