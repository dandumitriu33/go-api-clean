package controller

import (
	"encoding/json"
	"go-api-clean/entity"
	"go-api-clean/errors"
	"go-api-clean/service"
	"log"
	"net/http"
)

type controller struct{}

var (
	postService service.PostService
)

type PostController interface {
	GetPosts(resp http.ResponseWriter, req *http.Request)
	AddPost(resp http.ResponseWriter, req *http.Request)
}

func NewPostController(service service.PostService) PostController {
	postService = service
	return &controller{}
}

func (*controller) GetPosts(resp http.ResponseWriter, req *http.Request) {
	posts, err := postService.FindAll()
	if err != nil {
		log.Println(err)
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: "Error getting the post"})
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

func (*controller) AddPost(resp http.ResponseWriter, req *http.Request) {
	var post entity.Post
	err := json.NewDecoder(req.Body).Decode(&post)
	if err != nil {
		log.Println(err)
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: "Error unmarshalling the request"})
		return
	}
	err = postService.Validate(&post)
	if err != nil {
		log.Println(err)
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: err.Error()})
		return
	}
	result, err := postService.Create(&post)
	if err != nil {
		log.Println(err)
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: "Error saving the post"})
		return
	}
	resp.Header().Set("Content-type", "application/json")
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(result)
	// result, err := json.Marshal(post)
	// if err != nil {
	// 	log.Println(err)
	// 	resp.WriteHeader(http.StatusInternalServerError)
	// 	resp.Write([]byte(`{"error": "Error marshalling the post"}`))
	// 	return
	// }
	// resp.Write(result)
}
