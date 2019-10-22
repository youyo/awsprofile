# awsprofile

[![GitHubActionBadge](https://github.com/youyo/awsprofile/workflows/test/badge.svg)](https://github.com/youyo/awsprofile/actions)

awsprofile is library of parse to `~/.aws/credentials` and `~/.aws/config` .

## Basic usage

```go
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

// Get Configs
configs, err := awsProfile.GetConfigs()
if err != nil {
     log.Fatal(err)
}

// Get RoleArn
roleArn, err := configs.GetRoleArn("bar")
if err != nil {
      log.Fatal(err)
}

fmt.Println(roleArn)
// Output: arn:aws:iam::xxxxxxxxxxxx:role/bar
```

## Document

See https://godoc.org/github.com/youyo/awsprofile
