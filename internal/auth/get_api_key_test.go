package auth

import (
	"testing"
)

func TestNoHeader(t *testing.T) {
	ApiKey, err := GetAPIKey(nil)
	if err != ErrNoAuthHeaderIncluded {	
		t.Errorf("Expected ErrNoAuthHeaderIncluded, got %v", err)
	}
	t.Logf("ApiKey: %s, Error: %v\n", ApiKey, err)
}