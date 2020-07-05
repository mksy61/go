package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "postgres://learngo:2020@192.168.2.23/portal?sslmode=disable")
	if err != nil {
		log.Panicln("Error:", err)
	}
	defer db.Close()

	rows, err := db.Query("select * from animals where age > $1", 1)
	if err != nil {
		log.Panicln("Error:", err)
	}

	handleRows(rows, err)
	fmt.Println()

	/*	Imsert
		resp, err := db.Exec("insert into animals (animaltype, nickname, zone, age) values ('Mezovor', 'mez', 2, 3)")
		if err != nil {
			log.Panicln("Error:", err)
		}

		nrowseffected, err := resp.RowsAffected()
		fmt.Println(nrowseffected)
	*/
	//Prepared statement
	/*
		prpstmt, err := db.Prepare("select * from animals where age > $1")
		if err != nil {
			log.Panicln("Error:", err)
		}
		defer prpstmt.Close()

		rows, err = prpstmt.Query(10)
		handleRows(rows, err)
		fmt.Println()

		rows, err = prpstmt.Query(15)
		handleRows(rows, err)
		fmt.Println()
	*/
	testTransaction(db)
}

type animal struct {
	id         int
	animaltype string
	nickname   string
	zone       int
	age        int
}

func handleRows(rows *sql.Rows, err error) {
	if err != nil {
		rows.Close()
		log.Panicln("Error:", err)
	}

	animals := []animal{}
	for rows.Next() {
		a := animal{}
		err := rows.Scan(&a.id, &a.animaltype, &a.nickname, &a.zone, &a.age)
		if err != nil {
			log.Println(err)
			continue
		}
		animals = append(animals, a)
	}

	if err := rows.Err(); err != nil {
		rows.Close()
		log.Panicln("Error:", err)
	}

	for _, v := range animals {
		fmt.Printf("%+v\n", v)
	}

	defer rows.Close()
}

func testTransaction(db *sql.DB) {

	err := db.Ping()
	if err != nil {
		log.Panicln("Error:", err)
	}

	fmt.Println("transaction begins")
	tx, err := db.Begin()
	if err != nil {
		log.Panicln("Error:", err)
	}

	prpstmt, err := tx.Prepare("select * from animals where age > $1")
	if err != nil {
		log.Panicln("Error:", err)
	}
	defer prpstmt.Close()

	rows, err := prpstmt.Query(10)
	handleRows(rows, err)
	fmt.Println()

	rows, err = prpstmt.Query(15)
	handleRows(rows, err)
	fmt.Println()

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		log.Panicln("Error:", err)
	}
}
