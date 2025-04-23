package adguard

import (
	"os"
	"strconv"
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
		// default to 10 seconds if not set
		timeout = "10"
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
