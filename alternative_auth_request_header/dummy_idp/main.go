package main

import (
	"fmt"
	"net/http"
)

var setting = LoadFromFile("settings/settings.json")

// Loads the settings and starts up the server on port 8080
func main() {
	// Log out for debugging purposes
	setting.Output()

	challenge := http.HandlerFunc(handleChallengeRequest)
	webLogin := http.HandlerFunc(handleWebLoginRequest)
	defaultHandler := http.HandlerFunc(defaultRequest)

	// There are two paths and the all other route is catch all
	http.Handle("/challenges/oauth/authorize", challenge)
	http.Handle("/web-login/oauth/authorize", webLogin)
	http.Handle("/", defaultHandler)
	http.ListenAndServe(":8080", nil)
}

// Support for the Challenge Request
func handleChallengeRequest(w http.ResponseWriter, r *http.Request) {
	u, p, ok := r.BasicAuth()
	if !ok {
		w.Header().Set("WWW-Authenticate", `Basic realm="Demo"`)
		w.WriteHeader(401)
		return
	}

	if check := setting.Verify(u, p); !check {
		fmt.Printf("invalid username or password '%s' '%s'", u, p)
		w.Header().Set("WWW-Authenticate", `Basic realm="Demo"`)
		w.WriteHeader(401)
		return
	}

	for k, v := range r.URL.Query() {
		fmt.Println(k, " ", v)
	}

	// Dump the Query Parameters that were passed in.
	fmt.Println("Query Parameters:")
	for k, v := range r.URL.Query() {
		fmt.Println(k, " ", v)
	}

	// Another reason this is dev only (this should be hardened)
	queryParameters := r.URL.Query()
	clientId := queryParameters["client_id"][0]
	codeChallenge := queryParameters["code_challenge"][0]

	location := ChallengeRequest(setting, clientId, codeChallenge, u)
	w.Header().Add("Location", location)
	w.WriteHeader(302)
	return
}

// Support for the Web Login Request
func handleWebLoginRequest(w http.ResponseWriter, r *http.Request) {
	// Use Basic Auth to login the User (this is a dev-test)
	u, p, ok := r.BasicAuth()
	if !ok {
		w.Header().Set("WWW-Authenticate", `Basic realm="Demo"`)
		w.WriteHeader(401)
		return
	}

	if check := setting.Verify(u, p); !check {
		fmt.Printf("invalid username or password '%s' '%s'", u, p)
		w.Header().Set("WWW-Authenticate", `Basic realm="Demo"`)
		w.WriteHeader(401)
		return
	}

	// Dump the Query Parameters that were passed in.
	fmt.Println("Query Parameters:")
	for k, v := range r.URL.Query() {
		fmt.Println(k, " ", v)
	}

	// Another reason this is dev only (this should be hardened)
	queryParameters := r.URL.Query()
	clientId := queryParameters["client_id"][0]
	idp := queryParameters["idp"][0]
	state := queryParameters["state"][0]

	location := WebLoginRequest(setting, clientId, idp, state, u)
	w.Header().Add("Location", location)
	w.WriteHeader(302)
	return
}

// Catch All to hide business logic/other aspects of the app
func defaultRequest(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(401)
	return
}
