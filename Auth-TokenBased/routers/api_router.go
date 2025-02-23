package routers

import (
	"../controllers"
	"../core"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
)

func SetApiRouter(router *mux.Router) {
	commonHandler := alice.New(core.LoggingMiddleware)
	router.Handle("/", commonHandler.ThenFunc(controllers.ApiHandler)).Methods("GET")
}
