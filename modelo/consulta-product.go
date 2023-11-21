package modelo

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"log"
)

func AllProducts(db *sql.DB) []Product {
	query := "SELECT id_producto, id_subcategoria , nombre, precio , img1 , nuevo FROM productos"
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer rows.Close()
	var products []Product
	for rows.Next() {
		var product Product
		var Img1 []byte
		err := rows.Scan(&product.ID, &product.ID_Cat, &product.Nombre, &product.Precio, &Img1, &product.Isnew)
		if err != nil {
			log.Println(err)
			continue
		}
		product.Img1 = base64.StdEncoding.EncodeToString([]byte(Img1))
		products = append(products, product)
	}
	return products
}

func IdProduct(ids int, db *sql.DB) Product {
	stmt, err := db.Prepare("SELECT id_producto, nombre, descripcion, precio, cantidad, color, img1, img2, img3 FROM productos WHERE id_producto = ?")

	if err != nil {
		fmt.Println("Error al preparar la consulta:", err)
		return Product{}
	}
	defer stmt.Close()

	rows, err := stmt.Query(ids)
	if err != nil {
		fmt.Println("Error al ejecutar la consulta:", err)
		return Product{}
	}

	var product Product

	var Img1 []byte
	var Img2 []byte
	var Img3 []byte
	var Cantidad int
	for rows.Next() {
		err := rows.Scan(&product.ID, &product.Nombre, &product.Descripcion, &product.Precio,
			&Cantidad, &product.Color, &Img1, &Img2, &Img3)
		if err != nil {
			log.Println(err)
			continue
		}
		product.Img1 = base64.StdEncoding.EncodeToString([]byte(Img1))
		product.Img2 = base64.StdEncoding.EncodeToString([]byte(Img2))
		product.Img3 = base64.StdEncoding.EncodeToString([]byte(Img3))

	}
	product.Cantidad = ciclo(Cantidad)
	return product
}
func Filtrar_Productos(db *sql.DB, data []string) []Product {
	var sql string = ""
	for i, item := range data {
		if item != "Todos" {
			sql += "categorias.nombre_categoria = '" + item + "'"
		}
		if len(data)-1 != i && item != "Todos" {
			sql += " or "
		}
	}
	query := ""
	//fmt.Println(len(sql))
	if len(sql) != 0 {
		query = "SELECT id_producto,id_subcategoria, nombre, precio ,img1,nuevo FROM productos INNER JOIN categorias ON productos.id_subcategoria = categorias.id_categoria WHERE " + sql

	} else {
		query = "SELECT id_producto,id_subcategoria, nombre, precio ,img1,nuevo FROM productos"
	}
	//fmt.Println(query)
	rows, err := db.Query(query)
	if err != nil {
		log.Println(err)
		return nil
	}
	defer rows.Close()
	var products []Product
	for rows.Next() {
		var product Product
		var Img1 []byte
		err := rows.Scan(&product.ID, &product.ID_Cat, &product.Nombre, &product.Precio, &Img1, &product.Isnew)
		if err != nil {
			log.Println(err)
			continue
		}
		product.Img1 = base64.StdEncoding.EncodeToString([]byte(Img1))
		products = append(products, product)
	}
	return products
}

func ProductLimit(ids int, db *sql.DB) []Product {
	stmt, err := db.Prepare("SELECT id_producto,id_subcategoria, nombre, precio, img1, nuevo FROM productos	WHERE id_subcategoria = ? ORDER BY RAND() LIMIT 4")

	if err != nil {
		fmt.Println("Error al preparar la consulta:", err)
		return []Product{}
	}
	defer stmt.Close()

	rows, err := stmt.Query(ids)
	if err != nil {
		fmt.Println("Error al ejecutar la consulta:", err)
		return []Product{}
	}

	var products []Product
	for rows.Next() {
		var product Product
		var Img1 []byte
		err := rows.Scan(&product.ID, &product.ID_Cat, &product.Nombre, &product.Precio, &Img1, &product.Isnew)
		if err != nil {
			log.Println(err)
			continue
		}
		product.Img1 = base64.StdEncoding.EncodeToString([]byte(Img1))
		products = append(products, product)
	}
	return products
}

func ProductSuger(ids int, idcat int, db *sql.DB) []Product {
	stmt, err := db.Prepare(`SELECT id_producto, id_subcategoria, nombre, precio, img1, nuevo
	FROM productos
	WHERE id_producto != ? AND id_subcategoria = ?  
	ORDER BY RAND()
	LIMIT 4	
	`)

	if err != nil {
		fmt.Println("Error al preparar la consulta:", err)
		return []Product{}
	}
	defer stmt.Close()

	rows, err := stmt.Query(ids, idcat)
	if err != nil {
		fmt.Println("Error al ejecutar la consulta:", err)
		return []Product{}
	}

	var products []Product

	for rows.Next() {
		var product Product
		var Img1 []byte
		err := rows.Scan(&product.ID, &product.ID_Cat, &product.Nombre, &product.Precio, &Img1, &product.Isnew)
		if err != nil {
			log.Println(err)
			continue
		}
		product.Img1 = base64.StdEncoding.EncodeToString([]byte(Img1))
		products = append(products, product)
	}
	return products
}

func ciclo(n int) []int {

	slice := make([]int, n)

	for i := 0; i < n; i++ {
		slice[i] = i + 1
	}

	return slice

}
