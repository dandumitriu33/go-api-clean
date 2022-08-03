// Go-API-clean is a clean architecture simple GO API with Mux.
package main

import (
	"fmt"
	"net/http"

	"go-api-clean/controller"
	router "go-api-clean/http"
	"go-api-clean/repository"
	"go-api-clean/service"
)

var (
	postRepository repository.PostRepository = repository.NewMySQLRepository()
	postService    service.PostService       = service.NewPostService(postRepository)
	postController controller.PostController = controller.NewPostController(postService)
	httpRouter     router.Router             = router.NewChiRouter()
)

func main() {

	const port string = ":8000"
	httpRouter.GET("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "Up and running")
	})
	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/posts", postController.AddPost)
	httpRouter.SERVE(port)
}
