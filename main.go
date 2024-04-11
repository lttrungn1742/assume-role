package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials/stscreds"
)

func usage() {
	fmt.Fprintf(os.Stderr, "Usage: %s <role>\n", os.Args[0])
	flag.PrintDefaults()
}

func init() {
	flag.Usage = usage
}

func must(err error) {
	if err != nil {
		if _, ok := err.(*exec.ExitError); ok {
			// Errors are already on Stderr.
			os.Exit(1)
		}

		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func main() {

	flag.Parse()
	argv := flag.Args()
	if len(argv) < 1 {
		flag.Usage()
		os.Exit(1)
	}
	role := argv[0]

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

	fmt.Printf("export AWS_ACCESS_KEY_ID=%s\n", credentials.AccessKeyID)
	fmt.Printf("export AWS_SECRET_ACCESS_KEY=%s\n", credentials.SecretAccessKey)
	fmt.Printf("export AWS_SESSION_TOKEN=%s\n", credentials.SessionToken)
}
