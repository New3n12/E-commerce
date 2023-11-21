package modelo

import (
	"database/sql"
	"fmt"
	"log"
)

func Add_Direccion(db *sql.DB, formData FormData, id int) bool {

	stmt, err := db.Prepare("INSERT INTO informacionenvio (id_usuario,nombre, apellido, correo_electronico, direccion_envio_linea1, direccion_envio_linea2, ciudad_envio, estado_envio, codigo_postal_envio, pais_envio, telefono,saved) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,?)")
	if err != nil {
		log.Fatal(err)
		return false
	}
	defer stmt.Close()
	_, err = stmt.Exec(id, formData.FirstName, formData.LastName, formData.Email, formData.AddressOne, formData.AddressTwo, formData.City, formData.State, formData.Zip, formData.Country, formData.Telefono, formData.Check)
	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func Update_Direccion(db *sql.DB, formData FormData) error {
	query := `UPDATE informacionenvio 
		SET nombre = ?, apellido = ?, correo_electronico = ?, direccion_envio_linea1 = ?, 
			direccion_envio_linea2 = ?, ciudad_envio = ?, estado_envio = ?, 
			codigo_postal_envio = ?, pais_envio = ?, telefono = ? ,saved = ?
		WHERE id = ?`

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(formData.FirstName, formData.LastName, formData.Email, formData.AddressOne,
		formData.AddressTwo, formData.City, formData.State, formData.Zip,
		formData.Country, formData.Telefono, formData.Check, formData.ID)
	if err != nil {
		return err
	}

	return nil
}

func GetDireccionByID(db *sql.DB, id int) ([]FormData, error) {
	var addresses []FormData

	query := "SELECT id, nombre, apellido, correo_electronico, direccion_envio_linea1, direccion_envio_linea2, ciudad_envio, estado_envio, codigo_postal_envio, pais_envio, telefono, saved FROM informacionenvio WHERE id_usuario = ?"
	stmt, err := db.Prepare(query)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(id)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var address FormData
		if err := rows.Scan(&address.ID, &address.FirstName, &address.LastName, &address.Email, &address.AddressOne, &address.AddressTwo, &address.City, &address.State, &address.Zip, &address.Country, &address.Telefono, &address.Check); err != nil {
			fmt.Println(err)
			return nil, err
		}
		addresses = append(addresses, address)
	}

	if err := rows.Err(); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return addresses, nil
}
func GetDireccionByID2(db *sql.DB, id int) (FormData, error) {
	var address FormData

	query := `SELECT id, nombre, apellido, correo_electronico, direccion_envio_linea1, 
	direccion_envio_linea2, ciudad_envio, estado_envio, codigo_postal_envio, pais_envio, telefono, saved
	FROM informacionenvio
	WHERE id_usuario = ?
	ORDER BY id DESC 
	LIMIT 1
	`
	stmt, err := db.Prepare(query)
	if err != nil {
		fmt.Println(err)
		return address, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(&address.ID, &address.FirstName, &address.LastName, &address.Email, &address.AddressOne, &address.AddressTwo, &address.City, &address.State, &address.Zip, &address.Country, &address.Telefono, &address.Check)
	if err != nil {
		fmt.Println(err)
		return address, err
	}

	return address, nil
}
func GetDireccionByID3(db *sql.DB, id int) (FormData, error) {
	var address FormData

	query := `SELECT id, nombre, apellido, correo_electronico, direccion_envio_linea1, 
	direccion_envio_linea2, ciudad_envio, estado_envio, codigo_postal_envio, pais_envio, telefono, saved
	FROM informacionenvio
	WHERE id = ?
	`
	stmt, err := db.Prepare(query)
	if err != nil {
		fmt.Println(err)
		return address, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(id).Scan(&address.ID, &address.FirstName, &address.LastName, &address.Email, &address.AddressOne, &address.AddressTwo, &address.City, &address.State, &address.Zip, &address.Country, &address.Telefono, &address.Check)
	if err != nil {
		fmt.Println(err)
		return address, err
	}

	return address, nil
}

func EliminarDireccion(db *sql.DB, id int) error {
	query := `DELETE FROM informacionenvio WHERE id = ?`

	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	return nil
}

func ExistDirec(db *sql.DB, id int) int {
	query := "SELECT count(*) FROM informacionenvio WHERE id_usuario = ?"
	stmt, err := db.Prepare(query)
	if err != nil {
		return 0
	}
	defer stmt.Close()
	var tamño int
	err = stmt.QueryRow(id).Scan(&tamño)
	if err != nil {
		return 0
	}

	return tamño
}

func ActualizarDireccion(db *sql.DB, id int, id_usu int) bool {
	data, _ := GetDireccionByID3(db, id)
	addnew := Add_Direccion(db, data, id_usu)
	if addnew {
		delete := EliminarDireccion(db, id)
		if delete != nil {
			return false
		}
		return true
	}
	return false
}
