package controlador

import (
	"Go-Proyect/modelo"
)

func All_Products() []modelo.Product {
	produc := []modelo.Product{}
	if modelo.Conectar() {
		produc = modelo.AllProducts(modelo.Con)
		modelo.Close()
		return produc
	}

	return nil
}
func Id_Product(ids int) modelo.Product {
	produc := modelo.Product{}
	if modelo.Conectar() {
		produc = modelo.IdProduct(ids, modelo.Con)
		modelo.Close()
		return produc
	}

	return produc
}

func Filtro_Produc(sql []string) ([]modelo.Product, string) {
	produc := []modelo.Product{}
	if modelo.Conectar() {
		produc = modelo.Filtrar_Productos(modelo.Con, sql)
		if produc == nil {
			produc = All_Products()
			modelo.Close()
			return produc, "all"
		} else {
			modelo.Close()
			return produc, "ok"
		}
	}
	return nil, "bad"
}
func PrductosLimit4(id int) []modelo.Product {
	produc := []modelo.Product{}
	if modelo.Conectar() {
		produc = modelo.ProductLimit(id, modelo.Con)
		modelo.Close()
		return produc
	}
	return nil
}

func PrductosSugerencias(id int, idcat int) []modelo.Product {
	produc := []modelo.Product{}
	if modelo.Conectar() {
		produc = modelo.ProductSuger(id, idcat, modelo.Con)
		modelo.Close()
		return produc
	}
	return nil
}
