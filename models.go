package main

import (
	"time"

	"github.com/Numostanley/rssagg/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
}

func databaseUserToUser(dbUser database.User) User {
	return User{
		ID:        dbUser.ID,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Name:      dbUser.Name,
	}
}

func databaseUsersToUsers(dbUsers []database.User) []User {

	dbUserList := []User{}

	for _, dbUser := range dbUsers {
		user := User{
			ID:        dbUser.ID,
			CreatedAt: dbUser.CreatedAt,
			UpdatedAt: dbUser.UpdatedAt,
			Name:      dbUser.Name,
		}
		dbUserList = append(dbUserList, user)
	}
	return dbUserList
}
