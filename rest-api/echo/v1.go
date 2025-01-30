package main

import (
	"log"
	"net/http"
	"strconv"
	"github.com/labstack/echo/v4"
)

type Movie struct {
	ImdbID      string  `json:"imdbID"`
	Title       string  `json:"title"`
	Year        int     `json:"year"`
	Rating      float64 `json:"rating"`
	IsSuperHero bool    `json:"is_super_hero"`
}

var moviesV1 = []Movie{
	{
		ImdbID:      "tt4154796",
		Title:       "Avengers: Endgame",
		Year:        2019,
		Rating:      8.4,
		IsSuperHero: true,
	},
}

func getAllMoviesHandlerV1(c echo.Context) error {
	y := c.QueryParam("year")
	if y == "" {
		return c.JSON(http.StatusOK, moviesV1)
	}
	year, err := strconv.Atoi(y)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	ms := []Movie{}
	for _, m := range moviesV1 {
		if m.Year == year {
			ms = append(ms, m)
		}
	}
	return c.JSON(http.StatusOK, ms)
}

func getMovieByIdHandlerV1(c echo.Context) error {
	for _, v := range moviesV1 {
		if v.ImdbID == c.Param("id") {
			return c.JSON(http.StatusOK, v)
		}
	}

	return c.JSON(http.StatusOK, map[string]string{
		"message": "not found",
	})
}

func createMovieHandlerV1(c echo.Context) error {
	m := new(Movie)
	if err := c.Bind(m); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	moviesV1 = append(moviesV1, *m)
	return c.JSON(http.StatusCreated, map[string]string{
		"message": "created success!",
	})
}

func mainV1() {
	e := echo.New()
	e.GET("/movies", getAllMoviesHandlerV1)
	e.GET("/movies/:id", getMovieByIdHandlerV1)
	e.POST("/movies", createMovieHandlerV1)
	port := "2565"
	log.Println("starting... port:", port)
	log.Fatal(e.Start(":" + port))
}