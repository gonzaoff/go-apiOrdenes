package main

import (
	"fmt"
	"net/http"
)

func main() {

	server := &http.Server{
		Addr: ":3000",
		Handler: http.HandlerFunc(basicHandler),
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

