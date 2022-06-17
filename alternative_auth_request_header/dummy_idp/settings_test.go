package main

import (
	"golang.org/x/crypto/bcrypt"
	"testing"
)

var DemoSettings = GenerateExample()

// Generates a full fledged example with a salt and a hashed password
func GenerateExample() Settings {
	password := []byte("password")
	hashed, _ := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)

	var user1 User
	user1.DisplayName = "Paul B"
	user1.Email = "test1@bastide.org"
	user1.Login = "testuser1"
	user1.Password = string(hashed)

	var user2 User
	user2.DisplayName = "Peter B"
	user2.Email = "test2@bastide.org"
	user2.Login = "testuser2"
	user2.Password = string(hashed)

	var settings Settings
	settings.Users = make([]User, 2)
	settings.Users[0] = user1
	settings.Users[1] = user2

	settings.Output()

	return settings
}

func TestBadUser(t *testing.T) {
	got := DemoSettings.Verify("notreal", "badpass")
	if got {
		t.Errorf("BadUser = %t; wanted false", got)
	}
}

func TestBadPassword(t *testing.T) {
	got := DemoSettings.Verify("testuser2", "badpass")
	if got {
		t.Errorf("BadPassword = %t; wanted false", got)
	}
}

func TestBadUserGoodPassword(t *testing.T) {
	got := DemoSettings.Verify("fred", "password")
	if got {
		t.Errorf("BadUserGoodPassword = %t; wanted false", got)
	}
}

func TestGoodUserGoodPassword1(t *testing.T) {
	got := DemoSettings.Verify("testuser2", "password")
	if !got {
		t.Errorf("GoodUserGoodPassword = %t; wanted true", got)
	}
}

func TestGoodUserGoodPassword2(t *testing.T) {
	got := DemoSettings.Verify("testuser1", "password")
	if !got {
		t.Errorf("GoodUserGoodPassword2 = %t; wanted true", got)
	}
}
