package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

var DB *pgx.Conn

func getEnv(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("gagal membaca file .env ", err)
	}
	return os.Getenv(key)
}

func Init() {

	var (
		host     = getEnv("DB_HOST")
		port     = getEnv("DB_PORT")
		user     = getEnv("DB_USER")
		dbname   = getEnv("DB_NAME")
		password = getEnv("DB_PASSWORD")
	)

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=require",
		host,
		port,
		user,
		dbname,
		password,
	)
	var err error
	// Make sure not to shadow your global - just assign with = - don't initialise a new variable and assign with :=
	DB, err = pgx.Connect(context.Background(), dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Gagal terhubung dengan database %v\n", err)
	}
	fmt.Println("berhasil menjalanankan DB")

	err = DB.Ping(context.Background())
	if err != nil {
		fmt.Fprintf(os.Stderr, "Gagal ping ygy %v\n", err)
	}
	fmt.Println("berhasil ping")
}
