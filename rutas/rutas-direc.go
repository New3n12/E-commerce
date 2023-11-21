package rutas

import (
	controlador "Go-Proyect/Controlador"
	"Go-Proyect/config"
	"Go-Proyect/modelo"
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Add_Direccion(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
		return
	}
	id, err := strconv.Atoi(r.FormValue("id"))
	if err != nil {
		id = 0
	}
	temp_saved := 0
	if r.FormValue("checkbox") == "on" {
		temp_saved = 1
	}

	formData := modelo.FormData{
		ID:         id,
		FirstName:  r.FormValue("firstname"),
		LastName:   r.FormValue("lastname"),
		Email:      r.FormValue("email"),
		AddressOne: r.FormValue("address-one"),
		AddressTwo: r.FormValue("address-two"),
		City:       r.FormValue("city"),
		State:      r.FormValue("state"),
		Zip:        r.FormValue("zip"),
		Country:    r.FormValue("country"),
		Telefono:   r.FormValue("telefono"),
		Check:      temp_saved,
	}
	vars := mux.Vars(r)
	metodo, exists := vars["metodo"]
	exisid, id := config.ExitUser(r)
	if exists && exisid {
		if metodo == "add" {
			condi := controlador.Add_Direccion(formData, id)
			if condi {
				http.Redirect(w, r, "/productos/card/buy", http.StatusSeeOther)
			} else {
				Pagina404(w, r)
			}
		} else if metodo == "update" {
			condi := controlador.Update_Diccion(formData)
			if condi {
				http.Redirect(w, r, "/productos/card/buy", http.StatusSeeOther)
			}
		} else {
			Pagina404(w, r)
		}
	} else {
		Pagina404(w, r)
	}

}
func Direccion(w http.ResponseWriter, r *http.Request) {
	template := template.Must(template.ParseFiles("Vista/templates/direccion-form.html", "Vista/Layout/layout-compras.html"))
	data := Card_Produc{}
	exisid, id := config.ExitUser(r)
	if exisid {
		data.UserId = id
	}
	vars := mux.Vars(r)
	ids, exists := vars["id"]
	if exists && exisid {
		if ids == "new" {
			/*condi := controlador.ExistDirec(id)
			if condi {
				data.Mensaje_Error = "Existe"
				template.Execute(w, data)
				return
			}*/
			data.Mensaje_Error = "New-Direc"
			template.Execute(w, data)
			return
		} else if idss, err := strconv.Atoi(ids); err == nil {
			fmt.Println(idss)
			data.Descri2 = controlador.DireccionByID3(idss)
			if data.Descri2.FirstName != "" {
				data.Mensaje_Error = "Edit"
				template.Execute(w, data)
				return
			} else {
				http.Redirect(w, r, "/productos/card/direccion/new", http.StatusSeeOther)
			}
		} else {
			Pagina404(w, r)
			return
		}
	} else {
		Pagina404(w, r)
		return
	}
}

func Confirmar_Direccion(w http.ResponseWriter, r *http.Request) {
	exisid, ids := config.ExitUser(r)
	if exisid {
		vars := mux.Vars(r)
		id, exists := vars["id"]
		if exists {
			if idss, err := strconv.Atoi(id); err == nil {
				condi := controlador.ConfirmarDireccion(idss, ids)
				if condi {
					config.CrearMensajesFlash(w, r, "update")
					http.Redirect(w, r, "/productos/card/buy", http.StatusSeeOther)
					return
				}
			} else {
				Pagina404(w, r)
			}
		}
	} else {
		Pagina404(w, r)
	}
	return
}
func Delete_Direccion(w http.ResponseWriter, r *http.Request) {
	exisid, _ := config.ExitUser(r)
	if exisid {
		vars := mux.Vars(r)
		id, exists := vars["id"]
		if exists {
			if idss, err := strconv.Atoi(id); err == nil {
				condi := controlador.DeleteDirecion(idss)
				if condi {
					config.CrearMensajesFlash(w, r, "delete")
					http.Redirect(w, r, "/productos/card/buy", http.StatusSeeOther)
					return
				}
			} else {
				Pagina404(w, r)
			}
		}
	} else {
		Pagina404(w, r)
	}
	return
}
