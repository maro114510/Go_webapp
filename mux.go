//!/usr/local/go/bin/go
package main

import (
	"net/http"

	"github.com/maro114510/Go_webapp/handler"
	"github.com/maro114510/Go_webapp/store"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

func NewMux() http.Handler {
	mux := chi.NewRouter()
	mux.HandleFunc( "/health", func( w http.ResponseWriter, r *http.Request ) {
		w.Header().Set( "Content-Type", "application/json; charset=utf-8" )
		_, _ = w.Write( []byte( `{ "status": "OK" }` ) )
	} )

	v := validator.New()
	at := &handler.AddTask{ Store: store.Tasks, Validator: v }
	mux.Post( "/tasks", at.ServeHTTP )

	lt := &handler.ListTask{  Store: store.Tasks }
	mux.Get( "/tasks", lt.ServeHTTP )

	return mux
} /* NewMux */


// End_Of_Script