//contendrá un puntero a la estructura http.Server de la biblioteca estándar.
package server

import (
	"log"
	"net/http"
	"time"

	v1 "Sistema/internal/server/v1"

	"github.com/go-chi/chi"
)

//server es la base de la configuracion del servidor

type Server struct {
	server *http.Server
}

//New inicializa un nuevo servidor con su respectiva configuracion
/*Desde el paquete chi estamos usando la función
NewRouter, para obtener un nuevo Router y lo
agregamos como Handler en la estructura http.Server,
ademas de establecer el Addr con el port que recibe la función y
limitar el tiempo de lectura y escritura a 10 segundos, */

func New(port string) (*Server, error) {
	r := chi.NewRouter()

	r.Mount("api/v1", v1.New())

	serv := &http.Server{
		Addr:         ":" + port,
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	server := Server{server: serv}

	return &server, nil
}

//Para cerrar los recursos del servidor
/*Y ahora agregamos un par de métodos a la
 estructura Server, uno que llamaremos
 Close que utilizaremos para cerrar recursos
antes de terminar con la ejecución del programa
(por ejemplo, la conexión con la de base de datos,
aunque de momento este método estará vacío) y uno
llamado Start, en el que activamos el servidor.*/

func (serv *Server) Close() error {
	//agrega los recursos necesario para cerrar
	return nil
}

//Iniciar el servidor

func (serv *Server) Start() {
	log.Printf(" Corriendo el servidor en http://localhost%s", serv.server.Addr)
	log.Fatal(serv.server.ListenAndServe())
}
