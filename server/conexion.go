package server

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func conexionBD() (conexion *sql.DB) {
	Driver := "mysql"
	User := "root"
	Key := ""
	Name := "restaurantego"
	Host := "127.0.0.1"

	conexion, err := sql.Open(Driver, User+":"+Key+"@tcp("+Host+")/"+Name)
	Debug(err)
	return conexion
}

func Debug(err error) {
	if err != nil {
		panic(err.Error())
	}
}

func CloseConexion(conexion *sql.DB) {
	conexion.Close()
}
