//!/usr/local/go/bin/go
package main

import (
	"net/http"
)

func NewMux() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc( "/health", func( w http.ResponseWriter, r *http.Request ) {
		w.Header().Set( "Content-Type", "application/json; charset=utf-8" )
		_, _ = w.Write( []byte( `{ "status": "OK" }` ) )
	} )

	return mux
} /* NewMux */


// End_Of_Script