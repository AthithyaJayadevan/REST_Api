package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//BOOK STRUCT
type book struct {
	Title  string  `json:"title`
	Id     string  `json:"id"`
	Isbn   string  `json:"isbn`
	Author *Author `json:"author"`
}

//Author struct

type Author struct {
	F_name string `json:"first_name"`
	L_name string `json:"last_name"`
}

var books []book

//get all books
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(books)

}

//Get Single book
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	//Loop through books and find corresponding books
	for _, item := range books {
		if item.Id == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&book{})

}

func updateBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.Id == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var new_book book
			_ = json.NewDecoder(r.Body).Decode(&new_book)
			new_book.Id = params["id"]
			books = append(books, new_book)
			json.NewEncoder(w).Encode(new_book)
		}

	}
	json.NewEncoder(w).Encode(books)

}

func createBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var new_book book
	_ = json.NewDecoder(r.Body).Decode(&new_book)
	new_book.Id = strconv.Itoa(rand.Intn(1000000))
	books = append(books, new_book)
	json.NewEncoder(w).Encode(new_book)
}

func deleteBooks(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range books {
		if item.Id == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}

	}
	json.NewEncoder(w).Encode(books)
}

func main() {

	//Router variable
	router := mux.NewRouter()

	// MOCK DATA
	books = append(books, book{Id: "2984759", Isbn: "dcjwyhgb", Title: "gweiyg", Author: &Author{F_name: "John", L_name: "Doe"}})
	books = append(books, book{Id: "08248", Isbn: "kiuvh", Title: "wjdvn", Author: &Author{F_name: "Johnathan", L_name: "Does"}})
	books = append(books, book{Id: "098059", Isbn: "ikveiou", Title: "whjdbv", Author: &Author{F_name: "Johnny", L_name: "Doess"}})

	//Router Handlers /Endpoints
	router.HandleFunc("/api/books", getBooks).Methods("GET")
	router.HandleFunc("/api/books/{id}", getBook).Methods("GET")
	router.HandleFunc("/api/books", createBooks).Methods("POST")
	router.HandleFunc("/api/books/{id}", updateBooks).Methods("PUT")
	router.HandleFunc("/api/books/{id}", deleteBooks).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", router))

}
