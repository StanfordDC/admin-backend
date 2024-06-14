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

func(s* Store) GetAllByItem(item string) *firestore.DocumentSnapshot{
	iter := s.GetAll()
	for{
		doc, err := iter.Next()
		if err != nil{
			break
		}
		if doc.Data()["item"] == item{
			return doc
		}
	}
	return nil
}

func(s* Store) DeleteItemByName(item string) error{
	wasteCollection := s.db.Collection("wasteType")
	iter := wasteCollection.Documents(context.Background())
	for{
		doc, _ := iter.Next()
		if doc.Data()["item"] == item{
			_, err := doc.Ref.Delete(context.Background())
			return err
		}
	}
}
