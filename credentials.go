package awsprofile

import (
	"errors"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	ini "gopkg.in/ini.v1"
)

// constant
const (
	AwsSharedCredentialsFile string = "AWS_SHARED_CREDENTIALS_FILE"
	AwsCredentials           string = "~/.aws/credentials"
	AwsAccessKeyID           string = "aws_access_key_id"
	AwsSecretAccessKey       string = "aws_secret_access_key"
)

// error messages
var (
	ErrorNotFoundAwsAccessKeyID     = errors.New(AwsAccessKeyID + ErrorNotFound)
	ErrorNotFoundAwsSecretAccessKey = errors.New(AwsSecretAccessKey + ErrorNotFound)
)

// Credential provide credentials
type Credential struct {
	ProfileName        string
	AwsAccessKeyID     string
	AwsSecretAccessKey string
}

// Credentials has many Credential
type Credentials []Credential

// NewCredentials create a new Credentials instance
func NewCredentials() *Credentials {
	return new(Credentials)
}

// Parse credential file
func (c *Credentials) Parse(credentialsFile string) error {
	data, err := ini.Load(credentialsFile)
	if err != nil {
		return err
	}

	for _, section := range data.Sections() {
		if section.Name() == "DEFAULT" {
			continue
		}

		credential := Credential{}

		credential.ProfileName = section.Name()

		if section.HasKey(AwsAccessKeyID) {
			credential.AwsAccessKeyID = section.Key(AwsAccessKeyID).String()
		}

		if section.HasKey(AwsSecretAccessKey) {
			credential.AwsSecretAccessKey = section.Key(AwsSecretAccessKey).String()
		}

		*c = append(*c, credential)
	}

	return nil
}

// ProfileNames get name of profiles
func (c *Credentials) ProfileNames() ([]string, error) {
	var profileNames []string

	for _, credential := range *c {
		profileNames = append(profileNames, credential.ProfileName)
	}

	return profileNames, nil
}

// GetAwsAccessKeyID get aws_access_key_id
func (c *Credentials) GetAwsAccessKeyID(profileName string) (string, error) {
	for _, credential := range *c {
		if credential.ProfileName == profileName {
			return credential.AwsAccessKeyID, nil
		}
	}

	return EmptyString, ErrorNotFoundAwsAccessKeyID
}

// GetAwsSecretAccessKey get aws_secret_access_key
func (c *Credentials) GetAwsSecretAccessKey(profileName string) (string, error) {
	for _, credential := range *c {
		if credential.ProfileName == profileName {
			return credential.AwsSecretAccessKey, nil
		}
	}

	return EmptyString, ErrorNotFoundAwsSecretAccessKey
}

// GetCredentialsPath provide file path to credentials
func GetCredentialsPath() (string, error) {
	credentialsFile, err := homedir.Expand(AwsCredentials)
	if err != nil {
		return EmptyString, err
	}

	if os.Getenv(AwsSharedCredentialsFile) != EmptyString {
		credentialsFile, err = homedir.Expand(os.Getenv(AwsSharedCredentialsFile))
		if err != nil {
			return EmptyString, err
		}
	}

	return credentialsFile, nil
}
