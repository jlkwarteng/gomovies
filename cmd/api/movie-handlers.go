package main

import (
	"backend/models"
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
)

func (app *application) getOneMovie(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.Atoi(params.ByName("id"))

	if err != nil {
		app.logger.Println(errors.New("Invalid Id Param"))
		app.errorJSON(w, errors.New("Invalid Id Param"))
		return
	}
	app.logger.Println("Id is ", id)

	movie := models.Movie{
		ID:          id,
		Title:       "My Movie Title",
		Description: "My Mov Description",
		Year:        2022,
		ReleaseDate: time.Now(),
		Runtime:     100,
		Rating:      5,
		MPAARating:  "PG-13",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	err = app.writeJSON(w, http.StatusOK, movie, "movie")
}

func (app *application) getAllMovies(w http.ResponseWriter, r *http.Request) {

}
