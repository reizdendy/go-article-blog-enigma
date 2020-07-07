package configs

import (
	"blog-article-go/utils"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func CreateRouter() *mux.Router {
	router := mux.NewRouter()
	return router
}

func RunServer(router *mux.Router) {
	ServerHost := utils.GetCustomConf("ServerHost", "localhost")
	ServerPort := utils.GetCustomConf("ServerPort", "3500")

	serverSource := fmt.Sprintf("%s:%s", ServerHost, ServerPort)

	fmt.Printf("Starting Web Server At %s Port : %s", ServerHost, ServerPort)
	err := http.ListenAndServe(serverSource, router)
	if err != nil {
		log.Fatalf("%v", err)
	}
}
