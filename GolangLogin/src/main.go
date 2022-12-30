package main

import (
	"fmt"
	"net/http"

	authcontroller "github.com/jeypc/go-auth/controllers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", authcontroller.Index)
	mux.HandleFunc("/login", authcontroller.Login)
	mux.HandleFunc("/logout", authcontroller.Logout)
	mux.HandleFunc("/register", authcontroller.Register)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: mux,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
	fmt.Println("Server jalan di: http://localhost:3000")
}
