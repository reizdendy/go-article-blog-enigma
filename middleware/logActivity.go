package middleware

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func ActivityLogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//Bebas berkarya
		userAgent := r.Header.Get("User-Agent")
		writelog := fmt.Sprintf("Accessing path %v with application %v from %v", r.RequestURI, userAgent, r.RemoteAddr)

		log.Printf("Accessing path %v with application %v from %v", r.RequestURI, userAgent, r.RemoteAddr)

		file, err := os.OpenFile("../logData/data.txt", os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			log.Println(err)
		}
		defer file.Close()
		if _, err := file.WriteString(writelog); err != nil {
			log.Println(err)
		}
		next.ServeHTTP(w, r)
	})
}
