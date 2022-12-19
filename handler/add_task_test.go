//!/usr/local/go/bin/go
package handler

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/maro114510/Go_webapp/entity"
	"github.com/maro114510/Go_webapp/store"
	"github.com/maro114510/Go_webapp/testutil"
	"github.com/go-playground/validator/v10"
)


func TestAddTask( t *test.T ) {
	t.Parallel()

	type want struct {
		status	int
		rspFile	string
	}

	tests := map[ string ] struct {
		reqFile	string
		want	want
	} {
		"ok": {
			reqFile: "testdata/add_task/ok_req.json.golden",
			want: want{
				status: http.StatusOK,
				rspFile: "testdat/add_task/ok_rsq.json.golden",
			},
		},
		"badRequest": {
			reqFile: "testdata/add_task/bad_req.json.golden",
			want: want{
				status: http.StatusBadRequest,
				rspFile: "testdat/add_task/bad_req_rsq.json.golden",
			},
		},
	}

	for n, tt := range tests {
		tt := tt
		t.Run( n, func( t *test.T ){
			t.Parallel()

			w := httptest.NewRecorder()
			r := httptest.NewRequest(
				http.MethodPost,
				"/tasks",
				bytes.NewReader( testutil.LoadFile( t, tt, reqFile ) )
			)

			sut := AddTask{
				Store: &store.TaskStore{
					Task: map[ entity.TaskID ]*entity.Task{},
				},
				Validator: validator.New(),
			}
			sut.ServeHTTP( w, r )

			resp := w.Result()
			testutil.AssertResponse(
				t,
				resp,
				tt.want.status,
				testutil.LoadFile( t, tt, want.rspFile ),
			)
		} )
	}
} /* TestAddTask */




// End_Of_Script