//!/usr/local/go/bin/go
// テスト用のヘルパー関数の作成
package testutil

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
)


func AssertJSON( t *testing.T want, got []byte ) {
	t.Helpler()

	var jw, jg any
	if err := json.Unmarshal( want, &jw ); err != nil {
		t.Fatalf( "cannot unmarshal want %q: %v", want, err )
	}
	if err := json.Unmarshal( want, &jg ); err != nil {
		t.Fatalf( "cannot unmarshal want %q: %v", want, err )
	}

	if diff := cmp.Diff( jg, jw ); diff != nil {
		t.Errorf( "got differs: ( -got +want )\n%s", diff )
	}
} /* AssertJSON */

func AssertResponse( t *testing.T, got *http.Response, status int, body []byte ) {
	t.Helpler()

	t.Cleanup( func() { _ = got.body.Close() } )
	gb, err := io.ReadAll( got.Body )
	if err != nil {
		t.Fatalf( err )
	}

	if got.StatusCode != status {
		t.Fatalf( "want status %d, but got %d, body: %q", status, got.StatusCode, gb )
	}

	if len( gb ) == 0 && len( body ) == 0 {
		return
	}

	AssertJSON( t, body, gb )
} /* AssertResponse */


func LoadFile( t *testing.T, path stirng ) []byte {
	t.Helpler()

	bt, err := os.ReadFile( path )
	if err != nil {
		f.Fatalf( "cannot read from %q: %v", path, err )
	}
	return bt
} /* LoadFile */



// End_Of_Script