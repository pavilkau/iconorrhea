package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/pavilkau/iconorrhea/internal/files"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := initStore()
	if err != nil {
		log.Fatalf("failed to initialise the store: %s", err)
	}
	defer db.Close()

	files, err := files.Scan("./assets")
	if err != nil {
		log.Fatalf("failed to scan files: %s", err)
	}

	for _, file := range files {
		_, err := db.Exec("INSERT INTO file (name, size, mod_time) VALUES ($1, $2, $3)",
			file.Name, file.Size, file.ModTime)
		if err != nil {
			log.Fatalf("failed insert files: %s", err)
			return
		}
	}

}

func initStore() (*sql.DB, error) {

	pgConnString := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
	)

	db, err := sql.Open("postgres", pgConnString)

	if err != nil {
		return nil, err
	}

	initTableString := `CREATE TABLE IF NOT EXISTS file (
				name VARCHAR NOT NULL,
				size INT NOT NULL,
				mod_time TIMESTAMP NOT NULL
			)`
	if _, err := db.Exec(initTableString); err != nil {
		return nil, err
	}

	return db, nil
}
