package author

import (
    "fmt"   
"github.com/labstack/echo"
"net/http"
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

func Create (c echo.Context) error{
    payload := struct {
        FullName string `json:"fullName"`
    }{}
    if err := c.Bind(&payload); err != nil {
        return err
    }

    author, err := AuthorBusinessImpl.Create(payload.FullName)
    if err != nil {
        errors.New("error exists")
        return c.String(http.StatusOK, err.Error())
    }
    return c.JSON(http.StatusOK, author)
}

func Read (c echo.Context) error{  
    strmsg, err := AuthorBusinessImpl.Read()
    if err != nil {
        errors.New("error exists")
        return c.String(http.StatusOK, err.Error())
    }
    return c.JSON(http.StatusOK, strmsg)
}

func ReadID (c echo.Context) error{
    num := c.Param("identity")
    id,_ := strconv.Atoi(num)
    author, err := AuthorBusinessImpl.ReadID(id)
    if err != nil {
        errors.New("error exists")
        return c.String(http.StatusOK, err.Error())
    }
    return c.JSON(http.StatusOK, author)

}

func Update (c echo.Context) error{
    num := c.Param("identity")
    id,_ := strconv.Atoi(num)
    payload := struct {
        ID string `json:"id"`
        FullName string `json:"fullName"`
    }{}
    if err := c.Bind(&payload); err != nil {
        return err
    }
    idtwo,_ := strconv.Atoi(payload.ID)
    if idtwo != id{
        return c.String(http.StatusOK, "matching errors exist")
    }
    author, err := AuthorBusinessImpl.ReadID(id)
    if err != nil {
        return c.String(http.StatusOK, err.Error())
    }
    if author.ID == 0 {
        err = errors.New("ID does not exist")
        return c.String(http.StatusOK, err.Error())
    }
    author, err = AuthorBusinessImpl.Update(id, payload.FullName)
    if err != nil {
        errors.New("error exists")
        return c.String(http.StatusOK, err.Error())

    }
    return c.JSON(http.StatusOK, author)
}

func Delete (c echo.Context) error{
    num := c.Param("identity")
    id,_ := strconv.Atoi(num)
    fmt.Println(id)
    author, err := AuthorBusinessImpl.Delete(id)
    if err != nil {
        fmt.Println(err.Error())
        errors.New("error exists")
        return c.String(http.StatusOK, err.Error())
    }
    return c.JSON(http.StatusOK, author)
}