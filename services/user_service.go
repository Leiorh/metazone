package services

import "METAZONE/models"

// Lista de usuarios en memoria
var users []models.User

// Crear un nuevo usuario
func CreateUser(name, email, password string) (*models.User, error) {
	user := models.User{
		ID:    len(users) + 1,
		Name:  name,
		Email: email,
	}

	err := user.SetPassword(password)
	if err != nil {
		return nil, err
	}

	users = append(users, user)

	// Retorna el usuario recién agregado al slice
	return &users[len(users)-1], nil
}

// Obtener todos los usuarios
func GetUsers() []models.User {
	return users
}