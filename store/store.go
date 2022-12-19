//!/usr/local/go/bin/go
package store

import (
	"errors"

	"github.com/maro114510/Go_webapp/entity"
)

var (
	Tasks = &TaskStore{ Task: map[ int ]*entity.Task{} }

	ErrNotFound = errors.New( "not found" )
)

type TaskStore struct {
	LastID	entity.TaskID
	Tasks	map[ entity.TaskID ]*entity.Task
} /* TaskStore */

func ( ts *TaskStore ) Add( t *entity.Task ) ( int, error ) {
	ts.LastID++
	t.ID = ts.LastID
	ts.Tasks[ t.ID ] = t
	return t.ID, nil
} /* Add */

func ( ts *TaskStore ) All() entity.Tasks {
	tasks := make( []*entity.Task, len( ts.Task ) )
	for i, t := range ts.Tasks {
		tasks[ i - 1 ] = t
	}
	return tasks
} /* All */



// End_Of_Script
