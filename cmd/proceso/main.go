package main

import (
	//"fmt"
	"Sistema/internal/data"
	"Sistema/internal/server"
	"log"
	"os"
	"os/signal"

	_ "github.com/joho/godotenv/autoload"
)

func main() {

	port := os.Getenv("PORT")
	serv, err := server.New(port)
	if err != nil {
		log.Fatal(err)
	}

	//conexion a la base de datos
	d := data.New()
	if err := d.DB.Ping(); err != nil {
		log.Fatal(err)
	}

	//Inicia el servidor

	go serv.Start()

	//Espera por si hay interrupcion
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	//Intenta un cierre elegante
	serv.Close()
	data.New().Close()

}
