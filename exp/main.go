package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "Test_33"
	dbname   = "lenslocked_dev"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	var id int
	row := db.QueryRow(`
    INSERT INTO users(name, email)
    VALUES($1, $2)
    RETURNING id`,
		"Xavier", "smxxcastillo@gmail.com")
	err = row.Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("id is...", id)
}
