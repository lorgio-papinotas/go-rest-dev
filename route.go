package main

import (
	"encoding/json"
	"net/http"
)

type Post struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Text  string `json:"text"`
}

var (
	posts []Post
)

func init() {
	posts = []Post{Post{Id: 1, Title: "Title 1", Text: "Text 1"}}
}

func getPosts(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Context-type", "application/json")
	result, err := json.Marshal(posts)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"error": "Error marshalling the posts array"}`))
		return
	}
	res.WriteHeader(http.StatusOK)
	res.Write(result)
}

func addPost(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Context-type", "application/json")
	var post Post
	err := json.NewDecoder(req.Body).Decode(&post)
	if err != nil {
		res.WriteHeader(http.StatusInternalServerError)
		res.Write([]byte(`{"error": "Error unmarshalling the request"}`))
		return
	}
	post.Id = len(posts) + 1
	posts = append(posts, post)
	res.WriteHeader(http.StatusOK)
	result, err := json.Marshal(post)
	res.Write(result)
}
