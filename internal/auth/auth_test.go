package auth

import (
	"errors"
	"net/http"
	"reflect"
	"testing"
)

func TestGetAPIAuthFailure(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "http://", nil)
	_, err := GetAPIKey(req.Header)

	if !errors.Is(err, ErrNoAuthHeaderIncluded) {
		t.Fatalf("Expected: %v, got: %v", ErrNoAuthHeaderIncluded, err)
	}
}

func TestGetAPI(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "http://", nil)
	req.Header.Add("Authorization", "")
	req.Header.Add("ApiKey", "test")

	key, _ := GetAPIKey(req.Header)
	if !reflect.DeepEqual(key, "test") {
		t.Fatalf("Expected: test, got: %v", key)
	}
}
