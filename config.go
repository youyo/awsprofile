package awsprofile

import (
	"errors"
	"os"
	"strings"

	homedir "github.com/mitchellh/go-homedir"
	ini "gopkg.in/ini.v1"
)

const (
	AWS_CONFIG_FILE         string = "AWS_CONFIG_FILE"
	AWS_CONFIG              string = "~/.aws/config"
	ROLE_ARN                string = "role_arn"
	SOURCE_PROFILE          string = "source_profile"
	CREDENTIAL_SOURCE       string = "credential_source"
	ROLE_SESSION_NAME       string = "role_session_name"
	MFA_SERIAL              string = "mfa_serial"
	DURATION_SECONDS        string = "duration_seconds"
	AWS_SESSION_TOKEN       string = "aws_session_token"
	EXTERNAL_ID             string = "external_id"
	CA_BUNDLE               string = "ca_bundle"
	CLI_FOLLOW_URLPARAM     string = "cli_follow_urlparam"
	CLI_TIMESTAMP_FORMAT    string = "cli_timestamp_format"
	CREDENTIAL_PROCESS      string = "credential_process"
	WEB_IDENTITY_TOKEN_FILE string = "web_identity_token_file"
	OUTPUT                  string = "output"
	REGION                  string = "region"
)

var (
	ErrorNotFoundRoleArn              error = errors.New(ROLE_ARN + ErrorNotFound)
	ErrorNotFoundSourceProfile        error = errors.New(SOURCE_PROFILE + ErrorNotFound)
	ErrorNotFoundCredentialSource     error = errors.New(CREDENTIAL_SOURCE + ErrorNotFound)
	ErrorNotFoundRoleSessionName      error = errors.New(ROLE_SESSION_NAME + ErrorNotFound)
	ErrorNotFoundMfaSerial            error = errors.New(MFA_SERIAL + ErrorNotFound)
	ErrorNotFoundDurationSeconds      error = errors.New(DURATION_SECONDS + ErrorNotFound)
	ErrorNotFoundAwsSessionToken      error = errors.New(AWS_SESSION_TOKEN + ErrorNotFound)
	ErrorNotFoundExternalID           error = errors.New(EXTERNAL_ID + ErrorNotFound)
	ErrorNotFoundCaBundle             error = errors.New(CA_BUNDLE + ErrorNotFound)
	ErrorNotFoundCliFollowUrlparam    error = errors.New(CLI_FOLLOW_URLPARAM + ErrorNotFound)
	ErrorNotFoundCliTimestampFormat   error = errors.New(CLI_TIMESTAMP_FORMAT + ErrorNotFound)
	ErrorNotFoundCredentialProcess    error = errors.New(CREDENTIAL_PROCESS + ErrorNotFound)
	ErrorNotFoundWebIdentityTokenFile error = errors.New(WEB_IDENTITY_TOKEN_FILE + ErrorNotFound)
	ErrorNotFoundOutput               error = errors.New(OUTPUT + ErrorNotFound)
	ErrorNotFoundRegion               error = errors.New(REGION + ErrorNotFound)
)

type Config struct {
	ProfileName          string
	RoleArn              string
	SourceProfile        string
	CredentialSource     string
	RoleSessionName      string
	MfaSerial            string
	DurationSeconds      int
	AwsSessionToken      string
	ExternalID           int
	CaBundle             string
	CliFollowUrlparam    string
	CliTimestampFormat   string
	CredentialProcess    string
	WebIdentityTokenFile string
	Output               string
	Region               string
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

		config.ProfileName = strings.Replace(section.Name(), "profile ", EmptyString, 1)

		if section.HasKey(ROLE_ARN) {
			config.RoleArn = section.Key(ROLE_ARN).String()
		}

		if section.HasKey(SOURCE_PROFILE) {
			config.SourceProfile = section.Key(SOURCE_PROFILE).String()
		}

		if section.HasKey(CREDENTIAL_SOURCE) {
			config.CredentialSource = section.Key(CREDENTIAL_SOURCE).String()
		}

		if section.HasKey(ROLE_SESSION_NAME) {
			config.RoleSessionName = section.Key(ROLE_SESSION_NAME).String()
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

		if section.HasKey(AWS_SESSION_TOKEN) {
			config.AwsSessionToken = section.Key(AWS_SESSION_TOKEN).String()
		}

		if section.HasKey(EXTERNAL_ID) {
			externalID, err := section.Key(EXTERNAL_ID).Int()
			if err != nil {
				return err
			}

			config.ExternalID = externalID
		}

		if section.HasKey(CA_BUNDLE) {
			config.CaBundle = section.Key(CA_BUNDLE).String()
		}

		if section.HasKey(CLI_FOLLOW_URLPARAM) {
			config.CliFollowUrlparam = section.Key(CLI_FOLLOW_URLPARAM).String()
		}

		if section.HasKey(CLI_TIMESTAMP_FORMAT) {
			config.CliTimestampFormat = section.Key(CLI_TIMESTAMP_FORMAT).String()
		}

		if section.HasKey(CREDENTIAL_PROCESS) {
			config.CredentialProcess = section.Key(CREDENTIAL_PROCESS).String()
		}

		if section.HasKey(WEB_IDENTITY_TOKEN_FILE) {
			config.WebIdentityTokenFile = section.Key(WEB_IDENTITY_TOKEN_FILE).String()
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
	if config, ok := c.get(profileName); ok {
		return config.RoleArn, nil
	}

	return EmptyString, ErrorNotFoundRoleArn
}

func (c *Configs) GetSourceProfile(profileName string) (string, error) {
	if config, ok := c.get(profileName); ok {
		return config.SourceProfile, nil
	}

	return EmptyString, ErrorNotFoundSourceProfile
}

func (c *Configs) GetCredentialSource(profileName string) (string, error) {
	if config, ok := c.get(profileName); ok {
		return config.CredentialSource, nil
	}

	return EmptyString, ErrorNotFoundCredentialSource
}

func (c *Configs) GetRoleSessionName(profileName string) (string, error) {
	if config, ok := c.get(profileName); ok {
		return config.RoleSessionName, nil
	}

	return EmptyString, ErrorNotFoundRoleSessionName
}

func (c *Configs) GetMfaSerial(profileName string) (string, error) {
	if config, ok := c.get(profileName); ok {
		return config.MfaSerial, nil
	}

	return EmptyString, ErrorNotFoundMfaSerial
}

func (c *Configs) GetDurationSeconds(profileName string) (int, error) {
	if config, ok := c.get(profileName); ok {
		return config.DurationSeconds, nil
	}

	return ZeroInt, ErrorNotFoundDurationSeconds
}

func (c *Configs) GetAwsSessionToken(profileName string) (string, error) {
	if config, ok := c.get(profileName); ok {
		return config.AwsSessionToken, nil
	}

	return EmptyString, ErrorNotFoundAwsSessionToken
}

func (c *Configs) GetExternalID(profileName string) (int, error) {
	if config, ok := c.get(profileName); ok {
		return config.ExternalID, nil
	}

	return ZeroInt, ErrorNotFoundExternalID
}

func (c *Configs) GetCaBundle(profileName string) (string, error) {
	if config, ok := c.get(profileName); ok {
		return config.CaBundle, nil
	}

	return EmptyString, ErrorNotFoundCaBundle
}

func (c *Configs) GetCliFollowUrlparam(profileName string) (string, error) {
	if config, ok := c.get(profileName); ok {
		return config.CliFollowUrlparam, nil
	}

	return EmptyString, ErrorNotFoundCliFollowUrlparam
}

func (c *Configs) GetCliTimestampFormat(profileName string) (string, error) {
	if config, ok := c.get(profileName); ok {
		return config.CliTimestampFormat, nil
	}

	return EmptyString, ErrorNotFoundCliTimestampFormat
}

func (c *Configs) GetCredentialProcess(profileName string) (string, error) {
	if config, ok := c.get(profileName); ok {
		return config.CredentialProcess, nil
	}

	return EmptyString, ErrorNotFoundCredentialProcess
}

func (c *Configs) GetWebIdentityTokenFile(profileName string) (string, error) {
	if config, ok := c.get(profileName); ok {
		return config.WebIdentityTokenFile, nil
	}

	return EmptyString, ErrorNotFoundWebIdentityTokenFile
}

func (c *Configs) GetOutput(profileName string) (string, error) {
	if config, ok := c.get(profileName); ok {
		return config.Output, nil
	}

	return EmptyString, ErrorNotFoundOutput
}

func (c *Configs) GetRegion(profileName string) (string, error) {
	if config, ok := c.get(profileName); ok {
		return config.Region, nil
	}

	return EmptyString, ErrorNotFoundRegion
}

func (c *Configs) get(profileName string) (*Config, bool) {
	for _, config := range *c {
		if config.ProfileName == profileName {
			return &config, true
		}
	}

	return nil, false
}

func GetConfigsPath() (string, error) {
	configsFile, err := homedir.Expand(AWS_CONFIG)
	if err != nil {
		return EmptyString, err
	}

	if os.Getenv(AWS_CONFIG_FILE) != EmptyString {
		configsFile, err = homedir.Expand(os.Getenv(AWS_CONFIG_FILE))
		if err != nil {
			return EmptyString, err
		}
	}

	return configsFile, nil
}
