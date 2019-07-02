package main

import (
    "fmt"   
"github.com/gorilla/mux"
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
    db := open()
    defer db.Close()
    fmt.Println(db)
    author.SetDB(db)
    books.SetDB(db)

    router := mux.NewRouter()

    router.HandleFunc("/api/author/create", author.Create).Methods(http.MethodPost)
    router.HandleFunc("/api/author/read", author.Read).Methods(http.MethodGet)
    router.HandleFunc("/api/author/readID/{identity}", author.ReadID).Methods(http.MethodGet)
    router.HandleFunc("/api/author/update/{identity}", author.Update).Methods(http.MethodPut)
    router.HandleFunc("/api/v1/author/delete/{identity}", author.Delete).Methods(http.MethodDelete)

    router.HandleFunc("/api/books/create", books.Create).Methods(http.MethodPost)
    router.HandleFunc("/api/books/read", books.Read).Methods(http.MethodGet)
    router.HandleFunc("/api/books/readID/{identity}", books.ReadID).Methods(http.MethodGet)
    router.HandleFunc("/api/books/update/{identity}", books.Update).Methods(http.MethodPut)
    router.HandleFunc("/api/v1/books/delete/{identity}", books.Delete).Methods(http.MethodDelete)

    http.Handle("/", router)
    fmt.Println("Connected to port 8080")
    log.Fatal(http.ListenAndServe(":8080", router))

}