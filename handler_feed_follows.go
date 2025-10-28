package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/Adharsh112/Web-Scraper/internal/database"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handlerCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {

	type parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	decoder := json.NewDecoder(r.Body)

	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing JSON : %s", err))
		return
	}
	feed, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't create feed follow : %s", err))
		return
	}
	respondWithJSON(w, 201, databaseFeedFollowToFeedFollow(feed))

}

func (apicfg *apiConfig) handlerGetFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feedFollows, err := apicfg.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't your following feeds : %s", err))
	}

	respondWithJSON(w, 201, databaseFeedFollowsToFeedFollows(feedFollows))

}

func (apicfg *apiConfig) handlerDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	FeedFollowIDStr := chi.URLParam(r, "feedFollowID")
	FeedFollowID, err := uuid.Parse(FeedFollowIDStr)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't parse feed follow id : %v", err))
		return
	}

	err = apicfg.DB.DeleteFeedFollow(r.Context(), database.DeleteFeedFollowParams{
		ID:     FeedFollowID,
		UserID: user.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't delete feed follow id : %v", err))
		return
	}
	respondWithJSON(w, 200, struct{}{})

}
