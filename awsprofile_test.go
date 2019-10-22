package awsprofile_test

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/youyo/awsprofile"
)

func ExampleAwsProfile_ProfileNames() {
	// if you use non-default configuration file path
	if err := os.Setenv("AWS_SHARED_CREDENTIALS_FILE", "./tests/.aws/credentials"); err != nil {
		log.Fatal(err)
	}
	if err := os.Setenv("AWS_CONFIG_FILE", "./tests/.aws/config"); err != nil {
		log.Fatal(err)
	}

	// New
	awsProfile := awsprofile.New()

	// Parse AWS_SHARED_CREDENTIALS_FILE and AWS_CONFIG_FILE
	if err := awsProfile.Parse(); err != nil {
		log.Fatal(err)
	}

	// Get profile names
	profiles, err := awsProfile.ProfileNames()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(profiles)
	// Output: [default foo bar]
}

func ExampleAwsProfile_GetCredentials() {
	// if you use non-default configuration file path
	if err := os.Setenv("AWS_SHARED_CREDENTIALS_FILE", "./tests/.aws/credentials"); err != nil {
		log.Fatal(err)
	}
	if err := os.Setenv("AWS_CONFIG_FILE", "./tests/.aws/config"); err != nil {
		log.Fatal(err)
	}

	// New
	awsProfile := awsprofile.New()

	// Parse AWS_SHARED_CREDENTIALS_FILE and AWS_CONFIG_FILE
	if err := awsProfile.Parse(); err != nil {
		log.Fatal(err)
	}

	credentials := awsProfile.GetCredentials()

	profiles, err := credentials.ProfileNames()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(profiles)
	// Output: [default foo]
}

func ExampleAwsProfile_GetConfigs() {
	// if you use non-default configuration file path
	if err := os.Setenv("AWS_SHARED_CREDENTIALS_FILE", "./tests/.aws/credentials"); err != nil {
		log.Fatal(err)
	}
	if err := os.Setenv("AWS_CONFIG_FILE", "./tests/.aws/config"); err != nil {
		log.Fatal(err)
	}

	// New
	awsProfile := awsprofile.New()

	// Parse AWS_SHARED_CREDENTIALS_FILE and AWS_CONFIG_FILE
	if err := awsProfile.Parse(); err != nil {
		log.Fatal(err)
	}

	configs := awsProfile.GetConfigs()

	profiles, err := configs.ProfileNames()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(profiles)
	// Output: [default bar]
}

func TestAwsProfile_Parse(t *testing.T) {
	// if you use non-default configuration file path
	if err := os.Setenv("AWS_SHARED_CREDENTIALS_FILE", "./tests/.aws/credentials"); err != nil {
		log.Fatal(err)
	}
	if err := os.Setenv("AWS_CONFIG_FILE", "./tests/.aws/config"); err != nil {
		log.Fatal(err)
	}

	// New
	awsProfile := awsprofile.New()

	// Parse AWS_SHARED_CREDENTIALS_FILE and AWS_CONFIG_FILE
	if err := awsProfile.Parse(); err != nil {
		log.Fatal(err)
	}
}
