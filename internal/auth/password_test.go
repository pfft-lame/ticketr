package auth

import "testing"

func TestHashPassword(t *testing.T) {
	password := "mySecretPassword123"
	hashedPassword, err := HashPassword(password)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	err = CheckPasswordHash(password, hashedPassword)
	if err != nil {
		t.Fatalf("expected hashed password valid, got %v", err)
	}
}
