package auth

import (
	"net/http"
	"reflect"
	"testing"
)

func TestGetApiKey(t *testing.T) {
	header := http.Header{}
	header.Set("Authorization", "ApiKey example")
	got, err := GetAPIKey(header)
	want := "example"
	if err != nil {
		t.Fatalf("Could not call func")
	}
	if !reflect.DeepEqual(want, got) {
		t.Fatalf("expected: %v, got: %v", want, got)
	}
}
