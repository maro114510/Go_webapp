//!/usr/local/go/bin/go
package main

import (
	"os"
	"fmt"
	"context"
	"log"
	"net/http"

	"golang.org/x/sync/errgroup"
)

func execute( args int, argv []string ) int {
	if err := Run( context.Background() ); err != nil {
		log.Printf( "faild to terminate server:  %+v\n", err )
	}
	return 0
}

func Run( ctx context.Context ) error {
	s := &http.Server{
		Addr: ":18080",
		Handler: http.HandlerFunc( func ( w http.ResponseWriter, r *http.Request ) {
			fmt.Fprintf( w, "hello, %s!", r.URL.Path[ 1: ] )
		} ),
	}

	eg, ctx := errgroup.WithContext( ctx )
	eg.Go( func() error {
		if err := s.ListenAndServe(); err != nil &&
		err != http.ErrServerClosed {
			log.Printf( "faild to close: %+v\n", err )
			return err
		}
		return nil
	} )

	<-ctx.Done()
	if err := s.Shutdown( context.Background() ); err != nil {
		log.Printf( "faild to shutdown: %+v\n", err )
	}

	return eg.Wait()
}


// Entry Point

func main() {
	os.Exit( execute( len( os.Args ), os.Args ) )
}



// End_Of_Script