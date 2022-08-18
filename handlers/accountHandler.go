package handlers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/marutijogdand17/golang-bankapp/controllers"
	"github.com/marutijogdand17/golang-bankapp/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AuthenticationHandler(w http.ResponseWriter, r *http.Request) (string, error) {
	// Fetch token from headers
	token, err := r.Cookie("token")

	// Validate err
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return "", errors.New(err.Error())
	}

	// Build empty claims object
	claims := &Claims{}

	// Parse the token
	tkn, err := jwt.ParseWithClaims(token.Value, claims, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return "", errors.New(err.Error())
	}

	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return "", errors.New(err.Error())
	}

	return claims.Id, nil
}

func GetAccount(w http.ResponseWriter, r *http.Request) {
	id, err := AuthenticationHandler(w, r)

	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	db := controllers.NewDbCollection("Accounts")
	account, err := db.GetAccount(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(models.Message{Data: "user details with id " + id + " not found"})
		return
	}

	json.NewEncoder(w).Encode(&account)
	w.WriteHeader(http.StatusOK)
}

func CreateAccount(w http.ResponseWriter, r *http.Request) {

	var account models.Account
	json.NewDecoder(r.Body).Decode(&account)
	account.Id = primitive.NewObjectID().Hex()
	account.Password = primitive.NewObjectID().Hex()

	db := controllers.NewDbCollection("Accounts")
	if err := db.CreateAccount(account); err != nil {
		json.NewEncoder(w).Encode(&models.Message{Data: "Error while Creating account"})
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(&models.Message{Data: "Account Id: " + account.Id})
}

func UpdateAccount(w http.ResponseWriter, r *http.Request) {

	var id = mux.Vars(r)["id"]

	var account models.Account
	json.NewDecoder(r.Body).Decode(&account)
	account.Id = id

	log.Println(id, account)

	db := controllers.NewDbCollection("Accounts")
	if err := db.UpdateUserDetails(account); err != nil {
		json.NewEncoder(w).Encode(&models.Message{Data: "Error while Creating account"})
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(&models.Message{Data: "Account details updated"})
}

func DeleteAccount(w http.ResponseWriter, r *http.Request) {
	var id = mux.Vars(r)["id"]

	db := controllers.NewDbCollection("Accounts")
	if err := db.DeleteAccount(id); err != nil {
		json.NewEncoder(w).Encode(&models.Message{Data: "Error while Creating account"})
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	json.NewEncoder(w).Encode(&models.Message{Data: "Account Deleted"})
}
