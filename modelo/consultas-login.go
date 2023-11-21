package modelo

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

func CrearUsuario(db *sql.DB, usu User) bool {
	hashedPassword, _ := HashPassword(usu.Contra)
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()

	stmt, err := db.Prepare("INSERT INTO usuarios (nombre, correo, edad, fecha_creacion, contra) VALUES (?,?,?,?,?)")
	if err != nil {
		return false
	}
	defer stmt.Close()

	_, err = stmt.Exec(usu.Nombre, usu.Correo, usu.Edad, usu.FechaCreacion, hashedPassword)
	if err != nil {
		return false
	}

	return true
}

func HashPassword(password string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashed), nil
}

func ReturnUsu(db *sql.DB, correo string) User {
	stmt, err := db.Prepare("SELECT COUNT(*) as tam, id, nombre, correo, edad, fecha_creacion, contra FROM usuarios WHERE correo = ?")
	if err != nil {
		fmt.Println("Error al preparar la consulta:", err)
		return User{}
	}
	defer stmt.Close()

	rows, err := stmt.Query(correo)
	if err != nil {
		fmt.Println("Error al ejecutar la consulta:", err)
		return User{}
	}

	ususario := User{}

	for rows.Next() {
		var tam int
		err := rows.Scan(&tam, &ususario.ID, &ususario.Nombre, &ususario.Correo,
			&ususario.Edad, &ususario.FechaCreacion, &ususario.Contra)
		fmt.Println(tam)

		if err != nil {
			fmt.Println("Error al escanear fila:", err)
		}
		if tam == 0 {
			return ususario
		}
	}
	return ususario
}

func ActualizarUsuario(db *sql.DB, usu User) error {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()
	stmt, err := db.Prepare("UPDATE usuarios SET nombre = ?, correo = ? edad = ? fecha_creacion = ? WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(usu.Nombre, usu.Correo, usu.Edad, usu.FechaCreacion, usu.ID)
	if err != nil {
		return err
	}

	return nil
}

func EliminarUsuario(db *sql.DB, id int) error {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()

	stmt, err := db.Prepare("DELETE FROM usuarios WHERE id = ?")
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

func ReturnIdUsu(db *sql.DB) (bool, int) {
	rows, err := db.Query("SELECT id FROM usuarios ORDER BY id DESC LIMIT 1")
	if err != nil {
		log.Fatal(err)
		return false, 0
	}
	defer rows.Close()

	var id int

	for rows.Next() {

		err := rows.Scan(&id)
		if err != nil {
			log.Println(err)
			return false, 0
		}
	}

	return true, id
}
