package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func main() {

	// Connect to a database
	conn, err := sql.Open("pgx", "database_url")
	if err != nil {
		log.Fatal("it was not possible to connect to the database", err)
	}

	defer conn.Close()

	log.Println("successful connection to the database")

	// Testing connection
	err = conn.Ping()
	if err != nil {
		log.Fatal("cannot ping database! ", err)
	}
	log.Println("Pinged database!")

	// Getting rows from database
	err = getAllRows(conn)
	if err != nil {
		log.Fatal(err)
	}

	// Insert a row
	query := `insert into users (first_name, last_name) values ($1, $2)` 
	_, err = conn.Exec(query, "Jack", "Brown")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Inserted a row successfully!")

	// Getting rows from database
	err = getAllRows(conn)
	if err != nil {
		log.Fatal(err)
	}

	// Update a row
	stmt := `update users set first_name = $1 where first_name = $2`
	_, err = conn.Exec(stmt, "Jackie", "Jack")
	if err != nil {
		log.Fatal(err)
	}
	log.Println("updated a row successfully")

	// Getting rows from database
	err = getAllRows(conn)
	if err != nil {
		log.Fatal(err)
	}

	// Get row by id
	query = `select id, first_name, last_name from users where id = $1`
	var firstName, lastName string
	var id int
	row := conn.QueryRow(query, 1)
	err = row.Scan(&id, &firstName, &lastName)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(id, firstName, lastName)

	// Delete a row
	query = `delete from users where id = $1`
	_, err = conn.Exec(query, 6)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("row deleted successfully")

	// Getting rows from database
	err = getAllRows(conn)
	if err != nil {
		log.Fatal(err)
	}

}

func getAllRows(conn *sql.DB) error {
	rows, err := conn.Query("select id, first_name, last_name from users")
	if err != nil {
		log.Println(err)
		return err
	}
	defer rows.Close()
	
	var firstName, lastName string
	var id int

	for rows.Next() {
		err := rows.Scan(&id, &firstName, &lastName)
		if err != nil {
			log.Println(err)
			return err
		}

		fmt.Println(id, firstName, lastName)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return nil
}