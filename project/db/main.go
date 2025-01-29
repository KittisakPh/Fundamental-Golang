package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/proullon/ramsql/driver"
)

func main() {
	db, err := sql.Open("ramsql", "internet-movie")
	if err != nil {
		log.Fatal("error", err)
	}
	defer db.Close()

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

	_, err = db.Exec(createTb)
	if err != nil {
		log.Fatal("error", err)
	}
	fmt.Println("table created.")

	insertTb := `
	INSERT INTO goimdb(imdbID,title,year,rating,isSuperHero)
	VALUES ($1,$2,$3,$4,$5);
	`

	stmt, err := db.Prepare(insertTb)
	if err != nil {
		log.Fatal("prepare statement error:", err)
	}
	r, err := stmt.Exec("tt4154796", "Avenger", 2019, 8.4, true)
	if err != nil {
		log.Fatal("error", err)
	}
	l, err := r.LastInsertId()
	fmt.Println("lastINsertId", l, "err:", err)
	ef, err := r.RowsAffected()
	fmt.Println("RowsAffected", ef, "err:", err)

	rows, err := db.Query(`
	SELECT id, imdbID, title, year, rating 
	FROM goimdb
	`)
	if err != nil {
		log.Fatal("query err", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var imdbID, title string
		var year int
		var rating float32
		if err := rows.Scan(&id, &imdbID, &title, &year, &rating); err != nil {
			log.Fatal("for rows error", err)
		}
		fmt.Println("row", id, imdbID, title, year, rating)
	}

	stmt2, _ := db.Prepare(`UPDATE goimdb SET rating = $1 WHERE imdbID = $2`)
	_, err = stmt2.Exec(9.2, "tt4154796")
	if err != nil {
		log.Fatal("update error", err)
	}
	var id int
	var imdbID, title string
	var year int
	var rating float32
	var isSuperHero bool
	rowx := db.QueryRow(`SELECT id, imdbID, title, year, rating, isSuperHero 
	FROM goimdb WHERE imdbID=$1
	`, "tt4154796")
	if err := rowx.Scan(&id, &imdbID, &title, &year, &rating, &isSuperHero); err != nil {
		log.Fatal(err)
	}
	fmt.Println("one rowx", id, imdbID, title, year, rating, isSuperHero)
}