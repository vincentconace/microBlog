package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/vincentconace/microBlog/bd"
	"github.com/vincentconace/microBlog/models"
)

func GraboTweet(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet
	err := json.NewDecoder(r.Body).Decode(&mensaje)

	registro := models.GraboTweet{
		UsuarioID: IDUsuario,
		Mensaje:   mensaje.Mensaje,
		Fecha:     time.Now(),
	}

	_, status, err := bd.InsertoTweet(registro)
	if err != nil {
		http.Error(w, "Ocurrio un erro al insertar el registro reintente nuevamente"+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se a logrado insertar el Tweet", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
