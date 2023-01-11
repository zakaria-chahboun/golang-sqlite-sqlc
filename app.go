package main

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"
	"log"
	"os"
	"reflect"
	"tt/db"

	_ "embed"

	gonanoid "github.com/matoous/go-nanoid"

	sqlite "modernc.org/sqlite"
)

//go:generate sqlc generate

//go:embed models/schema/*.sql
var ddl string

const db_path = "./db/database.db"

func run() error {
	ctx := context.Background()

	// check if the database file exists?
	_, err := os.Stat(db_path)
	isDBFotFound := false
	if err != nil {
		isDBFotFound = true
	}

	// register a new sqlite function! easily: (e.g SELECT new_id('p');)
	sqlite.MustRegisterScalarFunction("new_id", 1, newID)

	// db connection!
	connection, err := sql.Open("sqlite", db_path)
	if err != nil {
		return err
	}
	defer connection.Close()

	// create tables
	if isDBFotFound {
		if _, err := connection.ExecContext(ctx, ddl); err != nil {
			return err
		}
	}
	// Make sure your foreign keys relationship is working fine
	if _, err := connection.ExecContext(ctx, "pragma foreign_keys=on;"); err != nil {
		return err
	}

	// set up the queries
	queries := db.New(connection)

	// create a post
	insertedPost, err := queries.CreatePost(ctx, db.CreatePostParams{
		Title: "Orange is sweet",
		Text:  "As you know, the orange is always good!",
	})
	if err != nil {
		return err
	}
	fmt.Println("New Post:", insertedPost)

	// get the post we just inserted
	fetchedPost, err := queries.GetPost(ctx, insertedPost.ID)
	if err != nil {
		return err
	}

	// prints true
	fmt.Println(reflect.DeepEqual(insertedPost, fetchedPost))

	// create comments for the post we just inserted
	for range [3]int{} {
		_, err = queries.CreateComment(ctx, db.CreateCommentParams{
			Text:   "Yest it is!",
			PostID: insertedPost.ID,
		})
		if err != nil {
			return err
		}
	}

	// list all posts with their comments
	fetchedPosts, err := queries.ListPosts(ctx)
	if err != nil {
		return err
	}
	for _, p := range fetchedPosts {
		fmt.Println()
		fmt.Println("Post:", p)
		comments, err := queries.ListCommentsInPost(ctx, p.ID)
		if err != nil {
			return err
		}
		for _, c := range comments {
			fmt.Println("Comment:", c)
		}
	}

	// delete the post we just inserted
	/*err = queries.DeletePost(ctx, insertedPost.ID)
	if err != nil {
		return err
	}*/

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func newID(ctx2 *sqlite.FunctionContext, args []driver.Value) (driver.Value, error) {
	const defaultRandomAlphabet = "abcdefghijklmnopqrstuvwxyz0123456789"
	id, _ := gonanoid.Generate(defaultRandomAlphabet, 7)
	return id, nil
}
