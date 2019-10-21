package awsprofile_test

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/youyo/awsprofile"
)

func ExampleCredentials_ProfileNames() {
	// if you use non-default configuration file path
	if err := os.Setenv("AWS_SHARED_CREDENTIALS_FILE", "./tests/.aws/credentials"); err != nil {
		log.Fatal(err)
	}

	// New
	creds := awsprofile.NewCredentials()

	// get credentials file path
	file, err := awsprofile.GetCredentialsPath()
	if err != nil {
		log.Fatal(err)
	}

	// Parse AWS_SHARED_CREDENTIALS_FILE
	if err := creds.Parse(file); err != nil {
		log.Fatal(err)
	}

	// Get profile names
	profiles, err := creds.ProfileNames()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(profiles)
	// Output: [default foo]
}

func TestGetCredentialsPath(t *testing.T) {
	if err := os.Setenv("AWS_SHARED_CREDENTIALS_FILE", "./tests/.aws/credentials"); err != nil {
		t.Fatal(err)
	}

	file, err := awsprofile.GetCredentialsPath()
	if err != nil {
		t.Fatal(err)
	}

	if file != "./tests/.aws/credentials" {
		t.Error("file", file)
		t.Error("expect", "./tests/.aws/credentials")
		t.Fatal("credentials file path unmatched")
	}
}

func TestCredentials_Parse(t *testing.T) {
	if err := os.Setenv("AWS_SHARED_CREDENTIALS_FILE", "./tests/.aws/credentials"); err != nil {
		t.Fatal(err)
	}

	creds := awsprofile.NewCredentials()
	file, err := awsprofile.GetCredentialsPath()
	if err != nil {
		t.Fatal(err)
	}

	if err = creds.Parse(file); err != nil {
		t.Fatal(err)
	}
}
