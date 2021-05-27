package main
 
import (
    "database/sql"
    "fmt"
    _ "github.com/lib/pq"
    "conn.go"
)
 
 
func main() {
    CheckError(err)
 
    defer db.Close()

    insertStmt := `insert into "Students"("Name", "Roll") values('John', 1)`
    _, e := db.Exec(insertStmt)
    CheckError(e)
 
    insertDynStmt := `insert into "Students"("Name", "Roll") values($1, $2)`
    _, e = db.Exec(insertDynStmt, "Jane", 2)
    CheckError(e)
}
