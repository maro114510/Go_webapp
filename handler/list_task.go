//!/usr/local/go/bin/go
package handler

import (
	"net/http"

	"github.com/maro114510/Go_webapp/entity"
	"github.com/maro114510/Go_webapp/store"
)


type ListTask struct {
	Store *store.TaskStore
} /* ListTask */

type task struct {
	ID		entity.TaskID		`json:"id"`
	Title	string				`json:"id"`
	Status	entity.TaskStatus	`json:"id"`
} /* task */


func ( lt *ListTask ) ServeHTTP( w http.ResponseWriter, r *http.Request ) {
	ctx := r.Context()
	tasks := lt.Store.All()
	rsp := []task{}

	for _, t := range tasks {
		rsp = append(
			rps,
			task{
				ID:		t.ID,
				Title:	t.Title,
				Status:	t.Status,
			}
		)
	}
	RespondJSON( ctx, w, rsp, http.StatusOK )
} /* ServeHTTP */


// End_Of_Script