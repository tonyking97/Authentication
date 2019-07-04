package routers

import (
	"../controllers"
	"../core"
	"github.com/gobuffalo/packr"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"net/http"
)

/*
* @Author : Arputha Tony King P @Created at : Apr 19
* Definition : Routes all the URL to their respective method
 */

func SetAuthenticationRouter(router *mux.Router) {

	//Load static files to the build file
	buildHandlerBox := packr.NewBox("../build")
	staticHandlerBox := packr.NewBox("../build/static")

	//Static routers for Authentication web page
	commonHandler := alice.New(core.LoggingMiddleware)
	//buildHandler := http.FileServer(http.Dir(buildHandlerPath))
	buildHandler := http.FileServer(buildHandlerBox)
	//staticHandler := http.StripPrefix("/static/", http.FileServer(http.Dir(staticHandlerPath)))
	staticHandler := http.StripPrefix("/static/", http.FileServer(staticHandlerBox))

	//Authentication routers
	router.Handle("/", buildHandler).Methods("GET")
	router.PathPrefix("/static/").Handler(staticHandler)
	//Login
	router.Handle("/", commonHandler.ThenFunc(controllers.LoginHandler)).Methods("POST")
	//Check Username is Available
	router.Handle("/checkUsername", commonHandler.ThenFunc(controllers.CheckUsername)).Methods("POST")
	//Get Username
	router.Handle("/getNameDetails", commonHandler.ThenFunc(controllers.GetNameDetails)).Methods("POST")
	//Check Username is Availability on signup
	router.Handle("/checkUsernameAvailability", commonHandler.ThenFunc(controllers.CheckUsernameAvailability)).Methods("POST")
	//Check Email is Availability on signup
	router.Handle("/checkEmailAvailability", commonHandler.ThenFunc(controllers.CheckEmailAvailability)).Methods("POST")
	//GetActiveSessionListDetails
	router.Handle("/getActiveSessionDetails", commonHandler.ThenFunc(controllers.GetActiveSessionDetails)).Methods("GET")
	//GetSessionListDetails
	router.Handle("/getSessionDetails", commonHandler.ThenFunc(controllers.GetSessionDetails)).Methods("POST")
	//GetInActiveSessionListDetails
	router.Handle("/getInActiveSessionDetails", commonHandler.ThenFunc(controllers.GetInActiveSessionDetails)).Methods("GET")
	//GetCurrentSessionDetails
	router.Handle("/getCurrentSessionDetails", commonHandler.ThenFunc(controllers.GetCurrentSessionDetails)).Methods("POST")
	//Signup
	router.Handle("/signup", commonHandler.ThenFunc(controllers.SignUpHandler)).Methods("POST")
	//RefreshToken
	router.Handle("/refreshtoken", commonHandler.ThenFunc(controllers.RefreshTokenHandler)).Methods("POST")
	//checkToken
	router.Handle("/checktoken", commonHandler.ThenFunc(controllers.CheckTokenHandler)).Methods("POST")
	//Logout
	router.Handle("/logout", commonHandler.ThenFunc(controllers.LogoutHandler)).Methods("POST")
}
