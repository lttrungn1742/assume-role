package main

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"time"

	"runtime"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
)

func must(err error) {
	if err != nil {
		if _, ok := err.(*exec.ExitError); ok {
			os.Exit(1)
		}

		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func main() {
	role := os.Args[0]
	var system string = ""
	if len(os.Args) == 2 {
		system = os.Args[1]
	}

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithSharedConfigProfile(role),
		config.WithAssumeRoleCredentialOptions(
			func(aro *stscreds.AssumeRoleOptions) {
				aro.TokenProvider = stscreds.StdinTokenProvider
				aro.Duration = time.Duration(60) * time.Minute
			}),
	)

	must(err)
	credentials, err := cfg.Credentials.Retrieve(context.TODO())

	must(err)

	if system == "--windows" {
		fmt.Printf("$env:AWS_ACCESS_KEY_ID=\"%s\"\n", credentials.AccessKeyID)
		fmt.Printf("$env:AWS_SECRET_ACCESS_KEY=\"%s\"\n", credentials.SecretAccessKey)
		fmt.Printf("$env:AWS_SESSION_TOKEN=\"%s\"\n", credentials.SessionToken)
	} else {
		fmt.Printf("export AWS_ACCESS_KEY_ID=%s\n", credentials.AccessKeyID)
		fmt.Printf("export AWS_SECRET_ACCESS_KEY=%s\n", credentials.SecretAccessKey)
		fmt.Printf("export AWS_SESSION_TOKEN=%s\n", credentials.SessionToken)
	}

	if runtime.GOOS == "windows" || system == "" {
		fmt.Printf("$env:AWS_ACCESS_KEY_ID=\"%s\"\n", credentials.AccessKeyID)
		fmt.Printf("$env:AWS_SECRET_ACCESS_KEY=\"%s\"\n", credentials.SecretAccessKey)
		fmt.Printf("$env:AWS_SESSION_TOKEN=\"%s\"\n", credentials.SessionToken)
	} else {
		fmt.Printf("export AWS_ACCESS_KEY_ID=%s\n", credentials.AccessKeyID)
		fmt.Printf("export AWS_SECRET_ACCESS_KEY=%s\n", credentials.SecretAccessKey)
		fmt.Printf("export AWS_SESSION_TOKEN=%s\n", credentials.SessionToken)
	}
}
