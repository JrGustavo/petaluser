package bd

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/petaluser/models"
	"github.com/petaluser/tools"
)

func SignUp(sig models.SignUp) error {
	fmt.Println("Comienza el registro")

	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	sentencia := "INSERT INTO users (User_Email, User_UUIDD, User_DateAdd) VALUES (?, ?, ?)"
	fmt.Println(sentencia)

	// Usa consultas preparadas para evitar inyección de SQL
	_, err = Db.Exec(sentencia, sig.UserEmail, sig.UserUUID, tools.FechaMySQL())
	if err != nil {
		// Manejar el error de manera adecuada
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			fmt.Println("Error MySQL:", mysqlErr.Number, mysqlErr.Message)
		}
		return err
	}

	fmt.Println("Sign Ejecución exitosa")
	return nil
}
