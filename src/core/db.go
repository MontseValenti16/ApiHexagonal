package core

import (
    "database/sql"
    "fmt"

    "os"

    _ "github.com/go-sql-driver/mysql" 
)

var DB *sql.DB

func InitDB() (*sql.DB, error) {
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbHost := os.Getenv("DB_HOST")
    dbPort := os.Getenv("DB_PORT")
    dbName := os.Getenv("DB_NAME")

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
        dbUser, dbPassword, dbHost, dbPort, dbName)

    var err error
    DB, err = sql.Open("mysql", dsn)
    if err != nil {
        return nil, fmt.Errorf("error al conectar a la base de datos: %v", err)
    }

    if err = DB.Ping(); err != nil {
        return nil, fmt.Errorf("error al hacer ping a la base de datos: %v", err)
    }

    fmt.Println("Conectado a la base de datos correctamente.")
    return DB, nil
}

func CloseDB() {
    if DB != nil {
        DB.Close()
    }
}