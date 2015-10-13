package main

import (
	//"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"runtime"
	"time"
)

var (
	db            gorm.DB
	sqlConnection string
)

type Person struct {
	Id       int64     `gorm:"column:Id;primary_key" sql:"AUTO_INCREMENT"`
	Name     string    `gorm:"column:Name"`
	Birthday time.Time `gorm:"column:Birthday"`
	InDate   time.Time `gorm:"column:InDate"`
}

//通过Struct的方法指定自定义的表名
func (p Person) TableName() string {
	return "Person"
}

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	var err error
	//loc=Local转换本地时间，否则写入mysql的时间是utc时间(当前时间-8小时)
	sqlConnection = "writeuser:writeuser@tcp(192.168.1.116:3306)/TestInsert?charset=utf8&parseTime=True&loc=Local"
	db, err = gorm.Open("mysql", sqlConnection)
	if err != nil {
		panic(err.Error())
		return
	}
	//defer db.Close()

	// Get database connection handle [*sql.DB](http://golang.org/pkg/database/sql/#DB)
	db.DB()

	// Then you could invoke `*sql.DB`'s functions with it
	db.DB().Ping()
	db.DB().SetMaxIdleConns(5)
	db.DB().SetMaxOpenConns(20)

	// Disable table name's pluralization, eg: user-->users
	//db.SingularTable(true)

	// Enable Logger
	db.LogMode(true)
}

func main() {
	defer db.Close()

	person := Person{
		Id:       0,
		Name:     "name9",
		Birthday: time.Now(),
	}

	//简单插入
	//db.Save(&person) //根据主健确定insert or update, 自增主健值会自动带回
	//fmt.Println(person.Id)

	//复杂插入
	res, err := db.DB().Exec("INSERT INTO Person(Name, Birthday) VALUES (?, ?)", person.Name, person.Birthday)
	if err != nil {
		println("Exec err:", err.Error())
	} else {
		id, err := res.LastInsertId() //.Raw("select LAST_INSERT_ID()", nil).First(&id)
		if err != nil {
			println("Error:", err.Error())
		} else {
			fmt.Println(id)
			person.Id = id
		}
	}

	//不能返回自增值的插入，不建议使用
	//db.Create(&person)
	//db.NewRecord(person)

	//简单修改
	//person.Name = "name9xx"
	//db.Save(&person)

	//复杂修改
	person.Name = "namexx"
	_, err = db.DB().Exec("UPDATE Person SET Name = ? WHERE Id = ?", person.Name, person.Id)
	if err != nil {
		println("Exec err:", err.Error())
	}

	//查询
	// Raw SQL
	var pp Person
	db.Raw("select Id,Name from Person where Id = ?", 10).Find(&pp)
	fmt.Println(pp.Name)

	var persons []Person
	db.Raw("select Id,Name from Person where Id > ?", 0).Scan(&persons)
	for _, p := range persons {
		fmt.Println(p.Name)
	}

	//删除
	db.DB().Exec("DELETE FROM Person WHERE Id = ?", 4)

	//使用原生事务
	tx, _ := db.DB().Begin()

	//使用gorm事务
	//tx := db.Begin()

	//事务操作
	tx.Exec("INSERT INTO Person(Name, Birthday) VALUES (?, ?)", person.Name+"tran", person.Birthday)

	// rollback in case of error
	//tx.Rollback()

	// Or commit if all is ok
	tx.Commit()
}

func checkErr(err error, msg string) {
	if err != nil {
		fmt.Println(msg, err)
	}
}
