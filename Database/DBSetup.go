package sqlDB

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)




type (

	Urlshort struct {
		id int
		Url string
		time string
		Token string 
	
	}

	Database struct {
		DBName string
		dbConnect *sql.DB
	}

)


const (

	create_db string = `
CREATE TABLE IF NOT EXISTS url_shortner (
id INTEGER NOT NULL PRIMARY KEY,
time DATETIME NOT NULL,
Url TEXT,
Token TEXT
);`
)

func(con *Database) Init(){
	
	db, err := sql.Open("sqlite3", con.DBName)
	
	
	if err != nil {
		fmt.Println(err)
	}
	
	if _, err := db.Exec(create_db); err != nil {
		
		fmt.Printf("An error occured: %s", err)
		os.Exit(1)
		
		} 
		con.dbConnect = db
}
	
	
func (con *Database) FetchOne(table_name, column, token string) (Urlshort, error) {
	var dbRow Urlshort
	var rowErr error
	
	str := fmt.Sprintf("SELECT * FROM %s where %s=?", table_name, column)
	fmt.Println(str)
	row := con.dbConnect.QueryRow(str, token)
	fmt.Println(row)
	rowErr = row.Scan(&dbRow.id, &dbRow.time, &dbRow.Url, &dbRow.Token)
	return dbRow, rowErr
}
