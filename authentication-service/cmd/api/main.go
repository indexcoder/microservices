package main

import (
	"authentication/data"
	"database/sql"
	"fmt"
	_ "github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	"log"
	"net/http"
	"os"
	"time"
)

const webPort = 80

var counts int64

type Config struct {
	DB     *sql.DB
	Models data.Models
}

func main() {

	log.Println("Starting authentication server on port ", webPort)

	// connect to DB
	conn := connectToDB()
	if conn == nil {
		log.Panic("Невозможно подключиться к postgres")
	}
	defer conn.Close()

	// set up config
	app := Config{
		DB:     conn,
		Models: data.New(conn),
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", webPort),
		Handler: app.routes(),
	}

	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}

}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func connectToDB() *sql.DB {
	dsn := os.Getenv("DSN")
	if dsn == "" {
		dsn = "host=postgres port=5439 user=postgres password=password dbname=users sslmode=disable"
		log.Println("Используется стандартный DSN:", dsn)
		log.Fatal("Переменная окружения DSN не задана")
	}

	for {
		connection, err := openDB(dsn)
		if err != nil {
			log.Printf("Не удалось подключиться к Postgres: %v\n", err)
			counts++
		} else {
			log.Println("Подключение к Postgres установлено")
			return connection
		}

		if counts > 10 {
			log.Println("Превышено количество попыток подключения к Postgres")
			os.Exit(1) // Программа завершится, если подключение не удалось
		}

		log.Println("Отступление на 5 секунд...")
		time.Sleep(5 * time.Second)
	}
}
