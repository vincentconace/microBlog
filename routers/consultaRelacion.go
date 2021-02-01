package routers

import (
	"encoding/json"
	"net/http"

	"github.com/vincentconace/microBlog/bd"
	"github.com/vincentconace/microBlog/models"
)

/*ConsultaRelacion chequea si hay relación entre dos usuarios*/
func ConsultaRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	var rel models.Relacion

	rel.UsuarioID = IDUsuario
	rel.UsuarioRelacionID = ID

	var resp models.RespuestaConsultaRelacion

	status, err := bd.ConsultoRelacion(rel)

	// No pierdo tiempo mostrando mensajes de error y me focalizo en responder true o false
	if err != nil || status == false {
		resp.Status = false
	} else {
		resp.Status = true
	}

	// Ahora seteo el Header indicando que la respuesta va a ser un json
	w.Header().Set("Content-type", "application/json")
	// escribo el header ahora con el status created
	w.WriteHeader(http.StatusCreated)
	// ahora hacemos nuestro encode para pasar a json nuestro modelo
	json.NewEncoder(w).Encode(resp)
}
