package main

import (
    "fmt"   
"github.com/labstack/echo"
"net/http"
"log"
"database/sql"
"errors"
  _ "github.com/lib/pq"
    "bookstore/author"
    "bookstore/books"
    
)

const (
  host     = "localhost"
  port     = 5432
  user     = "postgres"
  password = "jumble52"
  dbname   = "postgres"
)


func open() *sql.DB{
    psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)

    db, err := sql.Open("postgres", psqlInfo)

    if err != nil {
        errors.New("error exists")
        return db
    }

    err = db.Ping()
    if err != nil {
        errors.New("error exists")
        return db
    }
    fmt.Println("Successfully connected!")

    return db
}



func main (){
    e := echo.New()
    db := open()
    defer db.Close()
    fmt.Println(db)
    author.SetDB(db)
    books.SetDB(db)

    router := mux.NewRouter()

    e.POST("/api/author/create", author.Create)
    e.GET("/api/author/read", author.Read)
    e.GET("/api/author/readID/:identity", author.ReadID)
    e.PUT("/api/author/update/:identity", author.Update)
    e.DELETE("/api/v1/author/delete/:identity", author.Delete)

    e.POST("/api/books/create", books.Create)
    e.GET("/api/books/read", books.Read)
    e.GET("/api/books/readID/:identity", books.ReadID)
    e.PUT("/api/books/update/:identity", books.Update)
    e.DELETE("/api/v1/books/delete/:identity", books.Delete)

    http.Handle("/", router)
    fmt.Println("Connected to port 8080")
    log.Fatal(http.ListenAndServe(":8080", router))

}