package main
 
import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
)
 
const (
    host     = "localhost"
    port     = 5432
    user     = "postgres"
    dbname   = "goapi"
)
 
func main() {
    conn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
         
    db, err := sql.Open("postgres", conn)
    CheckError(err)
     
    defer db.Close()
 
    err = db.Ping()
    CheckError(err)
 
    fmt.Println("Connected!")
}
 
func CheckError(err error) {
    if err != nil {
        panic(err)
    }
}
