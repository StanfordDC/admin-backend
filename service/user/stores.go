package user

import (
	"admin-backend/types"
	"context"
	"strings"
	"cloud.google.com/go/firestore"
	"google.golang.org/api/iterator"
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
	})
	return err
}

func (s* Store) GetUserByUsername(username string) *firestore.DocumentSnapshot{
	users := s.db.Collection("user")
	iter := users.Documents(context.Background())
	for{
		doc, err := iter.Next()
		if err == iterator.Done{
			break
		}
		//assert item to be string
		target := doc.Data()["username"].(string)
		//Check if email from db and input are the same regardless of case
		if strings.EqualFold(target, username){
			return doc
		}
	}
	return nil
}

func (s* Store) UpdateUser(user types.User) error{
	ref := s.db.Collection("user").Doc(user.Id)
	_, err := ref.Set(context.Background(), map[string]interface{}{
		"id":user.Id,
		"username" : user.Username,
		"password" : user.Password,
	}, firestore.MergeAll)
	return err
}

func (s* Store) DeleteUserByUsername(username string) (bool, error){
	users := s.db.Collection("user")
	iter := users.Documents(context.Background())
	for{
		doc, err := iter.Next()
		if err != nil{
			break
		}
		//assert item to be string
		target := doc.Data()["username"].(string)
		//Check if item from db and input are the same regardless of case
		if strings.EqualFold(target, username){
			_, err := doc.Ref.Delete(context.Background())
			return true, err
		}
	}
	return false, nil
}