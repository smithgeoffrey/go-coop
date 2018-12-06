package main

import (
	"html/template"
	"net/http"
)

// Tpl is a pointer to template.Template in the html package.
// Create the var for a tpl.
var tpl *template.Template

// Initialize the tpl var, to load all our HTML templates
// that live under the top-level templates dir.
func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

// Having staged the tpl var with all our templates, we move on
// to defining each of our routes.

// Index is the home page Handler function
func index(w http.ResponseWriter, req *http.Request) {
	userData, _ := getUserAndSession(w, req)
	tpl.ExecuteTemplate(w, "index.gohtml", userData)
}

// Signup encrypts the provided password and registers a user struct
// in the dbUsers map.
func signup(w http.ResponseWriter, req *http.Request) {

	// Redirect to home page if already signed up
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
	tpl.ExecuteTemplate(w, "signup.gohtml", nil)
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
		dbSessions[sessionID] = userID

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

	// Remove sessionId
	delete(dbSessions, sessionID)

	cookie := &http.Cookie{
		Name:   "session",
		Value:  "deleteNow",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)

	http.Redirect(w, req, "/", http.StatusSeeOther)
	return
}

// Other is a test secondary Handler function, registered in DefaultServeMux.
func other(w http.ResponseWriter, req *http.Request) {

	// get cookie
	c, err := req.Cookie("session")
	if err != nil {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	un, ok := dbSessions[c.Value]
	if !ok {
		http.Redirect(w, req, "/", http.StatusSeeOther)
		return
	}
	u := dbUsers[un]
	tpl.ExecuteTemplate(w, "other.gohtml", u)

}
