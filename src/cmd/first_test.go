//!/usr/local/go/bin/go
package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"testing"

	"golang.org/x/sync/errgroup"
	"github.com/maro114510/Go_webapp/src"
)

func TestRun( t *testing.T ) {
	ctx, cancel := context.WithCancel( context.Background() )
	eg, ctx := errgroup.WithContext( ctx )

	eg.Go( func() error {
		return run( ctx )
	} )

	in := "message"
	rsp, err := http.Get( "http://localhost:18080" + in )
	if err != nil {
		t.Errorf( "faild to get: %+v\n", err )
	}
	defer rsp.Body.Close()

	got, err := io.ReadAll( rsp.Body )
	if err != nil {
		t.Faitalf( "faild to read body: %+v\n", err )
	}

	wont := fmt.Sprintf( "Hello, %s!", in )
	if string( got ) != wont {
		t.Errorf( "wont %q, but got %q\n", wont, got )
	}

	cancel()
	if err := eg.Wait(); err != nil {
		t.Faital( err )
	}
}


// End_Of_Script