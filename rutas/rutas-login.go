package rutas

import (
	controlador "Go-Proyect/Controlador"
	"Go-Proyect/modelo"
	"html/template"
	"net/http"
	"time"
)

type Mensaje_Error struct {
	Error string
}

func (p Mensaje_Error) GetMsgErr() string {
	return p.Error
}

func (p *Mensaje_Error) SetMsgErr(msg string) {
	p.Error = msg
}

var data = Mensaje_Error{
	Error: "",
}

func Login(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles("Vista/templates/login.html"))
	if data.GetMsgErr() != "" {
		template.Execute(w, data)
		data.SetMsgErr("")
	} else {
		data.Error = ""
		template.Execute(w, data)
	}
}

func NewRegistro(w http.ResponseWriter, r *http.Request) {
	fechaActual := time.Now()
	formato := "2006-01-02"
	fechaStr := fechaActual.Format(formato)
	usu := modelo.User{
		ID:            1,
		Nombre:        r.FormValue("nombre"),
		Correo:        r.FormValue("correo"),
		Edad:          string(r.FormValue("edad")),
		Contra:        string(r.FormValue("contra")),
		FechaCreacion: fechaStr,
	}

	registro := controlador.Registro(usu, r, w)

	if registro {
		w.Write([]byte("true"))
	} else {
		w.Write([]byte("false"))
	}

}

func ValidarUsu(w http.ResponseWriter, r *http.Request) {
	valida := controlador.Validar(r.FormValue("correo"), r.FormValue("password"), r, w)
	if valida {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	} else {
		data.SetMsgErr("Usuario / Contraseña Incorrecta")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}
