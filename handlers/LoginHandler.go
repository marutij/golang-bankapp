package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"

	"github.com/marutijogdand17/golang-bankapp/database"
	"github.com/marutijogdand17/golang-bankapp/models"
	"go.mongodb.org/mongo-driver/bson"
)

var jwtKey = []byte("secret-jwt-key")
var ctx = context.Background()

type Claims struct {
	Id string `json:"id"`
	jwt.StandardClaims
}

func Login(w http.ResponseWriter, r *http.Request) {
	var user models.Login
	json.NewDecoder(r.Body).Decode(&user)
	accCollection := database.GetCollectionInstance("Accounts")
	result, _ := accCollection.Find(ctx, bson.M{})
	var account []models.Account
	result.All(ctx, &account)
	expireTime := time.Now().Local().Add(15 * time.Minute)
	for _, val := range account {
		if val.Id == user.AccountId && val.Password == user.Password {

			claims := &Claims{
				Id: val.Id,
				StandardClaims: jwt.StandardClaims{
					ExpiresAt: expireTime.Unix(),
				},
			}

			token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

			tokenString, err := token.SignedString(jwtKey)

			if err != nil {
				json.NewEncoder(w).Encode(&models.Message{Data: "Internal Server Error"})
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			http.SetCookie(w, &http.Cookie{Name: "token", Value: tokenString, Path: "/", Expires: expireTime})

			json.NewEncoder(w).Encode(&models.Message{Data: "Logged In successfully"})
			w.WriteHeader(http.StatusOK)
			return
		}

	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(&models.Message{Data: "Invalid Credentials"})
}
