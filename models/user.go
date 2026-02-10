package models

import "fmt"

type User struct {
	ID int
	Name string
	Email string
	password string //privado encapsulado
}	

//Setter de contraseña con validación
func (u *User) SetPassword(pwd string) error {
	if len(pwd) < 8 {
		return fmt.Errorf("la contraseña debe tener al menos 8 caracteres")
	}
	u.password = pwd
	return nil
}

//Getter validando la contraseña
func (u *User) CheckPassword(pwd string) bool {
	return u.password == pwd
}
