package controlador

import (
	"Go-Proyect/modelo"
)

func Add_Direccion(data modelo.FormData, id int) bool {
	if modelo.Conectar() {
		condi := modelo.Add_Direccion(modelo.Con, data, id)
		if condi {
			modelo.Close()
			return true
		}
		modelo.Close()
	}
	return false
}
func Update_Diccion(data modelo.FormData) bool {
	if modelo.Conectar() {
		condi := modelo.Update_Direccion(modelo.Con, data)
		if condi == nil {
			modelo.Close()
			return true
		}
		modelo.Close()
	}
	return false
}

func DireccionByID(id int) []modelo.FormData {
	var data = []modelo.FormData{}
	if modelo.Conectar() {
		res, err := modelo.GetDireccionByID(modelo.Con, id)
		if err == nil {
			data = res
			modelo.Close()
			return data
		}
		modelo.Close()
	}
	return data
}
func DireccionByID2(id int) modelo.FormData {
	var data = modelo.FormData{}
	if modelo.Conectar() {
		res, err := modelo.GetDireccionByID2(modelo.Con, id)
		if err == nil {
			data = res
			modelo.Close()
			return data
		}
		modelo.Close()
	}
	return data
}
func DireccionByID3(id int) modelo.FormData {
	var data = modelo.FormData{}
	if modelo.Conectar() {
		res, err := modelo.GetDireccionByID3(modelo.Con, id)
		if err == nil {
			data = res
			modelo.Close()
			return data
		}
		modelo.Close()
	}
	return data
}

func ConfirmarDireccion(id int, id_usu int) bool {
	if modelo.Conectar() {
		condi := modelo.ActualizarDireccion(modelo.Con, id, id_usu)
		if condi {
			modelo.Close()
			return true
		}
	}
	modelo.Close()
	return false
}
func DeleteDirecion(id int) bool {
	if modelo.Conectar() {
		condi := modelo.EliminarDireccion(modelo.Con, id)
		if condi == nil {
			modelo.Close()
			return true
		}
	}
	modelo.Close()
	return false
}

func ExistDirec(id int) (bool, int) {
	if modelo.Conectar() {
		condi := modelo.ExistDirec(modelo.Con, id)
		if condi > 0 {
			modelo.Close()
			return true, condi
		}
		modelo.Close()
	}
	return false, 0
}
