package controllers

import (
	"context"
	"errors"

	"github.com/marutijogdand17/golang-bankapp/database"
	"github.com/marutijogdand17/golang-bankapp/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var ctx = context.Background()

type Db struct {
	Col *mongo.Collection
}

func NewDbCollection(collection string) *Db {
	return &Db{
		Col: database.GetCollectionInstance(collection),
	}
}

func (db *Db) CreateAccount(account models.Account) error {
	_, err := db.Col.InsertOne(ctx, account)
	if err != nil {
		errors.New(err.Error())
	}

	return nil
}

func (db *Db) GetAccount(id string) (models.Account, error) {
	var account models.Account
	db.Col.FindOne(ctx, bson.D{{"_id", id}}).Decode(&account)
	if account.Id == "" {
		return account, errors.New("User not found")
	}
	return account, nil
}

func (db *Db) UpdateUserDetails(account models.Account) error {

	update := bson.M{
		"$set": bson.M{
			"_id":            account.Id,
			"first_name":     account.FirstName,
			"last_name":      account.LastName,
			"email_id":       account.EmailId,
			"mobile_number":  account.MobileNumber,
			"address_line_1": account.AddressLine1,
			"city":           account.City,
			"district":       account.District,
			"state":          account.State,
			"country":        account.Country,
			"pincode":        account.Pincode,
		}}

	db.Col.FindOneAndUpdate(ctx, bson.M{"_id": account.Id}, update)
	if account.Id == "" {
		return errors.New("User not found")
	}
	return nil
}

func (db *Db) DeleteAccount(id string) error {

	_, err := db.Col.DeleteOne(ctx, bson.D{{"_id", id}})
	if err != nil {
		return errors.New("User not found")
	}
	return nil
}
