package types

import (
	"cloud.google.com/go/firestore"
)


type WasteTypeStore interface{
	Create(w WasteType) error
	GetAll() *firestore.DocumentIterator
	GetAllByItem(item string) *firestore.DocumentIterator
}

type WasteType struct{
	CanBePlaced bool `json:"canBePlaced"`
	Description string `json:"description"`
	ItemName string `json:"itemName"`
}