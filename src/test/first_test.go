//!/usr/local/go/bin/go
package main

import (
	"os"
	"fmt"
)

func execute( args int, argv []string ) int {
	fmt.Println( "Hello, World!" )
	return 0
}


// Entry Point
func main() {
	os.Exit( execute( len( os.Args ), os.Args ) )
}


// End_Of_Script