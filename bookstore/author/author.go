package author

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

var AuthorBusinessImpl *AuthorBusiness

func init() {
    AuthorBusinessImpl = &AuthorBusiness{}
}

func SetDB(db *sql.DB) {
    AuthorBusinessImpl.DB = db
}

type Authors struct {
  ID        int
  Name      string
  Createdat string
  Updatedat string
  Deletedat string
}

type AuthorBusiness struct {
    DB *sql.DB
}

func (a *AuthorBusiness) Create(fullName string) (*Authors, error) {
    sqlStatement := `
    INSERT INTO authors (full_name)
    VALUES ($1)
    RETURNING id`
    id := 0
    if err := AuthorBusinessImpl.DB.QueryRow(sqlStatement, fullName).Scan(&id); err != nil {
        return nil, err
    }
    t := time.Now()
    return &Authors{
        ID: id,
        Name: fullName,
        Createdat: t.String(),
        Updatedat: t.String(),
        Deletedat: "",
    }, nil
}

func (a *AuthorBusiness) Read()([]Authors, error) {
    sqlStatement := `SELECT id, full_name, createdat, updatedat, deletedat FROM authors;`
    rows, err := AuthorBusinessImpl.DB.Query(sqlStatement)
    if err != nil{
        return nil, errors.New("Query error exists")
    }
    s := make([]Authors, 0)
    for rows.Next(){
        var author Authors
        err := rows.Scan(&author.ID, &author.Name, &author.Createdat,
        &author.Updatedat, &author.Deletedat)
        if err != nil {
            return nil, errors.New("scanning error exists")
        }
        err = rows.Err()
        if err != nil {
            return nil, errors.New("rows error exists")
        }
        s = append(s, author)
    }
    return s, err
}

func (a *AuthorBusiness) ReadID(ID int) (*Authors, error) {
    sqlStatement := `SELECT id, full_name, createdat, updatedat, deletedat FROM authors WHERE id=$1;`
    row := AuthorBusinessImpl.DB.QueryRow(sqlStatement, ID)
    var author Authors
    err := row.Scan(&author.ID, &author.Name, &author.Createdat,
        &author.Updatedat, &author.Deletedat)
    switch err {
        case sql.ErrNoRows:
            fmt.Println("No rows were returned!")
            return nil, errors.New("row error exists")
        case nil:
            return &author, nil
    }
    return &author, nil
}

func (a *AuthorBusiness) Update(ID int, fullName string) (*Authors, error) {
    sqlStatement := `
    UPDATE authors
    SET full_name = $2
    WHERE id = $1;`
    _, err := AuthorBusinessImpl.DB.Exec(sqlStatement, ID, fullName)
    if err != nil {
        return nil, errors.New("error exists")
        log.Println("error exists")
    }
    t := time.Now()

    return &Authors{
        ID: ID,
        Name: fullName,
        Createdat: t.String(),
        Updatedat: t.String(),
        Deletedat: "",
    }, nil
}

func (a *AuthorBusiness) Delete(ID int) (*Authors, error) {
    sqlStatement := `SELECT id, full_name, createdat, updatedat, deletedat FROM authors WHERE id=$1;`
    row := AuthorBusinessImpl.DB.QueryRow(sqlStatement, ID)
    var author Authors
    var authortwo Authors
    err := row.Scan(&author.ID, &author.Name, &author.Createdat,
    &author.Updatedat, &author.Deletedat)
    authortwo = author
    switch err {
    case sql.ErrNoRows:
        fmt.Println("No rows were returned!")
        return nil, err
    case nil:
        sqlStatement := `
        DELETE FROM books
        WHERE author_id = $1;`
        _, err = AuthorBusinessImpl.DB.Exec(sqlStatement, ID)
        if err != nil {
            return nil, err
        }
        sqlStatement = `
        DELETE FROM authors
        WHERE id = $1;`
        _, err = AuthorBusinessImpl.DB.Exec(sqlStatement, ID)
        if err != nil {
            return nil, err
        }
        return &authortwo, nil
    }
    return &authortwo, nil
}

func Create (w http.ResponseWriter, r *http.Request){
    payload := struct {
        FullName string `json:"fullName"`
    }{}
    if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
        log.Fatalln(err)
        return
    }
    
    author, err := AuthorBusinessImpl.Create(payload.FullName)
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
        Result *Authors `json:"result:"`
        }{
            Result: author,
        }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(resp)
}

func Read (w http.ResponseWriter, r *http.Request){  
    strmsg, err := AuthorBusinessImpl.Read()
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
        Result []Authors `json:"result"`
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

    author, err := AuthorBusinessImpl.ReadID(id)
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
        Result *Authors `json:"result"`
    }{
        Result: author,
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(resp)
}

func Update (w http.ResponseWriter, r *http.Request){
    params := mux.Vars(r)
    num := params["identity"]
    id,_ := strconv.Atoi(num)
    payload := struct {
        ID string `json:"id"`
        FullName string `json:"fullName"`
    }{}
    if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
        log.Fatalln(err)
        return
    }
    idtwo,_ := strconv.Atoi(payload.ID)

    if idtwo != id{
        resp := struct {
            Message string `json:"message:"`
        }{
            Message: "matching errors exist",
        }
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(resp)
        return
    }
    
    author, err := AuthorBusinessImpl.Update(id, payload.FullName)
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
        Result *Authors `json:"result"`
    }{
        Result: author,
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(resp)
}

func Delete (w http.ResponseWriter, r *http.Request){
    params := mux.Vars(r)
    num := params["identity"]
    id,_ := strconv.Atoi(num)
    fmt.Println(id)
    author, err := AuthorBusinessImpl.Delete(id)
    if err != nil {
        fmt.Println(err.Error())
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
        Result *Authors `json:"result"`
    }{
        Result: author,
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(resp)
}