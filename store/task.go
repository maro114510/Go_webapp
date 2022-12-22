//!/usr/local/go/bin/go
package store

import (
	"context"

	"github.com/maro114510/Go_webapp/entity"
)


func ( r *Repository ) AddTask(
	ctx context.Context, db Execer, t *entity.Task,
) error {
	t.Created = r.Clocker.Now()
	t.Modified = r.Clocker.Now()
	sql := `
	INSERT INTO task
		( title, status, created, modified )
		VALUES (
			?,
			?,
			?,
			?
		);
	`
	result, err := db.ExecContext(
		ctx, sql, t.Title, t.Status,
		t.Created, t.Modified,
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	t.ID = entity.TaskID( id )
	return nil
} /* AddTask */

func ( r *Repository ) ListTasks (
	ctx context.Context, db Queryer,
) ( entity.Tasks, error ) {
	tasks := entity.Tasks{}
	sql := `
	SELECT 
		id,
		title,
		status,
		created,
		modified
	FROM task;
	`
	if err := db.SelectContext( ctx, &tasks, sql ); err != nil {
		return nil, err
	}
	return tasks, nil
} /* ListTasks */



// End_Of_Script