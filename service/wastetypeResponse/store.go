package wastetyperesponse

import (
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

func (s *Store) GetAll() *firestore.DocumentIterator{
	responseCollection := s.db.Collection("wastetypeResponse")
	iter := responseCollection.Documents(context.Background())
	return iter
}
