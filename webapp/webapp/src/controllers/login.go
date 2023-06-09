package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"webapp/src/config"
	"webapp/src/modelos"
	"webapp/src/respostas"
)

//FazerLogin utiliza o e-mail e senha do usuário para autenticar na aplicação
func FazerLogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	usuario, erro := json.Marshal(map[string]string{
		"email": r.FormValue("email"),
		"senha": r.FormValue("senha"),
	})

	if erro != nil {
		respostas.JSON(w, http.StatusBadRequest, respostas.ErroAPI{Erro: erro.Error()})
		return
	}

	url := fmt.Sprintf("%s/login", config.APIURL)
	response, erro := http.Post(url, "application/json", bytes.NewBuffer(usuario))
	if erro != nil {
		respostas.JSON(w, http.StatusInternalServerError, respostas.ErroAPI{Erro: erro.Error()})
		return
	}
	defer response.Body.Close()

	corpoResposta, _ := ioutil.ReadAll(response.Body)
	log.Println("Resposta do servidor:", string(corpoResposta))

	if response.StatusCode >= 400 {
		respostas.TratarStatusCodeDeErro(w, response)
		log.Printf("Erro ao fazer login. Status: %d - Resposta: %s", response.StatusCode, response.Body)
		return
	}

	var dadosAutenticacao modelos.DadosAutenticacao
	if erro = json.NewDecoder(response.Body).Decode(&dadosAutenticacao); erro != nil {
		respostas.JSON(w, http.StatusUnprocessableEntity, respostas.ErroAPI{Erro: erro.Error()})
		log.Printf("Erro ao fazer login. Status: %d - Resposta: %s", response.StatusCode, response.Body)
		return
	}

	//

	respostas.JSON(w, http.StatusOK, nil)
}
