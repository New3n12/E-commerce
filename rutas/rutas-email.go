package rutas

import (
	"Go-Proyect/config"
	"fmt"
	"net/http"

	gomail "gopkg.in/gomail.v2"
)

func Sen_Email(w http.ResponseWriter, r *http.Request) {
	msg := gomail.NewMessage()
	msg.SetHeader("From", r.FormValue("email"))
	msg.SetHeader("To", "noreply@agendahoras.com")
	msg.SetHeader("Subject", r.FormValue("subject"))
	msg.SetBody("text/html", fmt.Sprintf(`<h1>%s</h1><p>%s</p>`, r.FormValue("nombre"), r.FormValue("contenido")))
	//msg.Attach("/Vista/img/logo.png")
	//ahora se debe configurar la conexión con el SMTP
	n := gomail.NewDialer("sandbox.smtp.mailtrap.io", 587, "6a18e0e4ef0294", "fc66ec67548195")

	if err := n.DialAndSend(msg); err != nil {
		config.CrearMensajesFlash(w, r, "Correo_err")
	}
	config.CrearMensajesFlash(w, r, "Correo_ok")
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
