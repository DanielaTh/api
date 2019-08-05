package main

import (
	"crudPackages/api"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	r := mux.NewRouter()
	h, err := api.NewHandler(viper.GetString("db"))
	if err != nil {
		log.Fatalln(err)
	}

	r.HandleFunc("/get/subject/{id}", h.GetSubject).Methods("GET")
	r.HandleFunc("/add/subject", h.AddSubject).Methods("POST")
	logHandler := handlers.LoggingHandler(os.Stdout, r)

	port := viper.GetInt("port")
	log.Println("puerto", port)
	http.ListenAndServe(":"+strconv.Itoa(port), logHandler)
}
