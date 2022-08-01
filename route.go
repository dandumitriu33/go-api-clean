package main

import (
	"encoding/json"
	"go-api-clean/entity"
	"go-api-clean/repository"
	"log"
	"net/http"
)

var (
	// posts []entity.Post
	repo repository.PostRepository = repository.NewPostRepository()
)

// func init() {
// 	posts = []entity.Post{{ID: 1, Title: "Title One", Text: "Text One"}}
// }

func getPosts(resp http.ResponseWriter, req *http.Request) {
	posts, err := repo.FindAll()
	if err != nil {
		log.Println(err)
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error getting data."}`))
		return
	}
	resp.Header().Set("Content-type", "application/json")
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(posts)
	// result, err := json.Marshal(posts)
	// if err != nil {
	// 	log.Println(err)
	// 	resp.WriteHeader(http.StatusInternalServerError)
	// 	resp.Write([]byte(`{"error": "Error marshalling the posts array"}`))
	// 	return
	// }
	//resp.Write(result)
}

func addPost(resp http.ResponseWriter, req *http.Request) {
	var post entity.Post
	err := json.NewDecoder(req.Body).Decode(&post)
	if err != nil {
		log.Println(err)
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error unmarshalling the request"}`))
		return
	}
	// post.ID = len(posts) + 1
	post.ID = 0
	// posts = append(posts, post)
	repo.Save(&post)

	resp.Header().Set("Content-type", "application/json")
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(post)
	// result, err := json.Marshal(post)
	// if err != nil {
	// 	log.Println(err)
	// 	resp.WriteHeader(http.StatusInternalServerError)
	// 	resp.Write([]byte(`{"error": "Error marshalling the post"}`))
	// 	return
	// }
	// resp.Write(result)
}
