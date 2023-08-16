package main

import (
	"fmt"
	conn "programa3/database"
)

func main() {
	// Crear conexión de base de datos
	conn.Init()

	// Finalizar la conexión
	defer func() {
		if err := conn.Close(conn.GetConn()); err != nil {
			fmt.Println(err)
		}
	}()
}
