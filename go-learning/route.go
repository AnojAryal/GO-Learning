package main

import (
	"encoding/json"
	"net/http"
)

type Post struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

var posts = []Post{
	{ID: 1, Title: "Eat Pizza", Description: "GO and Eat Pizza", Completed: false},
}

func getPosts(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(resp).Encode(posts); err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "error marshalling the posts array"}`))
		return
	}
}

func addPost(resp http.ResponseWriter, req *http.Request) {
	var post Post
	err := json.NewDecoder(req.Body).Decode(&post)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "error unmarshalling the request"}`))
		return
	}

	post.ID = len(posts) + 1
	posts = append(posts, post)

	resp.Header().Set("Content-Type", "application/json")
	resp.WriteHeader(http.StatusOK)

	result, err := json.Marshal(post)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "error marshalling the post object"}`))
		return
	}

	resp.Write(result)
}