package pg_benchmark

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"
	"sync"

	_ "github.com/lib/pq"
)

const (
	DB_USER     = "admin"
	DB_PASSWORD = "secret"
	DB_NAME     = "postgres"
)

// Simple golang script to insert 1 million records in our db
func main() {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)
	checkErr(err)
	defer db.Close()

	var wg sync.WaitGroup
	wg.Add(2)
	fmt.Println("# Inserting values")
	go func() {
		defer wg.Done()
		for i := 0; i < 300000; i++ {
			var lastInsertId int
			err = db.QueryRow("INSERT INTO users(name) VALUES($1);",
				"sedhossein"+strconv.Itoa(i)).Scan(&lastInsertId)
			checkErr(err)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 300001; i < 7000000; i++ {
			var lastInsertId int
			err = db.QueryRow("INSERT INTO users(name) VALUES($1);",
				"sharez"+strconv.Itoa(i)).Scan(&lastInsertId)
			checkErr(err)
		}
	}()
	go func() {
		defer wg.Done()
		for i := 700001; i < 1000000; i++ {
			var lastInsertId int
			err = db.QueryRow("INSERT INTO users(name) VALUES($1);",
				"sedrez"+strconv.Itoa(i)).Scan(&lastInsertId)
			checkErr(err)
		}
	}()

	wg.Wait()

	//fmt.Println("# Updating")
	//stmt, err := db.Prepare("update userinfo set username=$1 where uid=$2")
	//checkErr(err)
	//
	//res, err := stmt.Exec("astaxieupdate", lastInsertId)
	//checkErr(err)
	//
	//affect, err := res.RowsAffected()
	//checkErr(err)
	//
	//fmt.Println(affect, "rows changed")
	//
	//fmt.Println("# Querying")
	//rows, err := db.Query("SELECT * FROM userinfo")
	//checkErr(err)
	//
	//for rows.Next() {
	//	var uid int
	//	var username string
	//	var department string
	//	var created time.Time
	//	err = rows.Scan(&uid, &username, &department, &created)
	//	checkErr(err)
	//	fmt.Println("uid | username | department | created ")
	//	fmt.Printf("%3v | %8v | %6v | %6v\n", uid, username, department, created)
	//}
	//
	//fmt.Println("# Deleting")
	//stmt, err = db.Prepare("delete from userinfo where uid=$1")
	//checkErr(err)
	//
	//res, err = stmt.Exec(lastInsertId)
	//checkErr(err)
	//
	//affect, err = res.RowsAffected()
	//checkErr(err)
	//
	//fmt.Println(affect, "rows changed")
}

func checkErr(err error) {
	if err != nil {
		log.Println(err)
	}
}
