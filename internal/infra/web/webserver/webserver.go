package webserver

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type WebServer struct {
	Router        chi.Router
	Handlers      []WebServerHandler
	WebServerPort string
}

type WebServerHandler struct {
	Path   string
	Func   http.HandlerFunc
	Method string
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:        chi.NewRouter(),
		Handlers:      []WebServerHandler{},
		WebServerPort: serverPort,
	}
}

func (s *WebServer) AddHandler(path string, handler http.HandlerFunc, method string) {
	s.Handlers = append(s.Handlers, WebServerHandler{
		Path:   path,
		Func:   handler,
		Method: method,
	})
}

// loop through the handlers and add them to the router
// register middeleware logger
// start the server
func (s *WebServer) Start() {
	s.Router.Use(middleware.Logger)
	for _, handler := range s.Handlers {
		switch handler.Method {
		case "POST":
			s.Router.Post(handler.Path, handler.Func)
		case "GET":
			s.Router.Get(handler.Path, handler.Func)
		}
	}
	http.ListenAndServe(s.WebServerPort, s.Router)
}
