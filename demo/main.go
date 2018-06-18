package main

import (
	"fmt"
	"errors"
)

type Human struct {
	id   int
	name string
}

func (h *Human) SetID(id int) { // = func setId(h *Human, id int)
	h.id = id
}

func newFirstErrorKeeper() func(err error) error {
	var e error
	return func(err error) error {
		if e == nil && err != nil {
			e = err
		}
		return e
	}
}

func updateSomething() {
	tx := db.Begin()
	defer tx.Rollback() // 如果提交了，这一句不会执行

	rows, err := tx.Query("do something")
	if err != nil {
		// handle error
		return
	}
	handleRows(rows)
	tx.Commit()
}

func main() {
	var keepFirstErr = newFirstErrorKeeper()
	fmt.Println(keepFirstErr(nil))
	fmt.Println(keepFirstErr(errors.New("1")))
	fmt.Println(keepFirstErr(errors.New("2")))
	// return keepFirstErr(nil)

}