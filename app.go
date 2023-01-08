package main

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"
	"log"
	"os"
	"reflect"

	"tt/tutorial"

	_ "embed"

	sqlite "modernc.org/sqlite"
)

//go:embed schema.sql
var ddl string

func run() error {
	ctx := context.Background()

	// check if the database file exists?
	_, err := os.Stat(`./db.db`)
	isDBFotFound := false
	if err != nil {
		isDBFotFound = true
	}

	// register a new sqlite function! easily: (e.g SELECT new_name();)
	sqlite.MustRegisterScalarFunction("new_name", 0, func(ctx2 *sqlite.FunctionContext, args []driver.Value) (driver.Value, error) {
		return "zaki!", nil
	})

	// db conenction!
	db, err := sql.Open("sqlite", "./db.db")
	if err != nil {
		return err
	}

	// create tables
	if isDBFotFound {
		if _, err := db.ExecContext(ctx, ddl); err != nil {
			return err
		}
	}

	// set up the queries
	queries := tutorial.New(db)

	// list all posts
	posts, err := queries.ListPosts(ctx)
	if err != nil {
		return err
	}
	for _, post := range posts {
		fmt.Println(post)
	}

	// create an post
	insertedPost, err := queries.CreatePost(ctx, tutorial.CreatePostParams{
		Title: fmt.Sprintf("Title [%d]", len(posts)+1),
		Text:  fmt.Sprintf("This is the post text [%d]", len(posts)+1),
	})
	if err != nil {
		return err
	}
	fmt.Println(insertedPost)

	// get the post we just inserted
	fetchedPost, err := queries.GetPost(ctx, insertedPost.ID)
	if err != nil {
		return err
	}

	// prints true
	fmt.Println(reflect.DeepEqual(insertedPost, fetchedPost))

	// get name from new_name() function 
	fetchedNewName, err := queries.GetNewName(ctx)
	if err != nil {
		return err
	}
	fmt.Println(fetchedNewName)

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
