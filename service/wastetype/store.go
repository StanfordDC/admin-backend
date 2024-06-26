package wastetype

import (
	"admin-backend/types"
	"context"
	"strings"

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
	ref := s.db.Collection("wasteType").NewDoc()
	_, err := ref.Set(context.Background(), map[string]interface{}{
		"id" : ref.ID,
		"instructions":wastetype.Instructions,
		"item":wastetype.Item,
		"link":wastetype.Link,
		"material":wastetype.Material,
		"recyclable":wastetype.Recyclable,
	})
	return err
}

func (s *Store) GetAll() *firestore.DocumentIterator{
	wasteCollection := s.db.Collection("wasteType")
	iter := wasteCollection.Documents(context.Background())
	return iter
}

func(s *Store) GetAllByItem(item string) *firestore.DocumentSnapshot{
	iter := s.GetAll()
	for{
		doc, err := iter.Next()
		if err != nil{
			break
		}
		//assert item to be string
		dbItem := doc.Data()["item"].(string)
		//Check if item from db and input are the same regardless of case
		if strings.EqualFold(dbItem, item){
			return doc
		}
	}
	return nil
}

func(s *Store) DeleteItemByName(item string) (bool, error){
	wasteCollection := s.db.Collection("wasteType")
	iter := wasteCollection.Documents(context.Background())
	for{
		doc, err := iter.Next()
		if err != nil{
			break
		}
		//assert item to be string
		dbItem := doc.Data()["item"].(string)
		//Check if item from db and input are the same regardless of case
		if strings.EqualFold(dbItem, item){
			_, err := doc.Ref.Delete(context.Background())
			return true, err
		}
	}
	return false, nil
}

func(s *Store) Update(wastetype types.WasteType) error{
	ref := s.db.Collection("wasteType").Doc(wastetype.Id)
	_, err := ref.Set(context.Background(), map[string]interface{}{
		"id":wastetype.Id,
		"instructions":wastetype.Instructions ,
		"item":wastetype.Item,
		"link":wastetype.Link,
		"material":wastetype.Material,
		"recyclable":wastetype.Recyclable,
	}, firestore.MergeAll)
	return err
}
