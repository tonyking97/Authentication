package core

import (
	"log"
	"net/http"
	"time"
)

/*
* @Author : Arputha Tony King P @Created at : Apr 19
* Definition : Logging middleware to log the incoming request
 */

func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		defer func() {
			log.Println(r.RemoteAddr+" --> "+r.RequestURI, r.Method, time.Since(start))
		}()
		next.ServeHTTP(w, r)
	})
}
