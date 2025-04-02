package links

import (
    database "my-gqlgen-app/internal/pkg/db/migrations/mysql"
    "my-gqlgen-app/internal/users"
    "log"
	"fmt"
)
// #1
type Link struct {
	ID      string
	Title   string
	Address string
	User    *users.User
}

//#2
// func (link Link) Save() int64 {
// 	//#3
// 	stmt, err := database.Db.Prepare("INSERT INTO Links(Title,Address, UserID) VALUES(?,?, ?)")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	//#4
// 	// res, err := stmt.Exec(link.Title, link.Address)
// 	res, err := stmt.Exec(link.Title, link.Address, link.User.ID)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	//#5
// 	id, err := res.LastInsertId()
// 	if err != nil {
// 		log.Fatal("Error:", err.Error())
// 	}
// 	log.Print("Row inserted!")
// 	return id
// }
// func GetAll() []Link {
// 	stmt, err := database.Db.Prepare("select id, title, address from Links")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer stmt.Close()
// 	rows, err := stmt.Query()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer rows.Close()
// 	var links []Link
// 	for rows.Next() {
// 		var link Link
// 		err := rows.Scan(&link.ID, &link.Title, &link.Address)
// 		if err != nil{
// 			log.Fatal(err)
// 		}
// 		links = append(links, link)
// 	}
// 	if err = rows.Err(); err != nil {
// 		log.Fatal(err)
// 	}
// 	return links
// }

func (link Link) Save() (int64, error) {
	// Prepare the SQL statement
	stmt, err := database.Db.Prepare("INSERT INTO Links(Title, Address, UserID) VALUES (?, ?, ?)")
	if err != nil {
		return 0, fmt.Errorf("failed to prepare statement: %w", err)
	}
	defer stmt.Close() // Ensure stmt is closed properly

	// Execute the statement
	res, err := stmt.Exec(link.Title, link.Address, link.User.ID)
	if err != nil {
		return 0, fmt.Errorf("failed to execute insert: %w", err)
	}

	// Get the last inserted ID
	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to get last insert ID: %w", err)
	}

	log.Print("Row inserted successfully!")
	return id, nil
}

func GetAll() []Link {
	stmt, err := database.Db.Prepare("select L.id, L.title, L.address, L.UserID, U.Username from Links L inner join Users U on L.UserID = U.ID") // changed
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var links []Link
	var username string
	var id string
	for rows.Next() {
		var link Link
		err := rows.Scan(&link.ID, &link.Title, &link.Address, &id, &username) // changed
		if err != nil{
			log.Fatal(err)
		}
		link.User = &users.User{
			ID:       id,
			Username: username,
		} // changed
		links = append(links, link)
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
	return links
}