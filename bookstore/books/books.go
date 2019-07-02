package books

import (
    "fmt"   
"github.com/gorilla/mux"
"net/http"
"log"
"encoding/json"
"database/sql"
_ "github.com/lib/pq"
"strconv"
"time"
"errors"
)

var BooksBusinessImpl *BooksBusiness

func init() {
    BooksBusinessImpl = &BooksBusiness{}
}

func SetDB(db *sql.DB) {
    BooksBusinessImpl.DB = db
}

type Books struct {
  ID        int
  ISBN      string 
  Title     string
  Category  string
  Author_id int
  Createdat string
  Updatedat string
  Deletedat string
}

type BooksBusiness struct {
    DB *sql.DB
}

func (a *BooksBusiness) Create(isbn string, Title string, Category string, Author_id string) (*Books, error) {
    sqlStatement := `
    INSERT INTO books (isbn, title, category, author_id)
    VALUES ($1, $2, $3, $4)
    RETURNING id`
    id := 0
    aid,_ := strconv.Atoi(Author_id)
    
    if err := BooksBusinessImpl.DB.QueryRow(sqlStatement, isbn, Title, Category, Author_id).Scan(&id); err != nil {
        return nil, err
    }
    t := time.Now()
    return &Books{
        ID: id,
        ISBN: isbn,
        Title: Title,
        Category: Category,
        Author_id: aid,
        Createdat: t.String(),
        Updatedat: t.String(),
        Deletedat: "",
    }, nil
}

func (a *BooksBusiness) Read() ([]Books, error) {
    sqlStatement := `SELECT id, isbn, title, category, author_id, createdat, updatedat, deletedat FROM books;`
    rows, err := BooksBusinessImpl.DB.Query(sqlStatement)
    if err != nil {
        return nil, errors.New("error exists")
    }
    s := make([]Books, 0)
    for rows.Next(){
        var book Books
        err := rows.Scan(&book.ID, &book.ISBN, &book.Title, &book.Category, &book.Author_id, & book.Createdat,
        &book.Updatedat, &book.Deletedat)
        if err != nil {
            return nil, errors.New("category error exists")
        }
        err = rows.Err()
        if err != nil {
            return nil, errors.New("row error exists")
        }
        s = append(s, book)
    }
    return s, nil  
}

func (a *BooksBusiness) ReadID(ID int) (*Books, error) {
    sqlStatement := `SELECT id, isbn, title, category, author_id, createdat, updatedat, deletedat FROM books WHERE id=$1;`
    row := BooksBusinessImpl.DB.QueryRow(sqlStatement, ID)
    var book Books
    err := row.Scan(&book.ID, &book.ISBN, &book.Title, &book.Category, &book.Author_id, &book.Createdat,
    &book.Updatedat, &book.Deletedat)
    switch err {
        case sql.ErrNoRows:
            fmt.Println("No rows were returned!")
            return nil, err
        case nil:
            return &book, nil
    }
    return &book, nil
}

func (a *BooksBusiness) Update(ID int, Isbn string, Title string, Category string, Author_id int) (*Books, error) {
    sqlStatement := `UPDATE books SET isbn = $1, title = $2, category = $3, author_id = $4 WHERE id = $5;`
    fmt.Println(Author_id)
    _, err := BooksBusinessImpl.DB.Exec(sqlStatement, Isbn, Title, Category, Author_id, ID)
    if err != nil {
        fmt.Println(err.Error())
        return nil, errors.New("execution error exists")
    }
    
    t := time.Now()
    return &Books{
        ID: ID,
        ISBN: Isbn,
        Title: Title,
        Category: Category,
        Author_id: Author_id,
        Createdat: t.String(),
        Updatedat: t.String(),
        Deletedat: "",
    }, nil
}

func (a *BooksBusiness) Delete(ID int) (*Books, error) {
    sqlStatement := `SELECT id, isbn, title, category, author_id, createdat, updatedat, deletedat FROM books WHERE id=$1;` 
    row := BooksBusinessImpl.DB.QueryRow(sqlStatement, ID)
    var book Books
    err := row.Scan(&book.ID, &book.ISBN, &book.Title, &book.Category, &book.Author_id, &book.Createdat,
    &book.Updatedat, &book.Deletedat)
    booktwo := book
    switch err {
        case sql.ErrNoRows:
            fmt.Println("No rows were returned!")
            return nil, err
        case nil:
            sqlStatement := `
            DELETE FROM books
            WHERE id = $1;`
            _, err = BooksBusinessImpl.DB.Exec(sqlStatement, ID)
            if err != nil {
                return nil, errors.New("error exists")
            }
            return &booktwo, nil
    }
    return &booktwo, nil
}

func Create (w http.ResponseWriter, r *http.Request){
    payload := struct {
        Isbn string `json:"isbn"`
        Title string `json:"Title"`
        Category string `json:"Category"`
        Author_id string `json:"Author_id"`
    }{}

    if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
        log.Fatalln(err)
        return
    }
    books, err := BooksBusinessImpl.Create(payload.Isbn, payload.Title, payload.Category, payload.Author_id)
    if err != nil {
        errors.New("error exists")
        resp := struct {
            Message string `json:"message:"`
        }{
            Message: err.Error(),
        }
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(resp)
        return 
    }
    resp := struct {
        Result *Books `json:"result:"`
    }{
        Result: books,
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(resp)
}

func Read (w http.ResponseWriter, r *http.Request){
    strmsg, err := BooksBusinessImpl.Read()
    if err != nil {
        errors.New("error exists")
        resp := struct {
            Message string `json:"message:"`
        }{
            Message: err.Error(),
        }
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(resp)
        return 
    }

    resp := struct {
        Result []Books `json:"result"`
    }{
        Result: strmsg,
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(resp)
}

func ReadID (w http.ResponseWriter, r *http.Request){
    params := mux.Vars(r)
    num := params["identity"]
    id,_ := strconv.Atoi(num)

    books, err := BooksBusinessImpl.ReadID(id)
    if err != nil {
        errors.New("error exists")
        resp := struct {
            Message string `json:"message:"`
        }{
            Message: err.Error(),
        }
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(resp)
        return 
    }
    
    resp := struct {
        Result *Books `json:"result"`
    }{
        Result: books,
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(resp)
}

func Update (w http.ResponseWriter, r *http.Request){
    params := mux.Vars(r)
    num := params["identity"]
    id,_ := strconv.Atoi(num)
    payload := struct {
        ID int `json:"id"`
        Isbn  string `json:"isbn"`
        Title  string `json:"title"`
        Category  string `json:"category"`
        Author_id  int `json:"Author_id"`
    }{}
    if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
        log.Fatalln(err)
        return
    }
    if payload.ID != id{
        resp := struct {
            Message string `json:"message:"`
        }{
            Message: "matching errors exist",
        }
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(resp)
        return
    }

    books, err := BooksBusinessImpl.ReadID(id)

    if err != nil {
        err = errors.New("reading error exists")
        resp := struct {
            Message string `json:"message:"`
        }{
            Message: err.Error(),
        }
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(resp)
        return 
    }

    if books.ID == 0 {
        err = errors.New("ID does not exist")
        resp := struct {
            Message string `json:"message:"`
        }{
            Message: err.Error(),
        }
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(resp)
        return 
    }

    books, err = BooksBusinessImpl.Update(id, payload.Isbn, payload.Title, payload.Category, payload.Author_id)
    if err != nil {
        err = errors.New("object error exists")
        resp := struct {
            Message string `json:"message:"`
        }{
            Message: err.Error(),
        }
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(resp)
        return 
    }
    
    resp := struct {
        Result *Books `json:"result"`
    }{
        Result: books,
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(resp)
}

func Delete (w http.ResponseWriter, r *http.Request){
    params := mux.Vars(r)
    num := params["identity"]
    id,_ := strconv.Atoi(num)

    books, err := BooksBusinessImpl.Delete(id)

    if err != nil {
        errors.New("error exists")
        resp := struct {
            Message string `json:"message:"`
        }{
            Message: err.Error(),
        }
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(resp)
        return 
    }
    resp := struct {
        Result *Books `json:"result"`
    }{
        Result: books,
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(resp)
}