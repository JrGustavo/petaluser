package bd

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/petaluser/models"
	"github.com/petaluser/tools"
)

func SignUp(sig models.Signup) error {
	fmt.Println("Comienza el registro")

	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	sentencia := "INSERT INTO users (User_Email, User_UUIDD, User_DateAdd) VALUES ('" + sig.UserEmail + "','" + sig.UserUUID + "','" + tools.FechaMySQL() + "')"
	fmt.Println(sentencia)

	_, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Sign Ejecución exitosa")
	return nil
}
