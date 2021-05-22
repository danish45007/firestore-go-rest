package router

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var (
	chiDispatcher = chi.NewRouter()
)

type chiRouter struct{}

func NewChiRouter() Router {
	return &chiRouter{}
}

func (*chiRouter) GET(uri string, f func(res http.ResponseWriter, req *http.Request)) {
	chiDispatcher.Get(uri, f)
}

func (*chiRouter) POST(uri string, f func(res http.ResponseWriter, req *http.Request)) {
	chiDispatcher.Post(uri, f)
}

func (*chiRouter) SERVE(port string) {
	fmt.Printf("Chi HTTP server running on port %v", port)
	http.ListenAndServe(port, chiDispatcher)
}

func (*chiRouter) MIDDLEWARE() {
	chiDispatcher.Use(middleware.Logger)
}
