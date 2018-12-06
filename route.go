package main

import (
	"html/template"
	"net/http"
	"time"
)

// Tpl is a pointer to template.Template in the html package. Create the
// var for it.
var tpl *template.Template

// Initialize the tpl var to load all our HTML templates that live under
// the top-level templates dir.
func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

// Having staged the tpl var containing all our templates, we move on to
// defining each of our routes registered in main.go.

// Favicon serves our favicon.ico file if requested by a client browser
func favicon(w http.ResponseWriter, req *http.Request) {
	http.ServeFile(w, req, "public/images/favicon.ico")
}

// Index is the home page Handler function
func index(w http.ResponseWriter, req *http.Request) {
	// Get user and session state
	userData, _ := getUserAndSession(w, req)

	// Redirect to login page if user is logged out
	if !userData.LoggedIn {
		http.Redirect(w, req, "/login", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "index.gohtml", userData)
}

// Signin encrypts the provided password and registers a user struct
// in the dbUsers map.
func signin(w http.ResponseWriter, req *http.Request) {

	// Redirect to home page if already signed in. We run getUserAndSession()
	// at the beginning of each route to capture user and session state
	// for processing.
	userData, _ := getUserAndSession(w, req)
	if _, ok := dbUsers[userData.UserID]; ok {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	// Process form submission
	if req.Method == http.MethodPost {

		userID := req.FormValue("userid")
		password := req.FormValue("password")
		first := req.FormValue("first")
		last := req.FormValue("last")

		// Check that the userID isn't already taken.  If it is, reply to
		// the request with the specified error message and HTTP code.
		if _, ok := dbUsers[userID]; ok {
			http.Error(w, "Username already taken", http.StatusForbidden)
			return
		}

		// Ecrypt password before storing it.  If the encryption errs, respond with
		// an Internal Server Error.
		encryptedPassword, err := getEncryptedPassword(password)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		// Register user in dbUsers map.  That is the test for being signed up,
		// so set the `signedUp` field to true when building the composite literal
		// for the user{}.
		newuser := user{userID, first, last, encryptedPassword, false}
		dbUsers[userID] = newuser

		// Redirect to home page after signup
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	// Render page if not a POST
	tpl.ExecuteTemplate(w, "signin.gohtml", nil)
}

func login(w http.ResponseWriter, req *http.Request) {

	// Redirect to home page if already logged in
	userData, sessionID := getUserAndSession(w, req)
	if userData.LoggedIn {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	// process form submission
	if req.Method == http.MethodPost {
		userID := req.FormValue("userid")
		password := req.FormValue("password")

		// ValidateLogin checks the provided password against the hashed password stored
		// in dbUsers for that user.  Respond forbidden if not matching.
		ok := validateLogin(userID, password)
		if !ok {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}

		// Add sessionID to dbSessions to effect the login
		s := session{userID, time.Now()}
		dbSessions[sessionID] = s

		// Redirect to home page after login
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "login.gohtml", nil)
}

func logout(w http.ResponseWriter, req *http.Request) {

	userData, sessionID := getUserAndSession(w, req)

	// Redirect to home page if already logged out
	if !userData.LoggedIn {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	// Unregister the sessionId in dbSession and remove the session cookie.
	delete(dbSessions, sessionID)
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "deleteNow",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)

	// Garbage collect stale sessions registered in dbSessions.
	sessionTimout := time.Second * time.Duration(sessionLength)
	if time.Now().Sub(dbSessionsCleaned) > sessionTimout {
		go cleanSessions()
	}

	http.Redirect(w, req, "/", http.StatusSeeOther)
	return
}

// ResetPassword is a page dedicated to just that
func resetPassword(w http.ResponseWriter, req *http.Request) {

	userData, _ := getUserAndSession(w, req)

	if !userData.LoggedIn {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "resetPassword.gohtml", userData)
}

// Other is a test secondary Handler function, registered in DefaultServeMux.
func other(w http.ResponseWriter, req *http.Request) {

	userData, _ := getUserAndSession(w, req)

	// Redirect to home page if logged out
	if !userData.LoggedIn {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}

	tpl.ExecuteTemplate(w, "other.gohtml", userData)
}
