package awsprofile

import (
	"errors"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	ini "gopkg.in/ini.v1"
)

const (
	AWS_SHARED_CREDENTIALS_FILE string = "AWS_SHARED_CREDENTIALS_FILE"
	AWS_CREDENTIALS             string = "~/.aws/credentials"
	AWS_ACCESS_KEY_ID           string = "aws_access_key_id"
	AWS_SECRET_ACCESS_KEY       string = "aws_secret_access_key"
)

var (
	ErrorNotFoundAwsAccessKeyID     error = errors.New(AWS_ACCESS_KEY_ID + ErrorNotFound)
	ErrorNotFoundAwsSecretAccessKey error = errors.New(AWS_SECRET_ACCESS_KEY + ErrorNotFound)
)

type Credential struct {
	ProfileName        string
	AwsAccessKeyID     string
	AwsSecretAccessKey string
}

type Credentials []Credential

func NewCredentials() *Credentials {
	return new(Credentials)
}

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

		if section.HasKey(AWS_ACCESS_KEY_ID) {
			credential.AwsAccessKeyID = section.Key(AWS_ACCESS_KEY_ID).String()
		}

		if section.HasKey(AWS_SECRET_ACCESS_KEY) {
			credential.AwsSecretAccessKey = section.Key(AWS_SECRET_ACCESS_KEY).String()
		}

		*c = append(*c, credential)
	}

	return nil
}

func (c *Credentials) ProfileNames() ([]string, error) {
	var profileNames []string

	for _, credential := range *c {
		profileNames = append(profileNames, credential.ProfileName)
	}

	return profileNames, nil
}

func (c *Credentials) GetAwsAccessKeyID(profileName string) (string, error) {
	for _, credential := range *c {
		if credential.ProfileName == profileName {
			return credential.AwsAccessKeyID, nil
		}
	}

	return EmptyString, ErrorNotFoundAwsAccessKeyID
}

func (c *Credentials) GetAwsSecretAccessKey(profileName string) (string, error) {
	for _, credential := range *c {
		if credential.ProfileName == profileName {
			return credential.AwsSecretAccessKey, nil
		}
	}

	return EmptyString, ErrorNotFoundAwsSecretAccessKey
}

func GetCredentialsPath() (string, error) {
	credentialsFile, err := homedir.Expand(AWS_CREDENTIALS)
	if err != nil {
		return EmptyString, err
	}

	if os.Getenv(AWS_SHARED_CREDENTIALS_FILE) != EmptyString {
		credentialsFile, err = homedir.Expand(os.Getenv(AWS_SHARED_CREDENTIALS_FILE))
		if err != nil {
			return EmptyString, err
		}
	}

	return credentialsFile, nil
}
