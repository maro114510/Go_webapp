//!/usr/local/go/bin/go
package handler

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"fmt"
)

type ErrResponse struct {
	Message		string		`json:"message"`
	Details		[]string	`json:"details, omitempty"`
} /* ErrResponse */


func RespondJSON(
	ctx context.Context,
	w http.ResponseWriter,
	body any,
	status int ) {
	w.Header().Set( "Content-Type", "application/json; charset=utf-8" )
	bodyBytes, err := json.Marshal( body )
	if err != nil {
		w.WriteHeader( http.StatusInternalServerError )
		rsp := ErrResponse{
			Message: http.StatusText( http.StatusInternalServerError ),
		}

		if err := json.NewEncoder( w ).Encode( rsp ); err != nil {
			fmt.Printf( "write error response error: %v", err )
		}
		return
	}

	w.WriteHeader( status )
	if _, err := fmt.Fprintf( w, "%s", bodyBytes ); err != nil {
		fmt.Printf( "write response error: %v", err )
	}
} /* RespondJSON */


// End_Of_Script