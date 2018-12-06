package main

import (
	"net/http"

	uuid "github.com/satori/go.uuid"
	"golang.org/x/crypto/bcrypt"
)

type user struct {
	UserID   string
	First    string
	Last     string
	Password []byte
	LoggedIn bool
}

var (
	dbUsers    = map[string]user{}   // user ID, user
	dbSessions = map[string]string{} // session ID, user ID
)

func validateLogin(userID, password string) bool {

	// Check that username exists
	userData, userExists := dbUsers[userID]

	// Check that the entered password matches the stored password
	passwordMatches := bcrypt.CompareHashAndPassword(userData.Password, []byte(password))

	if !userExists || passwordMatches != nil {
		return false
	}

	return true
}

func getEncryptedPassword(password string) ([]byte, error) {
	passwordBytes := []byte(password)
	encryptedPassword, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.MinCost)
	if err != nil {
		return []byte{}, err
	}

	return encryptedPassword, nil
}

// GetUserAndSession gets a cookie named session, initializing one if missing.
// The cookie's value is a UUID being the session Id.
//
// If the user has `signed up` that means the userID is registered in dbUsers map,
// with a value being a user struct.  If the user has `logged in, that means the
// the sessionID is registered in dbSessions map with a value being the user ID.
// Together, we can lookup a user struct from the session ID.
//
// The user struct contains the user's data: userID, first and last names, encrypted
// password, and login and signup status.
//
func getUserAndSession(w http.ResponseWriter, req *http.Request) (user, string) {

	// Ensure a cookie exists named session, having a value that is a UUID.
	cookie, err := req.Cookie("session")
	if err != nil {
		sessionID, _ := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "session",
			Value: sessionID.String(),
		}
		http.SetCookie(w, cookie)
	}
	sessionID := cookie.Value

	// Initialize a user struct.  Its fields will have zero values per
	// https://golang.org/ref/spec#The_zero_value.  Return that unless the
	// user is logged in, in which case fetch the user struct from dbUsers
	// and return that instead.  Either way, also return session Id.
	var newuser user
	if userID, ok := dbSessions[sessionID]; ok {
		preexistingUser := dbUsers[userID]
		preexistingUser.LoggedIn = true
		return preexistingUser, sessionID
	}

	return newuser, sessionID
}
