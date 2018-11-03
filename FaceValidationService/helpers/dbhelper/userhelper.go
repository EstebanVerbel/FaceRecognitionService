package dbhelper

import (
	"log"

	"github.com/gocql/gocql"
)

// RegisterUser registers users to application (insert into db)
func RegisterUser(email string, lastname string, firstName string) {

	session := GetSession()

	// Todo:
	// make sure user does not already exist

	var version = gocql.TimeUUID()

	if err := session.Query(`INSERT INTO users (email, last_name, name, version) VALUES (?,?,?,?)`,
		email, lastname, firstName, version).Exec(); err != nil {
		log.Fatal(err)
	}
}

// Todo: Implement following

// Get User Images
// Add user images
