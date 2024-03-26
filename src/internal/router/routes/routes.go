package routes

import (
	"CatRegistry/src/internal"
	"CatRegistry/src/internal/cat"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var (
	DBConnection *internal.MongoDBConnection
	uri          string
	port         string
)

func init() {
	uri = os.Getenv("MONGODB_URI")
	port = os.Getenv("MONGODB_PORT")
	if uri == "" {
		uri = "127.0.0.1"
	}
	if port == "" {
		port = "27017"
	}
	fullConnString := fmt.Sprintf("mongodb://%s:%s", uri, port)
	DBConnection = internal.CreateMongoDBConnection(fullConnString)
}

func RegisterRoutes(srv *http.ServeMux) {
	srv.HandleFunc("/getcats", GetAllCats)
	srv.HandleFunc("/filtercats", GetCatsByFilter)
	srv.HandleFunc("/insertcat", PostCat)
	srv.HandleFunc("/health", healthCheck)
	srv.HandleFunc("/dbhealth", dbCheck)
}

// GET ROUTES
func GetCatsByFilter(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(400)
		w.Write([]byte("Wrong method used. Use Get method."))
	}
	var filter cat.Cat

	err := json.NewDecoder(r.Body).Decode(&filter)
	if err != nil {
		fmt.Println("Error decoding json from body")
		fmt.Println(err)
		w.WriteHeader(400)
	}

	cat, err := DBConnection.GetCatsByFilter(filter)
	if err != nil {
		fmt.Println("Encountered error")
		fmt.Println(err)
		w.WriteHeader(400)
	}
	catsJson, err := json.Marshal(cat)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
	}
	w.Write([]byte(catsJson))
}

func GetAllCats(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(400)
		w.Write([]byte("Wrong method used. Use Get method."))
	}
	cats := DBConnection.GetAllCats()
	catsJson, err := json.Marshal(cats)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
	}
	w.WriteHeader(200)
	w.Write([]byte(catsJson))
}

// POST ROUTES

func PostCat(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(400)
		w.Write([]byte("Wrong method used. Use Post method."))
	}
	var input cat.Cat

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		fmt.Println("Error encountered during decoding body")
		fmt.Println(err)
		w.WriteHeader(400)
	}

	DBConnection.InsertCat(input)
	inputJson, err := json.Marshal(input)
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Error encountered"))
	}
	w.WriteHeader(201)
	w.Write([]byte(inputJson))
}

// HEALTH CHECKS

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte("Healthy"))
}

func dbCheck(w http.ResponseWriter, r *http.Request) {
	err := DBConnection.Connection.Ping(context.TODO(), readpref.Nearest())
	if err != nil {
		w.WriteHeader(500)
		w.Write([]byte("Error connection to database"))
	}
	w.WriteHeader(200)
	w.Write([]byte("Healthy"))
}
