package main

import (
	"flag"
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

func main() {
	godotenv.Load()

	var portNum string
	flag.StringVar(&portNum, "port", "", "Application's port")

	var fileLocation string
	flag.StringVar(&fileLocation, "location", "", "File's Location")
	flag.Parse()

	static_file_path := "./public"
	if path := os.Getenv("FILE_PATH"); path != "" {
		fmt.Println("environment variable detected, Changing location to ", path)
		static_file_path = path
	}

	if fileLocation != "" {
		fmt.Println("input detected, Changing location to ", fileLocation)
		static_file_path = fileLocation
	}

	fs := http.FileServer(http.Dir(static_file_path))
	http.Handle("/", fs)

	port := "3000"
	if port_env := os.Getenv("PORT"); port_env != "" {
		fmt.Println("environment variable detected, Changing port to ", port_env)
		port = port_env
	}

	if portNum != "" {
		fmt.Println("input detected, Changing port to ", portNum)
		port = portNum
	}

	log.Printf("Listening on :%v...", port)
	err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
	if err != nil {
		log.Fatal(err)
	}
}
