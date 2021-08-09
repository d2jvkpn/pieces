package main

import (
	"encoding/json"
	"fmt"
	"github.com/d2jvkpn/gorover/rover"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"os"
	"time"
)

const (
	HASHCOST     = 15
	CONN_ADDRESS = "localhost"
	CONN_PORT    = ":9000"
)

var DB = make(map[string]string, 5) // replace a database

func main() {
	var err error

	http.HandleFunc("/signup", UserSignup)
	http.HandleFunc("/login", UserLogin)

	log.Printf("Service %s%s", CONN_ADDRESS, CONN_PORT)
	err = http.ListenAndServe(CONN_ADDRESS+CONN_PORT, nil)
	if err != nil {
		log.Fatal(err)
	}
}

type Credentials struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type logWriter struct {
}

func (writer *logWriter) Write(bytes []byte) (int, error) {
	return fmt.Fprintf(os.Stderr, "%s %s",
		time.Now().Format("["+time.RFC3339+"] "), string(bytes))
}

func init() {
	log.SetFlags(0)
	log.SetOutput(new(logWriter))
}

func UserSignup(w http.ResponseWriter, r *http.Request) {
	var (
		ok             bool
		hashedPassword []byte
		err            error
	)

	creds, cms := new(Credentials), rover.NewCMS(200, "OK")
	defer cms.WJson(w)

	if err = json.NewDecoder(r.Body).Decode(creds); err != nil {
		cms.C = http.StatusBadRequest
		cms.M = "sign up failed: invalid parameters"
		return
	}
	r.Body.Close() //!!!

	if _, ok = DB[creds.Username]; ok {
		cms.C, cms.M = http.StatusConflict, "signUp failed: name was used"
		log.Printf("sign up failed: %s", creds.Username)
		return
	}

	// Salt and hash the password using the bcrypt algorithm
	hashedPassword, err = bcrypt.GenerateFromPassword(
		[]byte(creds.Password), HASHCOST)

	if err != nil {
		cms.C = http.StatusConflict
		cms.M = "sign up failed: to generate hashed password"
		return
	}

	DB[creds.Username] = string(hashedPassword)
	log.Println("sign up successed:", creds.Username)
}

func UserLogin(w http.ResponseWriter, r *http.Request) {
	// Parse and decode the request body into a new `Credentials` instance
	var (
		err    error
		result string
		ok     bool
	)

	creds, cms := new(Credentials), rover.NewCMS(200, "OK")
	defer cms.WJson(w)

	if err = json.NewDecoder(r.Body).Decode(creds); err != nil {
		cms.C, cms.M = http.StatusBadRequest, "login failed"
		return
	}
	r.Body.Close() //!!!

	if result, ok = DB[creds.Username]; !ok {
		cms.C = http.StatusBadRequest
		cms.M = "sign in failed: user not exists"
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(result), []byte(creds.Password))
	if err != nil {
		cms.C = http.StatusUnauthorized
		cms.M = "login failed: wrong password"
		log.Printf("login failed: %s", creds.Username)
		return
	}

	log.Println("login successed:", creds.Username)
}
