package rutas

import (
	controlador "Go-Proyect/Controlador"
	"Go-Proyect/config"
	"Go-Proyect/modelo"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"
)

func Card(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
		}
	}()
	template := template.Must(template.ParseFiles("Vista/templates/card.html", "Vista/Layout/layout.html"))
	data := Card_Produc{}
	exisid, id := config.ExitUser(r)
	if exisid {
		card := controlador.AllCards(id)
		data.Cantida_card = len(card)
		data.Card = card
		data.UserId = id
		data.Total = Total_Precie(card)
	}
	data.Mensaje_Error = config.RetornarMensajesFlash(w, r)
	template.Execute(w, data)
}

func Operaciones_card(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
		}
	}()
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}

	type Data struct {
		Case      int
		IdProduct int
		IdUser    int
		Quantity  int
		Price     int
	}
	var d Data
	d.Case = stringToInt(r.FormValue("caso"))
	switch d.Case {
	case 0:
		d.IdProduct = stringToInt(r.FormValue("id"))
		d.IdUser = stringToInt(r.FormValue("idusu"))
		d.Quantity = stringToInt(r.FormValue("cant"))
		//d.Price = stringToInt(r.FormValue("prec"))
		cond, msg, id := controlador.Add_Card(d.IdProduct, d.IdUser, d.Quantity)
		if cond {
			if msg == "ok" {
				produc := controlador.Id_Product(d.IdProduct)
				product := modelo.Produc_Card{
					ID_card:       id,
					Nombre_card:   produc.Nombre,
					Precio_card:   produc.Precio,
					Cantidad_card: d.Quantity,
					Img_card:      produc.Img1,
				}
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(product)
				return
			}
			w.Write([]byte(msg))
			return
		}
		break
	case 1:
		d.IdProduct = stringToInt(r.FormValue("id"))
		if controlador.Delete_Card(d.IdProduct) {
			data := Card_Produc{}
			exisid, id := config.ExitUser(r)
			if exisid {
				card := controlador.AllCards(id)
				data.Total = Total_Precie(card)
				floatAsString := strconv.FormatFloat(data.Total, 'f', -1, 64)
				w.Write([]byte(floatAsString))
				return
			}
			w.Write([]byte("ok"))
			return
		}
		break
	case 2:
		d.IdProduct = stringToInt(r.FormValue("id"))
		if controlador.More_Card(d.IdProduct) {
			data := Card_Produc{}
			exisid, id := config.ExitUser(r)
			if exisid {
				card := controlador.AllCards(id)
				data.Total = Total_Precie(card)
				floatAsString := strconv.FormatFloat(data.Total, 'f', -1, 64)
				w.Write([]byte(floatAsString))
				return
			}
			w.Write([]byte("ok"))
			return
		} else {
			w.Write([]byte("ok"))
			return
		}
		//break
	case 3:
		d.IdProduct = stringToInt(r.FormValue("id"))
		condi, mesg := controlador.Less_Card(d.IdProduct)
		if condi {
			data := Card_Produc{}
			exisid, id := config.ExitUser(r)
			if exisid {
				card := controlador.AllCards(id)
				data.Total = Total_Precie(card)
				floatAsString := strconv.FormatFloat(data.Total, 'f', -1, 64)
				data := struct {
					Mensaje string
					Total   string
				}{
					Mensaje: mesg,
					Total:   floatAsString,
				}
				w.Header().Set("Content-Type", "application/json")
				json.NewEncoder(w).Encode(data)
				return
			}
			w.Write([]byte(mesg))
			return
		}
		break
	default:
		http.Error(w, "Operación no válida", http.StatusBadRequest)
	}
}

func stringToInt(s string) int {
	i, _ := strconv.Atoi(s)
	return i
}
func Total_Precie(list []modelo.Produc_Card) float64 {
	total := 0.0

	for _, p := range list {
		total += p.Precio_card * float64(p.Cantidad_card)
	}

	return total
}
