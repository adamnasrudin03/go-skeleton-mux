package router

import (
	"net/http"
	"time"

	response_mapper "github.com/adamnasrudin03/go-helpers/response-mapper/v1"
	"github.com/adamnasrudin03/go-skeleton-mux/app/controller"

	"github.com/gorilla/mux"
)

type routes struct {
	HttpServer *mux.Router
}

func NewRoutes(h controller.Controllers) routes {
	var err error
	r := routes{
		HttpServer: mux.NewRouter(),
	}

	r.HttpServer.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response_mapper.RenderJSON(w, http.StatusOK, response_mapper.NewResponseMultiLang(response_mapper.MultiLanguages{
			ID: "selamat datang di server ini",
			EN: "Welcome this server",
		}))
	}).Methods("GET")

	r.HttpServer.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		err = response_mapper.ErrRouteNotFound()
		response_mapper.RenderJSON(w, http.StatusNotFound, err)
	})
	return r
}

func (r routes) Run(addr string) error {
	server := &http.Server{
		Addr:         addr,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r.HttpServer,
	}
	return server.ListenAndServe()
}
