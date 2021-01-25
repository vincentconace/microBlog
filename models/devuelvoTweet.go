package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//DevuelvoTweet es la estructura con la que devuelvo los Tweet
type DevuelvoTweets struct {
	ID        primitive.ObjectID `bson:"_id" json:"_id,omitemty"`
	UsuarioID string             `bson:"usuarioid" json:"usuarioId,omitemty"`
	Mensaje   string             `bson:"mesaje" json:"mensaje,omitemty"`
	Fecha     time.Time          `bson:"fecha" json:"fecha,omitemty"`
}
