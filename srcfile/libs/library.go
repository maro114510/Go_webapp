//!/usr/local/go/bin/go
package point

import (
	"fmt"
	"os"
)

func execute( args int, argv []string ) int {
	fmt.Println( "Execute!" )
	// greeting.Do()
	return 0
}

func Point() {
	fmt.Println( "package" )
}


// Entry Point

func main() {
	os.Exit( execute( len( os.Args ), os.Args ) )
}



// End_Of_Script