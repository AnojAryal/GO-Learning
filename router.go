package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"

	_ "go-learning/docs"
)

// Post represents a blog post
type Post struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

// getPosts godoc
// @Summary Get all posts
// @Description Get all posts
// @Tags posts
// @Produce json
// @Success 200 {array} Post
// @Router /posts [get]
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

// addPost godoc
// @Summary Add a new post
// @Description Add a new post
// @Tags posts
// @Accept json
// @Produce json
// @Param post body Post true "Add post"
// @Success 201 {object} Post
// @Router /posts [post]
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


// deletePost godoc
// @Summary Delete a post by ID
// @Description Delete a post by ID
// @Tags posts
// @Produce json
// @Param id path int true "Post ID"
// @Success 204 "No Content"
// @Router /posts/{id} [delete]
func deletePost(db *sql.DB) http.HandlerFunc {
	return func(resp http.ResponseWriter, req *http.Request) {
		// Get the ID from the URL parameters
		vars := mux.Vars(req)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			http.Error(resp, "Invalid post ID", http.StatusBadRequest)
			return
		}

		// Execute the delete query
		result, err := db.Exec("DELETE FROM posts WHERE id = $1", id)
		if err != nil {
			http.Error(resp, err.Error(), http.StatusInternalServerError)
			return
		}

		// Check if any rows were affected
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			http.Error(resp, err.Error(), http.StatusInternalServerError)
			return
		}
		if rowsAffected == 0 {
			http.Error(resp, "Post not found", http.StatusNotFound)
			return
		}

		// Set the response header and status code
		resp.Header().Set("Content-Type", "application/json")
		resp.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(resp).Encode(vars); err != nil {
			http.Error(resp, err.Error(), http.StatusInternalServerError)
		}
	}
}
