package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Numostanley/rssagg/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		Name string `json:"name"`
	}

	decoder := json.NewDecoder(r.Body)

	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON: %v \n", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error creating user: %v \n", err))
		return
	}

	respondWithJSON(w, 200, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handlerGetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := apiCfg.DB.GetAllUsers(r.Context())

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error getting all user: %v \n", err))
		return
	}

	respondWithJSON(w, 200, databaseUsersToUsers(users))
}

func (apiCfg *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	userId, err := uuid.Parse(id)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing uuid: %v \n", err))
		return
	}

	user, err := apiCfg.DB.GetUserByID(r.Context(), userId)

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error getting all user: %v \n", err))
		return
	}

	respondWithJSON(w, 200, databaseUserToUser(user))
}
