//!/usr/local/go/bin/go
package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/maro114510/Go_webapp/clock"
	"github.com/maro114510/Go_webapp/config"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmiron/sqlx"
)


func New( ctx context.Context, cfg *config.Config ) ( *sqlx.DB, func(), error ) {
	db, err := sql.Open( "mysql",
		fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?parseTime=true",
			cfg.DBUser,
			cfg.DBPassword,
			cfg.DBHost,
			cfg.DBPort,
			cfg.DBName,
		),
	)
	if err != nil {
		return nil, nil, err
	}

	ctx, cancel := context.WithTimeout( ctx, 2*time.Second )
	defer cancel()

	if err := db.PingContext( ctx ); err != nil {
		return nil, func() { _ = db.Close() }, err
	}

	xdb := sqlx.NewDb( db, "mysql" )
	return xdb, func() { _ = db.Close() }, nil
} /* New */

type Beginner interface {
	BeginTx( ctx context.Context, query string ) ( *sqlx.Stmt, error )
} /* Beginner */

type Preparer interface {
	Preparex( ctx context.Context, query string ) ( *sqlx.Stmt, error )
} /* Preparer */

type Execer interface {
	ExecContext( ctx context.Context, query string, args ...any ) ( sql.Result, error )
	NameExecContext( ctx context.Context, query string, arg interface{} ) ( sql.Result, error )
} /* Execer */

type Queryer interface {
	Preparer
	QueryContext( ctx context.Context, query string, args ..any ) ( *sqlx.Rows, error )
	QueryRowContext( ctx context.Context, query string, args ...any ) *sqlx.Row
	GetContent( ctx context.Context, dest interface{}, query string, args ...any ) error
	SelectContent( ctx context.Context, dest interface{}, query string, args ...any ) error
} /* Queryer */

var (
	_ Beginner = ( *sqlx.DB )( nil )
	_ Preparer = ( *sqlx.DB )( nil )
	_ Queryer = ( *sqlx.DB )( nil )
	_ Execer = ( *sqlx.DB )( nil )
	_ Execer = ( *sqlx.Tx )( nil )
)

type Repository struct {
	Clocker clock.Clocker
} /* Repository */



// End_Of_Script