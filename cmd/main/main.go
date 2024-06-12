package main

import (
	"admin-backend/cmd/api"
	"fmt"
	"context"
	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

func main(){
	db, db_err := NewFirestore()
	if db_err != nil{
		fmt.Println(db_err.Error())
	}

	server := api.NewAPIServer(":8080", db)
	err := server.Run()
	if err != nil{
		fmt.Println(err.Error())
	}
}

func NewFirestore() (*firestore.Client, error){
	opt := option.WithCredentialsFile("serviceAccountKey.json")
	app, _ := firebase.NewApp(context.Background(), nil, opt)
	client, err := app.Firestore(context.Background())
	return client, err
}