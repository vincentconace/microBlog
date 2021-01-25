package models

import "time"

//GRaboTweet es el formato o estructura que tendra nuestro Tweet
type GraboTweet struct {
	UsuarioID string    `bson:"usuarioid" json:"usuarioid,omitemty"`
	Mensaje   string    `bson:"mensaje" json:"mensaje,omitemty"`
	Fecha     time.Time `bson:"fecha" json:"fecha,omitemty"`
}
