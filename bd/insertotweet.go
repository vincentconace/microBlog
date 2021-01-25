package bd

import (
	"context"
	"time"

	"github.com/vincentconace/microBlog/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//InsertoTweet graba el Tweet en la BD
func InsertoTweet(u models.GraboTweet) (string, bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("microBlog")
	col := db.Collection("tweet")

	registro := bson.M{
		"usuarioid": u.UsuarioID,
		"mensaje":   u.Mensaje,
		"fecha":     u.Fecha,
	}
	result, err := col.InsertOne(ctx, registro)
	if err != nil {
		return "", false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)
	return objID.String(), true, nil
}
