package main

import (
	"flag"
	"fmt"
	"os"
)

func informAndExit(message string, code int) {
	fmt.Println(message)
	os.Exit(code)
}

func main() {
	var roleArn string
	var loadEnv bool

	linkCmd := flag.NewFlagSet("link", flag.ExitOnError)
	linkCmd.StringVar(&roleArn, "role-arn", "", "the role arn to assume for federating with AWS")

	credsCmd := flag.NewFlagSet("creds", flag.ExitOnError)
	credsCmd.StringVar(&roleArn, "role-arn", "", "the role arn to assume for federating with AWS")
	credsCmd.BoolVar(&loadEnv, "load", false, "include this flag to load the creds into the environment")

	if len(os.Args) < 2 {
		informAndExit("This executable needs a subcommand and options to work. Use -h for help.", 1)
	}
}
