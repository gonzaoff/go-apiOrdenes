package aplicacion

import (
	"context"
	"fmt"
	"net/http"
)

// Vamos a importar las librerias y crear un archivo App con una estructura para contener nuestro router

type App struct {
	router http.Handler
}

func New() *App {
	app := &App{
		router: loadRoutes(),
	}

	return app

}


// Vamos a crear una funcion Iniciadora que hara referencia a al puntero de datos de "App" con un unico contexto.


func (a *App) Start(ctx context.Context) error {
	// Funcion iniciadora(
	server := &http.Server{
		Addr: ":3000",
		Handler : a.router,
	}

	err := server.ListenAndServe()

	if err != nil {
		return fmt.Errorf("Fallo al iniciar el servidor: %w", err)
	}
	// )
	

	return nil //(?)
}

