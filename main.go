package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {

	router := chi.NewRouter()
	router.Use(middleware.Logger)


	router.Get("/hola", basicHandler)


	server := &http.Server{
		Addr: ":3000",
		Handler: router,
	}
	
	fmt.Println("Servidor: online.\nHost: http://localhost:3000")

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Fallo el servidor", err)
	}

	
}

func basicHandler(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("Hola Mundo!"))
}

