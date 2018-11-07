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

// GetUserImages gets images for specific user
func GetUserImages(username *string) {

	session := GetSession()

	m := map[string]interface{}{}

	var images [][]byte
	var faceIds []string

	query := `SELECT image, face_id FROM user_images WHERE "username" = ?`

	iter := session.Query(query, username).Iter()
	for iter.MapScan(m) {

		faceID := m["face_id"].(string)
		photo := m["image"].([]byte)

		images = append(images, photo)
		faceIds = append(faceIds, faceID)
	}

	// TODO: Define how to return images (what format)
}

// AddUserImage adds user image to database
func AddUserImage(username *string, image *[]byte) {

	session := GetSession()

	isUserExists := verifyIfUserExists(username)

	// if username exists in database
	if isUserExists == false {
		return
	}

	query := "INSERT INTO user_images (username, id, image) VALUES (?, ?, ?)"
	id := gocql.TimeUUID

	if err := session.Query(query, &username, id, &image).
		Exec(); err != nil {
		log.Fatal(err)
	}
}

// Todo: Implement following

// Add User Image
// Set face id for Image
