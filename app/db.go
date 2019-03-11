package app

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func dbConnect() (*sql.DB, error) {
	conf, err := getConfig()
	if err != nil {
		return nil, err
	}

	dbInfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable",
		conf.DBUser, conf.DBPass, conf.DBName)

	db, err := sql.Open("postgres", dbInfo)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func dbSearch(query string) ([]int, error) {
	db, err := dbConnect()
	if err != nil {
		return nil, err
	}
	defer db.Close()

	rows, err := db.Query(`
		SELECT num FROM xkcd WHERE transcript LIKE $1`, fmt.Sprintf("%%%s%%", query),
	)

	if err != nil {
		return nil, err
	}

	var result []int
	var num int
	for rows.Next() {
		rows.Scan(&num)
		result = append(result, num)
	}

	if err != nil {
		return nil, err
	}

	return result, nil
}

func dbCheck(num int) bool {
	db, err := dbConnect()
	if err != nil {
		return false
	}
	defer db.Close()

	result, err := db.Exec(`SELECT num FROM xkcd WHERE num = $1`, num)
	if err != nil {
		return false
	}

	row, err := result.RowsAffected()
	if err != nil {
		return false
	}

	if row == 0 {
		return false
	}

	return true
}
