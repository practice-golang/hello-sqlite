package main // import "hello-sqlite"

import (
	"database/sql"
	"fmt"
	"os"

	// sqlite
	_ "modernc.org/sqlite"
)

func tryInsertSelect() error {
	fn := "./data.db"
	db, err := sql.Open("sqlite", fn)
	if err != nil {
		return err
	}

	sql := `
	DROP TABLE IF EXISTS FAMILY;

	CREATE TABLE "FAMILY" (
		"IDX"	INTEGER,
		"NAME"	TEXT,
		"AGE"	INTEGER,
		PRIMARY KEY("IDX" AUTOINCREMENT)
	);

	INSERT INTO FAMILY
		(NAME, AGE)
	VAlUES
		("John", 11),
		("Jane", 10);
	`

	if _, err = db.Exec(sql); err != nil {
		return err
	}

	selSQL := `
	SELECT *
		FROM FAMILY
	WHERE NAME LIKE "Ja%"
		ORDER BY IDX;
	`
	rows, err := db.Query(selSQL)
	if err != nil {
		return err
	}

	for rows.Next() {
		var name string
		var idx, age int
		err = rows.Scan(&idx, &name, &age)
		if err != nil {
			return err
		}

		fmt.Println(idx, name, age)
	}

	if err = rows.Err(); err != nil {
		return err
	}

	if err = db.Close(); err != nil {
		return err
	}

	fi, err := os.Stat(fn)
	if err != nil {
		return err
	}

	fmt.Printf("%s size: %v\n", fn, fi.Size())
	return nil
}

func main() {
	if err := tryInsertSelect(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
