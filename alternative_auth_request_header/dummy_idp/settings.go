package main

import (
	"encoding/json"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"log"
)

type User struct {
	DisplayName string `json:"display_name"`
	Email       string `json:"email"`
	Login       string `json:"login"`
	Password    string `json:"hashed"`
}

type Settings struct {
	Users   []User `json:"users"`
	Cert    string `json:"cert"`
	Key     string `json:"key"`
	CA      string `json:"ca"`
	Backend string `json:"backend"`
}

// Serializes a Settings to JSON
func (s Settings) Output() {
	out, err := json.MarshalIndent(s, "", "\t")
	if err != nil {
		log.Print(err)
	}
	fmt.Println(string(out))
}

// Verify the Users Identity
func (s Settings) Verify(id string, pass string) bool {
	if id == "" {
		return false
	}
	if pass == "" {
		return false
	}

	// Dump to bytes so we can get a candidate to check
	password := []byte(pass)
	for _, user := range s.Users {
		err := bcrypt.CompareHashAndPassword([]byte(user.Password), password)
		if user.Login == id {
			// err == nil indicates there is a match
			return err == nil
		}
	}
	return false
}

// Load From File grabs the settings and marshals it to a JSON
func LoadFromFile(fn string) Settings {
	content, err := ioutil.ReadFile(fn)
	if err != nil {
		log.Fatal(err)
	}

	var settings Settings
	json.Unmarshal(content, &settings)
	return settings
}
