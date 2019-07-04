package routers

import (
	"github.com/gorilla/mux"
)

/*
* @Author : Arputha Tony King P @Created at : Apr 19
* Definition : Routes the URL to the specific Routers based on the URL
 */

//Initialize the router
func InitRouters() *mux.Router {
	router := mux.NewRouter()

	//apiRouter := router.PathPrefix("/api").Subrouter().StrictSlash(true)
	authenticationRouter := router.PathPrefix("/").Subrouter().StrictSlash(true)

	//SetApiRouter(apiRouter)
	SetAuthenticationRouter(authenticationRouter)

	return router
}
