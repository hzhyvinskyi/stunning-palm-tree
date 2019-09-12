package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/handler"
	"github.com/hzhyvinskyi/stunning-palm-tree"
	"github.com/hzhyvinskyi/stunning-palm-tree/api/dal"
)

const defaultPort = "8080"

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	db, err := dal.Connect()
	if err != nil {
		panic(err)
	}

	http.Handle("/", handler.Playground("GraphQL playground", "/query"))
	http.Handle("/query", handler.GraphQL(stunning_palm_tree.NewExecutableSchema(stunning_palm_tree.Config{Resolvers: &stunning_palm_tree.Resolver{}})))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func initDB(db *sql.DB) {
	dal.MustExec(db, "DROP TABLE IF EXISTS reviews")
	dal.MustExec(db, "DROP TABLE IF EXISTS screenshots")
	dal.MustExec(db, "DROP TABLE IF EXISTS users")
	dal.MustExec(db, "DROP TABLE IF EXISTS videos")
	dal.MustExec(db, "CREATE TABLE public.users (id SERIAL PRIMARY KEY, name varchar(255), email varchar(255))")
	dal.MustExec(db, "CREATE TABLE public.videos (id SERIAL PRIMARY KEY, name varchar(255), description varchar(255), url text,created_at TIMESTAMP, user_id int, FOREIGN KEY (user_id) REFERENCES users (id))")
	dal.MustExec(db, "CREATE TABLE public.screenshots (id SERIAL PRIMARY KEY, video_id int, url text, FOREIGN KEY (video_id) REFERENCES videos (id))")
	dal.MustExec(db, "CREATE TABLE public.reviews (id SERIAL PRIMARY KEY, video_id int,user_id int, description varchar(255), rating varchar(255), created_at TIMESTAMP, FOREIGN KEY (user_id) REFERENCES users (id), FOREIGN KEY (video_id) REFERENCES videos (id))")
	dal.MustExec(db, "INSERT INTO users(name, email) VALUES('DFK', 'test@dir.me')")
	dal.MustExec(db, "INSERT INTO users(name, email) VALUES('Tushar', 'erl@gmail.com')")
	dal.MustExec(db, "INSERT INTO users(name, email) VALUES('Dipen', 'asm@mail.com')")
	dal.MustExec(db, "INSERT INTO users(name, email) VALUES('Harsh', 'crush1z@gmail.com')")
	dal.MustExec(db, "INSERT INTO users(name, email) VALUES('Priyank', 'zwma3o@gmail.com')")
}
