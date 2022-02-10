package main

import (
	"encoding/json"
	"net/http"
	"log"
	"context"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/primitive"
)

collection := helper.ConnectDB()
func main() {

	r := mux.NewRouter()

	r.HandleFunc("/api/sneakers", getSneakers).Methods("GET")
	r.HandleFunc("/api/sneakers/{id}", getSneaker).Methods("GET")
	r.HandleFunc("/api/sneakers", createSneaker).Methods("POST")
	r.HandleFunc("/api/sneakers/{id}", updateSneaker).Methods("PUT")
	r.HandleFunc("/api/sneakers/{id}", deleteSneaker).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", r))

}

func getSneakers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var sneakers []models.Sneaker

	cur, err := collection.Find(context.TODO(), bson.M{})

	if err != nil {
		helper.errorHandler(err, w)
		return
	}

	defer cur.Close(context.TODO())

	for cur.Next(context.TODO()) {

		var sneaker models.Sneaker

		err := cur.Decode(&sneaker)
		if err != nil {
			log.Fatal(err)
		}

		sneakers = append(sneakers, sneaker)

	}

	ir err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	json.NewEncoder(w).Encode(sneakers)

}

func getSneaker(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("content-type", "application/json")

	var sneaker models.Sneaker

	var params = mux.Vars(r)

	id, _ := primitive.ObjectIDFromHex(params["id"])

	filter := bson.M{"_id": id}
	err := collection.FindOne(context.TODO(), filter).Decode(&sneaker)

	if err != nil {
		helper.errorHandler(err, w)
		return
	}

	json.NewEncoder(w).Encode(sneaker)

}

func createSneaker(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var sneaker models.Sneaker

	_= json.NewDecoder(r.Body).Decode(&sneaker)

	result, err := collection.InsertOne(context.TODO(), sneaker)

	if err != nil {
		helper.errorHandler(err, w)
		return
	}

	json.NewEncoder(w).Encode(result)

}

func updateSneaker(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")

	var params = mux.Vars(r)

	id, _ := primitive.ObjectIDFromHex(params["id"])

	var sneaker models.Sneaker

	filter := bson.M{"_id": id}

	_ = json.NewDecoder(r.Body).Decode(&book)

	update := bson.D{
		{"$set", bson.D{
			{"brand", sneaker.Brand},
			{"model", sneaker.Model},
			{"color", sneaker.Color},
			{"year", sneaker.Year},
			{"price", sneaker.Price},
		}},
	}

	err := collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&sneaker)

	if err != nil {
		helper.errorHandler(err, w)
		return
	}

	sneaker.ID = id

	json.NewEncoder(w).Encode(sneaker)
}

func deleteSneaker(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	v params = mux.Vars(r)


	id, er := primitive.ObjectIDFromHex(params["id"])

	filter := bson.M{"_id": id}

	deleteResult, err := collection.DeleteOne(context.TODO(), filter)

	if err != nil {
		helper.errorHandler(err, w)
		return
	}

	json.NewEncoder(w).Encode(deleteResult)

}
