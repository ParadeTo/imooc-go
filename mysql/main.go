package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"fmt"
	"time"
)

type Stu struct {
	Name	string
	Age		int
}

func (Stu) TableName() string {
	return "stu"
}

func main () {
	db1, err1 := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test?charset=utf8&parseTime=True&loc=Local")
	db2, err1 := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3307)/test?charset=utf8&parseTime=True&loc=Local")
	if err1 != nil {
		panic(err1)
	}

	stu := Stu{Name:"newuser", Age: 16}
	db1.Save(&stu)

	stu1 := Stu{}
	stu2 := Stu{}
	db1.Find(&stu1)
	db2.Find(&stu2)

	fmt.Println(stu1)
	fmt.Println(stu2)

	time.Sleep(5 * time.Second)
	db2.Find(&stu2)
	fmt.Println(stu2)
}