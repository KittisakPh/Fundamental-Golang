package main

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/echo/v4"
	_ "github.com/proullon/ramsql/driver"
)

type MovieApi struct {
	ID          int64   `json:"id"`
	ImdbID      string  `json:"imdbID"`
	Title       string  `json:"title"`
	Year        int     `json:"year"`
	Rating      float64 `json:"rating"`
	IsSuperHero bool    `json:"isSuperHero"`
}

func getAllMoviesHandler(c echo.Context) error {
	m := MovieApi{}
	ms := []MovieApi{}
	y := c.QueryParam("year")
	if y == "" {
		rows, err := db.Query(`
	SELECT id,imdbID,title,year,rating,isSuperHero
	FROM goimdb
		`)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, "query:"+err.Error())
		}
		defer rows.Close()
		for rows.Next() {
			if err := rows.Scan(&m.ID, &m.ImdbID, &m.Title, &m.Year, &m.Rating, &m.IsSuperHero); err != nil {
				return c.JSON(http.StatusInternalServerError, "scan:"+err.Error())
			}
			ms = append(ms, m)
		}
		if err := rows.Err(); err != nil{
			return c.JSON(http.StatusInternalServerError, "rows:"+err.Error())
		}
		return c.JSON(http.StatusOK, ms)
	}

	year, err := strconv.Atoi(y)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	rows, err := db.Query(`
	SELECT id,imdbID,title,year,rating,isSuperHero
	FROM goimdb
	WHERE year=$1
	`, year)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&m.ID, &m.ImdbID, &m.Title, &m.Year, &m.Rating, &m.IsSuperHero); err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		ms = append(ms, m)
	}
	if err := rows.Err(); err != nil{
		return c.JSON(http.StatusInternalServerError, "rows:"+err.Error())
	}
	return c.JSON(http.StatusOK, ms)
}

func getMoviesByIdHandler(c echo.Context) error {
	m := MovieApi{}
	imdbID := c.QueryParam("imdbID")
	rowx := db.QueryRow(`
	SELECT id,imdbID,title,year,rating,isSuperHero
	FROM goimdb
	WHERE imdbID = @1
	`, imdbID)
	err := rowx.Scan(&m.ID, &m.ImdbID, &m.Title, &m.Year, &m.Rating, &m.IsSuperHero)
	switch err {
	case nil:
		return c.JSON(http.StatusOK, m)
	case sql.ErrNoRows:
		return c.JSON(http.StatusNotFound, map[string]string{"message":"movie not found"})
	default:
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
}

func createMoviesHandler(c echo.Context) error {
	m := new(MovieApi)
	if err := c.Bind(m); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	stmt, err := db.Prepare(`
	INSERT INTO goimdb (imdbID,title,year,rating,isSuperHero)
	VALUES ($1,$2,$3,$4,$5)
	`)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	defer stmt.Close()

	r, err := stmt.Exec(m.ImdbID, m.Title, m.Year, m.Rating, m.IsSuperHero)
	if err != nil {
	}
	switch {
	case err == nil:
		rowsAfftected, err := r.RowsAffected()
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		if rowsAfftected < 1 {
			return c.JSON(http.StatusInternalServerError, "create failed")
		}
		id, _ := r.LastInsertId()
		m.ID = id
		return c.JSON(http.StatusCreated, m)
	case err.Error() == "UNIQUE constraint violation":
		return c.JSON(http.StatusConflict, "movie already exists")
	default:
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
}

var db *sql.DB

func conn() {
	var err error
	db, err = sql.Open("ramsql", "goimdb")
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	conn()
	createTb := `
	CREATE TABLE IF NOT EXISTS goimdb (
	id INT AUTO_INCREMENT,
	imdbID TEXT NOT NULL UNIQUE,
	title TEXT NOT NULL,
	year INT NOT NULL,
	rating FLOAT NOT NULL,
	isSuperHero BOOLEAN NOT NULL,
	PRIMARY KEY (id)
	);
	`
	if _, err := db.Exec(createTb); err != nil {
		log.Fatal("create table error", err)
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.GET("/movies", getAllMoviesHandler)
	e.GET("/movies/:imdbID", getMoviesByIdHandler)
	e.POST("/movies", createMoviesHandler)

	port := "2565"
	log.Fatal(e.Start(":" + port))
}