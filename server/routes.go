package server

import (
	"net/http"
)

func initRoutes() {
	http.HandleFunc("/", index)
	http.HandleFunc("/agregar", form)
	http.HandleFunc("/insertar", insertar)
	http.HandleFunc("/borrar", Borrar)
	http.HandleFunc("/editar", Editaregistro)
	http.HandleFunc("/actualizar", Actualizar)
	http.HandleFunc("/concurrencia", Filtro)
	http.HandleFunc("/concurrenciapool", Pool)
}
