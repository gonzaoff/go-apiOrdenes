package main

import (
	"fmt"

	// Vamos a importar la libreria context y el directorio donde se encuentran nuestras funciones

	"context"
	"github.com/gonzaoff/go-apiOrdenes/aplicacion"
)




// Vamos a limpiar el codigo 

func main() {

	// Vamos a hacer una nueva aplicacion basandonos en las funciones realizadas.

	app := aplicacion.New()
	
	fmt.Println("Iniciando servidor en: http://localhost:3000.")
	err := app.Start(context.TODO())


	if err != nil {
		fmt.Println("Fallo al iniciar la app: ", err)
	}

}

// Es el turno de los handlers (manipuladores de solicitud)

// De esa manera introdujimos(?) los terminos que se usan para manejar solicitudes, mañana se viene lo interesante... Redis
