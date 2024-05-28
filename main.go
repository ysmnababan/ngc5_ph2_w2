package main

import (
	"fmt"
	"log"
	"net/http"
	"pagi/config"
	"pagi/handler"

	"github.com/julienschmidt/httprouter"
)

func logRequest(message string) func(httprouter.Handle) httprouter.Handle {
	return func(next httprouter.Handle) httprouter.Handle {
		return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
			log.Println(message)
			next(w, r, p)
		}
	}
}

func main() {
	db := config.Connect()
	defer db.Close()

	handler := &handler.MysqlDB{DB: db}
	router, server := config.InitServer()
	router.POST("/register", logRequest("request sent to POST /register")(handler.Register))
	router.POST("/login", logRequest("request sent to POST /login")(handler.Login))

	fmt.Println("server running on localhost:8080")
	panic(server.ListenAndServe())
}
