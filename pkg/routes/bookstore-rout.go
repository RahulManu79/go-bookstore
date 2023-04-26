package routes

import (
    // "net/http"
   "github.com/gorilla/mux"
   "github.com/RahulManu79/go-bookstore/pkg/controllers"
)

var RegisterBookStoreRoute = func (router *mux.Route)  {
	router.HandlerFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandlerFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandlerFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandlerFunc("/book/{bookId}", controllers.updateBook).Methods("PUT")
	router.HandlerFunc("/book/{bookId}", controllers.deleteBook).Methods("DELETE")

}