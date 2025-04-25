package adguard

import (
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"
)

// testADG - Helper function to create a test ADG instance
func testADG(test_install ...bool) *ADG {
	// use an environment variable for the backend URL
	host := os.Getenv("ADGUARD_HOST")
	if host == "" {
		// default to localhost if not set
		host = "localhost:8080"
	}

	username := os.Getenv("ADGUARD_USERNAME")
	if username == "" {
		// default to admin if not set
		username = "admin"
	}
	password := os.Getenv("ADGUARD_PASSWORD")
	if password == "" {
		// default to a test password if not set
		password = "SecretP@ssw0rd"
	}

	scheme := os.Getenv("ADGUARD_SCHEME")
	if scheme == "" {
		// default to http if not set
		scheme = "http"
	}
	timeout := os.Getenv("ADGUARD_TIMEOUT")
	if timeout == "" {
		// default to 30 seconds if not set
		timeout = "30"
	}
	// convert timeout to int
	timeoutInt, err := strconv.Atoi(timeout)
	if err != nil {
		panic("Invalid timeout value, must be an integer")
	}

	// if test_install is true, adjust the port as
	// it's for the first time setup
	if len(test_install) > 0 && test_install[0] {
		host = "localhost:3000"
	}

	// create a new ADG instance
	adg, err := NewClient(&host, &username, &password, &scheme, &timeoutInt)
	if err != nil {
		panic(err)
	}
	return adg
}

// testADGWithNewRequestError - Helper function to create a test ADG instance with an initialization error
func testADGWithNewRequestError() *ADG {
	adg := testADG()
	// invalid URL to trigger an error
	adg.HostURL = "http://%invalid-url"

	return adg
}

// testADGWithDoRequestError - Helper function to create a test ADG instance with a request error
func testADGWithDoRequestError(test_install ...bool) *ADG {
	adg := testADG()

	// if test_install is true, adjust the port as
	// it's for the first time setup
	if len(test_install) > 0 && test_install[0] {
		adg.HostURL = "http://localhost:8000"
	} else {

		// set an invalid password to trigger a 403 error
		adg.Auth.Password = "wrongpassword"
	}
	return adg
}

// testADGWithInvalidJSON - Helper function to create a test ADG instance with an invalid JSON response
func testADGWithInvalidJSON(t *testing.T) (*ADG, *httptest.Server) {
	t.Helper()

	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{")) // intentionally invalid JSON
	})

	server := httptest.NewServer(handler)

	adg := testADG()
	adg.HostURL = server.URL
	adg.HTTPClient = server.Client()

	return adg, server
}
