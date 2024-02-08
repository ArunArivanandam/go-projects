package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v4/stdlib"
)

func getAllRowData(conn *sql.DB) error {
	rows, err := conn.Query("SELECT id, name, email from users")
	if err != nil {
		log.Println(err)
		return err
	}
	defer rows.Close()
	var name, email string 
	var id int 
	for rows.Next() {
		err := rows.Scan(&id, &name, &email)
		if err != nil {
			log.Println(err)
		}
		fmt.Println("Data: ",id, name, email)
	}
	if err != nil {
		log.Fatal("Error Reading Data :", err)
	}
	return nil
}

func insertNewUser(conn *sql.DB, name string, email string, pw string, uType int) error {
	query := fmt.Sprintf(`INSERT INTO users (name, email, password,acct_created, last_login, user_type) VALUES ('%s', '%s', '%s', current_timestamp, current_timestamp, %d)`, name, email, pw, uType)

	_, err := conn.Exec(query)
	if err != nil {
		log.Fatal(err)
		return err 
	}
	return nil 
}

func getUserData(conn *sql.DB, id int) error {
	var name, email, pw, uType string

	query := fmt.Sprintf(`SELECT id, name, email, password, user_type FROM users WHERE id = %d`, id)

	row := conn.QueryRow(query)
	err := row.Scan(&id, &name, &email, &pw, &uType)

	if err != nil {
		log.Fatal(err)
		return err
	}
	fmt.Println("ID: ", id)
	fmt.Println("Name: ", name)
	fmt.Println("Email: ", email)
	fmt.Println("Password: ", pw)
	fmt.Println("User Type: ", uType)
	return nil 


}

func updateUserEmail(conn *sql.DB, newEmail string, id int) error {
	query := fmt.Sprintf(`UPDATE users SET email = '%s' WHERE id = %d`, newEmail, id)
	_, err := conn.Exec(query)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func deleteUserByID(conn *sql.DB, id int) error {
	query := fmt.Sprintf(`DELETE FROM users WHERE id = %d`, id)
	_, err := conn.Exec(query)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func main() {
	conn, err := sql.Open("pgx", "host=localhost port=5432 dbname=blog_db user=postgres password=root")
	if err != nil {
		log.Fatalf(fmt.Sprintf("Couldn't connect to db : %v", err))
	}
	defer conn.Close()

	err = conn.Ping()
	if err != nil {
		log.Fatalf("Couldn't ping database")
	}

	// err = insertNewUser(conn, "arun a", "arun@gmail.com", "abcdefgh", 3)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// err = getAllRowData(conn)
	// if err != nil {
	// 	log.Fatal(err)
	// }


	fmt.Println(("------------------------"))
	getUserData(conn, 2)
	fmt.Println(("------------------------"))

	updateUserEmail(conn, "arunaruvanandam@example.com", 2)
	if err != nil {
		log.Fatal(err)
	}

	getUserData(conn, 2)

	err = deleteUserByID(conn, 2)
	if err != nil {
		log.Fatal(err)
	}
	err = getAllRowData(conn)
	if err != nil {
		log.Fatal(err)
	}
}

