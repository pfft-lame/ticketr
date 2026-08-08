package auth

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestMakeJWT(t *testing.T) {
	userId, tokenSecret := uuid.New(), "secret123"

	tokenStr, err := MakeJWT(userId, tokenSecret, time.Hour)
	if err != nil {
		t.Fatalf("expected no err in MakeJWT, got %v", err)
	}

	claims, err := ValidateJWT(tokenStr, tokenSecret)
	if err != nil {
		t.Fatalf("expected no err in ValidateJWT, got %v", err)
	}

	if userId != claims.UserId {
		t.Fatalf("expected userId and authId to be same, userId=%v authId=%v", userId, claims.UserId)
	}
}
