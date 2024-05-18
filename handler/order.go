package handler

import (
	"fmt"
	"net/http"
)

// Vamos a crear un controlador vacio con el que despues haremos operaciones cuadruples

type Order struct{}

func (o *Order) Crear(w http.ResponseWriter, r *http.Request){
	fmt.Println("Orden creada.")
}
func (o *Order) Lista(w http.ResponseWriter, r *http.Request){
	fmt.Println("Lista de ordenes.")
}
func (o *Order) BuscarPorID(w http.ResponseWriter, r *http.Request){
	fmt.Println("Esta es la orden {id}")
}
func (o *Order) ActualizarPorID(w http.ResponseWriter, r *http.Request){
	fmt.Println("Le haz hecho un put o haz actualizado la orden {id}")
}
func (o *Order) EliminarPorID(w http.ResponseWriter, r *http.Request){
	fmt.Println("Haz eliminado la orden {id}")
}

