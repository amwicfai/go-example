package main

import (
	"database/sql"
	"fmt"
	"github.com/go-gorp/gorp"
	_ "github.com/go-sql-driver/mysql"
	"runtime"
)

type Person struct {
	Id   int32
	Name string
}

var dbmap *gorp.DbMap

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	db, err := sql.Open("mysql", "writeuser:writeuser@tcp(192.168.1.116:3306)/TestInsert?charset=utf8")
	if err != nil {
		panic(err.Error())
	}
	//defer db.Close()

	err = db.Ping()
	if err != nil {
		// do something here
	}

	dbmap = &gorp.DbMap{Db: db, Dialect: gorp.MySQLDialect{"InnoDB", "UTF8"}}

	table := dbmap.AddTableWithName(Person{}, "Person").SetKeys(true, "Id")
	if table != nil {
		//table.ColMap("Username").SetMaxSize(25)
	}

}

func main() {
	defer dbmap.Db.Close()

	p1 := Person{
		Id:   1,
		Name: "名四",
	}

	err := dbmap.Insert(&p1)
	checkErr(err, "Insert failed")

	// update a row
	p1.Name = "名四update"
	count, err := dbmap.Update(&p1)
	checkErr(err, "Update failed")
	fmt.Println("Rows updated:", count)

	p2 := Person{
		Id: 1,
	}
	err2 := dbmap.SelectOne(&p2, "select Id,Name from Person where Id=?", 2)
	checkErr(err2, "SelectOne failed")
	fmt.Println(p2.Name)

	// delete row by PK
	count, err := dbmap.Delete(&p1)
	checkErr(err, "Delete failed")
	fmt.Println("Rows deleted:", count)

	// delete row manually via Exec
	_, err = dbmap.Exec("delete from Person where Id=?", p1.Id)
	checkErr(err, "Exec failed")
}

func checkErr(err error, msg string) {
	if err != nil {
		fmt.Println(msg, err)
	}
}
