package main

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
)


func main() {
	db, err := sql.Open("postgres", "host=localhost port=5432 dbname=dev_database user=user password=asdf sslmode=disable")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	userStore := NewUserStore(db)

	user := User{ Slug: "dfajsdlfk", InternalSlug: "sadflkj", Email: "stuff@man.com", Password: []byte("stuff") }
	dataerr := userStore.Insert(&user)
	if dataerr != nil {
		fmt.Println(dataerr)
		panic(dataerr)
	}
}
