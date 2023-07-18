package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Book represents a book entity.
type Book struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"Isbn"`
	Title    string    `json:"title"`
	Author   *Author   `json:"author"`
}

// Author represents the author of a book.
type Author struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var books []Book // Slice to store books

// getBooks returns all the books.
func getBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(books)
}

// deleteBook deletes a book by ID.
func deleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	// Find the book with the given ID and remove it from the slice
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}

// getBook returns a book by ID.
func getBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	// Find the book with the given ID and encode it as JSON
	for _, item := range books {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

// createBook creates a new book.
func createBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var book Book

	// Decode the JSON request body into a Book struct
	_ = json.NewDecoder(r.Body).Decode(&book)
	book.ID = strconv.Itoa(rand.Intn(100000000))
	books = append(books, book)
	json.NewEncoder(w).Encode(book)
}

// updateBook updates a book by ID.
func updateBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("content-type", "application/json")
	params := mux.Vars(r)

	// Find the book with the given ID and update it with the new data
	for index, item := range books {
		if item.ID == params["id"] {
			books = append(books[:index], books[index+1:]...)

			// Add a new book with the updated data
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.ID = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
}

func main() {
	r := mux.NewRouter()

	// Adding sample books to the slice
	books = append(books, Book{ID: "1", Isbn: "438227", Title: "Book One", Author: &Author{Firstname: "John", Lastname: "Doe"}})
	books = append(books, Book{ID: "2", Isbn: "45455", Title: "Book Two", Author: &Author{Firstname: "Steve", Lastname: "Smith"}})

	// Define the routes and their corresponding handlers
	r.HandleFunc("/books", getBooks).Methods("GET")
	r.HandleFunc("/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/books", createBook).Methods("POST")
	r.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/books/{id}", deleteBook).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
}
