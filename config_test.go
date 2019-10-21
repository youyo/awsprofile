package awsprofile_test

import (
	"fmt"
	"log"
	"os"

	"github.com/youyo/awsprofile"
)

func ExampleConfigs_ProfileNames() {
	// if you use non-default configuration file path
	if err := os.Setenv("AWS_CONFIG_FILE", "./tests/.aws/config"); err != nil {
		log.Fatal(err)
	}

	// New
	config := awsprofile.NewConfigs()

	// get config file path
	file, err := awsprofile.GetConfigsPath()
	if err != nil {
		log.Fatal(err)
	}

	// Parse AWS_CONFIG_FILE
	if err := config.Parse(file); err != nil {
		log.Fatal(err)
	}

	// Get profile names
	profiles, err := config.ProfileNames()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(profiles)
	// Output: [default bar]
}
