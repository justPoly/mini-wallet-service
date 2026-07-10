package services

import "testing"

func TestIdempotency(t *testing.T) {

	key := "abc123"

	if Exists(key) {
		t.Error("Key should not exist initially")
	}

	Save(key)

	if !Exists(key) {
		t.Error("Key should exist after saving")
	}
}