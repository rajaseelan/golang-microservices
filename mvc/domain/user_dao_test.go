package domain

import (
	"net/http"
	"testing"
)

func TestGetUserNoUserFound(t *testing.T) {
	// initialization

	// Execution
	user, err := GetUser(0)

	// Validation
	if user != nil {
		t.Error("We were not expecting a user with id 0")
	}

	if err == nil {
		t.Error("We were expecting and error when user id == 0")
	}

	if err.StatusCode != http.StatusNotFound {
		t.Error("We were expecting 404 when user not found")
	}

	if err.Code != "not_found" {
		t.Error("Expecting not_found for code text when 404")
	}

	if err.Message != "user 0 was not found" {
		t.Error("Expecting user 0 was not found")
	}
}

func TestGetUserNoError(t *testing.T) {
	user, err := GetUser(123)

	if err != nil {
		t.Error("There should be no error gettings userId 123")
	}

	if user.Id != 123 {
		t.Error("The userId returned should be 123")
	}

	if user.FirstName != "Fede" {
		t.Error("User.FirstName should be Fede")
	}

	if user.LastName != "Leon" {
		t.Error("The lastName should be Leon")
	}

	if user.Email != "f@w.com" {
		t.Error("The user.Email should be f@w.com")
	}

}
