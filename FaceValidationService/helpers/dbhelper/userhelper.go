package dbhelper

import (
	"fmt"
	"log"

	"github.com/gocql/gocql"
)

// RegisterUser registers users to application (insert into db)
func RegisterUser(username *string, lastname *string, firstName *string) {

	var isUserExistsAlready = verifyIfUserExists(*username)

	if isUserExistsAlready {
		// TODO Esteban: Return error message or something when user already exists
		// log it to error logs
		fmt.Println("User can't be registered because it already Exists")
		return
	}

	session := GetSession()
	defer session.Close()

	var version = gocql.TimeUUID()

	if err := session.Query(`INSERT INTO users (username, last_name, name, version) VALUES (?,?,?,?)`,
		*username, *lastname, *firstName, version).
		Exec(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("User Registered Successfully")
}

// GetUserImages gets images for specific user
func GetUserImages(username *string) {

	var isUserExistsAlready = verifyIfUserExists(*username)

	if isUserExistsAlready == false {
		// TODO Esteban: Return error message or something when user already exists
		// log it to error logs
		fmt.Println("This user is not registered")
		return
	}

	session := GetSession()

	m := map[string]interface{}{}

	var images [][]byte
	//var faceIds []string

	query := `SELECT image FROM user_images WHERE "username" = ?`

	iter := session.Query(query, *username).Iter()
	for iter.MapScan(m) {

		//faceID := m["face_id"].(string)
		photo := m["image"].([]byte)

		images = append(images, photo)
		// faceIds = append(faceIds, faceID)
	}

	fmt.Println("User images fetched successfully")
	// TODO: Define how to return images (what format)
}

// AddUserImage adds user image to database
func AddUserImage(username *string, image *[]byte) {

	session := GetSession()

	isUserExists := verifyIfUserExists(*username)

	// if username exists in database
	if isUserExists == false {
		fmt.Println("This user is not registered.")
		return
	}

	query := "INSERT INTO user_images (username, id, image) VALUES (?, ?, ?)"
	id := gocql.TimeUUID()

	if err := session.Query(query, *username, id, *image).
		Exec(); err != nil {
		fmt.Println("Error adding image. ", err)
	}

	fmt.Println("Image added successfully for user ", *username)
}

// Todo: Implement following

// Set face id for Image

// verifyIfUserExists checks if username exists in database
func verifyIfUserExists(username string) bool {

	session := GetSession()
	defer session.Close()
	var existingUsername string

	if err := session.Query(`SELECT "username" FROM users WHERE "username" = ?`, username).
		//Consistency(gocql.Quorum).
		Scan(&existingUsername); err != nil {

		fmt.Println("Error verifying user. ", err)
		//log. Fatal(err)
	}

	// if username exists in database
	if len(existingUsername) > 0 {
		return true
	}

	return false
}
