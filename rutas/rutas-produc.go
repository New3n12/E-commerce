package rutas

import (
	controlador "Go-Proyect/Controlador"
	"Go-Proyect/config"
	"Go-Proyect/modelo"
	"regexp"
	"strings"

	//"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// ruta para ver la descripcion de los productos
func Produc_descripcion(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
		}
	}()
	temp := template.Must(template.ParseFiles("Vista/templates/produc-descripcion.html", "Vista/Layout/layout.html"))
	data := Card_Produc{}
	exisid, id := config.ExitUser(r)
	if exisid {
		card := controlador.AllCards(id)
		data.Cantida_card = len(card)
		data.Card = card
		data.UserId = id
	}
	vars := mux.Vars(r)
	ids, err := strconv.Atoi(vars["id"])
	idcat, err2 := strconv.Atoi(vars["idcat"])
	data.Product = controlador.Id_Product(ids)
	data.Products = controlador.PrductosSugerencias(ids, idcat)
	if err != nil || err2 != nil {
		Pagina404(w, r)
	} else if data.Product.ID != 0 && ids == data.Product.ID {
		temp.Execute(w, data)
	} else {
		http.Redirect(w, r, "/productos", http.StatusSeeOther)
	}
}

func Prductos(w http.ResponseWriter, r *http.Request) {
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
		}
	}()
	template := template.Must(template.ParseFiles("Vista/templates/productos.html", "Vista/Layout/layout.html"))
	data := Card_Produc{}
	exisid, id := config.ExitUser(r)
	if exisid {
		card := controlador.AllCards(id)
		data.Cantida_card = len(card)
		data.Card = card
		data.UserId = id
	}
	body := r.URL.Query().Get("sql")

	if body != "" {
		newsql := LimpiarSQL(body)
		var sql []string
		sql = strings.Split(newsql, ",")
		data.Categoria = Rellenar(sql)
		pro, msg := controlador.Filtro_Produc(sql)
		if pro != nil && msg == "ok" {
			data.Products = pro
		} else if msg == "all" {
			data.Products = controlador.All_Products()
		}
	} else {
		data.Products = controlador.All_Products()
	}

	if data.Products != nil {
		template.Execute(w, data)
		return
	} else {
		template.Execute(w, nil)
	}
	return
}

func Rellenar(produc []string) modelo.Categoria {
	var data modelo.Categoria
	for _, item := range produc {
		switch item {
		case "Flores":
			data.Flores = "checked"
			break
		case "Plantas de Interior":
			data.Plantas_de_Interior = "checked"
			break
		case "Plantas de Sombra":
			data.Plantas_de_Sombra = "checked"
			break
		case "Plantas de Exterior":
			data.Plantas_de_Exterior = "checked"
			break
		case "Acuáticas":
			data.Acuaticas = "checked"
			break
		}
	}
	return data
}

func LimpiarSQL(body string) string {
	body = strings.ReplaceAll(body, "AND", "")
	body = strings.ReplaceAll(body, "OR", "")
	body = strings.ReplaceAll(body, "and", "")
	body = strings.ReplaceAll(body, "=", "")
	body = strings.ReplaceAll(body, "!=", "")
	expresionRegular := regexp.MustCompile("[0-9]+")
	cadenaLimpia := expresionRegular.ReplaceAllString(body, "")
	return cadenaLimpia
}
