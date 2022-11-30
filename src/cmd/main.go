//!/usr/local/go/bin/go
package main

import (
	"os"
	"fmt"

	"net/http"
)

func execute( args int, argv []string ) int {
	err := http.ListenAndServe(
		":18080",
		http.HandlerFunc( func( w http.ResponseWriter, r *http.Request ) {
			fmt.Fprintf( w, "Hello, World!" )
		} ),
	)

	if err != nil {
		fmt.Printf( "faild to terminate server: %v\n", err )
		os.Exit( 1 )
	}

	return 0
}


// Entry Point

func main() {
	os.Exit( execute( len( os.Args ), os.Args ) )
}



// End_Of_Script