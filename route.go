package main

import (
	"encoding/json"
	"log"
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
	posts = []Post{{Id: 1, Title: "Title One", Text: "Text One"}}
}

func getPosts(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	result, err := json.Marshal(posts)
	if err != nil {
		log.Println(err)
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error marshalling the posts array"}`))
		return
	}
	resp.WriteHeader(http.StatusOK)
	resp.Write(result)
}

func addPost(resp http.ResponseWriter, req *http.Request) {
	var post Post
	err := json.NewDecoder(req.Body).Decode(&post)
	if err != nil {
		log.Println(err)
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error unmarshalling the request"}`))
		return
	}
	post.Id = len(posts) + 1
	posts = append(posts, post)
	resp.Header().Set("Content-type", "application/json")
	resp.WriteHeader(http.StatusOK)
	result, err := json.Marshal(post)
	if err != nil {
		log.Println(err)
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error marshalling the post"}`))
		return
	}
	resp.Write(result)
}
