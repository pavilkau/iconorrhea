package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/pavilkau/iconorrhea/internal/files"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		panic(err)
	}

	initStore()
	err = saveFiles()
	if err != nil {
		panic(err)
	}
}

func saveFiles() error {
	files, err := files.Scan("./assets")
	if err != nil {
		return err
	}

	db, err := dbConn()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	query := "INSERT INTO files (name, file, size, mod_time) VALUES ($1, $2, $3, $4)"
	for _, file := range files {
		_, err := db.Exec(query,
			file.Name, file.File, file.Size, file.ModTime)

		if err != nil {
			return err
		}
	}

	return nil
}

func dbConn() (*sql.DB, error) {
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

	return db, nil
}

func initStore() {
	db, err := dbConn()
	if err != nil {
		panic(err)
	}

	initFilesTableString := `CREATE TABLE IF NOT EXISTS files (
				name VARCHAR NOT NULL,
				file BYTEA NOT NULL,
				size INT NOT NULL,
				mod_time TIME,
				seen BOOL DEFAULT FALSE,
				seen_at TIME
			)`
	if _, err := db.Exec(initFilesTableString); err != nil {
		panic(err)
	}
}
