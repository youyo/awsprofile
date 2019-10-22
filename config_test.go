package awsprofile_test

import (
	"errors"
	"fmt"
	"log"
	"os"
	"testing"

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

func TestGetConfigsPath(t *testing.T) {
	if err := os.Setenv("AWS_CONFIG_FILE", "./tests/.aws/config"); err != nil {
		t.Fatal(err)
	}

	file, err := awsprofile.GetConfigsPath()
	if err != nil {
		t.Fatal(err)
	}

	if file != "./tests/.aws/config" {
		t.Error("file", file)
		t.Error("expect", "./tests/.aws/config")
		t.Fatal("config file path unmatched")
	}
}

func TestConfigs_Parse(t *testing.T) {
	if err := os.Setenv("AWS_CONFIG_FILE", "./tests/.aws/config"); err != nil {
		t.Fatal(err)
	}

	config := awsprofile.NewConfigs()
	file, err := awsprofile.GetConfigsPath()
	if err != nil {
		t.Fatal(err)
	}

	if err = config.Parse(file); err != nil {
		t.Fatal(err)
	}
}

func TestConfigs_GetRoleArn(t *testing.T) {
	os.Setenv("AWS_CONFIG_FILE", "./tests/.aws/config")
	config := awsprofile.NewConfigs()
	file, _ := awsprofile.GetConfigsPath()
	config.Parse(file)

	if value, err := config.GetRoleArn("bar"); err != nil {
		t.Fatal(err)
	} else if value != "arn:aws:iam::xxxxxxxxxxxx:role/bar" {
		t.Fatal(errors.New("Unmatched RoleArn"))
	}
}

func TestConfigs_GetSourceProfile(t *testing.T) {
	os.Setenv("AWS_CONFIG_FILE", "./tests/.aws/config")
	config := awsprofile.NewConfigs()
	file, _ := awsprofile.GetConfigsPath()
	config.Parse(file)

	if value, err := config.GetSourceProfile("bar"); err != nil {
		t.Fatal(err)
	} else if value != "foo" {
		t.Fatal(errors.New("Unmatched SourceProfile"))
	}
}

func TestConfigs_GetCredentialSource(t *testing.T) {
	os.Setenv("AWS_CONFIG_FILE", "./tests/.aws/config")
	config := awsprofile.NewConfigs()
	file, _ := awsprofile.GetConfigsPath()
	config.Parse(file)

	if value, err := config.GetCredentialSource("bar"); err != nil {
		t.Fatal(err)
	} else if value != "Environment" {
		t.Fatal(errors.New("Unmatched CredentialSource"))
	}
}

func TestConfigs_GetRoleSessionName(t *testing.T) {
	os.Setenv("AWS_CONFIG_FILE", "./tests/.aws/config")
	config := awsprofile.NewConfigs()
	file, _ := awsprofile.GetConfigsPath()
	config.Parse(file)

	if value, err := config.GetRoleSessionName("bar"); err != nil {
		t.Fatal(err)
	} else if value != "foobar" {
		t.Fatal(errors.New("Unmatched RoleSessionName"))
	}
}

func TestConfigs_GetMfaSerial(t *testing.T) {
	os.Setenv("AWS_CONFIG_FILE", "./tests/.aws/config")
	config := awsprofile.NewConfigs()
	file, _ := awsprofile.GetConfigsPath()
	config.Parse(file)

	if value, err := config.GetMfaSerial("bar"); err != nil {
		t.Fatal(err)
	} else if value != "arn:aws:iam::123456789012:mfa/foo" {
		t.Fatal(errors.New("Unmatched MfaSerial"))
	}
}

func TestConfigs_GetDurationSeconds(t *testing.T) {
	os.Setenv("AWS_CONFIG_FILE", "./tests/.aws/config")
	config := awsprofile.NewConfigs()
	file, _ := awsprofile.GetConfigsPath()
	config.Parse(file)

	if value, err := config.GetDurationSeconds("bar"); err != nil {
		t.Fatal(err)
	} else if value != 43200 {
		t.Fatal(errors.New("Unmatched DurationSeconds"))
	}
}

func TestConfigs_GetAwsSessionToken(t *testing.T) {
	os.Setenv("AWS_CONFIG_FILE", "./tests/.aws/config")
	config := awsprofile.NewConfigs()
	file, _ := awsprofile.GetConfigsPath()
	config.Parse(file)

	if value, err := config.GetAwsSessionToken("bar"); err != nil {
		t.Fatal(err)
	} else if value != "AQoEXAMPLEH4aoAH0gNCAPy..." {
		t.Fatal(errors.New("Unmatched AwsSessionToken"))
	}
}

func TestConfigs_GetExternalID(t *testing.T) {
	os.Setenv("AWS_CONFIG_FILE", "./tests/.aws/config")
	config := awsprofile.NewConfigs()
	file, _ := awsprofile.GetConfigsPath()
	config.Parse(file)

	if value, err := config.GetExternalID("bar"); err != nil {
		t.Fatal(err)
	} else if value != 12345 {
		t.Fatal(errors.New("Unmatched ExternalID"))
	}
}

func TestConfigs_GetCaBundle(t *testing.T) {
	os.Setenv("AWS_CONFIG_FILE", "./tests/.aws/config")
	config := awsprofile.NewConfigs()
	file, _ := awsprofile.GetConfigsPath()
	config.Parse(file)

	if value, err := config.GetCaBundle("bar"); err != nil {
		t.Fatal(err)
	} else if value != "dev/apps/ca-certs/cabundle-2019mar05.pem" {
		t.Fatal(errors.New("Unmatched CaBundle"))
	}
}

func TestConfigs_GetCliFollowUrlparam(t *testing.T) {
	os.Setenv("AWS_CONFIG_FILE", "./tests/.aws/config")
	config := awsprofile.NewConfigs()
	file, _ := awsprofile.GetConfigsPath()
	config.Parse(file)

	if value, err := config.GetCliFollowUrlparam("bar"); err != nil {
		t.Fatal(err)
	} else if value != "false" {
		t.Fatal(errors.New("Unmatched CliFollowUrlparam"))
	}
}

func TestConfigs_GetCliTimestampFormat(t *testing.T) {
	os.Setenv("AWS_CONFIG_FILE", "./tests/.aws/config")
	config := awsprofile.NewConfigs()
	file, _ := awsprofile.GetConfigsPath()
	config.Parse(file)

	if value, err := config.GetCliTimestampFormat("bar"); err != nil {
		t.Fatal(err)
	} else if value != "iso8601" {
		t.Fatal(errors.New("Unmatched CliTimestampFormat"))
	}
}

func TestConfigs_GetCredentialProcess(t *testing.T) {
	os.Setenv("AWS_CONFIG_FILE", "./tests/.aws/config")
	config := awsprofile.NewConfigs()
	file, _ := awsprofile.GetConfigsPath()
	config.Parse(file)

	if value, err := config.GetCredentialProcess("bar"); err != nil {
		t.Fatal(err)
	} else if value != "/opt/bin/awscreds-retriever --username susan" {
		t.Fatal(errors.New("Unmatched CredentialProcess"))
	}
}

func TestConfigs_GetWebIdentityTokenFile(t *testing.T) {
	os.Setenv("AWS_CONFIG_FILE", "./tests/.aws/config")
	config := awsprofile.NewConfigs()
	file, _ := awsprofile.GetConfigsPath()
	config.Parse(file)

	if value, err := config.GetWebIdentityTokenFile("bar"); err != nil {
		t.Fatal(err)
	} else if value != "/path/to/a/token" {
		t.Fatal(errors.New("Unmatched WebIdentityTokenFile"))
	}
}

func TestConfigs_GetOutput(t *testing.T) {
	os.Setenv("AWS_CONFIG_FILE", "./tests/.aws/config")
	config := awsprofile.NewConfigs()
	file, _ := awsprofile.GetConfigsPath()
	config.Parse(file)

	if value, err := config.GetOutput("bar"); err != nil {
		t.Fatal(err)
	} else if value != "json" {
		t.Fatal(errors.New("Unmatched Output"))
	}
}

func TestConfigs_GetRegion(t *testing.T) {
	os.Setenv("AWS_CONFIG_FILE", "./tests/.aws/config")
	config := awsprofile.NewConfigs()
	file, _ := awsprofile.GetConfigsPath()
	config.Parse(file)

	if value, err := config.GetRegion("bar"); err != nil {
		t.Fatal(err)
	} else if value != "ap-northeast-1" {
		t.Fatal(errors.New("Unmatched Region"))
	}
}
