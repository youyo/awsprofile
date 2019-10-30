package awsprofile_test

import (
	"errors"
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
		t.Error(err)
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

func TestCredentials_GetAwsAccessKeyID(t *testing.T) {
	os.Setenv("AWS_SHARED_CREDENTIALS_FILE", "./tests/.aws/credentials")
	creds := awsprofile.NewCredentials()
	file, _ := awsprofile.GetCredentialsPath()
	creds.Parse(file)

	if awsAccessKeyID, err := creds.GetAwsAccessKeyID("foo"); err != nil {
		t.Fatal(err)
	} else if awsAccessKeyID != "ACCESS-2-XXXXXXXXXXXXX" {
		t.Fatal(errors.New("Unmatched AwsAccessKeyID"))
	}
}

func TestCredentials_GetAwsSecretAccessKey(t *testing.T) {
	os.Setenv("AWS_SHARED_CREDENTIALS_FILE", "./tests/.aws/credentials")
	creds := awsprofile.NewCredentials()
	file, _ := awsprofile.GetCredentialsPath()
	creds.Parse(file)

	if awsSecretAccessKey, err := creds.GetAwsSecretAccessKey("foo"); err != nil {
		t.Fatal(err)
	} else if awsSecretAccessKey != "SECRET-2-XXXXXXXXXXXXX" {
		t.Fatal(errors.New("Unmatched AwsSecretAccessKey"))
	}
}

func TestCredential_GetAwsAccessKeyID(t *testing.T) {
	os.Setenv("AWS_SHARED_CREDENTIALS_FILE", "./tests/.aws/credentials")
	os.Setenv("AWS_CONFIG_FILE", "./tests/.aws/config")

	awsProfile := awsprofile.New()
	awsProfile.Parse()

	if ok, cred := awsProfile.IsCredential("foo"); ok {
		awsAccessKeyID := cred.GetAwsAccessKeyID()
		if awsAccessKeyID != "ACCESS-2-XXXXXXXXXXXXX" {
			t.Fatal(errors.New("Unmatched AwsAccessKeyID"))
		}
	}
}

func TestCredential_GetAwsSecretAccessKey(t *testing.T) {
	os.Setenv("AWS_SHARED_CREDENTIALS_FILE", "./tests/.aws/credentials")
	os.Setenv("AWS_CONFIG_FILE", "./tests/.aws/config")

	awsProfile := awsprofile.New()
	awsProfile.Parse()

	if ok, cred := awsProfile.IsCredential("foo"); ok {
		awsSecretAccessKey := cred.GetAwsSecretAccessKey()
		if awsSecretAccessKey != "SECRET-2-XXXXXXXXXXXXX" {
			t.Fatal(errors.New("Unmatched AwsSecretAccessKey"))
		}
	}
}
