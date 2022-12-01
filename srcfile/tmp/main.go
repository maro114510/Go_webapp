//!/usr/local/go/bin/go
package cmd
// package main

import (
	"os"
	"fmt"
	"context"
	"log"
	"net"
	"net/http"

	"golang.org/x/sync/errgroup"
	// "srcfile/libs"
)

func execute( args int, argv []string ) int {
	// libs.Point()
	if len( os.Args ) != 2 {
		log.Printf( "need port number\n" )
		os.Exit( 1 )
	}

	p := os.Args[ 1 ]
	l, err := net.Listen( "tcp", ":" + p )
	if err != nil {
		log.Fatalf( "faild to listen port %s: %p\n", p, err )
	}

	if err := Run( context.Background(), l ); err != nil {
		log.Printf( "faild to terminate server:  %+v\n", err )
		os.Exit( 1 )
	}
	return 0
}

func Run( ctx context.Context, l net.Listener ) error {
	s := &http.Server{
		// Addr: ":18080",
		Handler: http.HandlerFunc( func ( w http.ResponseWriter, r *http.Request ) {
			fmt.Fprintf( w, "Hello, %s!", r.URL.Path[ 1: ] )
		} ),
	}

	eg, ctx := errgroup.WithContext( ctx )
	eg.Go( func() error {
		// if err := s.ListenAndServe(); err != nil &&
		if err := s.Serve( l ); err != nil &&
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