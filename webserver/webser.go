package webserver

import "net/http"

func MyWebServer() {

	// Esta función establece un controlador para una ruta determinada en el servidor web.
	// En este caso, se está estableciendo un controlador para la ruta raíz del servidor ("/") y 
	// se está asociando con la función home.
	http.HandleFunc("/", home)

	// Esta función inicia el servidor web en el puerto especificado.
	http.ListenAndServe(":3000", nil)
}

// w http.ResponseWriter -> Este objeto se utiliza para escribir la respuesta HTTP que se enviará al cliente. 
// r *http.Request -> Este es un objeto que representa la solicitud HTTP entrante. 
func home(w http.ResponseWriter, r *http.Request) {

	// Esta función se usa para servir un archivo HTML alojado localmente en la ruta.
	http.ServeFile(w, r, "./webserver/index.html")
}