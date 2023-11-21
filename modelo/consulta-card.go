package modelo

import (
	"database/sql"
	"encoding/base64"
	"fmt"
	"log"
)

func AllCard(db *sql.DB, id int) []Produc_Card {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()

	stmt, err := db.Prepare("SELECT carrito.id_card as id, productos.nombre as  nombre, productos.img1 as img, carrito.cantidad as cantidad, productos.precio as precio " +
		"" + "FROM `carrito` INNER JOIN productos ON carrito.id_producto = productos.id_producto WHERE carrito.id_usuario = ?")

	if err != nil {
		fmt.Println("Error al preparar la consulta:", err)
		return []Produc_Card{}
	}
	defer stmt.Close()

	rows, err := stmt.Query(id)

	if err != nil {
		fmt.Println("Error al ejecutar la consulta:", err)
		return []Produc_Card{}
	}

	var products []Produc_Card
	for rows.Next() {
		var product Produc_Card
		var Img1 []byte
		err := rows.Scan(&product.ID_card, &product.Nombre_card, &Img1, &product.Cantidad_card, &product.Precio_card)
		if err != nil {
			log.Println(err)
			continue
		}
		product.Img_card = base64.StdEncoding.EncodeToString([]byte(Img1))
		products = append(products, product)
	}
	return products
}

func Agregar_Card(db *sql.DB, idProduct int, idUser int, quantity int) (bool, string, int) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()

	stmtExists, err := db.Prepare("SELECT COUNT(*) FROM carrito WHERE id_producto = ? AND id_usuario = ?")

	if err != nil {
		log.Println(err)
		return false, "", 0
	}
	defer stmtExists.Close()

	var count int
	err = stmtExists.QueryRow(idProduct, idUser).Scan(&count)
	if err != nil {
		log.Println(err)
		return false, "", 0
	}

	if count == 0 {
		stmtInsert, err := db.Prepare("INSERT INTO carrito(id_producto, id_usuario, cantidad) VALUES (?, ?, ?)")
		if err != nil {
			log.Println(err)
			return false, "", 0
		}
		defer stmtInsert.Close()

		_, err = stmtInsert.Exec(idProduct, idUser, quantity)
		if err != nil {
			log.Println(err)
			return false, "", 0
		}

		rows, err := db.Query("SELECT id_card FROM carrito ORDER BY id_card DESC LIMIT 1")
		if err != nil {
			log.Fatal(err)
			return false, "", 0
		}
		defer rows.Close()

		var id int

		for rows.Next() {

			err := rows.Scan(&id)
			if err != nil {
				log.Println(err)
				return false, "", 0
			}
		}

		return true, "ok", id
	}

	return false, "exist", 0
}

func Delete_Card(db *sql.DB, id int) bool {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()

	stmt, err := db.Prepare("DELETE FROM carrito WHERE id_card = ?")
	if err != nil {
		log.Println(err)
		return false
	}

	_, err = stmt.Exec(id)
	if err != nil {
		return false
	}

	return true
}

func More_Card(db *sql.DB, id int) bool {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()

	sql := "UPDATE carrito INNER JOIN productos ON carrito.id_producto = productos.id_producto SET " +
		"" + "carrito.cantidad = carrito.cantidad + 1 WHERE carrito.cantidad + 1 <= productos.cantidad AND carrito.id_card = ?"
	stmt, err := db.Prepare(sql)
	if err != nil {
		log.Println(err)
		return false
	}

	result, err := stmt.Exec(id)

	if err != nil {
		return false
	}

	rowsAffected, err := result.RowsAffected()

	if err != nil {
		return false
	}

	if rowsAffected == 0 {
		return false
	}

	return true
}
func Less_Card(db *sql.DB, id int) (bool, string) {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()

	sql := "DELETE FROM carrito WHERE cantidad = 1 AND id_card = ?"
	stmt, err := db.Prepare(sql)

	if err != nil {
		log.Println(err)
		return false, ""
	}

	result, err := stmt.Exec(id)
	if err != nil {
		return false, ""
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return false, ""
	}

	if rowsAffected > 0 {
		return true, "DELETE"
	} else {
		sql = "UPDATE carrito SET cantidad = cantidad - 1 WHERE cantidad >= 2 AND id_card = ?"
		stmt, err = db.Prepare(sql)

		if err != nil {
			log.Println(err)
			return false, ""
		}

		result, err = stmt.Exec(id)
		if err != nil {
			return false, ""
		}
		return true, "UPDATE"
	}
}
func Existe_Card(db *sql.DB, id int) bool {
	stmtExists, err := db.Prepare("SELECT COUNT(*) FROM carrito WHERE id_usuario = ?")

	if err != nil {
		log.Println(err)
		return false
	}
	defer stmtExists.Close()

	var count int

	fmt.Println(count)
	err = stmtExists.QueryRow(id).Scan(&count)
	if err != nil {
		log.Println(err)
		return false
	}
	if count > 0 {
		return true
	}
	return false
}
