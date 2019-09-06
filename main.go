package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

// type db struct {
// 	*sql.DB
// }
//var db *sql.DB

type user struct {
	id   int
	name string
	age  int
}

func Add(db *sql.DB, name string, age int) {
	res, err := db.Exec("INSERT INTO users (name, age) VALUES ($1, $2)", name, age)
	if err != nil {
		panic(err)
	}
	fmt.Println(res.RowsAffected())

}

func Remove(db *sql.DB, id int) {
	_, err := db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		panic(err)
	}

}

func (u *user) Get(db *sql.DB, id int) {
	row := db.QueryRow("SELECT * FROM users WHERE id = $1", id)
	err := row.Scan(&u.id, &u.name, &u.age)
	if err != nil {
		panic(err)
	}
	fmt.Println(u.id, u.name, u.age)

}

func main() {
	str := "user=postgres password=Sharingan7 dbname=test sslmode=disable"
	conn, err := sql.Open("postgres", str)
	if err != nil {
		fmt.Errorf("no coonection", err)
	}
	defer conn.Close()
	err = conn.Ping()
	if err != nil {
		panic(err)
	}
	//Remove(conn, 4)
	users := []user{}
	rows, err := conn.Query("SELECT * FROM users WHERE age BETWEEN 10 AND 20")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		u := user{}
		err := rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Println(err)
			continue
		}
		users = append(users, u)
	}
	for _, v := range users {
		fmt.Println(v.id, v.name, v.age)
	}

	//u.Get(conn, 2)

	Add(conn, "Lena", 16)

}
