package controlador

import (
	"Go-Proyect/config"
	"Go-Proyect/modelo"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func Registro(usu modelo.User, r *http.Request, w http.ResponseWriter) bool {
	if modelo.Conectar() {
		err := modelo.CrearUsuario(modelo.Con, usu)
		condi, id := modelo.ReturnIdUsu(modelo.Con)
		if err && condi {
			session, _ := config.Store.Get(r, "session-name")
			session.Values["usuarioId"] = id
			session.Save(r, w)
			modelo.Close()
			return true
		}
		modelo.Close()
	}
	return false
}

func Validar(correo string, pass string, r *http.Request, w http.ResponseWriter) bool {
	if modelo.Conectar() {
		usu := modelo.ReturnUsu(modelo.Con, correo)
		modelo.Close()
		if usu.Contra != "" {
			errpass := bcrypt.CompareHashAndPassword([]byte(usu.Contra), []byte(pass))
			if errpass != nil {
				return false
			} else {
				session, _ := config.Store.Get(r, "session-name")
				session.Values["usuarioId"] = usu.ID
				session.Save(r, w)
				return true
			}
		}
	}
	return false
}
