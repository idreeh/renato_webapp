package router

import (
	"webapp/src/router/rotas"

	"github.com/gorilla/mux"
)

//Gerar retorna um router com todas as rotas configuradas
func Gerar() *mux.Router {
	router := mux.NewRouter()
	rotas.Configurar(router)

	return router
}
