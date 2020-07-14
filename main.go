package main

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	godotenv.Load()

	static_file_path := "./public"
	if path := os.Getenv("FILE_PATH"); path != ""{
		static_file_path = path
	}

	fs := http.FileServer(http.Dir(static_file_path))
	http.Handle("/", fs)

	port := "3000"
	if port_env := os.Getenv("PORT"); port_env != ""{
		port = port_env
	}

	log.Printf("Listening on :%v...",port)
	err := http.ListenAndServe(fmt.Sprintf(":%v",port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
