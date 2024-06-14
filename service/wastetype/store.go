package wastetype

import (
	"admin-backend/types"
	"context"
	"cloud.google.com/go/firestore"
)

type Store struct {
	db *firestore.Client
}

func NewStore(db *firestore.Client) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) Create(wastetype types.WasteType) error{
	wasteCollection := s.db.Collection("wasteType")
	_,_, err := wasteCollection.Add(context.Background(), map[string]interface{}{
		"instructions":wastetype.Instructions,
		"item":wastetype.Item,
		"link":wastetype.Link,
		"material":wastetype.Material,
		"recyclable":wastetype.Recyclable,
	})
	if err != nil{
		return err
	}
	return nil
}

func (s* Store) GetAll() *firestore.DocumentIterator{
	wasteCollection := s.db.Collection("wasteType")
	iter := wasteCollection.Documents(context.Background())
	return iter
}

func(s* Store) GetAllByItem(item string) *firestore.DocumentIterator{
	wasteCollection := s.db.Collection("wasteType").Where("item", "==", item)
	iter := wasteCollection.Documents(context.Background())
	return iter
}
