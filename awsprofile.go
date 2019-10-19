package awsprofile

const (
	emptyString   string = ""
	zeroInt       int    = 0
	errorNotFound string = " is not found"
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
