package main

import (
	"net/http"
	"time"

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

type session struct {
	userID       string
	lastActivity time.Time
}

var (
	dbUsers           = map[string]user{}    // user ID, user
	dbSessions        = map[string]session{} // session ID, user ID
	dbSessionsCleaned = time.Now()
)

const sessionLength int = 30

func validateLogin(userID, password string) bool {

	// Check that username exists.  This uses the the “comma ok” idiom discussed in
	// https://golang.org/doc/effective_go.html#maps.  The map lookup will return
	// a bool for the second return value, indicating if the key exists in the map.
	// If it does, the user is registered.
	userData, userExists := dbUsers[userID]

	// Check that the entered password matches the stored password.  Convert the provided
	// password from plain text to []byte.
	passwordMatches := bcrypt.CompareHashAndPassword(userData.Password, []byte(password))

	if !userExists || passwordMatches != nil {
		return false
	}

	return true
}

// GetEncryptedPassword takes a plain-text password and returns an encrypted version
// of it as byte slice, tegether with any error.
func getEncryptedPassword(password string) ([]byte, error) {

	// Use type conversion (https://tour.golang.org/basics/13) to get the string
	// into a []byte.
	passwordBytes := []byte(password)

	encryptedPassword, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.MinCost)
	if err != nil {
		// Return the error, preceeded by the zero-value of the Type.  The zero-value we
		// create using a https://golang.org/ref/spec#Composite_literals for []byte.
		return []byte{}, err
	}

	return encryptedPassword, nil
}

// GetUserAndSession gets a cookie named session, initializing one if missing.
// The cookie's value is a UUID being the session Id.
//
// If the user has `signed up` that means the userID is registered in dbUsers map,
// with a value being a user struct.  The user struct contains the user's data: userID,
// first and last names, encrypted password, and login status.
//
// If the user has `logged in, that means the sessionID is registered in dbSessions map
// with a value being a session struct that houses the user ID together with the time of
// last activity for that session.
//
// Together, we can lookup a user struct from the session ID and have the ability to
// expire the session if it goes stale.
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
	}
	cookie.MaxAge = sessionLength
	http.SetCookie(w, cookie)

	// If user is logged in (sessionID is registered in dbSessions), fetch the user
	// struct from dbUsers and return that together with the session Id.  First update
	// the .lastActivity field of the registered session struct.
	sessionID := cookie.Value
	if session, ok := dbSessions[sessionID]; ok {
		// Session stores userID as a field, use that to key dbUsers to get the user
		// struct
		preexistingUser := dbUsers[session.userID]
		preexistingUser.LoggedIn = true // TODO: may not be necessary

		// Refresh session's lastActivity field then update the session in dbSessions
		session.lastActivity = time.Now()
		dbSessions[sessionID] = session

		// Return the user and session
		return preexistingUser, sessionID
	}

	// Return same but where the user is a newly initialized one.  Its fields will have
	// zero values per https://golang.org/ref/spec#The_zero_value.
	var newuser user
	return newuser, sessionID
}

// CleanSessions deletes all sessions registered in dbSessions if last activity
// exceeds our sessionLength constant.
func cleanSessions() {
	for sessionID, session := range dbSessions {
		sessionTimeout := time.Second * time.Duration(sessionLength)
		if time.Now().Sub(session.lastActivity) > sessionTimeout {
			delete(dbSessions, sessionID)
		}
	}
	dbSessionsCleaned = time.Now()
}
