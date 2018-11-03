package dbhelper

import (
	"log"

	"github.com/gocql/gocql"
)

// RegisterUser registers users to application (insert into db)
func RegisterUser(username *string, lastname *string, firstName *string) {

	var isUserExistsAlready = verifyIfUserExists(username)

	if isUserExistsAlready {
		// TODO Esteban: Return error message or something when user already exists
		// log it to error logs
		return
	}

	session := GetSession()

	var version = gocql.TimeUUID()

	if err := session.Query(`INSERT INTO users (username, last_name, name, version) VALUES (?,?,?,?)`,
		*username, *lastname, *firstName, version).
		Exec(); err != nil {
		log.Fatal(err)
	}
}

// verifyIfUserExists checks if username exists in database
func verifyIfUserExists(username *string) bool {

	session := GetSession()
	var existingUsername string

	if err := session.Query(`SELECT "username" FROM users WHERE "username" = ?`, &username).
		Consistency(gocql.Quorum).
		Scan(&existingUsername); err != nil {
		log.Fatal(err)
	}

	// if username exists in database
	if len(existingUsername) > 0 {
		return true
	}

	return false
}

// Todo: Implement following

// Get User Images
// Add user images
// Add User Image
// Set face id for Image
