package helper

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

)

func connectDatabase() *mongo.Collection {

	clientOptions := options.Client().ApplyURI("")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connection to MongoDB successful.")

	collection := client.Database("").Collection("sneakers")

	return collection

}

type ErrorResponse struct {
	StatusCode int `json:"status"`
	ErrorMessage string `json:"message"`

}

func errorHandler(err error, w http.ResponseWriter){

	log.Fatal(err.Error())
	var response = ErrorResponse{
		ErrorMessage: err.Error(),
		StatusCode: http.StatusInternalServerError,
	}

	message, _ := json.Marshal(response)

	w.WriteHeader(response.StatusCode)
	w.Write(message)

}
