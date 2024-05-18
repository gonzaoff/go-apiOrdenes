package aplicacion

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gonzaoff/go-apiOrdenes/handler"
)

// Defino una funcion de ruta que devuelve un puntero al tipo de g.mux dentro del que creo una instancia parecida a la que hicimos en el main
func loadRoutes() *chi.Mux {


	router := chi.NewRouter()
	router.Use(middleware.Logger)


	// Definimos una funcion get (solicitud de datos) en la raiz de nuestra pagina web
						// Definimos la funcion que va a leer y escribir en la solicitud
	router.Get("/", func(w http.ResponseWriter, r *http.Request){

		// Le decimos que nos debe responder a la hora de solicitar datos a la raiz de la pagina web
		w.WriteHeader(http.StatusOK)
	})
 
	// Ahora vamos a cargar las ordenes del enrutador (post, get, put y delete [postear(?, solicitar, actualizar y eliminar]), Vamos a tener en cuenta las funciones creadas anteriormente
	// Para eso vamos a cargar los enrutadores en la ruta "ordenes" seguida de la funcion que cargara las ordenes.
	router.Route("/ordenes", rutasCargaOrdenes)

	// Retornamos el enrutador 
	return router
	
}

func rutasCargaOrdenes(router chi.Router) {
	orderHandler := &handler.Order{}

	router.Post("/", orderHandler.Crear)
	router.Get("/", orderHandler.Lista)

	// Tres de estas rutas seran vereficadas por un numero de identificacion {id}.
	router.Get("/{id}", orderHandler.BuscarPorID)
	router.Put("/{id}", orderHandler.ActualizarPorID)
	router.Delete("/{id}", orderHandler.EliminarPorID)
}