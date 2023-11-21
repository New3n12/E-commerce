package controlador

import (
	"Go-Proyect/modelo"
	"fmt"
)

func AllCards(id int) []modelo.Produc_Card {
	produc := []modelo.Produc_Card{}
	if modelo.Conectar() {
		produc = modelo.AllCard(modelo.Con, id)
		modelo.Close()
		return produc
	}
	return nil
}
func Add_Card(idProduct int, idUser int, quantity int) (bool, string, int) {
	if modelo.Conectar() {
		cond, msg, id := modelo.Agregar_Card(modelo.Con, idProduct, idUser, quantity)
		if cond {
			modelo.Close()
			return true, msg, id
		}
		modelo.Close()
	}
	return false, "", 0
}

func Delete_Card(id int) bool {
	if modelo.Conectar() {
		if modelo.Delete_Card(modelo.Con, id) {
			modelo.Close()
			return true
		}
		modelo.Close()
	}
	return false
}

func More_Card(id int) bool {
	if modelo.Conectar() {
		if modelo.More_Card(modelo.Con, id) {
			modelo.Close()
			return true
		}
		modelo.Close()
	}
	return false
}

func Less_Card(id int) (bool, string) {
	if modelo.Conectar() {
		condi, mesg := modelo.Less_Card(modelo.Con, id)
		if condi {
			modelo.Close()
			return true, mesg
		}
		modelo.Close()
	}
	return false, ""
}
func Existe_Card(id int) bool {
	if modelo.Conectar() {
		condi := modelo.Existe_Card(modelo.Con, id)
		fmt.Println(condi)
		if condi {
			modelo.Close()
			return true
		}
		modelo.Close()
	}
	return false
}
