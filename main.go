package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

func handler(w http.ResponseWriter, r *http.Request) {
	defer func(start time.Time) {
		log.Printf("/ handler. status: ok. duration: %v", time.Now().Sub(start))
	}(time.Now())

	fmt.Fprintf(w, "hello world")
}

func dbHandler(w http.ResponseWriter, r *http.Request) {
	defer func(start time.Time) {
		log.Printf("/db handler. status: ok. duration: %v", time.Now().Sub(start))
	}(time.Now())

	host := os.Getenv("GUINEAPIG_DB_SERVICE_HOST")
	user := os.Getenv("SECRET_USER")
	password := os.Getenv("SECRET_PASSWORD")
	dbName := os.Getenv("SECRET_DBNAME")
	connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", host, user, password, dbName)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT * FROM test LIMIT 1;")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()

	var out []string
	for rows.Next() {
		var n string
		if err := rows.Scan(&n); err != nil {
			log.Println(err)
		}

		out = append(out, n)
	}

	fmt.Fprintf(w, out[0])
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/db", dbHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
