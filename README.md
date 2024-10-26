# Repository lib

This lib is build on top of gorm: https://gorm.io/
Its goals are to make easier the entity management with a basic repository design pattern implementation

## Example

````go
package main

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"repositoryLib"
)

// Définition du modèle
type User struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string
	Email string
}

func main() {
	// Initialisation de la base de données SQLite (ou n'importe quel autre SGBD)
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migration automatique pour créer les tables
	db.AutoMigrate(&User{})

	// Création d'un repository pour le modèle User
	userRepo := repositoryLib.NewGormRepository[User](db)

	// Exemple de création d'un nouvel utilisateur
	newUser := User{Name: "John Doe", Email: "johndoe@example.com"}
	if err := userRepo.Create(&newUser); err != nil {
		fmt.Println("Erreur lors de la création de l'utilisateur :", err)
	}

	// Récupérer tous les utilisateurs
	users, err := userRepo.FindAll()
	if err != nil {
		fmt.Println("Erreur lors de la récupération des utilisateurs :", err)
	} else {
		fmt.Println("Liste des utilisateurs :", users)
	}

	// Trouver un utilisateur par ID
	user, err := userRepo.FindOneById(newUser.ID)
	if err != nil {
		fmt.Println("Erreur lors de la récupération de l'utilisateur :", err)
	} else {
		fmt.Println("Utilisateur trouvé :", user)
	}
}
````

## Install

Wait 'til I push, so I have a URL to give
