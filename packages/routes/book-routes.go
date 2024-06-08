package routes

import (
	"github.com/gorilla/mux"
	"github.com/Darshan016/gp-booksotre/packages/controllers"
)

var registerBookStore = func(router *mux.Router){
	router.HandleFunc("/book/",controllers.createBook).Methods("POST")
	router.HandleFunc("/book/",controllers.getBooks).Methods("GET")
	router.HandleFunc("/book/{bookId}",controllers.getBook).Methods("GET")
	router.HandleFunc("/book/{bookId}",controllers.updateBook).Methods("PUT")
	router.HandleFunc("/book/{bookId}",controllers.deleteBook).Methods("DELETE")
}