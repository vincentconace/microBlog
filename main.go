package main

import (
	"github.com/vincentconace/microBlog/bd"
	"github.com/vincentconace/microBlog/handlers"
	"log"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Manejadores()

}
