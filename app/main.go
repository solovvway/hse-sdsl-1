package main

import (
	"database/sql"
	_ "embed"
	"fmt"
	"log"
	"regexp"

	"gopkg.in/yaml.v3"

	_ "github.com/lib/pq"
)

//go:embed config.yaml
var config_file []byte

type Config struct {
	DB struct {
		POSTGRES_HOST     string `yaml:"POSTGRES_HOST"`
		POSTGRES_DB       string `yaml:"POSTGRES_DB"`
		POSTGRES_USERNAME string
		POSTGRES_PASSWORD string
	} `yaml:"db"`
}

func sanitizeInput(s string) string {
	re := regexp.MustCompile(`[^a-zA-Z0-9_-]`)
	return re.ReplaceAllString(s, "")
}
func main() {
	var username string
	var password string
	fmt.Println("Enter pq username:")
	fmt.Scanln(&username)
	fmt.Println("Enter pq password for user", username, ":")
	fmt.Scanln(&password)
	username = sanitizeInput(username)
	password = sanitizeInput(password)
	c := Config{}
	if err := yaml.Unmarshal(config_file, &c); err != nil {
		log.Fatalf("error: %v", err)
	}
	c.DB.POSTGRES_USERNAME = username
	c.DB.POSTGRES_PASSWORD = password
	// c.db.POSTGRES_HOST = "db"
	// c.db.POSTGRES_DB = "docker"
	connStr := fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s sslmode=disable",
		c.DB.POSTGRES_USERNAME,
		c.DB.POSTGRES_PASSWORD,
		c.DB.POSTGRES_DB,
		c.DB.POSTGRES_HOST,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var version string
	if err := db.QueryRow("SELECT VERSION()").Scan(&version); err != nil {
		log.Fatal(err)
	}

	log.Println("Postgres version:", version)
}
