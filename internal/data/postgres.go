package data

import (
	"database/sql"
	"io/ioutil"
	"os"

	//registrando el controlador de la base de datos
	_ "github.com/lib/pq"
)

func getConnection() (*sql.DB, error) {
	uri := os.Getenv("DATABASE_URI")

	return sql.Open("postgres", uri)
}

func New() *Data {
	once.Do(initDB)

	return data
}

func (data *Data) Close() error {
	//agrega los recursos necesario para cerrar
	return nil
}

func MakeMigration(db *sql.DB) error {
	b, err := ioutil.ReadFile("./database/models.sql")
	if err != nil {
		return err
	}

	rows, err := db.Query(string(b))
	if err != nil {
		return err
	}

	return rows.Close()
}
