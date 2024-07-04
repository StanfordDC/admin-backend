package types

import (
	"cloud.google.com/go/firestore"
)

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