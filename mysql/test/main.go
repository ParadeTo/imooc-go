package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"math/rand"
)

type Dept struct {
	Name string
	Id   int
}

func (Dept) TableName() string {
	return "dept"
}

type Staff struct {
	Id     int
	DeptId int
	Name   string
}

func (Staff) TableName() string {
	return "staff"
}

func main() {
	db, err := gorm.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/log-analysis?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	for i := 0; i < 5000; i++ {
		var deptId int
		r := rand.Float32()
		if r > 0.5 {
			deptId = 2
		} else {
			deptId = 1
		}
		stu := Staff{Name: fmt.Sprint("newuser%s", deptId), DeptId: deptId}
		db.Save(&stu)
	}
}
