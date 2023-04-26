package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

type Product struct {
	id        int
	name      string
	inventory int
	price     int
}

func main() {
	fmt.Println("welcome")
	db, err := sql.Open("sqlite3", "./practiceit.db")
	fmt.Println("got here")

	if err != nil {
		fmt.Println("got here")
		fmt.Println(err.Error())
	}
	rows, err := db.Query("SELECT id, name, inventory, pricce FROM products")
	if err != nil {
		fmt.Println(err.Error())
	}

	defer rows.Close()
	fmt.Println(rows)

	for rows.Next() {
		var p Product
		rows.Scan(&p.id, &p.name, &p.inventory, &p.price)
		fmt.Println("Product: \n", p)
	}
}
