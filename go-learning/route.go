package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
)

type Post struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

func getPosts(db *sql.DB) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		rows, err := db.Query("SELECT id, title, description, completed FROM posts")
		if err != nil {
			http.Error(resp, err.Error(), http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var posts []Post
		for rows.Next() {
			var post Post
			if err := rows.Scan(&post.ID, &post.Title, &post.Description, &post.Completed); err != nil {
				http.Error(resp, err.Error(), http.StatusInternalServerError)
				return
			}
			posts = append(posts, post)
		}
		if err := rows.Err(); err != nil {
			http.Error(resp, err.Error(), http.StatusInternalServerError)
			return
		}

		resp.Header().Set("Content-Type", "application/json")
		resp.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(resp).Encode(posts); err != nil {
			http.Error(resp, err.Error(), http.StatusInternalServerError)
		}
	}
}

func addPost(db *sql.DB) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		var post Post
		err := json.NewDecoder(req.Body).Decode(&post)
		if err != nil {
			http.Error(resp, err.Error(), http.StatusBadRequest)
			return
		}

		err = db.QueryRow(
			"INSERT INTO posts (title, description, completed) VALUES ($1, $2, $3) RETURNING id",
			post.Title, post.Description, post.Completed).Scan(&post.ID)
		if err != nil {
			http.Error(resp, err.Error(), http.StatusInternalServerError)
			return
		}

		resp.Header().Set("Content-Type", "application/json")
		resp.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(resp).Encode(post); err != nil {
			http.Error(resp, err.Error(), http.StatusInternalServerError)
		}
	}
}
