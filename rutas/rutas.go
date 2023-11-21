package rutas

import (
	controlador "Go-Proyect/Controlador"
	"Go-Proyect/config"
	"Go-Proyect/modelo"

	"html/template"
	"net/http"
)

type Card_Produc struct {
	Products        []modelo.Product
	Card            []modelo.Produc_Card
	Product         modelo.Product
	Cantida_card    int
	UserId          int
	Total           float64
	Descri          []modelo.FormData
	Descri2         modelo.FormData
	Condi           string
	Mensaje_Error   string
	Categoria       modelo.Categoria
	Products_1      []modelo.Product
	Products_2      []modelo.Product
	Products_3      []modelo.Product
	Products_oferta []modelo.Product
}

func Index(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles("Vista/templates/index.html", "Vista/Layout/layout.html"))
	data := Card_Produc{}
	exisid, id := config.ExitUser(r)
	if exisid {
		card := controlador.AllCards(id)
		data.Cantida_card = len(card)
		data.Card = card
		data.UserId = id
	}
	data.Products_1 = controlador.PrductosLimit4(1)
	data.Products_2 = controlador.PrductosLimit4(2)
	data.Products_3 = controlador.PrductosLimit4(4)
	data.Mensaje_Error = config.RetornarMensajesFlash(w, r)
	template.Execute(w, data)
}

func Servicio(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles("Vista/templates/servicio.html", "Vista/Layout/layout.html"))
	data := Card_Produc{}
	exisid, id := config.ExitUser(r)
	if exisid {
		card := controlador.AllCards(id)
		data.Cantida_card = len(card)
		data.Card = card
		data.UserId = id
	}
	template.Execute(w, data)
}

func Logout(w http.ResponseWriter, r *http.Request) {
	logout := config.DeadSession(r, w)
	if logout {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}
}

func Pagina404(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles("Vista/templates/error.html", "Vista/Layout/layout-compras.html"))
	template.Execute(w, nil)
}
