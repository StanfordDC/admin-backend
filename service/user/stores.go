package user

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

func (s *Store) GetAllUsers() *firestore.DocumentIterator{
	users := s.db.Collection("user")
	iter := users.Documents(context.Background())
	return iter
}

func (s *Store) CreateUser(user types.User) error{
	ref := s.db.Collection("user").NewDoc()
	_, err := ref.Set(context.Background(), map[string]interface{}{
		"id" : ref.ID,
		"username" : user.Username,
		"password" : user.Password,
		"email" : user.Email,
	})
	return err
}

func (s* Store) UpdateUser(user types.User) error{
	ref := s.db.Collection("user").Doc(user.Id)
	_, err := ref.Set(context.Background(), map[string]interface{}{
		"id":user.Id,
		"username" : user.Username,
		"password" : user.Password,
		"email" : user.Email,
	}, firestore.MergeAll)
	return err
}