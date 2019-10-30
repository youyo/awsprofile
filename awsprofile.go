package awsprofile

// Null values
const (
	EmptyString string = ""
	ZeroInt     int    = 0
)

// Error messages
const (
	ErrorNotFound string = " is not found"
)

// AwsProfile provide Credentials and Configs
type AwsProfile struct {
	Credentials *Credentials
	Configs     *Configs
}

// New create a AwsProfile instance
func New() *AwsProfile {
	awsProfile := &AwsProfile{
		Credentials: NewCredentials(),
		Configs:     NewConfigs(),
	}

	return awsProfile
}

// Parse credential file and config file
func (a *AwsProfile) Parse() error {
	credentialsFile, err := GetCredentialsPath()
	if err != nil {
		return err
	}

	if err = a.Credentials.Parse(credentialsFile); err != nil {
		return err
	}

	configsFile, err := GetConfigsPath()
	if err != nil {
		return err
	}

	if err = a.Configs.Parse(configsFile); err != nil {
		return err
	}

	return nil
}

// ProfileNames get name of profiles
func (a *AwsProfile) ProfileNames() ([]string, error) {
	var profileNames []string

	for _, credential := range *a.Credentials {
		profileNames = append(profileNames, credential.ProfileName)
	}

	for _, config := range *a.Configs {
		profileNames = append(profileNames, config.ProfileName)
	}

	profileNames = removeDuplicate(profileNames)

	return profileNames, nil
}

// GetCredentials get Credentials
func (a *AwsProfile) GetCredentials() *Credentials {
	return a.Credentials
}

// GetConfigs get Configs
func (a *AwsProfile) GetConfigs() *Configs {
	return a.Configs
}

// IsCredential
func (a *AwsProfile) IsCredential(profile string) (bool, *Credential) {
	var ok bool = false
	var cred *Credential

	for _, credential := range *a.Credentials {
		if credential.ProfileName == profile {
			ok = true
			cred = &credential
		}
	}

	return ok, cred
}

// IsConfig
func (a *AwsProfile) IsConfig(profile string) (bool, *Config) {
	var ok bool = false
	var conf *Config

	for _, config := range *a.Configs {
		if config.ProfileName == profile {
			ok = true
			conf = &config
		}
	}

	return ok, conf
}

func removeDuplicate(s []string) []string {
	var list []string

	m := make(map[string]bool)

	for _, v := range s {
		if !m[v] {
			m[v] = true
			list = append(list, v)
		}
	}

	return list
}
