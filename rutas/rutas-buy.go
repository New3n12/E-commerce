package rutas

import (
	controlador "Go-Proyect/Controlador"
	"Go-Proyect/config"
	"html/template"
	"log"
	"net/http"
)

func Buy(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
		}
	}()
	template := template.Must(template.ParseFiles("Vista/templates/buy.html", "Vista/Layout/layout-compras.html"))
	data := Card_Produc{}
	exisid, id := config.ExitUser(r)
	if exisid {
		condi := controlador.Existe_Card(id)
		if !condi {
			config.CrearMensajesFlash(w, r, "Error")
			http.Redirect(w, r, "/productos/card", http.StatusSeeOther)
			return
		}
		//card := controlador.AllCards(id)
		//data.Cantida_card = len(card)
		//data.Card = card
		data.UserId = id
		//data.Total = Total_Precie(card)
		condicion, conut := controlador.ExistDirec(id)
		if condicion {
			if conut == 1 {
				data.Descri2 = controlador.DireccionByID2(data.UserId)
			} else if conut > 1 {
				data.Descri2 = controlador.DireccionByID2(data.UserId)
				data.Descri = controlador.DireccionByID(data.UserId)
			}
		}
		data.Mensaje_Error = config.RetornarMensajesFlash(w, r)
		template.Execute(w, data)
	} else {
		Pagina404(w, r)
	}
	return
}
