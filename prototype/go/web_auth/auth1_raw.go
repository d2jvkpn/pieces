package main

import (
	"encoding/json"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
)

const gHashCost = 16

var gDB = make(map[string]string, 5)

func main() {
	http.HandleFunc("/signin", Signin)
	http.HandleFunc("/signup", Signup)

	log.Fatal(http.ListenAndServe(":8000", nil))
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func Signup(w http.ResponseWriter, r *http.Request) {
	var err error
	var hashedPassword []byte

	creds := new(Credentials)
	if err = json.NewDecoder(r.Body).Decode(creds); err != nil {
		log.Println("SignUp failed: invalid parameters")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Salt and hash the password using the bcrypt algorithm
	// The second argument is the cost of hashing, which we arbitrarily set as
	//   8 (this value can be more or less, depending on the computing power
	//   you wish to utilize)
	hashedPassword, err = bcrypt.GenerateFromPassword([]byte(creds.Password),
		gHashCost)

	if err != nil {
		log.Println("SignUp failed: to generate hashed password")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// save username and hashed password
	gDB[creds.Username] = string(hashedPassword)
	log.Println("SignUp successed:", creds.Username)
}

func Signin(w http.ResponseWriter, r *http.Request) {
	// Parse and decode the request body into a new `Credentials` instance
	var (
		err    error
		result string
		ok     bool
	)

	creds := new(Credentials)
	// If there is something wrong with the request body, return a 400 status
	if err = json.NewDecoder(r.Body).Decode(creds); err != nil {
		log.Println("SignIn failed")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get the existing entry present in the database for the given username
	if result, ok = gDB[creds.Username]; !ok {
		log.Println("SignIn failed: user not exists", creds.Username)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Compare the stored hashed password, with the hashed version of the
	//   password that was received
	err = bcrypt.CompareHashAndPassword([]byte(result), []byte(creds.Password))
	// If the two passwords don't match, return a 401 status
	if err != nil {
		log.Println("SignIn failed: wrong password for", creds.Username)
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	log.Println("SignIn successed:", creds.Username)
	// If we reach this point, that means the users password was correct,
	//   and that they are authorized
}
