package model

import (
	"testing"
)

func TestUser(t *testing.T) *User {
	return &User{
		Email:    "user@example.org",
		Password: "strongpassword",
	}
}

func TestPreferences(t *testing.T) Preferences {
	return Preferences{
		TrackedTokens: []string{"ETH"},
		FiatCurrency:  "USD",
	}
}
