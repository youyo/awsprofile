package awsprofile

import (
	"errors"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	ini "gopkg.in/ini.v1"
)

const (
	awsSharedCredentialsFile string = "AWS_SHARED_CREDENTIALS_FILE"
	awsCredentials           string = "~/.aws/credentials"
	awsAccessKeyId           string = "aws_access_key_id"
	awsSecretAccessKey       string = "aws_secret_access_key"
)

var (
	ErrorNotFoundAwsAccessKeyID     error = errors.New(awsAccessKeyId + errorNotFound)
	ErrorNotFoundAwsSecretAccessKey error = errors.New(awsSecretAccessKey + errorNotFound)
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

		if section.HasKey(awsAccessKeyId) {
			credential.AwsAccessKeyID = section.Key(awsAccessKeyId).String()
		}

		if section.HasKey(awsSecretAccessKey) {
			credential.AwsSecretAccessKey = section.Key(awsSecretAccessKey).String()
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

	return emptyString, ErrorNotFoundAwsAccessKeyID
}

func (c *Credentials) GetAwsSecretAccessKey(profileName string) (string, error) {
	for _, credential := range *c {
		if credential.ProfileName == profileName {
			return credential.AwsSecretAccessKey, nil
		}
	}

	return emptyString, ErrorNotFoundAwsSecretAccessKey
}

func GetCredentialsPath() (string, error) {
	credentialsFile, err := homedir.Expand(awsCredentials)
	if err != nil {
		return emptyString, err
	}

	if os.Getenv(awsSharedCredentialsFile) != emptyString {
		credentialsFile, err = homedir.Expand(os.Getenv(awsSharedCredentialsFile))
		if err != nil {
			return emptyString, err
		}
	}

	return credentialsFile, nil
}
