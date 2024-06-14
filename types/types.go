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
	Instructions string `json:"instructions"`
	Item string `json:"item"`
	Link string `json:"link"`
	Material string `json:"material"`
	Recyclable bool `json:"recyclable"`
}