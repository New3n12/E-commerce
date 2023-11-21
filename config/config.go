package config

import (
	"fmt"
	"net/http"

	"github.com/gorilla/sessions"
)

// inicializa la session
var Store = sessions.NewCookieStore([]byte("session-flowers"))

// Verifica si existe el usuario y retorna su id
func ExitUser(r *http.Request) (bool, int) {
	session, _ := Store.Get(r, "session-name")
	userIdRaw := session.Values["usuarioId"]
	userId, ok := userIdRaw.(int)
	if !ok {
		fmt.Println("No se pudo convertir usuarioId a int")
	}
	if userId != 0 {
		return true, userId
	} else {
		return false, 0
	}

}

// Elimina la session del usuario
func DeadSession(r *http.Request, w http.ResponseWriter) bool {
	session, err := Store.Get(r, "session-name")
	if err != nil {
		return false
	}
	delete(session.Values, "usuarioId")
	session.Options.MaxAge = -1
	session.Save(r, w)
	return true
}

// Se crea un flash para manejo de errores solo cadenas de texto
func CrearMensajesFlash(response http.ResponseWriter, request *http.Request, msg string) {
	session, err := Store.Get(request, "flash-session")
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	session.AddFlash(msg, "Error")
	session.Save(request, response)
}

// Se recupera el mensaje flash para manejo de errores
func RetornarMensajesFlash(response http.ResponseWriter, request *http.Request) string {
	session, _ := Store.Get(request, "flash-session")

	fm := session.Flashes("Error")
	session.Save(request, response)
	css_sesion := ""
	if len(fm) == 0 {
		css_sesion = ""
	} else {
		css_sesion = fm[0].(string)
	}

	return css_sesion
}

// Se crea un mesajje flash de tipo arreglo para la filtracion de productos
func CrearProductoFlash(response http.ResponseWriter, request *http.Request, sql []string) {
	session, err := Store.Get(request, "flash-session")
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
	session.AddFlash(sql, "Producto")
	err = session.Save(request, response)
	if err != nil {
		http.Error(response, err.Error(), http.StatusInternalServerError)
		return
	}
}

// se recuepra flash del arreglo para la filtarcion de productos
func RetornarProductoFlash(response http.ResponseWriter, request *http.Request) []string {
	session, _ := Store.Get(request, "flash-session")
	prod := session.Flashes("Producto")
	session.Save(request, response)
	var data []string

	if len(prod) == 0 {
		data = nil
	} else if prodSlice, ok := prod[0].([]string); ok {
		data = prodSlice
	} else {
		// Manejar el caso en el que el tipo no es []string
		// Puedes mostrar un error o tomar otra acción apropiada
	}

	return data
}
