package user

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

func (s *Store) CreateUser(user types.User) error{
	ref := s.db.Collection("wasteType").NewDoc()
	_, err := ref.Set(context.Background(), map[string]interface{}{
		"id" : ref.ID,
		"username" : user.Username,
		"password" : user.Password,
		"email" : user.Email,
	})
	return err
}

func (s *Store) GetAllUsers() *firestore.DocumentIterator{
	wasteCollection := s.db.Collection("user")
	iter := wasteCollection.Documents(context.Background())
	return iter
}
