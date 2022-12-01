//!/usr/local/go/bin/go
package main

import (
	"context"
	"fmt"
	"io"
	"net"
	"net/http"
	"testing"

	"golang.org/x/sync/errgroup"
	"srcfile/cmd"
)

func TestRun( t *testing.T ) {
	l, err := net.Listen( "tcp", "localhost:0" )
	if err != nil {
		t.Fatalf( "faild to listen port %v\n", err )
	}

	ctx, cancel := context.WithCancel( context.Background() )
	eg, ctx := errgroup.WithContext( ctx )

	eg.Go( func() error {
		return cmd.Run( ctx, l )
	} )

	in := "message1"
	url := fmt.Sprintf( "http://%s/%s", l.Addr().String(), in )
	t.Logf( "try request to %q\n", url )
	rsp, err := http.Get( url )
	if err != nil {
		t.Errorf( "faild to get: %+v\n", err )
	}
	defer rsp.Body.Close()

	got, err := io.ReadAll( rsp.Body )
	if err != nil {
		t.Fatalf( "faild to read body: %+v\n", err )
	}

	wont := fmt.Sprintf( "Hello, %s!", in )
	if string( got ) != wont {
		t.Errorf( "wont %q, but got %q\n", wont, got )
	}

	cancel()
	if err := eg.Wait(); err != nil {
		t.Fatal( err )
	}
}


// End_Of_Script