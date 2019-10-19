package awsprofile

import (
	"errors"
	"os"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
	ini "gopkg.in/ini.v1"
)

const (
	AWS_CONFIG_FILE  string = "AWS_CONFIG_FILE"
	AWS_CONFIG       string = "~/.aws/config"
	ROLE_ARN         string = "role_arn"
	SOURCE_PROFILE   string = "source_profile"
	MFA_SERIAL       string = "mfa_serial"
	DURATION_SECONDS string = "duration_seconds"
	EXTERNAL_ID      string = "external_id"
	OUTPUT           string = "output"
	REGION           string = "region"
)

var (
	ErrorNotFoundRoleArn         error = errors.New(ROLE_ARN + errorNotFound)
	ErrorNotFoundSourceProfile   error = errors.New(SOURCE_PROFILE + errorNotFound)
	ErrorNotFoundMfaSerial       error = errors.New(MFA_SERIAL + errorNotFound)
	ErrorNotFoundDurationSeconds error = errors.New(DURATION_SECONDS + errorNotFound)
	ErrorNotFoundExternalID      error = errors.New(EXTERNAL_ID + errorNotFound)
	ErrorNotFoundOutput          error = errors.New(OUTPUT + errorNotFound)
	ErrorNotFoundRegion          error = errors.New(REGION + errorNotFound)
)

type Config struct {
	ProfileName     string
	RoleArn         string
	SourceProfile   string
	MfaSerial       string
	DurationSeconds int
	ExternalID      int
	Output          string
	Region          string
}

type Configs []Config

func NewConfigs() *Configs {
	return new(Configs)
}

func (c *Configs) Parse(configFile string) error {
	data, err := ini.Load(configFile)
	if err != nil {
		return err
	}

	for _, section := range data.Sections() {
		if section.Name() == "DEFAULT" {
			continue
		}

		config := Config{}

		config.ProfileName = strings.Replace(section.Name(), "profile ", emptyString, 1)

		if section.HasKey(ROLE_ARN) {
			config.RoleArn = section.Key(ROLE_ARN).String()
		}

		if section.HasKey(SOURCE_PROFILE) {
			config.SourceProfile = section.Key(SOURCE_PROFILE).String()
		}

		if section.HasKey(MFA_SERIAL) {
			config.MfaSerial = section.Key(MFA_SERIAL).String()
		}

		if section.HasKey(DURATION_SECONDS) {
			durationSeconds, err := section.Key(DURATION_SECONDS).Int()
			if err != nil {
				return err
			}

			config.DurationSeconds = durationSeconds
		}

		if section.HasKey(EXTERNAL_ID) {
			externalID, err := section.Key(EXTERNAL_ID).Int()
			if err != nil {
				return err
			}

			config.ExternalID = externalID
		}

		if section.HasKey(OUTPUT) {
			config.Output = section.Key(OUTPUT).String()
		}

		if section.HasKey(REGION) {
			config.Region = section.Key(REGION).String()
		}

		*c = append(*c, config)
	}

	return nil
}

func (c *Configs) ProfileNames() ([]string, error) {
	var profileNames []string

	for _, config := range *c {
		profileNames = append(profileNames, config.ProfileName)
	}

	return profileNames, nil
}

func (c *Configs) GetRoleArn(profileName string) (string, error) {
	for _, config := range *c {
		if config.ProfileName == profileName {
			return config.RoleArn, nil
		}
	}

	return emptyString, ErrorNotFoundRoleArn
}

func (c *Configs) GetSourceProfile(profileName string) (string, error) {
	for _, config := range *c {
		if config.ProfileName == profileName {
			return config.SourceProfile, nil
		}
	}

	return emptyString, ErrorNotFoundSourceProfile
}

func (c *Configs) GetMfaSerial(profileName string) (string, error) {
	for _, config := range *c {
		if config.ProfileName == profileName {
			return config.MfaSerial, nil
		}
	}

	return emptyString, ErrorNotFoundMfaSerial
}

func (c *Configs) GetDurationSeconds(profileName string) (int, error) {
	for _, config := range *c {
		if config.ProfileName == profileName {
			return config.DurationSeconds, nil
		}
	}

	return zeroInt, ErrorNotFoundDurationSeconds
}

func (c *Configs) GetExternalID(profileName string) (int, error) {
	for _, config := range *c {
		if config.ProfileName == profileName {
			return config.ExternalID, nil
		}
	}

	return zeroInt, ErrorNotFoundExternalID
}

func (c *Configs) GetOutput(profileName string) (string, error) {
	for _, config := range *c {
		if config.ProfileName == profileName {
			return config.Output, nil
		}
	}

	return emptyString, ErrorNotFoundOutput
}

func (c *Configs) GetRegion(profileName string) (string, error) {
	for _, config := range *c {
		if config.ProfileName == profileName {
			return config.Region, nil
		}
	}

	return emptyString, ErrorNotFoundRegion
}

func GetConfigsPath() (string, error) {
	configsFile, err := homedir.Expand(AWS_CONFIG)
	if err != nil {
		return emptyString, err
	}

	if os.Getenv(AWS_CONFIG_FILE) != emptyString {
		configsFile, err = homedir.Expand(os.Getenv(AWS_CONFIG_FILE))
		if err != nil {
			return emptyString, err
		}
	}

	return configsFile, nil
}
